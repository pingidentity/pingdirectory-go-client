# ThirdPartyRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyRecurringTaskSchemaUrn**](EnumthirdPartyRecurringTaskSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Recurring Task. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewThirdPartyRecurringTaskShared

`func NewThirdPartyRecurringTaskShared(schemas []EnumthirdPartyRecurringTaskSchemaUrn, extensionClass string, ) *ThirdPartyRecurringTaskShared`

NewThirdPartyRecurringTaskShared instantiates a new ThirdPartyRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyRecurringTaskSharedWithDefaults

`func NewThirdPartyRecurringTaskSharedWithDefaults() *ThirdPartyRecurringTaskShared`

NewThirdPartyRecurringTaskSharedWithDefaults instantiates a new ThirdPartyRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyRecurringTaskShared) GetSchemas() []EnumthirdPartyRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyRecurringTaskShared) GetSchemasOk() (*[]EnumthirdPartyRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyRecurringTaskShared) SetSchemas(v []EnumthirdPartyRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyRecurringTaskShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyRecurringTaskShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyRecurringTaskShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyRecurringTaskShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyRecurringTaskShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyRecurringTaskShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyRecurringTaskShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *ThirdPartyRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *ThirdPartyRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *ThirdPartyRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *ThirdPartyRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *ThirdPartyRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *ThirdPartyRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *ThirdPartyRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *ThirdPartyRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *ThirdPartyRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *ThirdPartyRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *ThirdPartyRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *ThirdPartyRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *ThirdPartyRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *ThirdPartyRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *ThirdPartyRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *ThirdPartyRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *ThirdPartyRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *ThirdPartyRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *ThirdPartyRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *ThirdPartyRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *ThirdPartyRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *ThirdPartyRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *ThirdPartyRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *ThirdPartyRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *ThirdPartyRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *ThirdPartyRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *ThirdPartyRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *ThirdPartyRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


