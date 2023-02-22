# AddPeriodicStatsLoggerPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumperiodicStatsLoggerPluginSchemaUrn**](EnumperiodicStatsLoggerPluginSchemaUrn.md) |  | 
**LogInterval** | **string** | The duration between statistics collection and logging. A new line is logged to the output for each interval. Setting this value too small can have an impact on performance. | 
**CollectionInterval** | **string** | Some of the calculated statistics, such as the average and maximum queue sizes, can use multiple samples within a log interval. This value controls how often samples are gathered. It should be a multiple of the log-interval. | 
**SuppressIfIdle** | **bool** | If the server is idle during the specified interval, then do not log any output if this property is set to true. The server is idle if during the interval, no new connections were established, no operations were processed, and no operations are pending. | 
**HeaderPrefixPerColumn** | Pointer to **bool** | This property controls whether the header prefix, which applies to a group of columns, appears at the start of each column header or only the first column in a group. | [optional] 
**EmptyInsteadOfZero** | Pointer to **bool** | This property controls whether a value in the output is shown as empty if the value is zero. | [optional] 
**LinesBetweenHeader** | **int32** | The number of lines to log between logging the header line that summarizes the columns in the table. | 
**IncludedLDAPStat** | Pointer to [**[]EnumpluginIncludedLDAPStatProp**](EnumpluginIncludedLDAPStatProp.md) |  | [optional] 
**IncludedResourceStat** | Pointer to [**[]EnumpluginIncludedResourceStatProp**](EnumpluginIncludedResourceStatProp.md) |  | [optional] 
**HistogramFormat** | [**EnumpluginHistogramFormatProp**](EnumpluginHistogramFormatProp.md) |  | 
**HistogramOpType** | Pointer to [**[]EnumpluginHistogramOpTypeProp**](EnumpluginHistogramOpTypeProp.md) |  | [optional] 
**PerApplicationLDAPStats** | Pointer to [**EnumpluginPerApplicationLDAPStatsProp**](EnumpluginPerApplicationLDAPStatsProp.md) |  | [optional] 
**StatusSummaryInfo** | Pointer to [**EnumpluginStatusSummaryInfoProp**](EnumpluginStatusSummaryInfoProp.md) |  | [optional] 
**LdapChangelogInfo** | Pointer to [**EnumpluginLdapChangelogInfoProp**](EnumpluginLdapChangelogInfoProp.md) |  | [optional] 
**GaugeInfo** | Pointer to [**EnumpluginGaugeInfoProp**](EnumpluginGaugeInfoProp.md) |  | [optional] 
**LogFileFormat** | Pointer to [**EnumpluginLogFileFormatProp**](EnumpluginLogFileFormatProp.md) |  | [optional] 
**LogFile** | **string** | The file name to use for the log files generated by the Periodic Stats Logger Plugin. The path to the file can be specified either as relative to the server root or as an absolute path. | 
**LogFilePermissions** | **string** | The UNIX permissions of the log files created by this Periodic Stats Logger Plugin. | 
**Append** | Pointer to **bool** | Specifies whether to append to existing log files. | [optional] 
**RotationPolicy** | **[]string** | The rotation policy to use for the Periodic Stats Logger Plugin . | 
**RotationListener** | Pointer to **[]string** | A listener that should be notified whenever a log file is rotated out of service. | [optional] 
**RetentionPolicy** | **[]string** | The retention policy to use for the Periodic Stats Logger Plugin . | 
**LoggingErrorBehavior** | Pointer to [**EnumpluginLoggingErrorBehaviorProp**](EnumpluginLoggingErrorBehaviorProp.md) |  | [optional] 
**LocalDBBackendInfo** | Pointer to [**EnumpluginLocalDBBackendInfoProp**](EnumpluginLocalDBBackendInfoProp.md) |  | [optional] 
**ReplicationInfo** | Pointer to [**EnumpluginReplicationInfoProp**](EnumpluginReplicationInfoProp.md) |  | [optional] 
**EntryCacheInfo** | Pointer to [**EnumpluginEntryCacheInfoProp**](EnumpluginEntryCacheInfoProp.md) |  | [optional] 
**HostInfo** | Pointer to [**[]EnumpluginHostInfoProp**](EnumpluginHostInfoProp.md) |  | [optional] 
**IncludedLDAPApplication** | Pointer to **[]string** | If statistics should not be included for all applications, this property names the subset of applications that should be included. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewAddPeriodicStatsLoggerPluginRequest

