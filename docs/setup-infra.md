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
az group create --name <resource-group-name> --location <location>
```

## 1. Create an AKS cluster and attach ACR

Create an AKS cluster:

```sh

az feature register --namespace "Microsoft.ContainerService" --name "EnablePodIdentityPreview"

az aks create --resource-group <resource-group-name> --name <aks-name> --node-count 3 --location <location> --node-vm-size Standard_D4ds_v5 --tier free --enable-pod-identity --network-plugin azure --generate-ssh-keys
```

Create an Container Registry:

```sh
az acr create --name <acr-name> --resource-group <resource-group-name> --sku basic
```

Attach the Container Registry to AKS:

```sh
az aks update --name <aks-name> --resource-group <resource-group-name> --attach-acr <acr-name>
```

Get the access credentials for the AKS cluster:

```sh
az aks get-credentials --resource-group <resource-group-name> --name <aks-name> --overwrite-existing
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

### Dapr Dashboard

#### To access the Dapr dashboard, run the following command

```sh
dapr dashboard -k
```

#### Expected response

```sh
Dapr dashboard found in namespace: dapr-system
Dapr dashboard available at http://localhost:8080
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

## 3. Deploying  your applications on AKS

- [Setup App steps](setup-app.md)

## 4. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```sh
helm uninstall keda-add-ons-http -n keda-system
helm uninstall keda -n keda-system
helm uninstall dapr -n dapr-system
```

Delete all Azure resources:

```sh
az aks delete --name <aks-name> --resource-group <resource-group-name>
az acr delete --name <acr-name> --resource-group <resource-group-name>
az group delete --name <resource-group-name>
```
