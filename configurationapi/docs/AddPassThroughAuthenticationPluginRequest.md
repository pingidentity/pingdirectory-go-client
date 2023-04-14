# AddPassThroughAuthenticationPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumpassThroughAuthenticationPluginSchemaUrn**](EnumpassThroughAuthenticationPluginSchemaUrn.md) |  | 
**PluginType** | Pointer to [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | [optional] 
**Server** | **[]string** | Specifies the LDAP external server(s) to which authentication attempts should be forwarded. | 
**TryLocalBind** | Pointer to **bool** | Indicates whether the bind attempt should first be attempted against the local server. Depending on the value of the override-local-password property, the bind attempt may then be attempted against a remote server if the local bind fails. | [optional] 
**OverrideLocalPassword** | Pointer to **bool** | Indicates whether the bind attempt should be attempted against a remote server in the event that the local bind fails but the local password is present. | [optional] 
**UpdateLocalPassword** | Pointer to **bool** | Indicates whether the local password value should be updated to the value used in the bind request in the event that the local bind fails but the remote bind succeeds. | [optional] 
**AllowLaxPassThroughAuthenticationPasswords** | Pointer to **bool** | Indicates whether updates to the local password value should accept passwords that do not meet password policy constraints. | [optional] 
**ServerAccessMode** | Pointer to [**EnumpluginServerAccessModeProp**](EnumpluginServerAccessModeProp.md) |  | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to an alternate server. | [optional] 
**ConnectionCriteria** | Pointer to **string** | Specifies a set of connection criteria that must match the client associated with the bind request for the bind to be passed through to an alternate server. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria that must match the bind request for the bind to be passed through to an alternate server. | [optional] 
**DnMap** | Pointer to **[]string** | Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers. | [optional] 
**BindDNPattern** | Pointer to **string** | A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \&quot;cn&#x3D;{cn},ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot; indicates that the remote bind DN should be constructed from the text \&quot;cn&#x3D;\&quot; followed by the value of the local entry&#39;s cn attribute followed by the text \&quot;ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot;. If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \&quot;{seeAlso}\&quot; would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name. | [optional] 
**SearchBaseDN** | Pointer to **string** | The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN. | [optional] 
**SearchFilterPattern** | Pointer to **string** | A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \&quot;(mail&#x3D;{uid:ldapFilterEscape}@example.com)\&quot; would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \&quot;@example.com\&quot; in the mail attribute in the external server. Note that the \&quot;ldapFilterEscape\&quot; modifier should almost always be used with attributes specified in the pattern. | [optional] 
**InitialConnections** | Pointer to **int64** | Specifies the initial number of connections to establish to each external server against which authentication may be attempted. | [optional] 
**MaxConnections** | Pointer to **int64** | Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewAddPassThroughAuthenticationPluginRequest

`func NewAddPassThroughAuthenticationPluginRequest(pluginName string, schemas []EnumpassThroughAuthenticationPluginSchemaUrn, server []string, enabled bool, ) *AddPassThroughAuthenticationPluginRequest`

NewAddPassThroughAuthenticationPluginRequest instantiates a new AddPassThroughAuthenticationPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPassThroughAuthenticationPluginRequestWithDefaults

`func NewAddPassThroughAuthenticationPluginRequestWithDefaults() *AddPassThroughAuthenticationPluginRequest`

NewAddPassThroughAuthenticationPluginRequestWithDefaults instantiates a new AddPassThroughAuthenticationPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddPassThroughAuthenticationPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddPassThroughAuthenticationPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddPassThroughAuthenticationPluginRequest) GetSchemas() []EnumpassThroughAuthenticationPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetSchemasOk() (*[]EnumpassThroughAuthenticationPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPassThroughAuthenticationPluginRequest) SetSchemas(v []EnumpassThroughAuthenticationPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddPassThroughAuthenticationPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddPassThroughAuthenticationPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *AddPassThroughAuthenticationPluginRequest) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetServer

`func (o *AddPassThroughAuthenticationPluginRequest) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddPassThroughAuthenticationPluginRequest) SetServer(v []string)`

SetServer sets Server field to given value.


### GetTryLocalBind

`func (o *AddPassThroughAuthenticationPluginRequest) GetTryLocalBind() bool`

GetTryLocalBind returns the TryLocalBind field if non-nil, zero value otherwise.

### GetTryLocalBindOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetTryLocalBindOk() (*bool, bool)`

GetTryLocalBindOk returns a tuple with the TryLocalBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryLocalBind

`func (o *AddPassThroughAuthenticationPluginRequest) SetTryLocalBind(v bool)`

SetTryLocalBind sets TryLocalBind field to given value.

### HasTryLocalBind

`func (o *AddPassThroughAuthenticationPluginRequest) HasTryLocalBind() bool`

HasTryLocalBind returns a boolean if a field has been set.

### GetOverrideLocalPassword

`func (o *AddPassThroughAuthenticationPluginRequest) GetOverrideLocalPassword() bool`

GetOverrideLocalPassword returns the OverrideLocalPassword field if non-nil, zero value otherwise.

### GetOverrideLocalPasswordOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetOverrideLocalPasswordOk() (*bool, bool)`

GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLocalPassword

`func (o *AddPassThroughAuthenticationPluginRequest) SetOverrideLocalPassword(v bool)`

SetOverrideLocalPassword sets OverrideLocalPassword field to given value.

### HasOverrideLocalPassword

`func (o *AddPassThroughAuthenticationPluginRequest) HasOverrideLocalPassword() bool`

HasOverrideLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPassword

`func (o *AddPassThroughAuthenticationPluginRequest) GetUpdateLocalPassword() bool`

GetUpdateLocalPassword returns the UpdateLocalPassword field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordOk() (*bool, bool)`

GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPassword

`func (o *AddPassThroughAuthenticationPluginRequest) SetUpdateLocalPassword(v bool)`

SetUpdateLocalPassword sets UpdateLocalPassword field to given value.

### HasUpdateLocalPassword

`func (o *AddPassThroughAuthenticationPluginRequest) HasUpdateLocalPassword() bool`

HasUpdateLocalPassword returns a boolean if a field has been set.

### GetAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPassThroughAuthenticationPluginRequest) GetAllowLaxPassThroughAuthenticationPasswords() bool`

GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field if non-nil, zero value otherwise.

### GetAllowLaxPassThroughAuthenticationPasswordsOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool)`

GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPassThroughAuthenticationPluginRequest) SetAllowLaxPassThroughAuthenticationPasswords(v bool)`

SetAllowLaxPassThroughAuthenticationPasswords sets AllowLaxPassThroughAuthenticationPasswords field to given value.

### HasAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPassThroughAuthenticationPluginRequest) HasAllowLaxPassThroughAuthenticationPasswords() bool`

HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.

### GetServerAccessMode

`func (o *AddPassThroughAuthenticationPluginRequest) GetServerAccessMode() EnumpluginServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetServerAccessModeOk() (*EnumpluginServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *AddPassThroughAuthenticationPluginRequest) SetServerAccessMode(v EnumpluginServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.

### HasServerAccessMode

`func (o *AddPassThroughAuthenticationPluginRequest) HasServerAccessMode() bool`

HasServerAccessMode returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *AddPassThroughAuthenticationPluginRequest) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddPassThroughAuthenticationPluginRequest) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddPassThroughAuthenticationPluginRequest) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddPassThroughAuthenticationPluginRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddPassThroughAuthenticationPluginRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddPassThroughAuthenticationPluginRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddPassThroughAuthenticationPluginRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddPassThroughAuthenticationPluginRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddPassThroughAuthenticationPluginRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetDnMap

`func (o *AddPassThroughAuthenticationPluginRequest) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *AddPassThroughAuthenticationPluginRequest) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *AddPassThroughAuthenticationPluginRequest) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *AddPassThroughAuthenticationPluginRequest) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *AddPassThroughAuthenticationPluginRequest) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *AddPassThroughAuthenticationPluginRequest) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *AddPassThroughAuthenticationPluginRequest) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *AddPassThroughAuthenticationPluginRequest) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *AddPassThroughAuthenticationPluginRequest) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *AddPassThroughAuthenticationPluginRequest) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *AddPassThroughAuthenticationPluginRequest) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *AddPassThroughAuthenticationPluginRequest) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddPassThroughAuthenticationPluginRequest) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddPassThroughAuthenticationPluginRequest) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *AddPassThroughAuthenticationPluginRequest) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *AddPassThroughAuthenticationPluginRequest) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddPassThroughAuthenticationPluginRequest) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *AddPassThroughAuthenticationPluginRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDescription

`func (o *AddPassThroughAuthenticationPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPassThroughAuthenticationPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPassThroughAuthenticationPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPassThroughAuthenticationPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPassThroughAuthenticationPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddPassThroughAuthenticationPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddPassThroughAuthenticationPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddPassThroughAuthenticationPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddPassThroughAuthenticationPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


