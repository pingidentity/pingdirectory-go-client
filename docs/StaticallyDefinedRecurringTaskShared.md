# StaticallyDefinedRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumstaticallyDefinedRecurringTaskSchemaUrn**](EnumstaticallyDefinedRecurringTaskSchemaUrn.md) |  | 
**TaskJavaClass** | **string** | The fully-qualified name of the Java class that provides the logic for the task to be invoked. | 
**TaskObjectClass** | **[]string** |  | 
**TaskAttributeValue** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewStaticallyDefinedRecurringTaskShared

`func NewStaticallyDefinedRecurringTaskShared(schemas []EnumstaticallyDefinedRecurringTaskSchemaUrn, taskJavaClass string, taskObjectClass []string, ) *StaticallyDefinedRecurringTaskShared`

NewStaticallyDefinedRecurringTaskShared instantiates a new StaticallyDefinedRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticallyDefinedRecurringTaskSharedWithDefaults

`func NewStaticallyDefinedRecurringTaskSharedWithDefaults() *StaticallyDefinedRecurringTaskShared`

NewStaticallyDefinedRecurringTaskSharedWithDefaults instantiates a new StaticallyDefinedRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *StaticallyDefinedRecurringTaskShared) GetSchemas() []EnumstaticallyDefinedRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StaticallyDefinedRecurringTaskShared) GetSchemasOk() (*[]EnumstaticallyDefinedRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StaticallyDefinedRecurringTaskShared) SetSchemas(v []EnumstaticallyDefinedRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTaskJavaClass

`func (o *StaticallyDefinedRecurringTaskShared) GetTaskJavaClass() string`

GetTaskJavaClass returns the TaskJavaClass field if non-nil, zero value otherwise.

### GetTaskJavaClassOk

`func (o *StaticallyDefinedRecurringTaskShared) GetTaskJavaClassOk() (*string, bool)`

GetTaskJavaClassOk returns a tuple with the TaskJavaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJavaClass

`func (o *StaticallyDefinedRecurringTaskShared) SetTaskJavaClass(v string)`

SetTaskJavaClass sets TaskJavaClass field to given value.


### GetTaskObjectClass

`func (o *StaticallyDefinedRecurringTaskShared) GetTaskObjectClass() []string`

GetTaskObjectClass returns the TaskObjectClass field if non-nil, zero value otherwise.

### GetTaskObjectClassOk

`func (o *StaticallyDefinedRecurringTaskShared) GetTaskObjectClassOk() (*[]string, bool)`

GetTaskObjectClassOk returns a tuple with the TaskObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskObjectClass

`func (o *StaticallyDefinedRecurringTaskShared) SetTaskObjectClass(v []string)`

SetTaskObjectClass sets TaskObjectClass field to given value.


### GetTaskAttributeValue

`func (o *StaticallyDefinedRecurringTaskShared) GetTaskAttributeValue() []string`

GetTaskAttributeValue returns the TaskAttributeValue field if non-nil, zero value otherwise.

### GetTaskAttributeValueOk

`func (o *StaticallyDefinedRecurringTaskShared) GetTaskAttributeValueOk() (*[]string, bool)`

GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttributeValue

`func (o *StaticallyDefinedRecurringTaskShared) SetTaskAttributeValue(v []string)`

SetTaskAttributeValue sets TaskAttributeValue field to given value.

### HasTaskAttributeValue

`func (o *StaticallyDefinedRecurringTaskShared) HasTaskAttributeValue() bool`

HasTaskAttributeValue returns a boolean if a field has been set.

### GetDescription

`func (o *StaticallyDefinedRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticallyDefinedRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticallyDefinedRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticallyDefinedRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *StaticallyDefinedRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *StaticallyDefinedRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *StaticallyDefinedRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *StaticallyDefinedRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *StaticallyDefinedRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *StaticallyDefinedRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *StaticallyDefinedRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *StaticallyDefinedRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *StaticallyDefinedRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *StaticallyDefinedRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *StaticallyDefinedRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *StaticallyDefinedRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *StaticallyDefinedRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *StaticallyDefinedRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *StaticallyDefinedRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *StaticallyDefinedRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *StaticallyDefinedRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *StaticallyDefinedRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *StaticallyDefinedRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *StaticallyDefinedRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *StaticallyDefinedRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *StaticallyDefinedRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *StaticallyDefinedRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *StaticallyDefinedRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *StaticallyDefinedRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *StaticallyDefinedRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *StaticallyDefinedRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *StaticallyDefinedRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


