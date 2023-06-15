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

// checks if the ScimAttributeMappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimAttributeMappingResponse{}

// ScimAttributeMappingResponse struct for ScimAttributeMappingResponse
type ScimAttributeMappingResponse struct {
	// Name of the SCIM Attribute Mapping
	Id      string                              `json:"id"`
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
	Authoritative                                 *bool                                              `json:"authoritative,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewScimAttributeMappingResponse instantiates a new ScimAttributeMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimAttributeMappingResponse(id string, scimResourceTypeAttribute string, ldapAttribute string) *ScimAttributeMappingResponse {
	this := ScimAttributeMappingResponse{}
	this.Id = id
	this.ScimResourceTypeAttribute = scimResourceTypeAttribute
	this.LdapAttribute = ldapAttribute
	return &this
}

// NewScimAttributeMappingResponseWithDefaults instantiates a new ScimAttributeMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimAttributeMappingResponseWithDefaults() *ScimAttributeMappingResponse {
	this := ScimAttributeMappingResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ScimAttributeMappingResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScimAttributeMappingResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetSchemas() []EnumscimAttributeMappingSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumscimAttributeMappingSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetSchemasOk() ([]EnumscimAttributeMappingSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumscimAttributeMappingSchemaUrn and assigns it to the Schemas field.
func (o *ScimAttributeMappingResponse) SetSchemas(v []EnumscimAttributeMappingSchemaUrn) {
	o.Schemas = v
}

// GetCorrelatedLDAPDataView returns the CorrelatedLDAPDataView field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetCorrelatedLDAPDataView() string {
	if o == nil || IsNil(o.CorrelatedLDAPDataView) {
		var ret string
		return ret
	}
	return *o.CorrelatedLDAPDataView
}

// GetCorrelatedLDAPDataViewOk returns a tuple with the CorrelatedLDAPDataView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetCorrelatedLDAPDataViewOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelatedLDAPDataView) {
		return nil, false
	}
	return o.CorrelatedLDAPDataView, true
}

// HasCorrelatedLDAPDataView returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasCorrelatedLDAPDataView() bool {
	if o != nil && !IsNil(o.CorrelatedLDAPDataView) {
		return true
	}

	return false
}

// SetCorrelatedLDAPDataView gets a reference to the given string and assigns it to the CorrelatedLDAPDataView field.
func (o *ScimAttributeMappingResponse) SetCorrelatedLDAPDataView(v string) {
	o.CorrelatedLDAPDataView = &v
}

// GetScimResourceTypeAttribute returns the ScimResourceTypeAttribute field value
func (o *ScimAttributeMappingResponse) GetScimResourceTypeAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimResourceTypeAttribute
}

// GetScimResourceTypeAttributeOk returns a tuple with the ScimResourceTypeAttribute field value
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetScimResourceTypeAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimResourceTypeAttribute, true
}

// SetScimResourceTypeAttribute sets field value
func (o *ScimAttributeMappingResponse) SetScimResourceTypeAttribute(v string) {
	o.ScimResourceTypeAttribute = v
}

// GetLdapAttribute returns the LdapAttribute field value
func (o *ScimAttributeMappingResponse) GetLdapAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdapAttribute
}

// GetLdapAttributeOk returns a tuple with the LdapAttribute field value
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetLdapAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapAttribute, true
}

// SetLdapAttribute sets field value
func (o *ScimAttributeMappingResponse) SetLdapAttribute(v string) {
	o.LdapAttribute = v
}

// GetReadable returns the Readable field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetReadable() bool {
	if o == nil || IsNil(o.Readable) {
		var ret bool
		return ret
	}
	return *o.Readable
}

// GetReadableOk returns a tuple with the Readable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetReadableOk() (*bool, bool) {
	if o == nil || IsNil(o.Readable) {
		return nil, false
	}
	return o.Readable, true
}

// HasReadable returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasReadable() bool {
	if o != nil && !IsNil(o.Readable) {
		return true
	}

	return false
}

// SetReadable gets a reference to the given bool and assigns it to the Readable field.
func (o *ScimAttributeMappingResponse) SetReadable(v bool) {
	o.Readable = &v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetWritable() bool {
	if o == nil || IsNil(o.Writable) {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetWritableOk() (*bool, bool) {
	if o == nil || IsNil(o.Writable) {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasWritable() bool {
	if o != nil && !IsNil(o.Writable) {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *ScimAttributeMappingResponse) SetWritable(v bool) {
	o.Writable = &v
}

// GetSearchable returns the Searchable field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetSearchable() bool {
	if o == nil || IsNil(o.Searchable) {
		var ret bool
		return ret
	}
	return *o.Searchable
}

// GetSearchableOk returns a tuple with the Searchable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetSearchableOk() (*bool, bool) {
	if o == nil || IsNil(o.Searchable) {
		return nil, false
	}
	return o.Searchable, true
}

// HasSearchable returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasSearchable() bool {
	if o != nil && !IsNil(o.Searchable) {
		return true
	}

	return false
}

// SetSearchable gets a reference to the given bool and assigns it to the Searchable field.
func (o *ScimAttributeMappingResponse) SetSearchable(v bool) {
	o.Searchable = &v
}

// GetAuthoritative returns the Authoritative field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetAuthoritative() bool {
	if o == nil || IsNil(o.Authoritative) {
		var ret bool
		return ret
	}
	return *o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetAuthoritativeOk() (*bool, bool) {
	if o == nil || IsNil(o.Authoritative) {
		return nil, false
	}
	return o.Authoritative, true
}

// HasAuthoritative returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasAuthoritative() bool {
	if o != nil && !IsNil(o.Authoritative) {
		return true
	}

	return false
}

// SetAuthoritative gets a reference to the given bool and assigns it to the Authoritative field.
func (o *ScimAttributeMappingResponse) SetAuthoritative(v bool) {
	o.Authoritative = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ScimAttributeMappingResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ScimAttributeMappingResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimAttributeMappingResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ScimAttributeMappingResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ScimAttributeMappingResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ScimAttributeMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimAttributeMappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableScimAttributeMappingResponse struct {
	value *ScimAttributeMappingResponse
	isSet bool
}

func (v NullableScimAttributeMappingResponse) Get() *ScimAttributeMappingResponse {
	return v.value
}

func (v *NullableScimAttributeMappingResponse) Set(val *ScimAttributeMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScimAttributeMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScimAttributeMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimAttributeMappingResponse(val *ScimAttributeMappingResponse) *NullableScimAttributeMappingResponse {
	return &NullableScimAttributeMappingResponse{value: val, isSet: true}
}

func (v NullableScimAttributeMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimAttributeMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
