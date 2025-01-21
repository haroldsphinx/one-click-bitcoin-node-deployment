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
