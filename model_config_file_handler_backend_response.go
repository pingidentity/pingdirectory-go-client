/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConfigFileHandlerBackendResponse struct for ConfigFileHandlerBackendResponse
type ConfigFileHandlerBackendResponse struct {
	Schemas []EnumconfigFileHandlerBackendSchemaUrn `json:"schemas"`
	// Specifies a name to identify the associated backend.
	BackendID string `json:"backendID"`
	// Specifies the base DN(s) for the data that the backend handles.
	BaseDN []string `json:"baseDN"`
	WritabilityMode EnumbackendWritabilityModeProp `json:"writabilityMode"`
	// The name or OID of an attribute type that is considered insignificant for the purpose of maintaining the configuration archive.
	InsignificantConfigArchiveAttribute []string `json:"insignificantConfigArchiveAttribute,omitempty"`
	// Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait before polling the peer servers in the topology to determine if there are any changes in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data.
	MirroredSubtreePeerPollingInterval *string `json:"mirroredSubtreePeerPollingInterval,omitempty"`
	// Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for an update operation (add, delete, modify and modify-dn) on an entry to be applied on all servers in the topology. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data.
	MirroredSubtreeEntryUpdateTimeout *string `json:"mirroredSubtreeEntryUpdateTimeout,omitempty"`
	// Tells the server component that is responsible for mirroring configuration data across a topology of servers the maximum amount of time to wait for a search operation to complete. Mirrored data includes meta-data about the servers in the topology as well as cluster-wide configuration data. Search requests that take longer than this timeout will be canceled and considered failures.
	MirroredSubtreeSearchTimeout *string `json:"mirroredSubtreeSearchTimeout,omitempty"`
	// A description for this Backend
	Description *string `json:"description,omitempty"`
	// Indicates whether the backend is enabled in the server.
	Enabled bool `json:"enabled"`
	// Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled.
	SetDegradedAlertWhenDisabled *bool `json:"setDegradedAlertWhenDisabled,omitempty"`
	// Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled.
	ReturnUnavailableWhenDisabled *bool `json:"returnUnavailableWhenDisabled,omitempty"`
	// Specifies the permissions that should be applied to files and directories created by a backup of the backend.
	BackupFilePermissions *string `json:"backupFilePermissions,omitempty"`
	// Specifies a notification manager for changes resulting from operations processed through this Backend
	NotificationManager *string `json:"notificationManager,omitempty"`
}

// NewConfigFileHandlerBackendResponse instantiates a new ConfigFileHandlerBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigFileHandlerBackendResponse(schemas []EnumconfigFileHandlerBackendSchemaUrn, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, enabled bool) *ConfigFileHandlerBackendResponse {
	this := ConfigFileHandlerBackendResponse{}
	this.Schemas = schemas
	this.BackendID = backendID
	this.BaseDN = baseDN
	this.WritabilityMode = writabilityMode
	this.Enabled = enabled
	return &this
}

// NewConfigFileHandlerBackendResponseWithDefaults instantiates a new ConfigFileHandlerBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigFileHandlerBackendResponseWithDefaults() *ConfigFileHandlerBackendResponse {
	this := ConfigFileHandlerBackendResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ConfigFileHandlerBackendResponse) GetSchemas() []EnumconfigFileHandlerBackendSchemaUrn {
	if o == nil {
		var ret []EnumconfigFileHandlerBackendSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetSchemasOk() ([]EnumconfigFileHandlerBackendSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ConfigFileHandlerBackendResponse) SetSchemas(v []EnumconfigFileHandlerBackendSchemaUrn) {
	o.Schemas = v
}

// GetBackendID returns the BackendID field value
func (o *ConfigFileHandlerBackendResponse) GetBackendID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetBackendIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackendID, true
}

// SetBackendID sets field value
func (o *ConfigFileHandlerBackendResponse) SetBackendID(v string) {
	o.BackendID = v
}

// GetBaseDN returns the BaseDN field value
func (o *ConfigFileHandlerBackendResponse) GetBaseDN() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BaseDN, true
}

// SetBaseDN sets field value
func (o *ConfigFileHandlerBackendResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetWritabilityMode returns the WritabilityMode field value
func (o *ConfigFileHandlerBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp {
	if o == nil {
		var ret EnumbackendWritabilityModeProp
		return ret
	}

	return o.WritabilityMode
}

// GetWritabilityModeOk returns a tuple with the WritabilityMode field value
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WritabilityMode, true
}

// SetWritabilityMode sets field value
func (o *ConfigFileHandlerBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp) {
	o.WritabilityMode = v
}

// GetInsignificantConfigArchiveAttribute returns the InsignificantConfigArchiveAttribute field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetInsignificantConfigArchiveAttribute() []string {
	if o == nil || isNil(o.InsignificantConfigArchiveAttribute) {
		var ret []string
		return ret
	}
	return o.InsignificantConfigArchiveAttribute
}

