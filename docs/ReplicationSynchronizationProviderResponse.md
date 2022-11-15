# ReplicationSynchronizationProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumreplicationSynchronizationProviderSchemaUrn**](EnumreplicationSynchronizationProviderSchemaUrn.md) |  | [optional] 
**NumUpdateReplayThreads** | Pointer to **int32** | Specifies the number of update replay threads. | [optional] 
**Description** | Pointer to **string** | A description for this Synchronization Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Synchronization Provider is enabled for use. | 

## Methods

### NewReplicationSynchronizationProviderResponse

`func NewReplicationSynchronizationProviderResponse(enabled bool, ) *ReplicationSynchronizationProviderResponse`

NewReplicationSynchronizationProviderResponse instantiates a new ReplicationSynchronizationProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationSynchronizationProviderResponseWithDefaults

`func NewReplicationSynchronizationProviderResponseWithDefaults() *ReplicationSynchronizationProviderResponse`

NewReplicationSynchronizationProviderResponseWithDefaults instantiates a new ReplicationSynchronizationProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplicationSynchronizationProviderResponse) GetSchemas() []EnumreplicationSynchronizationProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationSynchronizationProviderResponse) GetSchemasOk() (*[]EnumreplicationSynchronizationProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationSynchronizationProviderResponse) SetSchemas(v []EnumreplicationSynchronizationProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ReplicationSynchronizationProviderResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetNumUpdateReplayThreads

`func (o *ReplicationSynchronizationProviderResponse) GetNumUpdateReplayThreads() int32`

GetNumUpdateReplayThreads returns the NumUpdateReplayThreads field if non-nil, zero value otherwise.

### GetNumUpdateReplayThreadsOk

`func (o *ReplicationSynchronizationProviderResponse) GetNumUpdateReplayThreadsOk() (*int32, bool)`

GetNumUpdateReplayThreadsOk returns a tuple with the NumUpdateReplayThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpdateReplayThreads

`func (o *ReplicationSynchronizationProviderResponse) SetNumUpdateReplayThreads(v int32)`

SetNumUpdateReplayThreads sets NumUpdateReplayThreads field to given value.

### HasNumUpdateReplayThreads

`func (o *ReplicationSynchronizationProviderResponse) HasNumUpdateReplayThreads() bool`

HasNumUpdateReplayThreads returns a boolean if a field has been set.

### GetDescription

`func (o *ReplicationSynchronizationProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplicationSynchronizationProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplicationSynchronizationProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplicationSynchronizationProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReplicationSynchronizationProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplicationSynchronizationProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplicationSynchronizationProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


