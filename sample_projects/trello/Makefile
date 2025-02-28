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
