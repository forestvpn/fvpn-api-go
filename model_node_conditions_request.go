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

// checks if the NodeConditionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeConditionsRequest{}

// NodeConditionsRequest struct for NodeConditionsRequest
type NodeConditionsRequest struct {
	Items []NodeConditionRequest `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _NodeConditionsRequest NodeConditionsRequest

// NewNodeConditionsRequest instantiates a new NodeConditionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeConditionsRequest(items []NodeConditionRequest) *NodeConditionsRequest {
	this := NodeConditionsRequest{}
	this.Items = items
	return &this
}

// NewNodeConditionsRequestWithDefaults instantiates a new NodeConditionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeConditionsRequestWithDefaults() *NodeConditionsRequest {
	this := NodeConditionsRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *NodeConditionsRequest) GetItems() []NodeConditionRequest {
	if o == nil {
		var ret []NodeConditionRequest
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NodeConditionsRequest) GetItemsOk() ([]NodeConditionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NodeConditionsRequest) SetItems(v []NodeConditionRequest) {
	o.Items = v
}

func (o NodeConditionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeConditionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NodeConditionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varNodeConditionsRequest := _NodeConditionsRequest{}

	err = json.Unmarshal(data, &varNodeConditionsRequest)

	if err != nil {
		return err
	}

	*o = NodeConditionsRequest(varNodeConditionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNodeConditionsRequest struct {
	value *NodeConditionsRequest
	isSet bool
}

func (v NullableNodeConditionsRequest) Get() *NodeConditionsRequest {
	return v.value
}

func (v *NullableNodeConditionsRequest) Set(val *NodeConditionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeConditionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeConditionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeConditionsRequest(val *NodeConditionsRequest) *NullableNodeConditionsRequest {
	return &NullableNodeConditionsRequest{value: val, isSet: true}
}

func (v NullableNodeConditionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeConditionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


