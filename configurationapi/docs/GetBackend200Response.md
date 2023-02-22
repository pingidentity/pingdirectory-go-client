# GetBackend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnummetricsBackendSchemaUrn**](EnummetricsBackendSchemaUrn.md) |  | 
**Id** | **string** | Name of the Backend | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**SchemaEntryDN** | Pointer to **[]string** | Defines the base DNs of the subtrees in which the schema information is published in addition to the value included in the base-dn property. | [optional] 
**ShowAllAttributes** | **bool** | Indicates whether to treat all attributes in the schema entry as if they were user attributes regardless of their configuration. | 
**ReadOnlySchemaFile** | Pointer to **[]string** | Specifies the name of a file (which must exist in the config/schema directory) containing schema elements that should be considered read-only. Any schema definitions contained in read-only files cannot be altered by external clients. | [optional] 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**BackupDirectory** | **[]string** | Specifies the path to a backup directory containing one or more backups for a particular backend. | 
**IsPrivateBackend** | Pointer to **bool** | Indicates whether this backend should be considered a private backend in the server. Private backends are meant for storing server-internal information and should not be used for user or application data. | [optional] 
**LdifFile** | **string** | Specifies the path to the LDIF file that serves as the backing file for this backend. | 
**TrustStoreFile** | **string** | Specifies the path to the file that stores the trust information. | 
**TrustStoreType** | Pointer to **string** | Specifies the format for the data in the key store file. | [optional] 
**TrustStorePin** | Pointer to **string** | Specifies the clear-text PIN needed to access the Trust Store Backend. | [optional] 
**TrustStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the Trust Store Backend. | [optional] 
**TrustStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the Trust Store Backend. | [optional] 
**DbDirectory** | **string** | Specifies the path to the filesystem directory that is used to hold the Berkeley DB Java Edition database files containing the data for this backend. The files for this backend are stored in a sub-directory named after the backend-id. | 
**DbDirectoryPermissions** | Pointer to **string** | Specifies the permissions that should be applied to the directory containing the backend database files and to directories and files created during backup or LDIF export of the backend. | [optional] 
**DbCachePercent** | Pointer to **int32** | Specifies the percentage of JVM memory to allocate to the database cache. | [optional] 
**JeProperty** | Pointer to **[]string** | Specifies the database and environment properties for the Berkeley DB Java Edition database serving the data for this backend. | [optional] 
**ChangelogWriteBatchSize** | Pointer to **int32** | Specifies the number of changelog entries written in a single database transaction. | [optional] 
**ChangelogPurgeBatchSize** | Pointer to **int32** | Specifies the number of changelog entries purged in a single database transaction. | [optional] 
**ChangelogWriteQueueCapacity** | Pointer to **int32** | Specifies the capacity of the changelog write queue in number of changes. | [optional] 
**IndexIncludeAttribute** | Pointer to **[]string** | Specifies which attribute types are to be specifically included in the set of attribute indexes maintained on the changelog. If this property does not have any values then no attribute types are indexed. | [optional] 
**IndexExcludeAttribute** | Pointer to **[]string** | Specifies which attribute types are to be specifically excluded from the set of attribute indexes maintained on the changelog. This property is useful when the index-include-attribute property contains one of the special values \&quot;*\&quot; and \&quot;+\&quot;. | [optional] 
**ChangelogMaximumAge** | **string** | Changes are guaranteed to be maintained in the changelog database for at least this duration. Setting target-database-size can allow additional changes to be maintained up to the configured size on disk. | 
**TargetDatabaseSize** | Pointer to **string** | The changelog database is allowed to grow up to this size on disk even if changes are older than the configured changelog-maximum-age. | [optional] 
**ChangelogEntryIncludeBaseDN** | Pointer to **[]string** | The base DNs for branches in the data for which to record changes in the changelog. | [optional] 
**ChangelogEntryExcludeBaseDN** | Pointer to **[]string** | The base DNs for branches in the data for which no changelog records should be generated. | [optional] 
**ChangelogEntryIncludeFilter** | Pointer to **[]string** | A filter that indicates which changelog entries should actually be stored in the changelog. Note that this filter is evaluated against the changelog entry itself and not against the entry that was the target of the change referenced by the changelog entry. This filter may target any attributes that appear in changelog entries with the exception of the changeNumber and entry-size-bytes attributes, since they will not be known at the time of the filter evaluation. | [optional] 
**ChangelogEntryExcludeFilter** | Pointer to **[]string** | A filter that indicates which changelog entries should be excluded from the changelog. Note that this filter is evaluated against the changelog entry itself and not against the entry that was the target of the change referenced by the changelog entry. This filter may target any attributes that appear in changelog entries with the exception of the changeNumber and entry-size-bytes attributes, since they will not be known at the time of the filter evaluation. | [optional] 
**ChangelogIncludeAttribute** | Pointer to **[]string** | Specifies which attribute types will be included in a changelog entry for ADD and MODIFY operations. | [optional] 
**ChangelogExcludeAttribute** | Pointer to **[]string** | Specifies a set of attribute types that should be excluded in a changelog entry for ADD and MODIFY operations. | [optional] 
**ChangelogDeletedEntryIncludeAttribute** | Pointer to **[]string** | Specifies a set of attribute types that should be included in a changelog entry for DELETE operations. | [optional] 
**ChangelogDeletedEntryExcludeAttribute** | Pointer to **[]string** | Specifies a set of attribute types that should be excluded from a changelog entry for DELETE operations. | [optional] 
**ChangelogIncludeKeyAttribute** | Pointer to **[]string** | Specifies which attribute types will be included in a changelog entry on every change. | [optional] 
**ChangelogMaxBeforeAfterValues** | Pointer to **int32** | This controls whether all attribute values for a modified attribute (even those values that have not changed) will be included in the changelog entry. If the number of attribute values does not exceed this limit, then all values for the modified attribute will be included in the changelog entry. | [optional] 
**WriteLastmodAttributes** | Pointer to **bool** | Specifies whether values of creatorsName, createTimestamp, modifiersName and modifyTimestamp attributes will be written to changelog entries. | [optional] 
**UseReversibleForm** | Pointer to **bool** | Specifies whether the changelog should provide enough information to be able to revert the changes if desired. | [optional] 
**IncludeVirtualAttributes** | Pointer to [**[]EnumbackendIncludeVirtualAttributesProp**](EnumbackendIncludeVirtualAttributesProp.md) |  | [optional] 
**ApplyAccessControlsToChangelogEntryContents** | Pointer to **bool** | Indicates whether the contents of changelog entries should be subject to access control and sensitive attribute evaluation such that the contents of attributes like changes, deletedEntryAttrs, ds-changelog-entry-key-attr-values, ds-changelog-before-values, and ds-changelog-after-values may be altered based on attributes the user can see in the target entry. | [optional] 
**ReportExcludedChangelogAttributes** | Pointer to [**EnumbackendReportExcludedChangelogAttributesProp**](EnumbackendReportExcludedChangelogAttributesProp.md) |  | [optional] 
**SoftDeleteEntryIncludedOperation** | Pointer to [**[]EnumbackendSoftDeleteEntryIncludedOperationProp**](EnumbackendSoftDeleteEntryIncludedOperationProp.md) |  | [optional] 
**UncachedId2entryCacheMode** | Pointer to [**EnumbackendUncachedId2entryCacheModeProp**](EnumbackendUncachedId2entryCacheModeProp.md) |  | [optional] 
**UncachedAttributeCriteria** | Pointer to **string** | The criteria that will be used to identify attributes that should be written into the uncached-id2entry database rather than the id2entry database. This will only be used for entries in which the associated uncached-entry-criteria does not indicate that the entire entry should be uncached. | [optional] 
**UncachedEntryCriteria** | Pointer to **string** | The criteria that will be used to identify entries that should be written into the uncached-id2entry database rather than the id2entry database. | [optional] 
**SetDegradedAlertForUntrustedIndex** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state when this Local DB Backend has an index whose contents cannot be trusted. | [optional] 
**ReturnUnavailableForUntrustedIndex** | Pointer to **bool** | Determines whether the Directory Server returns UNAVAILABLE for any LDAP search operation in this Local DB Backend that would use an index whose contents cannot be trusted. | [optional] 
**ProcessFiltersWithUndefinedAttributeTypes** | Pointer to **bool** | Determines whether the Directory Server should continue filter processing for LDAP search operations in this Local DB Backend that includes a search filter with an attribute that is not defined in the schema. This will only apply if check-schema is enabled in the global configuration. | [optional] 
**CompactCommonParentDN** | Pointer to **[]string** | Provides a DN of an entry that may be the parent for a large number of entries in the backend. This may be used to help increase the space efficiency when encoding entries for storage. | [optional] 
**CompressEntries** | Pointer to **bool** | Indicates whether the backend should attempt to compress entries before storing them in the database. | [optional] 
**HashEntries** | Pointer to **bool** | Indicates whether to calculate and store a message digest of the entry contents along with the entry data, in order to provide a means of verifying the integrity of the entry data. | [optional] 
**DbNumCleanerThreads** | Pointer to **int32** | Specifies the number of threads that the backend should maintain to keep the database log files at or near the desired utilization. A value of zero indicates that the number of cleaner threads should be automatically configured based on the number of available CPUs. | [optional] 
**DbCleanerMinUtilization** | Pointer to **int32** | Specifies the minimum percentage of \&quot;live\&quot; data that the database cleaner attempts to keep in database log files. | [optional] 
**DbEvictorCriticalPercentage** | Pointer to **int32** | Specifies the percentage over the configured maximum that the database cache is allowed to grow. It is recommended to set this value slightly above zero when the database is too large to fully cache in memory. In this case, a dedicated background evictor thread is used to perform evictions once the cache fills up reducing the possibility that server threads are blocked. | [optional] 
**DbCheckpointerWakeupInterval** | Pointer to **string** | Specifies the maximum length of time that should pass between checkpoints. | [optional] 
**DbBackgroundSyncInterval** | Pointer to **string** | Specifies the interval to use when performing background synchronous writes in the database environment in order to smooth overall write performance and increase data durability. A value of \&quot;0 s\&quot; will disable background synchronous writes. | [optional] 
**DbUseThreadLocalHandles** | Pointer to **bool** | Indicates whether to use thread-local database handles to reduce contention in the backend. | [optional] 
**DbLogFileMax** | Pointer to **string** | Specifies the maximum size for a database log file. | [optional] 
**DbLoggingLevel** | Pointer to **string** | Specifies the log level that should be used by the database when it is writing information into the je.info file. | [optional] 
**DefaultCacheMode** | Pointer to [**EnumbackendDefaultCacheModeProp**](EnumbackendDefaultCacheModeProp.md) |  | [optional] 
**Id2entryCacheMode** | Pointer to [**EnumbackendId2entryCacheModeProp**](EnumbackendId2entryCacheModeProp.md) |  | [optional] 
**Dn2idCacheMode** | Pointer to [**EnumbackendDn2idCacheModeProp**](EnumbackendDn2idCacheModeProp.md) |  | [optional] 
**Id2childrenCacheMode** | Pointer to [**EnumbackendId2childrenCacheModeProp**](EnumbackendId2childrenCacheModeProp.md) |  | [optional] 
**Id2subtreeCacheMode** | Pointer to [**EnumbackendId2subtreeCacheModeProp**](EnumbackendId2subtreeCacheModeProp.md) |  | [optional] 
**Dn2uriCacheMode** | Pointer to [**EnumbackendDn2uriCacheModeProp**](EnumbackendDn2uriCacheModeProp.md) |  | [optional] 
**PrimeMethod** | Pointer to [**[]EnumbackendPrimeMethodProp**](EnumbackendPrimeMethodProp.md) |  | [optional] 
**PrimeThreadCount** | Pointer to **int32** | Specifies the number of threads to use when priming. At present, this applies only to the preload and cursor-across-indexes prime methods. | [optional] 
**PrimeTimeLimit** | Pointer to **string** | Specifies the maximum length of time that the backend prime should be allowed to run. A duration of zero seconds indicates that there should not be a time limit. | [optional] 
**PrimeAllIndexes** | Pointer to **bool** | Indicates whether to prime all indexes associated with this backend, or to only prime the specified set of indexes (as configured with the system-index-to-prime property for the system indexes, and the prime-index property in the attribute index definition for attribute indexes). | [optional] 
**SystemIndexToPrime** | Pointer to [**[]EnumbackendSystemIndexToPrimeProp**](EnumbackendSystemIndexToPrimeProp.md) |  | [optional] 
**SystemIndexToPrimeInternalNodesOnly** | Pointer to [**[]EnumbackendSystemIndexToPrimeInternalNodesOnlyProp**](EnumbackendSystemIndexToPrimeInternalNodesOnlyProp.md) |  | [optional] 
**BackgroundPrime** | Pointer to **bool** | Indicates whether to attempt to perform the prime using a background thread if possible. If background priming is enabled, then the Directory Server may be allowed to accept client connections and process requests while the prime is in progress. | [optional] 
**IndexEntryLimit** | Pointer to **int32** | Specifies the maximum number of entries that are allowed to match a given index key before that particular index key is no longer maintained. | [optional] 
**CompositeIndexEntryLimit** | Pointer to **int32** | Specifies the maximum number of entries that are allowed to match a given composite index key before that particular composite index key is no longer maintained. | [optional] 
**Id2childrenIndexEntryLimit** | Pointer to **int32** | Specifies the maximum number of entry IDs to maintain for each entry in the id2children system index (which keeps track of the immediate children for an entry, to assist in otherwise unindexed searches with a single-level scope). A value of 0 means there is no limit, however this could have a big impact on database size on disk and on server performance. | [optional] 
**Id2subtreeIndexEntryLimit** | Pointer to **int32** | Specifies the maximum number of entry IDs to maintain for each entry in the id2subtree system index (which keeps track of all descendants below an entry, to assist in otherwise unindexed searches with a whole-subtree or subordinate subtree scope). A value of 0 means there is no limit, however this could have a big impact on database size on disk and on server performance. | [optional] 
**ImportTempDirectory** | **string** | Specifies the location of the directory that is used to hold temporary information during the index post-processing phase of an LDIF import. | 
**ImportThreadCount** | Pointer to **int32** | Specifies the number of threads to use for concurrent processing during an LDIF import. | [optional] 
**ExportThreadCount** | Pointer to **int32** | Specifies the number of threads to use for concurrently retrieving and encoding entries during an LDIF export. | [optional] 
**DbImportCachePercent** | Pointer to **int32** | The percentage of JVM memory to allocate to the database cache during import operations. | [optional] 
**DbTxnWriteNoSync** | Pointer to **bool** | Indicates whether the database should synchronously flush data as it is written to disk. | [optional] 
**DeadlockRetryLimit** | Pointer to **int32** | Specifies the number of times that the server should retry an attempted operation in the backend if a deadlock results from two concurrent requests that interfere with each other in a conflicting manner. | [optional] 
**ExternalTxnDefaultBackendLockBehavior** | Pointer to [**EnumbackendExternalTxnDefaultBackendLockBehaviorProp**](EnumbackendExternalTxnDefaultBackendLockBehaviorProp.md) |  | [optional] 
**SingleWriterLockBehavior** | Pointer to [**EnumbackendSingleWriterLockBehaviorProp**](EnumbackendSingleWriterLockBehaviorProp.md) |  | [optional] 
**SubtreeDeleteSizeLimit** | Pointer to **int32** | Specifies the maximum number of entries that may be deleted from the backend when using the subtree delete control. | [optional] 
**NumRecentChanges** | Pointer to **int32** | Specifies the number of recent LDAP entry changes per replica for which the backend keeps a record to allow replication to recover in the event that the server is abruptly terminated. Increasing this value can lead to an increased peak server modification rate as well as increased replication throughput. | [optional] 
**OfflineProcessDatabaseOpenTimeout** | Pointer to **string** | Specifies a timeout duration which will be used for opening the database environment by an offline process, such as export-ldif. | [optional] 
**InsignificantConfigArchiveAttribute** | Pointer to **[]string** | The name or OID of an attribute type that is considered insignificant for the purpose of maintaining the configuration archive. | [optional] 
**MirroredSubtreePeerPollingInterval** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait before polling the peer servers in the topology to determine if there are any changes in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. | [optional] 
**MirroredSubtreeEntryUpdateTimeout** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for an update operation (add, delete, modify and modify-dn) on an entry to be applied on all servers in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. | [optional] 
**MirroredSubtreeSearchTimeout** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for a search operation to complete. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. Search requests that take longer than this timeout will be canceled and considered failures. | [optional] 
**TaskBackingFile** | **string** | Specifies the path to the backing file for storing information about the tasks configured in the server. | 
**MaximumInitialTaskLogMessagesToRetain** | Pointer to **int32** | The maximum number of log messages to retain in each task entry from the beginning of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks. | [optional] 
**MaximumFinalTaskLogMessagesToRetain** | Pointer to **int32** | The maximum number of log messages to retain in each task entry from the end of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks. | [optional] 
**TaskRetentionTime** | Pointer to **string** | Specifies the length of time that task entries should be retained after processing on the associated task has been completed. | [optional] 
**NotificationSenderAddress** | Pointer to **string** | Specifies the email address to use as the sender address (that is, the \&quot;From:\&quot; address) for notification mail messages generated when a task completes execution. | [optional] 
**AlertRetentionTime** | **string** | Specifies the maximum length of time that information about generated alerts should be maintained before they will be purged. | 
**MaxAlerts** | Pointer to **int32** | Specifies the maximum number of alerts that should be retained. If more alerts than this configured maximum are generated within the alert retention time, then the oldest alerts will be purged to achieve this maximum. | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumbackendDisabledAlertTypeProp**](EnumbackendDisabledAlertTypeProp.md) |  | [optional] 
**AlarmRetentionTime** | **string** | Specifies the maximum length of time that information about raised alarms should be maintained before they will be purged. | 
**MaxAlarms** | Pointer to **int32** | Specifies the maximum number of alarms that should be retained. If more alarms than this configured maximum are generated within the alarm retention time, then the oldest alarms will be purged to achieve this maximum. Only alarms at normal severity will be purged. | [optional] 
**StorageDir** | **string** | Specifies the path to the directory that will be used to store queued samples. | 
**MetricsDir** | **string** | Specifies the path to the directory that contains metric definitions. | 
**SampleFlushInterval** | Pointer to **string** | Period when samples are flushed to disk. | [optional] 
**RetentionPolicy** | **[]string** | The retention policy to use for the Metrics Backend . | 

