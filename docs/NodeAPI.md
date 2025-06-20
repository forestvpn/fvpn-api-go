# \NodeAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNode**](NodeAPI.md#CreateNode) | **Post** /v1/nodes/ | Create a new node
[**CreateNodeCondition**](NodeAPI.md#CreateNodeCondition) | **Post** /v1/nodes/{node_id}/conditions/ | Create a new node condition
[**CreateNodeDataUsageReport**](NodeAPI.md#CreateNodeDataUsageReport) | **Post** /v1/nodes/{node_id}/data-usage-report/ | Create a new node data usage
[**DeleteNode**](NodeAPI.md#DeleteNode) | **Delete** /v1/nodes/{node_id}/ | Delete a node
[**GetNode**](NodeAPI.md#GetNode) | **Get** /v1/nodes/{node_id}/ | Retrieve a node
[**GetNodeConditions**](NodeAPI.md#GetNodeConditions) | **Get** /v1/nodes/{node_id}/conditions/ | List node conditions
[**GetNodeDevices**](NodeAPI.md#GetNodeDevices) | **Get** /v1/nodes/{node_id}/devices/ | List node devices
[**GetNodes**](NodeAPI.md#GetNodes) | **Get** /v1/nodes/ | List nodes
[**PartialUpdateNode**](NodeAPI.md#PartialUpdateNode) | **Patch** /v1/nodes/{node_id}/ | Partially update a device
[**SearchNodes**](NodeAPI.md#SearchNodes) | **Get** /v1/nodes/search/ | Search for nodes with various filters
[**UpdateNode**](NodeAPI.md#UpdateNode) | **Put** /v1/nodes/{node_id}/ | Update a node



## CreateNode

> Node CreateNode(ctx).NodeRequest(nodeRequest).Execute()

Create a new node

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
	nodeRequest := *openapiclient.NewNodeRequest("Hostname_example", []string{"IpAddresses_example"}, "PubKey_example", []string{"Ports_example"}) // NodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.CreateNode(context.Background()).NodeRequest(nodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.CreateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNode`: Node
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.CreateNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeRequest** | [**NodeRequest**](NodeRequest.md) |  | 

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


## CreateNodeCondition

> CreateNodeCondition(ctx, nodeId).NodeConditionsRequest(nodeConditionsRequest).Execute()

Create a new node condition

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	nodeConditionsRequest := *openapiclient.NewNodeConditionsRequest([]openapiclient.NodeConditionRequest{*openapiclient.NewNodeConditionRequest(openapiclient.Condition("ready"), false)}) // NodeConditionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodeAPI.CreateNodeCondition(context.Background(), nodeId).NodeConditionsRequest(nodeConditionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.CreateNodeCondition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeConditionsRequest** | [**NodeConditionsRequest**](NodeConditionsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodeDataUsageReport

> CreateNodeDataUsageReport(ctx, nodeId).NodeDataUsageReportRequest(nodeDataUsageReportRequest).Execute()

Create a new node data usage

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Node.
	nodeDataUsageReportRequest := *openapiclient.NewNodeDataUsageReportRequest(time.Now(), []openapiclient.NodeDataUsageItemRequest{*openapiclient.NewNodeDataUsageItemRequest("Id_example", "Proto_example")}) // NodeDataUsageReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodeAPI.CreateNodeDataUsageReport(context.Background(), nodeId).NodeDataUsageReportRequest(nodeDataUsageReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.CreateNodeDataUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | A UUID string identifying this Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeDataUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeDataUsageReportRequest** | [**NodeDataUsageReportRequest**](NodeDataUsageReportRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNode

> DeleteNode(ctx, nodeId).Execute()

Delete a node

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodeAPI.DeleteNode(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | A UUID string identifying this Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


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


## GetNode

> Node GetNode(ctx, nodeId).Execute()

Retrieve a node

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.GetNode(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNode`: Node
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | A UUID string identifying this Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Node**](Node.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeConditions

> []NodeCondition GetNodeConditions(ctx, nodeId).Status(status).Execute()

List node conditions

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	status := true // bool | Filter by status: true (active), false (inactive), or all (both). Default is to show active and recently failed.  * `true` - True * `false` - False * `all` - All (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.GetNodeConditions(context.Background(), nodeId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.GetNodeConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeConditions`: []NodeCondition
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.GetNodeConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **bool** | Filter by status: true (active), false (inactive), or all (both). Default is to show active and recently failed.  * &#x60;true&#x60; - True * &#x60;false&#x60; - False * &#x60;all&#x60; - All | 

### Return type

[**[]NodeCondition**](NodeCondition.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeDevices

> []NodeDevice GetNodeDevices(ctx, nodeId).Execute()

List node devices

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.GetNodeDevices(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.GetNodeDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeDevices`: []NodeDevice
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.GetNodeDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | A UUID string identifying this Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NodeDevice**](NodeDevice.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodes

> PaginatedNodeList GetNodes(ctx).Cursor(cursor).Limit(limit).Execute()

List nodes

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
	resp, r, err := apiClient.NodeAPI.GetNodes(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodes`: PaginatedNodeList
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedNodeList**](PaginatedNodeList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateNode

> Node PartialUpdateNode(ctx, nodeId).PatchedNodeRequest(patchedNodeRequest).Execute()

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Node.
	patchedNodeRequest := *openapiclient.NewPatchedNodeRequest() // PatchedNodeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.PartialUpdateNode(context.Background(), nodeId).PatchedNodeRequest(patchedNodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.PartialUpdateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateNode`: Node
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.PartialUpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | A UUID string identifying this Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNodeRequest** | [**PatchedNodeRequest**](PatchedNodeRequest.md) |  | 

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


## SearchNodes

> PaginatedNodeList SearchNodes(ctx).Countries(countries).Cursor(cursor).Distance(distance).IsExitNode(isExitNode).Lat(lat).Limit(limit).Lon(lon).OrderByDistance(orderByDistance).Q(q).Tags(tags).Execute()

Search for nodes with various filters

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
	countries := "countries_example" // string | Filter by country codes (optional)
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	distance := "distance_example" // string | Maximum distance (format: 10km, 5mi) (optional)
	isExitNode := true // bool | Filter by exit node status (optional)
	lat := float64(1.2) // float64 | Latitude for location search (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	lon := float64(1.2) // float64 | Longitude for location search (optional)
	orderByDistance := true // bool | Order by distance (optional)
	q := "q_example" // string | Search nodes by hostname, tags, OS, OS arch, distro, distro codename, or public key (optional)
	tags := "tags_example" // string | Filter by tags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.SearchNodes(context.Background()).Countries(countries).Cursor(cursor).Distance(distance).IsExitNode(isExitNode).Lat(lat).Limit(limit).Lon(lon).OrderByDistance(orderByDistance).Q(q).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.SearchNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNodes`: PaginatedNodeList
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.SearchNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countries** | **string** | Filter by country codes | 
 **cursor** | **string** | The pagination cursor value. | 
 **distance** | **string** | Maximum distance (format: 10km, 5mi) | 
 **isExitNode** | **bool** | Filter by exit node status | 
 **lat** | **float64** | Latitude for location search | 
 **limit** | **int32** | Number of results to return per page. | 
 **lon** | **float64** | Longitude for location search | 
 **orderByDistance** | **bool** | Order by distance | 
 **q** | **string** | Search nodes by hostname, tags, OS, OS arch, distro, distro codename, or public key | 
 **tags** | **string** | Filter by tags | 

### Return type

[**PaginatedNodeList**](PaginatedNodeList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> Node UpdateNode(ctx, nodeId).NodeRequest(nodeRequest).Execute()

Update a node

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
	nodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Node.
	nodeRequest := *openapiclient.NewNodeRequest("Hostname_example", []string{"IpAddresses_example"}, "PubKey_example", []string{"Ports_example"}) // NodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAPI.UpdateNode(context.Background(), nodeId).NodeRequest(nodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAPI.UpdateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNode`: Node
	fmt.Fprintf(os.Stdout, "Response from `NodeAPI.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | A UUID string identifying this Node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeRequest** | [**NodeRequest**](NodeRequest.md) |  | 

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

