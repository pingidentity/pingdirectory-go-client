# AddClientSecretAzureAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumclientSecretAzureAuthenticationMethodSchemaUrn**](EnumclientSecretAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**ClientSecret** | **string** | The client secret to use to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 
**MethodName** | **string** | Name of the new Azure Authentication Method | 

## Methods

### NewAddClientSecretAzureAuthenticationMethodRequest

`func NewAddClientSecretAzureAuthenticationMethodRequest(schemas []EnumclientSecretAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, clientSecret string, methodName string, ) *AddClientSecretAzureAuthenticationMethodRequest`

NewAddClientSecretAzureAuthenticationMethodRequest instantiates a new AddClientSecretAzureAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClientSecretAzureAuthenticationMethodRequestWithDefaults

`func NewAddClientSecretAzureAuthenticationMethodRequestWithDefaults() *AddClientSecretAzureAuthenticationMethodRequest`

NewAddClientSecretAzureAuthenticationMethodRequestWithDefaults instantiates a new AddClientSecretAzureAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetSchemas() []EnumclientSecretAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetSchemasOk() (*[]EnumclientSecretAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddClientSecretAzureAuthenticationMethodRequest) SetSchemas(v []EnumclientSecretAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AddClientSecretAzureAuthenticationMethodRequest) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddClientSecretAzureAuthenticationMethodRequest) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddClientSecretAzureAuthenticationMethodRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetDescription

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddClientSecretAzureAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddClientSecretAzureAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMethodName

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddClientSecretAzureAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddClientSecretAzureAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


