GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

BINARY_NAME=pfsense-conftool
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_MAC=$(BINARY_NAME)_macos
BINARY_WIN=$(BINARY_NAME)_win.exe

all: build-linux build-macos build-windows
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GOBUILD) -o bin/$(BINARY_LINUX) -v

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_MAC) -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o bin/$(BINARY_WIN) -v