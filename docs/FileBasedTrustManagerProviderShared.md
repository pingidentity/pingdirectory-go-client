# FileBasedTrustManagerProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileBasedTrustManagerProviderSchemaUrn**](EnumfileBasedTrustManagerProviderSchemaUrn.md) |  | 
**TrustStoreFile** | **string** | Specifies the path to the file containing the trust information. It can be an absolute path or a path that is relative to the Directory Server instance root. | 
**TrustStoreType** | Pointer to **string** | Specifies the format for the data in the trust store file. | [optional] 
**TrustStorePin** | Pointer to **string** | Specifies the clear-text PIN needed to access the File Based Trust Manager Provider. | [optional] 
**TrustStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Trust Manager Provider. | [optional] 
**TrustStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Trust Manager Provider. | [optional] 
**Enabled** | **bool** | Indicate whether the Trust Manager Provider is enabled for use. | 
**IncludeJVMDefaultIssuers** | Pointer to **bool** | Indicates whether certificates issued by an authority included in the JVM&#39;s set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider. | [optional] 

## Methods

### NewFileBasedTrustManagerProviderShared

`func NewFileBasedTrustManagerProviderShared(schemas []EnumfileBasedTrustManagerProviderSchemaUrn, trustStoreFile string, enabled bool, ) *FileBasedTrustManagerProviderShared`

NewFileBasedTrustManagerProviderShared instantiates a new FileBasedTrustManagerProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileBasedTrustManagerProviderSharedWithDefaults

`func NewFileBasedTrustManagerProviderSharedWithDefaults() *FileBasedTrustManagerProviderShared`

NewFileBasedTrustManagerProviderSharedWithDefaults instantiates a new FileBasedTrustManagerProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileBasedTrustManagerProviderShared) GetSchemas() []EnumfileBasedTrustManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileBasedTrustManagerProviderShared) GetSchemasOk() (*[]EnumfileBasedTrustManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileBasedTrustManagerProviderShared) SetSchemas(v []EnumfileBasedTrustManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTrustStoreFile

`func (o *FileBasedTrustManagerProviderShared) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *FileBasedTrustManagerProviderShared) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *FileBasedTrustManagerProviderShared) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.


### GetTrustStoreType

`func (o *FileBasedTrustManagerProviderShared) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *FileBasedTrustManagerProviderShared) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *FileBasedTrustManagerProviderShared) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *FileBasedTrustManagerProviderShared) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *FileBasedTrustManagerProviderShared) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *FileBasedTrustManagerProviderShared) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *FileBasedTrustManagerProviderShared) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *FileBasedTrustManagerProviderShared) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStorePinFile

`func (o *FileBasedTrustManagerProviderShared) GetTrustStorePinFile() string`

GetTrustStorePinFile returns the TrustStorePinFile field if non-nil, zero value otherwise.

### GetTrustStorePinFileOk

`func (o *FileBasedTrustManagerProviderShared) GetTrustStorePinFileOk() (*string, bool)`

GetTrustStorePinFileOk returns a tuple with the TrustStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinFile

`func (o *FileBasedTrustManagerProviderShared) SetTrustStorePinFile(v string)`

SetTrustStorePinFile sets TrustStorePinFile field to given value.

### HasTrustStorePinFile

`func (o *FileBasedTrustManagerProviderShared) HasTrustStorePinFile() bool`

HasTrustStorePinFile returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *FileBasedTrustManagerProviderShared) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *FileBasedTrustManagerProviderShared) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *FileBasedTrustManagerProviderShared) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *FileBasedTrustManagerProviderShared) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetEnabled

`func (o *FileBasedTrustManagerProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FileBasedTrustManagerProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FileBasedTrustManagerProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeJVMDefaultIssuers

`func (o *FileBasedTrustManagerProviderShared) GetIncludeJVMDefaultIssuers() bool`

GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field if non-nil, zero value otherwise.

### GetIncludeJVMDefaultIssuersOk

`func (o *FileBasedTrustManagerProviderShared) GetIncludeJVMDefaultIssuersOk() (*bool, bool)`

GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJVMDefaultIssuers

`func (o *FileBasedTrustManagerProviderShared) SetIncludeJVMDefaultIssuers(v bool)`

SetIncludeJVMDefaultIssuers sets IncludeJVMDefaultIssuers field to given value.

### HasIncludeJVMDefaultIssuers

`func (o *FileBasedTrustManagerProviderShared) HasIncludeJVMDefaultIssuers() bool`

HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


