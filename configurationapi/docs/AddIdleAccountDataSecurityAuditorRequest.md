# AddIdleAccountDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnumidleAccountDataSecurityAuditorSchemaUrn**](EnumidleAccountDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | Pointer to **string** | Specifies the name of the detailed report file. | [optional] 
**IdleAccountWarningInterval** | **string** | The length of time to use as the warning interval for idle accounts. If the length of time since a user last authenticated is greater than the warning interval but less than the error interval (or if it is greater than the warning interval and no error interval is defined), then a warning will be generated for that account. | 
**IdleAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for idle accounts. If the length of time since a user last authenticated is greater than the error interval, then an error will be generated for that account. If no error interval is defined, then only the warning interval will be used. | [optional] 
**NeverLoggedInAccountWarningInterval** | Pointer to **string** | The length of time to use as the warning interval for accounts that do not appear to have authenticated. If this is not specified, then the idle account warning interval will be used. | [optional] 
**NeverLoggedInAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for accounts that do not appear to have authenticated. If this is not specified, then the never-logged-in warning interval will be used. The idle account warning and error intervals will be used if no never-logged-in interval is configured. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAddIdleAccountDataSecurityAuditorRequest

`func NewAddIdleAccountDataSecurityAuditorRequest(auditorName string, schemas []EnumidleAccountDataSecurityAuditorSchemaUrn, idleAccountWarningInterval string, ) *AddIdleAccountDataSecurityAuditorRequest`

NewAddIdleAccountDataSecurityAuditorRequest instantiates a new AddIdleAccountDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdleAccountDataSecurityAuditorRequestWithDefaults

`func NewAddIdleAccountDataSecurityAuditorRequestWithDefaults() *AddIdleAccountDataSecurityAuditorRequest`

NewAddIdleAccountDataSecurityAuditorRequestWithDefaults instantiates a new AddIdleAccountDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetSchemas() []EnumidleAccountDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumidleAccountDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetSchemas(v []EnumidleAccountDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.

### HasReportFile

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasReportFile() bool`

HasReportFile returns a boolean if a field has been set.

### GetIdleAccountWarningInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountWarningInterval() string`

GetIdleAccountWarningInterval returns the IdleAccountWarningInterval field if non-nil, zero value otherwise.

### GetIdleAccountWarningIntervalOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountWarningIntervalOk() (*string, bool)`

GetIdleAccountWarningIntervalOk returns a tuple with the IdleAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountWarningInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetIdleAccountWarningInterval(v string)`

SetIdleAccountWarningInterval sets IdleAccountWarningInterval field to given value.


### GetIdleAccountErrorInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountErrorInterval() string`

GetIdleAccountErrorInterval returns the IdleAccountErrorInterval field if non-nil, zero value otherwise.

### GetIdleAccountErrorIntervalOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountErrorIntervalOk() (*string, bool)`

GetIdleAccountErrorIntervalOk returns a tuple with the IdleAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountErrorInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetIdleAccountErrorInterval(v string)`

SetIdleAccountErrorInterval sets IdleAccountErrorInterval field to given value.

### HasIdleAccountErrorInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasIdleAccountErrorInterval() bool`

HasIdleAccountErrorInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountWarningInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountWarningInterval() string`

GetNeverLoggedInAccountWarningInterval returns the NeverLoggedInAccountWarningInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountWarningIntervalOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountWarningIntervalOk() (*string, bool)`

GetNeverLoggedInAccountWarningIntervalOk returns a tuple with the NeverLoggedInAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountWarningInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetNeverLoggedInAccountWarningInterval(v string)`

SetNeverLoggedInAccountWarningInterval sets NeverLoggedInAccountWarningInterval field to given value.

### HasNeverLoggedInAccountWarningInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasNeverLoggedInAccountWarningInterval() bool`

HasNeverLoggedInAccountWarningInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountErrorInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountErrorInterval() string`

GetNeverLoggedInAccountErrorInterval returns the NeverLoggedInAccountErrorInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountErrorIntervalOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountErrorIntervalOk() (*string, bool)`

GetNeverLoggedInAccountErrorIntervalOk returns a tuple with the NeverLoggedInAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountErrorInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetNeverLoggedInAccountErrorInterval(v string)`

SetNeverLoggedInAccountErrorInterval sets NeverLoggedInAccountErrorInterval field to given value.

### HasNeverLoggedInAccountErrorInterval

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasNeverLoggedInAccountErrorInterval() bool`

HasNeverLoggedInAccountErrorInterval returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIncludeAttribute

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddIdleAccountDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddIdleAccountDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


