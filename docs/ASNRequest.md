# ASNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int64** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Country** | [**CountryRequest**](CountryRequest.md) |  | 
**Domain** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewASNRequest

`func NewASNRequest(asn int64, country CountryRequest, ) *ASNRequest`

NewASNRequest instantiates a new ASNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNRequestWithDefaults

`func NewASNRequestWithDefaults() *ASNRequest`

NewASNRequestWithDefaults instantiates a new ASNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *ASNRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ASNRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ASNRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.


### GetName

`func (o *ASNRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ASNRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ASNRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ASNRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ASNRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ASNRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCountry

`func (o *ASNRequest) GetCountry() CountryRequest`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ASNRequest) GetCountryOk() (*CountryRequest, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ASNRequest) SetCountry(v CountryRequest)`

SetCountry sets Country field to given value.


### GetDomain

`func (o *ASNRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ASNRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ASNRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ASNRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ASNRequest) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ASNRequest) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


