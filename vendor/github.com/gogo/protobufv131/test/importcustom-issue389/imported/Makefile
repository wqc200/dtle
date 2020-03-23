regenerate:
	go install github.com/gogo/protobufv131/protoc-gen-gogo
	go install github.com/gogo/protobufv131/protoc-min-version
	protoc-min-version --version="3.0.0" --proto_path=../../../../../../:../../../protobufv131/:. --gogo_out=. a.proto