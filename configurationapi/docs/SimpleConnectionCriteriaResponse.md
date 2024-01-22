# SimpleConnectionCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsimpleConnectionCriteriaSchemaUrn**](EnumsimpleConnectionCriteriaSchemaUrn.md) |  | 
**IncludedClientAddress** | Pointer to **[]string** | Specifies an address mask that may be used to specify a set of clients that should be included in this Simple Connection Criteria. | [optional] 
**ExcludedClientAddress** | Pointer to **[]string** | Specifies an address mask that may be used to specify a set of clients that should be excluded from this Simple Connection Criteria. | [optional] 
**IncludedConnectionHandler** | Pointer to **[]string** | Specifies a connection handler for clients that should be included in this Simple Connection Criteria. | [optional] 
**ExcludedConnectionHandler** | Pointer to **[]string** | Specifies a connection handler for clients that should be excluded from this Simple Connection Criteria. | [optional] 
**IncludedProtocol** | Pointer to **[]string** | Specifies the name of a communication protocol that should be used by clients included in this Simple Connection Criteria. | [optional] 
**ExcludedProtocol** | Pointer to **[]string** | Specifies the name of a communication protocol that should be used by clients excluded from this Simple Connection Criteria. | [optional] 
**CommunicationSecurityLevel** | Pointer to [**EnumconnectionCriteriaCommunicationSecurityLevelProp**](EnumconnectionCriteriaCommunicationSecurityLevelProp.md) |  | [optional] 
**UserAuthType** | Pointer to [**[]EnumconnectionCriteriaUserAuthTypeProp**](EnumconnectionCriteriaUserAuthTypeProp.md) |  | [optional] 
**AuthenticationSecurityLevel** | Pointer to [**EnumconnectionCriteriaAuthenticationSecurityLevelProp**](EnumconnectionCriteriaAuthenticationSecurityLevelProp.md) |  | [optional] 
**IncludedUserSASLMechanism** | Pointer to **[]string** | Specifies the name of a SASL mechanism that should be used by clients included in this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server using a SASL mechanism and will be ignored for unauthenticated client connections and for client connections that authenticated using some other method (e.g., those performing simple or internal authentication). | [optional] 
**ExcludedUserSASLMechanism** | Pointer to **[]string** | Specifies the name of a SASL mechanism that should be used by clients excluded from this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server using a SASL mechanism and will be ignored for unauthenticated client connections and for client connections that authenticated using some other method (e.g., those performing simple or internal authentication). | [optional] 
**IncludedUserBaseDN** | Pointer to **[]string** | Specifies a base DN below which authenticated user entries may exist for clients included in this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as). | [optional] 
**ExcludedUserBaseDN** | Pointer to **[]string** | Specifies a base DN below which authenticated user entries may exist for clients excluded from this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as). | [optional] 
**AllIncludedUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authenticated users must exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must be a member of all of those groups. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as). | [optional] 
**AnyIncludedUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authenticated users may exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must be a member of at least one of those groups. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as). | [optional] 
**NotAllIncludedUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authenticated users should not exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must not be a member of at least one of those groups (that is, the user may be a member of zero or more of those groups, but not of all of them). This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as). | [optional] 
**NoneIncludedUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authenticated users must not exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must not be a member any of those groups. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as). | [optional] 
**AllIncludedUserFilter** | Pointer to **[]string** | Specifies a search filter that must match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then all of those filters must match the authenticated user entry. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. | [optional] 
**AnyIncludedUserFilter** | Pointer to **[]string** | Specifies a search filter that may match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then at least one of those filters must match the authenticated user entry. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. | [optional] 
**NotAllIncludedUserFilter** | Pointer to **[]string** | Specifies a search filter that should not match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then at least one of those filters must not match the authenticated user entry (that is, the user entry may match zero or more of those filters, but not all of them). This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. | [optional] 
**NoneIncludedUserFilter** | Pointer to **[]string** | Specifies a search filter that must not match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then none of those filters may match the authenticated user entry. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. | [optional] 
**AllIncludedUserPrivilege** | Pointer to [**[]EnumconnectionCriteriaAllIncludedUserPrivilegeProp**](EnumconnectionCriteriaAllIncludedUserPrivilegeProp.md) |  | [optional] 
**AnyIncludedUserPrivilege** | Pointer to [**[]EnumconnectionCriteriaAnyIncludedUserPrivilegeProp**](EnumconnectionCriteriaAnyIncludedUserPrivilegeProp.md) |  | [optional] 
**NotAllIncludedUserPrivilege** | Pointer to [**[]EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp**](EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp.md) |  | [optional] 
**NoneIncludedUserPrivilege** | Pointer to [**[]EnumconnectionCriteriaNoneIncludedUserPrivilegeProp**](EnumconnectionCriteriaNoneIncludedUserPrivilegeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Connection Criteria | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Connection Criteria | 

## Methods

### NewSimpleConnectionCriteriaResponse

`func NewSimpleConnectionCriteriaResponse(schemas []EnumsimpleConnectionCriteriaSchemaUrn, id string, ) *SimpleConnectionCriteriaResponse`

NewSimpleConnectionCriteriaResponse instantiates a new SimpleConnectionCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleConnectionCriteriaResponseWithDefaults

`func NewSimpleConnectionCriteriaResponseWithDefaults() *SimpleConnectionCriteriaResponse`

NewSimpleConnectionCriteriaResponseWithDefaults instantiates a new SimpleConnectionCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SimpleConnectionCriteriaResponse) GetSchemas() []EnumsimpleConnectionCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleConnectionCriteriaResponse) GetSchemasOk() (*[]EnumsimpleConnectionCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleConnectionCriteriaResponse) SetSchemas(v []EnumsimpleConnectionCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIncludedClientAddress

`func (o *SimpleConnectionCriteriaResponse) GetIncludedClientAddress() []string`

GetIncludedClientAddress returns the IncludedClientAddress field if non-nil, zero value otherwise.

### GetIncludedClientAddressOk

`func (o *SimpleConnectionCriteriaResponse) GetIncludedClientAddressOk() (*[]string, bool)`

GetIncludedClientAddressOk returns a tuple with the IncludedClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientAddress

`func (o *SimpleConnectionCriteriaResponse) SetIncludedClientAddress(v []string)`

SetIncludedClientAddress sets IncludedClientAddress field to given value.

### HasIncludedClientAddress

`func (o *SimpleConnectionCriteriaResponse) HasIncludedClientAddress() bool`

HasIncludedClientAddress returns a boolean if a field has been set.

### GetExcludedClientAddress

`func (o *SimpleConnectionCriteriaResponse) GetExcludedClientAddress() []string`

GetExcludedClientAddress returns the ExcludedClientAddress field if non-nil, zero value otherwise.

### GetExcludedClientAddressOk

`func (o *SimpleConnectionCriteriaResponse) GetExcludedClientAddressOk() (*[]string, bool)`

GetExcludedClientAddressOk returns a tuple with the ExcludedClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientAddress

`func (o *SimpleConnectionCriteriaResponse) SetExcludedClientAddress(v []string)`

SetExcludedClientAddress sets ExcludedClientAddress field to given value.

### HasExcludedClientAddress

`func (o *SimpleConnectionCriteriaResponse) HasExcludedClientAddress() bool`

HasExcludedClientAddress returns a boolean if a field has been set.

### GetIncludedConnectionHandler

`func (o *SimpleConnectionCriteriaResponse) GetIncludedConnectionHandler() []string`

GetIncludedConnectionHandler returns the IncludedConnectionHandler field if non-nil, zero value otherwise.

### GetIncludedConnectionHandlerOk

`func (o *SimpleConnectionCriteriaResponse) GetIncludedConnectionHandlerOk() (*[]string, bool)`

GetIncludedConnectionHandlerOk returns a tuple with the IncludedConnectionHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedConnectionHandler

`func (o *SimpleConnectionCriteriaResponse) SetIncludedConnectionHandler(v []string)`

SetIncludedConnectionHandler sets IncludedConnectionHandler field to given value.

### HasIncludedConnectionHandler

`func (o *SimpleConnectionCriteriaResponse) HasIncludedConnectionHandler() bool`

HasIncludedConnectionHandler returns a boolean if a field has been set.

### GetExcludedConnectionHandler

`func (o *SimpleConnectionCriteriaResponse) GetExcludedConnectionHandler() []string`

GetExcludedConnectionHandler returns the ExcludedConnectionHandler field if non-nil, zero value otherwise.

### GetExcludedConnectionHandlerOk

`func (o *SimpleConnectionCriteriaResponse) GetExcludedConnectionHandlerOk() (*[]string, bool)`

GetExcludedConnectionHandlerOk returns a tuple with the ExcludedConnectionHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedConnectionHandler

`func (o *SimpleConnectionCriteriaResponse) SetExcludedConnectionHandler(v []string)`

SetExcludedConnectionHandler sets ExcludedConnectionHandler field to given value.

### HasExcludedConnectionHandler

`func (o *SimpleConnectionCriteriaResponse) HasExcludedConnectionHandler() bool`

HasExcludedConnectionHandler returns a boolean if a field has been set.

### GetIncludedProtocol

`func (o *SimpleConnectionCriteriaResponse) GetIncludedProtocol() []string`

GetIncludedProtocol returns the IncludedProtocol field if non-nil, zero value otherwise.

### GetIncludedProtocolOk

`func (o *SimpleConnectionCriteriaResponse) GetIncludedProtocolOk() (*[]string, bool)`

GetIncludedProtocolOk returns a tuple with the IncludedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedProtocol

`func (o *SimpleConnectionCriteriaResponse) SetIncludedProtocol(v []string)`

SetIncludedProtocol sets IncludedProtocol field to given value.

### HasIncludedProtocol

`func (o *SimpleConnectionCriteriaResponse) HasIncludedProtocol() bool`

HasIncludedProtocol returns a boolean if a field has been set.

### GetExcludedProtocol

`func (o *SimpleConnectionCriteriaResponse) GetExcludedProtocol() []string`

GetExcludedProtocol returns the ExcludedProtocol field if non-nil, zero value otherwise.

### GetExcludedProtocolOk

`func (o *SimpleConnectionCriteriaResponse) GetExcludedProtocolOk() (*[]string, bool)`

GetExcludedProtocolOk returns a tuple with the ExcludedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProtocol

`func (o *SimpleConnectionCriteriaResponse) SetExcludedProtocol(v []string)`

SetExcludedProtocol sets ExcludedProtocol field to given value.

### HasExcludedProtocol

`func (o *SimpleConnectionCriteriaResponse) HasExcludedProtocol() bool`

HasExcludedProtocol returns a boolean if a field has been set.

### GetCommunicationSecurityLevel

`func (o *SimpleConnectionCriteriaResponse) GetCommunicationSecurityLevel() EnumconnectionCriteriaCommunicationSecurityLevelProp`

GetCommunicationSecurityLevel returns the CommunicationSecurityLevel field if non-nil, zero value otherwise.

### GetCommunicationSecurityLevelOk

`func (o *SimpleConnectionCriteriaResponse) GetCommunicationSecurityLevelOk() (*EnumconnectionCriteriaCommunicationSecurityLevelProp, bool)`

GetCommunicationSecurityLevelOk returns a tuple with the CommunicationSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationSecurityLevel

`func (o *SimpleConnectionCriteriaResponse) SetCommunicationSecurityLevel(v EnumconnectionCriteriaCommunicationSecurityLevelProp)`

SetCommunicationSecurityLevel sets CommunicationSecurityLevel field to given value.

### HasCommunicationSecurityLevel

`func (o *SimpleConnectionCriteriaResponse) HasCommunicationSecurityLevel() bool`

HasCommunicationSecurityLevel returns a boolean if a field has been set.

### GetUserAuthType

`func (o *SimpleConnectionCriteriaResponse) GetUserAuthType() []EnumconnectionCriteriaUserAuthTypeProp`

GetUserAuthType returns the UserAuthType field if non-nil, zero value otherwise.

### GetUserAuthTypeOk

`func (o *SimpleConnectionCriteriaResponse) GetUserAuthTypeOk() (*[]EnumconnectionCriteriaUserAuthTypeProp, bool)`

GetUserAuthTypeOk returns a tuple with the UserAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthType

`func (o *SimpleConnectionCriteriaResponse) SetUserAuthType(v []EnumconnectionCriteriaUserAuthTypeProp)`

SetUserAuthType sets UserAuthType field to given value.

### HasUserAuthType

`func (o *SimpleConnectionCriteriaResponse) HasUserAuthType() bool`

HasUserAuthType returns a boolean if a field has been set.

### GetAuthenticationSecurityLevel

`func (o *SimpleConnectionCriteriaResponse) GetAuthenticationSecurityLevel() EnumconnectionCriteriaAuthenticationSecurityLevelProp`

GetAuthenticationSecurityLevel returns the AuthenticationSecurityLevel field if non-nil, zero value otherwise.

### GetAuthenticationSecurityLevelOk

`func (o *SimpleConnectionCriteriaResponse) GetAuthenticationSecurityLevelOk() (*EnumconnectionCriteriaAuthenticationSecurityLevelProp, bool)`

GetAuthenticationSecurityLevelOk returns a tuple with the AuthenticationSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSecurityLevel

`func (o *SimpleConnectionCriteriaResponse) SetAuthenticationSecurityLevel(v EnumconnectionCriteriaAuthenticationSecurityLevelProp)`

SetAuthenticationSecurityLevel sets AuthenticationSecurityLevel field to given value.

### HasAuthenticationSecurityLevel

`func (o *SimpleConnectionCriteriaResponse) HasAuthenticationSecurityLevel() bool`

HasAuthenticationSecurityLevel returns a boolean if a field has been set.

### GetIncludedUserSASLMechanism

`func (o *SimpleConnectionCriteriaResponse) GetIncludedUserSASLMechanism() []string`

GetIncludedUserSASLMechanism returns the IncludedUserSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedUserSASLMechanismOk

`func (o *SimpleConnectionCriteriaResponse) GetIncludedUserSASLMechanismOk() (*[]string, bool)`

GetIncludedUserSASLMechanismOk returns a tuple with the IncludedUserSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserSASLMechanism

`func (o *SimpleConnectionCriteriaResponse) SetIncludedUserSASLMechanism(v []string)`

SetIncludedUserSASLMechanism sets IncludedUserSASLMechanism field to given value.

### HasIncludedUserSASLMechanism

`func (o *SimpleConnectionCriteriaResponse) HasIncludedUserSASLMechanism() bool`

HasIncludedUserSASLMechanism returns a boolean if a field has been set.

### GetExcludedUserSASLMechanism

`func (o *SimpleConnectionCriteriaResponse) GetExcludedUserSASLMechanism() []string`

GetExcludedUserSASLMechanism returns the ExcludedUserSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedUserSASLMechanismOk

`func (o *SimpleConnectionCriteriaResponse) GetExcludedUserSASLMechanismOk() (*[]string, bool)`

GetExcludedUserSASLMechanismOk returns a tuple with the ExcludedUserSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserSASLMechanism

`func (o *SimpleConnectionCriteriaResponse) SetExcludedUserSASLMechanism(v []string)`

SetExcludedUserSASLMechanism sets ExcludedUserSASLMechanism field to given value.

### HasExcludedUserSASLMechanism

`func (o *SimpleConnectionCriteriaResponse) HasExcludedUserSASLMechanism() bool`

HasExcludedUserSASLMechanism returns a boolean if a field has been set.

### GetIncludedUserBaseDN

`func (o *SimpleConnectionCriteriaResponse) GetIncludedUserBaseDN() []string`

GetIncludedUserBaseDN returns the IncludedUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedUserBaseDNOk

`func (o *SimpleConnectionCriteriaResponse) GetIncludedUserBaseDNOk() (*[]string, bool)`

GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserBaseDN

`func (o *SimpleConnectionCriteriaResponse) SetIncludedUserBaseDN(v []string)`

SetIncludedUserBaseDN sets IncludedUserBaseDN field to given value.

### HasIncludedUserBaseDN

`func (o *SimpleConnectionCriteriaResponse) HasIncludedUserBaseDN() bool`

HasIncludedUserBaseDN returns a boolean if a field has been set.

### GetExcludedUserBaseDN

`func (o *SimpleConnectionCriteriaResponse) GetExcludedUserBaseDN() []string`

GetExcludedUserBaseDN returns the ExcludedUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedUserBaseDNOk

`func (o *SimpleConnectionCriteriaResponse) GetExcludedUserBaseDNOk() (*[]string, bool)`

GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserBaseDN

`func (o *SimpleConnectionCriteriaResponse) SetExcludedUserBaseDN(v []string)`

SetExcludedUserBaseDN sets ExcludedUserBaseDN field to given value.

### HasExcludedUserBaseDN

`func (o *SimpleConnectionCriteriaResponse) HasExcludedUserBaseDN() bool`

HasExcludedUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) GetAllIncludedUserGroupDN() []string`

GetAllIncludedUserGroupDN returns the AllIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedUserGroupDNOk

`func (o *SimpleConnectionCriteriaResponse) GetAllIncludedUserGroupDNOk() (*[]string, bool)`

GetAllIncludedUserGroupDNOk returns a tuple with the AllIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) SetAllIncludedUserGroupDN(v []string)`

