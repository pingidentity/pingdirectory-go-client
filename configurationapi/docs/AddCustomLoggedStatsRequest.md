# AddCustomLoggedStatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcustomLoggedStatsSchemaUrn**](EnumcustomLoggedStatsSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Custom Logged Stats | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Custom Logged Stats object is enabled. | [optional] 
**MonitorObjectclass** | **string** | The objectclass name of the monitor entries to examine for generating these statistics. | 
**IncludeFilter** | Pointer to **string** | An optional LDAP filter that can be used restrict which monitor entries are used to produce the output. | [optional] 
**AttributeToLog** | **[]string** | Specifies the attributes on the monitor entries that should be included in the output. | 
**ColumnName** | Pointer to **[]string** | Optionally, specifies an explicit name for each column header instead of having these names automatically generated from the monitored attribute name. | [optional] 
**StatisticType** | [**[]EnumcustomLoggedStatsStatisticTypeProp**](EnumcustomLoggedStatsStatisticTypeProp.md) |  | 
**HeaderPrefix** | Pointer to **string** | An optional prefix that is included in the header before the column name. | [optional] 
**HeaderPrefixAttribute** | Pointer to **string** | An optional attribute from the monitor entry that is included as a prefix before the column name in the column header. | [optional] 
**RegexPattern** | Pointer to **string** | An optional regular expression pattern, that when used in conjunction with regex-replacement, can alter the value of the attribute being monitored. | [optional] 
**RegexReplacement** | Pointer to **string** | An optional regular expression replacement value, that when used in conjunction with regex-pattern, can alter the value of the attribute being monitored. | [optional] 
**DivideValueBy** | Pointer to **string** | An optional floating point value that can be used to scale the resulting value. | [optional] 
**DivideValueByAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute in the monitored entry. | [optional] 
**DecimalFormat** | Pointer to **string** | This provides a way to format the monitored attribute value in the output to control the precision for instance. | [optional] 
**NonZeroImpliesNotIdle** | Pointer to **bool** | If this property is set to true, then the value of any of the monitored attributes here can contribute to whether an interval is considered \&quot;idle\&quot; by the Periodic Stats Logger. | [optional] 
**StatsName** | **string** | Name of the new Custom Logged Stats | 

## Methods

### NewAddCustomLoggedStatsRequest

`func NewAddCustomLoggedStatsRequest(schemas []EnumcustomLoggedStatsSchemaUrn, monitorObjectclass string, attributeToLog []string, statisticType []EnumcustomLoggedStatsStatisticTypeProp, statsName string, ) *AddCustomLoggedStatsRequest`

NewAddCustomLoggedStatsRequest instantiates a new AddCustomLoggedStatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomLoggedStatsRequestWithDefaults

`func NewAddCustomLoggedStatsRequestWithDefaults() *AddCustomLoggedStatsRequest`

NewAddCustomLoggedStatsRequestWithDefaults instantiates a new AddCustomLoggedStatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddCustomLoggedStatsRequest) GetSchemas() []EnumcustomLoggedStatsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCustomLoggedStatsRequest) GetSchemasOk() (*[]EnumcustomLoggedStatsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCustomLoggedStatsRequest) SetSchemas(v []EnumcustomLoggedStatsSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddCustomLoggedStatsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCustomLoggedStatsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCustomLoggedStatsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCustomLoggedStatsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCustomLoggedStatsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCustomLoggedStatsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCustomLoggedStatsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCustomLoggedStatsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *AddCustomLoggedStatsRequest) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *AddCustomLoggedStatsRequest) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *AddCustomLoggedStatsRequest) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetIncludeFilter

