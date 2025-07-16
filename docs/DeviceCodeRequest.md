# DeviceCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DeviceCodeStatus**](DeviceCodeStatus.md) |  | [optional] 

## Methods

### NewDeviceCodeRequest

`func NewDeviceCodeRequest() *DeviceCodeRequest`

NewDeviceCodeRequest instantiates a new DeviceCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCodeRequestWithDefaults

`func NewDeviceCodeRequestWithDefaults() *DeviceCodeRequest`

NewDeviceCodeRequestWithDefaults instantiates a new DeviceCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeviceCodeRequest) GetStatus() DeviceCodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceCodeRequest) GetStatusOk() (*DeviceCodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceCodeRequest) SetStatus(v DeviceCodeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceCodeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


