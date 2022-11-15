# LocalDbBackendShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlocalDbBackendSchemaUrn**](EnumlocalDbBackendSchemaUrn.md) |  | 
**UncachedId2entryCacheMode** | Pointer to [**EnumbackendUncachedId2entryCacheModeProp**](EnumbackendUncachedId2entryCacheModeProp.md) |  | [optional] 
**UncachedAttributeCriteria** | Pointer to **string** | The criteria that will be used to identify attributes that should be written into the uncached-id2entry database rather than the id2entry database. This will only be used for entries in which the associated uncached-entry-criteria does not indicate that the entire entry should be uncached. | [optional] 
**UncachedEntryCriteria** | Pointer to **string** | The criteria that will be used to identify entries that should be written into the uncached-id2entry database rather than the id2entry database. | [optional] 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**SetDegradedAlertForUntrustedIndex** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state when this Local DB Backend has an index whose contents cannot be trusted. | [optional] 
**ReturnUnavailableForUntrustedIndex** | Pointer to **bool** | Determines whether the Directory Server returns UNAVAILABLE for any LDAP search operation in this Local DB Backend that would use an index whose contents cannot be trusted. | [optional] 
**ProcessFiltersWithUndefinedAttributeTypes** | Pointer to **bool** | Determines whether the Directory Server should continue filter processing for LDAP search operations in this Local DB Backend that includes a search filter with an attribute that is not defined in the schema. This will only apply if check-schema is enabled in the global configuration. | [optional] 
**IsPrivateBackend** | Pointer to **bool** | Indicates whether this backend should be considered a private backend in the server. Private backends are meant for storing server-internal information and should not be used for user or application data. | [optional] 
**DbDirectory** | **string** | Specifies the path to the filesystem directory that is used to hold the Berkeley DB Java Edition database files containing the data for this backend. The files for this backend are stored in a sub-directory named after the backend-id. | 
**DbDirectoryPermissions** | Pointer to **string** | Specifies the permissions that should be applied to the directory containing the backend database files and to directories and files created during backup or LDIF export of the backend. | [optional] 
**CompactCommonParentDN** | Pointer to **[]string** |  | [optional] 
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
**JeProperty** | Pointer to **[]string** |  | [optional] 
**DbCachePercent** | Pointer to **int32** | Specifies the percentage of JVM memory to allocate to the database cache. | [optional] 
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
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**BaseDN** | **[]string** |  | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 

## Methods

### NewLocalDbBackendShared

`func NewLocalDbBackendShared(schemas []EnumlocalDbBackendSchemaUrn, writabilityMode EnumbackendWritabilityModeProp, dbDirectory string, importTempDirectory string, backendID string, enabled bool, baseDN []string, ) *LocalDbBackendShared`

NewLocalDbBackendShared instantiates a new LocalDbBackendShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDbBackendSharedWithDefaults

`func NewLocalDbBackendSharedWithDefaults() *LocalDbBackendShared`

