# VaultAuthenticationMethodListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Vault Authentication Method | 
**Schemas** | [**[]EnumstaticTokenVaultAuthenticationMethodSchemaUrn**](EnumstaticTokenVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultRoleID** | **string** | The role ID for the AppRole to authenticate. | 
**VaultSecretID** | **string** | The secret ID for the AppRole to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired UserPass authentication mechanism in the Vault server. | [optional] 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 
**VaultAccessToken** | **string** | The static token used to authenticate to the Vault server. | 

## Methods

### NewVaultAuthenticationMethodListResponseResourcesInner

`func NewVaultAuthenticationMethodListResponseResourcesInner(id string, schemas []EnumstaticTokenVaultAuthenticationMethodSchemaUrn, vaultRoleID string, vaultSecretID string, username string, password string, vaultAccessToken string, ) *VaultAuthenticationMethodListResponseResourcesInner`

NewVaultAuthenticationMethodListResponseResourcesInner instantiates a new VaultAuthenticationMethodListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultAuthenticationMethodListResponseResourcesInnerWithDefaults

`func NewVaultAuthenticationMethodListResponseResourcesInnerWithDefaults() *VaultAuthenticationMethodListResponseResourcesInner`

NewVaultAuthenticationMethodListResponseResourcesInnerWithDefaults instantiates a new VaultAuthenticationMethodListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetSchemas() []EnumstaticTokenVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetSchemasOk() (*[]EnumstaticTokenVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetSchemas(v []EnumstaticTokenVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultRoleID

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.


### GetVaultSecretID

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.


### GetLoginMechanismName

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *VaultAuthenticationMethodListResponseResourcesInner) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetDescription

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VaultAuthenticationMethodListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VaultAuthenticationMethodListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *VaultAuthenticationMethodListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetUsername

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetVaultAccessToken

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetVaultAccessToken() string`

GetVaultAccessToken returns the VaultAccessToken field if non-nil, zero value otherwise.

### GetVaultAccessTokenOk

`func (o *VaultAuthenticationMethodListResponseResourcesInner) GetVaultAccessTokenOk() (*string, bool)`

GetVaultAccessTokenOk returns a tuple with the VaultAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAccessToken

`func (o *VaultAuthenticationMethodListResponseResourcesInner) SetVaultAccessToken(v string)`

SetVaultAccessToken sets VaultAccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


