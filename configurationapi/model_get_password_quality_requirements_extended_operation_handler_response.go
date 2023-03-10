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

// checks if the GetPasswordQualityRequirementsExtendedOperationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPasswordQualityRequirementsExtendedOperationHandlerResponse{}

// GetPasswordQualityRequirementsExtendedOperationHandlerResponse struct for GetPasswordQualityRequirementsExtendedOperationHandlerResponse
type GetPasswordQualityRequirementsExtendedOperationHandlerResponse struct {
	Meta                                          *MetaMeta                                                             `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20                    `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewGetPasswordQualityRequirementsExtendedOperationHandlerResponse instantiates a new GetPasswordQualityRequirementsExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPasswordQualityRequirementsExtendedOperationHandlerResponse(schemas []EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn, id string, enabled bool) *GetPasswordQualityRequirementsExtendedOperationHandlerResponse {
	this := GetPasswordQualityRequirementsExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewGetPasswordQualityRequirementsExtendedOperationHandlerResponseWithDefaults instantiates a new GetPasswordQualityRequirementsExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPasswordQualityRequirementsExtendedOperationHandlerResponseWithDefaults() *GetPasswordQualityRequirementsExtendedOperationHandlerResponse {
	this := GetPasswordQualityRequirementsExtendedOperationHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetSchemas() []EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) SetSchemas(v []EnumgetPasswordQualityRequirementsExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GetPasswordQualityRequirementsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPasswordQualityRequirementsExtendedOperationHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse struct {
	value *GetPasswordQualityRequirementsExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse) Get() *GetPasswordQualityRequirementsExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse) Set(val *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse(val *GetPasswordQualityRequirementsExtendedOperationHandlerResponse) *NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse {
	return &NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPasswordQualityRequirementsExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
