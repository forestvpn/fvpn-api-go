# WalletTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**TransactionType** | [**TransactionType**](TransactionType.md) |  | [readonly] 
**Credits** | **float64** | Transaction amount in credits | [readonly] 
**Amount** | **float32** | Return the amount in the specified currency. | [readonly] 
**Currency** | **string** | Currency of the transaction amount | [readonly] 
**Description** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewWalletTransaction

`func NewWalletTransaction(id string, transactionType TransactionType, credits float64, amount float32, currency string, description string, createdAt time.Time, ) *WalletTransaction`

NewWalletTransaction instantiates a new WalletTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionWithDefaults

`func NewWalletTransactionWithDefaults() *WalletTransaction`

NewWalletTransactionWithDefaults instantiates a new WalletTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTransactionType

`func (o *WalletTransaction) GetTransactionType() TransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *WalletTransaction) GetTransactionTypeOk() (*TransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *WalletTransaction) SetTransactionType(v TransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetCredits

`func (o *WalletTransaction) GetCredits() float64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *WalletTransaction) GetCreditsOk() (*float64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *WalletTransaction) SetCredits(v float64)`

SetCredits sets Credits field to given value.


### GetAmount

`func (o *WalletTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *WalletTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WalletTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WalletTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *WalletTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WalletTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WalletTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *WalletTransaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletTransaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletTransaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


