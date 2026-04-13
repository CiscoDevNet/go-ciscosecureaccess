// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

/*
Cisco Secure Access Networks API

Testing NetworksAPI

*/

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/CiscoDevNet/go-ciscosecureaccess/client"
	"github.com/CiscoDevNet/go-ciscosecureaccess/sites"
)

func main() {
	auth := context.Background()
	factory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
	apiClient := factory.GetSitesClient(auth)

	// List existing sites
	listResp, httpRes, err := apiClient.SitesAPI.ListSites(auth).Page(1).Limit(10).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.ListSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.ListSites`: %v\n", listResp)

	// Create a new site
	createSiteRequest := *sites.NewCreateSiteRequest("go-sdk-site-example")
	createResp, httpRes, err := apiClient.SitesAPI.CreateSite(auth).CreateSiteRequest(createSiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.CreateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.CreateSite`: %v\n", createResp)

	// Get the new site
	getResp, httpRes, err := apiClient.SitesAPI.GetSite(auth, createResp.GetSiteId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSite`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSite`: %v\n", getResp)

	// Update the new site
	resp, r, err := apiClient.SitesAPI.UpdateSite(context.Background(), createResp.GetSiteId()).CreateSiteRequest(createSiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.UpdateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSite`: SiteObject
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.UpdateSite`: %v\n", resp)

	// Delete the new site
	deleteResp, err := apiClient.SitesAPI.DeleteSite(auth, createResp.GetSiteId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.DeleteSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.DeleteSite`: %v\n", deleteResp)
}
