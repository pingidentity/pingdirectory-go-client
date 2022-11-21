# AddConnectionCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Connection Criteria | 
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
**AllIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that must match the associated client connection in order to match the aggregate connection criteria. If one or more all-included connection criteria objects are provided, then a client connection must match all of them in order to match the aggregate connection criteria. | [optional] 
**AnyIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that may match the associated client connection in order to match the aggregate connection criteria. If one or more any-included connection criteria objects are provided, then a client connection must match at least one of them in order to match the aggregate connection criteria. | [optional] 
**NotAllIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that should not match the associated client connection in order to match the aggregate connection criteria. If one or more not-all-included connection criteria objects are provided, then a client connection must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate connection criteria. | [optional] 
**NoneIncludedConnectionCriteria** | Pointer to **[]string** | Specifies a connection criteria object that must not match the associated client connection in order to match the aggregate connection criteria. If one or more none-included connection criteria objects are provided, then a client connection must not match any of them in order to match the aggregate connection criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Connection Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Connection Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddConnectionCriteriaRequest

`func NewAddConnectionCriteriaRequest(criteriaName string, schemas []EnumthirdPartyConnectionCriteriaSchemaUrn, extensionClass string, ) *AddConnectionCriteriaRequest`

NewAddConnectionCriteriaRequest instantiates a new AddConnectionCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConnectionCriteriaRequestWithDefaults

`func NewAddConnectionCriteriaRequestWithDefaults() *AddConnectionCriteriaRequest`

