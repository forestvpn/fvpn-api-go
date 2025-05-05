# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Ipv4** | **string** |  | [readonly] 
**Ipv6** | **string** |  | [readonly] 
**PubKey** | **string** |  | 
**PsKey** | **string** |  | [readonly] 
**Password** | Pointer to **NullableString** | Bcrypt Hashed Password for Authentication, e.g. HTTP Proxy | [optional] 
**ConnectionStatus** | [**ConnectionStatus**](ConnectionStatus.md) | Determine connection status based on last activity time from data usage. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**LastActiveAt** | **time.Time** |  | [readonly] 
**Last30daysDataUsage** | [**DeviceDataUsage**](DeviceDataUsage.md) |  | [readonly] 
**QuickConfTemplate** | **string** | Generate a quick configuration string for the device. | [readonly] 

## Methods

### NewDevice

`func NewDevice(id string, name string, ipv4 string, ipv6 string, pubKey string, psKey string, connectionStatus ConnectionStatus, createdAt time.Time, lastActiveAt time.Time, last30daysDataUsage DeviceDataUsage, quickConfTemplate string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.


### GetIpv4

`func (o *Device) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Device) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Device) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *Device) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Device) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Device) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.


### GetPubKey

`func (o *Device) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *Device) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *Device) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetPsKey

`func (o *Device) GetPsKey() string`

GetPsKey returns the PsKey field if non-nil, zero value otherwise.

### GetPsKeyOk

`func (o *Device) GetPsKeyOk() (*string, bool)`

GetPsKeyOk returns a tuple with the PsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsKey

`func (o *Device) SetPsKey(v string)`

SetPsKey sets PsKey field to given value.


### GetPassword

`func (o *Device) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Device) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Device) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Device) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *Device) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Device) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetConnectionStatus

`func (o *Device) GetConnectionStatus() ConnectionStatus`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *Device) GetConnectionStatusOk() (*ConnectionStatus, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *Device) SetConnectionStatus(v ConnectionStatus)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetCreatedAt

`func (o *Device) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastActiveAt

`func (o *Device) GetLastActiveAt() time.Time`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *Device) GetLastActiveAtOk() (*time.Time, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *Device) SetLastActiveAt(v time.Time)`

SetLastActiveAt sets LastActiveAt field to given value.


### GetLast30daysDataUsage

`func (o *Device) GetLast30daysDataUsage() DeviceDataUsage`

GetLast30daysDataUsage returns the Last30daysDataUsage field if non-nil, zero value otherwise.

### GetLast30daysDataUsageOk

`func (o *Device) GetLast30daysDataUsageOk() (*DeviceDataUsage, bool)`

GetLast30daysDataUsageOk returns a tuple with the Last30daysDataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast30daysDataUsage

`func (o *Device) SetLast30daysDataUsage(v DeviceDataUsage)`

SetLast30daysDataUsage sets Last30daysDataUsage field to given value.


### GetQuickConfTemplate

`func (o *Device) GetQuickConfTemplate() string`

GetQuickConfTemplate returns the QuickConfTemplate field if non-nil, zero value otherwise.

### GetQuickConfTemplateOk

`func (o *Device) GetQuickConfTemplateOk() (*string, bool)`

GetQuickConfTemplateOk returns a tuple with the QuickConfTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickConfTemplate

`func (o *Device) SetQuickConfTemplate(v string)`

SetQuickConfTemplate sets QuickConfTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


