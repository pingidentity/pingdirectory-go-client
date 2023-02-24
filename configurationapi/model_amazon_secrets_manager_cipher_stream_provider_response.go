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

// checks if the AmazonSecretsManagerCipherStreamProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonSecretsManagerCipherStreamProviderResponse{}

// AmazonSecretsManagerCipherStreamProviderResponse struct for AmazonSecretsManagerCipherStreamProviderResponse
type AmazonSecretsManagerCipherStreamProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Cipher Stream Provider
	Id      string                                                  `json:"id"`
	Schemas []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn `json:"schemas"`
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
	EncryptionMetadataFile string `json:"encryptionMetadataFile"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAmazonSecretsManagerCipherStreamProviderResponse instantiates a new AmazonSecretsManagerCipherStreamProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSecretsManagerCipherStreamProviderResponse(id string, schemas []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, awsExternalServer string, secretID string, secretFieldName string, encryptionMetadataFile string, enabled bool) *AmazonSecretsManagerCipherStreamProviderResponse {
	this := AmazonSecretsManagerCipherStreamProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AwsExternalServer = awsExternalServer
	this.SecretID = secretID
	this.SecretFieldName = secretFieldName
	this.EncryptionMetadataFile = encryptionMetadataFile
	this.Enabled = enabled
	return &this
}

// NewAmazonSecretsManagerCipherStreamProviderResponseWithDefaults instantiates a new AmazonSecretsManagerCipherStreamProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSecretsManagerCipherStreamProviderResponseWithDefaults() *AmazonSecretsManagerCipherStreamProviderResponse {
	this := AmazonSecretsManagerCipherStreamProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSchemas() []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSchemasOk() ([]EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetSchemas(v []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetAwsExternalServer returns the AwsExternalServer field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetAwsExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsExternalServer
}

// GetAwsExternalServerOk returns a tuple with the AwsExternalServer field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetAwsExternalServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsExternalServer, true
}

// SetAwsExternalServer sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetAwsExternalServer(v string) {
	o.AwsExternalServer = v
}

// GetSecretID returns the SecretID field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretID
}

// GetSecretIDOk returns a tuple with the SecretID field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretID, true
}

// SetSecretID sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetSecretID(v string) {
	o.SecretID = v
}

// GetSecretFieldName returns the SecretFieldName field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretFieldName
}

// GetSecretFieldNameOk returns a tuple with the SecretFieldName field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretFieldName, true
}

// SetSecretFieldName sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetSecretFieldName(v string) {
	o.SecretFieldName = v
}

// GetSecretVersionID returns the SecretVersionID field value if set, zero value otherwise.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretVersionID() string {
	if o == nil || IsNil(o.SecretVersionID) {
		var ret string
		return ret
	}
	return *o.SecretVersionID
}

// GetSecretVersionIDOk returns a tuple with the SecretVersionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretVersionIDOk() (*string, bool) {
	if o == nil || IsNil(o.SecretVersionID) {
		return nil, false
	}
	return o.SecretVersionID, true
}

// HasSecretVersionID returns a boolean if a field has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) HasSecretVersionID() bool {
	if o != nil && !IsNil(o.SecretVersionID) {
		return true
	}

	return false
}

// SetSecretVersionID gets a reference to the given string and assigns it to the SecretVersionID field.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetSecretVersionID(v string) {
	o.SecretVersionID = &v
}

// GetSecretVersionStage returns the SecretVersionStage field value if set, zero value otherwise.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretVersionStage() string {
	if o == nil || IsNil(o.SecretVersionStage) {
		var ret string
		return ret
	}
	return *o.SecretVersionStage
}

// GetSecretVersionStageOk returns a tuple with the SecretVersionStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetSecretVersionStageOk() (*string, bool) {
	if o == nil || IsNil(o.SecretVersionStage) {
		return nil, false
	}
	return o.SecretVersionStage, true
}

// HasSecretVersionStage returns a boolean if a field has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) HasSecretVersionStage() bool {
	if o != nil && !IsNil(o.SecretVersionStage) {
		return true
	}

	return false
}

// SetSecretVersionStage gets a reference to the given string and assigns it to the SecretVersionStage field.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetSecretVersionStage(v string) {
	o.SecretVersionStage = &v
}

// GetEncryptionMetadataFile returns the EncryptionMetadataFile field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetEncryptionMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptionMetadataFile
}

// GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetEncryptionMetadataFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionMetadataFile, true
}

// SetEncryptionMetadataFile sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetEncryptionMetadataFile(v string) {
	o.EncryptionMetadataFile = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AmazonSecretsManagerCipherStreamProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AmazonSecretsManagerCipherStreamProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AmazonSecretsManagerCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmazonSecretsManagerCipherStreamProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
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
	toSerialize["encryptionMetadataFile"] = o.EncryptionMetadataFile
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAmazonSecretsManagerCipherStreamProviderResponse struct {
	value *AmazonSecretsManagerCipherStreamProviderResponse
	isSet bool
}

func (v NullableAmazonSecretsManagerCipherStreamProviderResponse) Get() *AmazonSecretsManagerCipherStreamProviderResponse {
	return v.value
}

func (v *NullableAmazonSecretsManagerCipherStreamProviderResponse) Set(val *AmazonSecretsManagerCipherStreamProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSecretsManagerCipherStreamProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSecretsManagerCipherStreamProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSecretsManagerCipherStreamProviderResponse(val *AmazonSecretsManagerCipherStreamProviderResponse) *NullableAmazonSecretsManagerCipherStreamProviderResponse {
	return &NullableAmazonSecretsManagerCipherStreamProviderResponse{value: val, isSet: true}
}

func (v NullableAmazonSecretsManagerCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSecretsManagerCipherStreamProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
