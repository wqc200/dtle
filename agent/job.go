package agent

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/docker/libkv/store"
	"github.com/hashicorp/serf/serf"

	uconf "udup/config"
	ulog "udup/logger"
)

var (
	ErrNoAgent = errors.New("No agent defined")
)

// JobStatus is the status of the Job.
type JobStatus int

const (
	Running JobStatus = iota
	Queued
	Stopped
	Failed
)

func (s JobStatus) String() string {
	switch s {
	case Running:
		return "running"
	case Queued:
		return "queued"
	case Stopped:
		return "stopped"
	case Failed:
		return "failed"
	default:
		return "unknown"
	}
}

type Job struct {
	// Job name. Must be unique, acts as the id.
	Name string `json:"name"`

	Failover bool `json:"failover"`

	// Job status
	Status JobStatus `json:"status"`

	// Tags of the target servers to run this job against.
	Tags map[string]string `json:"tags"`

	// Pointer to the calling agent.
	agent *Agent `json:"-"`

	running sync.Mutex

	lock store.Locker

	// Processors to use for this job
	Processors map[string]*uconf.DriverConfig `json:"processors"`
}

// Start the job
func (j *Job) Start() {
	j.running.Lock()
	defer j.running.Unlock()

	if j.agent != nil {
		if j.Status != Running {
			for _, m := range j.agent.serf.Members() {
				for k, v := range j.Processors {
					if m.Name == v.NodeName && m.Status != serf.StatusAlive {
						j.Status = Queued
						j.Processors[k].Running = false
						if err := j.agent.store.UpsertJob(j); err != nil {
							ulog.Logger.Fatal(err)
						}
						return
					}
				}
			}

			ulog.Logger.WithFields(logrus.Fields{
				"job": j.Name,
			}).Debug("job: Run job")

			for k, v := range j.Processors {
				if v.Running != true {
					go j.StartJobQuery(k)
				}
			}
		}
	}
}

// Stop the job
func (j *Job) Stop() {
	j.running.Lock()
	defer j.running.Unlock()

	if j.agent != nil && j.Status == Running {
		ulog.Logger.WithFields(logrus.Fields{
			"job": j.Name,
		}).Debug("job: Stop job")

		for k, _ := range j.Processors {
			go j.StopJobQuery(k)
		}
	}
}

// Lock the job in store
func (j *Job) Lock() error {
	// Maybe we are testing
	if j.agent == nil {
		return ErrNoAgent
	}

	lockKey := fmt.Sprintf("%s/job_locks/%s", keyspace, j.Name)
	// TODO: LockOptions empty is a temporary fix until https://github.com/docker/libkv/pull/99 is fixed
	l, err := j.agent.store.Client.NewLock(lockKey, &store.LockOptions{RenewLock: make(chan (struct{}))})
	if err != nil {
		return err
	}
	j.lock = l

	_, err = j.lock.Lock(nil)
	if err != nil {
		return err
	}

	return nil
}

// Unlock the job in store
func (j *Job) Unlock() error {
	// Maybe we are testing
	if j.agent == nil {
		return ErrNoAgent
	}

	if err := j.lock.Unlock(); err != nil {
		return err
	}

	return nil
}
