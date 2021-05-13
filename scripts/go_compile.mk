### packaging golang for aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
PRESIGN_GO_BINARY_PATH:=$(BUILDDIR)/$(PRESIGN_GO_BINARY)
PRESIGN_GO_ZIP_PATH:=$(PRESIGN_GO_BINARY_PATH).zip
PRESIGN_GO_AWS_ZIP_PATH:=fileb://$(PRESIGN_GO_ZIP_PATH)

$(PRESIGN_GO_BINARY_PATH): $(GO_SOURCES)
	$(GO_ENVVAR) go build -o $@ $(GO_LDFLAGS) $(PRESIGN_GO_PACKAGE)

# for testing the build
.PHONY:
target-build: $(PRESIGN_GO_BINARY_PATH) $(CH_GO_BINARY_PATH)

$(PRESIGN_GO_ZIP_PATH): $(PRESIGN_GO_BINARY_PATH)
	zip -j $@ $<

### convex hull lambda
CH_GO_BINARY_PATH:=$(BUILDDIR)/$(CH_GO_BINARY)
CH_GO_ZIP_PATH:=$(CH_GO_BINARY_PATH).zip
CH_GO_AWS_ZIP_PATH:=fileb://$(CH_GO_ZIP_PATH)

$(CH_GO_BINARY_PATH): $(GO_SOURCES)
	$(GO_ENVVAR) go build -o $@ $(GO_LDFLAGS) $(CH_GO_PACKAGE)

$(CH_GO_ZIP_PATH): $(CH_GO_BINARY_PATH)
	zip -j $@ $<

### general

.PHONY:
target-clean:
	rm -rf $(BUILDDIR)

.PHONY:
lint:
	golint src/$(PRESIGN_GO_PACKAGE)/**
	go fmt $(PRESIGN_GO_PACKAGE)