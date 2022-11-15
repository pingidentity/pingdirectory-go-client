# AddKeyManagerProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Key Manager Provider | 
**Schemas** | [**[]EnumthirdPartyKeyManagerProviderSchemaUrn**](EnumthirdPartyKeyManagerProviderSchemaUrn.md) |  | 
**KeyStoreFile** | **string** | Specifies the path to the file that contains the private key information. This may be an absolute path, or a path that is relative to the Directory Server instance root. | 
**KeyStoreType** | Pointer to **string** | Specifies the format for the data in the key store file. | [optional] 
**KeyStorePin** | Pointer to **string** | Specifies the PIN needed to access the PKCS11 Key Manager Provider. | [optional] 
**KeyStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the PKCS11 Key Manager Provider. | [optional] 
**KeyStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the PKCS11 Key Manager Provider. | [optional] 
**PrivateKeyPin** | Pointer to **string** | Specifies the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN. | [optional] 
**PrivateKeyPinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN. | [optional] 
**PrivateKeyPinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN. | [optional] 
**Description** | Pointer to **string** | A description for this Key Manager Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Key Manager Provider is enabled for use. | 
**Pkcs11ProviderClass** | Pointer to **string** | The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11ProviderConfigurationFile** | Pointer to **string** | The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11KeyStoreType** | Pointer to **string** | The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Key Manager Provider. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddKeyManagerProvider200Response

`func NewAddKeyManagerProvider200Response(id string, schemas []EnumthirdPartyKeyManagerProviderSchemaUrn, keyStoreFile string, enabled bool, extensionClass string, ) *AddKeyManagerProvider200Response`

NewAddKeyManagerProvider200Response instantiates a new AddKeyManagerProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyManagerProvider200ResponseWithDefaults

`func NewAddKeyManagerProvider200ResponseWithDefaults() *AddKeyManagerProvider200Response`

NewAddKeyManagerProvider200ResponseWithDefaults instantiates a new AddKeyManagerProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddKeyManagerProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddKeyManagerProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddKeyManagerProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddKeyManagerProvider200Response) GetSchemas() []EnumthirdPartyKeyManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddKeyManagerProvider200Response) GetSchemasOk() (*[]EnumthirdPartyKeyManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddKeyManagerProvider200Response) SetSchemas(v []EnumthirdPartyKeyManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetKeyStoreFile

`func (o *AddKeyManagerProvider200Response) GetKeyStoreFile() string`

GetKeyStoreFile returns the KeyStoreFile field if non-nil, zero value otherwise.

### GetKeyStoreFileOk

`func (o *AddKeyManagerProvider200Response) GetKeyStoreFileOk() (*string, bool)`

GetKeyStoreFileOk returns a tuple with the KeyStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStoreFile

`func (o *AddKeyManagerProvider200Response) SetKeyStoreFile(v string)`

SetKeyStoreFile sets KeyStoreFile field to given value.


### GetKeyStoreType

`func (o *AddKeyManagerProvider200Response) GetKeyStoreType() string`

GetKeyStoreType returns the KeyStoreType field if non-nil, zero value otherwise.

### GetKeyStoreTypeOk

`func (o *AddKeyManagerProvider200Response) GetKeyStoreTypeOk() (*string, bool)`

GetKeyStoreTypeOk returns a tuple with the KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStoreType

`func (o *AddKeyManagerProvider200Response) SetKeyStoreType(v string)`

SetKeyStoreType sets KeyStoreType field to given value.

### HasKeyStoreType

`func (o *AddKeyManagerProvider200Response) HasKeyStoreType() bool`

HasKeyStoreType returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *AddKeyManagerProvider200Response) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *AddKeyManagerProvider200Response) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *AddKeyManagerProvider200Response) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *AddKeyManagerProvider200Response) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *AddKeyManagerProvider200Response) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *AddKeyManagerProvider200Response) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *AddKeyManagerProvider200Response) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *AddKeyManagerProvider200Response) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinPassphraseProvider

`func (o *AddKeyManagerProvider200Response) GetKeyStorePinPassphraseProvider() string`

GetKeyStorePinPassphraseProvider returns the KeyStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetKeyStorePinPassphraseProviderOk

`func (o *AddKeyManagerProvider200Response) GetKeyStorePinPassphraseProviderOk() (*string, bool)`

GetKeyStorePinPassphraseProviderOk returns a tuple with the KeyStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinPassphraseProvider

`func (o *AddKeyManagerProvider200Response) SetKeyStorePinPassphraseProvider(v string)`

SetKeyStorePinPassphraseProvider sets KeyStorePinPassphraseProvider field to given value.

### HasKeyStorePinPassphraseProvider

`func (o *AddKeyManagerProvider200Response) HasKeyStorePinPassphraseProvider() bool`

HasKeyStorePinPassphraseProvider returns a boolean if a field has been set.

### GetPrivateKeyPin

`func (o *AddKeyManagerProvider200Response) GetPrivateKeyPin() string`

GetPrivateKeyPin returns the PrivateKeyPin field if non-nil, zero value otherwise.

### GetPrivateKeyPinOk

`func (o *AddKeyManagerProvider200Response) GetPrivateKeyPinOk() (*string, bool)`

