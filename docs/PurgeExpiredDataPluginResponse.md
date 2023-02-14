# PurgeExpiredDataPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 
**Schemas** | [**[]EnumpurgeExpiredDataPluginSchemaUrn**](EnumpurgeExpiredDataPluginSchemaUrn.md) |  | 
**DatetimeAttribute** | **string** | The LDAP attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted. | 
**DatetimeJSONField** | Pointer to **string** | The top-level JSON field within the configured datetime-attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted. | [optional] 
**DatetimeFormat** | [**EnumpluginDatetimeFormatProp**](EnumpluginDatetimeFormatProp.md) |  | 
**CustomDatetimeFormat** | Pointer to **string** | When the datetime-format property is configured with a value of \&quot;custom\&quot;, this specifies the format (using a string compatible with the java.text.SimpleDateFormat class) that will be used to search for expired data. | [optional] 
**CustomTimezone** | Pointer to **string** | Specifies the time zone to use when generating a date string using the configured custom-datetime-format value. The provided value must be accepted by java.util.TimeZone.getTimeZone. | [optional] 
**ExpirationOffset** | **string** | The duration to wait after the value specified in datetime-attribute (and optionally datetime-json-field) before purging the data. | 
**PurgeBehavior** | Pointer to [**EnumpluginPurgeBehaviorProp**](EnumpluginPurgeBehaviorProp.md) |  | [optional] 
**BaseDN** | Pointer to **string** | Only entries located within the subtree specified by this base DN are eligible for purging. | [optional] 
**Filter** | Pointer to **string** | Only entries that match this LDAP filter will be eligible for having data purged. | [optional] 
**PollingInterval** | **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | 
**MaxUpdatesPerSecond** | **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**NumDeleteThreads** | **int32** | The number of threads used to delete expired entries. | 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewPurgeExpiredDataPluginResponse

`func NewPurgeExpiredDataPluginResponse(id string, schemas []EnumpurgeExpiredDataPluginSchemaUrn, datetimeAttribute string, datetimeFormat EnumpluginDatetimeFormatProp, expirationOffset string, pollingInterval string, maxUpdatesPerSecond int32, numDeleteThreads int32, enabled bool, ) *PurgeExpiredDataPluginResponse`

NewPurgeExpiredDataPluginResponse instantiates a new PurgeExpiredDataPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeExpiredDataPluginResponseWithDefaults

`func NewPurgeExpiredDataPluginResponseWithDefaults() *PurgeExpiredDataPluginResponse`

