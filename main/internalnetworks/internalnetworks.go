// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

/*
Cisco Secure Access Internal Networks API

Example usage of the InternalNetworksAPI

*/

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/CiscoDevNet/go-ciscosecureaccess/client"
	"github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
)

func main() {
	auth := context.Background()
	factory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
	apiClient := factory.GetInternalNetworksClient(auth)

	// List existing internal networks
	listResp, httpRes, err := apiClient.InternalNetworksAPI.ListInternalNetworks(auth).Page(1).Limit(10).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.ListInternalNetworks`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.ListInternalNetworks`: %v\n", listResp)

	// Create a new internal network
	createRequest := *internalnetworks.NewCreateInternalNetworkRequest("go-sdk-internal-network-example", "198.51.100.0", 24)
	siteId := int64(1)
	createRequest.SetSiteId(siteId)
	createResp, httpRes, err := apiClient.InternalNetworksAPI.CreateInternalNetwork(auth).CreateInternalNetworkRequest(createRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.CreateInternalNetwork`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.CreateInternalNetwork`: %v\n", createResp)

	// Get the new internal network
	getResp, httpRes, err := apiClient.InternalNetworksAPI.GetInternalNetwork(auth, createResp.GetOriginId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.GetInternalNetwork`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.GetInternalNetwork`: %v\n", getResp)

	// Update the internal network
	createRequest.SetSiteId(siteId)
	updateResp, httpRes, err := apiClient.InternalNetworksAPI.UpdateInternalNetwork(auth, createResp.GetOriginId()).CreateInternalNetworkRequest(createRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.UpdateInternalNetwork`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.UpdateInternalNetwork`: %v\n", updateResp)

	// Delete the internal network
	httpRes, err = apiClient.InternalNetworksAPI.DeleteInternalNetwork(auth, createResp.GetOriginId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.DeleteInternalNetwork`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Internal network %d deleted successfully\n", createResp.GetOriginId())
}
