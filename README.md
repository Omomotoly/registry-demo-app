# Docker Registry Demo: Go Hello World Application

## Overview

This demo demonstrates the complete container image distribution lifecycle by building a simple Go web application, containerizing it with Docker, and pushing it to multiple container registries (Docker Hub and GitHub Container Registry).

## Demo Objectives

- Build a simple "Hello World" web application in Go
- Create a Dockerfile to containerize the application
- Build and test the Docker image locally
- Tag the image for multiple registries
- Push the image to Docker Hub
- Push the image to GitHub Container Registry (GHCR)
- Pull and run the image from both registries

## Prerequisites

- Basic understanding of command line operations
- Docker installed and running
- Docker Hub account
- GitHub account with personal access token (PAT) for GHCR
- Go installed (version 1.19 or higher)

## Project Structure

```
registry-demo-app/
├── main.go              ✓ Go application
├── Dockerfile           ✓ Multi-stage container build
├── go.mod               ✓ Go module definition
├── .dockerignore        ✓ Build optimization
├── .gitignore           ✓ Version control
├── demo-commands.sh     ✓ Command reference
└── README.md            ✓ Complete documentation
```

## Application Details

The Go application is a simple HTTP server that:
- Listens on port 2222
- Responds with `Hello, World! This is a containerized Go app used for DockerHub and GHCR Registry demo.` on the root path `/`.
- Displays the hostname (container ID) for demonstration purposes.
- Lightweight and perfect for demonstrating containerization and image distribution lifecycle.

## Demo Steps

### Step 1: Create the Go Application
Create `main.go` with a simple HTTP server that responds to requests on port `2222`.

### Step 2: Initialize Go Module
```bash
# Initialize the Go module to manage dependencies:
go mod init github.com/fykio/registry-demo-app
```

### Step 3: Test Locally (Optional)
```bash
# Run the application locally to verify it works:
go run main.go

# Visit the following URL in your browser:
- http://localhost:2222
- http://localhost:2222/health
```

### Step 4: Create the Dockerfile
Create a multi-stage Dockerfile that:
- Uses Go official image for building
- Compiles the application as a static binary
- Uses a minimal base image (alpine) for the final container
- Results in a small, secure production image

### Step 5: Build the Docker Image
```bash
# Build the image locally:
docker build -t registry-demo-app:latest .
```

### Step 6: Test the Container Locally
```bash
# Run the container to verify it works:
docker run -p 2222:2222 registry-demo-app:latest

# Visit these to verify it works:
# - http://localhost:2222
# - http://localhost:2222/health
```

### Step 7: Tag for Docker Hub
```bash
# Tag the image with your Docker Hub username:
docker tag registry-demo-app:latest fykio/registry-demo-app:latest
docker tag registry-demo-app:latest fykio/registry-demo-app:v1.0.0
```

### Step 8: Login to Docker Hub
```bash
# Authenticate with Docker Hub:
docker login

# Enter your Docker Hub username and Personal Access Token (PAT)
#...
```

### Step 9: Push to Docker Hub
```bash
# Push both tags to Docker Hub:
docker push fykio/registry-demo-app:latest

docker push fykio/registry-demo-app:v1.0.0
```

### Step 10: Tag for GitHub Container Registry
```bash
# Tag the image for GHCR:
docker tag registry-demo-app:latest ghcr.io/fykio/registry-demo-app:latest
docker tag registry-demo-app:latest ghcr.io/fykio/registry-demo-app:v1.0.0
```

### Step 11: Login to GitHub Container Registry
```bash
# Authenticate with GHCR using a personal access token interactively:
docker login ghcr.io
```

### Step 12: Push to GHCR
```bash
# Push both tags to GitHub Container Registry:
docker push ghcr.io/fykio/registry-demo-app:latest

docker push ghcr.io/fykio/registry-demo-app:v1.0.0
```

### Step 13: Verify Images in Registries
**Docker Hub:**
- Visit https://hub.docker.com/r/fykio/registry-demo-app
- Verify both tags are visible

**GitHub Container Registry:**
- Visit https://github.com/fykio?tab=packages
- Verify the package appears with both tags

### Step 14: Pull and Run from Docker Hub
```bash
# Clean local images and pull from Docker Hub:
docker rmi registry-demo-app:latest

docker rmi fykio/registry-demo-app:latest

docker pull fykio/registry-demo-app:latest

docker run -p 2222:2222 fykio/registry-demo-app:latest
```

### Step 15: Pull and Run from GHCR
```bash
# Pull from GitHub Container Registry:
docker pull ghcr.io/fykio/registry-demo-app:latest

docker run -p 2222:2222 ghcr.io/fykio/registry-demo-app:latest
```

## Key Commands Reference

### Image Building
```bash
docker build -t <image-name>:<tag> .
```

### Image Tagging
```bash
docker tag <source-image>:<tag> <target-image>:<tag>
```

### Registry Authentication
```bash
# Docker Hub
docker login

# GitHub Container Registry
docker login ghcr.io
```

### Image Pushing
```bash
docker push <registry>/<username>/<image>:<tag>
```

### Image Pulling
```bash
docker pull <registry>/<username>/<image>:<tag>
```

### Running Container
```bash
docker run -p <host-port>:<container-port> <image>:<tag>
```

## Expected Outcomes

After completing this demo, you will have:
- ✅ A working Go web application
- ✅ A multi-stage Dockerfile for optimal image size
- ✅ A locally built and tested Docker image
- ✅ The image pushed to Docker Hub with version tags
- ✅ The same image pushed to GitHub Container Registry
- ✅ Successfully pulled and run containers from both registries

## Image Size Comparison

The multi-stage build approach typically produces:
- **Builder stage:** ~800MB (Go compiler and tools)
- **Final image:** ~15-20MB (Alpine + compiled binary)

This demonstrates the efficiency of multi-stage builds for production deployments.

## Troubleshooting

### Authentication Issues

- Ensure you're using the correct username and PAT for Docker Hub and GHCR
- For GHCR, verify your PAT has the `write:packages` scope
- Check that you're logged into the correct registry

### Build Failures

- Verify Go is installed correctly
- Check that all files are in the correct directory
- Ensure Dockerfile syntax is correct

### Push Failures

- Confirm you're authenticated to the registry
- Verify the image is tagged correctly with your username
- Check your network connection

### Pull Failures

- Ensure the image exists in the registry
- Verify you have access to pull the image (especially for private repos)
- Check that the tag exists

## Best Practices Demonstrated

1. **Multi-stage builds** - Reduces final image size
2. **Semantic versioning** - Both `latest` and version-specific tags
3. **Multi-registry distribution** - Redundancy and flexibility
4. **Minimal base images** - Using Alpine Linux for security and size
5. **Port documentation** - Clearly exposing application ports
6. **Testing before pushing** - Always verify locally first

## Next Steps

After completing this demo, consider learning these:
- Automating the build and push process with GitHub Actions or GitLab CI
- Implementing image scanning and security checks
- Setting up private repositories for sensitive applications
- Exploring Docker Compose for multi-container applications
- Implementing health checks in your Dockerfile
- Adding metadata labels to your images

## Resources

- [Docker Hub Documentation](https://docs.docker.com/docker-hub/)
- [GitHub Container Registry Documentation](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
- [Docker Multi-stage Builds](https://docs.docker.com/build/building/multi-stage/)
- [Go Docker Best Practices](https://docs.docker.com/language/golang/)

## Conclusion

This demo provides hands-on experience with the complete container image distribution lifecycle, from local development to multi-registry deployment. It showcases how Docker enables consistent application packaging and distribution across different registry platforms.