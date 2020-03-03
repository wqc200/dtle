package client

import (
	"testing"

	"github.com/actiontech/dtle/acl"
	"github.com/actiontech/dtle/client/config"
	"github.com/actiontech/dtle/client/structs"
	"github.com/actiontech/dtle/nomad/mock"
	nstructs "github.com/actiontech/dtle/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestClientStats_Stats(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	client, cleanup := TestClient(t, nil)
	defer cleanup()

	req := &nstructs.NodeSpecificRequest{}
	var resp structs.ClientStatsResponse
	require.Nil(client.ClientRPC("ClientStats.Stats", &req, &resp))
	require.NotNil(resp.HostStats)
	require.NotNil(resp.HostStats.AllocDirStats)
	require.NotZero(resp.HostStats.Uptime)
}

func TestClientStats_Stats_ACL(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	server, addr, root := testACLServer(t, nil)
	defer server.Shutdown()

	client, cleanup := TestClient(t, func(c *config.Config) {
		c.Servers = []string{addr}
		c.ACLEnabled = true
	})
	defer cleanup()

	// Try request without a token and expect failure
	{
		req := &nstructs.NodeSpecificRequest{}
		var resp structs.ClientStatsResponse
		err := client.ClientRPC("ClientStats.Stats", &req, &resp)
		require.NotNil(err)
		require.EqualError(err, nstructs.ErrPermissionDenied.Error())
	}

	// Try request with an invalid token and expect failure
	{
		token := mock.CreatePolicyAndToken(t, server.State(), 1005, "invalid", mock.NodePolicy(acl.PolicyDeny))
		req := &nstructs.NodeSpecificRequest{}
		req.AuthToken = token.SecretID

		var resp structs.ClientStatsResponse
		err := client.ClientRPC("ClientStats.Stats", &req, &resp)

		require.NotNil(err)
		require.EqualError(err, nstructs.ErrPermissionDenied.Error())
	}

	// Try request with a valid token
	{
		token := mock.CreatePolicyAndToken(t, server.State(), 1007, "valid", mock.NodePolicy(acl.PolicyRead))
		req := &nstructs.NodeSpecificRequest{}
		req.AuthToken = token.SecretID

		var resp structs.ClientStatsResponse
		err := client.ClientRPC("ClientStats.Stats", &req, &resp)

		require.Nil(err)
		require.NotNil(resp.HostStats)
	}

	// Try request with a management token
	{
		req := &nstructs.NodeSpecificRequest{}
		req.AuthToken = root.SecretID

		var resp structs.ClientStatsResponse
		err := client.ClientRPC("ClientStats.Stats", &req, &resp)

		require.Nil(err)
		require.NotNil(resp.HostStats)
	}
}
