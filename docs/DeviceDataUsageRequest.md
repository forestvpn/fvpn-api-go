# DeviceDataUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rx** | **int64** |  | 
**Tx** | **int64** |  | 
**Total** | **int64** |  | 

## Methods

### NewDeviceDataUsageRequest

`func NewDeviceDataUsageRequest(rx int64, tx int64, total int64, ) *DeviceDataUsageRequest`

NewDeviceDataUsageRequest instantiates a new DeviceDataUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDataUsageRequestWithDefaults

`func NewDeviceDataUsageRequestWithDefaults() *DeviceDataUsageRequest`

NewDeviceDataUsageRequestWithDefaults instantiates a new DeviceDataUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRx

`func (o *DeviceDataUsageRequest) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *DeviceDataUsageRequest) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *DeviceDataUsageRequest) SetRx(v int64)`

SetRx sets Rx field to given value.


### GetTx

`func (o *DeviceDataUsageRequest) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *DeviceDataUsageRequest) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *DeviceDataUsageRequest) SetTx(v int64)`

SetTx sets Tx field to given value.


### GetTotal

`func (o *DeviceDataUsageRequest) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DeviceDataUsageRequest) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DeviceDataUsageRequest) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


