package fingerprint

import (
	"testing"

	"github.com/actiontech/dtle/client/config"
	"github.com/actiontech/dtle/helper/testlog"
	"github.com/actiontech/dtle/nomad/structs"
	"github.com/actiontech/dtle/testutil"
)

func TestVaultFingerprint(t *testing.T) {
	tv := testutil.NewTestVault(t)
	defer tv.Stop()

	fp := NewVaultFingerprint(testlog.HCLogger(t))
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	conf := config.DefaultConfig()
	conf.VaultConfig = tv.Config

	request := &FingerprintRequest{Config: conf, Node: node}
	var response FingerprintResponse
	err := fp.Fingerprint(request, &response)
	if err != nil {
		t.Fatalf("Failed to fingerprint: %s", err)
	}

	if !response.Detected {
		t.Fatalf("expected response to be applicable")
	}

	assertNodeAttributeContains(t, response.Attributes, "vault.accessible")
	assertNodeAttributeContains(t, response.Attributes, "vault.version")
	assertNodeAttributeContains(t, response.Attributes, "vault.cluster_id")
	assertNodeAttributeContains(t, response.Attributes, "vault.cluster_name")
}
