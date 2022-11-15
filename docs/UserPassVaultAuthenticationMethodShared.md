# UserPassVaultAuthenticationMethodShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumuserPassVaultAuthenticationMethodSchemaUrn**](EnumuserPassVaultAuthenticationMethodSchemaUrn.md) |  | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 
**LoginMechanismName** | Pointer to **string** | The name used when enabling the desired UserPass authentication mechanism in the Vault server. | [optional] 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 

## Methods

### NewUserPassVaultAuthenticationMethodShared

`func NewUserPassVaultAuthenticationMethodShared(schemas []EnumuserPassVaultAuthenticationMethodSchemaUrn, username string, password string, ) *UserPassVaultAuthenticationMethodShared`

NewUserPassVaultAuthenticationMethodShared instantiates a new UserPassVaultAuthenticationMethodShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPassVaultAuthenticationMethodSharedWithDefaults

`func NewUserPassVaultAuthenticationMethodSharedWithDefaults() *UserPassVaultAuthenticationMethodShared`

NewUserPassVaultAuthenticationMethodSharedWithDefaults instantiates a new UserPassVaultAuthenticationMethodShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UserPassVaultAuthenticationMethodShared) GetSchemas() []EnumuserPassVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UserPassVaultAuthenticationMethodShared) GetSchemasOk() (*[]EnumuserPassVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UserPassVaultAuthenticationMethodShared) SetSchemas(v []EnumuserPassVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetUsername

`func (o *UserPassVaultAuthenticationMethodShared) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPassVaultAuthenticationMethodShared) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPassVaultAuthenticationMethodShared) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserPassVaultAuthenticationMethodShared) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPassVaultAuthenticationMethodShared) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPassVaultAuthenticationMethodShared) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLoginMechanismName

`func (o *UserPassVaultAuthenticationMethodShared) GetLoginMechanismName() string`

GetLoginMechanismName returns the LoginMechanismName field if non-nil, zero value otherwise.

### GetLoginMechanismNameOk

`func (o *UserPassVaultAuthenticationMethodShared) GetLoginMechanismNameOk() (*string, bool)`

GetLoginMechanismNameOk returns a tuple with the LoginMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMechanismName

`func (o *UserPassVaultAuthenticationMethodShared) SetLoginMechanismName(v string)`

SetLoginMechanismName sets LoginMechanismName field to given value.

### HasLoginMechanismName

`func (o *UserPassVaultAuthenticationMethodShared) HasLoginMechanismName() bool`

HasLoginMechanismName returns a boolean if a field has been set.

### GetDescription

`func (o *UserPassVaultAuthenticationMethodShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserPassVaultAuthenticationMethodShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserPassVaultAuthenticationMethodShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserPassVaultAuthenticationMethodShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