`func NewAddPeriodicStatsLoggerPluginRequest(pluginName string, schemas []EnumperiodicStatsLoggerPluginSchemaUrn, logInterval string, collectionInterval string, suppressIfIdle bool, linesBetweenHeader int32, histogramFormat EnumpluginHistogramFormatProp, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, enabled bool, ) *AddPeriodicStatsLoggerPluginRequest`

NewAddPeriodicStatsLoggerPluginRequest instantiates a new AddPeriodicStatsLoggerPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPeriodicStatsLoggerPluginRequestWithDefaults

`func NewAddPeriodicStatsLoggerPluginRequestWithDefaults() *AddPeriodicStatsLoggerPluginRequest`

NewAddPeriodicStatsLoggerPluginRequestWithDefaults instantiates a new AddPeriodicStatsLoggerPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddPeriodicStatsLoggerPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddPeriodicStatsLoggerPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddPeriodicStatsLoggerPluginRequest) GetSchemas() []EnumperiodicStatsLoggerPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetSchemasOk() (*[]EnumperiodicStatsLoggerPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPeriodicStatsLoggerPluginRequest) SetSchemas(v []EnumperiodicStatsLoggerPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogInterval

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogInterval() string`

GetLogInterval returns the LogInterval field if non-nil, zero value otherwise.

### GetLogIntervalOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogIntervalOk() (*string, bool)`

GetLogIntervalOk returns a tuple with the LogInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInterval

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLogInterval(v string)`

SetLogInterval sets LogInterval field to given value.


### GetCollectionInterval

`func (o *AddPeriodicStatsLoggerPluginRequest) GetCollectionInterval() string`

GetCollectionInterval returns the CollectionInterval field if non-nil, zero value otherwise.

### GetCollectionIntervalOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetCollectionIntervalOk() (*string, bool)`

GetCollectionIntervalOk returns a tuple with the CollectionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInterval

`func (o *AddPeriodicStatsLoggerPluginRequest) SetCollectionInterval(v string)`

SetCollectionInterval sets CollectionInterval field to given value.


### GetSuppressIfIdle

`func (o *AddPeriodicStatsLoggerPluginRequest) GetSuppressIfIdle() bool`

GetSuppressIfIdle returns the SuppressIfIdle field if non-nil, zero value otherwise.

### GetSuppressIfIdleOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetSuppressIfIdleOk() (*bool, bool)`

GetSuppressIfIdleOk returns a tuple with the SuppressIfIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressIfIdle

`func (o *AddPeriodicStatsLoggerPluginRequest) SetSuppressIfIdle(v bool)`

SetSuppressIfIdle sets SuppressIfIdle field to given value.


### GetHeaderPrefixPerColumn

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHeaderPrefixPerColumn() bool`

GetHeaderPrefixPerColumn returns the HeaderPrefixPerColumn field if non-nil, zero value otherwise.

### GetHeaderPrefixPerColumnOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHeaderPrefixPerColumnOk() (*bool, bool)`

GetHeaderPrefixPerColumnOk returns a tuple with the HeaderPrefixPerColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefixPerColumn

`func (o *AddPeriodicStatsLoggerPluginRequest) SetHeaderPrefixPerColumn(v bool)`

SetHeaderPrefixPerColumn sets HeaderPrefixPerColumn field to given value.

### HasHeaderPrefixPerColumn

`func (o *AddPeriodicStatsLoggerPluginRequest) HasHeaderPrefixPerColumn() bool`

HasHeaderPrefixPerColumn returns a boolean if a field has been set.

### GetEmptyInsteadOfZero

`func (o *AddPeriodicStatsLoggerPluginRequest) GetEmptyInsteadOfZero() bool`

GetEmptyInsteadOfZero returns the EmptyInsteadOfZero field if non-nil, zero value otherwise.

### GetEmptyInsteadOfZeroOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetEmptyInsteadOfZeroOk() (*bool, bool)`

GetEmptyInsteadOfZeroOk returns a tuple with the EmptyInsteadOfZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyInsteadOfZero

`func (o *AddPeriodicStatsLoggerPluginRequest) SetEmptyInsteadOfZero(v bool)`

SetEmptyInsteadOfZero sets EmptyInsteadOfZero field to given value.

### HasEmptyInsteadOfZero

`func (o *AddPeriodicStatsLoggerPluginRequest) HasEmptyInsteadOfZero() bool`

HasEmptyInsteadOfZero returns a boolean if a field has been set.

### GetLinesBetweenHeader

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLinesBetweenHeader() int32`

GetLinesBetweenHeader returns the LinesBetweenHeader field if non-nil, zero value otherwise.

### GetLinesBetweenHeaderOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLinesBetweenHeaderOk() (*int32, bool)`

GetLinesBetweenHeaderOk returns a tuple with the LinesBetweenHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesBetweenHeader

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLinesBetweenHeader(v int32)`

SetLinesBetweenHeader sets LinesBetweenHeader field to given value.


### GetIncludedLDAPStat

`func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPStat() []EnumpluginIncludedLDAPStatProp`

GetIncludedLDAPStat returns the IncludedLDAPStat field if non-nil, zero value otherwise.

### GetIncludedLDAPStatOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPStatOk() (*[]EnumpluginIncludedLDAPStatProp, bool)`

GetIncludedLDAPStatOk returns a tuple with the IncludedLDAPStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPStat

`func (o *AddPeriodicStatsLoggerPluginRequest) SetIncludedLDAPStat(v []EnumpluginIncludedLDAPStatProp)`

SetIncludedLDAPStat sets IncludedLDAPStat field to given value.

### HasIncludedLDAPStat

`func (o *AddPeriodicStatsLoggerPluginRequest) HasIncludedLDAPStat() bool`

HasIncludedLDAPStat returns a boolean if a field has been set.

### GetIncludedResourceStat

`func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedResourceStat() []EnumpluginIncludedResourceStatProp`

GetIncludedResourceStat returns the IncludedResourceStat field if non-nil, zero value otherwise.

### GetIncludedResourceStatOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedResourceStatOk() (*[]EnumpluginIncludedResourceStatProp, bool)`

GetIncludedResourceStatOk returns a tuple with the IncludedResourceStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedResourceStat

`func (o *AddPeriodicStatsLoggerPluginRequest) SetIncludedResourceStat(v []EnumpluginIncludedResourceStatProp)`

SetIncludedResourceStat sets IncludedResourceStat field to given value.

### HasIncludedResourceStat

`func (o *AddPeriodicStatsLoggerPluginRequest) HasIncludedResourceStat() bool`

HasIncludedResourceStat returns a boolean if a field has been set.

### GetHistogramFormat

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramFormat() EnumpluginHistogramFormatProp`

GetHistogramFormat returns the HistogramFormat field if non-nil, zero value otherwise.

### GetHistogramFormatOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramFormatOk() (*EnumpluginHistogramFormatProp, bool)`

GetHistogramFormatOk returns a tuple with the HistogramFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramFormat

`func (o *AddPeriodicStatsLoggerPluginRequest) SetHistogramFormat(v EnumpluginHistogramFormatProp)`

SetHistogramFormat sets HistogramFormat field to given value.


### GetHistogramOpType

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramOpType() []EnumpluginHistogramOpTypeProp`

GetHistogramOpType returns the HistogramOpType field if non-nil, zero value otherwise.

### GetHistogramOpTypeOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramOpTypeOk() (*[]EnumpluginHistogramOpTypeProp, bool)`

GetHistogramOpTypeOk returns a tuple with the HistogramOpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramOpType

`func (o *AddPeriodicStatsLoggerPluginRequest) SetHistogramOpType(v []EnumpluginHistogramOpTypeProp)`

SetHistogramOpType sets HistogramOpType field to given value.

### HasHistogramOpType

`func (o *AddPeriodicStatsLoggerPluginRequest) HasHistogramOpType() bool`

HasHistogramOpType returns a boolean if a field has been set.

### GetPerApplicationLDAPStats

`func (o *AddPeriodicStatsLoggerPluginRequest) GetPerApplicationLDAPStats() EnumpluginPerApplicationLDAPStatsProp`

GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field if non-nil, zero value otherwise.

### GetPerApplicationLDAPStatsOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetPerApplicationLDAPStatsOk() (*EnumpluginPerApplicationLDAPStatsProp, bool)`

GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerApplicationLDAPStats

`func (o *AddPeriodicStatsLoggerPluginRequest) SetPerApplicationLDAPStats(v EnumpluginPerApplicationLDAPStatsProp)`

SetPerApplicationLDAPStats sets PerApplicationLDAPStats field to given value.

### HasPerApplicationLDAPStats

`func (o *AddPeriodicStatsLoggerPluginRequest) HasPerApplicationLDAPStats() bool`

HasPerApplicationLDAPStats returns a boolean if a field has been set.

### GetStatusSummaryInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp`

GetStatusSummaryInfo returns the StatusSummaryInfo field if non-nil, zero value otherwise.

### GetStatusSummaryInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool)`

GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSummaryInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp)`

SetStatusSummaryInfo sets StatusSummaryInfo field to given value.

### HasStatusSummaryInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasStatusSummaryInfo() bool`

