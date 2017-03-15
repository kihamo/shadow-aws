package aws

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
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

func (c *Component) GetDashboardMenu() *dashboard.Menu {
	return &dashboard.Menu{
		Name: "Aws",
		Icon: "cloud",
		SubMenu: []*dashboard.Menu{
			{
				Name: "SNS",
				Url:  "/sns",
			},
			{
				Name: "SES",
				Url:  "/ses",
			},
		},
	}
}

func (c *Component) GetDashboardRoutes() []*dashboard.Route {
	return []*dashboard.Route{
		{
			Methods: []string{http.MethodGet, http.MethodPost},
			Path:    "/sns",
			Handler: &SNSHandler{
				component: c,
			},
		},
		{
			Methods: []string{http.MethodGet, http.MethodPost},
			Path:    "/ses",
			Handler: &SESHandler{
				component: c,
			},
		},
	}
}