NewLocalDbBackendSharedWithDefaults instantiates a new LocalDbBackendShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LocalDbBackendShared) GetSchemas() []EnumlocalDbBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LocalDbBackendShared) GetSchemasOk() (*[]EnumlocalDbBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LocalDbBackendShared) SetSchemas(v []EnumlocalDbBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetUncachedId2entryCacheMode

`func (o *LocalDbBackendShared) GetUncachedId2entryCacheMode() EnumbackendUncachedId2entryCacheModeProp`

GetUncachedId2entryCacheMode returns the UncachedId2entryCacheMode field if non-nil, zero value otherwise.

### GetUncachedId2entryCacheModeOk

`func (o *LocalDbBackendShared) GetUncachedId2entryCacheModeOk() (*EnumbackendUncachedId2entryCacheModeProp, bool)`

GetUncachedId2entryCacheModeOk returns a tuple with the UncachedId2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedId2entryCacheMode

`func (o *LocalDbBackendShared) SetUncachedId2entryCacheMode(v EnumbackendUncachedId2entryCacheModeProp)`

SetUncachedId2entryCacheMode sets UncachedId2entryCacheMode field to given value.

### HasUncachedId2entryCacheMode

`func (o *LocalDbBackendShared) HasUncachedId2entryCacheMode() bool`

HasUncachedId2entryCacheMode returns a boolean if a field has been set.

### GetUncachedAttributeCriteria

`func (o *LocalDbBackendShared) GetUncachedAttributeCriteria() string`

GetUncachedAttributeCriteria returns the UncachedAttributeCriteria field if non-nil, zero value otherwise.

### GetUncachedAttributeCriteriaOk

`func (o *LocalDbBackendShared) GetUncachedAttributeCriteriaOk() (*string, bool)`

GetUncachedAttributeCriteriaOk returns a tuple with the UncachedAttributeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedAttributeCriteria

`func (o *LocalDbBackendShared) SetUncachedAttributeCriteria(v string)`

SetUncachedAttributeCriteria sets UncachedAttributeCriteria field to given value.

### HasUncachedAttributeCriteria

`func (o *LocalDbBackendShared) HasUncachedAttributeCriteria() bool`

HasUncachedAttributeCriteria returns a boolean if a field has been set.

### GetUncachedEntryCriteria

`func (o *LocalDbBackendShared) GetUncachedEntryCriteria() string`

GetUncachedEntryCriteria returns the UncachedEntryCriteria field if non-nil, zero value otherwise.

### GetUncachedEntryCriteriaOk

`func (o *LocalDbBackendShared) GetUncachedEntryCriteriaOk() (*string, bool)`

GetUncachedEntryCriteriaOk returns a tuple with the UncachedEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedEntryCriteria

`func (o *LocalDbBackendShared) SetUncachedEntryCriteria(v string)`

SetUncachedEntryCriteria sets UncachedEntryCriteria field to given value.

### HasUncachedEntryCriteria

`func (o *LocalDbBackendShared) HasUncachedEntryCriteria() bool`

HasUncachedEntryCriteria returns a boolean if a field has been set.

### GetWritabilityMode

`func (o *LocalDbBackendShared) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *LocalDbBackendShared) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *LocalDbBackendShared) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetSetDegradedAlertForUntrustedIndex

`func (o *LocalDbBackendShared) GetSetDegradedAlertForUntrustedIndex() bool`

GetSetDegradedAlertForUntrustedIndex returns the SetDegradedAlertForUntrustedIndex field if non-nil, zero value otherwise.

### GetSetDegradedAlertForUntrustedIndexOk

`func (o *LocalDbBackendShared) GetSetDegradedAlertForUntrustedIndexOk() (*bool, bool)`

GetSetDegradedAlertForUntrustedIndexOk returns a tuple with the SetDegradedAlertForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertForUntrustedIndex

`func (o *LocalDbBackendShared) SetSetDegradedAlertForUntrustedIndex(v bool)`

SetSetDegradedAlertForUntrustedIndex sets SetDegradedAlertForUntrustedIndex field to given value.

### HasSetDegradedAlertForUntrustedIndex

`func (o *LocalDbBackendShared) HasSetDegradedAlertForUntrustedIndex() bool`

HasSetDegradedAlertForUntrustedIndex returns a boolean if a field has been set.

### GetReturnUnavailableForUntrustedIndex

`func (o *LocalDbBackendShared) GetReturnUnavailableForUntrustedIndex() bool`

GetReturnUnavailableForUntrustedIndex returns the ReturnUnavailableForUntrustedIndex field if non-nil, zero value otherwise.

### GetReturnUnavailableForUntrustedIndexOk

`func (o *LocalDbBackendShared) GetReturnUnavailableForUntrustedIndexOk() (*bool, bool)`

GetReturnUnavailableForUntrustedIndexOk returns a tuple with the ReturnUnavailableForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableForUntrustedIndex

`func (o *LocalDbBackendShared) SetReturnUnavailableForUntrustedIndex(v bool)`

SetReturnUnavailableForUntrustedIndex sets ReturnUnavailableForUntrustedIndex field to given value.

### HasReturnUnavailableForUntrustedIndex

`func (o *LocalDbBackendShared) HasReturnUnavailableForUntrustedIndex() bool`

HasReturnUnavailableForUntrustedIndex returns a boolean if a field has been set.

### GetProcessFiltersWithUndefinedAttributeTypes

`func (o *LocalDbBackendShared) GetProcessFiltersWithUndefinedAttributeTypes() bool`

GetProcessFiltersWithUndefinedAttributeTypes returns the ProcessFiltersWithUndefinedAttributeTypes field if non-nil, zero value otherwise.

### GetProcessFiltersWithUndefinedAttributeTypesOk

`func (o *LocalDbBackendShared) GetProcessFiltersWithUndefinedAttributeTypesOk() (*bool, bool)`

GetProcessFiltersWithUndefinedAttributeTypesOk returns a tuple with the ProcessFiltersWithUndefinedAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFiltersWithUndefinedAttributeTypes

`func (o *LocalDbBackendShared) SetProcessFiltersWithUndefinedAttributeTypes(v bool)`

SetProcessFiltersWithUndefinedAttributeTypes sets ProcessFiltersWithUndefinedAttributeTypes field to given value.

### HasProcessFiltersWithUndefinedAttributeTypes

`func (o *LocalDbBackendShared) HasProcessFiltersWithUndefinedAttributeTypes() bool`

HasProcessFiltersWithUndefinedAttributeTypes returns a boolean if a field has been set.

### GetIsPrivateBackend

`func (o *LocalDbBackendShared) GetIsPrivateBackend() bool`

GetIsPrivateBackend returns the IsPrivateBackend field if non-nil, zero value otherwise.

### GetIsPrivateBackendOk

`func (o *LocalDbBackendShared) GetIsPrivateBackendOk() (*bool, bool)`

GetIsPrivateBackendOk returns a tuple with the IsPrivateBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateBackend

`func (o *LocalDbBackendShared) SetIsPrivateBackend(v bool)`

SetIsPrivateBackend sets IsPrivateBackend field to given value.

### HasIsPrivateBackend

`func (o *LocalDbBackendShared) HasIsPrivateBackend() bool`

HasIsPrivateBackend returns a boolean if a field has been set.

### GetDbDirectory

`func (o *LocalDbBackendShared) GetDbDirectory() string`

GetDbDirectory returns the DbDirectory field if non-nil, zero value otherwise.

### GetDbDirectoryOk

`func (o *LocalDbBackendShared) GetDbDirectoryOk() (*string, bool)`

GetDbDirectoryOk returns a tuple with the DbDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectory

`func (o *LocalDbBackendShared) SetDbDirectory(v string)`

SetDbDirectory sets DbDirectory field to given value.


### GetDbDirectoryPermissions

`func (o *LocalDbBackendShared) GetDbDirectoryPermissions() string`

GetDbDirectoryPermissions returns the DbDirectoryPermissions field if non-nil, zero value otherwise.

### GetDbDirectoryPermissionsOk

`func (o *LocalDbBackendShared) GetDbDirectoryPermissionsOk() (*string, bool)`

GetDbDirectoryPermissionsOk returns a tuple with the DbDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectoryPermissions

`func (o *LocalDbBackendShared) SetDbDirectoryPermissions(v string)`

SetDbDirectoryPermissions sets DbDirectoryPermissions field to given value.

### HasDbDirectoryPermissions

`func (o *LocalDbBackendShared) HasDbDirectoryPermissions() bool`

HasDbDirectoryPermissions returns a boolean if a field has been set.

### GetCompactCommonParentDN

`func (o *LocalDbBackendShared) GetCompactCommonParentDN() []string`

GetCompactCommonParentDN returns the CompactCommonParentDN field if non-nil, zero value otherwise.

### GetCompactCommonParentDNOk

`func (o *LocalDbBackendShared) GetCompactCommonParentDNOk() (*[]string, bool)`

GetCompactCommonParentDNOk returns a tuple with the CompactCommonParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactCommonParentDN

`func (o *LocalDbBackendShared) SetCompactCommonParentDN(v []string)`

SetCompactCommonParentDN sets CompactCommonParentDN field to given value.

### HasCompactCommonParentDN

`func (o *LocalDbBackendShared) HasCompactCommonParentDN() bool`

HasCompactCommonParentDN returns a boolean if a field has been set.

### GetCompressEntries

`func (o *LocalDbBackendShared) GetCompressEntries() bool`

GetCompressEntries returns the CompressEntries field if non-nil, zero value otherwise.

### GetCompressEntriesOk

`func (o *LocalDbBackendShared) GetCompressEntriesOk() (*bool, bool)`

GetCompressEntriesOk returns a tuple with the CompressEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressEntries

`func (o *LocalDbBackendShared) SetCompressEntries(v bool)`

SetCompressEntries sets CompressEntries field to given value.

### HasCompressEntries

`func (o *LocalDbBackendShared) HasCompressEntries() bool`

HasCompressEntries returns a boolean if a field has been set.

### GetHashEntries

`func (o *LocalDbBackendShared) GetHashEntries() bool`

GetHashEntries returns the HashEntries field if non-nil, zero value otherwise.

### GetHashEntriesOk

`func (o *LocalDbBackendShared) GetHashEntriesOk() (*bool, bool)`

GetHashEntriesOk returns a tuple with the HashEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashEntries

`func (o *LocalDbBackendShared) SetHashEntries(v bool)`

SetHashEntries sets HashEntries field to given value.

### HasHashEntries

`func (o *LocalDbBackendShared) HasHashEntries() bool`

HasHashEntries returns a boolean if a field has been set.

### GetDbNumCleanerThreads

`func (o *LocalDbBackendShared) GetDbNumCleanerThreads() int32`

GetDbNumCleanerThreads returns the DbNumCleanerThreads field if non-nil, zero value otherwise.

### GetDbNumCleanerThreadsOk

`func (o *LocalDbBackendShared) GetDbNumCleanerThreadsOk() (*int32, bool)`

GetDbNumCleanerThreadsOk returns a tuple with the DbNumCleanerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbNumCleanerThreads

`func (o *LocalDbBackendShared) SetDbNumCleanerThreads(v int32)`

SetDbNumCleanerThreads sets DbNumCleanerThreads field to given value.

### HasDbNumCleanerThreads

`func (o *LocalDbBackendShared) HasDbNumCleanerThreads() bool`

HasDbNumCleanerThreads returns a boolean if a field has been set.

### GetDbCleanerMinUtilization

`func (o *LocalDbBackendShared) GetDbCleanerMinUtilization() int32`

GetDbCleanerMinUtilization returns the DbCleanerMinUtilization field if non-nil, zero value otherwise.

### GetDbCleanerMinUtilizationOk

`func (o *LocalDbBackendShared) GetDbCleanerMinUtilizationOk() (*int32, bool)`

GetDbCleanerMinUtilizationOk returns a tuple with the DbCleanerMinUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCleanerMinUtilization

`func (o *LocalDbBackendShared) SetDbCleanerMinUtilization(v int32)`

SetDbCleanerMinUtilization sets DbCleanerMinUtilization field to given value.

### HasDbCleanerMinUtilization

`func (o *LocalDbBackendShared) HasDbCleanerMinUtilization() bool`

HasDbCleanerMinUtilization returns a boolean if a field has been set.

### GetDbEvictorCriticalPercentage

`func (o *LocalDbBackendShared) GetDbEvictorCriticalPercentage() int32`

GetDbEvictorCriticalPercentage returns the DbEvictorCriticalPercentage field if non-nil, zero value otherwise.

### GetDbEvictorCriticalPercentageOk

`func (o *LocalDbBackendShared) GetDbEvictorCriticalPercentageOk() (*int32, bool)`

GetDbEvictorCriticalPercentageOk returns a tuple with the DbEvictorCriticalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEvictorCriticalPercentage

`func (o *LocalDbBackendShared) SetDbEvictorCriticalPercentage(v int32)`

SetDbEvictorCriticalPercentage sets DbEvictorCriticalPercentage field to given value.

### HasDbEvictorCriticalPercentage

`func (o *LocalDbBackendShared) HasDbEvictorCriticalPercentage() bool`

HasDbEvictorCriticalPercentage returns a boolean if a field has been set.

### GetDbCheckpointerWakeupInterval

`func (o *LocalDbBackendShared) GetDbCheckpointerWakeupInterval() string`

GetDbCheckpointerWakeupInterval returns the DbCheckpointerWakeupInterval field if non-nil, zero value otherwise.

### GetDbCheckpointerWakeupIntervalOk

`func (o *LocalDbBackendShared) GetDbCheckpointerWakeupIntervalOk() (*string, bool)`

GetDbCheckpointerWakeupIntervalOk returns a tuple with the DbCheckpointerWakeupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCheckpointerWakeupInterval

`func (o *LocalDbBackendShared) SetDbCheckpointerWakeupInterval(v string)`

SetDbCheckpointerWakeupInterval sets DbCheckpointerWakeupInterval field to given value.

### HasDbCheckpointerWakeupInterval

`func (o *LocalDbBackendShared) HasDbCheckpointerWakeupInterval() bool`

HasDbCheckpointerWakeupInterval returns a boolean if a field has been set.

### GetDbBackgroundSyncInterval

`func (o *LocalDbBackendShared) GetDbBackgroundSyncInterval() string`

GetDbBackgroundSyncInterval returns the DbBackgroundSyncInterval field if non-nil, zero value otherwise.

### GetDbBackgroundSyncIntervalOk

`func (o *LocalDbBackendShared) GetDbBackgroundSyncIntervalOk() (*string, bool)`

GetDbBackgroundSyncIntervalOk returns a tuple with the DbBackgroundSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbBackgroundSyncInterval

`func (o *LocalDbBackendShared) SetDbBackgroundSyncInterval(v string)`

SetDbBackgroundSyncInterval sets DbBackgroundSyncInterval field to given value.

### HasDbBackgroundSyncInterval

`func (o *LocalDbBackendShared) HasDbBackgroundSyncInterval() bool`

HasDbBackgroundSyncInterval returns a boolean if a field has been set.

### GetDbUseThreadLocalHandles

`func (o *LocalDbBackendShared) GetDbUseThreadLocalHandles() bool`

GetDbUseThreadLocalHandles returns the DbUseThreadLocalHandles field if non-nil, zero value otherwise.

### GetDbUseThreadLocalHandlesOk

`func (o *LocalDbBackendShared) GetDbUseThreadLocalHandlesOk() (*bool, bool)`

GetDbUseThreadLocalHandlesOk returns a tuple with the DbUseThreadLocalHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUseThreadLocalHandles

`func (o *LocalDbBackendShared) SetDbUseThreadLocalHandles(v bool)`

SetDbUseThreadLocalHandles sets DbUseThreadLocalHandles field to given value.

### HasDbUseThreadLocalHandles

`func (o *LocalDbBackendShared) HasDbUseThreadLocalHandles() bool`

HasDbUseThreadLocalHandles returns a boolean if a field has been set.

### GetDbLogFileMax

`func (o *LocalDbBackendShared) GetDbLogFileMax() string`

GetDbLogFileMax returns the DbLogFileMax field if non-nil, zero value otherwise.

### GetDbLogFileMaxOk

`func (o *LocalDbBackendShared) GetDbLogFileMaxOk() (*string, bool)`

GetDbLogFileMaxOk returns a tuple with the DbLogFileMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLogFileMax

`func (o *LocalDbBackendShared) SetDbLogFileMax(v string)`

SetDbLogFileMax sets DbLogFileMax field to given value.

### HasDbLogFileMax

`func (o *LocalDbBackendShared) HasDbLogFileMax() bool`

HasDbLogFileMax returns a boolean if a field has been set.

### GetDbLoggingLevel

`func (o *LocalDbBackendShared) GetDbLoggingLevel() string`

GetDbLoggingLevel returns the DbLoggingLevel field if non-nil, zero value otherwise.

### GetDbLoggingLevelOk

`func (o *LocalDbBackendShared) GetDbLoggingLevelOk() (*string, bool)`

GetDbLoggingLevelOk returns a tuple with the DbLoggingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLoggingLevel

`func (o *LocalDbBackendShared) SetDbLoggingLevel(v string)`

SetDbLoggingLevel sets DbLoggingLevel field to given value.

### HasDbLoggingLevel

`func (o *LocalDbBackendShared) HasDbLoggingLevel() bool`

HasDbLoggingLevel returns a boolean if a field has been set.

### GetJeProperty

`func (o *LocalDbBackendShared) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *LocalDbBackendShared) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *LocalDbBackendShared) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *LocalDbBackendShared) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *LocalDbBackendShared) GetDbCachePercent() int32`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *LocalDbBackendShared) GetDbCachePercentOk() (*int32, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *LocalDbBackendShared) SetDbCachePercent(v int32)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *LocalDbBackendShared) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetDefaultCacheMode

`func (o *LocalDbBackendShared) GetDefaultCacheMode() EnumbackendDefaultCacheModeProp`

GetDefaultCacheMode returns the DefaultCacheMode field if non-nil, zero value otherwise.

### GetDefaultCacheModeOk

`func (o *LocalDbBackendShared) GetDefaultCacheModeOk() (*EnumbackendDefaultCacheModeProp, bool)`

GetDefaultCacheModeOk returns a tuple with the DefaultCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCacheMode

`func (o *LocalDbBackendShared) SetDefaultCacheMode(v EnumbackendDefaultCacheModeProp)`

SetDefaultCacheMode sets DefaultCacheMode field to given value.

### HasDefaultCacheMode

`func (o *LocalDbBackendShared) HasDefaultCacheMode() bool`

HasDefaultCacheMode returns a boolean if a field has been set.

### GetId2entryCacheMode

`func (o *LocalDbBackendShared) GetId2entryCacheMode() EnumbackendId2entryCacheModeProp`

GetId2entryCacheMode returns the Id2entryCacheMode field if non-nil, zero value otherwise.

### GetId2entryCacheModeOk

`func (o *LocalDbBackendShared) GetId2entryCacheModeOk() (*EnumbackendId2entryCacheModeProp, bool)`

GetId2entryCacheModeOk returns a tuple with the Id2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2entryCacheMode

`func (o *LocalDbBackendShared) SetId2entryCacheMode(v EnumbackendId2entryCacheModeProp)`

SetId2entryCacheMode sets Id2entryCacheMode field to given value.

### HasId2entryCacheMode

`func (o *LocalDbBackendShared) HasId2entryCacheMode() bool`

HasId2entryCacheMode returns a boolean if a field has been set.

### GetDn2idCacheMode

`func (o *LocalDbBackendShared) GetDn2idCacheMode() EnumbackendDn2idCacheModeProp`

GetDn2idCacheMode returns the Dn2idCacheMode field if non-nil, zero value otherwise.

### GetDn2idCacheModeOk

`func (o *LocalDbBackendShared) GetDn2idCacheModeOk() (*EnumbackendDn2idCacheModeProp, bool)`

GetDn2idCacheModeOk returns a tuple with the Dn2idCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2idCacheMode

`func (o *LocalDbBackendShared) SetDn2idCacheMode(v EnumbackendDn2idCacheModeProp)`

SetDn2idCacheMode sets Dn2idCacheMode field to given value.

### HasDn2idCacheMode

`func (o *LocalDbBackendShared) HasDn2idCacheMode() bool`

HasDn2idCacheMode returns a boolean if a field has been set.

### GetId2childrenCacheMode

`func (o *LocalDbBackendShared) GetId2childrenCacheMode() EnumbackendId2childrenCacheModeProp`

GetId2childrenCacheMode returns the Id2childrenCacheMode field if non-nil, zero value otherwise.

### GetId2childrenCacheModeOk

`func (o *LocalDbBackendShared) GetId2childrenCacheModeOk() (*EnumbackendId2childrenCacheModeProp, bool)`

GetId2childrenCacheModeOk returns a tuple with the Id2childrenCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenCacheMode

`func (o *LocalDbBackendShared) SetId2childrenCacheMode(v EnumbackendId2childrenCacheModeProp)`

SetId2childrenCacheMode sets Id2childrenCacheMode field to given value.

### HasId2childrenCacheMode

`func (o *LocalDbBackendShared) HasId2childrenCacheMode() bool`

HasId2childrenCacheMode returns a boolean if a field has been set.

### GetId2subtreeCacheMode

`func (o *LocalDbBackendShared) GetId2subtreeCacheMode() EnumbackendId2subtreeCacheModeProp`

GetId2subtreeCacheMode returns the Id2subtreeCacheMode field if non-nil, zero value otherwise.

### GetId2subtreeCacheModeOk

`func (o *LocalDbBackendShared) GetId2subtreeCacheModeOk() (*EnumbackendId2subtreeCacheModeProp, bool)`

GetId2subtreeCacheModeOk returns a tuple with the Id2subtreeCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeCacheMode

`func (o *LocalDbBackendShared) SetId2subtreeCacheMode(v EnumbackendId2subtreeCacheModeProp)`

SetId2subtreeCacheMode sets Id2subtreeCacheMode field to given value.

### HasId2subtreeCacheMode

`func (o *LocalDbBackendShared) HasId2subtreeCacheMode() bool`

HasId2subtreeCacheMode returns a boolean if a field has been set.

### GetDn2uriCacheMode

`func (o *LocalDbBackendShared) GetDn2uriCacheMode() EnumbackendDn2uriCacheModeProp`

GetDn2uriCacheMode returns the Dn2uriCacheMode field if non-nil, zero value otherwise.

### GetDn2uriCacheModeOk

`func (o *LocalDbBackendShared) GetDn2uriCacheModeOk() (*EnumbackendDn2uriCacheModeProp, bool)`

GetDn2uriCacheModeOk returns a tuple with the Dn2uriCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2uriCacheMode

`func (o *LocalDbBackendShared) SetDn2uriCacheMode(v EnumbackendDn2uriCacheModeProp)`

SetDn2uriCacheMode sets Dn2uriCacheMode field to given value.

### HasDn2uriCacheMode

`func (o *LocalDbBackendShared) HasDn2uriCacheMode() bool`

HasDn2uriCacheMode returns a boolean if a field has been set.

### GetPrimeMethod

`func (o *LocalDbBackendShared) GetPrimeMethod() []EnumbackendPrimeMethodProp`

GetPrimeMethod returns the PrimeMethod field if non-nil, zero value otherwise.

### GetPrimeMethodOk

`func (o *LocalDbBackendShared) GetPrimeMethodOk() (*[]EnumbackendPrimeMethodProp, bool)`

GetPrimeMethodOk returns a tuple with the PrimeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeMethod

`func (o *LocalDbBackendShared) SetPrimeMethod(v []EnumbackendPrimeMethodProp)`

SetPrimeMethod sets PrimeMethod field to given value.

### HasPrimeMethod

`func (o *LocalDbBackendShared) HasPrimeMethod() bool`

HasPrimeMethod returns a boolean if a field has been set.

### GetPrimeThreadCount

`func (o *LocalDbBackendShared) GetPrimeThreadCount() int32`

GetPrimeThreadCount returns the PrimeThreadCount field if non-nil, zero value otherwise.

### GetPrimeThreadCountOk

`func (o *LocalDbBackendShared) GetPrimeThreadCountOk() (*int32, bool)`

GetPrimeThreadCountOk returns a tuple with the PrimeThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeThreadCount

`func (o *LocalDbBackendShared) SetPrimeThreadCount(v int32)`

SetPrimeThreadCount sets PrimeThreadCount field to given value.

### HasPrimeThreadCount

`func (o *LocalDbBackendShared) HasPrimeThreadCount() bool`

HasPrimeThreadCount returns a boolean if a field has been set.

### GetPrimeTimeLimit

`func (o *LocalDbBackendShared) GetPrimeTimeLimit() string`

GetPrimeTimeLimit returns the PrimeTimeLimit field if non-nil, zero value otherwise.

### GetPrimeTimeLimitOk

`func (o *LocalDbBackendShared) GetPrimeTimeLimitOk() (*string, bool)`

GetPrimeTimeLimitOk returns a tuple with the PrimeTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeTimeLimit

`func (o *LocalDbBackendShared) SetPrimeTimeLimit(v string)`

SetPrimeTimeLimit sets PrimeTimeLimit field to given value.

### HasPrimeTimeLimit

`func (o *LocalDbBackendShared) HasPrimeTimeLimit() bool`

HasPrimeTimeLimit returns a boolean if a field has been set.

### GetPrimeAllIndexes

`func (o *LocalDbBackendShared) GetPrimeAllIndexes() bool`

GetPrimeAllIndexes returns the PrimeAllIndexes field if non-nil, zero value otherwise.

### GetPrimeAllIndexesOk

`func (o *LocalDbBackendShared) GetPrimeAllIndexesOk() (*bool, bool)`

GetPrimeAllIndexesOk returns a tuple with the PrimeAllIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeAllIndexes

`func (o *LocalDbBackendShared) SetPrimeAllIndexes(v bool)`

SetPrimeAllIndexes sets PrimeAllIndexes field to given value.

### HasPrimeAllIndexes

`func (o *LocalDbBackendShared) HasPrimeAllIndexes() bool`

HasPrimeAllIndexes returns a boolean if a field has been set.

### GetSystemIndexToPrime

`func (o *LocalDbBackendShared) GetSystemIndexToPrime() []EnumbackendSystemIndexToPrimeProp`

GetSystemIndexToPrime returns the SystemIndexToPrime field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeOk

`func (o *LocalDbBackendShared) GetSystemIndexToPrimeOk() (*[]EnumbackendSystemIndexToPrimeProp, bool)`

GetSystemIndexToPrimeOk returns a tuple with the SystemIndexToPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrime

`func (o *LocalDbBackendShared) SetSystemIndexToPrime(v []EnumbackendSystemIndexToPrimeProp)`

SetSystemIndexToPrime sets SystemIndexToPrime field to given value.

### HasSystemIndexToPrime

`func (o *LocalDbBackendShared) HasSystemIndexToPrime() bool`

HasSystemIndexToPrime returns a boolean if a field has been set.

### GetSystemIndexToPrimeInternalNodesOnly

`func (o *LocalDbBackendShared) GetSystemIndexToPrimeInternalNodesOnly() []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp`

GetSystemIndexToPrimeInternalNodesOnly returns the SystemIndexToPrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeInternalNodesOnlyOk

`func (o *LocalDbBackendShared) GetSystemIndexToPrimeInternalNodesOnlyOk() (*[]EnumbackendSystemIndexToPrimeInternalNodesOnlyProp, bool)`

GetSystemIndexToPrimeInternalNodesOnlyOk returns a tuple with the SystemIndexToPrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrimeInternalNodesOnly

`func (o *LocalDbBackendShared) SetSystemIndexToPrimeInternalNodesOnly(v []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp)`

SetSystemIndexToPrimeInternalNodesOnly sets SystemIndexToPrimeInternalNodesOnly field to given value.

### HasSystemIndexToPrimeInternalNodesOnly

`func (o *LocalDbBackendShared) HasSystemIndexToPrimeInternalNodesOnly() bool`

HasSystemIndexToPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetBackgroundPrime

`func (o *LocalDbBackendShared) GetBackgroundPrime() bool`

GetBackgroundPrime returns the BackgroundPrime field if non-nil, zero value otherwise.

### GetBackgroundPrimeOk

`func (o *LocalDbBackendShared) GetBackgroundPrimeOk() (*bool, bool)`

GetBackgroundPrimeOk returns a tuple with the BackgroundPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPrime

`func (o *LocalDbBackendShared) SetBackgroundPrime(v bool)`

SetBackgroundPrime sets BackgroundPrime field to given value.

### HasBackgroundPrime

`func (o *LocalDbBackendShared) HasBackgroundPrime() bool`

HasBackgroundPrime returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *LocalDbBackendShared) GetIndexEntryLimit() int32`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *LocalDbBackendShared) GetIndexEntryLimitOk() (*int32, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *LocalDbBackendShared) SetIndexEntryLimit(v int32)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *LocalDbBackendShared) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetCompositeIndexEntryLimit

