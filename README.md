# Probe

A tiny Go web server that echoes what you start it with - perfect for testing.

## Usage

### Local Development

```bash
# Run with default message
go run main.go

# Run with custom message
go run main.go -message "Probe is running!"

# Run on custom port
go run main.go -port 9090 -message "Custom message"

# Using environment variables
ECHO_MESSAGE="Hello from probe!" go run main.go
```

### Docker

```bash
# Build the image
docker build -t probeit/probe:latest .

# Run with default message
docker run -p 8080:8080 probeit/probe:latest

# Run with custom message
docker run -p 8080:8080 -e ECHO_MESSAGE="Hello from Docker!" probeit/probe:latest

# Run with command line arguments
docker run -p 8080:8080 probeit/probe:latest -message "Hello from args!"
```

### Kubernetes Deployment

#### Basic Deployment

```bash
# Apply all manifests
kubectl apply -f kubernetes/

# Or use kustomize
kubectl apply -k kubernetes/
```

#### ArgoCD Application

Create an ArgoCD application pointing to this repository:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: probe
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/probeit/probe
    targetRevision: HEAD
    path: kubernetes
  destination:
    server: https://kubernetes.default.svc
    namespace: default
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```

## API Endpoints

### `GET /`
Returns a JSON response with the configured message and request details:

```json
{
  "message": "Hello from probe!",
  "method": "GET", 
  "path": "/",
  "headers": {...},
  "host": "localhost:8080",
  "remote_addr": "127.0.0.1:12345"
}
```

### `GET /health`
Health check endpoint:

```json
{
  "status": "healthy"
}
```

## Building

```bash
# Build locally
go build -o probe main.go

# Build Docker image
docker build -t probeit/probe:latest .

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o probe-linux-amd64 main.go
```

## Testing

```bash
# Test locally
curl http://localhost:8080

# Test health endpoint
curl http://localhost:8080/health
```
