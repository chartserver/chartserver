package server

import (
	"time"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/repo"
)

func (h *ChartServer) refreshCache() error {
	namespaces, err := h.client.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to list namespaces")
	}

	chartCache := map[string]repo.ChartVersions{}

	for _, namespace := range namespaces.Items {
		chartVersions, err := h.chartserverClient.ChartVersions(namespace.Name).List(metav1.ListOptions{})
		if err != nil {
			return errors.Wrap(err, "failed to list chart versions in namespace")
		}

		for _, chartVersion := range chartVersions.Items {
			metadata := chart.Metadata{
				Name:        chartVersion.Spec.Name,
				Home:        chartVersion.Spec.Home,
				Sources:     chartVersion.Spec.Sources,
				Version:     chartVersion.Spec.ChartVersion,
				Description: chartVersion.Spec.Description,
				Keywords:    []string{},
				Maintainers: []*chart.Maintainer{},
				Icon:        chartVersion.Spec.Icon,
				AppVersion:  chartVersion.Spec.AppVersion,
			}
			for _, m := range chartVersion.Spec.Maintainers {
				metadata.Maintainers = append(metadata.Maintainers, &chart.Maintainer{
					Name: m,
				})
			}

			indexChartVersions, ok := chartCache[chartVersion.Spec.Name]
			if !ok {
				indexChartVersions = repo.ChartVersions{}
			}

			created, err := time.Parse(time.RFC3339, chartVersion.Spec.Created)
			if err != nil {
				return errors.Wrap(err, "unable to parse created time in chartversion")
			}

			indexChartVersion := repo.ChartVersion{
				Metadata: &metadata,
				Digest:   chartVersion.Spec.Digest,
				URLs:     chartVersion.Spec.URLs,
				Created:  created,
			}

			indexChartVersions = append(indexChartVersions, &indexChartVersion)

			chartCache[chartVersion.Spec.Name] = indexChartVersions
		}
	}

	h.chartCache = chartCache
	h.cacheGeneratedAt = time.Now()

	return nil
}