`func (o *LocalDbBackendShared) GetCompositeIndexEntryLimit() int32`

GetCompositeIndexEntryLimit returns the CompositeIndexEntryLimit field if non-nil, zero value otherwise.

### GetCompositeIndexEntryLimitOk

`func (o *LocalDbBackendShared) GetCompositeIndexEntryLimitOk() (*int32, bool)`

GetCompositeIndexEntryLimitOk returns a tuple with the CompositeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeIndexEntryLimit

`func (o *LocalDbBackendShared) SetCompositeIndexEntryLimit(v int32)`

SetCompositeIndexEntryLimit sets CompositeIndexEntryLimit field to given value.

### HasCompositeIndexEntryLimit

`func (o *LocalDbBackendShared) HasCompositeIndexEntryLimit() bool`

HasCompositeIndexEntryLimit returns a boolean if a field has been set.

### GetId2childrenIndexEntryLimit

`func (o *LocalDbBackendShared) GetId2childrenIndexEntryLimit() int32`

GetId2childrenIndexEntryLimit returns the Id2childrenIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2childrenIndexEntryLimitOk

`func (o *LocalDbBackendShared) GetId2childrenIndexEntryLimitOk() (*int32, bool)`

GetId2childrenIndexEntryLimitOk returns a tuple with the Id2childrenIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenIndexEntryLimit

`func (o *LocalDbBackendShared) SetId2childrenIndexEntryLimit(v int32)`

SetId2childrenIndexEntryLimit sets Id2childrenIndexEntryLimit field to given value.

### HasId2childrenIndexEntryLimit

`func (o *LocalDbBackendShared) HasId2childrenIndexEntryLimit() bool`

HasId2childrenIndexEntryLimit returns a boolean if a field has been set.

### GetId2subtreeIndexEntryLimit

`func (o *LocalDbBackendShared) GetId2subtreeIndexEntryLimit() int32`

GetId2subtreeIndexEntryLimit returns the Id2subtreeIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2subtreeIndexEntryLimitOk

`func (o *LocalDbBackendShared) GetId2subtreeIndexEntryLimitOk() (*int32, bool)`

GetId2subtreeIndexEntryLimitOk returns a tuple with the Id2subtreeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeIndexEntryLimit

`func (o *LocalDbBackendShared) SetId2subtreeIndexEntryLimit(v int32)`

SetId2subtreeIndexEntryLimit sets Id2subtreeIndexEntryLimit field to given value.

### HasId2subtreeIndexEntryLimit

`func (o *LocalDbBackendShared) HasId2subtreeIndexEntryLimit() bool`

HasId2subtreeIndexEntryLimit returns a boolean if a field has been set.

