// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

/*
Cisco Secure Access SWG Device Settings API

Testing SWGDeviceSettingsAPIService
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/CiscoDevNet/go-ciscosecureaccess/client"
	"github.com/CiscoDevNet/go-ciscosecureaccess/swg"
)

func main() {
	ctx := context.Background()
	factory := client.SSEClientFactory{
		KeyId:     os.Getenv("API_KEY_ID"),
		KeySecret: os.Getenv("API_KEY_SECRET"),
	}
	apiClient := factory.GetSwgClient(ctx)

	// List SWG device settings for specific origin IDs
	// Note: You need valid origin IDs from your roaming computers
	fmt.Println("=== Listing SWG Device Settings ===")

	// Example origin IDs - replace with actual ones from your roaming computers
	originIds := []int64{632556530, 632605193} // These are from the roaming computers test

	listRequest := swg.ListSecureWebGatewayDeviceSettingsRequest{
		OriginIds: originIds,
	}

	settings, httpRes, err := apiClient.SWGDeviceSettingsAPI.ListSecureWebGatewayDeviceSettings(ctx).
		ListSecureWebGatewayDeviceSettingsRequest(listRequest).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing SWG device settings: %s\n", err)
		if httpRes != nil {
			fmt.Fprintf(os.Stderr, "HTTP Response: %s\n", httpRes.Status)
		}
	} else {
		jsonResp, _ := json.MarshalIndent(settings, "", "  ")
		fmt.Printf("SWG Device Settings: %s\n", jsonResp)
	}

	fmt.Println("\n=== Test Complete ===")
}
