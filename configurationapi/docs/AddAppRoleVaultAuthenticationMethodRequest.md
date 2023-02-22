# AddAppRoleVaultAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodName** | **string** | Name of the new Vault Authentication Method | 
**Schemas** | [**[]EnumappRoleVaultAuthenticationMethodSchemaUrn**](EnumappRoleVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultRoleID** | **string** | The role ID for the AppRole to authenticate. | 
**VaultSecretID** | **string** | The secret ID for the AppRole to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired AppRole authentication mechanism in the Vault server. | [optional] 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 

## Methods

### NewAddAppRoleVaultAuthenticationMethodRequest

`func NewAddAppRoleVaultAuthenticationMethodRequest(methodName string, schemas []EnumappRoleVaultAuthenticationMethodSchemaUrn, vaultRoleID string, vaultSecretID string, ) *AddAppRoleVaultAuthenticationMethodRequest`

NewAddAppRoleVaultAuthenticationMethodRequest instantiates a new AddAppRoleVaultAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAppRoleVaultAuthenticationMethodRequestWithDefaults

`func NewAddAppRoleVaultAuthenticationMethodRequestWithDefaults() *AddAppRoleVaultAuthenticationMethodRequest`

NewAddAppRoleVaultAuthenticationMethodRequestWithDefaults instantiates a new AddAppRoleVaultAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodName

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddAppRoleVaultAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetSchemas

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetSchemas() []EnumappRoleVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetSchemasOk() (*[]EnumappRoleVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAppRoleVaultAuthenticationMethodRequest) SetSchemas(v []EnumappRoleVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultRoleID

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *AddAppRoleVaultAuthenticationMethodRequest) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.


### GetVaultSecretID

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *AddAppRoleVaultAuthenticationMethodRequest) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.


### GetLoginMechanismName

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *AddAppRoleVaultAuthenticationMethodRequest) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *AddAppRoleVaultAuthenticationMethodRequest) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetDescription

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAppRoleVaultAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAppRoleVaultAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAppRoleVaultAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


