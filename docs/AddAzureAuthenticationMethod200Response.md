# AddAzureAuthenticationMethod200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Azure Authentication Method | 
**Schemas** | [**[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn**](EnumusernamePasswordAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**ClientSecret** | **string** | The client secret to use to authenticate. | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 

## Methods

### NewAddAzureAuthenticationMethod200Response

`func NewAddAzureAuthenticationMethod200Response(id string, schemas []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, clientSecret string, username string, password string, ) *AddAzureAuthenticationMethod200Response`

NewAddAzureAuthenticationMethod200Response instantiates a new AddAzureAuthenticationMethod200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAzureAuthenticationMethod200ResponseWithDefaults

`func NewAddAzureAuthenticationMethod200ResponseWithDefaults() *AddAzureAuthenticationMethod200Response`

NewAddAzureAuthenticationMethod200ResponseWithDefaults instantiates a new AddAzureAuthenticationMethod200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddAzureAuthenticationMethod200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddAzureAuthenticationMethod200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddAzureAuthenticationMethod200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddAzureAuthenticationMethod200Response) GetSchemas() []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAzureAuthenticationMethod200Response) GetSchemasOk() (*[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAzureAuthenticationMethod200Response) SetSchemas(v []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *AddAzureAuthenticationMethod200Response) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AddAzureAuthenticationMethod200Response) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AddAzureAuthenticationMethod200Response) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *AddAzureAuthenticationMethod200Response) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddAzureAuthenticationMethod200Response) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddAzureAuthenticationMethod200Response) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetDescription

`func (o *AddAzureAuthenticationMethod200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAzureAuthenticationMethod200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAzureAuthenticationMethod200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAzureAuthenticationMethod200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddAzureAuthenticationMethod200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddAzureAuthenticationMethod200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddAzureAuthenticationMethod200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddAzureAuthenticationMethod200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetClientSecret

`func (o *AddAzureAuthenticationMethod200Response) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddAzureAuthenticationMethod200Response) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddAzureAuthenticationMethod200Response) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetUsername

`func (o *AddAzureAuthenticationMethod200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddAzureAuthenticationMethod200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddAzureAuthenticationMethod200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddAzureAuthenticationMethod200Response) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddAzureAuthenticationMethod200Response) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddAzureAuthenticationMethod200Response) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


