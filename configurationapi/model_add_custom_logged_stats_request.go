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

// checks if the AddCustomLoggedStatsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCustomLoggedStatsRequest{}

// AddCustomLoggedStatsRequest struct for AddCustomLoggedStatsRequest
type AddCustomLoggedStatsRequest struct {
	Schemas []EnumcustomLoggedStatsSchemaUrn `json:"schemas"`
	// A description for this Custom Logged Stats
	Description *string `json:"description,omitempty"`
	// Indicates whether the Custom Logged Stats object is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The objectclass name of the monitor entries to examine for generating these statistics.
	MonitorObjectclass string `json:"monitorObjectclass"`
	// An optional LDAP filter that can be used restrict which monitor entries are used to produce the output.
	IncludeFilter *string `json:"includeFilter,omitempty"`
	// Specifies the attributes on the monitor entries that should be included in the output.
	AttributeToLog []string `json:"attributeToLog"`
	// Optionally, specifies an explicit name for each column header instead of having these names automatically generated from the monitored attribute name.
	ColumnName    []string                                 `json:"columnName,omitempty"`
	StatisticType []EnumcustomLoggedStatsStatisticTypeProp `json:"statisticType"`
	// An optional prefix that is included in the header before the column name.
	HeaderPrefix *string `json:"headerPrefix,omitempty"`
	// An optional attribute from the monitor entry that is included as a prefix before the column name in the column header.
	HeaderPrefixAttribute *string `json:"headerPrefixAttribute,omitempty"`
	// An optional regular expression pattern, that when used in conjunction with regex-replacement, can alter the value of the attribute being monitored.
	RegexPattern *string `json:"regexPattern,omitempty"`
	// An optional regular expression replacement value, that when used in conjunction with regex-pattern, can alter the value of the attribute being monitored.
	RegexReplacement *string `json:"regexReplacement,omitempty"`
	// An optional floating point value that can be used to scale the resulting value.
	DivideValueBy *string `json:"divideValueBy,omitempty"`
	// An optional property that can scale the resulting value by another attribute in the monitored entry.
	DivideValueByAttribute *string `json:"divideValueByAttribute,omitempty"`
	// This provides a way to format the monitored attribute value in the output to control the precision for instance.
	DecimalFormat *string `json:"decimalFormat,omitempty"`
	// If this property is set to true, then the value of any of the monitored attributes here can contribute to whether an interval is considered \"idle\" by the Periodic Stats Logger.
	NonZeroImpliesNotIdle *bool `json:"nonZeroImpliesNotIdle,omitempty"`
	// Name of the new Custom Logged Stats
	StatsName string `json:"statsName"`
}

// NewAddCustomLoggedStatsRequest instantiates a new AddCustomLoggedStatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCustomLoggedStatsRequest(schemas []EnumcustomLoggedStatsSchemaUrn, monitorObjectclass string, attributeToLog []string, statisticType []EnumcustomLoggedStatsStatisticTypeProp, statsName string) *AddCustomLoggedStatsRequest {
	this := AddCustomLoggedStatsRequest{}
	this.Schemas = schemas
	this.MonitorObjectclass = monitorObjectclass
	this.AttributeToLog = attributeToLog
	this.StatisticType = statisticType
	this.StatsName = statsName
	return &this
}

