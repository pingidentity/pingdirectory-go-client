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

// EncryptionSettingsBackendResponse struct for EncryptionSettingsBackendResponse
type EncryptionSettingsBackendResponse struct {
	Schemas []EnumencryptionSettingsBackendSchemaUrn `json:"schemas"`
	BaseDN []string `json:"baseDN"`
	// Specifies a name to identify the associated backend.
	BackendID string `json:"backendID"`
	// A description for this Backend
	Description *string `json:"description,omitempty"`
	// Indicates whether the backend is enabled in the server.
	Enabled bool `json:"enabled"`
	WritabilityMode EnumbackendWritabilityModeProp `json:"writabilityMode"`
	// Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled.
	SetDegradedAlertWhenDisabled *bool `json:"setDegradedAlertWhenDisabled,omitempty"`
	// Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled.
	ReturnUnavailableWhenDisabled *bool `json:"returnUnavailableWhenDisabled,omitempty"`
	// Specifies the permissions that should be applied to files and directories created by a backup of the backend.
	BackupFilePermissions *string `json:"backupFilePermissions,omitempty"`
	// Specifies a notification manager for changes resulting from operations processed through this Backend
	NotificationManager *string `json:"notificationManager,omitempty"`
}

// NewEncryptionSettingsBackendResponse instantiates a new EncryptionSettingsBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionSettingsBackendResponse(schemas []EnumencryptionSettingsBackendSchemaUrn, baseDN []string, backendID string, enabled bool, writabilityMode EnumbackendWritabilityModeProp) *EncryptionSettingsBackendResponse {
	this := EncryptionSettingsBackendResponse{}
	this.Schemas = schemas
	this.BaseDN = baseDN
	this.BackendID = backendID
	this.Enabled = enabled
	this.WritabilityMode = writabilityMode
	return &this
}

// NewEncryptionSettingsBackendResponseWithDefaults instantiates a new EncryptionSettingsBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionSettingsBackendResponseWithDefaults() *EncryptionSettingsBackendResponse {
	this := EncryptionSettingsBackendResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *EncryptionSettingsBackendResponse) GetSchemas() []EnumencryptionSettingsBackendSchemaUrn {
	if o == nil {
		var ret []EnumencryptionSettingsBackendSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetSchemasOk() ([]EnumencryptionSettingsBackendSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *EncryptionSettingsBackendResponse) SetSchemas(v []EnumencryptionSettingsBackendSchemaUrn) {
	o.Schemas = v
}

// GetBaseDN returns the BaseDN field value
func (o *EncryptionSettingsBackendResponse) GetBaseDN() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BaseDN, true
}

// SetBaseDN sets field value
func (o *EncryptionSettingsBackendResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetBackendID returns the BackendID field value
func (o *EncryptionSettingsBackendResponse) GetBackendID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetBackendIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackendID, true
}

// SetBackendID sets field value
func (o *EncryptionSettingsBackendResponse) SetBackendID(v string) {
	o.BackendID = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EncryptionSettingsBackendResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EncryptionSettingsBackendResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EncryptionSettingsBackendResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *EncryptionSettingsBackendResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EncryptionSettingsBackendResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWritabilityMode returns the WritabilityMode field value
func (o *EncryptionSettingsBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp {
	if o == nil {
		var ret EnumbackendWritabilityModeProp
		return ret
	}

	return o.WritabilityMode
}

// GetWritabilityModeOk returns a tuple with the WritabilityMode field value
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WritabilityMode, true
}

// SetWritabilityMode sets field value
func (o *EncryptionSettingsBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp) {
	o.WritabilityMode = v
}

// GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field value if set, zero value otherwise.
func (o *EncryptionSettingsBackendResponse) GetSetDegradedAlertWhenDisabled() bool {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.SetDegradedAlertWhenDisabled
}

// GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
    return nil, false
	}
	return o.SetDegradedAlertWhenDisabled, true
}

// HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.
func (o *EncryptionSettingsBackendResponse) HasSetDegradedAlertWhenDisabled() bool {
	if o != nil && !isNil(o.SetDegradedAlertWhenDisabled) {
		return true
	}

	return false
}