## Methods

### NewGetBackend200Response

`func NewGetBackend200Response(schemas []EnummetricsBackendSchemaUrn, id string, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, showAllAttributes bool, enabled bool, backupDirectory []string, ldifFile string, trustStoreFile string, dbDirectory string, changelogMaximumAge string, importTempDirectory string, taskBackingFile string, alertRetentionTime string, alarmRetentionTime string, storageDir string, metricsDir string, retentionPolicy []string, ) *GetBackend200Response`

NewGetBackend200Response instantiates a new GetBackend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackend200ResponseWithDefaults

`func NewGetBackend200ResponseWithDefaults() *GetBackend200Response`

NewGetBackend200ResponseWithDefaults instantiates a new GetBackend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBackend200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBackend200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBackend200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetBackend200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetBackend200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetBackend200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetBackend200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetBackend200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GetBackend200Response) GetSchemas() []EnummetricsBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetBackend200Response) GetSchemasOk() (*[]EnummetricsBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetBackend200Response) SetSchemas(v []EnummetricsBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetBackend200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBackend200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBackend200Response) SetId(v string)`

SetId sets Id field to given value.


### GetBackendID

`func (o *GetBackend200Response) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *GetBackend200Response) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *GetBackend200Response) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *GetBackend200Response) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *GetBackend200Response) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *GetBackend200Response) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *GetBackend200Response) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *GetBackend200Response) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *GetBackend200Response) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetSchemaEntryDN

`func (o *GetBackend200Response) GetSchemaEntryDN() []string`

GetSchemaEntryDN returns the SchemaEntryDN field if non-nil, zero value otherwise.

### GetSchemaEntryDNOk

`func (o *GetBackend200Response) GetSchemaEntryDNOk() (*[]string, bool)`

GetSchemaEntryDNOk returns a tuple with the SchemaEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaEntryDN

`func (o *GetBackend200Response) SetSchemaEntryDN(v []string)`

SetSchemaEntryDN sets SchemaEntryDN field to given value.

### HasSchemaEntryDN

`func (o *GetBackend200Response) HasSchemaEntryDN() bool`

HasSchemaEntryDN returns a boolean if a field has been set.

### GetShowAllAttributes

`func (o *GetBackend200Response) GetShowAllAttributes() bool`

GetShowAllAttributes returns the ShowAllAttributes field if non-nil, zero value otherwise.

### GetShowAllAttributesOk

`func (o *GetBackend200Response) GetShowAllAttributesOk() (*bool, bool)`

GetShowAllAttributesOk returns a tuple with the ShowAllAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllAttributes

`func (o *GetBackend200Response) SetShowAllAttributes(v bool)`

SetShowAllAttributes sets ShowAllAttributes field to given value.


### GetReadOnlySchemaFile

`func (o *GetBackend200Response) GetReadOnlySchemaFile() []string`

GetReadOnlySchemaFile returns the ReadOnlySchemaFile field if non-nil, zero value otherwise.

### GetReadOnlySchemaFileOk

`func (o *GetBackend200Response) GetReadOnlySchemaFileOk() (*[]string, bool)`

GetReadOnlySchemaFileOk returns a tuple with the ReadOnlySchemaFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySchemaFile

`func (o *GetBackend200Response) SetReadOnlySchemaFile(v []string)`

SetReadOnlySchemaFile sets ReadOnlySchemaFile field to given value.

### HasReadOnlySchemaFile

`func (o *GetBackend200Response) HasReadOnlySchemaFile() bool`

HasReadOnlySchemaFile returns a boolean if a field has been set.

### GetDescription

`func (o *GetBackend200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBackend200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBackend200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetBackend200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetBackend200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetBackend200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetBackend200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *GetBackend200Response) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *GetBackend200Response) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *GetBackend200Response) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *GetBackend200Response) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *GetBackend200Response) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *GetBackend200Response) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *GetBackend200Response) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *GetBackend200Response) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *GetBackend200Response) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *GetBackend200Response) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *GetBackend200Response) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *GetBackend200Response) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *GetBackend200Response) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *GetBackend200Response) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *GetBackend200Response) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *GetBackend200Response) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetBackupDirectory

`func (o *GetBackend200Response) GetBackupDirectory() []string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *GetBackend200Response) GetBackupDirectoryOk() (*[]string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *GetBackend200Response) SetBackupDirectory(v []string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIsPrivateBackend

`func (o *GetBackend200Response) GetIsPrivateBackend() bool`

GetIsPrivateBackend returns the IsPrivateBackend field if non-nil, zero value otherwise.

### GetIsPrivateBackendOk

`func (o *GetBackend200Response) GetIsPrivateBackendOk() (*bool, bool)`

GetIsPrivateBackendOk returns a tuple with the IsPrivateBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateBackend

`func (o *GetBackend200Response) SetIsPrivateBackend(v bool)`

SetIsPrivateBackend sets IsPrivateBackend field to given value.

### HasIsPrivateBackend

`func (o *GetBackend200Response) HasIsPrivateBackend() bool`

HasIsPrivateBackend returns a boolean if a field has been set.

### GetLdifFile

`func (o *GetBackend200Response) GetLdifFile() string`

GetLdifFile returns the LdifFile field if non-nil, zero value otherwise.

### GetLdifFileOk

`func (o *GetBackend200Response) GetLdifFileOk() (*string, bool)`

GetLdifFileOk returns a tuple with the LdifFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifFile

`func (o *GetBackend200Response) SetLdifFile(v string)`

SetLdifFile sets LdifFile field to given value.


### GetTrustStoreFile

`func (o *GetBackend200Response) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *GetBackend200Response) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *GetBackend200Response) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.


### GetTrustStoreType

`func (o *GetBackend200Response) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *GetBackend200Response) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *GetBackend200Response) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *GetBackend200Response) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *GetBackend200Response) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *GetBackend200Response) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *GetBackend200Response) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *GetBackend200Response) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStorePinFile

`func (o *GetBackend200Response) GetTrustStorePinFile() string`

