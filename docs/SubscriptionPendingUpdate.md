# SubscriptionPendingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [**Plan**](Plan.md) |  | [readonly] 
**PlanPrice** | [**PlanPrice**](PlanPrice.md) |  | [readonly] 
**Amount** | **float64** |  | [readonly] 
**Currency** | **string** |  | 
**EffectiveAt** | **time.Time** | The date when the transition will take effect. | 

## Methods

### NewSubscriptionPendingUpdate

`func NewSubscriptionPendingUpdate(plan Plan, planPrice PlanPrice, amount float64, currency string, effectiveAt time.Time, ) *SubscriptionPendingUpdate`

NewSubscriptionPendingUpdate instantiates a new SubscriptionPendingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPendingUpdateWithDefaults

`func NewSubscriptionPendingUpdateWithDefaults() *SubscriptionPendingUpdate`

NewSubscriptionPendingUpdateWithDefaults instantiates a new SubscriptionPendingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *SubscriptionPendingUpdate) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionPendingUpdate) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionPendingUpdate) SetPlan(v Plan)`

SetPlan sets Plan field to given value.


### GetPlanPrice

`func (o *SubscriptionPendingUpdate) GetPlanPrice() PlanPrice`

GetPlanPrice returns the PlanPrice field if non-nil, zero value otherwise.

### GetPlanPriceOk

`func (o *SubscriptionPendingUpdate) GetPlanPriceOk() (*PlanPrice, bool)`

GetPlanPriceOk returns a tuple with the PlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPrice

`func (o *SubscriptionPendingUpdate) SetPlanPrice(v PlanPrice)`

SetPlanPrice sets PlanPrice field to given value.


### GetAmount

`func (o *SubscriptionPendingUpdate) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionPendingUpdate) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionPendingUpdate) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *SubscriptionPendingUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionPendingUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionPendingUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEffectiveAt

`func (o *SubscriptionPendingUpdate) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *SubscriptionPendingUpdate) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *SubscriptionPendingUpdate) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