HasStatusSummaryInfo returns a boolean if a field has been set.

### GetLdapChangelogInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp`

GetLdapChangelogInfo returns the LdapChangelogInfo field if non-nil, zero value otherwise.

### GetLdapChangelogInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool)`

GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapChangelogInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp)`

SetLdapChangelogInfo sets LdapChangelogInfo field to given value.

### HasLdapChangelogInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasLdapChangelogInfo() bool`

HasLdapChangelogInfo returns a boolean if a field has been set.

### GetGaugeInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetGaugeInfo() EnumpluginGaugeInfoProp`

GetGaugeInfo returns the GaugeInfo field if non-nil, zero value otherwise.

### GetGaugeInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetGaugeInfoOk() (*EnumpluginGaugeInfoProp, bool)`

GetGaugeInfoOk returns a tuple with the GaugeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetGaugeInfo(v EnumpluginGaugeInfoProp)`

SetGaugeInfo sets GaugeInfo field to given value.

### HasGaugeInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasGaugeInfo() bool`

HasGaugeInfo returns a boolean if a field has been set.

### GetLogFileFormat

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFileFormat() EnumpluginLogFileFormatProp`

GetLogFileFormat returns the LogFileFormat field if non-nil, zero value otherwise.

### GetLogFileFormatOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFileFormatOk() (*EnumpluginLogFileFormatProp, bool)`

