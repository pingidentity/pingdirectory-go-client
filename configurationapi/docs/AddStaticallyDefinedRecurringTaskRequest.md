# AddStaticallyDefinedRecurringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumstaticallyDefinedRecurringTaskSchemaUrn**](EnumstaticallyDefinedRecurringTaskSchemaUrn.md) |  | 
**TaskJavaClass** | **string** | The fully-qualified name of the Java class that provides the logic for the task to be invoked. | 
**TaskObjectClass** | **[]string** | The names or OIDs of the object classes to include in the tasks that are scheduled from this Statically Defined Recurring Task. All object classes must be defined in the server schema, and the combination of object classes must be valid for a task entry. | 
**TaskAttributeValue** | Pointer to **[]string** | The set of attribute values that should be included in the tasks that are scheduled from this Statically Defined Recurring Task. Each value must be in the form {attribute-type}&#x3D;{value}, where {attribute-type} is the name or OID of an attribute type that is defined in the schema and permitted with the configured set of object classes, and {value} is a value to assign to an attribute with that type. A multivalued attribute can be created by providing multiple name-value pairs with the same name and different values. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**TaskName** | **string** | Name of the new Recurring Task | 

## Methods

### NewAddStaticallyDefinedRecurringTaskRequest

`func NewAddStaticallyDefinedRecurringTaskRequest(schemas []EnumstaticallyDefinedRecurringTaskSchemaUrn, taskJavaClass string, taskObjectClass []string, taskName string, ) *AddStaticallyDefinedRecurringTaskRequest`

NewAddStaticallyDefinedRecurringTaskRequest instantiates a new AddStaticallyDefinedRecurringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStaticallyDefinedRecurringTaskRequestWithDefaults

`func NewAddStaticallyDefinedRecurringTaskRequestWithDefaults() *AddStaticallyDefinedRecurringTaskRequest`

NewAddStaticallyDefinedRecurringTaskRequestWithDefaults instantiates a new AddStaticallyDefinedRecurringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetSchemas() []EnumstaticallyDefinedRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetSchemasOk() (*[]EnumstaticallyDefinedRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetSchemas(v []EnumstaticallyDefinedRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTaskJavaClass

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskJavaClass() string`

GetTaskJavaClass returns the TaskJavaClass field if non-nil, zero value otherwise.

### GetTaskJavaClassOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskJavaClassOk() (*string, bool)`

GetTaskJavaClassOk returns a tuple with the TaskJavaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJavaClass

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetTaskJavaClass(v string)`

SetTaskJavaClass sets TaskJavaClass field to given value.


### GetTaskObjectClass

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskObjectClass() []string`

GetTaskObjectClass returns the TaskObjectClass field if non-nil, zero value otherwise.

### GetTaskObjectClassOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskObjectClassOk() (*[]string, bool)`

GetTaskObjectClassOk returns a tuple with the TaskObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskObjectClass

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetTaskObjectClass(v []string)`

SetTaskObjectClass sets TaskObjectClass field to given value.


### GetTaskAttributeValue

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskAttributeValue() []string`

GetTaskAttributeValue returns the TaskAttributeValue field if non-nil, zero value otherwise.

### GetTaskAttributeValueOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskAttributeValueOk() (*[]string, bool)`

GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttributeValue

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetTaskAttributeValue(v []string)`

SetTaskAttributeValue sets TaskAttributeValue field to given value.

### HasTaskAttributeValue

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasTaskAttributeValue() bool`

HasTaskAttributeValue returns a boolean if a field has been set.

### GetDescription

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddStaticallyDefinedRecurringTaskRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetTaskName

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AddStaticallyDefinedRecurringTaskRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AddStaticallyDefinedRecurringTaskRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


