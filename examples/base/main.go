package main // import "github.com/kihamo/shadow-aws/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-aws/resource/aws"
	"github.com/kihamo/shadow-aws/service"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/resource/alerts"
	"github.com/kihamo/shadow/resource/metrics"
	"github.com/kihamo/shadow/resource/workers"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/shadow/service/system"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(resource.Config),
			new(resource.Logger),
			new(resource.Template),
			new(alerts.Alerts),
			new(metrics.Metrics),
			new(workers.Workers),
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
