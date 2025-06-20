# \DeviceAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DeviceAPI.md#CreateDevice) | **Post** /v1/devices/ | Create a new account
[**DeleteDevice**](DeviceAPI.md#DeleteDevice) | **Delete** /v1/devices/{device_id}/ | Delete a device
[**GetDashboardAnalytics**](DeviceAPI.md#GetDashboardAnalytics) | **Get** /v1/devices/analytics/dashboard/ | Get device data usage analytics for dashboard
[**GetDevice**](DeviceAPI.md#GetDevice) | **Get** /v1/devices/{device_id}/ | Retrieve a device
[**GetDeviceAnalytics**](DeviceAPI.md#GetDeviceAnalytics) | **Get** /v1/devices/{device_id}/analytics/ | Get data usage analytics for a specific device
[**GetDeviceNode**](DeviceAPI.md#GetDeviceNode) | **Get** /v1/devices/{device_id}/node/ | Get the node of a device
[**GetDevices**](DeviceAPI.md#GetDevices) | **Get** /v1/devices/ | List all devices
[**PartialUpdateDevice**](DeviceAPI.md#PartialUpdateDevice) | **Patch** /v1/devices/{device_id}/ | Partially update a device
[**SetDeviceNode**](DeviceAPI.md#SetDeviceNode) | **Post** /v1/devices/{device_id}/set-node/ | Set a node on a device
[**UpdateDevice**](DeviceAPI.md#UpdateDevice) | **Put** /v1/devices/{device_id}/ | Update a device



## CreateDevice

> Device CreateDevice(ctx).DeviceRequest(deviceRequest).Execute()

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
	deviceRequest := *openapiclient.NewDeviceRequest("PubKey_example") // DeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.CreateDevice(context.Background()).DeviceRequest(deviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.CreateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDevice`: Device
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceRequest** | [**DeviceRequest**](DeviceRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, deviceId).Execute()

Delete a device

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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceAPI.DeleteDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## GetDashboardAnalytics

> DashboardAnalytics GetDashboardAnalytics(ctx).Top(top).Execute()

Get device data usage analytics for dashboard



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
	top := int32(56) // int32 | Number of top devices to include (default: 5) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDashboardAnalytics(context.Background()).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDashboardAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardAnalytics`: DashboardAnalytics
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDashboardAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Number of top devices to include (default: 5) | 

### Return type

[**DashboardAnalytics**](DashboardAnalytics.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> Device GetDevice(ctx, deviceId).Execute()

Retrieve a device

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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevice`: Device
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Device**](Device.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAnalytics

> DeviceAnalytics GetDeviceAnalytics(ctx, deviceId).ComparePrevious(comparePrevious).EndDate(endDate).StartDate(startDate).Execute()

Get data usage analytics for a specific device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.
	comparePrevious := true // bool | Whether to include comparison with previous period (true/false) (optional)
	endDate := "endDate_example" // string | End date in ISO format (default: current date and time) (optional)
	startDate := "startDate_example" // string | Start date in ISO format (default: 7 days ago) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDeviceAnalytics(context.Background(), deviceId).ComparePrevious(comparePrevious).EndDate(endDate).StartDate(startDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDeviceAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAnalytics`: DeviceAnalytics
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDeviceAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **comparePrevious** | **bool** | Whether to include comparison with previous period (true/false) | 
 **endDate** | **string** | End date in ISO format (default: current date and time) | 
 **startDate** | **string** | Start date in ISO format (default: 7 days ago) | 

### Return type

[**DeviceAnalytics**](DeviceAnalytics.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceNode

> Node GetDeviceNode(ctx, deviceId).Execute()

Get the node of a device

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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDeviceNode(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDeviceNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceNode`: Node
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDeviceNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> PaginatedDeviceList GetDevices(ctx).Cursor(cursor).Limit(limit).Execute()

List all devices

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
	resp, r, err := apiClient.DeviceAPI.GetDevices(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: PaginatedDeviceList
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedDeviceList**](PaginatedDeviceList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDevice

> Device PartialUpdateDevice(ctx, deviceId).PatchedDeviceRequest(patchedDeviceRequest).Execute()

Partially update a device

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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.
	patchedDeviceRequest := *openapiclient.NewPatchedDeviceRequest() // PatchedDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.PartialUpdateDevice(context.Background(), deviceId).PatchedDeviceRequest(patchedDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.PartialUpdateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDevice`: Device
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.PartialUpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDeviceRequest** | [**PatchedDeviceRequest**](PatchedDeviceRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceNode

> Node SetDeviceNode(ctx, deviceId).SetNodeRequest(setNodeRequest).Execute()

Set a node on a device



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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.
	setNodeRequest := *openapiclient.NewSetNodeRequest() // SetNodeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.SetDeviceNode(context.Background(), deviceId).SetNodeRequest(setNodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.SetDeviceNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDeviceNode`: Node
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.SetDeviceNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNodeRequest** | [**SetNodeRequest**](SetNodeRequest.md) |  | 

### Return type

[**Node**](Node.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, deviceId).DeviceRequest(deviceRequest).Execute()

Update a device

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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Device.
	deviceRequest := *openapiclient.NewDeviceRequest("PubKey_example") // DeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.UpdateDevice(context.Background(), deviceId).DeviceRequest(deviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UpdateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevice`: Device
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | A UUID string identifying this Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceRequest** | [**DeviceRequest**](DeviceRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