NewAddConnectionCriteriaRequestWithDefaults instantiates a new AddConnectionCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddConnectionCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddConnectionCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddConnectionCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddConnectionCriteriaRequest) GetSchemas() []EnumthirdPartyConnectionCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConnectionCriteriaRequest) GetSchemasOk() (*[]EnumthirdPartyConnectionCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConnectionCriteriaRequest) SetSchemas(v []EnumthirdPartyConnectionCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIncludedClientAddress

`func (o *AddConnectionCriteriaRequest) GetIncludedClientAddress() []string`

GetIncludedClientAddress returns the IncludedClientAddress field if non-nil, zero value otherwise.

### GetIncludedClientAddressOk

`func (o *AddConnectionCriteriaRequest) GetIncludedClientAddressOk() (*[]string, bool)`

GetIncludedClientAddressOk returns a tuple with the IncludedClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientAddress

`func (o *AddConnectionCriteriaRequest) SetIncludedClientAddress(v []string)`

SetIncludedClientAddress sets IncludedClientAddress field to given value.

### HasIncludedClientAddress

`func (o *AddConnectionCriteriaRequest) HasIncludedClientAddress() bool`

HasIncludedClientAddress returns a boolean if a field has been set.

### GetExcludedClientAddress

`func (o *AddConnectionCriteriaRequest) GetExcludedClientAddress() []string`

GetExcludedClientAddress returns the ExcludedClientAddress field if non-nil, zero value otherwise.

### GetExcludedClientAddressOk

`func (o *AddConnectionCriteriaRequest) GetExcludedClientAddressOk() (*[]string, bool)`

GetExcludedClientAddressOk returns a tuple with the ExcludedClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientAddress

`func (o *AddConnectionCriteriaRequest) SetExcludedClientAddress(v []string)`

SetExcludedClientAddress sets ExcludedClientAddress field to given value.

### HasExcludedClientAddress

`func (o *AddConnectionCriteriaRequest) HasExcludedClientAddress() bool`

HasExcludedClientAddress returns a boolean if a field has been set.

### GetIncludedConnectionHandler

`func (o *AddConnectionCriteriaRequest) GetIncludedConnectionHandler() []string`

GetIncludedConnectionHandler returns the IncludedConnectionHandler field if non-nil, zero value otherwise.

### GetIncludedConnectionHandlerOk

`func (o *AddConnectionCriteriaRequest) GetIncludedConnectionHandlerOk() (*[]string, bool)`

GetIncludedConnectionHandlerOk returns a tuple with the IncludedConnectionHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedConnectionHandler

`func (o *AddConnectionCriteriaRequest) SetIncludedConnectionHandler(v []string)`

SetIncludedConnectionHandler sets IncludedConnectionHandler field to given value.

### HasIncludedConnectionHandler

`func (o *AddConnectionCriteriaRequest) HasIncludedConnectionHandler() bool`

HasIncludedConnectionHandler returns a boolean if a field has been set.

### GetExcludedConnectionHandler

`func (o *AddConnectionCriteriaRequest) GetExcludedConnectionHandler() []string`

GetExcludedConnectionHandler returns the ExcludedConnectionHandler field if non-nil, zero value otherwise.

### GetExcludedConnectionHandlerOk

`func (o *AddConnectionCriteriaRequest) GetExcludedConnectionHandlerOk() (*[]string, bool)`

GetExcludedConnectionHandlerOk returns a tuple with the ExcludedConnectionHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedConnectionHandler

`func (o *AddConnectionCriteriaRequest) SetExcludedConnectionHandler(v []string)`

SetExcludedConnectionHandler sets ExcludedConnectionHandler field to given value.

### HasExcludedConnectionHandler

`func (o *AddConnectionCriteriaRequest) HasExcludedConnectionHandler() bool`

HasExcludedConnectionHandler returns a boolean if a field has been set.

### GetIncludedProtocol

`func (o *AddConnectionCriteriaRequest) GetIncludedProtocol() []string`

GetIncludedProtocol returns the IncludedProtocol field if non-nil, zero value otherwise.

### GetIncludedProtocolOk

`func (o *AddConnectionCriteriaRequest) GetIncludedProtocolOk() (*[]string, bool)`

GetIncludedProtocolOk returns a tuple with the IncludedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedProtocol

`func (o *AddConnectionCriteriaRequest) SetIncludedProtocol(v []string)`

SetIncludedProtocol sets IncludedProtocol field to given value.

### HasIncludedProtocol

`func (o *AddConnectionCriteriaRequest) HasIncludedProtocol() bool`

HasIncludedProtocol returns a boolean if a field has been set.

### GetExcludedProtocol

`func (o *AddConnectionCriteriaRequest) GetExcludedProtocol() []string`

GetExcludedProtocol returns the ExcludedProtocol field if non-nil, zero value otherwise.

### GetExcludedProtocolOk

`func (o *AddConnectionCriteriaRequest) GetExcludedProtocolOk() (*[]string, bool)`

GetExcludedProtocolOk returns a tuple with the ExcludedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProtocol

`func (o *AddConnectionCriteriaRequest) SetExcludedProtocol(v []string)`

SetExcludedProtocol sets ExcludedProtocol field to given value.

### HasExcludedProtocol

`func (o *AddConnectionCriteriaRequest) HasExcludedProtocol() bool`

HasExcludedProtocol returns a boolean if a field has been set.

### GetCommunicationSecurityLevel

`func (o *AddConnectionCriteriaRequest) GetCommunicationSecurityLevel() EnumconnectionCriteriaCommunicationSecurityLevelProp`

GetCommunicationSecurityLevel returns the CommunicationSecurityLevel field if non-nil, zero value otherwise.

### GetCommunicationSecurityLevelOk

`func (o *AddConnectionCriteriaRequest) GetCommunicationSecurityLevelOk() (*EnumconnectionCriteriaCommunicationSecurityLevelProp, bool)`

GetCommunicationSecurityLevelOk returns a tuple with the CommunicationSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationSecurityLevel

`func (o *AddConnectionCriteriaRequest) SetCommunicationSecurityLevel(v EnumconnectionCriteriaCommunicationSecurityLevelProp)`

SetCommunicationSecurityLevel sets CommunicationSecurityLevel field to given value.

### HasCommunicationSecurityLevel

`func (o *AddConnectionCriteriaRequest) HasCommunicationSecurityLevel() bool`

HasCommunicationSecurityLevel returns a boolean if a field has been set.

### GetUserAuthType

`func (o *AddConnectionCriteriaRequest) GetUserAuthType() []EnumconnectionCriteriaUserAuthTypeProp`

GetUserAuthType returns the UserAuthType field if non-nil, zero value otherwise.

### GetUserAuthTypeOk

`func (o *AddConnectionCriteriaRequest) GetUserAuthTypeOk() (*[]EnumconnectionCriteriaUserAuthTypeProp, bool)`

GetUserAuthTypeOk returns a tuple with the UserAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthType

`func (o *AddConnectionCriteriaRequest) SetUserAuthType(v []EnumconnectionCriteriaUserAuthTypeProp)`

SetUserAuthType sets UserAuthType field to given value.

### HasUserAuthType

`func (o *AddConnectionCriteriaRequest) HasUserAuthType() bool`

HasUserAuthType returns a boolean if a field has been set.

### GetAuthenticationSecurityLevel

`func (o *AddConnectionCriteriaRequest) GetAuthenticationSecurityLevel() EnumconnectionCriteriaAuthenticationSecurityLevelProp`

GetAuthenticationSecurityLevel returns the AuthenticationSecurityLevel field if non-nil, zero value otherwise.

### GetAuthenticationSecurityLevelOk

`func (o *AddConnectionCriteriaRequest) GetAuthenticationSecurityLevelOk() (*EnumconnectionCriteriaAuthenticationSecurityLevelProp, bool)`

GetAuthenticationSecurityLevelOk returns a tuple with the AuthenticationSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSecurityLevel

`func (o *AddConnectionCriteriaRequest) SetAuthenticationSecurityLevel(v EnumconnectionCriteriaAuthenticationSecurityLevelProp)`

SetAuthenticationSecurityLevel sets AuthenticationSecurityLevel field to given value.

### HasAuthenticationSecurityLevel

`func (o *AddConnectionCriteriaRequest) HasAuthenticationSecurityLevel() bool`

HasAuthenticationSecurityLevel returns a boolean if a field has been set.

### GetIncludedUserSASLMechanism

`func (o *AddConnectionCriteriaRequest) GetIncludedUserSASLMechanism() []string`

GetIncludedUserSASLMechanism returns the IncludedUserSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedUserSASLMechanismOk

`func (o *AddConnectionCriteriaRequest) GetIncludedUserSASLMechanismOk() (*[]string, bool)`

GetIncludedUserSASLMechanismOk returns a tuple with the IncludedUserSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserSASLMechanism

`func (o *AddConnectionCriteriaRequest) SetIncludedUserSASLMechanism(v []string)`

SetIncludedUserSASLMechanism sets IncludedUserSASLMechanism field to given value.

### HasIncludedUserSASLMechanism

`func (o *AddConnectionCriteriaRequest) HasIncludedUserSASLMechanism() bool`

HasIncludedUserSASLMechanism returns a boolean if a field has been set.

### GetExcludedUserSASLMechanism

`func (o *AddConnectionCriteriaRequest) GetExcludedUserSASLMechanism() []string`

GetExcludedUserSASLMechanism returns the ExcludedUserSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedUserSASLMechanismOk

`func (o *AddConnectionCriteriaRequest) GetExcludedUserSASLMechanismOk() (*[]string, bool)`

GetExcludedUserSASLMechanismOk returns a tuple with the ExcludedUserSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserSASLMechanism

`func (o *AddConnectionCriteriaRequest) SetExcludedUserSASLMechanism(v []string)`

SetExcludedUserSASLMechanism sets ExcludedUserSASLMechanism field to given value.

### HasExcludedUserSASLMechanism

`func (o *AddConnectionCriteriaRequest) HasExcludedUserSASLMechanism() bool`

HasExcludedUserSASLMechanism returns a boolean if a field has been set.

### GetIncludedUserBaseDN

`func (o *AddConnectionCriteriaRequest) GetIncludedUserBaseDN() []string`

GetIncludedUserBaseDN returns the IncludedUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedUserBaseDNOk

`func (o *AddConnectionCriteriaRequest) GetIncludedUserBaseDNOk() (*[]string, bool)`

GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserBaseDN

`func (o *AddConnectionCriteriaRequest) SetIncludedUserBaseDN(v []string)`

SetIncludedUserBaseDN sets IncludedUserBaseDN field to given value.

### HasIncludedUserBaseDN

`func (o *AddConnectionCriteriaRequest) HasIncludedUserBaseDN() bool`

HasIncludedUserBaseDN returns a boolean if a field has been set.

### GetExcludedUserBaseDN

`func (o *AddConnectionCriteriaRequest) GetExcludedUserBaseDN() []string`

GetExcludedUserBaseDN returns the ExcludedUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedUserBaseDNOk

`func (o *AddConnectionCriteriaRequest) GetExcludedUserBaseDNOk() (*[]string, bool)`

GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserBaseDN

`func (o *AddConnectionCriteriaRequest) SetExcludedUserBaseDN(v []string)`

SetExcludedUserBaseDN sets ExcludedUserBaseDN field to given value.

### HasExcludedUserBaseDN

`func (o *AddConnectionCriteriaRequest) HasExcludedUserBaseDN() bool`

HasExcludedUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) GetAllIncludedUserGroupDN() []string`

GetAllIncludedUserGroupDN returns the AllIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedUserGroupDNOk

`func (o *AddConnectionCriteriaRequest) GetAllIncludedUserGroupDNOk() (*[]string, bool)`

GetAllIncludedUserGroupDNOk returns a tuple with the AllIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) SetAllIncludedUserGroupDN(v []string)`

