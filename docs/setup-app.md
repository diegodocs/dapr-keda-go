# App Setup

Expected Results:

- Build images via dockerfile
- Push images to ACR - Azure Container Registry
- Deploy producer and consumer apps via Helm-Charts

## 1. Running Applications Locally

### Initialize DAPR locally

```powershell
dapr init
```

### Running consumer app

```powershell
dapr run --app-id consumer1 --app-protocol http --dapr-http-port 3500 --app-port 8080  --resources-path .dapr/resources -- go run ./cmd/consumer
```

### Running producer app

```powershell
dapr run --app-id producer1 --app-protocol http --dapr-http-port 3501 --resources-path .dapr/resources -- go run ./cmd/producer
```

### or you can run on unique terminal

```powershell
dapr run -f ./dapr.yaml
```

### You can also initialize DAPR on the provisioned AKS cluster (just for debugging purpose)

```powershell
dapr init --kubernetes --wait
dapr status -k
```

## 2. Deploy applications to AKS

### Building images

```powershell
az acr login --name $ContainerRegistryName
docker build -t "$ContainerRegistryName.azurecr.io/consumer-app:1.0.0" -f cmd/consumer/dockerfile .
docker build -t "$ContainerRegistryName.azurecr.io/producer-app:1.0.0" -f cmd/producer/dockerfile .
```

### Pushing images to ACR

```powershell

docker push "$ContainerRegistryName.azurecr.io/consumer-app:1.0.0" 
docker push "$ContainerRegistryName.azurecr.io/producer-app:1.0.0" 
```

### 2. Setup DAPR and KEDA Dependencies

```powershell
helm upgrade --install app .helmcharts/app -n tree --create-namespace
```

Verify if pods are running:

```powershell
kubectl get pods -n tree
```

## 3. Testing the application

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

## 4. Clean-up

```powershell
helm uninstall app - tree
```