SetAllIncludedUserGroupDN sets AllIncludedUserGroupDN field to given value.

### HasAllIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) HasAllIncludedUserGroupDN() bool`

HasAllIncludedUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) GetAnyIncludedUserGroupDN() []string`

GetAnyIncludedUserGroupDN returns the AnyIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedUserGroupDNOk

`func (o *SimpleConnectionCriteriaResponse) GetAnyIncludedUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedUserGroupDNOk returns a tuple with the AnyIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) SetAnyIncludedUserGroupDN(v []string)`

SetAnyIncludedUserGroupDN sets AnyIncludedUserGroupDN field to given value.

### HasAnyIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) HasAnyIncludedUserGroupDN() bool`

HasAnyIncludedUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) GetNotAllIncludedUserGroupDN() []string`

GetNotAllIncludedUserGroupDN returns the NotAllIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedUserGroupDNOk

`func (o *SimpleConnectionCriteriaResponse) GetNotAllIncludedUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedUserGroupDNOk returns a tuple with the NotAllIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) SetNotAllIncludedUserGroupDN(v []string)`

SetNotAllIncludedUserGroupDN sets NotAllIncludedUserGroupDN field to given value.

### HasNotAllIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) HasNotAllIncludedUserGroupDN() bool`

HasNotAllIncludedUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) GetNoneIncludedUserGroupDN() []string`

GetNoneIncludedUserGroupDN returns the NoneIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedUserGroupDNOk

`func (o *SimpleConnectionCriteriaResponse) GetNoneIncludedUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedUserGroupDNOk returns a tuple with the NoneIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) SetNoneIncludedUserGroupDN(v []string)`

SetNoneIncludedUserGroupDN sets NoneIncludedUserGroupDN field to given value.

### HasNoneIncludedUserGroupDN

`func (o *SimpleConnectionCriteriaResponse) HasNoneIncludedUserGroupDN() bool`

HasNoneIncludedUserGroupDN returns a boolean if a field has been set.

### GetAllIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) GetAllIncludedUserFilter() []string`

