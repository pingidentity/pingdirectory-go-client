# AddBackend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Backend | 
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
**JeProperty** | Pointer to **[]string** | Specifies the database and environment properties for the Berkeley DB Java Edition database serving the data for this backend. | [optional] 
**DbCachePercent** | Pointer to **int64** | Specifies the percentage of JVM memory to allocate to the database cache. | [optional] 
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
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewAddBackend200Response

`func NewAddBackend200Response(id string, schemas []EnumlocalDbBackendSchemaUrn, writabilityMode EnumbackendWritabilityModeProp, dbDirectory string, importTempDirectory string, backendID string, enabled bool, baseDN []string, ) *AddBackend200Response`

NewAddBackend200Response instantiates a new AddBackend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackend200ResponseWithDefaults

`func NewAddBackend200ResponseWithDefaults() *AddBackend200Response`

NewAddBackend200ResponseWithDefaults instantiates a new AddBackend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddBackend200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddBackend200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddBackend200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddBackend200Response) GetSchemas() []EnumlocalDbBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddBackend200Response) GetSchemasOk() (*[]EnumlocalDbBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddBackend200Response) SetSchemas(v []EnumlocalDbBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetUncachedId2entryCacheMode

`func (o *AddBackend200Response) GetUncachedId2entryCacheMode() EnumbackendUncachedId2entryCacheModeProp`

GetUncachedId2entryCacheMode returns the UncachedId2entryCacheMode field if non-nil, zero value otherwise.

### GetUncachedId2entryCacheModeOk

`func (o *AddBackend200Response) GetUncachedId2entryCacheModeOk() (*EnumbackendUncachedId2entryCacheModeProp, bool)`

GetUncachedId2entryCacheModeOk returns a tuple with the UncachedId2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedId2entryCacheMode

`func (o *AddBackend200Response) SetUncachedId2entryCacheMode(v EnumbackendUncachedId2entryCacheModeProp)`

SetUncachedId2entryCacheMode sets UncachedId2entryCacheMode field to given value.

### HasUncachedId2entryCacheMode

`func (o *AddBackend200Response) HasUncachedId2entryCacheMode() bool`

HasUncachedId2entryCacheMode returns a boolean if a field has been set.

### GetUncachedAttributeCriteria

`func (o *AddBackend200Response) GetUncachedAttributeCriteria() string`

GetUncachedAttributeCriteria returns the UncachedAttributeCriteria field if non-nil, zero value otherwise.

### GetUncachedAttributeCriteriaOk

`func (o *AddBackend200Response) GetUncachedAttributeCriteriaOk() (*string, bool)`

GetUncachedAttributeCriteriaOk returns a tuple with the UncachedAttributeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedAttributeCriteria

`func (o *AddBackend200Response) SetUncachedAttributeCriteria(v string)`

SetUncachedAttributeCriteria sets UncachedAttributeCriteria field to given value.

### HasUncachedAttributeCriteria

`func (o *AddBackend200Response) HasUncachedAttributeCriteria() bool`

HasUncachedAttributeCriteria returns a boolean if a field has been set.

### GetUncachedEntryCriteria

`func (o *AddBackend200Response) GetUncachedEntryCriteria() string`

GetUncachedEntryCriteria returns the UncachedEntryCriteria field if non-nil, zero value otherwise.

### GetUncachedEntryCriteriaOk

`func (o *AddBackend200Response) GetUncachedEntryCriteriaOk() (*string, bool)`

GetUncachedEntryCriteriaOk returns a tuple with the UncachedEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncachedEntryCriteria

`func (o *AddBackend200Response) SetUncachedEntryCriteria(v string)`

SetUncachedEntryCriteria sets UncachedEntryCriteria field to given value.

### HasUncachedEntryCriteria

`func (o *AddBackend200Response) HasUncachedEntryCriteria() bool`

HasUncachedEntryCriteria returns a boolean if a field has been set.

### GetWritabilityMode

`func (o *AddBackend200Response) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *AddBackend200Response) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *AddBackend200Response) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetSetDegradedAlertForUntrustedIndex

`func (o *AddBackend200Response) GetSetDegradedAlertForUntrustedIndex() bool`

GetSetDegradedAlertForUntrustedIndex returns the SetDegradedAlertForUntrustedIndex field if non-nil, zero value otherwise.

### GetSetDegradedAlertForUntrustedIndexOk

`func (o *AddBackend200Response) GetSetDegradedAlertForUntrustedIndexOk() (*bool, bool)`

GetSetDegradedAlertForUntrustedIndexOk returns a tuple with the SetDegradedAlertForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertForUntrustedIndex

`func (o *AddBackend200Response) SetSetDegradedAlertForUntrustedIndex(v bool)`

SetSetDegradedAlertForUntrustedIndex sets SetDegradedAlertForUntrustedIndex field to given value.

### HasSetDegradedAlertForUntrustedIndex

`func (o *AddBackend200Response) HasSetDegradedAlertForUntrustedIndex() bool`

HasSetDegradedAlertForUntrustedIndex returns a boolean if a field has been set.

### GetReturnUnavailableForUntrustedIndex

`func (o *AddBackend200Response) GetReturnUnavailableForUntrustedIndex() bool`

GetReturnUnavailableForUntrustedIndex returns the ReturnUnavailableForUntrustedIndex field if non-nil, zero value otherwise.

### GetReturnUnavailableForUntrustedIndexOk

`func (o *AddBackend200Response) GetReturnUnavailableForUntrustedIndexOk() (*bool, bool)`

GetReturnUnavailableForUntrustedIndexOk returns a tuple with the ReturnUnavailableForUntrustedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableForUntrustedIndex

`func (o *AddBackend200Response) SetReturnUnavailableForUntrustedIndex(v bool)`

SetReturnUnavailableForUntrustedIndex sets ReturnUnavailableForUntrustedIndex field to given value.

### HasReturnUnavailableForUntrustedIndex

`func (o *AddBackend200Response) HasReturnUnavailableForUntrustedIndex() bool`

HasReturnUnavailableForUntrustedIndex returns a boolean if a field has been set.

### GetProcessFiltersWithUndefinedAttributeTypes

`func (o *AddBackend200Response) GetProcessFiltersWithUndefinedAttributeTypes() bool`

GetProcessFiltersWithUndefinedAttributeTypes returns the ProcessFiltersWithUndefinedAttributeTypes field if non-nil, zero value otherwise.

### GetProcessFiltersWithUndefinedAttributeTypesOk

`func (o *AddBackend200Response) GetProcessFiltersWithUndefinedAttributeTypesOk() (*bool, bool)`

GetProcessFiltersWithUndefinedAttributeTypesOk returns a tuple with the ProcessFiltersWithUndefinedAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFiltersWithUndefinedAttributeTypes

`func (o *AddBackend200Response) SetProcessFiltersWithUndefinedAttributeTypes(v bool)`

SetProcessFiltersWithUndefinedAttributeTypes sets ProcessFiltersWithUndefinedAttributeTypes field to given value.

### HasProcessFiltersWithUndefinedAttributeTypes

`func (o *AddBackend200Response) HasProcessFiltersWithUndefinedAttributeTypes() bool`

HasProcessFiltersWithUndefinedAttributeTypes returns a boolean if a field has been set.

### GetIsPrivateBackend

`func (o *AddBackend200Response) GetIsPrivateBackend() bool`

GetIsPrivateBackend returns the IsPrivateBackend field if non-nil, zero value otherwise.

### GetIsPrivateBackendOk

`func (o *AddBackend200Response) GetIsPrivateBackendOk() (*bool, bool)`

GetIsPrivateBackendOk returns a tuple with the IsPrivateBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateBackend

`func (o *AddBackend200Response) SetIsPrivateBackend(v bool)`

SetIsPrivateBackend sets IsPrivateBackend field to given value.

### HasIsPrivateBackend

`func (o *AddBackend200Response) HasIsPrivateBackend() bool`

HasIsPrivateBackend returns a boolean if a field has been set.

### GetDbDirectory

`func (o *AddBackend200Response) GetDbDirectory() string`

GetDbDirectory returns the DbDirectory field if non-nil, zero value otherwise.

### GetDbDirectoryOk

`func (o *AddBackend200Response) GetDbDirectoryOk() (*string, bool)`

GetDbDirectoryOk returns a tuple with the DbDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectory

`func (o *AddBackend200Response) SetDbDirectory(v string)`

SetDbDirectory sets DbDirectory field to given value.


### GetDbDirectoryPermissions

`func (o *AddBackend200Response) GetDbDirectoryPermissions() string`

GetDbDirectoryPermissions returns the DbDirectoryPermissions field if non-nil, zero value otherwise.

### GetDbDirectoryPermissionsOk

`func (o *AddBackend200Response) GetDbDirectoryPermissionsOk() (*string, bool)`

GetDbDirectoryPermissionsOk returns a tuple with the DbDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectoryPermissions

`func (o *AddBackend200Response) SetDbDirectoryPermissions(v string)`

SetDbDirectoryPermissions sets DbDirectoryPermissions field to given value.

### HasDbDirectoryPermissions

`func (o *AddBackend200Response) HasDbDirectoryPermissions() bool`

HasDbDirectoryPermissions returns a boolean if a field has been set.

### GetCompactCommonParentDN

`func (o *AddBackend200Response) GetCompactCommonParentDN() []string`

GetCompactCommonParentDN returns the CompactCommonParentDN field if non-nil, zero value otherwise.

### GetCompactCommonParentDNOk

`func (o *AddBackend200Response) GetCompactCommonParentDNOk() (*[]string, bool)`

GetCompactCommonParentDNOk returns a tuple with the CompactCommonParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactCommonParentDN

`func (o *AddBackend200Response) SetCompactCommonParentDN(v []string)`

SetCompactCommonParentDN sets CompactCommonParentDN field to given value.

### HasCompactCommonParentDN

`func (o *AddBackend200Response) HasCompactCommonParentDN() bool`

HasCompactCommonParentDN returns a boolean if a field has been set.

### GetCompressEntries

`func (o *AddBackend200Response) GetCompressEntries() bool`

GetCompressEntries returns the CompressEntries field if non-nil, zero value otherwise.

### GetCompressEntriesOk

`func (o *AddBackend200Response) GetCompressEntriesOk() (*bool, bool)`

GetCompressEntriesOk returns a tuple with the CompressEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressEntries

`func (o *AddBackend200Response) SetCompressEntries(v bool)`

SetCompressEntries sets CompressEntries field to given value.

### HasCompressEntries

`func (o *AddBackend200Response) HasCompressEntries() bool`

HasCompressEntries returns a boolean if a field has been set.

### GetHashEntries

`func (o *AddBackend200Response) GetHashEntries() bool`

GetHashEntries returns the HashEntries field if non-nil, zero value otherwise.

### GetHashEntriesOk

`func (o *AddBackend200Response) GetHashEntriesOk() (*bool, bool)`

GetHashEntriesOk returns a tuple with the HashEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashEntries

`func (o *AddBackend200Response) SetHashEntries(v bool)`

SetHashEntries sets HashEntries field to given value.

### HasHashEntries

`func (o *AddBackend200Response) HasHashEntries() bool`

HasHashEntries returns a boolean if a field has been set.

### GetDbNumCleanerThreads

`func (o *AddBackend200Response) GetDbNumCleanerThreads() int64`

GetDbNumCleanerThreads returns the DbNumCleanerThreads field if non-nil, zero value otherwise.

### GetDbNumCleanerThreadsOk

`func (o *AddBackend200Response) GetDbNumCleanerThreadsOk() (*int64, bool)`

GetDbNumCleanerThreadsOk returns a tuple with the DbNumCleanerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbNumCleanerThreads

`func (o *AddBackend200Response) SetDbNumCleanerThreads(v int64)`

SetDbNumCleanerThreads sets DbNumCleanerThreads field to given value.

### HasDbNumCleanerThreads

`func (o *AddBackend200Response) HasDbNumCleanerThreads() bool`

HasDbNumCleanerThreads returns a boolean if a field has been set.

### GetDbCleanerMinUtilization

`func (o *AddBackend200Response) GetDbCleanerMinUtilization() int64`

GetDbCleanerMinUtilization returns the DbCleanerMinUtilization field if non-nil, zero value otherwise.

### GetDbCleanerMinUtilizationOk

`func (o *AddBackend200Response) GetDbCleanerMinUtilizationOk() (*int64, bool)`

GetDbCleanerMinUtilizationOk returns a tuple with the DbCleanerMinUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCleanerMinUtilization

`func (o *AddBackend200Response) SetDbCleanerMinUtilization(v int64)`

SetDbCleanerMinUtilization sets DbCleanerMinUtilization field to given value.

### HasDbCleanerMinUtilization

`func (o *AddBackend200Response) HasDbCleanerMinUtilization() bool`

HasDbCleanerMinUtilization returns a boolean if a field has been set.

### GetDbEvictorCriticalPercentage

`func (o *AddBackend200Response) GetDbEvictorCriticalPercentage() int64`

GetDbEvictorCriticalPercentage returns the DbEvictorCriticalPercentage field if non-nil, zero value otherwise.

### GetDbEvictorCriticalPercentageOk

`func (o *AddBackend200Response) GetDbEvictorCriticalPercentageOk() (*int64, bool)`

GetDbEvictorCriticalPercentageOk returns a tuple with the DbEvictorCriticalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEvictorCriticalPercentage

`func (o *AddBackend200Response) SetDbEvictorCriticalPercentage(v int64)`

SetDbEvictorCriticalPercentage sets DbEvictorCriticalPercentage field to given value.

### HasDbEvictorCriticalPercentage

`func (o *AddBackend200Response) HasDbEvictorCriticalPercentage() bool`

HasDbEvictorCriticalPercentage returns a boolean if a field has been set.

### GetDbCheckpointerWakeupInterval

`func (o *AddBackend200Response) GetDbCheckpointerWakeupInterval() string`

GetDbCheckpointerWakeupInterval returns the DbCheckpointerWakeupInterval field if non-nil, zero value otherwise.

### GetDbCheckpointerWakeupIntervalOk

`func (o *AddBackend200Response) GetDbCheckpointerWakeupIntervalOk() (*string, bool)`

GetDbCheckpointerWakeupIntervalOk returns a tuple with the DbCheckpointerWakeupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCheckpointerWakeupInterval

`func (o *AddBackend200Response) SetDbCheckpointerWakeupInterval(v string)`

SetDbCheckpointerWakeupInterval sets DbCheckpointerWakeupInterval field to given value.

### HasDbCheckpointerWakeupInterval

`func (o *AddBackend200Response) HasDbCheckpointerWakeupInterval() bool`

HasDbCheckpointerWakeupInterval returns a boolean if a field has been set.

### GetDbBackgroundSyncInterval

`func (o *AddBackend200Response) GetDbBackgroundSyncInterval() string`

GetDbBackgroundSyncInterval returns the DbBackgroundSyncInterval field if non-nil, zero value otherwise.

### GetDbBackgroundSyncIntervalOk

`func (o *AddBackend200Response) GetDbBackgroundSyncIntervalOk() (*string, bool)`

GetDbBackgroundSyncIntervalOk returns a tuple with the DbBackgroundSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbBackgroundSyncInterval

`func (o *AddBackend200Response) SetDbBackgroundSyncInterval(v string)`

SetDbBackgroundSyncInterval sets DbBackgroundSyncInterval field to given value.

### HasDbBackgroundSyncInterval

`func (o *AddBackend200Response) HasDbBackgroundSyncInterval() bool`

HasDbBackgroundSyncInterval returns a boolean if a field has been set.

### GetDbUseThreadLocalHandles

`func (o *AddBackend200Response) GetDbUseThreadLocalHandles() bool`

GetDbUseThreadLocalHandles returns the DbUseThreadLocalHandles field if non-nil, zero value otherwise.

### GetDbUseThreadLocalHandlesOk

`func (o *AddBackend200Response) GetDbUseThreadLocalHandlesOk() (*bool, bool)`

GetDbUseThreadLocalHandlesOk returns a tuple with the DbUseThreadLocalHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUseThreadLocalHandles

`func (o *AddBackend200Response) SetDbUseThreadLocalHandles(v bool)`

SetDbUseThreadLocalHandles sets DbUseThreadLocalHandles field to given value.

### HasDbUseThreadLocalHandles

`func (o *AddBackend200Response) HasDbUseThreadLocalHandles() bool`

HasDbUseThreadLocalHandles returns a boolean if a field has been set.

### GetDbLogFileMax

`func (o *AddBackend200Response) GetDbLogFileMax() string`

GetDbLogFileMax returns the DbLogFileMax field if non-nil, zero value otherwise.

### GetDbLogFileMaxOk

`func (o *AddBackend200Response) GetDbLogFileMaxOk() (*string, bool)`

GetDbLogFileMaxOk returns a tuple with the DbLogFileMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLogFileMax

`func (o *AddBackend200Response) SetDbLogFileMax(v string)`

SetDbLogFileMax sets DbLogFileMax field to given value.

### HasDbLogFileMax

`func (o *AddBackend200Response) HasDbLogFileMax() bool`

HasDbLogFileMax returns a boolean if a field has been set.

### GetDbLoggingLevel

`func (o *AddBackend200Response) GetDbLoggingLevel() string`

GetDbLoggingLevel returns the DbLoggingLevel field if non-nil, zero value otherwise.

### GetDbLoggingLevelOk

`func (o *AddBackend200Response) GetDbLoggingLevelOk() (*string, bool)`

GetDbLoggingLevelOk returns a tuple with the DbLoggingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbLoggingLevel

`func (o *AddBackend200Response) SetDbLoggingLevel(v string)`

SetDbLoggingLevel sets DbLoggingLevel field to given value.

### HasDbLoggingLevel

`func (o *AddBackend200Response) HasDbLoggingLevel() bool`

HasDbLoggingLevel returns a boolean if a field has been set.

### GetJeProperty

`func (o *AddBackend200Response) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *AddBackend200Response) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *AddBackend200Response) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *AddBackend200Response) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *AddBackend200Response) GetDbCachePercent() int64`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *AddBackend200Response) GetDbCachePercentOk() (*int64, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *AddBackend200Response) SetDbCachePercent(v int64)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *AddBackend200Response) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetDefaultCacheMode

`func (o *AddBackend200Response) GetDefaultCacheMode() EnumbackendDefaultCacheModeProp`

GetDefaultCacheMode returns the DefaultCacheMode field if non-nil, zero value otherwise.

### GetDefaultCacheModeOk

`func (o *AddBackend200Response) GetDefaultCacheModeOk() (*EnumbackendDefaultCacheModeProp, bool)`

GetDefaultCacheModeOk returns a tuple with the DefaultCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCacheMode

`func (o *AddBackend200Response) SetDefaultCacheMode(v EnumbackendDefaultCacheModeProp)`

SetDefaultCacheMode sets DefaultCacheMode field to given value.

### HasDefaultCacheMode

`func (o *AddBackend200Response) HasDefaultCacheMode() bool`

HasDefaultCacheMode returns a boolean if a field has been set.

### GetId2entryCacheMode

`func (o *AddBackend200Response) GetId2entryCacheMode() EnumbackendId2entryCacheModeProp`

GetId2entryCacheMode returns the Id2entryCacheMode field if non-nil, zero value otherwise.

### GetId2entryCacheModeOk

`func (o *AddBackend200Response) GetId2entryCacheModeOk() (*EnumbackendId2entryCacheModeProp, bool)`

GetId2entryCacheModeOk returns a tuple with the Id2entryCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2entryCacheMode

`func (o *AddBackend200Response) SetId2entryCacheMode(v EnumbackendId2entryCacheModeProp)`

SetId2entryCacheMode sets Id2entryCacheMode field to given value.

### HasId2entryCacheMode

`func (o *AddBackend200Response) HasId2entryCacheMode() bool`

HasId2entryCacheMode returns a boolean if a field has been set.

### GetDn2idCacheMode

`func (o *AddBackend200Response) GetDn2idCacheMode() EnumbackendDn2idCacheModeProp`

GetDn2idCacheMode returns the Dn2idCacheMode field if non-nil, zero value otherwise.

### GetDn2idCacheModeOk

`func (o *AddBackend200Response) GetDn2idCacheModeOk() (*EnumbackendDn2idCacheModeProp, bool)`

GetDn2idCacheModeOk returns a tuple with the Dn2idCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2idCacheMode

`func (o *AddBackend200Response) SetDn2idCacheMode(v EnumbackendDn2idCacheModeProp)`

SetDn2idCacheMode sets Dn2idCacheMode field to given value.

### HasDn2idCacheMode

`func (o *AddBackend200Response) HasDn2idCacheMode() bool`

HasDn2idCacheMode returns a boolean if a field has been set.

### GetId2childrenCacheMode

`func (o *AddBackend200Response) GetId2childrenCacheMode() EnumbackendId2childrenCacheModeProp`

GetId2childrenCacheMode returns the Id2childrenCacheMode field if non-nil, zero value otherwise.

### GetId2childrenCacheModeOk

`func (o *AddBackend200Response) GetId2childrenCacheModeOk() (*EnumbackendId2childrenCacheModeProp, bool)`

GetId2childrenCacheModeOk returns a tuple with the Id2childrenCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenCacheMode

`func (o *AddBackend200Response) SetId2childrenCacheMode(v EnumbackendId2childrenCacheModeProp)`

SetId2childrenCacheMode sets Id2childrenCacheMode field to given value.

### HasId2childrenCacheMode

`func (o *AddBackend200Response) HasId2childrenCacheMode() bool`

HasId2childrenCacheMode returns a boolean if a field has been set.

### GetId2subtreeCacheMode

`func (o *AddBackend200Response) GetId2subtreeCacheMode() EnumbackendId2subtreeCacheModeProp`

GetId2subtreeCacheMode returns the Id2subtreeCacheMode field if non-nil, zero value otherwise.

### GetId2subtreeCacheModeOk

`func (o *AddBackend200Response) GetId2subtreeCacheModeOk() (*EnumbackendId2subtreeCacheModeProp, bool)`

GetId2subtreeCacheModeOk returns a tuple with the Id2subtreeCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeCacheMode

`func (o *AddBackend200Response) SetId2subtreeCacheMode(v EnumbackendId2subtreeCacheModeProp)`

SetId2subtreeCacheMode sets Id2subtreeCacheMode field to given value.

### HasId2subtreeCacheMode

`func (o *AddBackend200Response) HasId2subtreeCacheMode() bool`

HasId2subtreeCacheMode returns a boolean if a field has been set.

### GetDn2uriCacheMode

`func (o *AddBackend200Response) GetDn2uriCacheMode() EnumbackendDn2uriCacheModeProp`

GetDn2uriCacheMode returns the Dn2uriCacheMode field if non-nil, zero value otherwise.

### GetDn2uriCacheModeOk

`func (o *AddBackend200Response) GetDn2uriCacheModeOk() (*EnumbackendDn2uriCacheModeProp, bool)`

GetDn2uriCacheModeOk returns a tuple with the Dn2uriCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn2uriCacheMode

`func (o *AddBackend200Response) SetDn2uriCacheMode(v EnumbackendDn2uriCacheModeProp)`

SetDn2uriCacheMode sets Dn2uriCacheMode field to given value.

### HasDn2uriCacheMode

`func (o *AddBackend200Response) HasDn2uriCacheMode() bool`

HasDn2uriCacheMode returns a boolean if a field has been set.

### GetPrimeMethod

`func (o *AddBackend200Response) GetPrimeMethod() []EnumbackendPrimeMethodProp`

GetPrimeMethod returns the PrimeMethod field if non-nil, zero value otherwise.

### GetPrimeMethodOk

`func (o *AddBackend200Response) GetPrimeMethodOk() (*[]EnumbackendPrimeMethodProp, bool)`

GetPrimeMethodOk returns a tuple with the PrimeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeMethod

`func (o *AddBackend200Response) SetPrimeMethod(v []EnumbackendPrimeMethodProp)`

SetPrimeMethod sets PrimeMethod field to given value.

### HasPrimeMethod

`func (o *AddBackend200Response) HasPrimeMethod() bool`

HasPrimeMethod returns a boolean if a field has been set.

### GetPrimeThreadCount

`func (o *AddBackend200Response) GetPrimeThreadCount() int64`

GetPrimeThreadCount returns the PrimeThreadCount field if non-nil, zero value otherwise.

### GetPrimeThreadCountOk

`func (o *AddBackend200Response) GetPrimeThreadCountOk() (*int64, bool)`

GetPrimeThreadCountOk returns a tuple with the PrimeThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeThreadCount

`func (o *AddBackend200Response) SetPrimeThreadCount(v int64)`

SetPrimeThreadCount sets PrimeThreadCount field to given value.

### HasPrimeThreadCount

`func (o *AddBackend200Response) HasPrimeThreadCount() bool`

HasPrimeThreadCount returns a boolean if a field has been set.

### GetPrimeTimeLimit

`func (o *AddBackend200Response) GetPrimeTimeLimit() string`

GetPrimeTimeLimit returns the PrimeTimeLimit field if non-nil, zero value otherwise.

### GetPrimeTimeLimitOk

`func (o *AddBackend200Response) GetPrimeTimeLimitOk() (*string, bool)`

GetPrimeTimeLimitOk returns a tuple with the PrimeTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeTimeLimit

`func (o *AddBackend200Response) SetPrimeTimeLimit(v string)`

SetPrimeTimeLimit sets PrimeTimeLimit field to given value.

### HasPrimeTimeLimit

`func (o *AddBackend200Response) HasPrimeTimeLimit() bool`

HasPrimeTimeLimit returns a boolean if a field has been set.

### GetPrimeAllIndexes

`func (o *AddBackend200Response) GetPrimeAllIndexes() bool`

GetPrimeAllIndexes returns the PrimeAllIndexes field if non-nil, zero value otherwise.

### GetPrimeAllIndexesOk

`func (o *AddBackend200Response) GetPrimeAllIndexesOk() (*bool, bool)`

GetPrimeAllIndexesOk returns a tuple with the PrimeAllIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeAllIndexes

`func (o *AddBackend200Response) SetPrimeAllIndexes(v bool)`

SetPrimeAllIndexes sets PrimeAllIndexes field to given value.

### HasPrimeAllIndexes

`func (o *AddBackend200Response) HasPrimeAllIndexes() bool`

HasPrimeAllIndexes returns a boolean if a field has been set.

### GetSystemIndexToPrime

`func (o *AddBackend200Response) GetSystemIndexToPrime() []EnumbackendSystemIndexToPrimeProp`

GetSystemIndexToPrime returns the SystemIndexToPrime field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeOk

`func (o *AddBackend200Response) GetSystemIndexToPrimeOk() (*[]EnumbackendSystemIndexToPrimeProp, bool)`

GetSystemIndexToPrimeOk returns a tuple with the SystemIndexToPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrime

`func (o *AddBackend200Response) SetSystemIndexToPrime(v []EnumbackendSystemIndexToPrimeProp)`

SetSystemIndexToPrime sets SystemIndexToPrime field to given value.

### HasSystemIndexToPrime

`func (o *AddBackend200Response) HasSystemIndexToPrime() bool`

HasSystemIndexToPrime returns a boolean if a field has been set.

### GetSystemIndexToPrimeInternalNodesOnly

`func (o *AddBackend200Response) GetSystemIndexToPrimeInternalNodesOnly() []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp`

GetSystemIndexToPrimeInternalNodesOnly returns the SystemIndexToPrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetSystemIndexToPrimeInternalNodesOnlyOk

`func (o *AddBackend200Response) GetSystemIndexToPrimeInternalNodesOnlyOk() (*[]EnumbackendSystemIndexToPrimeInternalNodesOnlyProp, bool)`

GetSystemIndexToPrimeInternalNodesOnlyOk returns a tuple with the SystemIndexToPrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIndexToPrimeInternalNodesOnly

`func (o *AddBackend200Response) SetSystemIndexToPrimeInternalNodesOnly(v []EnumbackendSystemIndexToPrimeInternalNodesOnlyProp)`

SetSystemIndexToPrimeInternalNodesOnly sets SystemIndexToPrimeInternalNodesOnly field to given value.

### HasSystemIndexToPrimeInternalNodesOnly

`func (o *AddBackend200Response) HasSystemIndexToPrimeInternalNodesOnly() bool`

HasSystemIndexToPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetBackgroundPrime

`func (o *AddBackend200Response) GetBackgroundPrime() bool`

GetBackgroundPrime returns the BackgroundPrime field if non-nil, zero value otherwise.

### GetBackgroundPrimeOk

`func (o *AddBackend200Response) GetBackgroundPrimeOk() (*bool, bool)`

GetBackgroundPrimeOk returns a tuple with the BackgroundPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPrime

`func (o *AddBackend200Response) SetBackgroundPrime(v bool)`

SetBackgroundPrime sets BackgroundPrime field to given value.

### HasBackgroundPrime

`func (o *AddBackend200Response) HasBackgroundPrime() bool`

HasBackgroundPrime returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *AddBackend200Response) GetIndexEntryLimit() int64`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *AddBackend200Response) GetIndexEntryLimitOk() (*int64, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *AddBackend200Response) SetIndexEntryLimit(v int64)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *AddBackend200Response) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetCompositeIndexEntryLimit

