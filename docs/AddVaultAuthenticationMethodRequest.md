# AddVaultAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodName** | **string** | Name of the new Vault Authentication Method | 
**Schemas** | [**[]EnumuserPassVaultAuthenticationMethodSchemaUrn**](EnumuserPassVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultAccessToken** | **string** | The static token used to authenticate to the Vault server. | 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 
**VaultRoleID** | **string** | The role ID for the AppRole to authenticate. | 
**VaultSecretID** | **string** | The secret ID for the AppRole to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired UserPass authentication mechanism in the Vault server. | [optional] 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 

## Methods

### NewAddVaultAuthenticationMethodRequest

`func NewAddVaultAuthenticationMethodRequest(methodName string, schemas []EnumuserPassVaultAuthenticationMethodSchemaUrn, vaultAccessToken string, vaultRoleID string, vaultSecretID string, username string, password string, ) *AddVaultAuthenticationMethodRequest`

NewAddVaultAuthenticationMethodRequest instantiates a new AddVaultAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVaultAuthenticationMethodRequestWithDefaults

`func NewAddVaultAuthenticationMethodRequestWithDefaults() *AddVaultAuthenticationMethodRequest`

NewAddVaultAuthenticationMethodRequestWithDefaults instantiates a new AddVaultAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodName

`func (o *AddVaultAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddVaultAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddVaultAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetSchemas

`func (o *AddVaultAuthenticationMethodRequest) GetSchemas() []EnumuserPassVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVaultAuthenticationMethodRequest) GetSchemasOk() (*[]EnumuserPassVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVaultAuthenticationMethodRequest) SetSchemas(v []EnumuserPassVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultAccessToken

`func (o *AddVaultAuthenticationMethodRequest) GetVaultAccessToken() string`

GetVaultAccessToken returns the VaultAccessToken field if non-nil, zero value otherwise.

### GetVaultAccessTokenOk

`func (o *AddVaultAuthenticationMethodRequest) GetVaultAccessTokenOk() (*string, bool)`

GetVaultAccessTokenOk returns a tuple with the VaultAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAccessToken

`func (o *AddVaultAuthenticationMethodRequest) SetVaultAccessToken(v string)`

SetVaultAccessToken sets VaultAccessToken field to given value.


### GetDescription

`func (o *AddVaultAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVaultAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVaultAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVaultAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVaultRoleID

`func (o *AddVaultAuthenticationMethodRequest) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *AddVaultAuthenticationMethodRequest) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *AddVaultAuthenticationMethodRequest) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.


### GetVaultSecretID

`func (o *AddVaultAuthenticationMethodRequest) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *AddVaultAuthenticationMethodRequest) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *AddVaultAuthenticationMethodRequest) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.


### GetLoginMechanismName

`func (o *AddVaultAuthenticationMethodRequest) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *AddVaultAuthenticationMethodRequest) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *AddVaultAuthenticationMethodRequest) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *AddVaultAuthenticationMethodRequest) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetUsername

`func (o *AddVaultAuthenticationMethodRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddVaultAuthenticationMethodRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddVaultAuthenticationMethodRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddVaultAuthenticationMethodRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddVaultAuthenticationMethodRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddVaultAuthenticationMethodRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


