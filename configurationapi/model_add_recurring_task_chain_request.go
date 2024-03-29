/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddRecurringTaskChainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRecurringTaskChainRequest{}

// AddRecurringTaskChainRequest struct for AddRecurringTaskChainRequest
type AddRecurringTaskChainRequest struct {
	Schemas []EnumrecurringTaskChainSchemaUrn `json:"schemas,omitempty"`
	// A description for this Recurring Task Chain
	Description *string `json:"description,omitempty"`
	// Indicates whether this Recurring Task Chain is enabled for use. Recurring Task Chains that are disabled will not have any new instances scheduled, but instances that are already scheduled will be preserved. Those instances may be manually canceled if desired.
	Enabled *bool `json:"enabled,omitempty"`
	// The set of recurring tasks that make up this chain. At least one value must be provided. If multiple values are given, then the task instances will be invoked in the order in which they are listed.
	RecurringTask              []string                                             `json:"recurringTask"`
	ScheduledMonth             []EnumrecurringTaskChainScheduledMonthProp           `json:"scheduledMonth,omitempty"`
	ScheduledDateSelectionType EnumrecurringTaskChainScheduledDateSelectionTypeProp `json:"scheduledDateSelectionType"`
	ScheduledDayOfTheWeek      []EnumrecurringTaskChainScheduledDayOfTheWeekProp    `json:"scheduledDayOfTheWeek,omitempty"`
	ScheduledDayOfTheMonth     []EnumrecurringTaskChainScheduledDayOfTheMonthProp   `json:"scheduledDayOfTheMonth,omitempty"`
	// The time of day at which instances of the Recurring Task Chain should be eligible to start running. Values should be in the format HH:MM (where HH is a two-digit representation of the hour of the day, between 00 and 23, inclusive), and MM is a two-digit representation of the minute of the hour (between 00 and 59, inclusive). Alternately, the value can be in the form *:MM, which indicates that the task should be eligible to start at the specified minute of every hour. At least one value must be provided, but multiple values may be given to indicate multiple start times within the same day.
	ScheduledTimeOfDay []string `json:"scheduledTimeOfDay"`
	// The time zone that will be used to interpret the scheduled-time-of-day values. If no value is provided, then the JVM's default time zone will be used.
	TimeZone                         *string                                                     `json:"timeZone,omitempty"`
	InterruptedByShutdownBehavior    *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp    `json:"interruptedByShutdownBehavior,omitempty"`
	ServerOfflineAtStartTimeBehavior *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp `json:"serverOfflineAtStartTimeBehavior,omitempty"`
	// Name of the new Recurring Task Chain
	ChainName string `json:"chainName"`
}

// NewAddRecurringTaskChainRequest instantiates a new AddRecurringTaskChainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRecurringTaskChainRequest(recurringTask []string, scheduledDateSelectionType EnumrecurringTaskChainScheduledDateSelectionTypeProp, scheduledTimeOfDay []string, chainName string) *AddRecurringTaskChainRequest {
	this := AddRecurringTaskChainRequest{}
	this.RecurringTask = recurringTask
	this.ScheduledDateSelectionType = scheduledDateSelectionType
	this.ScheduledTimeOfDay = scheduledTimeOfDay
	this.ChainName = chainName
	return &this
}

// NewAddRecurringTaskChainRequestWithDefaults instantiates a new AddRecurringTaskChainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRecurringTaskChainRequestWithDefaults() *AddRecurringTaskChainRequest {
	this := AddRecurringTaskChainRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetSchemas() []EnumrecurringTaskChainSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumrecurringTaskChainSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetSchemasOk() ([]EnumrecurringTaskChainSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumrecurringTaskChainSchemaUrn and assigns it to the Schemas field.
func (o *AddRecurringTaskChainRequest) SetSchemas(v []EnumrecurringTaskChainSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddRecurringTaskChainRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddRecurringTaskChainRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRecurringTask returns the RecurringTask field value
func (o *AddRecurringTaskChainRequest) GetRecurringTask() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecurringTask
}

// GetRecurringTaskOk returns a tuple with the RecurringTask field value
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetRecurringTaskOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringTask, true
}

// SetRecurringTask sets field value
func (o *AddRecurringTaskChainRequest) SetRecurringTask(v []string) {
	o.RecurringTask = v
}

// GetScheduledMonth returns the ScheduledMonth field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetScheduledMonth() []EnumrecurringTaskChainScheduledMonthProp {
	if o == nil || IsNil(o.ScheduledMonth) {
		var ret []EnumrecurringTaskChainScheduledMonthProp
		return ret
	}
	return o.ScheduledMonth
}

// GetScheduledMonthOk returns a tuple with the ScheduledMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetScheduledMonthOk() ([]EnumrecurringTaskChainScheduledMonthProp, bool) {
	if o == nil || IsNil(o.ScheduledMonth) {
		return nil, false
	}
	return o.ScheduledMonth, true
}

// HasScheduledMonth returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasScheduledMonth() bool {
	if o != nil && !IsNil(o.ScheduledMonth) {
		return true
	}

	return false
}

