# AddRecurringTaskChainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainName** | **string** | Name of the new Recurring Task Chain | 
**Schemas** | Pointer to [**[]EnumrecurringTaskChainSchemaUrn**](EnumrecurringTaskChainSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task Chain | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Recurring Task Chain is enabled for use. Recurring Task Chains that are disabled will not have any new instances scheduled, but instances that are already scheduled will be preserved. Those instances may be manually canceled if desired. | [optional] 
**RecurringTask** | **[]string** | The set of recurring tasks that make up this chain. At least one value must be provided. If multiple values are given, then the task instances will be invoked in the order in which they are listed. | 
**ScheduledMonth** | Pointer to [**[]EnumrecurringTaskChainScheduledMonthProp**](EnumrecurringTaskChainScheduledMonthProp.md) | The months of the year in which instances of this Recurring Task Chain may be scheduled to start. | [optional] 
**ScheduledDateSelectionType** | [**EnumrecurringTaskChainScheduledDateSelectionTypeProp**](EnumrecurringTaskChainScheduledDateSelectionTypeProp.md) |  | 
**ScheduledDayOfTheWeek** | Pointer to [**[]EnumrecurringTaskChainScheduledDayOfTheWeekProp**](EnumrecurringTaskChainScheduledDayOfTheWeekProp.md) | The specific days of the week on which instances of this Recurring Task Chain may be scheduled to start. If the scheduled-day-selection-type property has a value of selected-days-of-the-week, then this property must have one or more values; otherwise, it must be left undefined. | [optional] 
**ScheduledDayOfTheMonth** | Pointer to [**[]EnumrecurringTaskChainScheduledDayOfTheMonthProp**](EnumrecurringTaskChainScheduledDayOfTheMonthProp.md) | The specific days of the month on which instances of this Recurring Task Chain may be scheduled to start. If the scheduled-day-selection-type property has a value of selected-days-of-the-month, then this property must have one or more values; otherwise, it must be left undefined. | [optional] 
**ScheduledTimeOfDay** | **[]string** | The time of day at which instances of the Recurring Task Chain should be eligible to start running. Values should be in the format HH:MM (where HH is a two-digit representation of the hour of the day, between 00 and 23, inclusive), and MM is a two-digit representation of the minute of the hour (between 00 and 59, inclusive). Alternately, the value can be in the form *:MM, which indicates that the task should be eligible to start at the specified minute of every hour. At least one value must be provided, but multiple values may be given to indicate multiple start times within the same day. | 
**TimeZone** | Pointer to **string** | The time zone that will be used to interpret the scheduled-time-of-day values. If no value is provided, then the JVM&#39;s default time zone will be used. | [optional] 
**InterruptedByShutdownBehavior** | Pointer to [**EnumrecurringTaskChainInterruptedByShutdownBehaviorProp**](EnumrecurringTaskChainInterruptedByShutdownBehaviorProp.md) |  | [optional] 
**ServerOfflineAtStartTimeBehavior** | Pointer to [**EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp**](EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp.md) |  | [optional] 

## Methods

### NewAddRecurringTaskChainRequest

`func NewAddRecurringTaskChainRequest(chainName string, recurringTask []string, scheduledDateSelectionType EnumrecurringTaskChainScheduledDateSelectionTypeProp, scheduledTimeOfDay []string, ) *AddRecurringTaskChainRequest`

NewAddRecurringTaskChainRequest instantiates a new AddRecurringTaskChainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRecurringTaskChainRequestWithDefaults

`func NewAddRecurringTaskChainRequestWithDefaults() *AddRecurringTaskChainRequest`

NewAddRecurringTaskChainRequestWithDefaults instantiates a new AddRecurringTaskChainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainName

`func (o *AddRecurringTaskChainRequest) GetChainName() string`

GetChainName returns the ChainName field if non-nil, zero value otherwise.

### GetChainNameOk

`func (o *AddRecurringTaskChainRequest) GetChainNameOk() (*string, bool)`

GetChainNameOk returns a tuple with the ChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainName

`func (o *AddRecurringTaskChainRequest) SetChainName(v string)`

SetChainName sets ChainName field to given value.


### GetSchemas

`func (o *AddRecurringTaskChainRequest) GetSchemas() []EnumrecurringTaskChainSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRecurringTaskChainRequest) GetSchemasOk() (*[]EnumrecurringTaskChainSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRecurringTaskChainRequest) SetSchemas(v []EnumrecurringTaskChainSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddRecurringTaskChainRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddRecurringTaskChainRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRecurringTaskChainRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRecurringTaskChainRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRecurringTaskChainRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddRecurringTaskChainRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddRecurringTaskChainRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddRecurringTaskChainRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddRecurringTaskChainRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRecurringTask

`func (o *AddRecurringTaskChainRequest) GetRecurringTask() []string`

GetRecurringTask returns the RecurringTask field if non-nil, zero value otherwise.

### GetRecurringTaskOk

`func (o *AddRecurringTaskChainRequest) GetRecurringTaskOk() (*[]string, bool)`

GetRecurringTaskOk returns a tuple with the RecurringTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTask

`func (o *AddRecurringTaskChainRequest) SetRecurringTask(v []string)`

SetRecurringTask sets RecurringTask field to given value.


### GetScheduledMonth

`func (o *AddRecurringTaskChainRequest) GetScheduledMonth() []EnumrecurringTaskChainScheduledMonthProp`

GetScheduledMonth returns the ScheduledMonth field if non-nil, zero value otherwise.

### GetScheduledMonthOk

`func (o *AddRecurringTaskChainRequest) GetScheduledMonthOk() (*[]EnumrecurringTaskChainScheduledMonthProp, bool)`

GetScheduledMonthOk returns a tuple with the ScheduledMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMonth

`func (o *AddRecurringTaskChainRequest) SetScheduledMonth(v []EnumrecurringTaskChainScheduledMonthProp)`

SetScheduledMonth sets ScheduledMonth field to given value.

### HasScheduledMonth

`func (o *AddRecurringTaskChainRequest) HasScheduledMonth() bool`

HasScheduledMonth returns a boolean if a field has been set.

### GetScheduledDateSelectionType

`func (o *AddRecurringTaskChainRequest) GetScheduledDateSelectionType() EnumrecurringTaskChainScheduledDateSelectionTypeProp`

GetScheduledDateSelectionType returns the ScheduledDateSelectionType field if non-nil, zero value otherwise.

### GetScheduledDateSelectionTypeOk

`func (o *AddRecurringTaskChainRequest) GetScheduledDateSelectionTypeOk() (*EnumrecurringTaskChainScheduledDateSelectionTypeProp, bool)`

GetScheduledDateSelectionTypeOk returns a tuple with the ScheduledDateSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDateSelectionType

`func (o *AddRecurringTaskChainRequest) SetScheduledDateSelectionType(v EnumrecurringTaskChainScheduledDateSelectionTypeProp)`

SetScheduledDateSelectionType sets ScheduledDateSelectionType field to given value.


### GetScheduledDayOfTheWeek

`func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheWeek() []EnumrecurringTaskChainScheduledDayOfTheWeekProp`

GetScheduledDayOfTheWeek returns the ScheduledDayOfTheWeek field if non-nil, zero value otherwise.

### GetScheduledDayOfTheWeekOk

`func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheWeekOk() (*[]EnumrecurringTaskChainScheduledDayOfTheWeekProp, bool)`

GetScheduledDayOfTheWeekOk returns a tuple with the ScheduledDayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDayOfTheWeek

`func (o *AddRecurringTaskChainRequest) SetScheduledDayOfTheWeek(v []EnumrecurringTaskChainScheduledDayOfTheWeekProp)`

SetScheduledDayOfTheWeek sets ScheduledDayOfTheWeek field to given value.

### HasScheduledDayOfTheWeek

`func (o *AddRecurringTaskChainRequest) HasScheduledDayOfTheWeek() bool`

HasScheduledDayOfTheWeek returns a boolean if a field has been set.

### GetScheduledDayOfTheMonth

`func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheMonth() []EnumrecurringTaskChainScheduledDayOfTheMonthProp`

GetScheduledDayOfTheMonth returns the ScheduledDayOfTheMonth field if non-nil, zero value otherwise.

### GetScheduledDayOfTheMonthOk

`func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheMonthOk() (*[]EnumrecurringTaskChainScheduledDayOfTheMonthProp, bool)`

GetScheduledDayOfTheMonthOk returns a tuple with the ScheduledDayOfTheMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDayOfTheMonth

`func (o *AddRecurringTaskChainRequest) SetScheduledDayOfTheMonth(v []EnumrecurringTaskChainScheduledDayOfTheMonthProp)`

SetScheduledDayOfTheMonth sets ScheduledDayOfTheMonth field to given value.

### HasScheduledDayOfTheMonth

`func (o *AddRecurringTaskChainRequest) HasScheduledDayOfTheMonth() bool`

HasScheduledDayOfTheMonth returns a boolean if a field has been set.

### GetScheduledTimeOfDay

`func (o *AddRecurringTaskChainRequest) GetScheduledTimeOfDay() []string`

GetScheduledTimeOfDay returns the ScheduledTimeOfDay field if non-nil, zero value otherwise.

### GetScheduledTimeOfDayOk

`func (o *AddRecurringTaskChainRequest) GetScheduledTimeOfDayOk() (*[]string, bool)`

GetScheduledTimeOfDayOk returns a tuple with the ScheduledTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeOfDay

`func (o *AddRecurringTaskChainRequest) SetScheduledTimeOfDay(v []string)`

SetScheduledTimeOfDay sets ScheduledTimeOfDay field to given value.


### GetTimeZone

`func (o *AddRecurringTaskChainRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AddRecurringTaskChainRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AddRecurringTaskChainRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AddRecurringTaskChainRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetInterruptedByShutdownBehavior

`func (o *AddRecurringTaskChainRequest) GetInterruptedByShutdownBehavior() EnumrecurringTaskChainInterruptedByShutdownBehaviorProp`

GetInterruptedByShutdownBehavior returns the InterruptedByShutdownBehavior field if non-nil, zero value otherwise.

### GetInterruptedByShutdownBehaviorOk

`func (o *AddRecurringTaskChainRequest) GetInterruptedByShutdownBehaviorOk() (*EnumrecurringTaskChainInterruptedByShutdownBehaviorProp, bool)`

GetInterruptedByShutdownBehaviorOk returns a tuple with the InterruptedByShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedByShutdownBehavior

`func (o *AddRecurringTaskChainRequest) SetInterruptedByShutdownBehavior(v EnumrecurringTaskChainInterruptedByShutdownBehaviorProp)`

SetInterruptedByShutdownBehavior sets InterruptedByShutdownBehavior field to given value.

### HasInterruptedByShutdownBehavior

`func (o *AddRecurringTaskChainRequest) HasInterruptedByShutdownBehavior() bool`

HasInterruptedByShutdownBehavior returns a boolean if a field has been set.

### GetServerOfflineAtStartTimeBehavior

`func (o *AddRecurringTaskChainRequest) GetServerOfflineAtStartTimeBehavior() EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp`

GetServerOfflineAtStartTimeBehavior returns the ServerOfflineAtStartTimeBehavior field if non-nil, zero value otherwise.

### GetServerOfflineAtStartTimeBehaviorOk

`func (o *AddRecurringTaskChainRequest) GetServerOfflineAtStartTimeBehaviorOk() (*EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp, bool)`

GetServerOfflineAtStartTimeBehaviorOk returns a tuple with the ServerOfflineAtStartTimeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOfflineAtStartTimeBehavior

`func (o *AddRecurringTaskChainRequest) SetServerOfflineAtStartTimeBehavior(v EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp)`

SetServerOfflineAtStartTimeBehavior sets ServerOfflineAtStartTimeBehavior field to given value.

### HasServerOfflineAtStartTimeBehavior

`func (o *AddRecurringTaskChainRequest) HasServerOfflineAtStartTimeBehavior() bool`

HasServerOfflineAtStartTimeBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


