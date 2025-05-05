# NodeConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**Condition**](Condition.md) |  | 
**Status** | **bool** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNodeConditionRequest

`func NewNodeConditionRequest(condition Condition, status bool, ) *NodeConditionRequest`

NewNodeConditionRequest instantiates a new NodeConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConditionRequestWithDefaults

`func NewNodeConditionRequestWithDefaults() *NodeConditionRequest`

NewNodeConditionRequestWithDefaults instantiates a new NodeConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *NodeConditionRequest) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *NodeConditionRequest) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *NodeConditionRequest) SetCondition(v Condition)`

SetCondition sets Condition field to given value.


### GetStatus

`func (o *NodeConditionRequest) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeConditionRequest) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeConditionRequest) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *NodeConditionRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeConditionRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeConditionRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodeConditionRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NodeConditionRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NodeConditionRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


