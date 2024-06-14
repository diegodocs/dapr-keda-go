# Infra Setup - RabbitMQ

Expected Results:

- Deploy RabbitMQ Cluster, User, Queues, Exchanges and Bindings
- Deploy dapr and keda configuration via Helm-Charts

## 1. Deploying Helm-chart: RabbitMQ Cluster Operator

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

## 2. Setup Dapr and Keda Dependencies

```sh
helm upgrade --install rbmq .helmcharts/rbmq -n tree --create-namespace
```

Verify if pods are running:

```sh
kubectl get pods -n tree
kubectl get scaledobjects -n tree
kubectl get components -n tree
kubectl get queues -n tree
```

## 3. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```sh

helm uninstall rbmq -n tree
helm uninstall rabbitmq-cluster-operator -n rabbitmq-system

```
