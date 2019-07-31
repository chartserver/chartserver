package server

import (
	"fmt"
	"net/http"
	"time"

	chartserverclientv1beta1 "github.com/chartserver/chartserver/pkg/client/chartserverclientset/typed/chartserver/v1beta1"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/repo"
)

type ChartServer struct {
	chartserverClient *chartserverclientv1beta1.ChartserverV1beta1Client
	client            *kubernetes.Clientset

	chartCache       map[string]repo.ChartVersions
	cacheGeneratedAt time.Time
}

func StartWebServer(config *rest.Config, bindAddr string) error {
	g := gin.New()

	chartserverClient, err := chartserverclientv1beta1.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "failed to create chartserver clientset")
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes clientset")
	}

	h := ChartServer{
		chartserverClient: chartserverClient,
		client:            client,
	}

	if err := h.refreshCache(); err != nil {
		return errors.Wrap(err, "failed to init cache")
	}

	go func() {
		for {
			if err := h.refreshCache(); err != nil {
				fmt.Println("failed to refresh cache")
			}

			time.Sleep(time.Second * 10)
		}
	}()

	root := g.Group("/")

	root.GET("/healthz", h.Healthz)
	root.GET("/index.yaml", h.Index)

	server := http.Server{Addr: bindAddr, Handler: g}
	errChan := make(chan error)

	go func() {
		errChan <- server.ListenAndServe()
	}()

	return nil
}

func (h ChartServer) Healthz(c *gin.Context) {
	c.JSON(200, map[string]interface{}{})
}

func (h *ChartServer) Index(c *gin.Context) {
	indexFile := repo.NewIndexFile()
	indexFile.Entries = h.chartCache
	indexFile.Generated = h.cacheGeneratedAt

	b, err := yaml.Marshal(indexFile)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.Header("Content-Type", "application/x-yaml; charset=utf-8")
	c.String(200, string(b))
}
