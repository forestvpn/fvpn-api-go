# CloudPaymentsCardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First6** | Pointer to **NullableString** |  | [optional] 
**Last4** | **string** |  | 
**ExpYear** | **int64** |  | 
**ExpMonth** | **int64** |  | 
**ExpiresAt** | **time.Time** |  | 
**Brand** | **string** |  | 
**Country** | Pointer to [**NullableCountry**](Country.md) |  | [optional] 

## Methods

### NewCloudPaymentsCardDetails

`func NewCloudPaymentsCardDetails(last4 string, expYear int64, expMonth int64, expiresAt time.Time, brand string, ) *CloudPaymentsCardDetails`

NewCloudPaymentsCardDetails instantiates a new CloudPaymentsCardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPaymentsCardDetailsWithDefaults

`func NewCloudPaymentsCardDetailsWithDefaults() *CloudPaymentsCardDetails`

NewCloudPaymentsCardDetailsWithDefaults instantiates a new CloudPaymentsCardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst6

`func (o *CloudPaymentsCardDetails) GetFirst6() string`

GetFirst6 returns the First6 field if non-nil, zero value otherwise.

### GetFirst6Ok

`func (o *CloudPaymentsCardDetails) GetFirst6Ok() (*string, bool)`

GetFirst6Ok returns a tuple with the First6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst6

`func (o *CloudPaymentsCardDetails) SetFirst6(v string)`

SetFirst6 sets First6 field to given value.

### HasFirst6

`func (o *CloudPaymentsCardDetails) HasFirst6() bool`

HasFirst6 returns a boolean if a field has been set.

### SetFirst6Nil

`func (o *CloudPaymentsCardDetails) SetFirst6Nil(b bool)`

 SetFirst6Nil sets the value for First6 to be an explicit nil

### UnsetFirst6
`func (o *CloudPaymentsCardDetails) UnsetFirst6()`

UnsetFirst6 ensures that no value is present for First6, not even an explicit nil
### GetLast4

`func (o *CloudPaymentsCardDetails) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *CloudPaymentsCardDetails) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *CloudPaymentsCardDetails) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetExpYear

`func (o *CloudPaymentsCardDetails) GetExpYear() int64`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *CloudPaymentsCardDetails) GetExpYearOk() (*int64, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *CloudPaymentsCardDetails) SetExpYear(v int64)`

SetExpYear sets ExpYear field to given value.


### GetExpMonth

`func (o *CloudPaymentsCardDetails) GetExpMonth() int64`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *CloudPaymentsCardDetails) GetExpMonthOk() (*int64, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *CloudPaymentsCardDetails) SetExpMonth(v int64)`

SetExpMonth sets ExpMonth field to given value.


### GetExpiresAt

`func (o *CloudPaymentsCardDetails) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudPaymentsCardDetails) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudPaymentsCardDetails) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetBrand

`func (o *CloudPaymentsCardDetails) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CloudPaymentsCardDetails) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CloudPaymentsCardDetails) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetCountry

`func (o *CloudPaymentsCardDetails) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CloudPaymentsCardDetails) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CloudPaymentsCardDetails) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CloudPaymentsCardDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CloudPaymentsCardDetails) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CloudPaymentsCardDetails) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


