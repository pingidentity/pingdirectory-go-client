# ExpiredPasswordDataSecurityAuditorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnumexpiredPasswordDataSecurityAuditorSchemaUrn**](EnumexpiredPasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**PasswordEvaluationAge** | Pointer to **string** | If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled. | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewExpiredPasswordDataSecurityAuditorResponse

`func NewExpiredPasswordDataSecurityAuditorResponse(id string, schemas []EnumexpiredPasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *ExpiredPasswordDataSecurityAuditorResponse`

NewExpiredPasswordDataSecurityAuditorResponse instantiates a new ExpiredPasswordDataSecurityAuditorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiredPasswordDataSecurityAuditorResponseWithDefaults

`func NewExpiredPasswordDataSecurityAuditorResponseWithDefaults() *ExpiredPasswordDataSecurityAuditorResponse`

NewExpiredPasswordDataSecurityAuditorResponseWithDefaults instantiates a new ExpiredPasswordDataSecurityAuditorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetSchemas() []EnumexpiredPasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetSchemasOk() (*[]EnumexpiredPasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetSchemas(v []EnumexpiredPasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludeAttribute

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *ExpiredPasswordDataSecurityAuditorResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetPasswordEvaluationAge

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetPasswordEvaluationAge() string`

GetPasswordEvaluationAge returns the PasswordEvaluationAge field if non-nil, zero value otherwise.

### GetPasswordEvaluationAgeOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetPasswordEvaluationAgeOk() (*string, bool)`

GetPasswordEvaluationAgeOk returns a tuple with the PasswordEvaluationAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEvaluationAge

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetPasswordEvaluationAge(v string)`

SetPasswordEvaluationAge sets PasswordEvaluationAge field to given value.

### HasPasswordEvaluationAge

`func (o *ExpiredPasswordDataSecurityAuditorResponse) HasPasswordEvaluationAge() bool`

HasPasswordEvaluationAge returns a boolean if a field has been set.

### GetEnabled

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuditBackend

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *ExpiredPasswordDataSecurityAuditorResponse) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *ExpiredPasswordDataSecurityAuditorResponse) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetMeta

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExpiredPasswordDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExpiredPasswordDataSecurityAuditorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExpiredPasswordDataSecurityAuditorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


