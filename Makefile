# Project specific variables
KUBECTL = kubectl
MINIKUBE = minikube
HELM = helm

# Service names
SERVICE_NAME1 = hotel-booking
SERVICE_NAME2 = flight-booking

# Kubernetes manifests paths
SERVICE_NAME1_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME1)/deployment.yaml
SERVICE_NAME1_SERVICE = deployments/kubernetes/$(SERVICE_NAME1)/service.yaml
SERVICE_NAME2_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME2)/deployment.yaml
SERVICE_NAME2_SERVICE = deployments/kubernetes/$(SERVICE_NAME2)/service.yaml

# Helm chart paths
SERVICE_NAME1_HELM = deployments/helm/$(SERVICE_NAME1)
SERVICE_NAME2_HELM = deployments/helm/$(SERVICE_NAME2)

# Docker Compose path for local development
DOCKER_COMPOSE = deployments/docker-compose.yaml

# Default target
all: start

## Local Kubernetes Cluster Setup (using Minikube or Kind)
start: ## Start Kubernetes Cluster (Minikube or Kind)
	$(MINIKUBE) start

## Deploy all services using kubectl
deploy-all: deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2)

deploy-$(SERVICE_NAME1): ## Deploy service-name1
	$(KUBECTL) apply -f $(SERVICE_NAME1_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME1_SERVICE)

deploy-$(SERVICE_NAME2): ## Deploy service-name2
	$(KUBECTL) apply -f $(SERVICE_NAME2_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME2_SERVICE)

## Deploy all services using Helm (Optional)
deploy-helm-all: deploy-helm-$(SERVICE_NAME1) deploy-helm-$(SERVICE_NAME2)

deploy-helm-$(SERVICE_NAME1): ## Deploy service-name1 using Helm
	$(HELM) install $(SERVICE_NAME1) $(SERVICE_NAME1_HELM)

deploy-helm-$(SERVICE_NAME2): ## Deploy service-name2 using Helm
	$(HELM) install $(SERVICE_NAME2) $(SERVICE_NAME2_HELM)

## Expose services using kubectl port-forward
forward-$(SERVICE_NAME1): ## Forward service-name1 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME1) 8080:8080

forward-$(SERVICE_NAME2): ## Forward service-name2 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME2) 9090:9090

## Clean up Kubernetes resources
clean: ## Clean up all Kubernetes resources
	$(KUBECTL) delete -f $(SERVICE_NAME1_DEPLOYMENT)
	$(KUBECTL) delete -f $(SERVICE_NAME1_SERVICE)
	$(KUBECTL) delete -f $(SERVICE_NAME2_DEPLOYMENT)
	$(KUBECTL) delete -f $(SERVICE_NAME2_SERVICE)

## Clean up Helm releases (Optional)
clean-helm: ## Clean up Helm releases
	$(HELM) uninstall $(SERVICE_NAME1)
	$(HELM) uninstall $(SERVICE_NAME2)

## Stop the local Kubernetes cluster
stop: ## Stop the Kubernetes cluster (Minikube or Kind)
	$(MINIKUBE) stop

## Run Docker Compose (for local testing)
docker-compose-up: ## Start local services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) up -d

docker-compose-down: ## Stop services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) down

.PHONY: all start deploy-all deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2) deploy-helm-all \
    deploy-helm-$(SERVICE_NAME1) deploy-helm-$(SERVICE_NAME2) forward-$(SERVICE_NAME1) \
    forward-$(SERVICE_NAME2) clean clean-helm stop docker-compose-up docker-compose-down
