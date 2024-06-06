[<- back to Home](../readme.md)

# Backend integration using (DAPR) Distributed Application Runtime

The [Dapr](https://github.com/dapr/community/blob/master/README.md) is a portable, serverless, event-driven runtime that makes it easy for developers to build resilient, stateless and stateful microservices that run on the cloud and edge and embraces the diversity of languages and developer frameworks.

Dapr provides a set of APIs that run as a sidecar process next to your application code, and offer various capabilities such as service invocation, state management, pub/sub, bindings, actors, observability, and secrets.

# Any language, any framework, anywhere

![Dapr](https://docs.dapr.io/images/overview.png)

## Pros

- You can leverage the power and flexibility of Kubernetes, which is a fully managed service on Azure, to run your serverless applications on any environment, such as public cloud, private cloud, hybrid cloud, or edge devices.
- You can use AKS, Azure Functions, Azure Logic Apps, Azure Event Grid, Azure Event Hubs, Azure Monitor, and Azure Application Insights to automate, extend, integrate, and monitor your Dapr applications with various tools and services on Azure.
- You can use Dapr to simplify the development and deployment of microservices, as Dapr abstracts away the underlying technology choices and provides a consistent and portable API for your application logic. Its is clear implementation of strategy pattern which you can have different strategies implemented/defined outside your code(yaml files).
- You can use Dapr to improve the reliability, security, and resilience of your microservices, as Dapr enables automatic mTLS authentication and encryption, distributed tracing, and stateful actors for your application components.

## Cons

- You may need to learn and adopt new concepts and tools, such as Dapr components, sidecars, and CLI, to use Dapr effectively on Azure.
- You may face some Limits and challenges, such as compatibility issues, performance overhead, and debugging difficulties, when using Dapr with some Azure services or features.
- You may incur additional costs, such as storage, networking, and compute charges, depending on your Dapr usage and configuration.
- Dapr extension for Azure Functions is in preview and it is only supported in Azure Container Apps environments.

## References

- [Dapr and service meshes](https://docs.dapr.io/concepts/service-mesh/)
- [Dapr Extension for Azure Functions](https://learn.microsoft.com/en-us/azure/azure-functions/functions-bindings-dapr?tabs=in-process%2Cpreview-bundle-v4x%2Cbicep1&pivots=programming-language-csharp)