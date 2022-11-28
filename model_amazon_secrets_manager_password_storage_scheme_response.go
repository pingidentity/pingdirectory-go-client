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

// AmazonSecretsManagerPasswordStorageSchemeResponse struct for AmazonSecretsManagerPasswordStorageSchemeResponse
type AmazonSecretsManagerPasswordStorageSchemeResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
	Schemas []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The external server with information to use when interacting with the AWS Secrets Manager service.
	AwsExternalServer string `json:"awsExternalServer"`
	// The default name of the field in JSON objects contained in the AWS Secrets Manager service that contains the password for the target user.
	DefaultField *string `json:"defaultField,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAmazonSecretsManagerPasswordStorageSchemeResponse instantiates a new AmazonSecretsManagerPasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSecretsManagerPasswordStorageSchemeResponse(id string, schemas []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, awsExternalServer string, enabled bool) *AmazonSecretsManagerPasswordStorageSchemeResponse {
	this := AmazonSecretsManagerPasswordStorageSchemeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AwsExternalServer = awsExternalServer
	this.Enabled = enabled
	return &this
}

// NewAmazonSecretsManagerPasswordStorageSchemeResponseWithDefaults instantiates a new AmazonSecretsManagerPasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSecretsManagerPasswordStorageSchemeResponseWithDefaults() *AmazonSecretsManagerPasswordStorageSchemeResponse {
	this := AmazonSecretsManagerPasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetSchemas() []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetSchemasOk() ([]EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetSchemas(v []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetAwsExternalServer returns the AwsExternalServer field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetAwsExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsExternalServer
}

// GetAwsExternalServerOk returns a tuple with the AwsExternalServer field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetAwsExternalServerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AwsExternalServer, true
}

// SetAwsExternalServer sets field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetAwsExternalServer(v string) {
	o.AwsExternalServer = v
}

// GetDefaultField returns the DefaultField field value if set, zero value otherwise.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetDefaultField() string {
	if o == nil || isNil(o.DefaultField) {
		var ret string
		return ret
	}
	return *o.DefaultField
}

// GetDefaultFieldOk returns a tuple with the DefaultField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetDefaultFieldOk() (*string, bool) {
	if o == nil || isNil(o.DefaultField) {
    return nil, false
	}
	return o.DefaultField, true
}

// HasDefaultField returns a boolean if a field has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) HasDefaultField() bool {
	if o != nil && !isNil(o.DefaultField) {
		return true
	}

	return false
}

// SetDefaultField gets a reference to the given string and assigns it to the DefaultField field.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetDefaultField(v string) {
	o.DefaultField = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AmazonSecretsManagerPasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AmazonSecretsManagerPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
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

type NullableAmazonSecretsManagerPasswordStorageSchemeResponse struct {
	value *AmazonSecretsManagerPasswordStorageSchemeResponse
	isSet bool
}

func (v NullableAmazonSecretsManagerPasswordStorageSchemeResponse) Get() *AmazonSecretsManagerPasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableAmazonSecretsManagerPasswordStorageSchemeResponse) Set(val *AmazonSecretsManagerPasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSecretsManagerPasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSecretsManagerPasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSecretsManagerPasswordStorageSchemeResponse(val *AmazonSecretsManagerPasswordStorageSchemeResponse) *NullableAmazonSecretsManagerPasswordStorageSchemeResponse {
	return &NullableAmazonSecretsManagerPasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableAmazonSecretsManagerPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSecretsManagerPasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


