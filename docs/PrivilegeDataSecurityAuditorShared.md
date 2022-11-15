# PrivilegeDataSecurityAuditorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumprivilegeDataSecurityAuditorSchemaUrn**](EnumprivilegeDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludePrivilege** | Pointer to [**[]EnumdataSecurityAuditorIncludePrivilegeProp**](EnumdataSecurityAuditorIncludePrivilegeProp.md) |  | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewPrivilegeDataSecurityAuditorShared

`func NewPrivilegeDataSecurityAuditorShared(schemas []EnumprivilegeDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *PrivilegeDataSecurityAuditorShared`

NewPrivilegeDataSecurityAuditorShared instantiates a new PrivilegeDataSecurityAuditorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeDataSecurityAuditorSharedWithDefaults

`func NewPrivilegeDataSecurityAuditorSharedWithDefaults() *PrivilegeDataSecurityAuditorShared`

NewPrivilegeDataSecurityAuditorSharedWithDefaults instantiates a new PrivilegeDataSecurityAuditorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PrivilegeDataSecurityAuditorShared) GetSchemas() []EnumprivilegeDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PrivilegeDataSecurityAuditorShared) GetSchemasOk() (*[]EnumprivilegeDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PrivilegeDataSecurityAuditorShared) SetSchemas(v []EnumprivilegeDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *PrivilegeDataSecurityAuditorShared) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *PrivilegeDataSecurityAuditorShared) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *PrivilegeDataSecurityAuditorShared) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludePrivilege

`func (o *PrivilegeDataSecurityAuditorShared) GetIncludePrivilege() []EnumdataSecurityAuditorIncludePrivilegeProp`

GetIncludePrivilege returns the IncludePrivilege field if non-nil, zero value otherwise.

### GetIncludePrivilegeOk

`func (o *PrivilegeDataSecurityAuditorShared) GetIncludePrivilegeOk() (*[]EnumdataSecurityAuditorIncludePrivilegeProp, bool)`

GetIncludePrivilegeOk returns a tuple with the IncludePrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivilege

`func (o *PrivilegeDataSecurityAuditorShared) SetIncludePrivilege(v []EnumdataSecurityAuditorIncludePrivilegeProp)`

SetIncludePrivilege sets IncludePrivilege field to given value.

### HasIncludePrivilege

`func (o *PrivilegeDataSecurityAuditorShared) HasIncludePrivilege() bool`

HasIncludePrivilege returns a boolean if a field has been set.

### GetEnabled

`func (o *PrivilegeDataSecurityAuditorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PrivilegeDataSecurityAuditorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PrivilegeDataSecurityAuditorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *PrivilegeDataSecurityAuditorShared) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *PrivilegeDataSecurityAuditorShared) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *PrivilegeDataSecurityAuditorShared) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *PrivilegeDataSecurityAuditorShared) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *PrivilegeDataSecurityAuditorShared) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *PrivilegeDataSecurityAuditorShared) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *PrivilegeDataSecurityAuditorShared) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *PrivilegeDataSecurityAuditorShared) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *PrivilegeDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *PrivilegeDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *PrivilegeDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *PrivilegeDataSecurityAuditorShared) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


