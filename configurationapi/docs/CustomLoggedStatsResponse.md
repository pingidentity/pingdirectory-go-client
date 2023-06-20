# CustomLoggedStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Custom Logged Stats | 
**Schemas** | [**[]EnumcustomLoggedStatsSchemaUrn**](EnumcustomLoggedStatsSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Custom Logged Stats | [optional] 
**Enabled** | **bool** | Indicates whether the Custom Logged Stats object is enabled. | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewCustomLoggedStatsResponse

`func NewCustomLoggedStatsResponse(id string, schemas []EnumcustomLoggedStatsSchemaUrn, enabled bool, monitorObjectclass string, attributeToLog []string, statisticType []EnumcustomLoggedStatsStatisticTypeProp, ) *CustomLoggedStatsResponse`

NewCustomLoggedStatsResponse instantiates a new CustomLoggedStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLoggedStatsResponseWithDefaults

`func NewCustomLoggedStatsResponseWithDefaults() *CustomLoggedStatsResponse`

NewCustomLoggedStatsResponseWithDefaults instantiates a new CustomLoggedStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomLoggedStatsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomLoggedStatsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomLoggedStatsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CustomLoggedStatsResponse) GetSchemas() []EnumcustomLoggedStatsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CustomLoggedStatsResponse) GetSchemasOk() (*[]EnumcustomLoggedStatsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CustomLoggedStatsResponse) SetSchemas(v []EnumcustomLoggedStatsSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *CustomLoggedStatsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomLoggedStatsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomLoggedStatsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomLoggedStatsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomLoggedStatsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomLoggedStatsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomLoggedStatsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMonitorObjectclass

`func (o *CustomLoggedStatsResponse) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *CustomLoggedStatsResponse) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *CustomLoggedStatsResponse) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetIncludeFilter

`func (o *CustomLoggedStatsResponse) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *CustomLoggedStatsResponse) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *CustomLoggedStatsResponse) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *CustomLoggedStatsResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetAttributeToLog

`func (o *CustomLoggedStatsResponse) GetAttributeToLog() []string`

GetAttributeToLog returns the AttributeToLog field if non-nil, zero value otherwise.

### GetAttributeToLogOk

`func (o *CustomLoggedStatsResponse) GetAttributeToLogOk() (*[]string, bool)`

GetAttributeToLogOk returns a tuple with the AttributeToLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeToLog

`func (o *CustomLoggedStatsResponse) SetAttributeToLog(v []string)`

SetAttributeToLog sets AttributeToLog field to given value.


### GetColumnName

`func (o *CustomLoggedStatsResponse) GetColumnName() []string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *CustomLoggedStatsResponse) GetColumnNameOk() (*[]string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *CustomLoggedStatsResponse) SetColumnName(v []string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *CustomLoggedStatsResponse) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetStatisticType

`func (o *CustomLoggedStatsResponse) GetStatisticType() []EnumcustomLoggedStatsStatisticTypeProp`

GetStatisticType returns the StatisticType field if non-nil, zero value otherwise.

### GetStatisticTypeOk

`func (o *CustomLoggedStatsResponse) GetStatisticTypeOk() (*[]EnumcustomLoggedStatsStatisticTypeProp, bool)`

GetStatisticTypeOk returns a tuple with the StatisticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticType

`func (o *CustomLoggedStatsResponse) SetStatisticType(v []EnumcustomLoggedStatsStatisticTypeProp)`

SetStatisticType sets StatisticType field to given value.


### GetHeaderPrefix

`func (o *CustomLoggedStatsResponse) GetHeaderPrefix() string`

GetHeaderPrefix returns the HeaderPrefix field if non-nil, zero value otherwise.

### GetHeaderPrefixOk

`func (o *CustomLoggedStatsResponse) GetHeaderPrefixOk() (*string, bool)`

GetHeaderPrefixOk returns a tuple with the HeaderPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefix

`func (o *CustomLoggedStatsResponse) SetHeaderPrefix(v string)`

SetHeaderPrefix sets HeaderPrefix field to given value.

### HasHeaderPrefix

`func (o *CustomLoggedStatsResponse) HasHeaderPrefix() bool`

HasHeaderPrefix returns a boolean if a field has been set.

### GetHeaderPrefixAttribute

`func (o *CustomLoggedStatsResponse) GetHeaderPrefixAttribute() string`

GetHeaderPrefixAttribute returns the HeaderPrefixAttribute field if non-nil, zero value otherwise.

### GetHeaderPrefixAttributeOk

`func (o *CustomLoggedStatsResponse) GetHeaderPrefixAttributeOk() (*string, bool)`

GetHeaderPrefixAttributeOk returns a tuple with the HeaderPrefixAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefixAttribute

`func (o *CustomLoggedStatsResponse) SetHeaderPrefixAttribute(v string)`

SetHeaderPrefixAttribute sets HeaderPrefixAttribute field to given value.

### HasHeaderPrefixAttribute

`func (o *CustomLoggedStatsResponse) HasHeaderPrefixAttribute() bool`

HasHeaderPrefixAttribute returns a boolean if a field has been set.

### GetRegexPattern

`func (o *CustomLoggedStatsResponse) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *CustomLoggedStatsResponse) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *CustomLoggedStatsResponse) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *CustomLoggedStatsResponse) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### GetRegexReplacement

