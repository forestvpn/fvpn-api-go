# \StripeAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStripeSetupIntent**](StripeAPI.md#CreateStripeSetupIntent) | **Post** /v1/billing/stripe/setup-intents/ | Create a new stripe setup intent



## CreateStripeSetupIntent

> StripeSetupIntent CreateStripeSetupIntent(ctx).Execute()

Create a new stripe setup intent

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
	resp, r, err := apiClient.StripeAPI.CreateStripeSetupIntent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StripeAPI.CreateStripeSetupIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStripeSetupIntent`: StripeSetupIntent
	fmt.Fprintf(os.Stdout, "Response from `StripeAPI.CreateStripeSetupIntent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStripeSetupIntentRequest struct via the builder pattern


### Return type

[**StripeSetupIntent**](StripeSetupIntent.md)

### Authorization

[bearerToken](../README.md#bearerToken), [apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

