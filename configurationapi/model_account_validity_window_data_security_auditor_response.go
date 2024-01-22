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

// checks if the AccountValidityWindowDataSecurityAuditorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountValidityWindowDataSecurityAuditorResponse{}

// AccountValidityWindowDataSecurityAuditorResponse struct for AccountValidityWindowDataSecurityAuditorResponse
type AccountValidityWindowDataSecurityAuditorResponse struct {
	Schemas []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile string `json:"reportFile"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// If set, the auditor will report all users with account expiration times are in the future, but are within the specified length of time away from the current time.
	AccountExpirationWarningInterval *string `json:"accountExpirationWarningInterval,omitempty"`
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

// NewAccountValidityWindowDataSecurityAuditorResponse instantiates a new AccountValidityWindowDataSecurityAuditorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountValidityWindowDataSecurityAuditorResponse(schemas []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, id string) *AccountValidityWindowDataSecurityAuditorResponse {
	this := AccountValidityWindowDataSecurityAuditorResponse{}
	this.Schemas = schemas
	this.ReportFile = reportFile
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewAccountValidityWindowDataSecurityAuditorResponseWithDefaults instantiates a new AccountValidityWindowDataSecurityAuditorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountValidityWindowDataSecurityAuditorResponseWithDefaults() *AccountValidityWindowDataSecurityAuditorResponse {
	this := AccountValidityWindowDataSecurityAuditorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetSchemas() []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetSchemasOk() ([]EnumaccountValidityWindowDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetSchemas(v []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetReportFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetReportFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFile, true
}

// SetReportFile sets field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetReportFile(v string) {
	o.ReportFile = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAccountExpirationWarningInterval returns the AccountExpirationWarningInterval field value if set, zero value otherwise.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetAccountExpirationWarningInterval() string {
	if o == nil || IsNil(o.AccountExpirationWarningInterval) {
		var ret string
		return ret
	}
	return *o.AccountExpirationWarningInterval
}

// GetAccountExpirationWarningIntervalOk returns a tuple with the AccountExpirationWarningInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetAccountExpirationWarningIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.AccountExpirationWarningInterval) {
		return nil, false
	}
	return o.AccountExpirationWarningInterval, true
}

// HasAccountExpirationWarningInterval returns a boolean if a field has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) HasAccountExpirationWarningInterval() bool {
	if o != nil && !IsNil(o.AccountExpirationWarningInterval) {
		return true
	}

	return false
}

// SetAccountExpirationWarningInterval gets a reference to the given string and assigns it to the AccountExpirationWarningInterval field.
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetAccountExpirationWarningInterval(v string) {
	o.AccountExpirationWarningInterval = &v
}

// GetEnabled returns the Enabled field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountValidityWindowDataSecurityAuditorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountValidityWindowDataSecurityAuditorResponse) SetId(v string) {
	o.Id = v
}

func (o AccountValidityWindowDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountValidityWindowDataSecurityAuditorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["reportFile"] = o.ReportFile
	if !IsNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !IsNil(o.AccountExpirationWarningInterval) {
		toSerialize["accountExpirationWarningInterval"] = o.AccountExpirationWarningInterval
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

type NullableAccountValidityWindowDataSecurityAuditorResponse struct {
	value *AccountValidityWindowDataSecurityAuditorResponse
	isSet bool
}

func (v NullableAccountValidityWindowDataSecurityAuditorResponse) Get() *AccountValidityWindowDataSecurityAuditorResponse {
	return v.value
}

func (v *NullableAccountValidityWindowDataSecurityAuditorResponse) Set(val *AccountValidityWindowDataSecurityAuditorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountValidityWindowDataSecurityAuditorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountValidityWindowDataSecurityAuditorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountValidityWindowDataSecurityAuditorResponse(val *AccountValidityWindowDataSecurityAuditorResponse) *NullableAccountValidityWindowDataSecurityAuditorResponse {
	return &NullableAccountValidityWindowDataSecurityAuditorResponse{value: val, isSet: true}
}

func (v NullableAccountValidityWindowDataSecurityAuditorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountValidityWindowDataSecurityAuditorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