// GetInsignificantConfigArchiveAttributeOk returns a tuple with the InsignificantConfigArchiveAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetInsignificantConfigArchiveAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.InsignificantConfigArchiveAttribute) {
    return nil, false
	}
	return o.InsignificantConfigArchiveAttribute, true
}

// HasInsignificantConfigArchiveAttribute returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasInsignificantConfigArchiveAttribute() bool {
	if o != nil && !isNil(o.InsignificantConfigArchiveAttribute) {
		return true
	}

	return false
}

// SetInsignificantConfigArchiveAttribute gets a reference to the given []string and assigns it to the InsignificantConfigArchiveAttribute field.
func (o *ConfigFileHandlerBackendResponse) SetInsignificantConfigArchiveAttribute(v []string) {
	o.InsignificantConfigArchiveAttribute = v
}

// GetMirroredSubtreePeerPollingInterval returns the MirroredSubtreePeerPollingInterval field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreePeerPollingInterval() string {
	if o == nil || isNil(o.MirroredSubtreePeerPollingInterval) {
		var ret string
		return ret
	}
	return *o.MirroredSubtreePeerPollingInterval
}

// GetMirroredSubtreePeerPollingIntervalOk returns a tuple with the MirroredSubtreePeerPollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreePeerPollingIntervalOk() (*string, bool) {
	if o == nil || isNil(o.MirroredSubtreePeerPollingInterval) {
    return nil, false
	}
	return o.MirroredSubtreePeerPollingInterval, true
}

// HasMirroredSubtreePeerPollingInterval returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasMirroredSubtreePeerPollingInterval() bool {
	if o != nil && !isNil(o.MirroredSubtreePeerPollingInterval) {
		return true
	}

	return false
}

// SetMirroredSubtreePeerPollingInterval gets a reference to the given string and assigns it to the MirroredSubtreePeerPollingInterval field.
func (o *ConfigFileHandlerBackendResponse) SetMirroredSubtreePeerPollingInterval(v string) {
	o.MirroredSubtreePeerPollingInterval = &v
}

// GetMirroredSubtreeEntryUpdateTimeout returns the MirroredSubtreeEntryUpdateTimeout field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeEntryUpdateTimeout() string {
	if o == nil || isNil(o.MirroredSubtreeEntryUpdateTimeout) {
		var ret string
		return ret
	}
	return *o.MirroredSubtreeEntryUpdateTimeout
}

// GetMirroredSubtreeEntryUpdateTimeoutOk returns a tuple with the MirroredSubtreeEntryUpdateTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeEntryUpdateTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.MirroredSubtreeEntryUpdateTimeout) {
    return nil, false
	}
	return o.MirroredSubtreeEntryUpdateTimeout, true
}

// HasMirroredSubtreeEntryUpdateTimeout returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasMirroredSubtreeEntryUpdateTimeout() bool {
	if o != nil && !isNil(o.MirroredSubtreeEntryUpdateTimeout) {
		return true
	}

	return false
}

// SetMirroredSubtreeEntryUpdateTimeout gets a reference to the given string and assigns it to the MirroredSubtreeEntryUpdateTimeout field.
func (o *ConfigFileHandlerBackendResponse) SetMirroredSubtreeEntryUpdateTimeout(v string) {
	o.MirroredSubtreeEntryUpdateTimeout = &v
}

// GetMirroredSubtreeSearchTimeout returns the MirroredSubtreeSearchTimeout field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeSearchTimeout() string {
	if o == nil || isNil(o.MirroredSubtreeSearchTimeout) {
		var ret string
		return ret
	}
	return *o.MirroredSubtreeSearchTimeout
}

// GetMirroredSubtreeSearchTimeoutOk returns a tuple with the MirroredSubtreeSearchTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetMirroredSubtreeSearchTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.MirroredSubtreeSearchTimeout) {
    return nil, false
	}
	return o.MirroredSubtreeSearchTimeout, true
}

// HasMirroredSubtreeSearchTimeout returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasMirroredSubtreeSearchTimeout() bool {
	if o != nil && !isNil(o.MirroredSubtreeSearchTimeout) {
		return true
	}

	return false
}

// SetMirroredSubtreeSearchTimeout gets a reference to the given string and assigns it to the MirroredSubtreeSearchTimeout field.
func (o *ConfigFileHandlerBackendResponse) SetMirroredSubtreeSearchTimeout(v string) {
	o.MirroredSubtreeSearchTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigFileHandlerBackendResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ConfigFileHandlerBackendResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConfigFileHandlerBackendResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetSetDegradedAlertWhenDisabled() bool {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.SetDegradedAlertWhenDisabled
}

// GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
    return nil, false
	}
	return o.SetDegradedAlertWhenDisabled, true
}

// HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasSetDegradedAlertWhenDisabled() bool {
	if o != nil && !isNil(o.SetDegradedAlertWhenDisabled) {
		return true
	}

	return false
}

