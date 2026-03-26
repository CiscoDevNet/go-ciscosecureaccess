// Copyright 2025 Cisco Systems, Inc. and its affiliates
//
// SPDX-License-Identifier: Apache-2.0

/*
Cisco Secure Access Internal Domains API

Testing InternalDomainsAPIService
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
	apiClient := factory.GetInternalDomainsClient(ctx)

	// List internal domains
	fmt.Println("=== Listing Internal Domains ===")
	domains, httpRes, err := apiClient.InternalDomainsAPI.ListInternalDomains(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing internal domains: %s\n", err)
		if httpRes != nil {
			fmt.Fprintf(os.Stderr, "HTTP Response: %s\n", httpRes.Status)
		}
	} else {
		jsonResp, _ := json.MarshalIndent(domains, "", "  ")
		fmt.Printf("Internal Domains: %s\n", jsonResp)
	}

	fmt.Println("\n=== Test Complete ===")
}
