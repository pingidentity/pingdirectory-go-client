# BackupBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumbackupBackendSchemaUrn**](EnumbackupBackendSchemaUrn.md) |  | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**BaseDN** | **[]string** |  | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**BackupDirectory** | **[]string** |  | 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 

## Methods

### NewBackupBackendResponse

`func NewBackupBackendResponse(schemas []EnumbackupBackendSchemaUrn, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, backupDirectory []string, enabled bool, ) *BackupBackendResponse`

NewBackupBackendResponse instantiates a new BackupBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupBackendResponseWithDefaults

`func NewBackupBackendResponseWithDefaults() *BackupBackendResponse`

NewBackupBackendResponseWithDefaults instantiates a new BackupBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BackupBackendResponse) GetSchemas() []EnumbackupBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BackupBackendResponse) GetSchemasOk() (*[]EnumbackupBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BackupBackendResponse) SetSchemas(v []EnumbackupBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBackendID

`func (o *BackupBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *BackupBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *BackupBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *BackupBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *BackupBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *BackupBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *BackupBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *BackupBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *BackupBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetBackupDirectory

`func (o *BackupBackendResponse) GetBackupDirectory() []string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *BackupBackendResponse) GetBackupDirectoryOk() (*[]string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *BackupBackendResponse) SetBackupDirectory(v []string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetDescription

`func (o *BackupBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BackupBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *BackupBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *BackupBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *BackupBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *BackupBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *BackupBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *BackupBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *BackupBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *BackupBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *BackupBackendResponse) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *BackupBackendResponse) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *BackupBackendResponse) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *BackupBackendResponse) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *BackupBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *BackupBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *BackupBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *BackupBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