GetLogFileFormatOk returns a tuple with the LogFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileFormat

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLogFileFormat(v EnumpluginLogFileFormatProp)`

SetLogFileFormat sets LogFileFormat field to given value.

### HasLogFileFormat

`func (o *AddPeriodicStatsLoggerPluginRequest) HasLogFileFormat() bool`

HasLogFileFormat returns a boolean if a field has been set.

### GetLogFile

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetLogFilePermissions

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFilePermissions() string`

GetLogFilePermissions returns the LogFilePermissions field if non-nil, zero value otherwise.

### GetLogFilePermissionsOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFilePermissionsOk() (*string, bool)`

GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFilePermissions

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLogFilePermissions(v string)`

SetLogFilePermissions sets LogFilePermissions field to given value.


### GetAppend

`func (o *AddPeriodicStatsLoggerPluginRequest) GetAppend() bool`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetAppendOk() (*bool, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *AddPeriodicStatsLoggerPluginRequest) SetAppend(v bool)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *AddPeriodicStatsLoggerPluginRequest) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetRotationPolicy

`func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationPolicy() []string`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationPolicyOk() (*[]string, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *AddPeriodicStatsLoggerPluginRequest) SetRotationPolicy(v []string)`

SetRotationPolicy sets RotationPolicy field to given value.


