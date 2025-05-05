# PriceDiffInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** |  | 
**Currency** | **string** |  | 
**Interval** | **string** |  | 
**Discount** | **int64** |  | 
**SaveAmount** | **float64** |  | 
**SaveAmountTotal** | **int64** |  | 

## Methods

### NewPriceDiffInfo

`func NewPriceDiffInfo(amount float64, currency string, interval string, discount int64, saveAmount float64, saveAmountTotal int64, ) *PriceDiffInfo`

NewPriceDiffInfo instantiates a new PriceDiffInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDiffInfoWithDefaults

`func NewPriceDiffInfoWithDefaults() *PriceDiffInfo`

NewPriceDiffInfoWithDefaults instantiates a new PriceDiffInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PriceDiffInfo) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PriceDiffInfo) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PriceDiffInfo) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PriceDiffInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceDiffInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceDiffInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInterval

`func (o *PriceDiffInfo) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PriceDiffInfo) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PriceDiffInfo) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetDiscount

`func (o *PriceDiffInfo) GetDiscount() int64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *PriceDiffInfo) GetDiscountOk() (*int64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *PriceDiffInfo) SetDiscount(v int64)`

SetDiscount sets Discount field to given value.


### GetSaveAmount

`func (o *PriceDiffInfo) GetSaveAmount() float64`

GetSaveAmount returns the SaveAmount field if non-nil, zero value otherwise.

### GetSaveAmountOk

`func (o *PriceDiffInfo) GetSaveAmountOk() (*float64, bool)`

GetSaveAmountOk returns a tuple with the SaveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAmount

`func (o *PriceDiffInfo) SetSaveAmount(v float64)`

SetSaveAmount sets SaveAmount field to given value.


### GetSaveAmountTotal

`func (o *PriceDiffInfo) GetSaveAmountTotal() int64`

GetSaveAmountTotal returns the SaveAmountTotal field if non-nil, zero value otherwise.

### GetSaveAmountTotalOk

`func (o *PriceDiffInfo) GetSaveAmountTotalOk() (*int64, bool)`

GetSaveAmountTotalOk returns a tuple with the SaveAmountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAmountTotal

`func (o *PriceDiffInfo) SetSaveAmountTotal(v int64)`

SetSaveAmountTotal sets SaveAmountTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


