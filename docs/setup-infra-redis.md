# Infra Setup - Redis

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

## 6. Clean-up

```sh
helm uninstall redis-cluster-n redis-system
```
