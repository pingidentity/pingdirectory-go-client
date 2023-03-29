# AddDataSecurityAuditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditorName** | **string** | Name of the new Data Security Auditor | 
**Schemas** | [**[]EnumthirdPartyDataSecurityAuditorSchemaUrn**](EnumthirdPartyDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**PasswordEvaluationAge** | Pointer to **string** | If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Data Security Auditor is enabled for use. | [optional] 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**IdleAccountWarningInterval** | **string** | The length of time to use as the warning interval for idle accounts. If the length of time since a user last authenticated is greater than the warning interval but less than the error interval (or if it is greater than the warning interval and no error interval is defined), then a warning will be generated for that account. | 
**IdleAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for idle accounts. If the length of time since a user last authenticated is greater than the error interval, then an error will be generated for that account. If no error interval is defined, then only the warning interval will be used. | [optional] 
**NeverLoggedInAccountWarningInterval** | Pointer to **string** | The length of time to use as the warning interval for accounts that do not appear to have authenticated. If this is not specified, then the idle account warning interval will be used. | [optional] 
**NeverLoggedInAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for accounts that do not appear to have authenticated. If this is not specified, then the never-logged-in warning interval will be used. The idle account warning and error intervals will be used if no never-logged-in interval is configured. | [optional] 
**WeakPasswordStorageScheme** | Pointer to **[]string** | The password storage schemes that are considered weak. Users with any of the specified password storage schemes will be included in the report. | [optional] 
**WeakCryptEncoding** | Pointer to [**[]EnumdataSecurityAuditorWeakCryptEncodingProp**](EnumdataSecurityAuditorWeakCryptEncodingProp.md) |  | [optional] 
**IncludePrivilege** | Pointer to [**[]EnumdataSecurityAuditorIncludePrivilegeProp**](EnumdataSecurityAuditorIncludePrivilegeProp.md) |  | [optional] 
**MaximumIdleTime** | Pointer to **string** | If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled. | [optional] 
**Filter** | **[]string** | The filter to use to identify entries that should be reported. Multiple filters may be configured, and each reported entry will indicate which of these filter(s) matched that entry. | 
**AccountExpirationWarningInterval** | Pointer to **string** | If set, the auditor will report all users with account expiration times are in the future, but are within the specified length of time away from the current time. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Data Security Auditor. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Data Security Auditor. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddDataSecurityAuditorRequest

`func NewAddDataSecurityAuditorRequest(auditorName string, schemas []EnumthirdPartyDataSecurityAuditorSchemaUrn, reportFile string, idleAccountWarningInterval string, filter []string, extensionClass string, ) *AddDataSecurityAuditorRequest`

NewAddDataSecurityAuditorRequest instantiates a new AddDataSecurityAuditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDataSecurityAuditorRequestWithDefaults

`func NewAddDataSecurityAuditorRequestWithDefaults() *AddDataSecurityAuditorRequest`

NewAddDataSecurityAuditorRequestWithDefaults instantiates a new AddDataSecurityAuditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditorName

`func (o *AddDataSecurityAuditorRequest) GetAuditorName() string`

GetAuditorName returns the AuditorName field if non-nil, zero value otherwise.

### GetAuditorNameOk

`func (o *AddDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool)`

GetAuditorNameOk returns a tuple with the AuditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorName

`func (o *AddDataSecurityAuditorRequest) SetAuditorName(v string)`

SetAuditorName sets AuditorName field to given value.


### GetSchemas

`func (o *AddDataSecurityAuditorRequest) GetSchemas() []EnumthirdPartyDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDataSecurityAuditorRequest) GetSchemasOk() (*[]EnumthirdPartyDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDataSecurityAuditorRequest) SetSchemas(v []EnumthirdPartyDataSecurityAuditorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReportFile

`func (o *AddDataSecurityAuditorRequest) GetReportFile() string`

GetReportFile returns the ReportFile field if non-nil, zero value otherwise.

### GetReportFileOk

`func (o *AddDataSecurityAuditorRequest) GetReportFileOk() (*string, bool)`

GetReportFileOk returns a tuple with the ReportFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFile

`func (o *AddDataSecurityAuditorRequest) SetReportFile(v string)`

SetReportFile sets ReportFile field to given value.


### GetIncludeAttribute

`func (o *AddDataSecurityAuditorRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddDataSecurityAuditorRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddDataSecurityAuditorRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddDataSecurityAuditorRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetPasswordEvaluationAge

`func (o *AddDataSecurityAuditorRequest) GetPasswordEvaluationAge() string`

GetPasswordEvaluationAge returns the PasswordEvaluationAge field if non-nil, zero value otherwise.

### GetPasswordEvaluationAgeOk

`func (o *AddDataSecurityAuditorRequest) GetPasswordEvaluationAgeOk() (*string, bool)`

GetPasswordEvaluationAgeOk returns a tuple with the PasswordEvaluationAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEvaluationAge

`func (o *AddDataSecurityAuditorRequest) SetPasswordEvaluationAge(v string)`

SetPasswordEvaluationAge sets PasswordEvaluationAge field to given value.

### HasPasswordEvaluationAge

`func (o *AddDataSecurityAuditorRequest) HasPasswordEvaluationAge() bool`

HasPasswordEvaluationAge returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDataSecurityAuditorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDataSecurityAuditorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddDataSecurityAuditorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuditBackend

`func (o *AddDataSecurityAuditorRequest) GetAuditBackend() []string`

GetAuditBackend returns the AuditBackend field if non-nil, zero value otherwise.

### GetAuditBackendOk

`func (o *AddDataSecurityAuditorRequest) GetAuditBackendOk() (*[]string, bool)`

GetAuditBackendOk returns a tuple with the AuditBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBackend

`func (o *AddDataSecurityAuditorRequest) SetAuditBackend(v []string)`

SetAuditBackend sets AuditBackend field to given value.

### HasAuditBackend

`func (o *AddDataSecurityAuditorRequest) HasAuditBackend() bool`

HasAuditBackend returns a boolean if a field has been set.

### GetAuditSeverity

`func (o *AddDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp`

GetAuditSeverity returns the AuditSeverity field if non-nil, zero value otherwise.

### GetAuditSeverityOk

`func (o *AddDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool)`

GetAuditSeverityOk returns a tuple with the AuditSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSeverity

`func (o *AddDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp)`

SetAuditSeverity sets AuditSeverity field to given value.

### HasAuditSeverity

`func (o *AddDataSecurityAuditorRequest) HasAuditSeverity() bool`

HasAuditSeverity returns a boolean if a field has been set.

### GetIdleAccountWarningInterval

`func (o *AddDataSecurityAuditorRequest) GetIdleAccountWarningInterval() string`

GetIdleAccountWarningInterval returns the IdleAccountWarningInterval field if non-nil, zero value otherwise.

### GetIdleAccountWarningIntervalOk

`func (o *AddDataSecurityAuditorRequest) GetIdleAccountWarningIntervalOk() (*string, bool)`

GetIdleAccountWarningIntervalOk returns a tuple with the IdleAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountWarningInterval

`func (o *AddDataSecurityAuditorRequest) SetIdleAccountWarningInterval(v string)`

SetIdleAccountWarningInterval sets IdleAccountWarningInterval field to given value.


### GetIdleAccountErrorInterval

`func (o *AddDataSecurityAuditorRequest) GetIdleAccountErrorInterval() string`

GetIdleAccountErrorInterval returns the IdleAccountErrorInterval field if non-nil, zero value otherwise.

### GetIdleAccountErrorIntervalOk

`func (o *AddDataSecurityAuditorRequest) GetIdleAccountErrorIntervalOk() (*string, bool)`

GetIdleAccountErrorIntervalOk returns a tuple with the IdleAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountErrorInterval

`func (o *AddDataSecurityAuditorRequest) SetIdleAccountErrorInterval(v string)`

SetIdleAccountErrorInterval sets IdleAccountErrorInterval field to given value.

### HasIdleAccountErrorInterval

`func (o *AddDataSecurityAuditorRequest) HasIdleAccountErrorInterval() bool`

HasIdleAccountErrorInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountWarningInterval

`func (o *AddDataSecurityAuditorRequest) GetNeverLoggedInAccountWarningInterval() string`

GetNeverLoggedInAccountWarningInterval returns the NeverLoggedInAccountWarningInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountWarningIntervalOk

`func (o *AddDataSecurityAuditorRequest) GetNeverLoggedInAccountWarningIntervalOk() (*string, bool)`

GetNeverLoggedInAccountWarningIntervalOk returns a tuple with the NeverLoggedInAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountWarningInterval

`func (o *AddDataSecurityAuditorRequest) SetNeverLoggedInAccountWarningInterval(v string)`

SetNeverLoggedInAccountWarningInterval sets NeverLoggedInAccountWarningInterval field to given value.

### HasNeverLoggedInAccountWarningInterval

`func (o *AddDataSecurityAuditorRequest) HasNeverLoggedInAccountWarningInterval() bool`

HasNeverLoggedInAccountWarningInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountErrorInterval

`func (o *AddDataSecurityAuditorRequest) GetNeverLoggedInAccountErrorInterval() string`

GetNeverLoggedInAccountErrorInterval returns the NeverLoggedInAccountErrorInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountErrorIntervalOk

`func (o *AddDataSecurityAuditorRequest) GetNeverLoggedInAccountErrorIntervalOk() (*string, bool)`

GetNeverLoggedInAccountErrorIntervalOk returns a tuple with the NeverLoggedInAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountErrorInterval

`func (o *AddDataSecurityAuditorRequest) SetNeverLoggedInAccountErrorInterval(v string)`

SetNeverLoggedInAccountErrorInterval sets NeverLoggedInAccountErrorInterval field to given value.

### HasNeverLoggedInAccountErrorInterval

`func (o *AddDataSecurityAuditorRequest) HasNeverLoggedInAccountErrorInterval() bool`

HasNeverLoggedInAccountErrorInterval returns a boolean if a field has been set.

### GetWeakPasswordStorageScheme

`func (o *AddDataSecurityAuditorRequest) GetWeakPasswordStorageScheme() []string`

GetWeakPasswordStorageScheme returns the WeakPasswordStorageScheme field if non-nil, zero value otherwise.

### GetWeakPasswordStorageSchemeOk

`func (o *AddDataSecurityAuditorRequest) GetWeakPasswordStorageSchemeOk() (*[]string, bool)`

GetWeakPasswordStorageSchemeOk returns a tuple with the WeakPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswordStorageScheme

`func (o *AddDataSecurityAuditorRequest) SetWeakPasswordStorageScheme(v []string)`

SetWeakPasswordStorageScheme sets WeakPasswordStorageScheme field to given value.

### HasWeakPasswordStorageScheme

`func (o *AddDataSecurityAuditorRequest) HasWeakPasswordStorageScheme() bool`

HasWeakPasswordStorageScheme returns a boolean if a field has been set.

### GetWeakCryptEncoding

`func (o *AddDataSecurityAuditorRequest) GetWeakCryptEncoding() []EnumdataSecurityAuditorWeakCryptEncodingProp`

GetWeakCryptEncoding returns the WeakCryptEncoding field if non-nil, zero value otherwise.

### GetWeakCryptEncodingOk

`func (o *AddDataSecurityAuditorRequest) GetWeakCryptEncodingOk() (*[]EnumdataSecurityAuditorWeakCryptEncodingProp, bool)`

GetWeakCryptEncodingOk returns a tuple with the WeakCryptEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakCryptEncoding

`func (o *AddDataSecurityAuditorRequest) SetWeakCryptEncoding(v []EnumdataSecurityAuditorWeakCryptEncodingProp)`

SetWeakCryptEncoding sets WeakCryptEncoding field to given value.

### HasWeakCryptEncoding

`func (o *AddDataSecurityAuditorRequest) HasWeakCryptEncoding() bool`

HasWeakCryptEncoding returns a boolean if a field has been set.

### GetIncludePrivilege

`func (o *AddDataSecurityAuditorRequest) GetIncludePrivilege() []EnumdataSecurityAuditorIncludePrivilegeProp`

GetIncludePrivilege returns the IncludePrivilege field if non-nil, zero value otherwise.

### GetIncludePrivilegeOk

`func (o *AddDataSecurityAuditorRequest) GetIncludePrivilegeOk() (*[]EnumdataSecurityAuditorIncludePrivilegeProp, bool)`

GetIncludePrivilegeOk returns a tuple with the IncludePrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivilege

`func (o *AddDataSecurityAuditorRequest) SetIncludePrivilege(v []EnumdataSecurityAuditorIncludePrivilegeProp)`

SetIncludePrivilege sets IncludePrivilege field to given value.

### HasIncludePrivilege

`func (o *AddDataSecurityAuditorRequest) HasIncludePrivilege() bool`

HasIncludePrivilege returns a boolean if a field has been set.

### GetMaximumIdleTime

`func (o *AddDataSecurityAuditorRequest) GetMaximumIdleTime() string`

GetMaximumIdleTime returns the MaximumIdleTime field if non-nil, zero value otherwise.

### GetMaximumIdleTimeOk

`func (o *AddDataSecurityAuditorRequest) GetMaximumIdleTimeOk() (*string, bool)`

GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIdleTime

`func (o *AddDataSecurityAuditorRequest) SetMaximumIdleTime(v string)`

SetMaximumIdleTime sets MaximumIdleTime field to given value.

### HasMaximumIdleTime

`func (o *AddDataSecurityAuditorRequest) HasMaximumIdleTime() bool`

HasMaximumIdleTime returns a boolean if a field has been set.

### GetFilter

`func (o *AddDataSecurityAuditorRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddDataSecurityAuditorRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddDataSecurityAuditorRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.


### GetAccountExpirationWarningInterval

`func (o *AddDataSecurityAuditorRequest) GetAccountExpirationWarningInterval() string`

GetAccountExpirationWarningInterval returns the AccountExpirationWarningInterval field if non-nil, zero value otherwise.

### GetAccountExpirationWarningIntervalOk

`func (o *AddDataSecurityAuditorRequest) GetAccountExpirationWarningIntervalOk() (*string, bool)`

GetAccountExpirationWarningIntervalOk returns a tuple with the AccountExpirationWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationWarningInterval

`func (o *AddDataSecurityAuditorRequest) SetAccountExpirationWarningInterval(v string)`

SetAccountExpirationWarningInterval sets AccountExpirationWarningInterval field to given value.

### HasAccountExpirationWarningInterval

`func (o *AddDataSecurityAuditorRequest) HasAccountExpirationWarningInterval() bool`

HasAccountExpirationWarningInterval returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddDataSecurityAuditorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddDataSecurityAuditorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddDataSecurityAuditorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddDataSecurityAuditorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddDataSecurityAuditorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddDataSecurityAuditorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddDataSecurityAuditorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


