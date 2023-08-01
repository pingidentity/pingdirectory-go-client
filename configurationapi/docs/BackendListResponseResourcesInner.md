# BackendListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcannedResponseBackendSchemaUrn**](EnumcannedResponseBackendSchemaUrn.md) |  | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
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
**DbCachePercent** | Pointer to **int64** | Specifies the percentage of JVM memory to allocate to the database cache. | [optional] 
**JeProperty** | Pointer to **[]string** | Specifies the database and environment properties for the Berkeley DB Java Edition database serving the data for this backend. | [optional] 
**ChangelogWriteBatchSize** | Pointer to **int64** | Specifies the number of changelog entries written in a single database transaction. | [optional] 
**ChangelogPurgeBatchSize** | Pointer to **int64** | Specifies the number of changelog entries purged in a single database transaction. | [optional] 
**ChangelogWriteQueueCapacity** | Pointer to **int64** | Specifies the capacity of the changelog write queue in number of changes. | [optional] 
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
**ChangelogMaxBeforeAfterValues** | Pointer to **int64** | This controls whether all attribute values for a modified attribute (even those values that have not changed) will be included in the changelog entry. If the number of attribute values does not exceed this limit, then all values for the modified attribute will be included in the changelog entry. | [optional] 
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
**DbNumCleanerThreads** | Pointer to **int64** | Specifies the number of threads that the backend should maintain to keep the database log files at or near the desired utilization. A value of zero indicates that the number of cleaner threads should be automatically configured based on the number of available CPUs. | [optional] 
**DbCleanerMinUtilization** | Pointer to **int64** | Specifies the minimum percentage of \&quot;live\&quot; data that the database cleaner attempts to keep in database log files. | [optional] 
**DbEvictorCriticalPercentage** | Pointer to **int64** | Specifies the percentage over the configured maximum that the database cache is allowed to grow. It is recommended to set this value slightly above zero when the database is too large to fully cache in memory. In this case, a dedicated background evictor thread is used to perform evictions once the cache fills up reducing the possibility that server threads are blocked. | [optional] 
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
**PrimeThreadCount** | Pointer to **int64** | Specifies the number of threads to use when priming. At present, this applies only to the preload and cursor-across-indexes prime methods. | [optional] 
**PrimeTimeLimit** | Pointer to **string** | Specifies the maximum length of time that the backend prime should be allowed to run. A duration of zero seconds indicates that there should not be a time limit. | [optional] 
**PrimeAllIndexes** | Pointer to **bool** | Indicates whether to prime all indexes associated with this backend, or to only prime the specified set of indexes (as configured with the system-index-to-prime property for the system indexes, and the prime-index property in the attribute index definition for attribute indexes). | [optional] 
**SystemIndexToPrime** | Pointer to [**[]EnumbackendSystemIndexToPrimeProp**](EnumbackendSystemIndexToPrimeProp.md) |  | [optional] 
**SystemIndexToPrimeInternalNodesOnly** | Pointer to [**[]EnumbackendSystemIndexToPrimeInternalNodesOnlyProp**](EnumbackendSystemIndexToPrimeInternalNodesOnlyProp.md) |  | [optional] 
**BackgroundPrime** | Pointer to **bool** | Indicates whether to attempt to perform the prime using a background thread if possible. If background priming is enabled, then the Directory Server may be allowed to accept client connections and process requests while the prime is in progress. | [optional] 
**IndexEntryLimit** | Pointer to **int64** | Specifies the maximum number of entries that are allowed to match a given index key before that particular index key is no longer maintained. | [optional] 
**CompositeIndexEntryLimit** | Pointer to **int64** | Specifies the maximum number of entries that are allowed to match a given composite index key before that particular composite index key is no longer maintained. | [optional] 
**Id2childrenIndexEntryLimit** | Pointer to **int64** | Specifies the maximum number of entry IDs to maintain for each entry in the id2children system index (which keeps track of the immediate children for an entry, to assist in otherwise unindexed searches with a single-level scope). A value of 0 means there is no limit, however this could have a big impact on database size on disk and on server performance. | [optional] 
**Id2subtreeIndexEntryLimit** | Pointer to **int64** | Specifies the maximum number of entry IDs to maintain for each entry in the id2subtree system index (which keeps track of all descendants below an entry, to assist in otherwise unindexed searches with a whole-subtree or subordinate subtree scope). A value of 0 means there is no limit, however this could have a big impact on database size on disk and on server performance. | [optional] 
**ImportTempDirectory** | **string** | Specifies the location of the directory that is used to hold temporary information during the index post-processing phase of an LDIF import. | 
**ImportThreadCount** | Pointer to **int64** | Specifies the number of threads to use for concurrent processing during an LDIF import. | [optional] 
**ExportThreadCount** | Pointer to **int64** | Specifies the number of threads to use for concurrently retrieving and encoding entries during an LDIF export. | [optional] 
**DbImportCachePercent** | Pointer to **int64** | The percentage of JVM memory to allocate to the database cache during import operations. | [optional] 
**DbTxnWriteNoSync** | Pointer to **bool** | Indicates whether the database should synchronously flush data as it is written to disk. | [optional] 
**DeadlockRetryLimit** | Pointer to **int64** | Specifies the number of times that the server should retry an attempted operation in the backend if a deadlock results from two concurrent requests that interfere with each other in a conflicting manner. | [optional] 
**ExternalTxnDefaultBackendLockBehavior** | Pointer to [**EnumbackendExternalTxnDefaultBackendLockBehaviorProp**](EnumbackendExternalTxnDefaultBackendLockBehaviorProp.md) |  | [optional] 
**SingleWriterLockBehavior** | Pointer to [**EnumbackendSingleWriterLockBehaviorProp**](EnumbackendSingleWriterLockBehaviorProp.md) |  | [optional] 
**SubtreeDeleteSizeLimit** | Pointer to **int64** | Specifies the maximum number of entries that may be deleted from the backend when using the subtree delete control. | [optional] 
**NumRecentChanges** | Pointer to **int64** | Specifies the number of recent LDAP entry changes per replica for which the backend keeps a record to allow replication to recover in the event that the server is abruptly terminated. Increasing this value can lead to an increased peak server modification rate as well as increased replication throughput. | [optional] 
**OfflineProcessDatabaseOpenTimeout** | Pointer to **string** | Specifies a timeout duration which will be used for opening the database environment by an offline process, such as export-ldif. | [optional] 
**MirroredSubtreePeerPollingInterval** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait before polling the peer servers in the topology to determine if there are any changes in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. | [optional] 
**MirroredSubtreeEntryUpdateTimeout** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for an update operation (add, delete, modify and modify-dn) on an entry to be applied on all servers in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. | [optional] 
**MirroredSubtreeSearchTimeout** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for a search operation to complete. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. Search requests that take longer than this timeout will be canceled and considered failures. | [optional] 
**PeerServer** | Pointer to **[]string** | Specifies the set of peer servers onto which updates should be mirrored. The local server should not be included in this set, but if it is, then it will just be ignored. | [optional] 
**MirroredSubtreePreferredMasterType** | Pointer to [**[]EnumbackendMirroredSubtreePreferredMasterTypeProp**](EnumbackendMirroredSubtreePreferredMasterTypeProp.md) |  | [optional] 
**SimulatedResultCode** | Pointer to **int64** | Specifies the numeric value of the result code to be assumed for update operations (add, delete, modify and modify-dn) targeted to this backend. | [optional] 
**InsignificantConfigArchiveAttribute** | Pointer to **[]string** | The name or OID of an attribute type that is considered insignificant for the purpose of maintaining the configuration archive. | [optional] 
**InsignificantConfigArchiveBaseDN** | Pointer to **[]string** | The base DN that is considered insignificant for the purpose of maintaining the configuration archive. | [optional] 
**MaintainConfigArchive** | Pointer to **bool** | Indicates whether the server should maintain the config archive with new changes to the config backend. | [optional] 
**MaxConfigArchiveCount** | Pointer to **int64** | Indicates the maximum number of previous config files to keep as part of maintaining the config archive. | [optional] 
**TaskBackingFile** | **string** | Specifies the path to the backing file for storing information about the tasks configured in the server. | 
**MaximumInitialTaskLogMessagesToRetain** | Pointer to **int64** | The maximum number of log messages to retain in each task entry from the beginning of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks. | [optional] 
**MaximumFinalTaskLogMessagesToRetain** | Pointer to **int64** | The maximum number of log messages to retain in each task entry from the end of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks. | [optional] 
**TaskRetentionTime** | Pointer to **string** | Specifies the length of time that task entries should be retained after processing on the associated task has been completed. | [optional] 
**NotificationSenderAddress** | Pointer to **string** | Specifies the email address to use as the sender address (that is, the \&quot;From:\&quot; address) for notification mail messages generated when a task completes execution. | [optional] 
**AlertRetentionTime** | **string** | Specifies the maximum length of time that information about generated alerts should be maintained before they will be purged. | 
**MaxAlerts** | Pointer to **int64** | Specifies the maximum number of alerts that should be retained. If more alerts than this configured maximum are generated within the alert retention time, then the oldest alerts will be purged to achieve this maximum. | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumbackendDisabledAlertTypeProp**](EnumbackendDisabledAlertTypeProp.md) |  | [optional] 
**AlarmRetentionTime** | **string** | Specifies the maximum length of time that information about raised alarms should be maintained before they will be purged. | 
**MaxAlarms** | Pointer to **int64** | Specifies the maximum number of alarms that should be retained. If more alarms than this configured maximum are generated within the alarm retention time, then the oldest alarms will be purged to achieve this maximum. Only alarms at normal severity will be purged. | [optional] 
**StorageDir** | **string** | Specifies the path to the directory that will be used to store queued samples. | 
**MetricsDir** | **string** | Specifies the path to the directory that contains metric definitions. | 
**SampleFlushInterval** | Pointer to **string** | Period when samples are flushed to disk. | [optional] 
**RetentionPolicy** | **[]string** | The retention policy to use for the Metrics Backend . | 

## Methods

### NewBackendListResponseResourcesInner

`func NewBackendListResponseResourcesInner(schemas []EnumcannedResponseBackendSchemaUrn, id string, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, showAllAttributes bool, enabled bool, backupDirectory []string, ldifFile string, trustStoreFile string, dbDirectory string, changelogMaximumAge string, importTempDirectory string, taskBackingFile string, alertRetentionTime string, alarmRetentionTime string, storageDir string, metricsDir string, retentionPolicy []string, ) *BackendListResponseResourcesInner`

NewBackendListResponseResourcesInner instantiates a new BackendListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendListResponseResourcesInnerWithDefaults

`func NewBackendListResponseResourcesInnerWithDefaults() *BackendListResponseResourcesInner`

