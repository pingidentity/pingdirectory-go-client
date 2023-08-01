# SnmpMasterAgentPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsnmpMasterAgentPluginSchemaUrn**](EnumsnmpMasterAgentPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**ListenAddress** | **string** | The IP address on which the SNMP agent will listen for client requests. | 
**ListenPort** | **int64** | The port number on which the SNMP agent will listen for client requests. | 
**AgentxTransport** | [**EnumpluginAgentxTransportProp**](EnumpluginAgentxTransportProp.md) |  | 
**AgentxListenAddress** | **string** | The IP address on which the SNMP agent will listen for sub-agent sessions. | 
**AgentxListenPort** | **int64** | The port number on which the SNMP agent will listen for sub-agent sessions. | 
**NotificationTargetAddress** | Pointer to **string** | The IP address of the target to which the SNMP agent should send notifications (traps). | [optional] 
**NotificationTargetPort** | Pointer to **int64** | The port number of the target to which the SNMP agent should send notifications (traps). | [optional] 
**AgentSNMPVersion** | [**[]EnumpluginAgentSNMPVersionProp**](EnumpluginAgentSNMPVersionProp.md) |  | 
**CommunityName** | **string** | The name of the community to use for the SNMP agent. | 
**AgentSecurityName** | Pointer to **string** | The security name (i.e., username) to use for the SNMP agent. This must be defined if SNMPv3 is to be used. | [optional] 
**AgentSecurityLevel** | Pointer to [**EnumpluginAgentSecurityLevelProp**](EnumpluginAgentSecurityLevelProp.md) |  | [optional] 
**AgentAuthenticationProtocol** | Pointer to [**EnumpluginAgentAuthenticationProtocolProp**](EnumpluginAgentAuthenticationProtocolProp.md) |  | [optional] 
**AgentAuthenticationPassphrase** | Pointer to **string** | The authentication passphrase to use for SNMPv3 if authentication is to be performed. | [optional] 
**AgentPrivacyProtocol** | Pointer to [**EnumpluginAgentPrivacyProtocolProp**](EnumpluginAgentPrivacyProtocolProp.md) |  | [optional] 
**AgentPrivacyPassphrase** | Pointer to **string** | The privacy passphrase to use for SNMPv3 if privacy is to be used. | [optional] 
**NumWorkerThreads** | **int64** | The number of worker threads to use to handle SNMP requests. | 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewSnmpMasterAgentPluginResponse

`func NewSnmpMasterAgentPluginResponse(schemas []EnumsnmpMasterAgentPluginSchemaUrn, id string, listenAddress string, listenPort int64, agentxTransport EnumpluginAgentxTransportProp, agentxListenAddress string, agentxListenPort int64, agentSNMPVersion []EnumpluginAgentSNMPVersionProp, communityName string, numWorkerThreads int64, enabled bool, ) *SnmpMasterAgentPluginResponse`

NewSnmpMasterAgentPluginResponse instantiates a new SnmpMasterAgentPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpMasterAgentPluginResponseWithDefaults

`func NewSnmpMasterAgentPluginResponseWithDefaults() *SnmpMasterAgentPluginResponse`