`func (o *CustomLoggedStatsResponse) GetRegexReplacement() string`

GetRegexReplacement returns the RegexReplacement field if non-nil, zero value otherwise.

### GetRegexReplacementOk

`func (o *CustomLoggedStatsResponse) GetRegexReplacementOk() (*string, bool)`

GetRegexReplacementOk returns a tuple with the RegexReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexReplacement

`func (o *CustomLoggedStatsResponse) SetRegexReplacement(v string)`

SetRegexReplacement sets RegexReplacement field to given value.

### HasRegexReplacement

`func (o *CustomLoggedStatsResponse) HasRegexReplacement() bool`

HasRegexReplacement returns a boolean if a field has been set.

### GetDivideValueBy

`func (o *CustomLoggedStatsResponse) GetDivideValueBy() string`

GetDivideValueBy returns the DivideValueBy field if non-nil, zero value otherwise.

### GetDivideValueByOk

`func (o *CustomLoggedStatsResponse) GetDivideValueByOk() (*string, bool)`

GetDivideValueByOk returns a tuple with the DivideValueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueBy

`func (o *CustomLoggedStatsResponse) SetDivideValueBy(v string)`

SetDivideValueBy sets DivideValueBy field to given value.

### HasDivideValueBy

`func (o *CustomLoggedStatsResponse) HasDivideValueBy() bool`

HasDivideValueBy returns a boolean if a field has been set.

### GetDivideValueByAttribute

`func (o *CustomLoggedStatsResponse) GetDivideValueByAttribute() string`

GetDivideValueByAttribute returns the DivideValueByAttribute field if non-nil, zero value otherwise.

### GetDivideValueByAttributeOk

`func (o *CustomLoggedStatsResponse) GetDivideValueByAttributeOk() (*string, bool)`

GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByAttribute

`func (o *CustomLoggedStatsResponse) SetDivideValueByAttribute(v string)`

SetDivideValueByAttribute sets DivideValueByAttribute field to given value.

### HasDivideValueByAttribute

`func (o *CustomLoggedStatsResponse) HasDivideValueByAttribute() bool`

HasDivideValueByAttribute returns a boolean if a field has been set.

### GetDecimalFormat

`func (o *CustomLoggedStatsResponse) GetDecimalFormat() string`

GetDecimalFormat returns the DecimalFormat field if non-nil, zero value otherwise.

### GetDecimalFormatOk

`func (o *CustomLoggedStatsResponse) GetDecimalFormatOk() (*string, bool)`

GetDecimalFormatOk returns a tuple with the DecimalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalFormat

`func (o *CustomLoggedStatsResponse) SetDecimalFormat(v string)`

SetDecimalFormat sets DecimalFormat field to given value.

### HasDecimalFormat

`func (o *CustomLoggedStatsResponse) HasDecimalFormat() bool`

HasDecimalFormat returns a boolean if a field has been set.

### GetNonZeroImpliesNotIdle

`func (o *CustomLoggedStatsResponse) GetNonZeroImpliesNotIdle() bool`

GetNonZeroImpliesNotIdle returns the NonZeroImpliesNotIdle field if non-nil, zero value otherwise.

### GetNonZeroImpliesNotIdleOk

`func (o *CustomLoggedStatsResponse) GetNonZeroImpliesNotIdleOk() (*bool, bool)`

GetNonZeroImpliesNotIdleOk returns a tuple with the NonZeroImpliesNotIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonZeroImpliesNotIdle

`func (o *CustomLoggedStatsResponse) SetNonZeroImpliesNotIdle(v bool)`

SetNonZeroImpliesNotIdle sets NonZeroImpliesNotIdle field to given value.

### HasNonZeroImpliesNotIdle

`func (o *CustomLoggedStatsResponse) HasNonZeroImpliesNotIdle() bool`

HasNonZeroImpliesNotIdle returns a boolean if a field has been set.

### GetMeta

`func (o *CustomLoggedStatsResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomLoggedStatsResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomLoggedStatsResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CustomLoggedStatsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CustomLoggedStatsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CustomLoggedStatsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CustomLoggedStatsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CustomLoggedStatsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


