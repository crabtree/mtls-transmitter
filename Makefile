.PHONY: validate build

validate: format
	#==> Verify modules.
	go mod verify
	#==> It's worth running 'go test' to make sure things compile correctly.
	go test -race ./...

format:
	#==> Verify and update formatting, don't forget to commit on changes.
	go fmt ./...

build:
	go build -o mtls-transmitter ./cmd/transmitter
