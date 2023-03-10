# AddDataSecurityAuditor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Data Security Auditor | 
**Schemas** | [**[]EnumthirdPartyDataSecurityAuditorSchemaUrn**](EnumthirdPartyDataSecurityAuditorSchemaUrn.md) |  | 
**ReportFile** | **string** | Specifies the name of the detailed report file. | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included. | [optional] 
**PasswordEvaluationAge** | Pointer to **string** | If set, the auditor will report all users with passwords older than the specified value even if password expiration is not enabled. | [optional] 
**Enabled** | **bool** | Indicates whether the Data Security Auditor is enabled for use. | 
**AuditBackend** | Pointer to **[]string** | Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler). | [optional] 
**AuditSeverity** | Pointer to [**EnumdataSecurityAuditorAuditSeverityProp**](EnumdataSecurityAuditorAuditSeverityProp.md) |  | [optional] 
**IdleAccountWarningInterval** | **string** | The length of time to use as the warning interval for idle accounts. If the length of time since a user last authenticated is greater than the warning interval but less than the error interval (or if it is greater than the warning interval and no error interval is defined), then a warning will be generated for that account. | 
**IdleAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for idle accounts. If the length of time since a user last authenticated is greater than the error interval, then an error will be generated for that account. If no error interval is defined, then only the warning interval will be used. | [optional] 
**NeverLoggedInAccountWarningInterval** | Pointer to **string** | The length of time to use as the warning interval for accounts that do not appear to have authenticated. If this is not specified, then the idle account warning interval will be used. | [optional] 
**NeverLoggedInAccountErrorInterval** | Pointer to **string** | The length of time to use as the error interval for accounts that do not appear to have authenticated. If this is not specified, then the never-logged-in warning interval will be used. The idle account warning and error intervals will be used if no never-logged-in interval is configured. | [optional] 
**WeakPasswordStorageScheme** | **[]string** | The password storage schemes that are considered weak. Users with any of the specified password storage schemes will be included in the report. | 
**WeakCryptEncoding** | Pointer to [**[]EnumdataSecurityAuditorWeakCryptEncodingProp**](EnumdataSecurityAuditorWeakCryptEncodingProp.md) |  | [optional] 
**IncludePrivilege** | Pointer to [**[]EnumdataSecurityAuditorIncludePrivilegeProp**](EnumdataSecurityAuditorIncludePrivilegeProp.md) |  | [optional] 
**MaximumIdleTime** | Pointer to **string** | If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled. | [optional] 
**Filter** | **[]string** | The filter to use to identify entries that should be reported. Multiple filters may be configured, and each reported entry will indicate which of these filter(s) matched that entry. | 
**AccountExpirationWarningInterval** | Pointer to **string** | If set, the auditor will report all users with account expiration times are in the future, but are within the specified length of time away from the current time. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Data Security Auditor. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Data Security Auditor. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddDataSecurityAuditor200Response

`func NewAddDataSecurityAuditor200Response(id string, schemas []EnumthirdPartyDataSecurityAuditorSchemaUrn, reportFile string, enabled bool, idleAccountWarningInterval string, weakPasswordStorageScheme []string, filter []string, extensionClass string, ) *AddDataSecurityAuditor200Response`

NewAddDataSecurityAuditor200Response instantiates a new AddDataSecurityAuditor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDataSecurityAuditor200ResponseWithDefaults

`func NewAddDataSecurityAuditor200ResponseWithDefaults() *AddDataSecurityAuditor200Response`

NewAddDataSecurityAuditor200ResponseWithDefaults instantiates a new AddDataSecurityAuditor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddDataSecurityAuditor200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddDataSecurityAuditor200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddDataSecurityAuditor200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddDataSecurityAuditor200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

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

`func (o *AddDataSecurityAuditor200Response) GetSchemas() []EnumthirdPartyDataSecurityAuditorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDataSecurityAuditor200Response) GetSchemasOk() (*[]EnumthirdPartyDataSecurityAuditorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDataSecurityAuditor200Response) SetSchemas(v []EnumthirdPartyDataSecurityAuditorSchemaUrn)`

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

### GetIdleAccountWarningInterval

`func (o *AddDataSecurityAuditor200Response) GetIdleAccountWarningInterval() string`

GetIdleAccountWarningInterval returns the IdleAccountWarningInterval field if non-nil, zero value otherwise.

### GetIdleAccountWarningIntervalOk

`func (o *AddDataSecurityAuditor200Response) GetIdleAccountWarningIntervalOk() (*string, bool)`

GetIdleAccountWarningIntervalOk returns a tuple with the IdleAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountWarningInterval

