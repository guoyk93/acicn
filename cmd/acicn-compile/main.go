package main

import (
	"flag"
	"fmt"
	"github.com/guoyk93/acicn"
	"github.com/guoyk93/gg"
	"github.com/guoyk93/gg/ggos"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	rcSuffix = "-rc"
)

var (
	regexpNotSafe = regexp.MustCompile(`[^a-z0-9]`)
)

var (
	stepCheckout = gg.M{
		"name": "checkout",
		"uses": "actions/checkout@v3",
	}
	stepSetupBuildX = gg.M{
		"name": "setup docker buildx",
		"uses": "docker/setup-buildx-action@v2",
	}
	stepLoginGithubCR = gg.M{
		"name": "docker login - ghcr",
		"uses": "docker/login-action@v2",
		"with": gg.M{
			"registry": "ghcr.io",
			"username": "${{github.actor}}",
			"password": "${{secrets.GITHUB_TOKEN}}",
		},
	}
	stepLoginCodingCR = gg.M{
		"name": "docker login - coding registry",
		"uses": "docker/login-action@v2",
		"with": gg.M{
			"registry": "${{secrets.MIRROR_CODING_REGISTRY}}",
			"username": "${{secrets.MIRROR_CODING_USERNAME}}",
			"password": "${{secrets.MIRROR_CODING_PASSWORD}}",
		},
	}
	stepLoginTencentCR = gg.M{
		"name": "docker login - ccr registry",
		"uses": "docker/login-action@v2",
		"with": gg.M{
			"registry": "ccr.ccs.tencentyun.com",
			"username": "${{secrets.MIRROR_CCR_USERNAME}}",
			"password": "${{secrets.MIRROR_CCR_PASSWORD}}",
		},
	}
	stepLoginAliyunCR = gg.M{
		"name": "docker login - aliyun registry",
		"uses": "docker/login-action@v2",
		"with": gg.M{
			"registry": "registry.cn-shenzhen.aliyuncs.com",
			"username": "${{secrets.MIRROR_ALIYUN_USERNAME}}",
			"password": "${{secrets.MIRROR_ALIYUN_PASSWORD}}",
		},
	}
)

func releaseJobName(name string) string {
	return "r_" + regexpNotSafe.ReplaceAllString(path.Base(strings.ToLower(name)), "_") + "_r"
}

func mirrorJobName(name string) string {
	return "m_" + regexpNotSafe.ReplaceAllString(path.Base(strings.ToLower(name)), "_") + "_m"
}

type WorkflowMirrorOptions struct{}

func updateWorkflowMirror(repos []*acicn.Repo, opts WorkflowMirrorOptions) (err error) {
	defer gg.Guard(&err)

	jobs := gg.M{}

	for _, item := range repos {

		tags := gg.Map(item.Tags, func(tag string) string {
			return fmt.Sprintf("type=raw,value=%s", tag+rcSuffix)
		})

		job := gg.M{
			"if":      "inputs.job_name == 'all' || contains(inputs.job_name,'" + mirrorJobName(item.Name) + "')",
			"runs-on": "ubuntu-latest",
			"permissions": gg.M{
				"contents": "read",
				"packages": "read",
				"id-token": "write",
			},
			"steps": []gg.M{
				stepCheckout,
				{
					"name": "generate dockerfile",
					"uses": "DamianReeves/write-file-action@v1.2",
					"with": gg.M{
						"path":       "docker/Dockerfile",
						"write-mode": "overwrite",
						"contents":   "FROM " + item.Name + rcSuffix,
					},
				},
				stepSetupBuildX,
				stepLoginGithubCR,
				stepLoginCodingCR,
				stepLoginTencentCR,
				stepLoginAliyunCR,
				{
					"name": "meta for " + item.ShortName(),
					"id":   "meta",
					"uses": "docker/metadata-action@v4",
					"with": gg.M{
						"images": strings.Join([]string{
							"ccr.ccs.tencentyun.com/acicn/" + item.Repo,
							"registry.cn-shenzhen.aliyuncs.com/acicn/" + item.Repo,
							"${{secrets.MIRROR_CODING_REGISTRY}}/${{secrets.MIRROR_CODING_PREFIX}}/" + item.Repo,
						}, "\n"),
						"tags": strings.Join(tags, "\n"),
					},
				},
				{
					"name": "build for " + item.ShortName(),
					"uses": "docker/build-push-action@v3",
					"id":   "build",
					"with": gg.M{
						"context":    "docker",
						"pull":       true,
						"push":       true,
						"tags":       "${{steps.meta.outputs.tags}}",
						"labels":     "${{steps.meta.outputs.labels}}",
						"cache-from": "type=gha",
						"cache-to":   "type=gha,mode=max",
					},
				},
			},
		}

		jobs[mirrorJobName(item.Name)] = job
	}

	doc := gg.M{
		"name": "mirror",
		"on": gg.M{
			"workflow_dispatch": gg.M{
				"inputs": gg.M{
					"job_name": gg.M{
						"description": "names of jobs to execute, 'all' for all",
						"required":    true,
						"type":        "string",
					},
				},
			},
		},
		"jobs": jobs,
	}

	buf := gg.Must(yaml.Marshal(doc))
	gg.Must0(os.MkdirAll(filepath.Join(".github", "workflows"), 0755))
	gg.Must0(os.WriteFile(filepath.Join(".github", "workflows", "mirror.yaml"), buf, 0640))
	return
}

type WorkflowReleaseOptions struct {
	NoDep bool
}

