# AddAuditDataSecurityRecurringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** | Name of the new Recurring Task | 
**Schemas** | [**[]EnumauditDataSecurityRecurringTaskSchemaUrn**](EnumauditDataSecurityRecurringTaskSchemaUrn.md) |  | 
**BaseOutputDirectory** | Pointer to **string** | The base directory below which generated reports will be written. Each invocation of the audit-data-security task will create a new subdirectory below this base directory whose name is a timestamp indicating when the report was generated. | [optional] 
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

## Methods

### NewAddAuditDataSecurityRecurringTaskRequest

`func NewAddAuditDataSecurityRecurringTaskRequest(taskName string, schemas []EnumauditDataSecurityRecurringTaskSchemaUrn, ) *AddAuditDataSecurityRecurringTaskRequest`

NewAddAuditDataSecurityRecurringTaskRequest instantiates a new AddAuditDataSecurityRecurringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAuditDataSecurityRecurringTaskRequestWithDefaults

`func NewAddAuditDataSecurityRecurringTaskRequestWithDefaults() *AddAuditDataSecurityRecurringTaskRequest`

NewAddAuditDataSecurityRecurringTaskRequestWithDefaults instantiates a new AddAuditDataSecurityRecurringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetSchemas

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetSchemas() []EnumauditDataSecurityRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetSchemasOk() (*[]EnumauditDataSecurityRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetSchemas(v []EnumauditDataSecurityRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseOutputDirectory

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetBaseOutputDirectory() string`

GetBaseOutputDirectory returns the BaseOutputDirectory field if non-nil, zero value otherwise.

### GetBaseOutputDirectoryOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetBaseOutputDirectoryOk() (*string, bool)`

GetBaseOutputDirectoryOk returns a tuple with the BaseOutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOutputDirectory

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetBaseOutputDirectory(v string)`

SetBaseOutputDirectory sets BaseOutputDirectory field to given value.

### HasBaseOutputDirectory

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasBaseOutputDirectory() bool`

HasBaseOutputDirectory returns a boolean if a field has been set.

### GetDataSecurityAuditor

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetDataSecurityAuditor() []string`

GetDataSecurityAuditor returns the DataSecurityAuditor field if non-nil, zero value otherwise.

### GetDataSecurityAuditorOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetDataSecurityAuditorOk() (*[]string, bool)`

GetDataSecurityAuditorOk returns a tuple with the DataSecurityAuditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSecurityAuditor

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetDataSecurityAuditor(v []string)`

SetDataSecurityAuditor sets DataSecurityAuditor field to given value.

### HasDataSecurityAuditor

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasDataSecurityAuditor() bool`

HasDataSecurityAuditor returns a boolean if a field has been set.

### GetBackend

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetBackend() []string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetBackendOk() (*[]string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetBackend(v []string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetRetainPreviousReportCount

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetRetainPreviousReportCount() int64`

GetRetainPreviousReportCount returns the RetainPreviousReportCount field if non-nil, zero value otherwise.

### GetRetainPreviousReportCountOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetRetainPreviousReportCountOk() (*int64, bool)`

GetRetainPreviousReportCountOk returns a tuple with the RetainPreviousReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportCount

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetRetainPreviousReportCount(v int64)`

SetRetainPreviousReportCount sets RetainPreviousReportCount field to given value.

### HasRetainPreviousReportCount

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasRetainPreviousReportCount() bool`

HasRetainPreviousReportCount returns a boolean if a field has been set.

### GetRetainPreviousReportAge

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetRetainPreviousReportAge() string`

GetRetainPreviousReportAge returns the RetainPreviousReportAge field if non-nil, zero value otherwise.

### GetRetainPreviousReportAgeOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetRetainPreviousReportAgeOk() (*string, bool)`

GetRetainPreviousReportAgeOk returns a tuple with the RetainPreviousReportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportAge

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetRetainPreviousReportAge(v string)`

SetRetainPreviousReportAge sets RetainPreviousReportAge field to given value.

### HasRetainPreviousReportAge

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasRetainPreviousReportAge() bool`

HasRetainPreviousReportAge returns a boolean if a field has been set.

### GetDescription

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddAuditDataSecurityRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddAuditDataSecurityRecurringTaskRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


