regenerate:
	go install github.com/gogo/protobufv131/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobufv131/protobufv131/ \
	--gogo_out=. events.proto 
