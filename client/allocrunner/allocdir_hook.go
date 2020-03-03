package allocrunner

import (
	"github.com/actiontech/dtle/client/allocdir"
	log "github.com/hashicorp/go-hclog"
)

// allocDirHook creates and destroys the root directory and shared directories
// for an allocation.
type allocDirHook struct {
	allocDir *allocdir.AllocDir
	logger   log.Logger
}

func newAllocDirHook(logger log.Logger, allocDir *allocdir.AllocDir) *allocDirHook {
	ad := &allocDirHook{
		allocDir: allocDir,
	}
	ad.logger = logger.Named(ad.Name())
	return ad
}

func (h *allocDirHook) Name() string {
	return "alloc_dir"
}

func (h *allocDirHook) Prerun() error {
	return h.allocDir.Build()
}

func (h *allocDirHook) Destroy() error {
	return h.allocDir.Destroy()
}
