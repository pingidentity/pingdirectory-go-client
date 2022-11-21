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

// ConsentDefinitionResponse struct for ConsentDefinitionResponse
type ConsentDefinitionResponse struct {
	// Name of the Consent Definition
	Id string `json:"id"`
	Schemas []EnumconsentDefinitionSchemaUrn `json:"schemas,omitempty"`
	// A version-independent unique identifier for this Consent Definition.
	UniqueID string `json:"uniqueID"`
	// A human-readable display name for this Consent Definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Optional parameters for this Consent Definition.
	Parameter []string `json:"parameter,omitempty"`
	// A description for this Consent Definition
	Description *string `json:"description,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewConsentDefinitionResponse instantiates a new ConsentDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentDefinitionResponse(id string, uniqueID string) *ConsentDefinitionResponse {
	this := ConsentDefinitionResponse{}
	this.Id = id
	this.UniqueID = uniqueID
	return &this
}

// NewConsentDefinitionResponseWithDefaults instantiates a new ConsentDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentDefinitionResponseWithDefaults() *ConsentDefinitionResponse {
	this := ConsentDefinitionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ConsentDefinitionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsentDefinitionResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ConsentDefinitionResponse) GetSchemas() []EnumconsentDefinitionSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumconsentDefinitionSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetSchemasOk() ([]EnumconsentDefinitionSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ConsentDefinitionResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumconsentDefinitionSchemaUrn and assigns it to the Schemas field.
func (o *ConsentDefinitionResponse) SetSchemas(v []EnumconsentDefinitionSchemaUrn) {
	o.Schemas = v
}

// GetUniqueID returns the UniqueID field value
func (o *ConsentDefinitionResponse) GetUniqueID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueID
}

// GetUniqueIDOk returns a tuple with the UniqueID field value
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetUniqueIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UniqueID, true
}

// SetUniqueID sets field value
func (o *ConsentDefinitionResponse) SetUniqueID(v string) {
	o.UniqueID = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConsentDefinitionResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConsentDefinitionResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConsentDefinitionResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ConsentDefinitionResponse) GetParameter() []string {
	if o == nil || isNil(o.Parameter) {
		var ret []string
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetParameterOk() ([]string, bool) {
	if o == nil || isNil(o.Parameter) {
    return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ConsentDefinitionResponse) HasParameter() bool {
	if o != nil && !isNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given []string and assigns it to the Parameter field.
func (o *ConsentDefinitionResponse) SetParameter(v []string) {
	o.Parameter = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsentDefinitionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsentDefinitionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsentDefinitionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConsentDefinitionResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentDefinitionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConsentDefinitionResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ConsentDefinitionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o ConsentDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["uniqueID"] = o.UniqueID
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableConsentDefinitionResponse struct {
	value *ConsentDefinitionResponse
	isSet bool
}

func (v NullableConsentDefinitionResponse) Get() *ConsentDefinitionResponse {
	return v.value
}

func (v *NullableConsentDefinitionResponse) Set(val *ConsentDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentDefinitionResponse(val *ConsentDefinitionResponse) *NullableConsentDefinitionResponse {
	return &NullableConsentDefinitionResponse{value: val, isSet: true}
}

func (v NullableConsentDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


