regenerate:
	go install github.com/gogo/protobufv131/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/gogo/protobufv131/proto/test_proto,\
	Mgoogle/protobufv131/any.proto=github.com/gogo/protobufv131/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