`func (o *AddBackend200Response) GetCompositeIndexEntryLimit() int64`

GetCompositeIndexEntryLimit returns the CompositeIndexEntryLimit field if non-nil, zero value otherwise.

### GetCompositeIndexEntryLimitOk

`func (o *AddBackend200Response) GetCompositeIndexEntryLimitOk() (*int64, bool)`

GetCompositeIndexEntryLimitOk returns a tuple with the CompositeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeIndexEntryLimit

`func (o *AddBackend200Response) SetCompositeIndexEntryLimit(v int64)`

SetCompositeIndexEntryLimit sets CompositeIndexEntryLimit field to given value.

### HasCompositeIndexEntryLimit

`func (o *AddBackend200Response) HasCompositeIndexEntryLimit() bool`

HasCompositeIndexEntryLimit returns a boolean if a field has been set.

### GetId2childrenIndexEntryLimit

`func (o *AddBackend200Response) GetId2childrenIndexEntryLimit() int64`

GetId2childrenIndexEntryLimit returns the Id2childrenIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2childrenIndexEntryLimitOk

`func (o *AddBackend200Response) GetId2childrenIndexEntryLimitOk() (*int64, bool)`

GetId2childrenIndexEntryLimitOk returns a tuple with the Id2childrenIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2childrenIndexEntryLimit

