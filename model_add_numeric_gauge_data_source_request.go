/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AddNumericGaugeDataSourceRequest struct for AddNumericGaugeDataSourceRequest
type AddNumericGaugeDataSourceRequest struct {
	// Name of the new Gauge Data Source
	SourceName string `json:"sourceName"`
	Schemas []EnumnumericGaugeDataSourceSchemaUrn `json:"schemas"`
	DataOrientation *EnumgaugeDataSourceDataOrientationProp `json:"dataOrientation,omitempty"`
	StatisticType EnumgaugeDataSourceStatisticTypeProp `json:"statisticType"`
	// An optional floating point value that can be used to scale the resulting value.
	DivideValueBy *float32 `json:"divideValueBy,omitempty"`
	// An optional property that can scale the resulting value by another attribute in the monitored entry.
	DivideValueByAttribute *string `json:"divideValueByAttribute,omitempty"`
	// An optional property that can scale the resulting value by another attribute whose value represents a counter in the monitored entry.
	DivideValueByCounterAttribute *string `json:"divideValueByCounterAttribute,omitempty"`
	// A description for this Gauge Data Source
	Description *string `json:"description,omitempty"`
	// Additional information about the source of this data that is added to alerts sent as a result of gauges that use this Gauge Data Source.
	AdditionalText *string `json:"additionalText,omitempty"`
	// The object class name of the monitor entries to examine for generating gauge data.
	MonitorObjectclass string `json:"monitorObjectclass"`
	// Specifies the attribute on the monitor entries from which to derive the current gauge value.
	MonitorAttribute string `json:"monitorAttribute"`
	// An optional LDAP filter that can be used restrict which monitor entries are used to compute output.
	IncludeFilter *string `json:"includeFilter,omitempty"`
	// Specifies the attribute whose value is used to identify the specific resource being monitored (e.g. device name).
	ResourceAttribute *string `json:"resourceAttribute,omitempty"`
	// A string indicating the type of resource being monitored.
	ResourceType *string `json:"resourceType,omitempty"`
	// The minimum frequency with which gauges using this Gauge Data Source can be configured for update. In order to prevent undesirable side effects, some Gauge Data Sources may use this property to impose a higher bound on the update frequency of gauges.
	MinimumUpdateInterval *string `json:"minimumUpdateInterval,omitempty"`
}

// NewAddNumericGaugeDataSourceRequest instantiates a new AddNumericGaugeDataSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNumericGaugeDataSourceRequest(sourceName string, schemas []EnumnumericGaugeDataSourceSchemaUrn, statisticType EnumgaugeDataSourceStatisticTypeProp, monitorObjectclass string, monitorAttribute string) *AddNumericGaugeDataSourceRequest {
	this := AddNumericGaugeDataSourceRequest{}
	this.SourceName = sourceName
	this.Schemas = schemas
	this.StatisticType = statisticType
	this.MonitorObjectclass = monitorObjectclass
	this.MonitorAttribute = monitorAttribute
	return &this
}

// NewAddNumericGaugeDataSourceRequestWithDefaults instantiates a new AddNumericGaugeDataSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNumericGaugeDataSourceRequestWithDefaults() *AddNumericGaugeDataSourceRequest {
	this := AddNumericGaugeDataSourceRequest{}
	return &this
}

// GetSourceName returns the SourceName field value
func (o *AddNumericGaugeDataSourceRequest) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetSourceNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *AddNumericGaugeDataSourceRequest) SetSourceName(v string) {
	o.SourceName = v
}

// GetSchemas returns the Schemas field value
func (o *AddNumericGaugeDataSourceRequest) GetSchemas() []EnumnumericGaugeDataSourceSchemaUrn {
	if o == nil {
		var ret []EnumnumericGaugeDataSourceSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetSchemasOk() ([]EnumnumericGaugeDataSourceSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddNumericGaugeDataSourceRequest) SetSchemas(v []EnumnumericGaugeDataSourceSchemaUrn) {
	o.Schemas = v
}

// GetDataOrientation returns the DataOrientation field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetDataOrientation() EnumgaugeDataSourceDataOrientationProp {
	if o == nil || isNil(o.DataOrientation) {
		var ret EnumgaugeDataSourceDataOrientationProp
		return ret
	}
	return *o.DataOrientation
}

// GetDataOrientationOk returns a tuple with the DataOrientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetDataOrientationOk() (*EnumgaugeDataSourceDataOrientationProp, bool) {
	if o == nil || isNil(o.DataOrientation) {
    return nil, false
	}
	return o.DataOrientation, true
}

