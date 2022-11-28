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

// AlertBackendResponse struct for AlertBackendResponse
type AlertBackendResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumalertBackendSchemaUrn `json:"schemas"`
	// Name of the Backend
	Id *string `json:"id,omitempty"`
	// Specifies a name to identify the associated backend.
	BackendID string `json:"backendID"`
	// Specifies the base DN(s) for the data that the backend handles.
	BaseDN []string `json:"baseDN"`
	// Specifies the path to the LDIF file that serves as the backing file for this backend.
	LdifFile string `json:"ldifFile"`
	// Specifies the maximum length of time that information about generated alerts should be maintained before they will be purged.
	AlertRetentionTime string `json:"alertRetentionTime"`
	// Specifies the maximum number of alerts that should be retained. If more alerts than this configured maximum are generated within the alert retention time, then the oldest alerts will be purged to achieve this maximum.
	MaxAlerts *int32 `json:"maxAlerts,omitempty"`
	DisabledAlertType []EnumbackendDisabledAlertTypeProp `json:"disabledAlertType,omitempty"`
	WritabilityMode EnumbackendWritabilityModeProp `json:"writabilityMode"`
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

// NewAlertBackendResponse instantiates a new AlertBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertBackendResponse(schemas []EnumalertBackendSchemaUrn, backendID string, baseDN []string, ldifFile string, alertRetentionTime string, writabilityMode EnumbackendWritabilityModeProp, enabled bool) *AlertBackendResponse {
	this := AlertBackendResponse{}
	this.Schemas = schemas
	this.BackendID = backendID
	this.BaseDN = baseDN
	this.LdifFile = ldifFile
	this.AlertRetentionTime = alertRetentionTime
	this.WritabilityMode = writabilityMode
	this.Enabled = enabled
	return &this
}

// NewAlertBackendResponseWithDefaults instantiates a new AlertBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertBackendResponseWithDefaults() *AlertBackendResponse {
	this := AlertBackendResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AlertBackendResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AlertBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *AlertBackendResponse) GetSchemas() []EnumalertBackendSchemaUrn {
	if o == nil {
		var ret []EnumalertBackendSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetSchemasOk() ([]EnumalertBackendSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AlertBackendResponse) SetSchemas(v []EnumalertBackendSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertBackendResponse) SetId(v string) {
	o.Id = &v
}

// GetBackendID returns the BackendID field value
func (o *AlertBackendResponse) GetBackendID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetBackendIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackendID, true
}

// SetBackendID sets field value
func (o *AlertBackendResponse) SetBackendID(v string) {
	o.BackendID = v
}

// GetBaseDN returns the BaseDN field value
func (o *AlertBackendResponse) GetBaseDN() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BaseDN, true
}

// SetBaseDN sets field value
func (o *AlertBackendResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetLdifFile returns the LdifFile field value
func (o *AlertBackendResponse) GetLdifFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdifFile
}

// GetLdifFileOk returns a tuple with the LdifFile field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetLdifFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LdifFile, true
}

// SetLdifFile sets field value
func (o *AlertBackendResponse) SetLdifFile(v string) {
	o.LdifFile = v
}

// GetAlertRetentionTime returns the AlertRetentionTime field value
func (o *AlertBackendResponse) GetAlertRetentionTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertRetentionTime
}

// GetAlertRetentionTimeOk returns a tuple with the AlertRetentionTime field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetAlertRetentionTimeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlertRetentionTime, true
}

// SetAlertRetentionTime sets field value
func (o *AlertBackendResponse) SetAlertRetentionTime(v string) {
	o.AlertRetentionTime = v
}

// GetMaxAlerts returns the MaxAlerts field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetMaxAlerts() int32 {
	if o == nil || isNil(o.MaxAlerts) {
		var ret int32
		return ret
	}
	return *o.MaxAlerts
}

// GetMaxAlertsOk returns a tuple with the MaxAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetMaxAlertsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxAlerts) {
    return nil, false
	}
	return o.MaxAlerts, true
}

// HasMaxAlerts returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasMaxAlerts() bool {
	if o != nil && !isNil(o.MaxAlerts) {
		return true
	}

	return false
}

