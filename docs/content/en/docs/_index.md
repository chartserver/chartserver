---
title: "ChartServer Documentation"
linkTitle: "Documentation"
weight: 20
menu:
  main:
    weight: 30
---

ChartServer is the easiest way to deploy a self-hosted, Kubernetes-native Helm chart repository. A [Helm repository](https://github.com/helm/helm/blob/master/docs/chart_repository.md) is a hosted location where chart manifests are stored and can be shared.

The goal of ChartServer is to provide the easiest option for anyone needing to set up and host a Helm repository while supporting GitOps workflows and secure content-delivery pipelines to publish a chart.

While ChartServer hopes to be the easiest way to run a Helm repository using a GitOps pipeline, if you are looking for other options, there are some listed in the Helm docs, including:

#### S3 and GCS
Because a Helm repository is a web server that serves an `index.yaml` that references downloadable archives, this technically can be served from any easy-to-use site. The Helm repo has [documentation](https://github.com/helm/helm/blob/master/docs/chart_repository.md#google-cloud-storage) on how to use Google Cloud Storage to host a chart repository. This is the tool that the helm/charts/stable repo uses. The Helm CLI is used to publish charts here.

Using this option requires that you manage IAM keys and have the Helm CLI tool available anywhere that's needed to publish a chart.

#### ChartMuseum
ChartMuseum is an open source, self hosted repository server. It's mature and been around for a while. There's a (simple) proprietary API needed to publish charts to ChartMuseum. This is a solution for one-off, or relatively simple workflows that can support imperative deployments. ChartServer builds on the ideas of ChartMuseum by replacing the propietary API with a Kubernetes CRD which enables the declarative nature of Kubernetes deployments. ChartServer also enables a GitOps delivery of charts to the repository, which is not possible with ChartMuseum.

#### GitHub Pages
Using GitHub Pages is a unique and declarative way to host a Helm repository. GitHub Pages makes it easy to build chart publishing into a CI pipeline, and it's very transparent in how it works. GitHub Pages is a static site build/host tool. There's no API for it. ChartServer uses the Kubernetes API to allow easy create, edit and delete of Charts from the repo, while still allowing GitOps tools such as [Weave Flux](https://www.weave.works/oss/flux/) or [ArgoCD](https://argoproj.github.io/argo-cd/) to manage the published charts.

#### Other Options
There are other, good options such as Artifactory and ProGet. These tools support other package formats also and are worth investigating.
