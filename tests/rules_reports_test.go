//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/rules"
	"github.com/stretchr/testify/require"
)

func TestReportsFunctionalReadOnly(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetReportsClient(env.ctx)

	_, _, err := apiClient.UtilityAPI.GetCategories(env.ctx).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.UtilityAPI.GetIdentities(env.ctx).Limit(10).Offset(0).Execute()
	require.NoError(t, err)
}

func TestApplicationListsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetRulesClient(env.ctx)

	_, _, err := apiClient.ApplicationListsAPI.GetApplicationLists(env.ctx).Execute()
	require.NoError(t, err)

	appID, ok, err := envInt64("GO_CISCOSECUREACCESS_RULE_APPLICATION_ID")
	require.NoError(t, err)
	if !ok {
		appID = 64
	}

	name := uniqueName("go-sdk-applist")
	createReq := *rules.NewApplicationListRequest(name, false, []int64{appID})
	_, _, err = apiClient.ApplicationListsAPI.CreateApplicationList(env.ctx).ApplicationListRequest(createReq).Execute()
	require.NoError(t, err)

	listResp, _, err := apiClient.ApplicationListsAPI.GetApplicationLists(env.ctx).Execute()
	require.NoError(t, err)

	var applicationListID int64
	if rawResults, ok := listResp.AdditionalProperties["results"]; ok {
		if resultsSlice, ok := rawResults.([]interface{}); ok {
			for _, item := range resultsSlice {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if itemName, ok := itemMap["applicationListName"].(string); ok && itemName == name {
						if idVal, ok := itemMap["applicationListId"]; ok {
							if id, ok := idVal.(float64); ok {
								applicationListID = int64(id)
							}
						}
						break
					}
				}
			}
		}
	}
	require.NotZero(t, applicationListID)

	t.Cleanup(func() {
		if applicationListID == 0 {
			return
		}
		_, _, cleanupErr := apiClient.ApplicationListsAPI.DeleteApplicationList(env.ctx, applicationListID).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup application list %d: %v", applicationListID, cleanupErr)
		}
	})

	getResp, _, err := apiClient.ApplicationListsAPI.GetApplicationList(env.ctx, applicationListID).Execute()
	require.NoError(t, err)
	require.Equal(t, name, getResp.GetApplicationListName())

	updatedName := uniqueName("go-sdk-application-list-updated")
	updateReq := *rules.NewApplicationListRequest(updatedName, false, []int64{appID})
	updateResp, _, err := apiClient.ApplicationListsAPI.PutApplicationList(env.ctx, applicationListID).ApplicationListRequest(updateReq).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedName, updateResp.GetApplicationListName())

	_, _, err = apiClient.ApplicationListsAPI.DeleteApplicationList(env.ctx, applicationListID).Execute()
	require.NoError(t, err)
	applicationListID = 0
}
