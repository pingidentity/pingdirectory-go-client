# AddLeaveLockdownModeRecurringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** | Name of the new Recurring Task | 
**Schemas** | [**[]EnumleaveLockdownModeRecurringTaskSchemaUrn**](EnumleaveLockdownModeRecurringTaskSchemaUrn.md) |  | 
**Reason** | Pointer to **string** | The reason that the server is being taken out of in lockdown mode. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewAddLeaveLockdownModeRecurringTaskRequest

`func NewAddLeaveLockdownModeRecurringTaskRequest(taskName string, schemas []EnumleaveLockdownModeRecurringTaskSchemaUrn, ) *AddLeaveLockdownModeRecurringTaskRequest`

NewAddLeaveLockdownModeRecurringTaskRequest instantiates a new AddLeaveLockdownModeRecurringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLeaveLockdownModeRecurringTaskRequestWithDefaults

`func NewAddLeaveLockdownModeRecurringTaskRequestWithDefaults() *AddLeaveLockdownModeRecurringTaskRequest`

NewAddLeaveLockdownModeRecurringTaskRequestWithDefaults instantiates a new AddLeaveLockdownModeRecurringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetSchemas

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetSchemas() []EnumleaveLockdownModeRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetSchemasOk() (*[]EnumleaveLockdownModeRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetSchemas(v []EnumleaveLockdownModeRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReason

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDescription

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddLeaveLockdownModeRecurringTaskRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


