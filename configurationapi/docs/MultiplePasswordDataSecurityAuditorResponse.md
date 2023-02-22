# MultiplePasswordDataSecurityAuditorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnummultiplePasswordDataSecurityAuditorSchemaUrn**](EnummultiplePasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewMultiplePasswordDataSecurityAuditorResponse

`func NewMultiplePasswordDataSecurityAuditorResponse(id string, schemas []EnummultiplePasswordDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, ) *MultiplePasswordDataSecurityAuditorResponse`

NewMultiplePasswordDataSecurityAuditorResponse instantiates a new MultiplePasswordDataSecurityAuditorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiplePasswordDataSecurityAuditorResponseWithDefaults

`func NewMultiplePasswordDataSecurityAuditorResponseWithDefaults() *MultiplePasswordDataSecurityAuditorResponse`

NewMultiplePasswordDataSecurityAuditorResponseWithDefaults instantiates a new MultiplePasswordDataSecurityAuditorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MultiplePasswordDataSecurityAuditorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MultiplePasswordDataSecurityAuditorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetSchemas() []EnummultiplePasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetSchemasOk() (*[]EnummultiplePasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetSchemas(v []EnummultiplePasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetEnabled

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *MultiplePasswordDataSecurityAuditorResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *MultiplePasswordDataSecurityAuditorResponse) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *MultiplePasswordDataSecurityAuditorResponse) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *MultiplePasswordDataSecurityAuditorResponse) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *MultiplePasswordDataSecurityAuditorResponse) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


