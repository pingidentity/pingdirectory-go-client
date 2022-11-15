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

// LockedAccountDataSecurityAuditorShared struct for LockedAccountDataSecurityAuditorShared
type LockedAccountDataSecurityAuditorShared struct {
	Schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled.
	MaximumIdleTime *string `json:"maximumIdleTime,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled bool `json:"enabled"`
	AuditBackend []string `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewLockedAccountDataSecurityAuditorShared instantiates a new LockedAccountDataSecurityAuditorShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockedAccountDataSecurityAuditorShared(schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn, reportFile string, enabled bool) *LockedAccountDataSecurityAuditorShared {
	this := LockedAccountDataSecurityAuditorShared{}
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Enabled = enabled
	return &this
}

// NewLockedAccountDataSecurityAuditorSharedWithDefaults instantiates a new LockedAccountDataSecurityAuditorShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockedAccountDataSecurityAuditorSharedWithDefaults() *LockedAccountDataSecurityAuditorShared {
	this := LockedAccountDataSecurityAuditorShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *LockedAccountDataSecurityAuditorShared) GetSchemas() []EnumlockedAccountDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumlockedAccountDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetSchemasOk() ([]EnumlockedAccountDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LockedAccountDataSecurityAuditorShared) SetSchemas(v []EnumlockedAccountDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *LockedAccountDataSecurityAuditorShared) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetReportFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *LockedAccountDataSecurityAuditorShared) SetReportFile(v string) {
	o.ReportFile = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorShared) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
    return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorShared) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *LockedAccountDataSecurityAuditorShared) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetMaximumIdleTime returns the MaximumIdleTime field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorShared) GetMaximumIdleTime() string {
	if o == nil || isNil(o.MaximumIdleTime) {
		var ret string
		return ret
	}
	return *o.MaximumIdleTime
}

// GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetMaximumIdleTimeOk() (*string, bool) {
	if o == nil || isNil(o.MaximumIdleTime) {
    return nil, false
	}
	return o.MaximumIdleTime, true
}

// HasMaximumIdleTime returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorShared) HasMaximumIdleTime() bool {
	if o != nil && !isNil(o.MaximumIdleTime) {
		return true
	}

	return false
}

// SetMaximumIdleTime gets a reference to the given string and assigns it to the MaximumIdleTime field.
func (o *LockedAccountDataSecurityAuditorShared) SetMaximumIdleTime(v string) {
	o.MaximumIdleTime = &v
}

// GetEnabled returns the Enabled field value
func (o *LockedAccountDataSecurityAuditorShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LockedAccountDataSecurityAuditorShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorShared) GetAuditBackend() []string {
	if o == nil || isNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetAuditBackendOk() ([]string, bool) {
	if o == nil || isNil(o.AuditBackend) {
    return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorShared) HasAuditBackend() bool {
	if o != nil && !isNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *LockedAccountDataSecurityAuditorShared) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || isNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || isNil(o.AuditSeverity) {
    return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorShared) HasAuditSeverity() bool {
	if o != nil && !isNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *LockedAccountDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o LockedAccountDataSecurityAuditorShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["reportFile"] = o.ReportFile
	}
	if !isNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !isNil(o.MaximumIdleTime) {
		toSerialize["maximumIdleTime"] = o.MaximumIdleTime
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AuditBackend) {
		toSerialize["auditBackend"] = o.AuditBackend
	}
	if !isNil(o.AuditSeverity) {
		toSerialize["auditSeverity"] = o.AuditSeverity
	}
	return json.Marshal(toSerialize)
}

type NullableLockedAccountDataSecurityAuditorShared struct {
	value *LockedAccountDataSecurityAuditorShared
	isSet bool
}

func (v NullableLockedAccountDataSecurityAuditorShared) Get() *LockedAccountDataSecurityAuditorShared {
	return v.value
}

func (v *NullableLockedAccountDataSecurityAuditorShared) Set(val *LockedAccountDataSecurityAuditorShared) {
	v.value = val
	v.isSet = true
}

func (v NullableLockedAccountDataSecurityAuditorShared) IsSet() bool {
	return v.isSet
}

func (v *NullableLockedAccountDataSecurityAuditorShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockedAccountDataSecurityAuditorShared(val *LockedAccountDataSecurityAuditorShared) *NullableLockedAccountDataSecurityAuditorShared {
	return &NullableLockedAccountDataSecurityAuditorShared{value: val, isSet: true}
}

func (v NullableLockedAccountDataSecurityAuditorShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockedAccountDataSecurityAuditorShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


