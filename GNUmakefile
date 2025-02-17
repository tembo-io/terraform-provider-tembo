
HOSTNAME = tembo
NAME = tembo
VERSION ?= 1.9.0-alpha1
BINARY = terraform-provider-${NAME}
INSTALL_PATH = ~/.terraform.d/plugins/registry.terraform.io/${HOSTNAME}/${NAME}/${VERSION}/${OS_ARCH}

# Dynamic OS and ARCH detection
OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH = $(shell uname -m)
# Convert MacOS architecture names to match Terraform's naming
ifeq ($(ARCH),arm64)
    OS_ARCH = $(OS)_$(ARCH)
else ifeq ($(ARCH),x86_64)
    OS_ARCH = $(OS)_amd64
else
    OS_ARCH = $(OS)_$(ARCH)
endif

default: check

# Format, lint and generate targets
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: generate
generate:
	go generate ./...

# Combined check target
.PHONY: check
check: fmt lint generate
	git diff --exit-code || (echo "Found changes after running formatters and generators. Please commit these changes." && exit 1)

# Build and install provider locally
.PHONY: install
install: clean
	GORELEASER_CURRENT_TAG=v${VERSION} goreleaser build --snapshot --clean --single-target
	mkdir -p ${INSTALL_PATH}
	cp dist/${BINARY}_${OS_ARCH}_v1/${BINARY}_v${VERSION}-SNAPSHOT-* ${INSTALL_PATH}/${BINARY}_v${VERSION}
	chmod +x ${INSTALL_PATH}/${BINARY}_v${VERSION}

# Example of how to clean previous installations
.PHONY: clean
clean:
	rm -rf ${INSTALL_PATH}

# Build and install provider locally
.PHONY: install
install-local:
	GORELEASER_CURRENT_TAG=${VERSION} goreleaser build --snapshot --clean --single-target
	mkdir -p ${INSTALL_PATH}
	cp dist/terraform-provider-tembo_${OS_ARCH}_v1/${BINARY}_${VERSION}-SNAPSHOT-* ${INSTALL_PATH}/${BINARY}_${VERSION}
	chmod +x ${INSTALL_PATH}/${BINARY}_${VERSION}

# Example of how to clean previous installations
.PHONY: clean
clean:
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
