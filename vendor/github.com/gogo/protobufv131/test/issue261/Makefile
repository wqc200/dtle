regenerate:
	go install github.com/gogo/protobufv131/protoc-min-version
	go install github.com/gogo/protobufv131/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --gogoslick_out=\
	Mgoogle/protobufv131/duration.proto=github.com/gogo/protobufv131/types,\
	:. \
	--proto_path=../../../../../:../../protobufv131/:. issue261.proto
