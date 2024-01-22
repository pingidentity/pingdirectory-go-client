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

// checks if the AddScimAttributeMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddScimAttributeMappingRequest{}

// AddScimAttributeMappingRequest struct for AddScimAttributeMappingRequest
type AddScimAttributeMappingRequest struct {
	Schemas []EnumscimAttributeMappingSchemaUrn `json:"schemas,omitempty"`
	// The Correlated LDAP Data View that persists the mapped SCIM Resource Type attribute(s).
	CorrelatedLDAPDataView *string `json:"correlatedLDAPDataView,omitempty"`
	// The attribute path of SCIM Resource Type attributes to be mapped.
	ScimResourceTypeAttribute string `json:"scimResourceTypeAttribute"`
	// The LDAP attribute to be mapped, or the path to a specific field of an LDAP attribute with the JSON object attribute syntax.
	LdapAttribute string `json:"ldapAttribute"`
	// Specifies whether the mapping is used to map from LDAP attribute to SCIM Resource Type attribute in a read operation.
	Readable *bool `json:"readable,omitempty"`
	// Specifies that the mapping is used to map from SCIM Resource Type attribute to LDAP attribute in a write operation.
	Writable *bool `json:"writable,omitempty"`
	// Specifies that the mapping is used to map from SCIM Resource Type attribute to LDAP attribute in a search filter.
	Searchable *bool `json:"searchable,omitempty"`
	// Specifies that the mapping is authoritative over other mappings for the same SCIM Resource Type attribute (for read operations).
	Authoritative *bool `json:"authoritative,omitempty"`
	// Name of the new SCIM Attribute Mapping
	MappingName string `json:"mappingName"`
}

// NewAddScimAttributeMappingRequest instantiates a new AddScimAttributeMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScimAttributeMappingRequest(scimResourceTypeAttribute string, ldapAttribute string, mappingName string) *AddScimAttributeMappingRequest {
	this := AddScimAttributeMappingRequest{}
	this.ScimResourceTypeAttribute = scimResourceTypeAttribute
	this.LdapAttribute = ldapAttribute
	this.MappingName = mappingName
	return &this
}

// NewAddScimAttributeMappingRequestWithDefaults instantiates a new AddScimAttributeMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScimAttributeMappingRequestWithDefaults() *AddScimAttributeMappingRequest {
	this := AddScimAttributeMappingRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddScimAttributeMappingRequest) GetSchemas() []EnumscimAttributeMappingSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumscimAttributeMappingSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetSchemasOk() ([]EnumscimAttributeMappingSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddScimAttributeMappingRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumscimAttributeMappingSchemaUrn and assigns it to the Schemas field.
func (o *AddScimAttributeMappingRequest) SetSchemas(v []EnumscimAttributeMappingSchemaUrn) {
	o.Schemas = v
}

// GetCorrelatedLDAPDataView returns the CorrelatedLDAPDataView field value if set, zero value otherwise.
func (o *AddScimAttributeMappingRequest) GetCorrelatedLDAPDataView() string {
	if o == nil || IsNil(o.CorrelatedLDAPDataView) {
		var ret string
		return ret
	}
	return *o.CorrelatedLDAPDataView
}

// GetCorrelatedLDAPDataViewOk returns a tuple with the CorrelatedLDAPDataView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetCorrelatedLDAPDataViewOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelatedLDAPDataView) {
		return nil, false
	}
	return o.CorrelatedLDAPDataView, true
}

// HasCorrelatedLDAPDataView returns a boolean if a field has been set.
func (o *AddScimAttributeMappingRequest) HasCorrelatedLDAPDataView() bool {
	if o != nil && !IsNil(o.CorrelatedLDAPDataView) {
		return true
	}

	return false
}

// SetCorrelatedLDAPDataView gets a reference to the given string and assigns it to the CorrelatedLDAPDataView field.
func (o *AddScimAttributeMappingRequest) SetCorrelatedLDAPDataView(v string) {
	o.CorrelatedLDAPDataView = &v
}

// GetScimResourceTypeAttribute returns the ScimResourceTypeAttribute field value
func (o *AddScimAttributeMappingRequest) GetScimResourceTypeAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimResourceTypeAttribute
}

// GetScimResourceTypeAttributeOk returns a tuple with the ScimResourceTypeAttribute field value
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetScimResourceTypeAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimResourceTypeAttribute, true
}

// SetScimResourceTypeAttribute sets field value
func (o *AddScimAttributeMappingRequest) SetScimResourceTypeAttribute(v string) {
	o.ScimResourceTypeAttribute = v
}

// GetLdapAttribute returns the LdapAttribute field value
func (o *AddScimAttributeMappingRequest) GetLdapAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdapAttribute
}

// GetLdapAttributeOk returns a tuple with the LdapAttribute field value
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetLdapAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapAttribute, true
}

// SetLdapAttribute sets field value
func (o *AddScimAttributeMappingRequest) SetLdapAttribute(v string) {
	o.LdapAttribute = v
}

