---
date: 2019-05-25
linktitle: "Defining a Chart Version"
title: Defining a Chart Version
weight: 30010
---

Charts and chart versions are deployed to ChartServer using `kubectl`.

To deploy a Chart version, create a YAML spec that defines your chart:

```yaml
apiVersion: chartserver.io/v1beta1
kind: ChartVersion
metadata:
  name: my-app-name
  namespace: default
spec:
  appVersion: 1.0.1
  chartVersion: 0.1.0
  created: "2019-08-01T01:40:09Z"
  description: My-app is a new way to run AI on Kubernetes.
  digest: 1b463ced20823eaba77904015c546ca56c31fcc20d8a6a54179463fb12b78fc2
  home: "https://something.com"
  icon: ""
  maintainers:
    - "Your Name <you@gmail.com>"
  name: my-app-name
  urls:
  - https://github.com/my-app-name/chart/releases/download/v0.1.0/my-app-0.1.1.tgz
```

Once you have this YAML created, you can:

```shell
kubectl apply -f ./path/to.yaml
```

And the chart will be available in ChartServer for others to install.

Note that ChartServer doesn't manage the binary itself for you. You must provide this somewhere else (GitHub releases is a good option), and reference it in your ChartVersion manifest.
