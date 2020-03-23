test:
	go install github.com/gogo/protobufv131/protoc-gen-gogo
	go install github.com/gogo/protobufv131/protoc-min-version
	go get -u golang.org/x/net/context
	go get -u google.golang.org/grpc
	protoc-min-version --version="3.0.0" --gogo_out=plugins=grpc:. --proto_path=../../../../../:../../protobufv131/:. issue427.proto
	go test ./...
