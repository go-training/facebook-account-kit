GO ?= go
GOFILES := $(shell find . -name "*.go" -type f)
EXECUTABLE := kit

.PHONY: build
build: $(EXECUTABLE)

.PHONY: kit
$(EXECUTABLE): $(SOURCES)
	$(GO) build -v -o bin/$@

.PHONY: lint
lint:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config .revive.toml ./... || exit 1

clean:
	$(GO) clean -modcache -cache -x -i ./...
