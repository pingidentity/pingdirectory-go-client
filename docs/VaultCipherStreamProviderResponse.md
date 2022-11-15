# VaultCipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Cipher Stream Provider | 
**Schemas** | [**[]EnumvaultCipherStreamProviderSchemaUrn**](EnumvaultCipherStreamProviderSchemaUrn.md) |  | 
**VaultExternalServer** | Pointer to **string** | An external server definition with information needed to connect and authenticate to the Vault server. | [optional] 
**VaultServerBaseURI** | Pointer to **[]string** |  | [optional] 
**VaultAuthenticationMethod** | Pointer to **string** | The mechanism used to authenticate to the Vault server. | [optional] 
**VaultSecretPath** | **string** | The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server. | 
**VaultSecretFieldName** | **string** | The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key. | 
**VaultEncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this Vault Cipher Stream Provider. | 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Vault servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot; or \&quot;PKCS12\&quot;. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewVaultCipherStreamProviderResponse

`func NewVaultCipherStreamProviderResponse(id string, schemas []EnumvaultCipherStreamProviderSchemaUrn, vaultSecretPath string, vaultSecretFieldName string, vaultEncryptionMetadataFile string, enabled bool, ) *VaultCipherStreamProviderResponse`

NewVaultCipherStreamProviderResponse instantiates a new VaultCipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultCipherStreamProviderResponseWithDefaults

`func NewVaultCipherStreamProviderResponseWithDefaults() *VaultCipherStreamProviderResponse`

NewVaultCipherStreamProviderResponseWithDefaults instantiates a new VaultCipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VaultCipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultCipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultCipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *VaultCipherStreamProviderResponse) GetSchemas() []EnumvaultCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VaultCipherStreamProviderResponse) GetSchemasOk() (*[]EnumvaultCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VaultCipherStreamProviderResponse) SetSchemas(v []EnumvaultCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultExternalServer

`func (o *VaultCipherStreamProviderResponse) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *VaultCipherStreamProviderResponse) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *VaultCipherStreamProviderResponse) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.

### HasVaultExternalServer

`func (o *VaultCipherStreamProviderResponse) HasVaultExternalServer() bool`

HasVaultExternalServer returns a boolean if a field has been set.

### GetVaultServerBaseURI

`func (o *VaultCipherStreamProviderResponse) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *VaultCipherStreamProviderResponse) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *VaultCipherStreamProviderResponse) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.

### HasVaultServerBaseURI

`func (o *VaultCipherStreamProviderResponse) HasVaultServerBaseURI() bool`

HasVaultServerBaseURI returns a boolean if a field has been set.

### GetVaultAuthenticationMethod

`func (o *VaultCipherStreamProviderResponse) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *VaultCipherStreamProviderResponse) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *VaultCipherStreamProviderResponse) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.

### HasVaultAuthenticationMethod

`func (o *VaultCipherStreamProviderResponse) HasVaultAuthenticationMethod() bool`

HasVaultAuthenticationMethod returns a boolean if a field has been set.

### GetVaultSecretPath

`func (o *VaultCipherStreamProviderResponse) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *VaultCipherStreamProviderResponse) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *VaultCipherStreamProviderResponse) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *VaultCipherStreamProviderResponse) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *VaultCipherStreamProviderResponse) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *VaultCipherStreamProviderResponse) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetVaultEncryptionMetadataFile

`func (o *VaultCipherStreamProviderResponse) GetVaultEncryptionMetadataFile() string`

GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field if non-nil, zero value otherwise.

### GetVaultEncryptionMetadataFileOk

`func (o *VaultCipherStreamProviderResponse) GetVaultEncryptionMetadataFileOk() (*string, bool)`

GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultEncryptionMetadataFile

`func (o *VaultCipherStreamProviderResponse) SetVaultEncryptionMetadataFile(v string)`

SetVaultEncryptionMetadataFile sets VaultEncryptionMetadataFile field to given value.


### GetTrustStoreFile

`func (o *VaultCipherStreamProviderResponse) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *VaultCipherStreamProviderResponse) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *VaultCipherStreamProviderResponse) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *VaultCipherStreamProviderResponse) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *VaultCipherStreamProviderResponse) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *VaultCipherStreamProviderResponse) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *VaultCipherStreamProviderResponse) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *VaultCipherStreamProviderResponse) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *VaultCipherStreamProviderResponse) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *VaultCipherStreamProviderResponse) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *VaultCipherStreamProviderResponse) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *VaultCipherStreamProviderResponse) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetDescription

`func (o *VaultCipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VaultCipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VaultCipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VaultCipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VaultCipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VaultCipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VaultCipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


