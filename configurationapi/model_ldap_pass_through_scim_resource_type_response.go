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

// checks if the LdapPassThroughScimResourceTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapPassThroughScimResourceTypeResponse{}

// LdapPassThroughScimResourceTypeResponse struct for LdapPassThroughScimResourceTypeResponse
type LdapPassThroughScimResourceTypeResponse struct {
	// Name of the SCIM Resource Type
	Id      string                                         `json:"id"`
	Schemas []EnumldapPassThroughScimResourceTypeSchemaUrn `json:"schemas"`
	// A description for this SCIM Resource Type
	Description *string `json:"description,omitempty"`
	// Indicates whether the SCIM Resource Type is enabled.
	Enabled bool `json:"enabled"`
	// The HTTP addressable endpoint of this SCIM Resource Type relative to the '/scim/v2' base URL. Do not include a leading '/'.
	Endpoint string `json:"endpoint"`
	// The maximum number of resources that the SCIM Resource Type should \"look through\" in the course of processing a search request.
	LookthroughLimit     *int64                                         `json:"lookthroughLimit,omitempty"`
	SchemaCheckingOption []EnumscimResourceTypeSchemaCheckingOptionProp `json:"schemaCheckingOption,omitempty"`
	// Specifies the LDAP structural object class that should be exposed by this SCIM Resource Type.
	StructuralLDAPObjectclass *string `json:"structuralLDAPObjectclass,omitempty"`
	// Specifies an auxiliary LDAP object class that should be exposed by this SCIM Resource Type.
	AuxiliaryLDAPObjectclass []string `json:"auxiliaryLDAPObjectclass,omitempty"`
	// Specifies the base DN of the branch of the LDAP directory that can be accessed by this SCIM Resource Type.
	IncludeBaseDN *string `json:"includeBaseDN,omitempty"`
	// The set of LDAP filters that define the LDAP entries that should be included in this SCIM Resource Type.
	IncludeFilter []string `json:"includeFilter,omitempty"`
	// Specifies the set of operational LDAP attributes to be provided by this SCIM Resource Type.
	IncludeOperationalAttribute []string `json:"includeOperationalAttribute,omitempty"`
	// Specifies the template to use for the DN when creating new entries.
	CreateDNPattern                               *string                                            `json:"createDNPattern,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewLdapPassThroughScimResourceTypeResponse instantiates a new LdapPassThroughScimResourceTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapPassThroughScimResourceTypeResponse(id string, schemas []EnumldapPassThroughScimResourceTypeSchemaUrn, enabled bool, endpoint string) *LdapPassThroughScimResourceTypeResponse {
	this := LdapPassThroughScimResourceTypeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	this.Endpoint = endpoint
	return &this
}

// NewLdapPassThroughScimResourceTypeResponseWithDefaults instantiates a new LdapPassThroughScimResourceTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapPassThroughScimResourceTypeResponseWithDefaults() *LdapPassThroughScimResourceTypeResponse {
	this := LdapPassThroughScimResourceTypeResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LdapPassThroughScimResourceTypeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LdapPassThroughScimResourceTypeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *LdapPassThroughScimResourceTypeResponse) GetSchemas() []EnumldapPassThroughScimResourceTypeSchemaUrn {
	if o == nil {
		var ret []EnumldapPassThroughScimResourceTypeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetSchemasOk() ([]EnumldapPassThroughScimResourceTypeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LdapPassThroughScimResourceTypeResponse) SetSchemas(v []EnumldapPassThroughScimResourceTypeSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LdapPassThroughScimResourceTypeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *LdapPassThroughScimResourceTypeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LdapPassThroughScimResourceTypeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEndpoint returns the Endpoint field value
func (o *LdapPassThroughScimResourceTypeResponse) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *LdapPassThroughScimResourceTypeResponse) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetLookthroughLimit returns the LookthroughLimit field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetLookthroughLimit() int64 {
	if o == nil || IsNil(o.LookthroughLimit) {
		var ret int64
		return ret
	}
	return *o.LookthroughLimit
}

// GetLookthroughLimitOk returns a tuple with the LookthroughLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetLookthroughLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.LookthroughLimit) {
		return nil, false
	}
	return o.LookthroughLimit, true
}

// HasLookthroughLimit returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasLookthroughLimit() bool {
	if o != nil && !IsNil(o.LookthroughLimit) {
		return true
	}

	return false
}

// SetLookthroughLimit gets a reference to the given int64 and assigns it to the LookthroughLimit field.
func (o *LdapPassThroughScimResourceTypeResponse) SetLookthroughLimit(v int64) {
	o.LookthroughLimit = &v
}

// GetSchemaCheckingOption returns the SchemaCheckingOption field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp {
	if o == nil || IsNil(o.SchemaCheckingOption) {
		var ret []EnumscimResourceTypeSchemaCheckingOptionProp
		return ret
	}
	return o.SchemaCheckingOption
}

// GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetSchemaCheckingOptionOk() ([]EnumscimResourceTypeSchemaCheckingOptionProp, bool) {
	if o == nil || IsNil(o.SchemaCheckingOption) {
		return nil, false
	}
	return o.SchemaCheckingOption, true
}

// HasSchemaCheckingOption returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasSchemaCheckingOption() bool {
	if o != nil && !IsNil(o.SchemaCheckingOption) {
		return true
	}

	return false
}

// SetSchemaCheckingOption gets a reference to the given []EnumscimResourceTypeSchemaCheckingOptionProp and assigns it to the SchemaCheckingOption field.
func (o *LdapPassThroughScimResourceTypeResponse) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp) {
	o.SchemaCheckingOption = v
}

// GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetStructuralLDAPObjectclass() string {
	if o == nil || IsNil(o.StructuralLDAPObjectclass) {
		var ret string
		return ret
	}
	return *o.StructuralLDAPObjectclass
}

// GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetStructuralLDAPObjectclassOk() (*string, bool) {
	if o == nil || IsNil(o.StructuralLDAPObjectclass) {
		return nil, false
	}
	return o.StructuralLDAPObjectclass, true
}

// HasStructuralLDAPObjectclass returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasStructuralLDAPObjectclass() bool {
	if o != nil && !IsNil(o.StructuralLDAPObjectclass) {
		return true
	}

	return false
}

// SetStructuralLDAPObjectclass gets a reference to the given string and assigns it to the StructuralLDAPObjectclass field.
func (o *LdapPassThroughScimResourceTypeResponse) SetStructuralLDAPObjectclass(v string) {
	o.StructuralLDAPObjectclass = &v
}

// GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetAuxiliaryLDAPObjectclass() []string {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !IsNil(o.AuxiliaryLDAPObjectclass) {
		return true
	}

	return false
}

// SetAuxiliaryLDAPObjectclass gets a reference to the given []string and assigns it to the AuxiliaryLDAPObjectclass field.
func (o *LdapPassThroughScimResourceTypeResponse) SetAuxiliaryLDAPObjectclass(v []string) {
	o.AuxiliaryLDAPObjectclass = v
}

// GetIncludeBaseDN returns the IncludeBaseDN field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeBaseDN() string {
	if o == nil || IsNil(o.IncludeBaseDN) {
		var ret string
		return ret
	}
	return *o.IncludeBaseDN
}

// GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.IncludeBaseDN) {
		return nil, false
	}
	return o.IncludeBaseDN, true
}

// HasIncludeBaseDN returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasIncludeBaseDN() bool {
	if o != nil && !IsNil(o.IncludeBaseDN) {
		return true
	}

	return false
}

// SetIncludeBaseDN gets a reference to the given string and assigns it to the IncludeBaseDN field.
func (o *LdapPassThroughScimResourceTypeResponse) SetIncludeBaseDN(v string) {
	o.IncludeBaseDN = &v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *LdapPassThroughScimResourceTypeResponse) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeOperationalAttribute() []string {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeOperationalAttribute
}

// GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeOperationalAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		return nil, false
	}
	return o.IncludeOperationalAttribute, true
}

// HasIncludeOperationalAttribute returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasIncludeOperationalAttribute() bool {
	if o != nil && !IsNil(o.IncludeOperationalAttribute) {
		return true
	}

	return false
}

// SetIncludeOperationalAttribute gets a reference to the given []string and assigns it to the IncludeOperationalAttribute field.
func (o *LdapPassThroughScimResourceTypeResponse) SetIncludeOperationalAttribute(v []string) {
	o.IncludeOperationalAttribute = v
}

// GetCreateDNPattern returns the CreateDNPattern field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetCreateDNPattern() string {
	if o == nil || IsNil(o.CreateDNPattern) {
		var ret string
		return ret
	}
	return *o.CreateDNPattern
}

// GetCreateDNPatternOk returns a tuple with the CreateDNPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetCreateDNPatternOk() (*string, bool) {
	if o == nil || IsNil(o.CreateDNPattern) {
		return nil, false
	}
	return o.CreateDNPattern, true
}

// HasCreateDNPattern returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasCreateDNPattern() bool {
	if o != nil && !IsNil(o.CreateDNPattern) {
		return true
	}

	return false
}

// SetCreateDNPattern gets a reference to the given string and assigns it to the CreateDNPattern field.
func (o *LdapPassThroughScimResourceTypeResponse) SetCreateDNPattern(v string) {
	o.CreateDNPattern = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LdapPassThroughScimResourceTypeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LdapPassThroughScimResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapPassThroughScimResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LdapPassThroughScimResourceTypeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LdapPassThroughScimResourceTypeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o LdapPassThroughScimResourceTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapPassThroughScimResourceTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["endpoint"] = o.Endpoint
	if !IsNil(o.LookthroughLimit) {
		toSerialize["lookthroughLimit"] = o.LookthroughLimit
	}
	if !IsNil(o.SchemaCheckingOption) {
		toSerialize["schemaCheckingOption"] = o.SchemaCheckingOption
	}
	if !IsNil(o.StructuralLDAPObjectclass) {
		toSerialize["structuralLDAPObjectclass"] = o.StructuralLDAPObjectclass
	}
	if !IsNil(o.AuxiliaryLDAPObjectclass) {
		toSerialize["auxiliaryLDAPObjectclass"] = o.AuxiliaryLDAPObjectclass
	}
	if !IsNil(o.IncludeBaseDN) {
		toSerialize["includeBaseDN"] = o.IncludeBaseDN
	}
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !IsNil(o.IncludeOperationalAttribute) {
		toSerialize["includeOperationalAttribute"] = o.IncludeOperationalAttribute
	}
	if !IsNil(o.CreateDNPattern) {
		toSerialize["createDNPattern"] = o.CreateDNPattern
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableLdapPassThroughScimResourceTypeResponse struct {
	value *LdapPassThroughScimResourceTypeResponse
	isSet bool
}

func (v NullableLdapPassThroughScimResourceTypeResponse) Get() *LdapPassThroughScimResourceTypeResponse {
	return v.value
}

func (v *NullableLdapPassThroughScimResourceTypeResponse) Set(val *LdapPassThroughScimResourceTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapPassThroughScimResourceTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapPassThroughScimResourceTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapPassThroughScimResourceTypeResponse(val *LdapPassThroughScimResourceTypeResponse) *NullableLdapPassThroughScimResourceTypeResponse {
	return &NullableLdapPassThroughScimResourceTypeResponse{value: val, isSet: true}
}

func (v NullableLdapPassThroughScimResourceTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapPassThroughScimResourceTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
