# BackupRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewBackupRecurringTaskShared

`func NewBackupRecurringTaskShared(schemas []EnumbackupRecurringTaskSchemaUrn, backupDirectory string, ) *BackupRecurringTaskShared`

NewBackupRecurringTaskShared instantiates a new BackupRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRecurringTaskSharedWithDefaults

`func NewBackupRecurringTaskSharedWithDefaults() *BackupRecurringTaskShared`

NewBackupRecurringTaskSharedWithDefaults instantiates a new BackupRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BackupRecurringTaskShared) GetSchemas() []EnumbackupRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BackupRecurringTaskShared) GetSchemasOk() (*[]EnumbackupRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BackupRecurringTaskShared) SetSchemas(v []EnumbackupRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBackupDirectory

`func (o *BackupRecurringTaskShared) GetBackupDirectory() string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *BackupRecurringTaskShared) GetBackupDirectoryOk() (*string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *BackupRecurringTaskShared) SetBackupDirectory(v string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIncludedBackendID

`func (o *BackupRecurringTaskShared) GetIncludedBackendID() []string`

GetIncludedBackendID returns the IncludedBackendID field if non-nil, zero value otherwise.

### GetIncludedBackendIDOk

`func (o *BackupRecurringTaskShared) GetIncludedBackendIDOk() (*[]string, bool)`

GetIncludedBackendIDOk returns a tuple with the IncludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendID

`func (o *BackupRecurringTaskShared) SetIncludedBackendID(v []string)`

SetIncludedBackendID sets IncludedBackendID field to given value.

### HasIncludedBackendID

`func (o *BackupRecurringTaskShared) HasIncludedBackendID() bool`

HasIncludedBackendID returns a boolean if a field has been set.

### GetExcludedBackendID

`func (o *BackupRecurringTaskShared) GetExcludedBackendID() []string`

GetExcludedBackendID returns the ExcludedBackendID field if non-nil, zero value otherwise.

### GetExcludedBackendIDOk

`func (o *BackupRecurringTaskShared) GetExcludedBackendIDOk() (*[]string, bool)`

GetExcludedBackendIDOk returns a tuple with the ExcludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendID

`func (o *BackupRecurringTaskShared) SetExcludedBackendID(v []string)`

SetExcludedBackendID sets ExcludedBackendID field to given value.

### HasExcludedBackendID

`func (o *BackupRecurringTaskShared) HasExcludedBackendID() bool`

HasExcludedBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *BackupRecurringTaskShared) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *BackupRecurringTaskShared) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *BackupRecurringTaskShared) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *BackupRecurringTaskShared) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *BackupRecurringTaskShared) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *BackupRecurringTaskShared) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *BackupRecurringTaskShared) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *BackupRecurringTaskShared) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *BackupRecurringTaskShared) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *BackupRecurringTaskShared) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *BackupRecurringTaskShared) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *BackupRecurringTaskShared) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *BackupRecurringTaskShared) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *BackupRecurringTaskShared) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *BackupRecurringTaskShared) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *BackupRecurringTaskShared) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousFullBackupCount

`func (o *BackupRecurringTaskShared) GetRetainPreviousFullBackupCount() int32`

GetRetainPreviousFullBackupCount returns the RetainPreviousFullBackupCount field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupCountOk

`func (o *BackupRecurringTaskShared) GetRetainPreviousFullBackupCountOk() (*int32, bool)`

GetRetainPreviousFullBackupCountOk returns a tuple with the RetainPreviousFullBackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupCount

`func (o *BackupRecurringTaskShared) SetRetainPreviousFullBackupCount(v int32)`

SetRetainPreviousFullBackupCount sets RetainPreviousFullBackupCount field to given value.

### HasRetainPreviousFullBackupCount

`func (o *BackupRecurringTaskShared) HasRetainPreviousFullBackupCount() bool`

HasRetainPreviousFullBackupCount returns a boolean if a field has been set.

### GetRetainPreviousFullBackupAge

`func (o *BackupRecurringTaskShared) GetRetainPreviousFullBackupAge() string`

GetRetainPreviousFullBackupAge returns the RetainPreviousFullBackupAge field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupAgeOk

`func (o *BackupRecurringTaskShared) GetRetainPreviousFullBackupAgeOk() (*string, bool)`

GetRetainPreviousFullBackupAgeOk returns a tuple with the RetainPreviousFullBackupAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupAge

`func (o *BackupRecurringTaskShared) SetRetainPreviousFullBackupAge(v string)`

SetRetainPreviousFullBackupAge sets RetainPreviousFullBackupAge field to given value.

### HasRetainPreviousFullBackupAge

`func (o *BackupRecurringTaskShared) HasRetainPreviousFullBackupAge() bool`

HasRetainPreviousFullBackupAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *BackupRecurringTaskShared) GetMaxMegabytesPerSecond() int32`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *BackupRecurringTaskShared) GetMaxMegabytesPerSecondOk() (*int32, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *BackupRecurringTaskShared) SetMaxMegabytesPerSecond(v int32)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *BackupRecurringTaskShared) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetDescription

`func (o *BackupRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *BackupRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *BackupRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *BackupRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *BackupRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *BackupRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *BackupRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *BackupRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *BackupRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *BackupRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *BackupRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *BackupRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *BackupRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *BackupRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *BackupRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *BackupRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *BackupRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *BackupRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *BackupRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *BackupRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *BackupRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *BackupRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *BackupRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *BackupRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *BackupRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *BackupRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *BackupRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *BackupRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *BackupRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