// SetSetDegradedAlertWhenDisabled gets a reference to the given bool and assigns it to the SetDegradedAlertWhenDisabled field.
func (o *ConfigFileHandlerBackendResponse) SetSetDegradedAlertWhenDisabled(v bool) {
	o.SetDegradedAlertWhenDisabled = &v
}

// GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetReturnUnavailableWhenDisabled() bool {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.ReturnUnavailableWhenDisabled
}

// GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
    return nil, false
	}
	return o.ReturnUnavailableWhenDisabled, true
}

// HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasReturnUnavailableWhenDisabled() bool {
	if o != nil && !isNil(o.ReturnUnavailableWhenDisabled) {
		return true
	}

	return false
}

// SetReturnUnavailableWhenDisabled gets a reference to the given bool and assigns it to the ReturnUnavailableWhenDisabled field.
func (o *ConfigFileHandlerBackendResponse) SetReturnUnavailableWhenDisabled(v bool) {
	o.ReturnUnavailableWhenDisabled = &v
}

// GetBackupFilePermissions returns the BackupFilePermissions field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetBackupFilePermissions() string {
	if o == nil || isNil(o.BackupFilePermissions) {
		var ret string
		return ret
	}
	return *o.BackupFilePermissions
}

// GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetBackupFilePermissionsOk() (*string, bool) {
	if o == nil || isNil(o.BackupFilePermissions) {
    return nil, false
	}
	return o.BackupFilePermissions, true
}

// HasBackupFilePermissions returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasBackupFilePermissions() bool {
	if o != nil && !isNil(o.BackupFilePermissions) {
		return true
	}

	return false
}

// SetBackupFilePermissions gets a reference to the given string and assigns it to the BackupFilePermissions field.
func (o *ConfigFileHandlerBackendResponse) SetBackupFilePermissions(v string) {
	o.BackupFilePermissions = &v
}

// GetNotificationManager returns the NotificationManager field value if set, zero value otherwise.
func (o *ConfigFileHandlerBackendResponse) GetNotificationManager() string {
	if o == nil || isNil(o.NotificationManager) {
		var ret string
		return ret
	}
	return *o.NotificationManager
}

// GetNotificationManagerOk returns a tuple with the NotificationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFileHandlerBackendResponse) GetNotificationManagerOk() (*string, bool) {
	if o == nil || isNil(o.NotificationManager) {
    return nil, false
	}
	return o.NotificationManager, true
}

// HasNotificationManager returns a boolean if a field has been set.
func (o *ConfigFileHandlerBackendResponse) HasNotificationManager() bool {
	if o != nil && !isNil(o.NotificationManager) {
		return true
	}

	return false
}

// SetNotificationManager gets a reference to the given string and assigns it to the NotificationManager field.
func (o *ConfigFileHandlerBackendResponse) SetNotificationManager(v string) {
	o.NotificationManager = &v
}

func (o ConfigFileHandlerBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["backendID"] = o.BackendID
	}
	if true {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["writabilityMode"] = o.WritabilityMode
	}
	if !isNil(o.InsignificantConfigArchiveAttribute) {
		toSerialize["insignificantConfigArchiveAttribute"] = o.InsignificantConfigArchiveAttribute
	}
	if !isNil(o.MirroredSubtreePeerPollingInterval) {
		toSerialize["mirroredSubtreePeerPollingInterval"] = o.MirroredSubtreePeerPollingInterval
	}
	if !isNil(o.MirroredSubtreeEntryUpdateTimeout) {
		toSerialize["mirroredSubtreeEntryUpdateTimeout"] = o.MirroredSubtreeEntryUpdateTimeout
	}
	if !isNil(o.MirroredSubtreeSearchTimeout) {
		toSerialize["mirroredSubtreeSearchTimeout"] = o.MirroredSubtreeSearchTimeout
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SetDegradedAlertWhenDisabled) {
		toSerialize["setDegradedAlertWhenDisabled"] = o.SetDegradedAlertWhenDisabled
	}
	if !isNil(o.ReturnUnavailableWhenDisabled) {
		toSerialize["returnUnavailableWhenDisabled"] = o.ReturnUnavailableWhenDisabled
	}
	if !isNil(o.BackupFilePermissions) {
		toSerialize["backupFilePermissions"] = o.BackupFilePermissions
	}
	if !isNil(o.NotificationManager) {
		toSerialize["notificationManager"] = o.NotificationManager
	}
	return json.Marshal(toSerialize)
}

type NullableConfigFileHandlerBackendResponse struct {
	value *ConfigFileHandlerBackendResponse
	isSet bool
}

func (v NullableConfigFileHandlerBackendResponse) Get() *ConfigFileHandlerBackendResponse {
	return v.value
}

func (v *NullableConfigFileHandlerBackendResponse) Set(val *ConfigFileHandlerBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigFileHandlerBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigFileHandlerBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigFileHandlerBackendResponse(val *ConfigFileHandlerBackendResponse) *NullableConfigFileHandlerBackendResponse {
	return &NullableConfigFileHandlerBackendResponse{value: val, isSet: true}
}

func (v NullableConfigFileHandlerBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigFileHandlerBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


