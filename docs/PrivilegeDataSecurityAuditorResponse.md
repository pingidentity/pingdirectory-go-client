# PrivilegeDataSecurityAuditorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnumprivilegeDataSecurityAuditorSchemaUrn**](EnumprivilegeDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludePrivilege** | Pointer to [**[]EnumdataSecurityAuditorIncludePrivilegeProp**](EnumdataSecurityAuditorIncludePrivilegeProp.md) |  | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewPrivilegeDataSecurityAuditorResponse

`func NewPrivilegeDataSecurityAuditorResponse(id string, schemas []EnumprivilegeDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *PrivilegeDataSecurityAuditorResponse`

NewPrivilegeDataSecurityAuditorResponse instantiates a new PrivilegeDataSecurityAuditorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeDataSecurityAuditorResponseWithDefaults

`func NewPrivilegeDataSecurityAuditorResponseWithDefaults() *PrivilegeDataSecurityAuditorResponse`

NewPrivilegeDataSecurityAuditorResponseWithDefaults instantiates a new PrivilegeDataSecurityAuditorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivilegeDataSecurityAuditorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivilegeDataSecurityAuditorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PrivilegeDataSecurityAuditorResponse) GetSchemas() []EnumprivilegeDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetSchemasOk() (*[]EnumprivilegeDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PrivilegeDataSecurityAuditorResponse) SetSchemas(v []EnumprivilegeDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *PrivilegeDataSecurityAuditorResponse) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *PrivilegeDataSecurityAuditorResponse) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludePrivilege

`func (o *PrivilegeDataSecurityAuditorResponse) GetIncludePrivilege() []EnumdataSecurityAuditorIncludePrivilegeProp`

GetIncludePrivilege returns the IncludePrivilege field if non-nil, zero value otherwise.

### GetIncludePrivilegeOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetIncludePrivilegeOk() (*[]EnumdataSecurityAuditorIncludePrivilegeProp, bool)`

GetIncludePrivilegeOk returns a tuple with the IncludePrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivilege

`func (o *PrivilegeDataSecurityAuditorResponse) SetIncludePrivilege(v []EnumdataSecurityAuditorIncludePrivilegeProp)`

SetIncludePrivilege sets IncludePrivilege field to given value.

### HasIncludePrivilege

`func (o *PrivilegeDataSecurityAuditorResponse) HasIncludePrivilege() bool`

HasIncludePrivilege returns a boolean if a field has been set.

### GetEnabled

`func (o *PrivilegeDataSecurityAuditorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PrivilegeDataSecurityAuditorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *PrivilegeDataSecurityAuditorResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *PrivilegeDataSecurityAuditorResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *PrivilegeDataSecurityAuditorResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *PrivilegeDataSecurityAuditorResponse) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *PrivilegeDataSecurityAuditorResponse) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *PrivilegeDataSecurityAuditorResponse) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *PrivilegeDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *PrivilegeDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *PrivilegeDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *PrivilegeDataSecurityAuditorResponse) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