`func (o *AddBackend200Response) SetId2childrenIndexEntryLimit(v int64)`

SetId2childrenIndexEntryLimit sets Id2childrenIndexEntryLimit field to given value.

### HasId2childrenIndexEntryLimit

`func (o *AddBackend200Response) HasId2childrenIndexEntryLimit() bool`

HasId2childrenIndexEntryLimit returns a boolean if a field has been set.

### GetId2subtreeIndexEntryLimit

`func (o *AddBackend200Response) GetId2subtreeIndexEntryLimit() int64`

GetId2subtreeIndexEntryLimit returns the Id2subtreeIndexEntryLimit field if non-nil, zero value otherwise.

### GetId2subtreeIndexEntryLimitOk

`func (o *AddBackend200Response) GetId2subtreeIndexEntryLimitOk() (*int64, bool)`

GetId2subtreeIndexEntryLimitOk returns a tuple with the Id2subtreeIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId2subtreeIndexEntryLimit

`func (o *AddBackend200Response) SetId2subtreeIndexEntryLimit(v int64)`

SetId2subtreeIndexEntryLimit sets Id2subtreeIndexEntryLimit field to given value.

### HasId2subtreeIndexEntryLimit

`func (o *AddBackend200Response) HasId2subtreeIndexEntryLimit() bool`

