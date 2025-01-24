ENVIRONMENT ?= development
NAMESPACE ?= bitcoin-node-exporter


.PHONY: deploy delete validate

deploy: validate
	# kubectl create namespace ${NAMESPACE}
	kustomize build ./${ENVIRONMENT} | kubectl apply -f -

delete: validate
	kustomize build ./${ENVIRONMENT} | kubectl delete -f -

.PHONY: validate
validate:
ifeq (${ENVIRONMENT},)
	$(error ENVIRONMENT is undefined,  Should be testnet or mainnet)
endif
