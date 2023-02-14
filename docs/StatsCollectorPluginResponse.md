# StatsCollectorPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumstatsCollectorPluginSchemaUrn**](EnumstatsCollectorPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**SampleInterval** | **string** | The duration between statistics collections. Setting this value too small can have an impact on performance. This value should be a multiple of collection-interval. | 
**CollectionInterval** | **string** | Some of the calculated statistics, such as the average and maximum queue sizes, can use multiple samples within a log interval. This value controls how often samples are gathered, and setting this value too small can have an adverse impact on performance. | 
**LdapInfo** | Pointer to [**EnumpluginLdapInfoProp**](EnumpluginLdapInfoProp.md) |  | [optional] 
**ServerInfo** | Pointer to [**EnumpluginServerInfoProp**](EnumpluginServerInfoProp.md) |  | [optional] 
**PerApplicationLDAPStats** | Pointer to [**EnumpluginPerApplicationLDAPStatsProp**](EnumpluginPerApplicationLDAPStatsProp.md) |  | [optional] 
**LdapChangelogInfo** | Pointer to [**EnumpluginLdapChangelogInfoProp**](EnumpluginLdapChangelogInfoProp.md) |  | [optional] 
**StatusSummaryInfo** | Pointer to [**EnumpluginStatusSummaryInfoProp**](EnumpluginStatusSummaryInfoProp.md) |  | [optional] 
**GenerateCollectorFiles** | Pointer to **bool** | Indicates whether this plugin should store metric samples on disk for use by the Data Metrics Server. If the Stats Collector Plugin is only being used to collect metrics for one or more StatsD Monitoring Endpoints, then this can be set to false to prevent unnecessary I/O. | [optional] 
**LocalDBBackendInfo** | Pointer to [**EnumpluginLocalDBBackendInfoProp**](EnumpluginLocalDBBackendInfoProp.md) |  | [optional] 
**ReplicationInfo** | Pointer to [**EnumpluginReplicationInfoProp**](EnumpluginReplicationInfoProp.md) |  | [optional] 
**EntryCacheInfo** | Pointer to [**EnumpluginEntryCacheInfoProp**](EnumpluginEntryCacheInfoProp.md) |  | [optional] 
**HostInfo** | Pointer to [**[]EnumpluginHostInfoProp**](EnumpluginHostInfoProp.md) |  | [optional] 
**IncludedLDAPApplication** | Pointer to **[]string** | If statistics should not be included for all applications, this property names the subset of applications that should be included. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewStatsCollectorPluginResponse

`func NewStatsCollectorPluginResponse(schemas []EnumstatsCollectorPluginSchemaUrn, id string, sampleInterval string, collectionInterval string, enabled bool, ) *StatsCollectorPluginResponse`

NewStatsCollectorPluginResponse instantiates a new StatsCollectorPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsCollectorPluginResponseWithDefaults

`func NewStatsCollectorPluginResponseWithDefaults() *StatsCollectorPluginResponse`

