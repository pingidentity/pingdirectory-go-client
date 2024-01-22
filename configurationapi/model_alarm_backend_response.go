/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AlarmBackendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmBackendResponse{}

// AlarmBackendResponse struct for AlarmBackendResponse
type AlarmBackendResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumalarmBackendSchemaUrn                        `json:"schemas"`
	// Name of the Backend
	Id string `json:"id"`
	// Specifies a name to identify the associated backend.
	BackendID string `json:"backendID"`
	// Specifies the base DN(s) for the data that the backend handles.
	BaseDN []string `json:"baseDN"`
	// Specifies the path to the LDIF file that serves as the backing file for this backend.
	LdifFile string `json:"ldifFile"`
	// Specifies the maximum length of time that information about raised alarms should be maintained before they will be purged.
	AlarmRetentionTime string `json:"alarmRetentionTime"`
	// Specifies the maximum number of alarms that should be retained. If more alarms than this configured maximum are generated within the alarm retention time, then the oldest alarms will be purged to achieve this maximum. Only alarms at normal severity will be purged.
	MaxAlarms       *int64                         `json:"maxAlarms,omitempty"`
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

// NewAlarmBackendResponse instantiates a new AlarmBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmBackendResponse(schemas []EnumalarmBackendSchemaUrn, id string, backendID string, baseDN []string, ldifFile string, alarmRetentionTime string, writabilityMode EnumbackendWritabilityModeProp, enabled bool) *AlarmBackendResponse {
	this := AlarmBackendResponse{}
	this.Schemas = schemas
	this.Id = id
	this.BackendID = backendID
	this.BaseDN = baseDN
	this.LdifFile = ldifFile
	this.AlarmRetentionTime = alarmRetentionTime
	this.WritabilityMode = writabilityMode
	this.Enabled = enabled
	return &this
}

// NewAlarmBackendResponseWithDefaults instantiates a new AlarmBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmBackendResponseWithDefaults() *AlarmBackendResponse {
	this := AlarmBackendResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AlarmBackendResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AlarmBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *AlarmBackendResponse) GetSchemas() []EnumalarmBackendSchemaUrn {
	if o == nil {
		var ret []EnumalarmBackendSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetSchemasOk() ([]EnumalarmBackendSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AlarmBackendResponse) SetSchemas(v []EnumalarmBackendSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *AlarmBackendResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AlarmBackendResponse) SetId(v string) {
	o.Id = v
}

// GetBackendID returns the BackendID field value
func (o *AlarmBackendResponse) GetBackendID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetBackendIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendID, true
}

// SetBackendID sets field value
func (o *AlarmBackendResponse) SetBackendID(v string) {
	o.BackendID = v
}

// GetBaseDN returns the BaseDN field value
func (o *AlarmBackendResponse) GetBaseDN() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseDN, true
}

// SetBaseDN sets field value
func (o *AlarmBackendResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetLdifFile returns the LdifFile field value
func (o *AlarmBackendResponse) GetLdifFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdifFile
}

// GetLdifFileOk returns a tuple with the LdifFile field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetLdifFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdifFile, true
}

// SetLdifFile sets field value
func (o *AlarmBackendResponse) SetLdifFile(v string) {
	o.LdifFile = v
}

// GetAlarmRetentionTime returns the AlarmRetentionTime field value
func (o *AlarmBackendResponse) GetAlarmRetentionTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmRetentionTime
}

// GetAlarmRetentionTimeOk returns a tuple with the AlarmRetentionTime field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetAlarmRetentionTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmRetentionTime, true
}

// SetAlarmRetentionTime sets field value
func (o *AlarmBackendResponse) SetAlarmRetentionTime(v string) {
	o.AlarmRetentionTime = v
}

// GetMaxAlarms returns the MaxAlarms field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetMaxAlarms() int64 {
	if o == nil || IsNil(o.MaxAlarms) {
		var ret int64
		return ret
	}
	return *o.MaxAlarms
}

// GetMaxAlarmsOk returns a tuple with the MaxAlarms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetMaxAlarmsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxAlarms) {
		return nil, false
	}
	return o.MaxAlarms, true
}

// HasMaxAlarms returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasMaxAlarms() bool {
	if o != nil && !IsNil(o.MaxAlarms) {
		return true
	}

	return false
}

// SetMaxAlarms gets a reference to the given int64 and assigns it to the MaxAlarms field.
func (o *AlarmBackendResponse) SetMaxAlarms(v int64) {
	o.MaxAlarms = &v
}

// GetWritabilityMode returns the WritabilityMode field value
func (o *AlarmBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp {
	if o == nil {
		var ret EnumbackendWritabilityModeProp
		return ret
	}

	return o.WritabilityMode
}

// GetWritabilityModeOk returns a tuple with the WritabilityMode field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WritabilityMode, true
}

