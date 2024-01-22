# AddFileBasedKeyManagerProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileBasedKeyManagerProviderSchemaUrn**](EnumfileBasedKeyManagerProviderSchemaUrn.md) |  | 
**KeyStoreFile** | **string** | Specifies the path to the file that contains the private key information. This may be an absolute path, or a path that is relative to the Directory Server instance root. | 
**KeyStoreType** | Pointer to **string** | Specifies the format for the data in the key store file. | [optional] 
**KeyStorePin** | Pointer to **string** | Specifies the PIN needed to access the File Based Key Manager Provider. | [optional] 
**KeyStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Key Manager Provider. | [optional] 
**KeyStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Key Manager Provider. | [optional] 
**PrivateKeyPin** | Pointer to **string** | Specifies the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN. | [optional] 
**PrivateKeyPinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN. | [optional] 
**PrivateKeyPinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN. | [optional] 
**Description** | Pointer to **string** | A description for this Key Manager Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Key Manager Provider is enabled for use. | 
**ProviderName** | **string** | Name of the new Key Manager Provider | 

## Methods

### NewAddFileBasedKeyManagerProviderRequest

`func NewAddFileBasedKeyManagerProviderRequest(schemas []EnumfileBasedKeyManagerProviderSchemaUrn, keyStoreFile string, enabled bool, providerName string, ) *AddFileBasedKeyManagerProviderRequest`

NewAddFileBasedKeyManagerProviderRequest instantiates a new AddFileBasedKeyManagerProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFileBasedKeyManagerProviderRequestWithDefaults

`func NewAddFileBasedKeyManagerProviderRequestWithDefaults() *AddFileBasedKeyManagerProviderRequest`

NewAddFileBasedKeyManagerProviderRequestWithDefaults instantiates a new AddFileBasedKeyManagerProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddFileBasedKeyManagerProviderRequest) GetSchemas() []EnumfileBasedKeyManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetSchemasOk() (*[]EnumfileBasedKeyManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFileBasedKeyManagerProviderRequest) SetSchemas(v []EnumfileBasedKeyManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetKeyStoreFile

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStoreFile() string`

GetKeyStoreFile returns the KeyStoreFile field if non-nil, zero value otherwise.

### GetKeyStoreFileOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStoreFileOk() (*string, bool)`

GetKeyStoreFileOk returns a tuple with the KeyStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStoreFile

`func (o *AddFileBasedKeyManagerProviderRequest) SetKeyStoreFile(v string)`

SetKeyStoreFile sets KeyStoreFile field to given value.


### GetKeyStoreType

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStoreType() string`

GetKeyStoreType returns the KeyStoreType field if non-nil, zero value otherwise.

### GetKeyStoreTypeOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStoreTypeOk() (*string, bool)`

GetKeyStoreTypeOk returns a tuple with the KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStoreType

`func (o *AddFileBasedKeyManagerProviderRequest) SetKeyStoreType(v string)`

SetKeyStoreType sets KeyStoreType field to given value.

### HasKeyStoreType

`func (o *AddFileBasedKeyManagerProviderRequest) HasKeyStoreType() bool`

HasKeyStoreType returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *AddFileBasedKeyManagerProviderRequest) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *AddFileBasedKeyManagerProviderRequest) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *AddFileBasedKeyManagerProviderRequest) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *AddFileBasedKeyManagerProviderRequest) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinPassphraseProvider

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStorePinPassphraseProvider() string`

GetKeyStorePinPassphraseProvider returns the KeyStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetKeyStorePinPassphraseProviderOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetKeyStorePinPassphraseProviderOk() (*string, bool)`

GetKeyStorePinPassphraseProviderOk returns a tuple with the KeyStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinPassphraseProvider

`func (o *AddFileBasedKeyManagerProviderRequest) SetKeyStorePinPassphraseProvider(v string)`

SetKeyStorePinPassphraseProvider sets KeyStorePinPassphraseProvider field to given value.

### HasKeyStorePinPassphraseProvider

`func (o *AddFileBasedKeyManagerProviderRequest) HasKeyStorePinPassphraseProvider() bool`

HasKeyStorePinPassphraseProvider returns a boolean if a field has been set.

### GetPrivateKeyPin

`func (o *AddFileBasedKeyManagerProviderRequest) GetPrivateKeyPin() string`

GetPrivateKeyPin returns the PrivateKeyPin field if non-nil, zero value otherwise.

### GetPrivateKeyPinOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetPrivateKeyPinOk() (*string, bool)`

GetPrivateKeyPinOk returns a tuple with the PrivateKeyPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPin

`func (o *AddFileBasedKeyManagerProviderRequest) SetPrivateKeyPin(v string)`

SetPrivateKeyPin sets PrivateKeyPin field to given value.

### HasPrivateKeyPin

`func (o *AddFileBasedKeyManagerProviderRequest) HasPrivateKeyPin() bool`

HasPrivateKeyPin returns a boolean if a field has been set.

### GetPrivateKeyPinFile

`func (o *AddFileBasedKeyManagerProviderRequest) GetPrivateKeyPinFile() string`

GetPrivateKeyPinFile returns the PrivateKeyPinFile field if non-nil, zero value otherwise.

### GetPrivateKeyPinFileOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetPrivateKeyPinFileOk() (*string, bool)`

GetPrivateKeyPinFileOk returns a tuple with the PrivateKeyPinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPinFile

`func (o *AddFileBasedKeyManagerProviderRequest) SetPrivateKeyPinFile(v string)`

SetPrivateKeyPinFile sets PrivateKeyPinFile field to given value.

### HasPrivateKeyPinFile

`func (o *AddFileBasedKeyManagerProviderRequest) HasPrivateKeyPinFile() bool`

HasPrivateKeyPinFile returns a boolean if a field has been set.

### GetPrivateKeyPinPassphraseProvider

`func (o *AddFileBasedKeyManagerProviderRequest) GetPrivateKeyPinPassphraseProvider() string`

GetPrivateKeyPinPassphraseProvider returns the PrivateKeyPinPassphraseProvider field if non-nil, zero value otherwise.

### GetPrivateKeyPinPassphraseProviderOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetPrivateKeyPinPassphraseProviderOk() (*string, bool)`

GetPrivateKeyPinPassphraseProviderOk returns a tuple with the PrivateKeyPinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPinPassphraseProvider

`func (o *AddFileBasedKeyManagerProviderRequest) SetPrivateKeyPinPassphraseProvider(v string)`

SetPrivateKeyPinPassphraseProvider sets PrivateKeyPinPassphraseProvider field to given value.

### HasPrivateKeyPinPassphraseProvider

`func (o *AddFileBasedKeyManagerProviderRequest) HasPrivateKeyPinPassphraseProvider() bool`

HasPrivateKeyPinPassphraseProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddFileBasedKeyManagerProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFileBasedKeyManagerProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFileBasedKeyManagerProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddFileBasedKeyManagerProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddFileBasedKeyManagerProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *AddFileBasedKeyManagerProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddFileBasedKeyManagerProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddFileBasedKeyManagerProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


