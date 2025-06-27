# Probe

A tiny Go web server that you can probe with a custom message and port.

## Usage

```bash
# Run locally
go run main.go

# Run with a custom message
go run main.go -message "Hello World!"

# Run with a custom port
go run main.go -port 9090

# Test
curl http://localhost:8080
curl http://localhost:8080/health
```