// SetScheduledMonth gets a reference to the given []EnumrecurringTaskChainScheduledMonthProp and assigns it to the ScheduledMonth field.
func (o *AddRecurringTaskChainRequest) SetScheduledMonth(v []EnumrecurringTaskChainScheduledMonthProp) {
	o.ScheduledMonth = v
}

// GetScheduledDateSelectionType returns the ScheduledDateSelectionType field value
func (o *AddRecurringTaskChainRequest) GetScheduledDateSelectionType() EnumrecurringTaskChainScheduledDateSelectionTypeProp {
	if o == nil {
		var ret EnumrecurringTaskChainScheduledDateSelectionTypeProp
		return ret
	}

	return o.ScheduledDateSelectionType
}

// GetScheduledDateSelectionTypeOk returns a tuple with the ScheduledDateSelectionType field value
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetScheduledDateSelectionTypeOk() (*EnumrecurringTaskChainScheduledDateSelectionTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledDateSelectionType, true
}

// SetScheduledDateSelectionType sets field value
func (o *AddRecurringTaskChainRequest) SetScheduledDateSelectionType(v EnumrecurringTaskChainScheduledDateSelectionTypeProp) {
	o.ScheduledDateSelectionType = v
}

// GetScheduledDayOfTheWeek returns the ScheduledDayOfTheWeek field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheWeek() []EnumrecurringTaskChainScheduledDayOfTheWeekProp {
	if o == nil || IsNil(o.ScheduledDayOfTheWeek) {
		var ret []EnumrecurringTaskChainScheduledDayOfTheWeekProp
		return ret
	}
	return o.ScheduledDayOfTheWeek
}

// GetScheduledDayOfTheWeekOk returns a tuple with the ScheduledDayOfTheWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheWeekOk() ([]EnumrecurringTaskChainScheduledDayOfTheWeekProp, bool) {
	if o == nil || IsNil(o.ScheduledDayOfTheWeek) {
		return nil, false
	}
	return o.ScheduledDayOfTheWeek, true
}

// HasScheduledDayOfTheWeek returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasScheduledDayOfTheWeek() bool {
	if o != nil && !IsNil(o.ScheduledDayOfTheWeek) {
		return true
	}

	return false
}

// SetScheduledDayOfTheWeek gets a reference to the given []EnumrecurringTaskChainScheduledDayOfTheWeekProp and assigns it to the ScheduledDayOfTheWeek field.
func (o *AddRecurringTaskChainRequest) SetScheduledDayOfTheWeek(v []EnumrecurringTaskChainScheduledDayOfTheWeekProp) {
	o.ScheduledDayOfTheWeek = v
}

// GetScheduledDayOfTheMonth returns the ScheduledDayOfTheMonth field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheMonth() []EnumrecurringTaskChainScheduledDayOfTheMonthProp {
	if o == nil || IsNil(o.ScheduledDayOfTheMonth) {
		var ret []EnumrecurringTaskChainScheduledDayOfTheMonthProp
		return ret
	}
	return o.ScheduledDayOfTheMonth
}

// GetScheduledDayOfTheMonthOk returns a tuple with the ScheduledDayOfTheMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetScheduledDayOfTheMonthOk() ([]EnumrecurringTaskChainScheduledDayOfTheMonthProp, bool) {
	if o == nil || IsNil(o.ScheduledDayOfTheMonth) {
		return nil, false
	}
	return o.ScheduledDayOfTheMonth, true
}

// HasScheduledDayOfTheMonth returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasScheduledDayOfTheMonth() bool {
	if o != nil && !IsNil(o.ScheduledDayOfTheMonth) {
		return true
	}

	return false
}

// SetScheduledDayOfTheMonth gets a reference to the given []EnumrecurringTaskChainScheduledDayOfTheMonthProp and assigns it to the ScheduledDayOfTheMonth field.
func (o *AddRecurringTaskChainRequest) SetScheduledDayOfTheMonth(v []EnumrecurringTaskChainScheduledDayOfTheMonthProp) {
	o.ScheduledDayOfTheMonth = v
}

// GetScheduledTimeOfDay returns the ScheduledTimeOfDay field value
func (o *AddRecurringTaskChainRequest) GetScheduledTimeOfDay() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScheduledTimeOfDay
}

// GetScheduledTimeOfDayOk returns a tuple with the ScheduledTimeOfDay field value
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetScheduledTimeOfDayOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledTimeOfDay, true
}

// SetScheduledTimeOfDay sets field value
func (o *AddRecurringTaskChainRequest) SetScheduledTimeOfDay(v []string) {
	o.ScheduledTimeOfDay = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AddRecurringTaskChainRequest) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetInterruptedByShutdownBehavior returns the InterruptedByShutdownBehavior field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetInterruptedByShutdownBehavior() EnumrecurringTaskChainInterruptedByShutdownBehaviorProp {
	if o == nil || IsNil(o.InterruptedByShutdownBehavior) {
		var ret EnumrecurringTaskChainInterruptedByShutdownBehaviorProp
		return ret
	}
	return *o.InterruptedByShutdownBehavior
}

