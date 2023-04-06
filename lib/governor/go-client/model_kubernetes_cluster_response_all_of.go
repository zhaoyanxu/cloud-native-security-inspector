/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: 0.1.0
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// KubernetesClusterResponseAllOf Kubernetes cluster response
type KubernetesClusterResponseAllOf struct {
	// Creation date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Last updated date
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// User who created the cluster
	CreatedBy *string `json:"created_by,omitempty"`
	// User who last updated the cluster
	UpdatedBy            *string `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterResponseAllOf KubernetesClusterResponseAllOf

// NewKubernetesClusterResponseAllOf instantiates a new KubernetesClusterResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterResponseAllOf() *KubernetesClusterResponseAllOf {
	this := KubernetesClusterResponseAllOf{}
	return &this
}

// NewKubernetesClusterResponseAllOfWithDefaults instantiates a new KubernetesClusterResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterResponseAllOfWithDefaults() *KubernetesClusterResponseAllOf {
	this := KubernetesClusterResponseAllOf{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KubernetesClusterResponseAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterResponseAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KubernetesClusterResponseAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *KubernetesClusterResponseAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KubernetesClusterResponseAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterResponseAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KubernetesClusterResponseAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *KubernetesClusterResponseAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *KubernetesClusterResponseAllOf) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterResponseAllOf) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *KubernetesClusterResponseAllOf) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *KubernetesClusterResponseAllOf) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *KubernetesClusterResponseAllOf) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterResponseAllOf) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *KubernetesClusterResponseAllOf) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *KubernetesClusterResponseAllOf) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o KubernetesClusterResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterResponseAllOf := _KubernetesClusterResponseAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterResponseAllOf); err == nil {
		*o = KubernetesClusterResponseAllOf(varKubernetesClusterResponseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterResponseAllOf struct {
	value *KubernetesClusterResponseAllOf
	isSet bool
}

func (v NullableKubernetesClusterResponseAllOf) Get() *KubernetesClusterResponseAllOf {
	return v.value
}

func (v *NullableKubernetesClusterResponseAllOf) Set(val *KubernetesClusterResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterResponseAllOf(val *KubernetesClusterResponseAllOf) *NullableKubernetesClusterResponseAllOf {
	return &NullableKubernetesClusterResponseAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
