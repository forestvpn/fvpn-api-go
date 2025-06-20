# \BillingAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](BillingAPI.md#CancelSubscription) | **Delete** /v1/billing/subscriptions/{subscription_id}/ | Cancel subscription
[**ChangeSubscriptionPlan**](BillingAPI.md#ChangeSubscriptionPlan) | **Post** /v1/billing/subscriptions/{subscription_id}/change-plan/ | Change subscription plan
[**DeletePaymentMethod**](BillingAPI.md#DeletePaymentMethod) | **Delete** /v1/billing/payment-methods/{payment_method_id}/ | Delete a payment method
[**GetBillingAccount**](BillingAPI.md#GetBillingAccount) | **Get** /v1/billing/account/ | Retrieve the current account
[**GetBillingOptions**](BillingAPI.md#GetBillingOptions) | **Get** /v1/billing/options/ | Retrieve billing options
[**GetCurrencies**](BillingAPI.md#GetCurrencies) | **Get** /v1/billing/currencies/ | Retrieve all available currencies
[**GetCurrentSubscription**](BillingAPI.md#GetCurrentSubscription) | **Get** /v1/billing/current-subscription/ | Retrieve the current subscription
[**GetCurrentWallet**](BillingAPI.md#GetCurrentWallet) | **Get** /v1/billing/current-wallet/ | Retrieve the current wallet
[**GetEntitlement**](BillingAPI.md#GetEntitlement) | **Get** /v1/billing/entitlements/{entitlement_id}/ | Retrieve a specific entitlement
[**GetEntitlements**](BillingAPI.md#GetEntitlements) | **Get** /v1/billing/entitlements/ | List account entitlements
[**GetPaymentMethod**](BillingAPI.md#GetPaymentMethod) | **Get** /v1/billing/payment-methods/{payment_method_id}/ | Retrieve a payment method
[**GetPaymentMethods**](BillingAPI.md#GetPaymentMethods) | **Get** /v1/billing/payment-methods/ | List all payment methods
[**GetPlanBySlug**](BillingAPI.md#GetPlanBySlug) | **Get** /v1/billing/plans/{slug}/ | Retrieve a plan
[**GetPlanTimeline**](BillingAPI.md#GetPlanTimeline) | **Get** /v1/billing/plans/timeline/ | Get billing timeline for a plan
[**GetPlans**](BillingAPI.md#GetPlans) | **Get** /v1/billing/plans/ | List all plans
[**GetSubscription**](BillingAPI.md#GetSubscription) | **Get** /v1/billing/subscriptions/{subscription_id}/ | Retrieve subscription
[**GetSubscriptions**](BillingAPI.md#GetSubscriptions) | **Get** /v1/billing/subscriptions/ | List subscriptions
[**GetWallet**](BillingAPI.md#GetWallet) | **Get** /v1/billing/wallets/{wallet_id}/ | Retrieve wallet details
[**GetWalletTransaction**](BillingAPI.md#GetWalletTransaction) | **Get** /v1/billing/transactions/{transaction_id}/ | Retrieve transaction details
[**GetWalletTransactions**](BillingAPI.md#GetWalletTransactions) | **Get** /v1/billing/transactions/ | List wallet transactions
[**GetWallets**](BillingAPI.md#GetWallets) | **Get** /v1/billing/wallets/ | List wallets
[**ReactivateSubscription**](BillingAPI.md#ReactivateSubscription) | **Post** /v1/billing/subscriptions/{subscription_id}/reactivate/ | Reactivate subscription
[**SetDefaultPaymentMethod**](BillingAPI.md#SetDefaultPaymentMethod) | **Post** /v1/billing/account/set-default-payment-method/ | Set the default payment method for the account
[**SubscribeToPlan**](BillingAPI.md#SubscribeToPlan) | **Post** /v1/billing/subscriptions/ | Subscribe to a plan



## CancelSubscription

> Subscription CancelSubscription(ctx, subscriptionId).Execute()

Cancel subscription

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
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Subscription.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CancelSubscription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CancelSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | A UUID string identifying this Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeSubscriptionPlan

> Subscription ChangeSubscriptionPlan(ctx, subscriptionId).SubscriptionChangePlanRequest(subscriptionChangePlanRequest).Execute()

Change subscription plan

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
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Subscription.
	subscriptionChangePlanRequest := *openapiclient.NewSubscriptionChangePlanRequest("NewPlanPrice_example") // SubscriptionChangePlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ChangeSubscriptionPlan(context.Background(), subscriptionId).SubscriptionChangePlanRequest(subscriptionChangePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ChangeSubscriptionPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeSubscriptionPlan`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ChangeSubscriptionPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | A UUID string identifying this Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeSubscriptionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionChangePlanRequest** | [**SubscriptionChangePlanRequest**](SubscriptionChangePlanRequest.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentMethod

> DeletePaymentMethod(ctx, paymentMethodId).Execute()

Delete a payment method

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
	paymentMethodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Payment method.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.DeletePaymentMethod(context.Background(), paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeletePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | A UUID string identifying this Payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentMethodRequest struct via the builder pattern


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


## GetBillingAccount

> BillingAccount GetBillingAccount(ctx).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetBillingAccount(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingAccount`: BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**BillingAccount**](BillingAccount.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingOptions

> BillingOptions GetBillingOptions(ctx).Execute()

Retrieve billing options



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
	resp, r, err := apiClient.BillingAPI.GetBillingOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingOptions`: BillingOptions
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingOptionsRequest struct via the builder pattern


### Return type

[**BillingOptions**](BillingOptions.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencies

> []Currency GetCurrencies(ctx).Execute()

Retrieve all available currencies



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
	resp, r, err := apiClient.BillingAPI.GetCurrencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrencies`: []Currency
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetCurrencies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesRequest struct via the builder pattern


### Return type

[**[]Currency**](Currency.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentSubscription

> Subscription GetCurrentSubscription(ctx).Cursor(cursor).Limit(limit).Execute()

Retrieve the current subscription



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
	resp, r, err := apiClient.BillingAPI.GetCurrentSubscription(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetCurrentSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetCurrentSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentWallet

> Wallet GetCurrentWallet(ctx).Cursor(cursor).Limit(limit).Execute()

Retrieve the current wallet



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
	resp, r, err := apiClient.BillingAPI.GetCurrentWallet(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetCurrentWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentWallet`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetCurrentWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlement

> Entitlement GetEntitlement(ctx, entitlementId).Execute()

Retrieve a specific entitlement



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
	entitlementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Entitlement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetEntitlement(context.Background(), entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlement`: Entitlement
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | A UUID string identifying this Entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entitlement**](Entitlement.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlements

> []Entitlement GetEntitlements(ctx).LookupKey(lookupKey).Execute()

List account entitlements



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
	lookupKey := "lookupKey_example" // string | Filter entitlements by lookup key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetEntitlements(context.Background()).LookupKey(lookupKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlements`: []Entitlement
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupKey** | **string** | Filter entitlements by lookup key | 

### Return type

[**[]Entitlement**](Entitlement.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethod

> PaymentMethod GetPaymentMethod(ctx, paymentMethodId).Execute()

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
	paymentMethodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Payment method.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetPaymentMethod(context.Background(), paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentMethod`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | A UUID string identifying this Payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethods

> PaginatedPaymentMethodList GetPaymentMethods(ctx).Cursor(cursor).Limit(limit).Execute()

List all payment methods

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
	resp, r, err := apiClient.BillingAPI.GetPaymentMethods(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentMethods`: PaginatedPaymentMethodList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedPaymentMethodList**](PaginatedPaymentMethodList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanBySlug

> Plan GetPlanBySlug(ctx, slug).Execute()

Retrieve a plan

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetPlanBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetPlanBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanBySlug`: Plan
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetPlanBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanTimeline

> Timeline GetPlanTimeline(ctx).PlanPriceId(planPriceId).SubscriptionId(subscriptionId).Execute()

Get billing timeline for a plan



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
	planPriceId := "planPriceId_example" // string | The ID of the plan price to get the timeline for
	subscriptionId := "subscriptionId_example" // string | Optional subscription ID for upgrade/downgrade scenarios (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetPlanTimeline(context.Background()).PlanPriceId(planPriceId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetPlanTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanTimeline`: Timeline
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetPlanTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planPriceId** | **string** | The ID of the plan price to get the timeline for | 
 **subscriptionId** | **string** | Optional subscription ID for upgrade/downgrade scenarios | 

### Return type

[**Timeline**](Timeline.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlans

> []Plan GetPlans(ctx).Currency(currency).Execute()

List all plans

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
	currency := "currency_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetPlans(context.Background()).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlans`: []Plan
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** |  | 

### Return type

[**[]Plan**](Plan.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> Subscription GetSubscription(ctx, subscriptionId).Execute()

Retrieve subscription

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
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Subscription.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetSubscription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | A UUID string identifying this Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> PaginatedSubscriptionList GetSubscriptions(ctx).Cursor(cursor).Limit(limit).ManagedBy(managedBy).Status(status).Execute()

List subscriptions

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
	managedBy := []*string{"ManagedBy_example"} // []*string | * `apple` - Apple * `google` - Google * `stripe` - Stripe * `unknown` - Unknown (optional)
	status := []string{"Status_example"} // []string | * `incomplete` - Incomplete * `incomplete_expired` - Incomplete expired * `trialing` - Trialing * `active` - Active * `past_due` - Past due * `canceled` - Canceled * `unpaid` - Unpaid (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetSubscriptions(context.Background()).Cursor(cursor).Limit(limit).ManagedBy(managedBy).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: PaginatedSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 
 **managedBy** | **[]string** | * &#x60;apple&#x60; - Apple * &#x60;google&#x60; - Google * &#x60;stripe&#x60; - Stripe * &#x60;unknown&#x60; - Unknown | 
 **status** | **[]string** | * &#x60;incomplete&#x60; - Incomplete * &#x60;incomplete_expired&#x60; - Incomplete expired * &#x60;trialing&#x60; - Trialing * &#x60;active&#x60; - Active * &#x60;past_due&#x60; - Past due * &#x60;canceled&#x60; - Canceled * &#x60;unpaid&#x60; - Unpaid | 

### Return type

[**PaginatedSubscriptionList**](PaginatedSubscriptionList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallet

> Wallet GetWallet(ctx, walletId).Execute()

Retrieve wallet details



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Wallet.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetWallet(context.Background(), walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallet`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | A UUID string identifying this Wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTransaction

> WalletTransaction GetWalletTransaction(ctx, transactionId).Execute()

Retrieve transaction details



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
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Wallet Transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetWalletTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetWalletTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletTransaction`: WalletTransaction
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetWalletTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | A UUID string identifying this Wallet Transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletTransaction**](WalletTransaction.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTransactions

> PaginatedWalletTransactionList GetWalletTransactions(ctx).Cursor(cursor).Limit(limit).Execute()

List wallet transactions



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
	resp, r, err := apiClient.BillingAPI.GetWalletTransactions(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetWalletTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletTransactions`: PaginatedWalletTransactionList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetWalletTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedWalletTransactionList**](PaginatedWalletTransactionList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallets

> PaginatedWalletList GetWallets(ctx).Cursor(cursor).Limit(limit).Execute()

List wallets



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
	resp, r, err := apiClient.BillingAPI.GetWallets(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallets`: PaginatedWalletList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedWalletList**](PaginatedWalletList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateSubscription

> Subscription ReactivateSubscription(ctx, subscriptionId).Execute()

Reactivate subscription

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
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Subscription.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ReactivateSubscription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ReactivateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactivateSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ReactivateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | A UUID string identifying this Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultPaymentMethod

> BillingAccount SetDefaultPaymentMethod(ctx).SetDefaultPaymentMethodRequest(setDefaultPaymentMethodRequest).Execute()

Set the default payment method for the account



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
	setDefaultPaymentMethodRequest := *openapiclient.NewSetDefaultPaymentMethodRequest("PaymentMethod_example") // SetDefaultPaymentMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.SetDefaultPaymentMethod(context.Background()).SetDefaultPaymentMethodRequest(setDefaultPaymentMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.SetDefaultPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultPaymentMethod`: BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.SetDefaultPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultPaymentMethodRequest** | [**SetDefaultPaymentMethodRequest**](SetDefaultPaymentMethodRequest.md) |  | 

### Return type

[**BillingAccount**](BillingAccount.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeToPlan

> Subscription SubscribeToPlan(ctx).SubscriptionCreationRequest(subscriptionCreationRequest).Execute()

Subscribe to a plan

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
	subscriptionCreationRequest := *openapiclient.NewSubscriptionCreationRequest("PlanPrice_example") // SubscriptionCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.SubscribeToPlan(context.Background()).SubscriptionCreationRequest(subscriptionCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.SubscribeToPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeToPlan`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.SubscribeToPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreationRequest** | [**SubscriptionCreationRequest**](SubscriptionCreationRequest.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

