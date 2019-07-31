package controller

import (
	"github.com/chartserver/chartserver/pkg/controller/server"
	"k8s.io/client-go/rest"
)

func StartWebServer(config *rest.Config, bindAddr string) error {
	return server.StartWebServer(config, bindAddr)
}
