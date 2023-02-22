# RecurringTaskChainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Recurring Task Chain | 
**Schemas** | Pointer to [**[]EnumrecurringTaskChainSchemaUrn**](EnumrecurringTaskChainSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task Chain | [optional] 
**Enabled** | **bool** | Indicates whether this Recurring Task Chain is enabled for use. Recurring Task Chains that are disabled will not have any new instances scheduled, but instances that are already scheduled will be preserved. Those instances may be manually canceled if desired. | 
**RecurringTask** | **[]string** | The set of recurring tasks that make up this chain. At least one value must be provided. If multiple values are given, then the task instances will be invoked in the order in which they are listed. | 
**ScheduledMonth** | [**[]EnumrecurringTaskChainScheduledMonthProp**](EnumrecurringTaskChainScheduledMonthProp.md) |  | 
**ScheduledDateSelectionType** | [**EnumrecurringTaskChainScheduledDateSelectionTypeProp**](EnumrecurringTaskChainScheduledDateSelectionTypeProp.md) |  | 
**ScheduledDayOfTheWeek** | Pointer to [**[]EnumrecurringTaskChainScheduledDayOfTheWeekProp**](EnumrecurringTaskChainScheduledDayOfTheWeekProp.md) |  | [optional] 
**ScheduledDayOfTheMonth** | Pointer to [**[]EnumrecurringTaskChainScheduledDayOfTheMonthProp**](EnumrecurringTaskChainScheduledDayOfTheMonthProp.md) |  | [optional] 
**ScheduledTimeOfDay** | **[]string** | The time of day at which instances of the Recurring Task Chain should be eligible to start running. Values should be in the format HH:MM (where HH is a two-digit representation of the hour of the day, between 00 and 23, inclusive), and MM is a two-digit representation of the minute of the hour (between 00 and 59, inclusive). Alternately, the value can be in the form *:MM, which indicates that the task should be eligible to start at the specified minute of every hour. At least one value must be provided, but multiple values may be given to indicate multiple start times within the same day. | 
**TimeZone** | Pointer to **string** | The time zone that will be used to interpret the scheduled-time-of-day values. If no value is provided, then the JVM&#39;s default time zone will be used. | [optional] 
**InterruptedByShutdownBehavior** | Pointer to [**EnumrecurringTaskChainInterruptedByShutdownBehaviorProp**](EnumrecurringTaskChainInterruptedByShutdownBehaviorProp.md) |  | [optional] 
**ServerOfflineAtStartTimeBehavior** | Pointer to [**EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp**](EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp.md) |  | [optional] 

## Methods

### NewRecurringTaskChainResponse

`func NewRecurringTaskChainResponse(id string, enabled bool, recurringTask []string, scheduledMonth []EnumrecurringTaskChainScheduledMonthProp, scheduledDateSelectionType EnumrecurringTaskChainScheduledDateSelectionTypeProp, scheduledTimeOfDay []string, ) *RecurringTaskChainResponse`

NewRecurringTaskChainResponse instantiates a new RecurringTaskChainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringTaskChainResponseWithDefaults

`func NewRecurringTaskChainResponseWithDefaults() *RecurringTaskChainResponse`

NewRecurringTaskChainResponseWithDefaults instantiates a new RecurringTaskChainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *RecurringTaskChainResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RecurringTaskChainResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RecurringTaskChainResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RecurringTaskChainResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *RecurringTaskChainResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *RecurringTaskChainResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *RecurringTaskChainResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *RecurringTaskChainResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *RecurringTaskChainResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurringTaskChainResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurringTaskChainResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *RecurringTaskChainResponse) GetSchemas() []EnumrecurringTaskChainSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RecurringTaskChainResponse) GetSchemasOk() (*[]EnumrecurringTaskChainSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RecurringTaskChainResponse) SetSchemas(v []EnumrecurringTaskChainSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *RecurringTaskChainResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *RecurringTaskChainResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurringTaskChainResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurringTaskChainResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurringTaskChainResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RecurringTaskChainResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecurringTaskChainResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecurringTaskChainResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRecurringTask

`func (o *RecurringTaskChainResponse) GetRecurringTask() []string`

GetRecurringTask returns the RecurringTask field if non-nil, zero value otherwise.

### GetRecurringTaskOk

`func (o *RecurringTaskChainResponse) GetRecurringTaskOk() (*[]string, bool)`

GetRecurringTaskOk returns a tuple with the RecurringTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTask

`func (o *RecurringTaskChainResponse) SetRecurringTask(v []string)`

SetRecurringTask sets RecurringTask field to given value.


### GetScheduledMonth

`func (o *RecurringTaskChainResponse) GetScheduledMonth() []EnumrecurringTaskChainScheduledMonthProp`

GetScheduledMonth returns the ScheduledMonth field if non-nil, zero value otherwise.

### GetScheduledMonthOk

`func (o *RecurringTaskChainResponse) GetScheduledMonthOk() (*[]EnumrecurringTaskChainScheduledMonthProp, bool)`

GetScheduledMonthOk returns a tuple with the ScheduledMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMonth

`func (o *RecurringTaskChainResponse) SetScheduledMonth(v []EnumrecurringTaskChainScheduledMonthProp)`

SetScheduledMonth sets ScheduledMonth field to given value.


### GetScheduledDateSelectionType

`func (o *RecurringTaskChainResponse) GetScheduledDateSelectionType() EnumrecurringTaskChainScheduledDateSelectionTypeProp`

GetScheduledDateSelectionType returns the ScheduledDateSelectionType field if non-nil, zero value otherwise.

### GetScheduledDateSelectionTypeOk

`func (o *RecurringTaskChainResponse) GetScheduledDateSelectionTypeOk() (*EnumrecurringTaskChainScheduledDateSelectionTypeProp, bool)`

GetScheduledDateSelectionTypeOk returns a tuple with the ScheduledDateSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDateSelectionType

`func (o *RecurringTaskChainResponse) SetScheduledDateSelectionType(v EnumrecurringTaskChainScheduledDateSelectionTypeProp)`

SetScheduledDateSelectionType sets ScheduledDateSelectionType field to given value.


### GetScheduledDayOfTheWeek

`func (o *RecurringTaskChainResponse) GetScheduledDayOfTheWeek() []EnumrecurringTaskChainScheduledDayOfTheWeekProp`

GetScheduledDayOfTheWeek returns the ScheduledDayOfTheWeek field if non-nil, zero value otherwise.

### GetScheduledDayOfTheWeekOk

`func (o *RecurringTaskChainResponse) GetScheduledDayOfTheWeekOk() (*[]EnumrecurringTaskChainScheduledDayOfTheWeekProp, bool)`

GetScheduledDayOfTheWeekOk returns a tuple with the ScheduledDayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDayOfTheWeek

`func (o *RecurringTaskChainResponse) SetScheduledDayOfTheWeek(v []EnumrecurringTaskChainScheduledDayOfTheWeekProp)`

SetScheduledDayOfTheWeek sets ScheduledDayOfTheWeek field to given value.

### HasScheduledDayOfTheWeek

`func (o *RecurringTaskChainResponse) HasScheduledDayOfTheWeek() bool`

HasScheduledDayOfTheWeek returns a boolean if a field has been set.

### GetScheduledDayOfTheMonth

`func (o *RecurringTaskChainResponse) GetScheduledDayOfTheMonth() []EnumrecurringTaskChainScheduledDayOfTheMonthProp`

GetScheduledDayOfTheMonth returns the ScheduledDayOfTheMonth field if non-nil, zero value otherwise.

### GetScheduledDayOfTheMonthOk

`func (o *RecurringTaskChainResponse) GetScheduledDayOfTheMonthOk() (*[]EnumrecurringTaskChainScheduledDayOfTheMonthProp, bool)`

GetScheduledDayOfTheMonthOk returns a tuple with the ScheduledDayOfTheMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDayOfTheMonth

`func (o *RecurringTaskChainResponse) SetScheduledDayOfTheMonth(v []EnumrecurringTaskChainScheduledDayOfTheMonthProp)`

SetScheduledDayOfTheMonth sets ScheduledDayOfTheMonth field to given value.

### HasScheduledDayOfTheMonth

`func (o *RecurringTaskChainResponse) HasScheduledDayOfTheMonth() bool`

HasScheduledDayOfTheMonth returns a boolean if a field has been set.

### GetScheduledTimeOfDay

`func (o *RecurringTaskChainResponse) GetScheduledTimeOfDay() []string`

GetScheduledTimeOfDay returns the ScheduledTimeOfDay field if non-nil, zero value otherwise.

### GetScheduledTimeOfDayOk

`func (o *RecurringTaskChainResponse) GetScheduledTimeOfDayOk() (*[]string, bool)`

GetScheduledTimeOfDayOk returns a tuple with the ScheduledTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeOfDay

`func (o *RecurringTaskChainResponse) SetScheduledTimeOfDay(v []string)`

SetScheduledTimeOfDay sets ScheduledTimeOfDay field to given value.


### GetTimeZone

`func (o *RecurringTaskChainResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *RecurringTaskChainResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *RecurringTaskChainResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *RecurringTaskChainResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetInterruptedByShutdownBehavior

`func (o *RecurringTaskChainResponse) GetInterruptedByShutdownBehavior() EnumrecurringTaskChainInterruptedByShutdownBehaviorProp`

GetInterruptedByShutdownBehavior returns the InterruptedByShutdownBehavior field if non-nil, zero value otherwise.

### GetInterruptedByShutdownBehaviorOk

`func (o *RecurringTaskChainResponse) GetInterruptedByShutdownBehaviorOk() (*EnumrecurringTaskChainInterruptedByShutdownBehaviorProp, bool)`

GetInterruptedByShutdownBehaviorOk returns a tuple with the InterruptedByShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedByShutdownBehavior

`func (o *RecurringTaskChainResponse) SetInterruptedByShutdownBehavior(v EnumrecurringTaskChainInterruptedByShutdownBehaviorProp)`

SetInterruptedByShutdownBehavior sets InterruptedByShutdownBehavior field to given value.

### HasInterruptedByShutdownBehavior

`func (o *RecurringTaskChainResponse) HasInterruptedByShutdownBehavior() bool`

HasInterruptedByShutdownBehavior returns a boolean if a field has been set.

### GetServerOfflineAtStartTimeBehavior

`func (o *RecurringTaskChainResponse) GetServerOfflineAtStartTimeBehavior() EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp`

GetServerOfflineAtStartTimeBehavior returns the ServerOfflineAtStartTimeBehavior field if non-nil, zero value otherwise.

### GetServerOfflineAtStartTimeBehaviorOk

`func (o *RecurringTaskChainResponse) GetServerOfflineAtStartTimeBehaviorOk() (*EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp, bool)`

GetServerOfflineAtStartTimeBehaviorOk returns a tuple with the ServerOfflineAtStartTimeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOfflineAtStartTimeBehavior

`func (o *RecurringTaskChainResponse) SetServerOfflineAtStartTimeBehavior(v EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp)`

SetServerOfflineAtStartTimeBehavior sets ServerOfflineAtStartTimeBehavior field to given value.

### HasServerOfflineAtStartTimeBehavior

`func (o *RecurringTaskChainResponse) HasServerOfflineAtStartTimeBehavior() bool`

HasServerOfflineAtStartTimeBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


