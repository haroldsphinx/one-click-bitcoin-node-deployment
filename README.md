# one-click-bitcoin-node-deployment


# Briefs

## 1. Setup Minikube
There are two ways to setup minikube:
1. Using Ansible to setup minikube on ubuntu/macos
2. Using Minikube CLI to setup minikube on ubuntu/macos
for my setup, I'm using minkube cli on macos.

#### Using Ansible

```
ansible-playbook -i ansible/inventories/inventory.yml ansible/playbook.yml
```

#### Using Minikube CLI
```
make setup-minikube-macos
make start-minikube
```
P.S: Since its a local kubernetes cluster, I didnt want to use any elaborate tools like terraform to automate the setup so I just decided to use something basic, however my approach will be different for a production usecase.

## 2. Deploying Bitcoin Node
I kustomize to deploy the bitcoin node, you can find the kustomize files in the k8s/bitcoin-node directory.

The idea around this is for easy multi-environment deployment and to also avoid too many replicared k8s manifest files.

- Steps to Deploy
    - set environemnt variable to testnet or mainnet
    - set namespace variable to bitcoin-node
    - create namespace if not exists
    - deploy the bitcoin node

```
kubectl create namespace ${NAMESPACE}
make deploy
```

### 3. Creating a bitcoin-node exporter service in Golang
- to build locally on any architecture, you can use the following command:
```
make build-bitcoin-node-exporter
```
to run and test locally
```

./cmd/bitcoin-exporter/bitcoin-exporter
```

- to build with docker
```
make build
```

- to deploy to local kubernetes cluster, ensure you have a namespace called bitcoin-node-exporter
```
make deploy
```

### 4. Setup Monitoring
- Deploy Proemtheus
```
make deploy
```

# Things I will do differently in a production setup
# Reference
https://github.com/bitcoin/bitcoin/blob/master/doc/bitcoin-conf.md
github.com/btcsuite/btcd/rpcclient
https://prometheus.io/docs/guides/go-application/
https://developer.bitcoin.org/reference/rpc/
https://pkg.go.dev/github.com/btcsuite/btcd/rpcclient
https://prometheus.io/docs/instrumenting/writing_exporters/


## Notes
- For setting up Local Kubernetes Cluster, I decided to use Minikube, to setup minikube, I used two approaches: 1. Using Ansible to setup minikube on macos and 2. Using Minikube CLI to setup minikube on ubuntu.

```
make setup-minikube-macos
make start-minikube
```

for ansible setup, you can use the following command:
```
ansible-playbook -i ansible/inventories/inventory.yml ansible/playbook.yml
```


# References
- https://github.com/bitcoin/bitcoin/blob/master/doc/bitcoin-conf.md
- https://github.com/btcsuite/btcd/rpcclient
- https://prometheus.io/docs/guides/go-application/
- https://developer.bitcoin.org/reference/rpc/
- https://pkg.go.dev/github.com/btcsuite/btcd/rpcclient
- https://prometheus.io/docs/instrumenting/writing_exporters/
