//go:build functional

package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/networks"
	"github.com/stretchr/testify/require"
)

func TestNetworksFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetNetworksClient(env.ctx)

	ipv4Address := strings.TrimSpace(os.Getenv("GO_CISCOSECUREACCESS_VERIFIED_IP4"))
	if ipv4Address == "" {
		t.Skip("set GO_CISCOSECUREACCESS_VERIFIED_IP4 to a Secure Access verified IP address to run network CRUD functional test")
	}

	_, _, err := apiClient.NetworksAPI.ListNetworks(env.ctx).Page(1).Limit(10).Execute()
	require.NoError(t, err)

	name := uniqueName("go-sdk-network")
	createReq := *networks.NewCreateNetworkRequest(name, 32, false, "OPEN")
	createReq.SetIpAddress(ipv4Address)
	createResp, createHTTP, err := apiClient.NetworksAPI.CreateNetwork(env.ctx).CreateNetworkRequest(createReq).Execute()
	if err != nil && createHTTP != nil && (createHTTP.StatusCode == 401 || createHTTP.StatusCode == 403) {
		t.Skipf("skipping network CRUD: insufficient API permissions (%s)", createHTTP.Status)
	}
	require.NoError(t, err)
	networkID := createResp.GetOriginId()

	t.Cleanup(func() {
		if networkID == 0 {
			return
		}
		_, _, cleanupErr := apiClient.NetworksAPI.DeleteNetwork(env.ctx, networkID).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup network %d: %v", networkID, cleanupErr)
		}
	})

	getResp, _, err := apiClient.NetworksAPI.GetNetwork(env.ctx, networkID).Execute()
	require.NoError(t, err)
	require.Equal(t, name, getResp.GetName())

	updatedName := uniqueName("go-sdk-network-updated")
	updateReq := *networks.NewUpdateNetworkRequest(updatedName, true, "OPEN")
	updateReq.SetIpAddress(ipv4Address)
	updateReq.SetPrefixLength(32)
	updateResp, _, err := apiClient.NetworksAPI.UpdateNetwork(env.ctx, networkID).UpdateNetworkRequest(updateReq).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedName, updateResp.GetName())

	_, _, err = apiClient.NetworksAPI.DeleteNetwork(env.ctx, networkID).Execute()
	require.NoError(t, err)
	networkID = 0
}
