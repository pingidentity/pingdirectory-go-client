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

// GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse struct for GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse
type GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse struct {
	Meta                                          *MetaMeta                                                                `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20                       `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumgetSupportedOtpDeliveryMechanismsExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse instantiates a new GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse(schemas []EnumgetSupportedOtpDeliveryMechanismsExtendedOperationHandlerSchemaUrn, id string, enabled bool) *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse {
	this := GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponseWithDefaults instantiates a new GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponseWithDefaults() *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse {
	this := GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetSchemas() []EnumgetSupportedOtpDeliveryMechanismsExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgetSupportedOtpDeliveryMechanismsExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumgetSupportedOtpDeliveryMechanismsExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) SetSchemas(v []EnumgetSupportedOtpDeliveryMechanismsExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse struct {
	value *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) Get() *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) Set(val *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse(val *GetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) *NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse {
	return &NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSupportedOtpDeliveryMechanismsExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
