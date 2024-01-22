# GlobalConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnumglobalConfigurationSchemaUrn**](EnumglobalConfigurationSchemaUrn.md) |  | [optional] 
**InstanceName** | **string** | Specifies a name that may be used to uniquely identify this Directory Server instance among other instances in the environment. | 
**Location** | Pointer to **string** | Specifies the location for this Directory Server. Operations performed which involve communication with other servers may prefer servers in the same location to help ensure low-latency responses. | [optional] 
**ConfigurationServerGroup** | Pointer to **string** | When this property is set, changes made to this server using the console or dsconfig can be automatically applied to all servers in the specified server group. | [optional] 
**ForceAsMasterForMirroredData** | Pointer to **bool** | Indicates whether this server should be forced to assume the master role if no other suitable server is found to act as master or if multiple masters are detected. A master is only needed when changes are made to mirrored data, i.e. data specific to the topology itself and cluster-wide configuration data. | [optional] 
**EncryptData** | Pointer to **bool** | Indicates whether the Directory Server should encrypt the data that it stores in all components that support it. This may include certain types of backends (including local DB and large attribute backends), the LDAP changelog, and the replication server database. | [optional] 
**EncryptionSettingsCipherStreamProvider** | Pointer to **string** | Specifies the cipher stream provider that should be used to protect the contents of the encryption settings database. | [optional] 
**EncryptBackupsByDefault** | Pointer to **bool** | Indicates whether the server should encrypt backups by default. | [optional] 
**BackupEncryptionSettingsDefinitionID** | Pointer to **string** | The unique identifier for the encryption settings definition to use to generate the encryption key for encrypted backups by default. | [optional] 
**EncryptLDIFExportsByDefault** | Pointer to **bool** | Indicates whether the server should encrypt LDIF exports by default. | [optional] 
**LdifExportEncryptionSettingsDefinitionID** | Pointer to **string** | The unique identifier for the encryption settings definition to use to generate the encryption key for encrypted LDIF exports by default. | [optional] 
**AutomaticallyCompressEncryptedLDIFExports** | Pointer to **bool** | Indicates whether to automatically compress LDIF exports that are also encrypted. | [optional] 
**RedactSensitiveValuesInConfigLogs** | Pointer to **bool** | Indicates whether the values of sensitive configuration properties should be redacted when logging configuration changes, including in the configuration audit log, the error log, and the server.out log file. | [optional] 
**SensitiveAttribute** | Pointer to **[]string** | Provides the ability to indicate that some attributes should be considered sensitive and additional protection should be in place when interacting with those attributes. | [optional] 
**RejectInsecureRequests** | Pointer to **bool** | Indicates whether the Directory Server should reject any LDAP request (other than StartTLS) received from a client that is not using an encrypted connection. | [optional] 
**AllowedInsecureRequestCriteria** | Pointer to **string** | A set of criteria that may be used to match LDAP requests that may be permitted over an insecure connection even if reject-insecure-requests is true. Note that some types of requests will always be permitted, including StartTLS and start administrative session requests. | [optional] 
**RejectUnauthenticatedRequests** | Pointer to **bool** | Indicates whether the Directory Server should reject any LDAP request (other than bind or StartTLS requests) received from a client that has not yet been authenticated, whose last authentication attempt was unsuccessful, or whose last authentication attempt used anonymous authentication. | [optional] 
**AllowedUnauthenticatedRequestCriteria** | Pointer to **string** | A set of criteria that may be used to match LDAP requests that may be permitted over an unauthenticated connection even if reject-unauthenticated-requests is true. Note that some types of requests will always be permitted, including bind, StartTLS, and start administrative session requests. | [optional] 
**BindWithDNRequiresPassword** | Pointer to **bool** | Indicates whether the Directory Server should reject any simple bind request that contains a DN but no password. | [optional] 
**DisabledPrivilege** | Pointer to [**[]EnumglobalConfigurationDisabledPrivilegeProp**](EnumglobalConfigurationDisabledPrivilegeProp.md) |  | [optional] 
**DefaultPasswordPolicy** | **string** | Specifies the name of the password policy that is in effect for users whose entries do not specify an alternate password policy (either via a real or virtual attribute). | 
**MaximumUserDataPasswordPoliciesToCache** | Pointer to **int64** | Specifies the maximum number of password policies that are defined in the user data (that is, outside of the configuration) that the server should cache in memory for faster access. A value of zero indicates that the server should not cache any user data password policies. | [optional] 
**ProxiedAuthorizationIdentityMapper** | **string** | Specifies the name of the identity mapper to map authorization ID values (using the \&quot;u:\&quot; form) provided in the proxied authorization control to the corresponding user entry. | 
**VerifyEntryDigests** | Pointer to **bool** | Indicates whether the digest should always be verified whenever an entry containing a digest is decoded. If this is \&quot;true\&quot;, then if a digest exists, it will always be verified. Otherwise, the digest will be written when encoding entries but ignored when decoding entries but may still be available for other verification processing. | [optional] 
**AllowedInsecureTLSProtocol** | Pointer to [**[]EnumglobalConfigurationAllowedInsecureTLSProtocolProp**](EnumglobalConfigurationAllowedInsecureTLSProtocolProp.md) |  | [optional] 
**AllowInsecureLocalJMXConnections** | Pointer to **bool** | Indicates that processes attaching to this server&#39;s local JVM are allowed to access internal data through JMX without the authentication requirements that remote JMX connections are subject to. Please review and understand the data that this option will expose (such as cn&#x3D;monitor) to client applications to ensure there are no security concerns. | [optional] 
**DefaultInternalOperationClientConnectionPolicy** | Pointer to **string** | Specifies the client connection policy that will be used by default for internal operations. | [optional] 
**SizeLimit** | Pointer to **int64** | Specifies the maximum number of entries that the Directory Server should return to clients by default when processing a search operation. | [optional] 
**UnauthenticatedSizeLimit** | Pointer to **int64** | The size limit value that will apply for connections from unauthenticated clients. If this is not specified, then the value of the size-limit property will be applied for both authenticated and unauthenticated connections. | [optional] 
**TimeLimit** | Pointer to **string** | Specifies the maximum length of time that the Directory Server should be allowed to spend processing a search operation. | [optional] 
**UnauthenticatedTimeLimit** | Pointer to **string** | The time limit value that will apply for connections from unauthenticated clients. If this is not specified, then the value of the time-limit property will be applied for both authenticated and unauthenticated connections. | [optional] 
**IdleTimeLimit** | Pointer to **string** | Specifies the maximum length of time that a client connection may remain established since its last completed operation. | [optional] 
**UnauthenticatedIdleTimeLimit** | Pointer to **string** | The idle-time-limit limit value that will apply for connections from unauthenticated clients. If this is not specified, then the value of the idle-time-limit property will be applied for both authenticated and unauthenticated connections. | [optional] 
**LookthroughLimit** | Pointer to **int64** | Specifies the maximum number of entries that the Directory Server should \&quot;look through\&quot; in the course of processing a search request. | [optional] 
**UnauthenticatedLookthroughLimit** | Pointer to **int64** | The lookthrough limit value that will apply for connections from unauthenticated clients. If this is not specified, then the value of the lookthrough-limit property will be applied for both authenticated and unauthenticated connections. | [optional] 
**LdapJoinSizeLimit** | Pointer to **int64** | Specifies the maximum number of entries that may be directly joined with any individual search result entry. | [optional] 
**MaximumConcurrentConnections** | Pointer to **int64** | Specifies the maximum number of LDAP client connections which may be established to this Directory Server at the same time. | [optional] 
**MaximumConcurrentConnectionsPerIPAddress** | Pointer to **int64** | Specifies the maximum number of LDAP client connections originating from the same IP address which may be established to this Directory Server at the same time. | [optional] 
**MaximumConcurrentConnectionsPerBindDN** | Pointer to **int64** | Specifies the maximum number of LDAP client connections which may be established to this Directory Server at the same time and authenticated as the same user. | [optional] 
**MaximumConcurrentUnindexedSearches** | Pointer to **int64** | Specifies the maximum number of unindexed searches that may be in progress in this backend at any given time. Any unindexed searches requested while the maximum number of unindexed searches are already being processed will be rejected. A value of zero indicates that no limit will be enforced. | [optional] 
**MaximumAttributesPerAddRequest** | Pointer to **int64** | Specifies the maximum number of attributes that may be included in an add request. This property does not impose any limit on the number of values that an attribute may have. | [optional] 
**MaximumModificationsPerModifyRequest** | Pointer to **int64** | Specifies the maximum number of modifications that may be included in a modify request. This property does not impose any limit on the number of attribute values that a modification may have. | [optional] 
**BackgroundThreadForEachPersistentSearch** | Pointer to **bool** | Indicates whether the server should use a separate background thread for each persistent search. | [optional] 
**AllowAttributeNameExceptions** | Pointer to **bool** | Indicates whether the Directory Server should allow underscores in attribute names and allow attribute names to begin with numeric digits (both of which are violations of the LDAP standards). | [optional] 
**InvalidAttributeSyntaxBehavior** | Pointer to [**EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp**](EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp.md) |  | [optional] 
**PermitSyntaxViolationsForAttribute** | Pointer to **[]string** | Specifies a set of attribute types for which the server will permit values that do not conform to the associated attribute syntax. | [optional] 
**SingleStructuralObjectclassBehavior** | Pointer to [**EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp**](EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp.md) |  | [optional] 
**AttributesModifiableWithIgnoreNoUserModificationRequestControl** | Pointer to [**[]EnumglobalConfigurationAttributesModifiableWithIgnoreNoUserModificationRequestControlProp**](EnumglobalConfigurationAttributesModifiableWithIgnoreNoUserModificationRequestControlProp.md) |  | [optional] 
**MaximumServerOutLogFileSize** | Pointer to **string** | The maximum allowed size that the server.out log file will be allowed to have. If a write would cause the file to exceed this size, then the current file will be rotated out of place and a new empty file will be created and the message written to it. | [optional] 
**MaximumServerOutLogFileCount** | Pointer to **int64** | The maximum number of server.out log files (including the current active log file) that should be retained. When rotating the log file, if the total number of files exceeds this count, then the oldest file(s) will be removed so that the total number of log files is within this limit. | [optional] 
**StartupErrorLoggerOutputLocation** | Pointer to [**EnumglobalConfigurationStartupErrorLoggerOutputLocationProp**](EnumglobalConfigurationStartupErrorLoggerOutputLocationProp.md) |  | [optional] 
**ExitOnJVMError** | Pointer to **bool** | Indicates whether the Directory Server should be shut down if a severe error is raised (e.g., an out of memory error) which may prevent the JVM from continuing to run properly. | [optional] 
**ServerErrorResultCode** | Pointer to **int64** | Specifies the numeric value of the result code when request processing fails due to an internal server error. | [optional] 
**ResultCodeMap** | Pointer to **string** | Specifies a result code map that should be used for clients that do not have a map associated with their client connection policy. If the associated client connection policy has a result code map, then that map will be used instead. If no map is associated either with the client connection policy or the global configuration, then an internal default will be used. | [optional] 
**ReturnBindErrorMessages** | Pointer to **bool** | Indicates whether responses for failed bind operations should include a message string providing the reason for the authentication failure. | [optional] 
**NotifyAbandonedOperations** | Pointer to **bool** | Indicates whether the Directory Server should send a response to any operation that is interrupted via an abandon request. | [optional] 
**DuplicateErrorLogLimit** | **int64** | Specifies the maximum number of duplicate error log messages that should be logged in the time window specified by the duplicate-error-log-time-limit property. | 
**DuplicateErrorLogTimeLimit** | **string** | Specifies the length of time that must expire before duplicate log messages above the duplicate-error-log-limit threshold are logged again to the error log. | 
**DuplicateAlertLimit** | **int64** | Specifies the maximum number of duplicate alert messages that should be sent via the administrative alert framework in the time window specified by the duplicate-alert-time-limit property. | 
**DuplicateAlertTimeLimit** | **string** | Specifies the length of time that must expire before duplicate messages are sent via the administrative alert framework. | 
**WritabilityMode** | Pointer to [**EnumglobalConfigurationWritabilityModeProp**](EnumglobalConfigurationWritabilityModeProp.md) |  | [optional] 
**UseSharedDatabaseCacheAcrossAllLocalDBBackends** | Pointer to **bool** | Indicates whether the server should use a common database cache that is shared across all local DB backends instead of maintaining a separate cache for each backend. | [optional] 
**SharedLocalDBBackendDatabaseCachePercent** | Pointer to **int64** | Specifies the percentage of the JVM memory to allocate to the database cache that is shared across all local DB backends. | [optional] 
**UnrecoverableDatabaseErrorMode** | Pointer to [**EnumglobalConfigurationUnrecoverableDatabaseErrorModeProp**](EnumglobalConfigurationUnrecoverableDatabaseErrorModeProp.md) |  | [optional] 
**DatabaseOnVirtualizedOrNetworkStorage** | Pointer to **bool** | This setting provides data integrity options when the Directory Server is installed with a database on a network storage device. A storage device may be accessed directly by a physical server, or indirectly through a virtual machine running on a hypervisor. Enabling this setting will apply changes to all Local DB Backends, the LDAP Changelog Backend, and the replication changelog database. | [optional] 
**AutoNameWithEntryUUIDConnectionCriteria** | Pointer to **string** | Connection criteria that may be used to identify clients whose add requests should use entryUUID as the naming attribute. | [optional] 
**AutoNameWithEntryUUIDRequestCriteria** | Pointer to **string** | Request criteria that may be used to identify add requests that should use entryUUID as the naming attribute. | [optional] 
**SoftDeletePolicy** | Pointer to **string** | Specifies the soft delete policy that will be used by default for delete operations. Soft delete operations introduce the ability to control the server behavior of the delete operation. Instead of performing a permanent delete of an entry, deleted entries can be retained as soft deleted entries by their entryUUID values and are available for undelete at a later time. In addition to a soft delete policy enabling soft deletes, delete operations sent to the server must have the soft delete request control present with sufficient access privileges to access the soft delete request control. | [optional] 
**SubtreeAccessibilityAlertTimeLimit** | Pointer to **string** | Specifies the length of time that a subtree may remain hidden or read-only before an administrative alert is sent. | [optional] 
**WarnForBackendsWithMultipleBaseDns** | Pointer to **bool** | Indicates whether the server should issue a warning when enabling a backend that contains multiple base DNs. | [optional] 
**ForcedGCPrimeDuration** | Pointer to **string** | Specifies the minimum length of time required for backend or request processor initialization that will trigger the server to force an explicit garbage collection. A value of \&quot;0 seconds\&quot; indicates that the server should never invoke an explicit garbage collection regardless of the length of time required to initialize the server backends. | [optional] 
**ReplicationSetName** | Pointer to **string** | The name of the replication set assigned to this Directory Server. Restricted domains are only replicated within instances using the same replication set name. | [optional] 
**StartupMinReplicationBacklogCount** | **int64** | The number of outstanding changes any replica can have before the Directory Server will start accepting connections. The Directory Server may never accept connections if this setting is too low. If you are unsure which value to use, you can use the number of expected updates within a five second interval. | 
**ReplicationBacklogCountAlertThreshold** | **int64** | An alert is sent when the number of outstanding replication changes for the Directory Server has exceeded this threshold for longer than the replication backlog duration alert threshold. | 
**ReplicationBacklogDurationAlertThreshold** | **string** | An alert is sent when the number of outstanding replication changes for the Directory Server has exceeded the replication backlog count alert threshold for longer than this duration. | 
**ReplicationAssuranceSourceTimeoutSuspendDuration** | **string** | The amount of time a replication assurance source (i.e. a peer Directory Server) will be suspended from assurance requirements on this Directory Server if it experiences an assurance timeout. | 
**ReplicationAssuranceSourceBacklogFastStartThreshold** | **int64** | The maximum number of replication backlog updates a replication assurance source (i.e. a peer Directory Server) can have and be immediately recognized as an available assurance source by this Directory Server. | 
**ReplicationHistoryLimit** | Pointer to **int64** | Specifies the size limit for historical information. | [optional] 
**AllowInheritedReplicationOfSubordinateBackends** | **bool** | Allow replication to be inherited by subordinate/child backends. | 
**ReplicationPurgeObsoleteReplicas** | Pointer to **bool** | Indicates whether state about obsolete replicas is automatically purged. | [optional] 
**SmtpServer** | Pointer to **[]string** | Specifies the set of servers that will be used to send email messages. The order in which the servers are listed indicates the order in which the Directory Server will attempt to use them in the course of sending a message. The first attempt will always go to the server at the top of the list, and servers further down the list will only be used if none of the servers listed above it were able to successfully send the message. | [optional] 
**MaxSMTPConnectionCount** | Pointer to **int64** | The maximum number of SMTP connections that will be maintained for delivering email messages. | [optional] 
**MaxSMTPConnectionAge** | Pointer to **string** | The maximum length of time that a connection to an SMTP server should be considered valid. | [optional] 
**SmtpConnectionHealthCheckInterval** | Pointer to **string** | The length of time between checks to ensure that available SMTP connections are still valid. | [optional] 
**AllowedTask** | Pointer to **[]string** | Specifies the fully-qualified name of a Java class that may be invoked in the server. | [optional] 
**EnableSubOperationTimer** | Pointer to **bool** | Indicates whether the Directory Server should attempt to record information about the length of time required to process various phases of an operation. Enabling this feature may impact performance, but could make it easier to identify potential bottlenecks in operation processing. | [optional] 
**MaximumShutdownTime** | Pointer to **string** | Specifies the maximum amount of time the shutdown of Directory Server may take. | [optional] 
**NetworkAddressCacheTTL** | Pointer to **string** | Specifies the length of time that the Directory Server should cache the IP addresses associated with the names of systems with which it interacts. | [optional] 
**NetworkAddressOutageCacheEnabled** | Pointer to **bool** | Specifies whether the Directory Server should cache the last valid IP addresses associated with the names of systems with which it interacts with when the domain name service returns an unknown host exception. Java may return an unknown host exception when there is unexpected interruption in domain name service so this setting protects the Directory Server from temporary DNS server outages if previous results have been cached. | [optional] 
**TrackedApplication** | Pointer to **[]string** | Specifies criteria for identifying specific applications that access the server to enable tracking throughput and latency of LDAP operations issued by an application. | [optional] 
**JmxValueBehavior** | Pointer to [**EnumglobalConfigurationJmxValueBehaviorProp**](EnumglobalConfigurationJmxValueBehaviorProp.md) |  | [optional] 
**JmxUseLegacyMbeanNames** | Pointer to **bool** | When set to true, the server will use its original, non-standard JMX MBean names for the monitoring MBeans. These include RDN keys of \&quot;Rdn1\&quot; and \&quot;Rdn2\&quot; instead of the recommended \&quot;type\&quot; and \&quot;name\&quot; keys. This should option should only be enabled for installations that have monitoring infrastructure that depends on the old keys. | [optional] 

