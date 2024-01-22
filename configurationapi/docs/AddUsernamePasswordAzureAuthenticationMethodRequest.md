# AddUsernamePasswordAzureAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn**](EnumusernamePasswordAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | **string** | The tenant ID to use to authenticate. | 
**ClientID** | **string** | The client ID to use to authenticate. | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | **string** | The password for the user to authenticate. | 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 
**MethodName** | **string** | Name of the new Azure Authentication Method | 

## Methods

### NewAddUsernamePasswordAzureAuthenticationMethodRequest

`func NewAddUsernamePasswordAzureAuthenticationMethodRequest(schemas []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, tenantID string, clientID string, username string, password string, methodName string, ) *AddUsernamePasswordAzureAuthenticationMethodRequest`

NewAddUsernamePasswordAzureAuthenticationMethodRequest instantiates a new AddUsernamePasswordAzureAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUsernamePasswordAzureAuthenticationMethodRequestWithDefaults

`func NewAddUsernamePasswordAzureAuthenticationMethodRequestWithDefaults() *AddUsernamePasswordAzureAuthenticationMethodRequest`

NewAddUsernamePasswordAzureAuthenticationMethodRequestWithDefaults instantiates a new AddUsernamePasswordAzureAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetSchemas() []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetSchemasOk() (*[]EnumusernamePasswordAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetSchemas(v []EnumusernamePasswordAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetClientID

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetUsername

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDescription

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMethodName

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddUsernamePasswordAzureAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