GetTrustStorePinFile returns the TrustStorePinFile field if non-nil, zero value otherwise.

### GetTrustStorePinFileOk

`func (o *GetBackend200Response) GetTrustStorePinFileOk() (*string, bool)`

GetTrustStorePinFileOk returns a tuple with the TrustStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinFile

`func (o *GetBackend200Response) SetTrustStorePinFile(v string)`

SetTrustStorePinFile sets TrustStorePinFile field to given value.

### HasTrustStorePinFile

`func (o *GetBackend200Response) HasTrustStorePinFile() bool`

HasTrustStorePinFile returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *GetBackend200Response) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *GetBackend200Response) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *GetBackend200Response) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *GetBackend200Response) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetDbDirectory

`func (o *GetBackend200Response) GetDbDirectory() string`

GetDbDirectory returns the DbDirectory field if non-nil, zero value otherwise.

### GetDbDirectoryOk

`func (o *GetBackend200Response) GetDbDirectoryOk() (*string, bool)`

GetDbDirectoryOk returns a tuple with the DbDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectory

`func (o *GetBackend200Response) SetDbDirectory(v string)`

SetDbDirectory sets DbDirectory field to given value.


### GetDbDirectoryPermissions

`func (o *GetBackend200Response) GetDbDirectoryPermissions() string`

GetDbDirectoryPermissions returns the DbDirectoryPermissions field if non-nil, zero value otherwise.

### GetDbDirectoryPermissionsOk

`func (o *GetBackend200Response) GetDbDirectoryPermissionsOk() (*string, bool)`

GetDbDirectoryPermissionsOk returns a tuple with the DbDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectoryPermissions

`func (o *GetBackend200Response) SetDbDirectoryPermissions(v string)`

SetDbDirectoryPermissions sets DbDirectoryPermissions field to given value.

### HasDbDirectoryPermissions

`func (o *GetBackend200Response) HasDbDirectoryPermissions() bool`

HasDbDirectoryPermissions returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *GetBackend200Response) GetDbCachePercent() int32`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *GetBackend200Response) GetDbCachePercentOk() (*int32, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *GetBackend200Response) SetDbCachePercent(v int32)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *GetBackend200Response) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetJeProperty

`func (o *GetBackend200Response) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *GetBackend200Response) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *GetBackend200Response) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *GetBackend200Response) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetChangelogWriteBatchSize

`func (o *GetBackend200Response) GetChangelogWriteBatchSize() int32`

GetChangelogWriteBatchSize returns the ChangelogWriteBatchSize field if non-nil, zero value otherwise.

### GetChangelogWriteBatchSizeOk

`func (o *GetBackend200Response) GetChangelogWriteBatchSizeOk() (*int32, bool)`

GetChangelogWriteBatchSizeOk returns a tuple with the ChangelogWriteBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogWriteBatchSize

`func (o *GetBackend200Response) SetChangelogWriteBatchSize(v int32)`

SetChangelogWriteBatchSize sets ChangelogWriteBatchSize field to given value.

### HasChangelogWriteBatchSize

`func (o *GetBackend200Response) HasChangelogWriteBatchSize() bool`

HasChangelogWriteBatchSize returns a boolean if a field has been set.

### GetChangelogPurgeBatchSize

`func (o *GetBackend200Response) GetChangelogPurgeBatchSize() int32`

GetChangelogPurgeBatchSize returns the ChangelogPurgeBatchSize field if non-nil, zero value otherwise.

### GetChangelogPurgeBatchSizeOk

`func (o *GetBackend200Response) GetChangelogPurgeBatchSizeOk() (*int32, bool)`

GetChangelogPurgeBatchSizeOk returns a tuple with the ChangelogPurgeBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPurgeBatchSize

`func (o *GetBackend200Response) SetChangelogPurgeBatchSize(v int32)`

SetChangelogPurgeBatchSize sets ChangelogPurgeBatchSize field to given value.

### HasChangelogPurgeBatchSize

`func (o *GetBackend200Response) HasChangelogPurgeBatchSize() bool`

HasChangelogPurgeBatchSize returns a boolean if a field has been set.

### GetChangelogWriteQueueCapacity

`func (o *GetBackend200Response) GetChangelogWriteQueueCapacity() int32`

GetChangelogWriteQueueCapacity returns the ChangelogWriteQueueCapacity field if non-nil, zero value otherwise.

### GetChangelogWriteQueueCapacityOk

`func (o *GetBackend200Response) GetChangelogWriteQueueCapacityOk() (*int32, bool)`

GetChangelogWriteQueueCapacityOk returns a tuple with the ChangelogWriteQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogWriteQueueCapacity

`func (o *GetBackend200Response) SetChangelogWriteQueueCapacity(v int32)`

SetChangelogWriteQueueCapacity sets ChangelogWriteQueueCapacity field to given value.

### HasChangelogWriteQueueCapacity

`func (o *GetBackend200Response) HasChangelogWriteQueueCapacity() bool`

HasChangelogWriteQueueCapacity returns a boolean if a field has been set.

### GetIndexIncludeAttribute

`func (o *GetBackend200Response) GetIndexIncludeAttribute() []string`

GetIndexIncludeAttribute returns the IndexIncludeAttribute field if non-nil, zero value otherwise.

### GetIndexIncludeAttributeOk

`func (o *GetBackend200Response) GetIndexIncludeAttributeOk() (*[]string, bool)`

GetIndexIncludeAttributeOk returns a tuple with the IndexIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIncludeAttribute

`func (o *GetBackend200Response) SetIndexIncludeAttribute(v []string)`

SetIndexIncludeAttribute sets IndexIncludeAttribute field to given value.

### HasIndexIncludeAttribute

`func (o *GetBackend200Response) HasIndexIncludeAttribute() bool`

HasIndexIncludeAttribute returns a boolean if a field has been set.

### GetIndexExcludeAttribute

`func (o *GetBackend200Response) GetIndexExcludeAttribute() []string`

GetIndexExcludeAttribute returns the IndexExcludeAttribute field if non-nil, zero value otherwise.

### GetIndexExcludeAttributeOk

`func (o *GetBackend200Response) GetIndexExcludeAttributeOk() (*[]string, bool)`

GetIndexExcludeAttributeOk returns a tuple with the IndexExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexExcludeAttribute

`func (o *GetBackend200Response) SetIndexExcludeAttribute(v []string)`

SetIndexExcludeAttribute sets IndexExcludeAttribute field to given value.

### HasIndexExcludeAttribute

`func (o *GetBackend200Response) HasIndexExcludeAttribute() bool`

HasIndexExcludeAttribute returns a boolean if a field has been set.

### GetChangelogMaximumAge

`func (o *GetBackend200Response) GetChangelogMaximumAge() string`

GetChangelogMaximumAge returns the ChangelogMaximumAge field if non-nil, zero value otherwise.

### GetChangelogMaximumAgeOk

`func (o *GetBackend200Response) GetChangelogMaximumAgeOk() (*string, bool)`

GetChangelogMaximumAgeOk returns a tuple with the ChangelogMaximumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogMaximumAge

`func (o *GetBackend200Response) SetChangelogMaximumAge(v string)`

SetChangelogMaximumAge sets ChangelogMaximumAge field to given value.


### GetTargetDatabaseSize

`func (o *GetBackend200Response) GetTargetDatabaseSize() string`

GetTargetDatabaseSize returns the TargetDatabaseSize field if non-nil, zero value otherwise.

### GetTargetDatabaseSizeOk

`func (o *GetBackend200Response) GetTargetDatabaseSizeOk() (*string, bool)`

GetTargetDatabaseSizeOk returns a tuple with the TargetDatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatabaseSize

`func (o *GetBackend200Response) SetTargetDatabaseSize(v string)`

SetTargetDatabaseSize sets TargetDatabaseSize field to given value.

### HasTargetDatabaseSize

`func (o *GetBackend200Response) HasTargetDatabaseSize() bool`

HasTargetDatabaseSize returns a boolean if a field has been set.

### GetChangelogEntryIncludeBaseDN

`func (o *GetBackend200Response) GetChangelogEntryIncludeBaseDN() []string`

GetChangelogEntryIncludeBaseDN returns the ChangelogEntryIncludeBaseDN field if non-nil, zero value otherwise.

### GetChangelogEntryIncludeBaseDNOk

`func (o *GetBackend200Response) GetChangelogEntryIncludeBaseDNOk() (*[]string, bool)`

GetChangelogEntryIncludeBaseDNOk returns a tuple with the ChangelogEntryIncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryIncludeBaseDN

`func (o *GetBackend200Response) SetChangelogEntryIncludeBaseDN(v []string)`

SetChangelogEntryIncludeBaseDN sets ChangelogEntryIncludeBaseDN field to given value.

### HasChangelogEntryIncludeBaseDN

`func (o *GetBackend200Response) HasChangelogEntryIncludeBaseDN() bool`

HasChangelogEntryIncludeBaseDN returns a boolean if a field has been set.

### GetChangelogEntryExcludeBaseDN

`func (o *GetBackend200Response) GetChangelogEntryExcludeBaseDN() []string`

GetChangelogEntryExcludeBaseDN returns the ChangelogEntryExcludeBaseDN field if non-nil, zero value otherwise.

### GetChangelogEntryExcludeBaseDNOk

`func (o *GetBackend200Response) GetChangelogEntryExcludeBaseDNOk() (*[]string, bool)`

GetChangelogEntryExcludeBaseDNOk returns a tuple with the ChangelogEntryExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryExcludeBaseDN

`func (o *GetBackend200Response) SetChangelogEntryExcludeBaseDN(v []string)`

SetChangelogEntryExcludeBaseDN sets ChangelogEntryExcludeBaseDN field to given value.

### HasChangelogEntryExcludeBaseDN

`func (o *GetBackend200Response) HasChangelogEntryExcludeBaseDN() bool`

HasChangelogEntryExcludeBaseDN returns a boolean if a field has been set.

### GetChangelogEntryIncludeFilter

`func (o *GetBackend200Response) GetChangelogEntryIncludeFilter() []string`

GetChangelogEntryIncludeFilter returns the ChangelogEntryIncludeFilter field if non-nil, zero value otherwise.

### GetChangelogEntryIncludeFilterOk

`func (o *GetBackend200Response) GetChangelogEntryIncludeFilterOk() (*[]string, bool)`

GetChangelogEntryIncludeFilterOk returns a tuple with the ChangelogEntryIncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryIncludeFilter

`func (o *GetBackend200Response) SetChangelogEntryIncludeFilter(v []string)`

SetChangelogEntryIncludeFilter sets ChangelogEntryIncludeFilter field to given value.

### HasChangelogEntryIncludeFilter

`func (o *GetBackend200Response) HasChangelogEntryIncludeFilter() bool`

HasChangelogEntryIncludeFilter returns a boolean if a field has been set.

### GetChangelogEntryExcludeFilter

`func (o *GetBackend200Response) GetChangelogEntryExcludeFilter() []string`

GetChangelogEntryExcludeFilter returns the ChangelogEntryExcludeFilter field if non-nil, zero value otherwise.

### GetChangelogEntryExcludeFilterOk

`func (o *GetBackend200Response) GetChangelogEntryExcludeFilterOk() (*[]string, bool)`

GetChangelogEntryExcludeFilterOk returns a tuple with the ChangelogEntryExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryExcludeFilter

`func (o *GetBackend200Response) SetChangelogEntryExcludeFilter(v []string)`

SetChangelogEntryExcludeFilter sets ChangelogEntryExcludeFilter field to given value.

### HasChangelogEntryExcludeFilter

`func (o *GetBackend200Response) HasChangelogEntryExcludeFilter() bool`

HasChangelogEntryExcludeFilter returns a boolean if a field has been set.

### GetChangelogIncludeAttribute

`func (o *GetBackend200Response) GetChangelogIncludeAttribute() []string`

GetChangelogIncludeAttribute returns the ChangelogIncludeAttribute field if non-nil, zero value otherwise.

### GetChangelogIncludeAttributeOk

`func (o *GetBackend200Response) GetChangelogIncludeAttributeOk() (*[]string, bool)`

GetChangelogIncludeAttributeOk returns a tuple with the ChangelogIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogIncludeAttribute

