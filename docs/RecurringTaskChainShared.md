# RecurringTaskChainShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumrecurringTaskChainSchemaUrn**](EnumrecurringTaskChainSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task Chain | [optional] 
**Enabled** | **bool** | Indicates whether this Recurring Task Chain is enabled for use. Recurring Task Chains that are disabled will not have any new instances scheduled, but instances that are already scheduled will be preserved. Those instances may be manually canceled if desired. | 
**RecurringTask** | **[]string** |  | 
**ScheduledMonth** | [**[]EnumrecurringTaskChainScheduledMonthProp**](EnumrecurringTaskChainScheduledMonthProp.md) |  | 
**ScheduledDateSelectionType** | [**EnumrecurringTaskChainScheduledDateSelectionTypeProp**](EnumrecurringTaskChainScheduledDateSelectionTypeProp.md) |  | 
**ScheduledDayOfTheWeek** | Pointer to [**[]EnumrecurringTaskChainScheduledDayOfTheWeekProp**](EnumrecurringTaskChainScheduledDayOfTheWeekProp.md) |  | [optional] 
**ScheduledDayOfTheMonth** | Pointer to [**[]EnumrecurringTaskChainScheduledDayOfTheMonthProp**](EnumrecurringTaskChainScheduledDayOfTheMonthProp.md) |  | [optional] 
**ScheduledTimeOfDay** | **[]string** |  | 
**TimeZone** | Pointer to **string** | The time zone that will be used to interpret the scheduled-time-of-day values. If no value is provided, then the JVM&#39;s default time zone will be used. | [optional] 
**InterruptedByShutdownBehavior** | Pointer to [**EnumrecurringTaskChainInterruptedByShutdownBehaviorProp**](EnumrecurringTaskChainInterruptedByShutdownBehaviorProp.md) |  | [optional] 
**ServerOfflineAtStartTimeBehavior** | Pointer to [**EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp**](EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp.md) |  | [optional] 

## Methods

### NewRecurringTaskChainShared

`func NewRecurringTaskChainShared(enabled bool, recurringTask []string, scheduledMonth []EnumrecurringTaskChainScheduledMonthProp, scheduledDateSelectionType EnumrecurringTaskChainScheduledDateSelectionTypeProp, scheduledTimeOfDay []string, ) *RecurringTaskChainShared`

NewRecurringTaskChainShared instantiates a new RecurringTaskChainShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringTaskChainSharedWithDefaults

`func NewRecurringTaskChainSharedWithDefaults() *RecurringTaskChainShared`