SetAllIncludedUserGroupDN sets AllIncludedUserGroupDN field to given value.

### HasAllIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) HasAllIncludedUserGroupDN() bool`

HasAllIncludedUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedUserGroupDN() []string`

GetAnyIncludedUserGroupDN returns the AnyIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedUserGroupDNOk

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedUserGroupDNOk returns a tuple with the AnyIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) SetAnyIncludedUserGroupDN(v []string)`

SetAnyIncludedUserGroupDN sets AnyIncludedUserGroupDN field to given value.

### HasAnyIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) HasAnyIncludedUserGroupDN() bool`

HasAnyIncludedUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedUserGroupDN() []string`

GetNotAllIncludedUserGroupDN returns the NotAllIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedUserGroupDNOk

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedUserGroupDNOk returns a tuple with the NotAllIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) SetNotAllIncludedUserGroupDN(v []string)`

SetNotAllIncludedUserGroupDN sets NotAllIncludedUserGroupDN field to given value.

### HasNotAllIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) HasNotAllIncludedUserGroupDN() bool`

HasNotAllIncludedUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedUserGroupDN() []string`

GetNoneIncludedUserGroupDN returns the NoneIncludedUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedUserGroupDNOk

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedUserGroupDNOk returns a tuple with the NoneIncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) SetNoneIncludedUserGroupDN(v []string)`

