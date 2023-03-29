# AddConnectionCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Connection Criteria | 
**Schemas** | [**[]EnumthirdPartyConnectionCriteriaSchemaUrn**](EnumthirdPartyConnectionCriteriaSchemaUrn.md) |  | 
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
**AllIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that must match the associated client connection in order to match the aggregate connection criteria. If one or more all-included connection criteria objects are provided, then a client connection must match all of them in order to match the aggregate connection criteria. | [optional] 
**AnyIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that may match the associated client connection in order to match the aggregate connection criteria. If one or more any-included connection criteria objects are provided, then a client connection must match at least one of them in order to match the aggregate connection criteria. | [optional] 
**NotAllIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that should not match the associated client connection in order to match the aggregate connection criteria. If one or more not-all-included connection criteria objects are provided, then a client connection must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate connection criteria. | [optional] 
**NoneIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that must not match the associated client connection in order to match the aggregate connection criteria. If one or more none-included connection criteria objects are provided, then a client connection must not match any of them in order to match the aggregate connection criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Connection Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Connection Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddConnectionCriteria200Response

`func NewAddConnectionCriteria200Response(id string, schemas []EnumthirdPartyConnectionCriteriaSchemaUrn, extensionClass string, ) *AddConnectionCriteria200Response`

NewAddConnectionCriteria200Response instantiates a new AddConnectionCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConnectionCriteria200ResponseWithDefaults

`func NewAddConnectionCriteria200ResponseWithDefaults() *AddConnectionCriteria200Response`

