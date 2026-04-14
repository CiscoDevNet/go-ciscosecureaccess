//go:build functional

package tests

import (
	"net/http"
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/ntg"
	"github.com/stretchr/testify/require"
)

func TestNetworkTunnelGroupsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetNtgClient(env.ctx)

	_, _, err := apiClient.NetworkTunnelGroupsAPI.ListNetworkTunnelGroups(env.ctx).Limit(10).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.NetworkTunnelGroupsRegionsAPI.ListNetworkTunnelGroupRegions(env.ctx).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupBulkState(env.ctx).Limit(10).Execute()
	require.NoError(t, err)

	authPrefixValue := uniqueName("testtunnel")
	createReq := *ntg.NewAddNetworkTunnelGroupRequest(
		uniqueName("go-sdk-ntg"),
		"us-test-2",
		ntg.StringAsAddNetworkTunnelGroupRequestAuthIdPrefix(&authPrefixValue),
		"t3stingTunn3lNow",
	)
	createReq.SetDeviceType(ntg.OTHER)
	createReq.SetRouting(ntg.RoutingRequest{
		Type: "static",
		Data: ntg.RoutingRequestData{
			StaticDataRequestObj: &ntg.StaticDataRequestObj{NetworkCIDRs: []string{"192.16.30.2/32"}},
		},
	})

	createResp, _, err := apiClient.NetworkTunnelGroupsAPI.AddNetworkTunnelGroup(env.ctx).AddNetworkTunnelGroupRequest(createReq).Execute()
	require.NoError(t, err)
	id := createResp.GetId()

	t.Cleanup(func() {
		if id == 0 {
			return
		}
		_, cleanupErr := apiClient.NetworkTunnelGroupsAPI.DeleteNetworkTunnelGroup(env.ctx, id).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup tunnel group %d: %v", id, cleanupErr)
		}
	})

	getResp, _, err := apiClient.NetworkTunnelGroupsAPI.GetNetworkTunnelGroup(env.ctx, id).Execute()
	require.NoError(t, err)
	require.Equal(t, createReq.GetName(), getResp.GetName())

	updatedName := uniqueName("go-sdk-ntg-updated")
	patchValue := ntg.StringAsPatchNetworkTunnelGroupRequestInnerValue(&updatedName)
	patchReq := []ntg.PatchNetworkTunnelGroupRequestInner{
		*ntg.NewPatchNetworkTunnelGroupRequestInner("replace", "/name", patchValue),
	}
	patchResp, _, err := apiClient.NetworkTunnelGroupsAPI.PatchNetworkTunnelGroup(env.ctx, id).PatchNetworkTunnelGroupRequestInner(patchReq).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedName, patchResp.GetName())

	_, stateHTTP, stateErr := apiClient.NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupState(env.ctx, id).Execute()
	if stateErr != nil {
		require.NotNil(t, stateHTTP)
		require.Contains(t, []int{http.StatusOK, http.StatusNotFound}, stateHTTP.StatusCode)
	}

	_, peersHTTP, peersErr := apiClient.NetworkTunnelGroupsPeerStateAPI.GetPeersState(env.ctx, id).Execute()
	if peersErr != nil {
		require.NotNil(t, peersHTTP)
		require.Contains(t, []int{http.StatusOK, http.StatusNotFound}, peersHTTP.StatusCode)
	}

	_, err = apiClient.NetworkTunnelGroupsAPI.DeleteNetworkTunnelGroup(env.ctx, id).Execute()
	require.NoError(t, err)
	id = 0
}
