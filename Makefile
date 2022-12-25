VERSION = 0.0.6

GC = go build -ldflags="-X 'vpn-dns/cmd.Version=v$(VERSION)' -s -w"
ENTRYFILE = main.go
MEMTEST_ENTRYFILE = memtest/main.go

BUILD_DIR = build
BINARY_NAME = vpn-dns

DARWIN_ARM64 = $(BUILD_DIR)/$(BINARY_NAME)_arm64
DARWIN_AMD64 = $(BUILD_DIR)/$(BINARY_NAME)_amd64
DARWIN_MEMTEST = $(BUILD_DIR)/$(BINARY_NAME)_memtest

define build_binary
    env GOOS="darwin" GOARCH="$(2)" $(GC) -o "$(1)" $(ENTRYFILE)
endef

GOSRC := \
	$(wildcard cmd/*.go) \
	$(wildcard internal/**/*.go) \
	$(wildcard pkg/**/*.go)

$(DARWIN_ARM64): $(GOSRC)
	$(call build_binary,$(DARWIN_ARM64),arm64)

$(DARWIN_AMD64): $(GOSRC)
	$(call build_binary,$(DARWIN_AMD64),amd64)

$(MEMTEST_CURRENT): $(GOSRC)
	go build \
		-ldflags="-X 'vpn-dns/cmd.Version=memtest'" \
		-o "$(MEMTEST_CURRENT)" \
		"$(MEMTEST_ENTRYFILE)"

all: \
	$(DARWIN_ARM64) \
	$(DARWIN_AMD64) \
	$(MEMTEST_CURRENT)

.PHONY: test
test:
	go test -cover ./...

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

# Memory testing utils.

# Runs special version of app with memory debugging utils.
.PHONY: memtest-run
memtest-run:
	go run $(MEMTEST_ENTRYFILE) run

# Opens browser with app memory info. Special version of app must be running.
.PHONY: memtest-view
memtest-view:
	go tool pprof \
        -http=:8081 \
        -alloc_space \
        http://localhost:8080/debug/pprof/heap

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
