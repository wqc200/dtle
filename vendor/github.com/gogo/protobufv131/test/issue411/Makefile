regenerate:
	go install github.com/gogo/protobufv131/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobufv131/timestamp.proto=github.com/gogo/protobufv131/types:. --proto_path=../../../../../:../../protobufv131/:. issue411.proto
