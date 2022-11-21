# AmazonKeyManagementServiceCipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Cipher Stream Provider | 
**Schemas** | [**[]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn**](EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn.md) |  | 
**EncryptedPassphraseFile** | **string** | The path to a file that will hold the encrypted passphrase used by this cipher stream provider. | 
**AwsExternalServer** | Pointer to **string** | The external server with information to use when interacting with the Amazon Key Management Service. | [optional] 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsRegionName** | Pointer to **string** | The name of the Amazon Web Services region that holds the encryption key. This is optional, and if it is not provided, then the server will attempt to determine the region from the key ARN. | [optional] 
**KmsEncryptionKeyArn** | **string** | The Amazon resource name (ARN) for the KMS key that will be used to encrypt the contents of the passphrase file. This key must exist, and the AWS client must have access to encrypt and decrypt data using this key. | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewAmazonKeyManagementServiceCipherStreamProviderResponse

`func NewAmazonKeyManagementServiceCipherStreamProviderResponse(id string, schemas []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, encryptedPassphraseFile string, kmsEncryptionKeyArn string, enabled bool, ) *AmazonKeyManagementServiceCipherStreamProviderResponse`

NewAmazonKeyManagementServiceCipherStreamProviderResponse instantiates a new AmazonKeyManagementServiceCipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonKeyManagementServiceCipherStreamProviderResponseWithDefaults

`func NewAmazonKeyManagementServiceCipherStreamProviderResponseWithDefaults() *AmazonKeyManagementServiceCipherStreamProviderResponse`

NewAmazonKeyManagementServiceCipherStreamProviderResponseWithDefaults instantiates a new AmazonKeyManagementServiceCipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetSchemas() []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetSchemasOk() (*[]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetSchemas(v []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEncryptedPassphraseFile

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEncryptedPassphraseFile() string`

GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptedPassphraseFileOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEncryptedPassphraseFileOk() (*string, bool)`

GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassphraseFile

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetEncryptedPassphraseFile(v string)`

SetEncryptedPassphraseFile sets EncryptedPassphraseFile field to given value.


### GetAwsExternalServer

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.

### HasAwsExternalServer

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsExternalServer() bool`

HasAwsExternalServer returns a boolean if a field has been set.

### GetAwsAccessKeyID

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.

### HasAwsRegionName

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasAwsRegionName() bool`

HasAwsRegionName returns a boolean if a field has been set.

### GetKmsEncryptionKeyArn

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetKmsEncryptionKeyArn() string`

GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field if non-nil, zero value otherwise.

### GetKmsEncryptionKeyArnOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetKmsEncryptionKeyArnOk() (*string, bool)`

GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsEncryptionKeyArn

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetKmsEncryptionKeyArn(v string)`

SetKmsEncryptionKeyArn sets KmsEncryptionKeyArn field to given value.


### GetDescription

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AmazonKeyManagementServiceCipherStreamProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


