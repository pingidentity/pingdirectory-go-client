# AddPrivilegeDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnumprivilegeDataSecurityAuditorSchemaUrn**](EnumprivilegeDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludePrivilege** | Pointer to [**[]EnumdataSecurityAuditorIncludePrivilegeProp**](EnumdataSecurityAuditorIncludePrivilegeProp.md) |  | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAddPrivilegeDataSecurityAuditorRequest

`func NewAddPrivilegeDataSecurityAuditorRequest(auditorName string, schemas []EnumprivilegeDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *AddPrivilegeDataSecurityAuditorRequest`

NewAddPrivilegeDataSecurityAuditorRequest instantiates a new AddPrivilegeDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPrivilegeDataSecurityAuditorRequestWithDefaults

`func NewAddPrivilegeDataSecurityAuditorRequestWithDefaults() *AddPrivilegeDataSecurityAuditorRequest`

NewAddPrivilegeDataSecurityAuditorRequestWithDefaults instantiates a new AddPrivilegeDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetSchemas() []EnumprivilegeDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumprivilegeDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetSchemas(v []EnumprivilegeDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludePrivilege

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludePrivilege() []EnumdataSecurityAuditorIncludePrivilegeProp`

GetIncludePrivilege returns the IncludePrivilege field if non-nil, zero value otherwise.

### GetIncludePrivilegeOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludePrivilegeOk() (*[]EnumdataSecurityAuditorIncludePrivilegeProp, bool)`

GetIncludePrivilegeOk returns a tuple with the IncludePrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivilege

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetIncludePrivilege(v []EnumdataSecurityAuditorIncludePrivilegeProp)`

SetIncludePrivilege sets IncludePrivilege field to given value.

### HasIncludePrivilege

`func (o *AddPrivilegeDataSecurityAuditorRequest) HasIncludePrivilege() bool`

HasIncludePrivilege returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddPrivilegeDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddPrivilegeDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddPrivilegeDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddPrivilegeDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddPrivilegeDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


