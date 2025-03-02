#!/bin/bash

# Exit on any error
set -e

# Function to wait for pod readiness
wait_for_pod() {
    local label="$1"
    local timeout=120
    local start_time=$(date +%s)

    echo -e "${YELLOW}Waiting for pod with label $label to be ready...${NC}"
    
    while true; do
        if [ $(($(date +%s) - start_time)) -gt $timeout ]; then
            echo -e "${RED}Timeout waiting for pod${NC}"
            return 1
        fi

        local status=$(kubectl get pods -l app=$label -o jsonpath='{.items[0].status.phase}' 2>/dev/null)
        if [ "$status" = "Running" ]; then
            echo -e "${GREEN}Pod is ready!${NC}"
            return 0
        fi
        sleep 2
    done
}

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${YELLOW}Starting Minikube setup...${NC}"

# Start Minikube if it's not running
if ! minikube status > /dev/null 2>&1; then
    echo -e "${YELLOW}Starting Minikube cluster...${NC}"
    minikube start --addons=metrics-server
else
    echo -e "${GREEN}Minikube is already running${NC}"
fi

# Point shell to minikube's docker-env
eval $(minikube -p minikube docker-env)

echo -e "${YELLOW}Building test server Docker image...${NC}"
docker build -t skyflo-test-server:latest ./test-server

echo -e "${YELLOW}Deploying test server...${NC}"
kubectl apply -f k8s/test-server.yaml
wait_for_pod "skyflo-test-server"

echo -e "${YELLOW}Building Docker image...${NC}"
docker build -t skyflo-k8s-agent:latest .

echo -e "${YELLOW}Creating API key secret...${NC}"
# Create a dummy API key for development
kubectl create secret generic skyflo-agent-secret \
    --from-literal=api-key=dev-api-key-123 \
    --dry-run=client -o yaml | kubectl apply -f -

echo -e "${YELLOW}Applying Kubernetes manifests...${NC}"
kubectl apply -f k8s/agent.yaml

wait_for_pod "skyflo-k8s-agent"

echo -e "${GREEN}Setup complete! The agent is now running in your Minikube cluster.${NC}"
echo -e "${YELLOW}Useful commands:${NC}"
echo -e "  View agent logs:    ${GREEN}kubectl logs -f deployment/skyflo-k8s-agent${NC}"
echo -e "  View server logs:   ${GREEN}kubectl logs -f deployment/skyflo-test-server${NC}"
echo -e "  Clean up:          ${GREEN}kubectl delete -f k8s/agent.yaml -f k8s/test-server.yaml${NC}"

# Show pod status
echo -e "\n${YELLOW}Current pod status:${NC}"
kubectl get pods -l "app in (skyflo-k8s-agent,skyflo-test-server)"

# Show service status
echo -e "\n${YELLOW}Service status:${NC}"
kubectl get svc skyflo-test-server

# Show recent logs from test server
echo -e "\n${YELLOW}Recent test server logs:${NC}"
kubectl logs -l app=skyflo-test-server --tail=20 