# DeviceCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 6-character device authentication code | [readonly] 
**ExpiresAt** | **time.Time** | When this code expires | [readonly] 
**Status** | Pointer to [**DeviceCodeStatus**](DeviceCodeStatus.md) |  | [optional] 
**AccessToken** | **NullableString** | JWT access token generated after approval | [readonly] 
**RefreshToken** | **NullableString** | JWT refresh token generated after approval | [readonly] 
**ApprovalUrl** | **string** | Generate the approval URL for the device code. | [readonly] 

## Methods

### NewDeviceCode

`func NewDeviceCode(code string, expiresAt time.Time, accessToken NullableString, refreshToken NullableString, approvalUrl string, ) *DeviceCode`

NewDeviceCode instantiates a new DeviceCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCodeWithDefaults

`func NewDeviceCodeWithDefaults() *DeviceCode`

NewDeviceCodeWithDefaults instantiates a new DeviceCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DeviceCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeviceCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeviceCode) SetCode(v string)`

SetCode sets Code field to given value.


### GetExpiresAt

`func (o *DeviceCode) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DeviceCode) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DeviceCode) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetStatus

`func (o *DeviceCode) GetStatus() DeviceCodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceCode) GetStatusOk() (*DeviceCodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceCode) SetStatus(v DeviceCodeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceCode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccessToken

`func (o *DeviceCode) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DeviceCode) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DeviceCode) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### SetAccessTokenNil

`func (o *DeviceCode) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *DeviceCode) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetRefreshToken

`func (o *DeviceCode) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *DeviceCode) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *DeviceCode) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### SetRefreshTokenNil

`func (o *DeviceCode) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *DeviceCode) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetApprovalUrl

`func (o *DeviceCode) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *DeviceCode) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *DeviceCode) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