## Methods

### NewGlobalConfigurationResponse

`func NewGlobalConfigurationResponse(instanceName string, defaultPasswordPolicy string, proxiedAuthorizationIdentityMapper string, duplicateErrorLogLimit int64, duplicateErrorLogTimeLimit string, duplicateAlertLimit int64, duplicateAlertTimeLimit string, startupMinReplicationBacklogCount int64, replicationBacklogCountAlertThreshold int64, replicationBacklogDurationAlertThreshold string, replicationAssuranceSourceTimeoutSuspendDuration string, replicationAssuranceSourceBacklogFastStartThreshold int64, allowInheritedReplicationOfSubordinateBackends bool, ) *GlobalConfigurationResponse`

NewGlobalConfigurationResponse instantiates a new GlobalConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalConfigurationResponseWithDefaults

`func NewGlobalConfigurationResponseWithDefaults() *GlobalConfigurationResponse`

NewGlobalConfigurationResponseWithDefaults instantiates a new GlobalConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GlobalConfigurationResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GlobalConfigurationResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GlobalConfigurationResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GlobalConfigurationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GlobalConfigurationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GlobalConfigurationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GlobalConfigurationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GlobalConfigurationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GlobalConfigurationResponse) GetSchemas() []EnumglobalConfigurationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GlobalConfigurationResponse) GetSchemasOk() (*[]EnumglobalConfigurationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GlobalConfigurationResponse) SetSchemas(v []EnumglobalConfigurationSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *GlobalConfigurationResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetInstanceName

`func (o *GlobalConfigurationResponse) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *GlobalConfigurationResponse) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *GlobalConfigurationResponse) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetLocation

`func (o *GlobalConfigurationResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GlobalConfigurationResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GlobalConfigurationResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GlobalConfigurationResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConfigurationServerGroup

`func (o *GlobalConfigurationResponse) GetConfigurationServerGroup() string`

GetConfigurationServerGroup returns the ConfigurationServerGroup field if non-nil, zero value otherwise.

### GetConfigurationServerGroupOk

`func (o *GlobalConfigurationResponse) GetConfigurationServerGroupOk() (*string, bool)`

GetConfigurationServerGroupOk returns a tuple with the ConfigurationServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationServerGroup

`func (o *GlobalConfigurationResponse) SetConfigurationServerGroup(v string)`

SetConfigurationServerGroup sets ConfigurationServerGroup field to given value.

### HasConfigurationServerGroup

`func (o *GlobalConfigurationResponse) HasConfigurationServerGroup() bool`

HasConfigurationServerGroup returns a boolean if a field has been set.

### GetForceAsMasterForMirroredData

`func (o *GlobalConfigurationResponse) GetForceAsMasterForMirroredData() bool`

GetForceAsMasterForMirroredData returns the ForceAsMasterForMirroredData field if non-nil, zero value otherwise.

### GetForceAsMasterForMirroredDataOk

`func (o *GlobalConfigurationResponse) GetForceAsMasterForMirroredDataOk() (*bool, bool)`

GetForceAsMasterForMirroredDataOk returns a tuple with the ForceAsMasterForMirroredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAsMasterForMirroredData

`func (o *GlobalConfigurationResponse) SetForceAsMasterForMirroredData(v bool)`

SetForceAsMasterForMirroredData sets ForceAsMasterForMirroredData field to given value.

### HasForceAsMasterForMirroredData

`func (o *GlobalConfigurationResponse) HasForceAsMasterForMirroredData() bool`

HasForceAsMasterForMirroredData returns a boolean if a field has been set.

### GetEncryptData

`func (o *GlobalConfigurationResponse) GetEncryptData() bool`

GetEncryptData returns the EncryptData field if non-nil, zero value otherwise.

### GetEncryptDataOk

`func (o *GlobalConfigurationResponse) GetEncryptDataOk() (*bool, bool)`

GetEncryptDataOk returns a tuple with the EncryptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptData

`func (o *GlobalConfigurationResponse) SetEncryptData(v bool)`

SetEncryptData sets EncryptData field to given value.

### HasEncryptData

`func (o *GlobalConfigurationResponse) HasEncryptData() bool`

HasEncryptData returns a boolean if a field has been set.

### GetEncryptionSettingsCipherStreamProvider

`func (o *GlobalConfigurationResponse) GetEncryptionSettingsCipherStreamProvider() string`

GetEncryptionSettingsCipherStreamProvider returns the EncryptionSettingsCipherStreamProvider field if non-nil, zero value otherwise.

### GetEncryptionSettingsCipherStreamProviderOk

`func (o *GlobalConfigurationResponse) GetEncryptionSettingsCipherStreamProviderOk() (*string, bool)`

GetEncryptionSettingsCipherStreamProviderOk returns a tuple with the EncryptionSettingsCipherStreamProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsCipherStreamProvider

`func (o *GlobalConfigurationResponse) SetEncryptionSettingsCipherStreamProvider(v string)`

SetEncryptionSettingsCipherStreamProvider sets EncryptionSettingsCipherStreamProvider field to given value.

### HasEncryptionSettingsCipherStreamProvider

`func (o *GlobalConfigurationResponse) HasEncryptionSettingsCipherStreamProvider() bool`

HasEncryptionSettingsCipherStreamProvider returns a boolean if a field has been set.

### GetEncryptBackupsByDefault

`func (o *GlobalConfigurationResponse) GetEncryptBackupsByDefault() bool`

GetEncryptBackupsByDefault returns the EncryptBackupsByDefault field if non-nil, zero value otherwise.

### GetEncryptBackupsByDefaultOk

`func (o *GlobalConfigurationResponse) GetEncryptBackupsByDefaultOk() (*bool, bool)`

GetEncryptBackupsByDefaultOk returns a tuple with the EncryptBackupsByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptBackupsByDefault

`func (o *GlobalConfigurationResponse) SetEncryptBackupsByDefault(v bool)`

SetEncryptBackupsByDefault sets EncryptBackupsByDefault field to given value.

### HasEncryptBackupsByDefault

`func (o *GlobalConfigurationResponse) HasEncryptBackupsByDefault() bool`

HasEncryptBackupsByDefault returns a boolean if a field has been set.

### GetBackupEncryptionSettingsDefinitionID

`func (o *GlobalConfigurationResponse) GetBackupEncryptionSettingsDefinitionID() string`

GetBackupEncryptionSettingsDefinitionID returns the BackupEncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetBackupEncryptionSettingsDefinitionIDOk

`func (o *GlobalConfigurationResponse) GetBackupEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetBackupEncryptionSettingsDefinitionIDOk returns a tuple with the BackupEncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEncryptionSettingsDefinitionID

`func (o *GlobalConfigurationResponse) SetBackupEncryptionSettingsDefinitionID(v string)`

SetBackupEncryptionSettingsDefinitionID sets BackupEncryptionSettingsDefinitionID field to given value.

### HasBackupEncryptionSettingsDefinitionID

`func (o *GlobalConfigurationResponse) HasBackupEncryptionSettingsDefinitionID() bool`

HasBackupEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetEncryptLDIFExportsByDefault

`func (o *GlobalConfigurationResponse) GetEncryptLDIFExportsByDefault() bool`

GetEncryptLDIFExportsByDefault returns the EncryptLDIFExportsByDefault field if non-nil, zero value otherwise.

### GetEncryptLDIFExportsByDefaultOk

`func (o *GlobalConfigurationResponse) GetEncryptLDIFExportsByDefaultOk() (*bool, bool)`

GetEncryptLDIFExportsByDefaultOk returns a tuple with the EncryptLDIFExportsByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLDIFExportsByDefault

`func (o *GlobalConfigurationResponse) SetEncryptLDIFExportsByDefault(v bool)`

SetEncryptLDIFExportsByDefault sets EncryptLDIFExportsByDefault field to given value.

### HasEncryptLDIFExportsByDefault

`func (o *GlobalConfigurationResponse) HasEncryptLDIFExportsByDefault() bool`

HasEncryptLDIFExportsByDefault returns a boolean if a field has been set.

### GetLdifExportEncryptionSettingsDefinitionID

`func (o *GlobalConfigurationResponse) GetLdifExportEncryptionSettingsDefinitionID() string`

GetLdifExportEncryptionSettingsDefinitionID returns the LdifExportEncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetLdifExportEncryptionSettingsDefinitionIDOk

`func (o *GlobalConfigurationResponse) GetLdifExportEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetLdifExportEncryptionSettingsDefinitionIDOk returns a tuple with the LdifExportEncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifExportEncryptionSettingsDefinitionID

`func (o *GlobalConfigurationResponse) SetLdifExportEncryptionSettingsDefinitionID(v string)`

SetLdifExportEncryptionSettingsDefinitionID sets LdifExportEncryptionSettingsDefinitionID field to given value.

### HasLdifExportEncryptionSettingsDefinitionID

`func (o *GlobalConfigurationResponse) HasLdifExportEncryptionSettingsDefinitionID() bool`

HasLdifExportEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetAutomaticallyCompressEncryptedLDIFExports

`func (o *GlobalConfigurationResponse) GetAutomaticallyCompressEncryptedLDIFExports() bool`

GetAutomaticallyCompressEncryptedLDIFExports returns the AutomaticallyCompressEncryptedLDIFExports field if non-nil, zero value otherwise.

### GetAutomaticallyCompressEncryptedLDIFExportsOk

`func (o *GlobalConfigurationResponse) GetAutomaticallyCompressEncryptedLDIFExportsOk() (*bool, bool)`

GetAutomaticallyCompressEncryptedLDIFExportsOk returns a tuple with the AutomaticallyCompressEncryptedLDIFExports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyCompressEncryptedLDIFExports

`func (o *GlobalConfigurationResponse) SetAutomaticallyCompressEncryptedLDIFExports(v bool)`

SetAutomaticallyCompressEncryptedLDIFExports sets AutomaticallyCompressEncryptedLDIFExports field to given value.

### HasAutomaticallyCompressEncryptedLDIFExports

`func (o *GlobalConfigurationResponse) HasAutomaticallyCompressEncryptedLDIFExports() bool`

HasAutomaticallyCompressEncryptedLDIFExports returns a boolean if a field has been set.

### GetRedactSensitiveValuesInConfigLogs

`func (o *GlobalConfigurationResponse) GetRedactSensitiveValuesInConfigLogs() bool`

GetRedactSensitiveValuesInConfigLogs returns the RedactSensitiveValuesInConfigLogs field if non-nil, zero value otherwise.

### GetRedactSensitiveValuesInConfigLogsOk

`func (o *GlobalConfigurationResponse) GetRedactSensitiveValuesInConfigLogsOk() (*bool, bool)`

GetRedactSensitiveValuesInConfigLogsOk returns a tuple with the RedactSensitiveValuesInConfigLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactSensitiveValuesInConfigLogs

`func (o *GlobalConfigurationResponse) SetRedactSensitiveValuesInConfigLogs(v bool)`

SetRedactSensitiveValuesInConfigLogs sets RedactSensitiveValuesInConfigLogs field to given value.

### HasRedactSensitiveValuesInConfigLogs

`func (o *GlobalConfigurationResponse) HasRedactSensitiveValuesInConfigLogs() bool`

HasRedactSensitiveValuesInConfigLogs returns a boolean if a field has been set.

### GetSensitiveAttribute

`func (o *GlobalConfigurationResponse) GetSensitiveAttribute() []string`

GetSensitiveAttribute returns the SensitiveAttribute field if non-nil, zero value otherwise.

### GetSensitiveAttributeOk

`func (o *GlobalConfigurationResponse) GetSensitiveAttributeOk() (*[]string, bool)`

GetSensitiveAttributeOk returns a tuple with the SensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveAttribute

`func (o *GlobalConfigurationResponse) SetSensitiveAttribute(v []string)`

SetSensitiveAttribute sets SensitiveAttribute field to given value.

### HasSensitiveAttribute

`func (o *GlobalConfigurationResponse) HasSensitiveAttribute() bool`

HasSensitiveAttribute returns a boolean if a field has been set.

### GetRejectInsecureRequests

`func (o *GlobalConfigurationResponse) GetRejectInsecureRequests() bool`

GetRejectInsecureRequests returns the RejectInsecureRequests field if non-nil, zero value otherwise.

### GetRejectInsecureRequestsOk

`func (o *GlobalConfigurationResponse) GetRejectInsecureRequestsOk() (*bool, bool)`

GetRejectInsecureRequestsOk returns a tuple with the RejectInsecureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectInsecureRequests

`func (o *GlobalConfigurationResponse) SetRejectInsecureRequests(v bool)`

SetRejectInsecureRequests sets RejectInsecureRequests field to given value.

### HasRejectInsecureRequests

`func (o *GlobalConfigurationResponse) HasRejectInsecureRequests() bool`

HasRejectInsecureRequests returns a boolean if a field has been set.

### GetAllowedInsecureRequestCriteria

`func (o *GlobalConfigurationResponse) GetAllowedInsecureRequestCriteria() string`

GetAllowedInsecureRequestCriteria returns the AllowedInsecureRequestCriteria field if non-nil, zero value otherwise.

### GetAllowedInsecureRequestCriteriaOk

`func (o *GlobalConfigurationResponse) GetAllowedInsecureRequestCriteriaOk() (*string, bool)`

GetAllowedInsecureRequestCriteriaOk returns a tuple with the AllowedInsecureRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInsecureRequestCriteria

`func (o *GlobalConfigurationResponse) SetAllowedInsecureRequestCriteria(v string)`

SetAllowedInsecureRequestCriteria sets AllowedInsecureRequestCriteria field to given value.

### HasAllowedInsecureRequestCriteria

`func (o *GlobalConfigurationResponse) HasAllowedInsecureRequestCriteria() bool`

HasAllowedInsecureRequestCriteria returns a boolean if a field has been set.

### GetRejectUnauthenticatedRequests

`func (o *GlobalConfigurationResponse) GetRejectUnauthenticatedRequests() bool`

GetRejectUnauthenticatedRequests returns the RejectUnauthenticatedRequests field if non-nil, zero value otherwise.

### GetRejectUnauthenticatedRequestsOk

`func (o *GlobalConfigurationResponse) GetRejectUnauthenticatedRequestsOk() (*bool, bool)`

GetRejectUnauthenticatedRequestsOk returns a tuple with the RejectUnauthenticatedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUnauthenticatedRequests

`func (o *GlobalConfigurationResponse) SetRejectUnauthenticatedRequests(v bool)`

SetRejectUnauthenticatedRequests sets RejectUnauthenticatedRequests field to given value.

### HasRejectUnauthenticatedRequests

`func (o *GlobalConfigurationResponse) HasRejectUnauthenticatedRequests() bool`

HasRejectUnauthenticatedRequests returns a boolean if a field has been set.

### GetAllowedUnauthenticatedRequestCriteria

`func (o *GlobalConfigurationResponse) GetAllowedUnauthenticatedRequestCriteria() string`

GetAllowedUnauthenticatedRequestCriteria returns the AllowedUnauthenticatedRequestCriteria field if non-nil, zero value otherwise.

### GetAllowedUnauthenticatedRequestCriteriaOk

`func (o *GlobalConfigurationResponse) GetAllowedUnauthenticatedRequestCriteriaOk() (*string, bool)`

GetAllowedUnauthenticatedRequestCriteriaOk returns a tuple with the AllowedUnauthenticatedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUnauthenticatedRequestCriteria

`func (o *GlobalConfigurationResponse) SetAllowedUnauthenticatedRequestCriteria(v string)`

SetAllowedUnauthenticatedRequestCriteria sets AllowedUnauthenticatedRequestCriteria field to given value.

### HasAllowedUnauthenticatedRequestCriteria

`func (o *GlobalConfigurationResponse) HasAllowedUnauthenticatedRequestCriteria() bool`

HasAllowedUnauthenticatedRequestCriteria returns a boolean if a field has been set.

### GetBindWithDNRequiresPassword

`func (o *GlobalConfigurationResponse) GetBindWithDNRequiresPassword() bool`

GetBindWithDNRequiresPassword returns the BindWithDNRequiresPassword field if non-nil, zero value otherwise.

### GetBindWithDNRequiresPasswordOk

`func (o *GlobalConfigurationResponse) GetBindWithDNRequiresPasswordOk() (*bool, bool)`

GetBindWithDNRequiresPasswordOk returns a tuple with the BindWithDNRequiresPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindWithDNRequiresPassword

`func (o *GlobalConfigurationResponse) SetBindWithDNRequiresPassword(v bool)`

SetBindWithDNRequiresPassword sets BindWithDNRequiresPassword field to given value.

### HasBindWithDNRequiresPassword

`func (o *GlobalConfigurationResponse) HasBindWithDNRequiresPassword() bool`

HasBindWithDNRequiresPassword returns a boolean if a field has been set.

### GetDisabledPrivilege

`func (o *GlobalConfigurationResponse) GetDisabledPrivilege() []EnumglobalConfigurationDisabledPrivilegeProp`

GetDisabledPrivilege returns the DisabledPrivilege field if non-nil, zero value otherwise.

### GetDisabledPrivilegeOk

`func (o *GlobalConfigurationResponse) GetDisabledPrivilegeOk() (*[]EnumglobalConfigurationDisabledPrivilegeProp, bool)`

GetDisabledPrivilegeOk returns a tuple with the DisabledPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPrivilege

`func (o *GlobalConfigurationResponse) SetDisabledPrivilege(v []EnumglobalConfigurationDisabledPrivilegeProp)`

SetDisabledPrivilege sets DisabledPrivilege field to given value.

### HasDisabledPrivilege

`func (o *GlobalConfigurationResponse) HasDisabledPrivilege() bool`

HasDisabledPrivilege returns a boolean if a field has been set.

### GetDefaultPasswordPolicy

`func (o *GlobalConfigurationResponse) GetDefaultPasswordPolicy() string`

GetDefaultPasswordPolicy returns the DefaultPasswordPolicy field if non-nil, zero value otherwise.

### GetDefaultPasswordPolicyOk

`func (o *GlobalConfigurationResponse) GetDefaultPasswordPolicyOk() (*string, bool)`

GetDefaultPasswordPolicyOk returns a tuple with the DefaultPasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordPolicy

`func (o *GlobalConfigurationResponse) SetDefaultPasswordPolicy(v string)`

SetDefaultPasswordPolicy sets DefaultPasswordPolicy field to given value.


### GetMaximumUserDataPasswordPoliciesToCache

`func (o *GlobalConfigurationResponse) GetMaximumUserDataPasswordPoliciesToCache() int64`

GetMaximumUserDataPasswordPoliciesToCache returns the MaximumUserDataPasswordPoliciesToCache field if non-nil, zero value otherwise.

### GetMaximumUserDataPasswordPoliciesToCacheOk

`func (o *GlobalConfigurationResponse) GetMaximumUserDataPasswordPoliciesToCacheOk() (*int64, bool)`

GetMaximumUserDataPasswordPoliciesToCacheOk returns a tuple with the MaximumUserDataPasswordPoliciesToCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumUserDataPasswordPoliciesToCache

`func (o *GlobalConfigurationResponse) SetMaximumUserDataPasswordPoliciesToCache(v int64)`

SetMaximumUserDataPasswordPoliciesToCache sets MaximumUserDataPasswordPoliciesToCache field to given value.

### HasMaximumUserDataPasswordPoliciesToCache

`func (o *GlobalConfigurationResponse) HasMaximumUserDataPasswordPoliciesToCache() bool`

HasMaximumUserDataPasswordPoliciesToCache returns a boolean if a field has been set.

### GetProxiedAuthorizationIdentityMapper

`func (o *GlobalConfigurationResponse) GetProxiedAuthorizationIdentityMapper() string`

GetProxiedAuthorizationIdentityMapper returns the ProxiedAuthorizationIdentityMapper field if non-nil, zero value otherwise.

### GetProxiedAuthorizationIdentityMapperOk

`func (o *GlobalConfigurationResponse) GetProxiedAuthorizationIdentityMapperOk() (*string, bool)`

GetProxiedAuthorizationIdentityMapperOk returns a tuple with the ProxiedAuthorizationIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxiedAuthorizationIdentityMapper

`func (o *GlobalConfigurationResponse) SetProxiedAuthorizationIdentityMapper(v string)`

SetProxiedAuthorizationIdentityMapper sets ProxiedAuthorizationIdentityMapper field to given value.


### GetVerifyEntryDigests

`func (o *GlobalConfigurationResponse) GetVerifyEntryDigests() bool`

GetVerifyEntryDigests returns the VerifyEntryDigests field if non-nil, zero value otherwise.

### GetVerifyEntryDigestsOk

`func (o *GlobalConfigurationResponse) GetVerifyEntryDigestsOk() (*bool, bool)`

GetVerifyEntryDigestsOk returns a tuple with the VerifyEntryDigests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyEntryDigests

`func (o *GlobalConfigurationResponse) SetVerifyEntryDigests(v bool)`

SetVerifyEntryDigests sets VerifyEntryDigests field to given value.

### HasVerifyEntryDigests

`func (o *GlobalConfigurationResponse) HasVerifyEntryDigests() bool`

HasVerifyEntryDigests returns a boolean if a field has been set.

### GetAllowedInsecureTLSProtocol

`func (o *GlobalConfigurationResponse) GetAllowedInsecureTLSProtocol() []EnumglobalConfigurationAllowedInsecureTLSProtocolProp`

GetAllowedInsecureTLSProtocol returns the AllowedInsecureTLSProtocol field if non-nil, zero value otherwise.

### GetAllowedInsecureTLSProtocolOk

`func (o *GlobalConfigurationResponse) GetAllowedInsecureTLSProtocolOk() (*[]EnumglobalConfigurationAllowedInsecureTLSProtocolProp, bool)`

GetAllowedInsecureTLSProtocolOk returns a tuple with the AllowedInsecureTLSProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInsecureTLSProtocol

`func (o *GlobalConfigurationResponse) SetAllowedInsecureTLSProtocol(v []EnumglobalConfigurationAllowedInsecureTLSProtocolProp)`

SetAllowedInsecureTLSProtocol sets AllowedInsecureTLSProtocol field to given value.

### HasAllowedInsecureTLSProtocol

`func (o *GlobalConfigurationResponse) HasAllowedInsecureTLSProtocol() bool`

HasAllowedInsecureTLSProtocol returns a boolean if a field has been set.

### GetAllowInsecureLocalJMXConnections

`func (o *GlobalConfigurationResponse) GetAllowInsecureLocalJMXConnections() bool`

GetAllowInsecureLocalJMXConnections returns the AllowInsecureLocalJMXConnections field if non-nil, zero value otherwise.

### GetAllowInsecureLocalJMXConnectionsOk

`func (o *GlobalConfigurationResponse) GetAllowInsecureLocalJMXConnectionsOk() (*bool, bool)`

GetAllowInsecureLocalJMXConnectionsOk returns a tuple with the AllowInsecureLocalJMXConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecureLocalJMXConnections

`func (o *GlobalConfigurationResponse) SetAllowInsecureLocalJMXConnections(v bool)`

SetAllowInsecureLocalJMXConnections sets AllowInsecureLocalJMXConnections field to given value.

### HasAllowInsecureLocalJMXConnections

`func (o *GlobalConfigurationResponse) HasAllowInsecureLocalJMXConnections() bool`

HasAllowInsecureLocalJMXConnections returns a boolean if a field has been set.

### GetDefaultInternalOperationClientConnectionPolicy

`func (o *GlobalConfigurationResponse) GetDefaultInternalOperationClientConnectionPolicy() string`

GetDefaultInternalOperationClientConnectionPolicy returns the DefaultInternalOperationClientConnectionPolicy field if non-nil, zero value otherwise.

### GetDefaultInternalOperationClientConnectionPolicyOk

`func (o *GlobalConfigurationResponse) GetDefaultInternalOperationClientConnectionPolicyOk() (*string, bool)`

GetDefaultInternalOperationClientConnectionPolicyOk returns a tuple with the DefaultInternalOperationClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInternalOperationClientConnectionPolicy

`func (o *GlobalConfigurationResponse) SetDefaultInternalOperationClientConnectionPolicy(v string)`

SetDefaultInternalOperationClientConnectionPolicy sets DefaultInternalOperationClientConnectionPolicy field to given value.

### HasDefaultInternalOperationClientConnectionPolicy

`func (o *GlobalConfigurationResponse) HasDefaultInternalOperationClientConnectionPolicy() bool`

HasDefaultInternalOperationClientConnectionPolicy returns a boolean if a field has been set.

### GetSizeLimit

`func (o *GlobalConfigurationResponse) GetSizeLimit() int64`

GetSizeLimit returns the SizeLimit field if non-nil, zero value otherwise.

### GetSizeLimitOk

`func (o *GlobalConfigurationResponse) GetSizeLimitOk() (*int64, bool)`

GetSizeLimitOk returns a tuple with the SizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimit

`func (o *GlobalConfigurationResponse) SetSizeLimit(v int64)`

SetSizeLimit sets SizeLimit field to given value.

### HasSizeLimit

`func (o *GlobalConfigurationResponse) HasSizeLimit() bool`

HasSizeLimit returns a boolean if a field has been set.

### GetUnauthenticatedSizeLimit

`func (o *GlobalConfigurationResponse) GetUnauthenticatedSizeLimit() int64`

GetUnauthenticatedSizeLimit returns the UnauthenticatedSizeLimit field if non-nil, zero value otherwise.

### GetUnauthenticatedSizeLimitOk

`func (o *GlobalConfigurationResponse) GetUnauthenticatedSizeLimitOk() (*int64, bool)`

GetUnauthenticatedSizeLimitOk returns a tuple with the UnauthenticatedSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedSizeLimit

`func (o *GlobalConfigurationResponse) SetUnauthenticatedSizeLimit(v int64)`

SetUnauthenticatedSizeLimit sets UnauthenticatedSizeLimit field to given value.

### HasUnauthenticatedSizeLimit

`func (o *GlobalConfigurationResponse) HasUnauthenticatedSizeLimit() bool`

HasUnauthenticatedSizeLimit returns a boolean if a field has been set.

### GetTimeLimit

`func (o *GlobalConfigurationResponse) GetTimeLimit() string`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *GlobalConfigurationResponse) GetTimeLimitOk() (*string, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *GlobalConfigurationResponse) SetTimeLimit(v string)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *GlobalConfigurationResponse) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetUnauthenticatedTimeLimit

`func (o *GlobalConfigurationResponse) GetUnauthenticatedTimeLimit() string`

GetUnauthenticatedTimeLimit returns the UnauthenticatedTimeLimit field if non-nil, zero value otherwise.

### GetUnauthenticatedTimeLimitOk

`func (o *GlobalConfigurationResponse) GetUnauthenticatedTimeLimitOk() (*string, bool)`

GetUnauthenticatedTimeLimitOk returns a tuple with the UnauthenticatedTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedTimeLimit

`func (o *GlobalConfigurationResponse) SetUnauthenticatedTimeLimit(v string)`

SetUnauthenticatedTimeLimit sets UnauthenticatedTimeLimit field to given value.

### HasUnauthenticatedTimeLimit

`func (o *GlobalConfigurationResponse) HasUnauthenticatedTimeLimit() bool`

HasUnauthenticatedTimeLimit returns a boolean if a field has been set.

### GetIdleTimeLimit

`func (o *GlobalConfigurationResponse) GetIdleTimeLimit() string`

GetIdleTimeLimit returns the IdleTimeLimit field if non-nil, zero value otherwise.

### GetIdleTimeLimitOk

`func (o *GlobalConfigurationResponse) GetIdleTimeLimitOk() (*string, bool)`

GetIdleTimeLimitOk returns a tuple with the IdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimit

