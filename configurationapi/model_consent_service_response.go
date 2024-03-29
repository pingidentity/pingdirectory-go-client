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

// checks if the ConsentServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsentServiceResponse{}

// ConsentServiceResponse struct for ConsentServiceResponse
type ConsentServiceResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumconsentServiceSchemaUrn                      `json:"schemas,omitempty"`
	// Indicates whether the Consent Service is enabled.
	Enabled bool `json:"enabled"`
	// The base DN under which consent records are stored.
	BaseDN *string `json:"baseDN,omitempty"`
	// The DN of an internal service account used by the Consent Service to make internal LDAP requests.
	BindDN *string `json:"bindDN,omitempty"`
	// The maximum number of consent resources that may be returned from a search request.
	SearchSizeLimit *int64 `json:"searchSizeLimit,omitempty"`
	// If specified, the Identity Mapper(s) that may be used to map consent record subject and actor values to DNs. This is typically only needed if privileged API clients will be used.
	ConsentRecordIdentityMapper []string `json:"consentRecordIdentityMapper,omitempty"`
	// The set of account DNs that the Consent Service will consider to be privileged.
	ServiceAccountDN []string `json:"serviceAccountDN,omitempty"`
	// The name of a scope that must be present in an access token accepted by the Consent Service for unprivileged clients.
	UnprivilegedConsentScope *string `json:"unprivilegedConsentScope,omitempty"`
	// The name of a scope that must be present in an access token accepted by the Consent Service if the client is to be considered privileged.
	PrivilegedConsentScope *string `json:"privilegedConsentScope,omitempty"`
	// A string or URI that identifies the Consent Service in the context of OAuth2 authorization.
	Audience *string `json:"audience,omitempty"`
}

// NewConsentServiceResponse instantiates a new ConsentServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentServiceResponse(enabled bool) *ConsentServiceResponse {
	this := ConsentServiceResponse{}
	this.Enabled = enabled
	return &this
}

// NewConsentServiceResponseWithDefaults instantiates a new ConsentServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentServiceResponseWithDefaults() *ConsentServiceResponse {
	this := ConsentServiceResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ConsentServiceResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ConsentServiceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetSchemas() []EnumconsentServiceSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumconsentServiceSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetSchemasOk() ([]EnumconsentServiceSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumconsentServiceSchemaUrn and assigns it to the Schemas field.
func (o *ConsentServiceResponse) SetSchemas(v []EnumconsentServiceSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *ConsentServiceResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConsentServiceResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetBaseDN() string {
	if o == nil || IsNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *ConsentServiceResponse) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetBindDN returns the BindDN field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetBindDN() string {
	if o == nil || IsNil(o.BindDN) {
		var ret string
		return ret
	}
	return *o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetBindDNOk() (*string, bool) {
	if o == nil || IsNil(o.BindDN) {
		return nil, false
	}
	return o.BindDN, true
}

// HasBindDN returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasBindDN() bool {
	if o != nil && !IsNil(o.BindDN) {
		return true
	}

	return false
}

// SetBindDN gets a reference to the given string and assigns it to the BindDN field.
func (o *ConsentServiceResponse) SetBindDN(v string) {
	o.BindDN = &v
}

// GetSearchSizeLimit returns the SearchSizeLimit field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetSearchSizeLimit() int64 {
	if o == nil || IsNil(o.SearchSizeLimit) {
		var ret int64
		return ret
	}
	return *o.SearchSizeLimit
}

// GetSearchSizeLimitOk returns a tuple with the SearchSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetSearchSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SearchSizeLimit) {
		return nil, false
	}
	return o.SearchSizeLimit, true
}

// HasSearchSizeLimit returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasSearchSizeLimit() bool {
	if o != nil && !IsNil(o.SearchSizeLimit) {
		return true
	}

	return false
}

// SetSearchSizeLimit gets a reference to the given int64 and assigns it to the SearchSizeLimit field.
func (o *ConsentServiceResponse) SetSearchSizeLimit(v int64) {
	o.SearchSizeLimit = &v
}

// GetConsentRecordIdentityMapper returns the ConsentRecordIdentityMapper field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetConsentRecordIdentityMapper() []string {
	if o == nil || IsNil(o.ConsentRecordIdentityMapper) {
		var ret []string
		return ret
	}
	return o.ConsentRecordIdentityMapper
}

// GetConsentRecordIdentityMapperOk returns a tuple with the ConsentRecordIdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetConsentRecordIdentityMapperOk() ([]string, bool) {
	if o == nil || IsNil(o.ConsentRecordIdentityMapper) {
		return nil, false
	}
	return o.ConsentRecordIdentityMapper, true
}

// HasConsentRecordIdentityMapper returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasConsentRecordIdentityMapper() bool {
	if o != nil && !IsNil(o.ConsentRecordIdentityMapper) {
		return true
	}

	return false
}

