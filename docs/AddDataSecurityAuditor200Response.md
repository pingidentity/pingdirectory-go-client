# AddDataSecurityAuditor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnumaccessControlDataSecurityAuditorSchemaUrn**](EnumaccessControlDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**PasswordEvaluationAge** | Pointer to **string** | If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled. | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**WeakPasswordStorageScheme** | **[]string** | The password storage schemes that are considered weak. Users with any of the specified password storage schemes will be included in the report. | 
**WeakCryptEncoding** | Pointer to [**[]EnumdataSecurityAuditorWeakCryptEncodingProp**](EnumdataSecurityAuditorWeakCryptEncodingProp.md) |  | [optional] 
**IncludePrivilege** | Pointer to [**[]EnumdataSecurityAuditorIncludePrivilegeProp**](EnumdataSecurityAuditorIncludePrivilegeProp.md) |  | [optional] 
**MaximumIdleTime** | Pointer to **string** | If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled. | [optional] 

## Methods

### NewAddDataSecurityAuditor200Response

`func NewAddDataSecurityAuditor200Response(id string, schemas []EnumaccessControlDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, weakPasswordStorageScheme []string, ) *AddDataSecurityAuditor200Response`

NewAddDataSecurityAuditor200Response instantiates a new AddDataSecurityAuditor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDataSecurityAuditor200ResponseWithDefaults

`func NewAddDataSecurityAuditor200ResponseWithDefaults() *AddDataSecurityAuditor200Response`

NewAddDataSecurityAuditor200ResponseWithDefaults instantiates a new AddDataSecurityAuditor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddDataSecurityAuditor200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddDataSecurityAuditor200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddDataSecurityAuditor200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddDataSecurityAuditor200Response) GetSchemas() []EnumaccessControlDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDataSecurityAuditor200Response) GetSchemasOk() (*[]EnumaccessControlDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDataSecurityAuditor200Response) SetSchemas(v []EnumaccessControlDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddDataSecurityAuditor200Response) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddDataSecurityAuditor200Response) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddDataSecurityAuditor200Response) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludeAttribute

`func (o *AddDataSecurityAuditor200Response) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddDataSecurityAuditor200Response) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddDataSecurityAuditor200Response) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddDataSecurityAuditor200Response) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetPasswordEvaluationAge

`func (o *AddDataSecurityAuditor200Response) GetPasswordEvaluationAge() string`

GetPasswordEvaluationAge returns the PasswordEvaluationAge field if non-nil, zero value otherwise.

### GetPasswordEvaluationAgeOk

`func (o *AddDataSecurityAuditor200Response) GetPasswordEvaluationAgeOk() (*string, bool)`

GetPasswordEvaluationAgeOk returns a tuple with the PasswordEvaluationAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEvaluationAge

`func (o *AddDataSecurityAuditor200Response) SetPasswordEvaluationAge(v string)`

SetPasswordEvaluationAge sets PasswordEvaluationAge field to given value.

### HasPasswordEvaluationAge

`func (o *AddDataSecurityAuditor200Response) HasPasswordEvaluationAge() bool`

HasPasswordEvaluationAge returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDataSecurityAuditor200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDataSecurityAuditor200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDataSecurityAuditor200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuditBackend

`func (o *AddDataSecurityAuditor200Response) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddDataSecurityAuditor200Response) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddDataSecurityAuditor200Response) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddDataSecurityAuditor200Response) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddDataSecurityAuditor200Response) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddDataSecurityAuditor200Response) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddDataSecurityAuditor200Response) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddDataSecurityAuditor200Response) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetMeta

`func (o *AddDataSecurityAuditor200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddDataSecurityAuditor200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddDataSecurityAuditor200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddDataSecurityAuditor200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWeakPasswordStorageScheme

`func (o *AddDataSecurityAuditor200Response) GetWeakPasswordStorageScheme() []string`

GetWeakPasswordStorageScheme returns the WeakPasswordStorageScheme field if non-nil, zero value otherwise.

### GetWeakPasswordStorageSchemeOk

`func (o *AddDataSecurityAuditor200Response) GetWeakPasswordStorageSchemeOk() (*[]string, bool)`

GetWeakPasswordStorageSchemeOk returns a tuple with the WeakPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswordStorageScheme

`func (o *AddDataSecurityAuditor200Response) SetWeakPasswordStorageScheme(v []string)`

SetWeakPasswordStorageScheme sets WeakPasswordStorageScheme field to given value.


### GetWeakCryptEncoding

`func (o *AddDataSecurityAuditor200Response) GetWeakCryptEncoding() []EnumdataSecurityAuditorWeakCryptEncodingProp`

GetWeakCryptEncoding returns the WeakCryptEncoding field if non-nil, zero value otherwise.

### GetWeakCryptEncodingOk

`func (o *AddDataSecurityAuditor200Response) GetWeakCryptEncodingOk() (*[]EnumdataSecurityAuditorWeakCryptEncodingProp, bool)`

GetWeakCryptEncodingOk returns a tuple with the WeakCryptEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakCryptEncoding

`func (o *AddDataSecurityAuditor200Response) SetWeakCryptEncoding(v []EnumdataSecurityAuditorWeakCryptEncodingProp)`

SetWeakCryptEncoding sets WeakCryptEncoding field to given value.

### HasWeakCryptEncoding

`func (o *AddDataSecurityAuditor200Response) HasWeakCryptEncoding() bool`

HasWeakCryptEncoding returns a boolean if a field has been set.

### GetIncludePrivilege

`func (o *AddDataSecurityAuditor200Response) GetIncludePrivilege() []EnumdataSecurityAuditorIncludePrivilegeProp`

GetIncludePrivilege returns the IncludePrivilege field if non-nil, zero value otherwise.

### GetIncludePrivilegeOk

`func (o *AddDataSecurityAuditor200Response) GetIncludePrivilegeOk() (*[]EnumdataSecurityAuditorIncludePrivilegeProp, bool)`

GetIncludePrivilegeOk returns a tuple with the IncludePrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivilege

`func (o *AddDataSecurityAuditor200Response) SetIncludePrivilege(v []EnumdataSecurityAuditorIncludePrivilegeProp)`

SetIncludePrivilege sets IncludePrivilege field to given value.

### HasIncludePrivilege

`func (o *AddDataSecurityAuditor200Response) HasIncludePrivilege() bool`

HasIncludePrivilege returns a boolean if a field has been set.

### GetMaximumIdleTime

`func (o *AddDataSecurityAuditor200Response) GetMaximumIdleTime() string`

GetMaximumIdleTime returns the MaximumIdleTime field if non-nil, zero value otherwise.

### GetMaximumIdleTimeOk

`func (o *AddDataSecurityAuditor200Response) GetMaximumIdleTimeOk() (*string, bool)`

GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIdleTime

`func (o *AddDataSecurityAuditor200Response) SetMaximumIdleTime(v string)`

SetMaximumIdleTime sets MaximumIdleTime field to given value.

### HasMaximumIdleTime

`func (o *AddDataSecurityAuditor200Response) HasMaximumIdleTime() bool`

HasMaximumIdleTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


