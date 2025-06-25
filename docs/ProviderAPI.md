# \ProviderAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProvider**](ProviderAPI.md#GetProvider) | **Get** /v1/providers/{provider_id}/ | Retrieve a provider
[**GetProviders**](ProviderAPI.md#GetProviders) | **Get** /v1/providers/ | List all providers



## GetProvider

> ProviderDetail GetProvider(ctx, providerId).Execute()

Retrieve a provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forestvpn/fvpn-api-go"
)

func main() {
	providerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProvider(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: ProviderDetail
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | A UUID string identifying this Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderDetail**](ProviderDetail.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviders

> PaginatedProviderList GetProviders(ctx).Cursor(cursor).Limit(limit).Q(q).Rating(rating).Sort(sort).VerifiedAt(verifiedAt).Execute()

List all providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/forestvpn/fvpn-api-go"
)

func main() {
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	q := "q_example" // string | A search term. (optional)
	rating := float32(3.4) // float32 |  (optional)
	sort := "sort_example" // string | Which field to use when ordering the results. (optional)
	verifiedAt := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProviders(context.Background()).Cursor(cursor).Limit(limit).Q(q).Rating(rating).Sort(sort).VerifiedAt(verifiedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviders`: PaginatedProviderList
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 
 **q** | **string** | A search term. | 
 **rating** | **float32** |  | 
 **sort** | **string** | Which field to use when ordering the results. | 
 **verifiedAt** | **time.Time** |  | 

### Return type

[**PaginatedProviderList**](PaginatedProviderList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

