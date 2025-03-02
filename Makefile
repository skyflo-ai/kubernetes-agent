.PHONY: build test lint clean setup-cluster deploy-all clean-cluster build-all

# Build variables
BINARY_NAME=skyflo
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
BLUE=\033[0;34m
NC=\033[0m

# Build all binaries
build: build-watcher build-metrics

build-watcher:
	@echo "$(YELLOW)Building watcher binary...$(NC)"
	CGO_ENABLED=0 go build -o $(GOBIN)/watcher ./cmd/watcher

build-metrics:
	@echo "$(YELLOW)Building metrics binary...$(NC)"
	CGO_ENABLED=0 go build -o $(GOBIN)/metrics ./cmd/metrics

build-server:
	@echo "$(YELLOW)Building test server binary...$(NC)"
	CGO_ENABLED=0 go build -o $(GOBIN)/server ./cmd/testserver

# Run tests
test:
	go test -v -race ./...

# Run linter
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf $(GOBIN)

# Setup Minikube cluster
setup-cluster:
	@echo "$(YELLOW)Checking Minikube cluster status...$(NC)"
	@if ! minikube status > /dev/null 2>&1; then \
		echo "$(YELLOW)Setting up Minikube cluster...$(NC)"; \
		minikube start --nodes 2 --addons=metrics-server; \
	else \
		echo "$(BLUE)Minikube cluster is already running.$(NC)"; \
	fi
	@echo "$(YELLOW)Setting docker env...$(NC)"
	@eval $$(minikube -p minikube docker-env)

# Build Docker images
build-all:
	@echo "$(YELLOW)Building Docker images...$(NC)"
	@eval $$(minikube -p minikube docker-env)
	docker build -f dockerfiles/watcher.Dockerfile -t skyflo-k8s-watcher:latest .
	docker build -f dockerfiles/metrics.Dockerfile -t skyflo-k8s-metrics:latest .
	# docker build -f dockerfiles/testserver.Dockerfile -t skyflo-test-server:latest .

# Create development secrets
create-secrets:
	@echo "$(YELLOW)Creating development secrets...$(NC)"
	kubectl create secret generic skyflo-agent-secret \
		--from-literal=api-key=dev-api-key-123 \
		--dry-run=client -o yaml | kubectl apply -f -

# Deploy all components
deploy-all: create-secrets
	@echo "$(YELLOW)Deploying all components...$(NC)"
	kubectl delete -f manifests/agent.yaml --ignore-not-found
	kubectl apply -f manifests/agent.yaml
	@echo "$(YELLOW)Waiting for pods to be ready...$(NC)"
	kubectl wait --for=condition=ready pod -l app=skyflo-k8s-watcher --timeout=120s
	kubectl wait --for=condition=ready pod -l app=skyflo-k8s-metrics --timeout=120s

# Show status of all components
status:
	@echo "$(YELLOW)Pod Status:$(NC)"
	kubectl get pods -A
	@echo "\n$(YELLOW)Service Status:$(NC)"
	kubectl get svc
	@echo "\n$(YELLOW)Node Status:$(NC)"
	kubectl get nodes

# Clean up cluster resources
clean-cluster:
	@echo "$(YELLOW)Cleaning up cluster resources...$(NC)"
	kubectl delete -f k8s/test-server.yaml --ignore-not-found
	kubectl delete -f k8s/agent.yaml --ignore-not-found
	kubectl delete -f k8s/metrics-daemonset.yaml --ignore-not-found
	kubectl delete secret skyflo-agent-secret --ignore-not-found

# Stop Minikube cluster
stop-cluster:
	@echo "$(YELLOW)Stopping Minikube cluster...$(NC)"
	minikube stop

# All-in-one setup command
all: setup-cluster build-all deploy-all status
	@echo "$(GREEN)Setup complete! The cluster is now running with all components deployed.$(NC)"
	@echo "$(YELLOW)Useful commands:$(NC)"
	@echo "  View agent logs:    $(GREEN)kubectl logs -f deployment/skyflo-k8s-agent$(NC)"
	@echo "  View server logs:   $(GREEN)kubectl logs -f deployment/skyflo-test-server$(NC)"
	@echo "  View metrics logs:  $(GREEN)kubectl logs -f daemonset/skyflo-k8s-metrics$(NC)"
	@echo "  Check status:       $(GREEN)make status$(NC)"
	@echo "  Clean up:          $(GREEN)make clean-cluster$(NC)" 