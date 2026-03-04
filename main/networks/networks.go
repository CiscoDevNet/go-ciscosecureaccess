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
	"github.com/CiscoDevNet/go-ciscosecureaccess/networks"
)

func main() {
	auth := context.Background()
	factory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
	apiClient := factory.GetNetworksClient(auth)

	// List existing networks
	listResp, httpRes, err := apiClient.NetworksAPI.ListNetworks(auth).Page(1).Limit(10).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworks`: %v\n", listResp)

	// Create a new network
	createNetworkRequest := *networks.NewCreateNetworkRequest("go-sdk-network-example", 32, false, "OPEN")
	createNetworkRequest.SetIpAddress("198.51.100.11")
	createResp, httpRes, err := apiClient.NetworksAPI.CreateNetwork(auth).CreateNetworkRequest(createNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetwork`: %v\n", createResp)

	// Get the new network
	getResp, httpRes, err := apiClient.NetworksAPI.GetNetwork(auth, createResp.GetOriginId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetwork`: %v\n", getResp)

	// Update the new network
	updateNetworkRequest := *networks.NewUpdateNetworkRequest("go-sdk-network-example-updated", true, "OPEN")
	updateNetworkRequest.SetIpAddress("198.51.100.11")
	updateNetworkRequest.SetPrefixLength(32)
	updateResp, httpRes, err := apiClient.NetworksAPI.UpdateNetwork(auth, createResp.GetOriginId()).UpdateNetworkRequest(updateNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetwork`: %v\n", updateResp)

	// Delete the new network
	deleteResp, httpRes, err := apiClient.NetworksAPI.DeleteNetwork(auth, createResp.GetOriginId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
	}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetwork`: %v\n", deleteResp)
}
