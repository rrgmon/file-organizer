# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build

# Output directory
OUTDIR = bin

# Binary name
BINARY = sortzilla

# Release directory
RELEASEDIR = release

# Version number
VERSION = 1.0.0

# Clean target
clean:
	rm -rf $(OUTDIR) $(RELEASEDIR)

# Build targets
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTDIR)/linux/$(BINARY)

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(OUTDIR)/linux/$(BINARY)

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUTDIR)/macos/$(BINARY)

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(OUTDIR)/macos/$(BINARY)

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(OUTDIR)/windows/$(BINARY).exe

# Release target
release: clean build-linux-amd64 build-linux-arm64 build-macos-amd64 build-macos-arm64 build-windows-amd64
	mkdir -p $(RELEASEDIR)/linux
	cp $(OUTDIR)/linux/$(BINARY) $(RELEASEDIR)/linux/$(BINARY)
	zip -r $(RELEASEDIR)/$(BINARY)-$(VERSION)-linux.zip $(RELEASEDIR)/linux
	rm -rf $(RELEASEDIR)/linux
	mkdir -p $(RELEASEDIR)/macos
	cp $(OUTDIR)/macos/$(BINARY) $(RELEASEDIR)/macos/$(BINARY)
	zip -r $(RELEASEDIR)/$(BINARY)-$(VERSION)-macos.zip $(RELEASEDIR)/macos
	rm -rf $(RELEASEDIR)/macos
	mkdir -p $(RELEASEDIR)/windows
	cp $(OUTDIR)/windows/$(BINARY).exe $(RELEASEDIR)/windows/$(BINARY).exe
	zip -r $(RELEASEDIR)/$(BINARY)-$(VERSION)-windows.zip $(RELEASEDIR)/windows
	rm -rf $(RELEASEDIR)/windows
