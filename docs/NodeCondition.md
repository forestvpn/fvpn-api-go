# NodeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**Condition**](Condition.md) |  | 
**Status** | **bool** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**LastHeartbeatAt** | **time.Time** |  | [readonly] 
**LastTransitionAt** | **time.Time** |  | [readonly] 

## Methods

### NewNodeCondition

`func NewNodeCondition(condition Condition, status bool, lastHeartbeatAt time.Time, lastTransitionAt time.Time, ) *NodeCondition`

NewNodeCondition instantiates a new NodeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConditionWithDefaults

`func NewNodeConditionWithDefaults() *NodeCondition`

NewNodeConditionWithDefaults instantiates a new NodeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *NodeCondition) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *NodeCondition) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *NodeCondition) SetCondition(v Condition)`

SetCondition sets Condition field to given value.


### GetStatus

`func (o *NodeCondition) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeCondition) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeCondition) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *NodeCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodeCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NodeCondition) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NodeCondition) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLastHeartbeatAt

`func (o *NodeCondition) GetLastHeartbeatAt() time.Time`

GetLastHeartbeatAt returns the LastHeartbeatAt field if non-nil, zero value otherwise.

### GetLastHeartbeatAtOk

`func (o *NodeCondition) GetLastHeartbeatAtOk() (*time.Time, bool)`

GetLastHeartbeatAtOk returns a tuple with the LastHeartbeatAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatAt

`func (o *NodeCondition) SetLastHeartbeatAt(v time.Time)`

SetLastHeartbeatAt sets LastHeartbeatAt field to given value.


### GetLastTransitionAt

`func (o *NodeCondition) GetLastTransitionAt() time.Time`

GetLastTransitionAt returns the LastTransitionAt field if non-nil, zero value otherwise.

### GetLastTransitionAtOk

`func (o *NodeCondition) GetLastTransitionAtOk() (*time.Time, bool)`

GetLastTransitionAtOk returns a tuple with the LastTransitionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionAt

`func (o *NodeCondition) SetLastTransitionAt(v time.Time)`

SetLastTransitionAt sets LastTransitionAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