// GetReadable returns the Readable field value if set, zero value otherwise.
func (o *AddScimAttributeMappingRequest) GetReadable() bool {
	if o == nil || IsNil(o.Readable) {
		var ret bool
		return ret
	}
	return *o.Readable
}

// GetReadableOk returns a tuple with the Readable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetReadableOk() (*bool, bool) {
	if o == nil || IsNil(o.Readable) {
		return nil, false
	}
	return o.Readable, true
}

// HasReadable returns a boolean if a field has been set.
func (o *AddScimAttributeMappingRequest) HasReadable() bool {
	if o != nil && !IsNil(o.Readable) {
		return true
	}

	return false
}

// SetReadable gets a reference to the given bool and assigns it to the Readable field.
func (o *AddScimAttributeMappingRequest) SetReadable(v bool) {
	o.Readable = &v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *AddScimAttributeMappingRequest) GetWritable() bool {
	if o == nil || IsNil(o.Writable) {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetWritableOk() (*bool, bool) {
	if o == nil || IsNil(o.Writable) {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *AddScimAttributeMappingRequest) HasWritable() bool {
	if o != nil && !IsNil(o.Writable) {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *AddScimAttributeMappingRequest) SetWritable(v bool) {
	o.Writable = &v
}

// GetSearchable returns the Searchable field value if set, zero value otherwise.
func (o *AddScimAttributeMappingRequest) GetSearchable() bool {
	if o == nil || IsNil(o.Searchable) {
		var ret bool
		return ret
	}
	return *o.Searchable
}

// GetSearchableOk returns a tuple with the Searchable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetSearchableOk() (*bool, bool) {
	if o == nil || IsNil(o.Searchable) {
		return nil, false
	}
	return o.Searchable, true
}

// HasSearchable returns a boolean if a field has been set.
func (o *AddScimAttributeMappingRequest) HasSearchable() bool {
	if o != nil && !IsNil(o.Searchable) {
		return true
	}

	return false
}

// SetSearchable gets a reference to the given bool and assigns it to the Searchable field.
func (o *AddScimAttributeMappingRequest) SetSearchable(v bool) {
	o.Searchable = &v
}

// GetAuthoritative returns the Authoritative field value if set, zero value otherwise.
func (o *AddScimAttributeMappingRequest) GetAuthoritative() bool {
	if o == nil || IsNil(o.Authoritative) {
		var ret bool
		return ret
	}
	return *o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetAuthoritativeOk() (*bool, bool) {
	if o == nil || IsNil(o.Authoritative) {
		return nil, false
	}
	return o.Authoritative, true
}

// HasAuthoritative returns a boolean if a field has been set.
func (o *AddScimAttributeMappingRequest) HasAuthoritative() bool {
	if o != nil && !IsNil(o.Authoritative) {
		return true
	}

	return false
}

// SetAuthoritative gets a reference to the given bool and assigns it to the Authoritative field.
func (o *AddScimAttributeMappingRequest) SetAuthoritative(v bool) {
	o.Authoritative = &v
}

// GetMappingName returns the MappingName field value
func (o *AddScimAttributeMappingRequest) GetMappingName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MappingName
}

// GetMappingNameOk returns a tuple with the MappingName field value
// and a boolean to check if the value has been set.
func (o *AddScimAttributeMappingRequest) GetMappingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MappingName, true
}

// SetMappingName sets field value
func (o *AddScimAttributeMappingRequest) SetMappingName(v string) {
	o.MappingName = v
}

func (o AddScimAttributeMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddScimAttributeMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.CorrelatedLDAPDataView) {
		toSerialize["correlatedLDAPDataView"] = o.CorrelatedLDAPDataView
	}
	toSerialize["scimResourceTypeAttribute"] = o.ScimResourceTypeAttribute
	toSerialize["ldapAttribute"] = o.LdapAttribute
	if !IsNil(o.Readable) {
		toSerialize["readable"] = o.Readable
	}
	if !IsNil(o.Writable) {
		toSerialize["writable"] = o.Writable
	}
	if !IsNil(o.Searchable) {
		toSerialize["searchable"] = o.Searchable
	}
	if !IsNil(o.Authoritative) {
		toSerialize["authoritative"] = o.Authoritative
	}
	toSerialize["mappingName"] = o.MappingName
	return toSerialize, nil
}

type NullableAddScimAttributeMappingRequest struct {
	value *AddScimAttributeMappingRequest
	isSet bool
}

func (v NullableAddScimAttributeMappingRequest) Get() *AddScimAttributeMappingRequest {
	return v.value
}

func (v *NullableAddScimAttributeMappingRequest) Set(val *AddScimAttributeMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScimAttributeMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScimAttributeMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScimAttributeMappingRequest(val *AddScimAttributeMappingRequest) *NullableAddScimAttributeMappingRequest {
	return &NullableAddScimAttributeMappingRequest{value: val, isSet: true}
}

func (v NullableAddScimAttributeMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScimAttributeMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