// HasDataOrientation returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasDataOrientation() bool {
	if o != nil && !isNil(o.DataOrientation) {
		return true
	}

	return false
}

// SetDataOrientation gets a reference to the given EnumgaugeDataSourceDataOrientationProp and assigns it to the DataOrientation field.
func (o *AddNumericGaugeDataSourceRequest) SetDataOrientation(v EnumgaugeDataSourceDataOrientationProp) {
	o.DataOrientation = &v
}

// GetStatisticType returns the StatisticType field value
func (o *AddNumericGaugeDataSourceRequest) GetStatisticType() EnumgaugeDataSourceStatisticTypeProp {
	if o == nil {
		var ret EnumgaugeDataSourceStatisticTypeProp
		return ret
	}

	return o.StatisticType
}

// GetStatisticTypeOk returns a tuple with the StatisticType field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetStatisticTypeOk() (*EnumgaugeDataSourceStatisticTypeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StatisticType, true
}

// SetStatisticType sets field value
func (o *AddNumericGaugeDataSourceRequest) SetStatisticType(v EnumgaugeDataSourceStatisticTypeProp) {
	o.StatisticType = v
}

// GetDivideValueBy returns the DivideValueBy field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetDivideValueBy() float32 {
	if o == nil || isNil(o.DivideValueBy) {
		var ret float32
		return ret
	}
	return *o.DivideValueBy
}

// GetDivideValueByOk returns a tuple with the DivideValueBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetDivideValueByOk() (*float32, bool) {
	if o == nil || isNil(o.DivideValueBy) {
    return nil, false
	}
	return o.DivideValueBy, true
}

// HasDivideValueBy returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasDivideValueBy() bool {
	if o != nil && !isNil(o.DivideValueBy) {
		return true
	}

	return false
}

// SetDivideValueBy gets a reference to the given float32 and assigns it to the DivideValueBy field.
func (o *AddNumericGaugeDataSourceRequest) SetDivideValueBy(v float32) {
	o.DivideValueBy = &v
}

// GetDivideValueByAttribute returns the DivideValueByAttribute field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetDivideValueByAttribute() string {
	if o == nil || isNil(o.DivideValueByAttribute) {
		var ret string
		return ret
	}
	return *o.DivideValueByAttribute
}

// GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetDivideValueByAttributeOk() (*string, bool) {
	if o == nil || isNil(o.DivideValueByAttribute) {
    return nil, false
	}
	return o.DivideValueByAttribute, true
}

// HasDivideValueByAttribute returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasDivideValueByAttribute() bool {
	if o != nil && !isNil(o.DivideValueByAttribute) {
		return true
	}

	return false
}

// SetDivideValueByAttribute gets a reference to the given string and assigns it to the DivideValueByAttribute field.
func (o *AddNumericGaugeDataSourceRequest) SetDivideValueByAttribute(v string) {
	o.DivideValueByAttribute = &v
}

// GetDivideValueByCounterAttribute returns the DivideValueByCounterAttribute field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetDivideValueByCounterAttribute() string {
	if o == nil || isNil(o.DivideValueByCounterAttribute) {
		var ret string
		return ret
	}
	return *o.DivideValueByCounterAttribute
}

// GetDivideValueByCounterAttributeOk returns a tuple with the DivideValueByCounterAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetDivideValueByCounterAttributeOk() (*string, bool) {
	if o == nil || isNil(o.DivideValueByCounterAttribute) {
    return nil, false
	}
	return o.DivideValueByCounterAttribute, true
}

// HasDivideValueByCounterAttribute returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasDivideValueByCounterAttribute() bool {
	if o != nil && !isNil(o.DivideValueByCounterAttribute) {
		return true
	}

	return false
}

// SetDivideValueByCounterAttribute gets a reference to the given string and assigns it to the DivideValueByCounterAttribute field.
func (o *AddNumericGaugeDataSourceRequest) SetDivideValueByCounterAttribute(v string) {
	o.DivideValueByCounterAttribute = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddNumericGaugeDataSourceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *AddNumericGaugeDataSourceRequest) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetMonitorObjectclass returns the MonitorObjectclass field value
func (o *AddNumericGaugeDataSourceRequest) GetMonitorObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorObjectclass
}

// GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetMonitorObjectclassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MonitorObjectclass, true
}

// SetMonitorObjectclass sets field value
func (o *AddNumericGaugeDataSourceRequest) SetMonitorObjectclass(v string) {
	o.MonitorObjectclass = v
}

// GetMonitorAttribute returns the MonitorAttribute field value
func (o *AddNumericGaugeDataSourceRequest) GetMonitorAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorAttribute
}