NewAddConnectionCriteria200ResponseWithDefaults instantiates a new AddConnectionCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddConnectionCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddConnectionCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddConnectionCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddConnectionCriteria200Response) GetSchemas() []EnumthirdPartyConnectionCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConnectionCriteria200Response) GetSchemasOk() (*[]EnumthirdPartyConnectionCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConnectionCriteria200Response) SetSchemas(v []EnumthirdPartyConnectionCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIncludedClientAddress

`func (o *AddConnectionCriteria200Response) GetIncludedClientAddress() []string`

GetIncludedClientAddress returns the IncludedClientAddress field if non-nil, zero value otherwise.

### GetIncludedClientAddressOk

`func (o *AddConnectionCriteria200Response) GetIncludedClientAddressOk() (*[]string, bool)`

GetIncludedClientAddressOk returns a tuple with the IncludedClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientAddress

`func (o *AddConnectionCriteria200Response) SetIncludedClientAddress(v []string)`

SetIncludedClientAddress sets IncludedClientAddress field to given value.

### HasIncludedClientAddress

`func (o *AddConnectionCriteria200Response) HasIncludedClientAddress() bool`

HasIncludedClientAddress returns a boolean if a field has been set.

### GetExcludedClientAddress

`func (o *AddConnectionCriteria200Response) GetExcludedClientAddress() []string`

GetExcludedClientAddress returns the ExcludedClientAddress field if non-nil, zero value otherwise.

### GetExcludedClientAddressOk

`func (o *AddConnectionCriteria200Response) GetExcludedClientAddressOk() (*[]string, bool)`

GetExcludedClientAddressOk returns a tuple with the ExcludedClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientAddress

`func (o *AddConnectionCriteria200Response) SetExcludedClientAddress(v []string)`

SetExcludedClientAddress sets ExcludedClientAddress field to given value.

### HasExcludedClientAddress

`func (o *AddConnectionCriteria200Response) HasExcludedClientAddress() bool`

HasExcludedClientAddress returns a boolean if a field has been set.

### GetIncludedConnectionHandler

`func (o *AddConnectionCriteria200Response) GetIncludedConnectionHandler() []string`

GetIncludedConnectionHandler returns the IncludedConnectionHandler field if non-nil, zero value otherwise.

### GetIncludedConnectionHandlerOk

`func (o *AddConnectionCriteria200Response) GetIncludedConnectionHandlerOk() (*[]string, bool)`

GetIncludedConnectionHandlerOk returns a tuple with the IncludedConnectionHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedConnectionHandler

`func (o *AddConnectionCriteria200Response) SetIncludedConnectionHandler(v []string)`

SetIncludedConnectionHandler sets IncludedConnectionHandler field to given value.

### HasIncludedConnectionHandler

`func (o *AddConnectionCriteria200Response) HasIncludedConnectionHandler() bool`

HasIncludedConnectionHandler returns a boolean if a field has been set.

### GetExcludedConnectionHandler

`func (o *AddConnectionCriteria200Response) GetExcludedConnectionHandler() []string`

GetExcludedConnectionHandler returns the ExcludedConnectionHandler field if non-nil, zero value otherwise.

### GetExcludedConnectionHandlerOk

`func (o *AddConnectionCriteria200Response) GetExcludedConnectionHandlerOk() (*[]string, bool)`

GetExcludedConnectionHandlerOk returns a tuple with the ExcludedConnectionHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedConnectionHandler

`func (o *AddConnectionCriteria200Response) SetExcludedConnectionHandler(v []string)`

SetExcludedConnectionHandler sets ExcludedConnectionHandler field to given value.

### HasExcludedConnectionHandler

`func (o *AddConnectionCriteria200Response) HasExcludedConnectionHandler() bool`

HasExcludedConnectionHandler returns a boolean if a field has been set.

### GetIncludedProtocol

`func (o *AddConnectionCriteria200Response) GetIncludedProtocol() []string`

GetIncludedProtocol returns the IncludedProtocol field if non-nil, zero value otherwise.

### GetIncludedProtocolOk

`func (o *AddConnectionCriteria200Response) GetIncludedProtocolOk() (*[]string, bool)`

GetIncludedProtocolOk returns a tuple with the IncludedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedProtocol

`func (o *AddConnectionCriteria200Response) SetIncludedProtocol(v []string)`

SetIncludedProtocol sets IncludedProtocol field to given value.

### HasIncludedProtocol

`func (o *AddConnectionCriteria200Response) HasIncludedProtocol() bool`

HasIncludedProtocol returns a boolean if a field has been set.

### GetExcludedProtocol

`func (o *AddConnectionCriteria200Response) GetExcludedProtocol() []string`

GetExcludedProtocol returns the ExcludedProtocol field if non-nil, zero value otherwise.

### GetExcludedProtocolOk

`func (o *AddConnectionCriteria200Response) GetExcludedProtocolOk() (*[]string, bool)`

GetExcludedProtocolOk returns a tuple with the ExcludedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProtocol

`func (o *AddConnectionCriteria200Response) SetExcludedProtocol(v []string)`

SetExcludedProtocol sets ExcludedProtocol field to given value.

### HasExcludedProtocol

`func (o *AddConnectionCriteria200Response) HasExcludedProtocol() bool`

HasExcludedProtocol returns a boolean if a field has been set.

### GetCommunicationSecurityLevel

`func (o *AddConnectionCriteria200Response) GetCommunicationSecurityLevel() EnumconnectionCriteriaCommunicationSecurityLevelProp`

GetCommunicationSecurityLevel returns the CommunicationSecurityLevel field if non-nil, zero value otherwise.

### GetCommunicationSecurityLevelOk

`func (o *AddConnectionCriteria200Response) GetCommunicationSecurityLevelOk() (*EnumconnectionCriteriaCommunicationSecurityLevelProp, bool)`

GetCommunicationSecurityLevelOk returns a tuple with the CommunicationSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationSecurityLevel

`func (o *AddConnectionCriteria200Response) SetCommunicationSecurityLevel(v EnumconnectionCriteriaCommunicationSecurityLevelProp)`

SetCommunicationSecurityLevel sets CommunicationSecurityLevel field to given value.

### HasCommunicationSecurityLevel

`func (o *AddConnectionCriteria200Response) HasCommunicationSecurityLevel() bool`

HasCommunicationSecurityLevel returns a boolean if a field has been set.

### GetUserAuthType

`func (o *AddConnectionCriteria200Response) GetUserAuthType() []EnumconnectionCriteriaUserAuthTypeProp`

GetUserAuthType returns the UserAuthType field if non-nil, zero value otherwise.

### GetUserAuthTypeOk

`func (o *AddConnectionCriteria200Response) GetUserAuthTypeOk() (*[]EnumconnectionCriteriaUserAuthTypeProp, bool)`

GetUserAuthTypeOk returns a tuple with the UserAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthType

`func (o *AddConnectionCriteria200Response) SetUserAuthType(v []EnumconnectionCriteriaUserAuthTypeProp)`

SetUserAuthType sets UserAuthType field to given value.

### HasUserAuthType

`func (o *AddConnectionCriteria200Response) HasUserAuthType() bool`

HasUserAuthType returns a boolean if a field has been set.

### GetAuthenticationSecurityLevel

`func (o *AddConnectionCriteria200Response) GetAuthenticationSecurityLevel() EnumconnectionCriteriaAuthenticationSecurityLevelProp`

GetAuthenticationSecurityLevel returns the AuthenticationSecurityLevel field if non-nil, zero value otherwise.

### GetAuthenticationSecurityLevelOk

`func (o *AddConnectionCriteria200Response) GetAuthenticationSecurityLevelOk() (*EnumconnectionCriteriaAuthenticationSecurityLevelProp, bool)`

GetAuthenticationSecurityLevelOk returns a tuple with the AuthenticationSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSecurityLevel

`func (o *AddConnectionCriteria200Response) SetAuthenticationSecurityLevel(v EnumconnectionCriteriaAuthenticationSecurityLevelProp)`

SetAuthenticationSecurityLevel sets AuthenticationSecurityLevel field to given value.

### HasAuthenticationSecurityLevel

`func (o *AddConnectionCriteria200Response) HasAuthenticationSecurityLevel() bool`

HasAuthenticationSecurityLevel returns a boolean if a field has been set.

### GetIncludedUserSASLMechanism

`func (o *AddConnectionCriteria200Response) GetIncludedUserSASLMechanism() []string`

GetIncludedUserSASLMechanism returns the IncludedUserSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedUserSASLMechanismOk

`func (o *AddConnectionCriteria200Response) GetIncludedUserSASLMechanismOk() (*[]string, bool)`

GetIncludedUserSASLMechanismOk returns a tuple with the IncludedUserSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserSASLMechanism

`func (o *AddConnectionCriteria200Response) SetIncludedUserSASLMechanism(v []string)`

SetIncludedUserSASLMechanism sets IncludedUserSASLMechanism field to given value.

### HasIncludedUserSASLMechanism

`func (o *AddConnectionCriteria200Response) HasIncludedUserSASLMechanism() bool`

HasIncludedUserSASLMechanism returns a boolean if a field has been set.

### GetExcludedUserSASLMechanism

`func (o *AddConnectionCriteria200Response) GetExcludedUserSASLMechanism() []string`

GetExcludedUserSASLMechanism returns the ExcludedUserSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedUserSASLMechanismOk

`func (o *AddConnectionCriteria200Response) GetExcludedUserSASLMechanismOk() (*[]string, bool)`

GetExcludedUserSASLMechanismOk returns a tuple with the ExcludedUserSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserSASLMechanism

`func (o *AddConnectionCriteria200Response) SetExcludedUserSASLMechanism(v []string)`

SetExcludedUserSASLMechanism sets ExcludedUserSASLMechanism field to given value.

### HasExcludedUserSASLMechanism

`func (o *AddConnectionCriteria200Response) HasExcludedUserSASLMechanism() bool`

HasExcludedUserSASLMechanism returns a boolean if a field has been set.

### GetIncludedUserBaseDN

`func (o *AddConnectionCriteria200Response) GetIncludedUserBaseDN() []string`

GetIncludedUserBaseDN returns the IncludedUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedUserBaseDNOk

`func (o *AddConnectionCriteria200Response) GetIncludedUserBaseDNOk() (*[]string, bool)`

GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserBaseDN

`func (o *AddConnectionCriteria200Response) SetIncludedUserBaseDN(v []string)`

SetIncludedUserBaseDN sets IncludedUserBaseDN field to given value.

### HasIncludedUserBaseDN

`func (o *AddConnectionCriteria200Response) HasIncludedUserBaseDN() bool`

HasIncludedUserBaseDN returns a boolean if a field has been set.

### GetExcludedUserBaseDN

`func (o *AddConnectionCriteria200Response) GetExcludedUserBaseDN() []string`

GetExcludedUserBaseDN returns the ExcludedUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedUserBaseDNOk

`func (o *AddConnectionCriteria200Response) GetExcludedUserBaseDNOk() (*[]string, bool)`

GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserBaseDN

`func (o *AddConnectionCriteria200Response) SetExcludedUserBaseDN(v []string)`

SetExcludedUserBaseDN sets ExcludedUserBaseDN field to given value.

### HasExcludedUserBaseDN

`func (o *AddConnectionCriteria200Response) HasExcludedUserBaseDN() bool`

HasExcludedUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) GetAllIncludedUserGroupDN() []string`

GetAllIncludedUserGroupDN returns the AllIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedUserGroupDNOk

`func (o *AddConnectionCriteria200Response) GetAllIncludedUserGroupDNOk() (*[]string, bool)`

GetAllIncludedUserGroupDNOk returns a tuple with the AllIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) SetAllIncludedUserGroupDN(v []string)`

