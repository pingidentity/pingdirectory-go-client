# IdleAccountDataSecurityAuditorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnumidleAccountDataSecurityAuditorSchemaUrn**](EnumidleAccountDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IdleAccountWarningInterval** | **string** | The length of time to use as the warning interval for idle accounts. If the length of time since a user last authenticated is greater than the warning interval but less than the error interval (or if it is greater than the warning interval and no error interval is defined), then a warning will be generated for that account. | 
**IdleAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for idle accounts. If the length of time since a user last authenticated is greater than the error interval, then an error will be generated for that account. If no error interval is defined, then only the warning interval will be used. | [optional] 
**NeverLoggedInAccountWarningInterval** | Pointer to **string** | The length of time to use as the warning interval for accounts that do not appear to have authenticated. If this is not specified, then the idle account warning interval will be used. | [optional] 
**NeverLoggedInAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for accounts that do not appear to have authenticated. If this is not specified, then the never-logged-in warning interval will be used. The idle account warning and error intervals will be used if no never-logged-in interval is configured. | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewIdleAccountDataSecurityAuditorResponse

`func NewIdleAccountDataSecurityAuditorResponse(id string, schemas []EnumidleAccountDataSecurityAuditorSchemaUrn, reportFile string, idleAccountWarningInterval string, enabled bool, ) *IdleAccountDataSecurityAuditorResponse`

NewIdleAccountDataSecurityAuditorResponse instantiates a new IdleAccountDataSecurityAuditorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdleAccountDataSecurityAuditorResponseWithDefaults

`func NewIdleAccountDataSecurityAuditorResponseWithDefaults() *IdleAccountDataSecurityAuditorResponse`

NewIdleAccountDataSecurityAuditorResponseWithDefaults instantiates a new IdleAccountDataSecurityAuditorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *IdleAccountDataSecurityAuditorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IdleAccountDataSecurityAuditorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IdleAccountDataSecurityAuditorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *IdleAccountDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *IdleAccountDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *IdleAccountDataSecurityAuditorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *IdleAccountDataSecurityAuditorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *IdleAccountDataSecurityAuditorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdleAccountDataSecurityAuditorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *IdleAccountDataSecurityAuditorResponse) GetSchemas() []EnumidleAccountDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetSchemasOk() (*[]EnumidleAccountDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IdleAccountDataSecurityAuditorResponse) SetSchemas(v []EnumidleAccountDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *IdleAccountDataSecurityAuditorResponse) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *IdleAccountDataSecurityAuditorResponse) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIdleAccountWarningInterval

`func (o *IdleAccountDataSecurityAuditorResponse) GetIdleAccountWarningInterval() string`

GetIdleAccountWarningInterval returns the IdleAccountWarningInterval field if non-nil, zero value otherwise.

### GetIdleAccountWarningIntervalOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetIdleAccountWarningIntervalOk() (*string, bool)`

GetIdleAccountWarningIntervalOk returns a tuple with the IdleAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountWarningInterval

`func (o *IdleAccountDataSecurityAuditorResponse) SetIdleAccountWarningInterval(v string)`

SetIdleAccountWarningInterval sets IdleAccountWarningInterval field to given value.


### GetIdleAccountErrorInterval

`func (o *IdleAccountDataSecurityAuditorResponse) GetIdleAccountErrorInterval() string`

GetIdleAccountErrorInterval returns the IdleAccountErrorInterval field if non-nil, zero value otherwise.

### GetIdleAccountErrorIntervalOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetIdleAccountErrorIntervalOk() (*string, bool)`

GetIdleAccountErrorIntervalOk returns a tuple with the IdleAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountErrorInterval

`func (o *IdleAccountDataSecurityAuditorResponse) SetIdleAccountErrorInterval(v string)`

SetIdleAccountErrorInterval sets IdleAccountErrorInterval field to given value.

### HasIdleAccountErrorInterval

`func (o *IdleAccountDataSecurityAuditorResponse) HasIdleAccountErrorInterval() bool`

HasIdleAccountErrorInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountWarningInterval

`func (o *IdleAccountDataSecurityAuditorResponse) GetNeverLoggedInAccountWarningInterval() string`

GetNeverLoggedInAccountWarningInterval returns the NeverLoggedInAccountWarningInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountWarningIntervalOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetNeverLoggedInAccountWarningIntervalOk() (*string, bool)`

GetNeverLoggedInAccountWarningIntervalOk returns a tuple with the NeverLoggedInAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountWarningInterval

`func (o *IdleAccountDataSecurityAuditorResponse) SetNeverLoggedInAccountWarningInterval(v string)`

SetNeverLoggedInAccountWarningInterval sets NeverLoggedInAccountWarningInterval field to given value.

### HasNeverLoggedInAccountWarningInterval

`func (o *IdleAccountDataSecurityAuditorResponse) HasNeverLoggedInAccountWarningInterval() bool`

HasNeverLoggedInAccountWarningInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountErrorInterval

`func (o *IdleAccountDataSecurityAuditorResponse) GetNeverLoggedInAccountErrorInterval() string`

GetNeverLoggedInAccountErrorInterval returns the NeverLoggedInAccountErrorInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountErrorIntervalOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetNeverLoggedInAccountErrorIntervalOk() (*string, bool)`

GetNeverLoggedInAccountErrorIntervalOk returns a tuple with the NeverLoggedInAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountErrorInterval

`func (o *IdleAccountDataSecurityAuditorResponse) SetNeverLoggedInAccountErrorInterval(v string)`

SetNeverLoggedInAccountErrorInterval sets NeverLoggedInAccountErrorInterval field to given value.

### HasNeverLoggedInAccountErrorInterval

`func (o *IdleAccountDataSecurityAuditorResponse) HasNeverLoggedInAccountErrorInterval() bool`

HasNeverLoggedInAccountErrorInterval returns a boolean if a field has been set.

### GetEnabled

`func (o *IdleAccountDataSecurityAuditorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdleAccountDataSecurityAuditorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *IdleAccountDataSecurityAuditorResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *IdleAccountDataSecurityAuditorResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *IdleAccountDataSecurityAuditorResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *IdleAccountDataSecurityAuditorResponse) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *IdleAccountDataSecurityAuditorResponse) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *IdleAccountDataSecurityAuditorResponse) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *IdleAccountDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *IdleAccountDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *IdleAccountDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *IdleAccountDataSecurityAuditorResponse) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


