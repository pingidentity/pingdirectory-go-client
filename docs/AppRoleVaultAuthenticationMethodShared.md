# AppRoleVaultAuthenticationMethodShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumappRoleVaultAuthenticationMethodSchemaUrn**](EnumappRoleVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultRoleID** | **string** | The role ID for the AppRole to authenticate. | 
**VaultSecretID** | **string** | The secret ID for the AppRole to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired AppRole authentication mechanism in the Vault server. | [optional] 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 

## Methods

### NewAppRoleVaultAuthenticationMethodShared

`func NewAppRoleVaultAuthenticationMethodShared(schemas []EnumappRoleVaultAuthenticationMethodSchemaUrn, vaultRoleID string, vaultSecretID string, ) *AppRoleVaultAuthenticationMethodShared`

NewAppRoleVaultAuthenticationMethodShared instantiates a new AppRoleVaultAuthenticationMethodShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleVaultAuthenticationMethodSharedWithDefaults

`func NewAppRoleVaultAuthenticationMethodSharedWithDefaults() *AppRoleVaultAuthenticationMethodShared`

NewAppRoleVaultAuthenticationMethodSharedWithDefaults instantiates a new AppRoleVaultAuthenticationMethodShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AppRoleVaultAuthenticationMethodShared) GetSchemas() []EnumappRoleVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AppRoleVaultAuthenticationMethodShared) GetSchemasOk() (*[]EnumappRoleVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AppRoleVaultAuthenticationMethodShared) SetSchemas(v []EnumappRoleVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultRoleID

`func (o *AppRoleVaultAuthenticationMethodShared) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *AppRoleVaultAuthenticationMethodShared) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *AppRoleVaultAuthenticationMethodShared) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.


### GetVaultSecretID

`func (o *AppRoleVaultAuthenticationMethodShared) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *AppRoleVaultAuthenticationMethodShared) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *AppRoleVaultAuthenticationMethodShared) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.


### GetLoginMechanismName

`func (o *AppRoleVaultAuthenticationMethodShared) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *AppRoleVaultAuthenticationMethodShared) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *AppRoleVaultAuthenticationMethodShared) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *AppRoleVaultAuthenticationMethodShared) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetDescription

`func (o *AppRoleVaultAuthenticationMethodShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRoleVaultAuthenticationMethodShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRoleVaultAuthenticationMethodShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRoleVaultAuthenticationMethodShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


