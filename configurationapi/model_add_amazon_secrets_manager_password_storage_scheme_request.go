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

// AddAmazonSecretsManagerPasswordStorageSchemeRequest struct for AddAmazonSecretsManagerPasswordStorageSchemeRequest
type AddAmazonSecretsManagerPasswordStorageSchemeRequest struct {
	// Name of the new Password Storage Scheme
	SchemeName string                                                   `json:"schemeName"`
	Schemas    []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The external server with information to use when interacting with the AWS Secrets Manager service.
	AwsExternalServer string `json:"awsExternalServer"`
	// The default name of the field in JSON objects contained in the AWS Secrets Manager service that contains the password for the target user.
	DefaultField *string `json:"defaultField,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddAmazonSecretsManagerPasswordStorageSchemeRequest instantiates a new AddAmazonSecretsManagerPasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAmazonSecretsManagerPasswordStorageSchemeRequest(schemeName string, schemas []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, awsExternalServer string, enabled bool) *AddAmazonSecretsManagerPasswordStorageSchemeRequest {
	this := AddAmazonSecretsManagerPasswordStorageSchemeRequest{}
	this.SchemeName = schemeName
	this.Schemas = schemas
	this.AwsExternalServer = awsExternalServer
	this.Enabled = enabled
	return &this
}

// NewAddAmazonSecretsManagerPasswordStorageSchemeRequestWithDefaults instantiates a new AddAmazonSecretsManagerPasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAmazonSecretsManagerPasswordStorageSchemeRequestWithDefaults() *AddAmazonSecretsManagerPasswordStorageSchemeRequest {
	this := AddAmazonSecretsManagerPasswordStorageSchemeRequest{}
	return &this
}

// GetSchemeName returns the SchemeName field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemas() []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemasOk() ([]EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetSchemas(v []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetAwsExternalServer returns the AwsExternalServer field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetAwsExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsExternalServer
}

// GetAwsExternalServerOk returns a tuple with the AwsExternalServer field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetAwsExternalServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsExternalServer, true
}

// SetAwsExternalServer sets field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetAwsExternalServer(v string) {
	o.AwsExternalServer = v
}

// GetDefaultField returns the DefaultField field value if set, zero value otherwise.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDefaultField() string {
	if o == nil || isNil(o.DefaultField) {
		var ret string
		return ret
	}
	return *o.DefaultField
}

// GetDefaultFieldOk returns a tuple with the DefaultField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDefaultFieldOk() (*string, bool) {
	if o == nil || isNil(o.DefaultField) {
		return nil, false
	}
	return o.DefaultField, true
}

// HasDefaultField returns a boolean if a field has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) HasDefaultField() bool {
	if o != nil && !isNil(o.DefaultField) {
		return true
	}

	return false
}

// SetDefaultField gets a reference to the given string and assigns it to the DefaultField field.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetDefaultField(v string) {
	o.DefaultField = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddAmazonSecretsManagerPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemeName"] = o.SchemeName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["awsExternalServer"] = o.AwsExternalServer
	}
	if !isNil(o.DefaultField) {
		toSerialize["defaultField"] = o.DefaultField
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest struct {
	value *AddAmazonSecretsManagerPasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest) Get() *AddAmazonSecretsManagerPasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest) Set(val *AddAmazonSecretsManagerPasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAmazonSecretsManagerPasswordStorageSchemeRequest(val *AddAmazonSecretsManagerPasswordStorageSchemeRequest) *NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest {
	return &NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAmazonSecretsManagerPasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}