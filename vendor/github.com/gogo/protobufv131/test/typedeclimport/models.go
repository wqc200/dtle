package typedeclimport

import subpkg "github.com/gogo/protobufv131/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
