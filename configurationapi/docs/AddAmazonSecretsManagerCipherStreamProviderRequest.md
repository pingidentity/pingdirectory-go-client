# AddAmazonSecretsManagerCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]EnumamazonSecretsManagerCipherStreamProviderSchemaUrn**](EnumamazonSecretsManagerCipherStreamProviderSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be used to generate the encryption key for protecting the contents of the encryption settings database. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**EncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this Amazon Secrets Manager Cipher Stream Provider. | [optional] 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAddAmazonSecretsManagerCipherStreamProviderRequest

`func NewAddAmazonSecretsManagerCipherStreamProviderRequest(providerName string, schemas []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, awsExternalServer string, secretID string, secretFieldName string, enabled bool, ) *AddAmazonSecretsManagerCipherStreamProviderRequest`

NewAddAmazonSecretsManagerCipherStreamProviderRequest instantiates a new AddAmazonSecretsManagerCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAmazonSecretsManagerCipherStreamProviderRequestWithDefaults

`func NewAddAmazonSecretsManagerCipherStreamProviderRequestWithDefaults() *AddAmazonSecretsManagerCipherStreamProviderRequest`

NewAddAmazonSecretsManagerCipherStreamProviderRequestWithDefaults instantiates a new AddAmazonSecretsManagerCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSchemas() []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSchemasOk() (*[]EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSchemas(v []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetSecretID

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetEncryptionMetadataFile

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.

### HasEncryptionMetadataFile

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasEncryptionMetadataFile() bool`

HasEncryptionMetadataFile returns a boolean if a field has been set.

### GetIterationCount

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAmazonSecretsManagerCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


