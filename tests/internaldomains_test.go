//go:build functional

package tests

import (
	"testing"

	"github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
	"github.com/stretchr/testify/require"
)

func TestInternalDomainsFunctionalCRUD(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetInternalDomainsClient(env.ctx)

	_, _, err := apiClient.InternalDomainsAPI.ListInternalDomains(env.ctx).Page(1).Limit(10).Execute()
	require.NoError(t, err)

	domain := uniqueName("sdk-functional-domain") + ".example.com"
	createReq := *internaldomains.NewCreateInternalDomainRequest(domain)
	createResp, _, err := apiClient.InternalDomainsAPI.CreateInternalDomain(env.ctx).CreateInternalDomainRequest(createReq).Execute()
	require.NoError(t, err)
	id := createResp.GetId()

	t.Cleanup(func() {
		if id == 0 {
			return
		}
		_, _, cleanupErr := apiClient.InternalDomainsAPI.DeleteInternalDomain(env.ctx, id).Execute()
		if cleanupErr != nil {
			t.Logf("cleanup internal domain %d: %v", id, cleanupErr)
		}
	})

	getResp, _, err := apiClient.InternalDomainsAPI.GetInternalDomain(env.ctx, id).Execute()
	require.NoError(t, err)
	require.Equal(t, domain, getResp.GetDomain())

	updatedDomain := uniqueName("sdk-functional-domain-updated") + ".example.com"
	updateReq := *internaldomains.NewCreateInternalDomainRequest(updatedDomain)
	updateResp, _, err := apiClient.InternalDomainsAPI.UpdateInternalDomain(env.ctx, id).CreateInternalDomainRequest(updateReq).Execute()
	require.NoError(t, err)
	require.Equal(t, updatedDomain, updateResp.GetDomain())

	_, _, err = apiClient.InternalDomainsAPI.DeleteInternalDomain(env.ctx, id).Execute()
	require.NoError(t, err)
	id = 0
}
