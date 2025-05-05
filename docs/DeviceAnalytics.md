# DeviceAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**DeviceAnalyticsDevice**](DeviceAnalyticsDevice.md) |  | 
**Current** | [**SummaryPeriod**](SummaryPeriod.md) |  | 
**Previous** | Pointer to [**SummaryPeriod**](SummaryPeriod.md) |  | [optional] 
**PercentChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewDeviceAnalytics

`func NewDeviceAnalytics(device DeviceAnalyticsDevice, current SummaryPeriod, ) *DeviceAnalytics`

NewDeviceAnalytics instantiates a new DeviceAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAnalyticsWithDefaults

`func NewDeviceAnalyticsWithDefaults() *DeviceAnalytics`

NewDeviceAnalyticsWithDefaults instantiates a new DeviceAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *DeviceAnalytics) GetDevice() DeviceAnalyticsDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DeviceAnalytics) GetDeviceOk() (*DeviceAnalyticsDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DeviceAnalytics) SetDevice(v DeviceAnalyticsDevice)`

SetDevice sets Device field to given value.


### GetCurrent

`func (o *DeviceAnalytics) GetCurrent() SummaryPeriod`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *DeviceAnalytics) GetCurrentOk() (*SummaryPeriod, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *DeviceAnalytics) SetCurrent(v SummaryPeriod)`

SetCurrent sets Current field to given value.


### GetPrevious

`func (o *DeviceAnalytics) GetPrevious() SummaryPeriod`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DeviceAnalytics) GetPreviousOk() (*SummaryPeriod, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DeviceAnalytics) SetPrevious(v SummaryPeriod)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DeviceAnalytics) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetPercentChange

`func (o *DeviceAnalytics) GetPercentChange() float64`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *DeviceAnalytics) GetPercentChangeOk() (*float64, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *DeviceAnalytics) SetPercentChange(v float64)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *DeviceAnalytics) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


