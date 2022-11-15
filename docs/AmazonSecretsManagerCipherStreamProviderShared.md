# AmazonSecretsManagerCipherStreamProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonSecretsManagerCipherStreamProviderSchemaUrn**](EnumamazonSecretsManagerCipherStreamProviderSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be used to generate the encryption key for protecting the contents of the encryption settings database. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**EncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this Amazon Secrets Manager Cipher Stream Provider. | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAmazonSecretsManagerCipherStreamProviderShared

`func NewAmazonSecretsManagerCipherStreamProviderShared(schemas []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, awsExternalServer string, secretID string, secretFieldName string, encryptionMetadataFile string, enabled bool, ) *AmazonSecretsManagerCipherStreamProviderShared`

NewAmazonSecretsManagerCipherStreamProviderShared instantiates a new AmazonSecretsManagerCipherStreamProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSecretsManagerCipherStreamProviderSharedWithDefaults

`func NewAmazonSecretsManagerCipherStreamProviderSharedWithDefaults() *AmazonSecretsManagerCipherStreamProviderShared`

NewAmazonSecretsManagerCipherStreamProviderSharedWithDefaults instantiates a new AmazonSecretsManagerCipherStreamProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSchemas() []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSchemasOk() (*[]EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetSchemas(v []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetSecretID

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *AmazonSecretsManagerCipherStreamProviderShared) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *AmazonSecretsManagerCipherStreamProviderShared) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetEncryptionMetadataFile

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.


### GetDescription

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonSecretsManagerCipherStreamProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmazonSecretsManagerCipherStreamProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmazonSecretsManagerCipherStreamProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