`func (o *GetBackend200Response) SetChangelogIncludeAttribute(v []string)`

SetChangelogIncludeAttribute sets ChangelogIncludeAttribute field to given value.

### HasChangelogIncludeAttribute

`func (o *GetBackend200Response) HasChangelogIncludeAttribute() bool`

HasChangelogIncludeAttribute returns a boolean if a field has been set.

### GetChangelogExcludeAttribute

`func (o *GetBackend200Response) GetChangelogExcludeAttribute() []string`

GetChangelogExcludeAttribute returns the ChangelogExcludeAttribute field if non-nil, zero value otherwise.

### GetChangelogExcludeAttributeOk

`func (o *GetBackend200Response) GetChangelogExcludeAttributeOk() (*[]string, bool)`

GetChangelogExcludeAttributeOk returns a tuple with the ChangelogExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogExcludeAttribute

`func (o *GetBackend200Response) SetChangelogExcludeAttribute(v []string)`

SetChangelogExcludeAttribute sets ChangelogExcludeAttribute field to given value.

### HasChangelogExcludeAttribute

`func (o *GetBackend200Response) HasChangelogExcludeAttribute() bool`

HasChangelogExcludeAttribute returns a boolean if a field has been set.

### GetChangelogDeletedEntryIncludeAttribute

`func (o *GetBackend200Response) GetChangelogDeletedEntryIncludeAttribute() []string`

GetChangelogDeletedEntryIncludeAttribute returns the ChangelogDeletedEntryIncludeAttribute field if non-nil, zero value otherwise.

### GetChangelogDeletedEntryIncludeAttributeOk

`func (o *GetBackend200Response) GetChangelogDeletedEntryIncludeAttributeOk() (*[]string, bool)`

GetChangelogDeletedEntryIncludeAttributeOk returns a tuple with the ChangelogDeletedEntryIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogDeletedEntryIncludeAttribute

`func (o *GetBackend200Response) SetChangelogDeletedEntryIncludeAttribute(v []string)`

SetChangelogDeletedEntryIncludeAttribute sets ChangelogDeletedEntryIncludeAttribute field to given value.

### HasChangelogDeletedEntryIncludeAttribute

`func (o *GetBackend200Response) HasChangelogDeletedEntryIncludeAttribute() bool`

HasChangelogDeletedEntryIncludeAttribute returns a boolean if a field has been set.

### GetChangelogDeletedEntryExcludeAttribute

`func (o *GetBackend200Response) GetChangelogDeletedEntryExcludeAttribute() []string`

GetChangelogDeletedEntryExcludeAttribute returns the ChangelogDeletedEntryExcludeAttribute field if non-nil, zero value otherwise.

### GetChangelogDeletedEntryExcludeAttributeOk

`func (o *GetBackend200Response) GetChangelogDeletedEntryExcludeAttributeOk() (*[]string, bool)`

GetChangelogDeletedEntryExcludeAttributeOk returns a tuple with the ChangelogDeletedEntryExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogDeletedEntryExcludeAttribute

`func (o *GetBackend200Response) SetChangelogDeletedEntryExcludeAttribute(v []string)`

SetChangelogDeletedEntryExcludeAttribute sets ChangelogDeletedEntryExcludeAttribute field to given value.

### HasChangelogDeletedEntryExcludeAttribute

`func (o *GetBackend200Response) HasChangelogDeletedEntryExcludeAttribute() bool`

HasChangelogDeletedEntryExcludeAttribute returns a boolean if a field has been set.

### GetChangelogIncludeKeyAttribute

`func (o *GetBackend200Response) GetChangelogIncludeKeyAttribute() []string`

GetChangelogIncludeKeyAttribute returns the ChangelogIncludeKeyAttribute field if non-nil, zero value otherwise.

### GetChangelogIncludeKeyAttributeOk

`func (o *GetBackend200Response) GetChangelogIncludeKeyAttributeOk() (*[]string, bool)`

GetChangelogIncludeKeyAttributeOk returns a tuple with the ChangelogIncludeKeyAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogIncludeKeyAttribute

`func (o *GetBackend200Response) SetChangelogIncludeKeyAttribute(v []string)`

SetChangelogIncludeKeyAttribute sets ChangelogIncludeKeyAttribute field to given value.

### HasChangelogIncludeKeyAttribute

`func (o *GetBackend200Response) HasChangelogIncludeKeyAttribute() bool`

HasChangelogIncludeKeyAttribute returns a boolean if a field has been set.

### GetChangelogMaxBeforeAfterValues

`func (o *GetBackend200Response) GetChangelogMaxBeforeAfterValues() int32`

GetChangelogMaxBeforeAfterValues returns the ChangelogMaxBeforeAfterValues field if non-nil, zero value otherwise.

### GetChangelogMaxBeforeAfterValuesOk

`func (o *GetBackend200Response) GetChangelogMaxBeforeAfterValuesOk() (*int32, bool)`

GetChangelogMaxBeforeAfterValuesOk returns a tuple with the ChangelogMaxBeforeAfterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogMaxBeforeAfterValues

`func (o *GetBackend200Response) SetChangelogMaxBeforeAfterValues(v int32)`

SetChangelogMaxBeforeAfterValues sets ChangelogMaxBeforeAfterValues field to given value.

### HasChangelogMaxBeforeAfterValues

`func (o *GetBackend200Response) HasChangelogMaxBeforeAfterValues() bool`

HasChangelogMaxBeforeAfterValues returns a boolean if a field has been set.

### GetWriteLastmodAttributes

`func (o *GetBackend200Response) GetWriteLastmodAttributes() bool`

GetWriteLastmodAttributes returns the WriteLastmodAttributes field if non-nil, zero value otherwise.

### GetWriteLastmodAttributesOk

`func (o *GetBackend200Response) GetWriteLastmodAttributesOk() (*bool, bool)`

GetWriteLastmodAttributesOk returns a tuple with the WriteLastmodAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLastmodAttributes

`func (o *GetBackend200Response) SetWriteLastmodAttributes(v bool)`

SetWriteLastmodAttributes sets WriteLastmodAttributes field to given value.

### HasWriteLastmodAttributes

`func (o *GetBackend200Response) HasWriteLastmodAttributes() bool`

HasWriteLastmodAttributes returns a boolean if a field has been set.

### GetUseReversibleForm

`func (o *GetBackend200Response) GetUseReversibleForm() bool`

GetUseReversibleForm returns the UseReversibleForm field if non-nil, zero value otherwise.

### GetUseReversibleFormOk

`func (o *GetBackend200Response) GetUseReversibleFormOk() (*bool, bool)`

GetUseReversibleFormOk returns a tuple with the UseReversibleForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReversibleForm

`func (o *GetBackend200Response) SetUseReversibleForm(v bool)`

SetUseReversibleForm sets UseReversibleForm field to given value.

### HasUseReversibleForm

`func (o *GetBackend200Response) HasUseReversibleForm() bool`

HasUseReversibleForm returns a boolean if a field has been set.

### GetIncludeVirtualAttributes

`func (o *GetBackend200Response) GetIncludeVirtualAttributes() []EnumbackendIncludeVirtualAttributesProp`

GetIncludeVirtualAttributes returns the IncludeVirtualAttributes field if non-nil, zero value otherwise.

### GetIncludeVirtualAttributesOk

`func (o *GetBackend200Response) GetIncludeVirtualAttributesOk() (*[]EnumbackendIncludeVirtualAttributesProp, bool)`

GetIncludeVirtualAttributesOk returns a tuple with the IncludeVirtualAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVirtualAttributes

`func (o *GetBackend200Response) SetIncludeVirtualAttributes(v []EnumbackendIncludeVirtualAttributesProp)`

SetIncludeVirtualAttributes sets IncludeVirtualAttributes field to given value.

### HasIncludeVirtualAttributes

`func (o *GetBackend200Response) HasIncludeVirtualAttributes() bool`

HasIncludeVirtualAttributes returns a boolean if a field has been set.

### GetApplyAccessControlsToChangelogEntryContents

`func (o *GetBackend200Response) GetApplyAccessControlsToChangelogEntryContents() bool`

GetApplyAccessControlsToChangelogEntryContents returns the ApplyAccessControlsToChangelogEntryContents field if non-nil, zero value otherwise.

### GetApplyAccessControlsToChangelogEntryContentsOk

`func (o *GetBackend200Response) GetApplyAccessControlsToChangelogEntryContentsOk() (*bool, bool)`

GetApplyAccessControlsToChangelogEntryContentsOk returns a tuple with the ApplyAccessControlsToChangelogEntryContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAccessControlsToChangelogEntryContents

`func (o *GetBackend200Response) SetApplyAccessControlsToChangelogEntryContents(v bool)`

SetApplyAccessControlsToChangelogEntryContents sets ApplyAccessControlsToChangelogEntryContents field to given value.

### HasApplyAccessControlsToChangelogEntryContents

`func (o *GetBackend200Response) HasApplyAccessControlsToChangelogEntryContents() bool`

HasApplyAccessControlsToChangelogEntryContents returns a boolean if a field has been set.

### GetReportExcludedChangelogAttributes

`func (o *GetBackend200Response) GetReportExcludedChangelogAttributes() EnumbackendReportExcludedChangelogAttributesProp`

GetReportExcludedChangelogAttributes returns the ReportExcludedChangelogAttributes field if non-nil, zero value otherwise.

### GetReportExcludedChangelogAttributesOk

`func (o *GetBackend200Response) GetReportExcludedChangelogAttributesOk() (*EnumbackendReportExcludedChangelogAttributesProp, bool)`

GetReportExcludedChangelogAttributesOk returns a tuple with the ReportExcludedChangelogAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExcludedChangelogAttributes

`func (o *GetBackend200Response) SetReportExcludedChangelogAttributes(v EnumbackendReportExcludedChangelogAttributesProp)`

SetReportExcludedChangelogAttributes sets ReportExcludedChangelogAttributes field to given value.

### HasReportExcludedChangelogAttributes

`func (o *GetBackend200Response) HasReportExcludedChangelogAttributes() bool`

HasReportExcludedChangelogAttributes returns a boolean if a field has been set.

### GetSoftDeleteEntryIncludedOperation

`func (o *GetBackend200Response) GetSoftDeleteEntryIncludedOperation() []EnumbackendSoftDeleteEntryIncludedOperationProp`

GetSoftDeleteEntryIncludedOperation returns the SoftDeleteEntryIncludedOperation field if non-nil, zero value otherwise.

### GetSoftDeleteEntryIncludedOperationOk

`func (o *GetBackend200Response) GetSoftDeleteEntryIncludedOperationOk() (*[]EnumbackendSoftDeleteEntryIncludedOperationProp, bool)`

GetSoftDeleteEntryIncludedOperationOk returns a tuple with the SoftDeleteEntryIncludedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteEntryIncludedOperation

`func (o *GetBackend200Response) SetSoftDeleteEntryIncludedOperation(v []EnumbackendSoftDeleteEntryIncludedOperationProp)`

SetSoftDeleteEntryIncludedOperation sets SoftDeleteEntryIncludedOperation field to given value.

### HasSoftDeleteEntryIncludedOperation

`func (o *GetBackend200Response) HasSoftDeleteEntryIncludedOperation() bool`

HasSoftDeleteEntryIncludedOperation returns a boolean if a field has been set.

### GetUncachedId2entryCacheMode

`func (o *GetBackend200Response) GetUncachedId2entryCacheMode() EnumbackendUncachedId2entryCacheModeProp`

GetUncachedId2entryCacheMode returns the UncachedId2entryCacheMode field if non-nil, zero value otherwise.

### GetUncachedId2entryCacheModeOk

`func (o *GetBackend200Response) GetUncachedId2entryCacheModeOk() (*EnumbackendUncachedId2entryCacheModeProp, bool)`

GetUncachedId2entryCacheModeOk returns a tuple with the UncachedId2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedId2entryCacheMode

`func (o *GetBackend200Response) SetUncachedId2entryCacheMode(v EnumbackendUncachedId2entryCacheModeProp)`

SetUncachedId2entryCacheMode sets UncachedId2entryCacheMode field to given value.

### HasUncachedId2entryCacheMode

`func (o *GetBackend200Response) HasUncachedId2entryCacheMode() bool`