`func (o *AddCustomLoggedStatsRequest) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddCustomLoggedStatsRequest) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddCustomLoggedStatsRequest) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddCustomLoggedStatsRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetAttributeToLog

`func (o *AddCustomLoggedStatsRequest) GetAttributeToLog() []string`

GetAttributeToLog returns the AttributeToLog field if non-nil, zero value otherwise.

### GetAttributeToLogOk

`func (o *AddCustomLoggedStatsRequest) GetAttributeToLogOk() (*[]string, bool)`

GetAttributeToLogOk returns a tuple with the AttributeToLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeToLog

`func (o *AddCustomLoggedStatsRequest) SetAttributeToLog(v []string)`

SetAttributeToLog sets AttributeToLog field to given value.


### GetColumnName

`func (o *AddCustomLoggedStatsRequest) GetColumnName() []string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *AddCustomLoggedStatsRequest) GetColumnNameOk() (*[]string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *AddCustomLoggedStatsRequest) SetColumnName(v []string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *AddCustomLoggedStatsRequest) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetStatisticType

`func (o *AddCustomLoggedStatsRequest) GetStatisticType() []EnumcustomLoggedStatsStatisticTypeProp`

GetStatisticType returns the StatisticType field if non-nil, zero value otherwise.

### GetStatisticTypeOk

`func (o *AddCustomLoggedStatsRequest) GetStatisticTypeOk() (*[]EnumcustomLoggedStatsStatisticTypeProp, bool)`

GetStatisticTypeOk returns a tuple with the StatisticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticType

`func (o *AddCustomLoggedStatsRequest) SetStatisticType(v []EnumcustomLoggedStatsStatisticTypeProp)`

SetStatisticType sets StatisticType field to given value.


### GetHeaderPrefix

`func (o *AddCustomLoggedStatsRequest) GetHeaderPrefix() string`

GetHeaderPrefix returns the HeaderPrefix field if non-nil, zero value otherwise.

### GetHeaderPrefixOk

`func (o *AddCustomLoggedStatsRequest) GetHeaderPrefixOk() (*string, bool)`

GetHeaderPrefixOk returns a tuple with the HeaderPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefix

`func (o *AddCustomLoggedStatsRequest) SetHeaderPrefix(v string)`

SetHeaderPrefix sets HeaderPrefix field to given value.

### HasHeaderPrefix

`func (o *AddCustomLoggedStatsRequest) HasHeaderPrefix() bool`

HasHeaderPrefix returns a boolean if a field has been set.

### GetHeaderPrefixAttribute

`func (o *AddCustomLoggedStatsRequest) GetHeaderPrefixAttribute() string`

GetHeaderPrefixAttribute returns the HeaderPrefixAttribute field if non-nil, zero value otherwise.

### GetHeaderPrefixAttributeOk

`func (o *AddCustomLoggedStatsRequest) GetHeaderPrefixAttributeOk() (*string, bool)`

GetHeaderPrefixAttributeOk returns a tuple with the HeaderPrefixAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefixAttribute

`func (o *AddCustomLoggedStatsRequest) SetHeaderPrefixAttribute(v string)`

SetHeaderPrefixAttribute sets HeaderPrefixAttribute field to given value.

### HasHeaderPrefixAttribute

`func (o *AddCustomLoggedStatsRequest) HasHeaderPrefixAttribute() bool`

HasHeaderPrefixAttribute returns a boolean if a field has been set.

### GetRegexPattern