HasId2subtreeIndexEntryLimit returns a boolean if a field has been set.

### GetImportTempDirectory

`func (o *AddBackend200Response) GetImportTempDirectory() string`

GetImportTempDirectory returns the ImportTempDirectory field if non-nil, zero value otherwise.

### GetImportTempDirectoryOk

`func (o *AddBackend200Response) GetImportTempDirectoryOk() (*string, bool)`

GetImportTempDirectoryOk returns a tuple with the ImportTempDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTempDirectory

`func (o *AddBackend200Response) SetImportTempDirectory(v string)`

SetImportTempDirectory sets ImportTempDirectory field to given value.


### GetImportThreadCount

`func (o *AddBackend200Response) GetImportThreadCount() int64`

GetImportThreadCount returns the ImportThreadCount field if non-nil, zero value otherwise.

### GetImportThreadCountOk

`func (o *AddBackend200Response) GetImportThreadCountOk() (*int64, bool)`

GetImportThreadCountOk returns a tuple with the ImportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportThreadCount

`func (o *AddBackend200Response) SetImportThreadCount(v int64)`

SetImportThreadCount sets ImportThreadCount field to given value.

### HasImportThreadCount

`func (o *AddBackend200Response) HasImportThreadCount() bool`

HasImportThreadCount returns a boolean if a field has been set.

