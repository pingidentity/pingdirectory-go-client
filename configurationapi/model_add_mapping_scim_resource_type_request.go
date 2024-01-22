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

// checks if the AddMappingScimResourceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddMappingScimResourceTypeRequest{}

// AddMappingScimResourceTypeRequest struct for AddMappingScimResourceTypeRequest
type AddMappingScimResourceTypeRequest struct {
	Schemas []EnummappingScimResourceTypeSchemaUrn `json:"schemas"`
	// The core schema enforced on core attributes at the top level of a SCIM resource representation exposed by thisMapping SCIM Resource Type.
	CoreSchema string `json:"coreSchema"`
	// Required additive schemas that are enforced on extension attributes in a SCIM resource representation for this Mapping SCIM Resource Type.
	RequiredSchemaExtension []string `json:"requiredSchemaExtension,omitempty"`
	// Optional additive schemas that are enforced on extension attributes in a SCIM resource representation for this Mapping SCIM Resource Type.
	OptionalSchemaExtension []string `json:"optionalSchemaExtension,omitempty"`
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

// NewAddMappingScimResourceTypeRequest instantiates a new AddMappingScimResourceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMappingScimResourceTypeRequest(schemas []EnummappingScimResourceTypeSchemaUrn, coreSchema string, enabled bool, endpoint string, typeName string) *AddMappingScimResourceTypeRequest {
	this := AddMappingScimResourceTypeRequest{}
	this.Schemas = schemas
	this.CoreSchema = coreSchema
	this.Enabled = enabled
	this.Endpoint = endpoint
	this.TypeName = typeName
	return &this
}

// NewAddMappingScimResourceTypeRequestWithDefaults instantiates a new AddMappingScimResourceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMappingScimResourceTypeRequestWithDefaults() *AddMappingScimResourceTypeRequest {
	this := AddMappingScimResourceTypeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddMappingScimResourceTypeRequest) GetSchemas() []EnummappingScimResourceTypeSchemaUrn {
	if o == nil {
		var ret []EnummappingScimResourceTypeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetSchemasOk() ([]EnummappingScimResourceTypeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddMappingScimResourceTypeRequest) SetSchemas(v []EnummappingScimResourceTypeSchemaUrn) {
	o.Schemas = v
}

// GetCoreSchema returns the CoreSchema field value
func (o *AddMappingScimResourceTypeRequest) GetCoreSchema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CoreSchema
}

// GetCoreSchemaOk returns a tuple with the CoreSchema field value
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetCoreSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoreSchema, true
}

// SetCoreSchema sets field value
func (o *AddMappingScimResourceTypeRequest) SetCoreSchema(v string) {
	o.CoreSchema = v
}

// GetRequiredSchemaExtension returns the RequiredSchemaExtension field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetRequiredSchemaExtension() []string {
	if o == nil || IsNil(o.RequiredSchemaExtension) {
		var ret []string
		return ret
	}
	return o.RequiredSchemaExtension
}

// GetRequiredSchemaExtensionOk returns a tuple with the RequiredSchemaExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetRequiredSchemaExtensionOk() ([]string, bool) {
	if o == nil || IsNil(o.RequiredSchemaExtension) {
		return nil, false
	}
	return o.RequiredSchemaExtension, true
}

// HasRequiredSchemaExtension returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasRequiredSchemaExtension() bool {
	if o != nil && !IsNil(o.RequiredSchemaExtension) {
		return true
	}

	return false
}

// SetRequiredSchemaExtension gets a reference to the given []string and assigns it to the RequiredSchemaExtension field.
func (o *AddMappingScimResourceTypeRequest) SetRequiredSchemaExtension(v []string) {
	o.RequiredSchemaExtension = v
}

// GetOptionalSchemaExtension returns the OptionalSchemaExtension field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetOptionalSchemaExtension() []string {
	if o == nil || IsNil(o.OptionalSchemaExtension) {
		var ret []string
		return ret
	}
	return o.OptionalSchemaExtension
}

