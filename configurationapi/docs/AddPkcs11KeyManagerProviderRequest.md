# AddPkcs11KeyManagerProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumpkcs11KeyManagerProviderSchemaUrn**](Enumpkcs11KeyManagerProviderSchemaUrn.md) |  | 
**Pkcs11ProviderClass** | Pointer to **string** | The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11ProviderConfigurationFile** | Pointer to **string** | The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11KeyStoreType** | Pointer to **string** | The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token. | [optional] 
**Pkcs11MaxCacheDuration** | Pointer to **string** | The maximum length of time that data retrieved from PKCS #11 tokens may be cached for reuse. Caching might be necessary if there is noticable latency when accessing the token, for example if the token uses a remote key store. A value of zero milliseconds indicates that no caching should be performed. | [optional] 
**KeyStorePin** | Pointer to **string** | Specifies the PIN needed to access the PKCS11 Key Manager Provider. | [optional] 
**KeyStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the PKCS11 Key Manager Provider. | [optional] 
**KeyStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the PKCS11 Key Manager Provider. | [optional] 
**Description** | Pointer to **string** | A description for this Key Manager Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Key Manager Provider is enabled for use. | 
**ProviderName** | **string** | Name of the new Key Manager Provider | 

## Methods

### NewAddPkcs11KeyManagerProviderRequest

`func NewAddPkcs11KeyManagerProviderRequest(schemas []Enumpkcs11KeyManagerProviderSchemaUrn, enabled bool, providerName string, ) *AddPkcs11KeyManagerProviderRequest`

NewAddPkcs11KeyManagerProviderRequest instantiates a new AddPkcs11KeyManagerProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPkcs11KeyManagerProviderRequestWithDefaults

`func NewAddPkcs11KeyManagerProviderRequestWithDefaults() *AddPkcs11KeyManagerProviderRequest`

NewAddPkcs11KeyManagerProviderRequestWithDefaults instantiates a new AddPkcs11KeyManagerProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddPkcs11KeyManagerProviderRequest) GetSchemas() []Enumpkcs11KeyManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetSchemasOk() (*[]Enumpkcs11KeyManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPkcs11KeyManagerProviderRequest) SetSchemas(v []Enumpkcs11KeyManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPkcs11ProviderClass

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11ProviderClass() string`

GetPkcs11ProviderClass returns the Pkcs11ProviderClass field if non-nil, zero value otherwise.

### GetPkcs11ProviderClassOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11ProviderClassOk() (*string, bool)`

GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderClass

`func (o *AddPkcs11KeyManagerProviderRequest) SetPkcs11ProviderClass(v string)`

SetPkcs11ProviderClass sets Pkcs11ProviderClass field to given value.

### HasPkcs11ProviderClass

`func (o *AddPkcs11KeyManagerProviderRequest) HasPkcs11ProviderClass() bool`

HasPkcs11ProviderClass returns a boolean if a field has been set.

### GetPkcs11ProviderConfigurationFile

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11ProviderConfigurationFile() string`

GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field if non-nil, zero value otherwise.

### GetPkcs11ProviderConfigurationFileOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11ProviderConfigurationFileOk() (*string, bool)`

GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderConfigurationFile

`func (o *AddPkcs11KeyManagerProviderRequest) SetPkcs11ProviderConfigurationFile(v string)`

SetPkcs11ProviderConfigurationFile sets Pkcs11ProviderConfigurationFile field to given value.

### HasPkcs11ProviderConfigurationFile

`func (o *AddPkcs11KeyManagerProviderRequest) HasPkcs11ProviderConfigurationFile() bool`

HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.

### GetPkcs11KeyStoreType

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11KeyStoreType() string`

GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field if non-nil, zero value otherwise.

### GetPkcs11KeyStoreTypeOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11KeyStoreTypeOk() (*string, bool)`

GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11KeyStoreType

`func (o *AddPkcs11KeyManagerProviderRequest) SetPkcs11KeyStoreType(v string)`

SetPkcs11KeyStoreType sets Pkcs11KeyStoreType field to given value.

### HasPkcs11KeyStoreType

`func (o *AddPkcs11KeyManagerProviderRequest) HasPkcs11KeyStoreType() bool`

HasPkcs11KeyStoreType returns a boolean if a field has been set.

### GetPkcs11MaxCacheDuration

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11MaxCacheDuration() string`

GetPkcs11MaxCacheDuration returns the Pkcs11MaxCacheDuration field if non-nil, zero value otherwise.

### GetPkcs11MaxCacheDurationOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetPkcs11MaxCacheDurationOk() (*string, bool)`

GetPkcs11MaxCacheDurationOk returns a tuple with the Pkcs11MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11MaxCacheDuration

`func (o *AddPkcs11KeyManagerProviderRequest) SetPkcs11MaxCacheDuration(v string)`

SetPkcs11MaxCacheDuration sets Pkcs11MaxCacheDuration field to given value.

### HasPkcs11MaxCacheDuration

`func (o *AddPkcs11KeyManagerProviderRequest) HasPkcs11MaxCacheDuration() bool`

HasPkcs11MaxCacheDuration returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *AddPkcs11KeyManagerProviderRequest) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *AddPkcs11KeyManagerProviderRequest) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *AddPkcs11KeyManagerProviderRequest) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *AddPkcs11KeyManagerProviderRequest) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *AddPkcs11KeyManagerProviderRequest) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *AddPkcs11KeyManagerProviderRequest) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinPassphraseProvider

`func (o *AddPkcs11KeyManagerProviderRequest) GetKeyStorePinPassphraseProvider() string`

GetKeyStorePinPassphraseProvider returns the KeyStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetKeyStorePinPassphraseProviderOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetKeyStorePinPassphraseProviderOk() (*string, bool)`

GetKeyStorePinPassphraseProviderOk returns a tuple with the KeyStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinPassphraseProvider

`func (o *AddPkcs11KeyManagerProviderRequest) SetKeyStorePinPassphraseProvider(v string)`

SetKeyStorePinPassphraseProvider sets KeyStorePinPassphraseProvider field to given value.

### HasKeyStorePinPassphraseProvider

`func (o *AddPkcs11KeyManagerProviderRequest) HasKeyStorePinPassphraseProvider() bool`

HasKeyStorePinPassphraseProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddPkcs11KeyManagerProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPkcs11KeyManagerProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPkcs11KeyManagerProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPkcs11KeyManagerProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPkcs11KeyManagerProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *AddPkcs11KeyManagerProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddPkcs11KeyManagerProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddPkcs11KeyManagerProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