// NewAddCustomLoggedStatsRequestWithDefaults instantiates a new AddCustomLoggedStatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCustomLoggedStatsRequestWithDefaults() *AddCustomLoggedStatsRequest {
	this := AddCustomLoggedStatsRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddCustomLoggedStatsRequest) GetSchemas() []EnumcustomLoggedStatsSchemaUrn {
	if o == nil {
		var ret []EnumcustomLoggedStatsSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetSchemasOk() ([]EnumcustomLoggedStatsSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCustomLoggedStatsRequest) SetSchemas(v []EnumcustomLoggedStatsSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddCustomLoggedStatsRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddCustomLoggedStatsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMonitorObjectclass returns the MonitorObjectclass field value
func (o *AddCustomLoggedStatsRequest) GetMonitorObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorObjectclass
}

// GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field value
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetMonitorObjectclassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorObjectclass, true
}

// SetMonitorObjectclass sets field value
func (o *AddCustomLoggedStatsRequest) SetMonitorObjectclass(v string) {
	o.MonitorObjectclass = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetIncludeFilter() string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret string
		return ret
	}
	return *o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetIncludeFilterOk() (*string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given string and assigns it to the IncludeFilter field.
func (o *AddCustomLoggedStatsRequest) SetIncludeFilter(v string) {
	o.IncludeFilter = &v
}

// GetAttributeToLog returns the AttributeToLog field value
func (o *AddCustomLoggedStatsRequest) GetAttributeToLog() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttributeToLog
}

// GetAttributeToLogOk returns a tuple with the AttributeToLog field value
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetAttributeToLogOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeToLog, true
}

// SetAttributeToLog sets field value
func (o *AddCustomLoggedStatsRequest) SetAttributeToLog(v []string) {
	o.AttributeToLog = v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetColumnName() []string {
	if o == nil || IsNil(o.ColumnName) {
		var ret []string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetColumnNameOk() ([]string, bool) {
	if o == nil || IsNil(o.ColumnName) {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasColumnName() bool {
	if o != nil && !IsNil(o.ColumnName) {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given []string and assigns it to the ColumnName field.
func (o *AddCustomLoggedStatsRequest) SetColumnName(v []string) {
	o.ColumnName = v
}

// GetStatisticType returns the StatisticType field value
func (o *AddCustomLoggedStatsRequest) GetStatisticType() []EnumcustomLoggedStatsStatisticTypeProp {
	if o == nil {
		var ret []EnumcustomLoggedStatsStatisticTypeProp
		return ret
	}

	return o.StatisticType
}

// GetStatisticTypeOk returns a tuple with the StatisticType field value
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetStatisticTypeOk() ([]EnumcustomLoggedStatsStatisticTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatisticType, true
}

// SetStatisticType sets field value
func (o *AddCustomLoggedStatsRequest) SetStatisticType(v []EnumcustomLoggedStatsStatisticTypeProp) {
	o.StatisticType = v
}

// GetHeaderPrefix returns the HeaderPrefix field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetHeaderPrefix() string {
	if o == nil || IsNil(o.HeaderPrefix) {
		var ret string
		return ret
	}
	return *o.HeaderPrefix
}

// GetHeaderPrefixOk returns a tuple with the HeaderPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetHeaderPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderPrefix) {
		return nil, false
	}
	return o.HeaderPrefix, true
}

// HasHeaderPrefix returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasHeaderPrefix() bool {
	if o != nil && !IsNil(o.HeaderPrefix) {
		return true
	}

	return false
}

// SetHeaderPrefix gets a reference to the given string and assigns it to the HeaderPrefix field.
func (o *AddCustomLoggedStatsRequest) SetHeaderPrefix(v string) {
	o.HeaderPrefix = &v
}

// GetHeaderPrefixAttribute returns the HeaderPrefixAttribute field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetHeaderPrefixAttribute() string {
	if o == nil || IsNil(o.HeaderPrefixAttribute) {
		var ret string
		return ret
	}
	return *o.HeaderPrefixAttribute
}

// GetHeaderPrefixAttributeOk returns a tuple with the HeaderPrefixAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetHeaderPrefixAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderPrefixAttribute) {
		return nil, false
	}
	return o.HeaderPrefixAttribute, true
}

// HasHeaderPrefixAttribute returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasHeaderPrefixAttribute() bool {
	if o != nil && !IsNil(o.HeaderPrefixAttribute) {
		return true
	}

	return false
}

