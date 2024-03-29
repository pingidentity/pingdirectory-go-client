# AppRoleVaultAuthenticationMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumappRoleVaultAuthenticationMethodSchemaUrn**](EnumappRoleVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultRoleID** | **string** | The role ID for the AppRole to authenticate. | 
**VaultSecretID** | **string** | The secret ID for the AppRole to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired AppRole authentication mechanism in the Vault server. | [optional] 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Vault Authentication Method | 

## Methods

### NewAppRoleVaultAuthenticationMethodResponse

`func NewAppRoleVaultAuthenticationMethodResponse(schemas []EnumappRoleVaultAuthenticationMethodSchemaUrn, vaultRoleID string, vaultSecretID string, id string, ) *AppRoleVaultAuthenticationMethodResponse`

NewAppRoleVaultAuthenticationMethodResponse instantiates a new AppRoleVaultAuthenticationMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleVaultAuthenticationMethodResponseWithDefaults

`func NewAppRoleVaultAuthenticationMethodResponseWithDefaults() *AppRoleVaultAuthenticationMethodResponse`

NewAppRoleVaultAuthenticationMethodResponseWithDefaults instantiates a new AppRoleVaultAuthenticationMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AppRoleVaultAuthenticationMethodResponse) GetSchemas() []EnumappRoleVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetSchemasOk() (*[]EnumappRoleVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AppRoleVaultAuthenticationMethodResponse) SetSchemas(v []EnumappRoleVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultRoleID

`func (o *AppRoleVaultAuthenticationMethodResponse) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *AppRoleVaultAuthenticationMethodResponse) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.


### GetVaultSecretID

`func (o *AppRoleVaultAuthenticationMethodResponse) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *AppRoleVaultAuthenticationMethodResponse) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.


### GetLoginMechanismName

`func (o *AppRoleVaultAuthenticationMethodResponse) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *AppRoleVaultAuthenticationMethodResponse) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *AppRoleVaultAuthenticationMethodResponse) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetDescription

`func (o *AppRoleVaultAuthenticationMethodResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRoleVaultAuthenticationMethodResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRoleVaultAuthenticationMethodResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AppRoleVaultAuthenticationMethodResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppRoleVaultAuthenticationMethodResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppRoleVaultAuthenticationMethodResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AppRoleVaultAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AppRoleVaultAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AppRoleVaultAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AppRoleVaultAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AppRoleVaultAuthenticationMethodResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRoleVaultAuthenticationMethodResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRoleVaultAuthenticationMethodResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


