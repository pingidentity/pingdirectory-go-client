# MirroredLdifBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummirroredLdifBackendSchemaUrn**](EnummirroredLdifBackendSchemaUrn.md) |  | 
**Id** | **string** | Name of the Backend | 
**MirroredSubtreePeerPollingInterval** | Pointer to **string** | The amount of time to wait before polling the peer servers in the topology to determine if there are any changes in the topology. A lower value will make for a quicker failover in the event of a failure, but it will also cause more frequent traffic among the peers. | [optional] 
**MirroredSubtreeEntryUpdateTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for an update operation (add, delete, modify and modify-dn) on an entry to be applied on all servers in the topology. | [optional] 
**MirroredSubtreeSearchTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a search operation to complete. Search requests that take longer than this timeout will be canceled and considered failures. | [optional] 
**PeerServer** | Pointer to **[]string** | Specifies the set of peer servers onto which updates should be mirrored. The local server should not be included in this set, but if it is, then it will just be ignored. | [optional] 
**MirroredSubtreePreferredMasterType** | Pointer to [**[]EnumbackendMirroredSubtreePreferredMasterTypeProp**](EnumbackendMirroredSubtreePreferredMasterTypeProp.md) |  | [optional] 
**SimulatedResultCode** | Pointer to **int64** | Specifies the numeric value of the result code to be assumed for update operations (add, delete, modify and modify-dn) targeted to this backend. | [optional] 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**IsPrivateBackend** | Pointer to **bool** | Indicates whether the backend should be considered a private backend, which indicates that it is used for storing operational data rather than user-defined information. | [optional] 
**LdifFile** | **string** | Specifies the path to the LDIF file containing the data for this backend. | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewMirroredLdifBackendResponse

`func NewMirroredLdifBackendResponse(schemas []EnummirroredLdifBackendSchemaUrn, id string, writabilityMode EnumbackendWritabilityModeProp, ldifFile string, backendID string, enabled bool, baseDN []string, ) *MirroredLdifBackendResponse`

NewMirroredLdifBackendResponse instantiates a new MirroredLdifBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirroredLdifBackendResponseWithDefaults

`func NewMirroredLdifBackendResponseWithDefaults() *MirroredLdifBackendResponse`

