[<- back to Home](../readme.md)

# Infra Setup

Expected Results:

- Azure subscription with follow resources running
  - AKS - Azure Kubernetes Services
  - ACR - Azure Container Registry
- Resources installed on AKS via Helm-Charts
  - RabbitMq cluster
  - Dapr
  - Keda

Login to Azure using the CLI:

```sh
az login
```

Set the default subscription:

```sh
az account set --subscription <subscription-id>
```

Create a resource group:

```sh
az group create --name rgdev001 --location brazilsouth
```

## 1. Create an AKS cluster and attach ACR

Create an AKS cluster:

```sh
az aks create --resource-group rgdev001 --name aksdev001 --node-count 2 --location brazilsouth --node-vm-size Standard_D4ds_v5 --enable-managed-identity --generate-ssh-keys --tier free 
```

Create an Container Registry:
```sh
az acr create --name acrdev001br --resource-group rgdev001 --sku basic
```

Attach the Container Registry to AKS:
```sh
az aks update --name aksdev001 --resource-group rgdev001 --attach-acr acrdev001br
```

Get the access credentials for the AKS cluster:

```sh
az aks get-credentials --resource-group rgdev001 --name aksdev001 --overwrite-existing
```

Verify the connection to the cluster:

```sh
kubectl cluster-info
```

Verify the two nodes are deployed:

```sh
kubectl get pod -A
```

## 2. Setup Dapr on AKS

Add a reference:

```sh
helm repo add dapr https://dapr.github.io/helm-charts/   
helm repo update
helm upgrade --install dapr dapr/dapr --namespace dapr-system --create-namespace
helm upgrade --install dapr-dashboard dapr/dapr-dashboard --namespace dapr-system --create-namespace
```

Verify if pods are running:

```sh
kubectl get pods -n dapr-system
```

## 3. Add a RabbitMq cluster to AKS

Add a reference:

```sh
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm upgrade --install rabbitmq-cluster-operator bitnami/rabbitmq-cluster-operator -n rabbitmq-system --create-namespace
```

Verify if pods are running:

```sh
kubectl get pods -n rabbitmq-system
```

## 4. Add Keda to AKS

Add a reference :

```sh
helm repo add kedacore https://kedacore.github.io/charts
helm repo update
helm upgrade --install keda kedacore/keda -n keda-system --create-namespace
helm upgrade --install keda-add-ons-http kedacore/keda-add-ons-http -n keda-system --create-namespace
 
```

Verify if pods are running:

```sh
kubectl get pods -n keda-system
```

## 6. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```sh
helm uninstall plant-trees-rbmq -n plant-trees
helm uninstall plant-trees-app -n plant-trees
helm uninstall keda-add-ons-http -n keda-system
helm uninstall keda -n keda-system
helm uninstall dapr -n dapr-system
helm uninstall rabbitmq-cluster-operator
helm uninstall rabbitmq -n rabbitmq-system
```

Delete all Azure resources:

```sh
az aks delete --name aksdev001 --resource-group rgdev001
az acr delete --name acrdev001br --resource-group rgdev001
az group delete --name rgdev001
```