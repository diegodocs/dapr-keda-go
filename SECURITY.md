# Security Policy

## Security Improvements

This document outlines the security improvements implemented in this project.

### Application Security

#### Producer Service (`cmd/producer/main.go`)

1. **HTTP Server Hardening**
   - Added read/write/idle timeouts to prevent slowloris attacks
   - Configured `ReadHeaderTimeout` (5s) to prevent slow header attacks
   - Set `MaxHeaderBytes` (1MB) to limit header size
   - Implemented proper graceful shutdown handling

2. **Request Validation**
   - HTTP method validation (only POST allowed)
   - Request body size limiting (1MB max)
   - Input validation for `numberOfTrees` (positive and <= 10,000)
   - Proper error handling with appropriate HTTP status codes

3. **Context Management**
   - Request-level timeouts (30s)
   - Publish-level timeouts (5s per event)
   - Context cancellation handling
   - Proper resource cleanup with defer statements

4. **Error Handling**
   - All errors are properly checked and logged
   - No silent failures
   - Descriptive error messages without exposing internals
   - Proper HTTP status codes returned

#### Consumer Service (`cmd/consumer/main.go`)

1. **HTTP Server Hardening**
   - Added read/write/idle timeouts
   - Configured `ReadHeaderTimeout` and `MaxHeaderBytes`
   - Proper error handling on server startup

2. **Request Handling**
   - HTTP method validation (only POST allowed)
   - Request body size limiting (10MB max)
   - Proper error handling and logging
   - Resource cleanup with defer

### Container Security

#### Dockerfile Improvements (both services)

1. **Multi-stage Builds**
   - Separate builder and runtime stages
   - Minimal runtime image (Alpine Linux)

2. **Security Updates**
   - Automatic security updates via `apk upgrade`
   - CA certificates included for TLS

3. **Non-root User**
   - Created dedicated user `appuser` (UID 1000)
   - Application runs as non-root
   - Proper file ownership and permissions

4. **Build Optimizations**
   - Binary stripping with `-ldflags="-s -w"`
   - Go module verification with `go mod verify`
   - Minimal attack surface

5. **Dependencies Verification**
   - `go mod download && go mod verify` ensures integrity

### Build Security

#### .dockerignore

- Excludes sensitive files and unnecessary artifacts
- Reduces image size and attack surface
- Prevents accidental inclusion of secrets

### Dependency Management

1. **Up-to-date Dependencies**
   - All dependencies updated to latest versions
   - Go toolchain upgraded to 1.26.0
   - Regular dependency updates recommended

2. **Vulnerability Scanning**
   - Use `govulncheck` for vulnerability scanning
   - Use `gosec` for static security analysis

## Security Best Practices

### For Deployment

1. **Network Security**
   - Use network policies to restrict traffic
   - Enable TLS for all external communications
   - Use service mesh for mTLS between services

2. **Secrets Management**
   - Never commit secrets to version control
   - Use Kubernetes secrets or external secret managers
   - Rotate credentials regularly

3. **Resource Limits**
   - Set CPU and memory limits in Kubernetes
   - Configure proper pod security policies
   - Use resource quotas

4. **Monitoring**
   - Enable audit logging
   - Monitor for suspicious activity
   - Set up alerts for security events

### Running Security Scans

```bash
# Check for vulnerabilities in dependencies
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# Static security analysis
go install github.com/securego/gosec/v2/cmd/gosec@latest
gosec ./...

# Update dependencies
go get -u ./...
go mod tidy
```

## Reporting Security Issues

If you discover a security vulnerability, please email security@example.com instead of using the issue tracker.

## Security Checklist

- [x] HTTP timeouts configured
- [x] Request size limits enforced
- [x] Input validation implemented
- [x] Error handling comprehensive
- [x] Non-root container user
- [x] Security updates in containers
- [x] Dependencies up to date
- [x] Context timeouts for operations
- [x] Method validation on endpoints
- [x] Proper resource cleanup
- [x] Binary stripping in builds
- [x] .dockerignore configured

## Future Improvements

1. Implement rate limiting
2. Add authentication/authorization
3. Implement request signing
4. Add structured logging with correlation IDs
5. Implement circuit breakers
6. Add distributed tracing
7. Implement health checks and readiness probes
8. Add metrics and monitoring