NewMirroredLdifBackendResponseWithDefaults instantiates a new MirroredLdifBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MirroredLdifBackendResponse) GetSchemas() []EnummirroredLdifBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MirroredLdifBackendResponse) GetSchemasOk() (*[]EnummirroredLdifBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MirroredLdifBackendResponse) SetSchemas(v []EnummirroredLdifBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *MirroredLdifBackendResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MirroredLdifBackendResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MirroredLdifBackendResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMirroredSubtreePeerPollingInterval

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreePeerPollingInterval() string`

GetMirroredSubtreePeerPollingInterval returns the MirroredSubtreePeerPollingInterval field if non-nil, zero value otherwise.

### GetMirroredSubtreePeerPollingIntervalOk

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreePeerPollingIntervalOk() (*string, bool)`

GetMirroredSubtreePeerPollingIntervalOk returns a tuple with the MirroredSubtreePeerPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreePeerPollingInterval

`func (o *MirroredLdifBackendResponse) SetMirroredSubtreePeerPollingInterval(v string)`

SetMirroredSubtreePeerPollingInterval sets MirroredSubtreePeerPollingInterval field to given value.

### HasMirroredSubtreePeerPollingInterval

`func (o *MirroredLdifBackendResponse) HasMirroredSubtreePeerPollingInterval() bool`

HasMirroredSubtreePeerPollingInterval returns a boolean if a field has been set.

### GetMirroredSubtreeEntryUpdateTimeout

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreeEntryUpdateTimeout() string`

GetMirroredSubtreeEntryUpdateTimeout returns the MirroredSubtreeEntryUpdateTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeEntryUpdateTimeoutOk

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreeEntryUpdateTimeoutOk() (*string, bool)`

GetMirroredSubtreeEntryUpdateTimeoutOk returns a tuple with the MirroredSubtreeEntryUpdateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeEntryUpdateTimeout

`func (o *MirroredLdifBackendResponse) SetMirroredSubtreeEntryUpdateTimeout(v string)`

SetMirroredSubtreeEntryUpdateTimeout sets MirroredSubtreeEntryUpdateTimeout field to given value.

### HasMirroredSubtreeEntryUpdateTimeout

`func (o *MirroredLdifBackendResponse) HasMirroredSubtreeEntryUpdateTimeout() bool`

HasMirroredSubtreeEntryUpdateTimeout returns a boolean if a field has been set.

### GetMirroredSubtreeSearchTimeout

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreeSearchTimeout() string`

GetMirroredSubtreeSearchTimeout returns the MirroredSubtreeSearchTimeout field if non-nil, zero value otherwise.

### GetMirroredSubtreeSearchTimeoutOk

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreeSearchTimeoutOk() (*string, bool)`

GetMirroredSubtreeSearchTimeoutOk returns a tuple with the MirroredSubtreeSearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreeSearchTimeout

`func (o *MirroredLdifBackendResponse) SetMirroredSubtreeSearchTimeout(v string)`

SetMirroredSubtreeSearchTimeout sets MirroredSubtreeSearchTimeout field to given value.

### HasMirroredSubtreeSearchTimeout

`func (o *MirroredLdifBackendResponse) HasMirroredSubtreeSearchTimeout() bool`

HasMirroredSubtreeSearchTimeout returns a boolean if a field has been set.

### GetPeerServer

`func (o *MirroredLdifBackendResponse) GetPeerServer() []string`

GetPeerServer returns the PeerServer field if non-nil, zero value otherwise.

### GetPeerServerOk

`func (o *MirroredLdifBackendResponse) GetPeerServerOk() (*[]string, bool)`

GetPeerServerOk returns a tuple with the PeerServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServer

`func (o *MirroredLdifBackendResponse) SetPeerServer(v []string)`

SetPeerServer sets PeerServer field to given value.

### HasPeerServer

`func (o *MirroredLdifBackendResponse) HasPeerServer() bool`

HasPeerServer returns a boolean if a field has been set.

### GetMirroredSubtreePreferredMasterType

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreePreferredMasterType() []EnumbackendMirroredSubtreePreferredMasterTypeProp`

GetMirroredSubtreePreferredMasterType returns the MirroredSubtreePreferredMasterType field if non-nil, zero value otherwise.

### GetMirroredSubtreePreferredMasterTypeOk

`func (o *MirroredLdifBackendResponse) GetMirroredSubtreePreferredMasterTypeOk() (*[]EnumbackendMirroredSubtreePreferredMasterTypeProp, bool)`

GetMirroredSubtreePreferredMasterTypeOk returns a tuple with the MirroredSubtreePreferredMasterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredSubtreePreferredMasterType

`func (o *MirroredLdifBackendResponse) SetMirroredSubtreePreferredMasterType(v []EnumbackendMirroredSubtreePreferredMasterTypeProp)`

SetMirroredSubtreePreferredMasterType sets MirroredSubtreePreferredMasterType field to given value.

### HasMirroredSubtreePreferredMasterType

`func (o *MirroredLdifBackendResponse) HasMirroredSubtreePreferredMasterType() bool`

HasMirroredSubtreePreferredMasterType returns a boolean if a field has been set.

### GetSimulatedResultCode

`func (o *MirroredLdifBackendResponse) GetSimulatedResultCode() int64`

GetSimulatedResultCode returns the SimulatedResultCode field if non-nil, zero value otherwise.

### GetSimulatedResultCodeOk

`func (o *MirroredLdifBackendResponse) GetSimulatedResultCodeOk() (*int64, bool)`

GetSimulatedResultCodeOk returns a tuple with the SimulatedResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatedResultCode

`func (o *MirroredLdifBackendResponse) SetSimulatedResultCode(v int64)`

SetSimulatedResultCode sets SimulatedResultCode field to given value.

### HasSimulatedResultCode

`func (o *MirroredLdifBackendResponse) HasSimulatedResultCode() bool`

HasSimulatedResultCode returns a boolean if a field has been set.

### GetWritabilityMode

`func (o *MirroredLdifBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *MirroredLdifBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *MirroredLdifBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetIsPrivateBackend

`func (o *MirroredLdifBackendResponse) GetIsPrivateBackend() bool`

GetIsPrivateBackend returns the IsPrivateBackend field if non-nil, zero value otherwise.

### GetIsPrivateBackendOk

`func (o *MirroredLdifBackendResponse) GetIsPrivateBackendOk() (*bool, bool)`

GetIsPrivateBackendOk returns a tuple with the IsPrivateBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateBackend

`func (o *MirroredLdifBackendResponse) SetIsPrivateBackend(v bool)`

SetIsPrivateBackend sets IsPrivateBackend field to given value.

### HasIsPrivateBackend

`func (o *MirroredLdifBackendResponse) HasIsPrivateBackend() bool`

HasIsPrivateBackend returns a boolean if a field has been set.

### GetLdifFile

`func (o *MirroredLdifBackendResponse) GetLdifFile() string`

GetLdifFile returns the LdifFile field if non-nil, zero value otherwise.

### GetLdifFileOk

`func (o *MirroredLdifBackendResponse) GetLdifFileOk() (*string, bool)`

GetLdifFileOk returns a tuple with the LdifFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifFile

`func (o *MirroredLdifBackendResponse) SetLdifFile(v string)`

SetLdifFile sets LdifFile field to given value.


### GetBackendID

`func (o *MirroredLdifBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *MirroredLdifBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *MirroredLdifBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetDescription

`func (o *MirroredLdifBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MirroredLdifBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MirroredLdifBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MirroredLdifBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MirroredLdifBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MirroredLdifBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MirroredLdifBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *MirroredLdifBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *MirroredLdifBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *MirroredLdifBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *MirroredLdifBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *MirroredLdifBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *MirroredLdifBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *MirroredLdifBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *MirroredLdifBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *MirroredLdifBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *MirroredLdifBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *MirroredLdifBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *MirroredLdifBackendResponse) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *MirroredLdifBackendResponse) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *MirroredLdifBackendResponse) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *MirroredLdifBackendResponse) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *MirroredLdifBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *MirroredLdifBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *MirroredLdifBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *MirroredLdifBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *MirroredLdifBackendResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MirroredLdifBackendResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MirroredLdifBackendResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MirroredLdifBackendResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MirroredLdifBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MirroredLdifBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MirroredLdifBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MirroredLdifBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


