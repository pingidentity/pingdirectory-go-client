# TrustStoreBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtrustStoreBackendSchemaUrn**](EnumtrustStoreBackendSchemaUrn.md) |  | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**BaseDN** | **[]string** | Specifies the base DN(s) for the data that the backend handles. | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**TrustStoreFile** | **string** | Specifies the path to the file that stores the trust information. | 
**TrustStoreType** | Pointer to **string** | Specifies the format for the data in the key store file. | [optional] 
**TrustStorePin** | Pointer to **string** | Specifies the clear-text PIN needed to access the Trust Store Backend. | [optional] 
**TrustStorePinFile** | Pointer to **string** | Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the Trust Store Backend. | [optional] 
**TrustStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the clear-text PIN needed to access the Trust Store Backend. | [optional] 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**SetDegradedAlertWhenDisabled** | Pointer to **bool** | Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled. | [optional] 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**BackupFilePermissions** | Pointer to **string** | Specifies the permissions that should be applied to files and directories created by a backup of the backend. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 

## Methods

### NewTrustStoreBackendResponse

`func NewTrustStoreBackendResponse(schemas []EnumtrustStoreBackendSchemaUrn, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, trustStoreFile string, enabled bool, ) *TrustStoreBackendResponse`

NewTrustStoreBackendResponse instantiates a new TrustStoreBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustStoreBackendResponseWithDefaults

`func NewTrustStoreBackendResponseWithDefaults() *TrustStoreBackendResponse`

NewTrustStoreBackendResponseWithDefaults instantiates a new TrustStoreBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TrustStoreBackendResponse) GetSchemas() []EnumtrustStoreBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TrustStoreBackendResponse) GetSchemasOk() (*[]EnumtrustStoreBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TrustStoreBackendResponse) SetSchemas(v []EnumtrustStoreBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBackendID

`func (o *TrustStoreBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *TrustStoreBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *TrustStoreBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetBaseDN

`func (o *TrustStoreBackendResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *TrustStoreBackendResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *TrustStoreBackendResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetWritabilityMode

`func (o *TrustStoreBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *TrustStoreBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *TrustStoreBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetTrustStoreFile

`func (o *TrustStoreBackendResponse) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *TrustStoreBackendResponse) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *TrustStoreBackendResponse) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.


### GetTrustStoreType

`func (o *TrustStoreBackendResponse) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *TrustStoreBackendResponse) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *TrustStoreBackendResponse) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *TrustStoreBackendResponse) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *TrustStoreBackendResponse) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *TrustStoreBackendResponse) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *TrustStoreBackendResponse) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *TrustStoreBackendResponse) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStorePinFile

`func (o *TrustStoreBackendResponse) GetTrustStorePinFile() string`

GetTrustStorePinFile returns the TrustStorePinFile field if non-nil, zero value otherwise.

### GetTrustStorePinFileOk

`func (o *TrustStoreBackendResponse) GetTrustStorePinFileOk() (*string, bool)`

GetTrustStorePinFileOk returns a tuple with the TrustStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinFile

`func (o *TrustStoreBackendResponse) SetTrustStorePinFile(v string)`

SetTrustStorePinFile sets TrustStorePinFile field to given value.

### HasTrustStorePinFile

`func (o *TrustStoreBackendResponse) HasTrustStorePinFile() bool`

HasTrustStorePinFile returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *TrustStoreBackendResponse) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *TrustStoreBackendResponse) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *TrustStoreBackendResponse) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *TrustStoreBackendResponse) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetDescription

`func (o *TrustStoreBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrustStoreBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrustStoreBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrustStoreBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TrustStoreBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TrustStoreBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TrustStoreBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSetDegradedAlertWhenDisabled

`func (o *TrustStoreBackendResponse) GetSetDegradedAlertWhenDisabled() bool`

GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field if non-nil, zero value otherwise.

### GetSetDegradedAlertWhenDisabledOk

`func (o *TrustStoreBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool)`

GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDegradedAlertWhenDisabled

`func (o *TrustStoreBackendResponse) SetSetDegradedAlertWhenDisabled(v bool)`

SetSetDegradedAlertWhenDisabled sets SetDegradedAlertWhenDisabled field to given value.

### HasSetDegradedAlertWhenDisabled

`func (o *TrustStoreBackendResponse) HasSetDegradedAlertWhenDisabled() bool`

HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.

### GetReturnUnavailableWhenDisabled

`func (o *TrustStoreBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *TrustStoreBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *TrustStoreBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *TrustStoreBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetBackupFilePermissions

`func (o *TrustStoreBackendResponse) GetBackupFilePermissions() string`

GetBackupFilePermissions returns the BackupFilePermissions field if non-nil, zero value otherwise.

### GetBackupFilePermissionsOk

`func (o *TrustStoreBackendResponse) GetBackupFilePermissionsOk() (*string, bool)`

GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePermissions

`func (o *TrustStoreBackendResponse) SetBackupFilePermissions(v string)`

SetBackupFilePermissions sets BackupFilePermissions field to given value.

### HasBackupFilePermissions

`func (o *TrustStoreBackendResponse) HasBackupFilePermissions() bool`

HasBackupFilePermissions returns a boolean if a field has been set.

### GetNotificationManager

`func (o *TrustStoreBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *TrustStoreBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *TrustStoreBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *TrustStoreBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


