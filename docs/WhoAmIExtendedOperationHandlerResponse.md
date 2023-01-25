# WhoAmIExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumwhoAmIExtendedOperationHandlerSchemaUrn**](EnumwhoAmIExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Extended Operation Handler | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewWhoAmIExtendedOperationHandlerResponse

`func NewWhoAmIExtendedOperationHandlerResponse(schemas []EnumwhoAmIExtendedOperationHandlerSchemaUrn, id string, enabled bool, ) *WhoAmIExtendedOperationHandlerResponse`

NewWhoAmIExtendedOperationHandlerResponse instantiates a new WhoAmIExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoAmIExtendedOperationHandlerResponseWithDefaults

`func NewWhoAmIExtendedOperationHandlerResponseWithDefaults() *WhoAmIExtendedOperationHandlerResponse`

NewWhoAmIExtendedOperationHandlerResponseWithDefaults instantiates a new WhoAmIExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *WhoAmIExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WhoAmIExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WhoAmIExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WhoAmIExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *WhoAmIExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *WhoAmIExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *WhoAmIExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *WhoAmIExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *WhoAmIExtendedOperationHandlerResponse) GetSchemas() []EnumwhoAmIExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WhoAmIExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumwhoAmIExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WhoAmIExtendedOperationHandlerResponse) SetSchemas(v []EnumwhoAmIExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *WhoAmIExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhoAmIExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhoAmIExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *WhoAmIExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhoAmIExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhoAmIExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhoAmIExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WhoAmIExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WhoAmIExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WhoAmIExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


