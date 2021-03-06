# Image URL to use all building/pushing image targets
IMG ?= iter8-controller:latest
CRD_VERSION ?= v1alpha2

# ISTIO
ISTIO_NAMESPACE ?= istio-system
ISTIO_VERSION ?= $(shell kubectl -n ${ISTIO_NAMESPACE} get deploy --selector=istio=pilot -o jsonpath='{.items[0].spec.template.spec.containers[*].image}' | cut -d: -f2)
MIXER_DISABLED ?= $(shell kubectl -n ${ISTIO_NAMESPACE} get cm istio -o json | jq .data.mesh | grep -o 'disableMixerHttpReports: [A-Za-z]\+' | cut -d ' ' -f2)
ifndef TELEMETRY_VERSION
TELEMETRY_VERSION := $(shell if [ "${MIXER_DISABLED}" = "false" ]; then echo "v1"; else echo "v2"; fi)
endif
ifndef PROMETHEUS_JOB_LABEL
PROMETHEUS_JOB_LABEL := $(shell if [ "${MIXER_DISABLED}" = "false" ]; then echo "istio-mesh"; elif [ "-1" = $$(hack/semver.sh ${ISTIO_VERSION} 1.7.0) ]; then echo "envoy-stats"; else echo "kubernetes-pods"; fi)
endif

# HELM
HELM_VERSION ?= v$(shell helm version --client --short | sed 's/.*v\([0-9]*\).*/\1/')
ifeq ($(HELM_VERSION),v2)
HELM2_NAME := --name iter8-controller
HELM3_NAME := 
HELM_INCLUDE_OPTION := -x
else
HELM2_NAME := 
HELM3_NAME := iter8-controller
HELM_INCLUDE_OPTION := -s
endif

verify-env:
	echo "ISTIO_NAMESPACE = ${ISTIO_NAMESPACE}"
	echo "ISTIO_VERSION = ${ISTIO_VERSION}"
	echo "MIXER_DISABLED = ${MIXER_DISABLED}"
	echo "TELEMETRY_VERSION = ${TELEMETRY_VERSION}"
	echo "PROMETHEUS_JOB_LABEL = ${PROMETHEUS_JOB_LABEL}"
	echo "HELM_VERSION = ${HELM_VERSION}"
	echo "HELM2_NAME = ${HELM2_NAME}"
	echo "HELM3_NAME = ${HELM3_NAME}"
	echo "HELM_INCLUDE_OPTION = ${HELM_INCLUDE_OPTION}"

all: manager

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager github.com/iter8-tools/iter8-istio/cmd/manager

# Run against the Kubernetes cluster configured in $KUBECONFIG or ~/.kube/config
# TODO replace vet
run: generate fmt vet load
	go run ./cmd/manager/main.go

# Generate iter8 crds and rbac manifests
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd:allowDangerousTypes=true \
	  paths=./pkg/apis/iter8/${CRD_VERSION} output:crd:dir=./install/helm/iter8-controller/templates/crds/${CRD_VERSION}/
	./hack/crd_fix.sh ${CRD_VERSION}

# Prepare Kubernetes cluster for iter8 (running in cluster or locally):
#   install CRDs
#   install configmap/iter8-metrics is defined in namespace iter8 (creating namespace if needed)
load: manifests
	helm template ${HELM3_NAME} install/helm/iter8-controller ${HELM2_NAME} \
		${HELM_INCLUDE_OPTION} templates/default/namespace.yaml \
		${HELM_INCLUDE_OPTION} templates/crds/${CRD_VERSION}/iter8.tools_experiments.yaml \
		${HELM_INCLUDE_OPTION} templates/metrics/iter8_metrics.yaml \
		${HELM_INCLUDE_OPTION} templates/notifier/iter8_notifiers.yaml \
		--set istioTelemetry=${TELEMETRY_VERSION} \
		--set prometheusJobLabel=${PROMETHEUS_JOB_LABEL} \
	| kubectl apply -f -

# Deploy controller to the Kubernetes cluster configured in $KUBECONFIG or ~/.kube/config
deploy: manifests
	helm template ${HELM3_NAME} install/helm/iter8-controller ${HELM2_NAME} \
		--set image.repository=`echo ${IMG} | cut -f1 -d':'` \
		--set image.tag=`echo ${IMG} | cut -f2 -d':'` \
		${HELM_INCLUDE_OPTION} templates/default/namespace.yaml \
		${HELM_INCLUDE_OPTION} templates/default/serviceaccount.yaml \
		${HELM_INCLUDE_OPTION} templates/default/manager.yaml \
		${HELM_INCLUDE_OPTION} templates/crds/${CRD_VERSION}/iter8.tools_experiments.yaml \
		${HELM_INCLUDE_OPTION} templates/metrics/iter8_metrics.yaml \
		${HELM_INCLUDE_OPTION} templates/notifier/iter8_notifiers.yaml \
		${HELM_INCLUDE_OPTION} templates/rbac/role.yaml \
		${HELM_INCLUDE_OPTION} templates/rbac/role_binding.yaml \
		--set istioTelemetry=${TELEMETRY_VERSION} \
		--set prometheusJobLabel=${PROMETHEUS_JOB_LABEL} \
	| kubectl apply -f -

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	go generate ./pkg/... ./cmd/...

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

# Build default intall files. We generate three. They differ in the defintion of the default
# metrics. The label used in the job filter depends on the version of Istio.
build-default: manifests
	echo '# Generated by make build-default; DO NOT EDIT' > install/iter8-controller.yaml
	helm template ${HELM3_NAME} install/helm/iter8-controller ${HELM2_NAME} \
		--set istioTelemetry=v1 \
		--set prometheusJobLabel=istio-mesh \
	>> install/iter8-controller.yaml
	echo '# Generated by make build-default; DO NOT EDIT' > install/iter8-controller-telemetry-v2.yaml
	helm template ${HELM3_NAME} install/helm/iter8-controller ${HELM2_NAME} \
		--set istioTelemetry=v2 \
		--set prometheusJobLabel=envoy-stats \
	>> install/iter8-controller-telemetry-v2.yaml
	echo '# Generated by make build-default; DO NOT EDIT' > install/iter8-controller-telemetry-v2-17.yaml
	helm template ${HELM3_NAME} install/helm/iter8-controller ${HELM2_NAME} \
		--set istioTelemetry=v2 \
		--set prometheusJobLabel=kubernetes-pods \
	>> install/iter8-controller-telemetry-v2-17.yaml

.PHONY: changelog
changelog:
	@sed -n '/$(ver)/,/=====/p' CHANGELOG | grep -v $(ver) | grep -v "====="

tests:
	go test ./test/. -v
