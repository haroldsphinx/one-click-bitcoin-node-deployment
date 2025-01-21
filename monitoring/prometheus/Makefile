BASE_DIR=base
RELEASE_NAME=prometheus
CHART_NAME=prometheus-community/prometheus 
VERSION=15.8.0
NAMESPACE=monitoring
ENVIRONMENT=development

.PHONY: install-dependencies
install-dependencies:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add kube-state-metrics https://kubernetes.github.io/kube-state-metrics
	helm repo update

.PHONY: template
template: validate install-dependencies
	rm -rf ${BASE_DIR}/${RELEASE_NAME}
	helm template -n ${NAMESPACE} ${RELEASE_NAME} $(CHART_NAME) --version $(VERSION) --output-dir=./base -f ./values.yaml -f ./$(ENVIRONMENT)/values.yaml

.PHONY: deploy
deploy: validate install-dependencies
	helm upgrade --install -n ${NAMESPACE} $(RELEASE_NAME) $(CHART_NAME) --version $(VERSION) -f ./values.yaml -f ./$(ENVIRONMENT)/values.yaml

.PHONY: delete
delete: install-dependencies
	helm uninstall -n ${NAMESPACE} $(RELEASE_NAME)

.PHONY: validate
validate:
ifeq (${ENVIRONMENT},)
	$(error ENVIRONMENT is undefined. Should be development, staging or production)
endif
