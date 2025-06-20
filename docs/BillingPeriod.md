# BillingPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | [readonly] 
**End** | **time.Time** |  | [readonly] 

## Methods

### NewBillingPeriod

`func NewBillingPeriod(start time.Time, end time.Time, ) *BillingPeriod`

NewBillingPeriod instantiates a new BillingPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPeriodWithDefaults

`func NewBillingPeriodWithDefaults() *BillingPeriod`

NewBillingPeriodWithDefaults instantiates a new BillingPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *BillingPeriod) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BillingPeriod) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BillingPeriod) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *BillingPeriod) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *BillingPeriod) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *BillingPeriod) SetEnd(v time.Time)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


