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

// checks if the AddFilterDataSecurityAuditorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFilterDataSecurityAuditorRequest{}

// AddFilterDataSecurityAuditorRequest struct for AddFilterDataSecurityAuditorRequest
type AddFilterDataSecurityAuditorRequest struct {
	// Name of the new Data Security Auditor
	AuditorName string                                   `json:"auditorName"`
	Schemas     []EnumfilterDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	// The filter to use to identify entries that should be reported. Multiple filters may be configured, and each reported entry will indicate which of these filter(s) matched that entry.
	Filter []string `json:"filter"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend  []string                                  `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewAddFilterDataSecurityAuditorRequest instantiates a new AddFilterDataSecurityAuditorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFilterDataSecurityAuditorRequest(auditorName string, schemas []EnumfilterDataSecurityAuditorSchemaUrn, reportFile string, filter []string) *AddFilterDataSecurityAuditorRequest {
	this := AddFilterDataSecurityAuditorRequest{}
	this.AuditorName = auditorName
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Filter = filter
	return &this
}

// NewAddFilterDataSecurityAuditorRequestWithDefaults instantiates a new AddFilterDataSecurityAuditorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFilterDataSecurityAuditorRequestWithDefaults() *AddFilterDataSecurityAuditorRequest {
	this := AddFilterDataSecurityAuditorRequest{}
	return &this
}

// GetAuditorName returns the AuditorName field value
func (o *AddFilterDataSecurityAuditorRequest) GetAuditorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditorName
}

// GetAuditorNameOk returns a tuple with the AuditorName field value
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditorName, true
}

// SetAuditorName sets field value
func (o *AddFilterDataSecurityAuditorRequest) SetAuditorName(v string) {
	o.AuditorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddFilterDataSecurityAuditorRequest) GetSchemas() []EnumfilterDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumfilterDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetSchemasOk() ([]EnumfilterDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFilterDataSecurityAuditorRequest) SetSchemas(v []EnumfilterDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *AddFilterDataSecurityAuditorRequest) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetReportFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *AddFilterDataSecurityAuditorRequest) SetReportFile(v string) {
	o.ReportFile = v
}

// GetFilter returns the Filter field value
func (o *AddFilterDataSecurityAuditorRequest) GetFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetFilterOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter, true
}

// SetFilter sets field value
func (o *AddFilterDataSecurityAuditorRequest) SetFilter(v []string) {
	o.Filter = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddFilterDataSecurityAuditorRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddFilterDataSecurityAuditorRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddFilterDataSecurityAuditorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AddFilterDataSecurityAuditorRequest) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AddFilterDataSecurityAuditorRequest) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AddFilterDataSecurityAuditorRequest) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *AddFilterDataSecurityAuditorRequest) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *AddFilterDataSecurityAuditorRequest) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *AddFilterDataSecurityAuditorRequest) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *AddFilterDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFilterDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *AddFilterDataSecurityAuditorRequest) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *AddFilterDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o AddFilterDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFilterDataSecurityAuditorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auditorName"] = o.AuditorName
	toSerialize["schemas"] = o.Schemas
	toSerialize["reportFile"] = o.ReportFile
	toSerialize["filter"] = o.Filter
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !IsNil(o.AuditBackend) {
		toSerialize["auditBackend"] = o.AuditBackend
	}
	if !IsNil(o.AuditSeverity) {
		toSerialize["auditSeverity"] = o.AuditSeverity
	}
	return toSerialize, nil
}

type NullableAddFilterDataSecurityAuditorRequest struct {
	value *AddFilterDataSecurityAuditorRequest
	isSet bool
}

func (v NullableAddFilterDataSecurityAuditorRequest) Get() *AddFilterDataSecurityAuditorRequest {
	return v.value
}

func (v *NullableAddFilterDataSecurityAuditorRequest) Set(val *AddFilterDataSecurityAuditorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFilterDataSecurityAuditorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFilterDataSecurityAuditorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFilterDataSecurityAuditorRequest(val *AddFilterDataSecurityAuditorRequest) *NullableAddFilterDataSecurityAuditorRequest {
	return &NullableAddFilterDataSecurityAuditorRequest{value: val, isSet: true}
}

func (v NullableAddFilterDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFilterDataSecurityAuditorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