// GetOptionalSchemaExtensionOk returns a tuple with the OptionalSchemaExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetOptionalSchemaExtensionOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionalSchemaExtension) {
		return nil, false
	}
	return o.OptionalSchemaExtension, true
}

// HasOptionalSchemaExtension returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasOptionalSchemaExtension() bool {
	if o != nil && !IsNil(o.OptionalSchemaExtension) {
		return true
	}

	return false
}

// SetOptionalSchemaExtension gets a reference to the given []string and assigns it to the OptionalSchemaExtension field.
func (o *AddMappingScimResourceTypeRequest) SetOptionalSchemaExtension(v []string) {
	o.OptionalSchemaExtension = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddMappingScimResourceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddMappingScimResourceTypeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddMappingScimResourceTypeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEndpoint returns the Endpoint field value
func (o *AddMappingScimResourceTypeRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *AddMappingScimResourceTypeRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetLookthroughLimit returns the LookthroughLimit field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetLookthroughLimit() int64 {
	if o == nil || IsNil(o.LookthroughLimit) {
		var ret int64
		return ret
	}
	return *o.LookthroughLimit
}

// GetLookthroughLimitOk returns a tuple with the LookthroughLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetLookthroughLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.LookthroughLimit) {
		return nil, false
	}
	return o.LookthroughLimit, true
}

// HasLookthroughLimit returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasLookthroughLimit() bool {
	if o != nil && !IsNil(o.LookthroughLimit) {
		return true
	}

	return false
}

// SetLookthroughLimit gets a reference to the given int64 and assigns it to the LookthroughLimit field.
func (o *AddMappingScimResourceTypeRequest) SetLookthroughLimit(v int64) {
	o.LookthroughLimit = &v
}

// GetSchemaCheckingOption returns the SchemaCheckingOption field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp {
	if o == nil || IsNil(o.SchemaCheckingOption) {
		var ret []EnumscimResourceTypeSchemaCheckingOptionProp
		return ret
	}
	return o.SchemaCheckingOption
}

// GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetSchemaCheckingOptionOk() ([]EnumscimResourceTypeSchemaCheckingOptionProp, bool) {
	if o == nil || IsNil(o.SchemaCheckingOption) {
		return nil, false
	}
	return o.SchemaCheckingOption, true
}

// HasSchemaCheckingOption returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasSchemaCheckingOption() bool {
	if o != nil && !IsNil(o.SchemaCheckingOption) {
		return true
	}

	return false
}

// SetSchemaCheckingOption gets a reference to the given []EnumscimResourceTypeSchemaCheckingOptionProp and assigns it to the SchemaCheckingOption field.
func (o *AddMappingScimResourceTypeRequest) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp) {
	o.SchemaCheckingOption = v
}

// GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetStructuralLDAPObjectclass() string {
	if o == nil || IsNil(o.StructuralLDAPObjectclass) {
		var ret string
		return ret
	}
	return *o.StructuralLDAPObjectclass
}

// GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool) {
	if o == nil || IsNil(o.StructuralLDAPObjectclass) {
		return nil, false
	}
	return o.StructuralLDAPObjectclass, true
}

// HasStructuralLDAPObjectclass returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasStructuralLDAPObjectclass() bool {
	if o != nil && !IsNil(o.StructuralLDAPObjectclass) {
		return true
	}

	return false
}

// SetStructuralLDAPObjectclass gets a reference to the given string and assigns it to the StructuralLDAPObjectclass field.
func (o *AddMappingScimResourceTypeRequest) SetStructuralLDAPObjectclass(v string) {
	o.StructuralLDAPObjectclass = &v
}

// GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !IsNil(o.AuxiliaryLDAPObjectclass) {
		return true
	}

	return false
}

// SetAuxiliaryLDAPObjectclass gets a reference to the given []string and assigns it to the AuxiliaryLDAPObjectclass field.
func (o *AddMappingScimResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string) {
	o.AuxiliaryLDAPObjectclass = v
}

