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

// checks if the ConsoleJsonSyncLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsoleJsonSyncLogPublisherResponse{}

// ConsoleJsonSyncLogPublisherResponse struct for ConsoleJsonSyncLogPublisherResponse
type ConsoleJsonSyncLogPublisherResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumconsoleJsonSyncLogPublisherSchemaUrn         `json:"schemas"`
	// Name of the Log Publisher
	Id string `json:"id"`
	// Indicates whether the Console JSON Sync Log Publisher is enabled for use.
	Enabled        bool                                `json:"enabled"`
	OutputLocation *EnumlogPublisherOutputLocationProp `json:"outputLocation,omitempty"`
	// Indicates whether the JSON objects should use a multi-line representation (with each object field and array value on its own line) that may be easier for administrators to read, but each message will be larger (because of additional spaces and end-of-line markers), and it may be more difficult to consume and parse through some text-oriented tools.
	WriteMultiLineMessages *bool `json:"writeMultiLineMessages,omitempty"`
	// Indicates whether log messages should include the product name for the Directory Server.
	IncludeProductName *bool `json:"includeProductName,omitempty"`
	// Indicates whether log messages should include the instance name for the Directory Server.
	IncludeInstanceName *bool `json:"includeInstanceName,omitempty"`
	// Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted.
	IncludeStartupID *bool `json:"includeStartupID,omitempty"`
	// Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn=JVM Stack Trace,cn=monitor entry.
	IncludeThreadID *bool `json:"includeThreadID,omitempty"`
	// Specifies which Sync Pipes can log messages to this Sync Log Publisher.
	IncludeSyncPipe   []string                                `json:"includeSyncPipe,omitempty"`
	LoggedMessageType []EnumlogPublisherLoggedMessageTypeProp `json:"loggedMessageType,omitempty"`
	// A description for this Log Publisher
	Description          *string                                   `json:"description,omitempty"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewConsoleJsonSyncLogPublisherResponse instantiates a new ConsoleJsonSyncLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleJsonSyncLogPublisherResponse(schemas []EnumconsoleJsonSyncLogPublisherSchemaUrn, id string, enabled bool) *ConsoleJsonSyncLogPublisherResponse {
	this := ConsoleJsonSyncLogPublisherResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewConsoleJsonSyncLogPublisherResponseWithDefaults instantiates a new ConsoleJsonSyncLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleJsonSyncLogPublisherResponseWithDefaults() *ConsoleJsonSyncLogPublisherResponse {
	this := ConsoleJsonSyncLogPublisherResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *ConsoleJsonSyncLogPublisherResponse) GetSchemas() []EnumconsoleJsonSyncLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumconsoleJsonSyncLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetSchemasOk() ([]EnumconsoleJsonSyncLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ConsoleJsonSyncLogPublisherResponse) SetSchemas(v []EnumconsoleJsonSyncLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ConsoleJsonSyncLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsoleJsonSyncLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *ConsoleJsonSyncLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConsoleJsonSyncLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOutputLocation returns the OutputLocation field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetOutputLocation() EnumlogPublisherOutputLocationProp {
	if o == nil || IsNil(o.OutputLocation) {
		var ret EnumlogPublisherOutputLocationProp
		return ret
	}
	return *o.OutputLocation
}

// GetOutputLocationOk returns a tuple with the OutputLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetOutputLocationOk() (*EnumlogPublisherOutputLocationProp, bool) {
	if o == nil || IsNil(o.OutputLocation) {
		return nil, false
	}
	return o.OutputLocation, true
}

// HasOutputLocation returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasOutputLocation() bool {
	if o != nil && !IsNil(o.OutputLocation) {
		return true
	}

	return false
}

// SetOutputLocation gets a reference to the given EnumlogPublisherOutputLocationProp and assigns it to the OutputLocation field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetOutputLocation(v EnumlogPublisherOutputLocationProp) {
	o.OutputLocation = &v
}

// GetWriteMultiLineMessages returns the WriteMultiLineMessages field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetWriteMultiLineMessages() bool {
	if o == nil || IsNil(o.WriteMultiLineMessages) {
		var ret bool
		return ret
	}
	return *o.WriteMultiLineMessages
}

// GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.WriteMultiLineMessages) {
		return nil, false
	}
	return o.WriteMultiLineMessages, true
}

// HasWriteMultiLineMessages returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasWriteMultiLineMessages() bool {
	if o != nil && !IsNil(o.WriteMultiLineMessages) {
		return true
	}

	return false
}

// SetWriteMultiLineMessages gets a reference to the given bool and assigns it to the WriteMultiLineMessages field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetWriteMultiLineMessages(v bool) {
	o.WriteMultiLineMessages = &v
}

// GetIncludeProductName returns the IncludeProductName field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeProductName() bool {
	if o == nil || IsNil(o.IncludeProductName) {
		var ret bool
		return ret
	}
	return *o.IncludeProductName
}

// GetIncludeProductNameOk returns a tuple with the IncludeProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeProductName) {
		return nil, false
	}
	return o.IncludeProductName, true
}

// HasIncludeProductName returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasIncludeProductName() bool {
	if o != nil && !IsNil(o.IncludeProductName) {
		return true
	}

	return false
}

// SetIncludeProductName gets a reference to the given bool and assigns it to the IncludeProductName field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetIncludeProductName(v bool) {
	o.IncludeProductName = &v
}

// GetIncludeInstanceName returns the IncludeInstanceName field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeInstanceName() bool {
	if o == nil || IsNil(o.IncludeInstanceName) {
		var ret bool
		return ret
	}
	return *o.IncludeInstanceName
}

// GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeInstanceName) {
		return nil, false
	}
	return o.IncludeInstanceName, true
}

// HasIncludeInstanceName returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasIncludeInstanceName() bool {
	if o != nil && !IsNil(o.IncludeInstanceName) {
		return true
	}

	return false
}

// SetIncludeInstanceName gets a reference to the given bool and assigns it to the IncludeInstanceName field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetIncludeInstanceName(v bool) {
	o.IncludeInstanceName = &v
}

// GetIncludeStartupID returns the IncludeStartupID field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeStartupID() bool {
	if o == nil || IsNil(o.IncludeStartupID) {
		var ret bool
		return ret
	}
	return *o.IncludeStartupID
}

// GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeStartupID) {
		return nil, false
	}
	return o.IncludeStartupID, true
}

// HasIncludeStartupID returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasIncludeStartupID() bool {
	if o != nil && !IsNil(o.IncludeStartupID) {
		return true
	}

	return false
}

// SetIncludeStartupID gets a reference to the given bool and assigns it to the IncludeStartupID field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetIncludeStartupID(v bool) {
	o.IncludeStartupID = &v
}

// GetIncludeThreadID returns the IncludeThreadID field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeThreadID() bool {
	if o == nil || IsNil(o.IncludeThreadID) {
		var ret bool
		return ret
	}
	return *o.IncludeThreadID
}

// GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeThreadID) {
		return nil, false
	}
	return o.IncludeThreadID, true
}

// HasIncludeThreadID returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasIncludeThreadID() bool {
	if o != nil && !IsNil(o.IncludeThreadID) {
		return true
	}

	return false
}

// SetIncludeThreadID gets a reference to the given bool and assigns it to the IncludeThreadID field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetIncludeThreadID(v bool) {
	o.IncludeThreadID = &v
}

// GetIncludeSyncPipe returns the IncludeSyncPipe field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeSyncPipe() []string {
	if o == nil || IsNil(o.IncludeSyncPipe) {
		var ret []string
		return ret
	}
	return o.IncludeSyncPipe
}

// GetIncludeSyncPipeOk returns a tuple with the IncludeSyncPipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetIncludeSyncPipeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeSyncPipe) {
		return nil, false
	}
	return o.IncludeSyncPipe, true
}

// HasIncludeSyncPipe returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasIncludeSyncPipe() bool {
	if o != nil && !IsNil(o.IncludeSyncPipe) {
		return true
	}

	return false
}

// SetIncludeSyncPipe gets a reference to the given []string and assigns it to the IncludeSyncPipe field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetIncludeSyncPipe(v []string) {
	o.IncludeSyncPipe = v
}

// GetLoggedMessageType returns the LoggedMessageType field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetLoggedMessageType() []EnumlogPublisherLoggedMessageTypeProp {
	if o == nil || IsNil(o.LoggedMessageType) {
		var ret []EnumlogPublisherLoggedMessageTypeProp
		return ret
	}
	return o.LoggedMessageType
}

// GetLoggedMessageTypeOk returns a tuple with the LoggedMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetLoggedMessageTypeOk() ([]EnumlogPublisherLoggedMessageTypeProp, bool) {
	if o == nil || IsNil(o.LoggedMessageType) {
		return nil, false
	}
	return o.LoggedMessageType, true
}

// HasLoggedMessageType returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasLoggedMessageType() bool {
	if o != nil && !IsNil(o.LoggedMessageType) {
		return true
	}

	return false
}

// SetLoggedMessageType gets a reference to the given []EnumlogPublisherLoggedMessageTypeProp and assigns it to the LoggedMessageType field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetLoggedMessageType(v []EnumlogPublisherLoggedMessageTypeProp) {
	o.LoggedMessageType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *ConsoleJsonSyncLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *ConsoleJsonSyncLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *ConsoleJsonSyncLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o ConsoleJsonSyncLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsoleJsonSyncLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.OutputLocation) {
		toSerialize["outputLocation"] = o.OutputLocation
	}
	if !IsNil(o.WriteMultiLineMessages) {
		toSerialize["writeMultiLineMessages"] = o.WriteMultiLineMessages
	}
	if !IsNil(o.IncludeProductName) {
		toSerialize["includeProductName"] = o.IncludeProductName
	}
	if !IsNil(o.IncludeInstanceName) {
		toSerialize["includeInstanceName"] = o.IncludeInstanceName
	}
	if !IsNil(o.IncludeStartupID) {
		toSerialize["includeStartupID"] = o.IncludeStartupID
	}
	if !IsNil(o.IncludeThreadID) {
		toSerialize["includeThreadID"] = o.IncludeThreadID
	}
	if !IsNil(o.IncludeSyncPipe) {
		toSerialize["includeSyncPipe"] = o.IncludeSyncPipe
	}
	if !IsNil(o.LoggedMessageType) {
		toSerialize["loggedMessageType"] = o.LoggedMessageType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return toSerialize, nil
}

type NullableConsoleJsonSyncLogPublisherResponse struct {
	value *ConsoleJsonSyncLogPublisherResponse
	isSet bool
}

func (v NullableConsoleJsonSyncLogPublisherResponse) Get() *ConsoleJsonSyncLogPublisherResponse {
	return v.value
}

func (v *NullableConsoleJsonSyncLogPublisherResponse) Set(val *ConsoleJsonSyncLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleJsonSyncLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleJsonSyncLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleJsonSyncLogPublisherResponse(val *ConsoleJsonSyncLogPublisherResponse) *NullableConsoleJsonSyncLogPublisherResponse {
	return &NullableConsoleJsonSyncLogPublisherResponse{value: val, isSet: true}
}

func (v NullableConsoleJsonSyncLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleJsonSyncLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
