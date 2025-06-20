# IPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**Asn** | Pointer to [**NullableASN**](ASN.md) |  | [optional] 
**Country** | Pointer to [**NullableCountry**](Country.md) |  | [optional] 
**Location** | Pointer to [**NullableLocation**](Location.md) |  | [optional] 
**Lat** | Pointer to **NullableFloat64** |  | [optional] 
**Lon** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewIPInfo

`func NewIPInfo(ip string, ) *IPInfo`

NewIPInfo instantiates a new IPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPInfoWithDefaults

`func NewIPInfoWithDefaults() *IPInfo`

NewIPInfoWithDefaults instantiates a new IPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IPInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPInfo) SetIp(v string)`

SetIp sets Ip field to given value.


### GetAsn

`func (o *IPInfo) GetAsn() ASN`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *IPInfo) GetAsnOk() (*ASN, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *IPInfo) SetAsn(v ASN)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *IPInfo) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *IPInfo) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *IPInfo) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetCountry

`func (o *IPInfo) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IPInfo) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IPInfo) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IPInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *IPInfo) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *IPInfo) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLocation

`func (o *IPInfo) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IPInfo) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IPInfo) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IPInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *IPInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *IPInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLat

`func (o *IPInfo) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *IPInfo) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *IPInfo) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *IPInfo) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *IPInfo) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *IPInfo) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLon

`func (o *IPInfo) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *IPInfo) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *IPInfo) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *IPInfo) HasLon() bool`

HasLon returns a boolean if a field has been set.

### SetLonNil

`func (o *IPInfo) SetLonNil(b bool)`

 SetLonNil sets the value for Lon to be an explicit nil

### UnsetLon
`func (o *IPInfo) UnsetLon()`

UnsetLon ensures that no value is present for Lon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


