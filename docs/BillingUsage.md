# BillingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BillingDataQuota**](BillingDataQuota.md) |  | [readonly] 
**Device** | [**BillingDeviceQuota**](BillingDeviceQuota.md) |  | [readonly] 

## Methods

### NewBillingUsage

`func NewBillingUsage(data BillingDataQuota, device BillingDeviceQuota, ) *BillingUsage`

NewBillingUsage instantiates a new BillingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUsageWithDefaults

`func NewBillingUsageWithDefaults() *BillingUsage`

NewBillingUsageWithDefaults instantiates a new BillingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BillingUsage) GetData() BillingDataQuota`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingUsage) GetDataOk() (*BillingDataQuota, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingUsage) SetData(v BillingDataQuota)`

SetData sets Data field to given value.


### GetDevice

`func (o *BillingUsage) GetDevice() BillingDeviceQuota`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BillingUsage) GetDeviceOk() (*BillingDeviceQuota, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BillingUsage) SetDevice(v BillingDeviceQuota)`

SetDevice sets Device field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