NewPurgeExpiredDataPluginResponseWithDefaults instantiates a new PurgeExpiredDataPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PurgeExpiredDataPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PurgeExpiredDataPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PurgeExpiredDataPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PurgeExpiredDataPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PurgeExpiredDataPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PurgeExpiredDataPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PurgeExpiredDataPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PurgeExpiredDataPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *PurgeExpiredDataPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurgeExpiredDataPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurgeExpiredDataPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PurgeExpiredDataPluginResponse) GetSchemas() []EnumpurgeExpiredDataPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PurgeExpiredDataPluginResponse) GetSchemasOk() (*[]EnumpurgeExpiredDataPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PurgeExpiredDataPluginResponse) SetSchemas(v []EnumpurgeExpiredDataPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDatetimeAttribute

`func (o *PurgeExpiredDataPluginResponse) GetDatetimeAttribute() string`

GetDatetimeAttribute returns the DatetimeAttribute field if non-nil, zero value otherwise.

### GetDatetimeAttributeOk

`func (o *PurgeExpiredDataPluginResponse) GetDatetimeAttributeOk() (*string, bool)`

GetDatetimeAttributeOk returns a tuple with the DatetimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeAttribute

`func (o *PurgeExpiredDataPluginResponse) SetDatetimeAttribute(v string)`

SetDatetimeAttribute sets DatetimeAttribute field to given value.


### GetDatetimeJSONField

`func (o *PurgeExpiredDataPluginResponse) GetDatetimeJSONField() string`

GetDatetimeJSONField returns the DatetimeJSONField field if non-nil, zero value otherwise.

### GetDatetimeJSONFieldOk

`func (o *PurgeExpiredDataPluginResponse) GetDatetimeJSONFieldOk() (*string, bool)`

GetDatetimeJSONFieldOk returns a tuple with the DatetimeJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeJSONField

`func (o *PurgeExpiredDataPluginResponse) SetDatetimeJSONField(v string)`

SetDatetimeJSONField sets DatetimeJSONField field to given value.

### HasDatetimeJSONField

`func (o *PurgeExpiredDataPluginResponse) HasDatetimeJSONField() bool`

HasDatetimeJSONField returns a boolean if a field has been set.

### GetDatetimeFormat

`func (o *PurgeExpiredDataPluginResponse) GetDatetimeFormat() EnumpluginDatetimeFormatProp`

GetDatetimeFormat returns the DatetimeFormat field if non-nil, zero value otherwise.

### GetDatetimeFormatOk

`func (o *PurgeExpiredDataPluginResponse) GetDatetimeFormatOk() (*EnumpluginDatetimeFormatProp, bool)`

GetDatetimeFormatOk returns a tuple with the DatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeFormat

`func (o *PurgeExpiredDataPluginResponse) SetDatetimeFormat(v EnumpluginDatetimeFormatProp)`

SetDatetimeFormat sets DatetimeFormat field to given value.


### GetCustomDatetimeFormat

`func (o *PurgeExpiredDataPluginResponse) GetCustomDatetimeFormat() string`

GetCustomDatetimeFormat returns the CustomDatetimeFormat field if non-nil, zero value otherwise.

### GetCustomDatetimeFormatOk

`func (o *PurgeExpiredDataPluginResponse) GetCustomDatetimeFormatOk() (*string, bool)`

GetCustomDatetimeFormatOk returns a tuple with the CustomDatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDatetimeFormat

`func (o *PurgeExpiredDataPluginResponse) SetCustomDatetimeFormat(v string)`

SetCustomDatetimeFormat sets CustomDatetimeFormat field to given value.

### HasCustomDatetimeFormat

`func (o *PurgeExpiredDataPluginResponse) HasCustomDatetimeFormat() bool`

HasCustomDatetimeFormat returns a boolean if a field has been set.

### GetCustomTimezone

`func (o *PurgeExpiredDataPluginResponse) GetCustomTimezone() string`

GetCustomTimezone returns the CustomTimezone field if non-nil, zero value otherwise.

### GetCustomTimezoneOk

`func (o *PurgeExpiredDataPluginResponse) GetCustomTimezoneOk() (*string, bool)`

GetCustomTimezoneOk returns a tuple with the CustomTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimezone

`func (o *PurgeExpiredDataPluginResponse) SetCustomTimezone(v string)`

SetCustomTimezone sets CustomTimezone field to given value.

### HasCustomTimezone

`func (o *PurgeExpiredDataPluginResponse) HasCustomTimezone() bool`

HasCustomTimezone returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *PurgeExpiredDataPluginResponse) GetExpirationOffset() string`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *PurgeExpiredDataPluginResponse) GetExpirationOffsetOk() (*string, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *PurgeExpiredDataPluginResponse) SetExpirationOffset(v string)`

SetExpirationOffset sets ExpirationOffset field to given value.


### GetPurgeBehavior

`func (o *PurgeExpiredDataPluginResponse) GetPurgeBehavior() EnumpluginPurgeBehaviorProp`

GetPurgeBehavior returns the PurgeBehavior field if non-nil, zero value otherwise.

### GetPurgeBehaviorOk

`func (o *PurgeExpiredDataPluginResponse) GetPurgeBehaviorOk() (*EnumpluginPurgeBehaviorProp, bool)`

GetPurgeBehaviorOk returns a tuple with the PurgeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeBehavior

`func (o *PurgeExpiredDataPluginResponse) SetPurgeBehavior(v EnumpluginPurgeBehaviorProp)`

SetPurgeBehavior sets PurgeBehavior field to given value.

### HasPurgeBehavior

`func (o *PurgeExpiredDataPluginResponse) HasPurgeBehavior() bool`

HasPurgeBehavior returns a boolean if a field has been set.

### GetBaseDN

`func (o *PurgeExpiredDataPluginResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *PurgeExpiredDataPluginResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *PurgeExpiredDataPluginResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *PurgeExpiredDataPluginResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetFilter

`func (o *PurgeExpiredDataPluginResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PurgeExpiredDataPluginResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PurgeExpiredDataPluginResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PurgeExpiredDataPluginResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPollingInterval

`func (o *PurgeExpiredDataPluginResponse) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *PurgeExpiredDataPluginResponse) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *PurgeExpiredDataPluginResponse) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetMaxUpdatesPerSecond

`func (o *PurgeExpiredDataPluginResponse) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *PurgeExpiredDataPluginResponse) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *PurgeExpiredDataPluginResponse) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetPeerServerPriorityIndex

`func (o *PurgeExpiredDataPluginResponse) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *PurgeExpiredDataPluginResponse) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *PurgeExpiredDataPluginResponse) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *PurgeExpiredDataPluginResponse) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetNumDeleteThreads

`func (o *PurgeExpiredDataPluginResponse) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *PurgeExpiredDataPluginResponse) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *PurgeExpiredDataPluginResponse) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetDescription

`func (o *PurgeExpiredDataPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurgeExpiredDataPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurgeExpiredDataPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PurgeExpiredDataPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PurgeExpiredDataPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PurgeExpiredDataPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PurgeExpiredDataPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


