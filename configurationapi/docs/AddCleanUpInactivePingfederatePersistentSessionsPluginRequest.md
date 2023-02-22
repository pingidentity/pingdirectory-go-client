# AddCleanUpInactivePingfederatePersistentSessionsPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn**](EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn.md) |  | 
**ExpirationOffset** | **string** | Sessions whose last activity timestamp is older than this offset will be removed. | 
**PollingInterval** | **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**BaseDN** | Pointer to **string** | Only entries located within the subtree specified by this base DN are eligible for purging. | [optional] 
**MaxUpdatesPerSecond** | **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | 
**NumDeleteThreads** | **int32** | The number of threads used to delete expired entries. | 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequest

`func NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequest(pluginName string, schemas []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn, expirationOffset string, pollingInterval string, maxUpdatesPerSecond int32, numDeleteThreads int32, enabled bool, ) *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest`

NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequest instantiates a new AddCleanUpInactivePingfederatePersistentSessionsPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequestWithDefaults

`func NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequestWithDefaults() *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest`

NewAddCleanUpInactivePingfederatePersistentSessionsPluginRequestWithDefaults instantiates a new AddCleanUpInactivePingfederatePersistentSessionsPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetSchemas() []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetSchemasOk() (*[]EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetSchemas(v []EnumcleanUpInactivePingfederatePersistentSessionsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExpirationOffset

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetExpirationOffset() string`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetExpirationOffsetOk() (*string, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetExpirationOffset(v string)`

SetExpirationOffset sets ExpirationOffset field to given value.


### GetPollingInterval

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetPeerServerPriorityIndex

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetNumDeleteThreads

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetEnabled

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


