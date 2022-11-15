# AzureKeyVaultCipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Cipher Stream Provider | 
**Schemas** | [**[]EnumazureKeyVaultCipherStreamProviderSchemaUrn**](EnumazureKeyVaultCipherStreamProviderSchemaUrn.md) |  | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**SecretName** | **string** | The name of the secret to retrieve. | 
**EncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this Azure Key Vault Cipher Stream Provider. | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAzureKeyVaultCipherStreamProviderResponse

`func NewAzureKeyVaultCipherStreamProviderResponse(id string, schemas []EnumazureKeyVaultCipherStreamProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, encryptionMetadataFile string, enabled bool, ) *AzureKeyVaultCipherStreamProviderResponse`

NewAzureKeyVaultCipherStreamProviderResponse instantiates a new AzureKeyVaultCipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultCipherStreamProviderResponseWithDefaults

`func NewAzureKeyVaultCipherStreamProviderResponseWithDefaults() *AzureKeyVaultCipherStreamProviderResponse`

NewAzureKeyVaultCipherStreamProviderResponseWithDefaults instantiates a new AzureKeyVaultCipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetSchemas() []EnumazureKeyVaultCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetSchemasOk() (*[]EnumazureKeyVaultCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetSchemas(v []EnumazureKeyVaultCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetKeyVaultURI

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetSecretName

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetEncryptionMetadataFile

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.


### GetDescription

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureKeyVaultCipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureKeyVaultCipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureKeyVaultCipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


