# ExpiredPasswordDataSecurityAuditorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumexpiredPasswordDataSecurityAuditorSchemaUrn**](EnumexpiredPasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**PasswordEvaluationAge** | Pointer to **string** | If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled. | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewExpiredPasswordDataSecurityAuditorShared

`func NewExpiredPasswordDataSecurityAuditorShared(schemas []EnumexpiredPasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *ExpiredPasswordDataSecurityAuditorShared`

NewExpiredPasswordDataSecurityAuditorShared instantiates a new ExpiredPasswordDataSecurityAuditorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiredPasswordDataSecurityAuditorSharedWithDefaults

`func NewExpiredPasswordDataSecurityAuditorSharedWithDefaults() *ExpiredPasswordDataSecurityAuditorShared`

NewExpiredPasswordDataSecurityAuditorSharedWithDefaults instantiates a new ExpiredPasswordDataSecurityAuditorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetSchemas() []EnumexpiredPasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetSchemasOk() (*[]EnumexpiredPasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetSchemas(v []EnumexpiredPasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludeAttribute

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *ExpiredPasswordDataSecurityAuditorShared) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetPasswordEvaluationAge

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetPasswordEvaluationAge() string`

GetPasswordEvaluationAge returns the PasswordEvaluationAge field if non-nil, zero value otherwise.

### GetPasswordEvaluationAgeOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetPasswordEvaluationAgeOk() (*string, bool)`

GetPasswordEvaluationAgeOk returns a tuple with the PasswordEvaluationAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEvaluationAge

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetPasswordEvaluationAge(v string)`

SetPasswordEvaluationAge sets PasswordEvaluationAge field to given value.

### HasPasswordEvaluationAge

`func (o *ExpiredPasswordDataSecurityAuditorShared) HasPasswordEvaluationAge() bool`

HasPasswordEvaluationAge returns a boolean if a field has been set.

### GetEnabled

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuditBackend

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *ExpiredPasswordDataSecurityAuditorShared) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *ExpiredPasswordDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *ExpiredPasswordDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *ExpiredPasswordDataSecurityAuditorShared) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


