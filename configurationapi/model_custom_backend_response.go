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

// checks if the CustomBackendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomBackendResponse{}

// CustomBackendResponse struct for CustomBackendResponse
type CustomBackendResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumcustomBackendSchemaUrn                       `json:"schemas"`
	// Name of the Backend
	Id string `json:"id"`
	// Specifies a name to identify the associated backend.
	BackendID string `json:"backendID"`
	// A description for this Backend
	Description *string `json:"description,omitempty"`
	// Indicates whether the backend is enabled in the server.
	Enabled bool `json:"enabled"`
	// Specifies the base DN(s) for the data that the backend handles.
	BaseDN          []string                       `json:"baseDN"`
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

// NewCustomBackendResponse instantiates a new CustomBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBackendResponse(schemas []EnumcustomBackendSchemaUrn, id string, backendID string, enabled bool, baseDN []string, writabilityMode EnumbackendWritabilityModeProp) *CustomBackendResponse {
	this := CustomBackendResponse{}
	this.Schemas = schemas
	this.Id = id
	this.BackendID = backendID
	this.Enabled = enabled
	this.BaseDN = baseDN
	this.WritabilityMode = writabilityMode
	return &this
}

// NewCustomBackendResponseWithDefaults instantiates a new CustomBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBackendResponseWithDefaults() *CustomBackendResponse {
	this := CustomBackendResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CustomBackendResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CustomBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *CustomBackendResponse) GetSchemas() []EnumcustomBackendSchemaUrn {
	if o == nil {
		var ret []EnumcustomBackendSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetSchemasOk() ([]EnumcustomBackendSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CustomBackendResponse) SetSchemas(v []EnumcustomBackendSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *CustomBackendResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomBackendResponse) SetId(v string) {
	o.Id = v
}

// GetBackendID returns the BackendID field value
func (o *CustomBackendResponse) GetBackendID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetBackendIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendID, true
}

// SetBackendID sets field value
func (o *CustomBackendResponse) SetBackendID(v string) {
	o.BackendID = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomBackendResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CustomBackendResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CustomBackendResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetBaseDN returns the BaseDN field value
func (o *CustomBackendResponse) GetBaseDN() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseDN, true
}

// SetBaseDN sets field value
func (o *CustomBackendResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetWritabilityMode returns the WritabilityMode field value
func (o *CustomBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp {
	if o == nil {
		var ret EnumbackendWritabilityModeProp
		return ret
	}

	return o.WritabilityMode
}

// GetWritabilityModeOk returns a tuple with the WritabilityMode field value
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WritabilityMode, true
}

// SetWritabilityMode sets field value
func (o *CustomBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp) {
	o.WritabilityMode = v
}

// GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetSetDegradedAlertWhenDisabled() bool {
	if o == nil || IsNil(o.SetDegradedAlertWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.SetDegradedAlertWhenDisabled
}

// GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SetDegradedAlertWhenDisabled) {
		return nil, false
	}
	return o.SetDegradedAlertWhenDisabled, true
}

// HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasSetDegradedAlertWhenDisabled() bool {
	if o != nil && !IsNil(o.SetDegradedAlertWhenDisabled) {
		return true
	}

	return false
}

// SetSetDegradedAlertWhenDisabled gets a reference to the given bool and assigns it to the SetDegradedAlertWhenDisabled field.
func (o *CustomBackendResponse) SetSetDegradedAlertWhenDisabled(v bool) {
	o.SetDegradedAlertWhenDisabled = &v
}

// GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetReturnUnavailableWhenDisabled() bool {
	if o == nil || IsNil(o.ReturnUnavailableWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.ReturnUnavailableWhenDisabled
}

// GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnUnavailableWhenDisabled) {
		return nil, false
	}
	return o.ReturnUnavailableWhenDisabled, true
}

// HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasReturnUnavailableWhenDisabled() bool {
	if o != nil && !IsNil(o.ReturnUnavailableWhenDisabled) {
		return true
	}

	return false
}

// SetReturnUnavailableWhenDisabled gets a reference to the given bool and assigns it to the ReturnUnavailableWhenDisabled field.
func (o *CustomBackendResponse) SetReturnUnavailableWhenDisabled(v bool) {
	o.ReturnUnavailableWhenDisabled = &v
}

// GetBackupFilePermissions returns the BackupFilePermissions field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetBackupFilePermissions() string {
	if o == nil || IsNil(o.BackupFilePermissions) {
		var ret string
		return ret
	}
	return *o.BackupFilePermissions
}

// GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetBackupFilePermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.BackupFilePermissions) {
		return nil, false
	}
	return o.BackupFilePermissions, true
}

// HasBackupFilePermissions returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasBackupFilePermissions() bool {
	if o != nil && !IsNil(o.BackupFilePermissions) {
		return true
	}

	return false
}

// SetBackupFilePermissions gets a reference to the given string and assigns it to the BackupFilePermissions field.
func (o *CustomBackendResponse) SetBackupFilePermissions(v string) {
	o.BackupFilePermissions = &v
}

// GetNotificationManager returns the NotificationManager field value if set, zero value otherwise.
func (o *CustomBackendResponse) GetNotificationManager() string {
	if o == nil || IsNil(o.NotificationManager) {
		var ret string
		return ret
	}
	return *o.NotificationManager
}

// GetNotificationManagerOk returns a tuple with the NotificationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBackendResponse) GetNotificationManagerOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationManager) {
		return nil, false
	}
	return o.NotificationManager, true
}

// HasNotificationManager returns a boolean if a field has been set.
func (o *CustomBackendResponse) HasNotificationManager() bool {
	if o != nil && !IsNil(o.NotificationManager) {
		return true
	}

	return false
}

// SetNotificationManager gets a reference to the given string and assigns it to the NotificationManager field.
func (o *CustomBackendResponse) SetNotificationManager(v string) {
	o.NotificationManager = &v
}

func (o CustomBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomBackendResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["baseDN"] = o.BaseDN
	toSerialize["writabilityMode"] = o.WritabilityMode
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

type NullableCustomBackendResponse struct {
	value *CustomBackendResponse
	isSet bool
}

func (v NullableCustomBackendResponse) Get() *CustomBackendResponse {
	return v.value
}

func (v *NullableCustomBackendResponse) Set(val *CustomBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBackendResponse(val *CustomBackendResponse) *NullableCustomBackendResponse {
	return &NullableCustomBackendResponse{value: val, isSet: true}
}

func (v NullableCustomBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
