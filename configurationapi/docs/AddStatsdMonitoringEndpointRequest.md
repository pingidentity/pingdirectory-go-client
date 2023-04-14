# AddStatsdMonitoringEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | **string** | Name of the new Monitoring Endpoint | 
**Schemas** | Pointer to [**[]EnumstatsdMonitoringEndpointSchemaUrn**](EnumstatsdMonitoringEndpointSchemaUrn.md) |  | [optional] 
**Hostname** | **string** | The name of the host where this StatsD Monitoring Endpoint should send metric data. | 
**ServerPort** | Pointer to **int64** | Specifies the port number of the endpoint where metric data should be sent. | [optional] 
**ConnectionType** | Pointer to [**EnummonitoringEndpointConnectionTypeProp**](EnummonitoringEndpointConnectionTypeProp.md) |  | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL over TCP is to be used for connection-level security. | [optional] 
**AdditionalTags** | Pointer to **[]string** | Specifies any optional additional tags to include in StatsD messages. Any additional tags will be appended to the end of each StatsD message, separated by commas. Tags should be written in a [key]:[value] format (\&quot;host:server1\&quot;, for example). | [optional] 
**Enabled** | **bool** | Indicates whether this Monitoring Endpoint is enabled for use in the Directory Server. | 

## Methods

### NewAddStatsdMonitoringEndpointRequest

`func NewAddStatsdMonitoringEndpointRequest(endpointName string, hostname string, enabled bool, ) *AddStatsdMonitoringEndpointRequest`

NewAddStatsdMonitoringEndpointRequest instantiates a new AddStatsdMonitoringEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStatsdMonitoringEndpointRequestWithDefaults

`func NewAddStatsdMonitoringEndpointRequestWithDefaults() *AddStatsdMonitoringEndpointRequest`

NewAddStatsdMonitoringEndpointRequestWithDefaults instantiates a new AddStatsdMonitoringEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *AddStatsdMonitoringEndpointRequest) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *AddStatsdMonitoringEndpointRequest) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *AddStatsdMonitoringEndpointRequest) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.


### GetSchemas

`func (o *AddStatsdMonitoringEndpointRequest) GetSchemas() []EnumstatsdMonitoringEndpointSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddStatsdMonitoringEndpointRequest) GetSchemasOk() (*[]EnumstatsdMonitoringEndpointSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddStatsdMonitoringEndpointRequest) SetSchemas(v []EnumstatsdMonitoringEndpointSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddStatsdMonitoringEndpointRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetHostname

`func (o *AddStatsdMonitoringEndpointRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddStatsdMonitoringEndpointRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddStatsdMonitoringEndpointRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetServerPort

`func (o *AddStatsdMonitoringEndpointRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddStatsdMonitoringEndpointRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddStatsdMonitoringEndpointRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddStatsdMonitoringEndpointRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetConnectionType

`func (o *AddStatsdMonitoringEndpointRequest) GetConnectionType() EnummonitoringEndpointConnectionTypeProp`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AddStatsdMonitoringEndpointRequest) GetConnectionTypeOk() (*EnummonitoringEndpointConnectionTypeProp, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AddStatsdMonitoringEndpointRequest) SetConnectionType(v EnummonitoringEndpointConnectionTypeProp)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *AddStatsdMonitoringEndpointRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddStatsdMonitoringEndpointRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddStatsdMonitoringEndpointRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddStatsdMonitoringEndpointRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddStatsdMonitoringEndpointRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetAdditionalTags

`func (o *AddStatsdMonitoringEndpointRequest) GetAdditionalTags() []string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *AddStatsdMonitoringEndpointRequest) GetAdditionalTagsOk() (*[]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *AddStatsdMonitoringEndpointRequest) SetAdditionalTags(v []string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *AddStatsdMonitoringEndpointRequest) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetEnabled

`func (o *AddStatsdMonitoringEndpointRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddStatsdMonitoringEndpointRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddStatsdMonitoringEndpointRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