### GetExportThreadCount

`func (o *AddBackend200Response) GetExportThreadCount() int64`

GetExportThreadCount returns the ExportThreadCount field if non-nil, zero value otherwise.

### GetExportThreadCountOk

`func (o *AddBackend200Response) GetExportThreadCountOk() (*int64, bool)`

GetExportThreadCountOk returns a tuple with the ExportThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportThreadCount

`func (o *AddBackend200Response) SetExportThreadCount(v int64)`

SetExportThreadCount sets ExportThreadCount field to given value.

### HasExportThreadCount

`func (o *AddBackend200Response) HasExportThreadCount() bool`

HasExportThreadCount returns a boolean if a field has been set.

### GetDbImportCachePercent

`func (o *AddBackend200Response) GetDbImportCachePercent() int64`

GetDbImportCachePercent returns the DbImportCachePercent field if non-nil, zero value otherwise.

### GetDbImportCachePercentOk

`func (o *AddBackend200Response) GetDbImportCachePercentOk() (*int64, bool)`

GetDbImportCachePercentOk returns a tuple with the DbImportCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbImportCachePercent

`func (o *AddBackend200Response) SetDbImportCachePercent(v int64)`

SetDbImportCachePercent sets DbImportCachePercent field to given value.

### HasDbImportCachePercent

