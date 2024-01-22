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

// checks if the BooleanTokenClaimValidationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BooleanTokenClaimValidationResponse{}

// BooleanTokenClaimValidationResponse struct for BooleanTokenClaimValidationResponse
type BooleanTokenClaimValidationResponse struct {
	Schemas       []EnumbooleanTokenClaimValidationSchemaUrn `json:"schemas"`
	RequiredValue EnumtokenClaimValidationRequiredValueProp  `json:"requiredValue"`
	// A description for this Token Claim Validation
	Description *string `json:"description,omitempty"`
	// The name of the claim to be validated.
	ClaimName                                     string                                             `json:"claimName"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Token Claim Validation
	Id string `json:"id"`
}

// NewBooleanTokenClaimValidationResponse instantiates a new BooleanTokenClaimValidationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanTokenClaimValidationResponse(schemas []EnumbooleanTokenClaimValidationSchemaUrn, requiredValue EnumtokenClaimValidationRequiredValueProp, claimName string, id string) *BooleanTokenClaimValidationResponse {
	this := BooleanTokenClaimValidationResponse{}
	this.Schemas = schemas
	this.RequiredValue = requiredValue
	this.ClaimName = claimName
	this.Id = id
	return &this
}

// NewBooleanTokenClaimValidationResponseWithDefaults instantiates a new BooleanTokenClaimValidationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanTokenClaimValidationResponseWithDefaults() *BooleanTokenClaimValidationResponse {
	this := BooleanTokenClaimValidationResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *BooleanTokenClaimValidationResponse) GetSchemas() []EnumbooleanTokenClaimValidationSchemaUrn {
	if o == nil {
		var ret []EnumbooleanTokenClaimValidationSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetSchemasOk() ([]EnumbooleanTokenClaimValidationSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *BooleanTokenClaimValidationResponse) SetSchemas(v []EnumbooleanTokenClaimValidationSchemaUrn) {
	o.Schemas = v
}

// GetRequiredValue returns the RequiredValue field value
func (o *BooleanTokenClaimValidationResponse) GetRequiredValue() EnumtokenClaimValidationRequiredValueProp {
	if o == nil {
		var ret EnumtokenClaimValidationRequiredValueProp
		return ret
	}

	return o.RequiredValue
}

// GetRequiredValueOk returns a tuple with the RequiredValue field value
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetRequiredValueOk() (*EnumtokenClaimValidationRequiredValueProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredValue, true
}

// SetRequiredValue sets field value
func (o *BooleanTokenClaimValidationResponse) SetRequiredValue(v EnumtokenClaimValidationRequiredValueProp) {
	o.RequiredValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BooleanTokenClaimValidationResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BooleanTokenClaimValidationResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BooleanTokenClaimValidationResponse) SetDescription(v string) {
	o.Description = &v
}

// GetClaimName returns the ClaimName field value
func (o *BooleanTokenClaimValidationResponse) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimName, true
}

// SetClaimName sets field value
func (o *BooleanTokenClaimValidationResponse) SetClaimName(v string) {
	o.ClaimName = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BooleanTokenClaimValidationResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BooleanTokenClaimValidationResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *BooleanTokenClaimValidationResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *BooleanTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *BooleanTokenClaimValidationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *BooleanTokenClaimValidationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *BooleanTokenClaimValidationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BooleanTokenClaimValidationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BooleanTokenClaimValidationResponse) SetId(v string) {
	o.Id = v
}

func (o BooleanTokenClaimValidationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BooleanTokenClaimValidationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["requiredValue"] = o.RequiredValue
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["claimName"] = o.ClaimName
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableBooleanTokenClaimValidationResponse struct {
	value *BooleanTokenClaimValidationResponse
	isSet bool
}

func (v NullableBooleanTokenClaimValidationResponse) Get() *BooleanTokenClaimValidationResponse {
	return v.value
}

func (v *NullableBooleanTokenClaimValidationResponse) Set(val *BooleanTokenClaimValidationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanTokenClaimValidationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanTokenClaimValidationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanTokenClaimValidationResponse(val *BooleanTokenClaimValidationResponse) *NullableBooleanTokenClaimValidationResponse {
	return &NullableBooleanTokenClaimValidationResponse{value: val, isSet: true}
}

func (v NullableBooleanTokenClaimValidationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanTokenClaimValidationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
