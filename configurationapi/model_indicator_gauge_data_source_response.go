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

// IndicatorGaugeDataSourceResponse struct for IndicatorGaugeDataSourceResponse
type IndicatorGaugeDataSourceResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Gauge Data Source
	Id      string                                  `json:"id"`
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

// NewIndicatorGaugeDataSourceResponse instantiates a new IndicatorGaugeDataSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorGaugeDataSourceResponse(id string, schemas []EnumindicatorGaugeDataSourceSchemaUrn, monitorObjectclass string, monitorAttribute string) *IndicatorGaugeDataSourceResponse {
	this := IndicatorGaugeDataSourceResponse{}
	this.Id = id
	this.Schemas = schemas
	this.MonitorObjectclass = monitorObjectclass
	this.MonitorAttribute = monitorAttribute
	return &this
}

// NewIndicatorGaugeDataSourceResponseWithDefaults instantiates a new IndicatorGaugeDataSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorGaugeDataSourceResponseWithDefaults() *IndicatorGaugeDataSourceResponse {
	this := IndicatorGaugeDataSourceResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *IndicatorGaugeDataSourceResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *IndicatorGaugeDataSourceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *IndicatorGaugeDataSourceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IndicatorGaugeDataSourceResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *IndicatorGaugeDataSourceResponse) GetSchemas() []EnumindicatorGaugeDataSourceSchemaUrn {
	if o == nil {
		var ret []EnumindicatorGaugeDataSourceSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetSchemasOk() ([]EnumindicatorGaugeDataSourceSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *IndicatorGaugeDataSourceResponse) SetSchemas(v []EnumindicatorGaugeDataSourceSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IndicatorGaugeDataSourceResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *IndicatorGaugeDataSourceResponse) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetMonitorObjectclass returns the MonitorObjectclass field value
func (o *IndicatorGaugeDataSourceResponse) GetMonitorObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorObjectclass
}

// GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetMonitorObjectclassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorObjectclass, true
}

// SetMonitorObjectclass sets field value
func (o *IndicatorGaugeDataSourceResponse) SetMonitorObjectclass(v string) {
	o.MonitorObjectclass = v
}

// GetMonitorAttribute returns the MonitorAttribute field value
func (o *IndicatorGaugeDataSourceResponse) GetMonitorAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorAttribute
}

// GetMonitorAttributeOk returns a tuple with the MonitorAttribute field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetMonitorAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorAttribute, true
}

// SetMonitorAttribute sets field value
func (o *IndicatorGaugeDataSourceResponse) SetMonitorAttribute(v string) {
	o.MonitorAttribute = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetIncludeFilter() string {
	if o == nil || isNil(o.IncludeFilter) {
		var ret string
		return ret
	}
	return *o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetIncludeFilterOk() (*string, bool) {
	if o == nil || isNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasIncludeFilter() bool {
	if o != nil && !isNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given string and assigns it to the IncludeFilter field.
func (o *IndicatorGaugeDataSourceResponse) SetIncludeFilter(v string) {
	o.IncludeFilter = &v
}

// GetResourceAttribute returns the ResourceAttribute field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetResourceAttribute() string {
	if o == nil || isNil(o.ResourceAttribute) {
		var ret string
		return ret
	}
	return *o.ResourceAttribute
}

// GetResourceAttributeOk returns a tuple with the ResourceAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetResourceAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceAttribute) {
		return nil, false
	}
	return o.ResourceAttribute, true
}

// HasResourceAttribute returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasResourceAttribute() bool {
	if o != nil && !isNil(o.ResourceAttribute) {
		return true
	}

	return false
}

// SetResourceAttribute gets a reference to the given string and assigns it to the ResourceAttribute field.
func (o *IndicatorGaugeDataSourceResponse) SetResourceAttribute(v string) {
	o.ResourceAttribute = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *IndicatorGaugeDataSourceResponse) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetMinimumUpdateInterval returns the MinimumUpdateInterval field value if set, zero value otherwise.
func (o *IndicatorGaugeDataSourceResponse) GetMinimumUpdateInterval() string {
	if o == nil || isNil(o.MinimumUpdateInterval) {
		var ret string
		return ret
	}
	return *o.MinimumUpdateInterval
}

// GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeDataSourceResponse) GetMinimumUpdateIntervalOk() (*string, bool) {
	if o == nil || isNil(o.MinimumUpdateInterval) {
		return nil, false
	}
	return o.MinimumUpdateInterval, true
}

// HasMinimumUpdateInterval returns a boolean if a field has been set.
func (o *IndicatorGaugeDataSourceResponse) HasMinimumUpdateInterval() bool {
	if o != nil && !isNil(o.MinimumUpdateInterval) {
		return true
	}

	return false
}

// SetMinimumUpdateInterval gets a reference to the given string and assigns it to the MinimumUpdateInterval field.
func (o *IndicatorGaugeDataSourceResponse) SetMinimumUpdateInterval(v string) {
	o.MinimumUpdateInterval = &v
}

func (o IndicatorGaugeDataSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
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

type NullableIndicatorGaugeDataSourceResponse struct {
	value *IndicatorGaugeDataSourceResponse
	isSet bool
}

func (v NullableIndicatorGaugeDataSourceResponse) Get() *IndicatorGaugeDataSourceResponse {
	return v.value
}

func (v *NullableIndicatorGaugeDataSourceResponse) Set(val *IndicatorGaugeDataSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorGaugeDataSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorGaugeDataSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorGaugeDataSourceResponse(val *IndicatorGaugeDataSourceResponse) *NullableIndicatorGaugeDataSourceResponse {
	return &NullableIndicatorGaugeDataSourceResponse{value: val, isSet: true}
}

func (v NullableIndicatorGaugeDataSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorGaugeDataSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}