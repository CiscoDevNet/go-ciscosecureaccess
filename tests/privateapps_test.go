//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
	"github.com/stretchr/testify/require"
)

func TestPrivateResourcesAndGroupsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetPrivateAppsClient(env.ctx)

	_, _, err := apiClient.PrivateResourcesAPI.ListPrivateResources(env.ctx).Limit(10).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.ResourceGroupsAPI.ListResourceGroups(env.ctx).Limit(10).Execute()
	require.NoError(t, err)

	accessTypes := []privateapps.AccessTypesRequestInner{{NetworkBasedAccess: privateapps.NewNetworkBasedAccess("network")}}
	protocol := privateapps.ANY
	ports := "443"
	protocols := []privateapps.ResourceAddressesInnerProtocolPortsInner{{Protocol: &protocol, Ports: &ports}}
	addresses := []privateapps.ResourceAddressesInner{*privateapps.NewResourceAddressesInner([]string{"172.16.0.10/32"}, protocols)}
	resourceName := uniqueName("go-sdk-prires")
	resourceReq := *privateapps.NewPrivateResourceRequest(resourceName, accessTypes, addresses)

	createResourceResp, _, err := apiClient.PrivateResourcesAPI.AddPrivateResource(env.ctx).PrivateResourceRequest(resourceReq).Execute()
	require.NoError(t, err)
	resourceID := createResourceResp.GetResourceId()

	t.Cleanup(func() {
		if resourceID == 0 {
			return
		}
		_, _, cleanupErr := apiClient.PrivateResourcesAPI.DeletePrivateResource(env.ctx, resourceID).Force(true).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup private resource %d: %v", resourceID, cleanupErr)
		}
	})

	groupName := uniqueName("go-sdk-resgroup")
	groupReq := *privateapps.NewPrivateResourceGroupRequest(groupName, []int64{resourceID})
	groupResp, _, err := apiClient.ResourceGroupsAPI.AddPrivateResourceGroup(env.ctx).PrivateResourceGroupRequest(groupReq).Execute()
	require.NoError(t, err)
	groupID := groupResp.GetResourceGroupId()

	t.Cleanup(func() {
		if groupID == 0 {
			return
		}
		_, _, cleanupErr := apiClient.ResourceGroupsAPI.DeletePrivateResourceGroup(env.ctx, groupID).Force(true).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup private resource group %d: %v", groupID, cleanupErr)
		}
	})

	getResourceResp, _, err := apiClient.PrivateResourcesAPI.GetPrivateResource(env.ctx, resourceID).Execute()
	require.NoError(t, err)
	require.Equal(t, resourceID, getResourceResp.GetResourceId())

	getGroupResp, _, err := apiClient.ResourceGroupsAPI.GetPrivateResourceGroup(env.ctx, groupID).Execute()
	require.NoError(t, err)
	if getGroupResp != nil {
		require.Equal(t, groupID, getGroupResp.GetResourceGroupId())
	}

	updatedResourceReq := resourceReq
	updatedResourceReq.SetName(uniqueName("go-sdk-prires-upd"))
	updatedAddresses := []privateapps.ResourceAddressesInner{*privateapps.NewResourceAddressesInner([]string{"172.16.0.11/32"}, protocols)}
	updatedResourceReq.SetResourceAddresses(updatedAddresses)
	updatedResourceResp, _, err := apiClient.PrivateResourcesAPI.PutPrivateResource(env.ctx, resourceID).PrivateResourceRequest(updatedResourceReq).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedResourceReq.GetName(), updatedResourceResp.GetName())

	updatedGroupReq := groupReq
	updatedGroupReq.SetName(uniqueName("go-sdk-resgrp-upd"))
	updatedGroupResp, _, err := apiClient.ResourceGroupsAPI.PutPrivateResourceGroup(env.ctx, groupID).PrivateResourceGroupRequest(updatedGroupReq).Execute()
	require.NoError(t, err)
	if updatedGroupResp != nil {
		require.Equal(t, updatedGroupReq.GetName(), updatedGroupResp.GetName())
	}

	_, _, err = apiClient.PrivateResourcesAPI.DeletePrivateResource(env.ctx, resourceID).Force(true).Execute()
	require.NoError(t, err)
	resourceID = 0

	_, _, err = apiClient.ResourceGroupsAPI.DeletePrivateResourceGroup(env.ctx, groupID).Force(true).Execute()
	require.NoError(t, err)
	groupID = 0
}
