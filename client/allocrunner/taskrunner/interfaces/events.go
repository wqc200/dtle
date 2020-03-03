package interfaces

import "github.com/actiontech/dtle/nomad/structs"

type EventEmitter interface {
	EmitEvent(event *structs.TaskEvent)
}
