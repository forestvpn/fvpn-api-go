# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Hostname** | **string** |  | 
**IpAddresses** | **[]string** |  | 
**Subnets** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsExitNode** | **bool** |  | [readonly] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**PubKey** | **string** |  | 
**Ports** | **[]string** |  | 
**HttpsProxyPorts** | Pointer to **[]string** |  | [optional] 
**Conditions** | [**[]NodeCondition**](NodeCondition.md) |  | [readonly] 
**Status** | Pointer to [**NodeStatus**](NodeStatus.md) |  | [optional] 
**Os** | Pointer to **NullableString** |  | [optional] 
**OsVersion** | Pointer to **NullableString** |  | [optional] 
**OsArch** | Pointer to **NullableString** |  | [optional] 
**Distro** | Pointer to **NullableString** |  | [optional] 
**DistroVersion** | Pointer to **NullableString** |  | [optional] 
**DistroCodename** | Pointer to **NullableString** |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**AppBuild** | Pointer to **NullableString** |  | [optional] 
**LastActiveAt** | **NullableTime** |  | [readonly] 
**IpDetails** | [**[]IPInfo**](IPInfo.md) |  | [readonly] 

## Methods

### NewNode

`func NewNode(id string, hostname string, ipAddresses []string, isExitNode bool, pubKey string, ports []string, conditions []NodeCondition, lastActiveAt NullableTime, ipDetails []IPInfo, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.


### GetHostname

`func (o *Node) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Node) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Node) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIpAddresses

`func (o *Node) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Node) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Node) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetSubnets

`func (o *Node) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *Node) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *Node) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *Node) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *Node) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *Node) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetTags

`func (o *Node) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Node) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Node) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Node) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsExitNode

`func (o *Node) GetIsExitNode() bool`

GetIsExitNode returns the IsExitNode field if non-nil, zero value otherwise.

### GetIsExitNodeOk

`func (o *Node) GetIsExitNodeOk() (*bool, bool)`

GetIsExitNodeOk returns a tuple with the IsExitNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExitNode

`func (o *Node) SetIsExitNode(v bool)`

SetIsExitNode sets IsExitNode field to given value.


### GetIsPublic

`func (o *Node) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Node) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Node) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Node) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetPubKey

`func (o *Node) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *Node) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *Node) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetPorts

`func (o *Node) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Node) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Node) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetHttpsProxyPorts

`func (o *Node) GetHttpsProxyPorts() []string`

GetHttpsProxyPorts returns the HttpsProxyPorts field if non-nil, zero value otherwise.

### GetHttpsProxyPortsOk

`func (o *Node) GetHttpsProxyPortsOk() (*[]string, bool)`

GetHttpsProxyPortsOk returns a tuple with the HttpsProxyPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyPorts

`func (o *Node) SetHttpsProxyPorts(v []string)`

SetHttpsProxyPorts sets HttpsProxyPorts field to given value.

### HasHttpsProxyPorts

`func (o *Node) HasHttpsProxyPorts() bool`

HasHttpsProxyPorts returns a boolean if a field has been set.

### SetHttpsProxyPortsNil

`func (o *Node) SetHttpsProxyPortsNil(b bool)`

 SetHttpsProxyPortsNil sets the value for HttpsProxyPorts to be an explicit nil

### UnsetHttpsProxyPorts
`func (o *Node) UnsetHttpsProxyPorts()`

UnsetHttpsProxyPorts ensures that no value is present for HttpsProxyPorts, not even an explicit nil
### GetConditions

`func (o *Node) GetConditions() []NodeCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Node) GetConditionsOk() (*[]NodeCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Node) SetConditions(v []NodeCondition)`

SetConditions sets Conditions field to given value.


### GetStatus

`func (o *Node) GetStatus() NodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Node) GetStatusOk() (*NodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Node) SetStatus(v NodeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Node) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOs

`func (o *Node) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Node) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Node) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Node) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *Node) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *Node) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetOsVersion

