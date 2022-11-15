# LeaveLockdownModeRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Recurring Task | 
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

### NewLeaveLockdownModeRecurringTaskResponse

`func NewLeaveLockdownModeRecurringTaskResponse(id string, schemas []EnumleaveLockdownModeRecurringTaskSchemaUrn, ) *LeaveLockdownModeRecurringTaskResponse`

NewLeaveLockdownModeRecurringTaskResponse instantiates a new LeaveLockdownModeRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaveLockdownModeRecurringTaskResponseWithDefaults

`func NewLeaveLockdownModeRecurringTaskResponseWithDefaults() *LeaveLockdownModeRecurringTaskResponse`

NewLeaveLockdownModeRecurringTaskResponseWithDefaults instantiates a new LeaveLockdownModeRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LeaveLockdownModeRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeaveLockdownModeRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LeaveLockdownModeRecurringTaskResponse) GetSchemas() []EnumleaveLockdownModeRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetSchemasOk() (*[]EnumleaveLockdownModeRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LeaveLockdownModeRecurringTaskResponse) SetSchemas(v []EnumleaveLockdownModeRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReason

`func (o *LeaveLockdownModeRecurringTaskResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LeaveLockdownModeRecurringTaskResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LeaveLockdownModeRecurringTaskResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDescription

`func (o *LeaveLockdownModeRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LeaveLockdownModeRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LeaveLockdownModeRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *LeaveLockdownModeRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *LeaveLockdownModeRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *LeaveLockdownModeRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *LeaveLockdownModeRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *LeaveLockdownModeRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *LeaveLockdownModeRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *LeaveLockdownModeRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *LeaveLockdownModeRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *LeaveLockdownModeRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *LeaveLockdownModeRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *LeaveLockdownModeRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *LeaveLockdownModeRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *LeaveLockdownModeRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *LeaveLockdownModeRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