SetNoneIncludedUserGroupDN sets NoneIncludedUserGroupDN field to given value.

### HasNoneIncludedUserGroupDN

`func (o *AddConnectionCriteriaRequest) HasNoneIncludedUserGroupDN() bool`

HasNoneIncludedUserGroupDN returns a boolean if a field has been set.

### GetAllIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) GetAllIncludedUserFilter() []string`

GetAllIncludedUserFilter returns the AllIncludedUserFilter field if non-nil, zero value otherwise.

### GetAllIncludedUserFilterOk

`func (o *AddConnectionCriteriaRequest) GetAllIncludedUserFilterOk() (*[]string, bool)`

GetAllIncludedUserFilterOk returns a tuple with the AllIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) SetAllIncludedUserFilter(v []string)`

SetAllIncludedUserFilter sets AllIncludedUserFilter field to given value.

### HasAllIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) HasAllIncludedUserFilter() bool`

HasAllIncludedUserFilter returns a boolean if a field has been set.

### GetAnyIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedUserFilter() []string`

GetAnyIncludedUserFilter returns the AnyIncludedUserFilter field if non-nil, zero value otherwise.

### GetAnyIncludedUserFilterOk

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedUserFilterOk() (*[]string, bool)`

GetAnyIncludedUserFilterOk returns a tuple with the AnyIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) SetAnyIncludedUserFilter(v []string)`

