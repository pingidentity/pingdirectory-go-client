# GenerateServerProfileRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgenerateServerProfileRecurringTaskSchemaUrn**](EnumgenerateServerProfileRecurringTaskSchemaUrn.md) |  | 
**ProfileDirectory** | **string** | The directory in which the generated server profiles will be placed. The files will be named with the pattern \&quot;server-profile-{timestamp}.zip\&quot;, where \&quot;{timestamp}\&quot; represents the time that the profile was generated. | 
**IncludePath** | Pointer to **[]string** |  | [optional] 
**RetainPreviousProfileCount** | Pointer to **int32** | The minimum number of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**RetainPreviousProfileAge** | Pointer to **string** | The minimum age of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewGenerateServerProfileRecurringTaskShared

`func NewGenerateServerProfileRecurringTaskShared(schemas []EnumgenerateServerProfileRecurringTaskSchemaUrn, profileDirectory string, ) *GenerateServerProfileRecurringTaskShared`

NewGenerateServerProfileRecurringTaskShared instantiates a new GenerateServerProfileRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateServerProfileRecurringTaskSharedWithDefaults

`func NewGenerateServerProfileRecurringTaskSharedWithDefaults() *GenerateServerProfileRecurringTaskShared`

NewGenerateServerProfileRecurringTaskSharedWithDefaults instantiates a new GenerateServerProfileRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GenerateServerProfileRecurringTaskShared) GetSchemas() []EnumgenerateServerProfileRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GenerateServerProfileRecurringTaskShared) GetSchemasOk() (*[]EnumgenerateServerProfileRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GenerateServerProfileRecurringTaskShared) SetSchemas(v []EnumgenerateServerProfileRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetProfileDirectory

`func (o *GenerateServerProfileRecurringTaskShared) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *GenerateServerProfileRecurringTaskShared) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *GenerateServerProfileRecurringTaskShared) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetIncludePath

`func (o *GenerateServerProfileRecurringTaskShared) GetIncludePath() []string`

GetIncludePath returns the IncludePath field if non-nil, zero value otherwise.

### GetIncludePathOk

`func (o *GenerateServerProfileRecurringTaskShared) GetIncludePathOk() (*[]string, bool)`

GetIncludePathOk returns a tuple with the IncludePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePath

`func (o *GenerateServerProfileRecurringTaskShared) SetIncludePath(v []string)`

SetIncludePath sets IncludePath field to given value.

### HasIncludePath

`func (o *GenerateServerProfileRecurringTaskShared) HasIncludePath() bool`

HasIncludePath returns a boolean if a field has been set.

### GetRetainPreviousProfileCount

`func (o *GenerateServerProfileRecurringTaskShared) GetRetainPreviousProfileCount() int32`

GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field if non-nil, zero value otherwise.

### GetRetainPreviousProfileCountOk

`func (o *GenerateServerProfileRecurringTaskShared) GetRetainPreviousProfileCountOk() (*int32, bool)`

GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileCount

`func (o *GenerateServerProfileRecurringTaskShared) SetRetainPreviousProfileCount(v int32)`

SetRetainPreviousProfileCount sets RetainPreviousProfileCount field to given value.

### HasRetainPreviousProfileCount

`func (o *GenerateServerProfileRecurringTaskShared) HasRetainPreviousProfileCount() bool`

HasRetainPreviousProfileCount returns a boolean if a field has been set.

### GetRetainPreviousProfileAge

`func (o *GenerateServerProfileRecurringTaskShared) GetRetainPreviousProfileAge() string`

GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field if non-nil, zero value otherwise.

### GetRetainPreviousProfileAgeOk

`func (o *GenerateServerProfileRecurringTaskShared) GetRetainPreviousProfileAgeOk() (*string, bool)`

GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileAge

`func (o *GenerateServerProfileRecurringTaskShared) SetRetainPreviousProfileAge(v string)`

SetRetainPreviousProfileAge sets RetainPreviousProfileAge field to given value.

### HasRetainPreviousProfileAge

`func (o *GenerateServerProfileRecurringTaskShared) HasRetainPreviousProfileAge() bool`

HasRetainPreviousProfileAge returns a boolean if a field has been set.

### GetDescription

`func (o *GenerateServerProfileRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenerateServerProfileRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenerateServerProfileRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenerateServerProfileRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *GenerateServerProfileRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *GenerateServerProfileRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *GenerateServerProfileRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *GenerateServerProfileRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *GenerateServerProfileRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *GenerateServerProfileRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *GenerateServerProfileRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *GenerateServerProfileRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *GenerateServerProfileRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *GenerateServerProfileRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *GenerateServerProfileRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *GenerateServerProfileRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *GenerateServerProfileRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *GenerateServerProfileRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *GenerateServerProfileRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *GenerateServerProfileRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *GenerateServerProfileRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *GenerateServerProfileRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *GenerateServerProfileRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *GenerateServerProfileRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *GenerateServerProfileRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *GenerateServerProfileRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *GenerateServerProfileRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *GenerateServerProfileRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *GenerateServerProfileRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *GenerateServerProfileRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *GenerateServerProfileRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *GenerateServerProfileRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


