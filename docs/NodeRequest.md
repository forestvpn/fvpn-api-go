# NodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**IpAddresses** | **[]string** |  | 
**Subnets** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**PubKey** | **string** |  | 
**Ports** | **[]string** |  | 
**HttpsProxyPorts** | Pointer to **[]string** | Ports used for HTTPs Proxy | [optional] 
**SsBridgePorts** | Pointer to **[]string** | Ports used for Shadow Socket Bridge | [optional] 
**Status** | Pointer to [**NodeStatus**](NodeStatus.md) |  | [optional] 
**Os** | Pointer to **NullableString** |  | [optional] 
**OsVersion** | Pointer to **NullableString** |  | [optional] 
**OsArch** | Pointer to **NullableString** |  | [optional] 
**Distro** | Pointer to **NullableString** |  | [optional] 
**DistroVersion** | Pointer to **NullableString** |  | [optional] 
**DistroCodename** | Pointer to **NullableString** |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**AppBuild** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNodeRequest

`func NewNodeRequest(hostname string, ipAddresses []string, pubKey string, ports []string, ) *NodeRequest`

NewNodeRequest instantiates a new NodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRequestWithDefaults

`func NewNodeRequestWithDefaults() *NodeRequest`

NewNodeRequestWithDefaults instantiates a new NodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *NodeRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIpAddresses

`func (o *NodeRequest) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NodeRequest) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NodeRequest) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetSubnets

`func (o *NodeRequest) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *NodeRequest) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *NodeRequest) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *NodeRequest) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *NodeRequest) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *NodeRequest) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetTags

`func (o *NodeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NodeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsPublic

`func (o *NodeRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *NodeRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *NodeRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *NodeRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetPubKey

`func (o *NodeRequest) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *NodeRequest) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *NodeRequest) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetPorts

`func (o *NodeRequest) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NodeRequest) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NodeRequest) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetHttpsProxyPorts

`func (o *NodeRequest) GetHttpsProxyPorts() []string`

GetHttpsProxyPorts returns the HttpsProxyPorts field if non-nil, zero value otherwise.

### GetHttpsProxyPortsOk

`func (o *NodeRequest) GetHttpsProxyPortsOk() (*[]string, bool)`

GetHttpsProxyPortsOk returns a tuple with the HttpsProxyPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyPorts

`func (o *NodeRequest) SetHttpsProxyPorts(v []string)`

SetHttpsProxyPorts sets HttpsProxyPorts field to given value.

### HasHttpsProxyPorts

`func (o *NodeRequest) HasHttpsProxyPorts() bool`

HasHttpsProxyPorts returns a boolean if a field has been set.

### SetHttpsProxyPortsNil

`func (o *NodeRequest) SetHttpsProxyPortsNil(b bool)`

 SetHttpsProxyPortsNil sets the value for HttpsProxyPorts to be an explicit nil

### UnsetHttpsProxyPorts
`func (o *NodeRequest) UnsetHttpsProxyPorts()`

UnsetHttpsProxyPorts ensures that no value is present for HttpsProxyPorts, not even an explicit nil
### GetSsBridgePorts

`func (o *NodeRequest) GetSsBridgePorts() []string`

GetSsBridgePorts returns the SsBridgePorts field if non-nil, zero value otherwise.

### GetSsBridgePortsOk

`func (o *NodeRequest) GetSsBridgePortsOk() (*[]string, bool)`

GetSsBridgePortsOk returns a tuple with the SsBridgePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsBridgePorts

`func (o *NodeRequest) SetSsBridgePorts(v []string)`

SetSsBridgePorts sets SsBridgePorts field to given value.

### HasSsBridgePorts

`func (o *NodeRequest) HasSsBridgePorts() bool`

HasSsBridgePorts returns a boolean if a field has been set.

### SetSsBridgePortsNil

`func (o *NodeRequest) SetSsBridgePortsNil(b bool)`

 SetSsBridgePortsNil sets the value for SsBridgePorts to be an explicit nil

### UnsetSsBridgePorts
`func (o *NodeRequest) UnsetSsBridgePorts()`

UnsetSsBridgePorts ensures that no value is present for SsBridgePorts, not even an explicit nil
### GetStatus

`func (o *NodeRequest) GetStatus() NodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeRequest) GetStatusOk() (*NodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeRequest) SetStatus(v NodeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOs

`func (o *NodeRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *NodeRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *NodeRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *NodeRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *NodeRequest) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *NodeRequest) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetOsVersion

`func (o *NodeRequest) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *NodeRequest) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *NodeRequest) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *NodeRequest) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *NodeRequest) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *NodeRequest) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetOsArch

`func (o *NodeRequest) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *NodeRequest) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *NodeRequest) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *NodeRequest) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### SetOsArchNil

`func (o *NodeRequest) SetOsArchNil(b bool)`

 SetOsArchNil sets the value for OsArch to be an explicit nil

### UnsetOsArch
`func (o *NodeRequest) UnsetOsArch()`

UnsetOsArch ensures that no value is present for OsArch, not even an explicit nil
### GetDistro

`func (o *NodeRequest) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *NodeRequest) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *NodeRequest) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *NodeRequest) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *NodeRequest) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *NodeRequest) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *NodeRequest) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *NodeRequest) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *NodeRequest) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *NodeRequest) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### SetDistroVersionNil

`func (o *NodeRequest) SetDistroVersionNil(b bool)`

 SetDistroVersionNil sets the value for DistroVersion to be an explicit nil

### UnsetDistroVersion
`func (o *NodeRequest) UnsetDistroVersion()`

UnsetDistroVersion ensures that no value is present for DistroVersion, not even an explicit nil
### GetDistroCodename

`func (o *NodeRequest) GetDistroCodename() string`

GetDistroCodename returns the DistroCodename field if non-nil, zero value otherwise.

### GetDistroCodenameOk

`func (o *NodeRequest) GetDistroCodenameOk() (*string, bool)`

GetDistroCodenameOk returns a tuple with the DistroCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroCodename

`func (o *NodeRequest) SetDistroCodename(v string)`

SetDistroCodename sets DistroCodename field to given value.

### HasDistroCodename

`func (o *NodeRequest) HasDistroCodename() bool`

HasDistroCodename returns a boolean if a field has been set.

### SetDistroCodenameNil

`func (o *NodeRequest) SetDistroCodenameNil(b bool)`

 SetDistroCodenameNil sets the value for DistroCodename to be an explicit nil

### UnsetDistroCodename
`func (o *NodeRequest) UnsetDistroCodename()`

UnsetDistroCodename ensures that no value is present for DistroCodename, not even an explicit nil
### GetAppVersion

`func (o *NodeRequest) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *NodeRequest) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *NodeRequest) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *NodeRequest) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *NodeRequest) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *NodeRequest) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetAppBuild

`func (o *NodeRequest) GetAppBuild() string`

GetAppBuild returns the AppBuild field if non-nil, zero value otherwise.

### GetAppBuildOk

`func (o *NodeRequest) GetAppBuildOk() (*string, bool)`

GetAppBuildOk returns a tuple with the AppBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuild

`func (o *NodeRequest) SetAppBuild(v string)`

SetAppBuild sets AppBuild field to given value.

### HasAppBuild

`func (o *NodeRequest) HasAppBuild() bool`

HasAppBuild returns a boolean if a field has been set.

### SetAppBuildNil

`func (o *NodeRequest) SetAppBuildNil(b bool)`

 SetAppBuildNil sets the value for AppBuild to be an explicit nil

### UnsetAppBuild
`func (o *NodeRequest) UnsetAppBuild()`

UnsetAppBuild ensures that no value is present for AppBuild, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


