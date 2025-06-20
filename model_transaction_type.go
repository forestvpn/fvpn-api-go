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

// TransactionType * `credit` - Credit * `debit` - Debit
type TransactionType string

// List of TransactionType
const (
	TRANSACTIONTYPE_CREDIT TransactionType = "credit"
	TRANSACTIONTYPE_DEBIT TransactionType = "debit"
)

// All allowed values of TransactionType enum
var AllowedTransactionTypeEnumValues = []TransactionType{
	"credit",
	"debit",
}

func (v *TransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionType(value)
	for _, existing := range AllowedTransactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionType", value)
}

// NewTransactionTypeFromValue returns a pointer to a valid TransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionTypeFromValue(v string) (*TransactionType, error) {
	ev := TransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionType: valid values are %v", v, AllowedTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionType) IsValid() bool {
	for _, existing := range AllowedTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionType value
func (v TransactionType) Ptr() *TransactionType {
	return &v
}

type NullableTransactionType struct {
	value *TransactionType
	isSet bool
}

func (v NullableTransactionType) Get() *TransactionType {
	return v.value
}

func (v *NullableTransactionType) Set(val *TransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionType(val *TransactionType) *NullableTransactionType {
	return &NullableTransactionType{value: val, isSet: true}
}

func (v NullableTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

