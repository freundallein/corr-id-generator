
export BIN_DIR=bin
export PORT=7891
export MACHINE_ID=1

export IMAGE_NAME=freundallein/corridgen:latest

init:
	git config core.hooksPath .githooks
run:
	go run main.go
test:
	go test -cover ./...
build:
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/freundallein/corr-id-generator/ grpc/proto/corr-id-generator.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o $$BIN_DIR/corridgen
dockerbuild:
	docker build -t $$IMAGE_NAME -f Dockerfile .
distribute:
	echo "$$DOCKER_PASSWORD" | docker login -u "$$DOCKER_USERNAME" --password-stdin
	docker build -t $$IMAGE_NAME .
	docker push $$IMAGE_NAME