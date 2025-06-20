# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Name** | **string** |  | 
**Symbol** | Pointer to **string** |  | [optional] 
**Subunit** | Pointer to **int64** | Number of subunits in the currency | [optional] 

## Methods

### NewCurrency

`func NewCurrency(code string, name string, ) *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Currency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Currency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Currency) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *Currency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Currency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Currency) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *Currency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Currency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Currency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Currency) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSubunit

`func (o *Currency) GetSubunit() int64`

GetSubunit returns the Subunit field if non-nil, zero value otherwise.

### GetSubunitOk

`func (o *Currency) GetSubunitOk() (*int64, bool)`

GetSubunitOk returns a tuple with the Subunit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubunit

`func (o *Currency) SetSubunit(v int64)`

SetSubunit sets Subunit field to given value.

### HasSubunit

`func (o *Currency) HasSubunit() bool`

HasSubunit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


