# MultiUpdateExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnummultiUpdateExtendedOperationHandlerSchemaUrn**](EnummultiUpdateExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Extended Operation Handler | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewMultiUpdateExtendedOperationHandlerResponse

`func NewMultiUpdateExtendedOperationHandlerResponse(schemas []EnummultiUpdateExtendedOperationHandlerSchemaUrn, id string, enabled bool, ) *MultiUpdateExtendedOperationHandlerResponse`

NewMultiUpdateExtendedOperationHandlerResponse instantiates a new MultiUpdateExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiUpdateExtendedOperationHandlerResponseWithDefaults

`func NewMultiUpdateExtendedOperationHandlerResponseWithDefaults() *MultiUpdateExtendedOperationHandlerResponse`

NewMultiUpdateExtendedOperationHandlerResponseWithDefaults instantiates a new MultiUpdateExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MultiUpdateExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MultiUpdateExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetSchemas() []EnummultiUpdateExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnummultiUpdateExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetSchemas(v []EnummultiUpdateExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultiUpdateExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


