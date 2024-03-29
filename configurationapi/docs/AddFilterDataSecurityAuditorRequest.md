# AddFilterDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfilterDataSecurityAuditorSchemaUrn**](EnumfilterDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**Filter** | **[]string** | The filter to use to identify entries that should be reported. Multiple filters may be configured, and each reported entry will indicate which of these filter(s) matched that entry. | 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**AuditorName** | **string** | Name of the new Data Security Auditor | 

## Methods

### NewAddFilterDataSecurityAuditorRequest

`func NewAddFilterDataSecurityAuditorRequest(schemas []EnumfilterDataSecurityAuditorSchemaUrn, reportFile string, filter []string, auditorName string, ) *AddFilterDataSecurityAuditorRequest`

NewAddFilterDataSecurityAuditorRequest instantiates a new AddFilterDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFilterDataSecurityAuditorRequestWithDefaults

`func NewAddFilterDataSecurityAuditorRequestWithDefaults() *AddFilterDataSecurityAuditorRequest`

NewAddFilterDataSecurityAuditorRequestWithDefaults instantiates a new AddFilterDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddFilterDataSecurityAuditorRequest) GetSchemas() []EnumfilterDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFilterDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumfilterDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFilterDataSecurityAuditorRequest) SetSchemas(v []EnumfilterDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddFilterDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddFilterDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddFilterDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetFilter

`func (o *AddFilterDataSecurityAuditorRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddFilterDataSecurityAuditorRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddFilterDataSecurityAuditorRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.


### GetEnabled

`func (o *AddFilterDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddFilterDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddFilterDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddFilterDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIncludeAttribute

`func (o *AddFilterDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddFilterDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddFilterDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddFilterDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddFilterDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddFilterDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddFilterDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddFilterDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddFilterDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddFilterDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddFilterDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddFilterDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetAuditorName

`func (o *AddFilterDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddFilterDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddFilterDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


