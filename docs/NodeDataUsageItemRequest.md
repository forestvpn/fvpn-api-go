# NodeDataUsageItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the record, it may be pub_key, ipv4, ipv6 or device_id | 
**Proto** | **string** | Protocol used for the record | 
**Tx** | Pointer to **int64** | Total bytes sent | [optional] 
**Rx** | Pointer to **int64** | Total bytes received | [optional] 

## Methods

### NewNodeDataUsageItemRequest

`func NewNodeDataUsageItemRequest(id string, proto string, ) *NodeDataUsageItemRequest`

NewNodeDataUsageItemRequest instantiates a new NodeDataUsageItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDataUsageItemRequestWithDefaults

`func NewNodeDataUsageItemRequestWithDefaults() *NodeDataUsageItemRequest`

NewNodeDataUsageItemRequestWithDefaults instantiates a new NodeDataUsageItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeDataUsageItemRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeDataUsageItemRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeDataUsageItemRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProto

`func (o *NodeDataUsageItemRequest) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *NodeDataUsageItemRequest) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *NodeDataUsageItemRequest) SetProto(v string)`

SetProto sets Proto field to given value.


### GetTx

`func (o *NodeDataUsageItemRequest) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NodeDataUsageItemRequest) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NodeDataUsageItemRequest) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *NodeDataUsageItemRequest) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetRx

`func (o *NodeDataUsageItemRequest) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *NodeDataUsageItemRequest) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *NodeDataUsageItemRequest) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *NodeDataUsageItemRequest) HasRx() bool`

HasRx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