`func (o *GlobalConfigurationResponse) SetIdleTimeLimit(v string)`

SetIdleTimeLimit sets IdleTimeLimit field to given value.

### HasIdleTimeLimit

`func (o *GlobalConfigurationResponse) HasIdleTimeLimit() bool`

HasIdleTimeLimit returns a boolean if a field has been set.

### GetUnauthenticatedIdleTimeLimit

`func (o *GlobalConfigurationResponse) GetUnauthenticatedIdleTimeLimit() string`

GetUnauthenticatedIdleTimeLimit returns the UnauthenticatedIdleTimeLimit field if non-nil, zero value otherwise.

### GetUnauthenticatedIdleTimeLimitOk

`func (o *GlobalConfigurationResponse) GetUnauthenticatedIdleTimeLimitOk() (*string, bool)`

GetUnauthenticatedIdleTimeLimitOk returns a tuple with the UnauthenticatedIdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedIdleTimeLimit

`func (o *GlobalConfigurationResponse) SetUnauthenticatedIdleTimeLimit(v string)`

SetUnauthenticatedIdleTimeLimit sets UnauthenticatedIdleTimeLimit field to given value.

### HasUnauthenticatedIdleTimeLimit

`func (o *GlobalConfigurationResponse) HasUnauthenticatedIdleTimeLimit() bool`

HasUnauthenticatedIdleTimeLimit returns a boolean if a field has been set.

### GetLookthroughLimit

`func (o *GlobalConfigurationResponse) GetLookthroughLimit() int64`

GetLookthroughLimit returns the LookthroughLimit field if non-nil, zero value otherwise.

### GetLookthroughLimitOk

`func (o *GlobalConfigurationResponse) GetLookthroughLimitOk() (*int64, bool)`

GetLookthroughLimitOk returns a tuple with the LookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookthroughLimit

`func (o *GlobalConfigurationResponse) SetLookthroughLimit(v int64)`

SetLookthroughLimit sets LookthroughLimit field to given value.

### HasLookthroughLimit

`func (o *GlobalConfigurationResponse) HasLookthroughLimit() bool`

HasLookthroughLimit returns a boolean if a field has been set.

### GetUnauthenticatedLookthroughLimit

`func (o *GlobalConfigurationResponse) GetUnauthenticatedLookthroughLimit() int64`

GetUnauthenticatedLookthroughLimit returns the UnauthenticatedLookthroughLimit field if non-nil, zero value otherwise.

### GetUnauthenticatedLookthroughLimitOk

`func (o *GlobalConfigurationResponse) GetUnauthenticatedLookthroughLimitOk() (*int64, bool)`

GetUnauthenticatedLookthroughLimitOk returns a tuple with the UnauthenticatedLookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedLookthroughLimit

`func (o *GlobalConfigurationResponse) SetUnauthenticatedLookthroughLimit(v int64)`

SetUnauthenticatedLookthroughLimit sets UnauthenticatedLookthroughLimit field to given value.

### HasUnauthenticatedLookthroughLimit

`func (o *GlobalConfigurationResponse) HasUnauthenticatedLookthroughLimit() bool`

HasUnauthenticatedLookthroughLimit returns a boolean if a field has been set.

### GetLdapJoinSizeLimit

`func (o *GlobalConfigurationResponse) GetLdapJoinSizeLimit() int64`

GetLdapJoinSizeLimit returns the LdapJoinSizeLimit field if non-nil, zero value otherwise.

### GetLdapJoinSizeLimitOk

`func (o *GlobalConfigurationResponse) GetLdapJoinSizeLimitOk() (*int64, bool)`

GetLdapJoinSizeLimitOk returns a tuple with the LdapJoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapJoinSizeLimit

`func (o *GlobalConfigurationResponse) SetLdapJoinSizeLimit(v int64)`

SetLdapJoinSizeLimit sets LdapJoinSizeLimit field to given value.

### HasLdapJoinSizeLimit

`func (o *GlobalConfigurationResponse) HasLdapJoinSizeLimit() bool`

HasLdapJoinSizeLimit returns a boolean if a field has been set.

### GetMaximumConcurrentConnections

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentConnections() int64`

GetMaximumConcurrentConnections returns the MaximumConcurrentConnections field if non-nil, zero value otherwise.

### GetMaximumConcurrentConnectionsOk

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentConnectionsOk() (*int64, bool)`

GetMaximumConcurrentConnectionsOk returns a tuple with the MaximumConcurrentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentConnections

`func (o *GlobalConfigurationResponse) SetMaximumConcurrentConnections(v int64)`

SetMaximumConcurrentConnections sets MaximumConcurrentConnections field to given value.

### HasMaximumConcurrentConnections

`func (o *GlobalConfigurationResponse) HasMaximumConcurrentConnections() bool`

HasMaximumConcurrentConnections returns a boolean if a field has been set.

### GetMaximumConcurrentConnectionsPerIPAddress

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentConnectionsPerIPAddress() int64`

GetMaximumConcurrentConnectionsPerIPAddress returns the MaximumConcurrentConnectionsPerIPAddress field if non-nil, zero value otherwise.

### GetMaximumConcurrentConnectionsPerIPAddressOk

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentConnectionsPerIPAddressOk() (*int64, bool)`

