# AddSnmpSubagentPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsnmpSubagentPluginSchemaUrn**](EnumsnmpSubagentPluginSchemaUrn.md) |  | 
**ContextName** | Pointer to **string** | The SNMP context name for this sub-agent. The context name must not be longer than 30 ASCII characters. Each server in a topology must have a unique SNMP context name. | [optional] 
**AgentxAddress** | Pointer to **string** | The hostname or IP address of the SNMP master agent. | [optional] 
**AgentxPort** | Pointer to **int64** | The port number on which the SNMP master agent will be contacted. | [optional] 
**NumWorkerThreads** | Pointer to **int64** | The number of worker threads to use to handle SNMP requests. | [optional] 
**SessionTimeout** | Pointer to **string** | Specifies the maximum amount of time to wait for a session to the master agent to be established. | [optional] 
**ConnectRetryMaxWait** | Pointer to **string** | The maximum amount of time to wait between attempts to establish a connection to the master agent. | [optional] 
**PingInterval** | Pointer to **string** | The amount of time between consecutive pings sent by the sub-agent on its connection to the master agent. A value of zero disables the sending of pings by the sub-agent. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddSnmpSubagentPluginRequest

`func NewAddSnmpSubagentPluginRequest(schemas []EnumsnmpSubagentPluginSchemaUrn, enabled bool, pluginName string, ) *AddSnmpSubagentPluginRequest`

NewAddSnmpSubagentPluginRequest instantiates a new AddSnmpSubagentPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSnmpSubagentPluginRequestWithDefaults

`func NewAddSnmpSubagentPluginRequestWithDefaults() *AddSnmpSubagentPluginRequest`

NewAddSnmpSubagentPluginRequestWithDefaults instantiates a new AddSnmpSubagentPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSnmpSubagentPluginRequest) GetSchemas() []EnumsnmpSubagentPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSnmpSubagentPluginRequest) GetSchemasOk() (*[]EnumsnmpSubagentPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSnmpSubagentPluginRequest) SetSchemas(v []EnumsnmpSubagentPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetContextName

`func (o *AddSnmpSubagentPluginRequest) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *AddSnmpSubagentPluginRequest) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *AddSnmpSubagentPluginRequest) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *AddSnmpSubagentPluginRequest) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetAgentxAddress

`func (o *AddSnmpSubagentPluginRequest) GetAgentxAddress() string`

GetAgentxAddress returns the AgentxAddress field if non-nil, zero value otherwise.

### GetAgentxAddressOk

`func (o *AddSnmpSubagentPluginRequest) GetAgentxAddressOk() (*string, bool)`

GetAgentxAddressOk returns a tuple with the AgentxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxAddress

`func (o *AddSnmpSubagentPluginRequest) SetAgentxAddress(v string)`

SetAgentxAddress sets AgentxAddress field to given value.

### HasAgentxAddress

`func (o *AddSnmpSubagentPluginRequest) HasAgentxAddress() bool`

HasAgentxAddress returns a boolean if a field has been set.

### GetAgentxPort

`func (o *AddSnmpSubagentPluginRequest) GetAgentxPort() int64`

GetAgentxPort returns the AgentxPort field if non-nil, zero value otherwise.

### GetAgentxPortOk

`func (o *AddSnmpSubagentPluginRequest) GetAgentxPortOk() (*int64, bool)`

GetAgentxPortOk returns a tuple with the AgentxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxPort

`func (o *AddSnmpSubagentPluginRequest) SetAgentxPort(v int64)`

SetAgentxPort sets AgentxPort field to given value.

### HasAgentxPort

`func (o *AddSnmpSubagentPluginRequest) HasAgentxPort() bool`

HasAgentxPort returns a boolean if a field has been set.

### GetNumWorkerThreads

`func (o *AddSnmpSubagentPluginRequest) GetNumWorkerThreads() int64`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *AddSnmpSubagentPluginRequest) GetNumWorkerThreadsOk() (*int64, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *AddSnmpSubagentPluginRequest) SetNumWorkerThreads(v int64)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.

### HasNumWorkerThreads

`func (o *AddSnmpSubagentPluginRequest) HasNumWorkerThreads() bool`

HasNumWorkerThreads returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *AddSnmpSubagentPluginRequest) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *AddSnmpSubagentPluginRequest) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *AddSnmpSubagentPluginRequest) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *AddSnmpSubagentPluginRequest) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetConnectRetryMaxWait

`func (o *AddSnmpSubagentPluginRequest) GetConnectRetryMaxWait() string`

GetConnectRetryMaxWait returns the ConnectRetryMaxWait field if non-nil, zero value otherwise.

### GetConnectRetryMaxWaitOk

`func (o *AddSnmpSubagentPluginRequest) GetConnectRetryMaxWaitOk() (*string, bool)`

GetConnectRetryMaxWaitOk returns a tuple with the ConnectRetryMaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectRetryMaxWait

`func (o *AddSnmpSubagentPluginRequest) SetConnectRetryMaxWait(v string)`

SetConnectRetryMaxWait sets ConnectRetryMaxWait field to given value.

### HasConnectRetryMaxWait

`func (o *AddSnmpSubagentPluginRequest) HasConnectRetryMaxWait() bool`

HasConnectRetryMaxWait returns a boolean if a field has been set.

### GetPingInterval

`func (o *AddSnmpSubagentPluginRequest) GetPingInterval() string`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *AddSnmpSubagentPluginRequest) GetPingIntervalOk() (*string, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *AddSnmpSubagentPluginRequest) SetPingInterval(v string)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *AddSnmpSubagentPluginRequest) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetDescription

`func (o *AddSnmpSubagentPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSnmpSubagentPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSnmpSubagentPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSnmpSubagentPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSnmpSubagentPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSnmpSubagentPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSnmpSubagentPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddSnmpSubagentPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddSnmpSubagentPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddSnmpSubagentPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddSnmpSubagentPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddSnmpSubagentPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddSnmpSubagentPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddSnmpSubagentPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


