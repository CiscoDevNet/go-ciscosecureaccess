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
