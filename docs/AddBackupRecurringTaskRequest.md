# AddBackupRecurringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** | Name of the new Recurring Task | 
**Schemas** | [**[]EnumbackupRecurringTaskSchemaUrn**](EnumbackupRecurringTaskSchemaUrn.md) |  | 
**BackupDirectory** | **string** | The directory in which backup files will be placed. When backing up a single backend, the backup files will be placed directly in this directory. When backing up multiple backends, the backup files for each backend will be placed in a subdirectory whose name is the corresponding backend ID. | 
**IncludedBackendID** | Pointer to **[]string** |  | [optional] 
**ExcludedBackendID** | Pointer to **[]string** |  | [optional] 
**Compress** | Pointer to **bool** | Indicates whether to compress the data as it is written into the backup. | [optional] 
**Encrypt** | Pointer to **bool** | Indicates whether to encrypt the data as it is written into the backup. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The ID of an encryption settings definition to use to obtain the backup encryption key. | [optional] 
**Sign** | Pointer to **bool** | Indicates whether to cryptographically sign backups, which will make it possible to detect whether the backup has been altered since it was created. | [optional] 
**RetainPreviousFullBackupCount** | Pointer to **int32** | The minimum number of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**RetainPreviousFullBackupAge** | Pointer to **string** | The minimum age of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**MaxMegabytesPerSecond** | Pointer to **int32** | The maximum rate, in megabytes per second, at which backups should be written. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewAddBackupRecurringTaskRequest

`func NewAddBackupRecurringTaskRequest(taskName string, schemas []EnumbackupRecurringTaskSchemaUrn, backupDirectory string, ) *AddBackupRecurringTaskRequest`

NewAddBackupRecurringTaskRequest instantiates a new AddBackupRecurringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackupRecurringTaskRequestWithDefaults

`func NewAddBackupRecurringTaskRequestWithDefaults() *AddBackupRecurringTaskRequest`

NewAddBackupRecurringTaskRequestWithDefaults instantiates a new AddBackupRecurringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *AddBackupRecurringTaskRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AddBackupRecurringTaskRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AddBackupRecurringTaskRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetSchemas

`func (o *AddBackupRecurringTaskRequest) GetSchemas() []EnumbackupRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddBackupRecurringTaskRequest) GetSchemasOk() (*[]EnumbackupRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddBackupRecurringTaskRequest) SetSchemas(v []EnumbackupRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBackupDirectory

`func (o *AddBackupRecurringTaskRequest) GetBackupDirectory() string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *AddBackupRecurringTaskRequest) GetBackupDirectoryOk() (*string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *AddBackupRecurringTaskRequest) SetBackupDirectory(v string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIncludedBackendID

`func (o *AddBackupRecurringTaskRequest) GetIncludedBackendID() []string`

GetIncludedBackendID returns the IncludedBackendID field if non-nil, zero value otherwise.

### GetIncludedBackendIDOk

`func (o *AddBackupRecurringTaskRequest) GetIncludedBackendIDOk() (*[]string, bool)`

GetIncludedBackendIDOk returns a tuple with the IncludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendID

`func (o *AddBackupRecurringTaskRequest) SetIncludedBackendID(v []string)`

SetIncludedBackendID sets IncludedBackendID field to given value.

### HasIncludedBackendID

`func (o *AddBackupRecurringTaskRequest) HasIncludedBackendID() bool`

HasIncludedBackendID returns a boolean if a field has been set.

### GetExcludedBackendID

`func (o *AddBackupRecurringTaskRequest) GetExcludedBackendID() []string`

GetExcludedBackendID returns the ExcludedBackendID field if non-nil, zero value otherwise.

### GetExcludedBackendIDOk

`func (o *AddBackupRecurringTaskRequest) GetExcludedBackendIDOk() (*[]string, bool)`

GetExcludedBackendIDOk returns a tuple with the ExcludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendID

`func (o *AddBackupRecurringTaskRequest) SetExcludedBackendID(v []string)`

SetExcludedBackendID sets ExcludedBackendID field to given value.

### HasExcludedBackendID

`func (o *AddBackupRecurringTaskRequest) HasExcludedBackendID() bool`

HasExcludedBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *AddBackupRecurringTaskRequest) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *AddBackupRecurringTaskRequest) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *AddBackupRecurringTaskRequest) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *AddBackupRecurringTaskRequest) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *AddBackupRecurringTaskRequest) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *AddBackupRecurringTaskRequest) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *AddBackupRecurringTaskRequest) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *AddBackupRecurringTaskRequest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *AddBackupRecurringTaskRequest) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *AddBackupRecurringTaskRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *AddBackupRecurringTaskRequest) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *AddBackupRecurringTaskRequest) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *AddBackupRecurringTaskRequest) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *AddBackupRecurringTaskRequest) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *AddBackupRecurringTaskRequest) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *AddBackupRecurringTaskRequest) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousFullBackupCount

`func (o *AddBackupRecurringTaskRequest) GetRetainPreviousFullBackupCount() int32`

GetRetainPreviousFullBackupCount returns the RetainPreviousFullBackupCount field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupCountOk

`func (o *AddBackupRecurringTaskRequest) GetRetainPreviousFullBackupCountOk() (*int32, bool)`

GetRetainPreviousFullBackupCountOk returns a tuple with the RetainPreviousFullBackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupCount

`func (o *AddBackupRecurringTaskRequest) SetRetainPreviousFullBackupCount(v int32)`

SetRetainPreviousFullBackupCount sets RetainPreviousFullBackupCount field to given value.

### HasRetainPreviousFullBackupCount

`func (o *AddBackupRecurringTaskRequest) HasRetainPreviousFullBackupCount() bool`

HasRetainPreviousFullBackupCount returns a boolean if a field has been set.

### GetRetainPreviousFullBackupAge

`func (o *AddBackupRecurringTaskRequest) GetRetainPreviousFullBackupAge() string`

GetRetainPreviousFullBackupAge returns the RetainPreviousFullBackupAge field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupAgeOk

`func (o *AddBackupRecurringTaskRequest) GetRetainPreviousFullBackupAgeOk() (*string, bool)`

GetRetainPreviousFullBackupAgeOk returns a tuple with the RetainPreviousFullBackupAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupAge

`func (o *AddBackupRecurringTaskRequest) SetRetainPreviousFullBackupAge(v string)`

SetRetainPreviousFullBackupAge sets RetainPreviousFullBackupAge field to given value.

### HasRetainPreviousFullBackupAge

`func (o *AddBackupRecurringTaskRequest) HasRetainPreviousFullBackupAge() bool`

HasRetainPreviousFullBackupAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *AddBackupRecurringTaskRequest) GetMaxMegabytesPerSecond() int32`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *AddBackupRecurringTaskRequest) GetMaxMegabytesPerSecondOk() (*int32, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *AddBackupRecurringTaskRequest) SetMaxMegabytesPerSecond(v int32)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *AddBackupRecurringTaskRequest) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetDescription

`func (o *AddBackupRecurringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBackupRecurringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBackupRecurringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBackupRecurringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddBackupRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddBackupRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddBackupRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddBackupRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddBackupRecurringTaskRequest) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddBackupRecurringTaskRequest) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddBackupRecurringTaskRequest) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddBackupRecurringTaskRequest) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddBackupRecurringTaskRequest) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddBackupRecurringTaskRequest) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddBackupRecurringTaskRequest) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddBackupRecurringTaskRequest) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddBackupRecurringTaskRequest) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddBackupRecurringTaskRequest) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddBackupRecurringTaskRequest) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddBackupRecurringTaskRequest) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddBackupRecurringTaskRequest) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddBackupRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddBackupRecurringTaskRequest) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddBackupRecurringTaskRequest) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddBackupRecurringTaskRequest) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddBackupRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddBackupRecurringTaskRequest) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddBackupRecurringTaskRequest) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddBackupRecurringTaskRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddBackupRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddBackupRecurringTaskRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddBackupRecurringTaskRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