NewBackendListResponseResourcesInnerWithDefaults instantiates a new BackendListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BackendListResponseResourcesInner) GetSchemas() []EnumcannedResponseBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BackendListResponseResourcesInner) GetSchemasOk() (*[]EnumcannedResponseBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BackendListResponseResourcesInner) SetSchemas(v []EnumcannedResponseBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *BackendListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackendListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackendListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetBackendID

`func (o *BackendListResponseResourcesInner) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *BackendListResponseResourcesInner) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *BackendListResponseResourcesInner) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *BackendListResponseResourcesInner) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *BackendListResponseResourcesInner) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *BackendListResponseResourcesInner) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *BackendListResponseResourcesInner) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *BackendListResponseResourcesInner) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *BackendListResponseResourcesInner) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetSchemaEntryDN

`func (o *BackendListResponseResourcesInner) GetSchemaEntryDN() []string`

GetSchemaEntryDN returns the SchemaEntryDN field if non-nil, zero value otherwise.

### GetSchemaEntryDNOk

`func (o *BackendListResponseResourcesInner) GetSchemaEntryDNOk() (*[]string, bool)`

GetSchemaEntryDNOk returns a tuple with the SchemaEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaEntryDN

`func (o *BackendListResponseResourcesInner) SetSchemaEntryDN(v []string)`

SetSchemaEntryDN sets SchemaEntryDN field to given value.

### HasSchemaEntryDN

`func (o *BackendListResponseResourcesInner) HasSchemaEntryDN() bool`

HasSchemaEntryDN returns a boolean if a field has been set.

### GetShowAllAttributes

`func (o *BackendListResponseResourcesInner) GetShowAllAttributes() bool`

GetShowAllAttributes returns the ShowAllAttributes field if non-nil, zero value otherwise.

### GetShowAllAttributesOk

`func (o *BackendListResponseResourcesInner) GetShowAllAttributesOk() (*bool, bool)`

GetShowAllAttributesOk returns a tuple with the ShowAllAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllAttributes

`func (o *BackendListResponseResourcesInner) SetShowAllAttributes(v bool)`

SetShowAllAttributes sets ShowAllAttributes field to given value.


### GetReadOnlySchemaFile

`func (o *BackendListResponseResourcesInner) GetReadOnlySchemaFile() []string`

GetReadOnlySchemaFile returns the ReadOnlySchemaFile field if non-nil, zero value otherwise.

### GetReadOnlySchemaFileOk

`func (o *BackendListResponseResourcesInner) GetReadOnlySchemaFileOk() (*[]string, bool)`

GetReadOnlySchemaFileOk returns a tuple with the ReadOnlySchemaFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySchemaFile

`func (o *BackendListResponseResourcesInner) SetReadOnlySchemaFile(v []string)`

SetReadOnlySchemaFile sets ReadOnlySchemaFile field to given value.

### HasReadOnlySchemaFile

`func (o *BackendListResponseResourcesInner) HasReadOnlySchemaFile() bool`

HasReadOnlySchemaFile returns a boolean if a field has been set.

### GetDescription

`func (o *BackendListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackendListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackendListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackendListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BackendListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackendListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackendListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *BackendListResponseResourcesInner) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *BackendListResponseResourcesInner) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *BackendListResponseResourcesInner) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *BackendListResponseResourcesInner) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *BackendListResponseResourcesInner) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *BackendListResponseResourcesInner) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *BackendListResponseResourcesInner) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *BackendListResponseResourcesInner) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *BackendListResponseResourcesInner) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *BackendListResponseResourcesInner) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *BackendListResponseResourcesInner) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *BackendListResponseResourcesInner) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *BackendListResponseResourcesInner) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *BackendListResponseResourcesInner) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *BackendListResponseResourcesInner) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *BackendListResponseResourcesInner) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *BackendListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BackendListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BackendListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BackendListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *BackendListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *BackendListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *BackendListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *BackendListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetBackupDirectory

`func (o *BackendListResponseResourcesInner) GetBackupDirectory() []string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *BackendListResponseResourcesInner) GetBackupDirectoryOk() (*[]string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *BackendListResponseResourcesInner) SetBackupDirectory(v []string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIsPrivateBackend

`func (o *BackendListResponseResourcesInner) GetIsPrivateBackend() bool`

GetIsPrivateBackend returns the IsPrivateBackend field if non-nil, zero value otherwise.

### GetIsPrivateBackendOk

`func (o *BackendListResponseResourcesInner) GetIsPrivateBackendOk() (*bool, bool)`

GetIsPrivateBackendOk returns a tuple with the IsPrivateBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateBackend

`func (o *BackendListResponseResourcesInner) SetIsPrivateBackend(v bool)`

SetIsPrivateBackend sets IsPrivateBackend field to given value.

### HasIsPrivateBackend

`func (o *BackendListResponseResourcesInner) HasIsPrivateBackend() bool`

HasIsPrivateBackend returns a boolean if a field has been set.

### GetLdifFile

`func (o *BackendListResponseResourcesInner) GetLdifFile() string`

GetLdifFile returns the LdifFile field if non-nil, zero value otherwise.

### GetLdifFileOk

`func (o *BackendListResponseResourcesInner) GetLdifFileOk() (*string, bool)`

GetLdifFileOk returns a tuple with the LdifFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifFile

`func (o *BackendListResponseResourcesInner) SetLdifFile(v string)`

SetLdifFile sets LdifFile field to given value.


### GetTrustStoreFile

`func (o *BackendListResponseResourcesInner) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *BackendListResponseResourcesInner) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *BackendListResponseResourcesInner) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.


### GetTrustStoreType

`func (o *BackendListResponseResourcesInner) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *BackendListResponseResourcesInner) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *BackendListResponseResourcesInner) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *BackendListResponseResourcesInner) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *BackendListResponseResourcesInner) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *BackendListResponseResourcesInner) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *BackendListResponseResourcesInner) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *BackendListResponseResourcesInner) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStorePinFile

`func (o *BackendListResponseResourcesInner) GetTrustStorePinFile() string`

GetTrustStorePinFile returns the TrustStorePinFile field if non-nil, zero value otherwise.

### GetTrustStorePinFileOk

`func (o *BackendListResponseResourcesInner) GetTrustStorePinFileOk() (*string, bool)`

GetTrustStorePinFileOk returns a tuple with the TrustStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinFile

`func (o *BackendListResponseResourcesInner) SetTrustStorePinFile(v string)`

SetTrustStorePinFile sets TrustStorePinFile field to given value.

### HasTrustStorePinFile

`func (o *BackendListResponseResourcesInner) HasTrustStorePinFile() bool`

HasTrustStorePinFile returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *BackendListResponseResourcesInner) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *BackendListResponseResourcesInner) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *BackendListResponseResourcesInner) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *BackendListResponseResourcesInner) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetDbDirectory

`func (o *BackendListResponseResourcesInner) GetDbDirectory() string`

GetDbDirectory returns the DbDirectory field if non-nil, zero value otherwise.

### GetDbDirectoryOk

`func (o *BackendListResponseResourcesInner) GetDbDirectoryOk() (*string, bool)`

GetDbDirectoryOk returns a tuple with the DbDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectory

`func (o *BackendListResponseResourcesInner) SetDbDirectory(v string)`

SetDbDirectory sets DbDirectory field to given value.


### GetDbDirectoryPermissions

`func (o *BackendListResponseResourcesInner) GetDbDirectoryPermissions() string`

GetDbDirectoryPermissions returns the DbDirectoryPermissions field if non-nil, zero value otherwise.

### GetDbDirectoryPermissionsOk

`func (o *BackendListResponseResourcesInner) GetDbDirectoryPermissionsOk() (*string, bool)`

GetDbDirectoryPermissionsOk returns a tuple with the DbDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectoryPermissions

`func (o *BackendListResponseResourcesInner) SetDbDirectoryPermissions(v string)`

SetDbDirectoryPermissions sets DbDirectoryPermissions field to given value.

### HasDbDirectoryPermissions

`func (o *BackendListResponseResourcesInner) HasDbDirectoryPermissions() bool`

HasDbDirectoryPermissions returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *BackendListResponseResourcesInner) GetDbCachePercent() int64`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *BackendListResponseResourcesInner) GetDbCachePercentOk() (*int64, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *BackendListResponseResourcesInner) SetDbCachePercent(v int64)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *BackendListResponseResourcesInner) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetJeProperty

`func (o *BackendListResponseResourcesInner) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *BackendListResponseResourcesInner) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *BackendListResponseResourcesInner) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *BackendListResponseResourcesInner) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetChangelogWriteBatchSize

`func (o *BackendListResponseResourcesInner) GetChangelogWriteBatchSize() int64`

GetChangelogWriteBatchSize returns the ChangelogWriteBatchSize field if non-nil, zero value otherwise.

### GetChangelogWriteBatchSizeOk

`func (o *BackendListResponseResourcesInner) GetChangelogWriteBatchSizeOk() (*int64, bool)`

GetChangelogWriteBatchSizeOk returns a tuple with the ChangelogWriteBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogWriteBatchSize

`func (o *BackendListResponseResourcesInner) SetChangelogWriteBatchSize(v int64)`

SetChangelogWriteBatchSize sets ChangelogWriteBatchSize field to given value.

### HasChangelogWriteBatchSize

`func (o *BackendListResponseResourcesInner) HasChangelogWriteBatchSize() bool`

HasChangelogWriteBatchSize returns a boolean if a field has been set.

### GetChangelogPurgeBatchSize

`func (o *BackendListResponseResourcesInner) GetChangelogPurgeBatchSize() int64`

GetChangelogPurgeBatchSize returns the ChangelogPurgeBatchSize field if non-nil, zero value otherwise.

### GetChangelogPurgeBatchSizeOk

`func (o *BackendListResponseResourcesInner) GetChangelogPurgeBatchSizeOk() (*int64, bool)`

GetChangelogPurgeBatchSizeOk returns a tuple with the ChangelogPurgeBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPurgeBatchSize

`func (o *BackendListResponseResourcesInner) SetChangelogPurgeBatchSize(v int64)`

SetChangelogPurgeBatchSize sets ChangelogPurgeBatchSize field to given value.

### HasChangelogPurgeBatchSize

`func (o *BackendListResponseResourcesInner) HasChangelogPurgeBatchSize() bool`

HasChangelogPurgeBatchSize returns a boolean if a field has been set.

### GetChangelogWriteQueueCapacity

`func (o *BackendListResponseResourcesInner) GetChangelogWriteQueueCapacity() int64`

GetChangelogWriteQueueCapacity returns the ChangelogWriteQueueCapacity field if non-nil, zero value otherwise.

### GetChangelogWriteQueueCapacityOk

`func (o *BackendListResponseResourcesInner) GetChangelogWriteQueueCapacityOk() (*int64, bool)`

GetChangelogWriteQueueCapacityOk returns a tuple with the ChangelogWriteQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogWriteQueueCapacity

`func (o *BackendListResponseResourcesInner) SetChangelogWriteQueueCapacity(v int64)`

SetChangelogWriteQueueCapacity sets ChangelogWriteQueueCapacity field to given value.

### HasChangelogWriteQueueCapacity

`func (o *BackendListResponseResourcesInner) HasChangelogWriteQueueCapacity() bool`

HasChangelogWriteQueueCapacity returns a boolean if a field has been set.

### GetIndexIncludeAttribute

`func (o *BackendListResponseResourcesInner) GetIndexIncludeAttribute() []string`

GetIndexIncludeAttribute returns the IndexIncludeAttribute field if non-nil, zero value otherwise.

### GetIndexIncludeAttributeOk

`func (o *BackendListResponseResourcesInner) GetIndexIncludeAttributeOk() (*[]string, bool)`

GetIndexIncludeAttributeOk returns a tuple with the IndexIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIncludeAttribute

`func (o *BackendListResponseResourcesInner) SetIndexIncludeAttribute(v []string)`

SetIndexIncludeAttribute sets IndexIncludeAttribute field to given value.

### HasIndexIncludeAttribute

`func (o *BackendListResponseResourcesInner) HasIndexIncludeAttribute() bool`

HasIndexIncludeAttribute returns a boolean if a field has been set.

### GetIndexExcludeAttribute

`func (o *BackendListResponseResourcesInner) GetIndexExcludeAttribute() []string`

GetIndexExcludeAttribute returns the IndexExcludeAttribute field if non-nil, zero value otherwise.

### GetIndexExcludeAttributeOk

`func (o *BackendListResponseResourcesInner) GetIndexExcludeAttributeOk() (*[]string, bool)`

GetIndexExcludeAttributeOk returns a tuple with the IndexExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexExcludeAttribute

`func (o *BackendListResponseResourcesInner) SetIndexExcludeAttribute(v []string)`

SetIndexExcludeAttribute sets IndexExcludeAttribute field to given value.

### HasIndexExcludeAttribute

`func (o *BackendListResponseResourcesInner) HasIndexExcludeAttribute() bool`

HasIndexExcludeAttribute returns a boolean if a field has been set.

### GetChangelogMaximumAge

`func (o *BackendListResponseResourcesInner) GetChangelogMaximumAge() string`

GetChangelogMaximumAge returns the ChangelogMaximumAge field if non-nil, zero value otherwise.

### GetChangelogMaximumAgeOk

`func (o *BackendListResponseResourcesInner) GetChangelogMaximumAgeOk() (*string, bool)`

GetChangelogMaximumAgeOk returns a tuple with the ChangelogMaximumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogMaximumAge

`func (o *BackendListResponseResourcesInner) SetChangelogMaximumAge(v string)`

SetChangelogMaximumAge sets ChangelogMaximumAge field to given value.


### GetTargetDatabaseSize

`func (o *BackendListResponseResourcesInner) GetTargetDatabaseSize() string`

GetTargetDatabaseSize returns the TargetDatabaseSize field if non-nil, zero value otherwise.

### GetTargetDatabaseSizeOk

`func (o *BackendListResponseResourcesInner) GetTargetDatabaseSizeOk() (*string, bool)`

GetTargetDatabaseSizeOk returns a tuple with the TargetDatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatabaseSize

`func (o *BackendListResponseResourcesInner) SetTargetDatabaseSize(v string)`

SetTargetDatabaseSize sets TargetDatabaseSize field to given value.

### HasTargetDatabaseSize

`func (o *BackendListResponseResourcesInner) HasTargetDatabaseSize() bool`

HasTargetDatabaseSize returns a boolean if a field has been set.

### GetChangelogEntryIncludeBaseDN

`func (o *BackendListResponseResourcesInner) GetChangelogEntryIncludeBaseDN() []string`

GetChangelogEntryIncludeBaseDN returns the ChangelogEntryIncludeBaseDN field if non-nil, zero value otherwise.

### GetChangelogEntryIncludeBaseDNOk

`func (o *BackendListResponseResourcesInner) GetChangelogEntryIncludeBaseDNOk() (*[]string, bool)`

GetChangelogEntryIncludeBaseDNOk returns a tuple with the ChangelogEntryIncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryIncludeBaseDN

`func (o *BackendListResponseResourcesInner) SetChangelogEntryIncludeBaseDN(v []string)`

SetChangelogEntryIncludeBaseDN sets ChangelogEntryIncludeBaseDN field to given value.

### HasChangelogEntryIncludeBaseDN

`func (o *BackendListResponseResourcesInner) HasChangelogEntryIncludeBaseDN() bool`

HasChangelogEntryIncludeBaseDN returns a boolean if a field has been set.

### GetChangelogEntryExcludeBaseDN

`func (o *BackendListResponseResourcesInner) GetChangelogEntryExcludeBaseDN() []string`

GetChangelogEntryExcludeBaseDN returns the ChangelogEntryExcludeBaseDN field if non-nil, zero value otherwise.

### GetChangelogEntryExcludeBaseDNOk

`func (o *BackendListResponseResourcesInner) GetChangelogEntryExcludeBaseDNOk() (*[]string, bool)`

GetChangelogEntryExcludeBaseDNOk returns a tuple with the ChangelogEntryExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryExcludeBaseDN

`func (o *BackendListResponseResourcesInner) SetChangelogEntryExcludeBaseDN(v []string)`

SetChangelogEntryExcludeBaseDN sets ChangelogEntryExcludeBaseDN field to given value.

### HasChangelogEntryExcludeBaseDN

`func (o *BackendListResponseResourcesInner) HasChangelogEntryExcludeBaseDN() bool`

HasChangelogEntryExcludeBaseDN returns a boolean if a field has been set.

### GetChangelogEntryIncludeFilter

`func (o *BackendListResponseResourcesInner) GetChangelogEntryIncludeFilter() []string`

GetChangelogEntryIncludeFilter returns the ChangelogEntryIncludeFilter field if non-nil, zero value otherwise.

### GetChangelogEntryIncludeFilterOk

`func (o *BackendListResponseResourcesInner) GetChangelogEntryIncludeFilterOk() (*[]string, bool)`

GetChangelogEntryIncludeFilterOk returns a tuple with the ChangelogEntryIncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryIncludeFilter

`func (o *BackendListResponseResourcesInner) SetChangelogEntryIncludeFilter(v []string)`

SetChangelogEntryIncludeFilter sets ChangelogEntryIncludeFilter field to given value.

### HasChangelogEntryIncludeFilter

`func (o *BackendListResponseResourcesInner) HasChangelogEntryIncludeFilter() bool`

HasChangelogEntryIncludeFilter returns a boolean if a field has been set.

### GetChangelogEntryExcludeFilter

`func (o *BackendListResponseResourcesInner) GetChangelogEntryExcludeFilter() []string`

GetChangelogEntryExcludeFilter returns the ChangelogEntryExcludeFilter field if non-nil, zero value otherwise.

### GetChangelogEntryExcludeFilterOk

`func (o *BackendListResponseResourcesInner) GetChangelogEntryExcludeFilterOk() (*[]string, bool)`

GetChangelogEntryExcludeFilterOk returns a tuple with the ChangelogEntryExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryExcludeFilter

`func (o *BackendListResponseResourcesInner) SetChangelogEntryExcludeFilter(v []string)`

SetChangelogEntryExcludeFilter sets ChangelogEntryExcludeFilter field to given value.

### HasChangelogEntryExcludeFilter

`func (o *BackendListResponseResourcesInner) HasChangelogEntryExcludeFilter() bool`

HasChangelogEntryExcludeFilter returns a boolean if a field has been set.

### GetChangelogIncludeAttribute

`func (o *BackendListResponseResourcesInner) GetChangelogIncludeAttribute() []string`

GetChangelogIncludeAttribute returns the ChangelogIncludeAttribute field if non-nil, zero value otherwise.

### GetChangelogIncludeAttributeOk

`func (o *BackendListResponseResourcesInner) GetChangelogIncludeAttributeOk() (*[]string, bool)`

GetChangelogIncludeAttributeOk returns a tuple with the ChangelogIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogIncludeAttribute

`func (o *BackendListResponseResourcesInner) SetChangelogIncludeAttribute(v []string)`

SetChangelogIncludeAttribute sets ChangelogIncludeAttribute field to given value.

### HasChangelogIncludeAttribute

`func (o *BackendListResponseResourcesInner) HasChangelogIncludeAttribute() bool`

HasChangelogIncludeAttribute returns a boolean if a field has been set.

### GetChangelogExcludeAttribute

`func (o *BackendListResponseResourcesInner) GetChangelogExcludeAttribute() []string`

GetChangelogExcludeAttribute returns the ChangelogExcludeAttribute field if non-nil, zero value otherwise.

### GetChangelogExcludeAttributeOk

`func (o *BackendListResponseResourcesInner) GetChangelogExcludeAttributeOk() (*[]string, bool)`

GetChangelogExcludeAttributeOk returns a tuple with the ChangelogExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogExcludeAttribute

`func (o *BackendListResponseResourcesInner) SetChangelogExcludeAttribute(v []string)`

SetChangelogExcludeAttribute sets ChangelogExcludeAttribute field to given value.

### HasChangelogExcludeAttribute

`func (o *BackendListResponseResourcesInner) HasChangelogExcludeAttribute() bool`

HasChangelogExcludeAttribute returns a boolean if a field has been set.

### GetChangelogDeletedEntryIncludeAttribute

`func (o *BackendListResponseResourcesInner) GetChangelogDeletedEntryIncludeAttribute() []string`

GetChangelogDeletedEntryIncludeAttribute returns the ChangelogDeletedEntryIncludeAttribute field if non-nil, zero value otherwise.

### GetChangelogDeletedEntryIncludeAttributeOk

`func (o *BackendListResponseResourcesInner) GetChangelogDeletedEntryIncludeAttributeOk() (*[]string, bool)`

GetChangelogDeletedEntryIncludeAttributeOk returns a tuple with the ChangelogDeletedEntryIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogDeletedEntryIncludeAttribute

`func (o *BackendListResponseResourcesInner) SetChangelogDeletedEntryIncludeAttribute(v []string)`

SetChangelogDeletedEntryIncludeAttribute sets ChangelogDeletedEntryIncludeAttribute field to given value.

### HasChangelogDeletedEntryIncludeAttribute

`func (o *BackendListResponseResourcesInner) HasChangelogDeletedEntryIncludeAttribute() bool`

HasChangelogDeletedEntryIncludeAttribute returns a boolean if a field has been set.

### GetChangelogDeletedEntryExcludeAttribute

`func (o *BackendListResponseResourcesInner) GetChangelogDeletedEntryExcludeAttribute() []string`

GetChangelogDeletedEntryExcludeAttribute returns the ChangelogDeletedEntryExcludeAttribute field if non-nil, zero value otherwise.

### GetChangelogDeletedEntryExcludeAttributeOk

`func (o *BackendListResponseResourcesInner) GetChangelogDeletedEntryExcludeAttributeOk() (*[]string, bool)`

GetChangelogDeletedEntryExcludeAttributeOk returns a tuple with the ChangelogDeletedEntryExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogDeletedEntryExcludeAttribute

`func (o *BackendListResponseResourcesInner) SetChangelogDeletedEntryExcludeAttribute(v []string)`

SetChangelogDeletedEntryExcludeAttribute sets ChangelogDeletedEntryExcludeAttribute field to given value.

### HasChangelogDeletedEntryExcludeAttribute

`func (o *BackendListResponseResourcesInner) HasChangelogDeletedEntryExcludeAttribute() bool`

HasChangelogDeletedEntryExcludeAttribute returns a boolean if a field has been set.

### GetChangelogIncludeKeyAttribute

`func (o *BackendListResponseResourcesInner) GetChangelogIncludeKeyAttribute() []string`

GetChangelogIncludeKeyAttribute returns the ChangelogIncludeKeyAttribute field if non-nil, zero value otherwise.

### GetChangelogIncludeKeyAttributeOk

`func (o *BackendListResponseResourcesInner) GetChangelogIncludeKeyAttributeOk() (*[]string, bool)`

GetChangelogIncludeKeyAttributeOk returns a tuple with the ChangelogIncludeKeyAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogIncludeKeyAttribute

`func (o *BackendListResponseResourcesInner) SetChangelogIncludeKeyAttribute(v []string)`

SetChangelogIncludeKeyAttribute sets ChangelogIncludeKeyAttribute field to given value.

### HasChangelogIncludeKeyAttribute

`func (o *BackendListResponseResourcesInner) HasChangelogIncludeKeyAttribute() bool`

HasChangelogIncludeKeyAttribute returns a boolean if a field has been set.

### GetChangelogMaxBeforeAfterValues

`func (o *BackendListResponseResourcesInner) GetChangelogMaxBeforeAfterValues() int64`

GetChangelogMaxBeforeAfterValues returns the ChangelogMaxBeforeAfterValues field if non-nil, zero value otherwise.

### GetChangelogMaxBeforeAfterValuesOk

`func (o *BackendListResponseResourcesInner) GetChangelogMaxBeforeAfterValuesOk() (*int64, bool)`

GetChangelogMaxBeforeAfterValuesOk returns a tuple with the ChangelogMaxBeforeAfterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogMaxBeforeAfterValues

`func (o *BackendListResponseResourcesInner) SetChangelogMaxBeforeAfterValues(v int64)`

SetChangelogMaxBeforeAfterValues sets ChangelogMaxBeforeAfterValues field to given value.

### HasChangelogMaxBeforeAfterValues

`func (o *BackendListResponseResourcesInner) HasChangelogMaxBeforeAfterValues() bool`

HasChangelogMaxBeforeAfterValues returns a boolean if a field has been set.

### GetWriteLastmodAttributes

`func (o *BackendListResponseResourcesInner) GetWriteLastmodAttributes() bool`

GetWriteLastmodAttributes returns the WriteLastmodAttributes field if non-nil, zero value otherwise.

### GetWriteLastmodAttributesOk

`func (o *BackendListResponseResourcesInner) GetWriteLastmodAttributesOk() (*bool, bool)`

GetWriteLastmodAttributesOk returns a tuple with the WriteLastmodAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLastmodAttributes

`func (o *BackendListResponseResourcesInner) SetWriteLastmodAttributes(v bool)`

SetWriteLastmodAttributes sets WriteLastmodAttributes field to given value.

### HasWriteLastmodAttributes

`func (o *BackendListResponseResourcesInner) HasWriteLastmodAttributes() bool`

HasWriteLastmodAttributes returns a boolean if a field has been set.

### GetUseReversibleForm

`func (o *BackendListResponseResourcesInner) GetUseReversibleForm() bool`

GetUseReversibleForm returns the UseReversibleForm field if non-nil, zero value otherwise.

### GetUseReversibleFormOk

`func (o *BackendListResponseResourcesInner) GetUseReversibleFormOk() (*bool, bool)`

GetUseReversibleFormOk returns a tuple with the UseReversibleForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReversibleForm

`func (o *BackendListResponseResourcesInner) SetUseReversibleForm(v bool)`

SetUseReversibleForm sets UseReversibleForm field to given value.

### HasUseReversibleForm

`func (o *BackendListResponseResourcesInner) HasUseReversibleForm() bool`

HasUseReversibleForm returns a boolean if a field has been set.

### GetIncludeVirtualAttributes

`func (o *BackendListResponseResourcesInner) GetIncludeVirtualAttributes() []EnumbackendIncludeVirtualAttributesProp`

GetIncludeVirtualAttributes returns the IncludeVirtualAttributes field if non-nil, zero value otherwise.

### GetIncludeVirtualAttributesOk

`func (o *BackendListResponseResourcesInner) GetIncludeVirtualAttributesOk() (*[]EnumbackendIncludeVirtualAttributesProp, bool)`

GetIncludeVirtualAttributesOk returns a tuple with the IncludeVirtualAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVirtualAttributes

`func (o *BackendListResponseResourcesInner) SetIncludeVirtualAttributes(v []EnumbackendIncludeVirtualAttributesProp)`

SetIncludeVirtualAttributes sets IncludeVirtualAttributes field to given value.

### HasIncludeVirtualAttributes

`func (o *BackendListResponseResourcesInner) HasIncludeVirtualAttributes() bool`

HasIncludeVirtualAttributes returns a boolean if a field has been set.

### GetApplyAccessControlsToChangelogEntryContents

`func (o *BackendListResponseResourcesInner) GetApplyAccessControlsToChangelogEntryContents() bool`

GetApplyAccessControlsToChangelogEntryContents returns the ApplyAccessControlsToChangelogEntryContents field if non-nil, zero value otherwise.

### GetApplyAccessControlsToChangelogEntryContentsOk

`func (o *BackendListResponseResourcesInner) GetApplyAccessControlsToChangelogEntryContentsOk() (*bool, bool)`

GetApplyAccessControlsToChangelogEntryContentsOk returns a tuple with the ApplyAccessControlsToChangelogEntryContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAccessControlsToChangelogEntryContents

`func (o *BackendListResponseResourcesInner) SetApplyAccessControlsToChangelogEntryContents(v bool)`

SetApplyAccessControlsToChangelogEntryContents sets ApplyAccessControlsToChangelogEntryContents field to given value.

### HasApplyAccessControlsToChangelogEntryContents

`func (o *BackendListResponseResourcesInner) HasApplyAccessControlsToChangelogEntryContents() bool`

HasApplyAccessControlsToChangelogEntryContents returns a boolean if a field has been set.

### GetReportExcludedChangelogAttributes

`func (o *BackendListResponseResourcesInner) GetReportExcludedChangelogAttributes() EnumbackendReportExcludedChangelogAttributesProp`

GetReportExcludedChangelogAttributes returns the ReportExcludedChangelogAttributes field if non-nil, zero value otherwise.

### GetReportExcludedChangelogAttributesOk

`func (o *BackendListResponseResourcesInner) GetReportExcludedChangelogAttributesOk() (*EnumbackendReportExcludedChangelogAttributesProp, bool)`

GetReportExcludedChangelogAttributesOk returns a tuple with the ReportExcludedChangelogAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExcludedChangelogAttributes

`func (o *BackendListResponseResourcesInner) SetReportExcludedChangelogAttributes(v EnumbackendReportExcludedChangelogAttributesProp)`

SetReportExcludedChangelogAttributes sets ReportExcludedChangelogAttributes field to given value.

### HasReportExcludedChangelogAttributes

`func (o *BackendListResponseResourcesInner) HasReportExcludedChangelogAttributes() bool`

HasReportExcludedChangelogAttributes returns a boolean if a field has been set.

### GetSoftDeleteEntryIncludedOperation

`func (o *BackendListResponseResourcesInner) GetSoftDeleteEntryIncludedOperation() []EnumbackendSoftDeleteEntryIncludedOperationProp`

GetSoftDeleteEntryIncludedOperation returns the SoftDeleteEntryIncludedOperation field if non-nil, zero value otherwise.

### GetSoftDeleteEntryIncludedOperationOk

`func (o *BackendListResponseResourcesInner) GetSoftDeleteEntryIncludedOperationOk() (*[]EnumbackendSoftDeleteEntryIncludedOperationProp, bool)`

GetSoftDeleteEntryIncludedOperationOk returns a tuple with the SoftDeleteEntryIncludedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteEntryIncludedOperation

`func (o *BackendListResponseResourcesInner) SetSoftDeleteEntryIncludedOperation(v []EnumbackendSoftDeleteEntryIncludedOperationProp)`

SetSoftDeleteEntryIncludedOperation sets SoftDeleteEntryIncludedOperation field to given value.

### HasSoftDeleteEntryIncludedOperation

`func (o *BackendListResponseResourcesInner) HasSoftDeleteEntryIncludedOperation() bool`

HasSoftDeleteEntryIncludedOperation returns a boolean if a field has been set.

### GetUncachedId2entryCacheMode

`func (o *BackendListResponseResourcesInner) GetUncachedId2entryCacheMode() EnumbackendUncachedId2entryCacheModeProp`

GetUncachedId2entryCacheMode returns the UncachedId2entryCacheMode field if non-nil, zero value otherwise.

### GetUncachedId2entryCacheModeOk

`func (o *BackendListResponseResourcesInner) GetUncachedId2entryCacheModeOk() (*EnumbackendUncachedId2entryCacheModeProp, bool)`

GetUncachedId2entryCacheModeOk returns a tuple with the UncachedId2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedId2entryCacheMode

`func (o *BackendListResponseResourcesInner) SetUncachedId2entryCacheMode(v EnumbackendUncachedId2entryCacheModeProp)`

SetUncachedId2entryCacheMode sets UncachedId2entryCacheMode field to given value.

### HasUncachedId2entryCacheMode

`func (o *BackendListResponseResourcesInner) HasUncachedId2entryCacheMode() bool`

HasUncachedId2entryCacheMode returns a boolean if a field has been set.

### GetUncachedAttributeCriteria

`func (o *BackendListResponseResourcesInner) GetUncachedAttributeCriteria() string`

GetUncachedAttributeCriteria returns the UncachedAttributeCriteria field if non-nil, zero value otherwise.

### GetUncachedAttributeCriteriaOk

`func (o *BackendListResponseResourcesInner) GetUncachedAttributeCriteriaOk() (*string, bool)`

GetUncachedAttributeCriteriaOk returns a tuple with the UncachedAttributeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedAttributeCriteria

`func (o *BackendListResponseResourcesInner) SetUncachedAttributeCriteria(v string)`

SetUncachedAttributeCriteria sets UncachedAttributeCriteria field to given value.

### HasUncachedAttributeCriteria

`func (o *BackendListResponseResourcesInner) HasUncachedAttributeCriteria() bool`

HasUncachedAttributeCriteria returns a boolean if a field has been set.

### GetUncachedEntryCriteria

`func (o *BackendListResponseResourcesInner) GetUncachedEntryCriteria() string`

GetUncachedEntryCriteria returns the UncachedEntryCriteria field if non-nil, zero value otherwise.

### GetUncachedEntryCriteriaOk

`func (o *BackendListResponseResourcesInner) GetUncachedEntryCriteriaOk() (*string, bool)`

GetUncachedEntryCriteriaOk returns a tuple with the UncachedEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedEntryCriteria

`func (o *BackendListResponseResourcesInner) SetUncachedEntryCriteria(v string)`

SetUncachedEntryCriteria sets UncachedEntryCriteria field to given value.

### HasUncachedEntryCriteria

`func (o *BackendListResponseResourcesInner) HasUncachedEntryCriteria() bool`

HasUncachedEntryCriteria returns a boolean if a field has been set.

### GetSetDegradedAlertForUntrustedIndex

`func (o *BackendListResponseResourcesInner) GetSetDegradedAlertForUntrustedIndex() bool`

GetSetDegradedAlertForUntrustedIndex returns the SetDegradedAlertForUntrustedIndex field if non-nil, zero value otherwise.

### GetSetDegradedAlertForUntrustedIndexOk

`func (o *BackendListResponseResourcesInner) GetSetDegradedAlertForUntrustedIndexOk() (*bool, bool)`

GetSetDegradedAlertForUntrustedIndexOk returns a tuple with the SetDegradedAlertForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertForUntrustedIndex

`func (o *BackendListResponseResourcesInner) SetSetDegradedAlertForUntrustedIndex(v bool)`

SetSetDegradedAlertForUntrustedIndex sets SetDegradedAlertForUntrustedIndex field to given value.

### HasSetDegradedAlertForUntrustedIndex

`func (o *BackendListResponseResourcesInner) HasSetDegradedAlertForUntrustedIndex() bool`

HasSetDegradedAlertForUntrustedIndex returns a boolean if a field has been set.

### GetReturnUnavailableForUntrustedIndex

`func (o *BackendListResponseResourcesInner) GetReturnUnavailableForUntrustedIndex() bool`

GetReturnUnavailableForUntrustedIndex returns the ReturnUnavailableForUntrustedIndex field if non-nil, zero value otherwise.

### GetReturnUnavailableForUntrustedIndexOk

`func (o *BackendListResponseResourcesInner) GetReturnUnavailableForUntrustedIndexOk() (*bool, bool)`

GetReturnUnavailableForUntrustedIndexOk returns a tuple with the ReturnUnavailableForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableForUntrustedIndex

`func (o *BackendListResponseResourcesInner) SetReturnUnavailableForUntrustedIndex(v bool)`

SetReturnUnavailableForUntrustedIndex sets ReturnUnavailableForUntrustedIndex field to given value.

### HasReturnUnavailableForUntrustedIndex

`func (o *BackendListResponseResourcesInner) HasReturnUnavailableForUntrustedIndex() bool`

HasReturnUnavailableForUntrustedIndex returns a boolean if a field has been set.

### GetProcessFiltersWithUndefinedAttributeTypes

`func (o *BackendListResponseResourcesInner) GetProcessFiltersWithUndefinedAttributeTypes() bool`

GetProcessFiltersWithUndefinedAttributeTypes returns the ProcessFiltersWithUndefinedAttributeTypes field if non-nil, zero value otherwise.

### GetProcessFiltersWithUndefinedAttributeTypesOk

`func (o *BackendListResponseResourcesInner) GetProcessFiltersWithUndefinedAttributeTypesOk() (*bool, bool)`

GetProcessFiltersWithUndefinedAttributeTypesOk returns a tuple with the ProcessFiltersWithUndefinedAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFiltersWithUndefinedAttributeTypes

`func (o *BackendListResponseResourcesInner) SetProcessFiltersWithUndefinedAttributeTypes(v bool)`

SetProcessFiltersWithUndefinedAttributeTypes sets ProcessFiltersWithUndefinedAttributeTypes field to given value.

### HasProcessFiltersWithUndefinedAttributeTypes

`func (o *BackendListResponseResourcesInner) HasProcessFiltersWithUndefinedAttributeTypes() bool`

HasProcessFiltersWithUndefinedAttributeTypes returns a boolean if a field has been set.

### GetCompactCommonParentDN

`func (o *BackendListResponseResourcesInner) GetCompactCommonParentDN() []string`

GetCompactCommonParentDN returns the CompactCommonParentDN field if non-nil, zero value otherwise.

### GetCompactCommonParentDNOk

`func (o *BackendListResponseResourcesInner) GetCompactCommonParentDNOk() (*[]string, bool)`

GetCompactCommonParentDNOk returns a tuple with the CompactCommonParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactCommonParentDN

`func (o *BackendListResponseResourcesInner) SetCompactCommonParentDN(v []string)`

SetCompactCommonParentDN sets CompactCommonParentDN field to given value.

### HasCompactCommonParentDN

`func (o *BackendListResponseResourcesInner) HasCompactCommonParentDN() bool`

HasCompactCommonParentDN returns a boolean if a field has been set.

### GetCompressEntries

`func (o *BackendListResponseResourcesInner) GetCompressEntries() bool`

GetCompressEntries returns the CompressEntries field if non-nil, zero value otherwise.

### GetCompressEntriesOk

`func (o *BackendListResponseResourcesInner) GetCompressEntriesOk() (*bool, bool)`

GetCompressEntriesOk returns a tuple with the CompressEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressEntries

`func (o *BackendListResponseResourcesInner) SetCompressEntries(v bool)`

SetCompressEntries sets CompressEntries field to given value.

### HasCompressEntries

`func (o *BackendListResponseResourcesInner) HasCompressEntries() bool`

HasCompressEntries returns a boolean if a field has been set.

### GetHashEntries

`func (o *BackendListResponseResourcesInner) GetHashEntries() bool`

GetHashEntries returns the HashEntries field if non-nil, zero value otherwise.

### GetHashEntriesOk

`func (o *BackendListResponseResourcesInner) GetHashEntriesOk() (*bool, bool)`

GetHashEntriesOk returns a tuple with the HashEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashEntries

`func (o *BackendListResponseResourcesInner) SetHashEntries(v bool)`

SetHashEntries sets HashEntries field to given value.

### HasHashEntries

`func (o *BackendListResponseResourcesInner) HasHashEntries() bool`

HasHashEntries returns a boolean if a field has been set.

### GetDbNumCleanerThreads

`func (o *BackendListResponseResourcesInner) GetDbNumCleanerThreads() int64`

GetDbNumCleanerThreads returns the DbNumCleanerThreads field if non-nil, zero value otherwise.

### GetDbNumCleanerThreadsOk

`func (o *BackendListResponseResourcesInner) GetDbNumCleanerThreadsOk() (*int64, bool)`

GetDbNumCleanerThreadsOk returns a tuple with the DbNumCleanerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbNumCleanerThreads

`func (o *BackendListResponseResourcesInner) SetDbNumCleanerThreads(v int64)`

SetDbNumCleanerThreads sets DbNumCleanerThreads field to given value.

### HasDbNumCleanerThreads

`func (o *BackendListResponseResourcesInner) HasDbNumCleanerThreads() bool`

HasDbNumCleanerThreads returns a boolean if a field has been set.

### GetDbCleanerMinUtilization

`func (o *BackendListResponseResourcesInner) GetDbCleanerMinUtilization() int64`

GetDbCleanerMinUtilization returns the DbCleanerMinUtilization field if non-nil, zero value otherwise.

### GetDbCleanerMinUtilizationOk

`func (o *BackendListResponseResourcesInner) GetDbCleanerMinUtilizationOk() (*int64, bool)`

GetDbCleanerMinUtilizationOk returns a tuple with the DbCleanerMinUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCleanerMinUtilization

`func (o *BackendListResponseResourcesInner) SetDbCleanerMinUtilization(v int64)`

SetDbCleanerMinUtilization sets DbCleanerMinUtilization field to given value.

### HasDbCleanerMinUtilization

`func (o *BackendListResponseResourcesInner) HasDbCleanerMinUtilization() bool`

HasDbCleanerMinUtilization returns a boolean if a field has been set.

### GetDbEvictorCriticalPercentage

`func (o *BackendListResponseResourcesInner) GetDbEvictorCriticalPercentage() int64`

GetDbEvictorCriticalPercentage returns the DbEvictorCriticalPercentage field if non-nil, zero value otherwise.

### GetDbEvictorCriticalPercentageOk

`func (o *BackendListResponseResourcesInner) GetDbEvictorCriticalPercentageOk() (*int64, bool)`

GetDbEvictorCriticalPercentageOk returns a tuple with the DbEvictorCriticalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEvictorCriticalPercentage

`func (o *BackendListResponseResourcesInner) SetDbEvictorCriticalPercentage(v int64)`

SetDbEvictorCriticalPercentage sets DbEvictorCriticalPercentage field to given value.

### HasDbEvictorCriticalPercentage

`func (o *BackendListResponseResourcesInner) HasDbEvictorCriticalPercentage() bool`

HasDbEvictorCriticalPercentage returns a boolean if a field has been set.

### GetDbCheckpointerWakeupInterval

`func (o *BackendListResponseResourcesInner) GetDbCheckpointerWakeupInterval() string`

GetDbCheckpointerWakeupInterval returns the DbCheckpointerWakeupInterval field if non-nil, zero value otherwise.

### GetDbCheckpointerWakeupIntervalOk

`func (o *BackendListResponseResourcesInner) GetDbCheckpointerWakeupIntervalOk() (*string, bool)`

GetDbCheckpointerWakeupIntervalOk returns a tuple with the DbCheckpointerWakeupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCheckpointerWakeupInterval

`func (o *BackendListResponseResourcesInner) SetDbCheckpointerWakeupInterval(v string)`

SetDbCheckpointerWakeupInterval sets DbCheckpointerWakeupInterval field to given value.

### HasDbCheckpointerWakeupInterval

`func (o *BackendListResponseResourcesInner) HasDbCheckpointerWakeupInterval() bool`

HasDbCheckpointerWakeupInterval returns a boolean if a field has been set.

### GetDbBackgroundSyncInterval

`func (o *BackendListResponseResourcesInner) GetDbBackgroundSyncInterval() string`

GetDbBackgroundSyncInterval returns the DbBackgroundSyncInterval field if non-nil, zero value otherwise.

### GetDbBackgroundSyncIntervalOk

`func (o *BackendListResponseResourcesInner) GetDbBackgroundSyncIntervalOk() (*string, bool)`

GetDbBackgroundSyncIntervalOk returns a tuple with the DbBackgroundSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbBackgroundSyncInterval

`func (o *BackendListResponseResourcesInner) SetDbBackgroundSyncInterval(v string)`

SetDbBackgroundSyncInterval sets DbBackgroundSyncInterval field to given value.

### HasDbBackgroundSyncInterval

`func (o *BackendListResponseResourcesInner) HasDbBackgroundSyncInterval() bool`

HasDbBackgroundSyncInterval returns a boolean if a field has been set.

### GetDbUseThreadLocalHandles

`func (o *BackendListResponseResourcesInner) GetDbUseThreadLocalHandles() bool`

GetDbUseThreadLocalHandles returns the DbUseThreadLocalHandles field if non-nil, zero value otherwise.

### GetDbUseThreadLocalHandlesOk

`func (o *BackendListResponseResourcesInner) GetDbUseThreadLocalHandlesOk() (*bool, bool)`

GetDbUseThreadLocalHandlesOk returns a tuple with the DbUseThreadLocalHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUseThreadLocalHandles

`func (o *BackendListResponseResourcesInner) SetDbUseThreadLocalHandles(v bool)`

SetDbUseThreadLocalHandles sets DbUseThreadLocalHandles field to given value.

### HasDbUseThreadLocalHandles

`func (o *BackendListResponseResourcesInner) HasDbUseThreadLocalHandles() bool`

HasDbUseThreadLocalHandles returns a boolean if a field has been set.

### GetDbLogFileMax

`func (o *BackendListResponseResourcesInner) GetDbLogFileMax() string`

GetDbLogFileMax returns the DbLogFileMax field if non-nil, zero value otherwise.

### GetDbLogFileMaxOk

`func (o *BackendListResponseResourcesInner) GetDbLogFileMaxOk() (*string, bool)`

GetDbLogFileMaxOk returns a tuple with the DbLogFileMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLogFileMax

`func (o *BackendListResponseResourcesInner) SetDbLogFileMax(v string)`

SetDbLogFileMax sets DbLogFileMax field to given value.

### HasDbLogFileMax

`func (o *BackendListResponseResourcesInner) HasDbLogFileMax() bool`

HasDbLogFileMax returns a boolean if a field has been set.

### GetDbLoggingLevel

`func (o *BackendListResponseResourcesInner) GetDbLoggingLevel() string`

GetDbLoggingLevel returns the DbLoggingLevel field if non-nil, zero value otherwise.

### GetDbLoggingLevelOk

`func (o *BackendListResponseResourcesInner) GetDbLoggingLevelOk() (*string, bool)`

GetDbLoggingLevelOk returns a tuple with the DbLoggingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLoggingLevel

`func (o *BackendListResponseResourcesInner) SetDbLoggingLevel(v string)`

SetDbLoggingLevel sets DbLoggingLevel field to given value.

### HasDbLoggingLevel

`func (o *BackendListResponseResourcesInner) HasDbLoggingLevel() bool`

HasDbLoggingLevel returns a boolean if a field has been set.

### GetDefaultCacheMode

`func (o *BackendListResponseResourcesInner) GetDefaultCacheMode() EnumbackendDefaultCacheModeProp`

GetDefaultCacheMode returns the DefaultCacheMode field if non-nil, zero value otherwise.

### GetDefaultCacheModeOk

`func (o *BackendListResponseResourcesInner) GetDefaultCacheModeOk() (*EnumbackendDefaultCacheModeProp, bool)`

GetDefaultCacheModeOk returns a tuple with the DefaultCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCacheMode

`func (o *BackendListResponseResourcesInner) SetDefaultCacheMode(v EnumbackendDefaultCacheModeProp)`

SetDefaultCacheMode sets DefaultCacheMode field to given value.

### HasDefaultCacheMode

`func (o *BackendListResponseResourcesInner) HasDefaultCacheMode() bool`

HasDefaultCacheMode returns a boolean if a field has been set.

### GetId2entryCacheMode

`func (o *BackendListResponseResourcesInner) GetId2entryCacheMode() EnumbackendId2entryCacheModeProp`

GetId2entryCacheMode returns the Id2entryCacheMode field if non-nil, zero value otherwise.

### GetId2entryCacheModeOk

`func (o *BackendListResponseResourcesInner) GetId2entryCacheModeOk() (*EnumbackendId2entryCacheModeProp, bool)`

GetId2entryCacheModeOk returns a tuple with the Id2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2entryCacheMode

`func (o *BackendListResponseResourcesInner) SetId2entryCacheMode(v EnumbackendId2entryCacheModeProp)`

SetId2entryCacheMode sets Id2entryCacheMode field to given value.

### HasId2entryCacheMode

`func (o *BackendListResponseResourcesInner) HasId2entryCacheMode() bool`

HasId2entryCacheMode returns a boolean if a field has been set.

### GetDn2idCacheMode

`func (o *BackendListResponseResourcesInner) GetDn2idCacheMode() EnumbackendDn2idCacheModeProp`

GetDn2idCacheMode returns the Dn2idCacheMode field if non-nil, zero value otherwise.

### GetDn2idCacheModeOk

`func (o *BackendListResponseResourcesInner) GetDn2idCacheModeOk() (*EnumbackendDn2idCacheModeProp, bool)`

GetDn2idCacheModeOk returns a tuple with the Dn2idCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2idCacheMode

`func (o *BackendListResponseResourcesInner) SetDn2idCacheMode(v EnumbackendDn2idCacheModeProp)`

SetDn2idCacheMode sets Dn2idCacheMode field to given value.

### HasDn2idCacheMode

`func (o *BackendListResponseResourcesInner) HasDn2idCacheMode() bool`

HasDn2idCacheMode returns a boolean if a field has been set.

### GetId2childrenCacheMode

`func (o *BackendListResponseResourcesInner) GetId2childrenCacheMode() EnumbackendId2childrenCacheModeProp`

GetId2childrenCacheMode returns the Id2childrenCacheMode field if non-nil, zero value otherwise.

### GetId2childrenCacheModeOk

`func (o *BackendListResponseResourcesInner) GetId2childrenCacheModeOk() (*EnumbackendId2childrenCacheModeProp, bool)`

GetId2childrenCacheModeOk returns a tuple with the Id2childrenCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenCacheMode

`func (o *BackendListResponseResourcesInner) SetId2childrenCacheMode(v EnumbackendId2childrenCacheModeProp)`

SetId2childrenCacheMode sets Id2childrenCacheMode field to given value.

### HasId2childrenCacheMode

`func (o *BackendListResponseResourcesInner) HasId2childrenCacheMode() bool`

HasId2childrenCacheMode returns a boolean if a field has been set.

### GetId2subtreeCacheMode

`func (o *BackendListResponseResourcesInner) GetId2subtreeCacheMode() EnumbackendId2subtreeCacheModeProp`

GetId2subtreeCacheMode returns the Id2subtreeCacheMode field if non-nil, zero value otherwise.

### GetId2subtreeCacheModeOk

`func (o *BackendListResponseResourcesInner) GetId2subtreeCacheModeOk() (*EnumbackendId2subtreeCacheModeProp, bool)`

GetId2subtreeCacheModeOk returns a tuple with the Id2subtreeCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeCacheMode

`func (o *BackendListResponseResourcesInner) SetId2subtreeCacheMode(v EnumbackendId2subtreeCacheModeProp)`

SetId2subtreeCacheMode sets Id2subtreeCacheMode field to given value.

### HasId2subtreeCacheMode

`func (o *BackendListResponseResourcesInner) HasId2subtreeCacheMode() bool`

HasId2subtreeCacheMode returns a boolean if a field has been set.

### GetDn2uriCacheMode

`func (o *BackendListResponseResourcesInner) GetDn2uriCacheMode() EnumbackendDn2uriCacheModeProp`

GetDn2uriCacheMode returns the Dn2uriCacheMode field if non-nil, zero value otherwise.

### GetDn2uriCacheModeOk

`func (o *BackendListResponseResourcesInner) GetDn2uriCacheModeOk() (*EnumbackendDn2uriCacheModeProp, bool)`

GetDn2uriCacheModeOk returns a tuple with the Dn2uriCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2uriCacheMode

`func (o *BackendListResponseResourcesInner) SetDn2uriCacheMode(v EnumbackendDn2uriCacheModeProp)`

SetDn2uriCacheMode sets Dn2uriCacheMode field to given value.

### HasDn2uriCacheMode

`func (o *BackendListResponseResourcesInner) HasDn2uriCacheMode() bool`

HasDn2uriCacheMode returns a boolean if a field has been set.

### GetPrimeMethod

`func (o *BackendListResponseResourcesInner) GetPrimeMethod() []EnumbackendPrimeMethodProp`

GetPrimeMethod returns the PrimeMethod field if non-nil, zero value otherwise.

### GetPrimeMethodOk

`func (o *BackendListResponseResourcesInner) GetPrimeMethodOk() (*[]EnumbackendPrimeMethodProp, bool)`

GetPrimeMethodOk returns a tuple with the PrimeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeMethod

`func (o *BackendListResponseResourcesInner) SetPrimeMethod(v []EnumbackendPrimeMethodProp)`

SetPrimeMethod sets PrimeMethod field to given value.

### HasPrimeMethod

`func (o *BackendListResponseResourcesInner) HasPrimeMethod() bool`

HasPrimeMethod returns a boolean if a field has been set.

### GetPrimeThreadCount

`func (o *BackendListResponseResourcesInner) GetPrimeThreadCount() int64`

GetPrimeThreadCount returns the PrimeThreadCount field if non-nil, zero value otherwise.

### GetPrimeThreadCountOk

`func (o *BackendListResponseResourcesInner) GetPrimeThreadCountOk() (*int64, bool)`

GetPrimeThreadCountOk returns a tuple with the PrimeThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeThreadCount

`func (o *BackendListResponseResourcesInner) SetPrimeThreadCount(v int64)`

SetPrimeThreadCount sets PrimeThreadCount field to given value.

### HasPrimeThreadCount

`func (o *BackendListResponseResourcesInner) HasPrimeThreadCount() bool`

HasPrimeThreadCount returns a boolean if a field has been set.

### GetPrimeTimeLimit

`func (o *BackendListResponseResourcesInner) GetPrimeTimeLimit() string`

GetPrimeTimeLimit returns the PrimeTimeLimit field if non-nil, zero value otherwise.

### GetPrimeTimeLimitOk

`func (o *BackendListResponseResourcesInner) GetPrimeTimeLimitOk() (*string, bool)`

GetPrimeTimeLimitOk returns a tuple with the PrimeTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeTimeLimit

`func (o *BackendListResponseResourcesInner) SetPrimeTimeLimit(v string)`

SetPrimeTimeLimit sets PrimeTimeLimit field to given value.

### HasPrimeTimeLimit

`func (o *BackendListResponseResourcesInner) HasPrimeTimeLimit() bool`

HasPrimeTimeLimit returns a boolean if a field has been set.

### GetPrimeAllIndexes

`func (o *BackendListResponseResourcesInner) GetPrimeAllIndexes() bool`

GetPrimeAllIndexes returns the PrimeAllIndexes field if non-nil, zero value otherwise.

### GetPrimeAllIndexesOk

`func (o *BackendListResponseResourcesInner) GetPrimeAllIndexesOk() (*bool, bool)`

GetPrimeAllIndexesOk returns a tuple with the PrimeAllIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeAllIndexes

`func (o *BackendListResponseResourcesInner) SetPrimeAllIndexes(v bool)`

SetPrimeAllIndexes sets PrimeAllIndexes field to given value.

### HasPrimeAllIndexes

`func (o *BackendListResponseResourcesInner) HasPrimeAllIndexes() bool`

HasPrimeAllIndexes returns a boolean if a field has been set.

### GetSystemIndexToPrime

`func (o *BackendListResponseResourcesInner) GetSystemIndexToPrime() []EnumbackendSystemIndexToPrimeProp`

GetSystemIndexToPrime returns the SystemIndexToPrime field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeOk

`func (o *BackendListResponseResourcesInner) GetSystemIndexToPrimeOk() (*[]EnumbackendSystemIndexToPrimeProp, bool)`

GetSystemIndexToPrimeOk returns a tuple with the SystemIndexToPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrime

`func (o *BackendListResponseResourcesInner) SetSystemIndexToPrime(v []EnumbackendSystemIndexToPrimeProp)`

SetSystemIndexToPrime sets SystemIndexToPrime field to given value.

### HasSystemIndexToPrime

`func (o *BackendListResponseResourcesInner) HasSystemIndexToPrime() bool`

HasSystemIndexToPrime returns a boolean if a field has been set.

### GetSystemIndexToPrimeInternalNodesOnly

`func (o *BackendListResponseResourcesInner) GetSystemIndexToPrimeInternalNodesOnly() []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp`

GetSystemIndexToPrimeInternalNodesOnly returns the SystemIndexToPrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeInternalNodesOnlyOk

`func (o *BackendListResponseResourcesInner) GetSystemIndexToPrimeInternalNodesOnlyOk() (*[]EnumbackendSystemIndexToPrimeInternalNodesOnlyProp, bool)`

GetSystemIndexToPrimeInternalNodesOnlyOk returns a tuple with the SystemIndexToPrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrimeInternalNodesOnly

`func (o *BackendListResponseResourcesInner) SetSystemIndexToPrimeInternalNodesOnly(v []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp)`

SetSystemIndexToPrimeInternalNodesOnly sets SystemIndexToPrimeInternalNodesOnly field to given value.

### HasSystemIndexToPrimeInternalNodesOnly

`func (o *BackendListResponseResourcesInner) HasSystemIndexToPrimeInternalNodesOnly() bool`

HasSystemIndexToPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetBackgroundPrime

`func (o *BackendListResponseResourcesInner) GetBackgroundPrime() bool`

GetBackgroundPrime returns the BackgroundPrime field if non-nil, zero value otherwise.

### GetBackgroundPrimeOk

`func (o *BackendListResponseResourcesInner) GetBackgroundPrimeOk() (*bool, bool)`

GetBackgroundPrimeOk returns a tuple with the BackgroundPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPrime

`func (o *BackendListResponseResourcesInner) SetBackgroundPrime(v bool)`

SetBackgroundPrime sets BackgroundPrime field to given value.

### HasBackgroundPrime

`func (o *BackendListResponseResourcesInner) HasBackgroundPrime() bool`

HasBackgroundPrime returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *BackendListResponseResourcesInner) GetIndexEntryLimit() int64`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *BackendListResponseResourcesInner) GetIndexEntryLimitOk() (*int64, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *BackendListResponseResourcesInner) SetIndexEntryLimit(v int64)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *BackendListResponseResourcesInner) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetCompositeIndexEntryLimit

`func (o *BackendListResponseResourcesInner) GetCompositeIndexEntryLimit() int64`

GetCompositeIndexEntryLimit returns the CompositeIndexEntryLimit field if non-nil, zero value otherwise.

### GetCompositeIndexEntryLimitOk

`func (o *BackendListResponseResourcesInner) GetCompositeIndexEntryLimitOk() (*int64, bool)`

GetCompositeIndexEntryLimitOk returns a tuple with the CompositeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeIndexEntryLimit

`func (o *BackendListResponseResourcesInner) SetCompositeIndexEntryLimit(v int64)`

SetCompositeIndexEntryLimit sets CompositeIndexEntryLimit field to given value.

### HasCompositeIndexEntryLimit

`func (o *BackendListResponseResourcesInner) HasCompositeIndexEntryLimit() bool`

HasCompositeIndexEntryLimit returns a boolean if a field has been set.

### GetId2childrenIndexEntryLimit

`func (o *BackendListResponseResourcesInner) GetId2childrenIndexEntryLimit() int64`

GetId2childrenIndexEntryLimit returns the Id2childrenIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2childrenIndexEntryLimitOk

`func (o *BackendListResponseResourcesInner) GetId2childrenIndexEntryLimitOk() (*int64, bool)`

GetId2childrenIndexEntryLimitOk returns a tuple with the Id2childrenIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenIndexEntryLimit

