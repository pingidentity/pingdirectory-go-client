# GetSynchronizationProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumcustomSynchronizationProviderSchemaUrn**](EnumcustomSynchronizationProviderSchemaUrn.md) |  | 
**Id** | **string** | Name of the Synchronization Provider | 
**NumUpdateReplayThreads** | Pointer to **int32** | Specifies the number of update replay threads. | [optional] 
**Description** | Pointer to **string** | A description for this Synchronization Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Synchronization Provider is enabled for use. | 

## Methods

### NewGetSynchronizationProvider200Response

`func NewGetSynchronizationProvider200Response(schemas []EnumcustomSynchronizationProviderSchemaUrn, id string, enabled bool, ) *GetSynchronizationProvider200Response`

NewGetSynchronizationProvider200Response instantiates a new GetSynchronizationProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSynchronizationProvider200ResponseWithDefaults

`func NewGetSynchronizationProvider200ResponseWithDefaults() *GetSynchronizationProvider200Response`

NewGetSynchronizationProvider200ResponseWithDefaults instantiates a new GetSynchronizationProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetSynchronizationProvider200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSynchronizationProvider200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSynchronizationProvider200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetSynchronizationProvider200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetSynchronizationProvider200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetSynchronizationProvider200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetSynchronizationProvider200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetSynchronizationProvider200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GetSynchronizationProvider200Response) GetSchemas() []EnumcustomSynchronizationProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetSynchronizationProvider200Response) GetSchemasOk() (*[]EnumcustomSynchronizationProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetSynchronizationProvider200Response) SetSchemas(v []EnumcustomSynchronizationProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetSynchronizationProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSynchronizationProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSynchronizationProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetNumUpdateReplayThreads

`func (o *GetSynchronizationProvider200Response) GetNumUpdateReplayThreads() int32`

GetNumUpdateReplayThreads returns the NumUpdateReplayThreads field if non-nil, zero value otherwise.

### GetNumUpdateReplayThreadsOk

`func (o *GetSynchronizationProvider200Response) GetNumUpdateReplayThreadsOk() (*int32, bool)`

GetNumUpdateReplayThreadsOk returns a tuple with the NumUpdateReplayThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpdateReplayThreads

`func (o *GetSynchronizationProvider200Response) SetNumUpdateReplayThreads(v int32)`

SetNumUpdateReplayThreads sets NumUpdateReplayThreads field to given value.

### HasNumUpdateReplayThreads

`func (o *GetSynchronizationProvider200Response) HasNumUpdateReplayThreads() bool`

HasNumUpdateReplayThreads returns a boolean if a field has been set.

### GetDescription

`func (o *GetSynchronizationProvider200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSynchronizationProvider200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSynchronizationProvider200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSynchronizationProvider200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetSynchronizationProvider200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSynchronizationProvider200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSynchronizationProvider200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


