# StaticallyDefinedRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Recurring Task | 
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

### NewStaticallyDefinedRecurringTaskResponse

`func NewStaticallyDefinedRecurringTaskResponse(id string, schemas []EnumstaticallyDefinedRecurringTaskSchemaUrn, taskJavaClass string, taskObjectClass []string, ) *StaticallyDefinedRecurringTaskResponse`

NewStaticallyDefinedRecurringTaskResponse instantiates a new StaticallyDefinedRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticallyDefinedRecurringTaskResponseWithDefaults

`func NewStaticallyDefinedRecurringTaskResponseWithDefaults() *StaticallyDefinedRecurringTaskResponse`

NewStaticallyDefinedRecurringTaskResponseWithDefaults instantiates a new StaticallyDefinedRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StaticallyDefinedRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticallyDefinedRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *StaticallyDefinedRecurringTaskResponse) GetSchemas() []EnumstaticallyDefinedRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetSchemasOk() (*[]EnumstaticallyDefinedRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StaticallyDefinedRecurringTaskResponse) SetSchemas(v []EnumstaticallyDefinedRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTaskJavaClass

`func (o *StaticallyDefinedRecurringTaskResponse) GetTaskJavaClass() string`

GetTaskJavaClass returns the TaskJavaClass field if non-nil, zero value otherwise.

### GetTaskJavaClassOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetTaskJavaClassOk() (*string, bool)`

GetTaskJavaClassOk returns a tuple with the TaskJavaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJavaClass

`func (o *StaticallyDefinedRecurringTaskResponse) SetTaskJavaClass(v string)`

SetTaskJavaClass sets TaskJavaClass field to given value.


### GetTaskObjectClass

`func (o *StaticallyDefinedRecurringTaskResponse) GetTaskObjectClass() []string`

GetTaskObjectClass returns the TaskObjectClass field if non-nil, zero value otherwise.

### GetTaskObjectClassOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetTaskObjectClassOk() (*[]string, bool)`

GetTaskObjectClassOk returns a tuple with the TaskObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskObjectClass

`func (o *StaticallyDefinedRecurringTaskResponse) SetTaskObjectClass(v []string)`

SetTaskObjectClass sets TaskObjectClass field to given value.


### GetTaskAttributeValue

`func (o *StaticallyDefinedRecurringTaskResponse) GetTaskAttributeValue() []string`

GetTaskAttributeValue returns the TaskAttributeValue field if non-nil, zero value otherwise.

### GetTaskAttributeValueOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetTaskAttributeValueOk() (*[]string, bool)`

GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttributeValue

`func (o *StaticallyDefinedRecurringTaskResponse) SetTaskAttributeValue(v []string)`

SetTaskAttributeValue sets TaskAttributeValue field to given value.

### HasTaskAttributeValue

`func (o *StaticallyDefinedRecurringTaskResponse) HasTaskAttributeValue() bool`

HasTaskAttributeValue returns a boolean if a field has been set.

### GetDescription

`func (o *StaticallyDefinedRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticallyDefinedRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticallyDefinedRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *StaticallyDefinedRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *StaticallyDefinedRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *StaticallyDefinedRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *StaticallyDefinedRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *StaticallyDefinedRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *StaticallyDefinedRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *StaticallyDefinedRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *StaticallyDefinedRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *StaticallyDefinedRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *StaticallyDefinedRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *StaticallyDefinedRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *StaticallyDefinedRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *StaticallyDefinedRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *StaticallyDefinedRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *StaticallyDefinedRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


