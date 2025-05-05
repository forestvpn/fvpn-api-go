# PaymentMethodCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First6** | **NullableString** |  | 
**Last4** | **string** |  | 
**ExpMonth** | **int64** |  | 
**ExpYear** | **int64** |  | 
**ExpiresAt** | **time.Time** |  | 
**Brand** | **NullableString** |  | 
**Country** | [**NullableCountry**](Country.md) |  | 

## Methods

### NewPaymentMethodCard

`func NewPaymentMethodCard(first6 NullableString, last4 string, expMonth int64, expYear int64, expiresAt time.Time, brand NullableString, country NullableCountry, ) *PaymentMethodCard`

NewPaymentMethodCard instantiates a new PaymentMethodCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardWithDefaults

`func NewPaymentMethodCardWithDefaults() *PaymentMethodCard`

NewPaymentMethodCardWithDefaults instantiates a new PaymentMethodCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst6

`func (o *PaymentMethodCard) GetFirst6() string`

GetFirst6 returns the First6 field if non-nil, zero value otherwise.

### GetFirst6Ok

`func (o *PaymentMethodCard) GetFirst6Ok() (*string, bool)`

GetFirst6Ok returns a tuple with the First6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst6

`func (o *PaymentMethodCard) SetFirst6(v string)`

SetFirst6 sets First6 field to given value.


### SetFirst6Nil

`func (o *PaymentMethodCard) SetFirst6Nil(b bool)`

 SetFirst6Nil sets the value for First6 to be an explicit nil

### UnsetFirst6
`func (o *PaymentMethodCard) UnsetFirst6()`

UnsetFirst6 ensures that no value is present for First6, not even an explicit nil
### GetLast4

`func (o *PaymentMethodCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethodCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethodCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetExpMonth

`func (o *PaymentMethodCard) GetExpMonth() int64`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *PaymentMethodCard) GetExpMonthOk() (*int64, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *PaymentMethodCard) SetExpMonth(v int64)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *PaymentMethodCard) GetExpYear() int64`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *PaymentMethodCard) GetExpYearOk() (*int64, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *PaymentMethodCard) SetExpYear(v int64)`

SetExpYear sets ExpYear field to given value.


### GetExpiresAt

`func (o *PaymentMethodCard) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodCard) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodCard) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetBrand

`func (o *PaymentMethodCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethodCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethodCard) SetBrand(v string)`

SetBrand sets Brand field to given value.


### SetBrandNil

`func (o *PaymentMethodCard) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *PaymentMethodCard) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCountry

`func (o *PaymentMethodCard) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodCard) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodCard) SetCountry(v Country)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *PaymentMethodCard) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentMethodCard) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


