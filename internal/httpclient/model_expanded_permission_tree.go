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

// ExpandedPermissionTree struct for ExpandedPermissionTree
type ExpandedPermissionTree struct {
	// The children of this node.  This is never set if `node_type` == `NODE_TYPE_LEAF`.
	Children []ExpandedPermissionTree              `json:"children,omitempty"`
	Subject  *OryKetoRelationTuplesV1alpha2Subject `json:"subject,omitempty"`
	Tuple    *Relationship                         `json:"tuple,omitempty"`
	Type     OryKetoRelationTuplesV1alpha2NodeType `json:"type"`
}

// NewExpandedPermissionTree instantiates a new ExpandedPermissionTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandedPermissionTree(type_ OryKetoRelationTuplesV1alpha2NodeType) *ExpandedPermissionTree {
	this := ExpandedPermissionTree{}
	this.Type = type_
	return &this
}

// NewExpandedPermissionTreeWithDefaults instantiates a new ExpandedPermissionTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandedPermissionTreeWithDefaults() *ExpandedPermissionTree {
	this := ExpandedPermissionTree{}
	var type_ OryKetoRelationTuplesV1alpha2NodeType = ORYKETORELATIONTUPLESV1ALPHA2NODETYPE_UNSPECIFIED
	this.Type = type_
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ExpandedPermissionTree) GetChildren() []ExpandedPermissionTree {
	if o == nil || o.Children == nil {
		var ret []ExpandedPermissionTree
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedPermissionTree) GetChildrenOk() ([]ExpandedPermissionTree, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ExpandedPermissionTree) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ExpandedPermissionTree and assigns it to the Children field.
func (o *ExpandedPermissionTree) SetChildren(v []ExpandedPermissionTree) {
	o.Children = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ExpandedPermissionTree) GetSubject() OryKetoRelationTuplesV1alpha2Subject {
	if o == nil || o.Subject == nil {
		var ret OryKetoRelationTuplesV1alpha2Subject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedPermissionTree) GetSubjectOk() (*OryKetoRelationTuplesV1alpha2Subject, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ExpandedPermissionTree) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given OryKetoRelationTuplesV1alpha2Subject and assigns it to the Subject field.
func (o *ExpandedPermissionTree) SetSubject(v OryKetoRelationTuplesV1alpha2Subject) {
	o.Subject = &v
}

// GetTuple returns the Tuple field value if set, zero value otherwise.
func (o *ExpandedPermissionTree) GetTuple() Relationship {
	if o == nil || o.Tuple == nil {
		var ret Relationship
		return ret
	}
	return *o.Tuple
}

// GetTupleOk returns a tuple with the Tuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedPermissionTree) GetTupleOk() (*Relationship, bool) {
	if o == nil || o.Tuple == nil {
		return nil, false
	}
	return o.Tuple, true
}

// HasTuple returns a boolean if a field has been set.
func (o *ExpandedPermissionTree) HasTuple() bool {
	if o != nil && o.Tuple != nil {
		return true
	}

	return false
}

// SetTuple gets a reference to the given Relationship and assigns it to the Tuple field.
func (o *ExpandedPermissionTree) SetTuple(v Relationship) {
	o.Tuple = &v
}

// GetType returns the Type field value
func (o *ExpandedPermissionTree) GetType() OryKetoRelationTuplesV1alpha2NodeType {
	if o == nil {
		var ret OryKetoRelationTuplesV1alpha2NodeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExpandedPermissionTree) GetTypeOk() (*OryKetoRelationTuplesV1alpha2NodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExpandedPermissionTree) SetType(v OryKetoRelationTuplesV1alpha2NodeType) {
	o.Type = v
}

func (o ExpandedPermissionTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Tuple != nil {
		toSerialize["tuple"] = o.Tuple
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableExpandedPermissionTree struct {
	value *ExpandedPermissionTree
	isSet bool
}

func (v NullableExpandedPermissionTree) Get() *ExpandedPermissionTree {
	return v.value
}

func (v *NullableExpandedPermissionTree) Set(val *ExpandedPermissionTree) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandedPermissionTree) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandedPermissionTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandedPermissionTree(val *ExpandedPermissionTree) *NullableExpandedPermissionTree {
	return &NullableExpandedPermissionTree{value: val, isSet: true}
}

func (v NullableExpandedPermissionTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandedPermissionTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