// GetIncludeBaseDN returns the IncludeBaseDN field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetIncludeBaseDN() string {
	if o == nil || IsNil(o.IncludeBaseDN) {
		var ret string
		return ret
	}
	return *o.IncludeBaseDN
}

// GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetIncludeBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.IncludeBaseDN) {
		return nil, false
	}
	return o.IncludeBaseDN, true
}

// HasIncludeBaseDN returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasIncludeBaseDN() bool {
	if o != nil && !IsNil(o.IncludeBaseDN) {
		return true
	}

	return false
}

// SetIncludeBaseDN gets a reference to the given string and assigns it to the IncludeBaseDN field.
func (o *AddMappingScimResourceTypeRequest) SetIncludeBaseDN(v string) {
	o.IncludeBaseDN = &v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *AddMappingScimResourceTypeRequest) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetIncludeOperationalAttribute() []string {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeOperationalAttribute
}

// GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetIncludeOperationalAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		return nil, false
	}
	return o.IncludeOperationalAttribute, true
}

// HasIncludeOperationalAttribute returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasIncludeOperationalAttribute() bool {
	if o != nil && !IsNil(o.IncludeOperationalAttribute) {
		return true
	}

	return false
}

// SetIncludeOperationalAttribute gets a reference to the given []string and assigns it to the IncludeOperationalAttribute field.
func (o *AddMappingScimResourceTypeRequest) SetIncludeOperationalAttribute(v []string) {
	o.IncludeOperationalAttribute = v
}

// GetCreateDNPattern returns the CreateDNPattern field value if set, zero value otherwise.
func (o *AddMappingScimResourceTypeRequest) GetCreateDNPattern() string {
	if o == nil || IsNil(o.CreateDNPattern) {
		var ret string
		return ret
	}
	return *o.CreateDNPattern
}

// GetCreateDNPatternOk returns a tuple with the CreateDNPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetCreateDNPatternOk() (*string, bool) {
	if o == nil || IsNil(o.CreateDNPattern) {
		return nil, false
	}
	return o.CreateDNPattern, true
}

// HasCreateDNPattern returns a boolean if a field has been set.
func (o *AddMappingScimResourceTypeRequest) HasCreateDNPattern() bool {
	if o != nil && !IsNil(o.CreateDNPattern) {
		return true
	}

	return false
}

// SetCreateDNPattern gets a reference to the given string and assigns it to the CreateDNPattern field.
func (o *AddMappingScimResourceTypeRequest) SetCreateDNPattern(v string) {
	o.CreateDNPattern = &v
}

// GetTypeName returns the TypeName field value
func (o *AddMappingScimResourceTypeRequest) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *AddMappingScimResourceTypeRequest) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *AddMappingScimResourceTypeRequest) SetTypeName(v string) {
	o.TypeName = v
}

func (o AddMappingScimResourceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddMappingScimResourceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["coreSchema"] = o.CoreSchema
	if !IsNil(o.RequiredSchemaExtension) {
		toSerialize["requiredSchemaExtension"] = o.RequiredSchemaExtension
	}
	if !IsNil(o.OptionalSchemaExtension) {
		toSerialize["optionalSchemaExtension"] = o.OptionalSchemaExtension
	}
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

type NullableAddMappingScimResourceTypeRequest struct {
	value *AddMappingScimResourceTypeRequest
	isSet bool
}

func (v NullableAddMappingScimResourceTypeRequest) Get() *AddMappingScimResourceTypeRequest {
	return v.value
}

func (v *NullableAddMappingScimResourceTypeRequest) Set(val *AddMappingScimResourceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMappingScimResourceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMappingScimResourceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMappingScimResourceTypeRequest(val *AddMappingScimResourceTypeRequest) *NullableAddMappingScimResourceTypeRequest {
	return &NullableAddMappingScimResourceTypeRequest{value: val, isSet: true}
}

func (v NullableAddMappingScimResourceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMappingScimResourceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