GetMaximumConcurrentConnectionsPerIPAddressOk returns a tuple with the MaximumConcurrentConnectionsPerIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentConnectionsPerIPAddress

`func (o *GlobalConfigurationResponse) SetMaximumConcurrentConnectionsPerIPAddress(v int64)`

SetMaximumConcurrentConnectionsPerIPAddress sets MaximumConcurrentConnectionsPerIPAddress field to given value.

### HasMaximumConcurrentConnectionsPerIPAddress

`func (o *GlobalConfigurationResponse) HasMaximumConcurrentConnectionsPerIPAddress() bool`

HasMaximumConcurrentConnectionsPerIPAddress returns a boolean if a field has been set.

### GetMaximumConcurrentConnectionsPerBindDN

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentConnectionsPerBindDN() int64`

GetMaximumConcurrentConnectionsPerBindDN returns the MaximumConcurrentConnectionsPerBindDN field if non-nil, zero value otherwise.

### GetMaximumConcurrentConnectionsPerBindDNOk

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentConnectionsPerBindDNOk() (*int64, bool)`

GetMaximumConcurrentConnectionsPerBindDNOk returns a tuple with the MaximumConcurrentConnectionsPerBindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentConnectionsPerBindDN

`func (o *GlobalConfigurationResponse) SetMaximumConcurrentConnectionsPerBindDN(v int64)`

SetMaximumConcurrentConnectionsPerBindDN sets MaximumConcurrentConnectionsPerBindDN field to given value.

### HasMaximumConcurrentConnectionsPerBindDN

`func (o *GlobalConfigurationResponse) HasMaximumConcurrentConnectionsPerBindDN() bool`

HasMaximumConcurrentConnectionsPerBindDN returns a boolean if a field has been set.

