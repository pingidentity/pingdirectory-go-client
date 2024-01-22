# LdapServerInstanceListenerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumldapServerInstanceListenerSchemaUrn**](EnumldapServerInstanceListenerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Server Instance Listener | 
**ServerLDAPPort** | Pointer to **int64** | The TCP port number on which the LDAP server is listening. | [optional] 
**ConnectionSecurity** | Pointer to [**EnumserverInstanceListenerLdapConnectionSecurityProp**](EnumserverInstanceListenerLdapConnectionSecurityProp.md) |  | [optional] 
**ListenerCertificate** | Pointer to **string** | The public component of the certificate that the listener is expected to present to clients. When establishing a connection to this server, only the certificate(s) listed here will be trusted. | [optional] 
**Purpose** | Pointer to [**[]EnumserverInstanceListenerPurposeProp**](EnumserverInstanceListenerPurposeProp.md) |  | [optional] 

## Methods

### NewLdapServerInstanceListenerResponse

`func NewLdapServerInstanceListenerResponse(schemas []EnumldapServerInstanceListenerSchemaUrn, id string, ) *LdapServerInstanceListenerResponse`

NewLdapServerInstanceListenerResponse instantiates a new LdapServerInstanceListenerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapServerInstanceListenerResponseWithDefaults

`func NewLdapServerInstanceListenerResponseWithDefaults() *LdapServerInstanceListenerResponse`

NewLdapServerInstanceListenerResponseWithDefaults instantiates a new LdapServerInstanceListenerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *LdapServerInstanceListenerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdapServerInstanceListenerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdapServerInstanceListenerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdapServerInstanceListenerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdapServerInstanceListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapServerInstanceListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdapServerInstanceListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *LdapServerInstanceListenerResponse) GetSchemas() []EnumldapServerInstanceListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapServerInstanceListenerResponse) GetSchemasOk() (*[]EnumldapServerInstanceListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapServerInstanceListenerResponse) SetSchemas(v []EnumldapServerInstanceListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *LdapServerInstanceListenerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapServerInstanceListenerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapServerInstanceListenerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetServerLDAPPort

`func (o *LdapServerInstanceListenerResponse) GetServerLDAPPort() int64`

GetServerLDAPPort returns the ServerLDAPPort field if non-nil, zero value otherwise.

### GetServerLDAPPortOk

`func (o *LdapServerInstanceListenerResponse) GetServerLDAPPortOk() (*int64, bool)`

GetServerLDAPPortOk returns a tuple with the ServerLDAPPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLDAPPort

`func (o *LdapServerInstanceListenerResponse) SetServerLDAPPort(v int64)`

SetServerLDAPPort sets ServerLDAPPort field to given value.

### HasServerLDAPPort

`func (o *LdapServerInstanceListenerResponse) HasServerLDAPPort() bool`

HasServerLDAPPort returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *LdapServerInstanceListenerResponse) GetConnectionSecurity() EnumserverInstanceListenerLdapConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *LdapServerInstanceListenerResponse) GetConnectionSecurityOk() (*EnumserverInstanceListenerLdapConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *LdapServerInstanceListenerResponse) SetConnectionSecurity(v EnumserverInstanceListenerLdapConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *LdapServerInstanceListenerResponse) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetListenerCertificate

`func (o *LdapServerInstanceListenerResponse) GetListenerCertificate() string`

GetListenerCertificate returns the ListenerCertificate field if non-nil, zero value otherwise.

### GetListenerCertificateOk

`func (o *LdapServerInstanceListenerResponse) GetListenerCertificateOk() (*string, bool)`

GetListenerCertificateOk returns a tuple with the ListenerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerCertificate

`func (o *LdapServerInstanceListenerResponse) SetListenerCertificate(v string)`

SetListenerCertificate sets ListenerCertificate field to given value.

### HasListenerCertificate

`func (o *LdapServerInstanceListenerResponse) HasListenerCertificate() bool`

HasListenerCertificate returns a boolean if a field has been set.

### GetPurpose

`func (o *LdapServerInstanceListenerResponse) GetPurpose() []EnumserverInstanceListenerPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LdapServerInstanceListenerResponse) GetPurposeOk() (*[]EnumserverInstanceListenerPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LdapServerInstanceListenerResponse) SetPurpose(v []EnumserverInstanceListenerPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *LdapServerInstanceListenerResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