SetAllIncludedUserGroupDN sets AllIncludedUserGroupDN field to given value.

### HasAllIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) HasAllIncludedUserGroupDN() bool`

HasAllIncludedUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) GetAnyIncludedUserGroupDN() []string`

GetAnyIncludedUserGroupDN returns the AnyIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedUserGroupDNOk

`func (o *AddConnectionCriteria200Response) GetAnyIncludedUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedUserGroupDNOk returns a tuple with the AnyIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) SetAnyIncludedUserGroupDN(v []string)`

SetAnyIncludedUserGroupDN sets AnyIncludedUserGroupDN field to given value.

### HasAnyIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) HasAnyIncludedUserGroupDN() bool`

HasAnyIncludedUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedUserGroupDN() []string`

GetNotAllIncludedUserGroupDN returns the NotAllIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedUserGroupDNOk

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedUserGroupDNOk returns a tuple with the NotAllIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) SetNotAllIncludedUserGroupDN(v []string)`

SetNotAllIncludedUserGroupDN sets NotAllIncludedUserGroupDN field to given value.

### HasNotAllIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) HasNotAllIncludedUserGroupDN() bool`

HasNotAllIncludedUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) GetNoneIncludedUserGroupDN() []string`

GetNoneIncludedUserGroupDN returns the NoneIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedUserGroupDNOk

`func (o *AddConnectionCriteria200Response) GetNoneIncludedUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedUserGroupDNOk returns a tuple with the NoneIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) SetNoneIncludedUserGroupDN(v []string)`

SetNoneIncludedUserGroupDN sets NoneIncludedUserGroupDN field to given value.

### HasNoneIncludedUserGroupDN

`func (o *AddConnectionCriteria200Response) HasNoneIncludedUserGroupDN() bool`

HasNoneIncludedUserGroupDN returns a boolean if a field has been set.

### GetAllIncludedUserFilter

`func (o *AddConnectionCriteria200Response) GetAllIncludedUserFilter() []string`

GetAllIncludedUserFilter returns the AllIncludedUserFilter field if non-nil, zero value otherwise.

### GetAllIncludedUserFilterOk

`func (o *AddConnectionCriteria200Response) GetAllIncludedUserFilterOk() (*[]string, bool)`

GetAllIncludedUserFilterOk returns a tuple with the AllIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserFilter

`func (o *AddConnectionCriteria200Response) SetAllIncludedUserFilter(v []string)`

SetAllIncludedUserFilter sets AllIncludedUserFilter field to given value.

### HasAllIncludedUserFilter

`func (o *AddConnectionCriteria200Response) HasAllIncludedUserFilter() bool`

HasAllIncludedUserFilter returns a boolean if a field has been set.

### GetAnyIncludedUserFilter

`func (o *AddConnectionCriteria200Response) GetAnyIncludedUserFilter() []string`

GetAnyIncludedUserFilter returns the AnyIncludedUserFilter field if non-nil, zero value otherwise.

### GetAnyIncludedUserFilterOk

`func (o *AddConnectionCriteria200Response) GetAnyIncludedUserFilterOk() (*[]string, bool)`

GetAnyIncludedUserFilterOk returns a tuple with the AnyIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserFilter

`func (o *AddConnectionCriteria200Response) SetAnyIncludedUserFilter(v []string)`

SetAnyIncludedUserFilter sets AnyIncludedUserFilter field to given value.

### HasAnyIncludedUserFilter

`func (o *AddConnectionCriteria200Response) HasAnyIncludedUserFilter() bool`

HasAnyIncludedUserFilter returns a boolean if a field has been set.

### GetNotAllIncludedUserFilter

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedUserFilter() []string`

