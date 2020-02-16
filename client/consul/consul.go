package consul

import (
	"github.com/actiontech/dtle/command/agent/consul"
	"github.com/actiontech/dtle/nomad/structs"
)

// ConsulServiceAPI is the interface the Nomad Client uses to register and
// remove services and checks from Consul.
type ConsulServiceAPI interface {
	RegisterGroup(*structs.Allocation) error
	RemoveGroup(*structs.Allocation) error
	UpdateGroup(oldAlloc, newAlloc *structs.Allocation) error
	RegisterTask(*consul.TaskServices) error
	RemoveTask(*consul.TaskServices)
	UpdateTask(old, newTask *consul.TaskServices) error
	AllocRegistrations(allocID string) (*consul.AllocRegistration, error)
	UpdateTTL(id, output, status string) error
}
