# \GeoAPI

All URIs are relative to *https://api.fvpn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](GeoAPI.md#GetCountries) | **Get** /v1/countries/ | List countries
[**GetCountry**](GeoAPI.md#GetCountry) | **Get** /v1/countries/{country_id}/ | Retrieve country
[**GetLocation**](GeoAPI.md#GetLocation) | **Get** /v1/locations/{place_id}/ | Retrieve location
[**GetLocations**](GeoAPI.md#GetLocations) | **Get** /v1/locations/ | List locations
[**GetRegion**](GeoAPI.md#GetRegion) | **Get** /v1/regions/{region_id}/ | Retrieve region
[**GetRegions**](GeoAPI.md#GetRegions) | **Get** /v1/regions/ | List regions



## GetCountries

> []Country GetCountries(ctx).Execute()

List countries

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
	resp, r, err := apiClient.GeoAPI.GetCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: []Country
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


### Return type

[**[]Country**](Country.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountry

> Country GetCountry(ctx, countryId).Execute()

Retrieve country

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
	countryId := "countryId_example" // string | A unique value identifying this Country.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.GetCountry(context.Background(), countryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.GetCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountry`: Country
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.GetCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** | A unique value identifying this Country. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Country**](Country.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocation

> Location GetLocation(ctx, placeId).Execute()

Retrieve location

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
	placeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Place.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.GetLocation(context.Background(), placeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.GetLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocation`: Location
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string** | A UUID string identifying this Place. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocations

> PaginatedLocationList GetLocations(ctx).Country(country).Cursor(cursor).Limit(limit).Name(name).Q(q).Execute()

List locations

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
	country := "country_example" // string |  (optional)
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	q := "q_example" // string | Search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.GetLocations(context.Background()).Country(country).Cursor(cursor).Limit(limit).Name(name).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.GetLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocations`: PaginatedLocationList
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **cursor** | **string** | The pagination cursor value. | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **q** | **string** | Search | 

### Return type

[**PaginatedLocationList**](PaginatedLocationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegion

> RegionDetail GetRegion(ctx, regionId).Execute()

Retrieve region

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
	regionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Region.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.GetRegion(context.Background(), regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.GetRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegion`: RegionDetail
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.GetRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string** | A UUID string identifying this Region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegionDetail**](RegionDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegions

> []Region GetRegions(ctx).Code(code).Country(country).Name(name).Execute()

List regions

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
	code := "code_example" // string |  (optional)
	country := "country_example" // string |  (optional)
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.GetRegions(context.Background()).Code(code).Country(country).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.GetRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegions`: []Region
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.GetRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **country** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**[]Region**](Region.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

