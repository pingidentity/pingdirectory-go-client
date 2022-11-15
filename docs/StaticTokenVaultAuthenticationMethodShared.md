# StaticTokenVaultAuthenticationMethodShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumstaticTokenVaultAuthenticationMethodSchemaUrn**](EnumstaticTokenVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultAccessToken** | **string** | The static token used to authenticate to the Vault server. | 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 

## Methods

### NewStaticTokenVaultAuthenticationMethodShared

`func NewStaticTokenVaultAuthenticationMethodShared(schemas []EnumstaticTokenVaultAuthenticationMethodSchemaUrn, vaultAccessToken string, ) *StaticTokenVaultAuthenticationMethodShared`

NewStaticTokenVaultAuthenticationMethodShared instantiates a new StaticTokenVaultAuthenticationMethodShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticTokenVaultAuthenticationMethodSharedWithDefaults

`func NewStaticTokenVaultAuthenticationMethodSharedWithDefaults() *StaticTokenVaultAuthenticationMethodShared`

NewStaticTokenVaultAuthenticationMethodSharedWithDefaults instantiates a new StaticTokenVaultAuthenticationMethodShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *StaticTokenVaultAuthenticationMethodShared) GetSchemas() []EnumstaticTokenVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StaticTokenVaultAuthenticationMethodShared) GetSchemasOk() (*[]EnumstaticTokenVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StaticTokenVaultAuthenticationMethodShared) SetSchemas(v []EnumstaticTokenVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultAccessToken

`func (o *StaticTokenVaultAuthenticationMethodShared) GetVaultAccessToken() string`

GetVaultAccessToken returns the VaultAccessToken field if non-nil, zero value otherwise.

### GetVaultAccessTokenOk

`func (o *StaticTokenVaultAuthenticationMethodShared) GetVaultAccessTokenOk() (*string, bool)`

GetVaultAccessTokenOk returns a tuple with the VaultAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAccessToken

`func (o *StaticTokenVaultAuthenticationMethodShared) SetVaultAccessToken(v string)`

SetVaultAccessToken sets VaultAccessToken field to given value.


### GetDescription

`func (o *StaticTokenVaultAuthenticationMethodShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticTokenVaultAuthenticationMethodShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticTokenVaultAuthenticationMethodShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticTokenVaultAuthenticationMethodShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


