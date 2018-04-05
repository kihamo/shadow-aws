package main // import "github.com/kihamo/shadow-aws/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	aws "github.com/kihamo/shadow-aws/components/aws/instance"
	config "github.com/kihamo/shadow/components/config/instance"
	dashboard "github.com/kihamo/shadow/components/dashboard/instance"
	i18n "github.com/kihamo/shadow/components/i18n/instance"
	logger "github.com/kihamo/shadow/components/logger/instance"
	metrics "github.com/kihamo/shadow/components/metrics/instance"
	workers "github.com/kihamo/shadow/components/workers/instance"
)

func main() {
	application, err := shadow.NewApp(
		"Aws",
		"1.0",
		"12345-full",
		[]shadow.Component{
			aws.NewComponent(),
			config.NewComponent(),
			dashboard.NewComponent(),
			i18n.NewComponent(),
			logger.NewComponent(),
			metrics.NewComponent(),
			workers.NewComponent(),
		},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
