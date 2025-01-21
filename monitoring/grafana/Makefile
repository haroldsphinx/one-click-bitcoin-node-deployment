BASE_DIR=base
RELEASE_NAME=grafana
CHART_NAME=grafana/grafana 
VERSION=8.8.4
NAMESPACE=monitoring
ENVIRONMENT=development

.PHONY: install-dependencies
install-dependencies:
	helm repo add grafana https://grafana.github.io/helm-charts
	helm repo update

.PHONY: create-configmap
create-configmap:
	kubectl create configmap grafana-dashboards --from-file=./dashboards/ -n $(NAMESPACE)

.PHONY: template
template: validate install-dependencies
	rm -rf ${BASE_DIR}/${RELEASE_NAME}
	helm template -n ${NAMESPACE} ${RELEASE_NAME} $(CHART_NAME) --version $(VERSION) --output-dir=./base -f ./values.yaml -f ./$(ENVIRONMENT)/values.yaml

.PHONY: create-configmap deploy
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
