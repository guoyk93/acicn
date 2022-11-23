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

var (
	regexpNotSafe = regexp.MustCompile(`[^a-z0-9]`)
)

func jobName(name string) string {
	return "r_" + regexpNotSafe.ReplaceAllString(path.Base(strings.ToLower(name)), "_")
}

func updateWorkflow(repos []*acicn.Repo) (err error) {
	defer gg.Guard(&err)

	jobs := gg.M{}

	for _, item := range repos {

		tags := gg.Map(item.Tags, func(tag string) string {
			return fmt.Sprintf("type=raw,value=%s", tag)
		})

		upstream, _ := item.Vars["upstream"].(string)
		upstream = strings.TrimSpace(upstream)

		var pull any

		if upstream == "" {
			pull = "${{ inputs.force_pull }}"
			gg.Log("missing upstream for: " + item.Name)
		} else {
			pull = true
		}

		job := gg.M{
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
					"name": "setup go",
					"uses": "actions/setup-go@v3",
					"with": gg.M{
						"go-version": "1.19",
					},
				},
				{
					"name": "generate out",
					"run":  "go run -mod vendor ./cmd/acicn-build/main.go -name '" + item.ShortName() + "'",
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
					"name": "meta for " + item.Repo + ":" + item.Tags[0],
					"id":   "meta",
					"uses": "docker/metadata-action@v4",
					"with": gg.M{
						"images": strings.Join([]string{
							//"acicn/" + item.Repo,
							"ghcr.io/guoyk93/acicn/" + item.Repo,
						}, "\n"),
						"tags": strings.Join(tags, "\n"),
					},
				},
				{
					"name": "build for " + item.Repo + ":" + item.Tags[0],
					"uses": "docker/build-push-action@v3",
					"id":   "build",
					"with": gg.M{
						"context":    "out/" + item.Repo + ":" + item.Tags[0],
						"pull":       pull,
						"push":       true,
						"tags":       "${{steps.meta.outputs.tags}}",
						"labels":     "${{steps.meta.outputs.labels}}",
						"cache-from": "type=gha",
						"cache-to":   "type=gha,mode=max",
					},
				},
			},
		}

		if upstream != "" {
			job["needs"] = []string{jobName(gg.Must(item.LookupKnown(upstream)))}
		}

		jobs[jobName(item.Name)] = job
	}

	doc := gg.M{
		"name": "release",
		"on": gg.M{
			"workflow_dispatch": gg.M{
				"inputs": []gg.M{
					{
						"force_pull": gg.M{
							"description": "force pull upstream images",
							"required":    true,
							"type":        "boolean",
						},
					},
				},
			},
		},
		"jobs": jobs,
	}

	buf := gg.Must(yaml.Marshal(doc))
	gg.Must0(os.MkdirAll(filepath.Join(".github", "workflows"), 0755))
	gg.Must0(os.WriteFile(filepath.Join(".github", "workflows", "release.yaml"), buf, 0640))
	return
}

func main() {
	var err error
	defer ggos.Exit(&err)
	defer gg.Guard(&err)

	var (
		optUpdateWorkflow bool
		optUpdateImages   bool

		optName     string
		optOverride string
	)

	flag.BoolVar(&optUpdateWorkflow, "update-workflow", false, "update workflow")
	flag.BoolVar(&optUpdateImages, "update-images", false, "update images list")
	flag.StringVar(&optName, "name", "all", "name")
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
	if optUpdateWorkflow {
		gg.Must0(updateWorkflow(repos))
	}

	// collect image names
	if optUpdateImages {
		var imageNames []string
		{
			imageNameMap := map[string]struct{}{}
			for _, task := range repos {
				for _, tag := range task.Tags {
					imageNameMap[task.Repo+":"+tag] = struct{}{}
				}
			}
			imageNames = gg.Keys(imageNameMap)
			sort.Strings(imageNames)
		}

		// update IMAGES.txt
		gg.Must0(os.WriteFile("IMAGES.txt", []byte(strings.Join(imageNames, "\n")), 0644))
	}

	if optName != "" {
		// remove output dir
		gg.Must0(os.RemoveAll("out"))

		// generate
		for _, repo := range repos {
			if optName == "all" || optName == repo.ShortName() {
				gg.Log("generate: " + repo.ShortName())
				gg.Must0(repo.Generate())
			}
		}
	}

}
