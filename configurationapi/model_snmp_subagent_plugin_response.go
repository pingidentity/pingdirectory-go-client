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

// checks if the SnmpSubagentPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpSubagentPluginResponse{}

// SnmpSubagentPluginResponse struct for SnmpSubagentPluginResponse
type SnmpSubagentPluginResponse struct {
	// Name of the Plugin Root
	Id      string                            `json:"id"`
	Schemas []EnumsnmpSubagentPluginSchemaUrn `json:"schemas"`
	// The SNMP context name for this sub-agent. The context name must not be longer than 30 ASCII characters. Each server in a topology must have a unique SNMP context name.
	ContextName *string `json:"contextName,omitempty"`
	// The hostname or IP address of the SNMP master agent.
	AgentxAddress string `json:"agentxAddress"`
	// The port number on which the SNMP master agent will be contacted.
	AgentxPort int64 `json:"agentxPort"`
	// The number of worker threads to use to handle SNMP requests.
	NumWorkerThreads *int64 `json:"numWorkerThreads,omitempty"`
	// Specifies the maximum amount of time to wait for a session to the master agent to be established.
	SessionTimeout *string `json:"sessionTimeout,omitempty"`
	// The maximum amount of time to wait between attempts to establish a connection to the master agent.
	ConnectRetryMaxWait *string `json:"connectRetryMaxWait,omitempty"`
	// The amount of time between consecutive pings sent by the sub-agent on its connection to the master agent. A value of zero disables the sending of pings by the sub-agent.
	PingInterval *string `json:"pingInterval,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations                   *bool                                              `json:"invokeForInternalOperations,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewSnmpSubagentPluginResponse instantiates a new SnmpSubagentPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpSubagentPluginResponse(id string, schemas []EnumsnmpSubagentPluginSchemaUrn, agentxAddress string, agentxPort int64, enabled bool) *SnmpSubagentPluginResponse {
	this := SnmpSubagentPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AgentxAddress = agentxAddress
	this.AgentxPort = agentxPort
	this.Enabled = enabled
	return &this
}

// NewSnmpSubagentPluginResponseWithDefaults instantiates a new SnmpSubagentPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpSubagentPluginResponseWithDefaults() *SnmpSubagentPluginResponse {
	this := SnmpSubagentPluginResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SnmpSubagentPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SnmpSubagentPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SnmpSubagentPluginResponse) GetSchemas() []EnumsnmpSubagentPluginSchemaUrn {
	if o == nil {
		var ret []EnumsnmpSubagentPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetSchemasOk() ([]EnumsnmpSubagentPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SnmpSubagentPluginResponse) SetSchemas(v []EnumsnmpSubagentPluginSchemaUrn) {
	o.Schemas = v
}

// GetContextName returns the ContextName field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetContextName() string {
	if o == nil || IsNil(o.ContextName) {
		var ret string
		return ret
	}
	return *o.ContextName
}

// GetContextNameOk returns a tuple with the ContextName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetContextNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContextName) {
		return nil, false
	}
	return o.ContextName, true
}

// HasContextName returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasContextName() bool {
	if o != nil && !IsNil(o.ContextName) {
		return true
	}

	return false
}

// SetContextName gets a reference to the given string and assigns it to the ContextName field.
func (o *SnmpSubagentPluginResponse) SetContextName(v string) {
	o.ContextName = &v
}

// GetAgentxAddress returns the AgentxAddress field value
func (o *SnmpSubagentPluginResponse) GetAgentxAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentxAddress
}

// GetAgentxAddressOk returns a tuple with the AgentxAddress field value
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetAgentxAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentxAddress, true
}

// SetAgentxAddress sets field value
func (o *SnmpSubagentPluginResponse) SetAgentxAddress(v string) {
	o.AgentxAddress = v
}

// GetAgentxPort returns the AgentxPort field value
func (o *SnmpSubagentPluginResponse) GetAgentxPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AgentxPort
}

// GetAgentxPortOk returns a tuple with the AgentxPort field value
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetAgentxPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentxPort, true
}

// SetAgentxPort sets field value
func (o *SnmpSubagentPluginResponse) SetAgentxPort(v int64) {
	o.AgentxPort = v
}

// GetNumWorkerThreads returns the NumWorkerThreads field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetNumWorkerThreads() int64 {
	if o == nil || IsNil(o.NumWorkerThreads) {
		var ret int64
		return ret
	}
	return *o.NumWorkerThreads
}

// GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetNumWorkerThreadsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumWorkerThreads) {
		return nil, false
	}
	return o.NumWorkerThreads, true
}

// HasNumWorkerThreads returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasNumWorkerThreads() bool {
	if o != nil && !IsNil(o.NumWorkerThreads) {
		return true
	}

	return false
}

// SetNumWorkerThreads gets a reference to the given int64 and assigns it to the NumWorkerThreads field.
func (o *SnmpSubagentPluginResponse) SetNumWorkerThreads(v int64) {
	o.NumWorkerThreads = &v
}

// GetSessionTimeout returns the SessionTimeout field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetSessionTimeout() string {
	if o == nil || IsNil(o.SessionTimeout) {
		var ret string
		return ret
	}
	return *o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetSessionTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.SessionTimeout) {
		return nil, false
	}
	return o.SessionTimeout, true
}

// HasSessionTimeout returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasSessionTimeout() bool {
	if o != nil && !IsNil(o.SessionTimeout) {
		return true
	}

	return false
}

// SetSessionTimeout gets a reference to the given string and assigns it to the SessionTimeout field.
func (o *SnmpSubagentPluginResponse) SetSessionTimeout(v string) {
	o.SessionTimeout = &v
}

// GetConnectRetryMaxWait returns the ConnectRetryMaxWait field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetConnectRetryMaxWait() string {
	if o == nil || IsNil(o.ConnectRetryMaxWait) {
		var ret string
		return ret
	}
	return *o.ConnectRetryMaxWait
}

// GetConnectRetryMaxWaitOk returns a tuple with the ConnectRetryMaxWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetConnectRetryMaxWaitOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectRetryMaxWait) {
		return nil, false
	}
	return o.ConnectRetryMaxWait, true
}

// HasConnectRetryMaxWait returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasConnectRetryMaxWait() bool {
	if o != nil && !IsNil(o.ConnectRetryMaxWait) {
		return true
	}

	return false
}

// SetConnectRetryMaxWait gets a reference to the given string and assigns it to the ConnectRetryMaxWait field.
func (o *SnmpSubagentPluginResponse) SetConnectRetryMaxWait(v string) {
	o.ConnectRetryMaxWait = &v
}

// GetPingInterval returns the PingInterval field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetPingInterval() string {
	if o == nil || IsNil(o.PingInterval) {
		var ret string
		return ret
	}
	return *o.PingInterval
}

// GetPingIntervalOk returns a tuple with the PingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetPingIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.PingInterval) {
		return nil, false
	}
	return o.PingInterval, true
}

// HasPingInterval returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasPingInterval() bool {
	if o != nil && !IsNil(o.PingInterval) {
		return true
	}

	return false
}

// SetPingInterval gets a reference to the given string and assigns it to the PingInterval field.
func (o *SnmpSubagentPluginResponse) SetPingInterval(v string) {
	o.PingInterval = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SnmpSubagentPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SnmpSubagentPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SnmpSubagentPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *SnmpSubagentPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SnmpSubagentPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SnmpSubagentPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubagentPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SnmpSubagentPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SnmpSubagentPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o SnmpSubagentPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpSubagentPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.ContextName) {
		toSerialize["contextName"] = o.ContextName
	}
	toSerialize["agentxAddress"] = o.AgentxAddress
	toSerialize["agentxPort"] = o.AgentxPort
	if !IsNil(o.NumWorkerThreads) {
		toSerialize["numWorkerThreads"] = o.NumWorkerThreads
	}
	if !IsNil(o.SessionTimeout) {
		toSerialize["sessionTimeout"] = o.SessionTimeout
	}
	if !IsNil(o.ConnectRetryMaxWait) {
		toSerialize["connectRetryMaxWait"] = o.ConnectRetryMaxWait
	}
	if !IsNil(o.PingInterval) {
		toSerialize["pingInterval"] = o.PingInterval
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableSnmpSubagentPluginResponse struct {
	value *SnmpSubagentPluginResponse
	isSet bool
}

func (v NullableSnmpSubagentPluginResponse) Get() *SnmpSubagentPluginResponse {
	return v.value
}

func (v *NullableSnmpSubagentPluginResponse) Set(val *SnmpSubagentPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpSubagentPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpSubagentPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpSubagentPluginResponse(val *SnmpSubagentPluginResponse) *NullableSnmpSubagentPluginResponse {
	return &NullableSnmpSubagentPluginResponse{value: val, isSet: true}
}

func (v NullableSnmpSubagentPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpSubagentPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
