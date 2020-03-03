package loader

import (
	"github.com/actiontech/dtle/plugins/base"
	"github.com/actiontech/dtle/plugins/device"
)

var (
	// AgentSupportedApiVersions is the set of API versions supported by the
	// Nomad agent by plugin type.
	AgentSupportedApiVersions = map[string][]string{
		base.PluginTypeDevice: {device.ApiVersion010},
		base.PluginTypeDriver: {device.ApiVersion010},
	}
)
