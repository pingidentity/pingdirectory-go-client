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

// checks if the AddLdapPassThroughScimResourceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLdapPassThroughScimResourceTypeRequest{}

// AddLdapPassThroughScimResourceTypeRequest struct for AddLdapPassThroughScimResourceTypeRequest
type AddLdapPassThroughScimResourceTypeRequest struct {
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
	CreateDNPattern *string `json:"createDNPattern,omitempty"`
	// Name of the new SCIM Resource Type
	TypeName string `json:"typeName"`
}

// NewAddLdapPassThroughScimResourceTypeRequest instantiates a new AddLdapPassThroughScimResourceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLdapPassThroughScimResourceTypeRequest(schemas []EnumldapPassThroughScimResourceTypeSchemaUrn, enabled bool, endpoint string, typeName string) *AddLdapPassThroughScimResourceTypeRequest {
	this := AddLdapPassThroughScimResourceTypeRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.Endpoint = endpoint
	this.TypeName = typeName
	return &this
}

// NewAddLdapPassThroughScimResourceTypeRequestWithDefaults instantiates a new AddLdapPassThroughScimResourceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLdapPassThroughScimResourceTypeRequestWithDefaults() *AddLdapPassThroughScimResourceTypeRequest {
	this := AddLdapPassThroughScimResourceTypeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemas() []EnumldapPassThroughScimResourceTypeSchemaUrn {
	if o == nil {
		var ret []EnumldapPassThroughScimResourceTypeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemasOk() ([]EnumldapPassThroughScimResourceTypeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddLdapPassThroughScimResourceTypeRequest) SetSchemas(v []EnumldapPassThroughScimResourceTypeSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddLdapPassThroughScimResourceTypeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddLdapPassThroughScimResourceTypeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEndpoint returns the Endpoint field value
func (o *AddLdapPassThroughScimResourceTypeRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *AddLdapPassThroughScimResourceTypeRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetLookthroughLimit returns the LookthroughLimit field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetLookthroughLimit() int64 {
	if o == nil || IsNil(o.LookthroughLimit) {
		var ret int64
		return ret
	}
	return *o.LookthroughLimit
}

// GetLookthroughLimitOk returns a tuple with the LookthroughLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetLookthroughLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.LookthroughLimit) {
		return nil, false
	}
	return o.LookthroughLimit, true
}

// HasLookthroughLimit returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasLookthroughLimit() bool {
	if o != nil && !IsNil(o.LookthroughLimit) {
		return true
	}

	return false
}

// SetLookthroughLimit gets a reference to the given int64 and assigns it to the LookthroughLimit field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetLookthroughLimit(v int64) {
	o.LookthroughLimit = &v
}

// GetSchemaCheckingOption returns the SchemaCheckingOption field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp {
	if o == nil || IsNil(o.SchemaCheckingOption) {
		var ret []EnumscimResourceTypeSchemaCheckingOptionProp
		return ret
	}
	return o.SchemaCheckingOption
}

// GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemaCheckingOptionOk() ([]EnumscimResourceTypeSchemaCheckingOptionProp, bool) {
	if o == nil || IsNil(o.SchemaCheckingOption) {
		return nil, false
	}
	return o.SchemaCheckingOption, true
}

// HasSchemaCheckingOption returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasSchemaCheckingOption() bool {
	if o != nil && !IsNil(o.SchemaCheckingOption) {
		return true
	}

	return false
}

// SetSchemaCheckingOption gets a reference to the given []EnumscimResourceTypeSchemaCheckingOptionProp and assigns it to the SchemaCheckingOption field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp) {
	o.SchemaCheckingOption = v
}

// GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetStructuralLDAPObjectclass() string {
	if o == nil || IsNil(o.StructuralLDAPObjectclass) {
		var ret string
		return ret
	}
	return *o.StructuralLDAPObjectclass
}

// GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool) {
	if o == nil || IsNil(o.StructuralLDAPObjectclass) {
		return nil, false
	}
	return o.StructuralLDAPObjectclass, true
}

