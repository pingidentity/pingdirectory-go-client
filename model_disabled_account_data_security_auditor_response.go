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

// DisabledAccountDataSecurityAuditorResponse struct for DisabledAccountDataSecurityAuditorResponse
type DisabledAccountDataSecurityAuditorResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Data Security Auditor
	Id string `json:"id"`
	Schemas []EnumdisabledAccountDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend []string `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewDisabledAccountDataSecurityAuditorResponse instantiates a new DisabledAccountDataSecurityAuditorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisabledAccountDataSecurityAuditorResponse(id string, schemas []EnumdisabledAccountDataSecurityAuditorSchemaUrn, reportFile string, enabled bool) *DisabledAccountDataSecurityAuditorResponse {
	this := DisabledAccountDataSecurityAuditorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Enabled = enabled
	return &this
}

// NewDisabledAccountDataSecurityAuditorResponseWithDefaults instantiates a new DisabledAccountDataSecurityAuditorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisabledAccountDataSecurityAuditorResponseWithDefaults() *DisabledAccountDataSecurityAuditorResponse {
	this := DisabledAccountDataSecurityAuditorResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DisabledAccountDataSecurityAuditorResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DisabledAccountDataSecurityAuditorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DisabledAccountDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DisabledAccountDataSecurityAuditorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *DisabledAccountDataSecurityAuditorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DisabledAccountDataSecurityAuditorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *DisabledAccountDataSecurityAuditorResponse) GetSchemas() []EnumdisabledAccountDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumdisabledAccountDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetSchemasOk() ([]EnumdisabledAccountDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DisabledAccountDataSecurityAuditorResponse) SetSchemas(v []EnumdisabledAccountDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *DisabledAccountDataSecurityAuditorResponse) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetReportFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *DisabledAccountDataSecurityAuditorResponse) SetReportFile(v string) {
	o.ReportFile = v
}

// GetEnabled returns the Enabled field value
func (o *DisabledAccountDataSecurityAuditorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DisabledAccountDataSecurityAuditorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *DisabledAccountDataSecurityAuditorResponse) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
    return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *DisabledAccountDataSecurityAuditorResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditBackend() []string {
	if o == nil || isNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditBackendOk() ([]string, bool) {
	if o == nil || isNil(o.AuditBackend) {
    return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) HasAuditBackend() bool {
	if o != nil && !isNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *DisabledAccountDataSecurityAuditorResponse) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || isNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || isNil(o.AuditSeverity) {
    return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *DisabledAccountDataSecurityAuditorResponse) HasAuditSeverity() bool {
	if o != nil && !isNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *DisabledAccountDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o DisabledAccountDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableDisabledAccountDataSecurityAuditorResponse struct {
	value *DisabledAccountDataSecurityAuditorResponse
	isSet bool
}

func (v NullableDisabledAccountDataSecurityAuditorResponse) Get() *DisabledAccountDataSecurityAuditorResponse {
	return v.value
}

func (v *NullableDisabledAccountDataSecurityAuditorResponse) Set(val *DisabledAccountDataSecurityAuditorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisabledAccountDataSecurityAuditorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisabledAccountDataSecurityAuditorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisabledAccountDataSecurityAuditorResponse(val *DisabledAccountDataSecurityAuditorResponse) *NullableDisabledAccountDataSecurityAuditorResponse {
	return &NullableDisabledAccountDataSecurityAuditorResponse{value: val, isSet: true}
}

func (v NullableDisabledAccountDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisabledAccountDataSecurityAuditorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

