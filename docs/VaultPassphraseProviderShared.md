# VaultPassphraseProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvaultPassphraseProviderSchemaUrn**](EnumvaultPassphraseProviderSchemaUrn.md) |  | 
**VaultExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Vault instance containing the passphrase. | 
**VaultSecretPath** | **string** | The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server. | 
**VaultSecretFieldName** | **string** | The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Vault. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewVaultPassphraseProviderShared

`func NewVaultPassphraseProviderShared(schemas []EnumvaultPassphraseProviderSchemaUrn, vaultExternalServer string, vaultSecretPath string, vaultSecretFieldName string, enabled bool, ) *VaultPassphraseProviderShared`

NewVaultPassphraseProviderShared instantiates a new VaultPassphraseProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultPassphraseProviderSharedWithDefaults

`func NewVaultPassphraseProviderSharedWithDefaults() *VaultPassphraseProviderShared`

NewVaultPassphraseProviderSharedWithDefaults instantiates a new VaultPassphraseProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VaultPassphraseProviderShared) GetSchemas() []EnumvaultPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VaultPassphraseProviderShared) GetSchemasOk() (*[]EnumvaultPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VaultPassphraseProviderShared) SetSchemas(v []EnumvaultPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultExternalServer

`func (o *VaultPassphraseProviderShared) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *VaultPassphraseProviderShared) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *VaultPassphraseProviderShared) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.


### GetVaultSecretPath

`func (o *VaultPassphraseProviderShared) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *VaultPassphraseProviderShared) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *VaultPassphraseProviderShared) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *VaultPassphraseProviderShared) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *VaultPassphraseProviderShared) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *VaultPassphraseProviderShared) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetMaxCacheDuration

`func (o *VaultPassphraseProviderShared) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *VaultPassphraseProviderShared) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *VaultPassphraseProviderShared) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *VaultPassphraseProviderShared) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *VaultPassphraseProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VaultPassphraseProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VaultPassphraseProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VaultPassphraseProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VaultPassphraseProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VaultPassphraseProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VaultPassphraseProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


