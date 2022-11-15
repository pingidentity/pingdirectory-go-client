# AccessControlDataSecurityAuditorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaccessControlDataSecurityAuditorSchemaUrn**](EnumaccessControlDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAccessControlDataSecurityAuditorShared

`func NewAccessControlDataSecurityAuditorShared(schemas []EnumaccessControlDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *AccessControlDataSecurityAuditorShared`

NewAccessControlDataSecurityAuditorShared instantiates a new AccessControlDataSecurityAuditorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlDataSecurityAuditorSharedWithDefaults

`func NewAccessControlDataSecurityAuditorSharedWithDefaults() *AccessControlDataSecurityAuditorShared`

NewAccessControlDataSecurityAuditorSharedWithDefaults instantiates a new AccessControlDataSecurityAuditorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AccessControlDataSecurityAuditorShared) GetSchemas() []EnumaccessControlDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AccessControlDataSecurityAuditorShared) GetSchemasOk() (*[]EnumaccessControlDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AccessControlDataSecurityAuditorShared) SetSchemas(v []EnumaccessControlDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AccessControlDataSecurityAuditorShared) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AccessControlDataSecurityAuditorShared) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AccessControlDataSecurityAuditorShared) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludeAttribute

`func (o *AccessControlDataSecurityAuditorShared) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AccessControlDataSecurityAuditorShared) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AccessControlDataSecurityAuditorShared) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AccessControlDataSecurityAuditorShared) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetEnabled

`func (o *AccessControlDataSecurityAuditorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccessControlDataSecurityAuditorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccessControlDataSecurityAuditorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuditBackend

`func (o *AccessControlDataSecurityAuditorShared) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AccessControlDataSecurityAuditorShared) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AccessControlDataSecurityAuditorShared) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AccessControlDataSecurityAuditorShared) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AccessControlDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AccessControlDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AccessControlDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AccessControlDataSecurityAuditorShared) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


