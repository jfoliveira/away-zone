DOCKER_BUILD_CONTEXT=.
DOCKERFILE_NAME=Dockerfile
DOCKERFILE_TARGET=dev
DOCKER_COMPOSE_PROJECT_NAME=away-zone

build: ## TO DO: Build image for production release

list-containers: ## List running containers belonging to the docker compose project
	$(info Listing containers started by docker compose ...)
	@export DOCKERFILE_TARGET=$(DOCKERFILE_TARGET) && \
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) ps

run-dev: build
run-dev: ## Run in development mode. DO NOT use this in production environment!
	$(info Starting development environment...)
	@export DOCKERFILE_TARGET=$(DOCKERFILE_TARGET) && \
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) up -d

stop-dev: ## Stop running dev containers. Set `DELETE_IMAGES` to any value, e.g. DELETE_IMAGES=y, will delete image for current target
	$(info Stopping runnning containers in development environment...)
	@export DOCKERFILE_TARGET=$(DOCKERFILE_TARGET) && \
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) down
	$(if $(DELETE_IMAGES), docker rmi `docker images -q "$(DOCKER_COMPOSE_PROJECT_NAME)*:$(DOCKERFILE_TARGET)" | uniq`)

help: ## Show this help.
# `help' function obtained from GitHub gist: https://gist.github.com/prwhite/8168133?permalink_comment_id=4160123#gistcomment-4160123
	@echo Simple Health Check System
	@echo
	@awk 'BEGIN {FS = ": .*##"; \
	printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+(\\:[$$()% 0-9a-zA-Z_-]+)*:.*?##/ { gsub(/\\:/,":", $$1); \
	printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2 } /^##@/ { \
	printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL=help