SetAnyIncludedUserFilter sets AnyIncludedUserFilter field to given value.

### HasAnyIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) HasAnyIncludedUserFilter() bool`

HasAnyIncludedUserFilter returns a boolean if a field has been set.

### GetNotAllIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedUserFilter() []string`

GetNotAllIncludedUserFilter returns the NotAllIncludedUserFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedUserFilterOk

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedUserFilterOk() (*[]string, bool)`

GetNotAllIncludedUserFilterOk returns a tuple with the NotAllIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) SetNotAllIncludedUserFilter(v []string)`

SetNotAllIncludedUserFilter sets NotAllIncludedUserFilter field to given value.

### HasNotAllIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) HasNotAllIncludedUserFilter() bool`

HasNotAllIncludedUserFilter returns a boolean if a field has been set.

### GetNoneIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedUserFilter() []string`

GetNoneIncludedUserFilter returns the NoneIncludedUserFilter field if non-nil, zero value otherwise.

### GetNoneIncludedUserFilterOk

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedUserFilterOk() (*[]string, bool)`

GetNoneIncludedUserFilterOk returns a tuple with the NoneIncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) SetNoneIncludedUserFilter(v []string)`

SetNoneIncludedUserFilter sets NoneIncludedUserFilter field to given value.

### HasNoneIncludedUserFilter

`func (o *AddConnectionCriteriaRequest) HasNoneIncludedUserFilter() bool`

HasNoneIncludedUserFilter returns a boolean if a field has been set.

### GetAllIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) GetAllIncludedUserPrivilege() []EnumconnectionCriteriaAllIncludedUserPrivilegeProp`

GetAllIncludedUserPrivilege returns the AllIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetAllIncludedUserPrivilegeOk

`func (o *AddConnectionCriteriaRequest) GetAllIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaAllIncludedUserPrivilegeProp, bool)`

GetAllIncludedUserPrivilegeOk returns a tuple with the AllIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) SetAllIncludedUserPrivilege(v []EnumconnectionCriteriaAllIncludedUserPrivilegeProp)`

SetAllIncludedUserPrivilege sets AllIncludedUserPrivilege field to given value.

### HasAllIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) HasAllIncludedUserPrivilege() bool`

HasAllIncludedUserPrivilege returns a boolean if a field has been set.

### GetAnyIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedUserPrivilege() []EnumconnectionCriteriaAnyIncludedUserPrivilegeProp`

GetAnyIncludedUserPrivilege returns the AnyIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetAnyIncludedUserPrivilegeOk

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaAnyIncludedUserPrivilegeProp, bool)`

GetAnyIncludedUserPrivilegeOk returns a tuple with the AnyIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) SetAnyIncludedUserPrivilege(v []EnumconnectionCriteriaAnyIncludedUserPrivilegeProp)`

SetAnyIncludedUserPrivilege sets AnyIncludedUserPrivilege field to given value.

### HasAnyIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) HasAnyIncludedUserPrivilege() bool`

HasAnyIncludedUserPrivilege returns a boolean if a field has been set.

### GetNotAllIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedUserPrivilege() []EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp`

