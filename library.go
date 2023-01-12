package acicn

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/guoyk93/gg"
	cp "github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	PrefixDev = path.Join("dev-build", "acicn")
)

type ManifestGlobal struct {
	Dirs []string `yaml:"dirs"`
	Vars gg.M     `yaml:"vars"`
}

type ManifestRepo struct {
	Name string          `yaml:"name"`
	Tags []ManifestTag   `yaml:"tags"`
	Vars map[string]gg.M `yaml:"vars"`
}

type ManifestTag struct {
	Name       string   `yaml:"name"`
	Also       []string `yaml:"also"`
	Dockerfile string   `yaml:"dockerfile"`
	Vars       []string `yaml:"vars"`
}

type Repo struct {
	Dir               string   // src/jdk
	Repo              string   // jdk
	LongName          string   // dev-build/acicn/jdk:11-debian-11
	ShortName         string   // jdk:11-debian-11
	ShortNames        []string // jdk:11-debian-11,jdk:11
	Tags              []string // 11-debian-11,11
	Dockerfile        string   // Dockerfile
	DockerfileContent []byte   // FROM xxxx \n ...

	Vars gg.M

	Dependencies []string // debian:11

	known map[string]string
}

func (r *Repo) renderDockerfile() (err error) {
	var (
		tmpl = gg.Must(
			template.New("__main__").Option("missingkey=zero").Funcs(template.FuncMap{
				"Lookup": func(upstream string) (string, error) {
					r.Dependencies = append(r.Dependencies, upstream)
					if v, ok := r.known[upstream]; ok {
						return v, nil
					}
					return "", errors.New("no known: " + upstream)
				},
			}).Parse(string(
				gg.Must(os.ReadFile(filepath.Join(r.Dir, r.Dockerfile)))),
			),
		)
		out = &bytes.Buffer{}
	)

	gg.Must0(tmpl.Execute(out, r.Vars))

	r.DockerfileContent = cleanLines(out.Bytes())
	return
}

func (r *Repo) Generate() (err error) {
	defer gg.Guard(&err)

	dir := filepath.Join("out", r.ShortName)
	gg.Must0(os.RemoveAll(dir))

	gg.Must0(cp.Copy(r.Dir, dir, cp.Options{
		Sync:              true,
		PermissionControl: cp.PerservePermission,
		PreserveTimes:     true,
		PreserveOwner:     true,
	}))

	// render dockerfile to dockerfile.out
	gg.Must0(os.Remove(filepath.Join(dir, r.Dockerfile)))
	gg.Must0(os.Remove(filepath.Join(dir, "manifest.yml")))
	gg.Must0(os.Remove(filepath.Join(dir, "README.md")))
	// write new dockerfile
	gg.Must0(os.WriteFile(filepath.Join(dir, "Dockerfile"), r.DockerfileContent, 0640))
	return
}

func Load(overrides gg.M) (repos []*Repo, err error) {
	defer gg.Guard(&err)

	// create index
	known := map[string]string{}

	var mGlobal ManifestGlobal

	gg.Must0(yaml.Unmarshal(gg.Must(os.ReadFile("manifest.yml")), &mGlobal))

	// create build tasks
	for _, dir := range mGlobal.Dirs {

		dir = filepath.Join("src", filepath.Join(strings.Split(dir, "/")...))

		var mRepo ManifestRepo

		gg.Must0(yaml.Unmarshal(gg.Must(os.ReadFile(filepath.Join(dir, "manifest.yml"))), &mRepo))

		if mRepo.Name == "" {
			err = errors.New("missing field 'repo' for " + dir)
			return
		}

		for _, mTag := range mRepo.Tags {
			if mTag.Dockerfile == "" {
				mTag.Dockerfile = "Dockerfile"
			}
			if mTag.Name == "" {
				err = fmt.Errorf("missing name for tag in repo: %s", mRepo.Name)
				return
			}

			// create vars
			vars := gg.M{}
			{
				for k, v := range mGlobal.Vars {
					vars[k] = v
				}
				for _, kg := range mTag.Vars {
					if mRepo.Vars[kg] == nil {
						err = fmt.Errorf("missing vars group: %s in repo: %s ", kg, mRepo.Name)
						return
					}
					for k, v := range mRepo.Vars[kg] {
						vars[k] = v
					}
				}
				for k, v := range overrides {
					vars[k] = v
				}
			}

			// create repo
			repo := &Repo{
				Dir:        dir,
				Repo:       mRepo.Name,
				Tags:       append([]string{mTag.Name}, mTag.Also...),
				Dockerfile: mTag.Dockerfile,
				Vars:       vars,
				known:      known,
			}
			repo.ShortName = repo.Repo + ":" + repo.Tags[0]
			for _, item := range repo.Tags {
				repo.ShortNames = append(repo.ShortNames, repo.Repo+":"+item)
			}
			repo.LongName = path.Join(PrefixDev, repo.ShortName)
			gg.Must0(repo.renderDockerfile())

			// record known
			for _, item := range repo.ShortNames {
				known[item] = repo.LongName
			}

			// append repos
			repos = append(repos, repo)
		}
	}

	return
}

func cleanLines(buf []byte) []byte {
	lines := bytes.Split(buf, []byte{'\n'})
	out := make([][]byte, 0, len(lines))
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		out = append(out, line)
	}
	return bytes.Join(out, []byte{'\n'})
}
