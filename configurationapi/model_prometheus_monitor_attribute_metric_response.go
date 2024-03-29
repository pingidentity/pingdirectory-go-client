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

// checks if the PrometheusMonitorAttributeMetricResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusMonitorAttributeMetricResponse{}

// PrometheusMonitorAttributeMetricResponse struct for PrometheusMonitorAttributeMetricResponse
type PrometheusMonitorAttributeMetricResponse struct {
	Schemas []EnumprometheusMonitorAttributeMetricSchemaUrn `json:"schemas,omitempty"`
	// The name that will be used in the metric to be consumed by Prometheus.
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
	LabelNameValuePair                            []string                                           `json:"labelNameValuePair,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Prometheus Monitor Attribute Metric
	Id string `json:"id"`
}

// NewPrometheusMonitorAttributeMetricResponse instantiates a new PrometheusMonitorAttributeMetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusMonitorAttributeMetricResponse(metricName string, monitorAttributeName string, monitorObjectClassName string, metricType EnumprometheusMonitorAttributeMetricMetricTypeProp, id string) *PrometheusMonitorAttributeMetricResponse {
	this := PrometheusMonitorAttributeMetricResponse{}
	this.MetricName = metricName
	this.MonitorAttributeName = monitorAttributeName
	this.MonitorObjectClassName = monitorObjectClassName
	this.MetricType = metricType
	this.Id = id
	return &this
}

// NewPrometheusMonitorAttributeMetricResponseWithDefaults instantiates a new PrometheusMonitorAttributeMetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusMonitorAttributeMetricResponseWithDefaults() *PrometheusMonitorAttributeMetricResponse {
	this := PrometheusMonitorAttributeMetricResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *PrometheusMonitorAttributeMetricResponse) GetSchemas() []EnumprometheusMonitorAttributeMetricSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumprometheusMonitorAttributeMetricSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetSchemasOk() ([]EnumprometheusMonitorAttributeMetricSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *PrometheusMonitorAttributeMetricResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumprometheusMonitorAttributeMetricSchemaUrn and assigns it to the Schemas field.
func (o *PrometheusMonitorAttributeMetricResponse) SetSchemas(v []EnumprometheusMonitorAttributeMetricSchemaUrn) {
	o.Schemas = v
}

// GetMetricName returns the MetricName field value
func (o *PrometheusMonitorAttributeMetricResponse) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *PrometheusMonitorAttributeMetricResponse) SetMetricName(v string) {
	o.MetricName = v
}

// GetMonitorAttributeName returns the MonitorAttributeName field value
func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorAttributeName
}

// GetMonitorAttributeNameOk returns a tuple with the MonitorAttributeName field value
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorAttributeName, true
}

// SetMonitorAttributeName sets field value
func (o *PrometheusMonitorAttributeMetricResponse) SetMonitorAttributeName(v string) {
	o.MonitorAttributeName = v
}

// GetMonitorObjectClassName returns the MonitorObjectClassName field value
func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorObjectClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorObjectClassName
}

// GetMonitorObjectClassNameOk returns a tuple with the MonitorObjectClassName field value
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorObjectClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorObjectClassName, true
}

// SetMonitorObjectClassName sets field value
func (o *PrometheusMonitorAttributeMetricResponse) SetMonitorObjectClassName(v string) {
	o.MonitorObjectClassName = v
}

// GetMetricType returns the MetricType field value
func (o *PrometheusMonitorAttributeMetricResponse) GetMetricType() EnumprometheusMonitorAttributeMetricMetricTypeProp {
	if o == nil {
		var ret EnumprometheusMonitorAttributeMetricMetricTypeProp
		return ret
	}

	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetMetricTypeOk() (*EnumprometheusMonitorAttributeMetricMetricTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value
func (o *PrometheusMonitorAttributeMetricResponse) SetMetricType(v EnumprometheusMonitorAttributeMetricMetricTypeProp) {
	o.MetricType = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PrometheusMonitorAttributeMetricResponse) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PrometheusMonitorAttributeMetricResponse) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *PrometheusMonitorAttributeMetricResponse) SetFilter(v string) {
	o.Filter = &v
}

// GetMetricDescription returns the MetricDescription field value if set, zero value otherwise.
func (o *PrometheusMonitorAttributeMetricResponse) GetMetricDescription() string {
	if o == nil || IsNil(o.MetricDescription) {
		var ret string
		return ret
	}
	return *o.MetricDescription
}

// GetMetricDescriptionOk returns a tuple with the MetricDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetMetricDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetricDescription) {
		return nil, false
	}
	return o.MetricDescription, true
}

// HasMetricDescription returns a boolean if a field has been set.
func (o *PrometheusMonitorAttributeMetricResponse) HasMetricDescription() bool {
	if o != nil && !IsNil(o.MetricDescription) {
		return true
	}

	return false
}

// SetMetricDescription gets a reference to the given string and assigns it to the MetricDescription field.
func (o *PrometheusMonitorAttributeMetricResponse) SetMetricDescription(v string) {
	o.MetricDescription = &v
}

// GetLabelNameValuePair returns the LabelNameValuePair field value if set, zero value otherwise.
func (o *PrometheusMonitorAttributeMetricResponse) GetLabelNameValuePair() []string {
	if o == nil || IsNil(o.LabelNameValuePair) {
		var ret []string
		return ret
	}
	return o.LabelNameValuePair
}

// GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetLabelNameValuePairOk() ([]string, bool) {
	if o == nil || IsNil(o.LabelNameValuePair) {
		return nil, false
	}
	return o.LabelNameValuePair, true
}

// HasLabelNameValuePair returns a boolean if a field has been set.
func (o *PrometheusMonitorAttributeMetricResponse) HasLabelNameValuePair() bool {
	if o != nil && !IsNil(o.LabelNameValuePair) {
		return true
	}

	return false
}

// SetLabelNameValuePair gets a reference to the given []string and assigns it to the LabelNameValuePair field.
func (o *PrometheusMonitorAttributeMetricResponse) SetLabelNameValuePair(v []string) {
	o.LabelNameValuePair = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PrometheusMonitorAttributeMetricResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PrometheusMonitorAttributeMetricResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PrometheusMonitorAttributeMetricResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PrometheusMonitorAttributeMetricResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PrometheusMonitorAttributeMetricResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PrometheusMonitorAttributeMetricResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *PrometheusMonitorAttributeMetricResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrometheusMonitorAttributeMetricResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrometheusMonitorAttributeMetricResponse) SetId(v string) {
	o.Id = v
}

func (o PrometheusMonitorAttributeMetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusMonitorAttributeMetricResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullablePrometheusMonitorAttributeMetricResponse struct {
	value *PrometheusMonitorAttributeMetricResponse
	isSet bool
}

func (v NullablePrometheusMonitorAttributeMetricResponse) Get() *PrometheusMonitorAttributeMetricResponse {
	return v.value
}

func (v *NullablePrometheusMonitorAttributeMetricResponse) Set(val *PrometheusMonitorAttributeMetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusMonitorAttributeMetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusMonitorAttributeMetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusMonitorAttributeMetricResponse(val *PrometheusMonitorAttributeMetricResponse) *NullablePrometheusMonitorAttributeMetricResponse {
	return &NullablePrometheusMonitorAttributeMetricResponse{value: val, isSet: true}
}

func (v NullablePrometheusMonitorAttributeMetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusMonitorAttributeMetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
