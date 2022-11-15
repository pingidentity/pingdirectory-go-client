# AddVaultExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumvaultExternalServerSchemaUrn**](EnumvaultExternalServerSchemaUrn.md) |  | 
**VaultServerBaseURI** | **[]string** |  | 
**VaultAuthenticationMethod** | **string** | The mechanism used to authenticate to the Vault server. | 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Vault servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot;, \&quot;PKCS12\&quot;, or \&quot;BCFKS\&quot;. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewAddVaultExternalServerRequest

`func NewAddVaultExternalServerRequest(serverName string, schemas []EnumvaultExternalServerSchemaUrn, vaultServerBaseURI []string, vaultAuthenticationMethod string, ) *AddVaultExternalServerRequest`

NewAddVaultExternalServerRequest instantiates a new AddVaultExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVaultExternalServerRequestWithDefaults

`func NewAddVaultExternalServerRequestWithDefaults() *AddVaultExternalServerRequest`

NewAddVaultExternalServerRequestWithDefaults instantiates a new AddVaultExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddVaultExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddVaultExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddVaultExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddVaultExternalServerRequest) GetSchemas() []EnumvaultExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVaultExternalServerRequest) GetSchemasOk() (*[]EnumvaultExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVaultExternalServerRequest) SetSchemas(v []EnumvaultExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultServerBaseURI

`func (o *AddVaultExternalServerRequest) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *AddVaultExternalServerRequest) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *AddVaultExternalServerRequest) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.


### GetVaultAuthenticationMethod

`func (o *AddVaultExternalServerRequest) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *AddVaultExternalServerRequest) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *AddVaultExternalServerRequest) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.


### GetTrustStoreFile

`func (o *AddVaultExternalServerRequest) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *AddVaultExternalServerRequest) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *AddVaultExternalServerRequest) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *AddVaultExternalServerRequest) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *AddVaultExternalServerRequest) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *AddVaultExternalServerRequest) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *AddVaultExternalServerRequest) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *AddVaultExternalServerRequest) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *AddVaultExternalServerRequest) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *AddVaultExternalServerRequest) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *AddVaultExternalServerRequest) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *AddVaultExternalServerRequest) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetDescription

`func (o *AddVaultExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVaultExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVaultExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVaultExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