### GetImportTempDirectory

`func (o *LocalDbBackendShared) GetImportTempDirectory() string`

GetImportTempDirectory returns the ImportTempDirectory field if non-nil, zero value otherwise.

### GetImportTempDirectoryOk

`func (o *LocalDbBackendShared) GetImportTempDirectoryOk() (*string, bool)`

GetImportTempDirectoryOk returns a tuple with the ImportTempDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTempDirectory

`func (o *LocalDbBackendShared) SetImportTempDirectory(v string)`

SetImportTempDirectory sets ImportTempDirectory field to given value.


### GetImportThreadCount

`func (o *LocalDbBackendShared) GetImportThreadCount() int32`

GetImportThreadCount returns the ImportThreadCount field if non-nil, zero value otherwise.

### GetImportThreadCountOk

`func (o *LocalDbBackendShared) GetImportThreadCountOk() (*int32, bool)`

GetImportThreadCountOk returns a tuple with the ImportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportThreadCount

`func (o *LocalDbBackendShared) SetImportThreadCount(v int32)`

SetImportThreadCount sets ImportThreadCount field to given value.

### HasImportThreadCount

`func (o *LocalDbBackendShared) HasImportThreadCount() bool`

HasImportThreadCount returns a boolean if a field has been set.

### GetExportThreadCount

