# AddMultiplePasswordDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnummultiplePasswordDataSecurityAuditorSchemaUrn**](EnummultiplePasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAddMultiplePasswordDataSecurityAuditorRequest

`func NewAddMultiplePasswordDataSecurityAuditorRequest(auditorName string, schemas []EnummultiplePasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *AddMultiplePasswordDataSecurityAuditorRequest`

NewAddMultiplePasswordDataSecurityAuditorRequest instantiates a new AddMultiplePasswordDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMultiplePasswordDataSecurityAuditorRequestWithDefaults

`func NewAddMultiplePasswordDataSecurityAuditorRequestWithDefaults() *AddMultiplePasswordDataSecurityAuditorRequest`

NewAddMultiplePasswordDataSecurityAuditorRequestWithDefaults instantiates a new AddMultiplePasswordDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetSchemas() []EnummultiplePasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetSchemasOk() (*[]EnummultiplePasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetSchemas(v []EnummultiplePasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetEnabled

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddMultiplePasswordDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