`func (o *Node) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Node) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Node) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *Node) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *Node) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *Node) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetOsArch

`func (o *Node) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *Node) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *Node) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *Node) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### SetOsArchNil

`func (o *Node) SetOsArchNil(b bool)`

 SetOsArchNil sets the value for OsArch to be an explicit nil

### UnsetOsArch
`func (o *Node) UnsetOsArch()`

UnsetOsArch ensures that no value is present for OsArch, not even an explicit nil
### GetDistro

`func (o *Node) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *Node) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *Node) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *Node) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *Node) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *Node) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *Node) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *Node) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *Node) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *Node) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### SetDistroVersionNil

`func (o *Node) SetDistroVersionNil(b bool)`

 SetDistroVersionNil sets the value for DistroVersion to be an explicit nil

### UnsetDistroVersion
`func (o *Node) UnsetDistroVersion()`

UnsetDistroVersion ensures that no value is present for DistroVersion, not even an explicit nil
### GetDistroCodename

`func (o *Node) GetDistroCodename() string`

GetDistroCodename returns the DistroCodename field if non-nil, zero value otherwise.

### GetDistroCodenameOk

`func (o *Node) GetDistroCodenameOk() (*string, bool)`

GetDistroCodenameOk returns a tuple with the DistroCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroCodename

`func (o *Node) SetDistroCodename(v string)`

SetDistroCodename sets DistroCodename field to given value.

### HasDistroCodename

`func (o *Node) HasDistroCodename() bool`

HasDistroCodename returns a boolean if a field has been set.

### SetDistroCodenameNil

`func (o *Node) SetDistroCodenameNil(b bool)`

 SetDistroCodenameNil sets the value for DistroCodename to be an explicit nil

### UnsetDistroCodename
`func (o *Node) UnsetDistroCodename()`

UnsetDistroCodename ensures that no value is present for DistroCodename, not even an explicit nil
### GetAppVersion

`func (o *Node) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *Node) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *Node) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *Node) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *Node) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *Node) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetAppBuild

`func (o *Node) GetAppBuild() string`

GetAppBuild returns the AppBuild field if non-nil, zero value otherwise.

### GetAppBuildOk

`func (o *Node) GetAppBuildOk() (*string, bool)`

GetAppBuildOk returns a tuple with the AppBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuild

`func (o *Node) SetAppBuild(v string)`

SetAppBuild sets AppBuild field to given value.

### HasAppBuild

`func (o *Node) HasAppBuild() bool`

HasAppBuild returns a boolean if a field has been set.

### SetAppBuildNil

`func (o *Node) SetAppBuildNil(b bool)`

 SetAppBuildNil sets the value for AppBuild to be an explicit nil

### UnsetAppBuild
`func (o *Node) UnsetAppBuild()`

UnsetAppBuild ensures that no value is present for AppBuild, not even an explicit nil
### GetLastActiveAt

`func (o *Node) GetLastActiveAt() time.Time`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *Node) GetLastActiveAtOk() (*time.Time, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *Node) SetLastActiveAt(v time.Time)`

SetLastActiveAt sets LastActiveAt field to given value.


### SetLastActiveAtNil

`func (o *Node) SetLastActiveAtNil(b bool)`

 SetLastActiveAtNil sets the value for LastActiveAt to be an explicit nil

### UnsetLastActiveAt
`func (o *Node) UnsetLastActiveAt()`

UnsetLastActiveAt ensures that no value is present for LastActiveAt, not even an explicit nil
### GetIpDetails

`func (o *Node) GetIpDetails() []IPInfo`

GetIpDetails returns the IpDetails field if non-nil, zero value otherwise.

### GetIpDetailsOk

`func (o *Node) GetIpDetailsOk() (*[]IPInfo, bool)`

GetIpDetailsOk returns a tuple with the IpDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDetails

`func (o *Node) SetIpDetails(v []IPInfo)`

SetIpDetails sets IpDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


