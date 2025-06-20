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

// checks if the LocationRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationRegion{}

// LocationRegion struct for LocationRegion
type LocationRegion struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _LocationRegion LocationRegion

// NewLocationRegion instantiates a new LocationRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRegion(id string, name string) *LocationRegion {
	this := LocationRegion{}
	this.Id = id
	this.Name = name
	return &this
}

// NewLocationRegionWithDefaults instantiates a new LocationRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRegionWithDefaults() *LocationRegion {
	this := LocationRegion{}
	return &this
}

// GetId returns the Id field value
func (o *LocationRegion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocationRegion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocationRegion) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LocationRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LocationRegion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LocationRegion) SetName(v string) {
	o.Name = v
}

func (o LocationRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LocationRegion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varLocationRegion := _LocationRegion{}

	err = json.Unmarshal(data, &varLocationRegion)

	if err != nil {
		return err
	}

	*o = LocationRegion(varLocationRegion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocationRegion struct {
	value *LocationRegion
	isSet bool
}

func (v NullableLocationRegion) Get() *LocationRegion {
	return v.value
}

func (v *NullableLocationRegion) Set(val *LocationRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRegion(val *LocationRegion) *NullableLocationRegion {
	return &NullableLocationRegion{value: val, isSet: true}
}

func (v NullableLocationRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


