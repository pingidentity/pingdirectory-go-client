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

// checks if the AddAmazonSecretsManagerCipherStreamProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAmazonSecretsManagerCipherStreamProviderRequest{}

// AddAmazonSecretsManagerCipherStreamProviderRequest struct for AddAmazonSecretsManagerCipherStreamProviderRequest
type AddAmazonSecretsManagerCipherStreamProviderRequest struct {
	// Name of the new Cipher Stream Provider
	ProviderName string                                                  `json:"providerName"`
	Schemas      []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn `json:"schemas"`
	// The external server with information to use when interacting with the AWS Secrets Manager.
	AwsExternalServer string `json:"awsExternalServer"`
	// The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved.
	SecretID string `json:"secretID"`
	// The name of the JSON field whose value is the passphrase that will be used to generate the encryption key for protecting the contents of the encryption settings database.
	SecretFieldName string `json:"secretFieldName"`
	// The unique identifier for the version of the secret to be retrieved.
	SecretVersionID *string `json:"secretVersionID,omitempty"`
	// The staging label for the version of the secret to be retrieved.
	SecretVersionStage *string `json:"secretVersionStage,omitempty"`
	// The path to a file that will hold metadata about the encryption performed by this Amazon Secrets Manager Cipher Stream Provider.
	EncryptionMetadataFile *string `json:"encryptionMetadataFile,omitempty"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddAmazonSecretsManagerCipherStreamProviderRequest instantiates a new AddAmazonSecretsManagerCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAmazonSecretsManagerCipherStreamProviderRequest(providerName string, schemas []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, awsExternalServer string, secretID string, secretFieldName string, enabled bool) *AddAmazonSecretsManagerCipherStreamProviderRequest {
	this := AddAmazonSecretsManagerCipherStreamProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.AwsExternalServer = awsExternalServer
	this.SecretID = secretID
	this.SecretFieldName = secretFieldName
	this.Enabled = enabled
	return &this
}

// NewAddAmazonSecretsManagerCipherStreamProviderRequestWithDefaults instantiates a new AddAmazonSecretsManagerCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAmazonSecretsManagerCipherStreamProviderRequestWithDefaults() *AddAmazonSecretsManagerCipherStreamProviderRequest {
	this := AddAmazonSecretsManagerCipherStreamProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSchemas() []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSchemasOk() ([]EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSchemas(v []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetAwsExternalServer returns the AwsExternalServer field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetAwsExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsExternalServer
}

// GetAwsExternalServerOk returns a tuple with the AwsExternalServer field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetAwsExternalServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsExternalServer, true
}

// SetAwsExternalServer sets field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetAwsExternalServer(v string) {
	o.AwsExternalServer = v
}

// GetSecretID returns the SecretID field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretID
}

// GetSecretIDOk returns a tuple with the SecretID field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretID, true
}

// SetSecretID sets field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretID(v string) {
	o.SecretID = v
}

// GetSecretFieldName returns the SecretFieldName field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretFieldName
}

// GetSecretFieldNameOk returns a tuple with the SecretFieldName field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretFieldName, true
}

// SetSecretFieldName sets field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretFieldName(v string) {
	o.SecretFieldName = v
}

// GetSecretVersionID returns the SecretVersionID field value if set, zero value otherwise.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionID() string {
	if o == nil || IsNil(o.SecretVersionID) {
		var ret string
		return ret
	}
	return *o.SecretVersionID
}

// GetSecretVersionIDOk returns a tuple with the SecretVersionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionIDOk() (*string, bool) {
	if o == nil || IsNil(o.SecretVersionID) {
		return nil, false
	}
	return o.SecretVersionID, true
}

// HasSecretVersionID returns a boolean if a field has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasSecretVersionID() bool {
	if o != nil && !IsNil(o.SecretVersionID) {
		return true
	}

	return false
}

// SetSecretVersionID gets a reference to the given string and assigns it to the SecretVersionID field.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretVersionID(v string) {
	o.SecretVersionID = &v
}

// GetSecretVersionStage returns the SecretVersionStage field value if set, zero value otherwise.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionStage() string {
	if o == nil || IsNil(o.SecretVersionStage) {
		var ret string
		return ret
	}
	return *o.SecretVersionStage
}

// GetSecretVersionStageOk returns a tuple with the SecretVersionStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionStageOk() (*string, bool) {
	if o == nil || IsNil(o.SecretVersionStage) {
		return nil, false
	}
	return o.SecretVersionStage, true
}

// HasSecretVersionStage returns a boolean if a field has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasSecretVersionStage() bool {
	if o != nil && !IsNil(o.SecretVersionStage) {
		return true
	}

	return false
}

// SetSecretVersionStage gets a reference to the given string and assigns it to the SecretVersionStage field.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretVersionStage(v string) {
	o.SecretVersionStage = &v
}

// GetEncryptionMetadataFile returns the EncryptionMetadataFile field value if set, zero value otherwise.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEncryptionMetadataFile() string {
	if o == nil || IsNil(o.EncryptionMetadataFile) {
		var ret string
		return ret
	}
	return *o.EncryptionMetadataFile
}

// GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionMetadataFile) {
		return nil, false
	}
	return o.EncryptionMetadataFile, true
}

// HasEncryptionMetadataFile returns a boolean if a field has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasEncryptionMetadataFile() bool {
	if o != nil && !IsNil(o.EncryptionMetadataFile) {
		return true
	}

	return false
}

// SetEncryptionMetadataFile gets a reference to the given string and assigns it to the EncryptionMetadataFile field.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetEncryptionMetadataFile(v string) {
	o.EncryptionMetadataFile = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddAmazonSecretsManagerCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAmazonSecretsManagerCipherStreamProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerName"] = o.ProviderName
	toSerialize["schemas"] = o.Schemas
	toSerialize["awsExternalServer"] = o.AwsExternalServer
	toSerialize["secretID"] = o.SecretID
	toSerialize["secretFieldName"] = o.SecretFieldName
	if !IsNil(o.SecretVersionID) {
		toSerialize["secretVersionID"] = o.SecretVersionID
	}
	if !IsNil(o.SecretVersionStage) {
		toSerialize["secretVersionStage"] = o.SecretVersionStage
	}
	if !IsNil(o.EncryptionMetadataFile) {
		toSerialize["encryptionMetadataFile"] = o.EncryptionMetadataFile
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddAmazonSecretsManagerCipherStreamProviderRequest struct {
	value *AddAmazonSecretsManagerCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddAmazonSecretsManagerCipherStreamProviderRequest) Get() *AddAmazonSecretsManagerCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddAmazonSecretsManagerCipherStreamProviderRequest) Set(val *AddAmazonSecretsManagerCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAmazonSecretsManagerCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAmazonSecretsManagerCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAmazonSecretsManagerCipherStreamProviderRequest(val *AddAmazonSecretsManagerCipherStreamProviderRequest) *NullableAddAmazonSecretsManagerCipherStreamProviderRequest {
	return &NullableAddAmazonSecretsManagerCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddAmazonSecretsManagerCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAmazonSecretsManagerCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