GetAllIncludedUserFilter returns the AllIncludedUserFilter field if non-nil, zero value otherwise.

### GetAllIncludedUserFilterOk

`func (o *SimpleConnectionCriteriaResponse) GetAllIncludedUserFilterOk() (*[]string, bool)`

GetAllIncludedUserFilterOk returns a tuple with the AllIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) SetAllIncludedUserFilter(v []string)`

SetAllIncludedUserFilter sets AllIncludedUserFilter field to given value.

### HasAllIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) HasAllIncludedUserFilter() bool`

HasAllIncludedUserFilter returns a boolean if a field has been set.

### GetAnyIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) GetAnyIncludedUserFilter() []string`

GetAnyIncludedUserFilter returns the AnyIncludedUserFilter field if non-nil, zero value otherwise.

### GetAnyIncludedUserFilterOk

`func (o *SimpleConnectionCriteriaResponse) GetAnyIncludedUserFilterOk() (*[]string, bool)`

GetAnyIncludedUserFilterOk returns a tuple with the AnyIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) SetAnyIncludedUserFilter(v []string)`

SetAnyIncludedUserFilter sets AnyIncludedUserFilter field to given value.

### HasAnyIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) HasAnyIncludedUserFilter() bool`

HasAnyIncludedUserFilter returns a boolean if a field has been set.

### GetNotAllIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) GetNotAllIncludedUserFilter() []string`

GetNotAllIncludedUserFilter returns the NotAllIncludedUserFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedUserFilterOk

`func (o *SimpleConnectionCriteriaResponse) GetNotAllIncludedUserFilterOk() (*[]string, bool)`

GetNotAllIncludedUserFilterOk returns a tuple with the NotAllIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) SetNotAllIncludedUserFilter(v []string)`

SetNotAllIncludedUserFilter sets NotAllIncludedUserFilter field to given value.

### HasNotAllIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) HasNotAllIncludedUserFilter() bool`

HasNotAllIncludedUserFilter returns a boolean if a field has been set.

### GetNoneIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) GetNoneIncludedUserFilter() []string`

GetNoneIncludedUserFilter returns the NoneIncludedUserFilter field if non-nil, zero value otherwise.

### GetNoneIncludedUserFilterOk

`func (o *SimpleConnectionCriteriaResponse) GetNoneIncludedUserFilterOk() (*[]string, bool)`

GetNoneIncludedUserFilterOk returns a tuple with the NoneIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) SetNoneIncludedUserFilter(v []string)`

SetNoneIncludedUserFilter sets NoneIncludedUserFilter field to given value.

### HasNoneIncludedUserFilter

`func (o *SimpleConnectionCriteriaResponse) HasNoneIncludedUserFilter() bool`

HasNoneIncludedUserFilter returns a boolean if a field has been set.

### GetAllIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) GetAllIncludedUserPrivilege() []EnumconnectionCriteriaAllIncludedUserPrivilegeProp`

