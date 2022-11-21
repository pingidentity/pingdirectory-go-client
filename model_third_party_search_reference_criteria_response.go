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

// ThirdPartySearchReferenceCriteriaResponse struct for ThirdPartySearchReferenceCriteriaResponse
type ThirdPartySearchReferenceCriteriaResponse struct {
	// Name of the Search Reference Criteria
	Id string `json:"id"`
	Schemas []EnumthirdPartySearchReferenceCriteriaSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Search Reference Criteria.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Search Reference Criteria. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Search Reference Criteria
	Description *string `json:"description,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewThirdPartySearchReferenceCriteriaResponse instantiates a new ThirdPartySearchReferenceCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartySearchReferenceCriteriaResponse(id string, schemas []EnumthirdPartySearchReferenceCriteriaSchemaUrn, extensionClass string) *ThirdPartySearchReferenceCriteriaResponse {
	this := ThirdPartySearchReferenceCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	return &this
}

// NewThirdPartySearchReferenceCriteriaResponseWithDefaults instantiates a new ThirdPartySearchReferenceCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartySearchReferenceCriteriaResponseWithDefaults() *ThirdPartySearchReferenceCriteriaResponse {
	this := ThirdPartySearchReferenceCriteriaResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartySearchReferenceCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartySearchReferenceCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartySearchReferenceCriteriaResponse) GetSchemas() []EnumthirdPartySearchReferenceCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartySearchReferenceCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetSchemasOk() ([]EnumthirdPartySearchReferenceCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartySearchReferenceCriteriaResponse) SetSchemas(v []EnumthirdPartySearchReferenceCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartySearchReferenceCriteriaResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartySearchReferenceCriteriaResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartySearchReferenceCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartySearchReferenceCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartySearchReferenceCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o ThirdPartySearchReferenceCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartySearchReferenceCriteriaResponse struct {
	value *ThirdPartySearchReferenceCriteriaResponse
	isSet bool
}

func (v NullableThirdPartySearchReferenceCriteriaResponse) Get() *ThirdPartySearchReferenceCriteriaResponse {
	return v.value
}

func (v *NullableThirdPartySearchReferenceCriteriaResponse) Set(val *ThirdPartySearchReferenceCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartySearchReferenceCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartySearchReferenceCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartySearchReferenceCriteriaResponse(val *ThirdPartySearchReferenceCriteriaResponse) *NullableThirdPartySearchReferenceCriteriaResponse {
	return &NullableThirdPartySearchReferenceCriteriaResponse{value: val, isSet: true}
}

func (v NullableThirdPartySearchReferenceCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartySearchReferenceCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


