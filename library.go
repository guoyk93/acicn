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

type ManifestGlobal struct {
	Registries []string `yaml:"registries"`
	Doc        string   `yaml:"doc"`
	Upstreams  []string `yaml:"upstreams"`
	Dirs       []string `yaml:"dirs"`
	Vars       gg.M     `yaml:"vars"`
}

type ManifestRepo struct {
	Name string          `yaml:"name"`
	Tags []ManifestTag   `yaml:"tags"`
	Vars map[string]gg.M `yaml:"vars"`
}

type ManifestTag struct {
	Name       string       `yaml:"name"`
	Also       []string     `yaml:"also"`
	Dockerfile string       `yaml:"dockerfile"`
	Vars       []string     `yaml:"vars"`
	Test       ManifestTest `yaml:"test"`
}

type ManifestTest struct {
	Delay  int    `yaml:"delay"`
	Run    string `yaml:"run"`
	Exec   string `yaml:"exec"`
	Output string `yaml:"output"`
}

type Repo struct {
	Dir        string
	Name       string
	Repo       string
	Tags       []string
	Doc        string
	Dockerfile string
	Vars       gg.M
	Known      map[string]string
	Test       ManifestTest
}

func (b Repo) ShortName() string {
	return b.Repo + ":" + b.Tags[0]
}

func (b Repo) ShortNames() []string {
	return gg.Map(b.Tags, func(item string) string {
		return b.Repo + ":" + item
	})
}

func (b Repo) LookupKnown(upstream string) (string, error) {
	if v, ok := b.Known[upstream]; ok {
		return v, nil
	}
	return "", errors.New("no known: " + upstream)
}

func (b Repo) GenerateMirror() (err error) {
	defer gg.Guard(&err)

	dir := filepath.Join("out", b.ShortName())
	gg.Must0(os.RemoveAll(dir))
	gg.Must0(os.MkdirAll(dir, 0755))
	gg.Must0(os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte("FROM "+b.Name), 0644))

	return
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

	// create banner.minit.txt
	gg.Must0(
		os.WriteFile(
			filepath.Join(dir, "banner.minit.txt"),
			[]byte(fmt.Sprintf("本镜像基于 ACICN 镜像 %s 制作，详细信息参阅 %s", b.ShortName(), b.Doc)),
			0644,
		),
	)

	// render dockerfile to dockerfile.out
	{
		var (
			tmpl = gg.Must(
				template.New("__main__").Option("missingkey=zero").Funcs(template.FuncMap{
					"Known": b.LookupKnown,
				}).Parse(string(
					gg.Must(os.ReadFile(filepath.Join(dir, b.Dockerfile)))),
				),
			)
			out = &bytes.Buffer{}
		)

		// generate dockerfile
		gg.Must0(tmpl.Execute(out, b.Vars))
		out.WriteString("\nADD banner.minit.txt /etc/banner.minit.txt")
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
				Doc:        mGlobal.Doc,
				Dockerfile: mTag.Dockerfile,
				Vars:       vars,
				Known:      known,
				Test:       mTag.Test,
			}
			repo.Name = path.Join(mGlobal.Registries[0], repo.ShortName())

			// record known
			for _, item := range repo.ShortNames() {
				known[item] = repo.Name
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
