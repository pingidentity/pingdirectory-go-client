# LdapPassThroughAuthenticationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Pass Through Authentication Handler | 
**Schemas** | [**[]EnumldapPassThroughAuthenticationHandlerSchemaUrn**](EnumldapPassThroughAuthenticationHandlerSchemaUrn.md) |  | 
**Server** | **[]string** | Specifies the LDAP external server(s) to which authentication attempts should be forwarded. | 
**ServerAccessMode** | [**EnumpassThroughAuthenticationHandlerServerAccessModeProp**](EnumpassThroughAuthenticationHandlerServerAccessModeProp.md) |  | 
**DnMap** | Pointer to **[]string** | Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers. | [optional] 
**BindDNPattern** | Pointer to **string** | A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \&quot;cn&#x3D;{cn},ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot; indicates that the remote bind DN should be constructed from the text \&quot;cn&#x3D;\&quot; followed by the value of the local entry&#39;s cn attribute followed by the text \&quot;ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot;. If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \&quot;{seeAlso}\&quot; would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name. | [optional] 
**SearchBaseDN** | Pointer to **string** | The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN. | [optional] 
**SearchFilterPattern** | Pointer to **string** | A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \&quot;(mail&#x3D;{uid:ldapFilterEscape}@example.com)\&quot; would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \&quot;@example.com\&quot; in the mail attribute in the external server. Note that the \&quot;ldapFilterEscape\&quot; modifier should almost always be used with attributes specified in the pattern. | [optional] 
**InitialConnections** | **int32** | Specifies the initial number of connections to establish to each external server against which authentication may be attempted. | 
**MaxConnections** | **int32** | Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property. | 
**UseLocation** | Pointer to **bool** | Indicates whether to take server locations into account when prioritizing the servers to use for pass-through authentication attempts. | [optional] 
**MaximumAllowedLocalResponseTime** | Pointer to **string** | The maximum length of time to wait for a response from an external server in the same location as this Directory Server before considering it unavailable. | [optional] 
**MaximumAllowedNonlocalResponseTime** | Pointer to **string** | The maximum length of time to wait for a response from an external server in a different location from this Directory Server before considering it unavailable. | [optional] 
**UsePasswordPolicyControl** | Pointer to **bool** | Indicates whether to include the password policy request control (as defined in draft-behera-ldap-password-policy-10) in bind requests sent to the external server. | [optional] 
**Description** | Pointer to **string** | A description for this Pass Through Authentication Handler | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLdapPassThroughAuthenticationHandlerResponse

`func NewLdapPassThroughAuthenticationHandlerResponse(id string, schemas []EnumldapPassThroughAuthenticationHandlerSchemaUrn, server []string, serverAccessMode EnumpassThroughAuthenticationHandlerServerAccessModeProp, initialConnections int32, maxConnections int32, ) *LdapPassThroughAuthenticationHandlerResponse`

NewLdapPassThroughAuthenticationHandlerResponse instantiates a new LdapPassThroughAuthenticationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapPassThroughAuthenticationHandlerResponseWithDefaults

`func NewLdapPassThroughAuthenticationHandlerResponseWithDefaults() *LdapPassThroughAuthenticationHandlerResponse`

NewLdapPassThroughAuthenticationHandlerResponseWithDefaults instantiates a new LdapPassThroughAuthenticationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetSchemas() []EnumldapPassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetSchemasOk() (*[]EnumldapPassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetSchemas(v []EnumldapPassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetServer(v []string)`

SetServer sets Server field to given value.


### GetServerAccessMode

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetServerAccessMode() EnumpassThroughAuthenticationHandlerServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetServerAccessModeOk() (*EnumpassThroughAuthenticationHandlerServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetServerAccessMode(v EnumpassThroughAuthenticationHandlerServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.


### GetDnMap

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.


### GetMaxConnections

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.


### GetUseLocation

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetUseLocation() bool`

GetUseLocation returns the UseLocation field if non-nil, zero value otherwise.

### GetUseLocationOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetUseLocationOk() (*bool, bool)`

GetUseLocationOk returns a tuple with the UseLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocation

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetUseLocation(v bool)`

SetUseLocation sets UseLocation field to given value.

### HasUseLocation

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasUseLocation() bool`

HasUseLocation returns a boolean if a field has been set.

### GetMaximumAllowedLocalResponseTime

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMaximumAllowedLocalResponseTime() string`

GetMaximumAllowedLocalResponseTime returns the MaximumAllowedLocalResponseTime field if non-nil, zero value otherwise.

### GetMaximumAllowedLocalResponseTimeOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMaximumAllowedLocalResponseTimeOk() (*string, bool)`

GetMaximumAllowedLocalResponseTimeOk returns a tuple with the MaximumAllowedLocalResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAllowedLocalResponseTime

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetMaximumAllowedLocalResponseTime(v string)`

SetMaximumAllowedLocalResponseTime sets MaximumAllowedLocalResponseTime field to given value.

### HasMaximumAllowedLocalResponseTime

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasMaximumAllowedLocalResponseTime() bool`

HasMaximumAllowedLocalResponseTime returns a boolean if a field has been set.

### GetMaximumAllowedNonlocalResponseTime

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMaximumAllowedNonlocalResponseTime() string`

GetMaximumAllowedNonlocalResponseTime returns the MaximumAllowedNonlocalResponseTime field if non-nil, zero value otherwise.

### GetMaximumAllowedNonlocalResponseTimeOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMaximumAllowedNonlocalResponseTimeOk() (*string, bool)`

GetMaximumAllowedNonlocalResponseTimeOk returns a tuple with the MaximumAllowedNonlocalResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAllowedNonlocalResponseTime

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetMaximumAllowedNonlocalResponseTime(v string)`

SetMaximumAllowedNonlocalResponseTime sets MaximumAllowedNonlocalResponseTime field to given value.

### HasMaximumAllowedNonlocalResponseTime

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasMaximumAllowedNonlocalResponseTime() bool`

HasMaximumAllowedNonlocalResponseTime returns a boolean if a field has been set.

### GetUsePasswordPolicyControl

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetUsePasswordPolicyControl() bool`

GetUsePasswordPolicyControl returns the UsePasswordPolicyControl field if non-nil, zero value otherwise.

### GetUsePasswordPolicyControlOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetUsePasswordPolicyControlOk() (*bool, bool)`

GetUsePasswordPolicyControlOk returns a tuple with the UsePasswordPolicyControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasswordPolicyControl

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetUsePasswordPolicyControl(v bool)`

SetUsePasswordPolicyControl sets UsePasswordPolicyControl field to given value.

### HasUsePasswordPolicyControl

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasUsePasswordPolicyControl() bool`

HasUsePasswordPolicyControl returns a boolean if a field has been set.

### GetDescription

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdapPassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapPassThroughAuthenticationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdapPassThroughAuthenticationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


