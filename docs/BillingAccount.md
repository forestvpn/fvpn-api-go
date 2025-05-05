# BillingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DefaultPaymentMethod** | [**NullablePaymentMethod**](PaymentMethod.md) |  | 
**CurrentSubscription** | [**NullableSubscription**](Subscription.md) |  | [readonly] 

## Methods

### NewBillingAccount

`func NewBillingAccount(id string, defaultPaymentMethod NullablePaymentMethod, currentSubscription NullableSubscription, ) *BillingAccount`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


