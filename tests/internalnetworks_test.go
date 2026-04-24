//go:build functional

package tests

import (
	"net/http"
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
	"github.com/stretchr/testify/require"
)

func TestInternalNetworksFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetInternalNetworksClient(env.ctx)

	_, _, err := apiClient.InternalNetworksAPI.ListInternalNetworks(env.ctx).Page(1).Limit(10).Execute()
	require.NoError(t, err)

	siteID, hasSiteID, err := envInt64("GO_CISCOSECUREACCESS_INTERNAL_NETWORK_SITE_ID")
	require.NoError(t, err)
	if !hasSiteID {
		siteID = 1
	}

	name := uniqueName("go-sdk-internal-network")
	createReq := *internalnetworks.NewCreateInternalNetworkRequest(name, "198.51.100.0", 24)
	createReq.SetSiteId(siteID)

	createResp, createHTTP, err := apiClient.InternalNetworksAPI.CreateInternalNetwork(env.ctx).
		CreateInternalNetworkRequest(createReq).
		Execute()
	if err != nil && createHTTP != nil && (createHTTP.StatusCode == http.StatusBadRequest || createHTTP.StatusCode == http.StatusNotFound) {
		t.Skipf("skipping internal network create: siteId %d unavailable in this tenant; set GO_CISCOSECUREACCESS_INTERNAL_NETWORK_SITE_ID", siteID)
	}
	require.NoError(t, err)
	require.NotNil(t, createResp)
	internalNetworkID := createResp.GetOriginId()

	t.Cleanup(func() {
		if internalNetworkID == 0 {
			return
		}
		_, cleanupErr := apiClient.InternalNetworksAPI.DeleteInternalNetwork(env.ctx, internalNetworkID).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup internal network %d: %v", internalNetworkID, cleanupErr)
		}
	})

	getResp, _, err := apiClient.InternalNetworksAPI.GetInternalNetwork(env.ctx, internalNetworkID).Execute()
	require.NoError(t, err)
	if getResp != nil {
		require.Equal(t, name, getResp.GetName())
	}

	updatedName := uniqueName("go-sdk-internal-network-updated")
	updateReq := createReq
	updateReq.SetName(updatedName)
	updateResp, _, err := apiClient.InternalNetworksAPI.UpdateInternalNetwork(env.ctx, internalNetworkID).
		CreateInternalNetworkRequest(updateReq).
		Execute()
	require.NoError(t, err)
	if updateResp != nil && updateResp.GetName() != "" {
		require.Equal(t, updatedName, updateResp.GetName())
	}

	refetchResp, _, err := apiClient.InternalNetworksAPI.GetInternalNetwork(env.ctx, internalNetworkID).Execute()
	require.NoError(t, err)
	if refetchResp != nil {
		require.Equal(t, updatedName, refetchResp.GetName())
	}

	_, err = apiClient.InternalNetworksAPI.DeleteInternalNetwork(env.ctx, internalNetworkID).Execute()
	require.NoError(t, err)
	internalNetworkID = 0
}
