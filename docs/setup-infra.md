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
az group create --name ${rg-name} --location ${location}
```

## 1. Create an AKS cluster and attach ACR

Create an AKS cluster:

```sh
az aks create --resource-group ${rg-name} --name ${aks-name} --node-count 2 --location ${location} --node-vm-size Standard_D4ds_v5 --enable-managed-identity --generate-ssh-keys --tier free 
```

Create an Container Registry:

```sh
az acr create --name ${acr-name} --resource-group ${rg-name} --sku basic
```

Attach the Container Registry to AKS:

```sh
az aks update --name ${aks-name} --resource-group ${rg-name} --attach-acr ${acr-name}
```

Get the access credentials for the AKS cluster:

```sh
az aks get-credentials --resource-group ${rg-name} --name ${aks-name} --overwrite-existing
```

Verify the connection to the cluster:

```sh
kubectl cluster-info
```

Create Service Bus and Topic :

```sh
az servicebus namespace create --resource-group ${rg-name} --name ${svcbus-name} --location ${location}
az servicebus topic create --name events --namespace-name ${svcbus-name} --resource-group ${rg-name}
az servicebus namespace authorization-rule keys list --resource-group ${rg-name} --namespace-name ${svcbus-name} --name RootManageSharedAccessKey --query primaryConnectionString --output tsv
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

## 4. Add a Redis cluster to AKS

Add a reference:

```sh
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm upgrade --install redis-cluster bitnami/redis -n redis-system --create-namespace
```

Verify if pods are running:

```sh
kubectl get pods -n redis-system
```

## 5. Add Keda to AKS

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
helm uninstall rabbitmq-cluster-operator -n rabbitmq-system

```

Delete all Azure resources:

```sh
az servicebus namespace delete --resource-group ${rg-name} --name ${svcbus-name}
az aks delete --name ${aks-name} --resource-group ${rg-name}
az acr delete --name ${acr-name} --resource-group ${rg-name}
az group delete --name ${rg-name}
```
