start:
	go run main.go
start2:
	CGO_ENABLED=1 go build -o server main.go
	./server
test:
	gotestsum --format testname ./tests/
setup:
	@echo "Downloading dependencies..."
	go mod tidy
	go mod download

	@echo "Dependencies downloaded successfully."

	@echo "Downloading gotestsum..."
	go install gotest.tools/gotestsum@latest

	@echo "gotestsum downloaded successfully."