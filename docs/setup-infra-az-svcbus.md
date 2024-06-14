# Infra Setup - Azure Service Bus


## 1. Setup Azure Service Bus and Topic

Create Namespace:
```sh
az servicebus namespace create --resource-group {rgname} --name {svcbusname} --location {location}
```

Create Topic:

```sh
az servicebus topic create --name events --namespace-name {svcbusname} --resource-group {rgname}
```

Get ConnectionString value:

```sh
az servicebus namespace authorization-rule keys list --resource-group {rgname} --namespace-name {svcbusname} --name RootManageSharedAccessKey --query primaryConnectionString --output tsv
```

## 2. Setup Dapr and Keda Dependencies

Add a reference:

```sh
helm upgrade --install az-svcbus .helmcharts/az-svcbus -n tree --create-namespace
```

Verify if pods are running:

```sh
kubectl get scaledobjects -n tree
kubectl get components -n tree
```

## 6. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```sh
helm uninstall az-svcbus -n tree
```

Deleting azure resources:

```sh
az servicebus namespace delete --resource-group {rgname} --name {svcbusname}
```
