# MultiplePasswordDataSecurityAuditorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummultiplePasswordDataSecurityAuditorSchemaUrn**](EnummultiplePasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewMultiplePasswordDataSecurityAuditorShared

`func NewMultiplePasswordDataSecurityAuditorShared(schemas []EnummultiplePasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *MultiplePasswordDataSecurityAuditorShared`

NewMultiplePasswordDataSecurityAuditorShared instantiates a new MultiplePasswordDataSecurityAuditorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiplePasswordDataSecurityAuditorSharedWithDefaults

`func NewMultiplePasswordDataSecurityAuditorSharedWithDefaults() *MultiplePasswordDataSecurityAuditorShared`

NewMultiplePasswordDataSecurityAuditorSharedWithDefaults instantiates a new MultiplePasswordDataSecurityAuditorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MultiplePasswordDataSecurityAuditorShared) GetSchemas() []EnummultiplePasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MultiplePasswordDataSecurityAuditorShared) GetSchemasOk() (*[]EnummultiplePasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MultiplePasswordDataSecurityAuditorShared) SetSchemas(v []EnummultiplePasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *MultiplePasswordDataSecurityAuditorShared) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *MultiplePasswordDataSecurityAuditorShared) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *MultiplePasswordDataSecurityAuditorShared) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetEnabled

`func (o *MultiplePasswordDataSecurityAuditorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MultiplePasswordDataSecurityAuditorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MultiplePasswordDataSecurityAuditorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *MultiplePasswordDataSecurityAuditorShared) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *MultiplePasswordDataSecurityAuditorShared) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *MultiplePasswordDataSecurityAuditorShared) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *MultiplePasswordDataSecurityAuditorShared) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *MultiplePasswordDataSecurityAuditorShared) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *MultiplePasswordDataSecurityAuditorShared) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *MultiplePasswordDataSecurityAuditorShared) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *MultiplePasswordDataSecurityAuditorShared) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *MultiplePasswordDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *MultiplePasswordDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *MultiplePasswordDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *MultiplePasswordDataSecurityAuditorShared) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