HasUncachedId2entryCacheMode returns a boolean if a field has been set.

### GetUncachedAttributeCriteria

`func (o *GetBackend200Response) GetUncachedAttributeCriteria() string`

GetUncachedAttributeCriteria returns the UncachedAttributeCriteria field if non-nil, zero value otherwise.

### GetUncachedAttributeCriteriaOk

`func (o *GetBackend200Response) GetUncachedAttributeCriteriaOk() (*string, bool)`

GetUncachedAttributeCriteriaOk returns a tuple with the UncachedAttributeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedAttributeCriteria

`func (o *GetBackend200Response) SetUncachedAttributeCriteria(v string)`

SetUncachedAttributeCriteria sets UncachedAttributeCriteria field to given value.

### HasUncachedAttributeCriteria

`func (o *GetBackend200Response) HasUncachedAttributeCriteria() bool`

HasUncachedAttributeCriteria returns a boolean if a field has been set.

### GetUncachedEntryCriteria

`func (o *GetBackend200Response) GetUncachedEntryCriteria() string`

GetUncachedEntryCriteria returns the UncachedEntryCriteria field if non-nil, zero value otherwise.

### GetUncachedEntryCriteriaOk

`func (o *GetBackend200Response) GetUncachedEntryCriteriaOk() (*string, bool)`

GetUncachedEntryCriteriaOk returns a tuple with the UncachedEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedEntryCriteria

`func (o *GetBackend200Response) SetUncachedEntryCriteria(v string)`

SetUncachedEntryCriteria sets UncachedEntryCriteria field to given value.

### HasUncachedEntryCriteria

`func (o *GetBackend200Response) HasUncachedEntryCriteria() bool`

HasUncachedEntryCriteria returns a boolean if a field has been set.

### GetSetDegradedAlertForUntrustedIndex

`func (o *GetBackend200Response) GetSetDegradedAlertForUntrustedIndex() bool`

GetSetDegradedAlertForUntrustedIndex returns the SetDegradedAlertForUntrustedIndex field if non-nil, zero value otherwise.

### GetSetDegradedAlertForUntrustedIndexOk

`func (o *GetBackend200Response) GetSetDegradedAlertForUntrustedIndexOk() (*bool, bool)`

GetSetDegradedAlertForUntrustedIndexOk returns a tuple with the SetDegradedAlertForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertForUntrustedIndex

`func (o *GetBackend200Response) SetSetDegradedAlertForUntrustedIndex(v bool)`

SetSetDegradedAlertForUntrustedIndex sets SetDegradedAlertForUntrustedIndex field to given value.

### HasSetDegradedAlertForUntrustedIndex

`func (o *GetBackend200Response) HasSetDegradedAlertForUntrustedIndex() bool`

HasSetDegradedAlertForUntrustedIndex returns a boolean if a field has been set.

### GetReturnUnavailableForUntrustedIndex

`func (o *GetBackend200Response) GetReturnUnavailableForUntrustedIndex() bool`

GetReturnUnavailableForUntrustedIndex returns the ReturnUnavailableForUntrustedIndex field if non-nil, zero value otherwise.

### GetReturnUnavailableForUntrustedIndexOk

`func (o *GetBackend200Response) GetReturnUnavailableForUntrustedIndexOk() (*bool, bool)`

GetReturnUnavailableForUntrustedIndexOk returns a tuple with the ReturnUnavailableForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableForUntrustedIndex

`func (o *GetBackend200Response) SetReturnUnavailableForUntrustedIndex(v bool)`

SetReturnUnavailableForUntrustedIndex sets ReturnUnavailableForUntrustedIndex field to given value.

### HasReturnUnavailableForUntrustedIndex

`func (o *GetBackend200Response) HasReturnUnavailableForUntrustedIndex() bool`

HasReturnUnavailableForUntrustedIndex returns a boolean if a field has been set.

### GetProcessFiltersWithUndefinedAttributeTypes

`func (o *GetBackend200Response) GetProcessFiltersWithUndefinedAttributeTypes() bool`

GetProcessFiltersWithUndefinedAttributeTypes returns the ProcessFiltersWithUndefinedAttributeTypes field if non-nil, zero value otherwise.

### GetProcessFiltersWithUndefinedAttributeTypesOk

`func (o *GetBackend200Response) GetProcessFiltersWithUndefinedAttributeTypesOk() (*bool, bool)`

GetProcessFiltersWithUndefinedAttributeTypesOk returns a tuple with the ProcessFiltersWithUndefinedAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFiltersWithUndefinedAttributeTypes

`func (o *GetBackend200Response) SetProcessFiltersWithUndefinedAttributeTypes(v bool)`

SetProcessFiltersWithUndefinedAttributeTypes sets ProcessFiltersWithUndefinedAttributeTypes field to given value.

### HasProcessFiltersWithUndefinedAttributeTypes

`func (o *GetBackend200Response) HasProcessFiltersWithUndefinedAttributeTypes() bool`

HasProcessFiltersWithUndefinedAttributeTypes returns a boolean if a field has been set.

### GetCompactCommonParentDN

`func (o *GetBackend200Response) GetCompactCommonParentDN() []string`

GetCompactCommonParentDN returns the CompactCommonParentDN field if non-nil, zero value otherwise.

### GetCompactCommonParentDNOk

`func (o *GetBackend200Response) GetCompactCommonParentDNOk() (*[]string, bool)`

GetCompactCommonParentDNOk returns a tuple with the CompactCommonParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactCommonParentDN

`func (o *GetBackend200Response) SetCompactCommonParentDN(v []string)`

SetCompactCommonParentDN sets CompactCommonParentDN field to given value.

### HasCompactCommonParentDN

`func (o *GetBackend200Response) HasCompactCommonParentDN() bool`

HasCompactCommonParentDN returns a boolean if a field has been set.

### GetCompressEntries

`func (o *GetBackend200Response) GetCompressEntries() bool`

GetCompressEntries returns the CompressEntries field if non-nil, zero value otherwise.

### GetCompressEntriesOk

`func (o *GetBackend200Response) GetCompressEntriesOk() (*bool, bool)`

GetCompressEntriesOk returns a tuple with the CompressEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressEntries

`func (o *GetBackend200Response) SetCompressEntries(v bool)`

SetCompressEntries sets CompressEntries field to given value.

### HasCompressEntries

`func (o *GetBackend200Response) HasCompressEntries() bool`

HasCompressEntries returns a boolean if a field has been set.

### GetHashEntries

`func (o *GetBackend200Response) GetHashEntries() bool`

GetHashEntries returns the HashEntries field if non-nil, zero value otherwise.

### GetHashEntriesOk

`func (o *GetBackend200Response) GetHashEntriesOk() (*bool, bool)`

GetHashEntriesOk returns a tuple with the HashEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashEntries

`func (o *GetBackend200Response) SetHashEntries(v bool)`

SetHashEntries sets HashEntries field to given value.

### HasHashEntries

`func (o *GetBackend200Response) HasHashEntries() bool`

HasHashEntries returns a boolean if a field has been set.

### GetDbNumCleanerThreads

`func (o *GetBackend200Response) GetDbNumCleanerThreads() int32`

GetDbNumCleanerThreads returns the DbNumCleanerThreads field if non-nil, zero value otherwise.

### GetDbNumCleanerThreadsOk

`func (o *GetBackend200Response) GetDbNumCleanerThreadsOk() (*int32, bool)`

GetDbNumCleanerThreadsOk returns a tuple with the DbNumCleanerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbNumCleanerThreads

`func (o *GetBackend200Response) SetDbNumCleanerThreads(v int32)`

SetDbNumCleanerThreads sets DbNumCleanerThreads field to given value.

### HasDbNumCleanerThreads

`func (o *GetBackend200Response) HasDbNumCleanerThreads() bool`

HasDbNumCleanerThreads returns a boolean if a field has been set.

### GetDbCleanerMinUtilization

`func (o *GetBackend200Response) GetDbCleanerMinUtilization() int32`

GetDbCleanerMinUtilization returns the DbCleanerMinUtilization field if non-nil, zero value otherwise.

### GetDbCleanerMinUtilizationOk

`func (o *GetBackend200Response) GetDbCleanerMinUtilizationOk() (*int32, bool)`

GetDbCleanerMinUtilizationOk returns a tuple with the DbCleanerMinUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCleanerMinUtilization

`func (o *GetBackend200Response) SetDbCleanerMinUtilization(v int32)`

SetDbCleanerMinUtilization sets DbCleanerMinUtilization field to given value.

### HasDbCleanerMinUtilization

`func (o *GetBackend200Response) HasDbCleanerMinUtilization() bool`

HasDbCleanerMinUtilization returns a boolean if a field has been set.

### GetDbEvictorCriticalPercentage

`func (o *GetBackend200Response) GetDbEvictorCriticalPercentage() int32`

GetDbEvictorCriticalPercentage returns the DbEvictorCriticalPercentage field if non-nil, zero value otherwise.

### GetDbEvictorCriticalPercentageOk

`func (o *GetBackend200Response) GetDbEvictorCriticalPercentageOk() (*int32, bool)`

GetDbEvictorCriticalPercentageOk returns a tuple with the DbEvictorCriticalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEvictorCriticalPercentage

`func (o *GetBackend200Response) SetDbEvictorCriticalPercentage(v int32)`

SetDbEvictorCriticalPercentage sets DbEvictorCriticalPercentage field to given value.

### HasDbEvictorCriticalPercentage

`func (o *GetBackend200Response) HasDbEvictorCriticalPercentage() bool`

HasDbEvictorCriticalPercentage returns a boolean if a field has been set.

### GetDbCheckpointerWakeupInterval

`func (o *GetBackend200Response) GetDbCheckpointerWakeupInterval() string`

GetDbCheckpointerWakeupInterval returns the DbCheckpointerWakeupInterval field if non-nil, zero value otherwise.

### GetDbCheckpointerWakeupIntervalOk

`func (o *GetBackend200Response) GetDbCheckpointerWakeupIntervalOk() (*string, bool)`

GetDbCheckpointerWakeupIntervalOk returns a tuple with the DbCheckpointerWakeupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCheckpointerWakeupInterval

`func (o *GetBackend200Response) SetDbCheckpointerWakeupInterval(v string)`

SetDbCheckpointerWakeupInterval sets DbCheckpointerWakeupInterval field to given value.

### HasDbCheckpointerWakeupInterval

`func (o *GetBackend200Response) HasDbCheckpointerWakeupInterval() bool`

HasDbCheckpointerWakeupInterval returns a boolean if a field has been set.

### GetDbBackgroundSyncInterval

`func (o *GetBackend200Response) GetDbBackgroundSyncInterval() string`

GetDbBackgroundSyncInterval returns the DbBackgroundSyncInterval field if non-nil, zero value otherwise.

### GetDbBackgroundSyncIntervalOk

`func (o *GetBackend200Response) GetDbBackgroundSyncIntervalOk() (*string, bool)`

GetDbBackgroundSyncIntervalOk returns a tuple with the DbBackgroundSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbBackgroundSyncInterval

`func (o *GetBackend200Response) SetDbBackgroundSyncInterval(v string)`

SetDbBackgroundSyncInterval sets DbBackgroundSyncInterval field to given value.

### HasDbBackgroundSyncInterval

`func (o *GetBackend200Response) HasDbBackgroundSyncInterval() bool`

HasDbBackgroundSyncInterval returns a boolean if a field has been set.

### GetDbUseThreadLocalHandles

`func (o *GetBackend200Response) GetDbUseThreadLocalHandles() bool`

GetDbUseThreadLocalHandles returns the DbUseThreadLocalHandles field if non-nil, zero value otherwise.

### GetDbUseThreadLocalHandlesOk

`func (o *GetBackend200Response) GetDbUseThreadLocalHandlesOk() (*bool, bool)`

GetDbUseThreadLocalHandlesOk returns a tuple with the DbUseThreadLocalHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUseThreadLocalHandles

`func (o *GetBackend200Response) SetDbUseThreadLocalHandles(v bool)`

SetDbUseThreadLocalHandles sets DbUseThreadLocalHandles field to given value.

### HasDbUseThreadLocalHandles

`func (o *GetBackend200Response) HasDbUseThreadLocalHandles() bool`

