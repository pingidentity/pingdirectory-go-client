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

// checks if the AmazonKeyManagementServiceCipherStreamProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonKeyManagementServiceCipherStreamProviderResponse{}

// AmazonKeyManagementServiceCipherStreamProviderResponse struct for AmazonKeyManagementServiceCipherStreamProviderResponse
type AmazonKeyManagementServiceCipherStreamProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Cipher Stream Provider
	Id      string                                                        `json:"id"`
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

// NewAmazonKeyManagementServiceCipherStreamProviderResponse instantiates a new AmazonKeyManagementServiceCipherStreamProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonKeyManagementServiceCipherStreamProviderResponse(id string, schemas []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, encryptedPassphraseFile string, kmsEncryptionKeyArn string, enabled bool) *AmazonKeyManagementServiceCipherStreamProviderResponse {
	this := AmazonKeyManagementServiceCipherStreamProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.EncryptedPassphraseFile = encryptedPassphraseFile
	this.KmsEncryptionKeyArn = kmsEncryptionKeyArn
	this.Enabled = enabled
	return &this
}

// NewAmazonKeyManagementServiceCipherStreamProviderResponseWithDefaults instantiates a new AmazonKeyManagementServiceCipherStreamProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonKeyManagementServiceCipherStreamProviderResponseWithDefaults() *AmazonKeyManagementServiceCipherStreamProviderResponse {
	this := AmazonKeyManagementServiceCipherStreamProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetSchemas() []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetSchemasOk() ([]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetSchemas(v []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEncryptedPassphraseFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedPassphraseFile
}

// GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field value
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEncryptedPassphraseFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedPassphraseFile, true
}

// SetEncryptedPassphraseFile sets field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetEncryptedPassphraseFile(v string) {
	o.EncryptedPassphraseFile = v
}

// GetAwsExternalServer returns the AwsExternalServer field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsExternalServer() string {
	if o == nil || IsNil(o.AwsExternalServer) {
		var ret string
		return ret
	}
	return *o.AwsExternalServer
}

// GetAwsExternalServerOk returns a tuple with the AwsExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.AwsExternalServer) {
		return nil, false
	}
	return o.AwsExternalServer, true
}

// HasAwsExternalServer returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsExternalServer() bool {
	if o != nil && !IsNil(o.AwsExternalServer) {
		return true
	}

	return false
}

// SetAwsExternalServer gets a reference to the given string and assigns it to the AwsExternalServer field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsExternalServer(v string) {
	o.AwsExternalServer = &v
}

// GetAwsAccessKeyID returns the AwsAccessKeyID field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsAccessKeyID() string {
	if o == nil || IsNil(o.AwsAccessKeyID) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyID
}

// GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsAccessKeyIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccessKeyID) {
		return nil, false
	}
	return o.AwsAccessKeyID, true
}

// HasAwsAccessKeyID returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsAccessKeyID() bool {
	if o != nil && !IsNil(o.AwsAccessKeyID) {
		return true
	}

	return false
}

// SetAwsAccessKeyID gets a reference to the given string and assigns it to the AwsAccessKeyID field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsAccessKeyID(v string) {
	o.AwsAccessKeyID = &v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsSecretAccessKey() string {
	if o == nil || IsNil(o.AwsSecretAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsSecretAccessKey) {
		return nil, false
	}
	return o.AwsSecretAccessKey, true
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsSecretAccessKey() bool {
	if o != nil && !IsNil(o.AwsSecretAccessKey) {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given string and assigns it to the AwsSecretAccessKey field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = &v
}

// GetAwsRegionName returns the AwsRegionName field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsRegionName() string {
	if o == nil || IsNil(o.AwsRegionName) {
		var ret string
		return ret
	}
	return *o.AwsRegionName
}

// GetAwsRegionNameOk returns a tuple with the AwsRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.AwsRegionName) {
		return nil, false
	}
	return o.AwsRegionName, true
}

// HasAwsRegionName returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsRegionName() bool {
	if o != nil && !IsNil(o.AwsRegionName) {
		return true
	}

	return false
}

// SetAwsRegionName gets a reference to the given string and assigns it to the AwsRegionName field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsRegionName(v string) {
	o.AwsRegionName = &v
}

// GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetKmsEncryptionKeyArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KmsEncryptionKeyArn
}

// GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field value
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetKmsEncryptionKeyArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KmsEncryptionKeyArn, true
}

// SetKmsEncryptionKeyArn sets field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetKmsEncryptionKeyArn(v string) {
	o.KmsEncryptionKeyArn = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AmazonKeyManagementServiceCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmazonKeyManagementServiceCipherStreamProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["encryptedPassphraseFile"] = o.EncryptedPassphraseFile
	if !IsNil(o.AwsExternalServer) {
		toSerialize["awsExternalServer"] = o.AwsExternalServer
	}
	if !IsNil(o.AwsAccessKeyID) {
		toSerialize["awsAccessKeyID"] = o.AwsAccessKeyID
	}
	if !IsNil(o.AwsSecretAccessKey) {
		toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey
	}
	if !IsNil(o.AwsRegionName) {
		toSerialize["awsRegionName"] = o.AwsRegionName
	}
	toSerialize["kmsEncryptionKeyArn"] = o.KmsEncryptionKeyArn
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAmazonKeyManagementServiceCipherStreamProviderResponse struct {
	value *AmazonKeyManagementServiceCipherStreamProviderResponse
	isSet bool
}

func (v NullableAmazonKeyManagementServiceCipherStreamProviderResponse) Get() *AmazonKeyManagementServiceCipherStreamProviderResponse {
	return v.value
}

func (v *NullableAmazonKeyManagementServiceCipherStreamProviderResponse) Set(val *AmazonKeyManagementServiceCipherStreamProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonKeyManagementServiceCipherStreamProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonKeyManagementServiceCipherStreamProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonKeyManagementServiceCipherStreamProviderResponse(val *AmazonKeyManagementServiceCipherStreamProviderResponse) *NullableAmazonKeyManagementServiceCipherStreamProviderResponse {
	return &NullableAmazonKeyManagementServiceCipherStreamProviderResponse{value: val, isSet: true}
}

func (v NullableAmazonKeyManagementServiceCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonKeyManagementServiceCipherStreamProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
