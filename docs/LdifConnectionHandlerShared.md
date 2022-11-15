# LdifConnectionHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumldifConnectionHandlerSchemaUrn**](EnumldifConnectionHandlerSchemaUrn.md) |  | 
**AllowedClient** | Pointer to **[]string** |  | [optional] 
**DeniedClient** | Pointer to **[]string** |  | [optional] 
**LdifDirectory** | **string** | Specifies the path to the directory in which the LDIF files should be placed. | 
**PollInterval** | **string** | Specifies how frequently the LDIF connection handler should check the LDIF directory to determine whether a new LDIF file has been added. | 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 

## Methods

### NewLdifConnectionHandlerShared

`func NewLdifConnectionHandlerShared(schemas []EnumldifConnectionHandlerSchemaUrn, ldifDirectory string, pollInterval string, enabled bool, ) *LdifConnectionHandlerShared`

NewLdifConnectionHandlerShared instantiates a new LdifConnectionHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdifConnectionHandlerSharedWithDefaults

`func NewLdifConnectionHandlerSharedWithDefaults() *LdifConnectionHandlerShared`

NewLdifConnectionHandlerSharedWithDefaults instantiates a new LdifConnectionHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LdifConnectionHandlerShared) GetSchemas() []EnumldifConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdifConnectionHandlerShared) GetSchemasOk() (*[]EnumldifConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdifConnectionHandlerShared) SetSchemas(v []EnumldifConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedClient

`func (o *LdifConnectionHandlerShared) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *LdifConnectionHandlerShared) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *LdifConnectionHandlerShared) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *LdifConnectionHandlerShared) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *LdifConnectionHandlerShared) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *LdifConnectionHandlerShared) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *LdifConnectionHandlerShared) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *LdifConnectionHandlerShared) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.

### GetLdifDirectory

`func (o *LdifConnectionHandlerShared) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *LdifConnectionHandlerShared) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *LdifConnectionHandlerShared) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.


### GetPollInterval

`func (o *LdifConnectionHandlerShared) GetPollInterval() string`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *LdifConnectionHandlerShared) GetPollIntervalOk() (*string, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *LdifConnectionHandlerShared) SetPollInterval(v string)`

SetPollInterval sets PollInterval field to given value.


### GetDescription

`func (o *LdifConnectionHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdifConnectionHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdifConnectionHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdifConnectionHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LdifConnectionHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LdifConnectionHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LdifConnectionHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


