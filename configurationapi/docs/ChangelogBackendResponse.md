# ChangelogBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumchangelogBackendSchemaUrn**](EnumchangelogBackendSchemaUrn.md) |  | 
**Id** | **string** | Name of the Backend | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**DbDirectory** | Pointer to **string** | Specifies the path to the filesystem directory that is used to hold the Berkeley DB Java Edition database files containing the data for this backend. The files for this backend are stored in a sub-directory named after the backend-id. | [optional] 
**DbDirectoryPermissions** | Pointer to **string** | Specifies the permissions that should be applied to the directory containing the backend database files and to directories and files created during backup of the backend. | [optional] 
**DbCachePercent** | Pointer to **int32** | Specifies the percentage of JVM memory to allocate to the changelog database cache. | [optional] 
**JeProperty** | Pointer to **[]string** | Specifies the database and environment properties for the Berkeley DB Java Edition database for this changelog backend. | [optional] 
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
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewChangelogBackendResponse

`func NewChangelogBackendResponse(schemas []EnumchangelogBackendSchemaUrn, id string, baseDN []string, changelogMaximumAge string, backendID string, enabled bool, ) *ChangelogBackendResponse`

NewChangelogBackendResponse instantiates a new ChangelogBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogBackendResponseWithDefaults

`func NewChangelogBackendResponseWithDefaults() *ChangelogBackendResponse`

NewChangelogBackendResponseWithDefaults instantiates a new ChangelogBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ChangelogBackendResponse) GetSchemas() []EnumchangelogBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ChangelogBackendResponse) GetSchemasOk() (*[]EnumchangelogBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ChangelogBackendResponse) SetSchemas(v []EnumchangelogBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ChangelogBackendResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangelogBackendResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangelogBackendResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBaseDN

`func (o *ChangelogBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ChangelogBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ChangelogBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetDbDirectory

`func (o *ChangelogBackendResponse) GetDbDirectory() string`

GetDbDirectory returns the DbDirectory field if non-nil, zero value otherwise.

### GetDbDirectoryOk

`func (o *ChangelogBackendResponse) GetDbDirectoryOk() (*string, bool)`

GetDbDirectoryOk returns a tuple with the DbDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectory

`func (o *ChangelogBackendResponse) SetDbDirectory(v string)`

SetDbDirectory sets DbDirectory field to given value.

### HasDbDirectory

`func (o *ChangelogBackendResponse) HasDbDirectory() bool`

HasDbDirectory returns a boolean if a field has been set.

### GetDbDirectoryPermissions

`func (o *ChangelogBackendResponse) GetDbDirectoryPermissions() string`

GetDbDirectoryPermissions returns the DbDirectoryPermissions field if non-nil, zero value otherwise.

### GetDbDirectoryPermissionsOk

`func (o *ChangelogBackendResponse) GetDbDirectoryPermissionsOk() (*string, bool)`

GetDbDirectoryPermissionsOk returns a tuple with the DbDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDirectoryPermissions

`func (o *ChangelogBackendResponse) SetDbDirectoryPermissions(v string)`

SetDbDirectoryPermissions sets DbDirectoryPermissions field to given value.

### HasDbDirectoryPermissions

`func (o *ChangelogBackendResponse) HasDbDirectoryPermissions() bool`

HasDbDirectoryPermissions returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *ChangelogBackendResponse) GetDbCachePercent() int32`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *ChangelogBackendResponse) GetDbCachePercentOk() (*int32, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *ChangelogBackendResponse) SetDbCachePercent(v int32)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *ChangelogBackendResponse) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetJeProperty

