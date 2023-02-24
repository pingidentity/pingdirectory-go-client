# AddExpiredPasswordDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnumexpiredPasswordDataSecurityAuditorSchemaUrn**](EnumexpiredPasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | Pointer to **string** | Specifies the name of the detailed report file. | [optional] 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**PasswordEvaluationAge** | Pointer to **string** | If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAddExpiredPasswordDataSecurityAuditorRequest

`func NewAddExpiredPasswordDataSecurityAuditorRequest(auditorName string, schemas []EnumexpiredPasswordDataSecurityAuditorSchemaUrn, ) *AddExpiredPasswordDataSecurityAuditorRequest`

NewAddExpiredPasswordDataSecurityAuditorRequest instantiates a new AddExpiredPasswordDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExpiredPasswordDataSecurityAuditorRequestWithDefaults

`func NewAddExpiredPasswordDataSecurityAuditorRequestWithDefaults() *AddExpiredPasswordDataSecurityAuditorRequest`

NewAddExpiredPasswordDataSecurityAuditorRequestWithDefaults instantiates a new AddExpiredPasswordDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetSchemas() []EnumexpiredPasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumexpiredPasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetSchemas(v []EnumexpiredPasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.

### HasReportFile

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) HasReportFile() bool`

HasReportFile returns a boolean if a field has been set.

### GetIncludeAttribute

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetPasswordEvaluationAge

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetPasswordEvaluationAge() string`

GetPasswordEvaluationAge returns the PasswordEvaluationAge field if non-nil, zero value otherwise.

### GetPasswordEvaluationAgeOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetPasswordEvaluationAgeOk() (*string, bool)`

GetPasswordEvaluationAgeOk returns a tuple with the PasswordEvaluationAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEvaluationAge

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetPasswordEvaluationAge(v string)`

SetPasswordEvaluationAge sets PasswordEvaluationAge field to given value.

### HasPasswordEvaluationAge

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) HasPasswordEvaluationAge() bool`

HasPasswordEvaluationAge returns a boolean if a field has been set.

### GetEnabled

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddExpiredPasswordDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


