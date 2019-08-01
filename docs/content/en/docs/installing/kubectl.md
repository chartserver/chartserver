---
date: 2019-07-05
linktitle: "Installing ChartServer"
title: Installing ChartServer
weight: 20010
---

To install ChartServer to a cluster:

```shell
kubectl apply -f https://git.io/chartserver
```

This will create a new namespace named `chartserver-system` and install the ChartServer manager into this namespace.

This will install the latest version.

To verify the deployment, you can run:

```shell
kubectl get pods -n chartserver-system
```

There should be 1 pod running in this namespace:

```shell
$ kubectl get pods -n chartserver-system
NAME                              READY   STATUS    RESTARTS   AGE
chartserver-controller-manager-0  1/1     Running   0          3d2h
```
