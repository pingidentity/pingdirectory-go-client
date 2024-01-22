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

// checks if the PurgeExpiredDataPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurgeExpiredDataPluginResponse{}

// PurgeExpiredDataPluginResponse struct for PurgeExpiredDataPluginResponse
type PurgeExpiredDataPluginResponse struct {
	Schemas []EnumpurgeExpiredDataPluginSchemaUrn `json:"schemas"`
	// The LDAP attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted.
	DatetimeAttribute string `json:"datetimeAttribute"`
	// The top-level JSON field within the configured datetime-attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted.
	DatetimeJSONField *string                      `json:"datetimeJSONField,omitempty"`
	DatetimeFormat    EnumpluginDatetimeFormatProp `json:"datetimeFormat"`
	// When the datetime-format property is configured with a value of \"custom\", this specifies the format (using a string compatible with the java.text.SimpleDateFormat class) that will be used to search for expired data.
	CustomDatetimeFormat *string `json:"customDatetimeFormat,omitempty"`
	// Specifies the time zone to use when generating a date string using the configured custom-datetime-format value. The provided value must be accepted by java.util.TimeZone.getTimeZone.
	CustomTimezone *string `json:"customTimezone,omitempty"`
	// The duration to wait after the value specified in datetime-attribute (and optionally datetime-json-field) before purging the data.
	ExpirationOffset string                       `json:"expirationOffset"`
	PurgeBehavior    *EnumpluginPurgeBehaviorProp `json:"purgeBehavior,omitempty"`
	// Only entries located within the subtree specified by this base DN are eligible for purging.
	BaseDN *string `json:"baseDN,omitempty"`
	// Only entries that match this LDAP filter will be eligible for having data purged.
	Filter *string `json:"filter,omitempty"`
	// This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information).
	PollingInterval string `json:"pollingInterval"`
	// This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling.
	MaxUpdatesPerSecond int64 `json:"maxUpdatesPerSecond"`
	// In a replicated environment, this determines the order in which peer servers should attempt to purge data.
	PeerServerPriorityIndex *int64 `json:"peerServerPriorityIndex,omitempty"`
	// The number of threads used to delete expired entries.
	NumDeleteThreads int64 `json:"numDeleteThreads"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id string `json:"id"`
}

// NewPurgeExpiredDataPluginResponse instantiates a new PurgeExpiredDataPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeExpiredDataPluginResponse(schemas []EnumpurgeExpiredDataPluginSchemaUrn, datetimeAttribute string, datetimeFormat EnumpluginDatetimeFormatProp, expirationOffset string, pollingInterval string, maxUpdatesPerSecond int64, numDeleteThreads int64, enabled bool, id string) *PurgeExpiredDataPluginResponse {
	this := PurgeExpiredDataPluginResponse{}
	this.Schemas = schemas
	this.DatetimeAttribute = datetimeAttribute
	this.DatetimeFormat = datetimeFormat
	this.ExpirationOffset = expirationOffset
	this.PollingInterval = pollingInterval
	this.MaxUpdatesPerSecond = maxUpdatesPerSecond
	this.NumDeleteThreads = numDeleteThreads
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewPurgeExpiredDataPluginResponseWithDefaults instantiates a new PurgeExpiredDataPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeExpiredDataPluginResponseWithDefaults() *PurgeExpiredDataPluginResponse {
	this := PurgeExpiredDataPluginResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *PurgeExpiredDataPluginResponse) GetSchemas() []EnumpurgeExpiredDataPluginSchemaUrn {
	if o == nil {
		var ret []EnumpurgeExpiredDataPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetSchemasOk() ([]EnumpurgeExpiredDataPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PurgeExpiredDataPluginResponse) SetSchemas(v []EnumpurgeExpiredDataPluginSchemaUrn) {
	o.Schemas = v
}

// GetDatetimeAttribute returns the DatetimeAttribute field value
func (o *PurgeExpiredDataPluginResponse) GetDatetimeAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatetimeAttribute
}

// GetDatetimeAttributeOk returns a tuple with the DatetimeAttribute field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetDatetimeAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatetimeAttribute, true
}

// SetDatetimeAttribute sets field value
func (o *PurgeExpiredDataPluginResponse) SetDatetimeAttribute(v string) {
	o.DatetimeAttribute = v
}

// GetDatetimeJSONField returns the DatetimeJSONField field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetDatetimeJSONField() string {
	if o == nil || IsNil(o.DatetimeJSONField) {
		var ret string
		return ret
	}
	return *o.DatetimeJSONField
}

// GetDatetimeJSONFieldOk returns a tuple with the DatetimeJSONField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetDatetimeJSONFieldOk() (*string, bool) {
	if o == nil || IsNil(o.DatetimeJSONField) {
		return nil, false
	}
	return o.DatetimeJSONField, true
}

// HasDatetimeJSONField returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasDatetimeJSONField() bool {
	if o != nil && !IsNil(o.DatetimeJSONField) {
		return true
	}

	return false
}

// SetDatetimeJSONField gets a reference to the given string and assigns it to the DatetimeJSONField field.
func (o *PurgeExpiredDataPluginResponse) SetDatetimeJSONField(v string) {
	o.DatetimeJSONField = &v
}

// GetDatetimeFormat returns the DatetimeFormat field value
func (o *PurgeExpiredDataPluginResponse) GetDatetimeFormat() EnumpluginDatetimeFormatProp {
	if o == nil {
		var ret EnumpluginDatetimeFormatProp
		return ret
	}

	return o.DatetimeFormat
}

// GetDatetimeFormatOk returns a tuple with the DatetimeFormat field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetDatetimeFormatOk() (*EnumpluginDatetimeFormatProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatetimeFormat, true
}

// SetDatetimeFormat sets field value
func (o *PurgeExpiredDataPluginResponse) SetDatetimeFormat(v EnumpluginDatetimeFormatProp) {
	o.DatetimeFormat = v
}

// GetCustomDatetimeFormat returns the CustomDatetimeFormat field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetCustomDatetimeFormat() string {
	if o == nil || IsNil(o.CustomDatetimeFormat) {
		var ret string
		return ret
	}
	return *o.CustomDatetimeFormat
}

// GetCustomDatetimeFormatOk returns a tuple with the CustomDatetimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetCustomDatetimeFormatOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDatetimeFormat) {
		return nil, false
	}
	return o.CustomDatetimeFormat, true
}

// HasCustomDatetimeFormat returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasCustomDatetimeFormat() bool {
	if o != nil && !IsNil(o.CustomDatetimeFormat) {
		return true
	}

	return false
}

// SetCustomDatetimeFormat gets a reference to the given string and assigns it to the CustomDatetimeFormat field.
func (o *PurgeExpiredDataPluginResponse) SetCustomDatetimeFormat(v string) {
	o.CustomDatetimeFormat = &v
}

// GetCustomTimezone returns the CustomTimezone field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetCustomTimezone() string {
	if o == nil || IsNil(o.CustomTimezone) {
		var ret string
		return ret
	}
	return *o.CustomTimezone
}

// GetCustomTimezoneOk returns a tuple with the CustomTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetCustomTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTimezone) {
		return nil, false
	}
	return o.CustomTimezone, true
}

// HasCustomTimezone returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasCustomTimezone() bool {
	if o != nil && !IsNil(o.CustomTimezone) {
		return true
	}

	return false
}

// SetCustomTimezone gets a reference to the given string and assigns it to the CustomTimezone field.
func (o *PurgeExpiredDataPluginResponse) SetCustomTimezone(v string) {
	o.CustomTimezone = &v
}

// GetExpirationOffset returns the ExpirationOffset field value
func (o *PurgeExpiredDataPluginResponse) GetExpirationOffset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationOffset
}

// GetExpirationOffsetOk returns a tuple with the ExpirationOffset field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetExpirationOffsetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationOffset, true
}

// SetExpirationOffset sets field value
func (o *PurgeExpiredDataPluginResponse) SetExpirationOffset(v string) {
	o.ExpirationOffset = v
}

// GetPurgeBehavior returns the PurgeBehavior field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetPurgeBehavior() EnumpluginPurgeBehaviorProp {
	if o == nil || IsNil(o.PurgeBehavior) {
		var ret EnumpluginPurgeBehaviorProp
		return ret
	}
	return *o.PurgeBehavior
}

// GetPurgeBehaviorOk returns a tuple with the PurgeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetPurgeBehaviorOk() (*EnumpluginPurgeBehaviorProp, bool) {
	if o == nil || IsNil(o.PurgeBehavior) {
		return nil, false
	}
	return o.PurgeBehavior, true
}

// HasPurgeBehavior returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasPurgeBehavior() bool {
	if o != nil && !IsNil(o.PurgeBehavior) {
		return true
	}

	return false
}

// SetPurgeBehavior gets a reference to the given EnumpluginPurgeBehaviorProp and assigns it to the PurgeBehavior field.
func (o *PurgeExpiredDataPluginResponse) SetPurgeBehavior(v EnumpluginPurgeBehaviorProp) {
	o.PurgeBehavior = &v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetBaseDN() string {
	if o == nil || IsNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetBaseDNOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *PurgeExpiredDataPluginResponse) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *PurgeExpiredDataPluginResponse) SetFilter(v string) {
	o.Filter = &v
}

// GetPollingInterval returns the PollingInterval field value
func (o *PurgeExpiredDataPluginResponse) GetPollingInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PollingInterval
}

// GetPollingIntervalOk returns a tuple with the PollingInterval field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetPollingIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollingInterval, true
}

// SetPollingInterval sets field value
func (o *PurgeExpiredDataPluginResponse) SetPollingInterval(v string) {
	o.PollingInterval = v
}

// GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field value
func (o *PurgeExpiredDataPluginResponse) GetMaxUpdatesPerSecond() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxUpdatesPerSecond
}

// GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetMaxUpdatesPerSecondOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxUpdatesPerSecond, true
}

// SetMaxUpdatesPerSecond sets field value
func (o *PurgeExpiredDataPluginResponse) SetMaxUpdatesPerSecond(v int64) {
	o.MaxUpdatesPerSecond = v
}

// GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetPeerServerPriorityIndex() int64 {
	if o == nil || IsNil(o.PeerServerPriorityIndex) {
		var ret int64
		return ret
	}
	return *o.PeerServerPriorityIndex
}

// GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetPeerServerPriorityIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.PeerServerPriorityIndex) {
		return nil, false
	}
	return o.PeerServerPriorityIndex, true
}

// HasPeerServerPriorityIndex returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasPeerServerPriorityIndex() bool {
	if o != nil && !IsNil(o.PeerServerPriorityIndex) {
		return true
	}

	return false
}

// SetPeerServerPriorityIndex gets a reference to the given int64 and assigns it to the PeerServerPriorityIndex field.
func (o *PurgeExpiredDataPluginResponse) SetPeerServerPriorityIndex(v int64) {
	o.PeerServerPriorityIndex = &v
}

// GetNumDeleteThreads returns the NumDeleteThreads field value
func (o *PurgeExpiredDataPluginResponse) GetNumDeleteThreads() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumDeleteThreads
}

// GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetNumDeleteThreadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumDeleteThreads, true
}

// SetNumDeleteThreads sets field value
func (o *PurgeExpiredDataPluginResponse) SetNumDeleteThreads(v int64) {
	o.NumDeleteThreads = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PurgeExpiredDataPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PurgeExpiredDataPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PurgeExpiredDataPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PurgeExpiredDataPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PurgeExpiredDataPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PurgeExpiredDataPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PurgeExpiredDataPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *PurgeExpiredDataPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PurgeExpiredDataPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PurgeExpiredDataPluginResponse) SetId(v string) {
	o.Id = v
}

func (o PurgeExpiredDataPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurgeExpiredDataPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["datetimeAttribute"] = o.DatetimeAttribute
	if !IsNil(o.DatetimeJSONField) {
		toSerialize["datetimeJSONField"] = o.DatetimeJSONField
	}
	toSerialize["datetimeFormat"] = o.DatetimeFormat
	if !IsNil(o.CustomDatetimeFormat) {
		toSerialize["customDatetimeFormat"] = o.CustomDatetimeFormat
	}
	if !IsNil(o.CustomTimezone) {
		toSerialize["customTimezone"] = o.CustomTimezone
	}
	toSerialize["expirationOffset"] = o.ExpirationOffset
	if !IsNil(o.PurgeBehavior) {
		toSerialize["purgeBehavior"] = o.PurgeBehavior
	}
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["pollingInterval"] = o.PollingInterval
	toSerialize["maxUpdatesPerSecond"] = o.MaxUpdatesPerSecond
	if !IsNil(o.PeerServerPriorityIndex) {
		toSerialize["peerServerPriorityIndex"] = o.PeerServerPriorityIndex
	}
	toSerialize["numDeleteThreads"] = o.NumDeleteThreads
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullablePurgeExpiredDataPluginResponse struct {
	value *PurgeExpiredDataPluginResponse
	isSet bool
}

func (v NullablePurgeExpiredDataPluginResponse) Get() *PurgeExpiredDataPluginResponse {
	return v.value
}

func (v *NullablePurgeExpiredDataPluginResponse) Set(val *PurgeExpiredDataPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeExpiredDataPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeExpiredDataPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeExpiredDataPluginResponse(val *PurgeExpiredDataPluginResponse) *NullablePurgeExpiredDataPluginResponse {
	return &NullablePurgeExpiredDataPluginResponse{value: val, isSet: true}
}

func (v NullablePurgeExpiredDataPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeExpiredDataPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
