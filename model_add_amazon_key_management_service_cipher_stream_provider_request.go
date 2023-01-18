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

// AddAmazonKeyManagementServiceCipherStreamProviderRequest struct for AddAmazonKeyManagementServiceCipherStreamProviderRequest
type AddAmazonKeyManagementServiceCipherStreamProviderRequest struct {
	// Name of the new Cipher Stream Provider
	ProviderName string `json:"providerName"`
	Schemas []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn `json:"schemas"`
	// The path to a file that will hold the encrypted passphrase used by this cipher stream provider.
	EncryptedPassphraseFile string `json:"encryptedPassphraseFile"`
	// The external server with information to use when interacting with the Amazon Key Management Service.
	AwsExternalServer *string `json:"awsExternalServer,omitempty"`
	// The access key ID that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance.
	AwsAccessKeyID *string `json:"awsAccessKeyID,omitempty"`
	// The secret access key that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance.
	AwsSecretAccessKey *string `json:"awsSecretAccessKey,omitempty"`
	// The name of the Amazon Web Services region that holds the encryption key. This is optional, and if it is not provided, then the server will attempt to determine the region from the key ARN.
	AwsRegionName *string `json:"awsRegionName,omitempty"`
	// The Amazon resource name (ARN) for the KMS key that will be used to encrypt the contents of the passphrase file. This key must exist, and the AWS client must have access to encrypt and decrypt data using this key.
	KmsEncryptionKeyArn string `json:"kmsEncryptionKeyArn"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddAmazonKeyManagementServiceCipherStreamProviderRequest instantiates a new AddAmazonKeyManagementServiceCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAmazonKeyManagementServiceCipherStreamProviderRequest(providerName string, schemas []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, encryptedPassphraseFile string, kmsEncryptionKeyArn string, enabled bool) *AddAmazonKeyManagementServiceCipherStreamProviderRequest {
	this := AddAmazonKeyManagementServiceCipherStreamProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.EncryptedPassphraseFile = encryptedPassphraseFile
	this.KmsEncryptionKeyArn = kmsEncryptionKeyArn
	this.Enabled = enabled
	return &this
}

// NewAddAmazonKeyManagementServiceCipherStreamProviderRequestWithDefaults instantiates a new AddAmazonKeyManagementServiceCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAmazonKeyManagementServiceCipherStreamProviderRequestWithDefaults() *AddAmazonKeyManagementServiceCipherStreamProviderRequest {
	this := AddAmazonKeyManagementServiceCipherStreamProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetSchemas() []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetSchemasOk() ([]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetSchemas(v []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEncryptedPassphraseFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedPassphraseFile
}

// GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field value
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEncryptedPassphraseFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EncryptedPassphraseFile, true
}

// SetEncryptedPassphraseFile sets field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetEncryptedPassphraseFile(v string) {
	o.EncryptedPassphraseFile = v
}

// GetAwsExternalServer returns the AwsExternalServer field value if set, zero value otherwise.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsExternalServer() string {
	if o == nil || isNil(o.AwsExternalServer) {
		var ret string
		return ret
	}
	return *o.AwsExternalServer
}

// GetAwsExternalServerOk returns a tuple with the AwsExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsExternalServerOk() (*string, bool) {
	if o == nil || isNil(o.AwsExternalServer) {
    return nil, false
	}
	return o.AwsExternalServer, true
}

// HasAwsExternalServer returns a boolean if a field has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsExternalServer() bool {
	if o != nil && !isNil(o.AwsExternalServer) {
		return true
	}

	return false
}

// SetAwsExternalServer gets a reference to the given string and assigns it to the AwsExternalServer field.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsExternalServer(v string) {
	o.AwsExternalServer = &v
}

// GetAwsAccessKeyID returns the AwsAccessKeyID field value if set, zero value otherwise.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsAccessKeyID() string {
	if o == nil || isNil(o.AwsAccessKeyID) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyID
}

// GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsAccessKeyIDOk() (*string, bool) {
	if o == nil || isNil(o.AwsAccessKeyID) {
    return nil, false
	}
	return o.AwsAccessKeyID, true
}

// HasAwsAccessKeyID returns a boolean if a field has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsAccessKeyID() bool {
	if o != nil && !isNil(o.AwsAccessKeyID) {
		return true
	}

	return false
}

// SetAwsAccessKeyID gets a reference to the given string and assigns it to the AwsAccessKeyID field.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsAccessKeyID(v string) {
	o.AwsAccessKeyID = &v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsSecretAccessKey() string {
	if o == nil || isNil(o.AwsSecretAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil || isNil(o.AwsSecretAccessKey) {
    return nil, false
	}
	return o.AwsSecretAccessKey, true
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsSecretAccessKey() bool {
	if o != nil && !isNil(o.AwsSecretAccessKey) {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given string and assigns it to the AwsSecretAccessKey field.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = &v
}

// GetAwsRegionName returns the AwsRegionName field value if set, zero value otherwise.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsRegionName() string {
	if o == nil || isNil(o.AwsRegionName) {
		var ret string
		return ret
	}
	return *o.AwsRegionName
}

// GetAwsRegionNameOk returns a tuple with the AwsRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsRegionNameOk() (*string, bool) {
	if o == nil || isNil(o.AwsRegionName) {
    return nil, false
	}
	return o.AwsRegionName, true
}

// HasAwsRegionName returns a boolean if a field has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsRegionName() bool {
	if o != nil && !isNil(o.AwsRegionName) {
		return true
	}

	return false
}

// SetAwsRegionName gets a reference to the given string and assigns it to the AwsRegionName field.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsRegionName(v string) {
	o.AwsRegionName = &v
}

// GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetKmsEncryptionKeyArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KmsEncryptionKeyArn
}

// GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field value
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetKmsEncryptionKeyArnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.KmsEncryptionKeyArn, true
}

// SetKmsEncryptionKeyArn sets field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetKmsEncryptionKeyArn(v string) {
	o.KmsEncryptionKeyArn = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddAmazonKeyManagementServiceCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["encryptedPassphraseFile"] = o.EncryptedPassphraseFile
	}
	if !isNil(o.AwsExternalServer) {
		toSerialize["awsExternalServer"] = o.AwsExternalServer
	}
	if !isNil(o.AwsAccessKeyID) {
		toSerialize["awsAccessKeyID"] = o.AwsAccessKeyID
	}
	if !isNil(o.AwsSecretAccessKey) {
		toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey
	}
	if !isNil(o.AwsRegionName) {
		toSerialize["awsRegionName"] = o.AwsRegionName
	}
	if true {
		toSerialize["kmsEncryptionKeyArn"] = o.KmsEncryptionKeyArn
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest struct {
	value *AddAmazonKeyManagementServiceCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest) Get() *AddAmazonKeyManagementServiceCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest) Set(val *AddAmazonKeyManagementServiceCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAmazonKeyManagementServiceCipherStreamProviderRequest(val *AddAmazonKeyManagementServiceCipherStreamProviderRequest) *NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest {
	return &NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAmazonKeyManagementServiceCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