func updateWorkflowRelease(repos []*acicn.Repo, opts WorkflowReleaseOptions) (err error) {
	defer gg.Guard(&err)

	var soloSuffix string
	if opts.NoDep {
		soloSuffix = "-nodep"
	}

	jobs := gg.M{}

	for _, item := range repos {

		tags := gg.Map(item.Tags, func(tag string) string {
			return fmt.Sprintf("type=raw,value=%s", tag+rcSuffix)
		})

		job := gg.M{
			"if":      "inputs.job_name == 'all' || contains(inputs.job_name,'" + releaseJobName(item.Name) + "')",
			"runs-on": "ubuntu-latest",
			"permissions": gg.M{
				"contents": "read",
				"packages": "write",
				"id-token": "write",
			},
			"steps": []gg.M{
				{
					"name": "checkout",
					"uses": "actions/checkout@v3",
				},
				{
					"name": "setup docker buildx",
					"uses": "docker/setup-buildx-action@v2",
				},
				{
					"name": "docker login - ghcr",
					"uses": "docker/login-action@v2",
					"with": gg.M{
						"registry": "ghcr.io",
						"username": "${{github.actor}}",
						"password": "${{secrets.GITHUB_TOKEN}}",
					},
				},
				{
					"name": "docker login - dockerhub",
					"uses": "docker/login-action@v2",
					"with": gg.M{
						"username": "guoyk",
						"password": "${{secrets.DOCKERHUB_TOKEN}}",
					},
				},
				{
					"name": "meta for " + item.ShortName(),
					"id":   "meta",
					"uses": "docker/metadata-action@v4",
					"with": gg.M{
						"images": strings.Join([]string{
							"ghcr.io/guoyk93/acicn/" + item.Repo,
							"acicn/" + item.Repo,
						}, "\n"),
						"tags": strings.Join(tags, "\n"),
					},
				},
				{
					"name": "build for " + item.ShortName(),
					"uses": "docker/build-push-action@v3",
					"id":   "build",
					"with": gg.M{
						"context":    "out/" + item.ShortName(),
						"pull":       true,
						"push":       "${{ inputs.push }}",
						"tags":       "${{steps.meta.outputs.tags}}",
						"labels":     "${{steps.meta.outputs.labels}}",
						"cache-from": "type=gha",
						"cache-to":   "type=gha,mode=max",
					},
				},
			},
		}

		var needs []string

		for k, v := range item.Vars {
			if s, ok := v.(string); ok && s != "" {
				if strings.HasPrefix(k, "upstream") {
					needs = append(needs, releaseJobName(gg.Must(item.LookupKnown(s))))
				}
			}
		}

		sort.Strings(needs)

		if len(needs) > 0 && !opts.NoDep {
			job["needs"] = needs
		}

		jobs[releaseJobName(item.Name)] = job
	}

	doc := gg.M{
		"name": "release" + soloSuffix,
		"on": gg.M{
			"workflow_dispatch": gg.M{
				"inputs": gg.M{
					"push": gg.M{
						"description": "push to registry",
						"required":    true,
						"type":        "boolean",
					},
					"job_name": gg.M{
						"description": "names of jobs to execute, 'all' for all",
						"required":    true,
						"type":        "string",
					},
				},
			},
		},
		"jobs": jobs,
	}

	buf := gg.Must(yaml.Marshal(doc))
	gg.Must0(os.MkdirAll(filepath.Join(".github", "workflows"), 0755))
	gg.Must0(os.WriteFile(filepath.Join(".github", "workflows", "release"+soloSuffix+".yaml"), buf, 0640))
	return
}

type Record struct {
	Name string   `yaml:"name"`
	Also []string `yaml:"also"`
}

func main() {
	var err error
	defer ggos.Exit(&err)
	defer gg.Guard(&err)

	var (
		optOverride string
	)

	flag.StringVar(&optOverride, "override", "", "override values")
	flag.Parse()

	// update overrides
	overrides := gg.M{}
	{
		optOverrides := strings.Split(optOverride, ";")
		for _, override := range optOverrides {
			override = strings.TrimSpace(override)
			overrideKV := strings.SplitN(override, "=", 2)
			if len(overrideKV) != 2 {
				continue
			}
			overrides[strings.TrimSpace(overrideKV[0])] = strings.TrimSpace(overrideKV[1])
		}
	}

	// load the library
	repos := gg.Must(acicn.Load(overrides))

	// generate github workflow
	gg.Must0(updateWorkflowRelease(repos, WorkflowReleaseOptions{NoDep: false}))
	gg.Must0(updateWorkflowRelease(repos, WorkflowReleaseOptions{NoDep: true}))
	gg.Must0(updateWorkflowMirror(repos, WorkflowMirrorOptions{}))

	// collect image names
	// update IMAGES.txt
	var names []string
	{
		nameMap := map[string]struct{}{}
		for _, item := range repos {
			for _, tag := range item.Tags {
				nameMap[item.Repo+":"+tag] = struct{}{}
			}
		}
		names = gg.Keys(nameMap)
		sort.Strings(names)
	}
	gg.Must0(os.WriteFile("IMAGES.txt", []byte(strings.Join(names, "\n")), 0644))

	// update IMAGES.yml
	var records []*Record
	{
		for _, item := range repos {
			records = append(records, &Record{
				Name: item.ShortName(),
				Also: item.ShortNames()[1:],
			})
		}
		sort.Slice(records, func(i, j int) bool {
			return records[i].Name < records[j].Name
		})
		for _, item := range records {
			sort.Strings(item.Also)
		}
	}
	gg.Must0(os.WriteFile("IMAGES.yml", gg.Must(yaml.Marshal(records)), 0644))

	// remove output dir
	gg.Must0(os.RemoveAll("out"))

	// generate
	for _, repo := range repos {
		gg.Log("generate: " + repo.ShortName())
		gg.Must0(repo.Generate())
	}

}