`func (o *AddDataSecurityAuditor200Response) SetIdleAccountWarningInterval(v string)`

SetIdleAccountWarningInterval sets IdleAccountWarningInterval field to given value.


### GetIdleAccountErrorInterval

`func (o *AddDataSecurityAuditor200Response) GetIdleAccountErrorInterval() string`

GetIdleAccountErrorInterval returns the IdleAccountErrorInterval field if non-nil, zero value otherwise.

### GetIdleAccountErrorIntervalOk

`func (o *AddDataSecurityAuditor200Response) GetIdleAccountErrorIntervalOk() (*string, bool)`

GetIdleAccountErrorIntervalOk returns a tuple with the IdleAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleAccountErrorInterval

`func (o *AddDataSecurityAuditor200Response) SetIdleAccountErrorInterval(v string)`

SetIdleAccountErrorInterval sets IdleAccountErrorInterval field to given value.

### HasIdleAccountErrorInterval

`func (o *AddDataSecurityAuditor200Response) HasIdleAccountErrorInterval() bool`

HasIdleAccountErrorInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountWarningInterval

`func (o *AddDataSecurityAuditor200Response) GetNeverLoggedInAccountWarningInterval() string`

GetNeverLoggedInAccountWarningInterval returns the NeverLoggedInAccountWarningInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountWarningIntervalOk

`func (o *AddDataSecurityAuditor200Response) GetNeverLoggedInAccountWarningIntervalOk() (*string, bool)`

GetNeverLoggedInAccountWarningIntervalOk returns a tuple with the NeverLoggedInAccountWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountWarningInterval

`func (o *AddDataSecurityAuditor200Response) SetNeverLoggedInAccountWarningInterval(v string)`

SetNeverLoggedInAccountWarningInterval sets NeverLoggedInAccountWarningInterval field to given value.

### HasNeverLoggedInAccountWarningInterval

`func (o *AddDataSecurityAuditor200Response) HasNeverLoggedInAccountWarningInterval() bool`

HasNeverLoggedInAccountWarningInterval returns a boolean if a field has been set.

### GetNeverLoggedInAccountErrorInterval

`func (o *AddDataSecurityAuditor200Response) GetNeverLoggedInAccountErrorInterval() string`

GetNeverLoggedInAccountErrorInterval returns the NeverLoggedInAccountErrorInterval field if non-nil, zero value otherwise.

### GetNeverLoggedInAccountErrorIntervalOk

`func (o *AddDataSecurityAuditor200Response) GetNeverLoggedInAccountErrorIntervalOk() (*string, bool)`

GetNeverLoggedInAccountErrorIntervalOk returns a tuple with the NeverLoggedInAccountErrorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverLoggedInAccountErrorInterval

`func (o *AddDataSecurityAuditor200Response) SetNeverLoggedInAccountErrorInterval(v string)`

SetNeverLoggedInAccountErrorInterval sets NeverLoggedInAccountErrorInterval field to given value.

### HasNeverLoggedInAccountErrorInterval

`func (o *AddDataSecurityAuditor200Response) HasNeverLoggedInAccountErrorInterval() bool`

HasNeverLoggedInAccountErrorInterval returns a boolean if a field has been set.

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

### GetFilter

`func (o *AddDataSecurityAuditor200Response) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddDataSecurityAuditor200Response) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddDataSecurityAuditor200Response) SetFilter(v []string)`

SetFilter sets Filter field to given value.


### GetAccountExpirationWarningInterval

`func (o *AddDataSecurityAuditor200Response) GetAccountExpirationWarningInterval() string`

GetAccountExpirationWarningInterval returns the AccountExpirationWarningInterval field if non-nil, zero value otherwise.

### GetAccountExpirationWarningIntervalOk

`func (o *AddDataSecurityAuditor200Response) GetAccountExpirationWarningIntervalOk() (*string, bool)`

GetAccountExpirationWarningIntervalOk returns a tuple with the AccountExpirationWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationWarningInterval

`func (o *AddDataSecurityAuditor200Response) SetAccountExpirationWarningInterval(v string)`

SetAccountExpirationWarningInterval sets AccountExpirationWarningInterval field to given value.

### HasAccountExpirationWarningInterval

`func (o *AddDataSecurityAuditor200Response) HasAccountExpirationWarningInterval() bool`

HasAccountExpirationWarningInterval returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddDataSecurityAuditor200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddDataSecurityAuditor200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddDataSecurityAuditor200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddDataSecurityAuditor200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddDataSecurityAuditor200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddDataSecurityAuditor200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddDataSecurityAuditor200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


