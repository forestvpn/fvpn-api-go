# TimelineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**EventType** | [**TimelineEventType**](TimelineEventType.md) |  | 
**Description** | **string** |  | 
**Amount** | **NullableFloat64** |  | 
**Currency** | **NullableString** |  | 

## Methods

### NewTimelineEvent

`func NewTimelineEvent(date time.Time, eventType TimelineEventType, description string, amount NullableFloat64, currency NullableString, ) *TimelineEvent`

NewTimelineEvent instantiates a new TimelineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEventWithDefaults

`func NewTimelineEventWithDefaults() *TimelineEvent`

NewTimelineEventWithDefaults instantiates a new TimelineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TimelineEvent) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimelineEvent) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimelineEvent) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetEventType

`func (o *TimelineEvent) GetEventType() TimelineEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TimelineEvent) GetEventTypeOk() (*TimelineEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TimelineEvent) SetEventType(v TimelineEventType)`

SetEventType sets EventType field to given value.


### GetDescription

`func (o *TimelineEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimelineEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimelineEvent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *TimelineEvent) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TimelineEvent) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TimelineEvent) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *TimelineEvent) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *TimelineEvent) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *TimelineEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TimelineEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TimelineEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *TimelineEvent) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *TimelineEvent) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


