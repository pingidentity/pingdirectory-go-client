# AddPurgeExpiredDataPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
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

### NewAddPurgeExpiredDataPluginRequest

`func NewAddPurgeExpiredDataPluginRequest(pluginName string, schemas []EnumpurgeExpiredDataPluginSchemaUrn, datetimeAttribute string, datetimeFormat EnumpluginDatetimeFormatProp, expirationOffset string, pollingInterval string, maxUpdatesPerSecond int32, numDeleteThreads int32, enabled bool, ) *AddPurgeExpiredDataPluginRequest`

NewAddPurgeExpiredDataPluginRequest instantiates a new AddPurgeExpiredDataPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPurgeExpiredDataPluginRequestWithDefaults

`func NewAddPurgeExpiredDataPluginRequestWithDefaults() *AddPurgeExpiredDataPluginRequest`

NewAddPurgeExpiredDataPluginRequestWithDefaults instantiates a new AddPurgeExpiredDataPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddPurgeExpiredDataPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddPurgeExpiredDataPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddPurgeExpiredDataPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddPurgeExpiredDataPluginRequest) GetSchemas() []EnumpurgeExpiredDataPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPurgeExpiredDataPluginRequest) GetSchemasOk() (*[]EnumpurgeExpiredDataPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPurgeExpiredDataPluginRequest) SetSchemas(v []EnumpurgeExpiredDataPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDatetimeAttribute

`func (o *AddPurgeExpiredDataPluginRequest) GetDatetimeAttribute() string`

GetDatetimeAttribute returns the DatetimeAttribute field if non-nil, zero value otherwise.

### GetDatetimeAttributeOk

`func (o *AddPurgeExpiredDataPluginRequest) GetDatetimeAttributeOk() (*string, bool)`

GetDatetimeAttributeOk returns a tuple with the DatetimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeAttribute

`func (o *AddPurgeExpiredDataPluginRequest) SetDatetimeAttribute(v string)`

SetDatetimeAttribute sets DatetimeAttribute field to given value.


### GetDatetimeJSONField

`func (o *AddPurgeExpiredDataPluginRequest) GetDatetimeJSONField() string`

GetDatetimeJSONField returns the DatetimeJSONField field if non-nil, zero value otherwise.

### GetDatetimeJSONFieldOk

`func (o *AddPurgeExpiredDataPluginRequest) GetDatetimeJSONFieldOk() (*string, bool)`

GetDatetimeJSONFieldOk returns a tuple with the DatetimeJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeJSONField

`func (o *AddPurgeExpiredDataPluginRequest) SetDatetimeJSONField(v string)`

SetDatetimeJSONField sets DatetimeJSONField field to given value.

### HasDatetimeJSONField

`func (o *AddPurgeExpiredDataPluginRequest) HasDatetimeJSONField() bool`

HasDatetimeJSONField returns a boolean if a field has been set.

### GetDatetimeFormat

`func (o *AddPurgeExpiredDataPluginRequest) GetDatetimeFormat() EnumpluginDatetimeFormatProp`

GetDatetimeFormat returns the DatetimeFormat field if non-nil, zero value otherwise.

### GetDatetimeFormatOk

`func (o *AddPurgeExpiredDataPluginRequest) GetDatetimeFormatOk() (*EnumpluginDatetimeFormatProp, bool)`

GetDatetimeFormatOk returns a tuple with the DatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeFormat

`func (o *AddPurgeExpiredDataPluginRequest) SetDatetimeFormat(v EnumpluginDatetimeFormatProp)`

SetDatetimeFormat sets DatetimeFormat field to given value.


### GetCustomDatetimeFormat

`func (o *AddPurgeExpiredDataPluginRequest) GetCustomDatetimeFormat() string`

GetCustomDatetimeFormat returns the CustomDatetimeFormat field if non-nil, zero value otherwise.

### GetCustomDatetimeFormatOk

`func (o *AddPurgeExpiredDataPluginRequest) GetCustomDatetimeFormatOk() (*string, bool)`

GetCustomDatetimeFormatOk returns a tuple with the CustomDatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDatetimeFormat

`func (o *AddPurgeExpiredDataPluginRequest) SetCustomDatetimeFormat(v string)`

SetCustomDatetimeFormat sets CustomDatetimeFormat field to given value.

### HasCustomDatetimeFormat

`func (o *AddPurgeExpiredDataPluginRequest) HasCustomDatetimeFormat() bool`

HasCustomDatetimeFormat returns a boolean if a field has been set.

### GetCustomTimezone

`func (o *AddPurgeExpiredDataPluginRequest) GetCustomTimezone() string`

GetCustomTimezone returns the CustomTimezone field if non-nil, zero value otherwise.

### GetCustomTimezoneOk

`func (o *AddPurgeExpiredDataPluginRequest) GetCustomTimezoneOk() (*string, bool)`

GetCustomTimezoneOk returns a tuple with the CustomTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimezone

`func (o *AddPurgeExpiredDataPluginRequest) SetCustomTimezone(v string)`

SetCustomTimezone sets CustomTimezone field to given value.

### HasCustomTimezone

`func (o *AddPurgeExpiredDataPluginRequest) HasCustomTimezone() bool`

HasCustomTimezone returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *AddPurgeExpiredDataPluginRequest) GetExpirationOffset() string`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *AddPurgeExpiredDataPluginRequest) GetExpirationOffsetOk() (*string, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *AddPurgeExpiredDataPluginRequest) SetExpirationOffset(v string)`

