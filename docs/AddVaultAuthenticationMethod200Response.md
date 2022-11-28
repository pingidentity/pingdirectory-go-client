# AddVaultAuthenticationMethod200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Vault Authentication Method | 
**Schemas** | [**[]EnumuserPassVaultAuthenticationMethodSchemaUrn**](EnumuserPassVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultAccessToken** | **string** | The static token used to authenticate to the Vault server. | 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 
**VaultRoleID** | **string** | The role ID for the AppRole to authenticate. | 
**VaultSecretID** | **string** | The secret ID for the AppRole to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired UserPass authentication mechanism in the Vault server. | [optional] 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 

## Methods

### NewAddVaultAuthenticationMethod200Response

`func NewAddVaultAuthenticationMethod200Response(id string, schemas []EnumuserPassVaultAuthenticationMethodSchemaUrn, vaultAccessToken string, vaultRoleID string, vaultSecretID string, username string, password string, ) *AddVaultAuthenticationMethod200Response`

NewAddVaultAuthenticationMethod200Response instantiates a new AddVaultAuthenticationMethod200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVaultAuthenticationMethod200ResponseWithDefaults

`func NewAddVaultAuthenticationMethod200ResponseWithDefaults() *AddVaultAuthenticationMethod200Response`

NewAddVaultAuthenticationMethod200ResponseWithDefaults instantiates a new AddVaultAuthenticationMethod200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AddVaultAuthenticationMethod200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddVaultAuthenticationMethod200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddVaultAuthenticationMethod200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddVaultAuthenticationMethod200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddVaultAuthenticationMethod200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddVaultAuthenticationMethod200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddVaultAuthenticationMethod200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddVaultAuthenticationMethod200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AddVaultAuthenticationMethod200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddVaultAuthenticationMethod200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddVaultAuthenticationMethod200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddVaultAuthenticationMethod200Response) GetSchemas() []EnumuserPassVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVaultAuthenticationMethod200Response) GetSchemasOk() (*[]EnumuserPassVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVaultAuthenticationMethod200Response) SetSchemas(v []EnumuserPassVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultAccessToken

`func (o *AddVaultAuthenticationMethod200Response) GetVaultAccessToken() string`

GetVaultAccessToken returns the VaultAccessToken field if non-nil, zero value otherwise.

### GetVaultAccessTokenOk

`func (o *AddVaultAuthenticationMethod200Response) GetVaultAccessTokenOk() (*string, bool)`

GetVaultAccessTokenOk returns a tuple with the VaultAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAccessToken

`func (o *AddVaultAuthenticationMethod200Response) SetVaultAccessToken(v string)`

SetVaultAccessToken sets VaultAccessToken field to given value.


### GetDescription

`func (o *AddVaultAuthenticationMethod200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVaultAuthenticationMethod200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVaultAuthenticationMethod200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVaultAuthenticationMethod200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVaultRoleID

`func (o *AddVaultAuthenticationMethod200Response) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *AddVaultAuthenticationMethod200Response) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *AddVaultAuthenticationMethod200Response) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.


### GetVaultSecretID

`func (o *AddVaultAuthenticationMethod200Response) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *AddVaultAuthenticationMethod200Response) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *AddVaultAuthenticationMethod200Response) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.


### GetLoginMechanismName

`func (o *AddVaultAuthenticationMethod200Response) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *AddVaultAuthenticationMethod200Response) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *AddVaultAuthenticationMethod200Response) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *AddVaultAuthenticationMethod200Response) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetUsername

`func (o *AddVaultAuthenticationMethod200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddVaultAuthenticationMethod200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddVaultAuthenticationMethod200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddVaultAuthenticationMethod200Response) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddVaultAuthenticationMethod200Response) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddVaultAuthenticationMethod200Response) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


