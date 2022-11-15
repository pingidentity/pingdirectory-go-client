# StatsdMonitoringEndpointShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumstatsdMonitoringEndpointSchemaUrn**](EnumstatsdMonitoringEndpointSchemaUrn.md) |  | [optional] 
**Hostname** | **string** | The name of the host where this StatsD Monitoring Endpoint should send metric data. | 
**ServerPort** | **int32** | Specifies the port number of the endpoint where metric data should be sent. | 
**ConnectionType** | [**EnummonitoringEndpointConnectionTypeProp**](EnummonitoringEndpointConnectionTypeProp.md) |  | 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL over TCP is to be used for connection-level security. | [optional] 
**AdditionalTags** | Pointer to **[]string** |  | [optional] 
**Enabled** | **bool** | Indicates whether this Monitoring Endpoint is enabled for use in the Directory Server. | 

## Methods

### NewStatsdMonitoringEndpointShared

`func NewStatsdMonitoringEndpointShared(hostname string, serverPort int32, connectionType EnummonitoringEndpointConnectionTypeProp, enabled bool, ) *StatsdMonitoringEndpointShared`

NewStatsdMonitoringEndpointShared instantiates a new StatsdMonitoringEndpointShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsdMonitoringEndpointSharedWithDefaults

`func NewStatsdMonitoringEndpointSharedWithDefaults() *StatsdMonitoringEndpointShared`

NewStatsdMonitoringEndpointSharedWithDefaults instantiates a new StatsdMonitoringEndpointShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *StatsdMonitoringEndpointShared) GetSchemas() []EnumstatsdMonitoringEndpointSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StatsdMonitoringEndpointShared) GetSchemasOk() (*[]EnumstatsdMonitoringEndpointSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StatsdMonitoringEndpointShared) SetSchemas(v []EnumstatsdMonitoringEndpointSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *StatsdMonitoringEndpointShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetHostname

`func (o *StatsdMonitoringEndpointShared) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *StatsdMonitoringEndpointShared) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *StatsdMonitoringEndpointShared) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetServerPort

`func (o *StatsdMonitoringEndpointShared) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *StatsdMonitoringEndpointShared) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *StatsdMonitoringEndpointShared) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetConnectionType

`func (o *StatsdMonitoringEndpointShared) GetConnectionType() EnummonitoringEndpointConnectionTypeProp`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *StatsdMonitoringEndpointShared) GetConnectionTypeOk() (*EnummonitoringEndpointConnectionTypeProp, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *StatsdMonitoringEndpointShared) SetConnectionType(v EnummonitoringEndpointConnectionTypeProp)`

SetConnectionType sets ConnectionType field to given value.


### GetTrustManagerProvider

`func (o *StatsdMonitoringEndpointShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *StatsdMonitoringEndpointShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *StatsdMonitoringEndpointShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *StatsdMonitoringEndpointShared) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetAdditionalTags

`func (o *StatsdMonitoringEndpointShared) GetAdditionalTags() []string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *StatsdMonitoringEndpointShared) GetAdditionalTagsOk() (*[]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *StatsdMonitoringEndpointShared) SetAdditionalTags(v []string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *StatsdMonitoringEndpointShared) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetEnabled

`func (o *StatsdMonitoringEndpointShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StatsdMonitoringEndpointShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StatsdMonitoringEndpointShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


