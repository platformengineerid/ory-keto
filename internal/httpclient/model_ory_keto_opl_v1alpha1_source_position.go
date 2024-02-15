/*
 * Ory Keto API
 *
 * Documentation for all of Ory Keto's REST APIs. gRPC is documented separately.
 *
 * API version:
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OryKetoOplV1alpha1SourcePosition struct for OryKetoOplV1alpha1SourcePosition
type OryKetoOplV1alpha1SourcePosition struct {
	Line   *int64 `json:"Line,omitempty"`
	Column *int64 `json:"column,omitempty"`
}

// NewOryKetoOplV1alpha1SourcePosition instantiates a new OryKetoOplV1alpha1SourcePosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOryKetoOplV1alpha1SourcePosition() *OryKetoOplV1alpha1SourcePosition {
	this := OryKetoOplV1alpha1SourcePosition{}
	return &this
}

// NewOryKetoOplV1alpha1SourcePositionWithDefaults instantiates a new OryKetoOplV1alpha1SourcePosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOryKetoOplV1alpha1SourcePositionWithDefaults() *OryKetoOplV1alpha1SourcePosition {
	this := OryKetoOplV1alpha1SourcePosition{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *OryKetoOplV1alpha1SourcePosition) GetLine() int64 {
	if o == nil || o.Line == nil {
		var ret int64
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OryKetoOplV1alpha1SourcePosition) GetLineOk() (*int64, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *OryKetoOplV1alpha1SourcePosition) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int64 and assigns it to the Line field.
func (o *OryKetoOplV1alpha1SourcePosition) SetLine(v int64) {
	o.Line = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *OryKetoOplV1alpha1SourcePosition) GetColumn() int64 {
	if o == nil || o.Column == nil {
		var ret int64
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OryKetoOplV1alpha1SourcePosition) GetColumnOk() (*int64, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *OryKetoOplV1alpha1SourcePosition) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given int64 and assigns it to the Column field.
func (o *OryKetoOplV1alpha1SourcePosition) SetColumn(v int64) {
	o.Column = &v
}

func (o OryKetoOplV1alpha1SourcePosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Line != nil {
		toSerialize["Line"] = o.Line
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	return json.Marshal(toSerialize)
}

type NullableOryKetoOplV1alpha1SourcePosition struct {
	value *OryKetoOplV1alpha1SourcePosition
	isSet bool
}

func (v NullableOryKetoOplV1alpha1SourcePosition) Get() *OryKetoOplV1alpha1SourcePosition {
	return v.value
}

func (v *NullableOryKetoOplV1alpha1SourcePosition) Set(val *OryKetoOplV1alpha1SourcePosition) {
	v.value = val
	v.isSet = true
}

func (v NullableOryKetoOplV1alpha1SourcePosition) IsSet() bool {
	return v.isSet
}

func (v *NullableOryKetoOplV1alpha1SourcePosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOryKetoOplV1alpha1SourcePosition(val *OryKetoOplV1alpha1SourcePosition) *NullableOryKetoOplV1alpha1SourcePosition {
	return &NullableOryKetoOplV1alpha1SourcePosition{value: val, isSet: true}
}

func (v NullableOryKetoOplV1alpha1SourcePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOryKetoOplV1alpha1SourcePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
