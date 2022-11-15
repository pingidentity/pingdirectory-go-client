# FileBasedPassphraseProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileBasedPassphraseProviderSchemaUrn**](EnumfileBasedPassphraseProviderSchemaUrn.md) |  | 
**PasswordFile** | **string** | The path to the file containing the passphrase. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from the target file. A value of zero seconds indicates that the provider should always attempt to read the passphrase from the file. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewFileBasedPassphraseProviderShared

`func NewFileBasedPassphraseProviderShared(schemas []EnumfileBasedPassphraseProviderSchemaUrn, passwordFile string, enabled bool, ) *FileBasedPassphraseProviderShared`

NewFileBasedPassphraseProviderShared instantiates a new FileBasedPassphraseProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBasedPassphraseProviderSharedWithDefaults

`func NewFileBasedPassphraseProviderSharedWithDefaults() *FileBasedPassphraseProviderShared`

NewFileBasedPassphraseProviderSharedWithDefaults instantiates a new FileBasedPassphraseProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileBasedPassphraseProviderShared) GetSchemas() []EnumfileBasedPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileBasedPassphraseProviderShared) GetSchemasOk() (*[]EnumfileBasedPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileBasedPassphraseProviderShared) SetSchemas(v []EnumfileBasedPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordFile

`func (o *FileBasedPassphraseProviderShared) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *FileBasedPassphraseProviderShared) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *FileBasedPassphraseProviderShared) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetMaxCacheDuration

`func (o *FileBasedPassphraseProviderShared) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *FileBasedPassphraseProviderShared) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *FileBasedPassphraseProviderShared) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *FileBasedPassphraseProviderShared) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *FileBasedPassphraseProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileBasedPassphraseProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileBasedPassphraseProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileBasedPassphraseProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FileBasedPassphraseProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FileBasedPassphraseProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FileBasedPassphraseProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


