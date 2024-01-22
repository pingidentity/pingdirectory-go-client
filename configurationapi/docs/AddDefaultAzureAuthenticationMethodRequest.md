# AddDefaultAzureAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdefaultAzureAuthenticationMethodSchemaUrn**](EnumdefaultAzureAuthenticationMethodSchemaUrn.md) |  | 
**TenantID** | Pointer to **string** | The tenant ID to use to authenticate. If this is not provided, then it will be obtained from the AZURE_TENANT_ID environment variable. | [optional] 
**ClientID** | Pointer to **string** | The client ID to use to authenticate. If this is not provided, then it will be obtained from the AZURE_CLIENT_ID | [optional] 
**Description** | Pointer to **string** | A description for this Azure Authentication Method | [optional] 
**MethodName** | **string** | Name of the new Azure Authentication Method | 

## Methods

### NewAddDefaultAzureAuthenticationMethodRequest

`func NewAddDefaultAzureAuthenticationMethodRequest(schemas []EnumdefaultAzureAuthenticationMethodSchemaUrn, methodName string, ) *AddDefaultAzureAuthenticationMethodRequest`

NewAddDefaultAzureAuthenticationMethodRequest instantiates a new AddDefaultAzureAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDefaultAzureAuthenticationMethodRequestWithDefaults

`func NewAddDefaultAzureAuthenticationMethodRequestWithDefaults() *AddDefaultAzureAuthenticationMethodRequest`

NewAddDefaultAzureAuthenticationMethodRequestWithDefaults instantiates a new AddDefaultAzureAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetSchemas() []EnumdefaultAzureAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetSchemasOk() (*[]EnumdefaultAzureAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDefaultAzureAuthenticationMethodRequest) SetSchemas(v []EnumdefaultAzureAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTenantID

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AddDefaultAzureAuthenticationMethodRequest) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *AddDefaultAzureAuthenticationMethodRequest) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### GetClientID

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddDefaultAzureAuthenticationMethodRequest) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *AddDefaultAzureAuthenticationMethodRequest) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### GetDescription

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDefaultAzureAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDefaultAzureAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMethodName

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddDefaultAzureAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddDefaultAzureAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


