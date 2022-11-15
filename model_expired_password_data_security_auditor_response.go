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

// ExpiredPasswordDataSecurityAuditorResponse struct for ExpiredPasswordDataSecurityAuditorResponse
type ExpiredPasswordDataSecurityAuditorResponse struct {
	// Name of the Data Security Auditor
	Id string `json:"id"`
	Schemas []EnumexpiredPasswordDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled.
	PasswordEvaluationAge *string `json:"passwordEvaluationAge,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled bool `json:"enabled"`
	AuditBackend []string `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewExpiredPasswordDataSecurityAuditorResponse instantiates a new ExpiredPasswordDataSecurityAuditorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiredPasswordDataSecurityAuditorResponse(id string, schemas []EnumexpiredPasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool) *ExpiredPasswordDataSecurityAuditorResponse {
	this := ExpiredPasswordDataSecurityAuditorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Enabled = enabled
	return &this
}

// NewExpiredPasswordDataSecurityAuditorResponseWithDefaults instantiates a new ExpiredPasswordDataSecurityAuditorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiredPasswordDataSecurityAuditorResponseWithDefaults() *ExpiredPasswordDataSecurityAuditorResponse {
	this := ExpiredPasswordDataSecurityAuditorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetSchemas() []EnumexpiredPasswordDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumexpiredPasswordDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetSchemasOk() ([]EnumexpiredPasswordDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetSchemas(v []EnumexpiredPasswordDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetReportFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetReportFile(v string) {
	o.ReportFile = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
    return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetPasswordEvaluationAge returns the PasswordEvaluationAge field value if set, zero value otherwise.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetPasswordEvaluationAge() string {
	if o == nil || isNil(o.PasswordEvaluationAge) {
		var ret string
		return ret
	}
	return *o.PasswordEvaluationAge
}

// GetPasswordEvaluationAgeOk returns a tuple with the PasswordEvaluationAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetPasswordEvaluationAgeOk() (*string, bool) {
	if o == nil || isNil(o.PasswordEvaluationAge) {
    return nil, false
	}
	return o.PasswordEvaluationAge, true
}

// HasPasswordEvaluationAge returns a boolean if a field has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) HasPasswordEvaluationAge() bool {
	if o != nil && !isNil(o.PasswordEvaluationAge) {
		return true
	}

	return false
}

// SetPasswordEvaluationAge gets a reference to the given string and assigns it to the PasswordEvaluationAge field.
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetPasswordEvaluationAge(v string) {
	o.PasswordEvaluationAge = &v
}

// GetEnabled returns the Enabled field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditBackend() []string {
	if o == nil || isNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditBackendOk() ([]string, bool) {
	if o == nil || isNil(o.AuditBackend) {
    return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) HasAuditBackend() bool {
	if o != nil && !isNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || isNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || isNil(o.AuditSeverity) {
    return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *ExpiredPasswordDataSecurityAuditorResponse) HasAuditSeverity() bool {
	if o != nil && !isNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *ExpiredPasswordDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o ExpiredPasswordDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !isNil(o.PasswordEvaluationAge) {
		toSerialize["passwordEvaluationAge"] = o.PasswordEvaluationAge
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

type NullableExpiredPasswordDataSecurityAuditorResponse struct {
	value *ExpiredPasswordDataSecurityAuditorResponse
	isSet bool
}

func (v NullableExpiredPasswordDataSecurityAuditorResponse) Get() *ExpiredPasswordDataSecurityAuditorResponse {
	return v.value
}

func (v *NullableExpiredPasswordDataSecurityAuditorResponse) Set(val *ExpiredPasswordDataSecurityAuditorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiredPasswordDataSecurityAuditorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiredPasswordDataSecurityAuditorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiredPasswordDataSecurityAuditorResponse(val *ExpiredPasswordDataSecurityAuditorResponse) *NullableExpiredPasswordDataSecurityAuditorResponse {
	return &NullableExpiredPasswordDataSecurityAuditorResponse{value: val, isSet: true}
}

func (v NullableExpiredPasswordDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiredPasswordDataSecurityAuditorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


