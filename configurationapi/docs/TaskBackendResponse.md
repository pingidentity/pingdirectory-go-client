# TaskBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtaskBackendSchemaUrn**](EnumtaskBackendSchemaUrn.md) |  | 
**Id** | **string** | Name of the Backend | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**TaskBackingFile** | **string** | Specifies the path to the backing file for storing information about the tasks configured in the server. | 
**MaximumInitialTaskLogMessagesToRetain** | Pointer to **int32** | The maximum number of log messages to retain in each task entry from the beginning of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks. | [optional] 
**MaximumFinalTaskLogMessagesToRetain** | Pointer to **int32** | The maximum number of log messages to retain in each task entry from the end of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks. | [optional] 
**TaskRetentionTime** | Pointer to **string** | Specifies the length of time that task entries should be retained after processing on the associated task has been completed. | [optional] 
**NotificationSenderAddress** | Pointer to **string** | Specifies the email address to use as the sender address (that is, the \&quot;From:\&quot; address) for notification mail messages generated when a task completes execution. | [optional] 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewTaskBackendResponse

`func NewTaskBackendResponse(schemas []EnumtaskBackendSchemaUrn, id string, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, taskBackingFile string, enabled bool, ) *TaskBackendResponse`

NewTaskBackendResponse instantiates a new TaskBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskBackendResponseWithDefaults

`func NewTaskBackendResponseWithDefaults() *TaskBackendResponse`

NewTaskBackendResponseWithDefaults instantiates a new TaskBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TaskBackendResponse) GetSchemas() []EnumtaskBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TaskBackendResponse) GetSchemasOk() (*[]EnumtaskBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TaskBackendResponse) SetSchemas(v []EnumtaskBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *TaskBackendResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskBackendResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskBackendResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBackendID

`func (o *TaskBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *TaskBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *TaskBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *TaskBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *TaskBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *TaskBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *TaskBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *TaskBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *TaskBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetTaskBackingFile

`func (o *TaskBackendResponse) GetTaskBackingFile() string`

GetTaskBackingFile returns the TaskBackingFile field if non-nil, zero value otherwise.

### GetTaskBackingFileOk

`func (o *TaskBackendResponse) GetTaskBackingFileOk() (*string, bool)`

GetTaskBackingFileOk returns a tuple with the TaskBackingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskBackingFile

`func (o *TaskBackendResponse) SetTaskBackingFile(v string)`

SetTaskBackingFile sets TaskBackingFile field to given value.


### GetMaximumInitialTaskLogMessagesToRetain

`func (o *TaskBackendResponse) GetMaximumInitialTaskLogMessagesToRetain() int32`

GetMaximumInitialTaskLogMessagesToRetain returns the MaximumInitialTaskLogMessagesToRetain field if non-nil, zero value otherwise.

### GetMaximumInitialTaskLogMessagesToRetainOk

`func (o *TaskBackendResponse) GetMaximumInitialTaskLogMessagesToRetainOk() (*int32, bool)`

GetMaximumInitialTaskLogMessagesToRetainOk returns a tuple with the MaximumInitialTaskLogMessagesToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInitialTaskLogMessagesToRetain

`func (o *TaskBackendResponse) SetMaximumInitialTaskLogMessagesToRetain(v int32)`

SetMaximumInitialTaskLogMessagesToRetain sets MaximumInitialTaskLogMessagesToRetain field to given value.

### HasMaximumInitialTaskLogMessagesToRetain

`func (o *TaskBackendResponse) HasMaximumInitialTaskLogMessagesToRetain() bool`

HasMaximumInitialTaskLogMessagesToRetain returns a boolean if a field has been set.

### GetMaximumFinalTaskLogMessagesToRetain

`func (o *TaskBackendResponse) GetMaximumFinalTaskLogMessagesToRetain() int32`

GetMaximumFinalTaskLogMessagesToRetain returns the MaximumFinalTaskLogMessagesToRetain field if non-nil, zero value otherwise.

### GetMaximumFinalTaskLogMessagesToRetainOk

`func (o *TaskBackendResponse) GetMaximumFinalTaskLogMessagesToRetainOk() (*int32, bool)`

GetMaximumFinalTaskLogMessagesToRetainOk returns a tuple with the MaximumFinalTaskLogMessagesToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFinalTaskLogMessagesToRetain

`func (o *TaskBackendResponse) SetMaximumFinalTaskLogMessagesToRetain(v int32)`

SetMaximumFinalTaskLogMessagesToRetain sets MaximumFinalTaskLogMessagesToRetain field to given value.

### HasMaximumFinalTaskLogMessagesToRetain

`func (o *TaskBackendResponse) HasMaximumFinalTaskLogMessagesToRetain() bool`

HasMaximumFinalTaskLogMessagesToRetain returns a boolean if a field has been set.

### GetTaskRetentionTime

`func (o *TaskBackendResponse) GetTaskRetentionTime() string`

GetTaskRetentionTime returns the TaskRetentionTime field if non-nil, zero value otherwise.

### GetTaskRetentionTimeOk

`func (o *TaskBackendResponse) GetTaskRetentionTimeOk() (*string, bool)`

GetTaskRetentionTimeOk returns a tuple with the TaskRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRetentionTime

`func (o *TaskBackendResponse) SetTaskRetentionTime(v string)`

SetTaskRetentionTime sets TaskRetentionTime field to given value.

### HasTaskRetentionTime

`func (o *TaskBackendResponse) HasTaskRetentionTime() bool`

HasTaskRetentionTime returns a boolean if a field has been set.

### GetNotificationSenderAddress

`func (o *TaskBackendResponse) GetNotificationSenderAddress() string`

GetNotificationSenderAddress returns the NotificationSenderAddress field if non-nil, zero value otherwise.

### GetNotificationSenderAddressOk

`func (o *TaskBackendResponse) GetNotificationSenderAddressOk() (*string, bool)`

GetNotificationSenderAddressOk returns a tuple with the NotificationSenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSenderAddress

`func (o *TaskBackendResponse) SetNotificationSenderAddress(v string)`

SetNotificationSenderAddress sets NotificationSenderAddress field to given value.

### HasNotificationSenderAddress

`func (o *TaskBackendResponse) HasNotificationSenderAddress() bool`

HasNotificationSenderAddress returns a boolean if a field has been set.

### GetDescription

`func (o *TaskBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TaskBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TaskBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TaskBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *TaskBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *TaskBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *TaskBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *TaskBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *TaskBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *TaskBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *TaskBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *TaskBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *TaskBackendResponse) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *TaskBackendResponse) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *TaskBackendResponse) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *TaskBackendResponse) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *TaskBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *TaskBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *TaskBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *TaskBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *TaskBackendResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TaskBackendResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TaskBackendResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TaskBackendResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TaskBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TaskBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TaskBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TaskBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