`func (o *BackendListResponseResourcesInner) SetId2childrenIndexEntryLimit(v int64)`

SetId2childrenIndexEntryLimit sets Id2childrenIndexEntryLimit field to given value.

### HasId2childrenIndexEntryLimit

`func (o *BackendListResponseResourcesInner) HasId2childrenIndexEntryLimit() bool`

HasId2childrenIndexEntryLimit returns a boolean if a field has been set.

### GetId2subtreeIndexEntryLimit

`func (o *BackendListResponseResourcesInner) GetId2subtreeIndexEntryLimit() int64`

GetId2subtreeIndexEntryLimit returns the Id2subtreeIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2subtreeIndexEntryLimitOk

`func (o *BackendListResponseResourcesInner) GetId2subtreeIndexEntryLimitOk() (*int64, bool)`

GetId2subtreeIndexEntryLimitOk returns a tuple with the Id2subtreeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeIndexEntryLimit

`func (o *BackendListResponseResourcesInner) SetId2subtreeIndexEntryLimit(v int64)`

SetId2subtreeIndexEntryLimit sets Id2subtreeIndexEntryLimit field to given value.

### HasId2subtreeIndexEntryLimit

`func (o *BackendListResponseResourcesInner) HasId2subtreeIndexEntryLimit() bool`

