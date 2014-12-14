all:    install

install:
	go install

test:
	go test

clean:
	go clean ./...

nuke:
	go clean -i ./...

