# Tracking Order Service

## Dependencies

- minikube
- kustomize
- golang
- dotenv

## Settings

**Dotenv**

Create new file, named `.env` and the contents of the file must be set according to the contents of the `.env.example` file

## Tests

**With Coverage**

```shellscript
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Run

**Start Docker Service**

```shellscript
service docker start
```

**Start Minikube Service**

```shellscript
minikube start
```

**Set Minikube Docker Environment**

```shellscript
eval $(minikube docker-env)
```

**Build Docker Image**

```python
# go to directory (tracking-order-service)
cd /tracking-order-service/
docker build -t tracking-order -f service-tracking-order.Dockerfile .
```

**Run Docker Image With Kubernetes Using Kustomize**

```python
# go to parent directory
kubectl apply -k ./
```

**Check if the cronjob was created successfully**

```shellscript
kubectl get cronjobs
```

**Check if the cronjob is running successfully**

```shellscript
kubectl get jobs
```

## Teardown

```shellscript
kubectl delete -k ./
```

**Set Local Docker Environment**

```shellscript
eval "$(docker-machine env -u)"
```
