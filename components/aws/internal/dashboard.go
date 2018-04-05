package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow-aws/components/aws/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return dashboard.TemplatesFromAssetFS(c)
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenu("Aws").
		WithRoute(routes[0]).
		WithIcon("cloud").
		WithChild(dashboard.NewMenu("SNS").WithRoute(routes[0])).
		WithChild(dashboard.NewMenu("SES").WithRoute(routes[1]))
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRoute("/"+c.Name()+"/sns/", &handlers.SNSHandler{}).
				WithMethods([]string{http.MethodGet, http.MethodPost}).
				WithAuth(true),
			dashboard.NewRoute("/"+c.Name()+"/ses/", &handlers.SESHandler{}).
				WithMethods([]string{http.MethodGet, http.MethodPost}).
				WithAuth(true),
		}
	}

	return c.routes
}
