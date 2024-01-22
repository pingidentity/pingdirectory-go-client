# AddNonexistentPasswordPolicyDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumnonexistentPasswordPolicyDataSecurityAuditorSchemaUrn**](EnumnonexistentPasswordPolicyDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | Pointer to **string** | Specifies the name of the detailed report file. | [optional] 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**AuditorName** | **string** | Name of the new Data Security Auditor | 

## Methods

### NewAddNonexistentPasswordPolicyDataSecurityAuditorRequest

`func NewAddNonexistentPasswordPolicyDataSecurityAuditorRequest(schemas []EnumnonexistentPasswordPolicyDataSecurityAuditorSchemaUrn, auditorName string, ) *AddNonexistentPasswordPolicyDataSecurityAuditorRequest`

NewAddNonexistentPasswordPolicyDataSecurityAuditorRequest instantiates a new AddNonexistentPasswordPolicyDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNonexistentPasswordPolicyDataSecurityAuditorRequestWithDefaults

`func NewAddNonexistentPasswordPolicyDataSecurityAuditorRequestWithDefaults() *AddNonexistentPasswordPolicyDataSecurityAuditorRequest`

NewAddNonexistentPasswordPolicyDataSecurityAuditorRequestWithDefaults instantiates a new AddNonexistentPasswordPolicyDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetSchemas() []EnumnonexistentPasswordPolicyDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumnonexistentPasswordPolicyDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetSchemas(v []EnumnonexistentPasswordPolicyDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.

### HasReportFile

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) HasReportFile() bool`

HasReportFile returns a boolean if a field has been set.

### GetIncludeAttribute

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetEnabled

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetAuditorName

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddNonexistentPasswordPolicyDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