GetPrivateKeyPinOk returns a tuple with the PrivateKeyPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPin

`func (o *AddKeyManagerProvider200Response) SetPrivateKeyPin(v string)`

SetPrivateKeyPin sets PrivateKeyPin field to given value.

### HasPrivateKeyPin

`func (o *AddKeyManagerProvider200Response) HasPrivateKeyPin() bool`

HasPrivateKeyPin returns a boolean if a field has been set.

### GetPrivateKeyPinFile

`func (o *AddKeyManagerProvider200Response) GetPrivateKeyPinFile() string`

GetPrivateKeyPinFile returns the PrivateKeyPinFile field if non-nil, zero value otherwise.

### GetPrivateKeyPinFileOk

`func (o *AddKeyManagerProvider200Response) GetPrivateKeyPinFileOk() (*string, bool)`

GetPrivateKeyPinFileOk returns a tuple with the PrivateKeyPinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPinFile

`func (o *AddKeyManagerProvider200Response) SetPrivateKeyPinFile(v string)`

SetPrivateKeyPinFile sets PrivateKeyPinFile field to given value.

### HasPrivateKeyPinFile

`func (o *AddKeyManagerProvider200Response) HasPrivateKeyPinFile() bool`

HasPrivateKeyPinFile returns a boolean if a field has been set.

### GetPrivateKeyPinPassphraseProvider

`func (o *AddKeyManagerProvider200Response) GetPrivateKeyPinPassphraseProvider() string`

GetPrivateKeyPinPassphraseProvider returns the PrivateKeyPinPassphraseProvider field if non-nil, zero value otherwise.

### GetPrivateKeyPinPassphraseProviderOk

`func (o *AddKeyManagerProvider200Response) GetPrivateKeyPinPassphraseProviderOk() (*string, bool)`

GetPrivateKeyPinPassphraseProviderOk returns a tuple with the PrivateKeyPinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPinPassphraseProvider

`func (o *AddKeyManagerProvider200Response) SetPrivateKeyPinPassphraseProvider(v string)`

SetPrivateKeyPinPassphraseProvider sets PrivateKeyPinPassphraseProvider field to given value.

### HasPrivateKeyPinPassphraseProvider

`func (o *AddKeyManagerProvider200Response) HasPrivateKeyPinPassphraseProvider() bool`

HasPrivateKeyPinPassphraseProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddKeyManagerProvider200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddKeyManagerProvider200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddKeyManagerProvider200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddKeyManagerProvider200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddKeyManagerProvider200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddKeyManagerProvider200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddKeyManagerProvider200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPkcs11ProviderClass

`func (o *AddKeyManagerProvider200Response) GetPkcs11ProviderClass() string`

GetPkcs11ProviderClass returns the Pkcs11ProviderClass field if non-nil, zero value otherwise.

### GetPkcs11ProviderClassOk

`func (o *AddKeyManagerProvider200Response) GetPkcs11ProviderClassOk() (*string, bool)`

GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderClass

`func (o *AddKeyManagerProvider200Response) SetPkcs11ProviderClass(v string)`

SetPkcs11ProviderClass sets Pkcs11ProviderClass field to given value.

### HasPkcs11ProviderClass

`func (o *AddKeyManagerProvider200Response) HasPkcs11ProviderClass() bool`

HasPkcs11ProviderClass returns a boolean if a field has been set.

### GetPkcs11ProviderConfigurationFile

`func (o *AddKeyManagerProvider200Response) GetPkcs11ProviderConfigurationFile() string`

GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field if non-nil, zero value otherwise.

### GetPkcs11ProviderConfigurationFileOk

`func (o *AddKeyManagerProvider200Response) GetPkcs11ProviderConfigurationFileOk() (*string, bool)`

GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderConfigurationFile

`func (o *AddKeyManagerProvider200Response) SetPkcs11ProviderConfigurationFile(v string)`

SetPkcs11ProviderConfigurationFile sets Pkcs11ProviderConfigurationFile field to given value.

### HasPkcs11ProviderConfigurationFile

`func (o *AddKeyManagerProvider200Response) HasPkcs11ProviderConfigurationFile() bool`

HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.

### GetPkcs11KeyStoreType

`func (o *AddKeyManagerProvider200Response) GetPkcs11KeyStoreType() string`

GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field if non-nil, zero value otherwise.

### GetPkcs11KeyStoreTypeOk

`func (o *AddKeyManagerProvider200Response) GetPkcs11KeyStoreTypeOk() (*string, bool)`

GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11KeyStoreType

`func (o *AddKeyManagerProvider200Response) SetPkcs11KeyStoreType(v string)`

SetPkcs11KeyStoreType sets Pkcs11KeyStoreType field to given value.

### HasPkcs11KeyStoreType

`func (o *AddKeyManagerProvider200Response) HasPkcs11KeyStoreType() bool`

HasPkcs11KeyStoreType returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddKeyManagerProvider200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddKeyManagerProvider200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddKeyManagerProvider200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddKeyManagerProvider200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddKeyManagerProvider200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddKeyManagerProvider200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddKeyManagerProvider200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


