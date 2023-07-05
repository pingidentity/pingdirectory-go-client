# AddLdapPassThroughAuthenticationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Pass Through Authentication Handler | 
**Schemas** | [**[]EnumldapPassThroughAuthenticationHandlerSchemaUrn**](EnumldapPassThroughAuthenticationHandlerSchemaUrn.md) |  | 
**Server** | **[]string** | Specifies the LDAP external server(s) to which authentication attempts should be forwarded. | 
**ServerAccessMode** | Pointer to [**EnumpassThroughAuthenticationHandlerServerAccessModeProp**](EnumpassThroughAuthenticationHandlerServerAccessModeProp.md) |  | [optional] 
**DnMap** | Pointer to **[]string** | Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers. | [optional] 
**BindDNPattern** | Pointer to **string** | A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \&quot;cn&#x3D;{cn},ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot; indicates that the remote bind DN should be constructed from the text \&quot;cn&#x3D;\&quot; followed by the value of the local entry&#39;s cn attribute followed by the text \&quot;ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot;. If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \&quot;{seeAlso}\&quot; would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name. | [optional] 
**SearchBaseDN** | Pointer to **string** | The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN. | [optional] 
**SearchFilterPattern** | Pointer to **string** | A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \&quot;(mail&#x3D;{uid:ldapFilterEscape}@example.com)\&quot; would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \&quot;@example.com\&quot; in the mail attribute in the external server. Note that the \&quot;ldapFilterEscape\&quot; modifier should almost always be used with attributes specified in the pattern. | [optional] 
**InitialConnections** | Pointer to **int64** | Specifies the initial number of connections to establish to each external server against which authentication may be attempted. | [optional] 
**MaxConnections** | Pointer to **int64** | Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property. | [optional] 
**UseLocation** | Pointer to **bool** | Indicates whether to take server locations into account when prioritizing the servers to use for pass-through authentication attempts. | [optional] 
**MaximumAllowedLocalResponseTime** | Pointer to **string** | The maximum length of time to wait for a response from an external server in the same location as this Directory Server before considering it unavailable. | [optional] 
**MaximumAllowedNonlocalResponseTime** | Pointer to **string** | The maximum length of time to wait for a response from an external server in a different location from this Directory Server before considering it unavailable. | [optional] 
**UsePasswordPolicyControl** | Pointer to **bool** | Indicates whether to include the password policy request control (as defined in draft-behera-ldap-password-policy-10) in bind requests sent to the external server. | [optional] 
**Description** | Pointer to **string** | A description for this Pass Through Authentication Handler | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 

## Methods

### NewAddLdapPassThroughAuthenticationHandlerRequest

`func NewAddLdapPassThroughAuthenticationHandlerRequest(handlerName string, schemas []EnumldapPassThroughAuthenticationHandlerSchemaUrn, server []string, ) *AddLdapPassThroughAuthenticationHandlerRequest`

NewAddLdapPassThroughAuthenticationHandlerRequest instantiates a new AddLdapPassThroughAuthenticationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLdapPassThroughAuthenticationHandlerRequestWithDefaults

`func NewAddLdapPassThroughAuthenticationHandlerRequestWithDefaults() *AddLdapPassThroughAuthenticationHandlerRequest`

NewAddLdapPassThroughAuthenticationHandlerRequestWithDefaults instantiates a new AddLdapPassThroughAuthenticationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSchemas() []EnumldapPassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSchemasOk() (*[]EnumldapPassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetSchemas(v []EnumldapPassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetServer(v []string)`

SetServer sets Server field to given value.


### GetServerAccessMode

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServerAccessMode() EnumpassThroughAuthenticationHandlerServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServerAccessModeOk() (*EnumpassThroughAuthenticationHandlerServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetServerAccessMode(v EnumpassThroughAuthenticationHandlerServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.

### HasServerAccessMode

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasServerAccessMode() bool`

HasServerAccessMode returns a boolean if a field has been set.

### GetDnMap

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetUseLocation

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUseLocation() bool`

GetUseLocation returns the UseLocation field if non-nil, zero value otherwise.

### GetUseLocationOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUseLocationOk() (*bool, bool)`

GetUseLocationOk returns a tuple with the UseLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocation

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetUseLocation(v bool)`

SetUseLocation sets UseLocation field to given value.

### HasUseLocation

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasUseLocation() bool`

HasUseLocation returns a boolean if a field has been set.

### GetMaximumAllowedLocalResponseTime

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedLocalResponseTime() string`

GetMaximumAllowedLocalResponseTime returns the MaximumAllowedLocalResponseTime field if non-nil, zero value otherwise.

### GetMaximumAllowedLocalResponseTimeOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedLocalResponseTimeOk() (*string, bool)`

GetMaximumAllowedLocalResponseTimeOk returns a tuple with the MaximumAllowedLocalResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAllowedLocalResponseTime

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetMaximumAllowedLocalResponseTime(v string)`

SetMaximumAllowedLocalResponseTime sets MaximumAllowedLocalResponseTime field to given value.

### HasMaximumAllowedLocalResponseTime

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasMaximumAllowedLocalResponseTime() bool`

HasMaximumAllowedLocalResponseTime returns a boolean if a field has been set.

### GetMaximumAllowedNonlocalResponseTime

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedNonlocalResponseTime() string`

GetMaximumAllowedNonlocalResponseTime returns the MaximumAllowedNonlocalResponseTime field if non-nil, zero value otherwise.

### GetMaximumAllowedNonlocalResponseTimeOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedNonlocalResponseTimeOk() (*string, bool)`

GetMaximumAllowedNonlocalResponseTimeOk returns a tuple with the MaximumAllowedNonlocalResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAllowedNonlocalResponseTime

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetMaximumAllowedNonlocalResponseTime(v string)`

SetMaximumAllowedNonlocalResponseTime sets MaximumAllowedNonlocalResponseTime field to given value.

### HasMaximumAllowedNonlocalResponseTime

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasMaximumAllowedNonlocalResponseTime() bool`

HasMaximumAllowedNonlocalResponseTime returns a boolean if a field has been set.

### GetUsePasswordPolicyControl

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUsePasswordPolicyControl() bool`

GetUsePasswordPolicyControl returns the UsePasswordPolicyControl field if non-nil, zero value otherwise.

### GetUsePasswordPolicyControlOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUsePasswordPolicyControlOk() (*bool, bool)`

GetUsePasswordPolicyControlOk returns a tuple with the UsePasswordPolicyControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasswordPolicyControl

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetUsePasswordPolicyControl(v bool)`

SetUsePasswordPolicyControl sets UsePasswordPolicyControl field to given value.

### HasUsePasswordPolicyControl

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasUsePasswordPolicyControl() bool`

HasUsePasswordPolicyControl returns a boolean if a field has been set.

### GetDescription

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


