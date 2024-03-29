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

// checks if the WeaklyEncodedPasswordDataSecurityAuditorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeaklyEncodedPasswordDataSecurityAuditorResponse{}

// WeaklyEncodedPasswordDataSecurityAuditorResponse struct for WeaklyEncodedPasswordDataSecurityAuditorResponse
type WeaklyEncodedPasswordDataSecurityAuditorResponse struct {
	Schemas []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	// The password storage schemes that are considered weak. Users with any of the specified password storage schemes will be included in the report.
	WeakPasswordStorageScheme []string                                       `json:"weakPasswordStorageScheme"`
	WeakCryptEncoding         []EnumdataSecurityAuditorWeakCryptEncodingProp `json:"weakCryptEncoding,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend                                  []string                                           `json:"auditBackend,omitempty"`
	AuditSeverity                                 *EnumdataSecurityAuditorAuditSeverityProp          `json:"auditSeverity,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Data Security Auditor
	Id string `json:"id"`
}

// NewWeaklyEncodedPasswordDataSecurityAuditorResponse instantiates a new WeaklyEncodedPasswordDataSecurityAuditorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeaklyEncodedPasswordDataSecurityAuditorResponse(schemas []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, reportFile string, weakPasswordStorageScheme []string, enabled bool, id string) *WeaklyEncodedPasswordDataSecurityAuditorResponse {
	this := WeaklyEncodedPasswordDataSecurityAuditorResponse{}
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.WeakPasswordStorageScheme = weakPasswordStorageScheme
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewWeaklyEncodedPasswordDataSecurityAuditorResponseWithDefaults instantiates a new WeaklyEncodedPasswordDataSecurityAuditorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeaklyEncodedPasswordDataSecurityAuditorResponseWithDefaults() *WeaklyEncodedPasswordDataSecurityAuditorResponse {
	this := WeaklyEncodedPasswordDataSecurityAuditorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetSchemas() []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetSchemasOk() ([]EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetSchemas(v []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetReportFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetReportFile(v string) {
	o.ReportFile = v
}

// GetWeakPasswordStorageScheme returns the WeakPasswordStorageScheme field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetWeakPasswordStorageScheme() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WeakPasswordStorageScheme
}

// GetWeakPasswordStorageSchemeOk returns a tuple with the WeakPasswordStorageScheme field value
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetWeakPasswordStorageSchemeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeakPasswordStorageScheme, true
}

// SetWeakPasswordStorageScheme sets field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetWeakPasswordStorageScheme(v []string) {
	o.WeakPasswordStorageScheme = v
}

// GetWeakCryptEncoding returns the WeakCryptEncoding field value if set, zero value otherwise.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetWeakCryptEncoding() []EnumdataSecurityAuditorWeakCryptEncodingProp {
	if o == nil || IsNil(o.WeakCryptEncoding) {
		var ret []EnumdataSecurityAuditorWeakCryptEncodingProp
		return ret
	}
	return o.WeakCryptEncoding
}

// GetWeakCryptEncodingOk returns a tuple with the WeakCryptEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetWeakCryptEncodingOk() ([]EnumdataSecurityAuditorWeakCryptEncodingProp, bool) {
	if o == nil || IsNil(o.WeakCryptEncoding) {
		return nil, false
	}
	return o.WeakCryptEncoding, true
}

// HasWeakCryptEncoding returns a boolean if a field has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) HasWeakCryptEncoding() bool {
	if o != nil && !IsNil(o.WeakCryptEncoding) {
		return true
	}

	return false
}

// SetWeakCryptEncoding gets a reference to the given []EnumdataSecurityAuditorWeakCryptEncodingProp and assigns it to the WeakCryptEncoding field.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetWeakCryptEncoding(v []EnumdataSecurityAuditorWeakCryptEncodingProp) {
	o.WeakCryptEncoding = v
}

// GetEnabled returns the Enabled field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WeaklyEncodedPasswordDataSecurityAuditorResponse) SetId(v string) {
	o.Id = v
}

func (o WeaklyEncodedPasswordDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeaklyEncodedPasswordDataSecurityAuditorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["reportFile"] = o.ReportFile
	toSerialize["weakPasswordStorageScheme"] = o.WeakPasswordStorageScheme
	if !IsNil(o.WeakCryptEncoding) {
		toSerialize["weakCryptEncoding"] = o.WeakCryptEncoding
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
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

type NullableWeaklyEncodedPasswordDataSecurityAuditorResponse struct {
	value *WeaklyEncodedPasswordDataSecurityAuditorResponse
	isSet bool
}

func (v NullableWeaklyEncodedPasswordDataSecurityAuditorResponse) Get() *WeaklyEncodedPasswordDataSecurityAuditorResponse {
	return v.value
}

func (v *NullableWeaklyEncodedPasswordDataSecurityAuditorResponse) Set(val *WeaklyEncodedPasswordDataSecurityAuditorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWeaklyEncodedPasswordDataSecurityAuditorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWeaklyEncodedPasswordDataSecurityAuditorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeaklyEncodedPasswordDataSecurityAuditorResponse(val *WeaklyEncodedPasswordDataSecurityAuditorResponse) *NullableWeaklyEncodedPasswordDataSecurityAuditorResponse {
	return &NullableWeaklyEncodedPasswordDataSecurityAuditorResponse{value: val, isSet: true}
}

func (v NullableWeaklyEncodedPasswordDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeaklyEncodedPasswordDataSecurityAuditorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
