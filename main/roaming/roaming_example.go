// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

/*
Cisco Secure Access Roaming Computers API

Testing RoamingComputersAPIService
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/CiscoDevNet/go-ciscosecureaccess/client"
)

func main() {
	ctx := context.Background()
	factory := client.SSEClientFactory{
		KeyId:     os.Getenv("API_KEY_ID"),
		KeySecret: os.Getenv("API_KEY_SECRET"),
	}
	apiClient := factory.GetRoamingClient(ctx)

	// List roaming computers
	fmt.Println("=== Listing Roaming Computers ===")
	computers, httpRes, err := apiClient.RoamingComputersAPI.ListRoamingComputers(ctx).
		Limit(10).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing roaming computers: %s\n", err)
		if httpRes != nil {
			fmt.Fprintf(os.Stderr, "HTTP Response: %s\n", httpRes.Status)
		}
		os.Exit(1)
	}

	jsonResp, _ := json.MarshalIndent(computers, "", "  ")
	fmt.Printf("Roaming Computers: %s\n", jsonResp)
	fmt.Printf("Total computers found: %d\n", len(computers))

	// If we have at least one computer, get its details
	if len(computers) > 0 {
		deviceId := computers[0].DeviceId
		fmt.Printf("\n=== Getting Roaming Computer Details (deviceId: %s) ===\n", deviceId)

		computer, httpRes, err := apiClient.RoamingComputersAPI.GetRoamingComputer(ctx, deviceId).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting roaming computer: %s\n", err)
			if httpRes != nil {
				fmt.Fprintf(os.Stderr, "HTTP Response: %s\n", httpRes.Status)
			}
		} else {
			jsonResp, _ := json.MarshalIndent(computer, "", "  ")
			fmt.Printf("Roaming Computer Details: %s\n", jsonResp)
		}
	}

	// Get organization info
	fmt.Println("\n=== Getting Organization Info ===")
	orgInfo, httpRes, err := apiClient.OrganizationInformationAPI.GetOrganizationInfo(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting org info: %s\n", err)
		if httpRes != nil {
			fmt.Fprintf(os.Stderr, "HTTP Response: %s\n", httpRes.Status)
		}
	} else {
		jsonResp, _ := json.MarshalIndent(orgInfo, "", "  ")
		fmt.Printf("Organization Info: %s\n", jsonResp)
	}

	fmt.Println("\n=== Test Complete ===")
}
