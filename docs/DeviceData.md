# DeviceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Data** | [**[]DataPoint**](DataPoint.md) |  | 

## Methods

### NewDeviceData

`func NewDeviceData(id string, name string, data []DataPoint, ) *DeviceData`

NewDeviceData instantiates a new DeviceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDataWithDefaults

`func NewDeviceDataWithDefaults() *DeviceData`

NewDeviceDataWithDefaults instantiates a new DeviceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeviceData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceData) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *DeviceData) GetData() []DataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceData) GetDataOk() (*[]DataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceData) SetData(v []DataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


