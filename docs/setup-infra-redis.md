# Infra Setup - Redis

Expected Results:

- Deploy Redis Cluster on AKS
- Deploy dapr and keda configuration via Helm-Charts

## 1. Deploying Helm-chart: Redis Operator

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

## 2. Setup Dapr and Keda Dependencies

Add a reference:

```sh
helm upgrade --install redis .helmcharts/redis -n tree --create-namespace
```

Verify if pods are running:

```sh
kubectl get scaledobjects -n tree
kubectl get components -n tree
```

## 3. Clean-up

```sh
helm uninstall redis - tree
helm uninstall redis-cluster-n redis-system
```
