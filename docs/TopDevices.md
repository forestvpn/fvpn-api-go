# TopDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int64** |  | 
**Period** | [**Period**](Period.md) |  | 
**Devices** | [**[]DeviceData**](DeviceData.md) |  | 

## Methods

### NewTopDevices

`func NewTopDevices(limit int64, period Period, devices []DeviceData, ) *TopDevices`

NewTopDevices instantiates a new TopDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopDevicesWithDefaults

`func NewTopDevicesWithDefaults() *TopDevices`

NewTopDevicesWithDefaults instantiates a new TopDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *TopDevices) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TopDevices) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TopDevices) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetPeriod

`func (o *TopDevices) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TopDevices) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TopDevices) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetDevices

`func (o *TopDevices) GetDevices() []DeviceData`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *TopDevices) GetDevicesOk() (*[]DeviceData, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *TopDevices) SetDevices(v []DeviceData)`

SetDevices sets Devices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


