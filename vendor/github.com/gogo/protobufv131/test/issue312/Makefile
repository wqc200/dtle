regenerate:
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobufv131/protobufv131/ \
	--gogo_out=. issue312.proto 
	(cd events && make regenerate)
