# UsernamePasswordAzureAuthenticationMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn**](EnumusernamePasswordAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Azure Authentication Method | 

## Methods

### NewUsernamePasswordAzureAuthenticationMethodResponse

`func NewUsernamePasswordAzureAuthenticationMethodResponse(schemas []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, username string, password string, id string, ) *UsernamePasswordAzureAuthenticationMethodResponse`

NewUsernamePasswordAzureAuthenticationMethodResponse instantiates a new UsernamePasswordAzureAuthenticationMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordAzureAuthenticationMethodResponseWithDefaults

`func NewUsernamePasswordAzureAuthenticationMethodResponseWithDefaults() *UsernamePasswordAzureAuthenticationMethodResponse`

NewUsernamePasswordAzureAuthenticationMethodResponseWithDefaults instantiates a new UsernamePasswordAzureAuthenticationMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetSchemas() []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetSchemasOk() (*[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetSchemas(v []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetUsername

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDescription

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsernamePasswordAzureAuthenticationMethodResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