// GetMonitorAttributeOk returns a tuple with the MonitorAttribute field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetMonitorAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MonitorAttribute, true
}

// SetMonitorAttribute sets field value
func (o *AddNumericGaugeDataSourceRequest) SetMonitorAttribute(v string) {
	o.MonitorAttribute = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetIncludeFilter() string {
	if o == nil || isNil(o.IncludeFilter) {
		var ret string
		return ret
	}
	return *o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetIncludeFilterOk() (*string, bool) {
	if o == nil || isNil(o.IncludeFilter) {
    return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasIncludeFilter() bool {
	if o != nil && !isNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given string and assigns it to the IncludeFilter field.
func (o *AddNumericGaugeDataSourceRequest) SetIncludeFilter(v string) {
	o.IncludeFilter = &v
}

// GetResourceAttribute returns the ResourceAttribute field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetResourceAttribute() string {
	if o == nil || isNil(o.ResourceAttribute) {
		var ret string
		return ret
	}
	return *o.ResourceAttribute
}

// GetResourceAttributeOk returns a tuple with the ResourceAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetResourceAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceAttribute) {
    return nil, false
	}
	return o.ResourceAttribute, true
}

// HasResourceAttribute returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasResourceAttribute() bool {
	if o != nil && !isNil(o.ResourceAttribute) {
		return true
	}

	return false
}

// SetResourceAttribute gets a reference to the given string and assigns it to the ResourceAttribute field.
func (o *AddNumericGaugeDataSourceRequest) SetResourceAttribute(v string) {
	o.ResourceAttribute = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
    return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AddNumericGaugeDataSourceRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetMinimumUpdateInterval returns the MinimumUpdateInterval field value if set, zero value otherwise.
func (o *AddNumericGaugeDataSourceRequest) GetMinimumUpdateInterval() string {
	if o == nil || isNil(o.MinimumUpdateInterval) {
		var ret string
		return ret
	}
	return *o.MinimumUpdateInterval
}

// GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeDataSourceRequest) GetMinimumUpdateIntervalOk() (*string, bool) {
	if o == nil || isNil(o.MinimumUpdateInterval) {
    return nil, false
	}
	return o.MinimumUpdateInterval, true
}

// HasMinimumUpdateInterval returns a boolean if a field has been set.
func (o *AddNumericGaugeDataSourceRequest) HasMinimumUpdateInterval() bool {
	if o != nil && !isNil(o.MinimumUpdateInterval) {
		return true
	}

	return false
}

// SetMinimumUpdateInterval gets a reference to the given string and assigns it to the MinimumUpdateInterval field.
func (o *AddNumericGaugeDataSourceRequest) SetMinimumUpdateInterval(v string) {
	o.MinimumUpdateInterval = &v
}

func (o AddNumericGaugeDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceName"] = o.SourceName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.DataOrientation) {
		toSerialize["dataOrientation"] = o.DataOrientation
	}
	if true {
		toSerialize["statisticType"] = o.StatisticType
	}
	if !isNil(o.DivideValueBy) {
		toSerialize["divideValueBy"] = o.DivideValueBy
	}
	if !isNil(o.DivideValueByAttribute) {
		toSerialize["divideValueByAttribute"] = o.DivideValueByAttribute
	}
	if !isNil(o.DivideValueByCounterAttribute) {
		toSerialize["divideValueByCounterAttribute"] = o.DivideValueByCounterAttribute
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	if true {
		toSerialize["monitorObjectclass"] = o.MonitorObjectclass
	}
	if true {
		toSerialize["monitorAttribute"] = o.MonitorAttribute
	}
	if !isNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !isNil(o.ResourceAttribute) {
		toSerialize["resourceAttribute"] = o.ResourceAttribute
	}
	if !isNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !isNil(o.MinimumUpdateInterval) {
		toSerialize["minimumUpdateInterval"] = o.MinimumUpdateInterval
	}
	return json.Marshal(toSerialize)
}

type NullableAddNumericGaugeDataSourceRequest struct {
	value *AddNumericGaugeDataSourceRequest
	isSet bool
}

func (v NullableAddNumericGaugeDataSourceRequest) Get() *AddNumericGaugeDataSourceRequest {
	return v.value
}

func (v *NullableAddNumericGaugeDataSourceRequest) Set(val *AddNumericGaugeDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNumericGaugeDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNumericGaugeDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNumericGaugeDataSourceRequest(val *AddNumericGaugeDataSourceRequest) *NullableAddNumericGaugeDataSourceRequest {
	return &NullableAddNumericGaugeDataSourceRequest{value: val, isSet: true}
}

func (v NullableAddNumericGaugeDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNumericGaugeDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


