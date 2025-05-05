# ComparisonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**SummaryPeriod**](SummaryPeriod.md) |  | 
**Previous** | [**SummaryPeriod**](SummaryPeriod.md) |  | 
**PercentChange** | **float64** |  | 

## Methods

### NewComparisonData

`func NewComparisonData(current SummaryPeriod, previous SummaryPeriod, percentChange float64, ) *ComparisonData`

NewComparisonData instantiates a new ComparisonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonDataWithDefaults

`func NewComparisonDataWithDefaults() *ComparisonData`

NewComparisonDataWithDefaults instantiates a new ComparisonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *ComparisonData) GetCurrent() SummaryPeriod`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ComparisonData) GetCurrentOk() (*SummaryPeriod, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ComparisonData) SetCurrent(v SummaryPeriod)`

SetCurrent sets Current field to given value.


### GetPrevious

`func (o *ComparisonData) GetPrevious() SummaryPeriod`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ComparisonData) GetPreviousOk() (*SummaryPeriod, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ComparisonData) SetPrevious(v SummaryPeriod)`

SetPrevious sets Previous field to given value.


### GetPercentChange

`func (o *ComparisonData) GetPercentChange() float64`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *ComparisonData) GetPercentChangeOk() (*float64, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *ComparisonData) SetPercentChange(v float64)`

SetPercentChange sets PercentChange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


