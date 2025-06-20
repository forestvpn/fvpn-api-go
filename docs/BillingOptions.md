# BillingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableCurrencies** | Pointer to [**[]Currency**](Currency.md) |  | [optional] 
**DefaultCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewBillingOptions

`func NewBillingOptions() *BillingOptions`

NewBillingOptions instantiates a new BillingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingOptionsWithDefaults

`func NewBillingOptionsWithDefaults() *BillingOptions`

NewBillingOptionsWithDefaults instantiates a new BillingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableCurrencies

`func (o *BillingOptions) GetAvailableCurrencies() []Currency`

GetAvailableCurrencies returns the AvailableCurrencies field if non-nil, zero value otherwise.

### GetAvailableCurrenciesOk

`func (o *BillingOptions) GetAvailableCurrenciesOk() (*[]Currency, bool)`

GetAvailableCurrenciesOk returns a tuple with the AvailableCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCurrencies

`func (o *BillingOptions) SetAvailableCurrencies(v []Currency)`

SetAvailableCurrencies sets AvailableCurrencies field to given value.

### HasAvailableCurrencies

`func (o *BillingOptions) HasAvailableCurrencies() bool`

HasAvailableCurrencies returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *BillingOptions) GetDefaultCurrency() Currency`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *BillingOptions) GetDefaultCurrencyOk() (*Currency, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *BillingOptions) SetDefaultCurrency(v Currency)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *BillingOptions) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