HasId2subtreeIndexEntryLimit returns a boolean if a field has been set.

### GetImportTempDirectory

`func (o *BackendListResponseResourcesInner) GetImportTempDirectory() string`

GetImportTempDirectory returns the ImportTempDirectory field if non-nil, zero value otherwise.

### GetImportTempDirectoryOk

`func (o *BackendListResponseResourcesInner) GetImportTempDirectoryOk() (*string, bool)`

GetImportTempDirectoryOk returns a tuple with the ImportTempDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTempDirectory

`func (o *BackendListResponseResourcesInner) SetImportTempDirectory(v string)`

SetImportTempDirectory sets ImportTempDirectory field to given value.


### GetImportThreadCount

`func (o *BackendListResponseResourcesInner) GetImportThreadCount() int64`

GetImportThreadCount returns the ImportThreadCount field if non-nil, zero value otherwise.

### GetImportThreadCountOk

`func (o *BackendListResponseResourcesInner) GetImportThreadCountOk() (*int64, bool)`

GetImportThreadCountOk returns a tuple with the ImportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportThreadCount

`func (o *BackendListResponseResourcesInner) SetImportThreadCount(v int64)`

SetImportThreadCount sets ImportThreadCount field to given value.

### HasImportThreadCount

`func (o *BackendListResponseResourcesInner) HasImportThreadCount() bool`

HasImportThreadCount returns a boolean if a field has been set.

### GetExportThreadCount

`func (o *BackendListResponseResourcesInner) GetExportThreadCount() int64`

GetExportThreadCount returns the ExportThreadCount field if non-nil, zero value otherwise.

### GetExportThreadCountOk

`func (o *BackendListResponseResourcesInner) GetExportThreadCountOk() (*int64, bool)`

GetExportThreadCountOk returns a tuple with the ExportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportThreadCount

`func (o *BackendListResponseResourcesInner) SetExportThreadCount(v int64)`

SetExportThreadCount sets ExportThreadCount field to given value.

### HasExportThreadCount

`func (o *BackendListResponseResourcesInner) HasExportThreadCount() bool`

HasExportThreadCount returns a boolean if a field has been set.

### GetDbImportCachePercent

`func (o *BackendListResponseResourcesInner) GetDbImportCachePercent() int64`

GetDbImportCachePercent returns the DbImportCachePercent field if non-nil, zero value otherwise.

### GetDbImportCachePercentOk

`func (o *BackendListResponseResourcesInner) GetDbImportCachePercentOk() (*int64, bool)`

GetDbImportCachePercentOk returns a tuple with the DbImportCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbImportCachePercent

`func (o *BackendListResponseResourcesInner) SetDbImportCachePercent(v int64)`

SetDbImportCachePercent sets DbImportCachePercent field to given value.

### HasDbImportCachePercent

`func (o *BackendListResponseResourcesInner) HasDbImportCachePercent() bool`

HasDbImportCachePercent returns a boolean if a field has been set.

### GetDbTxnWriteNoSync

`func (o *BackendListResponseResourcesInner) GetDbTxnWriteNoSync() bool`

GetDbTxnWriteNoSync returns the DbTxnWriteNoSync field if non-nil, zero value otherwise.

### GetDbTxnWriteNoSyncOk

