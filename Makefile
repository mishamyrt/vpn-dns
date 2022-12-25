VERSION = 0.0.5

GC = go build -ldflags="-X 'vpn-dns/cmd.Version=v$(VERSION)' -s -w"
ENTRYFILE = main.go

BUILD_DIR = build
BINARY_NAME = vpn-dns

DARWIN_ARM64 = $(BUILD_DIR)/$(BINARY_NAME)_arm64
DARWIN_AMD64 = $(BUILD_DIR)/$(BINARY_NAME)_amd64

define build_binary
    env GOOS="$(2)" GOARCH="$(3)" $(GC) -o "$(1)" "$(ENTRYFILE)"
endef

GOSRC := \
	$(wildcard cmd/*.go) \
	$(wildcard internal/**/*.go) \
	$(wildcard pkg/**/*.go)

$(DARWIN_ARM64): $(GOSRC)
	$(call build_binary,$(DARWIN_ARM64),darwin,arm64)

$(DARWIN_AMD64): $(GOSRC)
	$(call build_binary,$(DARWIN_AMD64),darwin,amd64)

all: \
	$(DARWIN_ARM64) \
	$(DARWIN_AMD64)

.PHONY: test
test:
	go test -cover ./...

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: release
release:
	git tag "v$(VERSION)"
	git push --tags

.PHONY: build
build: all

.PHONY: clean
clean:
	rm -rf "$(BUILD_DIR)"

.PHONY: lint
lint:
	golangci-lint run
	revive -config revive.toml  ./...

.PHONY: install
install: build
	cp -f "build/vpn-dns_$$(arch)" /usr/local/bin/vpn-dns
