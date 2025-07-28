// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/CiscoDevNet/go-ciscosecureaccess/client"
)

func main() {

	// List existing rules
	auth := context.Background()
	factory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
	apiClient := factory.GetReportsClient(auth)
	resp, r, err := apiClient.UtilityAPI.GetIdentities(auth).Limit(100).Offset(0).Search("%qetest200@cisco.com%").Identitytypes("directory_user").Execute()
	//resp, r, err := apiClient.UtilityAPI.GetIdentities(auth).Limit(10).Offset(0).Search("qetest200").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourcesAPI.ListPrivateResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateResources`: PrivateResourceList
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetIdentiies`: %s\n", json)
}
