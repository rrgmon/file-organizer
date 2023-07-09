# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build

# Output directory
OUTDIR = bin

# Binary names
BINARY = sortzilla

# Build targets
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTDIR)/linux/$(BINARY)-amd64

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(OUTDIR)/linux/$(BINARY)-arm64

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUTDIR)/macos/$(BINARY)-amd64

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(OUTDIR)/macos/$(BINARY)-arm64

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(OUTDIR)/windows/$(BINARY)-amd64.exe

# Default target
build: build-linux-amd64 build-linux-arm64 build-macos-amd64 build-macos-arm64 build-windows-amd64

# Clean target
clean:
	rm -rf $(OUTDIR)

