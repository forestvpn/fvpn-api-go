# SubscriptionCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanPrice** | **string** |  | 
**PaymentMethod** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubscriptionCreationRequest

`func NewSubscriptionCreationRequest(planPrice string, ) *SubscriptionCreationRequest`

NewSubscriptionCreationRequest instantiates a new SubscriptionCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreationRequestWithDefaults

`func NewSubscriptionCreationRequestWithDefaults() *SubscriptionCreationRequest`

NewSubscriptionCreationRequestWithDefaults instantiates a new SubscriptionCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanPrice

`func (o *SubscriptionCreationRequest) GetPlanPrice() string`

GetPlanPrice returns the PlanPrice field if non-nil, zero value otherwise.

### GetPlanPriceOk

`func (o *SubscriptionCreationRequest) GetPlanPriceOk() (*string, bool)`

GetPlanPriceOk returns a tuple with the PlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPrice

`func (o *SubscriptionCreationRequest) SetPlanPrice(v string)`

SetPlanPrice sets PlanPrice field to given value.


### GetPaymentMethod

`func (o *SubscriptionCreationRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *SubscriptionCreationRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *SubscriptionCreationRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *SubscriptionCreationRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *SubscriptionCreationRequest) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *SubscriptionCreationRequest) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