// SetConsentRecordIdentityMapper gets a reference to the given []string and assigns it to the ConsentRecordIdentityMapper field.
func (o *ConsentServiceResponse) SetConsentRecordIdentityMapper(v []string) {
	o.ConsentRecordIdentityMapper = v
}

// GetServiceAccountDN returns the ServiceAccountDN field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetServiceAccountDN() []string {
	if o == nil || IsNil(o.ServiceAccountDN) {
		var ret []string
		return ret
	}
	return o.ServiceAccountDN
}

// GetServiceAccountDNOk returns a tuple with the ServiceAccountDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetServiceAccountDNOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceAccountDN) {
		return nil, false
	}
	return o.ServiceAccountDN, true
}

// HasServiceAccountDN returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasServiceAccountDN() bool {
	if o != nil && !IsNil(o.ServiceAccountDN) {
		return true
	}

	return false
}

// SetServiceAccountDN gets a reference to the given []string and assigns it to the ServiceAccountDN field.
func (o *ConsentServiceResponse) SetServiceAccountDN(v []string) {
	o.ServiceAccountDN = v
}

// GetUnprivilegedConsentScope returns the UnprivilegedConsentScope field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetUnprivilegedConsentScope() string {
	if o == nil || IsNil(o.UnprivilegedConsentScope) {
		var ret string
		return ret
	}
	return *o.UnprivilegedConsentScope
}

// GetUnprivilegedConsentScopeOk returns a tuple with the UnprivilegedConsentScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetUnprivilegedConsentScopeOk() (*string, bool) {
	if o == nil || IsNil(o.UnprivilegedConsentScope) {
		return nil, false
	}
	return o.UnprivilegedConsentScope, true
}

// HasUnprivilegedConsentScope returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasUnprivilegedConsentScope() bool {
	if o != nil && !IsNil(o.UnprivilegedConsentScope) {
		return true
	}

	return false
}

// SetUnprivilegedConsentScope gets a reference to the given string and assigns it to the UnprivilegedConsentScope field.
func (o *ConsentServiceResponse) SetUnprivilegedConsentScope(v string) {
	o.UnprivilegedConsentScope = &v
}

// GetPrivilegedConsentScope returns the PrivilegedConsentScope field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetPrivilegedConsentScope() string {
	if o == nil || IsNil(o.PrivilegedConsentScope) {
		var ret string
		return ret
	}
	return *o.PrivilegedConsentScope
}

// GetPrivilegedConsentScopeOk returns a tuple with the PrivilegedConsentScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetPrivilegedConsentScopeOk() (*string, bool) {
	if o == nil || IsNil(o.PrivilegedConsentScope) {
		return nil, false
	}
	return o.PrivilegedConsentScope, true
}

// HasPrivilegedConsentScope returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasPrivilegedConsentScope() bool {
	if o != nil && !IsNil(o.PrivilegedConsentScope) {
		return true
	}

	return false
}

// SetPrivilegedConsentScope gets a reference to the given string and assigns it to the PrivilegedConsentScope field.
func (o *ConsentServiceResponse) SetPrivilegedConsentScope(v string) {
	o.PrivilegedConsentScope = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *ConsentServiceResponse) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentServiceResponse) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *ConsentServiceResponse) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *ConsentServiceResponse) SetAudience(v string) {
	o.Audience = &v
}

func (o ConsentServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsentServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !IsNil(o.BindDN) {
		toSerialize["bindDN"] = o.BindDN
	}
	if !IsNil(o.SearchSizeLimit) {
		toSerialize["searchSizeLimit"] = o.SearchSizeLimit
	}
	if !IsNil(o.ConsentRecordIdentityMapper) {
		toSerialize["consentRecordIdentityMapper"] = o.ConsentRecordIdentityMapper
	}
	if !IsNil(o.ServiceAccountDN) {
		toSerialize["serviceAccountDN"] = o.ServiceAccountDN
	}
	if !IsNil(o.UnprivilegedConsentScope) {
		toSerialize["unprivilegedConsentScope"] = o.UnprivilegedConsentScope
	}
	if !IsNil(o.PrivilegedConsentScope) {
		toSerialize["privilegedConsentScope"] = o.PrivilegedConsentScope
	}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	return toSerialize, nil
}

type NullableConsentServiceResponse struct {
	value *ConsentServiceResponse
	isSet bool
}

func (v NullableConsentServiceResponse) Get() *ConsentServiceResponse {
	return v.value
}

func (v *NullableConsentServiceResponse) Set(val *ConsentServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentServiceResponse(val *ConsentServiceResponse) *NullableConsentServiceResponse {
	return &NullableConsentServiceResponse{value: val, isSet: true}
}

func (v NullableConsentServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