`func (o *AddBackend200Response) HasDbImportCachePercent() bool`

HasDbImportCachePercent returns a boolean if a field has been set.

### GetDbTxnWriteNoSync

`func (o *AddBackend200Response) GetDbTxnWriteNoSync() bool`

GetDbTxnWriteNoSync returns the DbTxnWriteNoSync field if non-nil, zero value otherwise.

### GetDbTxnWriteNoSyncOk

`func (o *AddBackend200Response) GetDbTxnWriteNoSyncOk() (*bool, bool)`

GetDbTxnWriteNoSyncOk returns a tuple with the DbTxnWriteNoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTxnWriteNoSync

`func (o *AddBackend200Response) SetDbTxnWriteNoSync(v bool)`

SetDbTxnWriteNoSync sets DbTxnWriteNoSync field to given value.

### HasDbTxnWriteNoSync

`func (o *AddBackend200Response) HasDbTxnWriteNoSync() bool`

HasDbTxnWriteNoSync returns a boolean if a field has been set.

### GetDeadlockRetryLimit

`func (o *AddBackend200Response) GetDeadlockRetryLimit() int64`

GetDeadlockRetryLimit returns the DeadlockRetryLimit field if non-nil, zero value otherwise.

### GetDeadlockRetryLimitOk

`func (o *AddBackend200Response) GetDeadlockRetryLimitOk() (*int64, bool)`

