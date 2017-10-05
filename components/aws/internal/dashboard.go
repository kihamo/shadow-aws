package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow-aws/components/aws/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) GetDashboardMenu() dashboard.Menu {
	routes := c.GetDashboardRoutes()

	return dashboard.NewMenuWithUrl(
		"Aws",
		"/"+c.GetName()+"/",
		"cloud",
		[]dashboard.Menu{
			dashboard.NewMenuWithRoute("SNS", routes[0], "", nil, nil),
			dashboard.NewMenuWithRoute("SES", routes[1], "", nil, nil),
		},
		nil)
}

func (c *Component) GetDashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRoute(
				c.GetName(),
				[]string{http.MethodGet, http.MethodPost},
				"/"+c.GetName()+"/sns/",
				&handlers.SNSHandler{
					Component: c,
				},
				"",
				false),
			dashboard.NewRoute(
				c.GetName(),
				[]string{http.MethodGet, http.MethodPost},
				"/"+c.GetName()+"/ses/",
				&handlers.SESHandler{
					Component: c,
				},
				"",
				false),
		}
	}

	return c.routes
}