// GetInterruptedByShutdownBehaviorOk returns a tuple with the InterruptedByShutdownBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetInterruptedByShutdownBehaviorOk() (*EnumrecurringTaskChainInterruptedByShutdownBehaviorProp, bool) {
	if o == nil || IsNil(o.InterruptedByShutdownBehavior) {
		return nil, false
	}
	return o.InterruptedByShutdownBehavior, true
}

// HasInterruptedByShutdownBehavior returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasInterruptedByShutdownBehavior() bool {
	if o != nil && !IsNil(o.InterruptedByShutdownBehavior) {
		return true
	}

	return false
}

// SetInterruptedByShutdownBehavior gets a reference to the given EnumrecurringTaskChainInterruptedByShutdownBehaviorProp and assigns it to the InterruptedByShutdownBehavior field.
func (o *AddRecurringTaskChainRequest) SetInterruptedByShutdownBehavior(v EnumrecurringTaskChainInterruptedByShutdownBehaviorProp) {
	o.InterruptedByShutdownBehavior = &v
}

// GetServerOfflineAtStartTimeBehavior returns the ServerOfflineAtStartTimeBehavior field value if set, zero value otherwise.
func (o *AddRecurringTaskChainRequest) GetServerOfflineAtStartTimeBehavior() EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp {
	if o == nil || IsNil(o.ServerOfflineAtStartTimeBehavior) {
		var ret EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp
		return ret
	}
	return *o.ServerOfflineAtStartTimeBehavior
}

// GetServerOfflineAtStartTimeBehaviorOk returns a tuple with the ServerOfflineAtStartTimeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetServerOfflineAtStartTimeBehaviorOk() (*EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp, bool) {
	if o == nil || IsNil(o.ServerOfflineAtStartTimeBehavior) {
		return nil, false
	}
	return o.ServerOfflineAtStartTimeBehavior, true
}

// HasServerOfflineAtStartTimeBehavior returns a boolean if a field has been set.
func (o *AddRecurringTaskChainRequest) HasServerOfflineAtStartTimeBehavior() bool {
	if o != nil && !IsNil(o.ServerOfflineAtStartTimeBehavior) {
		return true
	}

	return false
}

// SetServerOfflineAtStartTimeBehavior gets a reference to the given EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp and assigns it to the ServerOfflineAtStartTimeBehavior field.
func (o *AddRecurringTaskChainRequest) SetServerOfflineAtStartTimeBehavior(v EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) {
	o.ServerOfflineAtStartTimeBehavior = &v
}

// GetChainName returns the ChainName field value
func (o *AddRecurringTaskChainRequest) GetChainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainName
}

// GetChainNameOk returns a tuple with the ChainName field value
// and a boolean to check if the value has been set.
func (o *AddRecurringTaskChainRequest) GetChainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainName, true
}

// SetChainName sets field value
func (o *AddRecurringTaskChainRequest) SetChainName(v string) {
	o.ChainName = v
}

func (o AddRecurringTaskChainRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddRecurringTaskChainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["recurringTask"] = o.RecurringTask
	if !IsNil(o.ScheduledMonth) {
		toSerialize["scheduledMonth"] = o.ScheduledMonth
	}
	toSerialize["scheduledDateSelectionType"] = o.ScheduledDateSelectionType
	if !IsNil(o.ScheduledDayOfTheWeek) {
		toSerialize["scheduledDayOfTheWeek"] = o.ScheduledDayOfTheWeek
	}
	if !IsNil(o.ScheduledDayOfTheMonth) {
		toSerialize["scheduledDayOfTheMonth"] = o.ScheduledDayOfTheMonth
	}
	toSerialize["scheduledTimeOfDay"] = o.ScheduledTimeOfDay
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.InterruptedByShutdownBehavior) {
		toSerialize["interruptedByShutdownBehavior"] = o.InterruptedByShutdownBehavior
	}
	if !IsNil(o.ServerOfflineAtStartTimeBehavior) {
		toSerialize["serverOfflineAtStartTimeBehavior"] = o.ServerOfflineAtStartTimeBehavior
	}
	toSerialize["chainName"] = o.ChainName
	return toSerialize, nil
}

type NullableAddRecurringTaskChainRequest struct {
	value *AddRecurringTaskChainRequest
	isSet bool
}

func (v NullableAddRecurringTaskChainRequest) Get() *AddRecurringTaskChainRequest {
	return v.value
}

func (v *NullableAddRecurringTaskChainRequest) Set(val *AddRecurringTaskChainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRecurringTaskChainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRecurringTaskChainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRecurringTaskChainRequest(val *AddRecurringTaskChainRequest) *NullableAddRecurringTaskChainRequest {
	return &NullableAddRecurringTaskChainRequest{value: val, isSet: true}
}

func (v NullableAddRecurringTaskChainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRecurringTaskChainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
