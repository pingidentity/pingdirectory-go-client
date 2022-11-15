# FileBasedPassphraseProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumfileBasedPassphraseProviderSchemaUrn**](EnumfileBasedPassphraseProviderSchemaUrn.md) |  | 
**PasswordFile** | **string** | The path to the file containing the passphrase. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from the target file. A value of zero seconds indicates that the provider should always attempt to read the passphrase from the file. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewFileBasedPassphraseProviderResponse

`func NewFileBasedPassphraseProviderResponse(id string, schemas []EnumfileBasedPassphraseProviderSchemaUrn, passwordFile string, enabled bool, ) *FileBasedPassphraseProviderResponse`

NewFileBasedPassphraseProviderResponse instantiates a new FileBasedPassphraseProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBasedPassphraseProviderResponseWithDefaults

`func NewFileBasedPassphraseProviderResponseWithDefaults() *FileBasedPassphraseProviderResponse`

NewFileBasedPassphraseProviderResponseWithDefaults instantiates a new FileBasedPassphraseProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileBasedPassphraseProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileBasedPassphraseProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileBasedPassphraseProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FileBasedPassphraseProviderResponse) GetSchemas() []EnumfileBasedPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileBasedPassphraseProviderResponse) GetSchemasOk() (*[]EnumfileBasedPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileBasedPassphraseProviderResponse) SetSchemas(v []EnumfileBasedPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordFile

`func (o *FileBasedPassphraseProviderResponse) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *FileBasedPassphraseProviderResponse) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *FileBasedPassphraseProviderResponse) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetMaxCacheDuration

`func (o *FileBasedPassphraseProviderResponse) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *FileBasedPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *FileBasedPassphraseProviderResponse) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *FileBasedPassphraseProviderResponse) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *FileBasedPassphraseProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileBasedPassphraseProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileBasedPassphraseProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileBasedPassphraseProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FileBasedPassphraseProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FileBasedPassphraseProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FileBasedPassphraseProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