SetExpirationOffset sets ExpirationOffset field to given value.


### GetPurgeBehavior

`func (o *AddPurgeExpiredDataPluginRequest) GetPurgeBehavior() EnumpluginPurgeBehaviorProp`

GetPurgeBehavior returns the PurgeBehavior field if non-nil, zero value otherwise.

### GetPurgeBehaviorOk

`func (o *AddPurgeExpiredDataPluginRequest) GetPurgeBehaviorOk() (*EnumpluginPurgeBehaviorProp, bool)`

GetPurgeBehaviorOk returns a tuple with the PurgeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeBehavior

`func (o *AddPurgeExpiredDataPluginRequest) SetPurgeBehavior(v EnumpluginPurgeBehaviorProp)`

SetPurgeBehavior sets PurgeBehavior field to given value.

### HasPurgeBehavior

`func (o *AddPurgeExpiredDataPluginRequest) HasPurgeBehavior() bool`

HasPurgeBehavior returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddPurgeExpiredDataPluginRequest) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddPurgeExpiredDataPluginRequest) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddPurgeExpiredDataPluginRequest) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddPurgeExpiredDataPluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddPurgeExpiredDataPluginRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddPurgeExpiredDataPluginRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddPurgeExpiredDataPluginRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddPurgeExpiredDataPluginRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPollingInterval

`func (o *AddPurgeExpiredDataPluginRequest) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *AddPurgeExpiredDataPluginRequest) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *AddPurgeExpiredDataPluginRequest) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetMaxUpdatesPerSecond

`func (o *AddPurgeExpiredDataPluginRequest) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *AddPurgeExpiredDataPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *AddPurgeExpiredDataPluginRequest) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetPeerServerPriorityIndex

`func (o *AddPurgeExpiredDataPluginRequest) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *AddPurgeExpiredDataPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *AddPurgeExpiredDataPluginRequest) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *AddPurgeExpiredDataPluginRequest) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetNumDeleteThreads

`func (o *AddPurgeExpiredDataPluginRequest) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *AddPurgeExpiredDataPluginRequest) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *AddPurgeExpiredDataPluginRequest) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetDescription

`func (o *AddPurgeExpiredDataPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPurgeExpiredDataPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPurgeExpiredDataPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPurgeExpiredDataPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPurgeExpiredDataPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPurgeExpiredDataPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPurgeExpiredDataPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


