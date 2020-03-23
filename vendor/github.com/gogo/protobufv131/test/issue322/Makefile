regenerate:
	go install github.com/gogo/protobufv131/protoc-gen-gogofast
	protoc-min-version --version="3.0.0" --gogofast_out=Mgoogle/protobufv131/descriptor.proto=github.com/gogo/protobufv131/protoc-gen-gogo/descriptor:. --proto_path=../../../../../:../../protobufv131/:. *.proto
