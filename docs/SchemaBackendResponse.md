# SchemaBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumschemaBackendSchemaUrn**](EnumschemaBackendSchemaUrn.md) |  | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**BaseDN** | **[]string** |  | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**SchemaEntryDN** | Pointer to **[]string** |  | [optional] 
**ShowAllAttributes** | **bool** | Indicates whether to treat all attributes in the schema entry as if they were user attributes regardless of their configuration. | 
**ReadOnlySchemaFile** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 

## Methods

### NewSchemaBackendResponse

`func NewSchemaBackendResponse(schemas []EnumschemaBackendSchemaUrn, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, showAllAttributes bool, enabled bool, ) *SchemaBackendResponse`

NewSchemaBackendResponse instantiates a new SchemaBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaBackendResponseWithDefaults

`func NewSchemaBackendResponseWithDefaults() *SchemaBackendResponse`

NewSchemaBackendResponseWithDefaults instantiates a new SchemaBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SchemaBackendResponse) GetSchemas() []EnumschemaBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SchemaBackendResponse) GetSchemasOk() (*[]EnumschemaBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SchemaBackendResponse) SetSchemas(v []EnumschemaBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBackendID

`func (o *SchemaBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *SchemaBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *SchemaBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *SchemaBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *SchemaBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *SchemaBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *SchemaBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *SchemaBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *SchemaBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetSchemaEntryDN

`func (o *SchemaBackendResponse) GetSchemaEntryDN() []string`

GetSchemaEntryDN returns the SchemaEntryDN field if non-nil, zero value otherwise.

### GetSchemaEntryDNOk

`func (o *SchemaBackendResponse) GetSchemaEntryDNOk() (*[]string, bool)`

GetSchemaEntryDNOk returns a tuple with the SchemaEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaEntryDN

`func (o *SchemaBackendResponse) SetSchemaEntryDN(v []string)`

SetSchemaEntryDN sets SchemaEntryDN field to given value.

### HasSchemaEntryDN

`func (o *SchemaBackendResponse) HasSchemaEntryDN() bool`

HasSchemaEntryDN returns a boolean if a field has been set.

### GetShowAllAttributes

`func (o *SchemaBackendResponse) GetShowAllAttributes() bool`

GetShowAllAttributes returns the ShowAllAttributes field if non-nil, zero value otherwise.

### GetShowAllAttributesOk

`func (o *SchemaBackendResponse) GetShowAllAttributesOk() (*bool, bool)`

GetShowAllAttributesOk returns a tuple with the ShowAllAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllAttributes

`func (o *SchemaBackendResponse) SetShowAllAttributes(v bool)`

SetShowAllAttributes sets ShowAllAttributes field to given value.


### GetReadOnlySchemaFile

`func (o *SchemaBackendResponse) GetReadOnlySchemaFile() []string`

GetReadOnlySchemaFile returns the ReadOnlySchemaFile field if non-nil, zero value otherwise.

### GetReadOnlySchemaFileOk

`func (o *SchemaBackendResponse) GetReadOnlySchemaFileOk() (*[]string, bool)`

GetReadOnlySchemaFileOk returns a tuple with the ReadOnlySchemaFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySchemaFile

`func (o *SchemaBackendResponse) SetReadOnlySchemaFile(v []string)`

SetReadOnlySchemaFile sets ReadOnlySchemaFile field to given value.

### HasReadOnlySchemaFile

`func (o *SchemaBackendResponse) HasReadOnlySchemaFile() bool`

HasReadOnlySchemaFile returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SchemaBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchemaBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchemaBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *SchemaBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *SchemaBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *SchemaBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *SchemaBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *SchemaBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *SchemaBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *SchemaBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *SchemaBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *SchemaBackendResponse) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *SchemaBackendResponse) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *SchemaBackendResponse) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *SchemaBackendResponse) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *SchemaBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *SchemaBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *SchemaBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *SchemaBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


