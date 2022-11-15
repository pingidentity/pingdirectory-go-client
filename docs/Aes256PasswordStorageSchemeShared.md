# Aes256PasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumaes256PasswordStorageSchemeSchemaUrn**](Enumaes256PasswordStorageSchemeSchemaUrn.md) |  | 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The identifier for the encryption settings definition that should be used to derive the encryption key to use when encrypting new passwords. If this is not provided, the server&#39;s preferred encryption settings definition will be used. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewAes256PasswordStorageSchemeShared

`func NewAes256PasswordStorageSchemeShared(schemas []Enumaes256PasswordStorageSchemeSchemaUrn, enabled bool, ) *Aes256PasswordStorageSchemeShared`

NewAes256PasswordStorageSchemeShared instantiates a new Aes256PasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAes256PasswordStorageSchemeSharedWithDefaults

`func NewAes256PasswordStorageSchemeSharedWithDefaults() *Aes256PasswordStorageSchemeShared`

NewAes256PasswordStorageSchemeSharedWithDefaults instantiates a new Aes256PasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Aes256PasswordStorageSchemeShared) GetSchemas() []Enumaes256PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Aes256PasswordStorageSchemeShared) GetSchemasOk() (*[]Enumaes256PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Aes256PasswordStorageSchemeShared) SetSchemas(v []Enumaes256PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEncryptionSettingsDefinitionID

`func (o *Aes256PasswordStorageSchemeShared) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *Aes256PasswordStorageSchemeShared) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *Aes256PasswordStorageSchemeShared) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *Aes256PasswordStorageSchemeShared) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetDescription

`func (o *Aes256PasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Aes256PasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Aes256PasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Aes256PasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Aes256PasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Aes256PasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Aes256PasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


