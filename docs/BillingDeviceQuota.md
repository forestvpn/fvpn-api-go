# BillingDeviceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsage** | **int64** |  | [readonly] 
**Quota** | **int64** |  | [readonly] 
**UsagePercentage** | **float64** |  | [readonly] 
**IsWithinQuota** | **bool** |  | [readonly] 
**IsUnlimited** | **bool** |  | [readonly] 
**Message** | **string** |  | [readonly] 

## Methods

### NewBillingDeviceQuota

`func NewBillingDeviceQuota(currentUsage int64, quota int64, usagePercentage float64, isWithinQuota bool, isUnlimited bool, message string, ) *BillingDeviceQuota`

NewBillingDeviceQuota instantiates a new BillingDeviceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDeviceQuotaWithDefaults

`func NewBillingDeviceQuotaWithDefaults() *BillingDeviceQuota`

NewBillingDeviceQuotaWithDefaults instantiates a new BillingDeviceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsage

`func (o *BillingDeviceQuota) GetCurrentUsage() int64`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *BillingDeviceQuota) GetCurrentUsageOk() (*int64, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *BillingDeviceQuota) SetCurrentUsage(v int64)`

SetCurrentUsage sets CurrentUsage field to given value.


### GetQuota

`func (o *BillingDeviceQuota) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *BillingDeviceQuota) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *BillingDeviceQuota) SetQuota(v int64)`

SetQuota sets Quota field to given value.


### GetUsagePercentage

`func (o *BillingDeviceQuota) GetUsagePercentage() float64`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *BillingDeviceQuota) GetUsagePercentageOk() (*float64, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *BillingDeviceQuota) SetUsagePercentage(v float64)`

SetUsagePercentage sets UsagePercentage field to given value.


### GetIsWithinQuota

`func (o *BillingDeviceQuota) GetIsWithinQuota() bool`

GetIsWithinQuota returns the IsWithinQuota field if non-nil, zero value otherwise.

### GetIsWithinQuotaOk

`func (o *BillingDeviceQuota) GetIsWithinQuotaOk() (*bool, bool)`

GetIsWithinQuotaOk returns a tuple with the IsWithinQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWithinQuota

`func (o *BillingDeviceQuota) SetIsWithinQuota(v bool)`

SetIsWithinQuota sets IsWithinQuota field to given value.


### GetIsUnlimited

`func (o *BillingDeviceQuota) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *BillingDeviceQuota) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *BillingDeviceQuota) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.


### GetMessage

`func (o *BillingDeviceQuota) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BillingDeviceQuota) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BillingDeviceQuota) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