`func (o *LocalDbBackendShared) GetExportThreadCount() int32`

GetExportThreadCount returns the ExportThreadCount field if non-nil, zero value otherwise.

### GetExportThreadCountOk

`func (o *LocalDbBackendShared) GetExportThreadCountOk() (*int32, bool)`

GetExportThreadCountOk returns a tuple with the ExportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportThreadCount

`func (o *LocalDbBackendShared) SetExportThreadCount(v int32)`

SetExportThreadCount sets ExportThreadCount field to given value.

### HasExportThreadCount

`func (o *LocalDbBackendShared) HasExportThreadCount() bool`

HasExportThreadCount returns a boolean if a field has been set.

### GetDbImportCachePercent

`func (o *LocalDbBackendShared) GetDbImportCachePercent() int32`

GetDbImportCachePercent returns the DbImportCachePercent field if non-nil, zero value otherwise.

### GetDbImportCachePercentOk

`func (o *LocalDbBackendShared) GetDbImportCachePercentOk() (*int32, bool)`

GetDbImportCachePercentOk returns a tuple with the DbImportCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbImportCachePercent

`func (o *LocalDbBackendShared) SetDbImportCachePercent(v int32)`

SetDbImportCachePercent sets DbImportCachePercent field to given value.

### HasDbImportCachePercent

