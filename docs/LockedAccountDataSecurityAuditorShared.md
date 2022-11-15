# LockedAccountDataSecurityAuditorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlockedAccountDataSecurityAuditorSchemaUrn**](EnumlockedAccountDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**MaximumIdleTime** | Pointer to **string** | If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled. | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewLockedAccountDataSecurityAuditorShared

`func NewLockedAccountDataSecurityAuditorShared(schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *LockedAccountDataSecurityAuditorShared`

NewLockedAccountDataSecurityAuditorShared instantiates a new LockedAccountDataSecurityAuditorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockedAccountDataSecurityAuditorSharedWithDefaults

`func NewLockedAccountDataSecurityAuditorSharedWithDefaults() *LockedAccountDataSecurityAuditorShared`

NewLockedAccountDataSecurityAuditorSharedWithDefaults instantiates a new LockedAccountDataSecurityAuditorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LockedAccountDataSecurityAuditorShared) GetSchemas() []EnumlockedAccountDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LockedAccountDataSecurityAuditorShared) GetSchemasOk() (*[]EnumlockedAccountDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LockedAccountDataSecurityAuditorShared) SetSchemas(v []EnumlockedAccountDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *LockedAccountDataSecurityAuditorShared) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *LockedAccountDataSecurityAuditorShared) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *LockedAccountDataSecurityAuditorShared) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludeAttribute

`func (o *LockedAccountDataSecurityAuditorShared) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *LockedAccountDataSecurityAuditorShared) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *LockedAccountDataSecurityAuditorShared) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *LockedAccountDataSecurityAuditorShared) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetMaximumIdleTime

`func (o *LockedAccountDataSecurityAuditorShared) GetMaximumIdleTime() string`

GetMaximumIdleTime returns the MaximumIdleTime field if non-nil, zero value otherwise.

### GetMaximumIdleTimeOk

`func (o *LockedAccountDataSecurityAuditorShared) GetMaximumIdleTimeOk() (*string, bool)`

GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIdleTime

`func (o *LockedAccountDataSecurityAuditorShared) SetMaximumIdleTime(v string)`

SetMaximumIdleTime sets MaximumIdleTime field to given value.

### HasMaximumIdleTime

`func (o *LockedAccountDataSecurityAuditorShared) HasMaximumIdleTime() bool`

HasMaximumIdleTime returns a boolean if a field has been set.

### GetEnabled

`func (o *LockedAccountDataSecurityAuditorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LockedAccountDataSecurityAuditorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LockedAccountDataSecurityAuditorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuditBackend

`func (o *LockedAccountDataSecurityAuditorShared) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *LockedAccountDataSecurityAuditorShared) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *LockedAccountDataSecurityAuditorShared) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *LockedAccountDataSecurityAuditorShared) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *LockedAccountDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *LockedAccountDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *LockedAccountDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *LockedAccountDataSecurityAuditorShared) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