GetNotAllIncludedUserPrivilege returns the NotAllIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetNotAllIncludedUserPrivilegeOk

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp, bool)`

GetNotAllIncludedUserPrivilegeOk returns a tuple with the NotAllIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) SetNotAllIncludedUserPrivilege(v []EnumconnectionCriteriaNotAllIncludedUserPrivilegeProp)`

SetNotAllIncludedUserPrivilege sets NotAllIncludedUserPrivilege field to given value.

### HasNotAllIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) HasNotAllIncludedUserPrivilege() bool`

HasNotAllIncludedUserPrivilege returns a boolean if a field has been set.

### GetNoneIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedUserPrivilege() []EnumconnectionCriteriaNoneIncludedUserPrivilegeProp`

GetNoneIncludedUserPrivilege returns the NoneIncludedUserPrivilege field if non-nil, zero value otherwise.

### GetNoneIncludedUserPrivilegeOk

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedUserPrivilegeOk() (*[]EnumconnectionCriteriaNoneIncludedUserPrivilegeProp, bool)`

GetNoneIncludedUserPrivilegeOk returns a tuple with the NoneIncludedUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) SetNoneIncludedUserPrivilege(v []EnumconnectionCriteriaNoneIncludedUserPrivilegeProp)`

SetNoneIncludedUserPrivilege sets NoneIncludedUserPrivilege field to given value.

### HasNoneIncludedUserPrivilege

`func (o *AddConnectionCriteriaRequest) HasNoneIncludedUserPrivilege() bool`

HasNoneIncludedUserPrivilege returns a boolean if a field has been set.

### GetDescription

`func (o *AddConnectionCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddConnectionCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddConnectionCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddConnectionCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) GetAllIncludedConnectionCriteria() []string`

GetAllIncludedConnectionCriteria returns the AllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAllIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteriaRequest) GetAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAllIncludedConnectionCriteriaOk returns a tuple with the AllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) SetAllIncludedConnectionCriteria(v []string)`

SetAllIncludedConnectionCriteria sets AllIncludedConnectionCriteria field to given value.

### HasAllIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) HasAllIncludedConnectionCriteria() bool`

HasAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetAnyIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedConnectionCriteria() []string`

GetAnyIncludedConnectionCriteria returns the AnyIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteriaRequest) GetAnyIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAnyIncludedConnectionCriteriaOk returns a tuple with the AnyIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) SetAnyIncludedConnectionCriteria(v []string)`

SetAnyIncludedConnectionCriteria sets AnyIncludedConnectionCriteria field to given value.

### HasAnyIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) HasAnyIncludedConnectionCriteria() bool`

HasAnyIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNotAllIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedConnectionCriteria() []string`

GetNotAllIncludedConnectionCriteria returns the NotAllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteriaRequest) GetNotAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNotAllIncludedConnectionCriteriaOk returns a tuple with the NotAllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) SetNotAllIncludedConnectionCriteria(v []string)`

SetNotAllIncludedConnectionCriteria sets NotAllIncludedConnectionCriteria field to given value.

### HasNotAllIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) HasNotAllIncludedConnectionCriteria() bool`

HasNotAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNoneIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedConnectionCriteria() []string`

GetNoneIncludedConnectionCriteria returns the NoneIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedConnectionCriteriaOk

`func (o *AddConnectionCriteriaRequest) GetNoneIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNoneIncludedConnectionCriteriaOk returns a tuple with the NoneIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) SetNoneIncludedConnectionCriteria(v []string)`

SetNoneIncludedConnectionCriteria sets NoneIncludedConnectionCriteria field to given value.

### HasNoneIncludedConnectionCriteria

`func (o *AddConnectionCriteriaRequest) HasNoneIncludedConnectionCriteria() bool`

HasNoneIncludedConnectionCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddConnectionCriteriaRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddConnectionCriteriaRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddConnectionCriteriaRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddConnectionCriteriaRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddConnectionCriteriaRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddConnectionCriteriaRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddConnectionCriteriaRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