HasDbUseThreadLocalHandles returns a boolean if a field has been set.

### GetDbLogFileMax

`func (o *GetBackend200Response) GetDbLogFileMax() string`

GetDbLogFileMax returns the DbLogFileMax field if non-nil, zero value otherwise.

### GetDbLogFileMaxOk

`func (o *GetBackend200Response) GetDbLogFileMaxOk() (*string, bool)`

GetDbLogFileMaxOk returns a tuple with the DbLogFileMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLogFileMax

`func (o *GetBackend200Response) SetDbLogFileMax(v string)`

SetDbLogFileMax sets DbLogFileMax field to given value.

### HasDbLogFileMax

`func (o *GetBackend200Response) HasDbLogFileMax() bool`

HasDbLogFileMax returns a boolean if a field has been set.

### GetDbLoggingLevel

`func (o *GetBackend200Response) GetDbLoggingLevel() string`

GetDbLoggingLevel returns the DbLoggingLevel field if non-nil, zero value otherwise.

### GetDbLoggingLevelOk

`func (o *GetBackend200Response) GetDbLoggingLevelOk() (*string, bool)`

GetDbLoggingLevelOk returns a tuple with the DbLoggingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLoggingLevel

`func (o *GetBackend200Response) SetDbLoggingLevel(v string)`

SetDbLoggingLevel sets DbLoggingLevel field to given value.

### HasDbLoggingLevel

`func (o *GetBackend200Response) HasDbLoggingLevel() bool`

HasDbLoggingLevel returns a boolean if a field has been set.

### GetDefaultCacheMode

`func (o *GetBackend200Response) GetDefaultCacheMode() EnumbackendDefaultCacheModeProp`

GetDefaultCacheMode returns the DefaultCacheMode field if non-nil, zero value otherwise.

### GetDefaultCacheModeOk

`func (o *GetBackend200Response) GetDefaultCacheModeOk() (*EnumbackendDefaultCacheModeProp, bool)`

GetDefaultCacheModeOk returns a tuple with the DefaultCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCacheMode

`func (o *GetBackend200Response) SetDefaultCacheMode(v EnumbackendDefaultCacheModeProp)`

SetDefaultCacheMode sets DefaultCacheMode field to given value.

### HasDefaultCacheMode

`func (o *GetBackend200Response) HasDefaultCacheMode() bool`

HasDefaultCacheMode returns a boolean if a field has been set.

### GetId2entryCacheMode

`func (o *GetBackend200Response) GetId2entryCacheMode() EnumbackendId2entryCacheModeProp`

GetId2entryCacheMode returns the Id2entryCacheMode field if non-nil, zero value otherwise.

### GetId2entryCacheModeOk

`func (o *GetBackend200Response) GetId2entryCacheModeOk() (*EnumbackendId2entryCacheModeProp, bool)`

GetId2entryCacheModeOk returns a tuple with the Id2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2entryCacheMode

`func (o *GetBackend200Response) SetId2entryCacheMode(v EnumbackendId2entryCacheModeProp)`

SetId2entryCacheMode sets Id2entryCacheMode field to given value.

### HasId2entryCacheMode

`func (o *GetBackend200Response) HasId2entryCacheMode() bool`

HasId2entryCacheMode returns a boolean if a field has been set.

### GetDn2idCacheMode

`func (o *GetBackend200Response) GetDn2idCacheMode() EnumbackendDn2idCacheModeProp`

GetDn2idCacheMode returns the Dn2idCacheMode field if non-nil, zero value otherwise.

### GetDn2idCacheModeOk

`func (o *GetBackend200Response) GetDn2idCacheModeOk() (*EnumbackendDn2idCacheModeProp, bool)`

GetDn2idCacheModeOk returns a tuple with the Dn2idCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2idCacheMode

`func (o *GetBackend200Response) SetDn2idCacheMode(v EnumbackendDn2idCacheModeProp)`

SetDn2idCacheMode sets Dn2idCacheMode field to given value.

### HasDn2idCacheMode

`func (o *GetBackend200Response) HasDn2idCacheMode() bool`

HasDn2idCacheMode returns a boolean if a field has been set.

### GetId2childrenCacheMode

`func (o *GetBackend200Response) GetId2childrenCacheMode() EnumbackendId2childrenCacheModeProp`

GetId2childrenCacheMode returns the Id2childrenCacheMode field if non-nil, zero value otherwise.

### GetId2childrenCacheModeOk

`func (o *GetBackend200Response) GetId2childrenCacheModeOk() (*EnumbackendId2childrenCacheModeProp, bool)`

GetId2childrenCacheModeOk returns a tuple with the Id2childrenCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenCacheMode

`func (o *GetBackend200Response) SetId2childrenCacheMode(v EnumbackendId2childrenCacheModeProp)`

SetId2childrenCacheMode sets Id2childrenCacheMode field to given value.

### HasId2childrenCacheMode

`func (o *GetBackend200Response) HasId2childrenCacheMode() bool`

HasId2childrenCacheMode returns a boolean if a field has been set.

### GetId2subtreeCacheMode

`func (o *GetBackend200Response) GetId2subtreeCacheMode() EnumbackendId2subtreeCacheModeProp`

GetId2subtreeCacheMode returns the Id2subtreeCacheMode field if non-nil, zero value otherwise.

### GetId2subtreeCacheModeOk

`func (o *GetBackend200Response) GetId2subtreeCacheModeOk() (*EnumbackendId2subtreeCacheModeProp, bool)`

GetId2subtreeCacheModeOk returns a tuple with the Id2subtreeCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeCacheMode

`func (o *GetBackend200Response) SetId2subtreeCacheMode(v EnumbackendId2subtreeCacheModeProp)`

SetId2subtreeCacheMode sets Id2subtreeCacheMode field to given value.

### HasId2subtreeCacheMode

`func (o *GetBackend200Response) HasId2subtreeCacheMode() bool`

HasId2subtreeCacheMode returns a boolean if a field has been set.

### GetDn2uriCacheMode

`func (o *GetBackend200Response) GetDn2uriCacheMode() EnumbackendDn2uriCacheModeProp`

GetDn2uriCacheMode returns the Dn2uriCacheMode field if non-nil, zero value otherwise.

### GetDn2uriCacheModeOk

`func (o *GetBackend200Response) GetDn2uriCacheModeOk() (*EnumbackendDn2uriCacheModeProp, bool)`

GetDn2uriCacheModeOk returns a tuple with the Dn2uriCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2uriCacheMode

`func (o *GetBackend200Response) SetDn2uriCacheMode(v EnumbackendDn2uriCacheModeProp)`

SetDn2uriCacheMode sets Dn2uriCacheMode field to given value.

### HasDn2uriCacheMode

`func (o *GetBackend200Response) HasDn2uriCacheMode() bool`

HasDn2uriCacheMode returns a boolean if a field has been set.

### GetPrimeMethod

`func (o *GetBackend200Response) GetPrimeMethod() []EnumbackendPrimeMethodProp`

GetPrimeMethod returns the PrimeMethod field if non-nil, zero value otherwise.

### GetPrimeMethodOk

`func (o *GetBackend200Response) GetPrimeMethodOk() (*[]EnumbackendPrimeMethodProp, bool)`

GetPrimeMethodOk returns a tuple with the PrimeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeMethod

`func (o *GetBackend200Response) SetPrimeMethod(v []EnumbackendPrimeMethodProp)`

SetPrimeMethod sets PrimeMethod field to given value.

### HasPrimeMethod

`func (o *GetBackend200Response) HasPrimeMethod() bool`

HasPrimeMethod returns a boolean if a field has been set.

### GetPrimeThreadCount

`func (o *GetBackend200Response) GetPrimeThreadCount() int32`

GetPrimeThreadCount returns the PrimeThreadCount field if non-nil, zero value otherwise.

### GetPrimeThreadCountOk

`func (o *GetBackend200Response) GetPrimeThreadCountOk() (*int32, bool)`

GetPrimeThreadCountOk returns a tuple with the PrimeThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeThreadCount

`func (o *GetBackend200Response) SetPrimeThreadCount(v int32)`

SetPrimeThreadCount sets PrimeThreadCount field to given value.

### HasPrimeThreadCount

`func (o *GetBackend200Response) HasPrimeThreadCount() bool`

HasPrimeThreadCount returns a boolean if a field has been set.

### GetPrimeTimeLimit

`func (o *GetBackend200Response) GetPrimeTimeLimit() string`

GetPrimeTimeLimit returns the PrimeTimeLimit field if non-nil, zero value otherwise.

### GetPrimeTimeLimitOk

`func (o *GetBackend200Response) GetPrimeTimeLimitOk() (*string, bool)`

GetPrimeTimeLimitOk returns a tuple with the PrimeTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeTimeLimit

`func (o *GetBackend200Response) SetPrimeTimeLimit(v string)`

SetPrimeTimeLimit sets PrimeTimeLimit field to given value.

### HasPrimeTimeLimit

`func (o *GetBackend200Response) HasPrimeTimeLimit() bool`

HasPrimeTimeLimit returns a boolean if a field has been set.

### GetPrimeAllIndexes

`func (o *GetBackend200Response) GetPrimeAllIndexes() bool`

GetPrimeAllIndexes returns the PrimeAllIndexes field if non-nil, zero value otherwise.

### GetPrimeAllIndexesOk

`func (o *GetBackend200Response) GetPrimeAllIndexesOk() (*bool, bool)`

GetPrimeAllIndexesOk returns a tuple with the PrimeAllIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeAllIndexes

`func (o *GetBackend200Response) SetPrimeAllIndexes(v bool)`

SetPrimeAllIndexes sets PrimeAllIndexes field to given value.

### HasPrimeAllIndexes

`func (o *GetBackend200Response) HasPrimeAllIndexes() bool`

HasPrimeAllIndexes returns a boolean if a field has been set.

### GetSystemIndexToPrime

`func (o *GetBackend200Response) GetSystemIndexToPrime() []EnumbackendSystemIndexToPrimeProp`

GetSystemIndexToPrime returns the SystemIndexToPrime field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeOk

`func (o *GetBackend200Response) GetSystemIndexToPrimeOk() (*[]EnumbackendSystemIndexToPrimeProp, bool)`

GetSystemIndexToPrimeOk returns a tuple with the SystemIndexToPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrime

`func (o *GetBackend200Response) SetSystemIndexToPrime(v []EnumbackendSystemIndexToPrimeProp)`

SetSystemIndexToPrime sets SystemIndexToPrime field to given value.

### HasSystemIndexToPrime

`func (o *GetBackend200Response) HasSystemIndexToPrime() bool`

HasSystemIndexToPrime returns a boolean if a field has been set.

### GetSystemIndexToPrimeInternalNodesOnly

`func (o *GetBackend200Response) GetSystemIndexToPrimeInternalNodesOnly() []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp`

GetSystemIndexToPrimeInternalNodesOnly returns the SystemIndexToPrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeInternalNodesOnlyOk

`func (o *GetBackend200Response) GetSystemIndexToPrimeInternalNodesOnlyOk() (*[]EnumbackendSystemIndexToPrimeInternalNodesOnlyProp, bool)`

GetSystemIndexToPrimeInternalNodesOnlyOk returns a tuple with the SystemIndexToPrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrimeInternalNodesOnly

`func (o *GetBackend200Response) SetSystemIndexToPrimeInternalNodesOnly(v []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp)`

SetSystemIndexToPrimeInternalNodesOnly sets SystemIndexToPrimeInternalNodesOnly field to given value.

### HasSystemIndexToPrimeInternalNodesOnly

`func (o *GetBackend200Response) HasSystemIndexToPrimeInternalNodesOnly() bool`

HasSystemIndexToPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetBackgroundPrime

`func (o *GetBackend200Response) GetBackgroundPrime() bool`

GetBackgroundPrime returns the BackgroundPrime field if non-nil, zero value otherwise.

### GetBackgroundPrimeOk

`func (o *GetBackend200Response) GetBackgroundPrimeOk() (*bool, bool)`

GetBackgroundPrimeOk returns a tuple with the BackgroundPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPrime

`func (o *GetBackend200Response) SetBackgroundPrime(v bool)`

SetBackgroundPrime sets BackgroundPrime field to given value.

### HasBackgroundPrime

`func (o *GetBackend200Response) HasBackgroundPrime() bool`