// HasStructuralLDAPObjectclass returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasStructuralLDAPObjectclass() bool {
	if o != nil && !IsNil(o.StructuralLDAPObjectclass) {
		return true
	}

	return false
}

// SetStructuralLDAPObjectclass gets a reference to the given string and assigns it to the StructuralLDAPObjectclass field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetStructuralLDAPObjectclass(v string) {
	o.StructuralLDAPObjectclass = &v
}

// GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !IsNil(o.AuxiliaryLDAPObjectclass) {
		return true
	}

	return false
}

// SetAuxiliaryLDAPObjectclass gets a reference to the given []string and assigns it to the AuxiliaryLDAPObjectclass field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string) {
	o.AuxiliaryLDAPObjectclass = v
}

// GetIncludeBaseDN returns the IncludeBaseDN field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeBaseDN() string {
	if o == nil || IsNil(o.IncludeBaseDN) {
		var ret string
		return ret
	}
	return *o.IncludeBaseDN
}

// GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.IncludeBaseDN) {
		return nil, false
	}
	return o.IncludeBaseDN, true
}

// HasIncludeBaseDN returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasIncludeBaseDN() bool {
	if o != nil && !IsNil(o.IncludeBaseDN) {
		return true
	}

	return false
}

// SetIncludeBaseDN gets a reference to the given string and assigns it to the IncludeBaseDN field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetIncludeBaseDN(v string) {
	o.IncludeBaseDN = &v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeOperationalAttribute() []string {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeOperationalAttribute
}

// GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeOperationalAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		return nil, false
	}
	return o.IncludeOperationalAttribute, true
}

// HasIncludeOperationalAttribute returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasIncludeOperationalAttribute() bool {
	if o != nil && !IsNil(o.IncludeOperationalAttribute) {
		return true
	}

	return false
}

// SetIncludeOperationalAttribute gets a reference to the given []string and assigns it to the IncludeOperationalAttribute field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetIncludeOperationalAttribute(v []string) {
	o.IncludeOperationalAttribute = v
}

// GetCreateDNPattern returns the CreateDNPattern field value if set, zero value otherwise.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetCreateDNPattern() string {
	if o == nil || IsNil(o.CreateDNPattern) {
		var ret string
		return ret
	}
	return *o.CreateDNPattern
}

// GetCreateDNPatternOk returns a tuple with the CreateDNPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetCreateDNPatternOk() (*string, bool) {
	if o == nil || IsNil(o.CreateDNPattern) {
		return nil, false
	}
	return o.CreateDNPattern, true
}

// HasCreateDNPattern returns a boolean if a field has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) HasCreateDNPattern() bool {
	if o != nil && !IsNil(o.CreateDNPattern) {
		return true
	}

	return false
}

// SetCreateDNPattern gets a reference to the given string and assigns it to the CreateDNPattern field.
func (o *AddLdapPassThroughScimResourceTypeRequest) SetCreateDNPattern(v string) {
	o.CreateDNPattern = &v
}

// GetTypeName returns the TypeName field value
func (o *AddLdapPassThroughScimResourceTypeRequest) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughScimResourceTypeRequest) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *AddLdapPassThroughScimResourceTypeRequest) SetTypeName(v string) {
	o.TypeName = v
}

func (o AddLdapPassThroughScimResourceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLdapPassThroughScimResourceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableAddLdapPassThroughScimResourceTypeRequest struct {
	value *AddLdapPassThroughScimResourceTypeRequest
	isSet bool
}

func (v NullableAddLdapPassThroughScimResourceTypeRequest) Get() *AddLdapPassThroughScimResourceTypeRequest {
	return v.value
}

func (v *NullableAddLdapPassThroughScimResourceTypeRequest) Set(val *AddLdapPassThroughScimResourceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLdapPassThroughScimResourceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLdapPassThroughScimResourceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLdapPassThroughScimResourceTypeRequest(val *AddLdapPassThroughScimResourceTypeRequest) *NullableAddLdapPassThroughScimResourceTypeRequest {
	return &NullableAddLdapPassThroughScimResourceTypeRequest{value: val, isSet: true}
}

func (v NullableAddLdapPassThroughScimResourceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLdapPassThroughScimResourceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