### GetRotationListener

`func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationListener() []string`

GetRotationListener returns the RotationListener field if non-nil, zero value otherwise.

### GetRotationListenerOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationListenerOk() (*[]string, bool)`

GetRotationListenerOk returns a tuple with the RotationListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationListener

`func (o *AddPeriodicStatsLoggerPluginRequest) SetRotationListener(v []string)`

SetRotationListener sets RotationListener field to given value.

### HasRotationListener

`func (o *AddPeriodicStatsLoggerPluginRequest) HasRotationListener() bool`

HasRotationListener returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *AddPeriodicStatsLoggerPluginRequest) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *AddPeriodicStatsLoggerPluginRequest) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetLoggingErrorBehavior

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLoggingErrorBehavior() EnumpluginLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLoggingErrorBehaviorOk() (*EnumpluginLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLoggingErrorBehavior(v EnumpluginLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *AddPeriodicStatsLoggerPluginRequest) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetLocalDBBackendInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp`

GetLocalDBBackendInfo returns the LocalDBBackendInfo field if non-nil, zero value otherwise.

### GetLocalDBBackendInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool)`

GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDBBackendInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp)`

SetLocalDBBackendInfo sets LocalDBBackendInfo field to given value.

### HasLocalDBBackendInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasLocalDBBackendInfo() bool`

HasLocalDBBackendInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetReplicationInfo() EnumpluginReplicationInfoProp`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetReplicationInfo(v EnumpluginReplicationInfoProp)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetEntryCacheInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp`

GetEntryCacheInfo returns the EntryCacheInfo field if non-nil, zero value otherwise.

### GetEntryCacheInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool)`

GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCacheInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp)`

SetEntryCacheInfo sets EntryCacheInfo field to given value.

### HasEntryCacheInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasEntryCacheInfo() bool`

HasEntryCacheInfo returns a boolean if a field has been set.

### GetHostInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHostInfo() []EnumpluginHostInfoProp`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetHostInfoOk() (*[]EnumpluginHostInfoProp, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) SetHostInfo(v []EnumpluginHostInfoProp)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *AddPeriodicStatsLoggerPluginRequest) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetIncludedLDAPApplication

`func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPApplication() []string`

GetIncludedLDAPApplication returns the IncludedLDAPApplication field if non-nil, zero value otherwise.

### GetIncludedLDAPApplicationOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPApplicationOk() (*[]string, bool)`

GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPApplication

`func (o *AddPeriodicStatsLoggerPluginRequest) SetIncludedLDAPApplication(v []string)`

SetIncludedLDAPApplication sets IncludedLDAPApplication field to given value.

### HasIncludedLDAPApplication

`func (o *AddPeriodicStatsLoggerPluginRequest) HasIncludedLDAPApplication() bool`

HasIncludedLDAPApplication returns a boolean if a field has been set.

### GetDescription

`func (o *AddPeriodicStatsLoggerPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPeriodicStatsLoggerPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPeriodicStatsLoggerPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPeriodicStatsLoggerPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPeriodicStatsLoggerPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPeriodicStatsLoggerPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