GetNotAllIncludedUserFilter returns the NotAllIncludedUserFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedUserFilterOk

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedUserFilterOk() (*[]string, bool)`

GetNotAllIncludedUserFilterOk returns a tuple with the NotAllIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserFilter

`func (o *AddConnectionCriteria200Response) SetNotAllIncludedUserFilter(v []string)`

SetNotAllIncludedUserFilter sets NotAllIncludedUserFilter field to given value.

### HasNotAllIncludedUserFilter

`func (o *AddConnectionCriteria200Response) HasNotAllIncludedUserFilter() bool`

HasNotAllIncludedUserFilter returns a boolean if a field has been set.

### GetNoneIncludedUserFilter

`func (o *AddConnectionCriteria200Response) GetNoneIncludedUserFilter() []string`

GetNoneIncludedUserFilter returns the NoneIncludedUserFilter field if non-nil, zero value otherwise.

### GetNoneIncludedUserFilterOk

`func (o *AddConnectionCriteria200Response) GetNoneIncludedUserFilterOk() (*[]string, bool)`

GetNoneIncludedUserFilterOk returns a tuple with the NoneIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserFilter

`func (o *AddConnectionCriteria200Response) SetNoneIncludedUserFilter(v []string)`

SetNoneIncludedUserFilter sets NoneIncludedUserFilter field to given value.

### HasNoneIncludedUserFilter

`func (o *AddConnectionCriteria200Response) HasNoneIncludedUserFilter() bool`

HasNoneIncludedUserFilter returns a boolean if a field has been set.

### GetAllIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) GetAllIncludedUserPrivilege() []EnumconnectionCriteriaAllIncludedUserPrivilegeProp`

