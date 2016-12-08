package main // import "github.com/kihamo/shadow-aws/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-aws/resource/aws"
	"github.com/kihamo/shadow-aws/service"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/metrics"
	"github.com/kihamo/shadow/resource/template"
	"github.com/kihamo/shadow/resource/workers"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/shadow/service/system"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(config.Resource),
			new(logger.Resource),
			new(metrics.Resource),
			new(template.Resource),
			new(alerts.Resource),
			new(workers.Resource),
			new(aws.Aws),
		},
		[]shadow.Service{
			new(frontend.FrontendService),
			new(system.SystemService),
			new(service.AwsService),
		},
		"Aws",
		"1.0",
		"12345-full",
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