GetDeadlockRetryLimitOk returns a tuple with the DeadlockRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockRetryLimit

`func (o *AddBackend200Response) SetDeadlockRetryLimit(v int64)`

SetDeadlockRetryLimit sets DeadlockRetryLimit field to given value.

### HasDeadlockRetryLimit

`func (o *AddBackend200Response) HasDeadlockRetryLimit() bool`

HasDeadlockRetryLimit returns a boolean if a field has been set.

### GetExternalTxnDefaultBackendLockBehavior

`func (o *AddBackend200Response) GetExternalTxnDefaultBackendLockBehavior() EnumbackendExternalTxnDefaultBackendLockBehaviorProp`

GetExternalTxnDefaultBackendLockBehavior returns the ExternalTxnDefaultBackendLockBehavior field if non-nil, zero value otherwise.

### GetExternalTxnDefaultBackendLockBehaviorOk

`func (o *AddBackend200Response) GetExternalTxnDefaultBackendLockBehaviorOk() (*EnumbackendExternalTxnDefaultBackendLockBehaviorProp, bool)`

GetExternalTxnDefaultBackendLockBehaviorOk returns a tuple with the ExternalTxnDefaultBackendLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTxnDefaultBackendLockBehavior

`func (o *AddBackend200Response) SetExternalTxnDefaultBackendLockBehavior(v EnumbackendExternalTxnDefaultBackendLockBehaviorProp)`

