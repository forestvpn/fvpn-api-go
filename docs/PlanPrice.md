# PlanPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Interval** | Pointer to [**Interval**](Interval.md) |  | [optional] 
**Amount** | **float64** |  | [readonly] 
**Currency** | **string** |  | 
**TrialDays** | Pointer to **NullableInt64** | Number of trial days, must be a positive integer greater than 0 | [optional] 
**PriceDiffInfo** | [**NullablePriceDiffInfo**](PriceDiffInfo.md) |  | [readonly] 
**IsMostPopular** | Pointer to **bool** |  | [optional] 
**IsEligibleForTrial** | **bool** |  | [readonly] 

## Methods

### NewPlanPrice

`func NewPlanPrice(id string, amount float64, currency string, priceDiffInfo NullablePriceDiffInfo, isEligibleForTrial bool, ) *PlanPrice`

NewPlanPrice instantiates a new PlanPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPriceWithDefaults

`func NewPlanPriceWithDefaults() *PlanPrice`

NewPlanPriceWithDefaults instantiates a new PlanPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanPrice) SetId(v string)`

SetId sets Id field to given value.


### GetInterval

`func (o *PlanPrice) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanPrice) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanPrice) SetInterval(v Interval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PlanPrice) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetAmount

`func (o *PlanPrice) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlanPrice) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlanPrice) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PlanPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTrialDays

`func (o *PlanPrice) GetTrialDays() int64`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *PlanPrice) GetTrialDaysOk() (*int64, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *PlanPrice) SetTrialDays(v int64)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *PlanPrice) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### SetTrialDaysNil

`func (o *PlanPrice) SetTrialDaysNil(b bool)`

 SetTrialDaysNil sets the value for TrialDays to be an explicit nil

### UnsetTrialDays
`func (o *PlanPrice) UnsetTrialDays()`

UnsetTrialDays ensures that no value is present for TrialDays, not even an explicit nil
### GetPriceDiffInfo

`func (o *PlanPrice) GetPriceDiffInfo() PriceDiffInfo`

GetPriceDiffInfo returns the PriceDiffInfo field if non-nil, zero value otherwise.

### GetPriceDiffInfoOk

`func (o *PlanPrice) GetPriceDiffInfoOk() (*PriceDiffInfo, bool)`

GetPriceDiffInfoOk returns a tuple with the PriceDiffInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiffInfo

`func (o *PlanPrice) SetPriceDiffInfo(v PriceDiffInfo)`

SetPriceDiffInfo sets PriceDiffInfo field to given value.


### SetPriceDiffInfoNil

`func (o *PlanPrice) SetPriceDiffInfoNil(b bool)`

 SetPriceDiffInfoNil sets the value for PriceDiffInfo to be an explicit nil

### UnsetPriceDiffInfo
`func (o *PlanPrice) UnsetPriceDiffInfo()`

UnsetPriceDiffInfo ensures that no value is present for PriceDiffInfo, not even an explicit nil
### GetIsMostPopular

`func (o *PlanPrice) GetIsMostPopular() bool`

GetIsMostPopular returns the IsMostPopular field if non-nil, zero value otherwise.

### GetIsMostPopularOk

`func (o *PlanPrice) GetIsMostPopularOk() (*bool, bool)`

GetIsMostPopularOk returns a tuple with the IsMostPopular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMostPopular

`func (o *PlanPrice) SetIsMostPopular(v bool)`

SetIsMostPopular sets IsMostPopular field to given value.

### HasIsMostPopular

`func (o *PlanPrice) HasIsMostPopular() bool`

HasIsMostPopular returns a boolean if a field has been set.

### GetIsEligibleForTrial

`func (o *PlanPrice) GetIsEligibleForTrial() bool`

GetIsEligibleForTrial returns the IsEligibleForTrial field if non-nil, zero value otherwise.

### GetIsEligibleForTrialOk

`func (o *PlanPrice) GetIsEligibleForTrialOk() (*bool, bool)`

GetIsEligibleForTrialOk returns a tuple with the IsEligibleForTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForTrial

`func (o *PlanPrice) SetIsEligibleForTrial(v bool)`

SetIsEligibleForTrial sets IsEligibleForTrial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


