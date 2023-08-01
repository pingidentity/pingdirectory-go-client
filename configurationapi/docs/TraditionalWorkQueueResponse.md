# TraditionalWorkQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtraditionalWorkQueueSchemaUrn**](EnumtraditionalWorkQueueSchemaUrn.md) |  | 
**NumWorkerThreads** | **int64** | Specifies the number of worker threads to be used for processing operations placed in the queue. | 
**MaxWorkQueueCapacity** | Pointer to **int64** | Specifies the maximum number of queued operations that can be in the work queue at any given time. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewTraditionalWorkQueueResponse

`func NewTraditionalWorkQueueResponse(schemas []EnumtraditionalWorkQueueSchemaUrn, numWorkerThreads int64, ) *TraditionalWorkQueueResponse`

NewTraditionalWorkQueueResponse instantiates a new TraditionalWorkQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraditionalWorkQueueResponseWithDefaults

`func NewTraditionalWorkQueueResponseWithDefaults() *TraditionalWorkQueueResponse`

NewTraditionalWorkQueueResponseWithDefaults instantiates a new TraditionalWorkQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TraditionalWorkQueueResponse) GetSchemas() []EnumtraditionalWorkQueueSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TraditionalWorkQueueResponse) GetSchemasOk() (*[]EnumtraditionalWorkQueueSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TraditionalWorkQueueResponse) SetSchemas(v []EnumtraditionalWorkQueueSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetNumWorkerThreads

`func (o *TraditionalWorkQueueResponse) GetNumWorkerThreads() int64`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *TraditionalWorkQueueResponse) GetNumWorkerThreadsOk() (*int64, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *TraditionalWorkQueueResponse) SetNumWorkerThreads(v int64)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.


### GetMaxWorkQueueCapacity

`func (o *TraditionalWorkQueueResponse) GetMaxWorkQueueCapacity() int64`

GetMaxWorkQueueCapacity returns the MaxWorkQueueCapacity field if non-nil, zero value otherwise.

### GetMaxWorkQueueCapacityOk

`func (o *TraditionalWorkQueueResponse) GetMaxWorkQueueCapacityOk() (*int64, bool)`

GetMaxWorkQueueCapacityOk returns a tuple with the MaxWorkQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkQueueCapacity

`func (o *TraditionalWorkQueueResponse) SetMaxWorkQueueCapacity(v int64)`

SetMaxWorkQueueCapacity sets MaxWorkQueueCapacity field to given value.

### HasMaxWorkQueueCapacity

`func (o *TraditionalWorkQueueResponse) HasMaxWorkQueueCapacity() bool`

HasMaxWorkQueueCapacity returns a boolean if a field has been set.

### GetMeta

`func (o *TraditionalWorkQueueResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TraditionalWorkQueueResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TraditionalWorkQueueResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TraditionalWorkQueueResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TraditionalWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TraditionalWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TraditionalWorkQueueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TraditionalWorkQueueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


