package route

import "time"

type TaskGroup struct {

	Tasks            []*Task
	Meta             map[string]string
	Services         []*Service

}


// Task is a single process in a task group.
type Task struct {
	Name            string
	Driver          string
	User            string

	Config          map[string]interface{}
	Constraints     []*Constraint

	Env             map[string]string


	Services        []*Service


	Meta            map[string]string
	KillTimeout     *time.Duration `mapstructure:"kill_timeout"`
	ShutdownDelay   time.Duration `mapstructure:"shutdown_delay"`
	KillSignal      string        `mapstructure:"kill_signal"`
	Kind            string
}
// Service represents a Consul service definition.
type Service struct {

	Id                string

}


// Job is used to serialize a job.
type Job struct {
	Stop              *bool
	Region            *string
	Namespace         *string
	ID                *string
	ParentID          *string
	Name              *string
	Type              *string
	Priority          *int
	AllAtOnce         *bool `mapstructure:"all_at_once"`
	Datacenters       []string
	Constraints       []*Constraint
	TaskGroups        []*TaskGroup
	Dispatched        bool
	Payload           []byte
	Meta              map[string]string
	ConsulToken       *string `mapstructure:"consul_token"`
	VaultToken        *string `mapstructure:"vault_token"`
	Status            *string
	StatusDescription *string
	Stable            *bool
	Version           *uint64
	SubmitTime        *int64
	CreateIndex       *uint64
	ModifyIndex       *uint64
	JobModifyIndex    *uint64
}