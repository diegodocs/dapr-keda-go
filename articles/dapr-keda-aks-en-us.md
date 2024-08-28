# DAPR, KEDA on AKS (Azure Kubernetes Services): step by step

The CNCF (Cloud Native Computing Foundation) defines Cloud-Native Applications as software programs that consist of several small and interdependent services called microservices. These applications are designed to take full advantage of innovations in cloud computing, such as scalability, flexibility, and automation.

Some of the best-known Cloud-Native projects from the CNCF are: Kubernetes, Prometheus, Envoy, Jaeger, Helm, DAPR, KEDA, and others.

In this article, we will discuss DAPR, KEDA, and how their combination can bring efficiency and flexibility to building applications in Kubernetes.

## DAPR: Any language, any framework, anywhere

![DAPR](https://docs.dapr.io/images/overview.png)
The [DAPR](https://github.com/dapr/community/blob/master/README.md) is a portable, serverless, event-driven runtime that makes it easy for developers to build resilient, stateless and stateful microservices that run on the cloud and edge and embraces the diversity of languages and developer frameworks.

## KEDA: Kubernetes Event-driven Autoscaling

![KEDA](https://keda.sh/img/keda-arch.png)
KEDA makes application autoscaling simple by applying event-driven autoscaling to scale your application based on demand. It allows you to scale workloads from 0 to N instances efficiently (scale-to-zero), which means your application can dynamically scale down to zero instances when not in use, reducing costs.

<img src="https://github.com/diegodocs/dapr-keda-go/blob/main/docs/plant-tree.png?raw=true" alt="Plant Tree Apps" style="width:200px;"/>

**I am excited to share this [GitHub Repository](https://github.com/diegodocs/dapr-keda-go) that covers concepts about Cloud-Native Architecture combining follow technologies:**

- Go - Producer/Consumer App
- Distributed Application Runtime - DAPR
- Kubernetes Event Driven Autoscaling - KEDA
- Azure Kubernetes Service (AKS)
- Azure Container Registry (ACR)
- Azure Service Bus (ASB)

## Development Tools

- [Go SDK](https://go.dev/dl/)
- [Azure CLI](https://learn.microsoft.com/pt-br/cli/azure/install-azure-cli)
- [DAPR CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
- [Kubectl](https://kubernetes.io/pt-br/docs/tasks/tools/)
- [Helm CLI](https://github.com/helm/helm)
- [GIT bash](https://git-scm.com/downloads)
- [Visual Studio Code](https://code.visualstudio.com/download)

## Deploying Infrastructure

- [click here for complete step-by-step](https://github.com/diegodocs/dapr-keda-go/docs/setup-infra.md)

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

### 1. Create an AKS cluster and attach ACR

Create an AKS cluster:

```powershell
az aks create --resource-group $ResourceGroupName --name $AKSClusterName --node-count 3 --location $Location --node-vm-size Standard_D4ds_v5 --tier free --enable-pod-identity --network-plugin azure --generate-ssh-keys
```

- enabling [DAPR AKS Extension](https://learn.microsoft.com/pt-br/azure/aks/dapr?tabs=cli)
- enabling [KEDA AKS Addon](https://learn.microsoft.com/pt-br/azure/aks/keda-deploy-add-on-cli)

Create a Container Registry:

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

### 2. Setup DAPR on AKS

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

### 2.1  DAPR Dashboard

#### To access the DAPR dashboard, run the following command

```powershell
dapr dashboard -k
```

**Expected response:**

```powershell
DAPR dashboard found in namespace: dapr-system
DAPR dashboard available at http://localhost:8080
```

### 3. Add KEDA to AKS

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

### 4. Setup Transport Layer with DAPR and KEDA

In this project, we have three different options to configure the transport layer (choose one):

- [Azure Service Bus](https://github.com/diegodocs/dapr-keda-go/setup-infra-azsbus.md)
- [Redis](https://github.com/diegodocs/dapr-keda-go/setup-infra-redis.md)
- [RabbitMq](https://github.com/diegodocs/dapr-keda-go/setup-infra-rbmq.md)

### 5. Clean-up

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

## Deploy applications to AKS

- [click here for complete step-by-step](<https://github.com/diegodocs/dapr-keda-go/docs/setup-app.md>)

### 1. Building images

```powershell
az acr login --name $ContainerRegistryName
docker build -t "$ContainerRegistryName.azurecr.io/consumer-app:1.0.0" -f cmd/consumer/dockerfile .
docker build -t "$ContainerRegistryName.azurecr.io/producer-app:1.0.0" -f cmd/producer/dockerfile .
```

### 2. Pushing images to ACR

```powershell

docker push "$ContainerRegistryName.azurecr.io/consumer-app:1.0.0" 
docker push "$ContainerRegistryName.azurecr.io/producer-app:1.0.0" 
```

### 3. Setup DAPR and KEDA Dependencies

```powershell
helm upgrade --install app .helmcharts/app -n tree --create-namespace
```

Verify if pods are running:

```powershell
kubectl get pods -n tree
```

### 4. Testing the application

```powershell
# Reviewing Logs
kubectl logs -f -l app=consumer1 --all-containers=true -n tree

# Create a port
kubectl port-forward pod/producer1 8081 8081 -n tree

# Send post to producer app
- POST -> http://localhost:8081/plant
- Json Body: {"numberOfTrees":100}

# Review pod instances and status
kubectl get pod -l app=consumer1 -n tree
```

### 4. Clean-up

```powershell
helm uninstall app -n tree
```

## References

- [DAPR KEDA GO Project](https://github.com/diegodocs/dapr-keda-go)
- [DAPR - Pros/Cons](https://github.com/diegodocs/dapr-keda-go/docs/dapr-pros-cons.md)
- [KEDA  - Pros/Cons](https://github.com/diegodocs/dapr-keda-go/docs/keda-pros-cons.md)