`func (o *ChangelogBackendResponse) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *ChangelogBackendResponse) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *ChangelogBackendResponse) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *ChangelogBackendResponse) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetChangelogWriteBatchSize

`func (o *ChangelogBackendResponse) GetChangelogWriteBatchSize() int32`

GetChangelogWriteBatchSize returns the ChangelogWriteBatchSize field if non-nil, zero value otherwise.

### GetChangelogWriteBatchSizeOk

`func (o *ChangelogBackendResponse) GetChangelogWriteBatchSizeOk() (*int32, bool)`

GetChangelogWriteBatchSizeOk returns a tuple with the ChangelogWriteBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogWriteBatchSize

`func (o *ChangelogBackendResponse) SetChangelogWriteBatchSize(v int32)`

SetChangelogWriteBatchSize sets ChangelogWriteBatchSize field to given value.

### HasChangelogWriteBatchSize

`func (o *ChangelogBackendResponse) HasChangelogWriteBatchSize() bool`

HasChangelogWriteBatchSize returns a boolean if a field has been set.

### GetChangelogPurgeBatchSize

`func (o *ChangelogBackendResponse) GetChangelogPurgeBatchSize() int32`

GetChangelogPurgeBatchSize returns the ChangelogPurgeBatchSize field if non-nil, zero value otherwise.

### GetChangelogPurgeBatchSizeOk

`func (o *ChangelogBackendResponse) GetChangelogPurgeBatchSizeOk() (*int32, bool)`

GetChangelogPurgeBatchSizeOk returns a tuple with the ChangelogPurgeBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPurgeBatchSize

`func (o *ChangelogBackendResponse) SetChangelogPurgeBatchSize(v int32)`

SetChangelogPurgeBatchSize sets ChangelogPurgeBatchSize field to given value.

### HasChangelogPurgeBatchSize

`func (o *ChangelogBackendResponse) HasChangelogPurgeBatchSize() bool`

HasChangelogPurgeBatchSize returns a boolean if a field has been set.

### GetChangelogWriteQueueCapacity

`func (o *ChangelogBackendResponse) GetChangelogWriteQueueCapacity() int32`

GetChangelogWriteQueueCapacity returns the ChangelogWriteQueueCapacity field if non-nil, zero value otherwise.

### GetChangelogWriteQueueCapacityOk

`func (o *ChangelogBackendResponse) GetChangelogWriteQueueCapacityOk() (*int32, bool)`

GetChangelogWriteQueueCapacityOk returns a tuple with the ChangelogWriteQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogWriteQueueCapacity

`func (o *ChangelogBackendResponse) SetChangelogWriteQueueCapacity(v int32)`

SetChangelogWriteQueueCapacity sets ChangelogWriteQueueCapacity field to given value.

### HasChangelogWriteQueueCapacity

`func (o *ChangelogBackendResponse) HasChangelogWriteQueueCapacity() bool`

HasChangelogWriteQueueCapacity returns a boolean if a field has been set.

### GetIndexIncludeAttribute

`func (o *ChangelogBackendResponse) GetIndexIncludeAttribute() []string`

GetIndexIncludeAttribute returns the IndexIncludeAttribute field if non-nil, zero value otherwise.

### GetIndexIncludeAttributeOk

`func (o *ChangelogBackendResponse) GetIndexIncludeAttributeOk() (*[]string, bool)`

GetIndexIncludeAttributeOk returns a tuple with the IndexIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIncludeAttribute

`func (o *ChangelogBackendResponse) SetIndexIncludeAttribute(v []string)`

SetIndexIncludeAttribute sets IndexIncludeAttribute field to given value.

### HasIndexIncludeAttribute

`func (o *ChangelogBackendResponse) HasIndexIncludeAttribute() bool`

HasIndexIncludeAttribute returns a boolean if a field has been set.

### GetIndexExcludeAttribute

`func (o *ChangelogBackendResponse) GetIndexExcludeAttribute() []string`

GetIndexExcludeAttribute returns the IndexExcludeAttribute field if non-nil, zero value otherwise.

### GetIndexExcludeAttributeOk

`func (o *ChangelogBackendResponse) GetIndexExcludeAttributeOk() (*[]string, bool)`

GetIndexExcludeAttributeOk returns a tuple with the IndexExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexExcludeAttribute

`func (o *ChangelogBackendResponse) SetIndexExcludeAttribute(v []string)`

SetIndexExcludeAttribute sets IndexExcludeAttribute field to given value.

### HasIndexExcludeAttribute

`func (o *ChangelogBackendResponse) HasIndexExcludeAttribute() bool`

HasIndexExcludeAttribute returns a boolean if a field has been set.

### GetChangelogMaximumAge

`func (o *ChangelogBackendResponse) GetChangelogMaximumAge() string`

GetChangelogMaximumAge returns the ChangelogMaximumAge field if non-nil, zero value otherwise.

### GetChangelogMaximumAgeOk

`func (o *ChangelogBackendResponse) GetChangelogMaximumAgeOk() (*string, bool)`

GetChangelogMaximumAgeOk returns a tuple with the ChangelogMaximumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogMaximumAge

`func (o *ChangelogBackendResponse) SetChangelogMaximumAge(v string)`

SetChangelogMaximumAge sets ChangelogMaximumAge field to given value.


### GetTargetDatabaseSize

`func (o *ChangelogBackendResponse) GetTargetDatabaseSize() string`

GetTargetDatabaseSize returns the TargetDatabaseSize field if non-nil, zero value otherwise.

### GetTargetDatabaseSizeOk

`func (o *ChangelogBackendResponse) GetTargetDatabaseSizeOk() (*string, bool)`

GetTargetDatabaseSizeOk returns a tuple with the TargetDatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatabaseSize

`func (o *ChangelogBackendResponse) SetTargetDatabaseSize(v string)`

SetTargetDatabaseSize sets TargetDatabaseSize field to given value.

### HasTargetDatabaseSize

`func (o *ChangelogBackendResponse) HasTargetDatabaseSize() bool`

HasTargetDatabaseSize returns a boolean if a field has been set.

### GetChangelogEntryIncludeBaseDN

`func (o *ChangelogBackendResponse) GetChangelogEntryIncludeBaseDN() []string`

GetChangelogEntryIncludeBaseDN returns the ChangelogEntryIncludeBaseDN field if non-nil, zero value otherwise.

### GetChangelogEntryIncludeBaseDNOk

`func (o *ChangelogBackendResponse) GetChangelogEntryIncludeBaseDNOk() (*[]string, bool)`

GetChangelogEntryIncludeBaseDNOk returns a tuple with the ChangelogEntryIncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryIncludeBaseDN

`func (o *ChangelogBackendResponse) SetChangelogEntryIncludeBaseDN(v []string)`

SetChangelogEntryIncludeBaseDN sets ChangelogEntryIncludeBaseDN field to given value.

### HasChangelogEntryIncludeBaseDN

`func (o *ChangelogBackendResponse) HasChangelogEntryIncludeBaseDN() bool`

HasChangelogEntryIncludeBaseDN returns a boolean if a field has been set.

### GetChangelogEntryExcludeBaseDN

`func (o *ChangelogBackendResponse) GetChangelogEntryExcludeBaseDN() []string`

GetChangelogEntryExcludeBaseDN returns the ChangelogEntryExcludeBaseDN field if non-nil, zero value otherwise.

### GetChangelogEntryExcludeBaseDNOk

`func (o *ChangelogBackendResponse) GetChangelogEntryExcludeBaseDNOk() (*[]string, bool)`

GetChangelogEntryExcludeBaseDNOk returns a tuple with the ChangelogEntryExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryExcludeBaseDN

`func (o *ChangelogBackendResponse) SetChangelogEntryExcludeBaseDN(v []string)`

SetChangelogEntryExcludeBaseDN sets ChangelogEntryExcludeBaseDN field to given value.

### HasChangelogEntryExcludeBaseDN

`func (o *ChangelogBackendResponse) HasChangelogEntryExcludeBaseDN() bool`

HasChangelogEntryExcludeBaseDN returns a boolean if a field has been set.

### GetChangelogEntryIncludeFilter

`func (o *ChangelogBackendResponse) GetChangelogEntryIncludeFilter() []string`

GetChangelogEntryIncludeFilter returns the ChangelogEntryIncludeFilter field if non-nil, zero value otherwise.

### GetChangelogEntryIncludeFilterOk

`func (o *ChangelogBackendResponse) GetChangelogEntryIncludeFilterOk() (*[]string, bool)`

GetChangelogEntryIncludeFilterOk returns a tuple with the ChangelogEntryIncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryIncludeFilter

`func (o *ChangelogBackendResponse) SetChangelogEntryIncludeFilter(v []string)`

SetChangelogEntryIncludeFilter sets ChangelogEntryIncludeFilter field to given value.

### HasChangelogEntryIncludeFilter

`func (o *ChangelogBackendResponse) HasChangelogEntryIncludeFilter() bool`

HasChangelogEntryIncludeFilter returns a boolean if a field has been set.

### GetChangelogEntryExcludeFilter

`func (o *ChangelogBackendResponse) GetChangelogEntryExcludeFilter() []string`

GetChangelogEntryExcludeFilter returns the ChangelogEntryExcludeFilter field if non-nil, zero value otherwise.

### GetChangelogEntryExcludeFilterOk

`func (o *ChangelogBackendResponse) GetChangelogEntryExcludeFilterOk() (*[]string, bool)`

GetChangelogEntryExcludeFilterOk returns a tuple with the ChangelogEntryExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogEntryExcludeFilter

`func (o *ChangelogBackendResponse) SetChangelogEntryExcludeFilter(v []string)`

SetChangelogEntryExcludeFilter sets ChangelogEntryExcludeFilter field to given value.

### HasChangelogEntryExcludeFilter

`func (o *ChangelogBackendResponse) HasChangelogEntryExcludeFilter() bool`

HasChangelogEntryExcludeFilter returns a boolean if a field has been set.

### GetChangelogIncludeAttribute

`func (o *ChangelogBackendResponse) GetChangelogIncludeAttribute() []string`

GetChangelogIncludeAttribute returns the ChangelogIncludeAttribute field if non-nil, zero value otherwise.

### GetChangelogIncludeAttributeOk

`func (o *ChangelogBackendResponse) GetChangelogIncludeAttributeOk() (*[]string, bool)`

GetChangelogIncludeAttributeOk returns a tuple with the ChangelogIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogIncludeAttribute

`func (o *ChangelogBackendResponse) SetChangelogIncludeAttribute(v []string)`

SetChangelogIncludeAttribute sets ChangelogIncludeAttribute field to given value.

### HasChangelogIncludeAttribute

`func (o *ChangelogBackendResponse) HasChangelogIncludeAttribute() bool`

HasChangelogIncludeAttribute returns a boolean if a field has been set.

### GetChangelogExcludeAttribute

`func (o *ChangelogBackendResponse) GetChangelogExcludeAttribute() []string`

GetChangelogExcludeAttribute returns the ChangelogExcludeAttribute field if non-nil, zero value otherwise.

### GetChangelogExcludeAttributeOk

`func (o *ChangelogBackendResponse) GetChangelogExcludeAttributeOk() (*[]string, bool)`

GetChangelogExcludeAttributeOk returns a tuple with the ChangelogExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogExcludeAttribute

`func (o *ChangelogBackendResponse) SetChangelogExcludeAttribute(v []string)`

SetChangelogExcludeAttribute sets ChangelogExcludeAttribute field to given value.

### HasChangelogExcludeAttribute

`func (o *ChangelogBackendResponse) HasChangelogExcludeAttribute() bool`

HasChangelogExcludeAttribute returns a boolean if a field has been set.

### GetChangelogDeletedEntryIncludeAttribute

`func (o *ChangelogBackendResponse) GetChangelogDeletedEntryIncludeAttribute() []string`

GetChangelogDeletedEntryIncludeAttribute returns the ChangelogDeletedEntryIncludeAttribute field if non-nil, zero value otherwise.

### GetChangelogDeletedEntryIncludeAttributeOk

`func (o *ChangelogBackendResponse) GetChangelogDeletedEntryIncludeAttributeOk() (*[]string, bool)`

GetChangelogDeletedEntryIncludeAttributeOk returns a tuple with the ChangelogDeletedEntryIncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogDeletedEntryIncludeAttribute

`func (o *ChangelogBackendResponse) SetChangelogDeletedEntryIncludeAttribute(v []string)`

SetChangelogDeletedEntryIncludeAttribute sets ChangelogDeletedEntryIncludeAttribute field to given value.

### HasChangelogDeletedEntryIncludeAttribute

`func (o *ChangelogBackendResponse) HasChangelogDeletedEntryIncludeAttribute() bool`

HasChangelogDeletedEntryIncludeAttribute returns a boolean if a field has been set.

### GetChangelogDeletedEntryExcludeAttribute

`func (o *ChangelogBackendResponse) GetChangelogDeletedEntryExcludeAttribute() []string`

GetChangelogDeletedEntryExcludeAttribute returns the ChangelogDeletedEntryExcludeAttribute field if non-nil, zero value otherwise.

### GetChangelogDeletedEntryExcludeAttributeOk

`func (o *ChangelogBackendResponse) GetChangelogDeletedEntryExcludeAttributeOk() (*[]string, bool)`

GetChangelogDeletedEntryExcludeAttributeOk returns a tuple with the ChangelogDeletedEntryExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogDeletedEntryExcludeAttribute

`func (o *ChangelogBackendResponse) SetChangelogDeletedEntryExcludeAttribute(v []string)`

SetChangelogDeletedEntryExcludeAttribute sets ChangelogDeletedEntryExcludeAttribute field to given value.

### HasChangelogDeletedEntryExcludeAttribute

`func (o *ChangelogBackendResponse) HasChangelogDeletedEntryExcludeAttribute() bool`

HasChangelogDeletedEntryExcludeAttribute returns a boolean if a field has been set.

### GetChangelogIncludeKeyAttribute

`func (o *ChangelogBackendResponse) GetChangelogIncludeKeyAttribute() []string`

GetChangelogIncludeKeyAttribute returns the ChangelogIncludeKeyAttribute field if non-nil, zero value otherwise.

### GetChangelogIncludeKeyAttributeOk

`func (o *ChangelogBackendResponse) GetChangelogIncludeKeyAttributeOk() (*[]string, bool)`

GetChangelogIncludeKeyAttributeOk returns a tuple with the ChangelogIncludeKeyAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogIncludeKeyAttribute

`func (o *ChangelogBackendResponse) SetChangelogIncludeKeyAttribute(v []string)`

SetChangelogIncludeKeyAttribute sets ChangelogIncludeKeyAttribute field to given value.

### HasChangelogIncludeKeyAttribute

`func (o *ChangelogBackendResponse) HasChangelogIncludeKeyAttribute() bool`

HasChangelogIncludeKeyAttribute returns a boolean if a field has been set.

### GetChangelogMaxBeforeAfterValues

`func (o *ChangelogBackendResponse) GetChangelogMaxBeforeAfterValues() int32`

GetChangelogMaxBeforeAfterValues returns the ChangelogMaxBeforeAfterValues field if non-nil, zero value otherwise.

### GetChangelogMaxBeforeAfterValuesOk

`func (o *ChangelogBackendResponse) GetChangelogMaxBeforeAfterValuesOk() (*int32, bool)`

GetChangelogMaxBeforeAfterValuesOk returns a tuple with the ChangelogMaxBeforeAfterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogMaxBeforeAfterValues

`func (o *ChangelogBackendResponse) SetChangelogMaxBeforeAfterValues(v int32)`

SetChangelogMaxBeforeAfterValues sets ChangelogMaxBeforeAfterValues field to given value.

### HasChangelogMaxBeforeAfterValues

`func (o *ChangelogBackendResponse) HasChangelogMaxBeforeAfterValues() bool`

HasChangelogMaxBeforeAfterValues returns a boolean if a field has been set.

### GetWriteLastmodAttributes

`func (o *ChangelogBackendResponse) GetWriteLastmodAttributes() bool`

GetWriteLastmodAttributes returns the WriteLastmodAttributes field if non-nil, zero value otherwise.

### GetWriteLastmodAttributesOk

`func (o *ChangelogBackendResponse) GetWriteLastmodAttributesOk() (*bool, bool)`

GetWriteLastmodAttributesOk returns a tuple with the WriteLastmodAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLastmodAttributes

`func (o *ChangelogBackendResponse) SetWriteLastmodAttributes(v bool)`

SetWriteLastmodAttributes sets WriteLastmodAttributes field to given value.

### HasWriteLastmodAttributes

`func (o *ChangelogBackendResponse) HasWriteLastmodAttributes() bool`

HasWriteLastmodAttributes returns a boolean if a field has been set.

### GetUseReversibleForm

`func (o *ChangelogBackendResponse) GetUseReversibleForm() bool`

GetUseReversibleForm returns the UseReversibleForm field if non-nil, zero value otherwise.

### GetUseReversibleFormOk

`func (o *ChangelogBackendResponse) GetUseReversibleFormOk() (*bool, bool)`

GetUseReversibleFormOk returns a tuple with the UseReversibleForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReversibleForm

`func (o *ChangelogBackendResponse) SetUseReversibleForm(v bool)`

SetUseReversibleForm sets UseReversibleForm field to given value.

### HasUseReversibleForm

`func (o *ChangelogBackendResponse) HasUseReversibleForm() bool`

HasUseReversibleForm returns a boolean if a field has been set.

### GetIncludeVirtualAttributes

`func (o *ChangelogBackendResponse) GetIncludeVirtualAttributes() []EnumbackendIncludeVirtualAttributesProp`

GetIncludeVirtualAttributes returns the IncludeVirtualAttributes field if non-nil, zero value otherwise.

### GetIncludeVirtualAttributesOk

`func (o *ChangelogBackendResponse) GetIncludeVirtualAttributesOk() (*[]EnumbackendIncludeVirtualAttributesProp, bool)`

GetIncludeVirtualAttributesOk returns a tuple with the IncludeVirtualAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVirtualAttributes

`func (o *ChangelogBackendResponse) SetIncludeVirtualAttributes(v []EnumbackendIncludeVirtualAttributesProp)`

SetIncludeVirtualAttributes sets IncludeVirtualAttributes field to given value.

### HasIncludeVirtualAttributes

`func (o *ChangelogBackendResponse) HasIncludeVirtualAttributes() bool`

HasIncludeVirtualAttributes returns a boolean if a field has been set.

### GetApplyAccessControlsToChangelogEntryContents

`func (o *ChangelogBackendResponse) GetApplyAccessControlsToChangelogEntryContents() bool`

GetApplyAccessControlsToChangelogEntryContents returns the ApplyAccessControlsToChangelogEntryContents field if non-nil, zero value otherwise.

### GetApplyAccessControlsToChangelogEntryContentsOk

`func (o *ChangelogBackendResponse) GetApplyAccessControlsToChangelogEntryContentsOk() (*bool, bool)`

GetApplyAccessControlsToChangelogEntryContentsOk returns a tuple with the ApplyAccessControlsToChangelogEntryContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAccessControlsToChangelogEntryContents

`func (o *ChangelogBackendResponse) SetApplyAccessControlsToChangelogEntryContents(v bool)`

SetApplyAccessControlsToChangelogEntryContents sets ApplyAccessControlsToChangelogEntryContents field to given value.

### HasApplyAccessControlsToChangelogEntryContents

`func (o *ChangelogBackendResponse) HasApplyAccessControlsToChangelogEntryContents() bool`

HasApplyAccessControlsToChangelogEntryContents returns a boolean if a field has been set.

### GetReportExcludedChangelogAttributes

`func (o *ChangelogBackendResponse) GetReportExcludedChangelogAttributes() EnumbackendReportExcludedChangelogAttributesProp`

GetReportExcludedChangelogAttributes returns the ReportExcludedChangelogAttributes field if non-nil, zero value otherwise.

### GetReportExcludedChangelogAttributesOk

`func (o *ChangelogBackendResponse) GetReportExcludedChangelogAttributesOk() (*EnumbackendReportExcludedChangelogAttributesProp, bool)`

GetReportExcludedChangelogAttributesOk returns a tuple with the ReportExcludedChangelogAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExcludedChangelogAttributes

`func (o *ChangelogBackendResponse) SetReportExcludedChangelogAttributes(v EnumbackendReportExcludedChangelogAttributesProp)`

SetReportExcludedChangelogAttributes sets ReportExcludedChangelogAttributes field to given value.

### HasReportExcludedChangelogAttributes

`func (o *ChangelogBackendResponse) HasReportExcludedChangelogAttributes() bool`

HasReportExcludedChangelogAttributes returns a boolean if a field has been set.

### GetSoftDeleteEntryIncludedOperation

`func (o *ChangelogBackendResponse) GetSoftDeleteEntryIncludedOperation() []EnumbackendSoftDeleteEntryIncludedOperationProp`

GetSoftDeleteEntryIncludedOperation returns the SoftDeleteEntryIncludedOperation field if non-nil, zero value otherwise.

### GetSoftDeleteEntryIncludedOperationOk

`func (o *ChangelogBackendResponse) GetSoftDeleteEntryIncludedOperationOk() (*[]EnumbackendSoftDeleteEntryIncludedOperationProp, bool)`

GetSoftDeleteEntryIncludedOperationOk returns a tuple with the SoftDeleteEntryIncludedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteEntryIncludedOperation

`func (o *ChangelogBackendResponse) SetSoftDeleteEntryIncludedOperation(v []EnumbackendSoftDeleteEntryIncludedOperationProp)`

SetSoftDeleteEntryIncludedOperation sets SoftDeleteEntryIncludedOperation field to given value.

### HasSoftDeleteEntryIncludedOperation

`func (o *ChangelogBackendResponse) HasSoftDeleteEntryIncludedOperation() bool`

HasSoftDeleteEntryIncludedOperation returns a boolean if a field has been set.

### GetBackendID

`func (o *ChangelogBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *ChangelogBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *ChangelogBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetDescription

`func (o *ChangelogBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangelogBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangelogBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangelogBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ChangelogBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ChangelogBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ChangelogBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *ChangelogBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *ChangelogBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *ChangelogBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *ChangelogBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *ChangelogBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *ChangelogBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *ChangelogBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *ChangelogBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetNotificationManager

`func (o *ChangelogBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *ChangelogBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *ChangelogBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *ChangelogBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *ChangelogBackendResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChangelogBackendResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChangelogBackendResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChangelogBackendResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ChangelogBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ChangelogBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ChangelogBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ChangelogBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


