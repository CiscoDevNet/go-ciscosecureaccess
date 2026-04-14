//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/sites"
	"github.com/stretchr/testify/require"
)

func TestSitesFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetSitesClient(env.ctx)

	_, _, err := apiClient.SitesAPI.ListSites(env.ctx).Page(1).Limit(10).Execute()
	require.NoError(t, err)

	name := uniqueName("go-sdk-site")
	createReq := *sites.NewCreateSiteRequest(name)
	createResp, createHTTP, err := apiClient.SitesAPI.CreateSite(env.ctx).CreateSiteRequest(createReq).Execute()
	if err != nil && createHTTP != nil && (createHTTP.StatusCode == 401 || createHTTP.StatusCode == 403) {
		t.Skipf("skipping site create: insufficient API permissions (%s)", createHTTP.Status)
	}
	require.NoError(t, err)
	require.NotNil(t, createResp)
	siteID := createResp.GetSiteId()

	t.Cleanup(func() {
		if siteID == 0 {
			return
		}
		_, cleanupErr := apiClient.SitesAPI.DeleteSite(env.ctx, siteID).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup site %d: %v", siteID, cleanupErr)
		}
	})

	getResp, _, err := apiClient.SitesAPI.GetSite(env.ctx, siteID).Execute()
	require.NoError(t, err)
	if getResp != nil {
		require.Equal(t, name, getResp.GetName())
	}

	updatedName := uniqueName("go-sdk-site-updated")
	updateReq := *sites.NewCreateSiteRequest(updatedName)
	updateResp, _, err := apiClient.SitesAPI.UpdateSite(env.ctx, siteID).CreateSiteRequest(updateReq).Execute()
	require.NoError(t, err)
	if updateResp != nil && updateResp.GetName() != "" {
		require.Equal(t, updatedName, updateResp.GetName())
	}

	refetchResp, _, err := apiClient.SitesAPI.GetSite(env.ctx, siteID).Execute()
	require.NoError(t, err)
	if refetchResp != nil {
		require.Equal(t, updatedName, refetchResp.GetName())
	}

	_, err = apiClient.SitesAPI.DeleteSite(env.ctx, siteID).Execute()
	require.NoError(t, err)
	siteID = 0
}