`func (o *BackendListResponseResourcesInner) GetDbTxnWriteNoSyncOk() (*bool, bool)`

GetDbTxnWriteNoSyncOk returns a tuple with the DbTxnWriteNoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTxnWriteNoSync

`func (o *BackendListResponseResourcesInner) SetDbTxnWriteNoSync(v bool)`

SetDbTxnWriteNoSync sets DbTxnWriteNoSync field to given value.

### HasDbTxnWriteNoSync

`func (o *BackendListResponseResourcesInner) HasDbTxnWriteNoSync() bool`

HasDbTxnWriteNoSync returns a boolean if a field has been set.

### GetDeadlockRetryLimit

`func (o *BackendListResponseResourcesInner) GetDeadlockRetryLimit() int64`

GetDeadlockRetryLimit returns the DeadlockRetryLimit field if non-nil, zero value otherwise.

### GetDeadlockRetryLimitOk

`func (o *BackendListResponseResourcesInner) GetDeadlockRetryLimitOk() (*int64, bool)`

GetDeadlockRetryLimitOk returns a tuple with the DeadlockRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockRetryLimit

`func (o *BackendListResponseResourcesInner) SetDeadlockRetryLimit(v int64)`

SetDeadlockRetryLimit sets DeadlockRetryLimit field to given value.

### HasDeadlockRetryLimit

`func (o *BackendListResponseResourcesInner) HasDeadlockRetryLimit() bool`

HasDeadlockRetryLimit returns a boolean if a field has been set.

### GetExternalTxnDefaultBackendLockBehavior

`func (o *BackendListResponseResourcesInner) GetExternalTxnDefaultBackendLockBehavior() EnumbackendExternalTxnDefaultBackendLockBehaviorProp`

GetExternalTxnDefaultBackendLockBehavior returns the ExternalTxnDefaultBackendLockBehavior field if non-nil, zero value otherwise.

### GetExternalTxnDefaultBackendLockBehaviorOk

`func (o *BackendListResponseResourcesInner) GetExternalTxnDefaultBackendLockBehaviorOk() (*EnumbackendExternalTxnDefaultBackendLockBehaviorProp, bool)`

GetExternalTxnDefaultBackendLockBehaviorOk returns a tuple with the ExternalTxnDefaultBackendLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTxnDefaultBackendLockBehavior

`func (o *BackendListResponseResourcesInner) SetExternalTxnDefaultBackendLockBehavior(v EnumbackendExternalTxnDefaultBackendLockBehaviorProp)`

SetExternalTxnDefaultBackendLockBehavior sets ExternalTxnDefaultBackendLockBehavior field to given value.

### HasExternalTxnDefaultBackendLockBehavior

`func (o *BackendListResponseResourcesInner) HasExternalTxnDefaultBackendLockBehavior() bool`

HasExternalTxnDefaultBackendLockBehavior returns a boolean if a field has been set.

### GetSingleWriterLockBehavior

`func (o *BackendListResponseResourcesInner) GetSingleWriterLockBehavior() EnumbackendSingleWriterLockBehaviorProp`

GetSingleWriterLockBehavior returns the SingleWriterLockBehavior field if non-nil, zero value otherwise.

### GetSingleWriterLockBehaviorOk

`func (o *BackendListResponseResourcesInner) GetSingleWriterLockBehaviorOk() (*EnumbackendSingleWriterLockBehaviorProp, bool)`

GetSingleWriterLockBehaviorOk returns a tuple with the SingleWriterLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleWriterLockBehavior

`func (o *BackendListResponseResourcesInner) SetSingleWriterLockBehavior(v EnumbackendSingleWriterLockBehaviorProp)`

SetSingleWriterLockBehavior sets SingleWriterLockBehavior field to given value.

### HasSingleWriterLockBehavior

`func (o *BackendListResponseResourcesInner) HasSingleWriterLockBehavior() bool`

HasSingleWriterLockBehavior returns a boolean if a field has been set.

### GetSubtreeDeleteSizeLimit

`func (o *BackendListResponseResourcesInner) GetSubtreeDeleteSizeLimit() int64`

GetSubtreeDeleteSizeLimit returns the SubtreeDeleteSizeLimit field if non-nil, zero value otherwise.

### GetSubtreeDeleteSizeLimitOk

`func (o *BackendListResponseResourcesInner) GetSubtreeDeleteSizeLimitOk() (*int64, bool)`

GetSubtreeDeleteSizeLimitOk returns a tuple with the SubtreeDeleteSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeDeleteSizeLimit

`func (o *BackendListResponseResourcesInner) SetSubtreeDeleteSizeLimit(v int64)`

SetSubtreeDeleteSizeLimit sets SubtreeDeleteSizeLimit field to given value.

### HasSubtreeDeleteSizeLimit

`func (o *BackendListResponseResourcesInner) HasSubtreeDeleteSizeLimit() bool`

HasSubtreeDeleteSizeLimit returns a boolean if a field has been set.

### GetNumRecentChanges

`func (o *BackendListResponseResourcesInner) GetNumRecentChanges() int64`

GetNumRecentChanges returns the NumRecentChanges field if non-nil, zero value otherwise.

### GetNumRecentChangesOk

`func (o *BackendListResponseResourcesInner) GetNumRecentChangesOk() (*int64, bool)`

GetNumRecentChangesOk returns a tuple with the NumRecentChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecentChanges

`func (o *BackendListResponseResourcesInner) SetNumRecentChanges(v int64)`

SetNumRecentChanges sets NumRecentChanges field to given value.

### HasNumRecentChanges

`func (o *BackendListResponseResourcesInner) HasNumRecentChanges() bool`

HasNumRecentChanges returns a boolean if a field has been set.

### GetOfflineProcessDatabaseOpenTimeout

`func (o *BackendListResponseResourcesInner) GetOfflineProcessDatabaseOpenTimeout() string`

GetOfflineProcessDatabaseOpenTimeout returns the OfflineProcessDatabaseOpenTimeout field if non-nil, zero value otherwise.

### GetOfflineProcessDatabaseOpenTimeoutOk

`func (o *BackendListResponseResourcesInner) GetOfflineProcessDatabaseOpenTimeoutOk() (*string, bool)`

GetOfflineProcessDatabaseOpenTimeoutOk returns a tuple with the OfflineProcessDatabaseOpenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineProcessDatabaseOpenTimeout

`func (o *BackendListResponseResourcesInner) SetOfflineProcessDatabaseOpenTimeout(v string)`

SetOfflineProcessDatabaseOpenTimeout sets OfflineProcessDatabaseOpenTimeout field to given value.

### HasOfflineProcessDatabaseOpenTimeout

`func (o *BackendListResponseResourcesInner) HasOfflineProcessDatabaseOpenTimeout() bool`

HasOfflineProcessDatabaseOpenTimeout returns a boolean if a field has been set.

### GetMirroredSubtreePeerPollingInterval

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreePeerPollingInterval() string`

GetMirroredSubtreePeerPollingInterval returns the MirroredSubtreePeerPollingInterval field if non-nil, zero value otherwise.

### GetMirroredSubtreePeerPollingIntervalOk

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreePeerPollingIntervalOk() (*string, bool)`

GetMirroredSubtreePeerPollingIntervalOk returns a tuple with the MirroredSubtreePeerPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreePeerPollingInterval

`func (o *BackendListResponseResourcesInner) SetMirroredSubtreePeerPollingInterval(v string)`

SetMirroredSubtreePeerPollingInterval sets MirroredSubtreePeerPollingInterval field to given value.

### HasMirroredSubtreePeerPollingInterval

`func (o *BackendListResponseResourcesInner) HasMirroredSubtreePeerPollingInterval() bool`

HasMirroredSubtreePeerPollingInterval returns a boolean if a field has been set.

### GetMirroredSubtreeEntryUpdateTimeout

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreeEntryUpdateTimeout() string`

GetMirroredSubtreeEntryUpdateTimeout returns the MirroredSubtreeEntryUpdateTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeEntryUpdateTimeoutOk

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreeEntryUpdateTimeoutOk() (*string, bool)`

GetMirroredSubtreeEntryUpdateTimeoutOk returns a tuple with the MirroredSubtreeEntryUpdateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeEntryUpdateTimeout

`func (o *BackendListResponseResourcesInner) SetMirroredSubtreeEntryUpdateTimeout(v string)`

SetMirroredSubtreeEntryUpdateTimeout sets MirroredSubtreeEntryUpdateTimeout field to given value.

### HasMirroredSubtreeEntryUpdateTimeout

`func (o *BackendListResponseResourcesInner) HasMirroredSubtreeEntryUpdateTimeout() bool`

HasMirroredSubtreeEntryUpdateTimeout returns a boolean if a field has been set.

### GetMirroredSubtreeSearchTimeout

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreeSearchTimeout() string`

GetMirroredSubtreeSearchTimeout returns the MirroredSubtreeSearchTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeSearchTimeoutOk

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreeSearchTimeoutOk() (*string, bool)`

GetMirroredSubtreeSearchTimeoutOk returns a tuple with the MirroredSubtreeSearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeSearchTimeout

`func (o *BackendListResponseResourcesInner) SetMirroredSubtreeSearchTimeout(v string)`

SetMirroredSubtreeSearchTimeout sets MirroredSubtreeSearchTimeout field to given value.

### HasMirroredSubtreeSearchTimeout

`func (o *BackendListResponseResourcesInner) HasMirroredSubtreeSearchTimeout() bool`

HasMirroredSubtreeSearchTimeout returns a boolean if a field has been set.

### GetPeerServer

`func (o *BackendListResponseResourcesInner) GetPeerServer() []string`

GetPeerServer returns the PeerServer field if non-nil, zero value otherwise.

### GetPeerServerOk

`func (o *BackendListResponseResourcesInner) GetPeerServerOk() (*[]string, bool)`

GetPeerServerOk returns a tuple with the PeerServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServer

`func (o *BackendListResponseResourcesInner) SetPeerServer(v []string)`

SetPeerServer sets PeerServer field to given value.

### HasPeerServer

`func (o *BackendListResponseResourcesInner) HasPeerServer() bool`

HasPeerServer returns a boolean if a field has been set.

