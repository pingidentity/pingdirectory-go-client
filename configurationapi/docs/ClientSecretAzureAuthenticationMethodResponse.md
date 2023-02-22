# ClientSecretAzureAuthenticationMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Azure Authentication Method | 
**Schemas** | [**[]EnumclientSecretAzureAuthenticationMethodSchemaUrn**](EnumclientSecretAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**ClientSecret** | **string** | The client secret to use to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 

## Methods

### NewClientSecretAzureAuthenticationMethodResponse

`func NewClientSecretAzureAuthenticationMethodResponse(id string, schemas []EnumclientSecretAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, clientSecret string, ) *ClientSecretAzureAuthenticationMethodResponse`

NewClientSecretAzureAuthenticationMethodResponse instantiates a new ClientSecretAzureAuthenticationMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSecretAzureAuthenticationMethodResponseWithDefaults

`func NewClientSecretAzureAuthenticationMethodResponseWithDefaults() *ClientSecretAzureAuthenticationMethodResponse`

NewClientSecretAzureAuthenticationMethodResponseWithDefaults instantiates a new ClientSecretAzureAuthenticationMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClientSecretAzureAuthenticationMethodResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ClientSecretAzureAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetSchemas() []EnumclientSecretAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetSchemasOk() (*[]EnumclientSecretAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetSchemas(v []EnumclientSecretAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetDescription

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientSecretAzureAuthenticationMethodResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientSecretAzureAuthenticationMethodResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientSecretAzureAuthenticationMethodResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


