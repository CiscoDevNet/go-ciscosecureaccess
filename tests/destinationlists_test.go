//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/destinationlists"
	"github.com/stretchr/testify/require"
)

func TestDestinationListsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetDestinationListsClient(env.ctx)

	_, listHTTP, err := apiClient.DestinationListsAPI.GetDestinationLists(env.ctx).Page(1).Limit(10).Execute()
	if err != nil && !isHTTP2xx(listHTTP) {
		require.NoError(t, err)
	}

	name := uniqueName("go-sdk-destination-list")
	comment := "functional-test"
	destinationType := destinationlists.DOMAIN
	initialDestination := name + ".example.com"
	bundleType := destinationlists.BundleTypeId(2)
	createReq := destinationlists.DestinationListCreate{
		Access:       "none",
		IsGlobal:     false,
		Name:         name,
		BundleTypeId: &bundleType,
		Destinations: []destinationlists.DestinationListCreateDestinationsInner{{
			Comment:     &comment,
			Type:        &destinationType,
			Destination: &initialDestination,
		}},
	}

	createResp, _, err := apiClient.DestinationListsAPI.CreateDestinationList(env.ctx).DestinationListCreate(createReq).Execute()
	require.NoError(t, err)
	listID := createResp.Data.Id

	t.Cleanup(func() {
		if listID == 0 {
			return
		}
		_, _, cleanupErr := apiClient.DestinationListsAPI.DeleteDestinationList(env.ctx, listID).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup destination list %d: %v", listID, cleanupErr)
		}
	})

	getResp, _, err := apiClient.DestinationListsAPI.GetDestinationList(env.ctx, listID).Execute()
	require.NoError(t, err)
	require.Equal(t, name, getResp.Data.Name)

	updatedName := uniqueName("go-sdk-destination-list-updated")
	_, _, err = apiClient.DestinationListsAPI.UpdateDestinationLists(env.ctx, listID).
		DestinationListPatch(destinationlists.DestinationListPatch{Name: updatedName}).
		Execute()
	require.NoError(t, err)

	newDestinations := []destinationlists.DestinationCreateObject{
		*destinationlists.NewDestinationCreateObject("www." + updatedName + ".example.com"),
	}
	_, _, err = apiClient.DestinationsAPI.CreateDestinations(env.ctx, listID).DestinationCreateObject(newDestinations).Execute()
	require.NoError(t, err)

	destinationsResp, _, err := apiClient.DestinationsAPI.GetDestinations(env.ctx, listID).Page(1).Limit(100).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, destinationsResp.Data)

	updatedResp, _, err := apiClient.DestinationListsAPI.GetDestinationList(env.ctx, listID).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedName, updatedResp.Data.Name)

	_, _, err = apiClient.DestinationListsAPI.DeleteDestinationList(env.ctx, listID).Execute()
	require.NoError(t, err)
	listID = 0
	_ = listID
}
