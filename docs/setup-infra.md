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

```powershell
az login
```

Replace follow texts with correct values based on your environment:

```powershell
- $SubscriptionID = ''
- $Location = ''
- $ResourceGroupName = ''
- $AKSClusterName = ''
- $ContainerRegistryName = ''
- $ServiceBusNamespace = ''
```

Set the default subscription:

```powershell
az account set --subscription $SubscriptionID
```

Create a resource group:

```powershell
az group create --name $ResourceGroupName --location $Location
```

## 1. Create an AKS cluster and attach ACR

Create an AKS cluster:

```powershell
az aks create --resource-group $ResourceGroupName --name $AKSClusterName --node-count 3 --location $Location --node-vm-size Standard_D4ds_v5 --tier free --enable-pod-identity --network-plugin azure --generate-ssh-keys
```

Create an Container Registry:

```powershell
az acr create --name $ContainerRegistryName --resource-group $ResourceGroupName --sku basic
```

Attach the Container Registry to AKS:

```powershell
az aks update --name $AKSClusterName --resource-group $ResourceGroupName --attach-acr $ContainerRegistryName
```

Get the access credentials for the AKS cluster:

```powershell
az aks get-credentials --resource-group $ResourceGroupName --name $AKSClusterName --overwrite-existing
```

Verify the connection to the cluster:

```powershell
kubectl cluster-info
```

## 2. Setup Dapr on AKS

Add a reference:

```powershell
helm repo add dapr https://dapr.github.io/helm-charts/   
helm repo update
helm upgrade --install dapr dapr/dapr --namespace dapr-system --create-namespace
helm upgrade --install dapr-dashboard dapr/dapr-dashboard --namespace dapr-system --create-namespace
```

Verify if pods are running:

```powershell
kubectl get pods -n dapr-system
```

### Dapr Dashboard

#### To access the Dapr dashboard, run the following command

```powershell
dapr dashboard -k
```

#### Expected response

```powershell
Dapr dashboard found in namespace: dapr-system
Dapr dashboard available at http://localhost:8080
```

## 3. Add Keda to AKS

Add a reference :

```powershell
helm repo add kedacore https://kedacore.github.io/charts
helm repo update
helm upgrade --install keda kedacore/keda -n keda-system --create-namespace
helm upgrade --install keda-add-ons-http kedacore/keda-add-ons-http -n keda-system --create-namespace
 
```

Verify if pods are running:

```powershell
kubectl get pods -n keda-system
```

## 4. Setup Transport Layer with Dapr and Keda

In this project, we have 3 different options to configure the transport layer (choose one):

- [Azure Service Bus](setup-infra-azsbus.md)
- [Redis](setup-infra-redis.md)
- [RabbitMq](setup-infra-rbmq.md)

## 5. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```powershell
helm uninstall keda-add-ons-http -n keda-system
helm uninstall keda -n keda-system
helm uninstall dapr -n dapr-system
```

Delete all Azure resources:

```powershell
az aks delete --name $AKSClusterName --resource-group $ResourceGroupName
az acr delete --name $ContainerRegistryName --resource-group $ResourceGroupName
az group delete --name $ResourceGroupName
```
