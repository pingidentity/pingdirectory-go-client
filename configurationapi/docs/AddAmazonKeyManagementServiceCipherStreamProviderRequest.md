# AddAmazonKeyManagementServiceCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn**](EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn.md) |  | 
**EncryptedPassphraseFile** | Pointer to **string** | The path to a file that will hold the encrypted passphrase used by this cipher stream provider. | [optional] 
**AwsExternalServer** | Pointer to **string** | The external server with information to use when interacting with the Amazon Key Management Service. | [optional] 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsRegionName** | Pointer to **string** | The name of the Amazon Web Services region that holds the encryption key. This is optional, and if it is not provided, then the server will attempt to determine the region from the key ARN. | [optional] 
**KmsEncryptionKeyArn** | **string** | The Amazon resource name (ARN) for the KMS key that will be used to encrypt the contents of the passphrase file. This key must exist, and the AWS client must have access to encrypt and decrypt data using this key. | 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAddAmazonKeyManagementServiceCipherStreamProviderRequest

`func NewAddAmazonKeyManagementServiceCipherStreamProviderRequest(providerName string, schemas []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, kmsEncryptionKeyArn string, enabled bool, ) *AddAmazonKeyManagementServiceCipherStreamProviderRequest`

NewAddAmazonKeyManagementServiceCipherStreamProviderRequest instantiates a new AddAmazonKeyManagementServiceCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAmazonKeyManagementServiceCipherStreamProviderRequestWithDefaults

`func NewAddAmazonKeyManagementServiceCipherStreamProviderRequestWithDefaults() *AddAmazonKeyManagementServiceCipherStreamProviderRequest`

NewAddAmazonKeyManagementServiceCipherStreamProviderRequestWithDefaults instantiates a new AddAmazonKeyManagementServiceCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetSchemas() []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetSchemasOk() (*[]EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetSchemas(v []EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEncryptedPassphraseFile

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEncryptedPassphraseFile() string`

GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptedPassphraseFileOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEncryptedPassphraseFileOk() (*string, bool)`

GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassphraseFile

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetEncryptedPassphraseFile(v string)`

SetEncryptedPassphraseFile sets EncryptedPassphraseFile field to given value.

### HasEncryptedPassphraseFile

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasEncryptedPassphraseFile() bool`

HasEncryptedPassphraseFile returns a boolean if a field has been set.

### GetAwsExternalServer

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.

### HasAwsExternalServer

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsExternalServer() bool`

HasAwsExternalServer returns a boolean if a field has been set.

### GetAwsAccessKeyID

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.

### HasAwsRegionName

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasAwsRegionName() bool`

HasAwsRegionName returns a boolean if a field has been set.

### GetKmsEncryptionKeyArn

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetKmsEncryptionKeyArn() string`

GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field if non-nil, zero value otherwise.

### GetKmsEncryptionKeyArnOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetKmsEncryptionKeyArnOk() (*string, bool)`

GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsEncryptionKeyArn

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetKmsEncryptionKeyArn(v string)`

SetKmsEncryptionKeyArn sets KmsEncryptionKeyArn field to given value.


### GetIterationCount

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAmazonKeyManagementServiceCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


