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

// GeneratePasswordExtendedOperationHandlerResponse struct for GeneratePasswordExtendedOperationHandlerResponse
type GeneratePasswordExtendedOperationHandlerResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	// The default password policy that should be used when generating and validating passwords if the request does not specify an alternate policy. If this is not provided, then this Generate Password Extended Operation Handler will use the default password policy defined in the global configuration.
	DefaultPasswordPolicy *string `json:"defaultPasswordPolicy,omitempty"`
	// The default password generator that will be used if the selected password policy is not configured with a password generator.
	DefaultPasswordGenerator string `json:"defaultPasswordGenerator"`
	// The maximum number of passwords that may be generated and returned to the client for a single request.
	MaximumPasswordsPerRequest *int32 `json:"maximumPasswordsPerRequest,omitempty"`
	// The maximum number of attempts that the server may use to generate a password that passes validation.
	MaximumValidationAttemptsPerPassword *int32 `json:"maximumValidationAttemptsPerPassword,omitempty"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewGeneratePasswordExtendedOperationHandlerResponse instantiates a new GeneratePasswordExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePasswordExtendedOperationHandlerResponse(schemas []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn, id string, defaultPasswordGenerator string, enabled bool) *GeneratePasswordExtendedOperationHandlerResponse {
	this := GeneratePasswordExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.DefaultPasswordGenerator = defaultPasswordGenerator
	this.Enabled = enabled
	return &this
}

// NewGeneratePasswordExtendedOperationHandlerResponseWithDefaults instantiates a new GeneratePasswordExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePasswordExtendedOperationHandlerResponseWithDefaults() *GeneratePasswordExtendedOperationHandlerResponse {
	this := GeneratePasswordExtendedOperationHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetSchemas() []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumgeneratePasswordExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetSchemas(v []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetDefaultPasswordPolicy returns the DefaultPasswordPolicy field value if set, zero value otherwise.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordPolicy() string {
	if o == nil || isNil(o.DefaultPasswordPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPasswordPolicy
}

// GetDefaultPasswordPolicyOk returns a tuple with the DefaultPasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordPolicyOk() (*string, bool) {
	if o == nil || isNil(o.DefaultPasswordPolicy) {
    return nil, false
	}
	return o.DefaultPasswordPolicy, true
}

// HasDefaultPasswordPolicy returns a boolean if a field has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) HasDefaultPasswordPolicy() bool {
	if o != nil && !isNil(o.DefaultPasswordPolicy) {
		return true
	}

	return false
}

// SetDefaultPasswordPolicy gets a reference to the given string and assigns it to the DefaultPasswordPolicy field.
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetDefaultPasswordPolicy(v string) {
	o.DefaultPasswordPolicy = &v
}

// GetDefaultPasswordGenerator returns the DefaultPasswordGenerator field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordGenerator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultPasswordGenerator
}

// GetDefaultPasswordGeneratorOk returns a tuple with the DefaultPasswordGenerator field value
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordGeneratorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultPasswordGenerator, true
}

// SetDefaultPasswordGenerator sets field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetDefaultPasswordGenerator(v string) {
	o.DefaultPasswordGenerator = v
}

// GetMaximumPasswordsPerRequest returns the MaximumPasswordsPerRequest field value if set, zero value otherwise.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumPasswordsPerRequest() int32 {
	if o == nil || isNil(o.MaximumPasswordsPerRequest) {
		var ret int32
		return ret
	}
	return *o.MaximumPasswordsPerRequest
}

// GetMaximumPasswordsPerRequestOk returns a tuple with the MaximumPasswordsPerRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumPasswordsPerRequestOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumPasswordsPerRequest) {
    return nil, false
	}
	return o.MaximumPasswordsPerRequest, true
}

// HasMaximumPasswordsPerRequest returns a boolean if a field has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) HasMaximumPasswordsPerRequest() bool {
	if o != nil && !isNil(o.MaximumPasswordsPerRequest) {
		return true
	}

	return false
}

// SetMaximumPasswordsPerRequest gets a reference to the given int32 and assigns it to the MaximumPasswordsPerRequest field.
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetMaximumPasswordsPerRequest(v int32) {
	o.MaximumPasswordsPerRequest = &v
}

// GetMaximumValidationAttemptsPerPassword returns the MaximumValidationAttemptsPerPassword field value if set, zero value otherwise.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumValidationAttemptsPerPassword() int32 {
	if o == nil || isNil(o.MaximumValidationAttemptsPerPassword) {
		var ret int32
		return ret
	}
	return *o.MaximumValidationAttemptsPerPassword
}

// GetMaximumValidationAttemptsPerPasswordOk returns a tuple with the MaximumValidationAttemptsPerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumValidationAttemptsPerPasswordOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumValidationAttemptsPerPassword) {
    return nil, false
	}
	return o.MaximumValidationAttemptsPerPassword, true
}

// HasMaximumValidationAttemptsPerPassword returns a boolean if a field has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) HasMaximumValidationAttemptsPerPassword() bool {
	if o != nil && !isNil(o.MaximumValidationAttemptsPerPassword) {
		return true
	}

	return false
}

// SetMaximumValidationAttemptsPerPassword gets a reference to the given int32 and assigns it to the MaximumValidationAttemptsPerPassword field.
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetMaximumValidationAttemptsPerPassword(v int32) {
	o.MaximumValidationAttemptsPerPassword = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GeneratePasswordExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GeneratePasswordExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GeneratePasswordExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.DefaultPasswordPolicy) {
		toSerialize["defaultPasswordPolicy"] = o.DefaultPasswordPolicy
	}
	if true {
		toSerialize["defaultPasswordGenerator"] = o.DefaultPasswordGenerator
	}
	if !isNil(o.MaximumPasswordsPerRequest) {
		toSerialize["maximumPasswordsPerRequest"] = o.MaximumPasswordsPerRequest
	}
	if !isNil(o.MaximumValidationAttemptsPerPassword) {
		toSerialize["maximumValidationAttemptsPerPassword"] = o.MaximumValidationAttemptsPerPassword
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableGeneratePasswordExtendedOperationHandlerResponse struct {
	value *GeneratePasswordExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableGeneratePasswordExtendedOperationHandlerResponse) Get() *GeneratePasswordExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableGeneratePasswordExtendedOperationHandlerResponse) Set(val *GeneratePasswordExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePasswordExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePasswordExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePasswordExtendedOperationHandlerResponse(val *GeneratePasswordExtendedOperationHandlerResponse) *NullableGeneratePasswordExtendedOperationHandlerResponse {
	return &NullableGeneratePasswordExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableGeneratePasswordExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratePasswordExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


