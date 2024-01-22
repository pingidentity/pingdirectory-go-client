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

// checks if the ScimSchemaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimSchemaResponse{}

// ScimSchemaResponse struct for ScimSchemaResponse
type ScimSchemaResponse struct {
	Schemas []EnumscimSchemaSchemaUrn `json:"schemas,omitempty"`
	// A description for this SCIM Schema
	Description *string `json:"description,omitempty"`
	// The URN which identifies this SCIM Schema.
	SchemaURN string `json:"schemaURN"`
	// The human readable name for this SCIM Schema.
	DisplayName                                   *string                                            `json:"displayName,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the SCIM Schema
	Id string `json:"id"`
}

// NewScimSchemaResponse instantiates a new ScimSchemaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimSchemaResponse(schemaURN string, id string) *ScimSchemaResponse {
	this := ScimSchemaResponse{}
	this.SchemaURN = schemaURN
	this.Id = id
	return &this
}

// NewScimSchemaResponseWithDefaults instantiates a new ScimSchemaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimSchemaResponseWithDefaults() *ScimSchemaResponse {
	this := ScimSchemaResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimSchemaResponse) GetSchemas() []EnumscimSchemaSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumscimSchemaSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetSchemasOk() ([]EnumscimSchemaSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimSchemaResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumscimSchemaSchemaUrn and assigns it to the Schemas field.
func (o *ScimSchemaResponse) SetSchemas(v []EnumscimSchemaSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScimSchemaResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScimSchemaResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScimSchemaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetSchemaURN returns the SchemaURN field value
func (o *ScimSchemaResponse) GetSchemaURN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaURN
}

// GetSchemaURNOk returns a tuple with the SchemaURN field value
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetSchemaURNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaURN, true
}

// SetSchemaURN sets field value
func (o *ScimSchemaResponse) SetSchemaURN(v string) {
	o.SchemaURN = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ScimSchemaResponse) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScimSchemaResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ScimSchemaResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimSchemaResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimSchemaResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ScimSchemaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ScimSchemaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ScimSchemaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ScimSchemaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ScimSchemaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScimSchemaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScimSchemaResponse) SetId(v string) {
	o.Id = v
}

func (o ScimSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimSchemaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["schemaURN"] = o.SchemaURN
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
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

type NullableScimSchemaResponse struct {
	value *ScimSchemaResponse
	isSet bool
}

func (v NullableScimSchemaResponse) Get() *ScimSchemaResponse {
	return v.value
}

func (v *NullableScimSchemaResponse) Set(val *ScimSchemaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScimSchemaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScimSchemaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimSchemaResponse(val *ScimSchemaResponse) *NullableScimSchemaResponse {
	return &NullableScimSchemaResponse{value: val, isSet: true}
}

func (v NullableScimSchemaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimSchemaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
