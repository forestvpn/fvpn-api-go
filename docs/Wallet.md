# Wallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Status** | [**WalletStatus**](WalletStatus.md) |  | [readonly] 
**Currency** | **string** | Currency for this wallet in ISO 4217 format | [readonly] 
**Balance** | **float64** |  | [readonly] 
**CreditsBalance** | **float64** |  | [readonly] 
**ConsumedBalance** | **float64** |  | [readonly] 
**IsActive** | **bool** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**RecentTransactions** | [**[]WalletTransaction**](WalletTransaction.md) |  | [readonly] 

## Methods

### NewWallet

`func NewWallet(id string, name string, status WalletStatus, currency string, balance float64, creditsBalance float64, consumedBalance float64, isActive bool, createdAt time.Time, updatedAt time.Time, recentTransactions []WalletTransaction, ) *Wallet`

NewWallet instantiates a new Wallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithDefaults

`func NewWalletWithDefaults() *Wallet`

NewWalletWithDefaults instantiates a new Wallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Wallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wallet) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Wallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Wallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Wallet) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Wallet) GetStatus() WalletStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Wallet) GetStatusOk() (*WalletStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Wallet) SetStatus(v WalletStatus)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *Wallet) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Wallet) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Wallet) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *Wallet) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Wallet) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Wallet) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetCreditsBalance

`func (o *Wallet) GetCreditsBalance() float64`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *Wallet) GetCreditsBalanceOk() (*float64, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *Wallet) SetCreditsBalance(v float64)`

SetCreditsBalance sets CreditsBalance field to given value.


### GetConsumedBalance

`func (o *Wallet) GetConsumedBalance() float64`

GetConsumedBalance returns the ConsumedBalance field if non-nil, zero value otherwise.

### GetConsumedBalanceOk

`func (o *Wallet) GetConsumedBalanceOk() (*float64, bool)`

GetConsumedBalanceOk returns a tuple with the ConsumedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedBalance

`func (o *Wallet) SetConsumedBalance(v float64)`

SetConsumedBalance sets ConsumedBalance field to given value.


### GetIsActive

`func (o *Wallet) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Wallet) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Wallet) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedAt

`func (o *Wallet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Wallet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Wallet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Wallet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Wallet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Wallet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRecentTransactions

`func (o *Wallet) GetRecentTransactions() []WalletTransaction`

GetRecentTransactions returns the RecentTransactions field if non-nil, zero value otherwise.

### GetRecentTransactionsOk

`func (o *Wallet) GetRecentTransactionsOk() (*[]WalletTransaction, bool)`

GetRecentTransactionsOk returns a tuple with the RecentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTransactions

`func (o *Wallet) SetRecentTransactions(v []WalletTransaction)`

SetRecentTransactions sets RecentTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


