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

// checks if the ExactMatchIdentityMapperResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExactMatchIdentityMapperResponse{}

// ExactMatchIdentityMapperResponse struct for ExactMatchIdentityMapperResponse
type ExactMatchIdentityMapperResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Identity Mapper
	Id      string                                  `json:"id"`
	Schemas []EnumexactMatchIdentityMapperSchemaUrn `json:"schemas"`
	// Specifies the attribute whose value should exactly match the ID string provided to this identity mapper.
	MatchAttribute []string `json:"matchAttribute"`
	// Specifies the set of base DNs below which to search for users.
	MatchBaseDN []string `json:"matchBaseDN,omitempty"`
	// An optional filter that mapped users must match.
	MatchFilter *string `json:"matchFilter,omitempty"`
	// A description for this Identity Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Identity Mapper is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewExactMatchIdentityMapperResponse instantiates a new ExactMatchIdentityMapperResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExactMatchIdentityMapperResponse(id string, schemas []EnumexactMatchIdentityMapperSchemaUrn, matchAttribute []string, enabled bool) *ExactMatchIdentityMapperResponse {
	this := ExactMatchIdentityMapperResponse{}
	this.Id = id
	this.Schemas = schemas
	this.MatchAttribute = matchAttribute
	this.Enabled = enabled
	return &this
}

// NewExactMatchIdentityMapperResponseWithDefaults instantiates a new ExactMatchIdentityMapperResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExactMatchIdentityMapperResponseWithDefaults() *ExactMatchIdentityMapperResponse {
	this := ExactMatchIdentityMapperResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExactMatchIdentityMapperResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExactMatchIdentityMapperResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ExactMatchIdentityMapperResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ExactMatchIdentityMapperResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ExactMatchIdentityMapperResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ExactMatchIdentityMapperResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ExactMatchIdentityMapperResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExactMatchIdentityMapperResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ExactMatchIdentityMapperResponse) GetSchemas() []EnumexactMatchIdentityMapperSchemaUrn {
	if o == nil {
		var ret []EnumexactMatchIdentityMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetSchemasOk() ([]EnumexactMatchIdentityMapperSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ExactMatchIdentityMapperResponse) SetSchemas(v []EnumexactMatchIdentityMapperSchemaUrn) {
	o.Schemas = v
}

// GetMatchAttribute returns the MatchAttribute field value
func (o *ExactMatchIdentityMapperResponse) GetMatchAttribute() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MatchAttribute
}

// GetMatchAttributeOk returns a tuple with the MatchAttribute field value
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetMatchAttributeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MatchAttribute, true
}

// SetMatchAttribute sets field value
func (o *ExactMatchIdentityMapperResponse) SetMatchAttribute(v []string) {
	o.MatchAttribute = v
}

// GetMatchBaseDN returns the MatchBaseDN field value if set, zero value otherwise.
func (o *ExactMatchIdentityMapperResponse) GetMatchBaseDN() []string {
	if o == nil || IsNil(o.MatchBaseDN) {
		var ret []string
		return ret
	}
	return o.MatchBaseDN
}

// GetMatchBaseDNOk returns a tuple with the MatchBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetMatchBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.MatchBaseDN) {
		return nil, false
	}
	return o.MatchBaseDN, true
}

// HasMatchBaseDN returns a boolean if a field has been set.
func (o *ExactMatchIdentityMapperResponse) HasMatchBaseDN() bool {
	if o != nil && !IsNil(o.MatchBaseDN) {
		return true
	}

	return false
}

// SetMatchBaseDN gets a reference to the given []string and assigns it to the MatchBaseDN field.
func (o *ExactMatchIdentityMapperResponse) SetMatchBaseDN(v []string) {
	o.MatchBaseDN = v
}

// GetMatchFilter returns the MatchFilter field value if set, zero value otherwise.
func (o *ExactMatchIdentityMapperResponse) GetMatchFilter() string {
	if o == nil || IsNil(o.MatchFilter) {
		var ret string
		return ret
	}
	return *o.MatchFilter
}

// GetMatchFilterOk returns a tuple with the MatchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetMatchFilterOk() (*string, bool) {
	if o == nil || IsNil(o.MatchFilter) {
		return nil, false
	}
	return o.MatchFilter, true
}

// HasMatchFilter returns a boolean if a field has been set.
func (o *ExactMatchIdentityMapperResponse) HasMatchFilter() bool {
	if o != nil && !IsNil(o.MatchFilter) {
		return true
	}

	return false
}

// SetMatchFilter gets a reference to the given string and assigns it to the MatchFilter field.
func (o *ExactMatchIdentityMapperResponse) SetMatchFilter(v string) {
	o.MatchFilter = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExactMatchIdentityMapperResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExactMatchIdentityMapperResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExactMatchIdentityMapperResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ExactMatchIdentityMapperResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExactMatchIdentityMapperResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExactMatchIdentityMapperResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ExactMatchIdentityMapperResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExactMatchIdentityMapperResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["matchAttribute"] = o.MatchAttribute
	if !IsNil(o.MatchBaseDN) {
		toSerialize["matchBaseDN"] = o.MatchBaseDN
	}
	if !IsNil(o.MatchFilter) {
		toSerialize["matchFilter"] = o.MatchFilter
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableExactMatchIdentityMapperResponse struct {
	value *ExactMatchIdentityMapperResponse
	isSet bool
}

func (v NullableExactMatchIdentityMapperResponse) Get() *ExactMatchIdentityMapperResponse {
	return v.value
}

func (v *NullableExactMatchIdentityMapperResponse) Set(val *ExactMatchIdentityMapperResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExactMatchIdentityMapperResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExactMatchIdentityMapperResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExactMatchIdentityMapperResponse(val *ExactMatchIdentityMapperResponse) *NullableExactMatchIdentityMapperResponse {
	return &NullableExactMatchIdentityMapperResponse{value: val, isSet: true}
}

func (v NullableExactMatchIdentityMapperResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExactMatchIdentityMapperResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