NewSnmpMasterAgentPluginResponseWithDefaults instantiates a new SnmpMasterAgentPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SnmpMasterAgentPluginResponse) GetSchemas() []EnumsnmpMasterAgentPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SnmpMasterAgentPluginResponse) GetSchemasOk() (*[]EnumsnmpMasterAgentPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SnmpMasterAgentPluginResponse) SetSchemas(v []EnumsnmpMasterAgentPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SnmpMasterAgentPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnmpMasterAgentPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnmpMasterAgentPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetListenAddress

`func (o *SnmpMasterAgentPluginResponse) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *SnmpMasterAgentPluginResponse) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *SnmpMasterAgentPluginResponse) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.


### GetListenPort

`func (o *SnmpMasterAgentPluginResponse) GetListenPort() int64`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *SnmpMasterAgentPluginResponse) GetListenPortOk() (*int64, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *SnmpMasterAgentPluginResponse) SetListenPort(v int64)`

SetListenPort sets ListenPort field to given value.


### GetAgentxTransport

`func (o *SnmpMasterAgentPluginResponse) GetAgentxTransport() EnumpluginAgentxTransportProp`

GetAgentxTransport returns the AgentxTransport field if non-nil, zero value otherwise.

### GetAgentxTransportOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentxTransportOk() (*EnumpluginAgentxTransportProp, bool)`

GetAgentxTransportOk returns a tuple with the AgentxTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxTransport

`func (o *SnmpMasterAgentPluginResponse) SetAgentxTransport(v EnumpluginAgentxTransportProp)`

SetAgentxTransport sets AgentxTransport field to given value.


### GetAgentxListenAddress

`func (o *SnmpMasterAgentPluginResponse) GetAgentxListenAddress() string`

GetAgentxListenAddress returns the AgentxListenAddress field if non-nil, zero value otherwise.

### GetAgentxListenAddressOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentxListenAddressOk() (*string, bool)`

GetAgentxListenAddressOk returns a tuple with the AgentxListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxListenAddress

`func (o *SnmpMasterAgentPluginResponse) SetAgentxListenAddress(v string)`

SetAgentxListenAddress sets AgentxListenAddress field to given value.


### GetAgentxListenPort

`func (o *SnmpMasterAgentPluginResponse) GetAgentxListenPort() int64`

GetAgentxListenPort returns the AgentxListenPort field if non-nil, zero value otherwise.

### GetAgentxListenPortOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentxListenPortOk() (*int64, bool)`

GetAgentxListenPortOk returns a tuple with the AgentxListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxListenPort

`func (o *SnmpMasterAgentPluginResponse) SetAgentxListenPort(v int64)`

SetAgentxListenPort sets AgentxListenPort field to given value.


### GetNotificationTargetAddress

`func (o *SnmpMasterAgentPluginResponse) GetNotificationTargetAddress() string`

GetNotificationTargetAddress returns the NotificationTargetAddress field if non-nil, zero value otherwise.

### GetNotificationTargetAddressOk

`func (o *SnmpMasterAgentPluginResponse) GetNotificationTargetAddressOk() (*string, bool)`

GetNotificationTargetAddressOk returns a tuple with the NotificationTargetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTargetAddress

`func (o *SnmpMasterAgentPluginResponse) SetNotificationTargetAddress(v string)`

SetNotificationTargetAddress sets NotificationTargetAddress field to given value.

### HasNotificationTargetAddress

`func (o *SnmpMasterAgentPluginResponse) HasNotificationTargetAddress() bool`

HasNotificationTargetAddress returns a boolean if a field has been set.

### GetNotificationTargetPort

`func (o *SnmpMasterAgentPluginResponse) GetNotificationTargetPort() int64`

GetNotificationTargetPort returns the NotificationTargetPort field if non-nil, zero value otherwise.

### GetNotificationTargetPortOk

`func (o *SnmpMasterAgentPluginResponse) GetNotificationTargetPortOk() (*int64, bool)`

GetNotificationTargetPortOk returns a tuple with the NotificationTargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTargetPort

`func (o *SnmpMasterAgentPluginResponse) SetNotificationTargetPort(v int64)`

SetNotificationTargetPort sets NotificationTargetPort field to given value.

### HasNotificationTargetPort

`func (o *SnmpMasterAgentPluginResponse) HasNotificationTargetPort() bool`

HasNotificationTargetPort returns a boolean if a field has been set.

### GetAgentSNMPVersion

`func (o *SnmpMasterAgentPluginResponse) GetAgentSNMPVersion() []EnumpluginAgentSNMPVersionProp`

GetAgentSNMPVersion returns the AgentSNMPVersion field if non-nil, zero value otherwise.

### GetAgentSNMPVersionOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentSNMPVersionOk() (*[]EnumpluginAgentSNMPVersionProp, bool)`

GetAgentSNMPVersionOk returns a tuple with the AgentSNMPVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSNMPVersion

`func (o *SnmpMasterAgentPluginResponse) SetAgentSNMPVersion(v []EnumpluginAgentSNMPVersionProp)`

SetAgentSNMPVersion sets AgentSNMPVersion field to given value.


### GetCommunityName

`func (o *SnmpMasterAgentPluginResponse) GetCommunityName() string`

GetCommunityName returns the CommunityName field if non-nil, zero value otherwise.

### GetCommunityNameOk

`func (o *SnmpMasterAgentPluginResponse) GetCommunityNameOk() (*string, bool)`

GetCommunityNameOk returns a tuple with the CommunityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityName

`func (o *SnmpMasterAgentPluginResponse) SetCommunityName(v string)`

SetCommunityName sets CommunityName field to given value.


### GetAgentSecurityName

`func (o *SnmpMasterAgentPluginResponse) GetAgentSecurityName() string`

GetAgentSecurityName returns the AgentSecurityName field if non-nil, zero value otherwise.

### GetAgentSecurityNameOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentSecurityNameOk() (*string, bool)`

GetAgentSecurityNameOk returns a tuple with the AgentSecurityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSecurityName

`func (o *SnmpMasterAgentPluginResponse) SetAgentSecurityName(v string)`

SetAgentSecurityName sets AgentSecurityName field to given value.

### HasAgentSecurityName

`func (o *SnmpMasterAgentPluginResponse) HasAgentSecurityName() bool`

HasAgentSecurityName returns a boolean if a field has been set.

### GetAgentSecurityLevel

`func (o *SnmpMasterAgentPluginResponse) GetAgentSecurityLevel() EnumpluginAgentSecurityLevelProp`

GetAgentSecurityLevel returns the AgentSecurityLevel field if non-nil, zero value otherwise.

### GetAgentSecurityLevelOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentSecurityLevelOk() (*EnumpluginAgentSecurityLevelProp, bool)`

GetAgentSecurityLevelOk returns a tuple with the AgentSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSecurityLevel

`func (o *SnmpMasterAgentPluginResponse) SetAgentSecurityLevel(v EnumpluginAgentSecurityLevelProp)`

SetAgentSecurityLevel sets AgentSecurityLevel field to given value.

### HasAgentSecurityLevel

`func (o *SnmpMasterAgentPluginResponse) HasAgentSecurityLevel() bool`

HasAgentSecurityLevel returns a boolean if a field has been set.

### GetAgentAuthenticationProtocol

`func (o *SnmpMasterAgentPluginResponse) GetAgentAuthenticationProtocol() EnumpluginAgentAuthenticationProtocolProp`

GetAgentAuthenticationProtocol returns the AgentAuthenticationProtocol field if non-nil, zero value otherwise.

### GetAgentAuthenticationProtocolOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentAuthenticationProtocolOk() (*EnumpluginAgentAuthenticationProtocolProp, bool)`

GetAgentAuthenticationProtocolOk returns a tuple with the AgentAuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAuthenticationProtocol

`func (o *SnmpMasterAgentPluginResponse) SetAgentAuthenticationProtocol(v EnumpluginAgentAuthenticationProtocolProp)`

SetAgentAuthenticationProtocol sets AgentAuthenticationProtocol field to given value.

### HasAgentAuthenticationProtocol

`func (o *SnmpMasterAgentPluginResponse) HasAgentAuthenticationProtocol() bool`

HasAgentAuthenticationProtocol returns a boolean if a field has been set.

### GetAgentAuthenticationPassphrase

`func (o *SnmpMasterAgentPluginResponse) GetAgentAuthenticationPassphrase() string`

GetAgentAuthenticationPassphrase returns the AgentAuthenticationPassphrase field if non-nil, zero value otherwise.

### GetAgentAuthenticationPassphraseOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentAuthenticationPassphraseOk() (*string, bool)`

GetAgentAuthenticationPassphraseOk returns a tuple with the AgentAuthenticationPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAuthenticationPassphrase

`func (o *SnmpMasterAgentPluginResponse) SetAgentAuthenticationPassphrase(v string)`

SetAgentAuthenticationPassphrase sets AgentAuthenticationPassphrase field to given value.

### HasAgentAuthenticationPassphrase

`func (o *SnmpMasterAgentPluginResponse) HasAgentAuthenticationPassphrase() bool`

HasAgentAuthenticationPassphrase returns a boolean if a field has been set.

### GetAgentPrivacyProtocol

`func (o *SnmpMasterAgentPluginResponse) GetAgentPrivacyProtocol() EnumpluginAgentPrivacyProtocolProp`

GetAgentPrivacyProtocol returns the AgentPrivacyProtocol field if non-nil, zero value otherwise.

### GetAgentPrivacyProtocolOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentPrivacyProtocolOk() (*EnumpluginAgentPrivacyProtocolProp, bool)`

GetAgentPrivacyProtocolOk returns a tuple with the AgentPrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPrivacyProtocol

`func (o *SnmpMasterAgentPluginResponse) SetAgentPrivacyProtocol(v EnumpluginAgentPrivacyProtocolProp)`

SetAgentPrivacyProtocol sets AgentPrivacyProtocol field to given value.

### HasAgentPrivacyProtocol

`func (o *SnmpMasterAgentPluginResponse) HasAgentPrivacyProtocol() bool`

HasAgentPrivacyProtocol returns a boolean if a field has been set.

### GetAgentPrivacyPassphrase

`func (o *SnmpMasterAgentPluginResponse) GetAgentPrivacyPassphrase() string`

GetAgentPrivacyPassphrase returns the AgentPrivacyPassphrase field if non-nil, zero value otherwise.

### GetAgentPrivacyPassphraseOk

`func (o *SnmpMasterAgentPluginResponse) GetAgentPrivacyPassphraseOk() (*string, bool)`

GetAgentPrivacyPassphraseOk returns a tuple with the AgentPrivacyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPrivacyPassphrase

`func (o *SnmpMasterAgentPluginResponse) SetAgentPrivacyPassphrase(v string)`

SetAgentPrivacyPassphrase sets AgentPrivacyPassphrase field to given value.

### HasAgentPrivacyPassphrase

`func (o *SnmpMasterAgentPluginResponse) HasAgentPrivacyPassphrase() bool`

HasAgentPrivacyPassphrase returns a boolean if a field has been set.

### GetNumWorkerThreads

`func (o *SnmpMasterAgentPluginResponse) GetNumWorkerThreads() int64`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *SnmpMasterAgentPluginResponse) GetNumWorkerThreadsOk() (*int64, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *SnmpMasterAgentPluginResponse) SetNumWorkerThreads(v int64)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.


### GetDescription

`func (o *SnmpMasterAgentPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnmpMasterAgentPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnmpMasterAgentPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnmpMasterAgentPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpMasterAgentPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpMasterAgentPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpMasterAgentPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *SnmpMasterAgentPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *SnmpMasterAgentPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *SnmpMasterAgentPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *SnmpMasterAgentPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *SnmpMasterAgentPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SnmpMasterAgentPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SnmpMasterAgentPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SnmpMasterAgentPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SnmpMasterAgentPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SnmpMasterAgentPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SnmpMasterAgentPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SnmpMasterAgentPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


