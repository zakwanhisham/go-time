all: build

build:
	@echo "Building....."
	@go build -o gotime main.go

install: build
	@echo "Installing....."
	@cp -iv gotime ~/.local/bin

clean:
	@echo "Cleaning....."
	@rm -rvf gotime

.PHONY: all build install clean