GetAllIncludedUserPrivilege returns the AllIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetAllIncludedUserPrivilegeOk

`func (o *SimpleConnectionCriteriaResponse) GetAllIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaAllIncludedUserPrivilegeProp, bool)`

GetAllIncludedUserPrivilegeOk returns a tuple with the AllIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) SetAllIncludedUserPrivilege(v []EnumconnectionCriteriaAllIncludedUserPrivilegeProp)`

SetAllIncludedUserPrivilege sets AllIncludedUserPrivilege field to given value.

### HasAllIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) HasAllIncludedUserPrivilege() bool`

HasAllIncludedUserPrivilege returns a boolean if a field has been set.

### GetAnyIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) GetAnyIncludedUserPrivilege() []EnumconnectionCriteriaAnyIncludedUserPrivilegeProp`

GetAnyIncludedUserPrivilege returns the AnyIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetAnyIncludedUserPrivilegeOk

`func (o *SimpleConnectionCriteriaResponse) GetAnyIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaAnyIncludedUserPrivilegeProp, bool)`

GetAnyIncludedUserPrivilegeOk returns a tuple with the AnyIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) SetAnyIncludedUserPrivilege(v []EnumconnectionCriteriaAnyIncludedUserPrivilegeProp)`

SetAnyIncludedUserPrivilege sets AnyIncludedUserPrivilege field to given value.

### HasAnyIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) HasAnyIncludedUserPrivilege() bool`

