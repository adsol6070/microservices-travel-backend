# Project specific variables
KUBECTL = kubectl
MINIKUBE = minikube
MIGRATION_TOOL = bin/migrate
MIGRATION_DIR = migrations
SERVICE_NAME1 = hotel-booking
SERVICE_NAME2 = flight-booking
SERVICE_NAME3 = user-service
# DATABASE_URL = "postgres://postgres:royal-dusk-20@travel-db.cd2uyuqoiqtz.ap-south-1.rds.amazonaws.com:5432"
DATABASE_URL = "postgres://postgres:admin@localhost:5432/traveltest"

# Kubernetes manifests paths
SERVICE_NAME1_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME1)/deployment.yaml
SERVICE_NAME1_SERVICE = deployments/kubernetes/$(SERVICE_NAME1)/service.yaml
SERVICE_NAME2_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME2)/deployment.yaml
SERVICE_NAME2_SERVICE = deployments/kubernetes/$(SERVICE_NAME2)/service.yaml
SERVICE_NAME3_DEPLOYMENT = deployments/kubernetes/$(SERVICE_NAME3)/deployment.yaml
SERVICE_NAME3_SERVICE = deployments/kubernetes/$(SERVICE_NAME3)/service.yaml

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
deploy-all: deploy-secrets deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2) deploy-$(SERVICE_NAME3)

deploy-secrets: ## Apply secrets
	$(KUBECTL) apply -f deployments/kubernetes/shared/shared-secret.yaml

deploy-$(SERVICE_NAME1): ## Deploy service-name1
	$(KUBECTL) apply -f $(SERVICE_NAME1_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME1_SERVICE)

deploy-$(SERVICE_NAME2): ## Deploy service-name2
	$(KUBECTL) apply -f $(SERVICE_NAME2_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME2_SERVICE)

deploy-$(SERVICE_NAME3): ## Deploy service-name2
	$(KUBECTL) apply -f $(SERVICE_NAME3_DEPLOYMENT)
	$(KUBECTL) apply -f $(SERVICE_NAME3_SERVICE)

## Expose services using kubectl port-forward
forward-$(SERVICE_NAME1): ## Forward service-name1 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME1) 5100:5100

forward-$(SERVICE_NAME2): ## Forward service-name2 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME2) 6100:6100

forward-$(SERVICE_NAME3): ## Forward service-name2 port
	$(KUBECTL) port-forward service/$(SERVICE_NAME3) 7100:7100

## Clean up Kubernetes resources
clean: ## Clean up all Kubernetes resources
	$(KUBECTL) delete -f $(SERVICE_NAME1_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME1_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME2_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME2_SERVICE) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME3_DEPLOYMENT) --ignore-not-found
	$(KUBECTL) delete -f $(SERVICE_NAME3_SERVICE) --ignore-not-found

## Stop the local Kubernetes cluster
stop: ## Stop the Kubernetes cluster (Minikube or Kind)
	$(MINIKUBE) stop

## Run Docker Compose (for local testing)
docker-compose-up: ## Start local services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_OVERRIDE) up -d

docker-compose-down: ## Stop services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_OVERRIDE) down	

## Run Docker Compose for production
docker-compose-prod-up: ## Start production services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_PROD) up -d

docker-compose-prod-down: ## Stop production services using Docker Compose
	docker-compose -f $(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_PROD) down

## Migrations

# Migration base command setup
MIGRATION_DIR_FLIGHT = $(MIGRATION_DIR)/flight-booking
MIGRATION_DIR_HOTEL = $(MIGRATION_DIR)/hotel-booking
MIGRATION_DIR_USER = $(MIGRATION_DIR)/user-service

# Run a migration for the hotel-booking service with a filename argument
migrate-hotel: ## Run migration for the hotel-booking service
	@echo "Running migration for hotel-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_HOTEL) -database $(DATABASE_URL) up

# Run a migration for the flight-booking service with a filename argument
migrate-flight: ## Run migration for flight-booking
	@echo "Running migration for flight-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_FLIGHT) -database $(DATABASE_URL) up

# Run a migration for the user service with a filename argument
migrate-user: ## Run migration for the user service
	@echo "Running migration for user-service"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_USER) -database $(DATABASE_URL) up

# Revert a migration for the hotel-booking service with a filename argument
migrate-hotel-down: ## Revert migration for the hotel-booking service
	@echo "Reverting migration for hotel-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_HOTEL) -database $(DATABASE_URL) down

# Revert a migration for the flight-booking service with a filename argument
migrate-flight-down: ## Revert migration for flight-booking
	@echo "Reverting migration for flight-booking"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_FLIGHT) -database $(DATABASE_URL) down

# Revert a migration for the user service with a filename argument
migrate-user-down: ## Revert migration for the user service
	@echo "Reverting migration for user-service"
	$(MIGRATION_TOOL) -path=$(MIGRATION_DIR_USER) -database $(DATABASE_URL) down

## Generate a migration file (up and down) for hotel-booking
generate-migration-hotel: ## Generate a migration for hotel-booking
	@read -p "Enter migration name: " MIGRATION_NAME; \
	echo "Generating migration for hotel-booking"; \
	mkdir -p $(MIGRATION_DIR_HOTEL); \
	$(MIGRATION_TOOL) create -ext sql -dir $(MIGRATION_DIR_HOTEL) -seq $$MIGRATION_NAME

## Generate a migration file (up and down) for flight-booking
generate-migration-flight: ## Generate a migration for flight-booking
	@read -p "Enter migration name: " MIGRATION_NAME; \
	echo "Generating migration for flight-booking"; \
	mkdir -p $(MIGRATION_DIR_FLIGHT); \
	$(MIGRATION_TOOL) create -ext sql -dir $(MIGRATION_DIR_FLIGHT) -seq $$MIGRATION_NAME

## Generate a migration file (up and down) for flight-booking
generate-migration-user: ## Generate a migration for flight-booking
	@read -p "Enter migration name: " MIGRATION_NAME; \
	echo "Generating migration for user-service"; \
	mkdir -p $(MIGRATION_DIR_USER); \
	$(MIGRATION_TOOL) create -ext sql -dir $(MIGRATION_DIR_USER) -seq $$MIGRATION_NAME

# The user should provide the name of the migration when running these commands:
# Example: make generate-migration-hotel MIGRATION_NAME=create_hotels_table

.PHONY: all start deploy-all deploy-$(SERVICE_NAME1) deploy-$(SERVICE_NAME2) \
    forward-$(SERVICE_NAME1) forward-$(SERVICE_NAME2) clean stop docker-compose-up docker-compose-down \
    docker-compose-prod-up docker-compose-prod-down migrate-hotel migrate-flight migrate-hotel-down migrate-flight-down \
    generate-migration-hotel generate-migration-flight
