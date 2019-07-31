
# Image URL to use all building/pushing image targets
IMG ?= controller:latest
export GO111MODULE=on

all: test manager

# Run tests
test: generate fmt vet manifests
	go test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: manager
manager: generate fmt vet
	go build -o bin/manager github.com/chartserver/chartserver/cmd/manager

.PHONY: run
run: generate fmt vet
	go run ./cmd/manager/main.go

.PHONY: install
install: manifests
	kubectl apply -f config/crds

.PHONY: deploy
deploy: manifests
	kubectl apply -f config/crds
	kustomize build config/default | kubectl apply -f -

.PHONY: manifests
manifests:
	controller-gen paths=./pkg/apis/... output:dir=./config/crds

.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/...

.PHONY: generate
generate: controller-gen # client-gen
	controller-gen object:headerFile=./hack/boilerplate.go.txt paths=./pkg/apis/...
	# client-gen --output-package=github.com/chartserver/chartserver/pkg/client --clientset-name chartserverclientset --input-base github.com/chartserver/chartserver/pkg/apis --input chartserver/v1beta1 -h ./hack/boilerplate.go.txt

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.0-beta.2
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# find or download client-gen
client-gen:
ifeq (, $(shell which client-gen))
	go get k8s.io/code-generator/cmd/client-gen@kubernetes-1.13.5
CLIENT_GEN=$(shell go env GOPATH)/bin/client-gen
else
CLIENT_GEN=$(shell which client-gen)
endif
