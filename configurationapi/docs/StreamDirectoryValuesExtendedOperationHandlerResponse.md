# StreamDirectoryValuesExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumstreamDirectoryValuesExtendedOperationHandlerSchemaUrn**](EnumstreamDirectoryValuesExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Extended Operation Handler | 
**ValuesPerStreamResponse** | Pointer to **int64** | The maximum number of values to include per response when responding to a stream values extended request, when the client does not specify a value. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewStreamDirectoryValuesExtendedOperationHandlerResponse

`func NewStreamDirectoryValuesExtendedOperationHandlerResponse(schemas []EnumstreamDirectoryValuesExtendedOperationHandlerSchemaUrn, id string, enabled bool, ) *StreamDirectoryValuesExtendedOperationHandlerResponse`

NewStreamDirectoryValuesExtendedOperationHandlerResponse instantiates a new StreamDirectoryValuesExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamDirectoryValuesExtendedOperationHandlerResponseWithDefaults

`func NewStreamDirectoryValuesExtendedOperationHandlerResponseWithDefaults() *StreamDirectoryValuesExtendedOperationHandlerResponse`

NewStreamDirectoryValuesExtendedOperationHandlerResponseWithDefaults instantiates a new StreamDirectoryValuesExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetSchemas() []EnumstreamDirectoryValuesExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumstreamDirectoryValuesExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetSchemas(v []EnumstreamDirectoryValuesExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetValuesPerStreamResponse

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetValuesPerStreamResponse() int64`

GetValuesPerStreamResponse returns the ValuesPerStreamResponse field if non-nil, zero value otherwise.

### GetValuesPerStreamResponseOk

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetValuesPerStreamResponseOk() (*int64, bool)`

GetValuesPerStreamResponseOk returns a tuple with the ValuesPerStreamResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPerStreamResponse

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetValuesPerStreamResponse(v int64)`

SetValuesPerStreamResponse sets ValuesPerStreamResponse field to given value.

### HasValuesPerStreamResponse

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) HasValuesPerStreamResponse() bool`

HasValuesPerStreamResponse returns a boolean if a field has been set.

### GetDescription

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *StreamDirectoryValuesExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


