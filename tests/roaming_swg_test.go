//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
	"github.com/CiscoDevNet/go-ciscosecureaccess/swg"
	"github.com/stretchr/testify/require"
)

func TestRoamingFunctionalCoverage(t *testing.T) {
	t.Skip()
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetRoamingClient(env.ctx)

	computers, _, err := apiClient.RoamingComputersAPI.ListRoamingComputers(env.ctx).Limit(10).Execute()
	require.NoError(t, err)
	_, _, err = apiClient.OrganizationInformationAPI.GetOrganizationInfo(env.ctx).Execute()
	require.NoError(t, err)

	if len(computers) == 0 {
		t.Skip("no roaming computers available for get/update coverage")
	}

	deviceID := computers[0].GetDeviceId()
	currentName := computers[0].GetName()
	getResp, _, err := apiClient.RoamingComputersAPI.GetRoamingComputer(env.ctx, deviceID).Execute()
	require.NoError(t, err)
	require.Equal(t, deviceID, getResp.GetDeviceId())

	updatedName := uniqueName("go-sdk-roaming")
	updateReq := *roaming.NewUpdateRoamingComputerRequest(updatedName)
	updatedResp, _, err := apiClient.RoamingComputersAPI.UpdateRoamingComputer(env.ctx, deviceID).UpdateRoamingComputerRequest(updateReq).Execute()
	require.NoError(t, err)
	if updatedResp != nil && updatedResp.GetName() != "" {
		require.Equal(t, updatedName, updatedResp.GetName())
	}

	refetched, _, err := apiClient.RoamingComputersAPI.GetRoamingComputer(env.ctx, deviceID).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedName, refetched.GetName())

	_, _, err = apiClient.RoamingComputersAPI.UpdateRoamingComputer(env.ctx, deviceID).
		UpdateRoamingComputerRequest(*roaming.NewUpdateRoamingComputerRequest(currentName)).
		Execute()
	require.NoError(t, err)
}

func TestSWGDeviceSettingsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	roamingClient := env.factory.GetRoamingClient(env.ctx)
	computers, _, err := roamingClient.RoamingComputersAPI.ListRoamingComputers(env.ctx).Limit(1).Execute()
	require.NoError(t, err)
	if len(computers) == 0 {
		t.Skip("no roaming computers available for SWG device settings coverage")
	}

	originID := computers[0].GetOriginId()
	apiClient := env.factory.GetSwgClient(env.ctx)
	listReq := *swg.NewListSecureWebGatewayDeviceSettingsRequest([]int64{originID})
	value, err := swg.NewValueFromValue("1")
	require.NoError(t, err)

	createReq := *swg.NewCreateSecureWebGatewayDeviceSettingsRequest([]int64{originID}, *value)
	createResp, _, err := apiClient.SWGDeviceSettingsAPI.CreateSecureWebGatewayDeviceSettings(env.ctx).CreateSecureWebGatewayDeviceSettingsRequest(createReq).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, createResp.GetItems())

	t.Cleanup(func() {
		_, _, cleanupErr := apiClient.SWGDeviceSettingsAPI.DeleteSecureWebGatewayDeviceSettings(env.ctx).ListSecureWebGatewayDeviceSettingsRequest(listReq).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup swg override for origin %d: %v", originID, cleanupErr)
		}
	})

	listResp, _, err := apiClient.SWGDeviceSettingsAPI.ListSecureWebGatewayDeviceSettings(env.ctx).ListSecureWebGatewayDeviceSettingsRequest(listReq).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, listResp)

	_, _, err = apiClient.SWGDeviceSettingsAPI.DeleteSecureWebGatewayDeviceSettings(env.ctx).ListSecureWebGatewayDeviceSettingsRequest(listReq).Execute()
	require.NoError(t, err)
}
