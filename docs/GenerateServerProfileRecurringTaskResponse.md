# GenerateServerProfileRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Recurring Task | 
**Schemas** | [**[]EnumgenerateServerProfileRecurringTaskSchemaUrn**](EnumgenerateServerProfileRecurringTaskSchemaUrn.md) |  | 
**ProfileDirectory** | **string** | The directory in which the generated server profiles will be placed. The files will be named with the pattern \&quot;server-profile-{timestamp}.zip\&quot;, where \&quot;{timestamp}\&quot; represents the time that the profile was generated. | 
**IncludePath** | Pointer to **[]string** | An optional set of additional paths to files within the instance root that should be included in the generated server profile. All paths must be within the instance root, and relative paths will be relative to the instance root. | [optional] 
**RetainPreviousProfileCount** | Pointer to **int32** | The minimum number of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**RetainPreviousProfileAge** | Pointer to **string** | The minimum age of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewGenerateServerProfileRecurringTaskResponse

`func NewGenerateServerProfileRecurringTaskResponse(id string, schemas []EnumgenerateServerProfileRecurringTaskSchemaUrn, profileDirectory string, ) *GenerateServerProfileRecurringTaskResponse`

NewGenerateServerProfileRecurringTaskResponse instantiates a new GenerateServerProfileRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateServerProfileRecurringTaskResponseWithDefaults

`func NewGenerateServerProfileRecurringTaskResponseWithDefaults() *GenerateServerProfileRecurringTaskResponse`

NewGenerateServerProfileRecurringTaskResponseWithDefaults instantiates a new GenerateServerProfileRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenerateServerProfileRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenerateServerProfileRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GenerateServerProfileRecurringTaskResponse) GetSchemas() []EnumgenerateServerProfileRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetSchemasOk() (*[]EnumgenerateServerProfileRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GenerateServerProfileRecurringTaskResponse) SetSchemas(v []EnumgenerateServerProfileRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetProfileDirectory

`func (o *GenerateServerProfileRecurringTaskResponse) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *GenerateServerProfileRecurringTaskResponse) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetIncludePath

`func (o *GenerateServerProfileRecurringTaskResponse) GetIncludePath() []string`

GetIncludePath returns the IncludePath field if non-nil, zero value otherwise.

### GetIncludePathOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetIncludePathOk() (*[]string, bool)`

GetIncludePathOk returns a tuple with the IncludePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePath

`func (o *GenerateServerProfileRecurringTaskResponse) SetIncludePath(v []string)`

SetIncludePath sets IncludePath field to given value.

### HasIncludePath

`func (o *GenerateServerProfileRecurringTaskResponse) HasIncludePath() bool`

HasIncludePath returns a boolean if a field has been set.

### GetRetainPreviousProfileCount

`func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileCount() int32`

GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field if non-nil, zero value otherwise.

### GetRetainPreviousProfileCountOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileCountOk() (*int32, bool)`

GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileCount

`func (o *GenerateServerProfileRecurringTaskResponse) SetRetainPreviousProfileCount(v int32)`

SetRetainPreviousProfileCount sets RetainPreviousProfileCount field to given value.

### HasRetainPreviousProfileCount

`func (o *GenerateServerProfileRecurringTaskResponse) HasRetainPreviousProfileCount() bool`

HasRetainPreviousProfileCount returns a boolean if a field has been set.

### GetRetainPreviousProfileAge

`func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileAge() string`

GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field if non-nil, zero value otherwise.

### GetRetainPreviousProfileAgeOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileAgeOk() (*string, bool)`

GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileAge

`func (o *GenerateServerProfileRecurringTaskResponse) SetRetainPreviousProfileAge(v string)`

SetRetainPreviousProfileAge sets RetainPreviousProfileAge field to given value.

### HasRetainPreviousProfileAge

`func (o *GenerateServerProfileRecurringTaskResponse) HasRetainPreviousProfileAge() bool`

HasRetainPreviousProfileAge returns a boolean if a field has been set.

### GetDescription

`func (o *GenerateServerProfileRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenerateServerProfileRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenerateServerProfileRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *GenerateServerProfileRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *GenerateServerProfileRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *GenerateServerProfileRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *GenerateServerProfileRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *GenerateServerProfileRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *GenerateServerProfileRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *GenerateServerProfileRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *GenerateServerProfileRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *GenerateServerProfileRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *GenerateServerProfileRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *GenerateServerProfileRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *GenerateServerProfileRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *GenerateServerProfileRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *GenerateServerProfileRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *GenerateServerProfileRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetMeta

`func (o *GenerateServerProfileRecurringTaskResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenerateServerProfileRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenerateServerProfileRecurringTaskResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenerateServerProfileRecurringTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