NewRecurringTaskChainSharedWithDefaults instantiates a new RecurringTaskChainShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *RecurringTaskChainShared) GetSchemas() []EnumrecurringTaskChainSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RecurringTaskChainShared) GetSchemasOk() (*[]EnumrecurringTaskChainSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RecurringTaskChainShared) SetSchemas(v []EnumrecurringTaskChainSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *RecurringTaskChainShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *RecurringTaskChainShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurringTaskChainShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurringTaskChainShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurringTaskChainShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RecurringTaskChainShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecurringTaskChainShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecurringTaskChainShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRecurringTask

`func (o *RecurringTaskChainShared) GetRecurringTask() []string`

GetRecurringTask returns the RecurringTask field if non-nil, zero value otherwise.

### GetRecurringTaskOk

`func (o *RecurringTaskChainShared) GetRecurringTaskOk() (*[]string, bool)`

GetRecurringTaskOk returns a tuple with the RecurringTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTask

`func (o *RecurringTaskChainShared) SetRecurringTask(v []string)`

SetRecurringTask sets RecurringTask field to given value.


### GetScheduledMonth

`func (o *RecurringTaskChainShared) GetScheduledMonth() []EnumrecurringTaskChainScheduledMonthProp`

GetScheduledMonth returns the ScheduledMonth field if non-nil, zero value otherwise.

### GetScheduledMonthOk

`func (o *RecurringTaskChainShared) GetScheduledMonthOk() (*[]EnumrecurringTaskChainScheduledMonthProp, bool)`

GetScheduledMonthOk returns a tuple with the ScheduledMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMonth

`func (o *RecurringTaskChainShared) SetScheduledMonth(v []EnumrecurringTaskChainScheduledMonthProp)`

SetScheduledMonth sets ScheduledMonth field to given value.


### GetScheduledDateSelectionType

`func (o *RecurringTaskChainShared) GetScheduledDateSelectionType() EnumrecurringTaskChainScheduledDateSelectionTypeProp`

GetScheduledDateSelectionType returns the ScheduledDateSelectionType field if non-nil, zero value otherwise.

### GetScheduledDateSelectionTypeOk

`func (o *RecurringTaskChainShared) GetScheduledDateSelectionTypeOk() (*EnumrecurringTaskChainScheduledDateSelectionTypeProp, bool)`

GetScheduledDateSelectionTypeOk returns a tuple with the ScheduledDateSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDateSelectionType

`func (o *RecurringTaskChainShared) SetScheduledDateSelectionType(v EnumrecurringTaskChainScheduledDateSelectionTypeProp)`

SetScheduledDateSelectionType sets ScheduledDateSelectionType field to given value.


### GetScheduledDayOfTheWeek

`func (o *RecurringTaskChainShared) GetScheduledDayOfTheWeek() []EnumrecurringTaskChainScheduledDayOfTheWeekProp`

GetScheduledDayOfTheWeek returns the ScheduledDayOfTheWeek field if non-nil, zero value otherwise.

### GetScheduledDayOfTheWeekOk

`func (o *RecurringTaskChainShared) GetScheduledDayOfTheWeekOk() (*[]EnumrecurringTaskChainScheduledDayOfTheWeekProp, bool)`

GetScheduledDayOfTheWeekOk returns a tuple with the ScheduledDayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDayOfTheWeek

`func (o *RecurringTaskChainShared) SetScheduledDayOfTheWeek(v []EnumrecurringTaskChainScheduledDayOfTheWeekProp)`

SetScheduledDayOfTheWeek sets ScheduledDayOfTheWeek field to given value.

### HasScheduledDayOfTheWeek

`func (o *RecurringTaskChainShared) HasScheduledDayOfTheWeek() bool`

HasScheduledDayOfTheWeek returns a boolean if a field has been set.

### GetScheduledDayOfTheMonth

`func (o *RecurringTaskChainShared) GetScheduledDayOfTheMonth() []EnumrecurringTaskChainScheduledDayOfTheMonthProp`

GetScheduledDayOfTheMonth returns the ScheduledDayOfTheMonth field if non-nil, zero value otherwise.

### GetScheduledDayOfTheMonthOk

`func (o *RecurringTaskChainShared) GetScheduledDayOfTheMonthOk() (*[]EnumrecurringTaskChainScheduledDayOfTheMonthProp, bool)`

GetScheduledDayOfTheMonthOk returns a tuple with the ScheduledDayOfTheMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDayOfTheMonth

`func (o *RecurringTaskChainShared) SetScheduledDayOfTheMonth(v []EnumrecurringTaskChainScheduledDayOfTheMonthProp)`

SetScheduledDayOfTheMonth sets ScheduledDayOfTheMonth field to given value.

### HasScheduledDayOfTheMonth

`func (o *RecurringTaskChainShared) HasScheduledDayOfTheMonth() bool`

HasScheduledDayOfTheMonth returns a boolean if a field has been set.

### GetScheduledTimeOfDay

`func (o *RecurringTaskChainShared) GetScheduledTimeOfDay() []string`

GetScheduledTimeOfDay returns the ScheduledTimeOfDay field if non-nil, zero value otherwise.

### GetScheduledTimeOfDayOk

`func (o *RecurringTaskChainShared) GetScheduledTimeOfDayOk() (*[]string, bool)`

GetScheduledTimeOfDayOk returns a tuple with the ScheduledTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeOfDay

`func (o *RecurringTaskChainShared) SetScheduledTimeOfDay(v []string)`

SetScheduledTimeOfDay sets ScheduledTimeOfDay field to given value.


### GetTimeZone

`func (o *RecurringTaskChainShared) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *RecurringTaskChainShared) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *RecurringTaskChainShared) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *RecurringTaskChainShared) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetInterruptedByShutdownBehavior

`func (o *RecurringTaskChainShared) GetInterruptedByShutdownBehavior() EnumrecurringTaskChainInterruptedByShutdownBehaviorProp`

GetInterruptedByShutdownBehavior returns the InterruptedByShutdownBehavior field if non-nil, zero value otherwise.

### GetInterruptedByShutdownBehaviorOk

`func (o *RecurringTaskChainShared) GetInterruptedByShutdownBehaviorOk() (*EnumrecurringTaskChainInterruptedByShutdownBehaviorProp, bool)`

GetInterruptedByShutdownBehaviorOk returns a tuple with the InterruptedByShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedByShutdownBehavior

`func (o *RecurringTaskChainShared) SetInterruptedByShutdownBehavior(v EnumrecurringTaskChainInterruptedByShutdownBehaviorProp)`

SetInterruptedByShutdownBehavior sets InterruptedByShutdownBehavior field to given value.

### HasInterruptedByShutdownBehavior

`func (o *RecurringTaskChainShared) HasInterruptedByShutdownBehavior() bool`

HasInterruptedByShutdownBehavior returns a boolean if a field has been set.

### GetServerOfflineAtStartTimeBehavior

`func (o *RecurringTaskChainShared) GetServerOfflineAtStartTimeBehavior() EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp`

GetServerOfflineAtStartTimeBehavior returns the ServerOfflineAtStartTimeBehavior field if non-nil, zero value otherwise.

### GetServerOfflineAtStartTimeBehaviorOk

`func (o *RecurringTaskChainShared) GetServerOfflineAtStartTimeBehaviorOk() (*EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp, bool)`

GetServerOfflineAtStartTimeBehaviorOk returns a tuple with the ServerOfflineAtStartTimeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOfflineAtStartTimeBehavior

`func (o *RecurringTaskChainShared) SetServerOfflineAtStartTimeBehavior(v EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp)`

SetServerOfflineAtStartTimeBehavior sets ServerOfflineAtStartTimeBehavior field to given value.

### HasServerOfflineAtStartTimeBehavior

`func (o *RecurringTaskChainShared) HasServerOfflineAtStartTimeBehavior() bool`

HasServerOfflineAtStartTimeBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


