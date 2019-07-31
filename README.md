

ChartServer is an in-cluster Helm Chart repository that maintains and serves an index of Helm charts. ChartServer uses Custom Resource Definitions (CRDs) to deploy new Charts and Chart Versions.

# Quick Start

### Install ChartServer

Deploy the ChartServer operator to a cluster:

```shell
helm repo add chartserver https://chart.sh/charts
helm install chartserver/chartserver
```

or

```shell
kubectl apply -f https://get.chart.sh
```

### Create a Chart
A chart is not an installable chart. A chart will not be installable until at least 1 chart version has been deployed.

To create a chart, write a YAML descriptor of the chart:

```yaml
apiVersion: chart.sh/v1beta1
kind: Chart
metadata:
  name: sample-chart
spec:
  name: my-chart-name
  location:
    download:
      uri: https://something
```

Now, `kubectl apply -f` this new file. You can verify it worked with:

```shell
kubectl get charts.chart.sh
```

### Deploy a Chart Version

Create the first chart version:

```yaml
apiVersion: chart.sh/v1beta1
kind: ChartVersion
metadata:
  name: sample-chart
spec:
  name: my-chart-name
  apPVersion: 1.0.0
  chartVersion: 0.1.0
```

Deploy this YAML using `kubectl` and list chart versions:

```
kubectl get chartversions.chart.sh
```

### Install

ChartServer is now serving an `index.yaml` file at https://your-ingress-or-service/index.yaml

This will be updated every time a new chart is published.
