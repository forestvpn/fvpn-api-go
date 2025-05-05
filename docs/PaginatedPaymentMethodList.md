# PaginatedPaymentMethodList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]PaymentMethod**](PaymentMethod.md) |  | 
**Count** | Pointer to **int32** | Total count of items | [optional] 

## Methods

### NewPaginatedPaymentMethodList

`func NewPaginatedPaymentMethodList(results []PaymentMethod, ) *PaginatedPaymentMethodList`

NewPaginatedPaymentMethodList instantiates a new PaginatedPaymentMethodList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPaymentMethodListWithDefaults

`func NewPaginatedPaymentMethodListWithDefaults() *PaginatedPaymentMethodList`

NewPaginatedPaymentMethodListWithDefaults instantiates a new PaginatedPaymentMethodList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedPaymentMethodList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedPaymentMethodList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedPaymentMethodList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedPaymentMethodList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedPaymentMethodList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedPaymentMethodList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedPaymentMethodList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedPaymentMethodList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedPaymentMethodList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedPaymentMethodList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedPaymentMethodList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedPaymentMethodList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedPaymentMethodList) GetResults() []PaymentMethod`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedPaymentMethodList) GetResultsOk() (*[]PaymentMethod, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedPaymentMethodList) SetResults(v []PaymentMethod)`

SetResults sets Results field to given value.


### GetCount

`func (o *PaginatedPaymentMethodList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedPaymentMethodList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedPaymentMethodList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedPaymentMethodList) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


