package plugins

import (
	"fmt"

	"github.com/actiontech/dtle/plugins/device"
	"github.com/actiontech/dtle/plugins/drivers"
	log "github.com/hashicorp/go-hclog"
)

// PluginFactory returns a new plugin instance
type PluginFactory func(log log.Logger) interface{}

// Serve is used to serve a new Nomad plugin
func Serve(f PluginFactory) {
	logger := log.New(&log.LoggerOptions{
		Level:      log.Trace,
		JSONFormat: true,
	})

	plugin := f(logger)
	switch p := plugin.(type) {
	case device.DevicePlugin:
		device.Serve(p, logger)
	case drivers.DriverPlugin:
		drivers.Serve(p, logger)
	default:
		fmt.Println("Unsupported plugin type")
	}
}
