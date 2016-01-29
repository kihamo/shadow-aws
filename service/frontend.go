package service

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/service/frontend"
)

func (s *AwsService) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (s *AwsService) GetFrontendMenu() *frontend.FrontendMenu {
	return &frontend.FrontendMenu{
		Name: "Aws",
		Url:  "/aws",
		Icon: "cloud",
	}
}

func (s *AwsService) SetFrontendHandlers(router *frontend.Router) {
	router.GET(s, "/aws", &IndexHandler{})
}
