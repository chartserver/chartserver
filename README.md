
<p align="center">
  <h3 align="center">ChartServer</h3>
  <p align="center">The easiest way to host your own Helm chart repository.</p>
</p>

---

ChartServer is self-hosted, Kubernetes-native Helm Chart repository. ChartServer's responsibility is to maintain and serve an index of Helm charts that are available to install.


# Quick Start

### Install ChartServer

Deploy the ChartServer operator to a cluster:

```shell
helm repo add chartserver https://chartserver.io/charts
helm install chartserver/chartserver
```

or

```shell
kubectl apply -f https://get.chartserver.io
```

### Deploy a Chart Version

We strongly encourage using [HelmReleaser](https://helmreleaser.com) to automate the work required to release a Helm chart. ChartServer is fully integrated into the HelmReleaser workflow.

Create the first chart version:

```yaml
apiVersion: chartserver.io/v1beta1
kind: ChartVersion
metadata:
  name: sample-chart
spec:
  name: my-chart-name
  appVersion: 0.0.2
  chartVersion: 0.0.2
  created: "2019-07-31T14:26:20Z"
  description: A Helm chart for Kubernetes
  digest: b388d25a08ab7c9fe9ea20fe19165dcf8e53689e4ae369d13c9140d2a6a047ef
  home: "https://github.com/my/repo"
  icon: ""
  maintainers:
    - "Somebody <somebody@gmail.com>"
  sources:
    - https://github.com/my/repo
  urls:
  - https://github.com/my/repo/releases/download/v0.0.2/my-chart-0.0.2.tgz
```

Deploy this YAML using `kubectl` and list chart versions:

```
kubectl get chartversions.chartserver.io
```

### Install

ChartServer is now serving an `index.yaml` file at https://your-ingress-or-service/index.yaml

This will be updated every time a new chart is published.
