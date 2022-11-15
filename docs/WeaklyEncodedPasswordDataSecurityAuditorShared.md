# WeaklyEncodedPasswordDataSecurityAuditorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn**](EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**WeakPasswordStorageScheme** | **[]string** |  | 
**WeakCryptEncoding** | Pointer to [**[]EnumdataSecurityAuditorWeakCryptEncodingProp**](EnumdataSecurityAuditorWeakCryptEncodingProp.md) |  | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**IncludeAttribute** | Pointer to **[]string** |  | [optional] 
**AuditBackend** | Pointer to **[]string** |  | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 

## Methods

### NewWeaklyEncodedPasswordDataSecurityAuditorShared

`func NewWeaklyEncodedPasswordDataSecurityAuditorShared(schemas []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, reportFile string, weakPasswordStorageScheme []string, enabled bool, ) *WeaklyEncodedPasswordDataSecurityAuditorShared`

NewWeaklyEncodedPasswordDataSecurityAuditorShared instantiates a new WeaklyEncodedPasswordDataSecurityAuditorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeaklyEncodedPasswordDataSecurityAuditorSharedWithDefaults

`func NewWeaklyEncodedPasswordDataSecurityAuditorSharedWithDefaults() *WeaklyEncodedPasswordDataSecurityAuditorShared`

NewWeaklyEncodedPasswordDataSecurityAuditorSharedWithDefaults instantiates a new WeaklyEncodedPasswordDataSecurityAuditorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetSchemas() []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetSchemasOk() (*[]EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetSchemas(v []EnumweaklyEncodedPasswordDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetWeakPasswordStorageScheme

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetWeakPasswordStorageScheme() []string`

GetWeakPasswordStorageScheme returns the WeakPasswordStorageScheme field if non-nil, zero value otherwise.

### GetWeakPasswordStorageSchemeOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetWeakPasswordStorageSchemeOk() (*[]string, bool)`

GetWeakPasswordStorageSchemeOk returns a tuple with the WeakPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswordStorageScheme

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetWeakPasswordStorageScheme(v []string)`

SetWeakPasswordStorageScheme sets WeakPasswordStorageScheme field to given value.


### GetWeakCryptEncoding

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetWeakCryptEncoding() []EnumdataSecurityAuditorWeakCryptEncodingProp`

GetWeakCryptEncoding returns the WeakCryptEncoding field if non-nil, zero value otherwise.

### GetWeakCryptEncodingOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetWeakCryptEncodingOk() (*[]EnumdataSecurityAuditorWeakCryptEncodingProp, bool)`

GetWeakCryptEncodingOk returns a tuple with the WeakCryptEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakCryptEncoding

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetWeakCryptEncoding(v []EnumdataSecurityAuditorWeakCryptEncodingProp)`

SetWeakCryptEncoding sets WeakCryptEncoding field to given value.

### HasWeakCryptEncoding

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) HasWeakCryptEncoding() bool`

HasWeakCryptEncoding returns a boolean if a field has been set.

### GetEnabled

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeAttribute

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetAuditBackend

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *WeaklyEncodedPasswordDataSecurityAuditorShared) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


