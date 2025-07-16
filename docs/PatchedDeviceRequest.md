# PatchedDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalKey** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PubKey** | Pointer to **string** |  | [optional] 
**RawPassword** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **NullableString** | Bcrypt Hashed Password for Authentication, e.g. HTTP Proxy | [optional] 
**Os** | Pointer to **NullableString** |  | [optional] 
**OsVersion** | Pointer to **NullableString** |  | [optional] 
**OsArch** | Pointer to **NullableString** |  | [optional] 
**Distro** | Pointer to **NullableString** |  | [optional] 
**DistroVersion** | Pointer to **NullableString** |  | [optional] 
**DistroCodename** | Pointer to **NullableString** |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**AppBuild** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedDeviceRequest

`func NewPatchedDeviceRequest() *PatchedDeviceRequest`

NewPatchedDeviceRequest instantiates a new PatchedDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDeviceRequestWithDefaults

`func NewPatchedDeviceRequestWithDefaults() *PatchedDeviceRequest`

NewPatchedDeviceRequestWithDefaults instantiates a new PatchedDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalKey

`func (o *PatchedDeviceRequest) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *PatchedDeviceRequest) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *PatchedDeviceRequest) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *PatchedDeviceRequest) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *PatchedDeviceRequest) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *PatchedDeviceRequest) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetName

`func (o *PatchedDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPubKey

`func (o *PatchedDeviceRequest) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *PatchedDeviceRequest) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *PatchedDeviceRequest) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.

### HasPubKey

`func (o *PatchedDeviceRequest) HasPubKey() bool`

HasPubKey returns a boolean if a field has been set.

### GetRawPassword

`func (o *PatchedDeviceRequest) GetRawPassword() string`

GetRawPassword returns the RawPassword field if non-nil, zero value otherwise.

### GetRawPasswordOk

`func (o *PatchedDeviceRequest) GetRawPasswordOk() (*string, bool)`

GetRawPasswordOk returns a tuple with the RawPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPassword

`func (o *PatchedDeviceRequest) SetRawPassword(v string)`

SetRawPassword sets RawPassword field to given value.

### HasRawPassword

`func (o *PatchedDeviceRequest) HasRawPassword() bool`

HasRawPassword returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedDeviceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedDeviceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedDeviceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedDeviceRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchedDeviceRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchedDeviceRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetOs

`func (o *PatchedDeviceRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PatchedDeviceRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PatchedDeviceRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *PatchedDeviceRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *PatchedDeviceRequest) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *PatchedDeviceRequest) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetOsVersion

`func (o *PatchedDeviceRequest) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *PatchedDeviceRequest) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *PatchedDeviceRequest) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *PatchedDeviceRequest) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *PatchedDeviceRequest) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *PatchedDeviceRequest) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetOsArch

`func (o *PatchedDeviceRequest) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *PatchedDeviceRequest) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *PatchedDeviceRequest) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *PatchedDeviceRequest) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### SetOsArchNil

`func (o *PatchedDeviceRequest) SetOsArchNil(b bool)`

 SetOsArchNil sets the value for OsArch to be an explicit nil

### UnsetOsArch
`func (o *PatchedDeviceRequest) UnsetOsArch()`

UnsetOsArch ensures that no value is present for OsArch, not even an explicit nil
### GetDistro

`func (o *PatchedDeviceRequest) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *PatchedDeviceRequest) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *PatchedDeviceRequest) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *PatchedDeviceRequest) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *PatchedDeviceRequest) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *PatchedDeviceRequest) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *PatchedDeviceRequest) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *PatchedDeviceRequest) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *PatchedDeviceRequest) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *PatchedDeviceRequest) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### SetDistroVersionNil

`func (o *PatchedDeviceRequest) SetDistroVersionNil(b bool)`

 SetDistroVersionNil sets the value for DistroVersion to be an explicit nil

### UnsetDistroVersion
`func (o *PatchedDeviceRequest) UnsetDistroVersion()`

UnsetDistroVersion ensures that no value is present for DistroVersion, not even an explicit nil
### GetDistroCodename

`func (o *PatchedDeviceRequest) GetDistroCodename() string`

GetDistroCodename returns the DistroCodename field if non-nil, zero value otherwise.

### GetDistroCodenameOk

`func (o *PatchedDeviceRequest) GetDistroCodenameOk() (*string, bool)`

GetDistroCodenameOk returns a tuple with the DistroCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroCodename

`func (o *PatchedDeviceRequest) SetDistroCodename(v string)`

SetDistroCodename sets DistroCodename field to given value.

### HasDistroCodename

`func (o *PatchedDeviceRequest) HasDistroCodename() bool`

HasDistroCodename returns a boolean if a field has been set.

### SetDistroCodenameNil

`func (o *PatchedDeviceRequest) SetDistroCodenameNil(b bool)`

 SetDistroCodenameNil sets the value for DistroCodename to be an explicit nil

### UnsetDistroCodename
`func (o *PatchedDeviceRequest) UnsetDistroCodename()`

UnsetDistroCodename ensures that no value is present for DistroCodename, not even an explicit nil
### GetAppVersion

`func (o *PatchedDeviceRequest) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *PatchedDeviceRequest) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *PatchedDeviceRequest) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *PatchedDeviceRequest) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *PatchedDeviceRequest) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *PatchedDeviceRequest) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetAppBuild

`func (o *PatchedDeviceRequest) GetAppBuild() string`

GetAppBuild returns the AppBuild field if non-nil, zero value otherwise.

### GetAppBuildOk

`func (o *PatchedDeviceRequest) GetAppBuildOk() (*string, bool)`

GetAppBuildOk returns a tuple with the AppBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuild

`func (o *PatchedDeviceRequest) SetAppBuild(v string)`

SetAppBuild sets AppBuild field to given value.

### HasAppBuild

`func (o *PatchedDeviceRequest) HasAppBuild() bool`

HasAppBuild returns a boolean if a field has been set.

### SetAppBuildNil

`func (o *PatchedDeviceRequest) SetAppBuildNil(b bool)`

 SetAppBuildNil sets the value for AppBuild to be an explicit nil

### UnsetAppBuild
`func (o *PatchedDeviceRequest) UnsetAppBuild()`

UnsetAppBuild ensures that no value is present for AppBuild, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


