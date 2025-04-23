# Project specific variables
KUBECTL = kubectl
MINIKUBE = minikube
MIGRATION_TOOL = bin/migrate
MIGRATION_DIR = migrations
SERVICE_NAME1 = hotel-booking
SERVICE_NAME2 = flight-booking
SERVICE_NAME3 = user-service
SERVICE_NAME4 = blog-service
SERVICE_NAME5 = invoice-service
# DATABASE_URL = "postgres://postgres:royal-dusk-20@travel-db.cd2uyuqoiqtz.ap-south-1.rds.amazonaws.com:5432"
DATABASE_URL = "postgres://devuser:devpassword@localhost:5432/devdb?sslmode=disable"

# Kubernetes manifests paths
SERVICE_NAME1_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME1)/deployment.yaml
SERVICE_NAME1_SERVICE = deployments/kubernetes/$(SERVICE_NAME1)/service.yaml
SERVICE_NAME2_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME2)/deployment.yaml
SERVICE_NAME2_SERVICE = deployments/kubernetes/$(SERVICE_NAME2)/service.yaml
SERVICE_NAME3_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME3)/deployment.yaml
SERVICE_NAME3_SERVICE = deployments/kubernetes/$(SERVICE_NAME3)/service.yaml
SERVICE_NAME4_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME4)/deployment.yaml
SERVICE_NAME4_SERVICE = deployments/kubernetes/$(SERVICE_NAME4)/service.yaml
SERVICE_NAME5_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME5)/deployment.yaml
SERVICE_NAME5_SERVICE = deployments/kubernetes/$(SERVICE_NAME5)/service.yaml

# Docker Compose path for local and production development
DOCKER_COMPOSE = deployments/docker-compose.yaml
DOCKER_COMPOSE_OVERRIDE = deployments/docker-compose.override.yaml
DOCKER_COMPOSE_PROD = deployments/docker-compose.prod.yaml

# Exporting bin folder to the path for Makefile
export PATH   := $(PWD)/bin:$(PATH)
export SHELL  := bash
export OSTYPE := $(shell uname -s | tr A-Z a-z)
export ARCH := $(shell uname -m)

# --- Tools & Variables ---
include ./misc/make/tools.Makefile

deps: $(MIGRATE) $(AIR)

# Default target
all: start

## Local Kubernetes Cluster Setup (using Minikube or Kind)
start: ## Start Kubernetes Cluster (Minikube or Kind)
	$(MINIKUBE) start

## Deploy all services using kubectl
deploy-all: deploy-secrets deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2) deploy-$(SERVICE_NAME3) deploy-$(SERVICE_NAME4) deploy-$(SERVICE_NAME5)

deploy-secrets: ## Apply secrets
	$(KUBECTL) apply -f deployments/kubernetes/shared/shared-secret.yaml

deploy-$(SERVICE_NAME1): ## Deploy service-name1
	$(KUBECTL) apply -f $(SERVICE_NAME1_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME1_SERVICE)

deploy-$(SERVICE_NAME2): ## Deploy service-name2
	$(KUBECTL) apply -f $(SERVICE_NAME2_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME2_SERVICE)

deploy-$(SERVICE_NAME3): ## Deploy service-name3
	$(KUBECTL) apply -f $(SERVICE_NAME3_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME3_SERVICE)

deploy-$(SERVICE_NAME4): ## Deploy service-name4
	$(KUBECTL) apply -f $(SERVICE_NAME4_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME4_SERVICE)

deploy-$(SERVICE_NAME5): ## Deploy service-name5
	$(KUBECTL) apply -f $(SERVICE_NAME5_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME5_SERVICE)

## Expose services using kubectl port-forward
forward-$(SERVICE_NAME1): ## Forward service-name1 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME1) 5100:5100

forward-$(SERVICE_NAME2): ## Forward service-name2 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME2) 6100:6100

forward-$(SERVICE_NAME3): ## Forward service-name3 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME3) 7100:7100


forward-$(SERVICE_NAME4): ## Forward service-name4 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME4) 7200:7200

forward-$(SERVICE_NAME5): ## Forward service-name5 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME5) 8200:8200

## Clean up Kubernetes resources
clean: ## Clean up all Kubernetes resources
	$(KUBECTL) delete -f $(SERVICE_NAME1_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME1_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME2_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME2_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME3_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME3_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME4_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME4_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME5_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME5_SERVICE) --ignore-not-found

## Stop the local Kubernetes cluster
stop: ## Stop the Kubernetes cluster (Minikube or Kind)
	$(MINIKUBE) stop

## Run Docker Compose (for local testing)
docker-compose-up: ## Start local services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_OVERRIDE) up --build -d

docker-compose-down: ## Stop services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_OVERRIDE) down	

## Run Docker Compose for production
docker-compose-prod-up: ## Start production services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_PROD) up -d

docker-compose-prod-down: ## Stop production services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_PROD) down

## Run a migration for all services
run-migrate: ## Run all migrations
	@echo "Running all migrations"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR) -database $(DATABASE_URL) up

## Revert a migration for all services
migrate-down: ## Revert all migrations
	@echo "Reverting all migrations"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR) -database $(DATABASE_URL) down

## Generate a migration file (up and down)
generate-migration: ## Generate a migration (usage: make generate-migration)
	@read -p "Enter migration name: " MIGRATION_NAME; \
	echo "Generating migration for $$MIGRATION_NAME in $(MIGRATION_DIR)"; \
	$(MIGRATION_TOOL) create -ext sql -dir $(MIGRATION_DIR) -seq $$MIGRATION_NAME

# The user should provide the name of the migration when running these commands:
# Example: make generate-migration-hotel MIGRATION_NAME=create_hotels_table

.PHONY: all start deploy-all deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2) deploy-$(SERVICE_NAME3) deploy-$(SERVICE_NAME4) deploy-$(SERVICE_NAME5) \
    forward-$(SERVICE_NAME1) forward-$(SERVICE_NAME2) forward-$(SERVICE_NAME3) forward-$(SERVICE_NAME4) forward-$(SERVICE_NAME5) clean stop docker-compose-up docker-compose-down \
    docker-compose-prod-up docker-compose-prod-down migrate migrate-down \
    generate-migration
