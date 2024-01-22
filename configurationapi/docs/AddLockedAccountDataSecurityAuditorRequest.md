# AddLockedAccountDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlockedAccountDataSecurityAuditorSchemaUrn**](EnumlockedAccountDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | Pointer to **string** | Specifies the name of the detailed report file. | [optional] 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**MaximumIdleTime** | Pointer to **string** | If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**AuditorName** | **string** | Name of the new Data Security Auditor | 

## Methods

### NewAddLockedAccountDataSecurityAuditorRequest

`func NewAddLockedAccountDataSecurityAuditorRequest(schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn, auditorName string, ) *AddLockedAccountDataSecurityAuditorRequest`

NewAddLockedAccountDataSecurityAuditorRequest instantiates a new AddLockedAccountDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLockedAccountDataSecurityAuditorRequestWithDefaults

`func NewAddLockedAccountDataSecurityAuditorRequestWithDefaults() *AddLockedAccountDataSecurityAuditorRequest`

NewAddLockedAccountDataSecurityAuditorRequestWithDefaults instantiates a new AddLockedAccountDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetSchemas() []EnumlockedAccountDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumlockedAccountDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetSchemas(v []EnumlockedAccountDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.

### HasReportFile

`func (o *AddLockedAccountDataSecurityAuditorRequest) HasReportFile() bool`

HasReportFile returns a boolean if a field has been set.

### GetIncludeAttribute

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddLockedAccountDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetMaximumIdleTime

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetMaximumIdleTime() string`

GetMaximumIdleTime returns the MaximumIdleTime field if non-nil, zero value otherwise.

### GetMaximumIdleTimeOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetMaximumIdleTimeOk() (*string, bool)`

GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIdleTime

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetMaximumIdleTime(v string)`

SetMaximumIdleTime sets MaximumIdleTime field to given value.

### HasMaximumIdleTime

`func (o *AddLockedAccountDataSecurityAuditorRequest) HasMaximumIdleTime() bool`

HasMaximumIdleTime returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddLockedAccountDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddLockedAccountDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddLockedAccountDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetAuditorName

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddLockedAccountDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


