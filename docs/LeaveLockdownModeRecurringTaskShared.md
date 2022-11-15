# LeaveLockdownModeRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumleaveLockdownModeRecurringTaskSchemaUrn**](EnumleaveLockdownModeRecurringTaskSchemaUrn.md) |  | 
**Reason** | Pointer to **string** | The reason that the server is being taken out of in lockdown mode. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewLeaveLockdownModeRecurringTaskShared

`func NewLeaveLockdownModeRecurringTaskShared(schemas []EnumleaveLockdownModeRecurringTaskSchemaUrn, ) *LeaveLockdownModeRecurringTaskShared`

NewLeaveLockdownModeRecurringTaskShared instantiates a new LeaveLockdownModeRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaveLockdownModeRecurringTaskSharedWithDefaults

`func NewLeaveLockdownModeRecurringTaskSharedWithDefaults() *LeaveLockdownModeRecurringTaskShared`

NewLeaveLockdownModeRecurringTaskSharedWithDefaults instantiates a new LeaveLockdownModeRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LeaveLockdownModeRecurringTaskShared) GetSchemas() []EnumleaveLockdownModeRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetSchemasOk() (*[]EnumleaveLockdownModeRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LeaveLockdownModeRecurringTaskShared) SetSchemas(v []EnumleaveLockdownModeRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReason

`func (o *LeaveLockdownModeRecurringTaskShared) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LeaveLockdownModeRecurringTaskShared) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LeaveLockdownModeRecurringTaskShared) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDescription

`func (o *LeaveLockdownModeRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LeaveLockdownModeRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LeaveLockdownModeRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *LeaveLockdownModeRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *LeaveLockdownModeRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *LeaveLockdownModeRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *LeaveLockdownModeRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *LeaveLockdownModeRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *LeaveLockdownModeRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *LeaveLockdownModeRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *LeaveLockdownModeRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *LeaveLockdownModeRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *LeaveLockdownModeRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *LeaveLockdownModeRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *LeaveLockdownModeRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *LeaveLockdownModeRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *LeaveLockdownModeRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *LeaveLockdownModeRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *LeaveLockdownModeRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *LeaveLockdownModeRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *LeaveLockdownModeRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *LeaveLockdownModeRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *LeaveLockdownModeRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *LeaveLockdownModeRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *LeaveLockdownModeRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


