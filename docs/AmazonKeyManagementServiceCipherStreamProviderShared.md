# AmazonKeyManagementServiceCipherStreamProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn**](EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn.md) |  | 
**EncryptedPassphraseFile** | **string** | The path to a file that will hold the encrypted passphrase used by this cipher stream provider. | 
**AwsExternalServer** | Pointer to **string** | The external server with information to use when interacting with the Amazon Key Management Service. | [optional] 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsRegionName** | Pointer to **string** | The name of the Amazon Web Services region that holds the encryption key. This is optional, and if it is not provided, then the server will attempt to determine the region from the key ARN. | [optional] 
**KmsEncryptionKeyArn** | **string** | The Amazon resource name (ARN) for the KMS key that will be used to encrypt the contents of the passphrase file. This key must exist, and the AWS client must have access to encrypt and decrypt data using this key. | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAmazonKeyManagementServiceCipherStreamProviderShared

`func NewAmazonKeyManagementServiceCipherStreamProviderShared(schemas []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, encryptedPassphraseFile string, kmsEncryptionKeyArn string, enabled bool, ) *AmazonKeyManagementServiceCipherStreamProviderShared`

NewAmazonKeyManagementServiceCipherStreamProviderShared instantiates a new AmazonKeyManagementServiceCipherStreamProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonKeyManagementServiceCipherStreamProviderSharedWithDefaults

`func NewAmazonKeyManagementServiceCipherStreamProviderSharedWithDefaults() *AmazonKeyManagementServiceCipherStreamProviderShared`

NewAmazonKeyManagementServiceCipherStreamProviderSharedWithDefaults instantiates a new AmazonKeyManagementServiceCipherStreamProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetSchemas() []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetSchemasOk() (*[]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetSchemas(v []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEncryptedPassphraseFile

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetEncryptedPassphraseFile() string`

GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptedPassphraseFileOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetEncryptedPassphraseFileOk() (*string, bool)`

GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassphraseFile

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetEncryptedPassphraseFile(v string)`

SetEncryptedPassphraseFile sets EncryptedPassphraseFile field to given value.


### GetAwsExternalServer

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.

### HasAwsExternalServer

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) HasAwsExternalServer() bool`

HasAwsExternalServer returns a boolean if a field has been set.

### GetAwsAccessKeyID

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.

### HasAwsRegionName

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) HasAwsRegionName() bool`

HasAwsRegionName returns a boolean if a field has been set.

### GetKmsEncryptionKeyArn

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetKmsEncryptionKeyArn() string`

GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field if non-nil, zero value otherwise.

### GetKmsEncryptionKeyArnOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetKmsEncryptionKeyArnOk() (*string, bool)`

GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsEncryptionKeyArn

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetKmsEncryptionKeyArn(v string)`

SetKmsEncryptionKeyArn sets KmsEncryptionKeyArn field to given value.


### GetDescription

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmazonKeyManagementServiceCipherStreamProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


