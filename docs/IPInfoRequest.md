# IPInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**Asn** | Pointer to [**NullableASNRequest**](ASNRequest.md) |  | [optional] 
**Country** | Pointer to [**NullableCountryRequest**](CountryRequest.md) |  | [optional] 
**Location** | Pointer to [**NullableLocationRequest**](LocationRequest.md) |  | [optional] 
**Lat** | Pointer to **NullableFloat64** |  | [optional] 
**Lon** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewIPInfoRequest

`func NewIPInfoRequest(ip string, ) *IPInfoRequest`

NewIPInfoRequest instantiates a new IPInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPInfoRequestWithDefaults

`func NewIPInfoRequestWithDefaults() *IPInfoRequest`

NewIPInfoRequestWithDefaults instantiates a new IPInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IPInfoRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPInfoRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPInfoRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### GetAsn

`func (o *IPInfoRequest) GetAsn() ASNRequest`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *IPInfoRequest) GetAsnOk() (*ASNRequest, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *IPInfoRequest) SetAsn(v ASNRequest)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *IPInfoRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *IPInfoRequest) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *IPInfoRequest) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetCountry

`func (o *IPInfoRequest) GetCountry() CountryRequest`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IPInfoRequest) GetCountryOk() (*CountryRequest, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IPInfoRequest) SetCountry(v CountryRequest)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IPInfoRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *IPInfoRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *IPInfoRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLocation

`func (o *IPInfoRequest) GetLocation() LocationRequest`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IPInfoRequest) GetLocationOk() (*LocationRequest, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IPInfoRequest) SetLocation(v LocationRequest)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IPInfoRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *IPInfoRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *IPInfoRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLat

`func (o *IPInfoRequest) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *IPInfoRequest) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *IPInfoRequest) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *IPInfoRequest) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *IPInfoRequest) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *IPInfoRequest) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLon

`func (o *IPInfoRequest) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *IPInfoRequest) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *IPInfoRequest) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *IPInfoRequest) HasLon() bool`

HasLon returns a boolean if a field has been set.

### SetLonNil

`func (o *IPInfoRequest) SetLonNil(b bool)`

 SetLonNil sets the value for Lon to be an explicit nil

### UnsetLon
`func (o *IPInfoRequest) UnsetLon()`

UnsetLon ensures that no value is present for Lon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


