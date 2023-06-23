# ReplicationServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumreplicationServerSchemaUrn**](EnumreplicationServerSchemaUrn.md) |  | [optional] 
**ReplicationServerID** | **int64** | Specifies a unique identifier for the Replication Server. | 
**ReplicationDBDirectory** | **string** | The path where the Replication Server stores all persistent information. | 
**JeProperty** | Pointer to **[]string** | Specifies the database and environment properties for the Berkeley DB Java Edition database for the replication changelog. | [optional] 
**ReplicationPurgeDelay** | Pointer to **string** | Changes are guaranteed to be maintained in the changelog database for at least this duration. Setting target-database-size can allow additional changes to be maintained up to the configured size on disk. | [optional] 
**TargetDatabaseSize** | Pointer to **string** | The replication changelog database is allowed to grow up to this size even if changes are older than the configured replication-purge-delay. | [optional] 
**ReplicationPort** | **int64** | The port on which this Replication Server waits for connections from other Replication Servers or Directory Server instances. | 
**ListenOnAllAddresses** | Pointer to **bool** | Indicates whether the Replication Server should listen on all addresses for this host. If set to FALSE, then the Replication Server will listen only to the address resolved from the hostname provided. | [optional] 
**CompressionCriteria** | Pointer to [**EnumreplicationServerCompressionCriteriaProp**](EnumreplicationServerCompressionCriteriaProp.md) |  | [optional] 
**HeartbeatInterval** | Pointer to **string** | Specifies the heartbeat interval that the Directory Server will use when communicating with Replication Servers. | [optional] 
**RemoteMonitorUpdateInterval** | Pointer to **string** | Specifies the duration that topology monitor data will be cached before it is requested again from a remote server. | [optional] 
**RestrictedDomain** | Pointer to **[]string** | Specifies the base DN of domains that are only replicated between server instances that belong to the same replication set. | [optional] 
**GatewayPriority** | **int64** | Specifies the gateway priority of the Replication Server in the current location. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewReplicationServerResponse

`func NewReplicationServerResponse(replicationServerID int64, replicationDBDirectory string, replicationPort int64, gatewayPriority int64, ) *ReplicationServerResponse`

NewReplicationServerResponse instantiates a new ReplicationServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationServerResponseWithDefaults

`func NewReplicationServerResponseWithDefaults() *ReplicationServerResponse`

