# AddAccountValidityWindowDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnumaccountValidityWindowDataSecurityAuditorSchemaUrn**](EnumaccountValidityWindowDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | Pointer to **string** | Specifies the name of the detailed report file. | [optional] 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AccountExpirationWarningInterval** | Pointer to **string** | If set, the auditor will report all users with account expiration times are in the future, but are within the specified length of time away from the current time. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAddAccountValidityWindowDataSecurityAuditorRequest

`func NewAddAccountValidityWindowDataSecurityAuditorRequest(auditorName string, schemas []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn, ) *AddAccountValidityWindowDataSecurityAuditorRequest`

NewAddAccountValidityWindowDataSecurityAuditorRequest instantiates a new AddAccountValidityWindowDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountValidityWindowDataSecurityAuditorRequestWithDefaults

`func NewAddAccountValidityWindowDataSecurityAuditorRequestWithDefaults() *AddAccountValidityWindowDataSecurityAuditorRequest`

NewAddAccountValidityWindowDataSecurityAuditorRequestWithDefaults instantiates a new AddAccountValidityWindowDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetSchemas() []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumaccountValidityWindowDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetSchemas(v []EnumaccountValidityWindowDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.

### HasReportFile

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) HasReportFile() bool`

HasReportFile returns a boolean if a field has been set.

### GetIncludeAttribute

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAccountExpirationWarningInterval

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAccountExpirationWarningInterval() string`

GetAccountExpirationWarningInterval returns the AccountExpirationWarningInterval field if non-nil, zero value otherwise.

### GetAccountExpirationWarningIntervalOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAccountExpirationWarningIntervalOk() (*string, bool)`

GetAccountExpirationWarningIntervalOk returns a tuple with the AccountExpirationWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationWarningInterval

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetAccountExpirationWarningInterval(v string)`

SetAccountExpirationWarningInterval sets AccountExpirationWarningInterval field to given value.

### HasAccountExpirationWarningInterval

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) HasAccountExpirationWarningInterval() bool`

HasAccountExpirationWarningInterval returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddAccountValidityWindowDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


