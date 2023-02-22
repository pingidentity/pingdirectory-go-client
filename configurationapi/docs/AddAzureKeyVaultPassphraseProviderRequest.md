# AddAzureKeyVaultPassphraseProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Passphrase Provider | 
**Schemas** | [**[]EnumazureKeyVaultPassphraseProviderSchemaUrn**](EnumazureKeyVaultPassphraseProviderSchemaUrn.md) |  | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**SecretName** | **string** | The name of the secret to retrieve. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Azure Key Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from the Azure service. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewAddAzureKeyVaultPassphraseProviderRequest

`func NewAddAzureKeyVaultPassphraseProviderRequest(providerName string, schemas []EnumazureKeyVaultPassphraseProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, enabled bool, ) *AddAzureKeyVaultPassphraseProviderRequest`

NewAddAzureKeyVaultPassphraseProviderRequest instantiates a new AddAzureKeyVaultPassphraseProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAzureKeyVaultPassphraseProviderRequestWithDefaults

`func NewAddAzureKeyVaultPassphraseProviderRequestWithDefaults() *AddAzureKeyVaultPassphraseProviderRequest`

NewAddAzureKeyVaultPassphraseProviderRequestWithDefaults instantiates a new AddAzureKeyVaultPassphraseProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetSchemas() []EnumazureKeyVaultPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetSchemasOk() (*[]EnumazureKeyVaultPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetSchemas(v []EnumazureKeyVaultPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetKeyVaultURI

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetSecretName

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetMaxCacheDuration

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *AddAzureKeyVaultPassphraseProviderRequest) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAzureKeyVaultPassphraseProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAzureKeyVaultPassphraseProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAzureKeyVaultPassphraseProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


