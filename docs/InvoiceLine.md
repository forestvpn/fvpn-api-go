# InvoiceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Description** | **string** |  | 
**UnitAmount** | **float64** |  | [readonly] 
**TaxAmount** | **float64** |  | [readonly] 
**DiscountAmount** | **float64** |  | [readonly] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Amount** | **float64** |  | [readonly] 
**Currency** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewInvoiceLine

`func NewInvoiceLine(id string, description string, unitAmount float64, taxAmount float64, discountAmount float64, amount float64, currency string, createdAt time.Time, ) *InvoiceLine`

NewInvoiceLine instantiates a new InvoiceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineWithDefaults

`func NewInvoiceLineWithDefaults() *InvoiceLine`

NewInvoiceLineWithDefaults instantiates a new InvoiceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLine) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *InvoiceLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnitAmount

`func (o *InvoiceLine) GetUnitAmount() float64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *InvoiceLine) GetUnitAmountOk() (*float64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *InvoiceLine) SetUnitAmount(v float64)`

SetUnitAmount sets UnitAmount field to given value.


### GetTaxAmount

`func (o *InvoiceLine) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *InvoiceLine) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *InvoiceLine) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.


### GetDiscountAmount

`func (o *InvoiceLine) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InvoiceLine) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InvoiceLine) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetQuantity

`func (o *InvoiceLine) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLine) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLine) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAmount

`func (o *InvoiceLine) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceLine) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceLine) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *InvoiceLine) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLine) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLine) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *InvoiceLine) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceLine) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceLine) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


