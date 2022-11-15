# UsernamePasswordAzureAuthenticationMethodShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn**](EnumusernamePasswordAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 

## Methods

### NewUsernamePasswordAzureAuthenticationMethodShared

`func NewUsernamePasswordAzureAuthenticationMethodShared(schemas []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, username string, password string, ) *UsernamePasswordAzureAuthenticationMethodShared`

NewUsernamePasswordAzureAuthenticationMethodShared instantiates a new UsernamePasswordAzureAuthenticationMethodShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordAzureAuthenticationMethodSharedWithDefaults

`func NewUsernamePasswordAzureAuthenticationMethodSharedWithDefaults() *UsernamePasswordAzureAuthenticationMethodShared`

NewUsernamePasswordAzureAuthenticationMethodSharedWithDefaults instantiates a new UsernamePasswordAzureAuthenticationMethodShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetSchemas() []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetSchemasOk() (*[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UsernamePasswordAzureAuthenticationMethodShared) SetSchemas(v []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *UsernamePasswordAzureAuthenticationMethodShared) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *UsernamePasswordAzureAuthenticationMethodShared) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetUsername

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsernamePasswordAzureAuthenticationMethodShared) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsernamePasswordAzureAuthenticationMethodShared) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDescription

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsernamePasswordAzureAuthenticationMethodShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsernamePasswordAzureAuthenticationMethodShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UsernamePasswordAzureAuthenticationMethodShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