// SetHeaderPrefixAttribute gets a reference to the given string and assigns it to the HeaderPrefixAttribute field.
func (o *AddCustomLoggedStatsRequest) SetHeaderPrefixAttribute(v string) {
	o.HeaderPrefixAttribute = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetRegexPattern() string {
	if o == nil || IsNil(o.RegexPattern) {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetRegexPatternOk() (*string, bool) {
	if o == nil || IsNil(o.RegexPattern) {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasRegexPattern() bool {
	if o != nil && !IsNil(o.RegexPattern) {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *AddCustomLoggedStatsRequest) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetRegexReplacement returns the RegexReplacement field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetRegexReplacement() string {
	if o == nil || IsNil(o.RegexReplacement) {
		var ret string
		return ret
	}
	return *o.RegexReplacement
}

// GetRegexReplacementOk returns a tuple with the RegexReplacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetRegexReplacementOk() (*string, bool) {
	if o == nil || IsNil(o.RegexReplacement) {
		return nil, false
	}
	return o.RegexReplacement, true
}

// HasRegexReplacement returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasRegexReplacement() bool {
	if o != nil && !IsNil(o.RegexReplacement) {
		return true
	}

	return false
}

// SetRegexReplacement gets a reference to the given string and assigns it to the RegexReplacement field.
func (o *AddCustomLoggedStatsRequest) SetRegexReplacement(v string) {
	o.RegexReplacement = &v
}

// GetDivideValueBy returns the DivideValueBy field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetDivideValueBy() string {
	if o == nil || IsNil(o.DivideValueBy) {
		var ret string
		return ret
	}
	return *o.DivideValueBy
}

// GetDivideValueByOk returns a tuple with the DivideValueBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetDivideValueByOk() (*string, bool) {
	if o == nil || IsNil(o.DivideValueBy) {
		return nil, false
	}
	return o.DivideValueBy, true
}

// HasDivideValueBy returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasDivideValueBy() bool {
	if o != nil && !IsNil(o.DivideValueBy) {
		return true
	}

	return false
}

// SetDivideValueBy gets a reference to the given string and assigns it to the DivideValueBy field.
func (o *AddCustomLoggedStatsRequest) SetDivideValueBy(v string) {
	o.DivideValueBy = &v
}

// GetDivideValueByAttribute returns the DivideValueByAttribute field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetDivideValueByAttribute() string {
	if o == nil || IsNil(o.DivideValueByAttribute) {
		var ret string
		return ret
	}
	return *o.DivideValueByAttribute
}

// GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetDivideValueByAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.DivideValueByAttribute) {
		return nil, false
	}
	return o.DivideValueByAttribute, true
}

// HasDivideValueByAttribute returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasDivideValueByAttribute() bool {
	if o != nil && !IsNil(o.DivideValueByAttribute) {
		return true
	}

	return false
}

// SetDivideValueByAttribute gets a reference to the given string and assigns it to the DivideValueByAttribute field.
func (o *AddCustomLoggedStatsRequest) SetDivideValueByAttribute(v string) {
	o.DivideValueByAttribute = &v
}

// GetDecimalFormat returns the DecimalFormat field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetDecimalFormat() string {
	if o == nil || IsNil(o.DecimalFormat) {
		var ret string
		return ret
	}
	return *o.DecimalFormat
}

// GetDecimalFormatOk returns a tuple with the DecimalFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetDecimalFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DecimalFormat) {
		return nil, false
	}
	return o.DecimalFormat, true
}

// HasDecimalFormat returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasDecimalFormat() bool {
	if o != nil && !IsNil(o.DecimalFormat) {
		return true
	}

	return false
}

// SetDecimalFormat gets a reference to the given string and assigns it to the DecimalFormat field.
func (o *AddCustomLoggedStatsRequest) SetDecimalFormat(v string) {
	o.DecimalFormat = &v
}

