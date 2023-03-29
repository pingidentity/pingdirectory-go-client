# ConfigFileHandlerBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumconfigFileHandlerBackendSchemaUrn**](EnumconfigFileHandlerBackendSchemaUrn.md) |  | 
**Id** | **string** | Name of the Backend | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**InsignificantConfigArchiveAttribute** | Pointer to **[]string** | The name or OID of an attribute type that is considered insignificant for the purpose of maintaining the configuration archive. | [optional] 
**MirroredSubtreePeerPollingInterval** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait before polling the peer servers in the topology to determine if there are any changes in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. | [optional] 
**MirroredSubtreeEntryUpdateTimeout** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for an update operation (add, delete, modify and modify-dn) on an entry to be applied on all servers in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. | [optional] 
**MirroredSubtreeSearchTimeout** | Pointer to **string** | Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for a search operation to complete. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. Search requests that take longer than this timeout will be canceled and considered failures. | [optional] 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewConfigFileHandlerBackendResponse

`func NewConfigFileHandlerBackendResponse(schemas []EnumconfigFileHandlerBackendSchemaUrn, id string, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, enabled bool, ) *ConfigFileHandlerBackendResponse`

NewConfigFileHandlerBackendResponse instantiates a new ConfigFileHandlerBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFileHandlerBackendResponseWithDefaults

`func NewConfigFileHandlerBackendResponseWithDefaults() *ConfigFileHandlerBackendResponse`

NewConfigFileHandlerBackendResponseWithDefaults instantiates a new ConfigFileHandlerBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConfigFileHandlerBackendResponse) GetSchemas() []EnumconfigFileHandlerBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConfigFileHandlerBackendResponse) GetSchemasOk() (*[]EnumconfigFileHandlerBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConfigFileHandlerBackendResponse) SetSchemas(v []EnumconfigFileHandlerBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ConfigFileHandlerBackendResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigFileHandlerBackendResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigFileHandlerBackendResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBackendID

`func (o *ConfigFileHandlerBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *ConfigFileHandlerBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *ConfigFileHandlerBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *ConfigFileHandlerBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ConfigFileHandlerBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ConfigFileHandlerBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *ConfigFileHandlerBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *ConfigFileHandlerBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *ConfigFileHandlerBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetInsignificantConfigArchiveAttribute

`func (o *ConfigFileHandlerBackendResponse) GetInsignificantConfigArchiveAttribute() []string`

GetInsignificantConfigArchiveAttribute returns the InsignificantConfigArchiveAttribute field if non-nil, zero value otherwise.

### GetInsignificantConfigArchiveAttributeOk

`func (o *ConfigFileHandlerBackendResponse) GetInsignificantConfigArchiveAttributeOk() (*[]string, bool)`

GetInsignificantConfigArchiveAttributeOk returns a tuple with the InsignificantConfigArchiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsignificantConfigArchiveAttribute

`func (o *ConfigFileHandlerBackendResponse) SetInsignificantConfigArchiveAttribute(v []string)`

SetInsignificantConfigArchiveAttribute sets InsignificantConfigArchiveAttribute field to given value.

### HasInsignificantConfigArchiveAttribute

`func (o *ConfigFileHandlerBackendResponse) HasInsignificantConfigArchiveAttribute() bool`

HasInsignificantConfigArchiveAttribute returns a boolean if a field has been set.

### GetMirroredSubtreePeerPollingInterval

`func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreePeerPollingInterval() string`

GetMirroredSubtreePeerPollingInterval returns the MirroredSubtreePeerPollingInterval field if non-nil, zero value otherwise.

### GetMirroredSubtreePeerPollingIntervalOk

`func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreePeerPollingIntervalOk() (*string, bool)`

GetMirroredSubtreePeerPollingIntervalOk returns a tuple with the MirroredSubtreePeerPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreePeerPollingInterval

`func (o *ConfigFileHandlerBackendResponse) SetMirroredSubtreePeerPollingInterval(v string)`

SetMirroredSubtreePeerPollingInterval sets MirroredSubtreePeerPollingInterval field to given value.

### HasMirroredSubtreePeerPollingInterval

`func (o *ConfigFileHandlerBackendResponse) HasMirroredSubtreePeerPollingInterval() bool`

HasMirroredSubtreePeerPollingInterval returns a boolean if a field has been set.

### GetMirroredSubtreeEntryUpdateTimeout

`func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeEntryUpdateTimeout() string`

GetMirroredSubtreeEntryUpdateTimeout returns the MirroredSubtreeEntryUpdateTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeEntryUpdateTimeoutOk

`func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeEntryUpdateTimeoutOk() (*string, bool)`

GetMirroredSubtreeEntryUpdateTimeoutOk returns a tuple with the MirroredSubtreeEntryUpdateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeEntryUpdateTimeout

`func (o *ConfigFileHandlerBackendResponse) SetMirroredSubtreeEntryUpdateTimeout(v string)`

SetMirroredSubtreeEntryUpdateTimeout sets MirroredSubtreeEntryUpdateTimeout field to given value.

### HasMirroredSubtreeEntryUpdateTimeout

`func (o *ConfigFileHandlerBackendResponse) HasMirroredSubtreeEntryUpdateTimeout() bool`

HasMirroredSubtreeEntryUpdateTimeout returns a boolean if a field has been set.

### GetMirroredSubtreeSearchTimeout

`func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeSearchTimeout() string`

GetMirroredSubtreeSearchTimeout returns the MirroredSubtreeSearchTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeSearchTimeoutOk

`func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeSearchTimeoutOk() (*string, bool)`

GetMirroredSubtreeSearchTimeoutOk returns a tuple with the MirroredSubtreeSearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeSearchTimeout

`func (o *ConfigFileHandlerBackendResponse) SetMirroredSubtreeSearchTimeout(v string)`

SetMirroredSubtreeSearchTimeout sets MirroredSubtreeSearchTimeout field to given value.

### HasMirroredSubtreeSearchTimeout

`func (o *ConfigFileHandlerBackendResponse) HasMirroredSubtreeSearchTimeout() bool`

HasMirroredSubtreeSearchTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigFileHandlerBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigFileHandlerBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigFileHandlerBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigFileHandlerBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ConfigFileHandlerBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConfigFileHandlerBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConfigFileHandlerBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *ConfigFileHandlerBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *ConfigFileHandlerBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *ConfigFileHandlerBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *ConfigFileHandlerBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *ConfigFileHandlerBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *ConfigFileHandlerBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *ConfigFileHandlerBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *ConfigFileHandlerBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *ConfigFileHandlerBackendResponse) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *ConfigFileHandlerBackendResponse) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *ConfigFileHandlerBackendResponse) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *ConfigFileHandlerBackendResponse) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *ConfigFileHandlerBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *ConfigFileHandlerBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *ConfigFileHandlerBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *ConfigFileHandlerBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *ConfigFileHandlerBackendResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConfigFileHandlerBackendResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConfigFileHandlerBackendResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConfigFileHandlerBackendResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConfigFileHandlerBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConfigFileHandlerBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConfigFileHandlerBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConfigFileHandlerBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


