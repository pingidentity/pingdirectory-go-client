# VaultPasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvaultPasswordStorageSchemeSchemaUrn**](EnumvaultPasswordStorageSchemeSchemaUrn.md) |  | 
**VaultExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Vault instance containing the passphrase. | 
**DefaultField** | Pointer to **string** | The default name of the field in JSON objects contained in the AWS Secrets Manager service that contains the password for the target user. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewVaultPasswordStorageSchemeShared

`func NewVaultPasswordStorageSchemeShared(schemas []EnumvaultPasswordStorageSchemeSchemaUrn, vaultExternalServer string, enabled bool, ) *VaultPasswordStorageSchemeShared`

NewVaultPasswordStorageSchemeShared instantiates a new VaultPasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultPasswordStorageSchemeSharedWithDefaults

`func NewVaultPasswordStorageSchemeSharedWithDefaults() *VaultPasswordStorageSchemeShared`

NewVaultPasswordStorageSchemeSharedWithDefaults instantiates a new VaultPasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VaultPasswordStorageSchemeShared) GetSchemas() []EnumvaultPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VaultPasswordStorageSchemeShared) GetSchemasOk() (*[]EnumvaultPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VaultPasswordStorageSchemeShared) SetSchemas(v []EnumvaultPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultExternalServer

`func (o *VaultPasswordStorageSchemeShared) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *VaultPasswordStorageSchemeShared) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *VaultPasswordStorageSchemeShared) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.


### GetDefaultField

`func (o *VaultPasswordStorageSchemeShared) GetDefaultField() string`

GetDefaultField returns the DefaultField field if non-nil, zero value otherwise.

### GetDefaultFieldOk

`func (o *VaultPasswordStorageSchemeShared) GetDefaultFieldOk() (*string, bool)`

GetDefaultFieldOk returns a tuple with the DefaultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultField

`func (o *VaultPasswordStorageSchemeShared) SetDefaultField(v string)`

SetDefaultField sets DefaultField field to given value.

### HasDefaultField

`func (o *VaultPasswordStorageSchemeShared) HasDefaultField() bool`

HasDefaultField returns a boolean if a field has been set.

### GetDescription

`func (o *VaultPasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VaultPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VaultPasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VaultPasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VaultPasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VaultPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VaultPasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


