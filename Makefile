.PHONY: clear run lint release

VERSION = 0.0.3

GC = go build -ldflags="-X 'vpn-dns/cmd.Version=v$(VERSION)' -s -w"
ENTRYFILE = main.go

BUILD_DIR = build
BINARY_NAME = vpn-dns

CONFIG_DIR = $(HOME)/.config/vpn-dns
CONFIG_PATH = $(CONFIG_DIR)/config.yaml

DARWIN_ARM64 = $(BUILD_DIR)/$(BINARY_NAME)_arm64
DARWIN_AMD64 = $(BUILD_DIR)/$(BINARY_NAME)_amd64

all: \
	$(DARWIN_ARM64) \
	$(DARWIN_AMD64)

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

release:
	git tag "v$(VERSION)"
	git push --tags

build: all

run:
	go run "$(ENTRYFILE)"

clear:
	rm -rf "$(BUILD_DIR)"

lint:
	golangci-lint run -E lll -E misspell -E prealloc -E stylecheck -E gocritic
