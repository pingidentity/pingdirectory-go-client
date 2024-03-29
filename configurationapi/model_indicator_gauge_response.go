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

// checks if the IndicatorGaugeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorGaugeResponse{}

// IndicatorGaugeResponse struct for IndicatorGaugeResponse
type IndicatorGaugeResponse struct {
	Schemas []EnumindicatorGaugeSchemaUrn `json:"schemas"`
	// Specifies the source of data to use in determining this Indicator Gauge's severity and status.
	GaugeDataSource string `json:"gaugeDataSource"`
	// A regular expression pattern that is used to determine whether the current monitored value indicates this gauge's severity should be critical.
	CriticalValue *string `json:"criticalValue,omitempty"`
	// A regular expression pattern that is used to determine whether the current monitored value indicates this gauge's severity will be 'major'.
	MajorValue *string `json:"majorValue,omitempty"`
	// A regular expression pattern that is used to determine whether the current monitored value indicates this gauge's severity will be 'minor'.
	MinorValue *string `json:"minorValue,omitempty"`
	// A regular expression pattern that is used to determine whether the current monitored value indicates this gauge's severity will be 'warning'.
	WarningValue *string `json:"warningValue,omitempty"`
	// A description for this Gauge
	Description *string `json:"description,omitempty"`
	// Indicates whether this Gauge is enabled.
	Enabled          bool                           `json:"enabled"`
	OverrideSeverity *EnumgaugeOverrideSeverityProp `json:"overrideSeverity,omitempty"`
	AlertLevel       *EnumgaugeAlertLevelProp       `json:"alertLevel,omitempty"`
	// The frequency with which this Gauge is updated.
	UpdateInterval *string `json:"updateInterval,omitempty"`
	// Indicates the number of times the monitor data source value will be collected during the update interval.
	SamplesPerUpdateInterval *int64 `json:"samplesPerUpdateInterval,omitempty"`
	// Specifies set of resources to be monitored.
	IncludeResource []string `json:"includeResource,omitempty"`
	// Specifies resources to exclude from being monitored.
	ExcludeResource                               []string                                           `json:"excludeResource,omitempty"`
	ServerUnavailableSeverityLevel                *EnumgaugeServerUnavailableSeverityLevelProp       `json:"serverUnavailableSeverityLevel,omitempty"`
	ServerDegradedSeverityLevel                   *EnumgaugeServerDegradedSeverityLevelProp          `json:"serverDegradedSeverityLevel,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Gauge
	Id string `json:"id"`
}

// NewIndicatorGaugeResponse instantiates a new IndicatorGaugeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorGaugeResponse(schemas []EnumindicatorGaugeSchemaUrn, gaugeDataSource string, enabled bool, id string) *IndicatorGaugeResponse {
	this := IndicatorGaugeResponse{}
	this.Schemas = schemas
	this.GaugeDataSource = gaugeDataSource
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewIndicatorGaugeResponseWithDefaults instantiates a new IndicatorGaugeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorGaugeResponseWithDefaults() *IndicatorGaugeResponse {
	this := IndicatorGaugeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *IndicatorGaugeResponse) GetSchemas() []EnumindicatorGaugeSchemaUrn {
	if o == nil {
		var ret []EnumindicatorGaugeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetSchemasOk() ([]EnumindicatorGaugeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *IndicatorGaugeResponse) SetSchemas(v []EnumindicatorGaugeSchemaUrn) {
	o.Schemas = v
}

// GetGaugeDataSource returns the GaugeDataSource field value
func (o *IndicatorGaugeResponse) GetGaugeDataSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GaugeDataSource
}

// GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetGaugeDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GaugeDataSource, true
}

// SetGaugeDataSource sets field value
func (o *IndicatorGaugeResponse) SetGaugeDataSource(v string) {
	o.GaugeDataSource = v
}

// GetCriticalValue returns the CriticalValue field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetCriticalValue() string {
	if o == nil || IsNil(o.CriticalValue) {
		var ret string
		return ret
	}
	return *o.CriticalValue
}

// GetCriticalValueOk returns a tuple with the CriticalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetCriticalValueOk() (*string, bool) {
	if o == nil || IsNil(o.CriticalValue) {
		return nil, false
	}
	return o.CriticalValue, true
}

// HasCriticalValue returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasCriticalValue() bool {
	if o != nil && !IsNil(o.CriticalValue) {
		return true
	}

	return false
}

// SetCriticalValue gets a reference to the given string and assigns it to the CriticalValue field.
func (o *IndicatorGaugeResponse) SetCriticalValue(v string) {
	o.CriticalValue = &v
}

// GetMajorValue returns the MajorValue field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetMajorValue() string {
	if o == nil || IsNil(o.MajorValue) {
		var ret string
		return ret
	}
	return *o.MajorValue
}

// GetMajorValueOk returns a tuple with the MajorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetMajorValueOk() (*string, bool) {
	if o == nil || IsNil(o.MajorValue) {
		return nil, false
	}
	return o.MajorValue, true
}

// HasMajorValue returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasMajorValue() bool {
	if o != nil && !IsNil(o.MajorValue) {
		return true
	}

	return false
}

// SetMajorValue gets a reference to the given string and assigns it to the MajorValue field.
func (o *IndicatorGaugeResponse) SetMajorValue(v string) {
	o.MajorValue = &v
}

// GetMinorValue returns the MinorValue field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetMinorValue() string {
	if o == nil || IsNil(o.MinorValue) {
		var ret string
		return ret
	}
	return *o.MinorValue
}

// GetMinorValueOk returns a tuple with the MinorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetMinorValueOk() (*string, bool) {
	if o == nil || IsNil(o.MinorValue) {
		return nil, false
	}
	return o.MinorValue, true
}

// HasMinorValue returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasMinorValue() bool {
	if o != nil && !IsNil(o.MinorValue) {
		return true
	}

	return false
}

// SetMinorValue gets a reference to the given string and assigns it to the MinorValue field.
func (o *IndicatorGaugeResponse) SetMinorValue(v string) {
	o.MinorValue = &v
}

// GetWarningValue returns the WarningValue field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetWarningValue() string {
	if o == nil || IsNil(o.WarningValue) {
		var ret string
		return ret
	}
	return *o.WarningValue
}

// GetWarningValueOk returns a tuple with the WarningValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetWarningValueOk() (*string, bool) {
	if o == nil || IsNil(o.WarningValue) {
		return nil, false
	}
	return o.WarningValue, true
}

// HasWarningValue returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasWarningValue() bool {
	if o != nil && !IsNil(o.WarningValue) {
		return true
	}

	return false
}

// SetWarningValue gets a reference to the given string and assigns it to the WarningValue field.
func (o *IndicatorGaugeResponse) SetWarningValue(v string) {
	o.WarningValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IndicatorGaugeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *IndicatorGaugeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IndicatorGaugeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetOverrideSeverity() EnumgaugeOverrideSeverityProp {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret EnumgaugeOverrideSeverityProp
		return ret
	}
	return *o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given EnumgaugeOverrideSeverityProp and assigns it to the OverrideSeverity field.
func (o *IndicatorGaugeResponse) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp) {
	o.OverrideSeverity = &v
}

// GetAlertLevel returns the AlertLevel field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetAlertLevel() EnumgaugeAlertLevelProp {
	if o == nil || IsNil(o.AlertLevel) {
		var ret EnumgaugeAlertLevelProp
		return ret
	}
	return *o.AlertLevel
}

// GetAlertLevelOk returns a tuple with the AlertLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool) {
	if o == nil || IsNil(o.AlertLevel) {
		return nil, false
	}
	return o.AlertLevel, true
}

// HasAlertLevel returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasAlertLevel() bool {
	if o != nil && !IsNil(o.AlertLevel) {
		return true
	}

	return false
}

// SetAlertLevel gets a reference to the given EnumgaugeAlertLevelProp and assigns it to the AlertLevel field.
func (o *IndicatorGaugeResponse) SetAlertLevel(v EnumgaugeAlertLevelProp) {
	o.AlertLevel = &v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetUpdateInterval() string {
	if o == nil || IsNil(o.UpdateInterval) {
		var ret string
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetUpdateIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateInterval) {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasUpdateInterval() bool {
	if o != nil && !IsNil(o.UpdateInterval) {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given string and assigns it to the UpdateInterval field.
func (o *IndicatorGaugeResponse) SetUpdateInterval(v string) {
	o.UpdateInterval = &v
}

// GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetSamplesPerUpdateInterval() int64 {
	if o == nil || IsNil(o.SamplesPerUpdateInterval) {
		var ret int64
		return ret
	}
	return *o.SamplesPerUpdateInterval
}

// GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetSamplesPerUpdateIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.SamplesPerUpdateInterval) {
		return nil, false
	}
	return o.SamplesPerUpdateInterval, true
}

// HasSamplesPerUpdateInterval returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasSamplesPerUpdateInterval() bool {
	if o != nil && !IsNil(o.SamplesPerUpdateInterval) {
		return true
	}

	return false
}

// SetSamplesPerUpdateInterval gets a reference to the given int64 and assigns it to the SamplesPerUpdateInterval field.
func (o *IndicatorGaugeResponse) SetSamplesPerUpdateInterval(v int64) {
	o.SamplesPerUpdateInterval = &v
}

// GetIncludeResource returns the IncludeResource field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetIncludeResource() []string {
	if o == nil || IsNil(o.IncludeResource) {
		var ret []string
		return ret
	}
	return o.IncludeResource
}

// GetIncludeResourceOk returns a tuple with the IncludeResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetIncludeResourceOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeResource) {
		return nil, false
	}
	return o.IncludeResource, true
}

// HasIncludeResource returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasIncludeResource() bool {
	if o != nil && !IsNil(o.IncludeResource) {
		return true
	}

	return false
}

// SetIncludeResource gets a reference to the given []string and assigns it to the IncludeResource field.
func (o *IndicatorGaugeResponse) SetIncludeResource(v []string) {
	o.IncludeResource = v
}

// GetExcludeResource returns the ExcludeResource field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetExcludeResource() []string {
	if o == nil || IsNil(o.ExcludeResource) {
		var ret []string
		return ret
	}
	return o.ExcludeResource
}

// GetExcludeResourceOk returns a tuple with the ExcludeResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetExcludeResourceOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeResource) {
		return nil, false
	}
	return o.ExcludeResource, true
}

// HasExcludeResource returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasExcludeResource() bool {
	if o != nil && !IsNil(o.ExcludeResource) {
		return true
	}

	return false
}

// SetExcludeResource gets a reference to the given []string and assigns it to the ExcludeResource field.
func (o *IndicatorGaugeResponse) SetExcludeResource(v []string) {
	o.ExcludeResource = v
}

// GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp {
	if o == nil || IsNil(o.ServerUnavailableSeverityLevel) {
		var ret EnumgaugeServerUnavailableSeverityLevelProp
		return ret
	}
	return *o.ServerUnavailableSeverityLevel
}

// GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool) {
	if o == nil || IsNil(o.ServerUnavailableSeverityLevel) {
		return nil, false
	}
	return o.ServerUnavailableSeverityLevel, true
}

// HasServerUnavailableSeverityLevel returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasServerUnavailableSeverityLevel() bool {
	if o != nil && !IsNil(o.ServerUnavailableSeverityLevel) {
		return true
	}

	return false
}

// SetServerUnavailableSeverityLevel gets a reference to the given EnumgaugeServerUnavailableSeverityLevelProp and assigns it to the ServerUnavailableSeverityLevel field.
func (o *IndicatorGaugeResponse) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp) {
	o.ServerUnavailableSeverityLevel = &v
}

// GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp {
	if o == nil || IsNil(o.ServerDegradedSeverityLevel) {
		var ret EnumgaugeServerDegradedSeverityLevelProp
		return ret
	}
	return *o.ServerDegradedSeverityLevel
}

// GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool) {
	if o == nil || IsNil(o.ServerDegradedSeverityLevel) {
		return nil, false
	}
	return o.ServerDegradedSeverityLevel, true
}

// HasServerDegradedSeverityLevel returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasServerDegradedSeverityLevel() bool {
	if o != nil && !IsNil(o.ServerDegradedSeverityLevel) {
		return true
	}

	return false
}

// SetServerDegradedSeverityLevel gets a reference to the given EnumgaugeServerDegradedSeverityLevelProp and assigns it to the ServerDegradedSeverityLevel field.
func (o *IndicatorGaugeResponse) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp) {
	o.ServerDegradedSeverityLevel = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *IndicatorGaugeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *IndicatorGaugeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *IndicatorGaugeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *IndicatorGaugeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *IndicatorGaugeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IndicatorGaugeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IndicatorGaugeResponse) SetId(v string) {
	o.Id = v
}

func (o IndicatorGaugeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorGaugeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["gaugeDataSource"] = o.GaugeDataSource
	if !IsNil(o.CriticalValue) {
		toSerialize["criticalValue"] = o.CriticalValue
	}
	if !IsNil(o.MajorValue) {
		toSerialize["majorValue"] = o.MajorValue
	}
	if !IsNil(o.MinorValue) {
		toSerialize["minorValue"] = o.MinorValue
	}
	if !IsNil(o.WarningValue) {
		toSerialize["warningValue"] = o.WarningValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.OverrideSeverity) {
		toSerialize["overrideSeverity"] = o.OverrideSeverity
	}
	if !IsNil(o.AlertLevel) {
		toSerialize["alertLevel"] = o.AlertLevel
	}
	if !IsNil(o.UpdateInterval) {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if !IsNil(o.SamplesPerUpdateInterval) {
		toSerialize["samplesPerUpdateInterval"] = o.SamplesPerUpdateInterval
	}
	if !IsNil(o.IncludeResource) {
		toSerialize["includeResource"] = o.IncludeResource
	}
	if !IsNil(o.ExcludeResource) {
		toSerialize["excludeResource"] = o.ExcludeResource
	}
	if !IsNil(o.ServerUnavailableSeverityLevel) {
		toSerialize["serverUnavailableSeverityLevel"] = o.ServerUnavailableSeverityLevel
	}
	if !IsNil(o.ServerDegradedSeverityLevel) {
		toSerialize["serverDegradedSeverityLevel"] = o.ServerDegradedSeverityLevel
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

type NullableIndicatorGaugeResponse struct {
	value *IndicatorGaugeResponse
	isSet bool
}

func (v NullableIndicatorGaugeResponse) Get() *IndicatorGaugeResponse {
	return v.value
}

func (v *NullableIndicatorGaugeResponse) Set(val *IndicatorGaugeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorGaugeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorGaugeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorGaugeResponse(val *IndicatorGaugeResponse) *NullableIndicatorGaugeResponse {
	return &NullableIndicatorGaugeResponse{value: val, isSet: true}
}

func (v NullableIndicatorGaugeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorGaugeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