// SetWritabilityMode sets field value
func (o *AlarmBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp) {
	o.WritabilityMode = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlarmBackendResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AlarmBackendResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AlarmBackendResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetSetDegradedAlertWhenDisabled() bool {
	if o == nil || IsNil(o.SetDegradedAlertWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.SetDegradedAlertWhenDisabled
}

// GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SetDegradedAlertWhenDisabled) {
		return nil, false
	}
	return o.SetDegradedAlertWhenDisabled, true
}

// HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasSetDegradedAlertWhenDisabled() bool {
	if o != nil && !IsNil(o.SetDegradedAlertWhenDisabled) {
		return true
	}

	return false
}

// SetSetDegradedAlertWhenDisabled gets a reference to the given bool and assigns it to the SetDegradedAlertWhenDisabled field.
func (o *AlarmBackendResponse) SetSetDegradedAlertWhenDisabled(v bool) {
	o.SetDegradedAlertWhenDisabled = &v
}

// GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetReturnUnavailableWhenDisabled() bool {
	if o == nil || IsNil(o.ReturnUnavailableWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.ReturnUnavailableWhenDisabled
}

// GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnUnavailableWhenDisabled) {
		return nil, false
	}
	return o.ReturnUnavailableWhenDisabled, true
}

// HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasReturnUnavailableWhenDisabled() bool {
	if o != nil && !IsNil(o.ReturnUnavailableWhenDisabled) {
		return true
	}

	return false
}

// SetReturnUnavailableWhenDisabled gets a reference to the given bool and assigns it to the ReturnUnavailableWhenDisabled field.
func (o *AlarmBackendResponse) SetReturnUnavailableWhenDisabled(v bool) {
	o.ReturnUnavailableWhenDisabled = &v
}

// GetBackupFilePermissions returns the BackupFilePermissions field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetBackupFilePermissions() string {
	if o == nil || IsNil(o.BackupFilePermissions) {
		var ret string
		return ret
	}
	return *o.BackupFilePermissions
}

// GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetBackupFilePermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.BackupFilePermissions) {
		return nil, false
	}
	return o.BackupFilePermissions, true
}

// HasBackupFilePermissions returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasBackupFilePermissions() bool {
	if o != nil && !IsNil(o.BackupFilePermissions) {
		return true
	}

	return false
}

// SetBackupFilePermissions gets a reference to the given string and assigns it to the BackupFilePermissions field.
func (o *AlarmBackendResponse) SetBackupFilePermissions(v string) {
	o.BackupFilePermissions = &v
}

// GetNotificationManager returns the NotificationManager field value if set, zero value otherwise.
func (o *AlarmBackendResponse) GetNotificationManager() string {
	if o == nil || IsNil(o.NotificationManager) {
		var ret string
		return ret
	}
	return *o.NotificationManager
}

// GetNotificationManagerOk returns a tuple with the NotificationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmBackendResponse) GetNotificationManagerOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationManager) {
		return nil, false
	}
	return o.NotificationManager, true
}

// HasNotificationManager returns a boolean if a field has been set.
func (o *AlarmBackendResponse) HasNotificationManager() bool {
	if o != nil && !IsNil(o.NotificationManager) {
		return true
	}

	return false
}

// SetNotificationManager gets a reference to the given string and assigns it to the NotificationManager field.
func (o *AlarmBackendResponse) SetNotificationManager(v string) {
	o.NotificationManager = &v
}

func (o AlarmBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmBackendResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["backendID"] = o.BackendID
	toSerialize["baseDN"] = o.BaseDN
	toSerialize["ldifFile"] = o.LdifFile
	toSerialize["alarmRetentionTime"] = o.AlarmRetentionTime
	if !IsNil(o.MaxAlarms) {
		toSerialize["maxAlarms"] = o.MaxAlarms
	}
	toSerialize["writabilityMode"] = o.WritabilityMode
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.SetDegradedAlertWhenDisabled) {
		toSerialize["setDegradedAlertWhenDisabled"] = o.SetDegradedAlertWhenDisabled
	}
	if !IsNil(o.ReturnUnavailableWhenDisabled) {
		toSerialize["returnUnavailableWhenDisabled"] = o.ReturnUnavailableWhenDisabled
	}
	if !IsNil(o.BackupFilePermissions) {
		toSerialize["backupFilePermissions"] = o.BackupFilePermissions
	}
	if !IsNil(o.NotificationManager) {
		toSerialize["notificationManager"] = o.NotificationManager
	}
	return toSerialize, nil
}

type NullableAlarmBackendResponse struct {
	value *AlarmBackendResponse
	isSet bool
}

func (v NullableAlarmBackendResponse) Get() *AlarmBackendResponse {
	return v.value
}

func (v *NullableAlarmBackendResponse) Set(val *AlarmBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmBackendResponse(val *AlarmBackendResponse) *NullableAlarmBackendResponse {
	return &NullableAlarmBackendResponse{value: val, isSet: true}
}

func (v NullableAlarmBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
