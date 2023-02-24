# AddLdifConnectionHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Connection Handler | 
**Schemas** | [**[]EnumldifConnectionHandlerSchemaUrn**](EnumldifConnectionHandlerSchemaUrn.md) |  | 
**AllowedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are allowed to establish connections to this connection handler. | [optional] 
**DeniedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are not allowed to establish connections to this connection handler. | [optional] 
**LdifDirectory** | Pointer to **string** | Specifies the path to the directory in which the LDIF files should be placed. | [optional] 
**PollInterval** | Pointer to **string** | Specifies how frequently the LDIF connection handler should check the LDIF directory to determine whether a new LDIF file has been added. | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 

## Methods

### NewAddLdifConnectionHandlerRequest

`func NewAddLdifConnectionHandlerRequest(handlerName string, schemas []EnumldifConnectionHandlerSchemaUrn, enabled bool, ) *AddLdifConnectionHandlerRequest`

NewAddLdifConnectionHandlerRequest instantiates a new AddLdifConnectionHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLdifConnectionHandlerRequestWithDefaults

`func NewAddLdifConnectionHandlerRequestWithDefaults() *AddLdifConnectionHandlerRequest`

NewAddLdifConnectionHandlerRequestWithDefaults instantiates a new AddLdifConnectionHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddLdifConnectionHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddLdifConnectionHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddLdifConnectionHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddLdifConnectionHandlerRequest) GetSchemas() []EnumldifConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLdifConnectionHandlerRequest) GetSchemasOk() (*[]EnumldifConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLdifConnectionHandlerRequest) SetSchemas(v []EnumldifConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedClient

`func (o *AddLdifConnectionHandlerRequest) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *AddLdifConnectionHandlerRequest) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *AddLdifConnectionHandlerRequest) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *AddLdifConnectionHandlerRequest) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *AddLdifConnectionHandlerRequest) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *AddLdifConnectionHandlerRequest) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *AddLdifConnectionHandlerRequest) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *AddLdifConnectionHandlerRequest) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.

### GetLdifDirectory

`func (o *AddLdifConnectionHandlerRequest) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *AddLdifConnectionHandlerRequest) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *AddLdifConnectionHandlerRequest) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.

### HasLdifDirectory

`func (o *AddLdifConnectionHandlerRequest) HasLdifDirectory() bool`

HasLdifDirectory returns a boolean if a field has been set.

### GetPollInterval

`func (o *AddLdifConnectionHandlerRequest) GetPollInterval() string`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *AddLdifConnectionHandlerRequest) GetPollIntervalOk() (*string, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *AddLdifConnectionHandlerRequest) SetPollInterval(v string)`

SetPollInterval sets PollInterval field to given value.

### HasPollInterval

`func (o *AddLdifConnectionHandlerRequest) HasPollInterval() bool`

HasPollInterval returns a boolean if a field has been set.

### GetDescription

`func (o *AddLdifConnectionHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLdifConnectionHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLdifConnectionHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLdifConnectionHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLdifConnectionHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLdifConnectionHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLdifConnectionHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


