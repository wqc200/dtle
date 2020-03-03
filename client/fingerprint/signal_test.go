package fingerprint

import (
	"testing"

	"github.com/actiontech/dtle/helper/testlog"
	"github.com/actiontech/dtle/nomad/structs"
)

func TestSignalFingerprint(t *testing.T) {
	fp := NewSignalFingerprint(testlog.HCLogger(t))
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	response := assertFingerprintOK(t, fp, node)
	assertNodeAttributeContains(t, response.Attributes, "os.signals")
}
