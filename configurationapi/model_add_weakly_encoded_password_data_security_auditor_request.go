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

// checks if the AddWeaklyEncodedPasswordDataSecurityAuditorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddWeaklyEncodedPasswordDataSecurityAuditorRequest{}

// AddWeaklyEncodedPasswordDataSecurityAuditorRequest struct for AddWeaklyEncodedPasswordDataSecurityAuditorRequest
type AddWeaklyEncodedPasswordDataSecurityAuditorRequest struct {
	// Name of the new Data Security Auditor
	AuditorName string                                                  `json:"auditorName"`
	Schemas     []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile *string `json:"reportFile,omitempty"`
	// The password storage schemes that are considered weak. Users with any of the specified password storage schemes will be included in the report.
	WeakPasswordStorageScheme []string `json:"weakPasswordStorageScheme,omitempty"`
	// Reporting on users with passwords encoded using the Crypt Password Storage scheme may be further limited by selecting one or more encoding mechanisms that are considered weak.
	WeakCryptEncoding []EnumdataSecurityAuditorWeakCryptEncodingProp `json:"weakCryptEncoding,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend  []string                                  `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewAddWeaklyEncodedPasswordDataSecurityAuditorRequest instantiates a new AddWeaklyEncodedPasswordDataSecurityAuditorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddWeaklyEncodedPasswordDataSecurityAuditorRequest(auditorName string, schemas []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn) *AddWeaklyEncodedPasswordDataSecurityAuditorRequest {
	this := AddWeaklyEncodedPasswordDataSecurityAuditorRequest{}
	this.AuditorName = auditorName
	this.Schemas = schemas
	return &this
}

// NewAddWeaklyEncodedPasswordDataSecurityAuditorRequestWithDefaults instantiates a new AddWeaklyEncodedPasswordDataSecurityAuditorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddWeaklyEncodedPasswordDataSecurityAuditorRequestWithDefaults() *AddWeaklyEncodedPasswordDataSecurityAuditorRequest {
	this := AddWeaklyEncodedPasswordDataSecurityAuditorRequest{}
	return &this
}

// GetAuditorName returns the AuditorName field value
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditorName
}

// GetAuditorNameOk returns a tuple with the AuditorName field value
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditorName, true
}

// SetAuditorName sets field value
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetAuditorName(v string) {
	o.AuditorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetSchemas() []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetSchemasOk() ([]EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetSchemas(v []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetReportFile() string {
	if o == nil || IsNil(o.ReportFile) {
		var ret string
		return ret
	}
	return *o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetReportFileOk() (*string, bool) {
	if o == nil || IsNil(o.ReportFile) {
		return nil, false
	}
	return o.ReportFile, true
}

// HasReportFile returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasReportFile() bool {
	if o != nil && !IsNil(o.ReportFile) {
		return true
	}

	return false
}

// SetReportFile gets a reference to the given string and assigns it to the ReportFile field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetReportFile(v string) {
	o.ReportFile = &v
}

// GetWeakPasswordStorageScheme returns the WeakPasswordStorageScheme field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakPasswordStorageScheme() []string {
	if o == nil || IsNil(o.WeakPasswordStorageScheme) {
		var ret []string
		return ret
	}
	return o.WeakPasswordStorageScheme
}

// GetWeakPasswordStorageSchemeOk returns a tuple with the WeakPasswordStorageScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakPasswordStorageSchemeOk() ([]string, bool) {
	if o == nil || IsNil(o.WeakPasswordStorageScheme) {
		return nil, false
	}
	return o.WeakPasswordStorageScheme, true
}

// HasWeakPasswordStorageScheme returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasWeakPasswordStorageScheme() bool {
	if o != nil && !IsNil(o.WeakPasswordStorageScheme) {
		return true
	}

	return false
}

// SetWeakPasswordStorageScheme gets a reference to the given []string and assigns it to the WeakPasswordStorageScheme field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetWeakPasswordStorageScheme(v []string) {
	o.WeakPasswordStorageScheme = v
}

// GetWeakCryptEncoding returns the WeakCryptEncoding field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakCryptEncoding() []EnumdataSecurityAuditorWeakCryptEncodingProp {
	if o == nil || IsNil(o.WeakCryptEncoding) {
		var ret []EnumdataSecurityAuditorWeakCryptEncodingProp
		return ret
	}
	return o.WeakCryptEncoding
}

// GetWeakCryptEncodingOk returns a tuple with the WeakCryptEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakCryptEncodingOk() ([]EnumdataSecurityAuditorWeakCryptEncodingProp, bool) {
	if o == nil || IsNil(o.WeakCryptEncoding) {
		return nil, false
	}
	return o.WeakCryptEncoding, true
}

// HasWeakCryptEncoding returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasWeakCryptEncoding() bool {
	if o != nil && !IsNil(o.WeakCryptEncoding) {
		return true
	}

	return false
}

// SetWeakCryptEncoding gets a reference to the given []EnumdataSecurityAuditorWeakCryptEncodingProp and assigns it to the WeakCryptEncoding field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetWeakCryptEncoding(v []EnumdataSecurityAuditorWeakCryptEncodingProp) {
	o.WeakCryptEncoding = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o AddWeaklyEncodedPasswordDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddWeaklyEncodedPasswordDataSecurityAuditorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auditorName"] = o.AuditorName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.ReportFile) {
		toSerialize["reportFile"] = o.ReportFile
	}
	if !IsNil(o.WeakPasswordStorageScheme) {
		toSerialize["weakPasswordStorageScheme"] = o.WeakPasswordStorageScheme
	}
	if !IsNil(o.WeakCryptEncoding) {
		toSerialize["weakCryptEncoding"] = o.WeakCryptEncoding
	}
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

type NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest struct {
	value *AddWeaklyEncodedPasswordDataSecurityAuditorRequest
	isSet bool
}

func (v NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest) Get() *AddWeaklyEncodedPasswordDataSecurityAuditorRequest {
	return v.value
}

func (v *NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest) Set(val *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest(val *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) *NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest {
	return &NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest{value: val, isSet: true}
}

func (v NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddWeaklyEncodedPasswordDataSecurityAuditorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
