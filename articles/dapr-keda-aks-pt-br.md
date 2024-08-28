# DAPR, KEDA no AKS (Azure Kubernetes Services): step by step

A **CNCF (Cloud Native Computing Foundation)** define **Aplicações Cloud-Native** como software que consistem em vários serviços pequenos e interdependentes chamados **microsserviços**. Essas aplicações são projetadas para aproveitar ao máximo as inovações em computação em nuvem, como **escalabilidade**, **segurança**, **flexibilidade** e **automação**.

Alguns dos projetos **Cloud-Native** mais conhecidos da **CNCF** são: Kubernetes, Prometheus, Envoy, Jaeger, Helm, DAPR, KEDA e etc.

Neste artigo, falaremos sobre DAPR, KEDA e como esta combinação pode trazer eficiência e flexibilidade na construção de aplicações em Kubernetes.

## DAPR: Distributed Application Runtime

![DAPR](https://docs.dapr.io/images/overview.png)
O [DAPR](https://github.com/dapr/community/blob/master/README.md) é um runtime portátil, serverless orientado a eventos que facilita a vida dos desenvolvedores na construção de microsserviços resilientes, principalmente no gerenciamento de estado, envio e consumo de mensagens via broker etc, usando uma abordagem de mesh, que simplifique bastante seu gerenciamento e também abraça a diversidade de linguagens e frameworks.

## KEDA: Kubernetes Event Driven Autoscaling

![KEDA](https://keda.sh/img/keda-arch.png)
O KEDA torna o escalonamento automático de aplicativos simples, aplicando escalonamento automático orientado a eventos para escalar seu aplicativo com base na demanda ( filas, tópicos etc). Ele permite que você escale cargas de trabalho de 0 a N instâncias de forma eficiente (escala para zero), o que significa que seu aplicativo pode escalar dinamicamente para zero instâncias quando não estiver em uso, ajudando bastante em cenários de otimização de custos.

<img src="https://github.com/diegodocs/dapr-keda-go/blob/main/docs/plant-tree.png?raw=true" alt="Plant Tree Apps" style="width:200px;"/>

**Desta forma, aproveitei para montar este repositório no [GitHub](https://github.com/diegodocs/dapr-keda-go) chamado "App-Plant-Tree" que cobre conceitos sobre Arquitetura Cloud-Native combinando as seguintes tecnologias:**

- Go - Producer/Consumer App
- Distributed Application Runtime - DAPR
- Kubernetes Event Driven Autoscaling - KEDA
- Azure Kubernetes Service (AKS)
- Azure Container Registry (ACR)
- Azure Service Bus (ASB)

## Ferramentas para desenvolvimento

- [Go SDK](https://go.dev/dl/)
- [Azure CLI](https://learn.microsoft.com/pt-br/cli/azure/install-azure-cli)
- [DAPR CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
- [Kubectl](https://kubernetes.io/pt-br/docs/tasks/tools/)
- [Helm CLI](https://github.com/helm/helm)
- [GIT bash](https://git-scm.com/downloads)
- [Visual Studio Code](https://code.visualstudio.com/download)

## Configurando a Infraestrutura

- [veja passo a passo completo](https://github.com/diegodocs/dapr-keda-go/docs/setup-infra.md)

Logando no Azure usando a linha de comando(CLI):

```powershell
az login
```

Substitua os valores das variáveis abaixo conforme seu ambiente:

```powershell
- $SubscriptionID = ''
- $Location = ''
- $ResourceGroupName = ''
- $AKSClusterName = ''
- $ContainerRegistryName = ''
- $ServiceBusNamespace = ''
```

Aponte para sua subscription:

```powershell
az account set --subscription $SubscriptionID
```

Criando seu resource group:

```powershell
az group create --name $ResourceGroupName --location $Location
```

### 1. Criando seu cluster AKS e conecte o ACR

Crie o cluster AKS (Azure Kubernetes Service):

```powershell
az aks create --resource-group $ResourceGroupName --name $AKSClusterName --node-count 3 --location $Location --node-vm-size Standard_D4ds_v5 --tier free --enable-pod-identity --network-plugin azure --generate-ssh-keys
```

De forma opcional, mas benéfica por trazer o suporte da Microsoft(veja detalhes nos links abaixo), você pode você pode instalar DAPR e KEDA via extension/addons.

- [DAPR AKS Extension](https://learn.microsoft.com/pt-br/azure/aks/dapr?tabs=cli)
- [KEDA AKS Addon](https://learn.microsoft.com/pt-br/azure/aks/keda-deploy-add-on-cli)

Crie o ACR (Azure Container Registry):

```powershell
az acr create --name $ContainerRegistryName --resource-group $ResourceGroupName --sku basic
```

Conecte o AKS ao ACR :

```powershell
az aks update --name $AKSClusterName --resource-group $ResourceGroupName --attach-acr $ContainerRegistryName
```

Pegue as credenciais para se conectar ao cluster AKS:

```powershell
az aks get-credentials --resource-group $ResourceGroupName --name $AKSClusterName --overwrite-existing
```

Verifique a conexão com o cluster:

```powershell
kubectl cluster-info
```

### 2. Configurando DAPR no AKS via helmcharts

Adicione a referência:

```powershell
helm repo add dapr https://dapr.github.io/helm-charts/   
helm repo update
helm upgrade --install dapr dapr/dapr --namespace dapr-system --create-namespace
helm upgrade --install dapr-dashboard dapr/dapr-dashboard --namespace dapr-system --create-namespace
```

Verifique se os pods estão rodando:

```powershell
kubectl get pods -n dapr-system
```

### 2.1 Configurando o DAPR Dashboard

#### para acessar o DAPR dashboard, execute o seguinte comando

```powershell
dapr dashboard -k
```

**Resultado esperado:**

```powershell
DAPR dashboard found in namespace: dapr-system
DAPR dashboard available at http://localhost:8080
```

### 3. Configurando KEDA no AKS via helmcharts

Adicione a referência:

```powershell
helm repo add kedacore https://kedacore.github.io/charts
helm repo update
helm upgrade --install keda kedacore/keda -n keda-system --create-namespace
helm upgrade --install keda-add-ons-http kedacore/keda-add-ons-http -n keda-system --create-namespace
 
```

Verifique se os pods estão rodando:

```powershell
kubectl get pods -n keda-system
```

### 4. Configurando a camada de transporte para DAPR e KEDA

Neste projeto, temos três diferentes opções para configurar a camada de transporte (escolha uma):

- [Azure Service Bus](https://github.com/diegodocs/dapr-keda-go/setup-infra-azsbus.md)
- [Redis](https://github.com/diegodocs/dapr-keda-go/setup-infra-redis.md)
- [RabbitMq](https://github.com/diegodocs/dapr-keda-go/setup-infra-rbmq.md)

## Deploy das aplicações no AKS

- [veja passo a passo completo](<https://github.com/diegodocs/dapr-keda-go/docs/setup-app.md>)

### 1. Criando as imagens docker

```powershell
az acr login --name $ContainerRegistryName
docker build -t "$ContainerRegistryName.azurecr.io/consumer-app:1.0.0" -f cmd/consumer/dockerfile .
docker build -t "$ContainerRegistryName.azurecr.io/producer-app:1.0.0" -f cmd/producer/dockerfile .
```

### 2. Subindo as imagens no Container Registry ACR

```powershell

docker push "$ContainerRegistryName.azurecr.io/consumer-app:1.0.0" 
docker push "$ContainerRegistryName.azurecr.io/producer-app:1.0.0" 
```

### 3. Aplicando Helmchart das Aplicações

```powershell
helm upgrade --install app .helmcharts/app -n tree --create-namespace
```

Verifique se os pods estão rodando::

```powershell
kubectl get pods -n tree
```

### 4. Testando as aplicações

```powershell
# Revisar Logs
kubectl logs -f -l app=consumer1 --all-containers=true -n tree

# Fazer o forward da porta
kubectl port-forward pod/producer1 8081 8081 -n tree

# Enviar um post para o producer
- POST -> http://localhost:8081/plant
- Json Body: {"numberOfTrees":100}

# Validar os pods e seus status
kubectl get pod -l app=consumer1 -n tree
```

### 4. Removendo recursos

Removendo os helm-charts:

```powershell
helm uninstall app -n tree
helm uninstall keda-add-ons-http -n keda-system
helm uninstall keda -n keda-system
helm uninstall dapr -n dapr-system
```

Excluir todos os recursos Azure:

```powershell
az aks delete --name $AKSClusterName --resource-group $ResourceGroupName
az acr delete --name $ContainerRegistryName --resource-group $ResourceGroupName
az group delete --name $ResourceGroupName
```

## Referências

- [DAPR KEDA GO - Plant Tree App](https://github.com/diegodocs/dapr-keda-go)
- [DAPR - Pros/Cons](https://github.com/diegodocs/dapr-keda-go/docs/dapr-pros-cons.md)
- [KEDA  - Pros/Cons](https://github.com/diegodocs/dapr-keda-go/docs/keda-pros-cons.md)
