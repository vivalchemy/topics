DOCKER_COMPOSE := docker compose

# Default target
.DEFAULT_GOAL := help

# 🛠️ Utility
help: ## Show this help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?##' Makefile | awk 'BEGIN {FS = ":.*?##"} {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# 🧑‍💻 Development
dev-up: ## Start development stack
	$(DOCKER_COMPOSE) --profile development up -d

dev-up-build: ## Rebuild and start development stack
	$(DOCKER_COMPOSE) --profile development up --build -d

dev-down: ## Stop development stack
	$(DOCKER_COMPOSE) --profile development down

dev-logs: ## Tail logs for development stack
	$(DOCKER_COMPOSE) --profile development logs -f

dev-restart: ## Restart development stack
	$(DOCKER_COMPOSE) --profile development down
	$(DOCKER_COMPOSE) --profile development up -d

# 🚀 Production
prod-up: ## Start production stack
	$(DOCKER_COMPOSE) --profile production up -d

prod-up-build: ## Rebuild and start production stack
	$(DOCKER_COMPOSE) --profile production up --build -d

prod-down: ## Stop production stack
	$(DOCKER_COMPOSE) --profile production down

prod-logs: ## Tail logs for production stack
	$(DOCKER_COMPOSE) --profile production logs -f

prod-restart: ## Restart production stack
	$(DOCKER_COMPOSE) --profile production down
	$(DOCKER_COMPOSE) --profile production up -d

# 🔍 Debugging
ps: ## Show running containers
	$(DOCKER_COMPOSE) ps

logs: ## Tail logs for all services (any profile)
	$(DOCKER_COMPOSE) logs -f

build: ## Build all images (no containers started)
	$(DOCKER_COMPOSE) build

clean: ## Remove all containers, networks, and volumes
	$(DOCKER_COMPOSE) down -v --remove-orphans