GetAllIncludedUserPrivilege returns the AllIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetAllIncludedUserPrivilegeOk

`func (o *AddConnectionCriteria200Response) GetAllIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaAllIncludedUserPrivilegeProp, bool)`

GetAllIncludedUserPrivilegeOk returns a tuple with the AllIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) SetAllIncludedUserPrivilege(v []EnumconnectionCriteriaAllIncludedUserPrivilegeProp)`

SetAllIncludedUserPrivilege sets AllIncludedUserPrivilege field to given value.

### HasAllIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) HasAllIncludedUserPrivilege() bool`

HasAllIncludedUserPrivilege returns a boolean if a field has been set.

### GetAnyIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) GetAnyIncludedUserPrivilege() []EnumconnectionCriteriaAnyIncludedUserPrivilegeProp`

GetAnyIncludedUserPrivilege returns the AnyIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetAnyIncludedUserPrivilegeOk

`func (o *AddConnectionCriteria200Response) GetAnyIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaAnyIncludedUserPrivilegeProp, bool)`

GetAnyIncludedUserPrivilegeOk returns a tuple with the AnyIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) SetAnyIncludedUserPrivilege(v []EnumconnectionCriteriaAnyIncludedUserPrivilegeProp)`

SetAnyIncludedUserPrivilege sets AnyIncludedUserPrivilege field to given value.

### HasAnyIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) HasAnyIncludedUserPrivilege() bool`

HasAnyIncludedUserPrivilege returns a boolean if a field has been set.

### GetNotAllIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedUserPrivilege() []EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp`

GetNotAllIncludedUserPrivilege returns the NotAllIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetNotAllIncludedUserPrivilegeOk

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp, bool)`

GetNotAllIncludedUserPrivilegeOk returns a tuple with the NotAllIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) SetNotAllIncludedUserPrivilege(v []EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp)`

SetNotAllIncludedUserPrivilege sets NotAllIncludedUserPrivilege field to given value.

### HasNotAllIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) HasNotAllIncludedUserPrivilege() bool`

HasNotAllIncludedUserPrivilege returns a boolean if a field has been set.

### GetNoneIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) GetNoneIncludedUserPrivilege() []EnumconnectionCriteriaNoneIncludedUserPrivilegeProp`

GetNoneIncludedUserPrivilege returns the NoneIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetNoneIncludedUserPrivilegeOk

`func (o *AddConnectionCriteria200Response) GetNoneIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaNoneIncludedUserPrivilegeProp, bool)`

GetNoneIncludedUserPrivilegeOk returns a tuple with the NoneIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) SetNoneIncludedUserPrivilege(v []EnumconnectionCriteriaNoneIncludedUserPrivilegeProp)`

SetNoneIncludedUserPrivilege sets NoneIncludedUserPrivilege field to given value.

### HasNoneIncludedUserPrivilege

`func (o *AddConnectionCriteria200Response) HasNoneIncludedUserPrivilege() bool`

HasNoneIncludedUserPrivilege returns a boolean if a field has been set.

### GetDescription

`func (o *AddConnectionCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddConnectionCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddConnectionCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddConnectionCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddConnectionCriteria200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddConnectionCriteria200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddConnectionCriteria200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddConnectionCriteria200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddConnectionCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddConnectionCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddConnectionCriteria200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddConnectionCriteria200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetAllIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) GetAllIncludedConnectionCriteria() []string`

GetAllIncludedConnectionCriteria returns the AllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAllIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteria200Response) GetAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAllIncludedConnectionCriteriaOk returns a tuple with the AllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) SetAllIncludedConnectionCriteria(v []string)`

SetAllIncludedConnectionCriteria sets AllIncludedConnectionCriteria field to given value.

### HasAllIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) HasAllIncludedConnectionCriteria() bool`

HasAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetAnyIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) GetAnyIncludedConnectionCriteria() []string`

GetAnyIncludedConnectionCriteria returns the AnyIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteria200Response) GetAnyIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAnyIncludedConnectionCriteriaOk returns a tuple with the AnyIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) SetAnyIncludedConnectionCriteria(v []string)`

SetAnyIncludedConnectionCriteria sets AnyIncludedConnectionCriteria field to given value.

### HasAnyIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) HasAnyIncludedConnectionCriteria() bool`

HasAnyIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNotAllIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedConnectionCriteria() []string`

GetNotAllIncludedConnectionCriteria returns the NotAllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteria200Response) GetNotAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNotAllIncludedConnectionCriteriaOk returns a tuple with the NotAllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) SetNotAllIncludedConnectionCriteria(v []string)`

SetNotAllIncludedConnectionCriteria sets NotAllIncludedConnectionCriteria field to given value.

### HasNotAllIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) HasNotAllIncludedConnectionCriteria() bool`

HasNotAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNoneIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) GetNoneIncludedConnectionCriteria() []string`

GetNoneIncludedConnectionCriteria returns the NoneIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteria200Response) GetNoneIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNoneIncludedConnectionCriteriaOk returns a tuple with the NoneIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) SetNoneIncludedConnectionCriteria(v []string)`

SetNoneIncludedConnectionCriteria sets NoneIncludedConnectionCriteria field to given value.

### HasNoneIncludedConnectionCriteria

`func (o *AddConnectionCriteria200Response) HasNoneIncludedConnectionCriteria() bool`

HasNoneIncludedConnectionCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddConnectionCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddConnectionCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddConnectionCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddConnectionCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddConnectionCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddConnectionCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddConnectionCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


