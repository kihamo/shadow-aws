package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow-aws/components/aws/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenuWithUrl(
		"Aws",
		"/"+c.Name()+"/",
		"cloud",
		[]dashboard.Menu{
			dashboard.NewMenuWithRoute("SNS", routes[0], "", nil, nil),
			dashboard.NewMenuWithRoute("SES", routes[1], "", nil, nil),
		},
		nil)
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet, http.MethodPost},
				"/"+c.Name()+"/sns/",
				&handlers.SNSHandler{
					Component: c,
				},
				"",
				false),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet, http.MethodPost},
				"/"+c.Name()+"/ses/",
				&handlers.SESHandler{
					Component: c,
				},
				"",
				false),
		}
	}

	return c.routes
}