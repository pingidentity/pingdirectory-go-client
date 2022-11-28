/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConjurPasswordStorageSchemeResponse struct for ConjurPasswordStorageSchemeResponse
type ConjurPasswordStorageSchemeResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
	Schemas []EnumconjurPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// An external server definition with information needed to connect and authenticate to the Conjur instance containing user passwords.
	ConjurExternalServer string `json:"conjurExternalServer"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewConjurPasswordStorageSchemeResponse instantiates a new ConjurPasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConjurPasswordStorageSchemeResponse(id string, schemas []EnumconjurPasswordStorageSchemeSchemaUrn, conjurExternalServer string, enabled bool) *ConjurPasswordStorageSchemeResponse {
	this := ConjurPasswordStorageSchemeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ConjurExternalServer = conjurExternalServer
	this.Enabled = enabled
	return &this
}

// NewConjurPasswordStorageSchemeResponseWithDefaults instantiates a new ConjurPasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConjurPasswordStorageSchemeResponseWithDefaults() *ConjurPasswordStorageSchemeResponse {
	this := ConjurPasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConjurPasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConjurPasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ConjurPasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ConjurPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ConjurPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ConjurPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ConjurPasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConjurPasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ConjurPasswordStorageSchemeResponse) GetSchemas() []EnumconjurPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumconjurPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetSchemasOk() ([]EnumconjurPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ConjurPasswordStorageSchemeResponse) SetSchemas(v []EnumconjurPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetConjurExternalServer returns the ConjurExternalServer field value
func (o *ConjurPasswordStorageSchemeResponse) GetConjurExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurExternalServer
}

// GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field value
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetConjurExternalServerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConjurExternalServer, true
}

// SetConjurExternalServer sets field value
func (o *ConjurPasswordStorageSchemeResponse) SetConjurExternalServer(v string) {
	o.ConjurExternalServer = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConjurPasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConjurPasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConjurPasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ConjurPasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConjurPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConjurPasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ConjurPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["conjurExternalServer"] = o.ConjurExternalServer
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableConjurPasswordStorageSchemeResponse struct {
	value *ConjurPasswordStorageSchemeResponse
	isSet bool
}

func (v NullableConjurPasswordStorageSchemeResponse) Get() *ConjurPasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableConjurPasswordStorageSchemeResponse) Set(val *ConjurPasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConjurPasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConjurPasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConjurPasswordStorageSchemeResponse(val *ConjurPasswordStorageSchemeResponse) *NullableConjurPasswordStorageSchemeResponse {
	return &NullableConjurPasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableConjurPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConjurPasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


