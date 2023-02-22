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

// AddNumericGaugeRequest struct for AddNumericGaugeRequest
type AddNumericGaugeRequest struct {
	// Name of the new Gauge
	GaugeName string                      `json:"gaugeName"`
	Schemas   []EnumnumericGaugeSchemaUrn `json:"schemas"`
	// Specifies the source of data to use in determining this gauge's current severity.
	GaugeDataSource string `json:"gaugeDataSource"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should be 'critical'.
	CriticalValue *float32 `json:"criticalValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should no longer be 'critical'.
	CriticalExitValue *float32 `json:"criticalExitValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should be 'major'.
	MajorValue *float32 `json:"majorValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should no longer be 'major'.
	MajorExitValue *float32 `json:"majorExitValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should be 'minor'.
	MinorValue *float32 `json:"minorValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should no longer be 'minor'.
	MinorExitValue *float32 `json:"minorExitValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should be 'warning'.
	WarningValue *float32 `json:"warningValue,omitempty"`
	// A value that is used to determine whether the current monitored value indicates this gauge's severity should no longer be 'warning'.
	WarningExitValue *float32 `json:"warningExitValue,omitempty"`
	// A description for this Gauge
	Description *string `json:"description,omitempty"`
	// Indicates whether this Gauge is enabled.
	Enabled          bool                           `json:"enabled"`
	OverrideSeverity *EnumgaugeOverrideSeverityProp `json:"overrideSeverity,omitempty"`
	AlertLevel       *EnumgaugeAlertLevelProp       `json:"alertLevel,omitempty"`
	// The frequency with which this Gauge is updated.
	UpdateInterval *string `json:"updateInterval,omitempty"`
	// Indicates the number of times the monitor data source value will be collected during the update interval.
	SamplesPerUpdateInterval *int32 `json:"samplesPerUpdateInterval,omitempty"`
	// Specifies set of resources to be monitored.
	IncludeResource []string `json:"includeResource,omitempty"`
	// Specifies resources to exclude from being monitored.
	ExcludeResource                []string                                     `json:"excludeResource,omitempty"`
	ServerUnavailableSeverityLevel *EnumgaugeServerUnavailableSeverityLevelProp `json:"serverUnavailableSeverityLevel,omitempty"`
	ServerDegradedSeverityLevel    *EnumgaugeServerDegradedSeverityLevelProp    `json:"serverDegradedSeverityLevel,omitempty"`
}

// NewAddNumericGaugeRequest instantiates a new AddNumericGaugeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNumericGaugeRequest(gaugeName string, schemas []EnumnumericGaugeSchemaUrn, gaugeDataSource string, enabled bool) *AddNumericGaugeRequest {
	this := AddNumericGaugeRequest{}
	this.GaugeName = gaugeName
	this.Schemas = schemas
	this.GaugeDataSource = gaugeDataSource
	this.Enabled = enabled
	return &this
}

// NewAddNumericGaugeRequestWithDefaults instantiates a new AddNumericGaugeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNumericGaugeRequestWithDefaults() *AddNumericGaugeRequest {
	this := AddNumericGaugeRequest{}
	return &this
}

// GetGaugeName returns the GaugeName field value
func (o *AddNumericGaugeRequest) GetGaugeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GaugeName
}

// GetGaugeNameOk returns a tuple with the GaugeName field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetGaugeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GaugeName, true
}

// SetGaugeName sets field value
func (o *AddNumericGaugeRequest) SetGaugeName(v string) {
	o.GaugeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddNumericGaugeRequest) GetSchemas() []EnumnumericGaugeSchemaUrn {
	if o == nil {
		var ret []EnumnumericGaugeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetSchemasOk() ([]EnumnumericGaugeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddNumericGaugeRequest) SetSchemas(v []EnumnumericGaugeSchemaUrn) {
	o.Schemas = v
}

// GetGaugeDataSource returns the GaugeDataSource field value
func (o *AddNumericGaugeRequest) GetGaugeDataSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GaugeDataSource
}

// GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetGaugeDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GaugeDataSource, true
}

// SetGaugeDataSource sets field value
func (o *AddNumericGaugeRequest) SetGaugeDataSource(v string) {
	o.GaugeDataSource = v
}

// GetCriticalValue returns the CriticalValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetCriticalValue() float32 {
	if o == nil || isNil(o.CriticalValue) {
		var ret float32
		return ret
	}
	return *o.CriticalValue
}

// GetCriticalValueOk returns a tuple with the CriticalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetCriticalValueOk() (*float32, bool) {
	if o == nil || isNil(o.CriticalValue) {
		return nil, false
	}
	return o.CriticalValue, true
}

// HasCriticalValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasCriticalValue() bool {
	if o != nil && !isNil(o.CriticalValue) {
		return true
	}

	return false
}

// SetCriticalValue gets a reference to the given float32 and assigns it to the CriticalValue field.
func (o *AddNumericGaugeRequest) SetCriticalValue(v float32) {
	o.CriticalValue = &v
}

// GetCriticalExitValue returns the CriticalExitValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetCriticalExitValue() float32 {
	if o == nil || isNil(o.CriticalExitValue) {
		var ret float32
		return ret
	}
	return *o.CriticalExitValue
}

// GetCriticalExitValueOk returns a tuple with the CriticalExitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetCriticalExitValueOk() (*float32, bool) {
	if o == nil || isNil(o.CriticalExitValue) {
		return nil, false
	}
	return o.CriticalExitValue, true
}

// HasCriticalExitValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasCriticalExitValue() bool {
	if o != nil && !isNil(o.CriticalExitValue) {
		return true
	}

	return false
}

// SetCriticalExitValue gets a reference to the given float32 and assigns it to the CriticalExitValue field.
func (o *AddNumericGaugeRequest) SetCriticalExitValue(v float32) {
	o.CriticalExitValue = &v
}

// GetMajorValue returns the MajorValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetMajorValue() float32 {
	if o == nil || isNil(o.MajorValue) {
		var ret float32
		return ret
	}
	return *o.MajorValue
}

// GetMajorValueOk returns a tuple with the MajorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetMajorValueOk() (*float32, bool) {
	if o == nil || isNil(o.MajorValue) {
		return nil, false
	}
	return o.MajorValue, true
}

// HasMajorValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasMajorValue() bool {
	if o != nil && !isNil(o.MajorValue) {
		return true
	}

	return false
}

// SetMajorValue gets a reference to the given float32 and assigns it to the MajorValue field.
func (o *AddNumericGaugeRequest) SetMajorValue(v float32) {
	o.MajorValue = &v
}

// GetMajorExitValue returns the MajorExitValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetMajorExitValue() float32 {
	if o == nil || isNil(o.MajorExitValue) {
		var ret float32
		return ret
	}
	return *o.MajorExitValue
}

// GetMajorExitValueOk returns a tuple with the MajorExitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetMajorExitValueOk() (*float32, bool) {
	if o == nil || isNil(o.MajorExitValue) {
		return nil, false
	}
	return o.MajorExitValue, true
}

// HasMajorExitValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasMajorExitValue() bool {
	if o != nil && !isNil(o.MajorExitValue) {
		return true
	}

	return false
}

// SetMajorExitValue gets a reference to the given float32 and assigns it to the MajorExitValue field.
func (o *AddNumericGaugeRequest) SetMajorExitValue(v float32) {
	o.MajorExitValue = &v
}

// GetMinorValue returns the MinorValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetMinorValue() float32 {
	if o == nil || isNil(o.MinorValue) {
		var ret float32
		return ret
	}
	return *o.MinorValue
}

// GetMinorValueOk returns a tuple with the MinorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetMinorValueOk() (*float32, bool) {
	if o == nil || isNil(o.MinorValue) {
		return nil, false
	}
	return o.MinorValue, true
}

// HasMinorValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasMinorValue() bool {
	if o != nil && !isNil(o.MinorValue) {
		return true
	}

	return false
}

// SetMinorValue gets a reference to the given float32 and assigns it to the MinorValue field.
func (o *AddNumericGaugeRequest) SetMinorValue(v float32) {
	o.MinorValue = &v
}

// GetMinorExitValue returns the MinorExitValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetMinorExitValue() float32 {
	if o == nil || isNil(o.MinorExitValue) {
		var ret float32
		return ret
	}
	return *o.MinorExitValue
}

// GetMinorExitValueOk returns a tuple with the MinorExitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetMinorExitValueOk() (*float32, bool) {
	if o == nil || isNil(o.MinorExitValue) {
		return nil, false
	}
	return o.MinorExitValue, true
}

// HasMinorExitValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasMinorExitValue() bool {
	if o != nil && !isNil(o.MinorExitValue) {
		return true
	}

	return false
}

// SetMinorExitValue gets a reference to the given float32 and assigns it to the MinorExitValue field.
func (o *AddNumericGaugeRequest) SetMinorExitValue(v float32) {
	o.MinorExitValue = &v
}

// GetWarningValue returns the WarningValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetWarningValue() float32 {
	if o == nil || isNil(o.WarningValue) {
		var ret float32
		return ret
	}
	return *o.WarningValue
}

// GetWarningValueOk returns a tuple with the WarningValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetWarningValueOk() (*float32, bool) {
	if o == nil || isNil(o.WarningValue) {
		return nil, false
	}
	return o.WarningValue, true
}

// HasWarningValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasWarningValue() bool {
	if o != nil && !isNil(o.WarningValue) {
		return true
	}

	return false
}

// SetWarningValue gets a reference to the given float32 and assigns it to the WarningValue field.
func (o *AddNumericGaugeRequest) SetWarningValue(v float32) {
	o.WarningValue = &v
}

// GetWarningExitValue returns the WarningExitValue field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetWarningExitValue() float32 {
	if o == nil || isNil(o.WarningExitValue) {
		var ret float32
		return ret
	}
	return *o.WarningExitValue
}

// GetWarningExitValueOk returns a tuple with the WarningExitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetWarningExitValueOk() (*float32, bool) {
	if o == nil || isNil(o.WarningExitValue) {
		return nil, false
	}
	return o.WarningExitValue, true
}

// HasWarningExitValue returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasWarningExitValue() bool {
	if o != nil && !isNil(o.WarningExitValue) {
		return true
	}

	return false
}

// SetWarningExitValue gets a reference to the given float32 and assigns it to the WarningExitValue field.
func (o *AddNumericGaugeRequest) SetWarningExitValue(v float32) {
	o.WarningExitValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddNumericGaugeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddNumericGaugeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddNumericGaugeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetOverrideSeverity() EnumgaugeOverrideSeverityProp {
	if o == nil || isNil(o.OverrideSeverity) {
		var ret EnumgaugeOverrideSeverityProp
		return ret
	}
	return *o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool) {
	if o == nil || isNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasOverrideSeverity() bool {
	if o != nil && !isNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given EnumgaugeOverrideSeverityProp and assigns it to the OverrideSeverity field.
func (o *AddNumericGaugeRequest) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp) {
	o.OverrideSeverity = &v
}

// GetAlertLevel returns the AlertLevel field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetAlertLevel() EnumgaugeAlertLevelProp {
	if o == nil || isNil(o.AlertLevel) {
		var ret EnumgaugeAlertLevelProp
		return ret
	}
	return *o.AlertLevel
}

// GetAlertLevelOk returns a tuple with the AlertLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool) {
	if o == nil || isNil(o.AlertLevel) {
		return nil, false
	}
	return o.AlertLevel, true
}

// HasAlertLevel returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasAlertLevel() bool {
	if o != nil && !isNil(o.AlertLevel) {
		return true
	}

	return false
}

// SetAlertLevel gets a reference to the given EnumgaugeAlertLevelProp and assigns it to the AlertLevel field.
func (o *AddNumericGaugeRequest) SetAlertLevel(v EnumgaugeAlertLevelProp) {
	o.AlertLevel = &v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetUpdateInterval() string {
	if o == nil || isNil(o.UpdateInterval) {
		var ret string
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetUpdateIntervalOk() (*string, bool) {
	if o == nil || isNil(o.UpdateInterval) {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasUpdateInterval() bool {
	if o != nil && !isNil(o.UpdateInterval) {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given string and assigns it to the UpdateInterval field.
func (o *AddNumericGaugeRequest) SetUpdateInterval(v string) {
	o.UpdateInterval = &v
}

// GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetSamplesPerUpdateInterval() int32 {
	if o == nil || isNil(o.SamplesPerUpdateInterval) {
		var ret int32
		return ret
	}
	return *o.SamplesPerUpdateInterval
}

// GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetSamplesPerUpdateIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.SamplesPerUpdateInterval) {
		return nil, false
	}
	return o.SamplesPerUpdateInterval, true
}

// HasSamplesPerUpdateInterval returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasSamplesPerUpdateInterval() bool {
	if o != nil && !isNil(o.SamplesPerUpdateInterval) {
		return true
	}

	return false
}

// SetSamplesPerUpdateInterval gets a reference to the given int32 and assigns it to the SamplesPerUpdateInterval field.
func (o *AddNumericGaugeRequest) SetSamplesPerUpdateInterval(v int32) {
	o.SamplesPerUpdateInterval = &v
}

// GetIncludeResource returns the IncludeResource field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetIncludeResource() []string {
	if o == nil || isNil(o.IncludeResource) {
		var ret []string
		return ret
	}
	return o.IncludeResource
}

// GetIncludeResourceOk returns a tuple with the IncludeResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetIncludeResourceOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeResource) {
		return nil, false
	}
	return o.IncludeResource, true
}

// HasIncludeResource returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasIncludeResource() bool {
	if o != nil && !isNil(o.IncludeResource) {
		return true
	}

	return false
}

// SetIncludeResource gets a reference to the given []string and assigns it to the IncludeResource field.
func (o *AddNumericGaugeRequest) SetIncludeResource(v []string) {
	o.IncludeResource = v
}

// GetExcludeResource returns the ExcludeResource field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetExcludeResource() []string {
	if o == nil || isNil(o.ExcludeResource) {
		var ret []string
		return ret
	}
	return o.ExcludeResource
}

// GetExcludeResourceOk returns a tuple with the ExcludeResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetExcludeResourceOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeResource) {
		return nil, false
	}
	return o.ExcludeResource, true
}

// HasExcludeResource returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasExcludeResource() bool {
	if o != nil && !isNil(o.ExcludeResource) {
		return true
	}

	return false
}

// SetExcludeResource gets a reference to the given []string and assigns it to the ExcludeResource field.
func (o *AddNumericGaugeRequest) SetExcludeResource(v []string) {
	o.ExcludeResource = v
}

// GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp {
	if o == nil || isNil(o.ServerUnavailableSeverityLevel) {
		var ret EnumgaugeServerUnavailableSeverityLevelProp
		return ret
	}
	return *o.ServerUnavailableSeverityLevel
}

// GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool) {
	if o == nil || isNil(o.ServerUnavailableSeverityLevel) {
		return nil, false
	}
	return o.ServerUnavailableSeverityLevel, true
}

// HasServerUnavailableSeverityLevel returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasServerUnavailableSeverityLevel() bool {
	if o != nil && !isNil(o.ServerUnavailableSeverityLevel) {
		return true
	}

	return false
}

// SetServerUnavailableSeverityLevel gets a reference to the given EnumgaugeServerUnavailableSeverityLevelProp and assigns it to the ServerUnavailableSeverityLevel field.
func (o *AddNumericGaugeRequest) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp) {
	o.ServerUnavailableSeverityLevel = &v
}

// GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field value if set, zero value otherwise.
func (o *AddNumericGaugeRequest) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp {
	if o == nil || isNil(o.ServerDegradedSeverityLevel) {
		var ret EnumgaugeServerDegradedSeverityLevelProp
		return ret
	}
	return *o.ServerDegradedSeverityLevel
}

// GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNumericGaugeRequest) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool) {
	if o == nil || isNil(o.ServerDegradedSeverityLevel) {
		return nil, false
	}
	return o.ServerDegradedSeverityLevel, true
}

// HasServerDegradedSeverityLevel returns a boolean if a field has been set.
func (o *AddNumericGaugeRequest) HasServerDegradedSeverityLevel() bool {
	if o != nil && !isNil(o.ServerDegradedSeverityLevel) {
		return true
	}

	return false
}

// SetServerDegradedSeverityLevel gets a reference to the given EnumgaugeServerDegradedSeverityLevelProp and assigns it to the ServerDegradedSeverityLevel field.
func (o *AddNumericGaugeRequest) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp) {
	o.ServerDegradedSeverityLevel = &v
}

func (o AddNumericGaugeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gaugeName"] = o.GaugeName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["gaugeDataSource"] = o.GaugeDataSource
	}
	if !isNil(o.CriticalValue) {
		toSerialize["criticalValue"] = o.CriticalValue
	}
	if !isNil(o.CriticalExitValue) {
		toSerialize["criticalExitValue"] = o.CriticalExitValue
	}
	if !isNil(o.MajorValue) {
		toSerialize["majorValue"] = o.MajorValue
	}
	if !isNil(o.MajorExitValue) {
		toSerialize["majorExitValue"] = o.MajorExitValue
	}
	if !isNil(o.MinorValue) {
		toSerialize["minorValue"] = o.MinorValue
	}
	if !isNil(o.MinorExitValue) {
		toSerialize["minorExitValue"] = o.MinorExitValue
	}
	if !isNil(o.WarningValue) {
		toSerialize["warningValue"] = o.WarningValue
	}
	if !isNil(o.WarningExitValue) {
		toSerialize["warningExitValue"] = o.WarningExitValue
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.OverrideSeverity) {
		toSerialize["overrideSeverity"] = o.OverrideSeverity
	}
	if !isNil(o.AlertLevel) {
		toSerialize["alertLevel"] = o.AlertLevel
	}
	if !isNil(o.UpdateInterval) {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if !isNil(o.SamplesPerUpdateInterval) {
		toSerialize["samplesPerUpdateInterval"] = o.SamplesPerUpdateInterval
	}
	if !isNil(o.IncludeResource) {
		toSerialize["includeResource"] = o.IncludeResource
	}
	if !isNil(o.ExcludeResource) {
		toSerialize["excludeResource"] = o.ExcludeResource
	}
	if !isNil(o.ServerUnavailableSeverityLevel) {
		toSerialize["serverUnavailableSeverityLevel"] = o.ServerUnavailableSeverityLevel
	}
	if !isNil(o.ServerDegradedSeverityLevel) {
		toSerialize["serverDegradedSeverityLevel"] = o.ServerDegradedSeverityLevel
	}
	return json.Marshal(toSerialize)
}

type NullableAddNumericGaugeRequest struct {
	value *AddNumericGaugeRequest
	isSet bool
}

func (v NullableAddNumericGaugeRequest) Get() *AddNumericGaugeRequest {
	return v.value
}

func (v *NullableAddNumericGaugeRequest) Set(val *AddNumericGaugeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNumericGaugeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNumericGaugeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNumericGaugeRequest(val *AddNumericGaugeRequest) *NullableAddNumericGaugeRequest {
	return &NullableAddNumericGaugeRequest{value: val, isSet: true}
}

func (v NullableAddNumericGaugeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNumericGaugeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}