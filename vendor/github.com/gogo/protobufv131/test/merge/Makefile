regenerate:
	go install github.com/gogo/protobufv131/protoc-min-version
	go install github.com/gogo/protobufv131/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. \
	--proto_path=../../../../../:../../protobufv131/:. merge.proto
