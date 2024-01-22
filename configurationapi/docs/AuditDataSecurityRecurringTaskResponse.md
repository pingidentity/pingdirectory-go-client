# AuditDataSecurityRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumauditDataSecurityRecurringTaskSchemaUrn**](EnumauditDataSecurityRecurringTaskSchemaUrn.md) |  | 
**BaseOutputDirectory** | **string** | The base directory below which generated reports will be written. Each invocation of the audit-data-security task will create a new subdirectory below this base directory whose name is a timestamp indicating when the report was generated. | 
**DataSecurityAuditor** | Pointer to **[]string** | The set of data security auditors that should be invoked. If no auditors are specified, then all auditors defined in the configuration will be used. | [optional] 
**Backend** | Pointer to **[]string** | The set of backends that should be examined. If no backends are specified, then all backends that support this functionality will be included. | [optional] 
**IncludeFilter** | Pointer to **[]string** | A filter that will be used to identify entries that may be included in the generated report. If multiple filters are specified, then any entry that matches at least one of the filters will be included. If no filters are specified, then all entries will be included. | [optional] 
**RetainPreviousReportCount** | Pointer to **int64** | The minimum number of previous reports that should be preserved after a new report is generated. | [optional] 
**RetainPreviousReportAge** | Pointer to **string** | The minimum age of previous reports that should be preserved after a new report completes successfully. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Recurring Task | 

## Methods

### NewAuditDataSecurityRecurringTaskResponse

`func NewAuditDataSecurityRecurringTaskResponse(schemas []EnumauditDataSecurityRecurringTaskSchemaUrn, baseOutputDirectory string, id string, ) *AuditDataSecurityRecurringTaskResponse`

NewAuditDataSecurityRecurringTaskResponse instantiates a new AuditDataSecurityRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditDataSecurityRecurringTaskResponseWithDefaults

`func NewAuditDataSecurityRecurringTaskResponseWithDefaults() *AuditDataSecurityRecurringTaskResponse`

NewAuditDataSecurityRecurringTaskResponseWithDefaults instantiates a new AuditDataSecurityRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AuditDataSecurityRecurringTaskResponse) GetSchemas() []EnumauditDataSecurityRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetSchemasOk() (*[]EnumauditDataSecurityRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AuditDataSecurityRecurringTaskResponse) SetSchemas(v []EnumauditDataSecurityRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseOutputDirectory

`func (o *AuditDataSecurityRecurringTaskResponse) GetBaseOutputDirectory() string`

GetBaseOutputDirectory returns the BaseOutputDirectory field if non-nil, zero value otherwise.

### GetBaseOutputDirectoryOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetBaseOutputDirectoryOk() (*string, bool)`

GetBaseOutputDirectoryOk returns a tuple with the BaseOutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOutputDirectory

`func (o *AuditDataSecurityRecurringTaskResponse) SetBaseOutputDirectory(v string)`

SetBaseOutputDirectory sets BaseOutputDirectory field to given value.


### GetDataSecurityAuditor

`func (o *AuditDataSecurityRecurringTaskResponse) GetDataSecurityAuditor() []string`

GetDataSecurityAuditor returns the DataSecurityAuditor field if non-nil, zero value otherwise.

### GetDataSecurityAuditorOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetDataSecurityAuditorOk() (*[]string, bool)`

GetDataSecurityAuditorOk returns a tuple with the DataSecurityAuditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSecurityAuditor

`func (o *AuditDataSecurityRecurringTaskResponse) SetDataSecurityAuditor(v []string)`

SetDataSecurityAuditor sets DataSecurityAuditor field to given value.

### HasDataSecurityAuditor

`func (o *AuditDataSecurityRecurringTaskResponse) HasDataSecurityAuditor() bool`

HasDataSecurityAuditor returns a boolean if a field has been set.

### GetBackend

`func (o *AuditDataSecurityRecurringTaskResponse) GetBackend() []string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetBackendOk() (*[]string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *AuditDataSecurityRecurringTaskResponse) SetBackend(v []string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *AuditDataSecurityRecurringTaskResponse) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AuditDataSecurityRecurringTaskResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AuditDataSecurityRecurringTaskResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AuditDataSecurityRecurringTaskResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetRetainPreviousReportCount

`func (o *AuditDataSecurityRecurringTaskResponse) GetRetainPreviousReportCount() int64`

GetRetainPreviousReportCount returns the RetainPreviousReportCount field if non-nil, zero value otherwise.

### GetRetainPreviousReportCountOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetRetainPreviousReportCountOk() (*int64, bool)`

GetRetainPreviousReportCountOk returns a tuple with the RetainPreviousReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportCount

`func (o *AuditDataSecurityRecurringTaskResponse) SetRetainPreviousReportCount(v int64)`

SetRetainPreviousReportCount sets RetainPreviousReportCount field to given value.

### HasRetainPreviousReportCount

`func (o *AuditDataSecurityRecurringTaskResponse) HasRetainPreviousReportCount() bool`

HasRetainPreviousReportCount returns a boolean if a field has been set.

### GetRetainPreviousReportAge

`func (o *AuditDataSecurityRecurringTaskResponse) GetRetainPreviousReportAge() string`

GetRetainPreviousReportAge returns the RetainPreviousReportAge field if non-nil, zero value otherwise.

### GetRetainPreviousReportAgeOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetRetainPreviousReportAgeOk() (*string, bool)`

GetRetainPreviousReportAgeOk returns a tuple with the RetainPreviousReportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportAge

`func (o *AuditDataSecurityRecurringTaskResponse) SetRetainPreviousReportAge(v string)`

SetRetainPreviousReportAge sets RetainPreviousReportAge field to given value.

### HasRetainPreviousReportAge

`func (o *AuditDataSecurityRecurringTaskResponse) HasRetainPreviousReportAge() bool`

HasRetainPreviousReportAge returns a boolean if a field has been set.

### GetDescription

`func (o *AuditDataSecurityRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditDataSecurityRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditDataSecurityRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AuditDataSecurityRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AuditDataSecurityRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AuditDataSecurityRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AuditDataSecurityRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AuditDataSecurityRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AuditDataSecurityRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AuditDataSecurityRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AuditDataSecurityRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AuditDataSecurityRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AuditDataSecurityRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AuditDataSecurityRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AuditDataSecurityRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AuditDataSecurityRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AuditDataSecurityRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AuditDataSecurityRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AuditDataSecurityRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AuditDataSecurityRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AuditDataSecurityRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AuditDataSecurityRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AuditDataSecurityRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AuditDataSecurityRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetMeta

`func (o *AuditDataSecurityRecurringTaskResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuditDataSecurityRecurringTaskResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuditDataSecurityRecurringTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AuditDataSecurityRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AuditDataSecurityRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AuditDataSecurityRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AuditDataSecurityRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AuditDataSecurityRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditDataSecurityRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditDataSecurityRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


