# GetServerInstanceListener200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumhttpServerInstanceListenerSchemaUrn**](EnumhttpServerInstanceListenerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Server Instance | 
**ServerLDAPPort** | Pointer to **int32** | The TCP port number on which the LDAP server is listening. | [optional] 
**ConnectionSecurity** | Pointer to [**EnumserverInstanceListenerHttpConnectionSecurityProp**](EnumserverInstanceListenerHttpConnectionSecurityProp.md) |  | [optional] 
**ListenerCertificate** | Pointer to **string** | The public component of the certificate that the listener is expected to present to clients. When establishing a connection to this server, only the certificate(s) listed here will be trusted. | [optional] 
**Purpose** | Pointer to [**[]EnumserverInstanceListenerPurposeProp**](EnumserverInstanceListenerPurposeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ListenAddress** | Pointer to **string** | If the server is listening on a particular address different from the hostname, then this property may be used to specify the address on which to listen for connections from HTTP clients. | [optional] 
**ServerHTTPPort** | Pointer to **int32** | The TCP port number on which the HTTP server is listening. | [optional] 

## Methods

### NewGetServerInstanceListener200Response

`func NewGetServerInstanceListener200Response(schemas []EnumhttpServerInstanceListenerSchemaUrn, id string, ) *GetServerInstanceListener200Response`

NewGetServerInstanceListener200Response instantiates a new GetServerInstanceListener200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerInstanceListener200ResponseWithDefaults

`func NewGetServerInstanceListener200ResponseWithDefaults() *GetServerInstanceListener200Response`

NewGetServerInstanceListener200ResponseWithDefaults instantiates a new GetServerInstanceListener200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetServerInstanceListener200Response) GetSchemas() []EnumhttpServerInstanceListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetServerInstanceListener200Response) GetSchemasOk() (*[]EnumhttpServerInstanceListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetServerInstanceListener200Response) SetSchemas(v []EnumhttpServerInstanceListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetServerInstanceListener200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServerInstanceListener200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServerInstanceListener200Response) SetId(v string)`

SetId sets Id field to given value.


### GetServerLDAPPort

`func (o *GetServerInstanceListener200Response) GetServerLDAPPort() int32`

GetServerLDAPPort returns the ServerLDAPPort field if non-nil, zero value otherwise.

### GetServerLDAPPortOk

`func (o *GetServerInstanceListener200Response) GetServerLDAPPortOk() (*int32, bool)`

GetServerLDAPPortOk returns a tuple with the ServerLDAPPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLDAPPort

`func (o *GetServerInstanceListener200Response) SetServerLDAPPort(v int32)`

SetServerLDAPPort sets ServerLDAPPort field to given value.

### HasServerLDAPPort

`func (o *GetServerInstanceListener200Response) HasServerLDAPPort() bool`

HasServerLDAPPort returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *GetServerInstanceListener200Response) GetConnectionSecurity() EnumserverInstanceListenerHttpConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *GetServerInstanceListener200Response) GetConnectionSecurityOk() (*EnumserverInstanceListenerHttpConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *GetServerInstanceListener200Response) SetConnectionSecurity(v EnumserverInstanceListenerHttpConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *GetServerInstanceListener200Response) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetListenerCertificate

`func (o *GetServerInstanceListener200Response) GetListenerCertificate() string`

GetListenerCertificate returns the ListenerCertificate field if non-nil, zero value otherwise.

### GetListenerCertificateOk

`func (o *GetServerInstanceListener200Response) GetListenerCertificateOk() (*string, bool)`

GetListenerCertificateOk returns a tuple with the ListenerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerCertificate

`func (o *GetServerInstanceListener200Response) SetListenerCertificate(v string)`

SetListenerCertificate sets ListenerCertificate field to given value.

### HasListenerCertificate

`func (o *GetServerInstanceListener200Response) HasListenerCertificate() bool`

HasListenerCertificate returns a boolean if a field has been set.

### GetPurpose

`func (o *GetServerInstanceListener200Response) GetPurpose() []EnumserverInstanceListenerPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *GetServerInstanceListener200Response) GetPurposeOk() (*[]EnumserverInstanceListenerPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *GetServerInstanceListener200Response) SetPurpose(v []EnumserverInstanceListenerPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *GetServerInstanceListener200Response) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMeta

`func (o *GetServerInstanceListener200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServerInstanceListener200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServerInstanceListener200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetServerInstanceListener200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetServerInstanceListener200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetServerInstanceListener200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetServerInstanceListener200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetServerInstanceListener200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetListenAddress

`func (o *GetServerInstanceListener200Response) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *GetServerInstanceListener200Response) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *GetServerInstanceListener200Response) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *GetServerInstanceListener200Response) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetServerHTTPPort

`func (o *GetServerInstanceListener200Response) GetServerHTTPPort() int32`

GetServerHTTPPort returns the ServerHTTPPort field if non-nil, zero value otherwise.

### GetServerHTTPPortOk

`func (o *GetServerInstanceListener200Response) GetServerHTTPPortOk() (*int32, bool)`

GetServerHTTPPortOk returns a tuple with the ServerHTTPPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHTTPPort

`func (o *GetServerInstanceListener200Response) SetServerHTTPPort(v int32)`

SetServerHTTPPort sets ServerHTTPPort field to given value.

### HasServerHTTPPort

`func (o *GetServerInstanceListener200Response) HasServerHTTPPort() bool`

HasServerHTTPPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


