[<- back to Home](../readme.md)

# App Setup

Expected Results:

- Build images via dockerfile
- Push images to ACR - Azure Container Registry
- Deploy producer and consumer apps via Helm-Charts

## 1. Running Applications Locally

### Initialize Dapr locally:

```sh
dapr init
```

### Runing consumer app:

```sh
dapr run --app-id consumer1 --app-protocol http --dapr-http-port 3500 --app-port 8080  --resources-path .helmcharts/app/templates -- go run ./cmd/consumer
```

### Runing producer app:

```sh
dapr run --app-id producer1 --app-protocol http --dapr-http-port 3501 --resources-path .helmcharts/app/templates -- go run ./cmd/producer
```

### or you can run on unique terminal:

```sh
dapr run -f ./dapr.yaml
```

### You can also initialize Dapr on the provisioned AKS cluster (just for debugging purposal)

```sh
dapr init --kubernetes --wait
dapr status -k
```

## 2. Deploy applications to AKS

### Building images

```sh
docker build -t "acrdev001br.azurecr.io/consumer-app:1.0.0" -f cmd/consumer/dockerfile .
docker build -t "acrdev001br.azurecr.io/producer-app:1.0.0" -f cmd/producer/dockerfile .
```

### Pushing images to ACR

```sh
az acr login --name acrdev001br
docker push "acrdev001br.azurecr.io/consumer-app:1.0.0" 
docker push "acrdev001br.azurecr.io/producer-app:1.0.0" 
```

### Deploying App helm-charts

```sh
helm upgrade --install plant-trees-rbmq .helmcharts/rbmq -n plant-trees --create-namespace
helm upgrade --install plant-trees-app .helmcharts/app -n plant-trees --create-namespace
```

### Dapr Dashboard

#### To access the Dapr dashboard, run the following command:

```sh
dapr dashboard -k
```

#### Expected response:

```sh
Dapr dashboard found in namespace: dapr-system
Dapr dashboard available at http://localhost:8080
```

## 3. Testing the application

### Building images

```sh
# Reviewing Logs
kubectl logs -f -l app=consumer1 --all-containers=true -n plant-trees

# Create a port
kubectl port-forward pod/producer1 8081 8081 -n plant-trees

# Send post to producer app
- POST -> http://localhost:8081/plant
- Json Body: {"numberOfTrees":1}

# Review pod instances and status
kubectl get pod -l app=consumer1 -n plant-trees
```

Explore the dashboard to drill down into the applications, components, and services.
