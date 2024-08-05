# Dapr Keda Go App

[![App-Go-Build-Test](https://github.com/diegodocs/go-dapr-plant-trees/actions/workflows/app-go-build-test.yml/badge.svg?branch=main)](https://github.com/diegodocs/go-dapr-plant-trees/actions/workflows/app-go-build-test.yml)

## Summary

This project cover concepts about Distributed Architecture combining follow technologies:

- Go - Producer/Consumer App
- Distributed Application Runtime - DAPR
- Kubernetes Event Driven Autoscaling - KEDA
- Azure Kubernetes Service (AKS)
- Azure Container Registry (ACR)
- Azure Service Bus  (ASB)
- RabbitMQ:(exchange, binding and queues)
- Cache Redis

## Development Tools

- [Go SDK](https://go.dev/dl/)
- [Azure CLI](https://learn.microsoft.com/pt-br/cli/azure/install-azure-cli)
- [DAPR CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
- [Kubectl](https://kubernetes.io/pt-br/docs/tasks/tools/)
- [Helm CLI](https://github.com/helm/helm)
- [GIT bash](https://git-scm.com/downloads)
- [Visual Studio Code](https://code.visualstudio.com/download)

## Deploying Infra and Apps

- [Setup your environment(infra)](./docs/setup-infra.md)
- [Deploying  your applications on AKS](./docs/setup-app.md)

You can see the folder `.github/workflows` the pipelines (actions) for build and deploy:

- [Configuring a federated credential and connect GitHub Actions to Azure](https://learn.microsoft.com/en-us/azure/developer/github/connect-from-azure)

## References

- [DAPR - Pros/Cons](./docs/dapr-pros-cons.md)
- [KEDA  - Pros/Cons](./docs/keda-pros-cons.md)

## You can find in this repository

- `cmd`: main application code
- `internal`: reusable library code
- `tests`: automated tests(Integration, Smoke, Acceptance)
- `.docker`: app dockerfile
- `.helmcharts`: helmcharts for deployment

## You shouldn't find

- `src` directory (considering is not standard for Go projects)
- Binaries committed to source control.
- Unnecessary project/library references or third party frameworks.
