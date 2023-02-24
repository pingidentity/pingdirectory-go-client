# AddAzureKeyVaultCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]EnumazureKeyVaultCipherStreamProviderSchemaUrn**](EnumazureKeyVaultCipherStreamProviderSchemaUrn.md) |  | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**SecretName** | **string** | The name of the secret to retrieve. | 
**EncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this Azure Key Vault Cipher Stream Provider. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAddAzureKeyVaultCipherStreamProviderRequest

`func NewAddAzureKeyVaultCipherStreamProviderRequest(providerName string, schemas []EnumazureKeyVaultCipherStreamProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, enabled bool, ) *AddAzureKeyVaultCipherStreamProviderRequest`

NewAddAzureKeyVaultCipherStreamProviderRequest instantiates a new AddAzureKeyVaultCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAzureKeyVaultCipherStreamProviderRequestWithDefaults

`func NewAddAzureKeyVaultCipherStreamProviderRequestWithDefaults() *AddAzureKeyVaultCipherStreamProviderRequest`

NewAddAzureKeyVaultCipherStreamProviderRequestWithDefaults instantiates a new AddAzureKeyVaultCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSchemas() []EnumazureKeyVaultCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSchemasOk() (*[]EnumazureKeyVaultCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetSchemas(v []EnumazureKeyVaultCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetKeyVaultURI

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetSecretName

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetEncryptionMetadataFile

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.

### HasEncryptionMetadataFile

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) HasEncryptionMetadataFile() bool`

HasEncryptionMetadataFile returns a boolean if a field has been set.

### GetDescription

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


