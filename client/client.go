// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/CiscoDevNet/go-ciscosecureaccess/destinationlists"
	"github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
	"github.com/CiscoDevNet/go-ciscosecureaccess/networks"
	"github.com/CiscoDevNet/go-ciscosecureaccess/ntg"
	"github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
	"github.com/CiscoDevNet/go-ciscosecureaccess/reports"
	"github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
	"github.com/CiscoDevNet/go-ciscosecureaccess/rules"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const ciscoSseApiEndpoint = "api.sse.cisco.com"

//const ciscoSseTokenUri = fmt.Sprintf("https://%s/auth/v2/token", ciscoSseApiEndpoint)

type SSEClientFactory struct {
	KeyId         string
	KeySecret     string
	ApiEndpoint   string
	SSEHttpClient *http.Client
	mu            sync.Mutex
}

func (c *SSEClientFactory) GetHttpClient(ctx context.Context) *http.Client {

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ApiEndpoint == "" {
		c.ApiEndpoint = ciscoSseApiEndpoint
	}
	if c.SSEHttpClient == nil {
		// Create a retryable HTTP client for token requests
		tokenRetryClient := retryablehttp.NewClient()
		tokenRetryClient.RetryMax = 10

		sSEAuthconfig := &clientcredentials.Config{
			ClientID:     c.KeyId,
			ClientSecret: c.KeySecret,
			TokenURL:     fmt.Sprintf("https://%s/auth/v2/token", c.ApiEndpoint)}

		// Set the HTTP client for OAuth2 to use our retryable client
		ctx = context.WithValue(ctx, oauth2.HTTPClient, tokenRetryClient.StandardClient())

		initial, _ := sSEAuthconfig.TokenSource(ctx).Token()
		oauthHttpClient := oauth2.NewClient(ctx, oauth2.ReuseTokenSource(initial, sSEAuthconfig.TokenSource(ctx)))

		// Create another retryable client for API requests
		retryHttpClient := retryablehttp.NewClient()
		retryHttpClient.HTTPClient = oauthHttpClient
		retryHttpClient.RetryMax = 10

		c.SSEHttpClient = retryHttpClient.StandardClient()
	}
	return c.SSEHttpClient
}
func (c *SSEClientFactory) GetURLString(suffix string) string {
	var hostname string
	if c.ApiEndpoint == "" {
		hostname = ciscoSseApiEndpoint
	} else {
		hostname = c.ApiEndpoint
	}

	return fmt.Sprintf("https://%s/%s", hostname, suffix)

}

func (c *SSEClientFactory) GetDestinationListsClient(ctx context.Context) *destinationlists.APIClient {
	configuration := destinationlists.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return destinationlists.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetNtgClient(ctx context.Context) *ntg.APIClient {
	configuration := ntg.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return ntg.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetResConnClient(ctx context.Context) *resconn.APIClient {
	configuration := resconn.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return resconn.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetRulesClient(ctx context.Context) *rules.APIClient {
	configuration := rules.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return rules.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetReportsClient(ctx context.Context) *reports.APIClient {
	configuration := reports.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("")
	return reports.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetPrivateAppsClient(ctx context.Context) *privateapps.APIClient {
	configuration := privateapps.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return privateapps.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetNetworksClient(ctx context.Context) *networks.APIClient {
	configuration := networks.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return networks.NewAPIClient(configuration)
}

func (c *SSEClientFactory) GetInternalDomainsClient(ctx context.Context) *internaldomains.APIClient {
	configuration := internaldomains.NewConfiguration()
	configuration.HTTPClient = c.GetHttpClient(ctx)
	configuration.Servers[0].URL = c.GetURLString("{basePath}")
	return internaldomains.NewAPIClient(configuration)
}
