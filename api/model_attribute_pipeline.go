/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AttributePipeline struct for AttributePipeline
type AttributePipeline struct {
	// Pipeline calculation mode
	Mode *string `json:"mode,omitempty"`
	// Pipeline calculation intervals
	Raster []string `json:"raster,omitempty"`
}

// NewAttributePipeline instantiates a new AttributePipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributePipeline() *AttributePipeline {
	this := AttributePipeline{}
	return &this
}

// NewAttributePipelineWithDefaults instantiates a new AttributePipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributePipelineWithDefaults() *AttributePipeline {
	this := AttributePipeline{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AttributePipeline) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePipeline) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AttributePipeline) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *AttributePipeline) SetMode(v string) {
	o.Mode = &v
}

// GetRaster returns the Raster field value if set, zero value otherwise.
func (o *AttributePipeline) GetRaster() []string {
	if o == nil || o.Raster == nil {
		var ret []string
		return ret
	}
	return o.Raster
}

// GetRasterOk returns a tuple with the Raster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePipeline) GetRasterOk() ([]string, bool) {
	if o == nil || o.Raster == nil {
		return nil, false
	}
	return o.Raster, true
}

// HasRaster returns a boolean if a field has been set.
func (o *AttributePipeline) HasRaster() bool {
	if o != nil && o.Raster != nil {
		return true
	}

	return false
}

// SetRaster gets a reference to the given []string and assigns it to the Raster field.
func (o *AttributePipeline) SetRaster(v []string) {
	o.Raster = v
}

func (o AttributePipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Raster != nil {
		toSerialize["raster"] = o.Raster
	}
	return json.Marshal(toSerialize)
}

type NullableAttributePipeline struct {
	value *AttributePipeline
	isSet bool
}

func (v NullableAttributePipeline) Get() *AttributePipeline {
	return v.value
}

func (v *NullableAttributePipeline) Set(val *AttributePipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributePipeline(val *AttributePipeline) *NullableAttributePipeline {
	return &NullableAttributePipeline{value: val, isSet: true}
}

func (v NullableAttributePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}