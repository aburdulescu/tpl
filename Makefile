VERSION=0.1

build:
	go build -ldflags "-s -w -X main.version="$(VERSION)

install:
	go install -ldflags "-s -w -X main.version="$(VERSION)

clean:
	go clean
