# NodeDataUsageReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Items** | [**[]NodeDataUsageItemRequest**](NodeDataUsageItemRequest.md) |  | 

## Methods

### NewNodeDataUsageReportRequest

`func NewNodeDataUsageReportRequest(timestamp time.Time, items []NodeDataUsageItemRequest, ) *NodeDataUsageReportRequest`

NewNodeDataUsageReportRequest instantiates a new NodeDataUsageReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDataUsageReportRequestWithDefaults

`func NewNodeDataUsageReportRequestWithDefaults() *NodeDataUsageReportRequest`

NewNodeDataUsageReportRequestWithDefaults instantiates a new NodeDataUsageReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *NodeDataUsageReportRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NodeDataUsageReportRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NodeDataUsageReportRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetItems

`func (o *NodeDataUsageReportRequest) GetItems() []NodeDataUsageItemRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NodeDataUsageReportRequest) GetItemsOk() (*[]NodeDataUsageItemRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NodeDataUsageReportRequest) SetItems(v []NodeDataUsageItemRequest)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


