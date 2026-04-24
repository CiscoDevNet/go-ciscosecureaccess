//go:build functional

package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
	"github.com/stretchr/testify/require"
)

func TestConnectorGroupsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetResConnClient(env.ctx)

	groupsResp, _, err := apiClient.ConnectorGroupsAPI.ListConnectorGroups(env.ctx).Limit(10).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.ConnectorGroupsAPI.GetConnectorGroupCount(env.ctx).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.ConnectorsAPI.ListConnectors(env.ctx).Limit(10).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.ConnectorsAPI.GetConnectorAgentCount(env.ctx).Execute()
	require.NoError(t, err)

	location := strings.TrimSpace(os.Getenv("GO_CISCOSECUREACCESS_CONNECTOR_GROUP_LOCATION"))
	if location == "" && len(groupsResp.GetData()) > 0 {
		location = groupsResp.GetData()[0].GetLocation()
	}
	if location == "" {
		location = "us-east-1"
	}

	createReq := *resconn.NewConnectorGroupReq(uniqueName("go-sdk-connector-group"), location)
	createResp, _, err := apiClient.ConnectorGroupsAPI.CreateConnectorGroup(env.ctx).ConnectorGroupReq(createReq).Execute()
	require.NoError(t, err)
	id := createResp.GetId()

	t.Cleanup(func() {
		if id == 0 {
			return
		}
		_, _, cleanupErr := apiClient.ConnectorGroupsAPI.DeleteConnectorGroup(env.ctx, id).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup connector group %d: %v", id, cleanupErr)
		}
	})

	getResp, _, err := apiClient.ConnectorGroupsAPI.GetConnectorGroup(env.ctx, id).IncludeProvisioningKey(true).Execute()
	require.NoError(t, err)
	require.Equal(t, id, getResp.GetId())

	updatedName := uniqueName("go-sdk-connector-group-updated")
	patchReq := []resconn.ConnectorGroupPatchReqInner{
		*resconn.NewConnectorGroupPatchReqInner(resconn.REPLACE, "name", updatedName),
	}
	patchResp, patchHTTP, err := apiClient.ConnectorGroupsAPI.PatchConnectorGroup(env.ctx, id).ConnectorGroupPatchReqInner(patchReq).Execute()
	if err != nil && patchHTTP != nil && patchHTTP.StatusCode == 400 {
		patchReq = []resconn.ConnectorGroupPatchReqInner{
			*resconn.NewConnectorGroupPatchReqInner(resconn.REPLACE, "/name", updatedName),
		}
		patchResp, patchHTTP, err = apiClient.ConnectorGroupsAPI.PatchConnectorGroup(env.ctx, id).ConnectorGroupPatchReqInner(patchReq).Execute()
	}
	require.NoError(t, err)
	require.Equal(t, updatedName, patchResp.GetName())

	_, _, err = apiClient.ConnectorGroupsAPI.DeleteConnectorGroup(env.ctx, id).Execute()
	require.NoError(t, err)
	id = 0
}
