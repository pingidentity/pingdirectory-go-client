/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddPrometheusMonitorAttributeMetricRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPrometheusMonitorAttributeMetricRequest{}

// AddPrometheusMonitorAttributeMetricRequest struct for AddPrometheusMonitorAttributeMetricRequest
type AddPrometheusMonitorAttributeMetricRequest struct {
	Schemas []EnumprometheusMonitorAttributeMetricSchemaUrn `json:"schemas,omitempty"`
	// Name of the new Prometheus Monitor Attribute Metric
	MetricName string `json:"metricName"`
	// The name of the monitor attribute that contains the numeric value to be published.
	MonitorAttributeName string `json:"monitorAttributeName"`
	// The name of the object class for monitor entries that contain the monitor attribute.
	MonitorObjectClassName string                                             `json:"monitorObjectClassName"`
	MetricType             EnumprometheusMonitorAttributeMetricMetricTypeProp `json:"metricType"`
	// A filter that may be used to restrict the set of monitor entries for which the metric should be generated.
	Filter *string `json:"filter,omitempty"`
	// A human-readable description that should be published as part of the metric definition.
	MetricDescription *string `json:"metricDescription,omitempty"`
	// A set of name-value pairs for labels that should be included in the published metric for the target attribute.
	LabelNameValuePair []string `json:"labelNameValuePair,omitempty"`
}

// NewAddPrometheusMonitorAttributeMetricRequest instantiates a new AddPrometheusMonitorAttributeMetricRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPrometheusMonitorAttributeMetricRequest(metricName string, monitorAttributeName string, monitorObjectClassName string, metricType EnumprometheusMonitorAttributeMetricMetricTypeProp) *AddPrometheusMonitorAttributeMetricRequest {
	this := AddPrometheusMonitorAttributeMetricRequest{}
	this.MetricName = metricName
	this.MonitorAttributeName = monitorAttributeName
	this.MonitorObjectClassName = monitorObjectClassName
	this.MetricType = metricType
	return &this
}

// NewAddPrometheusMonitorAttributeMetricRequestWithDefaults instantiates a new AddPrometheusMonitorAttributeMetricRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPrometheusMonitorAttributeMetricRequestWithDefaults() *AddPrometheusMonitorAttributeMetricRequest {
	this := AddPrometheusMonitorAttributeMetricRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetSchemas() []EnumprometheusMonitorAttributeMetricSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumprometheusMonitorAttributeMetricSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetSchemasOk() ([]EnumprometheusMonitorAttributeMetricSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumprometheusMonitorAttributeMetricSchemaUrn and assigns it to the Schemas field.
func (o *AddPrometheusMonitorAttributeMetricRequest) SetSchemas(v []EnumprometheusMonitorAttributeMetricSchemaUrn) {
	o.Schemas = v
}

// GetMetricName returns the MetricName field value
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *AddPrometheusMonitorAttributeMetricRequest) SetMetricName(v string) {
	o.MetricName = v
}

// GetMonitorAttributeName returns the MonitorAttributeName field value
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorAttributeName
}

// GetMonitorAttributeNameOk returns a tuple with the MonitorAttributeName field value
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorAttributeName, true
}

// SetMonitorAttributeName sets field value
func (o *AddPrometheusMonitorAttributeMetricRequest) SetMonitorAttributeName(v string) {
	o.MonitorAttributeName = v
}

// GetMonitorObjectClassName returns the MonitorObjectClassName field value
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorObjectClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorObjectClassName
}

// GetMonitorObjectClassNameOk returns a tuple with the MonitorObjectClassName field value
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorObjectClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorObjectClassName, true
}

// SetMonitorObjectClassName sets field value
func (o *AddPrometheusMonitorAttributeMetricRequest) SetMonitorObjectClassName(v string) {
	o.MonitorObjectClassName = v
}

// GetMetricType returns the MetricType field value
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricType() EnumprometheusMonitorAttributeMetricMetricTypeProp {
	if o == nil {
		var ret EnumprometheusMonitorAttributeMetricMetricTypeProp
		return ret
	}

	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricTypeOk() (*EnumprometheusMonitorAttributeMetricMetricTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value
func (o *AddPrometheusMonitorAttributeMetricRequest) SetMetricType(v EnumprometheusMonitorAttributeMetricMetricTypeProp) {
	o.MetricType = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *AddPrometheusMonitorAttributeMetricRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetMetricDescription returns the MetricDescription field value if set, zero value otherwise.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricDescription() string {
	if o == nil || IsNil(o.MetricDescription) {
		var ret string
		return ret
	}
	return *o.MetricDescription
}

// GetMetricDescriptionOk returns a tuple with the MetricDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetricDescription) {
		return nil, false
	}
	return o.MetricDescription, true
}

// HasMetricDescription returns a boolean if a field has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) HasMetricDescription() bool {
	if o != nil && !IsNil(o.MetricDescription) {
		return true
	}

	return false
}

// SetMetricDescription gets a reference to the given string and assigns it to the MetricDescription field.
func (o *AddPrometheusMonitorAttributeMetricRequest) SetMetricDescription(v string) {
	o.MetricDescription = &v
}

// GetLabelNameValuePair returns the LabelNameValuePair field value if set, zero value otherwise.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetLabelNameValuePair() []string {
	if o == nil || IsNil(o.LabelNameValuePair) {
		var ret []string
		return ret
	}
	return o.LabelNameValuePair
}

// GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) GetLabelNameValuePairOk() ([]string, bool) {
	if o == nil || IsNil(o.LabelNameValuePair) {
		return nil, false
	}
	return o.LabelNameValuePair, true
}

// HasLabelNameValuePair returns a boolean if a field has been set.
func (o *AddPrometheusMonitorAttributeMetricRequest) HasLabelNameValuePair() bool {
	if o != nil && !IsNil(o.LabelNameValuePair) {
		return true
	}

	return false
}

// SetLabelNameValuePair gets a reference to the given []string and assigns it to the LabelNameValuePair field.
func (o *AddPrometheusMonitorAttributeMetricRequest) SetLabelNameValuePair(v []string) {
	o.LabelNameValuePair = v
}

func (o AddPrometheusMonitorAttributeMetricRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPrometheusMonitorAttributeMetricRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["metricName"] = o.MetricName
	toSerialize["monitorAttributeName"] = o.MonitorAttributeName
	toSerialize["monitorObjectClassName"] = o.MonitorObjectClassName
	toSerialize["metricType"] = o.MetricType
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.MetricDescription) {
		toSerialize["metricDescription"] = o.MetricDescription
	}
	if !IsNil(o.LabelNameValuePair) {
		toSerialize["labelNameValuePair"] = o.LabelNameValuePair
	}
	return toSerialize, nil
}

type NullableAddPrometheusMonitorAttributeMetricRequest struct {
	value *AddPrometheusMonitorAttributeMetricRequest
	isSet bool
}

func (v NullableAddPrometheusMonitorAttributeMetricRequest) Get() *AddPrometheusMonitorAttributeMetricRequest {
	return v.value
}

func (v *NullableAddPrometheusMonitorAttributeMetricRequest) Set(val *AddPrometheusMonitorAttributeMetricRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPrometheusMonitorAttributeMetricRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPrometheusMonitorAttributeMetricRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPrometheusMonitorAttributeMetricRequest(val *AddPrometheusMonitorAttributeMetricRequest) *NullableAddPrometheusMonitorAttributeMetricRequest {
	return &NullableAddPrometheusMonitorAttributeMetricRequest{value: val, isSet: true}
}

func (v NullableAddPrometheusMonitorAttributeMetricRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPrometheusMonitorAttributeMetricRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
