//+build !linux

package allocrunner

import (
	clientconfig "github.com/actiontech/dtle/client/config"
	"github.com/actiontech/dtle/client/pluginmanager/drivermanager"
	"github.com/actiontech/dtle/nomad/structs"
	"github.com/actiontech/dtle/plugins/drivers"
	hclog "github.com/hashicorp/go-hclog"
)

// TODO: Support windows shared networking
func newNetworkManager(alloc *structs.Allocation, driverManager drivermanager.Manager) (nm drivers.DriverNetworkManager, err error) {
	return nil, nil
}

func newNetworkConfigurator(log hclog.Logger, alloc *structs.Allocation, config *clientconfig.Config) (NetworkConfigurator, error) {
	return &hostNetworkConfigurator{}, nil
}
