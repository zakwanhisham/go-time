all: build

build:
	@echo "Building....."
	@go build -o gotime main.go

run:
	@echo "Running....."
	@go run main.go

clean:
	@echo "Cleaning....."
	@rm -rvf gotime

.PHONY: all build run
