# VaultPassphraseProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumvaultPassphraseProviderSchemaUrn**](EnumvaultPassphraseProviderSchemaUrn.md) |  | 
**VaultExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Vault instance containing the passphrase. | 
**VaultSecretPath** | **string** | The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server. | 
**VaultSecretFieldName** | **string** | The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Vault. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewVaultPassphraseProviderResponse

`func NewVaultPassphraseProviderResponse(id string, schemas []EnumvaultPassphraseProviderSchemaUrn, vaultExternalServer string, vaultSecretPath string, vaultSecretFieldName string, enabled bool, ) *VaultPassphraseProviderResponse`

NewVaultPassphraseProviderResponse instantiates a new VaultPassphraseProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultPassphraseProviderResponseWithDefaults

`func NewVaultPassphraseProviderResponseWithDefaults() *VaultPassphraseProviderResponse`

NewVaultPassphraseProviderResponseWithDefaults instantiates a new VaultPassphraseProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VaultPassphraseProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultPassphraseProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultPassphraseProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *VaultPassphraseProviderResponse) GetSchemas() []EnumvaultPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VaultPassphraseProviderResponse) GetSchemasOk() (*[]EnumvaultPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VaultPassphraseProviderResponse) SetSchemas(v []EnumvaultPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultExternalServer

`func (o *VaultPassphraseProviderResponse) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *VaultPassphraseProviderResponse) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *VaultPassphraseProviderResponse) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.


### GetVaultSecretPath

`func (o *VaultPassphraseProviderResponse) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *VaultPassphraseProviderResponse) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *VaultPassphraseProviderResponse) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *VaultPassphraseProviderResponse) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *VaultPassphraseProviderResponse) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *VaultPassphraseProviderResponse) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetMaxCacheDuration

`func (o *VaultPassphraseProviderResponse) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *VaultPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *VaultPassphraseProviderResponse) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *VaultPassphraseProviderResponse) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *VaultPassphraseProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VaultPassphraseProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VaultPassphraseProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VaultPassphraseProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VaultPassphraseProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VaultPassphraseProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VaultPassphraseProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *VaultPassphraseProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VaultPassphraseProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VaultPassphraseProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VaultPassphraseProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


