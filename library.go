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
	Dir        string
	LongName   string
	Repo       string
	Tags       []string
	Dockerfile string
	Vars       gg.M
	Known      map[string]string
}

func (b Repo) ShortName() string {
	return b.Repo + ":" + b.Tags[0]
}

func (b Repo) ShortNames() []string {
	return gg.Map(b.Tags, func(item string) string {
		return b.Repo + ":" + item
	})
}

func (b Repo) Lookup(upstream string) (string, error) {
	if v, ok := b.Known[upstream]; ok {
		return v, nil
	}
	return "", errors.New("no known: " + upstream)
}

func (b Repo) LookupUpstream(upstream string) (string, error) {
	if v, ok := b.Known[upstream]; ok {
		return v, nil
	}
	return "", errors.New("no known: " + upstream)
}

func (b Repo) Generate() (err error) {
	defer gg.Guard(&err)

	dir := filepath.Join("out", b.ShortName())
	gg.Must0(os.RemoveAll(dir))

	gg.Must0(cp.Copy(b.Dir, dir, cp.Options{
		Sync:              true,
		PermissionControl: cp.PerservePermission,
		PreserveTimes:     true,
		PreserveOwner:     true,
	}))

	// render dockerfile to dockerfile.out
	{
		var (
			tmpl = gg.Must(
				template.New("__main__").Option("missingkey=zero").Funcs(template.FuncMap{
					"Known": b.LookupUpstream,
				}).Parse(string(
					gg.Must(os.ReadFile(filepath.Join(dir, b.Dockerfile)))),
				),
			)
			out = &bytes.Buffer{}
		)

		// generate dockerfile
		gg.Must0(tmpl.Execute(out, b.Vars))
		// remove existed dockerfile
		gg.Must0(os.Remove(filepath.Join(dir, b.Dockerfile)))
		gg.Must0(os.Remove(filepath.Join(dir, "manifest.yml")))
		gg.Must0(os.Remove(filepath.Join(dir, "README.md")))
		// write new dockerfile
		gg.Must0(os.WriteFile(filepath.Join(dir, "Dockerfile"), cleanLines(out.Bytes()), 0640))
	}
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
				Known:      known,
			}
			repo.LongName = path.Join(PrefixDev, repo.ShortName())

			// record known
			for _, item := range repo.ShortNames() {
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