NewStatsCollectorPluginResponseWithDefaults instantiates a new StatsCollectorPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *StatsCollectorPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StatsCollectorPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StatsCollectorPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StatsCollectorPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *StatsCollectorPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *StatsCollectorPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *StatsCollectorPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *StatsCollectorPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *StatsCollectorPluginResponse) GetSchemas() []EnumstatsCollectorPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StatsCollectorPluginResponse) GetSchemasOk() (*[]EnumstatsCollectorPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StatsCollectorPluginResponse) SetSchemas(v []EnumstatsCollectorPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *StatsCollectorPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsCollectorPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsCollectorPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSampleInterval

`func (o *StatsCollectorPluginResponse) GetSampleInterval() string`

GetSampleInterval returns the SampleInterval field if non-nil, zero value otherwise.

### GetSampleIntervalOk

`func (o *StatsCollectorPluginResponse) GetSampleIntervalOk() (*string, bool)`

GetSampleIntervalOk returns a tuple with the SampleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleInterval

`func (o *StatsCollectorPluginResponse) SetSampleInterval(v string)`

SetSampleInterval sets SampleInterval field to given value.


### GetCollectionInterval

`func (o *StatsCollectorPluginResponse) GetCollectionInterval() string`

GetCollectionInterval returns the CollectionInterval field if non-nil, zero value otherwise.

### GetCollectionIntervalOk

`func (o *StatsCollectorPluginResponse) GetCollectionIntervalOk() (*string, bool)`

GetCollectionIntervalOk returns a tuple with the CollectionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInterval

`func (o *StatsCollectorPluginResponse) SetCollectionInterval(v string)`

SetCollectionInterval sets CollectionInterval field to given value.


### GetLdapInfo

`func (o *StatsCollectorPluginResponse) GetLdapInfo() EnumpluginLdapInfoProp`

GetLdapInfo returns the LdapInfo field if non-nil, zero value otherwise.

### GetLdapInfoOk

`func (o *StatsCollectorPluginResponse) GetLdapInfoOk() (*EnumpluginLdapInfoProp, bool)`

GetLdapInfoOk returns a tuple with the LdapInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapInfo

`func (o *StatsCollectorPluginResponse) SetLdapInfo(v EnumpluginLdapInfoProp)`

SetLdapInfo sets LdapInfo field to given value.

### HasLdapInfo

`func (o *StatsCollectorPluginResponse) HasLdapInfo() bool`

HasLdapInfo returns a boolean if a field has been set.

### GetServerInfo

`func (o *StatsCollectorPluginResponse) GetServerInfo() EnumpluginServerInfoProp`

GetServerInfo returns the ServerInfo field if non-nil, zero value otherwise.

### GetServerInfoOk

`func (o *StatsCollectorPluginResponse) GetServerInfoOk() (*EnumpluginServerInfoProp, bool)`

GetServerInfoOk returns a tuple with the ServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInfo

`func (o *StatsCollectorPluginResponse) SetServerInfo(v EnumpluginServerInfoProp)`

SetServerInfo sets ServerInfo field to given value.

### HasServerInfo

`func (o *StatsCollectorPluginResponse) HasServerInfo() bool`

HasServerInfo returns a boolean if a field has been set.

### GetPerApplicationLDAPStats

`func (o *StatsCollectorPluginResponse) GetPerApplicationLDAPStats() EnumpluginPerApplicationLDAPStatsProp`

GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field if non-nil, zero value otherwise.

### GetPerApplicationLDAPStatsOk

`func (o *StatsCollectorPluginResponse) GetPerApplicationLDAPStatsOk() (*EnumpluginPerApplicationLDAPStatsProp, bool)`

GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerApplicationLDAPStats

`func (o *StatsCollectorPluginResponse) SetPerApplicationLDAPStats(v EnumpluginPerApplicationLDAPStatsProp)`

SetPerApplicationLDAPStats sets PerApplicationLDAPStats field to given value.

### HasPerApplicationLDAPStats

`func (o *StatsCollectorPluginResponse) HasPerApplicationLDAPStats() bool`

HasPerApplicationLDAPStats returns a boolean if a field has been set.

### GetLdapChangelogInfo

`func (o *StatsCollectorPluginResponse) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp`

GetLdapChangelogInfo returns the LdapChangelogInfo field if non-nil, zero value otherwise.

### GetLdapChangelogInfoOk

`func (o *StatsCollectorPluginResponse) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool)`

GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapChangelogInfo

`func (o *StatsCollectorPluginResponse) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp)`

SetLdapChangelogInfo sets LdapChangelogInfo field to given value.

### HasLdapChangelogInfo

`func (o *StatsCollectorPluginResponse) HasLdapChangelogInfo() bool`

HasLdapChangelogInfo returns a boolean if a field has been set.

### GetStatusSummaryInfo

`func (o *StatsCollectorPluginResponse) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp`

GetStatusSummaryInfo returns the StatusSummaryInfo field if non-nil, zero value otherwise.

### GetStatusSummaryInfoOk

`func (o *StatsCollectorPluginResponse) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool)`

GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSummaryInfo

`func (o *StatsCollectorPluginResponse) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp)`

SetStatusSummaryInfo sets StatusSummaryInfo field to given value.

### HasStatusSummaryInfo

`func (o *StatsCollectorPluginResponse) HasStatusSummaryInfo() bool`

HasStatusSummaryInfo returns a boolean if a field has been set.

### GetGenerateCollectorFiles

`func (o *StatsCollectorPluginResponse) GetGenerateCollectorFiles() bool`

GetGenerateCollectorFiles returns the GenerateCollectorFiles field if non-nil, zero value otherwise.

### GetGenerateCollectorFilesOk

`func (o *StatsCollectorPluginResponse) GetGenerateCollectorFilesOk() (*bool, bool)`

GetGenerateCollectorFilesOk returns a tuple with the GenerateCollectorFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCollectorFiles

`func (o *StatsCollectorPluginResponse) SetGenerateCollectorFiles(v bool)`

SetGenerateCollectorFiles sets GenerateCollectorFiles field to given value.

### HasGenerateCollectorFiles

`func (o *StatsCollectorPluginResponse) HasGenerateCollectorFiles() bool`

HasGenerateCollectorFiles returns a boolean if a field has been set.

### GetLocalDBBackendInfo

`func (o *StatsCollectorPluginResponse) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp`

GetLocalDBBackendInfo returns the LocalDBBackendInfo field if non-nil, zero value otherwise.

### GetLocalDBBackendInfoOk

`func (o *StatsCollectorPluginResponse) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool)`

GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDBBackendInfo

`func (o *StatsCollectorPluginResponse) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp)`

SetLocalDBBackendInfo sets LocalDBBackendInfo field to given value.

### HasLocalDBBackendInfo

`func (o *StatsCollectorPluginResponse) HasLocalDBBackendInfo() bool`

HasLocalDBBackendInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *StatsCollectorPluginResponse) GetReplicationInfo() EnumpluginReplicationInfoProp`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *StatsCollectorPluginResponse) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *StatsCollectorPluginResponse) SetReplicationInfo(v EnumpluginReplicationInfoProp)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *StatsCollectorPluginResponse) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetEntryCacheInfo

`func (o *StatsCollectorPluginResponse) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp`

GetEntryCacheInfo returns the EntryCacheInfo field if non-nil, zero value otherwise.

### GetEntryCacheInfoOk

`func (o *StatsCollectorPluginResponse) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool)`

GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCacheInfo

`func (o *StatsCollectorPluginResponse) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp)`

SetEntryCacheInfo sets EntryCacheInfo field to given value.

### HasEntryCacheInfo

`func (o *StatsCollectorPluginResponse) HasEntryCacheInfo() bool`

HasEntryCacheInfo returns a boolean if a field has been set.

### GetHostInfo

`func (o *StatsCollectorPluginResponse) GetHostInfo() []EnumpluginHostInfoProp`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *StatsCollectorPluginResponse) GetHostInfoOk() (*[]EnumpluginHostInfoProp, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *StatsCollectorPluginResponse) SetHostInfo(v []EnumpluginHostInfoProp)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *StatsCollectorPluginResponse) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetIncludedLDAPApplication

`func (o *StatsCollectorPluginResponse) GetIncludedLDAPApplication() []string`

GetIncludedLDAPApplication returns the IncludedLDAPApplication field if non-nil, zero value otherwise.

### GetIncludedLDAPApplicationOk

`func (o *StatsCollectorPluginResponse) GetIncludedLDAPApplicationOk() (*[]string, bool)`

GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPApplication

`func (o *StatsCollectorPluginResponse) SetIncludedLDAPApplication(v []string)`

SetIncludedLDAPApplication sets IncludedLDAPApplication field to given value.

### HasIncludedLDAPApplication

`func (o *StatsCollectorPluginResponse) HasIncludedLDAPApplication() bool`

HasIncludedLDAPApplication returns a boolean if a field has been set.

### GetDescription

`func (o *StatsCollectorPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatsCollectorPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatsCollectorPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatsCollectorPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *StatsCollectorPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StatsCollectorPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StatsCollectorPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


