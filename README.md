# Planting Trees App

## Summary

This project cover concepts about Distributed Architecture combining follow technologies:

- Go Producer/Consumer App
- Dapr
- Keda
- Azure Kubernetes Service (AKS)

## Long story short

As a developer, I want to implement Distributed Architecture Solution :

- open to work with multiple languages/tech-stacks based on business context
- keep code testable
- keep ecosystem observable
- easy/transparent troubleshooting
- open for abstraction instead tech-stack high-coupling(SDKs, 3rd party libraries and etc )

## Distributed Application Runtime - DAPR

![dapr](https://docs.dapr.io/images/overview.png)

**[DAPR - Pros/Cons](./docs/dapr-pros-cons.md)**

## Kubernetes Event Driven Autoscaling - KEDA

![keda](https://keda.sh/img/keda-arch.png)

**[KEDA  - Pros/Cons](./docs/keda-pros-cons.md)**

## Development Tools

- [Go 1.22.3](https://go.dev/dl/)
- [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
- [kubectl](https://kubernetes.io/pt-br/docs/tasks/tools/)
- [helm](https://github.com/helm/helm)
- [git bash](https://git-scm.com/downloads)

## Setup your environment

- [Infra steps](./docs/setup-infra.md)
- [App steps](./docs/setup-app.md)

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
  - RabbitMq:(exchange, binding and queues)
  - Dapr(bindings, pubsub, sub, resilience etc)
  - Keda(ScaledObjects)

## You shouldn't find

- `src` directory (considering is not standard for Go projects)
- Binaries committed to source control.
- Unnecessary project/library references or third party frameworks.