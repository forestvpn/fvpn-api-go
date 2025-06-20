# BillingDataQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsageGb** | **float64** |  | [readonly] 
**QuotaGb** | **int64** |  | [readonly] 
**UsagePercentage** | **float64** |  | [readonly] 
**IsWithinQuota** | **bool** |  | [readonly] 
**IsUnlimited** | **bool** |  | [readonly] 
**WarningLevel** | **NullableString** |  | [readonly] 
**Message** | **string** |  | [readonly] 
**Period** | [**BillingPeriod**](BillingPeriod.md) |  | [readonly] 

## Methods

### NewBillingDataQuota

`func NewBillingDataQuota(currentUsageGb float64, quotaGb int64, usagePercentage float64, isWithinQuota bool, isUnlimited bool, warningLevel NullableString, message string, period BillingPeriod, ) *BillingDataQuota`

NewBillingDataQuota instantiates a new BillingDataQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDataQuotaWithDefaults

`func NewBillingDataQuotaWithDefaults() *BillingDataQuota`

NewBillingDataQuotaWithDefaults instantiates a new BillingDataQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsageGb

`func (o *BillingDataQuota) GetCurrentUsageGb() float64`

GetCurrentUsageGb returns the CurrentUsageGb field if non-nil, zero value otherwise.

### GetCurrentUsageGbOk

`func (o *BillingDataQuota) GetCurrentUsageGbOk() (*float64, bool)`

GetCurrentUsageGbOk returns a tuple with the CurrentUsageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsageGb

`func (o *BillingDataQuota) SetCurrentUsageGb(v float64)`

SetCurrentUsageGb sets CurrentUsageGb field to given value.


### GetQuotaGb

`func (o *BillingDataQuota) GetQuotaGb() int64`

GetQuotaGb returns the QuotaGb field if non-nil, zero value otherwise.

### GetQuotaGbOk

`func (o *BillingDataQuota) GetQuotaGbOk() (*int64, bool)`

GetQuotaGbOk returns a tuple with the QuotaGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaGb

`func (o *BillingDataQuota) SetQuotaGb(v int64)`

SetQuotaGb sets QuotaGb field to given value.


### GetUsagePercentage

`func (o *BillingDataQuota) GetUsagePercentage() float64`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *BillingDataQuota) GetUsagePercentageOk() (*float64, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *BillingDataQuota) SetUsagePercentage(v float64)`

SetUsagePercentage sets UsagePercentage field to given value.


### GetIsWithinQuota

`func (o *BillingDataQuota) GetIsWithinQuota() bool`

GetIsWithinQuota returns the IsWithinQuota field if non-nil, zero value otherwise.

### GetIsWithinQuotaOk

`func (o *BillingDataQuota) GetIsWithinQuotaOk() (*bool, bool)`

GetIsWithinQuotaOk returns a tuple with the IsWithinQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWithinQuota

`func (o *BillingDataQuota) SetIsWithinQuota(v bool)`

SetIsWithinQuota sets IsWithinQuota field to given value.


### GetIsUnlimited

`func (o *BillingDataQuota) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *BillingDataQuota) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *BillingDataQuota) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.


### GetWarningLevel

`func (o *BillingDataQuota) GetWarningLevel() string`

GetWarningLevel returns the WarningLevel field if non-nil, zero value otherwise.

### GetWarningLevelOk

`func (o *BillingDataQuota) GetWarningLevelOk() (*string, bool)`

GetWarningLevelOk returns a tuple with the WarningLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLevel

`func (o *BillingDataQuota) SetWarningLevel(v string)`

SetWarningLevel sets WarningLevel field to given value.


### SetWarningLevelNil

`func (o *BillingDataQuota) SetWarningLevelNil(b bool)`

 SetWarningLevelNil sets the value for WarningLevel to be an explicit nil

### UnsetWarningLevel
`func (o *BillingDataQuota) UnsetWarningLevel()`

UnsetWarningLevel ensures that no value is present for WarningLevel, not even an explicit nil
### GetMessage

`func (o *BillingDataQuota) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BillingDataQuota) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BillingDataQuota) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPeriod

`func (o *BillingDataQuota) GetPeriod() BillingPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BillingDataQuota) GetPeriodOk() (*BillingPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BillingDataQuota) SetPeriod(v BillingPeriod)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