HasBackgroundPrime returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *GetBackend200Response) GetIndexEntryLimit() int32`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *GetBackend200Response) GetIndexEntryLimitOk() (*int32, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *GetBackend200Response) SetIndexEntryLimit(v int32)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *GetBackend200Response) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetCompositeIndexEntryLimit

`func (o *GetBackend200Response) GetCompositeIndexEntryLimit() int32`

GetCompositeIndexEntryLimit returns the CompositeIndexEntryLimit field if non-nil, zero value otherwise.

### GetCompositeIndexEntryLimitOk

`func (o *GetBackend200Response) GetCompositeIndexEntryLimitOk() (*int32, bool)`

GetCompositeIndexEntryLimitOk returns a tuple with the CompositeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeIndexEntryLimit

`func (o *GetBackend200Response) SetCompositeIndexEntryLimit(v int32)`

SetCompositeIndexEntryLimit sets CompositeIndexEntryLimit field to given value.

### HasCompositeIndexEntryLimit

`func (o *GetBackend200Response) HasCompositeIndexEntryLimit() bool`

HasCompositeIndexEntryLimit returns a boolean if a field has been set.

### GetId2childrenIndexEntryLimit

`func (o *GetBackend200Response) GetId2childrenIndexEntryLimit() int32`

GetId2childrenIndexEntryLimit returns the Id2childrenIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2childrenIndexEntryLimitOk

`func (o *GetBackend200Response) GetId2childrenIndexEntryLimitOk() (*int32, bool)`

GetId2childrenIndexEntryLimitOk returns a tuple with the Id2childrenIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenIndexEntryLimit

`func (o *GetBackend200Response) SetId2childrenIndexEntryLimit(v int32)`

SetId2childrenIndexEntryLimit sets Id2childrenIndexEntryLimit field to given value.

### HasId2childrenIndexEntryLimit

`func (o *GetBackend200Response) HasId2childrenIndexEntryLimit() bool`

HasId2childrenIndexEntryLimit returns a boolean if a field has been set.

### GetId2subtreeIndexEntryLimit

`func (o *GetBackend200Response) GetId2subtreeIndexEntryLimit() int32`

GetId2subtreeIndexEntryLimit returns the Id2subtreeIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2subtreeIndexEntryLimitOk

`func (o *GetBackend200Response) GetId2subtreeIndexEntryLimitOk() (*int32, bool)`

GetId2subtreeIndexEntryLimitOk returns a tuple with the Id2subtreeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeIndexEntryLimit

`func (o *GetBackend200Response) SetId2subtreeIndexEntryLimit(v int32)`

SetId2subtreeIndexEntryLimit sets Id2subtreeIndexEntryLimit field to given value.

### HasId2subtreeIndexEntryLimit

`func (o *GetBackend200Response) HasId2subtreeIndexEntryLimit() bool`

HasId2subtreeIndexEntryLimit returns a boolean if a field has been set.

### GetImportTempDirectory

`func (o *GetBackend200Response) GetImportTempDirectory() string`

GetImportTempDirectory returns the ImportTempDirectory field if non-nil, zero value otherwise.

### GetImportTempDirectoryOk

`func (o *GetBackend200Response) GetImportTempDirectoryOk() (*string, bool)`

GetImportTempDirectoryOk returns a tuple with the ImportTempDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTempDirectory

`func (o *GetBackend200Response) SetImportTempDirectory(v string)`

SetImportTempDirectory sets ImportTempDirectory field to given value.


### GetImportThreadCount

`func (o *GetBackend200Response) GetImportThreadCount() int32`

GetImportThreadCount returns the ImportThreadCount field if non-nil, zero value otherwise.

### GetImportThreadCountOk

`func (o *GetBackend200Response) GetImportThreadCountOk() (*int32, bool)`

GetImportThreadCountOk returns a tuple with the ImportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportThreadCount

`func (o *GetBackend200Response) SetImportThreadCount(v int32)`

SetImportThreadCount sets ImportThreadCount field to given value.

### HasImportThreadCount

`func (o *GetBackend200Response) HasImportThreadCount() bool`

HasImportThreadCount returns a boolean if a field has been set.

### GetExportThreadCount

`func (o *GetBackend200Response) GetExportThreadCount() int32`

GetExportThreadCount returns the ExportThreadCount field if non-nil, zero value otherwise.

### GetExportThreadCountOk

`func (o *GetBackend200Response) GetExportThreadCountOk() (*int32, bool)`

GetExportThreadCountOk returns a tuple with the ExportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportThreadCount

`func (o *GetBackend200Response) SetExportThreadCount(v int32)`

SetExportThreadCount sets ExportThreadCount field to given value.

### HasExportThreadCount

`func (o *GetBackend200Response) HasExportThreadCount() bool`

HasExportThreadCount returns a boolean if a field has been set.

### GetDbImportCachePercent

`func (o *GetBackend200Response) GetDbImportCachePercent() int32`

GetDbImportCachePercent returns the DbImportCachePercent field if non-nil, zero value otherwise.

### GetDbImportCachePercentOk

`func (o *GetBackend200Response) GetDbImportCachePercentOk() (*int32, bool)`

GetDbImportCachePercentOk returns a tuple with the DbImportCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbImportCachePercent

`func (o *GetBackend200Response) SetDbImportCachePercent(v int32)`

SetDbImportCachePercent sets DbImportCachePercent field to given value.

### HasDbImportCachePercent

`func (o *GetBackend200Response) HasDbImportCachePercent() bool`

HasDbImportCachePercent returns a boolean if a field has been set.

### GetDbTxnWriteNoSync

`func (o *GetBackend200Response) GetDbTxnWriteNoSync() bool`

GetDbTxnWriteNoSync returns the DbTxnWriteNoSync field if non-nil, zero value otherwise.

### GetDbTxnWriteNoSyncOk

`func (o *GetBackend200Response) GetDbTxnWriteNoSyncOk() (*bool, bool)`

GetDbTxnWriteNoSyncOk returns a tuple with the DbTxnWriteNoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTxnWriteNoSync

`func (o *GetBackend200Response) SetDbTxnWriteNoSync(v bool)`

SetDbTxnWriteNoSync sets DbTxnWriteNoSync field to given value.

### HasDbTxnWriteNoSync

`func (o *GetBackend200Response) HasDbTxnWriteNoSync() bool`

HasDbTxnWriteNoSync returns a boolean if a field has been set.

### GetDeadlockRetryLimit

`func (o *GetBackend200Response) GetDeadlockRetryLimit() int32`

GetDeadlockRetryLimit returns the DeadlockRetryLimit field if non-nil, zero value otherwise.

### GetDeadlockRetryLimitOk

`func (o *GetBackend200Response) GetDeadlockRetryLimitOk() (*int32, bool)`

GetDeadlockRetryLimitOk returns a tuple with the DeadlockRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockRetryLimit

`func (o *GetBackend200Response) SetDeadlockRetryLimit(v int32)`

SetDeadlockRetryLimit sets DeadlockRetryLimit field to given value.

### HasDeadlockRetryLimit

`func (o *GetBackend200Response) HasDeadlockRetryLimit() bool`

HasDeadlockRetryLimit returns a boolean if a field has been set.

### GetExternalTxnDefaultBackendLockBehavior

`func (o *GetBackend200Response) GetExternalTxnDefaultBackendLockBehavior() EnumbackendExternalTxnDefaultBackendLockBehaviorProp`

GetExternalTxnDefaultBackendLockBehavior returns the ExternalTxnDefaultBackendLockBehavior field if non-nil, zero value otherwise.

### GetExternalTxnDefaultBackendLockBehaviorOk

`func (o *GetBackend200Response) GetExternalTxnDefaultBackendLockBehaviorOk() (*EnumbackendExternalTxnDefaultBackendLockBehaviorProp, bool)`

GetExternalTxnDefaultBackendLockBehaviorOk returns a tuple with the ExternalTxnDefaultBackendLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTxnDefaultBackendLockBehavior

`func (o *GetBackend200Response) SetExternalTxnDefaultBackendLockBehavior(v EnumbackendExternalTxnDefaultBackendLockBehaviorProp)`

SetExternalTxnDefaultBackendLockBehavior sets ExternalTxnDefaultBackendLockBehavior field to given value.

### HasExternalTxnDefaultBackendLockBehavior

`func (o *GetBackend200Response) HasExternalTxnDefaultBackendLockBehavior() bool`

HasExternalTxnDefaultBackendLockBehavior returns a boolean if a field has been set.

### GetSingleWriterLockBehavior

`func (o *GetBackend200Response) GetSingleWriterLockBehavior() EnumbackendSingleWriterLockBehaviorProp`

GetSingleWriterLockBehavior returns the SingleWriterLockBehavior field if non-nil, zero value otherwise.

### GetSingleWriterLockBehaviorOk

`func (o *GetBackend200Response) GetSingleWriterLockBehaviorOk() (*EnumbackendSingleWriterLockBehaviorProp, bool)`

GetSingleWriterLockBehaviorOk returns a tuple with the SingleWriterLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleWriterLockBehavior

`func (o *GetBackend200Response) SetSingleWriterLockBehavior(v EnumbackendSingleWriterLockBehaviorProp)`

SetSingleWriterLockBehavior sets SingleWriterLockBehavior field to given value.

### HasSingleWriterLockBehavior

`func (o *GetBackend200Response) HasSingleWriterLockBehavior() bool`

HasSingleWriterLockBehavior returns a boolean if a field has been set.

### GetSubtreeDeleteSizeLimit

`func (o *GetBackend200Response) GetSubtreeDeleteSizeLimit() int32`

GetSubtreeDeleteSizeLimit returns the SubtreeDeleteSizeLimit field if non-nil, zero value otherwise.

### GetSubtreeDeleteSizeLimitOk

`func (o *GetBackend200Response) GetSubtreeDeleteSizeLimitOk() (*int32, bool)`

GetSubtreeDeleteSizeLimitOk returns a tuple with the SubtreeDeleteSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeDeleteSizeLimit

`func (o *GetBackend200Response) SetSubtreeDeleteSizeLimit(v int32)`

SetSubtreeDeleteSizeLimit sets SubtreeDeleteSizeLimit field to given value.

### HasSubtreeDeleteSizeLimit

`func (o *GetBackend200Response) HasSubtreeDeleteSizeLimit() bool`

HasSubtreeDeleteSizeLimit returns a boolean if a field has been set.

### GetNumRecentChanges

`func (o *GetBackend200Response) GetNumRecentChanges() int32`

GetNumRecentChanges returns the NumRecentChanges field if non-nil, zero value otherwise.

### GetNumRecentChangesOk

`func (o *GetBackend200Response) GetNumRecentChangesOk() (*int32, bool)`

GetNumRecentChangesOk returns a tuple with the NumRecentChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecentChanges

`func (o *GetBackend200Response) SetNumRecentChanges(v int32)`

SetNumRecentChanges sets NumRecentChanges field to given value.

### HasNumRecentChanges

`func (o *GetBackend200Response) HasNumRecentChanges() bool`

HasNumRecentChanges returns a boolean if a field has been set.

### GetOfflineProcessDatabaseOpenTimeout

`func (o *GetBackend200Response) GetOfflineProcessDatabaseOpenTimeout() string`

GetOfflineProcessDatabaseOpenTimeout returns the OfflineProcessDatabaseOpenTimeout field if non-nil, zero value otherwise.

### GetOfflineProcessDatabaseOpenTimeoutOk

`func (o *GetBackend200Response) GetOfflineProcessDatabaseOpenTimeoutOk() (*string, bool)`

GetOfflineProcessDatabaseOpenTimeoutOk returns a tuple with the OfflineProcessDatabaseOpenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineProcessDatabaseOpenTimeout

`func (o *GetBackend200Response) SetOfflineProcessDatabaseOpenTimeout(v string)`

SetOfflineProcessDatabaseOpenTimeout sets OfflineProcessDatabaseOpenTimeout field to given value.

### HasOfflineProcessDatabaseOpenTimeout

`func (o *GetBackend200Response) HasOfflineProcessDatabaseOpenTimeout() bool`

HasOfflineProcessDatabaseOpenTimeout returns a boolean if a field has been set.

### GetInsignificantConfigArchiveAttribute

`func (o *GetBackend200Response) GetInsignificantConfigArchiveAttribute() []string`

GetInsignificantConfigArchiveAttribute returns the InsignificantConfigArchiveAttribute field if non-nil, zero value otherwise.

### GetInsignificantConfigArchiveAttributeOk

`func (o *GetBackend200Response) GetInsignificantConfigArchiveAttributeOk() (*[]string, bool)`

GetInsignificantConfigArchiveAttributeOk returns a tuple with the InsignificantConfigArchiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsignificantConfigArchiveAttribute

`func (o *GetBackend200Response) SetInsignificantConfigArchiveAttribute(v []string)`

SetInsignificantConfigArchiveAttribute sets InsignificantConfigArchiveAttribute field to given value.

### HasInsignificantConfigArchiveAttribute

`func (o *GetBackend200Response) HasInsignificantConfigArchiveAttribute() bool`

HasInsignificantConfigArchiveAttribute returns a boolean if a field has been set.

### GetMirroredSubtreePeerPollingInterval

`func (o *GetBackend200Response) GetMirroredSubtreePeerPollingInterval() string`

GetMirroredSubtreePeerPollingInterval returns the MirroredSubtreePeerPollingInterval field if non-nil, zero value otherwise.

### GetMirroredSubtreePeerPollingIntervalOk

`func (o *GetBackend200Response) GetMirroredSubtreePeerPollingIntervalOk() (*string, bool)`

GetMirroredSubtreePeerPollingIntervalOk returns a tuple with the MirroredSubtreePeerPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreePeerPollingInterval

`func (o *GetBackend200Response) SetMirroredSubtreePeerPollingInterval(v string)`

SetMirroredSubtreePeerPollingInterval sets MirroredSubtreePeerPollingInterval field to given value.

### HasMirroredSubtreePeerPollingInterval

`func (o *GetBackend200Response) HasMirroredSubtreePeerPollingInterval() bool`

HasMirroredSubtreePeerPollingInterval returns a boolean if a field has been set.

### GetMirroredSubtreeEntryUpdateTimeout

`func (o *GetBackend200Response) GetMirroredSubtreeEntryUpdateTimeout() string`

GetMirroredSubtreeEntryUpdateTimeout returns the MirroredSubtreeEntryUpdateTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeEntryUpdateTimeoutOk

`func (o *GetBackend200Response) GetMirroredSubtreeEntryUpdateTimeoutOk() (*string, bool)`

GetMirroredSubtreeEntryUpdateTimeoutOk returns a tuple with the MirroredSubtreeEntryUpdateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeEntryUpdateTimeout

`func (o *GetBackend200Response) SetMirroredSubtreeEntryUpdateTimeout(v string)`

SetMirroredSubtreeEntryUpdateTimeout sets MirroredSubtreeEntryUpdateTimeout field to given value.

### HasMirroredSubtreeEntryUpdateTimeout

`func (o *GetBackend200Response) HasMirroredSubtreeEntryUpdateTimeout() bool`

HasMirroredSubtreeEntryUpdateTimeout returns a boolean if a field has been set.

### GetMirroredSubtreeSearchTimeout

`func (o *GetBackend200Response) GetMirroredSubtreeSearchTimeout() string`

GetMirroredSubtreeSearchTimeout returns the MirroredSubtreeSearchTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeSearchTimeoutOk

`func (o *GetBackend200Response) GetMirroredSubtreeSearchTimeoutOk() (*string, bool)`

GetMirroredSubtreeSearchTimeoutOk returns a tuple with the MirroredSubtreeSearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeSearchTimeout

`func (o *GetBackend200Response) SetMirroredSubtreeSearchTimeout(v string)`

SetMirroredSubtreeSearchTimeout sets MirroredSubtreeSearchTimeout field to given value.

### HasMirroredSubtreeSearchTimeout

`func (o *GetBackend200Response) HasMirroredSubtreeSearchTimeout() bool`

HasMirroredSubtreeSearchTimeout returns a boolean if a field has been set.

### GetTaskBackingFile

`func (o *GetBackend200Response) GetTaskBackingFile() string`

GetTaskBackingFile returns the TaskBackingFile field if non-nil, zero value otherwise.

### GetTaskBackingFileOk

`func (o *GetBackend200Response) GetTaskBackingFileOk() (*string, bool)`

GetTaskBackingFileOk returns a tuple with the TaskBackingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskBackingFile

`func (o *GetBackend200Response) SetTaskBackingFile(v string)`

SetTaskBackingFile sets TaskBackingFile field to given value.


### GetMaximumInitialTaskLogMessagesToRetain

`func (o *GetBackend200Response) GetMaximumInitialTaskLogMessagesToRetain() int32`

GetMaximumInitialTaskLogMessagesToRetain returns the MaximumInitialTaskLogMessagesToRetain field if non-nil, zero value otherwise.

### GetMaximumInitialTaskLogMessagesToRetainOk

`func (o *GetBackend200Response) GetMaximumInitialTaskLogMessagesToRetainOk() (*int32, bool)`

GetMaximumInitialTaskLogMessagesToRetainOk returns a tuple with the MaximumInitialTaskLogMessagesToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInitialTaskLogMessagesToRetain

`func (o *GetBackend200Response) SetMaximumInitialTaskLogMessagesToRetain(v int32)`

SetMaximumInitialTaskLogMessagesToRetain sets MaximumInitialTaskLogMessagesToRetain field to given value.

### HasMaximumInitialTaskLogMessagesToRetain

`func (o *GetBackend200Response) HasMaximumInitialTaskLogMessagesToRetain() bool`

HasMaximumInitialTaskLogMessagesToRetain returns a boolean if a field has been set.

### GetMaximumFinalTaskLogMessagesToRetain

`func (o *GetBackend200Response) GetMaximumFinalTaskLogMessagesToRetain() int32`

GetMaximumFinalTaskLogMessagesToRetain returns the MaximumFinalTaskLogMessagesToRetain field if non-nil, zero value otherwise.

### GetMaximumFinalTaskLogMessagesToRetainOk

`func (o *GetBackend200Response) GetMaximumFinalTaskLogMessagesToRetainOk() (*int32, bool)`

GetMaximumFinalTaskLogMessagesToRetainOk returns a tuple with the MaximumFinalTaskLogMessagesToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFinalTaskLogMessagesToRetain

`func (o *GetBackend200Response) SetMaximumFinalTaskLogMessagesToRetain(v int32)`

SetMaximumFinalTaskLogMessagesToRetain sets MaximumFinalTaskLogMessagesToRetain field to given value.

### HasMaximumFinalTaskLogMessagesToRetain

`func (o *GetBackend200Response) HasMaximumFinalTaskLogMessagesToRetain() bool`

HasMaximumFinalTaskLogMessagesToRetain returns a boolean if a field has been set.

### GetTaskRetentionTime

`func (o *GetBackend200Response) GetTaskRetentionTime() string`

GetTaskRetentionTime returns the TaskRetentionTime field if non-nil, zero value otherwise.

### GetTaskRetentionTimeOk

`func (o *GetBackend200Response) GetTaskRetentionTimeOk() (*string, bool)`

GetTaskRetentionTimeOk returns a tuple with the TaskRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRetentionTime

`func (o *GetBackend200Response) SetTaskRetentionTime(v string)`

SetTaskRetentionTime sets TaskRetentionTime field to given value.

### HasTaskRetentionTime

`func (o *GetBackend200Response) HasTaskRetentionTime() bool`

HasTaskRetentionTime returns a boolean if a field has been set.

### GetNotificationSenderAddress

`func (o *GetBackend200Response) GetNotificationSenderAddress() string`

GetNotificationSenderAddress returns the NotificationSenderAddress field if non-nil, zero value otherwise.

### GetNotificationSenderAddressOk

`func (o *GetBackend200Response) GetNotificationSenderAddressOk() (*string, bool)`

GetNotificationSenderAddressOk returns a tuple with the NotificationSenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSenderAddress

`func (o *GetBackend200Response) SetNotificationSenderAddress(v string)`

SetNotificationSenderAddress sets NotificationSenderAddress field to given value.

### HasNotificationSenderAddress

`func (o *GetBackend200Response) HasNotificationSenderAddress() bool`

HasNotificationSenderAddress returns a boolean if a field has been set.

### GetAlertRetentionTime

`func (o *GetBackend200Response) GetAlertRetentionTime() string`

GetAlertRetentionTime returns the AlertRetentionTime field if non-nil, zero value otherwise.

### GetAlertRetentionTimeOk

`func (o *GetBackend200Response) GetAlertRetentionTimeOk() (*string, bool)`

GetAlertRetentionTimeOk returns a tuple with the AlertRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRetentionTime

`func (o *GetBackend200Response) SetAlertRetentionTime(v string)`

SetAlertRetentionTime sets AlertRetentionTime field to given value.


### GetMaxAlerts

`func (o *GetBackend200Response) GetMaxAlerts() int32`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *GetBackend200Response) GetMaxAlertsOk() (*int32, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *GetBackend200Response) SetMaxAlerts(v int32)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *GetBackend200Response) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *GetBackend200Response) GetDisabledAlertType() []EnumbackendDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *GetBackend200Response) GetDisabledAlertTypeOk() (*[]EnumbackendDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *GetBackend200Response) SetDisabledAlertType(v []EnumbackendDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *GetBackend200Response) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.

### GetAlarmRetentionTime

`func (o *GetBackend200Response) GetAlarmRetentionTime() string`

GetAlarmRetentionTime returns the AlarmRetentionTime field if non-nil, zero value otherwise.

### GetAlarmRetentionTimeOk

`func (o *GetBackend200Response) GetAlarmRetentionTimeOk() (*string, bool)`

GetAlarmRetentionTimeOk returns a tuple with the AlarmRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRetentionTime

`func (o *GetBackend200Response) SetAlarmRetentionTime(v string)`

SetAlarmRetentionTime sets AlarmRetentionTime field to given value.


### GetMaxAlarms

`func (o *GetBackend200Response) GetMaxAlarms() int32`

GetMaxAlarms returns the MaxAlarms field if non-nil, zero value otherwise.

### GetMaxAlarmsOk

`func (o *GetBackend200Response) GetMaxAlarmsOk() (*int32, bool)`

GetMaxAlarmsOk returns a tuple with the MaxAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlarms

`func (o *GetBackend200Response) SetMaxAlarms(v int32)`

SetMaxAlarms sets MaxAlarms field to given value.

### HasMaxAlarms

`func (o *GetBackend200Response) HasMaxAlarms() bool`

HasMaxAlarms returns a boolean if a field has been set.

### GetStorageDir

`func (o *GetBackend200Response) GetStorageDir() string`

GetStorageDir returns the StorageDir field if non-nil, zero value otherwise.

### GetStorageDirOk

`func (o *GetBackend200Response) GetStorageDirOk() (*string, bool)`

GetStorageDirOk returns a tuple with the StorageDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDir

`func (o *GetBackend200Response) SetStorageDir(v string)`

SetStorageDir sets StorageDir field to given value.


### GetMetricsDir

`func (o *GetBackend200Response) GetMetricsDir() string`

GetMetricsDir returns the MetricsDir field if non-nil, zero value otherwise.

### GetMetricsDirOk

`func (o *GetBackend200Response) GetMetricsDirOk() (*string, bool)`

GetMetricsDirOk returns a tuple with the MetricsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsDir

`func (o *GetBackend200Response) SetMetricsDir(v string)`

SetMetricsDir sets MetricsDir field to given value.


### GetSampleFlushInterval

`func (o *GetBackend200Response) GetSampleFlushInterval() string`

GetSampleFlushInterval returns the SampleFlushInterval field if non-nil, zero value otherwise.

### GetSampleFlushIntervalOk

`func (o *GetBackend200Response) GetSampleFlushIntervalOk() (*string, bool)`

GetSampleFlushIntervalOk returns a tuple with the SampleFlushInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFlushInterval

`func (o *GetBackend200Response) SetSampleFlushInterval(v string)`

SetSampleFlushInterval sets SampleFlushInterval field to given value.

### HasSampleFlushInterval

`func (o *GetBackend200Response) HasSampleFlushInterval() bool`

HasSampleFlushInterval returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *GetBackend200Response) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *GetBackend200Response) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *GetBackend200Response) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


