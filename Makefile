.PHONY: clear run

GC = go build -ldflags="-s -w"
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

build: all

run:
	go run "$(ENTRYFILE)"

clear:
	rm -rf "$(BUILD_DIR)"
