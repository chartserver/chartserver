---
date: 2019-07-05
linktitle: "Missing Charts"
title: Missing Charts
weight: 70010
---

Charts are stored in Kubernetes (in etcd) as the custom resource that you've deployed. Any problems or errors that ChartServer had parsing the fields will be reported in the Status field of the spec.

```yaml
apiVersion: chartserver.io/v1beta1
kind: ChartVersion
metadata:
  name: an-app-0-0-1
  namespace: default
spec:
  ...
```

Given the above ChartVersion spec, if there was an error parsing it, you can find the error by:

```shell
kubectl describe chartversion.chartserver.io an-app-0-0-1
```

This will show the fields that you deployed, including a `Status` field near the bottom. If there are any known errors with the spec, they will be reported here.
