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

// checks if the AddConsentDefinitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddConsentDefinitionRequest{}

// AddConsentDefinitionRequest struct for AddConsentDefinitionRequest
type AddConsentDefinitionRequest struct {
	Schemas []EnumconsentDefinitionSchemaUrn `json:"schemas,omitempty"`
	// A version-independent unique identifier for this Consent Definition.
	UniqueID string `json:"uniqueID"`
	// A human-readable display name for this Consent Definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Optional parameters for this Consent Definition.
	Parameter []string `json:"parameter,omitempty"`
	// A description for this Consent Definition
	Description *string `json:"description,omitempty"`
	// Name of the new Consent Definition
	DefinitionName string `json:"definitionName"`
}

// NewAddConsentDefinitionRequest instantiates a new AddConsentDefinitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConsentDefinitionRequest(uniqueID string, definitionName string) *AddConsentDefinitionRequest {
	this := AddConsentDefinitionRequest{}
	this.UniqueID = uniqueID
	this.DefinitionName = definitionName
	return &this
}

// NewAddConsentDefinitionRequestWithDefaults instantiates a new AddConsentDefinitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConsentDefinitionRequestWithDefaults() *AddConsentDefinitionRequest {
	this := AddConsentDefinitionRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddConsentDefinitionRequest) GetSchemas() []EnumconsentDefinitionSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumconsentDefinitionSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionRequest) GetSchemasOk() ([]EnumconsentDefinitionSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddConsentDefinitionRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumconsentDefinitionSchemaUrn and assigns it to the Schemas field.
func (o *AddConsentDefinitionRequest) SetSchemas(v []EnumconsentDefinitionSchemaUrn) {
	o.Schemas = v
}

// GetUniqueID returns the UniqueID field value
func (o *AddConsentDefinitionRequest) GetUniqueID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueID
}

// GetUniqueIDOk returns a tuple with the UniqueID field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionRequest) GetUniqueIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueID, true
}

// SetUniqueID sets field value
func (o *AddConsentDefinitionRequest) SetUniqueID(v string) {
	o.UniqueID = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AddConsentDefinitionRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AddConsentDefinitionRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AddConsentDefinitionRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *AddConsentDefinitionRequest) GetParameter() []string {
	if o == nil || IsNil(o.Parameter) {
		var ret []string
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionRequest) GetParameterOk() ([]string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *AddConsentDefinitionRequest) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given []string and assigns it to the Parameter field.
func (o *AddConsentDefinitionRequest) SetParameter(v []string) {
	o.Parameter = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddConsentDefinitionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddConsentDefinitionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddConsentDefinitionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDefinitionName returns the DefinitionName field value
func (o *AddConsentDefinitionRequest) GetDefinitionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefinitionName
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionRequest) GetDefinitionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionName, true
}

// SetDefinitionName sets field value
func (o *AddConsentDefinitionRequest) SetDefinitionName(v string) {
	o.DefinitionName = v
}

func (o AddConsentDefinitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddConsentDefinitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["uniqueID"] = o.UniqueID
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["definitionName"] = o.DefinitionName
	return toSerialize, nil
}

type NullableAddConsentDefinitionRequest struct {
	value *AddConsentDefinitionRequest
	isSet bool
}

func (v NullableAddConsentDefinitionRequest) Get() *AddConsentDefinitionRequest {
	return v.value
}

func (v *NullableAddConsentDefinitionRequest) Set(val *AddConsentDefinitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConsentDefinitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConsentDefinitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConsentDefinitionRequest(val *AddConsentDefinitionRequest) *NullableAddConsentDefinitionRequest {
	return &NullableAddConsentDefinitionRequest{value: val, isSet: true}
}

func (v NullableAddConsentDefinitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConsentDefinitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
