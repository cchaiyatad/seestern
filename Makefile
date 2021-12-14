build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/osx main.go

build-window:
	GOOS=darwin GOARCH=amd64 go build -o bin/osx main.go

build: build-mac build-window

run:
	go run main.go $(ARGS)

clean:
	rm -rf bin
	rm -rf out