// GetNonZeroImpliesNotIdle returns the NonZeroImpliesNotIdle field value if set, zero value otherwise.
func (o *AddCustomLoggedStatsRequest) GetNonZeroImpliesNotIdle() bool {
	if o == nil || IsNil(o.NonZeroImpliesNotIdle) {
		var ret bool
		return ret
	}
	return *o.NonZeroImpliesNotIdle
}

// GetNonZeroImpliesNotIdleOk returns a tuple with the NonZeroImpliesNotIdle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetNonZeroImpliesNotIdleOk() (*bool, bool) {
	if o == nil || IsNil(o.NonZeroImpliesNotIdle) {
		return nil, false
	}
	return o.NonZeroImpliesNotIdle, true
}

// HasNonZeroImpliesNotIdle returns a boolean if a field has been set.
func (o *AddCustomLoggedStatsRequest) HasNonZeroImpliesNotIdle() bool {
	if o != nil && !IsNil(o.NonZeroImpliesNotIdle) {
		return true
	}

	return false
}

// SetNonZeroImpliesNotIdle gets a reference to the given bool and assigns it to the NonZeroImpliesNotIdle field.
func (o *AddCustomLoggedStatsRequest) SetNonZeroImpliesNotIdle(v bool) {
	o.NonZeroImpliesNotIdle = &v
}

// GetStatsName returns the StatsName field value
func (o *AddCustomLoggedStatsRequest) GetStatsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatsName
}

// GetStatsNameOk returns a tuple with the StatsName field value
// and a boolean to check if the value has been set.
func (o *AddCustomLoggedStatsRequest) GetStatsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatsName, true
}

// SetStatsName sets field value
func (o *AddCustomLoggedStatsRequest) SetStatsName(v string) {
	o.StatsName = v
}

func (o AddCustomLoggedStatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCustomLoggedStatsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["monitorObjectclass"] = o.MonitorObjectclass
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	toSerialize["attributeToLog"] = o.AttributeToLog
	if !IsNil(o.ColumnName) {
		toSerialize["columnName"] = o.ColumnName
	}
	toSerialize["statisticType"] = o.StatisticType
	if !IsNil(o.HeaderPrefix) {
		toSerialize["headerPrefix"] = o.HeaderPrefix
	}
	if !IsNil(o.HeaderPrefixAttribute) {
		toSerialize["headerPrefixAttribute"] = o.HeaderPrefixAttribute
	}
	if !IsNil(o.RegexPattern) {
		toSerialize["regexPattern"] = o.RegexPattern
	}
	if !IsNil(o.RegexReplacement) {
		toSerialize["regexReplacement"] = o.RegexReplacement
	}
	if !IsNil(o.DivideValueBy) {
		toSerialize["divideValueBy"] = o.DivideValueBy
	}
	if !IsNil(o.DivideValueByAttribute) {
		toSerialize["divideValueByAttribute"] = o.DivideValueByAttribute
	}
	if !IsNil(o.DecimalFormat) {
		toSerialize["decimalFormat"] = o.DecimalFormat
	}
	if !IsNil(o.NonZeroImpliesNotIdle) {
		toSerialize["nonZeroImpliesNotIdle"] = o.NonZeroImpliesNotIdle
	}
	toSerialize["statsName"] = o.StatsName
	return toSerialize, nil
}

type NullableAddCustomLoggedStatsRequest struct {
	value *AddCustomLoggedStatsRequest
	isSet bool
}

func (v NullableAddCustomLoggedStatsRequest) Get() *AddCustomLoggedStatsRequest {
	return v.value
}

func (v *NullableAddCustomLoggedStatsRequest) Set(val *AddCustomLoggedStatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCustomLoggedStatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCustomLoggedStatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCustomLoggedStatsRequest(val *AddCustomLoggedStatsRequest) *NullableAddCustomLoggedStatsRequest {
	return &NullableAddCustomLoggedStatsRequest{value: val, isSet: true}
}

func (v NullableAddCustomLoggedStatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCustomLoggedStatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