NewReplicationServerResponseWithDefaults instantiates a new ReplicationServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplicationServerResponse) GetSchemas() []EnumreplicationServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationServerResponse) GetSchemasOk() (*[]EnumreplicationServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationServerResponse) SetSchemas(v []EnumreplicationServerSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ReplicationServerResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetReplicationServerID

`func (o *ReplicationServerResponse) GetReplicationServerID() int64`

GetReplicationServerID returns the ReplicationServerID field if non-nil, zero value otherwise.

### GetReplicationServerIDOk

`func (o *ReplicationServerResponse) GetReplicationServerIDOk() (*int64, bool)`

GetReplicationServerIDOk returns a tuple with the ReplicationServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationServerID

`func (o *ReplicationServerResponse) SetReplicationServerID(v int64)`

SetReplicationServerID sets ReplicationServerID field to given value.


### GetReplicationDBDirectory

`func (o *ReplicationServerResponse) GetReplicationDBDirectory() string`

GetReplicationDBDirectory returns the ReplicationDBDirectory field if non-nil, zero value otherwise.

### GetReplicationDBDirectoryOk

`func (o *ReplicationServerResponse) GetReplicationDBDirectoryOk() (*string, bool)`

GetReplicationDBDirectoryOk returns a tuple with the ReplicationDBDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDBDirectory

`func (o *ReplicationServerResponse) SetReplicationDBDirectory(v string)`

SetReplicationDBDirectory sets ReplicationDBDirectory field to given value.


### GetJeProperty

`func (o *ReplicationServerResponse) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *ReplicationServerResponse) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *ReplicationServerResponse) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *ReplicationServerResponse) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetReplicationPurgeDelay

`func (o *ReplicationServerResponse) GetReplicationPurgeDelay() string`

GetReplicationPurgeDelay returns the ReplicationPurgeDelay field if non-nil, zero value otherwise.

### GetReplicationPurgeDelayOk

`func (o *ReplicationServerResponse) GetReplicationPurgeDelayOk() (*string, bool)`

GetReplicationPurgeDelayOk returns a tuple with the ReplicationPurgeDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPurgeDelay

`func (o *ReplicationServerResponse) SetReplicationPurgeDelay(v string)`

SetReplicationPurgeDelay sets ReplicationPurgeDelay field to given value.

### HasReplicationPurgeDelay

`func (o *ReplicationServerResponse) HasReplicationPurgeDelay() bool`

HasReplicationPurgeDelay returns a boolean if a field has been set.

### GetTargetDatabaseSize

`func (o *ReplicationServerResponse) GetTargetDatabaseSize() string`

GetTargetDatabaseSize returns the TargetDatabaseSize field if non-nil, zero value otherwise.

### GetTargetDatabaseSizeOk

`func (o *ReplicationServerResponse) GetTargetDatabaseSizeOk() (*string, bool)`

GetTargetDatabaseSizeOk returns a tuple with the TargetDatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatabaseSize

`func (o *ReplicationServerResponse) SetTargetDatabaseSize(v string)`

SetTargetDatabaseSize sets TargetDatabaseSize field to given value.

### HasTargetDatabaseSize

`func (o *ReplicationServerResponse) HasTargetDatabaseSize() bool`

HasTargetDatabaseSize returns a boolean if a field has been set.

### GetReplicationPort

`func (o *ReplicationServerResponse) GetReplicationPort() int64`

GetReplicationPort returns the ReplicationPort field if non-nil, zero value otherwise.

### GetReplicationPortOk

`func (o *ReplicationServerResponse) GetReplicationPortOk() (*int64, bool)`

GetReplicationPortOk returns a tuple with the ReplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPort

`func (o *ReplicationServerResponse) SetReplicationPort(v int64)`

SetReplicationPort sets ReplicationPort field to given value.


### GetListenOnAllAddresses

`func (o *ReplicationServerResponse) GetListenOnAllAddresses() bool`

GetListenOnAllAddresses returns the ListenOnAllAddresses field if non-nil, zero value otherwise.

### GetListenOnAllAddressesOk

`func (o *ReplicationServerResponse) GetListenOnAllAddressesOk() (*bool, bool)`

GetListenOnAllAddressesOk returns a tuple with the ListenOnAllAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenOnAllAddresses

`func (o *ReplicationServerResponse) SetListenOnAllAddresses(v bool)`

SetListenOnAllAddresses sets ListenOnAllAddresses field to given value.

### HasListenOnAllAddresses

`func (o *ReplicationServerResponse) HasListenOnAllAddresses() bool`

HasListenOnAllAddresses returns a boolean if a field has been set.

### GetCompressionCriteria

`func (o *ReplicationServerResponse) GetCompressionCriteria() EnumreplicationServerCompressionCriteriaProp`

GetCompressionCriteria returns the CompressionCriteria field if non-nil, zero value otherwise.

### GetCompressionCriteriaOk

`func (o *ReplicationServerResponse) GetCompressionCriteriaOk() (*EnumreplicationServerCompressionCriteriaProp, bool)`

GetCompressionCriteriaOk returns a tuple with the CompressionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCriteria

`func (o *ReplicationServerResponse) SetCompressionCriteria(v EnumreplicationServerCompressionCriteriaProp)`

SetCompressionCriteria sets CompressionCriteria field to given value.

### HasCompressionCriteria

`func (o *ReplicationServerResponse) HasCompressionCriteria() bool`

HasCompressionCriteria returns a boolean if a field has been set.

### GetHeartbeatInterval

`func (o *ReplicationServerResponse) GetHeartbeatInterval() string`

GetHeartbeatInterval returns the HeartbeatInterval field if non-nil, zero value otherwise.

### GetHeartbeatIntervalOk

`func (o *ReplicationServerResponse) GetHeartbeatIntervalOk() (*string, bool)`

GetHeartbeatIntervalOk returns a tuple with the HeartbeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatInterval

`func (o *ReplicationServerResponse) SetHeartbeatInterval(v string)`

SetHeartbeatInterval sets HeartbeatInterval field to given value.

### HasHeartbeatInterval

`func (o *ReplicationServerResponse) HasHeartbeatInterval() bool`

HasHeartbeatInterval returns a boolean if a field has been set.

### GetRemoteMonitorUpdateInterval

`func (o *ReplicationServerResponse) GetRemoteMonitorUpdateInterval() string`

GetRemoteMonitorUpdateInterval returns the RemoteMonitorUpdateInterval field if non-nil, zero value otherwise.

### GetRemoteMonitorUpdateIntervalOk

`func (o *ReplicationServerResponse) GetRemoteMonitorUpdateIntervalOk() (*string, bool)`

GetRemoteMonitorUpdateIntervalOk returns a tuple with the RemoteMonitorUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMonitorUpdateInterval

`func (o *ReplicationServerResponse) SetRemoteMonitorUpdateInterval(v string)`

SetRemoteMonitorUpdateInterval sets RemoteMonitorUpdateInterval field to given value.

### HasRemoteMonitorUpdateInterval

`func (o *ReplicationServerResponse) HasRemoteMonitorUpdateInterval() bool`

HasRemoteMonitorUpdateInterval returns a boolean if a field has been set.

### GetRestrictedDomain

`func (o *ReplicationServerResponse) GetRestrictedDomain() []string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *ReplicationServerResponse) GetRestrictedDomainOk() (*[]string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *ReplicationServerResponse) SetRestrictedDomain(v []string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *ReplicationServerResponse) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetGatewayPriority

`func (o *ReplicationServerResponse) GetGatewayPriority() int64`

GetGatewayPriority returns the GatewayPriority field if non-nil, zero value otherwise.

### GetGatewayPriorityOk

`func (o *ReplicationServerResponse) GetGatewayPriorityOk() (*int64, bool)`

GetGatewayPriorityOk returns a tuple with the GatewayPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPriority

`func (o *ReplicationServerResponse) SetGatewayPriority(v int64)`

SetGatewayPriority sets GatewayPriority field to given value.


### GetMeta

`func (o *ReplicationServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReplicationServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReplicationServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReplicationServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ReplicationServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


