# AddUserPassVaultAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumuserPassVaultAuthenticationMethodSchemaUrn**](EnumuserPassVaultAuthenticationMethodSchemaUrn.md) |  | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired UserPass authentication mechanism in the Vault server. | [optional] 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 
**MethodName** | **string** | Name of the new Vault Authentication Method | 

## Methods

### NewAddUserPassVaultAuthenticationMethodRequest

`func NewAddUserPassVaultAuthenticationMethodRequest(schemas []EnumuserPassVaultAuthenticationMethodSchemaUrn, username string, password string, methodName string, ) *AddUserPassVaultAuthenticationMethodRequest`

NewAddUserPassVaultAuthenticationMethodRequest instantiates a new AddUserPassVaultAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserPassVaultAuthenticationMethodRequestWithDefaults

`func NewAddUserPassVaultAuthenticationMethodRequestWithDefaults() *AddUserPassVaultAuthenticationMethodRequest`

NewAddUserPassVaultAuthenticationMethodRequestWithDefaults instantiates a new AddUserPassVaultAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetSchemas() []EnumuserPassVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetSchemasOk() (*[]EnumuserPassVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUserPassVaultAuthenticationMethodRequest) SetSchemas(v []EnumuserPassVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetUsername

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUserPassVaultAuthenticationMethodRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUserPassVaultAuthenticationMethodRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLoginMechanismName

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *AddUserPassVaultAuthenticationMethodRequest) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *AddUserPassVaultAuthenticationMethodRequest) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetDescription

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUserPassVaultAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUserPassVaultAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMethodName

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddUserPassVaultAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddUserPassVaultAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


