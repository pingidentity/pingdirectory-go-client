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

// AddPrivilegeDataSecurityAuditorRequest struct for AddPrivilegeDataSecurityAuditorRequest
type AddPrivilegeDataSecurityAuditorRequest struct {
	// Name of the new Data Security Auditor
	AuditorName string                                      `json:"auditorName"`
	Schemas     []EnumprivilegeDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile       *string                                       `json:"reportFile,omitempty"`
	IncludePrivilege []EnumdataSecurityAuditorIncludePrivilegeProp `json:"includePrivilege,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend  []string                                  `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewAddPrivilegeDataSecurityAuditorRequest instantiates a new AddPrivilegeDataSecurityAuditorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPrivilegeDataSecurityAuditorRequest(auditorName string, schemas []EnumprivilegeDataSecurityAuditorSchemaUrn) *AddPrivilegeDataSecurityAuditorRequest {
	this := AddPrivilegeDataSecurityAuditorRequest{}
	this.AuditorName = auditorName
	this.Schemas = schemas
	return &this
}

// NewAddPrivilegeDataSecurityAuditorRequestWithDefaults instantiates a new AddPrivilegeDataSecurityAuditorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPrivilegeDataSecurityAuditorRequestWithDefaults() *AddPrivilegeDataSecurityAuditorRequest {
	this := AddPrivilegeDataSecurityAuditorRequest{}
	return &this
}

// GetAuditorName returns the AuditorName field value
func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditorName
}

// GetAuditorNameOk returns a tuple with the AuditorName field value
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditorName, true
}

// SetAuditorName sets field value
func (o *AddPrivilegeDataSecurityAuditorRequest) SetAuditorName(v string) {
	o.AuditorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPrivilegeDataSecurityAuditorRequest) GetSchemas() []EnumprivilegeDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumprivilegeDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetSchemasOk() ([]EnumprivilegeDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPrivilegeDataSecurityAuditorRequest) SetSchemas(v []EnumprivilegeDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value if set, zero value otherwise.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetReportFile() string {
	if o == nil || isNil(o.ReportFile) {
		var ret string
		return ret
	}
	return *o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetReportFileOk() (*string, bool) {
	if o == nil || isNil(o.ReportFile) {
		return nil, false
	}
	return o.ReportFile, true
}

// HasReportFile returns a boolean if a field has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) HasReportFile() bool {
	if o != nil && !isNil(o.ReportFile) {
		return true
	}

	return false
}

// SetReportFile gets a reference to the given string and assigns it to the ReportFile field.
func (o *AddPrivilegeDataSecurityAuditorRequest) SetReportFile(v string) {
	o.ReportFile = &v
}

// GetIncludePrivilege returns the IncludePrivilege field value if set, zero value otherwise.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludePrivilege() []EnumdataSecurityAuditorIncludePrivilegeProp {
	if o == nil || isNil(o.IncludePrivilege) {
		var ret []EnumdataSecurityAuditorIncludePrivilegeProp
		return ret
	}
	return o.IncludePrivilege
}

// GetIncludePrivilegeOk returns a tuple with the IncludePrivilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludePrivilegeOk() ([]EnumdataSecurityAuditorIncludePrivilegeProp, bool) {
	if o == nil || isNil(o.IncludePrivilege) {
		return nil, false
	}
	return o.IncludePrivilege, true
}

// HasIncludePrivilege returns a boolean if a field has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) HasIncludePrivilege() bool {
	if o != nil && !isNil(o.IncludePrivilege) {
		return true
	}

	return false
}

// SetIncludePrivilege gets a reference to the given []EnumdataSecurityAuditorIncludePrivilegeProp and assigns it to the IncludePrivilege field.
func (o *AddPrivilegeDataSecurityAuditorRequest) SetIncludePrivilege(v []EnumdataSecurityAuditorIncludePrivilegeProp) {
	o.IncludePrivilege = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddPrivilegeDataSecurityAuditorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AddPrivilegeDataSecurityAuditorRequest) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditBackend() []string {
	if o == nil || isNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditBackendOk() ([]string, bool) {
	if o == nil || isNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) HasAuditBackend() bool {
	if o != nil && !isNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *AddPrivilegeDataSecurityAuditorRequest) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || isNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || isNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *AddPrivilegeDataSecurityAuditorRequest) HasAuditSeverity() bool {
	if o != nil && !isNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *AddPrivilegeDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o AddPrivilegeDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auditorName"] = o.AuditorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.ReportFile) {
		toSerialize["reportFile"] = o.ReportFile
	}
	if !isNil(o.IncludePrivilege) {
		toSerialize["includePrivilege"] = o.IncludePrivilege
	}
	if !isNil(o.Enabled) {
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
	return json.Marshal(toSerialize)
}

type NullableAddPrivilegeDataSecurityAuditorRequest struct {
	value *AddPrivilegeDataSecurityAuditorRequest
	isSet bool
}

func (v NullableAddPrivilegeDataSecurityAuditorRequest) Get() *AddPrivilegeDataSecurityAuditorRequest {
	return v.value
}

func (v *NullableAddPrivilegeDataSecurityAuditorRequest) Set(val *AddPrivilegeDataSecurityAuditorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPrivilegeDataSecurityAuditorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPrivilegeDataSecurityAuditorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPrivilegeDataSecurityAuditorRequest(val *AddPrivilegeDataSecurityAuditorRequest) *NullableAddPrivilegeDataSecurityAuditorRequest {
	return &NullableAddPrivilegeDataSecurityAuditorRequest{value: val, isSet: true}
}

func (v NullableAddPrivilegeDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPrivilegeDataSecurityAuditorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
