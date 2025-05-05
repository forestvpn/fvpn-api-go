# \CloudpaymentsAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteCloudpaymentsPaymentMethodThreeDs**](CloudpaymentsAPI.md#CompleteCloudpaymentsPaymentMethodThreeDs) | **Post** /v1/billing/cloudpayments/payment-methods/{payment_method_id}/complete-3ds/ | Complete 3D Secure authentication
[**CreateCloudpaymentsPaymentMethod**](CloudpaymentsAPI.md#CreateCloudpaymentsPaymentMethod) | **Post** /v1/billing/cloudpayments/payment-methods/ | Create a new payment method
[**GetCloudpaymentsPaymentMethod**](CloudpaymentsAPI.md#GetCloudpaymentsPaymentMethod) | **Get** /v1/billing/cloudpayments/payment-methods/{payment_method_id}/ | Retrieve a payment method



## CompleteCloudpaymentsPaymentMethodThreeDs

> CloudPaymentsPaymentMethod CompleteCloudpaymentsPaymentMethodThreeDs(ctx, paymentMethodId).CloudPaymentsThreeDSecureCompleteRequest(cloudPaymentsThreeDSecureCompleteRequest).Execute()

Complete 3D Secure authentication



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
	paymentMethodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Payment Method.
	cloudPaymentsThreeDSecureCompleteRequest := *openapiclient.NewCloudPaymentsThreeDSecureCompleteRequest("PaRes_example") // CloudPaymentsThreeDSecureCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudpaymentsAPI.CompleteCloudpaymentsPaymentMethodThreeDs(context.Background(), paymentMethodId).CloudPaymentsThreeDSecureCompleteRequest(cloudPaymentsThreeDSecureCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudpaymentsAPI.CompleteCloudpaymentsPaymentMethodThreeDs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteCloudpaymentsPaymentMethodThreeDs`: CloudPaymentsPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `CloudpaymentsAPI.CompleteCloudpaymentsPaymentMethodThreeDs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | A UUID string identifying this Payment Method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteCloudpaymentsPaymentMethodThreeDsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPaymentsThreeDSecureCompleteRequest** | [**CloudPaymentsThreeDSecureCompleteRequest**](CloudPaymentsThreeDSecureCompleteRequest.md) |  | 

### Return type

[**CloudPaymentsPaymentMethod**](CloudPaymentsPaymentMethod.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudpaymentsPaymentMethod

> CloudPaymentsPaymentMethod CreateCloudpaymentsPaymentMethod(ctx).CloudPaymentsPaymentMethodCreateRequest(cloudPaymentsPaymentMethodCreateRequest).Execute()

Create a new payment method



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
	cloudPaymentsPaymentMethodCreateRequest := *openapiclient.NewCloudPaymentsPaymentMethodCreateRequest("Cryptogram_example") // CloudPaymentsPaymentMethodCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudpaymentsAPI.CreateCloudpaymentsPaymentMethod(context.Background()).CloudPaymentsPaymentMethodCreateRequest(cloudPaymentsPaymentMethodCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudpaymentsAPI.CreateCloudpaymentsPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudpaymentsPaymentMethod`: CloudPaymentsPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `CloudpaymentsAPI.CreateCloudpaymentsPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudpaymentsPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPaymentsPaymentMethodCreateRequest** | [**CloudPaymentsPaymentMethodCreateRequest**](CloudPaymentsPaymentMethodCreateRequest.md) |  | 

### Return type

[**CloudPaymentsPaymentMethod**](CloudPaymentsPaymentMethod.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudpaymentsPaymentMethod

> CloudPaymentsPaymentMethod GetCloudpaymentsPaymentMethod(ctx, paymentMethodId).Execute()

Retrieve a payment method



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
	paymentMethodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Payment Method.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudpaymentsAPI.GetCloudpaymentsPaymentMethod(context.Background(), paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudpaymentsAPI.GetCloudpaymentsPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudpaymentsPaymentMethod`: CloudPaymentsPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `CloudpaymentsAPI.GetCloudpaymentsPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | A UUID string identifying this Payment Method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudpaymentsPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPaymentsPaymentMethod**](CloudPaymentsPaymentMethod.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

