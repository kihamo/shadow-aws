package main // import "github.com/kihamo/shadow-aws/examples/base"

import (
	"log"

	"github.com/kihamo/shadow"
	r "github.com/kihamo/shadow-aws/resource"
	s "github.com/kihamo/shadow-aws/service"
	"github.com/kihamo/shadow/resource"
	"github.com/kihamo/shadow/service/frontend"
	"github.com/kihamo/shadow/service/system"
)

func main() {
	application, err := shadow.NewApplication(
		[]shadow.Resource{
			new(resource.Config),
			new(resource.Logger),
			new(resource.Template),
			new(resource.Workers),
			new(r.Aws),
		},
		[]shadow.Service{
			new(frontend.FrontendService),
			new(system.SystemService),
			new(s.AwsService),
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