### GetMirroredSubtreePreferredMasterType

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreePreferredMasterType() []EnumbackendMirroredSubtreePreferredMasterTypeProp`

GetMirroredSubtreePreferredMasterType returns the MirroredSubtreePreferredMasterType field if non-nil, zero value otherwise.

### GetMirroredSubtreePreferredMasterTypeOk

`func (o *BackendListResponseResourcesInner) GetMirroredSubtreePreferredMasterTypeOk() (*[]EnumbackendMirroredSubtreePreferredMasterTypeProp, bool)`

GetMirroredSubtreePreferredMasterTypeOk returns a tuple with the MirroredSubtreePreferredMasterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreePreferredMasterType

`func (o *BackendListResponseResourcesInner) SetMirroredSubtreePreferredMasterType(v []EnumbackendMirroredSubtreePreferredMasterTypeProp)`

SetMirroredSubtreePreferredMasterType sets MirroredSubtreePreferredMasterType field to given value.

### HasMirroredSubtreePreferredMasterType

`func (o *BackendListResponseResourcesInner) HasMirroredSubtreePreferredMasterType() bool`

HasMirroredSubtreePreferredMasterType returns a boolean if a field has been set.

### GetSimulatedResultCode

`func (o *BackendListResponseResourcesInner) GetSimulatedResultCode() int64`

GetSimulatedResultCode returns the SimulatedResultCode field if non-nil, zero value otherwise.

### GetSimulatedResultCodeOk

`func (o *BackendListResponseResourcesInner) GetSimulatedResultCodeOk() (*int64, bool)`

GetSimulatedResultCodeOk returns a tuple with the SimulatedResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatedResultCode

`func (o *BackendListResponseResourcesInner) SetSimulatedResultCode(v int64)`

SetSimulatedResultCode sets SimulatedResultCode field to given value.

### HasSimulatedResultCode

`func (o *BackendListResponseResourcesInner) HasSimulatedResultCode() bool`

HasSimulatedResultCode returns a boolean if a field has been set.

### GetInsignificantConfigArchiveAttribute

`func (o *BackendListResponseResourcesInner) GetInsignificantConfigArchiveAttribute() []string`

GetInsignificantConfigArchiveAttribute returns the InsignificantConfigArchiveAttribute field if non-nil, zero value otherwise.

### GetInsignificantConfigArchiveAttributeOk

`func (o *BackendListResponseResourcesInner) GetInsignificantConfigArchiveAttributeOk() (*[]string, bool)`

GetInsignificantConfigArchiveAttributeOk returns a tuple with the InsignificantConfigArchiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsignificantConfigArchiveAttribute

`func (o *BackendListResponseResourcesInner) SetInsignificantConfigArchiveAttribute(v []string)`

SetInsignificantConfigArchiveAttribute sets InsignificantConfigArchiveAttribute field to given value.

### HasInsignificantConfigArchiveAttribute

`func (o *BackendListResponseResourcesInner) HasInsignificantConfigArchiveAttribute() bool`

HasInsignificantConfigArchiveAttribute returns a boolean if a field has been set.

### GetInsignificantConfigArchiveBaseDN

`func (o *BackendListResponseResourcesInner) GetInsignificantConfigArchiveBaseDN() []string`

GetInsignificantConfigArchiveBaseDN returns the InsignificantConfigArchiveBaseDN field if non-nil, zero value otherwise.

### GetInsignificantConfigArchiveBaseDNOk

`func (o *BackendListResponseResourcesInner) GetInsignificantConfigArchiveBaseDNOk() (*[]string, bool)`

GetInsignificantConfigArchiveBaseDNOk returns a tuple with the InsignificantConfigArchiveBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsignificantConfigArchiveBaseDN

`func (o *BackendListResponseResourcesInner) SetInsignificantConfigArchiveBaseDN(v []string)`

SetInsignificantConfigArchiveBaseDN sets InsignificantConfigArchiveBaseDN field to given value.

### HasInsignificantConfigArchiveBaseDN

`func (o *BackendListResponseResourcesInner) HasInsignificantConfigArchiveBaseDN() bool`

HasInsignificantConfigArchiveBaseDN returns a boolean if a field has been set.

### GetMaintainConfigArchive

`func (o *BackendListResponseResourcesInner) GetMaintainConfigArchive() bool`

GetMaintainConfigArchive returns the MaintainConfigArchive field if non-nil, zero value otherwise.

### GetMaintainConfigArchiveOk

`func (o *BackendListResponseResourcesInner) GetMaintainConfigArchiveOk() (*bool, bool)`

GetMaintainConfigArchiveOk returns a tuple with the MaintainConfigArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainConfigArchive

`func (o *BackendListResponseResourcesInner) SetMaintainConfigArchive(v bool)`

SetMaintainConfigArchive sets MaintainConfigArchive field to given value.

### HasMaintainConfigArchive

`func (o *BackendListResponseResourcesInner) HasMaintainConfigArchive() bool`

HasMaintainConfigArchive returns a boolean if a field has been set.

### GetMaxConfigArchiveCount

`func (o *BackendListResponseResourcesInner) GetMaxConfigArchiveCount() int64`

GetMaxConfigArchiveCount returns the MaxConfigArchiveCount field if non-nil, zero value otherwise.

### GetMaxConfigArchiveCountOk

`func (o *BackendListResponseResourcesInner) GetMaxConfigArchiveCountOk() (*int64, bool)`

GetMaxConfigArchiveCountOk returns a tuple with the MaxConfigArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConfigArchiveCount

`func (o *BackendListResponseResourcesInner) SetMaxConfigArchiveCount(v int64)`

SetMaxConfigArchiveCount sets MaxConfigArchiveCount field to given value.

### HasMaxConfigArchiveCount

`func (o *BackendListResponseResourcesInner) HasMaxConfigArchiveCount() bool`

HasMaxConfigArchiveCount returns a boolean if a field has been set.

### GetTaskBackingFile

`func (o *BackendListResponseResourcesInner) GetTaskBackingFile() string`

GetTaskBackingFile returns the TaskBackingFile field if non-nil, zero value otherwise.

### GetTaskBackingFileOk

`func (o *BackendListResponseResourcesInner) GetTaskBackingFileOk() (*string, bool)`

GetTaskBackingFileOk returns a tuple with the TaskBackingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskBackingFile

`func (o *BackendListResponseResourcesInner) SetTaskBackingFile(v string)`

SetTaskBackingFile sets TaskBackingFile field to given value.


### GetMaximumInitialTaskLogMessagesToRetain

`func (o *BackendListResponseResourcesInner) GetMaximumInitialTaskLogMessagesToRetain() int64`

GetMaximumInitialTaskLogMessagesToRetain returns the MaximumInitialTaskLogMessagesToRetain field if non-nil, zero value otherwise.

### GetMaximumInitialTaskLogMessagesToRetainOk

`func (o *BackendListResponseResourcesInner) GetMaximumInitialTaskLogMessagesToRetainOk() (*int64, bool)`

GetMaximumInitialTaskLogMessagesToRetainOk returns a tuple with the MaximumInitialTaskLogMessagesToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInitialTaskLogMessagesToRetain

`func (o *BackendListResponseResourcesInner) SetMaximumInitialTaskLogMessagesToRetain(v int64)`

SetMaximumInitialTaskLogMessagesToRetain sets MaximumInitialTaskLogMessagesToRetain field to given value.

### HasMaximumInitialTaskLogMessagesToRetain

`func (o *BackendListResponseResourcesInner) HasMaximumInitialTaskLogMessagesToRetain() bool`

HasMaximumInitialTaskLogMessagesToRetain returns a boolean if a field has been set.

### GetMaximumFinalTaskLogMessagesToRetain

`func (o *BackendListResponseResourcesInner) GetMaximumFinalTaskLogMessagesToRetain() int64`

GetMaximumFinalTaskLogMessagesToRetain returns the MaximumFinalTaskLogMessagesToRetain field if non-nil, zero value otherwise.

### GetMaximumFinalTaskLogMessagesToRetainOk

`func (o *BackendListResponseResourcesInner) GetMaximumFinalTaskLogMessagesToRetainOk() (*int64, bool)`

GetMaximumFinalTaskLogMessagesToRetainOk returns a tuple with the MaximumFinalTaskLogMessagesToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFinalTaskLogMessagesToRetain

`func (o *BackendListResponseResourcesInner) SetMaximumFinalTaskLogMessagesToRetain(v int64)`

SetMaximumFinalTaskLogMessagesToRetain sets MaximumFinalTaskLogMessagesToRetain field to given value.

### HasMaximumFinalTaskLogMessagesToRetain

`func (o *BackendListResponseResourcesInner) HasMaximumFinalTaskLogMessagesToRetain() bool`

HasMaximumFinalTaskLogMessagesToRetain returns a boolean if a field has been set.

### GetTaskRetentionTime

`func (o *BackendListResponseResourcesInner) GetTaskRetentionTime() string`

GetTaskRetentionTime returns the TaskRetentionTime field if non-nil, zero value otherwise.

### GetTaskRetentionTimeOk

`func (o *BackendListResponseResourcesInner) GetTaskRetentionTimeOk() (*string, bool)`

GetTaskRetentionTimeOk returns a tuple with the TaskRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRetentionTime

`func (o *BackendListResponseResourcesInner) SetTaskRetentionTime(v string)`

SetTaskRetentionTime sets TaskRetentionTime field to given value.

### HasTaskRetentionTime

`func (o *BackendListResponseResourcesInner) HasTaskRetentionTime() bool`

HasTaskRetentionTime returns a boolean if a field has been set.

### GetNotificationSenderAddress

`func (o *BackendListResponseResourcesInner) GetNotificationSenderAddress() string`

GetNotificationSenderAddress returns the NotificationSenderAddress field if non-nil, zero value otherwise.

### GetNotificationSenderAddressOk

`func (o *BackendListResponseResourcesInner) GetNotificationSenderAddressOk() (*string, bool)`

GetNotificationSenderAddressOk returns a tuple with the NotificationSenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSenderAddress

`func (o *BackendListResponseResourcesInner) SetNotificationSenderAddress(v string)`

SetNotificationSenderAddress sets NotificationSenderAddress field to given value.

### HasNotificationSenderAddress

`func (o *BackendListResponseResourcesInner) HasNotificationSenderAddress() bool`

HasNotificationSenderAddress returns a boolean if a field has been set.

### GetAlertRetentionTime

`func (o *BackendListResponseResourcesInner) GetAlertRetentionTime() string`

GetAlertRetentionTime returns the AlertRetentionTime field if non-nil, zero value otherwise.

### GetAlertRetentionTimeOk

`func (o *BackendListResponseResourcesInner) GetAlertRetentionTimeOk() (*string, bool)`

GetAlertRetentionTimeOk returns a tuple with the AlertRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRetentionTime

`func (o *BackendListResponseResourcesInner) SetAlertRetentionTime(v string)`

SetAlertRetentionTime sets AlertRetentionTime field to given value.


### GetMaxAlerts

`func (o *BackendListResponseResourcesInner) GetMaxAlerts() int64`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *BackendListResponseResourcesInner) GetMaxAlertsOk() (*int64, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *BackendListResponseResourcesInner) SetMaxAlerts(v int64)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *BackendListResponseResourcesInner) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *BackendListResponseResourcesInner) GetDisabledAlertType() []EnumbackendDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *BackendListResponseResourcesInner) GetDisabledAlertTypeOk() (*[]EnumbackendDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *BackendListResponseResourcesInner) SetDisabledAlertType(v []EnumbackendDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *BackendListResponseResourcesInner) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.

### GetAlarmRetentionTime

`func (o *BackendListResponseResourcesInner) GetAlarmRetentionTime() string`

GetAlarmRetentionTime returns the AlarmRetentionTime field if non-nil, zero value otherwise.

### GetAlarmRetentionTimeOk

`func (o *BackendListResponseResourcesInner) GetAlarmRetentionTimeOk() (*string, bool)`

GetAlarmRetentionTimeOk returns a tuple with the AlarmRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRetentionTime

`func (o *BackendListResponseResourcesInner) SetAlarmRetentionTime(v string)`

SetAlarmRetentionTime sets AlarmRetentionTime field to given value.


### GetMaxAlarms

`func (o *BackendListResponseResourcesInner) GetMaxAlarms() int64`

GetMaxAlarms returns the MaxAlarms field if non-nil, zero value otherwise.

### GetMaxAlarmsOk

`func (o *BackendListResponseResourcesInner) GetMaxAlarmsOk() (*int64, bool)`

GetMaxAlarmsOk returns a tuple with the MaxAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlarms

`func (o *BackendListResponseResourcesInner) SetMaxAlarms(v int64)`

SetMaxAlarms sets MaxAlarms field to given value.

### HasMaxAlarms

`func (o *BackendListResponseResourcesInner) HasMaxAlarms() bool`

HasMaxAlarms returns a boolean if a field has been set.

### GetStorageDir

`func (o *BackendListResponseResourcesInner) GetStorageDir() string`

GetStorageDir returns the StorageDir field if non-nil, zero value otherwise.

### GetStorageDirOk

`func (o *BackendListResponseResourcesInner) GetStorageDirOk() (*string, bool)`

GetStorageDirOk returns a tuple with the StorageDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDir

`func (o *BackendListResponseResourcesInner) SetStorageDir(v string)`

SetStorageDir sets StorageDir field to given value.


### GetMetricsDir

`func (o *BackendListResponseResourcesInner) GetMetricsDir() string`

GetMetricsDir returns the MetricsDir field if non-nil, zero value otherwise.

### GetMetricsDirOk

`func (o *BackendListResponseResourcesInner) GetMetricsDirOk() (*string, bool)`

GetMetricsDirOk returns a tuple with the MetricsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsDir

`func (o *BackendListResponseResourcesInner) SetMetricsDir(v string)`

SetMetricsDir sets MetricsDir field to given value.


### GetSampleFlushInterval

`func (o *BackendListResponseResourcesInner) GetSampleFlushInterval() string`

GetSampleFlushInterval returns the SampleFlushInterval field if non-nil, zero value otherwise.

### GetSampleFlushIntervalOk

`func (o *BackendListResponseResourcesInner) GetSampleFlushIntervalOk() (*string, bool)`

GetSampleFlushIntervalOk returns a tuple with the SampleFlushInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFlushInterval

`func (o *BackendListResponseResourcesInner) SetSampleFlushInterval(v string)`

SetSampleFlushInterval sets SampleFlushInterval field to given value.

### HasSampleFlushInterval

`func (o *BackendListResponseResourcesInner) HasSampleFlushInterval() bool`

HasSampleFlushInterval returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *BackendListResponseResourcesInner) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *BackendListResponseResourcesInner) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *BackendListResponseResourcesInner) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


