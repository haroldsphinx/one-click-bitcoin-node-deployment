BASE_DIR=base
RELEASE_NAME=loki
CHART_NAME=grafana/loki
VERSION=3.0.0
NAMESPACE=monitoring
ENVIRONMENT=development

.PHONY: install-dependencies
install-dependencies:
	helm repo add grafana https://grafana.github.io/helm-charts
	helm repo update

.PHONY: template
template: validate
	rm -rf ${BASE_DIR}/${RELEASE_NAME}
	helm template -n ${NAMESPACE} ${RELEASE_NAME} $(CHART_NAME) --version $(VERSION) -f ./values.yaml -f ./$(ENVIRONMENT)/values.yaml

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
