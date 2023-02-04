build:
	mkdir -p build
	CGO_ENABLED=0 go build -o build/devops-security ./cmd/devops-security