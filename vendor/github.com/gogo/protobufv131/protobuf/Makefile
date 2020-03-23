VERSION=3.9.1
URL="https://raw.githubusercontent.com/protocolbuffers/protobufv131/v${VERSION}/src/google/protobuf"

regenerate:
	go install github.com/gogo/protobufv131/protoc-gen-gogotypes
	go install github.com/gogo/protobufv131/protoc-min-version

	protoc-min-version \
	--version="3.0.0" \
	--gogotypes_out=../types/ \
	-I=. \
	google/protobufv131/any.proto \
	google/protobufv131/type.proto \
	google/protobufv131/empty.proto \
	google/protobufv131/api.proto \
	google/protobufv131/timestamp.proto \
	google/protobufv131/duration.proto \
	google/protobufv131/struct.proto \
	google/protobufv131/wrappers.proto \
	google/protobufv131/field_mask.proto \
	google/protobufv131/source_context.proto

	mv ../types/google/protobufv131/*.pb.go ../types/ || true
	rmdir ../types/google/protobuf || true
	rmdir ../types/google || true

update:
	go install github.com/gogo/protobufv131/gogoreplace

	(cd ./google/protobuf && rm descriptor.proto; wget ${URL}/descriptor.proto)
	# gogoprotobuf requires users to import gogo.proto which imports descriptor.proto
	# The descriptor.proto is only compatible with proto3 just because of the reserved keyword.
	# We remove it to stay compatible with previous versions of protoc before proto3
	gogoreplace 'reserved 38;' '//reserved 38;' ./google/protobufv131/descriptor.proto
	gogoreplace 'reserved 8;' '//reserved 8;' ./google/protobufv131/descriptor.proto
	gogoreplace 'reserved 9;' '//reserved 9;' ./google/protobufv131/descriptor.proto
	gogoreplace 'reserved 4;' '//reserved 4;' ./google/protobufv131/descriptor.proto
	gogoreplace 'reserved 5;' '//reserved 5;' ./google/protobufv131/descriptor.proto
	gogoreplace 'option go_package = "github.com/golang/protobufv131/protoc-gen-go/descriptor;descriptor";' 'option go_package = "descriptor";' ./google/protobufv131/descriptor.proto

	(cd ./google/protobufv131/compiler && rm plugin.proto; wget ${URL}/compiler/plugin.proto)
	gogoreplace 'option go_package = "github.com/golang/protobufv131/protoc-gen-go/plugin;plugin_go";' 'option go_package = "plugin_go";' ./google/protobufv131/compiler/plugin.proto

	(cd ./google/protobuf && rm any.proto; wget ${URL}/any.proto)
	gogoreplace 'go_package = "github.com/golang/protobufv131/ptypes/any";' 'go_package = "types";' ./google/protobufv131/any.proto
	(cd ./google/protobuf && rm empty.proto; wget ${URL}/empty.proto)
	gogoreplace 'go_package = "github.com/golang/protobufv131/ptypes/empty";' 'go_package = "types";' ./google/protobufv131/empty.proto
	(cd ./google/protobuf && rm timestamp.proto; wget ${URL}/timestamp.proto)
	gogoreplace 'go_package = "github.com/golang/protobufv131/ptypes/timestamp";' 'go_package = "types";' ./google/protobufv131/timestamp.proto
	(cd ./google/protobuf && rm duration.proto; wget ${URL}/duration.proto)
	gogoreplace 'go_package = "github.com/golang/protobufv131/ptypes/duration";' 'go_package = "types";' ./google/protobufv131/duration.proto
	(cd ./google/protobuf && rm struct.proto; wget ${URL}/struct.proto)
	gogoreplace 'go_package = "github.com/golang/protobufv131/ptypes/struct;structpb";' 'go_package = "types";' ./google/protobufv131/struct.proto
	(cd ./google/protobuf && rm wrappers.proto; wget ${URL}/wrappers.proto)
	gogoreplace 'go_package = "github.com/golang/protobufv131/ptypes/wrappers";' 'go_package = "types";' ./google/protobufv131/wrappers.proto
	(cd ./google/protobuf && rm field_mask.proto; wget ${URL}/field_mask.proto)
	gogoreplace 'option go_package = "google.golang.org/genproto/protobufv131/field_mask;field_mask";' 'option go_package = "types";' ./google/protobufv131/field_mask.proto
	(cd ./google/protobuf && rm api.proto; wget ${URL}/api.proto)
	gogoreplace 'option go_package = "google.golang.org/genproto/protobufv131/api;api";' 'option go_package = "types";' ./google/protobufv131/api.proto
	(cd ./google/protobuf && rm type.proto; wget ${URL}/type.proto)
	gogoreplace 'option go_package = "google.golang.org/genproto/protobufv131/ptype;ptype";' 'option go_package = "types";' ./google/protobufv131/type.proto
	(cd ./google/protobuf && rm source_context.proto; wget ${URL}/source_context.proto)
	gogoreplace 'option go_package = "google.golang.org/genproto/protobufv131/source_context;source_context";' 'option go_package = "types";' ./google/protobufv131/source_context.proto


