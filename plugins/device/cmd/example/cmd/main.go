package main

import (
	log "github.com/hashicorp/go-hclog"

	"github.com/actiontech/dtle/plugins"
	"github.com/actiontech/dtle/plugins/device/cmd/example"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of our example device plugin
func factory(log log.Logger) interface{} {
	return example.NewExampleDevice(log)
}
