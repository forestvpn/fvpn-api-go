# ASN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int64** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Country** | [**Country**](Country.md) |  | 
**Domain** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewASN

`func NewASN(asn int64, country Country, ) *ASN`

NewASN instantiates a new ASN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNWithDefaults

`func NewASNWithDefaults() *ASN`

NewASNWithDefaults instantiates a new ASN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *ASN) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ASN) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ASN) SetAsn(v int64)`

SetAsn sets Asn field to given value.


### GetName

`func (o *ASN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ASN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ASN) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ASN) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ASN) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ASN) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCountry

`func (o *ASN) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ASN) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ASN) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetDomain

`func (o *ASN) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ASN) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ASN) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ASN) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ASN) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ASN) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


