package action

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func GetReleaseStrings() []string {
	var settings = cli.New()

	debug := func(format string, v ...interface{}) {
		if settings.Debug {
			format = fmt.Sprintf("[debug] %s\n", format)
			log.Output(2, fmt.Sprintf(format, v...))
		}
	}

	cfg := new(action.Configuration)
	cfg.Init(settings.RESTClientGetter(), "", "", debug)

	lister := action.NewList(cfg)
	releases, err := lister.Run()
	if err != nil {
		log.Fatal(err)
	}

	releaseStrings := make([]string, 0)
	for _, r := range releases {
		releaseStrings = append(releaseStrings, fmt.Sprintf("%s", r))
	}

	return releaseStrings
}
