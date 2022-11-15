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

// AddIndicatorGaugeDataSourceRequest struct for AddIndicatorGaugeDataSourceRequest
type AddIndicatorGaugeDataSourceRequest struct {
	// Name of the new Gauge Data Source
	SourceName string `json:"sourceName"`
	Schemas []EnumindicatorGaugeDataSourceSchemaUrn `json:"schemas"`
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

// NewAddIndicatorGaugeDataSourceRequest instantiates a new AddIndicatorGaugeDataSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIndicatorGaugeDataSourceRequest(sourceName string, schemas []EnumindicatorGaugeDataSourceSchemaUrn, monitorObjectclass string, monitorAttribute string) *AddIndicatorGaugeDataSourceRequest {
	this := AddIndicatorGaugeDataSourceRequest{}
	this.SourceName = sourceName
	this.Schemas = schemas
	this.MonitorObjectclass = monitorObjectclass
	this.MonitorAttribute = monitorAttribute
	return &this
}

// NewAddIndicatorGaugeDataSourceRequestWithDefaults instantiates a new AddIndicatorGaugeDataSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIndicatorGaugeDataSourceRequestWithDefaults() *AddIndicatorGaugeDataSourceRequest {
	this := AddIndicatorGaugeDataSourceRequest{}
	return &this
}

// GetSourceName returns the SourceName field value
func (o *AddIndicatorGaugeDataSourceRequest) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetSourceNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *AddIndicatorGaugeDataSourceRequest) SetSourceName(v string) {
	o.SourceName = v
}

// GetSchemas returns the Schemas field value
func (o *AddIndicatorGaugeDataSourceRequest) GetSchemas() []EnumindicatorGaugeDataSourceSchemaUrn {
	if o == nil {
		var ret []EnumindicatorGaugeDataSourceSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetSchemasOk() ([]EnumindicatorGaugeDataSourceSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddIndicatorGaugeDataSourceRequest) SetSchemas(v []EnumindicatorGaugeDataSourceSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddIndicatorGaugeDataSourceRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddIndicatorGaugeDataSourceRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddIndicatorGaugeDataSourceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *AddIndicatorGaugeDataSourceRequest) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *AddIndicatorGaugeDataSourceRequest) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *AddIndicatorGaugeDataSourceRequest) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetMonitorObjectclass returns the MonitorObjectclass field value
func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorObjectclass
}

// GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field value
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorObjectclassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MonitorObjectclass, true
}

// SetMonitorObjectclass sets field value
func (o *AddIndicatorGaugeDataSourceRequest) SetMonitorObjectclass(v string) {
	o.MonitorObjectclass = v
}

// GetMonitorAttribute returns the MonitorAttribute field value
func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorAttribute
}

// GetMonitorAttributeOk returns a tuple with the MonitorAttribute field value
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MonitorAttribute, true
}

// SetMonitorAttribute sets field value
func (o *AddIndicatorGaugeDataSourceRequest) SetMonitorAttribute(v string) {
	o.MonitorAttribute = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddIndicatorGaugeDataSourceRequest) GetIncludeFilter() string {
	if o == nil || isNil(o.IncludeFilter) {
		var ret string
		return ret
	}
	return *o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetIncludeFilterOk() (*string, bool) {
	if o == nil || isNil(o.IncludeFilter) {
    return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddIndicatorGaugeDataSourceRequest) HasIncludeFilter() bool {
	if o != nil && !isNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given string and assigns it to the IncludeFilter field.
func (o *AddIndicatorGaugeDataSourceRequest) SetIncludeFilter(v string) {
	o.IncludeFilter = &v
}

// GetResourceAttribute returns the ResourceAttribute field value if set, zero value otherwise.
func (o *AddIndicatorGaugeDataSourceRequest) GetResourceAttribute() string {
	if o == nil || isNil(o.ResourceAttribute) {
		var ret string
		return ret
	}
	return *o.ResourceAttribute
}

// GetResourceAttributeOk returns a tuple with the ResourceAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetResourceAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceAttribute) {
    return nil, false
	}
	return o.ResourceAttribute, true
}

// HasResourceAttribute returns a boolean if a field has been set.
func (o *AddIndicatorGaugeDataSourceRequest) HasResourceAttribute() bool {
	if o != nil && !isNil(o.ResourceAttribute) {
		return true
	}

	return false
}

// SetResourceAttribute gets a reference to the given string and assigns it to the ResourceAttribute field.
func (o *AddIndicatorGaugeDataSourceRequest) SetResourceAttribute(v string) {
	o.ResourceAttribute = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AddIndicatorGaugeDataSourceRequest) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
    return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AddIndicatorGaugeDataSourceRequest) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AddIndicatorGaugeDataSourceRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetMinimumUpdateInterval returns the MinimumUpdateInterval field value if set, zero value otherwise.
func (o *AddIndicatorGaugeDataSourceRequest) GetMinimumUpdateInterval() string {
	if o == nil || isNil(o.MinimumUpdateInterval) {
		var ret string
		return ret
	}
	return *o.MinimumUpdateInterval
}

// GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIndicatorGaugeDataSourceRequest) GetMinimumUpdateIntervalOk() (*string, bool) {
	if o == nil || isNil(o.MinimumUpdateInterval) {
    return nil, false
	}
	return o.MinimumUpdateInterval, true
}

// HasMinimumUpdateInterval returns a boolean if a field has been set.
func (o *AddIndicatorGaugeDataSourceRequest) HasMinimumUpdateInterval() bool {
	if o != nil && !isNil(o.MinimumUpdateInterval) {
		return true
	}

	return false
}

// SetMinimumUpdateInterval gets a reference to the given string and assigns it to the MinimumUpdateInterval field.
func (o *AddIndicatorGaugeDataSourceRequest) SetMinimumUpdateInterval(v string) {
	o.MinimumUpdateInterval = &v
}

func (o AddIndicatorGaugeDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceName"] = o.SourceName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
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

type NullableAddIndicatorGaugeDataSourceRequest struct {
	value *AddIndicatorGaugeDataSourceRequest
	isSet bool
}

func (v NullableAddIndicatorGaugeDataSourceRequest) Get() *AddIndicatorGaugeDataSourceRequest {
	return v.value
}

func (v *NullableAddIndicatorGaugeDataSourceRequest) Set(val *AddIndicatorGaugeDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIndicatorGaugeDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIndicatorGaugeDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIndicatorGaugeDataSourceRequest(val *AddIndicatorGaugeDataSourceRequest) *NullableAddIndicatorGaugeDataSourceRequest {
	return &NullableAddIndicatorGaugeDataSourceRequest{value: val, isSet: true}
}

func (v NullableAddIndicatorGaugeDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIndicatorGaugeDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


