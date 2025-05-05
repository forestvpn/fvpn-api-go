# Timeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]TimelineEvent**](TimelineEvent.md) |  | 
**ProratedAmount** | **NullableFloat64** |  | 
**Currency** | **NullableString** |  | 

## Methods

### NewTimeline

`func NewTimeline(events []TimelineEvent, proratedAmount NullableFloat64, currency NullableString, ) *Timeline`

NewTimeline instantiates a new Timeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineWithDefaults

`func NewTimelineWithDefaults() *Timeline`

NewTimelineWithDefaults instantiates a new Timeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Timeline) GetEvents() []TimelineEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Timeline) GetEventsOk() (*[]TimelineEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Timeline) SetEvents(v []TimelineEvent)`

SetEvents sets Events field to given value.


### GetProratedAmount

`func (o *Timeline) GetProratedAmount() float64`

GetProratedAmount returns the ProratedAmount field if non-nil, zero value otherwise.

### GetProratedAmountOk

`func (o *Timeline) GetProratedAmountOk() (*float64, bool)`

GetProratedAmountOk returns a tuple with the ProratedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedAmount

`func (o *Timeline) SetProratedAmount(v float64)`

SetProratedAmount sets ProratedAmount field to given value.


### SetProratedAmountNil

`func (o *Timeline) SetProratedAmountNil(b bool)`

 SetProratedAmountNil sets the value for ProratedAmount to be an explicit nil

### UnsetProratedAmount
`func (o *Timeline) UnsetProratedAmount()`

UnsetProratedAmount ensures that no value is present for ProratedAmount, not even an explicit nil
### GetCurrency

`func (o *Timeline) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Timeline) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Timeline) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *Timeline) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Timeline) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


