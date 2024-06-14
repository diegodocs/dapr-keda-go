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
az account set --subscription {subid}
```

Create a resource group:

```sh
az group create --name {rgname} --location {location}
```

## 1. Create an AKS cluster and attach ACR

Create an AKS cluster:

```sh
az aks create --resource-group {rgname} --name {aksname} --node-count 2 --location {location} --node-vm-size Standard_D4ds_v5 --tier free 
```

Create an Container Registry:

```sh
az acr create --name {acrname} --resource-group {rgname} --sku basic
```

Attach the Container Registry to AKS:

```sh
az aks update --name {aksname} --resource-group {rgname} --attach-acr {acrname}
```

Get the access credentials for the AKS cluster:

```sh
az aks get-credentials --resource-group {rgname} --name {aksname} --overwrite-existing
```

Verify the connection to the cluster:

```sh
kubectl cluster-info
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

## 3. Add Keda to AKS

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

## 3. Setup Transport Layer

In this project, we have 3 different options to configure the transport layer (choose one):

- [Azure Service Bus](setup-infra-az-svcbus.md)
- [Redis](setup-infra-redis.md)
- [RabbitMq](setup-infra-rbmq.md)

## 6. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```sh
helm uninstall keda-add-ons-http -n keda-system
helm uninstall keda -n keda-system
helm uninstall dapr -n dapr-system
```

Delete all Azure resources:

```sh
az servicebus namespace delete --resource-group {rgname} --name {svcbusname}
az aks delete --name {aksname} --resource-group {rgname}
az acr delete --name {acrname} --resource-group {rgname}
az group delete --name {rgname}
```
