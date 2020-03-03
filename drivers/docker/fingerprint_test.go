package docker

import (
	"testing"

	"github.com/actiontech/dtle/client/testutil"
	"github.com/actiontech/dtle/helper/testlog"
	"github.com/actiontech/dtle/plugins/drivers"
	tu "github.com/actiontech/dtle/testutil"
	"github.com/stretchr/testify/require"
)

// TestDockerDriver_FingerprintHealth asserts that docker reports healthy
// whenever Docker is supported.
//
// In Linux CI and AppVeyor Windows environment, it should be enabled.
func TestDockerDriver_FingerprintHealth(t *testing.T) {
	if !tu.IsCI() {
		t.Parallel()
	}
	testutil.DockerCompatible(t)

	d := NewDockerDriver(testlog.HCLogger(t)).(*Driver)

	fp := d.buildFingerprint()
	require.Equal(t, drivers.HealthStateHealthy, fp.Health)
}
