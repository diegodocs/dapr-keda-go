# Planting Trees App

## Summary

This project cover concepts about Distributed Architecture combining follow technologies:

- Go - Producer/Consumer App
- Dapr
- Keda
- Azure Kubernetes Service (AKS)
- Azure Container Registry (ACR)
- RabbitMq:(exchange, binding and queues)
- Redis

## Distributed Application Runtime - DAPR

![dapr](https://docs.dapr.io/images/overview.png)

**[DAPR - Pros/Cons](./docs/dapr-pros-cons.md)**

## Kubernetes Event Driven Autoscaling - KEDA

![keda](https://keda.sh/img/keda-arch.png)

**[KEDA  - Pros/Cons](./docs/keda-pros-cons.md)**

## Development Tools

- [Go 1.22.3](https://go.dev/dl/)
- [Azure CLI](https://learn.microsoft.com/pt-br/cli/azure/install-azure-cli)
- [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
- [Kubectl](https://kubernetes.io/pt-br/docs/tasks/tools/)
- [Helm CLI](https://github.com/helm/helm)
- [GIT bash](https://git-scm.com/downloads)

## Setup your environment

- [Setup Infra steps](./docs/setup-infra.md)

## Restore, Build and Test

```sh
go vet ./...
go test ./...
go build ./...
```

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