// SetMaxAlerts gets a reference to the given int32 and assigns it to the MaxAlerts field.
func (o *AlertBackendResponse) SetMaxAlerts(v int32) {
	o.MaxAlerts = &v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetDisabledAlertType() []EnumbackendDisabledAlertTypeProp {
	if o == nil || isNil(o.DisabledAlertType) {
		var ret []EnumbackendDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetDisabledAlertTypeOk() ([]EnumbackendDisabledAlertTypeProp, bool) {
	if o == nil || isNil(o.DisabledAlertType) {
    return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasDisabledAlertType() bool {
	if o != nil && !isNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumbackendDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *AlertBackendResponse) SetDisabledAlertType(v []EnumbackendDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

// GetWritabilityMode returns the WritabilityMode field value
func (o *AlertBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp {
	if o == nil {
		var ret EnumbackendWritabilityModeProp
		return ret
	}

	return o.WritabilityMode
}

// GetWritabilityModeOk returns a tuple with the WritabilityMode field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WritabilityMode, true
}

// SetWritabilityMode sets field value
func (o *AlertBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp) {
	o.WritabilityMode = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertBackendResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AlertBackendResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AlertBackendResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetSetDegradedAlertWhenDisabled() bool {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.SetDegradedAlertWhenDisabled
}

// GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
    return nil, false
	}
	return o.SetDegradedAlertWhenDisabled, true
}

// HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasSetDegradedAlertWhenDisabled() bool {
	if o != nil && !isNil(o.SetDegradedAlertWhenDisabled) {
		return true
	}

	return false
}

// SetSetDegradedAlertWhenDisabled gets a reference to the given bool and assigns it to the SetDegradedAlertWhenDisabled field.
func (o *AlertBackendResponse) SetSetDegradedAlertWhenDisabled(v bool) {
	o.SetDegradedAlertWhenDisabled = &v
}

// GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetReturnUnavailableWhenDisabled() bool {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.ReturnUnavailableWhenDisabled
}

// GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
    return nil, false
	}
	return o.ReturnUnavailableWhenDisabled, true
}

// HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasReturnUnavailableWhenDisabled() bool {
	if o != nil && !isNil(o.ReturnUnavailableWhenDisabled) {
		return true
	}

	return false
}

// SetReturnUnavailableWhenDisabled gets a reference to the given bool and assigns it to the ReturnUnavailableWhenDisabled field.
func (o *AlertBackendResponse) SetReturnUnavailableWhenDisabled(v bool) {
	o.ReturnUnavailableWhenDisabled = &v
}

// GetBackupFilePermissions returns the BackupFilePermissions field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetBackupFilePermissions() string {
	if o == nil || isNil(o.BackupFilePermissions) {
		var ret string
		return ret
	}
	return *o.BackupFilePermissions
}

// GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetBackupFilePermissionsOk() (*string, bool) {
	if o == nil || isNil(o.BackupFilePermissions) {
    return nil, false
	}
	return o.BackupFilePermissions, true
}

// HasBackupFilePermissions returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasBackupFilePermissions() bool {
	if o != nil && !isNil(o.BackupFilePermissions) {
		return true
	}

	return false
}

// SetBackupFilePermissions gets a reference to the given string and assigns it to the BackupFilePermissions field.
func (o *AlertBackendResponse) SetBackupFilePermissions(v string) {
	o.BackupFilePermissions = &v
}

// GetNotificationManager returns the NotificationManager field value if set, zero value otherwise.
func (o *AlertBackendResponse) GetNotificationManager() string {
	if o == nil || isNil(o.NotificationManager) {
		var ret string
		return ret
	}
	return *o.NotificationManager
}

// GetNotificationManagerOk returns a tuple with the NotificationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertBackendResponse) GetNotificationManagerOk() (*string, bool) {
	if o == nil || isNil(o.NotificationManager) {
    return nil, false
	}
	return o.NotificationManager, true
}

// HasNotificationManager returns a boolean if a field has been set.
func (o *AlertBackendResponse) HasNotificationManager() bool {
	if o != nil && !isNil(o.NotificationManager) {
		return true
	}

	return false
}

// SetNotificationManager gets a reference to the given string and assigns it to the NotificationManager field.
func (o *AlertBackendResponse) SetNotificationManager(v string) {
	o.NotificationManager = &v
}

func (o AlertBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["backendID"] = o.BackendID
	}
	if true {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["ldifFile"] = o.LdifFile
	}
	if true {
		toSerialize["alertRetentionTime"] = o.AlertRetentionTime
	}
	if !isNil(o.MaxAlerts) {
		toSerialize["maxAlerts"] = o.MaxAlerts
	}
	if !isNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	if true {
		toSerialize["writabilityMode"] = o.WritabilityMode
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

type NullableAlertBackendResponse struct {
	value *AlertBackendResponse
	isSet bool
}

func (v NullableAlertBackendResponse) Get() *AlertBackendResponse {
	return v.value
}

func (v *NullableAlertBackendResponse) Set(val *AlertBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertBackendResponse(val *AlertBackendResponse) *NullableAlertBackendResponse {
	return &NullableAlertBackendResponse{value: val, isSet: true}
}

func (v NullableAlertBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


