# AddStaticTokenVaultAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodName** | **string** | Name of the new Vault Authentication Method | 
**Schemas** | [**[]EnumstaticTokenVaultAuthenticationMethodSchemaUrn**](EnumstaticTokenVaultAuthenticationMethodSchemaUrn.md) |  | 
**VaultAccessToken** | **string** | The static token used to authenticate to the Vault server. | 
**Description** | Pointer to **string** | A description for this Vault Authentication Method | [optional] 

## Methods

### NewAddStaticTokenVaultAuthenticationMethodRequest

`func NewAddStaticTokenVaultAuthenticationMethodRequest(methodName string, schemas []EnumstaticTokenVaultAuthenticationMethodSchemaUrn, vaultAccessToken string, ) *AddStaticTokenVaultAuthenticationMethodRequest`

NewAddStaticTokenVaultAuthenticationMethodRequest instantiates a new AddStaticTokenVaultAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStaticTokenVaultAuthenticationMethodRequestWithDefaults

`func NewAddStaticTokenVaultAuthenticationMethodRequestWithDefaults() *AddStaticTokenVaultAuthenticationMethodRequest`

NewAddStaticTokenVaultAuthenticationMethodRequestWithDefaults instantiates a new AddStaticTokenVaultAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodName

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetSchemas

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetSchemas() []EnumstaticTokenVaultAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetSchemasOk() (*[]EnumstaticTokenVaultAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetSchemas(v []EnumstaticTokenVaultAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVaultAccessToken

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetVaultAccessToken() string`

GetVaultAccessToken returns the VaultAccessToken field if non-nil, zero value otherwise.

### GetVaultAccessTokenOk

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetVaultAccessTokenOk() (*string, bool)`

GetVaultAccessTokenOk returns a tuple with the VaultAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAccessToken

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetVaultAccessToken(v string)`

SetVaultAccessToken sets VaultAccessToken field to given value.


### GetDescription

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStaticTokenVaultAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


