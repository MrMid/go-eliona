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

// Widget A widget on a frontend dashboard
type Widget struct {
	// The name for the type of this widget
	WidgetTypeName *string `json:"widgetTypeName,omitempty"`
	// The width of this widget on dashboard
	Width string `json:"width"`
	// The number of days if the widget type uses timespan
	Timespan *int32 `json:"timespan,omitempty"`
	// Detailed configuration depending on the widget type
	Details map[string]interface{} `json:"details,omitempty"`
	// The master asset id of this widget
	AssetId *int32 `json:"assetId,omitempty"`
	// List of data for the elements of widget
	Data []WidgetData `json:"data,omitempty"`
}

// NewWidget instantiates a new Widget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidget(width string) *Widget {
	this := Widget{}
	this.Width = width
	return &this
}

// NewWidgetWithDefaults instantiates a new Widget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetWithDefaults() *Widget {
	this := Widget{}
	return &this
}

// GetWidgetTypeName returns the WidgetTypeName field value if set, zero value otherwise.
func (o *Widget) GetWidgetTypeName() string {
	if o == nil || o.WidgetTypeName == nil {
		var ret string
		return ret
	}
	return *o.WidgetTypeName
}

// GetWidgetTypeNameOk returns a tuple with the WidgetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetWidgetTypeNameOk() (*string, bool) {
	if o == nil || o.WidgetTypeName == nil {
		return nil, false
	}
	return o.WidgetTypeName, true
}

// HasWidgetTypeName returns a boolean if a field has been set.
func (o *Widget) HasWidgetTypeName() bool {
	if o != nil && o.WidgetTypeName != nil {
		return true
	}

	return false
}

// SetWidgetTypeName gets a reference to the given string and assigns it to the WidgetTypeName field.
func (o *Widget) SetWidgetTypeName(v string) {
	o.WidgetTypeName = &v
}

// GetWidth returns the Width field value
func (o *Widget) GetWidth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Widget) GetWidthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Widget) SetWidth(v string) {
	o.Width = v
}

// GetTimespan returns the Timespan field value if set, zero value otherwise.
func (o *Widget) GetTimespan() int32 {
	if o == nil || o.Timespan == nil {
		var ret int32
		return ret
	}
	return *o.Timespan
}

// GetTimespanOk returns a tuple with the Timespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetTimespanOk() (*int32, bool) {
	if o == nil || o.Timespan == nil {
		return nil, false
	}
	return o.Timespan, true
}

// HasTimespan returns a boolean if a field has been set.
func (o *Widget) HasTimespan() bool {
	if o != nil && o.Timespan != nil {
		return true
	}

	return false
}

// SetTimespan gets a reference to the given int32 and assigns it to the Timespan field.
func (o *Widget) SetTimespan(v int32) {
	o.Timespan = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Widget) GetDetails() map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Widget) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *Widget) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *Widget) GetAssetId() int32 {
	if o == nil || o.AssetId == nil {
		var ret int32
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetAssetIdOk() (*int32, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *Widget) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given int32 and assigns it to the AssetId field.
func (o *Widget) SetAssetId(v int32) {
	o.AssetId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Widget) GetData() []WidgetData {
	if o == nil || o.Data == nil {
		var ret []WidgetData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetDataOk() ([]WidgetData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Widget) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WidgetData and assigns it to the Data field.
func (o *Widget) SetData(v []WidgetData) {
	o.Data = v
}

func (o Widget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WidgetTypeName != nil {
		toSerialize["widgetTypeName"] = o.WidgetTypeName
	}
	if true {
		toSerialize["width"] = o.Width
	}
	if o.Timespan != nil {
		toSerialize["timespan"] = o.Timespan
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.AssetId != nil {
		toSerialize["assetId"] = o.AssetId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableWidget struct {
	value *Widget
	isSet bool
}

func (v NullableWidget) Get() *Widget {
	return v.value
}

func (v *NullableWidget) Set(val *Widget) {
	v.value = val
	v.isSet = true
}

func (v NullableWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidget(val *Widget) *NullableWidget {
	return &NullableWidget{value: val, isSet: true}
}

func (v NullableWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}