/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AssetRelation A related asset as child or parent
type AssetRelation struct {
	// the id of the child or parent asset
	RelatedAssetId int32 `json:"relatedAssetId"`
	// type of relation
	Type *string `json:"type,omitempty"`
}

// NewAssetRelation instantiates a new AssetRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetRelation(relatedAssetId int32) *AssetRelation {
	this := AssetRelation{}
	this.RelatedAssetId = relatedAssetId
	var type_ string = "location"
	this.Type = &type_
	return &this
}

// NewAssetRelationWithDefaults instantiates a new AssetRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetRelationWithDefaults() *AssetRelation {
	this := AssetRelation{}
	var type_ string = "location"
	this.Type = &type_
	return &this
}

// GetRelatedAssetId returns the RelatedAssetId field value
func (o *AssetRelation) GetRelatedAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelatedAssetId
}

// GetRelatedAssetIdOk returns a tuple with the RelatedAssetId field value
// and a boolean to check if the value has been set.
func (o *AssetRelation) GetRelatedAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedAssetId, true
}

// SetRelatedAssetId sets field value
func (o *AssetRelation) SetRelatedAssetId(v int32) {
	o.RelatedAssetId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssetRelation) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetRelation) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssetRelation) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssetRelation) SetType(v string) {
	o.Type = &v
}

func (o AssetRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relatedAssetId"] = o.RelatedAssetId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAssetRelation struct {
	value *AssetRelation
	isSet bool
}

func (v NullableAssetRelation) Get() *AssetRelation {
	return v.value
}

func (v *NullableAssetRelation) Set(val *AssetRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetRelation(val *AssetRelation) *NullableAssetRelation {
	return &NullableAssetRelation{value: val, isSet: true}
}

func (v NullableAssetRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}