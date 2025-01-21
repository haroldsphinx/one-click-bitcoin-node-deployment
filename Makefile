.PHONY: setup-minikube-macos setup-minikube-ubuntu bitcoin-exporter

setup-minikube-macos:
	curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-darwin-arm64
	sudo install minikube-darwin-arm64 /usr/local/bin/minikube

setup-minikube-ubuntu:
	curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64
	sudo install minikube-linux-amd64 /usr/local/bin/minikube

start-minikube:
	minikube start 

nginx-ingress:
	minikube addons enable ingress

all:
	@echo "Setting up minikube"
	make setup-minikube-macos
	@echo "Starting minikube"
	make start-minikube
	@echo "Building bitcoin node exporter"
	cd bitcoin-node-exporter && make build-bitcoin-node-exporter
	@echo "Deploying bitcoin node exporter"
	cd k8s/bitcoin-exporter && make deploy
	@echo "Deploying prometheus"
	cd monitoring/prometheus && make deploy
	@echo "Deploying grafana"
	cd monitoring/grafana && make deploy
	@echo "Deploying promtail"
	cd monitoring/promtail && make deploy
	@echo "Deploying loki"
	cd monitoring/loki && make deploy


