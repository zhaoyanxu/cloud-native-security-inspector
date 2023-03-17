/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: ${project.version}
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConstraintsViolation It represents a constraints violation error
type ConstraintsViolation struct {
	// The field that is causing the constraints violation
	Field string `json:"field"`
	// The human-readable constraints violation description
	Message string `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _ConstraintsViolation ConstraintsViolation

// NewConstraintsViolation instantiates a new ConstraintsViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraintsViolation(field string, message string) *ConstraintsViolation {
	this := ConstraintsViolation{}
	this.Field = field
	this.Message = message
	return &this
}

// NewConstraintsViolationWithDefaults instantiates a new ConstraintsViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintsViolationWithDefaults() *ConstraintsViolation {
	this := ConstraintsViolation{}
	return &this
}

// GetField returns the Field field value
func (o *ConstraintsViolation) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ConstraintsViolation) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *ConstraintsViolation) SetField(v string) {
	o.Field = v
}

// GetMessage returns the Message field value
func (o *ConstraintsViolation) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ConstraintsViolation) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ConstraintsViolation) SetMessage(v string) {
	o.Message = v
}

func (o ConstraintsViolation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConstraintsViolation) UnmarshalJSON(bytes []byte) (err error) {
	varConstraintsViolation := _ConstraintsViolation{}

	if err = json.Unmarshal(bytes, &varConstraintsViolation); err == nil {
		*o = ConstraintsViolation(varConstraintsViolation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstraintsViolation struct {
	value *ConstraintsViolation
	isSet bool
}

func (v NullableConstraintsViolation) Get() *ConstraintsViolation {
	return v.value
}

func (v *NullableConstraintsViolation) Set(val *ConstraintsViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintsViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintsViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintsViolation(val *ConstraintsViolation) *NullableConstraintsViolation {
	return &NullableConstraintsViolation{value: val, isSet: true}
}

func (v NullableConstraintsViolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraintsViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


