# FileBasedCipherStreamProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileBasedCipherStreamProviderSchemaUrn**](EnumfileBasedCipherStreamProviderSchemaUrn.md) |  | 
**PasswordFile** | **string** | The path to the file containing the password to use when generating ciphers. | 
**WaitForPasswordFile** | Pointer to **bool** | Indicates whether the server should wait for the password file to become available if it does not exist. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewFileBasedCipherStreamProviderShared

`func NewFileBasedCipherStreamProviderShared(schemas []EnumfileBasedCipherStreamProviderSchemaUrn, passwordFile string, enabled bool, ) *FileBasedCipherStreamProviderShared`

NewFileBasedCipherStreamProviderShared instantiates a new FileBasedCipherStreamProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBasedCipherStreamProviderSharedWithDefaults

`func NewFileBasedCipherStreamProviderSharedWithDefaults() *FileBasedCipherStreamProviderShared`

NewFileBasedCipherStreamProviderSharedWithDefaults instantiates a new FileBasedCipherStreamProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileBasedCipherStreamProviderShared) GetSchemas() []EnumfileBasedCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileBasedCipherStreamProviderShared) GetSchemasOk() (*[]EnumfileBasedCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileBasedCipherStreamProviderShared) SetSchemas(v []EnumfileBasedCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordFile

`func (o *FileBasedCipherStreamProviderShared) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *FileBasedCipherStreamProviderShared) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *FileBasedCipherStreamProviderShared) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetWaitForPasswordFile

`func (o *FileBasedCipherStreamProviderShared) GetWaitForPasswordFile() bool`

GetWaitForPasswordFile returns the WaitForPasswordFile field if non-nil, zero value otherwise.

### GetWaitForPasswordFileOk

`func (o *FileBasedCipherStreamProviderShared) GetWaitForPasswordFileOk() (*bool, bool)`

GetWaitForPasswordFileOk returns a tuple with the WaitForPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPasswordFile

`func (o *FileBasedCipherStreamProviderShared) SetWaitForPasswordFile(v bool)`

SetWaitForPasswordFile sets WaitForPasswordFile field to given value.

### HasWaitForPasswordFile

`func (o *FileBasedCipherStreamProviderShared) HasWaitForPasswordFile() bool`

HasWaitForPasswordFile returns a boolean if a field has been set.

### GetDescription

`func (o *FileBasedCipherStreamProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileBasedCipherStreamProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileBasedCipherStreamProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileBasedCipherStreamProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FileBasedCipherStreamProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FileBasedCipherStreamProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FileBasedCipherStreamProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


