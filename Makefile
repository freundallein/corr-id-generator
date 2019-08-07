build:
	 protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/gitlab.com/freundallein/corr-id-generator/ grpc/proto/corr-id-generator.proto
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o bin/corr-id-generator
	 docker build -t corridgen:latest --no-cache .