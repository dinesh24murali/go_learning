.PHONY: build run clean docker-build docker-run

# Build the Go application
build:
	go build -o gin-server

# Run the Go application
run:
	find . -name "*.go" | entr -r go run .

# Clean up compiled files
clean:
	rm -f gin-server

# Build Docker image
docker-build:
	docker build -t go-gin-app .

# Run Docker container
docker-run:
	docker run -p 8080:8080 go-gin-app