`func (o *LocalDbBackendShared) HasDbImportCachePercent() bool`

HasDbImportCachePercent returns a boolean if a field has been set.

### GetDbTxnWriteNoSync

`func (o *LocalDbBackendShared) GetDbTxnWriteNoSync() bool`

GetDbTxnWriteNoSync returns the DbTxnWriteNoSync field if non-nil, zero value otherwise.

### GetDbTxnWriteNoSyncOk

`func (o *LocalDbBackendShared) GetDbTxnWriteNoSyncOk() (*bool, bool)`

GetDbTxnWriteNoSyncOk returns a tuple with the DbTxnWriteNoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTxnWriteNoSync

`func (o *LocalDbBackendShared) SetDbTxnWriteNoSync(v bool)`

SetDbTxnWriteNoSync sets DbTxnWriteNoSync field to given value.

### HasDbTxnWriteNoSync

`func (o *LocalDbBackendShared) HasDbTxnWriteNoSync() bool`

HasDbTxnWriteNoSync returns a boolean if a field has been set.

### GetDeadlockRetryLimit

`func (o *LocalDbBackendShared) GetDeadlockRetryLimit() int32`

GetDeadlockRetryLimit returns the DeadlockRetryLimit field if non-nil, zero value otherwise.

### GetDeadlockRetryLimitOk

`func (o *LocalDbBackendShared) GetDeadlockRetryLimitOk() (*int32, bool)`

GetDeadlockRetryLimitOk returns a tuple with the DeadlockRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockRetryLimit

`func (o *LocalDbBackendShared) SetDeadlockRetryLimit(v int32)`

SetDeadlockRetryLimit sets DeadlockRetryLimit field to given value.

### HasDeadlockRetryLimit

`func (o *LocalDbBackendShared) HasDeadlockRetryLimit() bool`

HasDeadlockRetryLimit returns a boolean if a field has been set.

### GetExternalTxnDefaultBackendLockBehavior

`func (o *LocalDbBackendShared) GetExternalTxnDefaultBackendLockBehavior() EnumbackendExternalTxnDefaultBackendLockBehaviorProp`

GetExternalTxnDefaultBackendLockBehavior returns the ExternalTxnDefaultBackendLockBehavior field if non-nil, zero value otherwise.

### GetExternalTxnDefaultBackendLockBehaviorOk

