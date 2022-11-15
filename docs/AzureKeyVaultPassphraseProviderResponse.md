# AzureKeyVaultPassphraseProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumazureKeyVaultPassphraseProviderSchemaUrn**](EnumazureKeyVaultPassphraseProviderSchemaUrn.md) |  | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**SecretName** | **string** | The name of the secret to retrieve. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Azure Key Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from the Azure service. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewAzureKeyVaultPassphraseProviderResponse

`func NewAzureKeyVaultPassphraseProviderResponse(id string, schemas []EnumazureKeyVaultPassphraseProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, enabled bool, ) *AzureKeyVaultPassphraseProviderResponse`

NewAzureKeyVaultPassphraseProviderResponse instantiates a new AzureKeyVaultPassphraseProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultPassphraseProviderResponseWithDefaults

`func NewAzureKeyVaultPassphraseProviderResponseWithDefaults() *AzureKeyVaultPassphraseProviderResponse`

NewAzureKeyVaultPassphraseProviderResponseWithDefaults instantiates a new AzureKeyVaultPassphraseProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureKeyVaultPassphraseProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureKeyVaultPassphraseProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AzureKeyVaultPassphraseProviderResponse) GetSchemas() []EnumazureKeyVaultPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetSchemasOk() (*[]EnumazureKeyVaultPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AzureKeyVaultPassphraseProviderResponse) SetSchemas(v []EnumazureKeyVaultPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetKeyVaultURI

`func (o *AzureKeyVaultPassphraseProviderResponse) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AzureKeyVaultPassphraseProviderResponse) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AzureKeyVaultPassphraseProviderResponse) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AzureKeyVaultPassphraseProviderResponse) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetSecretName

`func (o *AzureKeyVaultPassphraseProviderResponse) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AzureKeyVaultPassphraseProviderResponse) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetMaxCacheDuration

`func (o *AzureKeyVaultPassphraseProviderResponse) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *AzureKeyVaultPassphraseProviderResponse) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *AzureKeyVaultPassphraseProviderResponse) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AzureKeyVaultPassphraseProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureKeyVaultPassphraseProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureKeyVaultPassphraseProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AzureKeyVaultPassphraseProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureKeyVaultPassphraseProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureKeyVaultPassphraseProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


