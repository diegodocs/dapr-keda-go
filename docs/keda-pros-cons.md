# Kubernetes Cluster scaling using KEDA

It's a Cloud Native Computing Foundation (CNCF) graduated project.

![](https://keda.sh/img/keda-arch.png)

## Pros

- Simplified Autoscaling: KEDA makes application autoscaling simple by applying event-driven autoscaling to scale your application based on demand. It allows you to scale workloads from 0 to N instances efficiently.
- Cost-Efficient: KEDA enables scale-to-zero, which means your application can dynamically scale down to zero instances when not in use, reducing costs.
- Rich Catalog of Scalability Options: The KEDA add-on provides a rich catalog of Azure KEDA scalers that you can use to scale your applications. These include built-in scalers for various Azure services.
- Production-Grade Security: KEDA decouples autoscaling authentication from workloads, ensuring secure scaling operations.
- Custom Scalability Decisions: You can bring your own external scaler to tailor autoscaling decisions according to your specific requirements.
- Architectural standardization: all applications running inside k8s

## Cons

- HTTP Add-on Limitation: The KEDA HTTP add-on (preview) for scaling HTTP workloads isn’t installed with the extension by default but can be deployed separately.
- Azure Cosmos DB External Scaler: The external scaler for Azure Cosmos DB, which scales based on Azure Cosmos DB change feed, isn’t installed with the extension but can also be deployed separately.
- Single Metric Server: Only one metric server is allowed in the Kubernetes cluster, and KEDA must be the only installed metric adapter.
- Multiple KEDA Installations Not Supported: Deploying multiple KEDA installations within the same cluster is not supported.
- Manual integration with SDKs: for input/output bindings, for example

## References

- [Simplified application autoscaling with Kubernetes Event-driven Autoscaling (KEDA) add-on](https://learn.microsoft.com/en-us/azure/aks/keda-about)
- [Dotnet Sample Producer/Consumer using KEDA on github](https://github.com/kedacore/sample-dotnet-worker-servicebus-queue)