.PHONY: clean test generate check build.local build.linux build.osx build.docker build.push

APP_BINARY              ?= rest-api
SOURCES                 = $(shell find . -name '*.go')

default: build.local

clean:
	rm -rf build

build.local: build/$(APP_BINARY)
build.linux: build/linux/$(APP_BINARY)
build.osx: build/osx/$(APP_BINARY)

build/$(APP_BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o $(APP_BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" cmd/main.go

build/linux/$(APP_BINARY): $(SOURCES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o build/linux/$(APP_BINARY) -ldflags "$(LDFLAGS)" cmd/main.go

build/osx/$(APP_BINARY): $(SOURCES)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o build/osx/$(APP_BINARY) -ldflags "$(LDFLAGS)" cmd/main.go
