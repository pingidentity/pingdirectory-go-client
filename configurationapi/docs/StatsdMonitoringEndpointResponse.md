# StatsdMonitoringEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Monitoring Endpoint | 
**Schemas** | Pointer to [**[]EnumstatsdMonitoringEndpointSchemaUrn**](EnumstatsdMonitoringEndpointSchemaUrn.md) |  | [optional] 
**Hostname** | **string** | The name of the host where this StatsD Monitoring Endpoint should send metric data. | 
**ServerPort** | **int32** | Specifies the port number of the endpoint where metric data should be sent. | 
**ConnectionType** | [**EnummonitoringEndpointConnectionTypeProp**](EnummonitoringEndpointConnectionTypeProp.md) |  | 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL over TCP is to be used for connection-level security. | [optional] 
**AdditionalTags** | Pointer to **[]string** | Specifies any optional additional tags to include in StatsD messages. Any additional tags will be appended to the end of each StatsD message, separated by commas. Tags should be written in a [key]:[value] format (\&quot;host:server1\&quot;, for example). | [optional] 
**Enabled** | **bool** | Indicates whether this Monitoring Endpoint is enabled for use in the Directory Server. | 

## Methods

### NewStatsdMonitoringEndpointResponse

`func NewStatsdMonitoringEndpointResponse(id string, hostname string, serverPort int32, connectionType EnummonitoringEndpointConnectionTypeProp, enabled bool, ) *StatsdMonitoringEndpointResponse`

NewStatsdMonitoringEndpointResponse instantiates a new StatsdMonitoringEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsdMonitoringEndpointResponseWithDefaults

`func NewStatsdMonitoringEndpointResponseWithDefaults() *StatsdMonitoringEndpointResponse`

NewStatsdMonitoringEndpointResponseWithDefaults instantiates a new StatsdMonitoringEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *StatsdMonitoringEndpointResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StatsdMonitoringEndpointResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StatsdMonitoringEndpointResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StatsdMonitoringEndpointResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *StatsdMonitoringEndpointResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *StatsdMonitoringEndpointResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *StatsdMonitoringEndpointResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *StatsdMonitoringEndpointResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *StatsdMonitoringEndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsdMonitoringEndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsdMonitoringEndpointResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *StatsdMonitoringEndpointResponse) GetSchemas() []EnumstatsdMonitoringEndpointSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StatsdMonitoringEndpointResponse) GetSchemasOk() (*[]EnumstatsdMonitoringEndpointSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StatsdMonitoringEndpointResponse) SetSchemas(v []EnumstatsdMonitoringEndpointSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *StatsdMonitoringEndpointResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetHostname

`func (o *StatsdMonitoringEndpointResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *StatsdMonitoringEndpointResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *StatsdMonitoringEndpointResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetServerPort

`func (o *StatsdMonitoringEndpointResponse) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *StatsdMonitoringEndpointResponse) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *StatsdMonitoringEndpointResponse) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetConnectionType

`func (o *StatsdMonitoringEndpointResponse) GetConnectionType() EnummonitoringEndpointConnectionTypeProp`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *StatsdMonitoringEndpointResponse) GetConnectionTypeOk() (*EnummonitoringEndpointConnectionTypeProp, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *StatsdMonitoringEndpointResponse) SetConnectionType(v EnummonitoringEndpointConnectionTypeProp)`

SetConnectionType sets ConnectionType field to given value.


### GetTrustManagerProvider

`func (o *StatsdMonitoringEndpointResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *StatsdMonitoringEndpointResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *StatsdMonitoringEndpointResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *StatsdMonitoringEndpointResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetAdditionalTags

`func (o *StatsdMonitoringEndpointResponse) GetAdditionalTags() []string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *StatsdMonitoringEndpointResponse) GetAdditionalTagsOk() (*[]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *StatsdMonitoringEndpointResponse) SetAdditionalTags(v []string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *StatsdMonitoringEndpointResponse) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetEnabled

`func (o *StatsdMonitoringEndpointResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StatsdMonitoringEndpointResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StatsdMonitoringEndpointResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


