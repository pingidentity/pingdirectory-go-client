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

// checks if the LockedAccountDataSecurityAuditorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockedAccountDataSecurityAuditorResponse{}

// LockedAccountDataSecurityAuditorResponse struct for LockedAccountDataSecurityAuditorResponse
type LockedAccountDataSecurityAuditorResponse struct {
	Schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled.
	MaximumIdleTime *string `json:"maximumIdleTime,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend                                  []string                                           `json:"auditBackend,omitempty"`
	AuditSeverity                                 *EnumdataSecurityAuditorAuditSeverityProp          `json:"auditSeverity,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Data Security Auditor
	Id string `json:"id"`
}

// NewLockedAccountDataSecurityAuditorResponse instantiates a new LockedAccountDataSecurityAuditorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockedAccountDataSecurityAuditorResponse(schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, id string) *LockedAccountDataSecurityAuditorResponse {
	this := LockedAccountDataSecurityAuditorResponse{}
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewLockedAccountDataSecurityAuditorResponseWithDefaults instantiates a new LockedAccountDataSecurityAuditorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockedAccountDataSecurityAuditorResponseWithDefaults() *LockedAccountDataSecurityAuditorResponse {
	this := LockedAccountDataSecurityAuditorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *LockedAccountDataSecurityAuditorResponse) GetSchemas() []EnumlockedAccountDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumlockedAccountDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetSchemasOk() ([]EnumlockedAccountDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LockedAccountDataSecurityAuditorResponse) SetSchemas(v []EnumlockedAccountDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *LockedAccountDataSecurityAuditorResponse) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetReportFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *LockedAccountDataSecurityAuditorResponse) SetReportFile(v string) {
	o.ReportFile = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorResponse) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorResponse) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *LockedAccountDataSecurityAuditorResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetMaximumIdleTime returns the MaximumIdleTime field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorResponse) GetMaximumIdleTime() string {
	if o == nil || IsNil(o.MaximumIdleTime) {
		var ret string
		return ret
	}
	return *o.MaximumIdleTime
}

// GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetMaximumIdleTimeOk() (*string, bool) {
	if o == nil || IsNil(o.MaximumIdleTime) {
		return nil, false
	}
	return o.MaximumIdleTime, true
}

// HasMaximumIdleTime returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorResponse) HasMaximumIdleTime() bool {
	if o != nil && !IsNil(o.MaximumIdleTime) {
		return true
	}

	return false
}

// SetMaximumIdleTime gets a reference to the given string and assigns it to the MaximumIdleTime field.
func (o *LockedAccountDataSecurityAuditorResponse) SetMaximumIdleTime(v string) {
	o.MaximumIdleTime = &v
}

// GetEnabled returns the Enabled field value
func (o *LockedAccountDataSecurityAuditorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LockedAccountDataSecurityAuditorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorResponse) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorResponse) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *LockedAccountDataSecurityAuditorResponse) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorResponse) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *LockedAccountDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LockedAccountDataSecurityAuditorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LockedAccountDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LockedAccountDataSecurityAuditorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LockedAccountDataSecurityAuditorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *LockedAccountDataSecurityAuditorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LockedAccountDataSecurityAuditorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LockedAccountDataSecurityAuditorResponse) SetId(v string) {
	o.Id = v
}

func (o LockedAccountDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockedAccountDataSecurityAuditorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["reportFile"] = o.ReportFile
	if !IsNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !IsNil(o.MaximumIdleTime) {
		toSerialize["maximumIdleTime"] = o.MaximumIdleTime
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.AuditBackend) {
		toSerialize["auditBackend"] = o.AuditBackend
	}
	if !IsNil(o.AuditSeverity) {
		toSerialize["auditSeverity"] = o.AuditSeverity
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

type NullableLockedAccountDataSecurityAuditorResponse struct {
	value *LockedAccountDataSecurityAuditorResponse
	isSet bool
}

func (v NullableLockedAccountDataSecurityAuditorResponse) Get() *LockedAccountDataSecurityAuditorResponse {
	return v.value
}

func (v *NullableLockedAccountDataSecurityAuditorResponse) Set(val *LockedAccountDataSecurityAuditorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLockedAccountDataSecurityAuditorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLockedAccountDataSecurityAuditorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockedAccountDataSecurityAuditorResponse(val *LockedAccountDataSecurityAuditorResponse) *NullableLockedAccountDataSecurityAuditorResponse {
	return &NullableLockedAccountDataSecurityAuditorResponse{value: val, isSet: true}
}

func (v NullableLockedAccountDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockedAccountDataSecurityAuditorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
