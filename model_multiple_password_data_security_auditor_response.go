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

// MultiplePasswordDataSecurityAuditorResponse struct for MultiplePasswordDataSecurityAuditorResponse
type MultiplePasswordDataSecurityAuditorResponse struct {
	// Name of the Data Security Auditor
	Id string `json:"id"`
	Schemas []EnummultiplePasswordDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend []string `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewMultiplePasswordDataSecurityAuditorResponse instantiates a new MultiplePasswordDataSecurityAuditorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiplePasswordDataSecurityAuditorResponse(id string, schemas []EnummultiplePasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool) *MultiplePasswordDataSecurityAuditorResponse {
	this := MultiplePasswordDataSecurityAuditorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Enabled = enabled
	return &this
}

// NewMultiplePasswordDataSecurityAuditorResponseWithDefaults instantiates a new MultiplePasswordDataSecurityAuditorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiplePasswordDataSecurityAuditorResponseWithDefaults() *MultiplePasswordDataSecurityAuditorResponse {
	this := MultiplePasswordDataSecurityAuditorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MultiplePasswordDataSecurityAuditorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MultiplePasswordDataSecurityAuditorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *MultiplePasswordDataSecurityAuditorResponse) GetSchemas() []EnummultiplePasswordDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnummultiplePasswordDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetSchemasOk() ([]EnummultiplePasswordDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *MultiplePasswordDataSecurityAuditorResponse) SetSchemas(v []EnummultiplePasswordDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *MultiplePasswordDataSecurityAuditorResponse) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetReportFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *MultiplePasswordDataSecurityAuditorResponse) SetReportFile(v string) {
	o.ReportFile = v
}

// GetEnabled returns the Enabled field value
func (o *MultiplePasswordDataSecurityAuditorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MultiplePasswordDataSecurityAuditorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
    return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *MultiplePasswordDataSecurityAuditorResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditBackend() []string {
	if o == nil || isNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditBackendOk() ([]string, bool) {
	if o == nil || isNil(o.AuditBackend) {
    return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) HasAuditBackend() bool {
	if o != nil && !isNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *MultiplePasswordDataSecurityAuditorResponse) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || isNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || isNil(o.AuditSeverity) {
    return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) HasAuditSeverity() bool {
	if o != nil && !isNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *MultiplePasswordDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MultiplePasswordDataSecurityAuditorResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *MultiplePasswordDataSecurityAuditorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o MultiplePasswordDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["reportFile"] = o.ReportFile
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !isNil(o.AuditBackend) {
		toSerialize["auditBackend"] = o.AuditBackend
	}
	if !isNil(o.AuditSeverity) {
		toSerialize["auditSeverity"] = o.AuditSeverity
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableMultiplePasswordDataSecurityAuditorResponse struct {
	value *MultiplePasswordDataSecurityAuditorResponse
	isSet bool
}

func (v NullableMultiplePasswordDataSecurityAuditorResponse) Get() *MultiplePasswordDataSecurityAuditorResponse {
	return v.value
}

func (v *NullableMultiplePasswordDataSecurityAuditorResponse) Set(val *MultiplePasswordDataSecurityAuditorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiplePasswordDataSecurityAuditorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiplePasswordDataSecurityAuditorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiplePasswordDataSecurityAuditorResponse(val *MultiplePasswordDataSecurityAuditorResponse) *NullableMultiplePasswordDataSecurityAuditorResponse {
	return &NullableMultiplePasswordDataSecurityAuditorResponse{value: val, isSet: true}
}

func (v NullableMultiplePasswordDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiplePasswordDataSecurityAuditorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


