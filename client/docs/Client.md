# Cisco Secure Access (SSE) OAuth2 Client


## GetHTTPClient

>  *http.Client GetHttpClient()

Retrieve a pointer to a generic SSE OAuth2 HTTP client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
http_client = clientFactory.GetHttpClient()
```

## *ntg.APIClient GetNtgClient() 

Retrieve a pointer to an SSE Network Tunnel Group API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetNtgClient()
```

## *privateapps.APIClient GetPrivateAppsClient()

Retrieve a pointer to an SSE Private Resources API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetPrivateAppsClient()
```

## *resconn.APIClient GetResConnClient()

Retrieve a pointer to an SSE Resource Connector API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetResConnClient()
```

## *rules.APIClient GetRulesClient()

Retrieve a pointer to an SSE Access Policies API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetRulesClient()
```

## *destinationlists.APIClient GetDestinationListsClient()

Retrieve a pointer to an SSE Destination Lists API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetDestinationListsClient()
```

## *internaldomains.APIClient GetInternalDomainsClient()

Retrieve a pointer to an SSE Internal Domains API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetInternalDomainsClient()
```

## *networks.APIClient GetNetworksClient()

Retrieve a pointer to an SSE Networks API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetNetworksClient()
```

## *reports.APIClient GetReportsClient()

Retrieve a pointer to an SSE Reports API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetReportsClient()
```

## *roaming.APIClient GetRoamingClient()

Retrieve a pointer to an SSE Roaming API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetRoamingClient()
```

## *swg.APIClient GetSwgClient()

Retrieve a pointer to an SSE Secure Web Gateway API client

### Example

```go

clientFactory := client.SSEClientFactory{KeyId: os.Getenv("API_KEY_ID"), KeySecret: os.Getenv("API_KEY_SECRET")}
apiClient := clientFactory.GetSwgClient()
```
