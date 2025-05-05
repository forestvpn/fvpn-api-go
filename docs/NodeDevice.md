# NodeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | **string** |  | [readonly] 
**Ipv6** | **string** |  | [readonly] 
**PubKey** | **string** |  | [readonly] 
**PsKey** | **string** |  | [readonly] 

## Methods

### NewNodeDevice

`func NewNodeDevice(ipv4 string, ipv6 string, pubKey string, psKey string, ) *NodeDevice`

NewNodeDevice instantiates a new NodeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDeviceWithDefaults

`func NewNodeDeviceWithDefaults() *NodeDevice`

NewNodeDeviceWithDefaults instantiates a new NodeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *NodeDevice) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *NodeDevice) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *NodeDevice) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *NodeDevice) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *NodeDevice) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *NodeDevice) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.


### GetPubKey

`func (o *NodeDevice) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *NodeDevice) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *NodeDevice) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetPsKey

`func (o *NodeDevice) GetPsKey() string`

GetPsKey returns the PsKey field if non-nil, zero value otherwise.

### GetPsKeyOk

`func (o *NodeDevice) GetPsKeyOk() (*string, bool)`

GetPsKeyOk returns a tuple with the PsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsKey

`func (o *NodeDevice) SetPsKey(v string)`

SetPsKey sets PsKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


