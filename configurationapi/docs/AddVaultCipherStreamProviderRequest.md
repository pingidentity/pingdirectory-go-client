# AddVaultCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvaultCipherStreamProviderSchemaUrn**](EnumvaultCipherStreamProviderSchemaUrn.md) |  | 
**VaultExternalServer** | Pointer to **string** | An external server definition with information needed to connect and authenticate to the Vault server. | [optional] 
**VaultServerBaseURI** | Pointer to **[]string** | The base URL needed to access the Vault server. The base URL should consist of the protocol (\&quot;http\&quot; or \&quot;https\&quot;), the server address (resolvable name or IP address), and the port number. For example, \&quot;https://vault.example.com:8200/\&quot;. | [optional] 
**VaultAuthenticationMethod** | Pointer to **string** | The mechanism used to authenticate to the Vault server. | [optional] 
**VaultSecretPath** | **string** | The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server. | 
**VaultSecretFieldName** | **string** | The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key. | 
**VaultEncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this Vault Cipher Stream Provider. | [optional] 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Vault servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot; or \&quot;PKCS12\&quot;. | [optional] 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 

## Methods

### NewAddVaultCipherStreamProviderRequest

`func NewAddVaultCipherStreamProviderRequest(schemas []EnumvaultCipherStreamProviderSchemaUrn, vaultSecretPath string, vaultSecretFieldName string, enabled bool, providerName string, ) *AddVaultCipherStreamProviderRequest`

NewAddVaultCipherStreamProviderRequest instantiates a new AddVaultCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVaultCipherStreamProviderRequestWithDefaults

`func NewAddVaultCipherStreamProviderRequestWithDefaults() *AddVaultCipherStreamProviderRequest`

NewAddVaultCipherStreamProviderRequestWithDefaults instantiates a new AddVaultCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddVaultCipherStreamProviderRequest) GetSchemas() []EnumvaultCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVaultCipherStreamProviderRequest) GetSchemasOk() (*[]EnumvaultCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVaultCipherStreamProviderRequest) SetSchemas(v []EnumvaultCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultExternalServer

`func (o *AddVaultCipherStreamProviderRequest) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *AddVaultCipherStreamProviderRequest) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *AddVaultCipherStreamProviderRequest) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.

### HasVaultExternalServer

`func (o *AddVaultCipherStreamProviderRequest) HasVaultExternalServer() bool`

HasVaultExternalServer returns a boolean if a field has been set.

### GetVaultServerBaseURI

`func (o *AddVaultCipherStreamProviderRequest) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *AddVaultCipherStreamProviderRequest) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *AddVaultCipherStreamProviderRequest) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.

### HasVaultServerBaseURI

`func (o *AddVaultCipherStreamProviderRequest) HasVaultServerBaseURI() bool`

HasVaultServerBaseURI returns a boolean if a field has been set.

### GetVaultAuthenticationMethod

`func (o *AddVaultCipherStreamProviderRequest) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *AddVaultCipherStreamProviderRequest) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *AddVaultCipherStreamProviderRequest) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.

### HasVaultAuthenticationMethod

`func (o *AddVaultCipherStreamProviderRequest) HasVaultAuthenticationMethod() bool`

HasVaultAuthenticationMethod returns a boolean if a field has been set.

### GetVaultSecretPath

`func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *AddVaultCipherStreamProviderRequest) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *AddVaultCipherStreamProviderRequest) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetVaultEncryptionMetadataFile

`func (o *AddVaultCipherStreamProviderRequest) GetVaultEncryptionMetadataFile() string`

GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field if non-nil, zero value otherwise.

### GetVaultEncryptionMetadataFileOk

`func (o *AddVaultCipherStreamProviderRequest) GetVaultEncryptionMetadataFileOk() (*string, bool)`

GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultEncryptionMetadataFile

`func (o *AddVaultCipherStreamProviderRequest) SetVaultEncryptionMetadataFile(v string)`

SetVaultEncryptionMetadataFile sets VaultEncryptionMetadataFile field to given value.

### HasVaultEncryptionMetadataFile

`func (o *AddVaultCipherStreamProviderRequest) HasVaultEncryptionMetadataFile() bool`

HasVaultEncryptionMetadataFile returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *AddVaultCipherStreamProviderRequest) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *AddVaultCipherStreamProviderRequest) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *AddVaultCipherStreamProviderRequest) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *AddVaultCipherStreamProviderRequest) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *AddVaultCipherStreamProviderRequest) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *AddVaultCipherStreamProviderRequest) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *AddVaultCipherStreamProviderRequest) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *AddVaultCipherStreamProviderRequest) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetIterationCount

`func (o *AddVaultCipherStreamProviderRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddVaultCipherStreamProviderRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddVaultCipherStreamProviderRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *AddVaultCipherStreamProviderRequest) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *AddVaultCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVaultCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVaultCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVaultCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVaultCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVaultCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVaultCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *AddVaultCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddVaultCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddVaultCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


