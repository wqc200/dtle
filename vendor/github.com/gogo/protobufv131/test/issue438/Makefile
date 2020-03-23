regenerate:
	go install github.com/gogo/protobufv131/protoc-gen-gogo
	go install github.com/gogo/protobufv131/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mgoogle/protobufv131/any.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/api.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/duration.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/empty.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/field_mask.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/source_context.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/struct.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/timestamp.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/type.proto=github.com/gogo/protobufv131/types,\
	Mgoogle/protobufv131/wrappers.proto=github.com/gogo/protobufv131/types,\
	:. --proto_path=../../../../../:../../protobufv131/:. issue438.proto
