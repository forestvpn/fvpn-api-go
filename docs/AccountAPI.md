# \AccountAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveDeviceCode**](AccountAPI.md#ApproveDeviceCode) | **Post** /v1/account/codes/{code}/approve/ | Approve a device authentication code
[**CreateAccount**](AccountAPI.md#CreateAccount) | **Post** /v1/accounts/ | Create a new account
[**CreateDeviceCode**](AccountAPI.md#CreateDeviceCode) | **Post** /v1/account/codes/ | Create a device authentication code
[**DeleteCurrentAccount**](AccountAPI.md#DeleteCurrentAccount) | **Delete** /v1/account/ | Delete the current account
[**GetAccountTokenPair**](AccountAPI.md#GetAccountTokenPair) | **Post** /v1/account/token/ | Get account token pair
[**GetCurrentAccount**](AccountAPI.md#GetCurrentAccount) | **Get** /v1/account/ | Retrieve the current account
[**GetCurrentAccountBillingUsage**](AccountAPI.md#GetCurrentAccountBillingUsage) | **Get** /v1/account/billing-usage/ | Retrieve current account billing usage
[**GetDeviceCode**](AccountAPI.md#GetDeviceCode) | **Get** /v1/account/codes/{code}/ | Check device code
[**RefreshAccountToken**](AccountAPI.md#RefreshAccountToken) | **Post** /v1/account/token/refresh/ | Refresh account token



## ApproveDeviceCode

> ApproveDeviceCode(ctx, code).Execute()

Approve a device authentication code



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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.ApproveDeviceCode(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ApproveDeviceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveDeviceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> Account CreateAccount(ctx).Execute()

Create a new account

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceCode

> DeviceCode CreateDeviceCode(ctx).DeviceCodeRequest(deviceCodeRequest).Execute()

Create a device authentication code



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
	deviceCodeRequest := *openapiclient.NewDeviceCodeRequest() // DeviceCodeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateDeviceCode(context.Background()).DeviceCodeRequest(deviceCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateDeviceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceCode`: DeviceCode
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateDeviceCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCodeRequest** | [**DeviceCodeRequest**](DeviceCodeRequest.md) |  | 

### Return type

[**DeviceCode**](DeviceCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCurrentAccount

> DeleteCurrentAccount(ctx).Last4(last4).Execute()

Delete the current account

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
	last4 := "last4_example" // string | Last 4 digits of the account number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.DeleteCurrentAccount(context.Background()).Last4(last4).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteCurrentAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurrentAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **last4** | **string** | Last 4 digits of the account number | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTokenPair

> AccountTokenPair GetAccountTokenPair(ctx).AccountTokenPairRequest(accountTokenPairRequest).Execute()

Get account token pair



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
	accountTokenPairRequest := *openapiclient.NewAccountTokenPairRequest("Number_example") // AccountTokenPairRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountTokenPair(context.Background()).AccountTokenPairRequest(accountTokenPairRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountTokenPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTokenPair`: AccountTokenPair
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountTokenPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTokenPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountTokenPairRequest** | [**AccountTokenPairRequest**](AccountTokenPairRequest.md) |  | 

### Return type

[**AccountTokenPair**](AccountTokenPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentAccount

> Account GetCurrentAccount(ctx).Cursor(cursor).Limit(limit).Execute()

Retrieve the current account

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
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetCurrentAccount(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetCurrentAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetCurrentAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentAccountBillingUsage

> BillingUsage GetCurrentAccountBillingUsage(ctx).Cursor(cursor).Limit(limit).Execute()

Retrieve current account billing usage

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
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetCurrentAccountBillingUsage(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetCurrentAccountBillingUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentAccountBillingUsage`: BillingUsage
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetCurrentAccountBillingUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentAccountBillingUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**BillingUsage**](BillingUsage.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCode

> DeviceCode GetDeviceCode(ctx, code).Execute()

Check device code



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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetDeviceCode(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetDeviceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCode`: DeviceCode
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetDeviceCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceCode**](DeviceCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAccountToken

> AccountTokenRefresh RefreshAccountToken(ctx).AccountTokenRefreshRequest(accountTokenRefreshRequest).Execute()

Refresh account token



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
	accountTokenRefreshRequest := *openapiclient.NewAccountTokenRefreshRequest() // AccountTokenRefreshRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RefreshAccountToken(context.Background()).AccountTokenRefreshRequest(accountTokenRefreshRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RefreshAccountToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshAccountToken`: AccountTokenRefresh
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RefreshAccountToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountTokenRefreshRequest** | [**AccountTokenRefreshRequest**](AccountTokenRefreshRequest.md) |  | 

### Return type

[**AccountTokenRefresh**](AccountTokenRefresh.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

