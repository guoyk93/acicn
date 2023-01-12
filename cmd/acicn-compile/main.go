package main

import (
	"flag"
	"github.com/guoyk93/acicn"
	"github.com/guoyk93/gg"
	"github.com/guoyk93/gg/ggos"
	"os"
	"strings"
)

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

	// remove output dir
	gg.Must0(os.RemoveAll("out"))

	// generate
	for _, repo := range repos {
		gg.Log("generate: " + repo.ShortName())
		gg.Must0(repo.Generate())
	}

}
