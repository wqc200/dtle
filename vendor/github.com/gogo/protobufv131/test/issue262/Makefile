regenerate:
	go install github.com/gogo/protobufv131/protoc-min-version
	go install github.com/gogo/protobufv131/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobufv131/protobufv131/ \
	--gogoslick_out=Mgoogle/protobufv131/timestamp.proto=github.com/gogo/protobufv131/types:. timefail.proto
