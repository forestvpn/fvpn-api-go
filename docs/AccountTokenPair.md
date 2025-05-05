# AccountTokenPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** |  | [readonly] 
**Refresh** | **string** |  | [readonly] 

## Methods

### NewAccountTokenPair

`func NewAccountTokenPair(access string, refresh string, ) *AccountTokenPair`

NewAccountTokenPair instantiates a new AccountTokenPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenPairWithDefaults

`func NewAccountTokenPairWithDefaults() *AccountTokenPair`

NewAccountTokenPairWithDefaults instantiates a new AccountTokenPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AccountTokenPair) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccountTokenPair) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccountTokenPair) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetRefresh

`func (o *AccountTokenPair) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *AccountTokenPair) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *AccountTokenPair) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


