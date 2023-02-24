# AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn**](EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn.md) |  | 
**PollingInterval** | Pointer to **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | [optional] 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**BaseDN** | Pointer to **string** | Only entries located within the subtree specified by this base DN are eligible for purging. | [optional] 
**MaxUpdatesPerSecond** | Pointer to **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | [optional] 
**NumDeleteThreads** | Pointer to **int32** | The number of threads used to delete expired entries. | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest

`func NewAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest(pluginName string, schemas []EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn, enabled bool, ) *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest`

NewAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest instantiates a new AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequestWithDefaults

`func NewAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequestWithDefaults() *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest`

NewAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequestWithDefaults instantiates a new AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetSchemas() []EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetSchemasOk() (*[]EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetSchemas(v []EnumcleanUpExpiredPingfederatePersistentAccessGrantsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPollingInterval

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPeerServerPriorityIndex

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.

### HasMaxUpdatesPerSecond

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) HasMaxUpdatesPerSecond() bool`

HasMaxUpdatesPerSecond returns a boolean if a field has been set.

### GetNumDeleteThreads

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.

### HasNumDeleteThreads

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) HasNumDeleteThreads() bool`

HasNumDeleteThreads returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


