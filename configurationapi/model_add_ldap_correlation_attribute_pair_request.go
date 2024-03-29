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

// checks if the AddLdapCorrelationAttributePairRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLdapCorrelationAttributePairRequest{}

// AddLdapCorrelationAttributePairRequest struct for AddLdapCorrelationAttributePairRequest
type AddLdapCorrelationAttributePairRequest struct {
	Schemas []EnumldapCorrelationAttributePairSchemaUrn `json:"schemas,omitempty"`
	// The LDAP attribute from the base SCIM Resource Type whose value will be used to match objects in the Correlated LDAP Data View.
	PrimaryCorrelationAttribute string `json:"primaryCorrelationAttribute"`
	// The LDAP attribute from the Correlated LDAP Data View whose value will be matched.
	SecondaryCorrelationAttribute string `json:"secondaryCorrelationAttribute"`
	// Name of the new LDAP Correlation Attribute Pair
	PairName string `json:"pairName"`
}

// NewAddLdapCorrelationAttributePairRequest instantiates a new AddLdapCorrelationAttributePairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLdapCorrelationAttributePairRequest(primaryCorrelationAttribute string, secondaryCorrelationAttribute string, pairName string) *AddLdapCorrelationAttributePairRequest {
	this := AddLdapCorrelationAttributePairRequest{}
	this.PrimaryCorrelationAttribute = primaryCorrelationAttribute
	this.SecondaryCorrelationAttribute = secondaryCorrelationAttribute
	this.PairName = pairName
	return &this
}

// NewAddLdapCorrelationAttributePairRequestWithDefaults instantiates a new AddLdapCorrelationAttributePairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLdapCorrelationAttributePairRequestWithDefaults() *AddLdapCorrelationAttributePairRequest {
	this := AddLdapCorrelationAttributePairRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddLdapCorrelationAttributePairRequest) GetSchemas() []EnumldapCorrelationAttributePairSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumldapCorrelationAttributePairSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapCorrelationAttributePairRequest) GetSchemasOk() ([]EnumldapCorrelationAttributePairSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddLdapCorrelationAttributePairRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumldapCorrelationAttributePairSchemaUrn and assigns it to the Schemas field.
func (o *AddLdapCorrelationAttributePairRequest) SetSchemas(v []EnumldapCorrelationAttributePairSchemaUrn) {
	o.Schemas = v
}

// GetPrimaryCorrelationAttribute returns the PrimaryCorrelationAttribute field value
func (o *AddLdapCorrelationAttributePairRequest) GetPrimaryCorrelationAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryCorrelationAttribute
}

// GetPrimaryCorrelationAttributeOk returns a tuple with the PrimaryCorrelationAttribute field value
// and a boolean to check if the value has been set.
func (o *AddLdapCorrelationAttributePairRequest) GetPrimaryCorrelationAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryCorrelationAttribute, true
}

// SetPrimaryCorrelationAttribute sets field value
func (o *AddLdapCorrelationAttributePairRequest) SetPrimaryCorrelationAttribute(v string) {
	o.PrimaryCorrelationAttribute = v
}

// GetSecondaryCorrelationAttribute returns the SecondaryCorrelationAttribute field value
func (o *AddLdapCorrelationAttributePairRequest) GetSecondaryCorrelationAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryCorrelationAttribute
}

// GetSecondaryCorrelationAttributeOk returns a tuple with the SecondaryCorrelationAttribute field value
// and a boolean to check if the value has been set.
func (o *AddLdapCorrelationAttributePairRequest) GetSecondaryCorrelationAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryCorrelationAttribute, true
}

// SetSecondaryCorrelationAttribute sets field value
func (o *AddLdapCorrelationAttributePairRequest) SetSecondaryCorrelationAttribute(v string) {
	o.SecondaryCorrelationAttribute = v
}

// GetPairName returns the PairName field value
func (o *AddLdapCorrelationAttributePairRequest) GetPairName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PairName
}

// GetPairNameOk returns a tuple with the PairName field value
// and a boolean to check if the value has been set.
func (o *AddLdapCorrelationAttributePairRequest) GetPairNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PairName, true
}

// SetPairName sets field value
func (o *AddLdapCorrelationAttributePairRequest) SetPairName(v string) {
	o.PairName = v
}

func (o AddLdapCorrelationAttributePairRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLdapCorrelationAttributePairRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["primaryCorrelationAttribute"] = o.PrimaryCorrelationAttribute
	toSerialize["secondaryCorrelationAttribute"] = o.SecondaryCorrelationAttribute
	toSerialize["pairName"] = o.PairName
	return toSerialize, nil
}

type NullableAddLdapCorrelationAttributePairRequest struct {
	value *AddLdapCorrelationAttributePairRequest
	isSet bool
}

func (v NullableAddLdapCorrelationAttributePairRequest) Get() *AddLdapCorrelationAttributePairRequest {
	return v.value
}

func (v *NullableAddLdapCorrelationAttributePairRequest) Set(val *AddLdapCorrelationAttributePairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLdapCorrelationAttributePairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLdapCorrelationAttributePairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLdapCorrelationAttributePairRequest(val *AddLdapCorrelationAttributePairRequest) *NullableAddLdapCorrelationAttributePairRequest {
	return &NullableAddLdapCorrelationAttributePairRequest{value: val, isSet: true}
}

func (v NullableAddLdapCorrelationAttributePairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLdapCorrelationAttributePairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
