# AddWeaklyEncodedPasswordDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn**](EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**WeakPasswordStorageScheme** | **[]string** | The password storage schemes that are considered weak. Users with any of the specified password storage schemes will be included in the report. | 
**WeakCryptEncoding** | Pointer to [**[]EnumdataSecurityAuditorWeakCryptEncodingProp**](EnumdataSecurityAuditorWeakCryptEncodingProp.md) |  | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewAddWeaklyEncodedPasswordDataSecurityAuditorRequest

`func NewAddWeaklyEncodedPasswordDataSecurityAuditorRequest(auditorName string, schemas []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, reportFile string, weakPasswordStorageScheme []string, enabled bool, ) *AddWeaklyEncodedPasswordDataSecurityAuditorRequest`

NewAddWeaklyEncodedPasswordDataSecurityAuditorRequest instantiates a new AddWeaklyEncodedPasswordDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWeaklyEncodedPasswordDataSecurityAuditorRequestWithDefaults

`func NewAddWeaklyEncodedPasswordDataSecurityAuditorRequestWithDefaults() *AddWeaklyEncodedPasswordDataSecurityAuditorRequest`

NewAddWeaklyEncodedPasswordDataSecurityAuditorRequestWithDefaults instantiates a new AddWeaklyEncodedPasswordDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetSchemas() []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetSchemas(v []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetWeakPasswordStorageScheme

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakPasswordStorageScheme() []string`

GetWeakPasswordStorageScheme returns the WeakPasswordStorageScheme field if non-nil, zero value otherwise.

### GetWeakPasswordStorageSchemeOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakPasswordStorageSchemeOk() (*[]string, bool)`

GetWeakPasswordStorageSchemeOk returns a tuple with the WeakPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswordStorageScheme

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetWeakPasswordStorageScheme(v []string)`

SetWeakPasswordStorageScheme sets WeakPasswordStorageScheme field to given value.


### GetWeakCryptEncoding

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakCryptEncoding() []EnumdataSecurityAuditorWeakCryptEncodingProp`

GetWeakCryptEncoding returns the WeakCryptEncoding field if non-nil, zero value otherwise.

### GetWeakCryptEncodingOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetWeakCryptEncodingOk() (*[]EnumdataSecurityAuditorWeakCryptEncodingProp, bool)`

GetWeakCryptEncodingOk returns a tuple with the WeakCryptEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakCryptEncoding

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetWeakCryptEncoding(v []EnumdataSecurityAuditorWeakCryptEncodingProp)`

SetWeakCryptEncoding sets WeakCryptEncoding field to given value.

### HasWeakCryptEncoding

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasWeakCryptEncoding() bool`

HasWeakCryptEncoding returns a boolean if a field has been set.

### GetEnabled

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddWeaklyEncodedPasswordDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