HasAnyIncludedUserPrivilege returns a boolean if a field has been set.

### GetNotAllIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) GetNotAllIncludedUserPrivilege() []EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp`

GetNotAllIncludedUserPrivilege returns the NotAllIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetNotAllIncludedUserPrivilegeOk

`func (o *SimpleConnectionCriteriaResponse) GetNotAllIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp, bool)`

GetNotAllIncludedUserPrivilegeOk returns a tuple with the NotAllIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) SetNotAllIncludedUserPrivilege(v []EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp)`

SetNotAllIncludedUserPrivilege sets NotAllIncludedUserPrivilege field to given value.

### HasNotAllIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) HasNotAllIncludedUserPrivilege() bool`

HasNotAllIncludedUserPrivilege returns a boolean if a field has been set.

### GetNoneIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) GetNoneIncludedUserPrivilege() []EnumconnectionCriteriaNoneIncludedUserPrivilegeProp`

GetNoneIncludedUserPrivilege returns the NoneIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetNoneIncludedUserPrivilegeOk

`func (o *SimpleConnectionCriteriaResponse) GetNoneIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaNoneIncludedUserPrivilegeProp, bool)`

GetNoneIncludedUserPrivilegeOk returns a tuple with the NoneIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) SetNoneIncludedUserPrivilege(v []EnumconnectionCriteriaNoneIncludedUserPrivilegeProp)`

SetNoneIncludedUserPrivilege sets NoneIncludedUserPrivilege field to given value.

### HasNoneIncludedUserPrivilege

`func (o *SimpleConnectionCriteriaResponse) HasNoneIncludedUserPrivilege() bool`

HasNoneIncludedUserPrivilege returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleConnectionCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleConnectionCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleConnectionCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleConnectionCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *SimpleConnectionCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SimpleConnectionCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SimpleConnectionCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SimpleConnectionCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleConnectionCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SimpleConnectionCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleConnectionCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleConnectionCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *SimpleConnectionCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleConnectionCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleConnectionCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


