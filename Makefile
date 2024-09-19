all: build

build:
	@echo "Building....."
	@go build -o gotime main.go

clean:
	@echo "Cleaning....."
	@rm -rvf gotime

.PHONY: all build