### GetMaximumConcurrentUnindexedSearches

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentUnindexedSearches() int64`

GetMaximumConcurrentUnindexedSearches returns the MaximumConcurrentUnindexedSearches field if non-nil, zero value otherwise.

### GetMaximumConcurrentUnindexedSearchesOk

`func (o *GlobalConfigurationResponse) GetMaximumConcurrentUnindexedSearchesOk() (*int64, bool)`

GetMaximumConcurrentUnindexedSearchesOk returns a tuple with the MaximumConcurrentUnindexedSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentUnindexedSearches

`func (o *GlobalConfigurationResponse) SetMaximumConcurrentUnindexedSearches(v int64)`

SetMaximumConcurrentUnindexedSearches sets MaximumConcurrentUnindexedSearches field to given value.

### HasMaximumConcurrentUnindexedSearches

`func (o *GlobalConfigurationResponse) HasMaximumConcurrentUnindexedSearches() bool`

HasMaximumConcurrentUnindexedSearches returns a boolean if a field has been set.

### GetMaximumAttributesPerAddRequest

`func (o *GlobalConfigurationResponse) GetMaximumAttributesPerAddRequest() int64`

GetMaximumAttributesPerAddRequest returns the MaximumAttributesPerAddRequest field if non-nil, zero value otherwise.

### GetMaximumAttributesPerAddRequestOk

`func (o *GlobalConfigurationResponse) GetMaximumAttributesPerAddRequestOk() (*int64, bool)`

GetMaximumAttributesPerAddRequestOk returns a tuple with the MaximumAttributesPerAddRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAttributesPerAddRequest

`func (o *GlobalConfigurationResponse) SetMaximumAttributesPerAddRequest(v int64)`

SetMaximumAttributesPerAddRequest sets MaximumAttributesPerAddRequest field to given value.

### HasMaximumAttributesPerAddRequest

`func (o *GlobalConfigurationResponse) HasMaximumAttributesPerAddRequest() bool`

HasMaximumAttributesPerAddRequest returns a boolean if a field has been set.

### GetMaximumModificationsPerModifyRequest

`func (o *GlobalConfigurationResponse) GetMaximumModificationsPerModifyRequest() int64`

GetMaximumModificationsPerModifyRequest returns the MaximumModificationsPerModifyRequest field if non-nil, zero value otherwise.

### GetMaximumModificationsPerModifyRequestOk

`func (o *GlobalConfigurationResponse) GetMaximumModificationsPerModifyRequestOk() (*int64, bool)`

GetMaximumModificationsPerModifyRequestOk returns a tuple with the MaximumModificationsPerModifyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumModificationsPerModifyRequest

`func (o *GlobalConfigurationResponse) SetMaximumModificationsPerModifyRequest(v int64)`

SetMaximumModificationsPerModifyRequest sets MaximumModificationsPerModifyRequest field to given value.

### HasMaximumModificationsPerModifyRequest

`func (o *GlobalConfigurationResponse) HasMaximumModificationsPerModifyRequest() bool`

HasMaximumModificationsPerModifyRequest returns a boolean if a field has been set.

### GetBackgroundThreadForEachPersistentSearch

`func (o *GlobalConfigurationResponse) GetBackgroundThreadForEachPersistentSearch() bool`

GetBackgroundThreadForEachPersistentSearch returns the BackgroundThreadForEachPersistentSearch field if non-nil, zero value otherwise.

### GetBackgroundThreadForEachPersistentSearchOk

`func (o *GlobalConfigurationResponse) GetBackgroundThreadForEachPersistentSearchOk() (*bool, bool)`

GetBackgroundThreadForEachPersistentSearchOk returns a tuple with the BackgroundThreadForEachPersistentSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundThreadForEachPersistentSearch

`func (o *GlobalConfigurationResponse) SetBackgroundThreadForEachPersistentSearch(v bool)`

SetBackgroundThreadForEachPersistentSearch sets BackgroundThreadForEachPersistentSearch field to given value.

### HasBackgroundThreadForEachPersistentSearch

`func (o *GlobalConfigurationResponse) HasBackgroundThreadForEachPersistentSearch() bool`

HasBackgroundThreadForEachPersistentSearch returns a boolean if a field has been set.

### GetAllowAttributeNameExceptions

`func (o *GlobalConfigurationResponse) GetAllowAttributeNameExceptions() bool`

GetAllowAttributeNameExceptions returns the AllowAttributeNameExceptions field if non-nil, zero value otherwise.

### GetAllowAttributeNameExceptionsOk

`func (o *GlobalConfigurationResponse) GetAllowAttributeNameExceptionsOk() (*bool, bool)`

GetAllowAttributeNameExceptionsOk returns a tuple with the AllowAttributeNameExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAttributeNameExceptions

`func (o *GlobalConfigurationResponse) SetAllowAttributeNameExceptions(v bool)`

SetAllowAttributeNameExceptions sets AllowAttributeNameExceptions field to given value.

### HasAllowAttributeNameExceptions

`func (o *GlobalConfigurationResponse) HasAllowAttributeNameExceptions() bool`

HasAllowAttributeNameExceptions returns a boolean if a field has been set.

### GetInvalidAttributeSyntaxBehavior

`func (o *GlobalConfigurationResponse) GetInvalidAttributeSyntaxBehavior() EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp`

GetInvalidAttributeSyntaxBehavior returns the InvalidAttributeSyntaxBehavior field if non-nil, zero value otherwise.

### GetInvalidAttributeSyntaxBehaviorOk

`func (o *GlobalConfigurationResponse) GetInvalidAttributeSyntaxBehaviorOk() (*EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp, bool)`

GetInvalidAttributeSyntaxBehaviorOk returns a tuple with the InvalidAttributeSyntaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidAttributeSyntaxBehavior

`func (o *GlobalConfigurationResponse) SetInvalidAttributeSyntaxBehavior(v EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp)`

SetInvalidAttributeSyntaxBehavior sets InvalidAttributeSyntaxBehavior field to given value.

### HasInvalidAttributeSyntaxBehavior

`func (o *GlobalConfigurationResponse) HasInvalidAttributeSyntaxBehavior() bool`

HasInvalidAttributeSyntaxBehavior returns a boolean if a field has been set.

### GetPermitSyntaxViolationsForAttribute

`func (o *GlobalConfigurationResponse) GetPermitSyntaxViolationsForAttribute() []string`

GetPermitSyntaxViolationsForAttribute returns the PermitSyntaxViolationsForAttribute field if non-nil, zero value otherwise.

### GetPermitSyntaxViolationsForAttributeOk

`func (o *GlobalConfigurationResponse) GetPermitSyntaxViolationsForAttributeOk() (*[]string, bool)`

GetPermitSyntaxViolationsForAttributeOk returns a tuple with the PermitSyntaxViolationsForAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitSyntaxViolationsForAttribute

`func (o *GlobalConfigurationResponse) SetPermitSyntaxViolationsForAttribute(v []string)`

SetPermitSyntaxViolationsForAttribute sets PermitSyntaxViolationsForAttribute field to given value.

### HasPermitSyntaxViolationsForAttribute

`func (o *GlobalConfigurationResponse) HasPermitSyntaxViolationsForAttribute() bool`

HasPermitSyntaxViolationsForAttribute returns a boolean if a field has been set.

### GetSingleStructuralObjectclassBehavior

`func (o *GlobalConfigurationResponse) GetSingleStructuralObjectclassBehavior() EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp`

GetSingleStructuralObjectclassBehavior returns the SingleStructuralObjectclassBehavior field if non-nil, zero value otherwise.

### GetSingleStructuralObjectclassBehaviorOk

`func (o *GlobalConfigurationResponse) GetSingleStructuralObjectclassBehaviorOk() (*EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp, bool)`

GetSingleStructuralObjectclassBehaviorOk returns a tuple with the SingleStructuralObjectclassBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleStructuralObjectclassBehavior

`func (o *GlobalConfigurationResponse) SetSingleStructuralObjectclassBehavior(v EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp)`

SetSingleStructuralObjectclassBehavior sets SingleStructuralObjectclassBehavior field to given value.

### HasSingleStructuralObjectclassBehavior

`func (o *GlobalConfigurationResponse) HasSingleStructuralObjectclassBehavior() bool`

HasSingleStructuralObjectclassBehavior returns a boolean if a field has been set.

### GetAttributesModifiableWithIgnoreNoUserModificationRequestControl

`func (o *GlobalConfigurationResponse) GetAttributesModifiableWithIgnoreNoUserModificationRequestControl() []EnumglobalConfigurationAttributesModifiableWithIgnoreNoUserModificationRequestControlProp`

GetAttributesModifiableWithIgnoreNoUserModificationRequestControl returns the AttributesModifiableWithIgnoreNoUserModificationRequestControl field if non-nil, zero value otherwise.

### GetAttributesModifiableWithIgnoreNoUserModificationRequestControlOk

`func (o *GlobalConfigurationResponse) GetAttributesModifiableWithIgnoreNoUserModificationRequestControlOk() (*[]EnumglobalConfigurationAttributesModifiableWithIgnoreNoUserModificationRequestControlProp, bool)`

GetAttributesModifiableWithIgnoreNoUserModificationRequestControlOk returns a tuple with the AttributesModifiableWithIgnoreNoUserModificationRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesModifiableWithIgnoreNoUserModificationRequestControl

`func (o *GlobalConfigurationResponse) SetAttributesModifiableWithIgnoreNoUserModificationRequestControl(v []EnumglobalConfigurationAttributesModifiableWithIgnoreNoUserModificationRequestControlProp)`

SetAttributesModifiableWithIgnoreNoUserModificationRequestControl sets AttributesModifiableWithIgnoreNoUserModificationRequestControl field to given value.

### HasAttributesModifiableWithIgnoreNoUserModificationRequestControl

`func (o *GlobalConfigurationResponse) HasAttributesModifiableWithIgnoreNoUserModificationRequestControl() bool`

HasAttributesModifiableWithIgnoreNoUserModificationRequestControl returns a boolean if a field has been set.

### GetMaximumServerOutLogFileSize

`func (o *GlobalConfigurationResponse) GetMaximumServerOutLogFileSize() string`

GetMaximumServerOutLogFileSize returns the MaximumServerOutLogFileSize field if non-nil, zero value otherwise.

### GetMaximumServerOutLogFileSizeOk

`func (o *GlobalConfigurationResponse) GetMaximumServerOutLogFileSizeOk() (*string, bool)`

GetMaximumServerOutLogFileSizeOk returns a tuple with the MaximumServerOutLogFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumServerOutLogFileSize

`func (o *GlobalConfigurationResponse) SetMaximumServerOutLogFileSize(v string)`

SetMaximumServerOutLogFileSize sets MaximumServerOutLogFileSize field to given value.

### HasMaximumServerOutLogFileSize

`func (o *GlobalConfigurationResponse) HasMaximumServerOutLogFileSize() bool`

HasMaximumServerOutLogFileSize returns a boolean if a field has been set.

### GetMaximumServerOutLogFileCount

`func (o *GlobalConfigurationResponse) GetMaximumServerOutLogFileCount() int64`

GetMaximumServerOutLogFileCount returns the MaximumServerOutLogFileCount field if non-nil, zero value otherwise.

### GetMaximumServerOutLogFileCountOk

`func (o *GlobalConfigurationResponse) GetMaximumServerOutLogFileCountOk() (*int64, bool)`

GetMaximumServerOutLogFileCountOk returns a tuple with the MaximumServerOutLogFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumServerOutLogFileCount

`func (o *GlobalConfigurationResponse) SetMaximumServerOutLogFileCount(v int64)`

SetMaximumServerOutLogFileCount sets MaximumServerOutLogFileCount field to given value.

### HasMaximumServerOutLogFileCount

`func (o *GlobalConfigurationResponse) HasMaximumServerOutLogFileCount() bool`

HasMaximumServerOutLogFileCount returns a boolean if a field has been set.

### GetStartupErrorLoggerOutputLocation

`func (o *GlobalConfigurationResponse) GetStartupErrorLoggerOutputLocation() EnumglobalConfigurationStartupErrorLoggerOutputLocationProp`

GetStartupErrorLoggerOutputLocation returns the StartupErrorLoggerOutputLocation field if non-nil, zero value otherwise.

### GetStartupErrorLoggerOutputLocationOk

`func (o *GlobalConfigurationResponse) GetStartupErrorLoggerOutputLocationOk() (*EnumglobalConfigurationStartupErrorLoggerOutputLocationProp, bool)`

GetStartupErrorLoggerOutputLocationOk returns a tuple with the StartupErrorLoggerOutputLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupErrorLoggerOutputLocation

`func (o *GlobalConfigurationResponse) SetStartupErrorLoggerOutputLocation(v EnumglobalConfigurationStartupErrorLoggerOutputLocationProp)`

SetStartupErrorLoggerOutputLocation sets StartupErrorLoggerOutputLocation field to given value.

### HasStartupErrorLoggerOutputLocation

`func (o *GlobalConfigurationResponse) HasStartupErrorLoggerOutputLocation() bool`

HasStartupErrorLoggerOutputLocation returns a boolean if a field has been set.

### GetExitOnJVMError

`func (o *GlobalConfigurationResponse) GetExitOnJVMError() bool`

GetExitOnJVMError returns the ExitOnJVMError field if non-nil, zero value otherwise.

### GetExitOnJVMErrorOk

`func (o *GlobalConfigurationResponse) GetExitOnJVMErrorOk() (*bool, bool)`

GetExitOnJVMErrorOk returns a tuple with the ExitOnJVMError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnJVMError

`func (o *GlobalConfigurationResponse) SetExitOnJVMError(v bool)`

SetExitOnJVMError sets ExitOnJVMError field to given value.

### HasExitOnJVMError

`func (o *GlobalConfigurationResponse) HasExitOnJVMError() bool`

HasExitOnJVMError returns a boolean if a field has been set.

### GetServerErrorResultCode

`func (o *GlobalConfigurationResponse) GetServerErrorResultCode() int64`

GetServerErrorResultCode returns the ServerErrorResultCode field if non-nil, zero value otherwise.

### GetServerErrorResultCodeOk

`func (o *GlobalConfigurationResponse) GetServerErrorResultCodeOk() (*int64, bool)`

GetServerErrorResultCodeOk returns a tuple with the ServerErrorResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorResultCode

`func (o *GlobalConfigurationResponse) SetServerErrorResultCode(v int64)`

SetServerErrorResultCode sets ServerErrorResultCode field to given value.

### HasServerErrorResultCode

`func (o *GlobalConfigurationResponse) HasServerErrorResultCode() bool`

HasServerErrorResultCode returns a boolean if a field has been set.

### GetResultCodeMap

`func (o *GlobalConfigurationResponse) GetResultCodeMap() string`

GetResultCodeMap returns the ResultCodeMap field if non-nil, zero value otherwise.

### GetResultCodeMapOk

`func (o *GlobalConfigurationResponse) GetResultCodeMapOk() (*string, bool)`

GetResultCodeMapOk returns a tuple with the ResultCodeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeMap

`func (o *GlobalConfigurationResponse) SetResultCodeMap(v string)`

SetResultCodeMap sets ResultCodeMap field to given value.

### HasResultCodeMap

`func (o *GlobalConfigurationResponse) HasResultCodeMap() bool`

HasResultCodeMap returns a boolean if a field has been set.

### GetReturnBindErrorMessages

`func (o *GlobalConfigurationResponse) GetReturnBindErrorMessages() bool`

GetReturnBindErrorMessages returns the ReturnBindErrorMessages field if non-nil, zero value otherwise.

### GetReturnBindErrorMessagesOk

`func (o *GlobalConfigurationResponse) GetReturnBindErrorMessagesOk() (*bool, bool)`

GetReturnBindErrorMessagesOk returns a tuple with the ReturnBindErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnBindErrorMessages

`func (o *GlobalConfigurationResponse) SetReturnBindErrorMessages(v bool)`

SetReturnBindErrorMessages sets ReturnBindErrorMessages field to given value.

### HasReturnBindErrorMessages

`func (o *GlobalConfigurationResponse) HasReturnBindErrorMessages() bool`

HasReturnBindErrorMessages returns a boolean if a field has been set.

### GetNotifyAbandonedOperations

`func (o *GlobalConfigurationResponse) GetNotifyAbandonedOperations() bool`

GetNotifyAbandonedOperations returns the NotifyAbandonedOperations field if non-nil, zero value otherwise.

### GetNotifyAbandonedOperationsOk

`func (o *GlobalConfigurationResponse) GetNotifyAbandonedOperationsOk() (*bool, bool)`

GetNotifyAbandonedOperationsOk returns a tuple with the NotifyAbandonedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAbandonedOperations

`func (o *GlobalConfigurationResponse) SetNotifyAbandonedOperations(v bool)`

SetNotifyAbandonedOperations sets NotifyAbandonedOperations field to given value.

### HasNotifyAbandonedOperations

`func (o *GlobalConfigurationResponse) HasNotifyAbandonedOperations() bool`

HasNotifyAbandonedOperations returns a boolean if a field has been set.

### GetDuplicateErrorLogLimit

`func (o *GlobalConfigurationResponse) GetDuplicateErrorLogLimit() int64`

GetDuplicateErrorLogLimit returns the DuplicateErrorLogLimit field if non-nil, zero value otherwise.

### GetDuplicateErrorLogLimitOk

`func (o *GlobalConfigurationResponse) GetDuplicateErrorLogLimitOk() (*int64, bool)`

GetDuplicateErrorLogLimitOk returns a tuple with the DuplicateErrorLogLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateErrorLogLimit

`func (o *GlobalConfigurationResponse) SetDuplicateErrorLogLimit(v int64)`

SetDuplicateErrorLogLimit sets DuplicateErrorLogLimit field to given value.


### GetDuplicateErrorLogTimeLimit

`func (o *GlobalConfigurationResponse) GetDuplicateErrorLogTimeLimit() string`

GetDuplicateErrorLogTimeLimit returns the DuplicateErrorLogTimeLimit field if non-nil, zero value otherwise.

### GetDuplicateErrorLogTimeLimitOk

`func (o *GlobalConfigurationResponse) GetDuplicateErrorLogTimeLimitOk() (*string, bool)`

GetDuplicateErrorLogTimeLimitOk returns a tuple with the DuplicateErrorLogTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateErrorLogTimeLimit

`func (o *GlobalConfigurationResponse) SetDuplicateErrorLogTimeLimit(v string)`

SetDuplicateErrorLogTimeLimit sets DuplicateErrorLogTimeLimit field to given value.


### GetDuplicateAlertLimit

`func (o *GlobalConfigurationResponse) GetDuplicateAlertLimit() int64`

GetDuplicateAlertLimit returns the DuplicateAlertLimit field if non-nil, zero value otherwise.

### GetDuplicateAlertLimitOk

`func (o *GlobalConfigurationResponse) GetDuplicateAlertLimitOk() (*int64, bool)`

GetDuplicateAlertLimitOk returns a tuple with the DuplicateAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateAlertLimit

`func (o *GlobalConfigurationResponse) SetDuplicateAlertLimit(v int64)`

SetDuplicateAlertLimit sets DuplicateAlertLimit field to given value.


### GetDuplicateAlertTimeLimit

`func (o *GlobalConfigurationResponse) GetDuplicateAlertTimeLimit() string`

GetDuplicateAlertTimeLimit returns the DuplicateAlertTimeLimit field if non-nil, zero value otherwise.

### GetDuplicateAlertTimeLimitOk

`func (o *GlobalConfigurationResponse) GetDuplicateAlertTimeLimitOk() (*string, bool)`

GetDuplicateAlertTimeLimitOk returns a tuple with the DuplicateAlertTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateAlertTimeLimit

`func (o *GlobalConfigurationResponse) SetDuplicateAlertTimeLimit(v string)`

SetDuplicateAlertTimeLimit sets DuplicateAlertTimeLimit field to given value.


### GetWritabilityMode

`func (o *GlobalConfigurationResponse) GetWritabilityMode() EnumglobalConfigurationWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *GlobalConfigurationResponse) GetWritabilityModeOk() (*EnumglobalConfigurationWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *GlobalConfigurationResponse) SetWritabilityMode(v EnumglobalConfigurationWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.

### HasWritabilityMode

`func (o *GlobalConfigurationResponse) HasWritabilityMode() bool`

HasWritabilityMode returns a boolean if a field has been set.

### GetUseSharedDatabaseCacheAcrossAllLocalDBBackends

`func (o *GlobalConfigurationResponse) GetUseSharedDatabaseCacheAcrossAllLocalDBBackends() bool`

GetUseSharedDatabaseCacheAcrossAllLocalDBBackends returns the UseSharedDatabaseCacheAcrossAllLocalDBBackends field if non-nil, zero value otherwise.

### GetUseSharedDatabaseCacheAcrossAllLocalDBBackendsOk

`func (o *GlobalConfigurationResponse) GetUseSharedDatabaseCacheAcrossAllLocalDBBackendsOk() (*bool, bool)`

GetUseSharedDatabaseCacheAcrossAllLocalDBBackendsOk returns a tuple with the UseSharedDatabaseCacheAcrossAllLocalDBBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSharedDatabaseCacheAcrossAllLocalDBBackends

`func (o *GlobalConfigurationResponse) SetUseSharedDatabaseCacheAcrossAllLocalDBBackends(v bool)`

SetUseSharedDatabaseCacheAcrossAllLocalDBBackends sets UseSharedDatabaseCacheAcrossAllLocalDBBackends field to given value.

### HasUseSharedDatabaseCacheAcrossAllLocalDBBackends

`func (o *GlobalConfigurationResponse) HasUseSharedDatabaseCacheAcrossAllLocalDBBackends() bool`

HasUseSharedDatabaseCacheAcrossAllLocalDBBackends returns a boolean if a field has been set.

### GetSharedLocalDBBackendDatabaseCachePercent

`func (o *GlobalConfigurationResponse) GetSharedLocalDBBackendDatabaseCachePercent() int64`

GetSharedLocalDBBackendDatabaseCachePercent returns the SharedLocalDBBackendDatabaseCachePercent field if non-nil, zero value otherwise.

### GetSharedLocalDBBackendDatabaseCachePercentOk

`func (o *GlobalConfigurationResponse) GetSharedLocalDBBackendDatabaseCachePercentOk() (*int64, bool)`

GetSharedLocalDBBackendDatabaseCachePercentOk returns a tuple with the SharedLocalDBBackendDatabaseCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedLocalDBBackendDatabaseCachePercent

`func (o *GlobalConfigurationResponse) SetSharedLocalDBBackendDatabaseCachePercent(v int64)`

SetSharedLocalDBBackendDatabaseCachePercent sets SharedLocalDBBackendDatabaseCachePercent field to given value.

### HasSharedLocalDBBackendDatabaseCachePercent

`func (o *GlobalConfigurationResponse) HasSharedLocalDBBackendDatabaseCachePercent() bool`

HasSharedLocalDBBackendDatabaseCachePercent returns a boolean if a field has been set.

### GetUnrecoverableDatabaseErrorMode

`func (o *GlobalConfigurationResponse) GetUnrecoverableDatabaseErrorMode() EnumglobalConfigurationUnrecoverableDatabaseErrorModeProp`

GetUnrecoverableDatabaseErrorMode returns the UnrecoverableDatabaseErrorMode field if non-nil, zero value otherwise.

### GetUnrecoverableDatabaseErrorModeOk

`func (o *GlobalConfigurationResponse) GetUnrecoverableDatabaseErrorModeOk() (*EnumglobalConfigurationUnrecoverableDatabaseErrorModeProp, bool)`

GetUnrecoverableDatabaseErrorModeOk returns a tuple with the UnrecoverableDatabaseErrorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrecoverableDatabaseErrorMode

`func (o *GlobalConfigurationResponse) SetUnrecoverableDatabaseErrorMode(v EnumglobalConfigurationUnrecoverableDatabaseErrorModeProp)`

SetUnrecoverableDatabaseErrorMode sets UnrecoverableDatabaseErrorMode field to given value.

### HasUnrecoverableDatabaseErrorMode

`func (o *GlobalConfigurationResponse) HasUnrecoverableDatabaseErrorMode() bool`

HasUnrecoverableDatabaseErrorMode returns a boolean if a field has been set.

### GetDatabaseOnVirtualizedOrNetworkStorage

`func (o *GlobalConfigurationResponse) GetDatabaseOnVirtualizedOrNetworkStorage() bool`

GetDatabaseOnVirtualizedOrNetworkStorage returns the DatabaseOnVirtualizedOrNetworkStorage field if non-nil, zero value otherwise.

### GetDatabaseOnVirtualizedOrNetworkStorageOk

`func (o *GlobalConfigurationResponse) GetDatabaseOnVirtualizedOrNetworkStorageOk() (*bool, bool)`

GetDatabaseOnVirtualizedOrNetworkStorageOk returns a tuple with the DatabaseOnVirtualizedOrNetworkStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOnVirtualizedOrNetworkStorage

`func (o *GlobalConfigurationResponse) SetDatabaseOnVirtualizedOrNetworkStorage(v bool)`

SetDatabaseOnVirtualizedOrNetworkStorage sets DatabaseOnVirtualizedOrNetworkStorage field to given value.

### HasDatabaseOnVirtualizedOrNetworkStorage

`func (o *GlobalConfigurationResponse) HasDatabaseOnVirtualizedOrNetworkStorage() bool`

HasDatabaseOnVirtualizedOrNetworkStorage returns a boolean if a field has been set.

### GetAutoNameWithEntryUUIDConnectionCriteria

`func (o *GlobalConfigurationResponse) GetAutoNameWithEntryUUIDConnectionCriteria() string`

GetAutoNameWithEntryUUIDConnectionCriteria returns the AutoNameWithEntryUUIDConnectionCriteria field if non-nil, zero value otherwise.

### GetAutoNameWithEntryUUIDConnectionCriteriaOk

`func (o *GlobalConfigurationResponse) GetAutoNameWithEntryUUIDConnectionCriteriaOk() (*string, bool)`

GetAutoNameWithEntryUUIDConnectionCriteriaOk returns a tuple with the AutoNameWithEntryUUIDConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNameWithEntryUUIDConnectionCriteria

`func (o *GlobalConfigurationResponse) SetAutoNameWithEntryUUIDConnectionCriteria(v string)`

SetAutoNameWithEntryUUIDConnectionCriteria sets AutoNameWithEntryUUIDConnectionCriteria field to given value.

### HasAutoNameWithEntryUUIDConnectionCriteria

`func (o *GlobalConfigurationResponse) HasAutoNameWithEntryUUIDConnectionCriteria() bool`

HasAutoNameWithEntryUUIDConnectionCriteria returns a boolean if a field has been set.

### GetAutoNameWithEntryUUIDRequestCriteria

`func (o *GlobalConfigurationResponse) GetAutoNameWithEntryUUIDRequestCriteria() string`

GetAutoNameWithEntryUUIDRequestCriteria returns the AutoNameWithEntryUUIDRequestCriteria field if non-nil, zero value otherwise.

### GetAutoNameWithEntryUUIDRequestCriteriaOk

`func (o *GlobalConfigurationResponse) GetAutoNameWithEntryUUIDRequestCriteriaOk() (*string, bool)`

GetAutoNameWithEntryUUIDRequestCriteriaOk returns a tuple with the AutoNameWithEntryUUIDRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNameWithEntryUUIDRequestCriteria

`func (o *GlobalConfigurationResponse) SetAutoNameWithEntryUUIDRequestCriteria(v string)`

SetAutoNameWithEntryUUIDRequestCriteria sets AutoNameWithEntryUUIDRequestCriteria field to given value.

### HasAutoNameWithEntryUUIDRequestCriteria

`func (o *GlobalConfigurationResponse) HasAutoNameWithEntryUUIDRequestCriteria() bool`

HasAutoNameWithEntryUUIDRequestCriteria returns a boolean if a field has been set.

### GetSoftDeletePolicy

`func (o *GlobalConfigurationResponse) GetSoftDeletePolicy() string`

GetSoftDeletePolicy returns the SoftDeletePolicy field if non-nil, zero value otherwise.

### GetSoftDeletePolicyOk

`func (o *GlobalConfigurationResponse) GetSoftDeletePolicyOk() (*string, bool)`

GetSoftDeletePolicyOk returns a tuple with the SoftDeletePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeletePolicy

`func (o *GlobalConfigurationResponse) SetSoftDeletePolicy(v string)`

SetSoftDeletePolicy sets SoftDeletePolicy field to given value.

### HasSoftDeletePolicy

`func (o *GlobalConfigurationResponse) HasSoftDeletePolicy() bool`

HasSoftDeletePolicy returns a boolean if a field has been set.

### GetSubtreeAccessibilityAlertTimeLimit

`func (o *GlobalConfigurationResponse) GetSubtreeAccessibilityAlertTimeLimit() string`

GetSubtreeAccessibilityAlertTimeLimit returns the SubtreeAccessibilityAlertTimeLimit field if non-nil, zero value otherwise.

### GetSubtreeAccessibilityAlertTimeLimitOk

`func (o *GlobalConfigurationResponse) GetSubtreeAccessibilityAlertTimeLimitOk() (*string, bool)`

GetSubtreeAccessibilityAlertTimeLimitOk returns a tuple with the SubtreeAccessibilityAlertTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeAccessibilityAlertTimeLimit

`func (o *GlobalConfigurationResponse) SetSubtreeAccessibilityAlertTimeLimit(v string)`

SetSubtreeAccessibilityAlertTimeLimit sets SubtreeAccessibilityAlertTimeLimit field to given value.

### HasSubtreeAccessibilityAlertTimeLimit

`func (o *GlobalConfigurationResponse) HasSubtreeAccessibilityAlertTimeLimit() bool`

HasSubtreeAccessibilityAlertTimeLimit returns a boolean if a field has been set.

### GetWarnForBackendsWithMultipleBaseDns

`func (o *GlobalConfigurationResponse) GetWarnForBackendsWithMultipleBaseDns() bool`

GetWarnForBackendsWithMultipleBaseDns returns the WarnForBackendsWithMultipleBaseDns field if non-nil, zero value otherwise.

### GetWarnForBackendsWithMultipleBaseDnsOk

`func (o *GlobalConfigurationResponse) GetWarnForBackendsWithMultipleBaseDnsOk() (*bool, bool)`

GetWarnForBackendsWithMultipleBaseDnsOk returns a tuple with the WarnForBackendsWithMultipleBaseDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnForBackendsWithMultipleBaseDns

`func (o *GlobalConfigurationResponse) SetWarnForBackendsWithMultipleBaseDns(v bool)`

SetWarnForBackendsWithMultipleBaseDns sets WarnForBackendsWithMultipleBaseDns field to given value.

### HasWarnForBackendsWithMultipleBaseDns

`func (o *GlobalConfigurationResponse) HasWarnForBackendsWithMultipleBaseDns() bool`

HasWarnForBackendsWithMultipleBaseDns returns a boolean if a field has been set.

### GetForcedGCPrimeDuration

`func (o *GlobalConfigurationResponse) GetForcedGCPrimeDuration() string`

GetForcedGCPrimeDuration returns the ForcedGCPrimeDuration field if non-nil, zero value otherwise.

### GetForcedGCPrimeDurationOk

`func (o *GlobalConfigurationResponse) GetForcedGCPrimeDurationOk() (*string, bool)`

GetForcedGCPrimeDurationOk returns a tuple with the ForcedGCPrimeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedGCPrimeDuration

`func (o *GlobalConfigurationResponse) SetForcedGCPrimeDuration(v string)`

SetForcedGCPrimeDuration sets ForcedGCPrimeDuration field to given value.

### HasForcedGCPrimeDuration

`func (o *GlobalConfigurationResponse) HasForcedGCPrimeDuration() bool`

HasForcedGCPrimeDuration returns a boolean if a field has been set.

### GetReplicationSetName

`func (o *GlobalConfigurationResponse) GetReplicationSetName() string`

GetReplicationSetName returns the ReplicationSetName field if non-nil, zero value otherwise.

### GetReplicationSetNameOk

`func (o *GlobalConfigurationResponse) GetReplicationSetNameOk() (*string, bool)`

GetReplicationSetNameOk returns a tuple with the ReplicationSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSetName

`func (o *GlobalConfigurationResponse) SetReplicationSetName(v string)`

SetReplicationSetName sets ReplicationSetName field to given value.

### HasReplicationSetName

`func (o *GlobalConfigurationResponse) HasReplicationSetName() bool`

HasReplicationSetName returns a boolean if a field has been set.

### GetStartupMinReplicationBacklogCount

`func (o *GlobalConfigurationResponse) GetStartupMinReplicationBacklogCount() int64`

GetStartupMinReplicationBacklogCount returns the StartupMinReplicationBacklogCount field if non-nil, zero value otherwise.

### GetStartupMinReplicationBacklogCountOk

`func (o *GlobalConfigurationResponse) GetStartupMinReplicationBacklogCountOk() (*int64, bool)`

GetStartupMinReplicationBacklogCountOk returns a tuple with the StartupMinReplicationBacklogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupMinReplicationBacklogCount

`func (o *GlobalConfigurationResponse) SetStartupMinReplicationBacklogCount(v int64)`

SetStartupMinReplicationBacklogCount sets StartupMinReplicationBacklogCount field to given value.


### GetReplicationBacklogCountAlertThreshold

`func (o *GlobalConfigurationResponse) GetReplicationBacklogCountAlertThreshold() int64`

GetReplicationBacklogCountAlertThreshold returns the ReplicationBacklogCountAlertThreshold field if non-nil, zero value otherwise.

### GetReplicationBacklogCountAlertThresholdOk

`func (o *GlobalConfigurationResponse) GetReplicationBacklogCountAlertThresholdOk() (*int64, bool)`

GetReplicationBacklogCountAlertThresholdOk returns a tuple with the ReplicationBacklogCountAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBacklogCountAlertThreshold

`func (o *GlobalConfigurationResponse) SetReplicationBacklogCountAlertThreshold(v int64)`

SetReplicationBacklogCountAlertThreshold sets ReplicationBacklogCountAlertThreshold field to given value.


### GetReplicationBacklogDurationAlertThreshold

`func (o *GlobalConfigurationResponse) GetReplicationBacklogDurationAlertThreshold() string`

GetReplicationBacklogDurationAlertThreshold returns the ReplicationBacklogDurationAlertThreshold field if non-nil, zero value otherwise.

### GetReplicationBacklogDurationAlertThresholdOk

`func (o *GlobalConfigurationResponse) GetReplicationBacklogDurationAlertThresholdOk() (*string, bool)`

GetReplicationBacklogDurationAlertThresholdOk returns a tuple with the ReplicationBacklogDurationAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBacklogDurationAlertThreshold

`func (o *GlobalConfigurationResponse) SetReplicationBacklogDurationAlertThreshold(v string)`

SetReplicationBacklogDurationAlertThreshold sets ReplicationBacklogDurationAlertThreshold field to given value.


### GetReplicationAssuranceSourceTimeoutSuspendDuration

`func (o *GlobalConfigurationResponse) GetReplicationAssuranceSourceTimeoutSuspendDuration() string`

GetReplicationAssuranceSourceTimeoutSuspendDuration returns the ReplicationAssuranceSourceTimeoutSuspendDuration field if non-nil, zero value otherwise.

### GetReplicationAssuranceSourceTimeoutSuspendDurationOk

`func (o *GlobalConfigurationResponse) GetReplicationAssuranceSourceTimeoutSuspendDurationOk() (*string, bool)`

GetReplicationAssuranceSourceTimeoutSuspendDurationOk returns a tuple with the ReplicationAssuranceSourceTimeoutSuspendDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationAssuranceSourceTimeoutSuspendDuration

`func (o *GlobalConfigurationResponse) SetReplicationAssuranceSourceTimeoutSuspendDuration(v string)`

SetReplicationAssuranceSourceTimeoutSuspendDuration sets ReplicationAssuranceSourceTimeoutSuspendDuration field to given value.


### GetReplicationAssuranceSourceBacklogFastStartThreshold

`func (o *GlobalConfigurationResponse) GetReplicationAssuranceSourceBacklogFastStartThreshold() int64`

GetReplicationAssuranceSourceBacklogFastStartThreshold returns the ReplicationAssuranceSourceBacklogFastStartThreshold field if non-nil, zero value otherwise.

### GetReplicationAssuranceSourceBacklogFastStartThresholdOk

`func (o *GlobalConfigurationResponse) GetReplicationAssuranceSourceBacklogFastStartThresholdOk() (*int64, bool)`

GetReplicationAssuranceSourceBacklogFastStartThresholdOk returns a tuple with the ReplicationAssuranceSourceBacklogFastStartThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationAssuranceSourceBacklogFastStartThreshold

`func (o *GlobalConfigurationResponse) SetReplicationAssuranceSourceBacklogFastStartThreshold(v int64)`

SetReplicationAssuranceSourceBacklogFastStartThreshold sets ReplicationAssuranceSourceBacklogFastStartThreshold field to given value.


### GetReplicationHistoryLimit

`func (o *GlobalConfigurationResponse) GetReplicationHistoryLimit() int64`

GetReplicationHistoryLimit returns the ReplicationHistoryLimit field if non-nil, zero value otherwise.

### GetReplicationHistoryLimitOk

`func (o *GlobalConfigurationResponse) GetReplicationHistoryLimitOk() (*int64, bool)`

GetReplicationHistoryLimitOk returns a tuple with the ReplicationHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationHistoryLimit

`func (o *GlobalConfigurationResponse) SetReplicationHistoryLimit(v int64)`

SetReplicationHistoryLimit sets ReplicationHistoryLimit field to given value.

### HasReplicationHistoryLimit

`func (o *GlobalConfigurationResponse) HasReplicationHistoryLimit() bool`

HasReplicationHistoryLimit returns a boolean if a field has been set.

### GetAllowInheritedReplicationOfSubordinateBackends

`func (o *GlobalConfigurationResponse) GetAllowInheritedReplicationOfSubordinateBackends() bool`

GetAllowInheritedReplicationOfSubordinateBackends returns the AllowInheritedReplicationOfSubordinateBackends field if non-nil, zero value otherwise.

### GetAllowInheritedReplicationOfSubordinateBackendsOk

`func (o *GlobalConfigurationResponse) GetAllowInheritedReplicationOfSubordinateBackendsOk() (*bool, bool)`

GetAllowInheritedReplicationOfSubordinateBackendsOk returns a tuple with the AllowInheritedReplicationOfSubordinateBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInheritedReplicationOfSubordinateBackends

`func (o *GlobalConfigurationResponse) SetAllowInheritedReplicationOfSubordinateBackends(v bool)`

SetAllowInheritedReplicationOfSubordinateBackends sets AllowInheritedReplicationOfSubordinateBackends field to given value.


### GetReplicationPurgeObsoleteReplicas

`func (o *GlobalConfigurationResponse) GetReplicationPurgeObsoleteReplicas() bool`

GetReplicationPurgeObsoleteReplicas returns the ReplicationPurgeObsoleteReplicas field if non-nil, zero value otherwise.

### GetReplicationPurgeObsoleteReplicasOk

`func (o *GlobalConfigurationResponse) GetReplicationPurgeObsoleteReplicasOk() (*bool, bool)`

GetReplicationPurgeObsoleteReplicasOk returns a tuple with the ReplicationPurgeObsoleteReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPurgeObsoleteReplicas

`func (o *GlobalConfigurationResponse) SetReplicationPurgeObsoleteReplicas(v bool)`

SetReplicationPurgeObsoleteReplicas sets ReplicationPurgeObsoleteReplicas field to given value.

### HasReplicationPurgeObsoleteReplicas

`func (o *GlobalConfigurationResponse) HasReplicationPurgeObsoleteReplicas() bool`

HasReplicationPurgeObsoleteReplicas returns a boolean if a field has been set.

### GetSmtpServer

`func (o *GlobalConfigurationResponse) GetSmtpServer() []string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *GlobalConfigurationResponse) GetSmtpServerOk() (*[]string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *GlobalConfigurationResponse) SetSmtpServer(v []string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *GlobalConfigurationResponse) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetMaxSMTPConnectionCount

`func (o *GlobalConfigurationResponse) GetMaxSMTPConnectionCount() int64`

GetMaxSMTPConnectionCount returns the MaxSMTPConnectionCount field if non-nil, zero value otherwise.

### GetMaxSMTPConnectionCountOk

`func (o *GlobalConfigurationResponse) GetMaxSMTPConnectionCountOk() (*int64, bool)`

GetMaxSMTPConnectionCountOk returns a tuple with the MaxSMTPConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSMTPConnectionCount

`func (o *GlobalConfigurationResponse) SetMaxSMTPConnectionCount(v int64)`

SetMaxSMTPConnectionCount sets MaxSMTPConnectionCount field to given value.

### HasMaxSMTPConnectionCount

`func (o *GlobalConfigurationResponse) HasMaxSMTPConnectionCount() bool`

HasMaxSMTPConnectionCount returns a boolean if a field has been set.

### GetMaxSMTPConnectionAge

`func (o *GlobalConfigurationResponse) GetMaxSMTPConnectionAge() string`

GetMaxSMTPConnectionAge returns the MaxSMTPConnectionAge field if non-nil, zero value otherwise.

### GetMaxSMTPConnectionAgeOk

`func (o *GlobalConfigurationResponse) GetMaxSMTPConnectionAgeOk() (*string, bool)`

GetMaxSMTPConnectionAgeOk returns a tuple with the MaxSMTPConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSMTPConnectionAge

`func (o *GlobalConfigurationResponse) SetMaxSMTPConnectionAge(v string)`

SetMaxSMTPConnectionAge sets MaxSMTPConnectionAge field to given value.

### HasMaxSMTPConnectionAge

`func (o *GlobalConfigurationResponse) HasMaxSMTPConnectionAge() bool`

HasMaxSMTPConnectionAge returns a boolean if a field has been set.

### GetSmtpConnectionHealthCheckInterval

`func (o *GlobalConfigurationResponse) GetSmtpConnectionHealthCheckInterval() string`

GetSmtpConnectionHealthCheckInterval returns the SmtpConnectionHealthCheckInterval field if non-nil, zero value otherwise.

### GetSmtpConnectionHealthCheckIntervalOk

`func (o *GlobalConfigurationResponse) GetSmtpConnectionHealthCheckIntervalOk() (*string, bool)`

GetSmtpConnectionHealthCheckIntervalOk returns a tuple with the SmtpConnectionHealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpConnectionHealthCheckInterval

`func (o *GlobalConfigurationResponse) SetSmtpConnectionHealthCheckInterval(v string)`

SetSmtpConnectionHealthCheckInterval sets SmtpConnectionHealthCheckInterval field to given value.

### HasSmtpConnectionHealthCheckInterval

`func (o *GlobalConfigurationResponse) HasSmtpConnectionHealthCheckInterval() bool`

HasSmtpConnectionHealthCheckInterval returns a boolean if a field has been set.

### GetAllowedTask

`func (o *GlobalConfigurationResponse) GetAllowedTask() []string`

GetAllowedTask returns the AllowedTask field if non-nil, zero value otherwise.

### GetAllowedTaskOk

`func (o *GlobalConfigurationResponse) GetAllowedTaskOk() (*[]string, bool)`

GetAllowedTaskOk returns a tuple with the AllowedTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTask

`func (o *GlobalConfigurationResponse) SetAllowedTask(v []string)`

SetAllowedTask sets AllowedTask field to given value.

### HasAllowedTask

`func (o *GlobalConfigurationResponse) HasAllowedTask() bool`

HasAllowedTask returns a boolean if a field has been set.

### GetEnableSubOperationTimer

`func (o *GlobalConfigurationResponse) GetEnableSubOperationTimer() bool`

GetEnableSubOperationTimer returns the EnableSubOperationTimer field if non-nil, zero value otherwise.

### GetEnableSubOperationTimerOk

`func (o *GlobalConfigurationResponse) GetEnableSubOperationTimerOk() (*bool, bool)`

GetEnableSubOperationTimerOk returns a tuple with the EnableSubOperationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubOperationTimer

`func (o *GlobalConfigurationResponse) SetEnableSubOperationTimer(v bool)`

SetEnableSubOperationTimer sets EnableSubOperationTimer field to given value.

### HasEnableSubOperationTimer

`func (o *GlobalConfigurationResponse) HasEnableSubOperationTimer() bool`

HasEnableSubOperationTimer returns a boolean if a field has been set.

### GetMaximumShutdownTime

`func (o *GlobalConfigurationResponse) GetMaximumShutdownTime() string`

GetMaximumShutdownTime returns the MaximumShutdownTime field if non-nil, zero value otherwise.

### GetMaximumShutdownTimeOk

`func (o *GlobalConfigurationResponse) GetMaximumShutdownTimeOk() (*string, bool)`

GetMaximumShutdownTimeOk returns a tuple with the MaximumShutdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumShutdownTime

`func (o *GlobalConfigurationResponse) SetMaximumShutdownTime(v string)`

SetMaximumShutdownTime sets MaximumShutdownTime field to given value.

### HasMaximumShutdownTime

`func (o *GlobalConfigurationResponse) HasMaximumShutdownTime() bool`

HasMaximumShutdownTime returns a boolean if a field has been set.

### GetNetworkAddressCacheTTL

`func (o *GlobalConfigurationResponse) GetNetworkAddressCacheTTL() string`

GetNetworkAddressCacheTTL returns the NetworkAddressCacheTTL field if non-nil, zero value otherwise.

### GetNetworkAddressCacheTTLOk

`func (o *GlobalConfigurationResponse) GetNetworkAddressCacheTTLOk() (*string, bool)`

GetNetworkAddressCacheTTLOk returns a tuple with the NetworkAddressCacheTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressCacheTTL

`func (o *GlobalConfigurationResponse) SetNetworkAddressCacheTTL(v string)`

SetNetworkAddressCacheTTL sets NetworkAddressCacheTTL field to given value.

### HasNetworkAddressCacheTTL

`func (o *GlobalConfigurationResponse) HasNetworkAddressCacheTTL() bool`

HasNetworkAddressCacheTTL returns a boolean if a field has been set.

### GetNetworkAddressOutageCacheEnabled

`func (o *GlobalConfigurationResponse) GetNetworkAddressOutageCacheEnabled() bool`

GetNetworkAddressOutageCacheEnabled returns the NetworkAddressOutageCacheEnabled field if non-nil, zero value otherwise.

### GetNetworkAddressOutageCacheEnabledOk

`func (o *GlobalConfigurationResponse) GetNetworkAddressOutageCacheEnabledOk() (*bool, bool)`

GetNetworkAddressOutageCacheEnabledOk returns a tuple with the NetworkAddressOutageCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressOutageCacheEnabled

`func (o *GlobalConfigurationResponse) SetNetworkAddressOutageCacheEnabled(v bool)`

SetNetworkAddressOutageCacheEnabled sets NetworkAddressOutageCacheEnabled field to given value.

### HasNetworkAddressOutageCacheEnabled

`func (o *GlobalConfigurationResponse) HasNetworkAddressOutageCacheEnabled() bool`

HasNetworkAddressOutageCacheEnabled returns a boolean if a field has been set.

### GetTrackedApplication

`func (o *GlobalConfigurationResponse) GetTrackedApplication() []string`

GetTrackedApplication returns the TrackedApplication field if non-nil, zero value otherwise.

### GetTrackedApplicationOk

`func (o *GlobalConfigurationResponse) GetTrackedApplicationOk() (*[]string, bool)`

GetTrackedApplicationOk returns a tuple with the TrackedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedApplication

`func (o *GlobalConfigurationResponse) SetTrackedApplication(v []string)`

SetTrackedApplication sets TrackedApplication field to given value.

### HasTrackedApplication

`func (o *GlobalConfigurationResponse) HasTrackedApplication() bool`

HasTrackedApplication returns a boolean if a field has been set.

### GetJmxValueBehavior

`func (o *GlobalConfigurationResponse) GetJmxValueBehavior() EnumglobalConfigurationJmxValueBehaviorProp`

GetJmxValueBehavior returns the JmxValueBehavior field if non-nil, zero value otherwise.

### GetJmxValueBehaviorOk

`func (o *GlobalConfigurationResponse) GetJmxValueBehaviorOk() (*EnumglobalConfigurationJmxValueBehaviorProp, bool)`

GetJmxValueBehaviorOk returns a tuple with the JmxValueBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxValueBehavior

`func (o *GlobalConfigurationResponse) SetJmxValueBehavior(v EnumglobalConfigurationJmxValueBehaviorProp)`

SetJmxValueBehavior sets JmxValueBehavior field to given value.

### HasJmxValueBehavior

`func (o *GlobalConfigurationResponse) HasJmxValueBehavior() bool`

HasJmxValueBehavior returns a boolean if a field has been set.

### GetJmxUseLegacyMbeanNames

`func (o *GlobalConfigurationResponse) GetJmxUseLegacyMbeanNames() bool`

GetJmxUseLegacyMbeanNames returns the JmxUseLegacyMbeanNames field if non-nil, zero value otherwise.

### GetJmxUseLegacyMbeanNamesOk

`func (o *GlobalConfigurationResponse) GetJmxUseLegacyMbeanNamesOk() (*bool, bool)`

GetJmxUseLegacyMbeanNamesOk returns a tuple with the JmxUseLegacyMbeanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxUseLegacyMbeanNames

`func (o *GlobalConfigurationResponse) SetJmxUseLegacyMbeanNames(v bool)`

SetJmxUseLegacyMbeanNames sets JmxUseLegacyMbeanNames field to given value.

### HasJmxUseLegacyMbeanNames

`func (o *GlobalConfigurationResponse) HasJmxUseLegacyMbeanNames() bool`

HasJmxUseLegacyMbeanNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


