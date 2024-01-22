# WaitNotifyWorkQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumwaitNotifyWorkQueueSchemaUrn**](EnumwaitNotifyWorkQueueSchemaUrn.md) |  | 
**NumWorkerThreads** | Pointer to **int64** | Specifies the number of worker threads that should be used within the server in order to process requested operations. | [optional] 
**MaxWorkQueueCapacity** | Pointer to **int64** | Specifies the maximum number of pending operations that may be held in the work queue at any given time. | [optional] 

## Methods

### NewWaitNotifyWorkQueueResponse

`func NewWaitNotifyWorkQueueResponse(schemas []EnumwaitNotifyWorkQueueSchemaUrn, ) *WaitNotifyWorkQueueResponse`

NewWaitNotifyWorkQueueResponse instantiates a new WaitNotifyWorkQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitNotifyWorkQueueResponseWithDefaults

`func NewWaitNotifyWorkQueueResponseWithDefaults() *WaitNotifyWorkQueueResponse`

NewWaitNotifyWorkQueueResponseWithDefaults instantiates a new WaitNotifyWorkQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *WaitNotifyWorkQueueResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WaitNotifyWorkQueueResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WaitNotifyWorkQueueResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WaitNotifyWorkQueueResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *WaitNotifyWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *WaitNotifyWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *WaitNotifyWorkQueueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *WaitNotifyWorkQueueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *WaitNotifyWorkQueueResponse) GetSchemas() []EnumwaitNotifyWorkQueueSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WaitNotifyWorkQueueResponse) GetSchemasOk() (*[]EnumwaitNotifyWorkQueueSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WaitNotifyWorkQueueResponse) SetSchemas(v []EnumwaitNotifyWorkQueueSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetNumWorkerThreads

`func (o *WaitNotifyWorkQueueResponse) GetNumWorkerThreads() int64`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *WaitNotifyWorkQueueResponse) GetNumWorkerThreadsOk() (*int64, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *WaitNotifyWorkQueueResponse) SetNumWorkerThreads(v int64)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.

### HasNumWorkerThreads

`func (o *WaitNotifyWorkQueueResponse) HasNumWorkerThreads() bool`

HasNumWorkerThreads returns a boolean if a field has been set.

### GetMaxWorkQueueCapacity

`func (o *WaitNotifyWorkQueueResponse) GetMaxWorkQueueCapacity() int64`

GetMaxWorkQueueCapacity returns the MaxWorkQueueCapacity field if non-nil, zero value otherwise.

### GetMaxWorkQueueCapacityOk

`func (o *WaitNotifyWorkQueueResponse) GetMaxWorkQueueCapacityOk() (*int64, bool)`

GetMaxWorkQueueCapacityOk returns a tuple with the MaxWorkQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkQueueCapacity

`func (o *WaitNotifyWorkQueueResponse) SetMaxWorkQueueCapacity(v int64)`

SetMaxWorkQueueCapacity sets MaxWorkQueueCapacity field to given value.

### HasMaxWorkQueueCapacity

`func (o *WaitNotifyWorkQueueResponse) HasMaxWorkQueueCapacity() bool`

HasMaxWorkQueueCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


