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

// checks if the StringArrayTokenClaimValidationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringArrayTokenClaimValidationResponse{}

// StringArrayTokenClaimValidationResponse struct for StringArrayTokenClaimValidationResponse
type StringArrayTokenClaimValidationResponse struct {
	// Name of the ID Token Validator
	Id      string                                         `json:"id"`
	Schemas []EnumstringArrayTokenClaimValidationSchemaUrn `json:"schemas"`
	// The set of all values that the claim must have to be considered valid.
	AllRequiredValue []string `json:"allRequiredValue,omitempty"`
	// The set of values that the claim may have to be considered valid.
	AnyRequiredValue []string `json:"anyRequiredValue,omitempty"`
	// A description for this Token Claim Validation
	Description *string `json:"description,omitempty"`
	// The name of the claim to be validated.
	ClaimName                                     string                                             `json:"claimName"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewStringArrayTokenClaimValidationResponse instantiates a new StringArrayTokenClaimValidationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringArrayTokenClaimValidationResponse(id string, schemas []EnumstringArrayTokenClaimValidationSchemaUrn, claimName string) *StringArrayTokenClaimValidationResponse {
	this := StringArrayTokenClaimValidationResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ClaimName = claimName
	return &this
}

// NewStringArrayTokenClaimValidationResponseWithDefaults instantiates a new StringArrayTokenClaimValidationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringArrayTokenClaimValidationResponseWithDefaults() *StringArrayTokenClaimValidationResponse {
	this := StringArrayTokenClaimValidationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *StringArrayTokenClaimValidationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StringArrayTokenClaimValidationResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *StringArrayTokenClaimValidationResponse) GetSchemas() []EnumstringArrayTokenClaimValidationSchemaUrn {
	if o == nil {
		var ret []EnumstringArrayTokenClaimValidationSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetSchemasOk() ([]EnumstringArrayTokenClaimValidationSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *StringArrayTokenClaimValidationResponse) SetSchemas(v []EnumstringArrayTokenClaimValidationSchemaUrn) {
	o.Schemas = v
}

// GetAllRequiredValue returns the AllRequiredValue field value if set, zero value otherwise.
func (o *StringArrayTokenClaimValidationResponse) GetAllRequiredValue() []string {
	if o == nil || IsNil(o.AllRequiredValue) {
		var ret []string
		return ret
	}
	return o.AllRequiredValue
}

// GetAllRequiredValueOk returns a tuple with the AllRequiredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetAllRequiredValueOk() ([]string, bool) {
	if o == nil || IsNil(o.AllRequiredValue) {
		return nil, false
	}
	return o.AllRequiredValue, true
}

// HasAllRequiredValue returns a boolean if a field has been set.
func (o *StringArrayTokenClaimValidationResponse) HasAllRequiredValue() bool {
	if o != nil && !IsNil(o.AllRequiredValue) {
		return true
	}

	return false
}

// SetAllRequiredValue gets a reference to the given []string and assigns it to the AllRequiredValue field.
func (o *StringArrayTokenClaimValidationResponse) SetAllRequiredValue(v []string) {
	o.AllRequiredValue = v
}

// GetAnyRequiredValue returns the AnyRequiredValue field value if set, zero value otherwise.
func (o *StringArrayTokenClaimValidationResponse) GetAnyRequiredValue() []string {
	if o == nil || IsNil(o.AnyRequiredValue) {
		var ret []string
		return ret
	}
	return o.AnyRequiredValue
}

// GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetAnyRequiredValueOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyRequiredValue) {
		return nil, false
	}
	return o.AnyRequiredValue, true
}

// HasAnyRequiredValue returns a boolean if a field has been set.
func (o *StringArrayTokenClaimValidationResponse) HasAnyRequiredValue() bool {
	if o != nil && !IsNil(o.AnyRequiredValue) {
		return true
	}

	return false
}

// SetAnyRequiredValue gets a reference to the given []string and assigns it to the AnyRequiredValue field.
func (o *StringArrayTokenClaimValidationResponse) SetAnyRequiredValue(v []string) {
	o.AnyRequiredValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StringArrayTokenClaimValidationResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StringArrayTokenClaimValidationResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StringArrayTokenClaimValidationResponse) SetDescription(v string) {
	o.Description = &v
}

// GetClaimName returns the ClaimName field value
func (o *StringArrayTokenClaimValidationResponse) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimName, true
}

// SetClaimName sets field value
func (o *StringArrayTokenClaimValidationResponse) SetClaimName(v string) {
	o.ClaimName = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *StringArrayTokenClaimValidationResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *StringArrayTokenClaimValidationResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *StringArrayTokenClaimValidationResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *StringArrayTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringArrayTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *StringArrayTokenClaimValidationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *StringArrayTokenClaimValidationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o StringArrayTokenClaimValidationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringArrayTokenClaimValidationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllRequiredValue) {
		toSerialize["allRequiredValue"] = o.AllRequiredValue
	}
	if !IsNil(o.AnyRequiredValue) {
		toSerialize["anyRequiredValue"] = o.AnyRequiredValue
	}
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
	return toSerialize, nil
}

type NullableStringArrayTokenClaimValidationResponse struct {
	value *StringArrayTokenClaimValidationResponse
	isSet bool
}

func (v NullableStringArrayTokenClaimValidationResponse) Get() *StringArrayTokenClaimValidationResponse {
	return v.value
}

func (v *NullableStringArrayTokenClaimValidationResponse) Set(val *StringArrayTokenClaimValidationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStringArrayTokenClaimValidationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStringArrayTokenClaimValidationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringArrayTokenClaimValidationResponse(val *StringArrayTokenClaimValidationResponse) *NullableStringArrayTokenClaimValidationResponse {
	return &NullableStringArrayTokenClaimValidationResponse{value: val, isSet: true}
}

func (v NullableStringArrayTokenClaimValidationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringArrayTokenClaimValidationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
