# Project specific variables
KUBECTL = kubectl
MINIKUBE = minikube
HELM = helm
MIGRATION_TOOL = migrate
MIGRATION_DIR = migrations
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
deploy-all: deploy-secrets deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2)

deploy-secrets: ## Apply secrets
	$(KUBECTL) apply -f deployments/kubernetes/shared/shared-secret.yaml

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
	$(KUBECTL) port-forward service/$(SERVICE_NAME1) 5000:5000

forward-$(SERVICE_NAME2): ## Forward service-name2 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME2) 9090:9090

## Clean up Kubernetes resources
clean: ## Clean up all Kubernetes resources
	$(KUBECTL) delete -f $(SERVICE_NAME1_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME1_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME2_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME2_SERVICE) --ignore-not-found

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

## Migrations

# Migration base command setup
MIGRATION_DIR_FLIGHT = $(MIGRATION_DIR)/flight-booking
MIGRATION_DIR_HOTEL = $(MIGRATION_DIR)/hotel-booking

# Run a migration for the hotel-booking service with a filename argument
migrate-hotel: ## Run migration for the hotel-booking service
	@echo "Running migration for hotel-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_HOTEL) -database $(DATABASE_URL) up

# Run a migration for the flight-booking service with a filename argument
migrate-flight: ## Run migration for the flight-booking service
	@echo "Running migration for flight-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_FLIGHT) -database $(DATABASE_URL) up

# Revert a migration for the hotel-booking service with a filename argument
migrate-hotel-down: ## Revert migration for the hotel-booking service
	@echo "Reverting migration for hotel-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_HOTEL) -database $(DATABASE_URL) down

# Revert a migration for the flight-booking service with a filename argument
migrate-flight-down: ## Revert migration for the flight-booking service
	@echo "Reverting migration for flight-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_FLIGHT) -database $(DATABASE_URL) down

## Generate a migration file (up and down) for hotel-booking
generate-migration-hotel: ## Generate a migration for hotel-booking
	@echo "Generating migration for hotel-booking"
	$(MIGRATION_TOOL) create -ext sql -dir $(MIGRATION_DIR_HOTEL) -seq $(MIGRATION_NAME)

## Generate a migration file (up and down) for flight-booking
generate-migration-flight: ## Generate a migration for flight-booking
	@echo "Generating migration for flight-booking"
	$(MIGRATION_TOOL) create -ext sql -dir $(MIGRATION_DIR_FLIGHT) -seq $(MIGRATION_NAME)

# The user should provide the name of the migration when running these commands:
# Example: make generate-migration-hotel MIGRATION_NAME=create_hotels_table

.PHONY: all start deploy-all deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2) deploy-helm-all \
    deploy-helm-$(SERVICE_NAME1) deploy-helm-$(SERVICE_NAME2) forward-$(SERVICE_NAME1) \
    forward-$(SERVICE_NAME2) clean clean-helm stop docker-compose-up docker-compose-down \
    migrate-hotel migrate-flight migrate-hotel-down migrate-flight-down generate-migration-hotel generate-migration-flight
