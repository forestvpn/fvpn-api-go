# AccountTokenRefresh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refresh** | Pointer to **string** |  | [optional] 
**Access** | **string** |  | [readonly] 

## Methods

### NewAccountTokenRefresh

`func NewAccountTokenRefresh(access string, ) *AccountTokenRefresh`

NewAccountTokenRefresh instantiates a new AccountTokenRefresh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenRefreshWithDefaults

`func NewAccountTokenRefreshWithDefaults() *AccountTokenRefresh`

NewAccountTokenRefreshWithDefaults instantiates a new AccountTokenRefresh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefresh

`func (o *AccountTokenRefresh) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *AccountTokenRefresh) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *AccountTokenRefresh) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *AccountTokenRefresh) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetAccess

`func (o *AccountTokenRefresh) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccountTokenRefresh) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccountTokenRefresh) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


