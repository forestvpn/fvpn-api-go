/*
ForestVPN API

ForestVPN API Documentation

API version: 3.0.0
Contact: support@forestvpn.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fvpn

import (
	"encoding/json"
	"fmt"
)

// checks if the NodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeRequest{}

// NodeRequest struct for NodeRequest
type NodeRequest struct {
	Hostname string `json:"hostname" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	IpAddresses []string `json:"ip_addresses"`
	Subnets []string `json:"subnets,omitempty"`
	Tags []string `json:"tags"`
	IsPublic *bool `json:"is_public,omitempty"`
	PubKey string `json:"pub_key"`
	Ports []string `json:"ports"`
	// Ports used for HTTPs Proxy
	HttpsProxyPorts []string `json:"https_proxy_ports,omitempty"`
	// Ports used for Shadow Socket Bridge
	SsBridgePorts []string `json:"ss_bridge_ports,omitempty"`
	Status *NodeStatus `json:"status,omitempty"`
	Os NullableString `json:"os,omitempty"`
	OsVersion NullableString `json:"os_version,omitempty"`
	OsArch NullableString `json:"os_arch,omitempty"`
	Distro NullableString `json:"distro,omitempty"`
	DistroVersion NullableString `json:"distro_version,omitempty"`
	DistroCodename NullableString `json:"distro_codename,omitempty"`
	AppVersion NullableString `json:"app_version,omitempty"`
	AppBuild NullableString `json:"app_build,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NodeRequest NodeRequest

// NewNodeRequest instantiates a new NodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeRequest(hostname string, ipAddresses []string, tags []string, pubKey string, ports []string) *NodeRequest {
	this := NodeRequest{}
	this.Hostname = hostname
	this.IpAddresses = ipAddresses
	this.Tags = tags
	this.PubKey = pubKey
	this.Ports = ports
	return &this
}

// NewNodeRequestWithDefaults instantiates a new NodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeRequestWithDefaults() *NodeRequest {
	this := NodeRequest{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *NodeRequest) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *NodeRequest) SetHostname(v string) {
	o.Hostname = v
}

// GetIpAddresses returns the IpAddresses field value
func (o *NodeRequest) GetIpAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetIpAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// SetIpAddresses sets field value
func (o *NodeRequest) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *NodeRequest) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *NodeRequest) SetSubnets(v []string) {
	o.Subnets = v
}

// GetTags returns the Tags field value
func (o *NodeRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *NodeRequest) SetTags(v []string) {
	o.Tags = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *NodeRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *NodeRequest) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *NodeRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetPubKey returns the PubKey field value
func (o *NodeRequest) GetPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubKey, true
}

// SetPubKey sets field value
func (o *NodeRequest) SetPubKey(v string) {
	o.PubKey = v
}

// GetPorts returns the Ports field value
func (o *NodeRequest) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetPortsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *NodeRequest) SetPorts(v []string) {
	o.Ports = v
}

// GetHttpsProxyPorts returns the HttpsProxyPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetHttpsProxyPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HttpsProxyPorts
}

// GetHttpsProxyPortsOk returns a tuple with the HttpsProxyPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetHttpsProxyPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.HttpsProxyPorts) {
		return nil, false
	}
	return o.HttpsProxyPorts, true
}

// HasHttpsProxyPorts returns a boolean if a field has been set.
func (o *NodeRequest) HasHttpsProxyPorts() bool {
	if o != nil && !IsNil(o.HttpsProxyPorts) {
		return true
	}

	return false
}

// SetHttpsProxyPorts gets a reference to the given []string and assigns it to the HttpsProxyPorts field.
func (o *NodeRequest) SetHttpsProxyPorts(v []string) {
	o.HttpsProxyPorts = v
}

// GetSsBridgePorts returns the SsBridgePorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetSsBridgePorts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SsBridgePorts
}

// GetSsBridgePortsOk returns a tuple with the SsBridgePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetSsBridgePortsOk() ([]string, bool) {
	if o == nil || IsNil(o.SsBridgePorts) {
		return nil, false
	}
	return o.SsBridgePorts, true
}

// HasSsBridgePorts returns a boolean if a field has been set.
func (o *NodeRequest) HasSsBridgePorts() bool {
	if o != nil && !IsNil(o.SsBridgePorts) {
		return true
	}

	return false
}

// SetSsBridgePorts gets a reference to the given []string and assigns it to the SsBridgePorts field.
func (o *NodeRequest) SetSsBridgePorts(v []string) {
	o.SsBridgePorts = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NodeRequest) GetStatus() NodeStatus {
	if o == nil || IsNil(o.Status) {
		var ret NodeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeRequest) GetStatusOk() (*NodeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NodeRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NodeStatus and assigns it to the Status field.
func (o *NodeRequest) SetStatus(v NodeStatus) {
	o.Status = &v
}

// GetOs returns the Os field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetOs() string {
	if o == nil || IsNil(o.Os.Get()) {
		var ret string
		return ret
	}
	return *o.Os.Get()
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Os.Get(), o.Os.IsSet()
}

// HasOs returns a boolean if a field has been set.
func (o *NodeRequest) HasOs() bool {
	if o != nil && o.Os.IsSet() {
		return true
	}

	return false
}

// SetOs gets a reference to the given NullableString and assigns it to the Os field.
func (o *NodeRequest) SetOs(v string) {
	o.Os.Set(&v)
}
// SetOsNil sets the value for Os to be an explicit nil
func (o *NodeRequest) SetOsNil() {
	o.Os.Set(nil)
}

// UnsetOs ensures that no value is present for Os, not even an explicit nil
func (o *NodeRequest) UnsetOs() {
	o.Os.Unset()
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion.Get()) {
		var ret string
		return ret
	}
	return *o.OsVersion.Get()
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetOsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsVersion.Get(), o.OsVersion.IsSet()
}

// HasOsVersion returns a boolean if a field has been set.
func (o *NodeRequest) HasOsVersion() bool {
	if o != nil && o.OsVersion.IsSet() {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given NullableString and assigns it to the OsVersion field.
func (o *NodeRequest) SetOsVersion(v string) {
	o.OsVersion.Set(&v)
}
// SetOsVersionNil sets the value for OsVersion to be an explicit nil
func (o *NodeRequest) SetOsVersionNil() {
	o.OsVersion.Set(nil)
}

// UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
func (o *NodeRequest) UnsetOsVersion() {
	o.OsVersion.Unset()
}

// GetOsArch returns the OsArch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetOsArch() string {
	if o == nil || IsNil(o.OsArch.Get()) {
		var ret string
		return ret
	}
	return *o.OsArch.Get()
}

// GetOsArchOk returns a tuple with the OsArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetOsArchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsArch.Get(), o.OsArch.IsSet()
}

// HasOsArch returns a boolean if a field has been set.
func (o *NodeRequest) HasOsArch() bool {
	if o != nil && o.OsArch.IsSet() {
		return true
	}

	return false
}

// SetOsArch gets a reference to the given NullableString and assigns it to the OsArch field.
func (o *NodeRequest) SetOsArch(v string) {
	o.OsArch.Set(&v)
}
// SetOsArchNil sets the value for OsArch to be an explicit nil
func (o *NodeRequest) SetOsArchNil() {
	o.OsArch.Set(nil)
}

// UnsetOsArch ensures that no value is present for OsArch, not even an explicit nil
func (o *NodeRequest) UnsetOsArch() {
	o.OsArch.Unset()
}

// GetDistro returns the Distro field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetDistro() string {
	if o == nil || IsNil(o.Distro.Get()) {
		var ret string
		return ret
	}
	return *o.Distro.Get()
}

// GetDistroOk returns a tuple with the Distro field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetDistroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Distro.Get(), o.Distro.IsSet()
}

// HasDistro returns a boolean if a field has been set.
func (o *NodeRequest) HasDistro() bool {
	if o != nil && o.Distro.IsSet() {
		return true
	}

	return false
}

// SetDistro gets a reference to the given NullableString and assigns it to the Distro field.
func (o *NodeRequest) SetDistro(v string) {
	o.Distro.Set(&v)
}
// SetDistroNil sets the value for Distro to be an explicit nil
func (o *NodeRequest) SetDistroNil() {
	o.Distro.Set(nil)
}

// UnsetDistro ensures that no value is present for Distro, not even an explicit nil
func (o *NodeRequest) UnsetDistro() {
	o.Distro.Unset()
}

// GetDistroVersion returns the DistroVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetDistroVersion() string {
	if o == nil || IsNil(o.DistroVersion.Get()) {
		var ret string
		return ret
	}
	return *o.DistroVersion.Get()
}

// GetDistroVersionOk returns a tuple with the DistroVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetDistroVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistroVersion.Get(), o.DistroVersion.IsSet()
}

// HasDistroVersion returns a boolean if a field has been set.
func (o *NodeRequest) HasDistroVersion() bool {
	if o != nil && o.DistroVersion.IsSet() {
		return true
	}

	return false
}

// SetDistroVersion gets a reference to the given NullableString and assigns it to the DistroVersion field.
func (o *NodeRequest) SetDistroVersion(v string) {
	o.DistroVersion.Set(&v)
}
// SetDistroVersionNil sets the value for DistroVersion to be an explicit nil
func (o *NodeRequest) SetDistroVersionNil() {
	o.DistroVersion.Set(nil)
}

// UnsetDistroVersion ensures that no value is present for DistroVersion, not even an explicit nil
func (o *NodeRequest) UnsetDistroVersion() {
	o.DistroVersion.Unset()
}

// GetDistroCodename returns the DistroCodename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetDistroCodename() string {
	if o == nil || IsNil(o.DistroCodename.Get()) {
		var ret string
		return ret
	}
	return *o.DistroCodename.Get()
}

// GetDistroCodenameOk returns a tuple with the DistroCodename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetDistroCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistroCodename.Get(), o.DistroCodename.IsSet()
}

// HasDistroCodename returns a boolean if a field has been set.
func (o *NodeRequest) HasDistroCodename() bool {
	if o != nil && o.DistroCodename.IsSet() {
		return true
	}

	return false
}

// SetDistroCodename gets a reference to the given NullableString and assigns it to the DistroCodename field.
func (o *NodeRequest) SetDistroCodename(v string) {
	o.DistroCodename.Set(&v)
}
// SetDistroCodenameNil sets the value for DistroCodename to be an explicit nil
func (o *NodeRequest) SetDistroCodenameNil() {
	o.DistroCodename.Set(nil)
}

// UnsetDistroCodename ensures that no value is present for DistroCodename, not even an explicit nil
func (o *NodeRequest) UnsetDistroCodename() {
	o.DistroCodename.Unset()
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// HasAppVersion returns a boolean if a field has been set.
func (o *NodeRequest) HasAppVersion() bool {
	if o != nil && o.AppVersion.IsSet() {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given NullableString and assigns it to the AppVersion field.
func (o *NodeRequest) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}
// SetAppVersionNil sets the value for AppVersion to be an explicit nil
func (o *NodeRequest) SetAppVersionNil() {
	o.AppVersion.Set(nil)
}

// UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
func (o *NodeRequest) UnsetAppVersion() {
	o.AppVersion.Unset()
}

// GetAppBuild returns the AppBuild field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeRequest) GetAppBuild() string {
	if o == nil || IsNil(o.AppBuild.Get()) {
		var ret string
		return ret
	}
	return *o.AppBuild.Get()
}

// GetAppBuildOk returns a tuple with the AppBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRequest) GetAppBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppBuild.Get(), o.AppBuild.IsSet()
}

// HasAppBuild returns a boolean if a field has been set.
func (o *NodeRequest) HasAppBuild() bool {
	if o != nil && o.AppBuild.IsSet() {
		return true
	}

	return false
}

// SetAppBuild gets a reference to the given NullableString and assigns it to the AppBuild field.
func (o *NodeRequest) SetAppBuild(v string) {
	o.AppBuild.Set(&v)
}
// SetAppBuildNil sets the value for AppBuild to be an explicit nil
func (o *NodeRequest) SetAppBuildNil() {
	o.AppBuild.Set(nil)
}

// UnsetAppBuild ensures that no value is present for AppBuild, not even an explicit nil
func (o *NodeRequest) UnsetAppBuild() {
	o.AppBuild.Unset()
}

func (o NodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostname"] = o.Hostname
	toSerialize["ip_addresses"] = o.IpAddresses
	if o.Subnets != nil {
		toSerialize["subnets"] = o.Subnets
	}
	toSerialize["tags"] = o.Tags
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	toSerialize["pub_key"] = o.PubKey
	toSerialize["ports"] = o.Ports
	if o.HttpsProxyPorts != nil {
		toSerialize["https_proxy_ports"] = o.HttpsProxyPorts
	}
	if o.SsBridgePorts != nil {
		toSerialize["ss_bridge_ports"] = o.SsBridgePorts
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Os.IsSet() {
		toSerialize["os"] = o.Os.Get()
	}
	if o.OsVersion.IsSet() {
		toSerialize["os_version"] = o.OsVersion.Get()
	}
	if o.OsArch.IsSet() {
		toSerialize["os_arch"] = o.OsArch.Get()
	}
	if o.Distro.IsSet() {
		toSerialize["distro"] = o.Distro.Get()
	}
	if o.DistroVersion.IsSet() {
		toSerialize["distro_version"] = o.DistroVersion.Get()
	}
	if o.DistroCodename.IsSet() {
		toSerialize["distro_codename"] = o.DistroCodename.Get()
	}
	if o.AppVersion.IsSet() {
		toSerialize["app_version"] = o.AppVersion.Get()
	}
	if o.AppBuild.IsSet() {
		toSerialize["app_build"] = o.AppBuild.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NodeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hostname",
		"ip_addresses",
		"tags",
		"pub_key",
		"ports",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNodeRequest := _NodeRequest{}

	err = json.Unmarshal(data, &varNodeRequest)

	if err != nil {
		return err
	}

	*o = NodeRequest(varNodeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ip_addresses")
		delete(additionalProperties, "subnets")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "pub_key")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "https_proxy_ports")
		delete(additionalProperties, "ss_bridge_ports")
		delete(additionalProperties, "status")
		delete(additionalProperties, "os")
		delete(additionalProperties, "os_version")
		delete(additionalProperties, "os_arch")
		delete(additionalProperties, "distro")
		delete(additionalProperties, "distro_version")
		delete(additionalProperties, "distro_codename")
		delete(additionalProperties, "app_version")
		delete(additionalProperties, "app_build")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNodeRequest struct {
	value *NodeRequest
	isSet bool
}

func (v NullableNodeRequest) Get() *NodeRequest {
	return v.value
}

func (v *NullableNodeRequest) Set(val *NodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeRequest(val *NodeRequest) *NullableNodeRequest {
	return &NullableNodeRequest{value: val, isSet: true}
}

func (v NullableNodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


