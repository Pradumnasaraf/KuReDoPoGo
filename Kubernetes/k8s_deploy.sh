#!/bin/bash

# Set an error flag
set -e

# Deploy API components
echo "Creating namespace for API and deploying API components..."
kubectl apply -f Kubernetes/API/Api.Namespace.yml
kubectl apply -f Kubernetes/API/Api.ConfigMap.yml
kubectl apply -f Kubernetes/API/Api.Deployment.yml
# kubectl apply -f Kubernetes/API/Api.Ingress.yml # Uncomment this line if you are using Ingress
kubectl apply -f Kubernetes/API/Api.Secret.yml
kubectl apply -f Kubernetes/API/Api.Service.yml

# Deploy Postgres
echo "Creating namespace for Mongo and deploying Mongo components..."
kubectl apply -f Kubernetes/Postgres/Postgres.Namespace.yml
kubectl apply -f Kubernetes/Postgres/Postgres.Secret.yml
kubectl apply -f Kubernetes/Postgres/Postgres.HeadlessService.yml
kubectl apply -f Kubernetes/Postgres/Postgres.Statefulset.yml

# Deploy Redis
echo "Creating namespace for Redis and deploying Redis components..."
kubectl apply -f Kubernetes/Redis/Redis.Namespace.yml
kubectl apply -f Kubernetes/Redis/Redis.HeadlessService.yml
kubectl apply -f Kubernetes/Redis/Redis.Statefulset.yml

echo "All components deployed successfully!"
