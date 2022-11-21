# DisabledAccountDataSecurityAuditorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnumdisabledAccountDataSecurityAuditorSchemaUrn**](EnumdisabledAccountDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewDisabledAccountDataSecurityAuditorResponse

`func NewDisabledAccountDataSecurityAuditorResponse(id string, schemas []EnumdisabledAccountDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *DisabledAccountDataSecurityAuditorResponse`

NewDisabledAccountDataSecurityAuditorResponse instantiates a new DisabledAccountDataSecurityAuditorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisabledAccountDataSecurityAuditorResponseWithDefaults

`func NewDisabledAccountDataSecurityAuditorResponseWithDefaults() *DisabledAccountDataSecurityAuditorResponse`

NewDisabledAccountDataSecurityAuditorResponseWithDefaults instantiates a new DisabledAccountDataSecurityAuditorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisabledAccountDataSecurityAuditorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisabledAccountDataSecurityAuditorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DisabledAccountDataSecurityAuditorResponse) GetSchemas() []EnumdisabledAccountDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetSchemasOk() (*[]EnumdisabledAccountDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DisabledAccountDataSecurityAuditorResponse) SetSchemas(v []EnumdisabledAccountDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *DisabledAccountDataSecurityAuditorResponse) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *DisabledAccountDataSecurityAuditorResponse) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetEnabled

`func (o *DisabledAccountDataSecurityAuditorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DisabledAccountDataSecurityAuditorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *DisabledAccountDataSecurityAuditorResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *DisabledAccountDataSecurityAuditorResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *DisabledAccountDataSecurityAuditorResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *DisabledAccountDataSecurityAuditorResponse) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *DisabledAccountDataSecurityAuditorResponse) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *DisabledAccountDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *DisabledAccountDataSecurityAuditorResponse) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetMeta

`func (o *DisabledAccountDataSecurityAuditorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DisabledAccountDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DisabledAccountDataSecurityAuditorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DisabledAccountDataSecurityAuditorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


