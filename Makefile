GO		?= go
DOCKER		?= docker
# -s removes symbol table and -ldflags -w debugging symbols
LDFLAGS		?= -asmflags -trimpath -ldflags "-s -w"
GOARCH		?= amd64
BINARY		?= expanses-feeder
CGO_ENABLED	?= 1

.PHONY: all escape-analysis test

build: test lint
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=$(GOARCH) \
	     $(GO) build $(LDFLAGS) -o $(BINARY)

lint:
	golangci-lint run ./...

escape-analysis:
	$(GO) build -gcflags="-m" 2>&1

docker-build:
	$(DOCKER) build --rm --target app -t .

docker-run:
	docker run --rm -v $(shell pwd):/app/config $(BINARY) &

test:
	go test ./...

clean:
	rm -f $(BINARY)