SetExternalTxnDefaultBackendLockBehavior sets ExternalTxnDefaultBackendLockBehavior field to given value.

### HasExternalTxnDefaultBackendLockBehavior

`func (o *AddBackend200Response) HasExternalTxnDefaultBackendLockBehavior() bool`

HasExternalTxnDefaultBackendLockBehavior returns a boolean if a field has been set.

### GetSingleWriterLockBehavior

`func (o *AddBackend200Response) GetSingleWriterLockBehavior() EnumbackendSingleWriterLockBehaviorProp`

GetSingleWriterLockBehavior returns the SingleWriterLockBehavior field if non-nil, zero value otherwise.

### GetSingleWriterLockBehaviorOk

`func (o *AddBackend200Response) GetSingleWriterLockBehaviorOk() (*EnumbackendSingleWriterLockBehaviorProp, bool)`

GetSingleWriterLockBehaviorOk returns a tuple with the SingleWriterLockBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleWriterLockBehavior

`func (o *AddBackend200Response) SetSingleWriterLockBehavior(v EnumbackendSingleWriterLockBehaviorProp)`

SetSingleWriterLockBehavior sets SingleWriterLockBehavior field to given value.

### HasSingleWriterLockBehavior

`func (o *AddBackend200Response) HasSingleWriterLockBehavior() bool`

HasSingleWriterLockBehavior returns a boolean if a field has been set.

### GetSubtreeDeleteSizeLimit

`func (o *AddBackend200Response) GetSubtreeDeleteSizeLimit() int64`

GetSubtreeDeleteSizeLimit returns the SubtreeDeleteSizeLimit field if non-nil, zero value otherwise.

### GetSubtreeDeleteSizeLimitOk

`func (o *AddBackend200Response) GetSubtreeDeleteSizeLimitOk() (*int64, bool)`

GetSubtreeDeleteSizeLimitOk returns a tuple with the SubtreeDeleteSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeDeleteSizeLimit

`func (o *AddBackend200Response) SetSubtreeDeleteSizeLimit(v int64)`

SetSubtreeDeleteSizeLimit sets SubtreeDeleteSizeLimit field to given value.

### HasSubtreeDeleteSizeLimit

`func (o *AddBackend200Response) HasSubtreeDeleteSizeLimit() bool`

HasSubtreeDeleteSizeLimit returns a boolean if a field has been set.

### GetNumRecentChanges

`func (o *AddBackend200Response) GetNumRecentChanges() int64`

GetNumRecentChanges returns the NumRecentChanges field if non-nil, zero value otherwise.

### GetNumRecentChangesOk

`func (o *AddBackend200Response) GetNumRecentChangesOk() (*int64, bool)`

GetNumRecentChangesOk returns a tuple with the NumRecentChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecentChanges

`func (o *AddBackend200Response) SetNumRecentChanges(v int64)`

SetNumRecentChanges sets NumRecentChanges field to given value.

### HasNumRecentChanges

`func (o *AddBackend200Response) HasNumRecentChanges() bool`

HasNumRecentChanges returns a boolean if a field has been set.

### GetOfflineProcessDatabaseOpenTimeout

`func (o *AddBackend200Response) GetOfflineProcessDatabaseOpenTimeout() string`

GetOfflineProcessDatabaseOpenTimeout returns the OfflineProcessDatabaseOpenTimeout field if non-nil, zero value otherwise.

### GetOfflineProcessDatabaseOpenTimeoutOk

`func (o *AddBackend200Response) GetOfflineProcessDatabaseOpenTimeoutOk() (*string, bool)`

GetOfflineProcessDatabaseOpenTimeoutOk returns a tuple with the OfflineProcessDatabaseOpenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineProcessDatabaseOpenTimeout

`func (o *AddBackend200Response) SetOfflineProcessDatabaseOpenTimeout(v string)`

SetOfflineProcessDatabaseOpenTimeout sets OfflineProcessDatabaseOpenTimeout field to given value.

### HasOfflineProcessDatabaseOpenTimeout

`func (o *AddBackend200Response) HasOfflineProcessDatabaseOpenTimeout() bool`

HasOfflineProcessDatabaseOpenTimeout returns a boolean if a field has been set.

### GetBackendID

`func (o *AddBackend200Response) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *AddBackend200Response) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *AddBackend200Response) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetDescription

`func (o *AddBackend200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBackend200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBackend200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBackend200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddBackend200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddBackend200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddBackend200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *AddBackend200Response) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddBackend200Response) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddBackend200Response) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *AddBackend200Response) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *AddBackend200Response) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *AddBackend200Response) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *AddBackend200Response) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *AddBackend200Response) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *AddBackend200Response) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *AddBackend200Response) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *AddBackend200Response) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetNotificationManager

`func (o *AddBackend200Response) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *AddBackend200Response) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *AddBackend200Response) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *AddBackend200Response) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *AddBackend200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddBackend200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddBackend200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddBackend200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddBackend200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddBackend200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddBackend200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddBackend200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


