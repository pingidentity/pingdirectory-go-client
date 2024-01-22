# BackupRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumbackupRecurringTaskSchemaUrn**](EnumbackupRecurringTaskSchemaUrn.md) |  | 
**BackupDirectory** | **string** | The directory in which backup files will be placed. When backing up a single backend, the backup files will be placed directly in this directory. When backing up multiple backends, the backup files for each backend will be placed in a subdirectory whose name is the corresponding backend ID. | 
**IncludedBackendID** | Pointer to **[]string** | The backend IDs of any backends that should be included in the backup. | [optional] 
**ExcludedBackendID** | Pointer to **[]string** | The backend IDs of any backends that should be excluded from the backup. All backends that support backups and are not listed will be included. | [optional] 
**Compress** | Pointer to **bool** | Indicates whether to compress the data as it is written into the backup. | [optional] 
**Encrypt** | Pointer to **bool** | Indicates whether to encrypt the data as it is written into the backup. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The ID of an encryption settings definition to use to obtain the backup encryption key. | [optional] 
**Sign** | Pointer to **bool** | Indicates whether to cryptographically sign backups, which will make it possible to detect whether the backup has been altered since it was created. | [optional] 
**RetainPreviousFullBackupCount** | Pointer to **int64** | The minimum number of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**RetainPreviousFullBackupAge** | Pointer to **string** | The minimum age of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**MaxMegabytesPerSecond** | Pointer to **int64** | The maximum rate, in megabytes per second, at which backups should be written. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Recurring Task | 

## Methods

### NewBackupRecurringTaskResponse

`func NewBackupRecurringTaskResponse(schemas []EnumbackupRecurringTaskSchemaUrn, backupDirectory string, id string, ) *BackupRecurringTaskResponse`

NewBackupRecurringTaskResponse instantiates a new BackupRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRecurringTaskResponseWithDefaults

`func NewBackupRecurringTaskResponseWithDefaults() *BackupRecurringTaskResponse`

NewBackupRecurringTaskResponseWithDefaults instantiates a new BackupRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BackupRecurringTaskResponse) GetSchemas() []EnumbackupRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BackupRecurringTaskResponse) GetSchemasOk() (*[]EnumbackupRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BackupRecurringTaskResponse) SetSchemas(v []EnumbackupRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBackupDirectory

`func (o *BackupRecurringTaskResponse) GetBackupDirectory() string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *BackupRecurringTaskResponse) GetBackupDirectoryOk() (*string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *BackupRecurringTaskResponse) SetBackupDirectory(v string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIncludedBackendID

`func (o *BackupRecurringTaskResponse) GetIncludedBackendID() []string`

GetIncludedBackendID returns the IncludedBackendID field if non-nil, zero value otherwise.

### GetIncludedBackendIDOk

`func (o *BackupRecurringTaskResponse) GetIncludedBackendIDOk() (*[]string, bool)`

GetIncludedBackendIDOk returns a tuple with the IncludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendID

`func (o *BackupRecurringTaskResponse) SetIncludedBackendID(v []string)`

SetIncludedBackendID sets IncludedBackendID field to given value.

### HasIncludedBackendID

`func (o *BackupRecurringTaskResponse) HasIncludedBackendID() bool`

HasIncludedBackendID returns a boolean if a field has been set.

### GetExcludedBackendID

`func (o *BackupRecurringTaskResponse) GetExcludedBackendID() []string`

GetExcludedBackendID returns the ExcludedBackendID field if non-nil, zero value otherwise.

### GetExcludedBackendIDOk

`func (o *BackupRecurringTaskResponse) GetExcludedBackendIDOk() (*[]string, bool)`

GetExcludedBackendIDOk returns a tuple with the ExcludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendID

`func (o *BackupRecurringTaskResponse) SetExcludedBackendID(v []string)`

SetExcludedBackendID sets ExcludedBackendID field to given value.

### HasExcludedBackendID

`func (o *BackupRecurringTaskResponse) HasExcludedBackendID() bool`

HasExcludedBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *BackupRecurringTaskResponse) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *BackupRecurringTaskResponse) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *BackupRecurringTaskResponse) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *BackupRecurringTaskResponse) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *BackupRecurringTaskResponse) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *BackupRecurringTaskResponse) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *BackupRecurringTaskResponse) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *BackupRecurringTaskResponse) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *BackupRecurringTaskResponse) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *BackupRecurringTaskResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *BackupRecurringTaskResponse) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *BackupRecurringTaskResponse) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *BackupRecurringTaskResponse) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *BackupRecurringTaskResponse) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *BackupRecurringTaskResponse) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *BackupRecurringTaskResponse) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousFullBackupCount

`func (o *BackupRecurringTaskResponse) GetRetainPreviousFullBackupCount() int64`

GetRetainPreviousFullBackupCount returns the RetainPreviousFullBackupCount field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupCountOk

`func (o *BackupRecurringTaskResponse) GetRetainPreviousFullBackupCountOk() (*int64, bool)`

GetRetainPreviousFullBackupCountOk returns a tuple with the RetainPreviousFullBackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupCount

`func (o *BackupRecurringTaskResponse) SetRetainPreviousFullBackupCount(v int64)`

SetRetainPreviousFullBackupCount sets RetainPreviousFullBackupCount field to given value.

### HasRetainPreviousFullBackupCount

`func (o *BackupRecurringTaskResponse) HasRetainPreviousFullBackupCount() bool`

HasRetainPreviousFullBackupCount returns a boolean if a field has been set.

### GetRetainPreviousFullBackupAge

`func (o *BackupRecurringTaskResponse) GetRetainPreviousFullBackupAge() string`

GetRetainPreviousFullBackupAge returns the RetainPreviousFullBackupAge field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupAgeOk

`func (o *BackupRecurringTaskResponse) GetRetainPreviousFullBackupAgeOk() (*string, bool)`

GetRetainPreviousFullBackupAgeOk returns a tuple with the RetainPreviousFullBackupAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupAge

`func (o *BackupRecurringTaskResponse) SetRetainPreviousFullBackupAge(v string)`

SetRetainPreviousFullBackupAge sets RetainPreviousFullBackupAge field to given value.

### HasRetainPreviousFullBackupAge

`func (o *BackupRecurringTaskResponse) HasRetainPreviousFullBackupAge() bool`

HasRetainPreviousFullBackupAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *BackupRecurringTaskResponse) GetMaxMegabytesPerSecond() int64`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *BackupRecurringTaskResponse) GetMaxMegabytesPerSecondOk() (*int64, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *BackupRecurringTaskResponse) SetMaxMegabytesPerSecond(v int64)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *BackupRecurringTaskResponse) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetDescription

`func (o *BackupRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *BackupRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *BackupRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *BackupRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *BackupRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *BackupRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *BackupRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *BackupRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *BackupRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *BackupRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *BackupRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *BackupRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *BackupRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *BackupRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *BackupRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *BackupRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *BackupRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *BackupRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *BackupRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *BackupRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *BackupRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *BackupRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *BackupRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *BackupRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *BackupRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *BackupRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *BackupRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *BackupRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *BackupRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetMeta

`func (o *BackupRecurringTaskResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BackupRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BackupRecurringTaskResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BackupRecurringTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *BackupRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *BackupRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *BackupRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *BackupRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *BackupRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


