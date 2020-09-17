build:
	go build -ldflags "-s -w"

install:
	go install -ldflags "-s -w"

clean:
	go clean
