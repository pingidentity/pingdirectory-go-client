# ClientSecretAzureAuthenticationMethodShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumclientSecretAzureAuthenticationMethodSchemaUrn**](EnumclientSecretAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**ClientSecret** | **string** | The client secret to use to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 

## Methods

### NewClientSecretAzureAuthenticationMethodShared

`func NewClientSecretAzureAuthenticationMethodShared(schemas []EnumclientSecretAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, clientSecret string, ) *ClientSecretAzureAuthenticationMethodShared`

NewClientSecretAzureAuthenticationMethodShared instantiates a new ClientSecretAzureAuthenticationMethodShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSecretAzureAuthenticationMethodSharedWithDefaults

`func NewClientSecretAzureAuthenticationMethodSharedWithDefaults() *ClientSecretAzureAuthenticationMethodShared`

NewClientSecretAzureAuthenticationMethodSharedWithDefaults instantiates a new ClientSecretAzureAuthenticationMethodShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ClientSecretAzureAuthenticationMethodShared) GetSchemas() []EnumclientSecretAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ClientSecretAzureAuthenticationMethodShared) GetSchemasOk() (*[]EnumclientSecretAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ClientSecretAzureAuthenticationMethodShared) SetSchemas(v []EnumclientSecretAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *ClientSecretAzureAuthenticationMethodShared) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *ClientSecretAzureAuthenticationMethodShared) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *ClientSecretAzureAuthenticationMethodShared) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *ClientSecretAzureAuthenticationMethodShared) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *ClientSecretAzureAuthenticationMethodShared) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *ClientSecretAzureAuthenticationMethodShared) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *ClientSecretAzureAuthenticationMethodShared) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientSecretAzureAuthenticationMethodShared) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientSecretAzureAuthenticationMethodShared) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetDescription

`func (o *ClientSecretAzureAuthenticationMethodShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientSecretAzureAuthenticationMethodShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientSecretAzureAuthenticationMethodShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientSecretAzureAuthenticationMethodShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