`func (o *LocalDbBackendShared) GetExternalTxnDefaultBackendLockBehaviorOk() (*EnumbackendExternalTxnDefaultBackendLockBehaviorProp, bool)`

GetExternalTxnDefaultBackendLockBehaviorOk returns a tuple with the ExternalTxnDefaultBackendLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTxnDefaultBackendLockBehavior

`func (o *LocalDbBackendShared) SetExternalTxnDefaultBackendLockBehavior(v EnumbackendExternalTxnDefaultBackendLockBehaviorProp)`

SetExternalTxnDefaultBackendLockBehavior sets ExternalTxnDefaultBackendLockBehavior field to given value.

### HasExternalTxnDefaultBackendLockBehavior

`func (o *LocalDbBackendShared) HasExternalTxnDefaultBackendLockBehavior() bool`

HasExternalTxnDefaultBackendLockBehavior returns a boolean if a field has been set.

### GetSingleWriterLockBehavior

`func (o *LocalDbBackendShared) GetSingleWriterLockBehavior() EnumbackendSingleWriterLockBehaviorProp`

GetSingleWriterLockBehavior returns the SingleWriterLockBehavior field if non-nil, zero value otherwise.

### GetSingleWriterLockBehaviorOk

`func (o *LocalDbBackendShared) GetSingleWriterLockBehaviorOk() (*EnumbackendSingleWriterLockBehaviorProp, bool)`

GetSingleWriterLockBehaviorOk returns a tuple with the SingleWriterLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleWriterLockBehavior

`func (o *LocalDbBackendShared) SetSingleWriterLockBehavior(v EnumbackendSingleWriterLockBehaviorProp)`

SetSingleWriterLockBehavior sets SingleWriterLockBehavior field to given value.

### HasSingleWriterLockBehavior

`func (o *LocalDbBackendShared) HasSingleWriterLockBehavior() bool`

HasSingleWriterLockBehavior returns a boolean if a field has been set.

### GetSubtreeDeleteSizeLimit

`func (o *LocalDbBackendShared) GetSubtreeDeleteSizeLimit() int32`

GetSubtreeDeleteSizeLimit returns the SubtreeDeleteSizeLimit field if non-nil, zero value otherwise.

### GetSubtreeDeleteSizeLimitOk

`func (o *LocalDbBackendShared) GetSubtreeDeleteSizeLimitOk() (*int32, bool)`

GetSubtreeDeleteSizeLimitOk returns a tuple with the SubtreeDeleteSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeDeleteSizeLimit

`func (o *LocalDbBackendShared) SetSubtreeDeleteSizeLimit(v int32)`

SetSubtreeDeleteSizeLimit sets SubtreeDeleteSizeLimit field to given value.

### HasSubtreeDeleteSizeLimit

`func (o *LocalDbBackendShared) HasSubtreeDeleteSizeLimit() bool`

HasSubtreeDeleteSizeLimit returns a boolean if a field has been set.

### GetNumRecentChanges

`func (o *LocalDbBackendShared) GetNumRecentChanges() int32`

GetNumRecentChanges returns the NumRecentChanges field if non-nil, zero value otherwise.

### GetNumRecentChangesOk

`func (o *LocalDbBackendShared) GetNumRecentChangesOk() (*int32, bool)`

GetNumRecentChangesOk returns a tuple with the NumRecentChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecentChanges

`func (o *LocalDbBackendShared) SetNumRecentChanges(v int32)`

SetNumRecentChanges sets NumRecentChanges field to given value.

### HasNumRecentChanges

`func (o *LocalDbBackendShared) HasNumRecentChanges() bool`

HasNumRecentChanges returns a boolean if a field has been set.

### GetOfflineProcessDatabaseOpenTimeout

`func (o *LocalDbBackendShared) GetOfflineProcessDatabaseOpenTimeout() string`

GetOfflineProcessDatabaseOpenTimeout returns the OfflineProcessDatabaseOpenTimeout field if non-nil, zero value otherwise.

### GetOfflineProcessDatabaseOpenTimeoutOk

`func (o *LocalDbBackendShared) GetOfflineProcessDatabaseOpenTimeoutOk() (*string, bool)`

GetOfflineProcessDatabaseOpenTimeoutOk returns a tuple with the OfflineProcessDatabaseOpenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineProcessDatabaseOpenTimeout

`func (o *LocalDbBackendShared) SetOfflineProcessDatabaseOpenTimeout(v string)`

SetOfflineProcessDatabaseOpenTimeout sets OfflineProcessDatabaseOpenTimeout field to given value.

### HasOfflineProcessDatabaseOpenTimeout

`func (o *LocalDbBackendShared) HasOfflineProcessDatabaseOpenTimeout() bool`

HasOfflineProcessDatabaseOpenTimeout returns a boolean if a field has been set.

### GetBackendID

`func (o *LocalDbBackendShared) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *LocalDbBackendShared) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *LocalDbBackendShared) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetDescription

`func (o *LocalDbBackendShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocalDbBackendShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocalDbBackendShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocalDbBackendShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LocalDbBackendShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LocalDbBackendShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LocalDbBackendShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *LocalDbBackendShared) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *LocalDbBackendShared) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *LocalDbBackendShared) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *LocalDbBackendShared) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *LocalDbBackendShared) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *LocalDbBackendShared) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *LocalDbBackendShared) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *LocalDbBackendShared) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *LocalDbBackendShared) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *LocalDbBackendShared) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *LocalDbBackendShared) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *LocalDbBackendShared) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *LocalDbBackendShared) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *LocalDbBackendShared) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *LocalDbBackendShared) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *LocalDbBackendShared) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *LocalDbBackendShared) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *LocalDbBackendShared) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *LocalDbBackendShared) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


