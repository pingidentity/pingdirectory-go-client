# FileBasedCipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Cipher Stream Provider | 
**Schemas** | [**[]EnumfileBasedCipherStreamProviderSchemaUrn**](EnumfileBasedCipherStreamProviderSchemaUrn.md) |  | 
**PasswordFile** | **string** | The path to the file containing the password to use when generating ciphers. | 
**WaitForPasswordFile** | Pointer to **bool** | Indicates whether the server should wait for the password file to become available if it does not exist. | [optional] 
**EncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this File Based Cipher Stream Provider. | [optional] 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewFileBasedCipherStreamProviderResponse

`func NewFileBasedCipherStreamProviderResponse(id string, schemas []EnumfileBasedCipherStreamProviderSchemaUrn, passwordFile string, enabled bool, ) *FileBasedCipherStreamProviderResponse`

NewFileBasedCipherStreamProviderResponse instantiates a new FileBasedCipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBasedCipherStreamProviderResponseWithDefaults

`func NewFileBasedCipherStreamProviderResponseWithDefaults() *FileBasedCipherStreamProviderResponse`

NewFileBasedCipherStreamProviderResponseWithDefaults instantiates a new FileBasedCipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileBasedCipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileBasedCipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileBasedCipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FileBasedCipherStreamProviderResponse) GetSchemas() []EnumfileBasedCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileBasedCipherStreamProviderResponse) GetSchemasOk() (*[]EnumfileBasedCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileBasedCipherStreamProviderResponse) SetSchemas(v []EnumfileBasedCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordFile

`func (o *FileBasedCipherStreamProviderResponse) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *FileBasedCipherStreamProviderResponse) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *FileBasedCipherStreamProviderResponse) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetWaitForPasswordFile

`func (o *FileBasedCipherStreamProviderResponse) GetWaitForPasswordFile() bool`

GetWaitForPasswordFile returns the WaitForPasswordFile field if non-nil, zero value otherwise.

### GetWaitForPasswordFileOk

`func (o *FileBasedCipherStreamProviderResponse) GetWaitForPasswordFileOk() (*bool, bool)`

GetWaitForPasswordFileOk returns a tuple with the WaitForPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPasswordFile

`func (o *FileBasedCipherStreamProviderResponse) SetWaitForPasswordFile(v bool)`

SetWaitForPasswordFile sets WaitForPasswordFile field to given value.

### HasWaitForPasswordFile

`func (o *FileBasedCipherStreamProviderResponse) HasWaitForPasswordFile() bool`

HasWaitForPasswordFile returns a boolean if a field has been set.

### GetEncryptionMetadataFile

`func (o *FileBasedCipherStreamProviderResponse) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *FileBasedCipherStreamProviderResponse) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *FileBasedCipherStreamProviderResponse) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.

### HasEncryptionMetadataFile

`func (o *FileBasedCipherStreamProviderResponse) HasEncryptionMetadataFile() bool`

HasEncryptionMetadataFile returns a boolean if a field has been set.

### GetIterationCount

`func (o *FileBasedCipherStreamProviderResponse) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *FileBasedCipherStreamProviderResponse) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *FileBasedCipherStreamProviderResponse) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *FileBasedCipherStreamProviderResponse) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *FileBasedCipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileBasedCipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileBasedCipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileBasedCipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FileBasedCipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FileBasedCipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FileBasedCipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *FileBasedCipherStreamProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileBasedCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileBasedCipherStreamProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileBasedCipherStreamProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FileBasedCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FileBasedCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FileBasedCipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FileBasedCipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


