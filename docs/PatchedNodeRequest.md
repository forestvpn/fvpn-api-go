# PatchedNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**Subnets** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**PubKey** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **[]string** |  | [optional] 
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

### NewPatchedNodeRequest

`func NewPatchedNodeRequest() *PatchedNodeRequest`

NewPatchedNodeRequest instantiates a new PatchedNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNodeRequestWithDefaults

`func NewPatchedNodeRequestWithDefaults() *PatchedNodeRequest`

NewPatchedNodeRequestWithDefaults instantiates a new PatchedNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *PatchedNodeRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PatchedNodeRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PatchedNodeRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PatchedNodeRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddresses

`func (o *PatchedNodeRequest) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *PatchedNodeRequest) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *PatchedNodeRequest) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *PatchedNodeRequest) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetSubnets

`func (o *PatchedNodeRequest) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *PatchedNodeRequest) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *PatchedNodeRequest) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *PatchedNodeRequest) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *PatchedNodeRequest) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *PatchedNodeRequest) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetTags

`func (o *PatchedNodeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedNodeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedNodeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedNodeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsPublic

`func (o *PatchedNodeRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PatchedNodeRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PatchedNodeRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *PatchedNodeRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetPubKey

`func (o *PatchedNodeRequest) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *PatchedNodeRequest) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *PatchedNodeRequest) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.

### HasPubKey

`func (o *PatchedNodeRequest) HasPubKey() bool`

HasPubKey returns a boolean if a field has been set.

### GetPorts

`func (o *PatchedNodeRequest) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PatchedNodeRequest) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PatchedNodeRequest) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PatchedNodeRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetHttpsProxyPorts

`func (o *PatchedNodeRequest) GetHttpsProxyPorts() []string`

GetHttpsProxyPorts returns the HttpsProxyPorts field if non-nil, zero value otherwise.

### GetHttpsProxyPortsOk

`func (o *PatchedNodeRequest) GetHttpsProxyPortsOk() (*[]string, bool)`

GetHttpsProxyPortsOk returns a tuple with the HttpsProxyPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyPorts

`func (o *PatchedNodeRequest) SetHttpsProxyPorts(v []string)`

SetHttpsProxyPorts sets HttpsProxyPorts field to given value.

### HasHttpsProxyPorts

`func (o *PatchedNodeRequest) HasHttpsProxyPorts() bool`

HasHttpsProxyPorts returns a boolean if a field has been set.

### SetHttpsProxyPortsNil

`func (o *PatchedNodeRequest) SetHttpsProxyPortsNil(b bool)`

 SetHttpsProxyPortsNil sets the value for HttpsProxyPorts to be an explicit nil

### UnsetHttpsProxyPorts
`func (o *PatchedNodeRequest) UnsetHttpsProxyPorts()`

UnsetHttpsProxyPorts ensures that no value is present for HttpsProxyPorts, not even an explicit nil
### GetSsBridgePorts

`func (o *PatchedNodeRequest) GetSsBridgePorts() []string`

GetSsBridgePorts returns the SsBridgePorts field if non-nil, zero value otherwise.

### GetSsBridgePortsOk

`func (o *PatchedNodeRequest) GetSsBridgePortsOk() (*[]string, bool)`

GetSsBridgePortsOk returns a tuple with the SsBridgePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsBridgePorts

`func (o *PatchedNodeRequest) SetSsBridgePorts(v []string)`

SetSsBridgePorts sets SsBridgePorts field to given value.

### HasSsBridgePorts

`func (o *PatchedNodeRequest) HasSsBridgePorts() bool`

HasSsBridgePorts returns a boolean if a field has been set.

### SetSsBridgePortsNil

`func (o *PatchedNodeRequest) SetSsBridgePortsNil(b bool)`

 SetSsBridgePortsNil sets the value for SsBridgePorts to be an explicit nil

### UnsetSsBridgePorts
`func (o *PatchedNodeRequest) UnsetSsBridgePorts()`

UnsetSsBridgePorts ensures that no value is present for SsBridgePorts, not even an explicit nil
### GetStatus

