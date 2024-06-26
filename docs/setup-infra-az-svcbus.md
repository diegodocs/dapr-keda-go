# Infra Setup - Azure Service Bus

Expected Results:

- Deploy Azure Service Bus Namespace, Topics
- Deploy dapr and keda configuration via Helm-Charts

## 1. Setup Azure Service Bus and Topic

Create Namespace:

```powershell
az servicebus namespace create --resource-group $ResourceGroupName --name $ServiceBusNamespace --location $Location --sku basic
```

Create Topic:

```powershell
az servicebus queue create --name events --namespace-name $ServiceBusNamespace --resource-group $ResourceGroupName
```

Get ConnectionString value:

```powershell
az servicebus namespace authorization-rule keys list --resource-group $ResourceGroupName --namespace-name $ServiceBusNamespace --name RootManageSharedAccessKey --query primaryConnectionString --output tsv


Replace text '$ServiceBusEndPoint' by value above

```

Add authorization to KEDA monitor the queue:

```powershell
az servicebus queue authorization-rule create --resource-group $ResourceGroupName --namespace-name $ServiceBusNamespace --queue-name events --name keda-monitor --rights Listen


```

## 2. Setup Dapr and Keda Dependencies

Add a reference:

```powershell
helm upgrade --install azsbus .helmcharts/azsbus -n tree --create-namespace
```

Verify if pods are running:

```powershell
kubectl get scaledobjects -n tree
kubectl get components -n tree
```

## 3. Clean-up

Follow these steps to remove all the apps, components and cloud resources created in this how-to guide.

```powershell
helm uninstall azsbus -n tree
```

Deleting azure resources:

```powershell
az servicebus namespace delete --resource-group $ResourceGroupName --name $ServiceBusNamespace
```
