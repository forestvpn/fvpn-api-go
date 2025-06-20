# BillingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DefaultPaymentMethod** | [**NullablePaymentMethod**](PaymentMethod.md) |  | 
**CurrentSubscription** | [**NullableSubscription**](Subscription.md) |  | [readonly] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**Company** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **string** | Default language for this account in ISO 639-1 format | [optional] 
**Currency** | Pointer to **string** | Default currency for this account in ISO 4217 format | [optional] 
**Status** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewBillingAccount

`func NewBillingAccount(id string, defaultPaymentMethod NullablePaymentMethod, currentSubscription NullableSubscription, createdAt time.Time, ) *BillingAccount`

NewBillingAccount instantiates a new BillingAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAccountWithDefaults

`func NewBillingAccountWithDefaults() *BillingAccount`

NewBillingAccountWithDefaults instantiates a new BillingAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingAccount) SetId(v string)`

SetId sets Id field to given value.


### GetDefaultPaymentMethod

`func (o *BillingAccount) GetDefaultPaymentMethod() PaymentMethod`

GetDefaultPaymentMethod returns the DefaultPaymentMethod field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodOk

`func (o *BillingAccount) GetDefaultPaymentMethodOk() (*PaymentMethod, bool)`

GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethod

`func (o *BillingAccount) SetDefaultPaymentMethod(v PaymentMethod)`

SetDefaultPaymentMethod sets DefaultPaymentMethod field to given value.


### SetDefaultPaymentMethodNil

`func (o *BillingAccount) SetDefaultPaymentMethodNil(b bool)`

 SetDefaultPaymentMethodNil sets the value for DefaultPaymentMethod to be an explicit nil

### UnsetDefaultPaymentMethod
`func (o *BillingAccount) UnsetDefaultPaymentMethod()`

UnsetDefaultPaymentMethod ensures that no value is present for DefaultPaymentMethod, not even an explicit nil
### GetCurrentSubscription

`func (o *BillingAccount) GetCurrentSubscription() Subscription`

GetCurrentSubscription returns the CurrentSubscription field if non-nil, zero value otherwise.

### GetCurrentSubscriptionOk

`func (o *BillingAccount) GetCurrentSubscriptionOk() (*Subscription, bool)`

GetCurrentSubscriptionOk returns a tuple with the CurrentSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSubscription

`func (o *BillingAccount) SetCurrentSubscription(v Subscription)`

SetCurrentSubscription sets CurrentSubscription field to given value.


### SetCurrentSubscriptionNil

`func (o *BillingAccount) SetCurrentSubscriptionNil(b bool)`

 SetCurrentSubscriptionNil sets the value for CurrentSubscription to be an explicit nil

### UnsetCurrentSubscription
`func (o *BillingAccount) UnsetCurrentSubscription()`

UnsetCurrentSubscription ensures that no value is present for CurrentSubscription, not even an explicit nil
### GetFullName

`func (o *BillingAccount) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *BillingAccount) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *BillingAccount) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *BillingAccount) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *BillingAccount) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *BillingAccount) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetCompany

`func (o *BillingAccount) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BillingAccount) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BillingAccount) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BillingAccount) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *BillingAccount) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *BillingAccount) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetEmail

`func (o *BillingAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BillingAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *BillingAccount) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *BillingAccount) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *BillingAccount) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BillingAccount) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BillingAccount) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BillingAccount) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *BillingAccount) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *BillingAccount) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAddress

`func (o *BillingAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingAccount) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BillingAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *BillingAccount) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *BillingAccount) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCity

`func (o *BillingAccount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingAccount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingAccount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillingAccount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *BillingAccount) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *BillingAccount) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *BillingAccount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingAccount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingAccount) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BillingAccount) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *BillingAccount) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BillingAccount) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *BillingAccount) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingAccount) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingAccount) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BillingAccount) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *BillingAccount) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *BillingAccount) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *BillingAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingAccount) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillingAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *BillingAccount) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *BillingAccount) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLanguage

`func (o *BillingAccount) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BillingAccount) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BillingAccount) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BillingAccount) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *BillingAccount) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingAccount) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingAccount) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