`func (o *PatchedNodeRequest) GetStatus() NodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedNodeRequest) GetStatusOk() (*NodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedNodeRequest) SetStatus(v NodeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedNodeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOs

`func (o *PatchedNodeRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PatchedNodeRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PatchedNodeRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *PatchedNodeRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *PatchedNodeRequest) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *PatchedNodeRequest) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetOsVersion

`func (o *PatchedNodeRequest) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *PatchedNodeRequest) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *PatchedNodeRequest) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *PatchedNodeRequest) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *PatchedNodeRequest) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *PatchedNodeRequest) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetOsArch

`func (o *PatchedNodeRequest) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *PatchedNodeRequest) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *PatchedNodeRequest) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *PatchedNodeRequest) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### SetOsArchNil

`func (o *PatchedNodeRequest) SetOsArchNil(b bool)`

 SetOsArchNil sets the value for OsArch to be an explicit nil

### UnsetOsArch
`func (o *PatchedNodeRequest) UnsetOsArch()`

UnsetOsArch ensures that no value is present for OsArch, not even an explicit nil
### GetDistro

`func (o *PatchedNodeRequest) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *PatchedNodeRequest) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *PatchedNodeRequest) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *PatchedNodeRequest) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *PatchedNodeRequest) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *PatchedNodeRequest) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *PatchedNodeRequest) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *PatchedNodeRequest) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *PatchedNodeRequest) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *PatchedNodeRequest) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### SetDistroVersionNil

`func (o *PatchedNodeRequest) SetDistroVersionNil(b bool)`

 SetDistroVersionNil sets the value for DistroVersion to be an explicit nil

### UnsetDistroVersion
`func (o *PatchedNodeRequest) UnsetDistroVersion()`

UnsetDistroVersion ensures that no value is present for DistroVersion, not even an explicit nil
### GetDistroCodename

`func (o *PatchedNodeRequest) GetDistroCodename() string`

GetDistroCodename returns the DistroCodename field if non-nil, zero value otherwise.

### GetDistroCodenameOk

`func (o *PatchedNodeRequest) GetDistroCodenameOk() (*string, bool)`

GetDistroCodenameOk returns a tuple with the DistroCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroCodename

`func (o *PatchedNodeRequest) SetDistroCodename(v string)`

SetDistroCodename sets DistroCodename field to given value.

### HasDistroCodename

`func (o *PatchedNodeRequest) HasDistroCodename() bool`

HasDistroCodename returns a boolean if a field has been set.

### SetDistroCodenameNil

`func (o *PatchedNodeRequest) SetDistroCodenameNil(b bool)`

 SetDistroCodenameNil sets the value for DistroCodename to be an explicit nil

### UnsetDistroCodename
`func (o *PatchedNodeRequest) UnsetDistroCodename()`

UnsetDistroCodename ensures that no value is present for DistroCodename, not even an explicit nil
### GetAppVersion

`func (o *PatchedNodeRequest) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *PatchedNodeRequest) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *PatchedNodeRequest) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *PatchedNodeRequest) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *PatchedNodeRequest) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *PatchedNodeRequest) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetAppBuild

`func (o *PatchedNodeRequest) GetAppBuild() string`

GetAppBuild returns the AppBuild field if non-nil, zero value otherwise.

### GetAppBuildOk

`func (o *PatchedNodeRequest) GetAppBuildOk() (*string, bool)`

GetAppBuildOk returns a tuple with the AppBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuild

`func (o *PatchedNodeRequest) SetAppBuild(v string)`

SetAppBuild sets AppBuild field to given value.

### HasAppBuild

`func (o *PatchedNodeRequest) HasAppBuild() bool`

HasAppBuild returns a boolean if a field has been set.

### SetAppBuildNil

`func (o *PatchedNodeRequest) SetAppBuildNil(b bool)`

 SetAppBuildNil sets the value for AppBuild to be an explicit nil

### UnsetAppBuild
`func (o *PatchedNodeRequest) UnsetAppBuild()`

UnsetAppBuild ensures that no value is present for AppBuild, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