// SetSetDegradedAlertWhenDisabled gets a reference to the given bool and assigns it to the SetDegradedAlertWhenDisabled field.
func (o *EncryptionSettingsBackendResponse) SetSetDegradedAlertWhenDisabled(v bool) {
	o.SetDegradedAlertWhenDisabled = &v
}

// GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field value if set, zero value otherwise.
func (o *EncryptionSettingsBackendResponse) GetReturnUnavailableWhenDisabled() bool {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.ReturnUnavailableWhenDisabled
}

// GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
    return nil, false
	}
	return o.ReturnUnavailableWhenDisabled, true
}

// HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.
func (o *EncryptionSettingsBackendResponse) HasReturnUnavailableWhenDisabled() bool {
	if o != nil && !isNil(o.ReturnUnavailableWhenDisabled) {
		return true
	}

	return false
}

// SetReturnUnavailableWhenDisabled gets a reference to the given bool and assigns it to the ReturnUnavailableWhenDisabled field.
func (o *EncryptionSettingsBackendResponse) SetReturnUnavailableWhenDisabled(v bool) {
	o.ReturnUnavailableWhenDisabled = &v
}

// GetBackupFilePermissions returns the BackupFilePermissions field value if set, zero value otherwise.
func (o *EncryptionSettingsBackendResponse) GetBackupFilePermissions() string {
	if o == nil || isNil(o.BackupFilePermissions) {
		var ret string
		return ret
	}
	return *o.BackupFilePermissions
}

// GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetBackupFilePermissionsOk() (*string, bool) {
	if o == nil || isNil(o.BackupFilePermissions) {
    return nil, false
	}
	return o.BackupFilePermissions, true
}

// HasBackupFilePermissions returns a boolean if a field has been set.
func (o *EncryptionSettingsBackendResponse) HasBackupFilePermissions() bool {
	if o != nil && !isNil(o.BackupFilePermissions) {
		return true
	}

	return false
}

// SetBackupFilePermissions gets a reference to the given string and assigns it to the BackupFilePermissions field.
func (o *EncryptionSettingsBackendResponse) SetBackupFilePermissions(v string) {
	o.BackupFilePermissions = &v
}

// GetNotificationManager returns the NotificationManager field value if set, zero value otherwise.
func (o *EncryptionSettingsBackendResponse) GetNotificationManager() string {
	if o == nil || isNil(o.NotificationManager) {
		var ret string
		return ret
	}
	return *o.NotificationManager
}

// GetNotificationManagerOk returns a tuple with the NotificationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionSettingsBackendResponse) GetNotificationManagerOk() (*string, bool) {
	if o == nil || isNil(o.NotificationManager) {
    return nil, false
	}
	return o.NotificationManager, true
}

// HasNotificationManager returns a boolean if a field has been set.
func (o *EncryptionSettingsBackendResponse) HasNotificationManager() bool {
	if o != nil && !isNil(o.NotificationManager) {
		return true
	}

	return false
}

// SetNotificationManager gets a reference to the given string and assigns it to the NotificationManager field.
func (o *EncryptionSettingsBackendResponse) SetNotificationManager(v string) {
	o.NotificationManager = &v
}

func (o EncryptionSettingsBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["backendID"] = o.BackendID
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["writabilityMode"] = o.WritabilityMode
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

type NullableEncryptionSettingsBackendResponse struct {
	value *EncryptionSettingsBackendResponse
	isSet bool
}

func (v NullableEncryptionSettingsBackendResponse) Get() *EncryptionSettingsBackendResponse {
	return v.value
}

func (v *NullableEncryptionSettingsBackendResponse) Set(val *EncryptionSettingsBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionSettingsBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionSettingsBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionSettingsBackendResponse(val *EncryptionSettingsBackendResponse) *NullableEncryptionSettingsBackendResponse {
	return &NullableEncryptionSettingsBackendResponse{value: val, isSet: true}
}

func (v NullableEncryptionSettingsBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionSettingsBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


