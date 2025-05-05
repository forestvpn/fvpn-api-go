# SummaryPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**Period**](Period.md) |  | 
**Data** | [**[]DataPoint**](DataPoint.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewSummaryPeriod

`func NewSummaryPeriod(period Period, data []DataPoint, total int64, ) *SummaryPeriod`

NewSummaryPeriod instantiates a new SummaryPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryPeriodWithDefaults

`func NewSummaryPeriodWithDefaults() *SummaryPeriod`

NewSummaryPeriodWithDefaults instantiates a new SummaryPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *SummaryPeriod) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SummaryPeriod) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SummaryPeriod) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetData

`func (o *SummaryPeriod) GetData() []DataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SummaryPeriod) GetDataOk() (*[]DataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SummaryPeriod) SetData(v []DataPoint)`

SetData sets Data field to given value.


### GetTotal

`func (o *SummaryPeriod) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SummaryPeriod) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SummaryPeriod) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


