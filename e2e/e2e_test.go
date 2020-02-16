package e2e

import (
	"testing"

	_ "github.com/actiontech/dtle/e2e/affinities"
	_ "github.com/actiontech/dtle/e2e/allocstats"
	_ "github.com/actiontech/dtle/e2e/clientstate"
	_ "github.com/actiontech/dtle/e2e/connect"
	_ "github.com/actiontech/dtle/e2e/consul"
	_ "github.com/actiontech/dtle/e2e/consultemplate"
	_ "github.com/actiontech/dtle/e2e/deployment"
	_ "github.com/actiontech/dtle/e2e/example"
	_ "github.com/actiontech/dtle/e2e/hostvolumes"
	_ "github.com/actiontech/dtle/e2e/nomad09upgrade"
	_ "github.com/actiontech/dtle/e2e/nomadexec"
	_ "github.com/actiontech/dtle/e2e/spread"
	_ "github.com/actiontech/dtle/e2e/taskevents"
)

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}
