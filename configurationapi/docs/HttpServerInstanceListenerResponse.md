# HttpServerInstanceListenerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumhttpServerInstanceListenerSchemaUrn**](EnumhttpServerInstanceListenerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Server Instance Listener | 
**ListenAddress** | Pointer to **string** | If the server is listening on a particular address different from the hostname, then this property may be used to specify the address on which to listen for connections from HTTP clients. | [optional] 
**ServerHTTPPort** | Pointer to **int32** | The TCP port number on which the HTTP server is listening. | [optional] 
**ConnectionSecurity** | Pointer to [**EnumserverInstanceListenerHttpConnectionSecurityProp**](EnumserverInstanceListenerHttpConnectionSecurityProp.md) |  | [optional] 
**Purpose** | Pointer to [**[]EnumserverInstanceListenerPurposeProp**](EnumserverInstanceListenerPurposeProp.md) |  | [optional] 

## Methods

### NewHttpServerInstanceListenerResponse

`func NewHttpServerInstanceListenerResponse(schemas []EnumhttpServerInstanceListenerSchemaUrn, id string, ) *HttpServerInstanceListenerResponse`

NewHttpServerInstanceListenerResponse instantiates a new HttpServerInstanceListenerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpServerInstanceListenerResponseWithDefaults

`func NewHttpServerInstanceListenerResponseWithDefaults() *HttpServerInstanceListenerResponse`

NewHttpServerInstanceListenerResponseWithDefaults instantiates a new HttpServerInstanceListenerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *HttpServerInstanceListenerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpServerInstanceListenerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpServerInstanceListenerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpServerInstanceListenerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HttpServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpServerInstanceListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HttpServerInstanceListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *HttpServerInstanceListenerResponse) GetSchemas() []EnumhttpServerInstanceListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpServerInstanceListenerResponse) GetSchemasOk() (*[]EnumhttpServerInstanceListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpServerInstanceListenerResponse) SetSchemas(v []EnumhttpServerInstanceListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *HttpServerInstanceListenerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpServerInstanceListenerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpServerInstanceListenerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetListenAddress

`func (o *HttpServerInstanceListenerResponse) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *HttpServerInstanceListenerResponse) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *HttpServerInstanceListenerResponse) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *HttpServerInstanceListenerResponse) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetServerHTTPPort

`func (o *HttpServerInstanceListenerResponse) GetServerHTTPPort() int32`

GetServerHTTPPort returns the ServerHTTPPort field if non-nil, zero value otherwise.

### GetServerHTTPPortOk

`func (o *HttpServerInstanceListenerResponse) GetServerHTTPPortOk() (*int32, bool)`

GetServerHTTPPortOk returns a tuple with the ServerHTTPPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHTTPPort

`func (o *HttpServerInstanceListenerResponse) SetServerHTTPPort(v int32)`

SetServerHTTPPort sets ServerHTTPPort field to given value.

### HasServerHTTPPort

`func (o *HttpServerInstanceListenerResponse) HasServerHTTPPort() bool`

HasServerHTTPPort returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *HttpServerInstanceListenerResponse) GetConnectionSecurity() EnumserverInstanceListenerHttpConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *HttpServerInstanceListenerResponse) GetConnectionSecurityOk() (*EnumserverInstanceListenerHttpConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *HttpServerInstanceListenerResponse) SetConnectionSecurity(v EnumserverInstanceListenerHttpConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *HttpServerInstanceListenerResponse) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetPurpose

`func (o *HttpServerInstanceListenerResponse) GetPurpose() []EnumserverInstanceListenerPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *HttpServerInstanceListenerResponse) GetPurposeOk() (*[]EnumserverInstanceListenerPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *HttpServerInstanceListenerResponse) SetPurpose(v []EnumserverInstanceListenerPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *HttpServerInstanceListenerResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


