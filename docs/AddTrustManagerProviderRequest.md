# AddTrustManagerProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Trust Manager Provider | 
**Schemas** | [**[]EnumthirdPartyTrustManagerProviderSchemaUrn**](EnumthirdPartyTrustManagerProviderSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicate whether the Trust Manager Provider is enabled for use. | 
**IncludeJVMDefaultIssuers** | Pointer to **bool** | Indicates whether certificates issued by an authority included in the JVM&#39;s set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider. | [optional] 
**TrustStoreFile** | **string** | Specifies the path to the file containing the trust information. It can be an absolute path or a path that is relative to the Directory Server instance root. | 
**TrustStoreType** | Pointer to **string** | Specifies the format for the data in the trust store file. | [optional] 
**TrustStorePin** | Pointer to **string** | Specifies the clear-text PIN needed to access the File Based Trust Manager Provider. | [optional] 
**TrustStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Trust Manager Provider. | [optional] 
**TrustStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Trust Manager Provider. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Trust Manager Provider. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddTrustManagerProviderRequest

`func NewAddTrustManagerProviderRequest(providerName string, schemas []EnumthirdPartyTrustManagerProviderSchemaUrn, enabled bool, trustStoreFile string, extensionClass string, ) *AddTrustManagerProviderRequest`

NewAddTrustManagerProviderRequest instantiates a new AddTrustManagerProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTrustManagerProviderRequestWithDefaults

`func NewAddTrustManagerProviderRequestWithDefaults() *AddTrustManagerProviderRequest`

NewAddTrustManagerProviderRequestWithDefaults instantiates a new AddTrustManagerProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddTrustManagerProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddTrustManagerProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddTrustManagerProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddTrustManagerProviderRequest) GetSchemas() []EnumthirdPartyTrustManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTrustManagerProviderRequest) GetSchemasOk() (*[]EnumthirdPartyTrustManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTrustManagerProviderRequest) SetSchemas(v []EnumthirdPartyTrustManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *AddTrustManagerProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddTrustManagerProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddTrustManagerProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeJVMDefaultIssuers

`func (o *AddTrustManagerProviderRequest) GetIncludeJVMDefaultIssuers() bool`

GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field if non-nil, zero value otherwise.

### GetIncludeJVMDefaultIssuersOk

`func (o *AddTrustManagerProviderRequest) GetIncludeJVMDefaultIssuersOk() (*bool, bool)`

GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJVMDefaultIssuers

`func (o *AddTrustManagerProviderRequest) SetIncludeJVMDefaultIssuers(v bool)`

SetIncludeJVMDefaultIssuers sets IncludeJVMDefaultIssuers field to given value.

### HasIncludeJVMDefaultIssuers

`func (o *AddTrustManagerProviderRequest) HasIncludeJVMDefaultIssuers() bool`

HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *AddTrustManagerProviderRequest) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *AddTrustManagerProviderRequest) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *AddTrustManagerProviderRequest) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.


### GetTrustStoreType

`func (o *AddTrustManagerProviderRequest) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *AddTrustManagerProviderRequest) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *AddTrustManagerProviderRequest) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *AddTrustManagerProviderRequest) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *AddTrustManagerProviderRequest) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *AddTrustManagerProviderRequest) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *AddTrustManagerProviderRequest) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *AddTrustManagerProviderRequest) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStorePinFile

`func (o *AddTrustManagerProviderRequest) GetTrustStorePinFile() string`

GetTrustStorePinFile returns the TrustStorePinFile field if non-nil, zero value otherwise.

### GetTrustStorePinFileOk

`func (o *AddTrustManagerProviderRequest) GetTrustStorePinFileOk() (*string, bool)`

GetTrustStorePinFileOk returns a tuple with the TrustStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinFile

`func (o *AddTrustManagerProviderRequest) SetTrustStorePinFile(v string)`

SetTrustStorePinFile sets TrustStorePinFile field to given value.

### HasTrustStorePinFile

`func (o *AddTrustManagerProviderRequest) HasTrustStorePinFile() bool`

HasTrustStorePinFile returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *AddTrustManagerProviderRequest) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *AddTrustManagerProviderRequest) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *AddTrustManagerProviderRequest) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *AddTrustManagerProviderRequest) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddTrustManagerProviderRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddTrustManagerProviderRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddTrustManagerProviderRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddTrustManagerProviderRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddTrustManagerProviderRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddTrustManagerProviderRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddTrustManagerProviderRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


