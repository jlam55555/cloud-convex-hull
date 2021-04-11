### packaging golang for aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
GO_BINARY_PATH:=$(BUILDDIR)/$(GO_BINARY)
GO_ZIP_PATH:=$(GO_BINARY_PATH).zip
GO_AWS_ZIP_PATH:=fileb://$(GO_ZIP_PATH)

$(GO_BINARY_PATH): $(GO_SOURCES)
	$(GO_ENVVAR) go build -o $@ $(GO_LDFLAGS) $(GO_PACKAGE)

# for testing the build
.PHONY:
target-build: $(GO_BINARY_PATH)

$(GO_ZIP_PATH): $(GO_BINARY_PATH)
	zip -j $@ $<

.PHONY:
target-clean:
	rm -rf $(BUILDDIR)

.PHONY:
lint:
	golint src/$(GO_PACKAGE)/**
	go fmt $(GO_PACKAGE)