`func (o *AddCustomLoggedStatsRequest) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *AddCustomLoggedStatsRequest) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *AddCustomLoggedStatsRequest) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *AddCustomLoggedStatsRequest) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### GetRegexReplacement

`func (o *AddCustomLoggedStatsRequest) GetRegexReplacement() string`

GetRegexReplacement returns the RegexReplacement field if non-nil, zero value otherwise.

### GetRegexReplacementOk

`func (o *AddCustomLoggedStatsRequest) GetRegexReplacementOk() (*string, bool)`

GetRegexReplacementOk returns a tuple with the RegexReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexReplacement

`func (o *AddCustomLoggedStatsRequest) SetRegexReplacement(v string)`

SetRegexReplacement sets RegexReplacement field to given value.

### HasRegexReplacement

`func (o *AddCustomLoggedStatsRequest) HasRegexReplacement() bool`

HasRegexReplacement returns a boolean if a field has been set.

### GetDivideValueBy

`func (o *AddCustomLoggedStatsRequest) GetDivideValueBy() string`

GetDivideValueBy returns the DivideValueBy field if non-nil, zero value otherwise.

### GetDivideValueByOk

`func (o *AddCustomLoggedStatsRequest) GetDivideValueByOk() (*string, bool)`

GetDivideValueByOk returns a tuple with the DivideValueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueBy

`func (o *AddCustomLoggedStatsRequest) SetDivideValueBy(v string)`

SetDivideValueBy sets DivideValueBy field to given value.

### HasDivideValueBy

`func (o *AddCustomLoggedStatsRequest) HasDivideValueBy() bool`

HasDivideValueBy returns a boolean if a field has been set.

### GetDivideValueByAttribute

`func (o *AddCustomLoggedStatsRequest) GetDivideValueByAttribute() string`

GetDivideValueByAttribute returns the DivideValueByAttribute field if non-nil, zero value otherwise.

### GetDivideValueByAttributeOk

`func (o *AddCustomLoggedStatsRequest) GetDivideValueByAttributeOk() (*string, bool)`

GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByAttribute

`func (o *AddCustomLoggedStatsRequest) SetDivideValueByAttribute(v string)`

SetDivideValueByAttribute sets DivideValueByAttribute field to given value.

### HasDivideValueByAttribute

`func (o *AddCustomLoggedStatsRequest) HasDivideValueByAttribute() bool`

HasDivideValueByAttribute returns a boolean if a field has been set.

### GetDecimalFormat

`func (o *AddCustomLoggedStatsRequest) GetDecimalFormat() string`

GetDecimalFormat returns the DecimalFormat field if non-nil, zero value otherwise.

### GetDecimalFormatOk

`func (o *AddCustomLoggedStatsRequest) GetDecimalFormatOk() (*string, bool)`

GetDecimalFormatOk returns a tuple with the DecimalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalFormat

`func (o *AddCustomLoggedStatsRequest) SetDecimalFormat(v string)`

SetDecimalFormat sets DecimalFormat field to given value.

### HasDecimalFormat

`func (o *AddCustomLoggedStatsRequest) HasDecimalFormat() bool`

HasDecimalFormat returns a boolean if a field has been set.

### GetNonZeroImpliesNotIdle

`func (o *AddCustomLoggedStatsRequest) GetNonZeroImpliesNotIdle() bool`

GetNonZeroImpliesNotIdle returns the NonZeroImpliesNotIdle field if non-nil, zero value otherwise.

### GetNonZeroImpliesNotIdleOk

`func (o *AddCustomLoggedStatsRequest) GetNonZeroImpliesNotIdleOk() (*bool, bool)`

GetNonZeroImpliesNotIdleOk returns a tuple with the NonZeroImpliesNotIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonZeroImpliesNotIdle

`func (o *AddCustomLoggedStatsRequest) SetNonZeroImpliesNotIdle(v bool)`

SetNonZeroImpliesNotIdle sets NonZeroImpliesNotIdle field to given value.

### HasNonZeroImpliesNotIdle

`func (o *AddCustomLoggedStatsRequest) HasNonZeroImpliesNotIdle() bool`

HasNonZeroImpliesNotIdle returns a boolean if a field has been set.

### GetStatsName

`func (o *AddCustomLoggedStatsRequest) GetStatsName() string`

GetStatsName returns the StatsName field if non-nil, zero value otherwise.

### GetStatsNameOk

`func (o *AddCustomLoggedStatsRequest) GetStatsNameOk() (*string, bool)`

GetStatsNameOk returns a tuple with the StatsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsName

`func (o *AddCustomLoggedStatsRequest) SetStatsName(v string)`

SetStatsName sets StatsName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


