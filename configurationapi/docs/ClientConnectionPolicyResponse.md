# ClientConnectionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Client Connection Policy | 
**Schemas** | Pointer to [**[]EnumclientConnectionPolicySchemaUrn**](EnumclientConnectionPolicySchemaUrn.md) |  | [optional] 
**PolicyID** | **string** | Specifies a name which uniquely identifies this Client Connection Policy in the server. | 
**Description** | Pointer to **string** | A description for this Client Connection Policy | [optional] 
**Enabled** | **bool** | Indicates whether this Client Connection Policy is enabled for use in the server. If a Client Connection Policy is disabled, then no new client connections will be associated with it. | 
**EvaluationOrderIndex** | **int32** | Specifies the order in which Client Connection Policy definitions will be evaluated. A Client Connection Policy with a lower index will be evaluated before one with a higher index, and the first Client Connection Policy evaluated which may apply to a client connection will be used for that connection. Each Client Connection Policy must be assigned a unique evaluation order index value. | 
**ConnectionCriteria** | Pointer to **string** | Specifies a set of connection criteria that must match the associated client connection for it to be associated with this Client Connection Policy. | [optional] 
**TerminateConnection** | Pointer to **bool** | Indicates whether any client connection for which this Client Connection Policy is selected should be terminated. This makes it possible to define fine-grained criteria for clients that should not be allowed to connect to this Directory Server. | [optional] 
**SensitiveAttribute** | Pointer to **[]string** | Provides the ability to indicate that some attributes should be considered sensitive and additional protection should be in place when interacting with those attributes. | [optional] 
**ExcludeGlobalSensitiveAttribute** | Pointer to **[]string** | Specifies the set of global sensitive attribute definitions that should not apply to this client connection policy. | [optional] 
**ResultCodeMap** | Pointer to **string** | Specifies the result code map that should be used for clients associated with this Client Connection Policy. If a value is defined for this property, then it will override any result code map referenced in the global configuration. | [optional] 
**IncludedBackendBaseDN** | Pointer to **[]string** | Specifies the set of backend base DNs for which subtree views should be included in this Client Connection Policy. | [optional] 
**ExcludedBackendBaseDN** | Pointer to **[]string** | Specifies the set of backend base DNs for which subtree views should be excluded from this Client Connection Policy. | [optional] 
**AllowedOperation** | [**[]EnumclientConnectionPolicyAllowedOperationProp**](EnumclientConnectionPolicyAllowedOperationProp.md) |  | 
**RequiredOperationRequestCriteria** | Pointer to **string** | Specifies a request criteria object that will be required to match all requests submitted by clients associated with this Client Connection Policy. If a client submits a request that does not satisfy this request criteria object, then that request will be rejected. | [optional] 
**ProhibitedOperationRequestCriteria** | Pointer to **string** | Specifies a request criteria object that must not match any requests submitted by clients associated with this Client Connection Policy. If a client submits a request that satisfies this request criteria object, then that request will be rejected. | [optional] 
**AllowedRequestControl** | Pointer to **[]string** | Specifies the OIDs of the controls that clients associated with this Client Connection Policy will be allowed to include in requests. | [optional] 
**DeniedRequestControl** | Pointer to **[]string** | Specifies the OIDs of the controls that clients associated with this Client Connection Policy will not be allowed to include in requests. | [optional] 
**AllowedExtendedOperation** | Pointer to **[]string** | Specifies the OIDs of the extended operations that clients associated with this Client Connection Policy will be allowed to request. | [optional] 
**DeniedExtendedOperation** | Pointer to **[]string** | Specifies the OIDs of the extended operations that clients associated with this Client Connection Policy will not be allowed to request. | [optional] 
**AllowedAuthType** | [**[]EnumclientConnectionPolicyAllowedAuthTypeProp**](EnumclientConnectionPolicyAllowedAuthTypeProp.md) |  | 
**AllowedSASLMechanism** | Pointer to **[]string** | Specifies the names of the SASL mechanisms that clients associated with this Client Connection Policy will be allowed to request. | [optional] 
**DeniedSASLMechanism** | Pointer to **[]string** | Specifies the names of the SASL mechanisms that clients associated with this Client Connection Policy will not be allowed to request. | [optional] 
**AllowedFilterType** | Pointer to [**[]EnumclientConnectionPolicyAllowedFilterTypeProp**](EnumclientConnectionPolicyAllowedFilterTypeProp.md) |  | [optional] 
**AllowUnindexedSearches** | Pointer to **bool** | Indicates whether clients will be allowed to request search operations that cannot be efficiently processed using the set of indexes defined in the corresponding backend. Note that even if this is false, some clients may be able to request unindexed searches if the allow-unindexed-searches-with-control property has a value of true and the necessary conditions are satisfied. | [optional] 
**AllowUnindexedSearchesWithControl** | Pointer to **bool** | Indicates whether clients will be allowed to request search operations that cannot be efficiently processed using the set of indexes defined in the corresponding backend, as long as the search request also includes the permit unindexed search request control and the requester has the unindexed-search-with-control privilege (or that privilege is disabled in the global configuration). | [optional] 
**MinimumSubstringLength** | Pointer to **int32** | Specifies the minimum number of consecutive bytes that must be present in any subInitial, subAny, or subFinal element of a substring filter component (i.e., the minimum number of consecutive bytes between wildcard characters in a substring filter). Any attempt to use a substring search with an element containing fewer than this number of bytes will be rejected. | [optional] 
**MaximumConcurrentConnections** | Pointer to **int32** | Specifies the maximum number of client connections which may be associated with this Client Connection Policy at any given time. | [optional] 
**MaximumConnectionDuration** | Pointer to **string** | Specifies the maximum length of time that a connection associated with this Client Connection Policy may be established. Any connection which is associated with this Client Connection Policy and has been established for longer than this period of time may be terminated. | [optional] 
**MaximumIdleConnectionDuration** | Pointer to **string** | Specifies the maximum length of time that a connection associated with this Client Connection Policy may remain established after the completion of the last operation processed on that connection. Any new operation requested on the connection will reset this timer. Any connection associated with this Client Connection Policy which has been idle for longer than this length of time may be terminated. | [optional] 
**MaximumOperationCountPerConnection** | Pointer to **int32** | Specifies the maximum number of operations that may be requested by any client connection associated with this Client Connection Policy. If an attempt is made to process more than this number of operations on a client connection, then that connection will be terminated. | [optional] 
**MaximumConcurrentOperationsPerConnection** | Pointer to **int32** | Specifies the maximum number of concurrent operations that can be in progress for any connection. This can help prevent a single client connection from monopolizing server processing resources by sending a large number of concurrent asynchronous requests. A value of zero indicates that no limit will be placed on the number of concurrent requests for a single client. | [optional] 
**MaximumConcurrentOperationWaitTimeBeforeRejecting** | Pointer to **string** | Specifies the maximum length of time that the server should wait for an outstanding operation to complete before rejecting a new request received when the maximum number of outstanding operations are already in progress on that connection. If an existing outstanding operation on the connection completes before this time, then the operation will be processed. Otherwise, the operation will be rejected with a \&quot;busy\&quot; result. | [optional] 
**MaximumConcurrentOperationsPerConnectionExceededBehavior** | Pointer to [**EnumclientConnectionPolicyMaximumConcurrentOperationsPerConnectionExceededBehaviorProp**](EnumclientConnectionPolicyMaximumConcurrentOperationsPerConnectionExceededBehaviorProp.md) |  | [optional] 
**MaximumConnectionOperationRate** | Pointer to **[]string** | Specifies the maximum rate at which a client associated with this Client Connection Policy may issue requests to the Directory Server. If any client attempts to request operations at a rate higher than this limit, then the server will exhibit the behavior described in the connection-operation-rate-exceeded-behavior property. | [optional] 
**ConnectionOperationRateExceededBehavior** | Pointer to [**EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp**](EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp.md) |  | [optional] 
**MaximumPolicyOperationRate** | Pointer to **[]string** | Specifies the maximum rate at which all clients associated with this Client Connection Policy, as a collective set, may issue requests to the Directory Server. If this limit is exceeded, then the server will exhibit the behavior described in the policy-operation-rate-exceeded-behavior property. | [optional] 
**PolicyOperationRateExceededBehavior** | Pointer to [**EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp**](EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp.md) |  | [optional] 
**MaximumSearchSizeLimit** | Pointer to **int32** | Specifies the maximum number of entries that may be returned for a search performed by a client associated with this Client Connection Policy. | [optional] 
**MaximumSearchTimeLimit** | Pointer to **string** | Specifies the maximum length of time that the server should spend processing search operations requested by clients associated with this Client Connection Policy. | [optional] 
**MaximumSearchLookthroughLimit** | Pointer to **int32** | Specifies the maximum number of entries that may be examined by a backend in the course of processing a search requested by clients associated with this Client Connection Policy. | [optional] 
**MaximumLDAPJoinSizeLimit** | Pointer to **int32** | Specifies the maximum number of entries that may be joined with any single search result entry for a search request performed by a client associated with this Client Connection Policy. | [optional] 
**MaximumSortSizeLimitWithoutVLVIndex** | Pointer to **int32** | Specifies the maximum number of entries that the server will attempt to sort without the benefit of a VLV index. A value of zero indicates that no limit should be enforced. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewClientConnectionPolicyResponse

`func NewClientConnectionPolicyResponse(id string, policyID string, enabled bool, evaluationOrderIndex int32, allowedOperation []EnumclientConnectionPolicyAllowedOperationProp, allowedAuthType []EnumclientConnectionPolicyAllowedAuthTypeProp, ) *ClientConnectionPolicyResponse`

NewClientConnectionPolicyResponse instantiates a new ClientConnectionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConnectionPolicyResponseWithDefaults

`func NewClientConnectionPolicyResponseWithDefaults() *ClientConnectionPolicyResponse`

NewClientConnectionPolicyResponseWithDefaults instantiates a new ClientConnectionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientConnectionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientConnectionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientConnectionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ClientConnectionPolicyResponse) GetSchemas() []EnumclientConnectionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ClientConnectionPolicyResponse) GetSchemasOk() (*[]EnumclientConnectionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ClientConnectionPolicyResponse) SetSchemas(v []EnumclientConnectionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ClientConnectionPolicyResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetPolicyID

`func (o *ClientConnectionPolicyResponse) GetPolicyID() string`

GetPolicyID returns the PolicyID field if non-nil, zero value otherwise.

### GetPolicyIDOk

`func (o *ClientConnectionPolicyResponse) GetPolicyIDOk() (*string, bool)`

GetPolicyIDOk returns a tuple with the PolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyID

`func (o *ClientConnectionPolicyResponse) SetPolicyID(v string)`

SetPolicyID sets PolicyID field to given value.


### GetDescription

`func (o *ClientConnectionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientConnectionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientConnectionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientConnectionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ClientConnectionPolicyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClientConnectionPolicyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClientConnectionPolicyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEvaluationOrderIndex

`func (o *ClientConnectionPolicyResponse) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *ClientConnectionPolicyResponse) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *ClientConnectionPolicyResponse) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetConnectionCriteria

`func (o *ClientConnectionPolicyResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ClientConnectionPolicyResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ClientConnectionPolicyResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ClientConnectionPolicyResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetTerminateConnection

`func (o *ClientConnectionPolicyResponse) GetTerminateConnection() bool`

GetTerminateConnection returns the TerminateConnection field if non-nil, zero value otherwise.

### GetTerminateConnectionOk

`func (o *ClientConnectionPolicyResponse) GetTerminateConnectionOk() (*bool, bool)`

GetTerminateConnectionOk returns a tuple with the TerminateConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateConnection

`func (o *ClientConnectionPolicyResponse) SetTerminateConnection(v bool)`

SetTerminateConnection sets TerminateConnection field to given value.

### HasTerminateConnection

`func (o *ClientConnectionPolicyResponse) HasTerminateConnection() bool`

HasTerminateConnection returns a boolean if a field has been set.

### GetSensitiveAttribute

`func (o *ClientConnectionPolicyResponse) GetSensitiveAttribute() []string`

GetSensitiveAttribute returns the SensitiveAttribute field if non-nil, zero value otherwise.

### GetSensitiveAttributeOk

`func (o *ClientConnectionPolicyResponse) GetSensitiveAttributeOk() (*[]string, bool)`

GetSensitiveAttributeOk returns a tuple with the SensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveAttribute

`func (o *ClientConnectionPolicyResponse) SetSensitiveAttribute(v []string)`

SetSensitiveAttribute sets SensitiveAttribute field to given value.

### HasSensitiveAttribute

`func (o *ClientConnectionPolicyResponse) HasSensitiveAttribute() bool`

HasSensitiveAttribute returns a boolean if a field has been set.

### GetExcludeGlobalSensitiveAttribute

`func (o *ClientConnectionPolicyResponse) GetExcludeGlobalSensitiveAttribute() []string`

GetExcludeGlobalSensitiveAttribute returns the ExcludeGlobalSensitiveAttribute field if non-nil, zero value otherwise.

### GetExcludeGlobalSensitiveAttributeOk

`func (o *ClientConnectionPolicyResponse) GetExcludeGlobalSensitiveAttributeOk() (*[]string, bool)`

GetExcludeGlobalSensitiveAttributeOk returns a tuple with the ExcludeGlobalSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGlobalSensitiveAttribute

`func (o *ClientConnectionPolicyResponse) SetExcludeGlobalSensitiveAttribute(v []string)`

SetExcludeGlobalSensitiveAttribute sets ExcludeGlobalSensitiveAttribute field to given value.

### HasExcludeGlobalSensitiveAttribute

`func (o *ClientConnectionPolicyResponse) HasExcludeGlobalSensitiveAttribute() bool`

HasExcludeGlobalSensitiveAttribute returns a boolean if a field has been set.

### GetResultCodeMap

`func (o *ClientConnectionPolicyResponse) GetResultCodeMap() string`

GetResultCodeMap returns the ResultCodeMap field if non-nil, zero value otherwise.

### GetResultCodeMapOk

`func (o *ClientConnectionPolicyResponse) GetResultCodeMapOk() (*string, bool)`

GetResultCodeMapOk returns a tuple with the ResultCodeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeMap

`func (o *ClientConnectionPolicyResponse) SetResultCodeMap(v string)`

SetResultCodeMap sets ResultCodeMap field to given value.

### HasResultCodeMap

`func (o *ClientConnectionPolicyResponse) HasResultCodeMap() bool`

HasResultCodeMap returns a boolean if a field has been set.

### GetIncludedBackendBaseDN

`func (o *ClientConnectionPolicyResponse) GetIncludedBackendBaseDN() []string`

GetIncludedBackendBaseDN returns the IncludedBackendBaseDN field if non-nil, zero value otherwise.

### GetIncludedBackendBaseDNOk

`func (o *ClientConnectionPolicyResponse) GetIncludedBackendBaseDNOk() (*[]string, bool)`

GetIncludedBackendBaseDNOk returns a tuple with the IncludedBackendBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendBaseDN

`func (o *ClientConnectionPolicyResponse) SetIncludedBackendBaseDN(v []string)`

SetIncludedBackendBaseDN sets IncludedBackendBaseDN field to given value.

### HasIncludedBackendBaseDN

`func (o *ClientConnectionPolicyResponse) HasIncludedBackendBaseDN() bool`

HasIncludedBackendBaseDN returns a boolean if a field has been set.

### GetExcludedBackendBaseDN

`func (o *ClientConnectionPolicyResponse) GetExcludedBackendBaseDN() []string`

GetExcludedBackendBaseDN returns the ExcludedBackendBaseDN field if non-nil, zero value otherwise.

### GetExcludedBackendBaseDNOk

`func (o *ClientConnectionPolicyResponse) GetExcludedBackendBaseDNOk() (*[]string, bool)`

GetExcludedBackendBaseDNOk returns a tuple with the ExcludedBackendBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendBaseDN

`func (o *ClientConnectionPolicyResponse) SetExcludedBackendBaseDN(v []string)`

SetExcludedBackendBaseDN sets ExcludedBackendBaseDN field to given value.

### HasExcludedBackendBaseDN

`func (o *ClientConnectionPolicyResponse) HasExcludedBackendBaseDN() bool`

HasExcludedBackendBaseDN returns a boolean if a field has been set.

### GetAllowedOperation

`func (o *ClientConnectionPolicyResponse) GetAllowedOperation() []EnumclientConnectionPolicyAllowedOperationProp`

GetAllowedOperation returns the AllowedOperation field if non-nil, zero value otherwise.

### GetAllowedOperationOk

`func (o *ClientConnectionPolicyResponse) GetAllowedOperationOk() (*[]EnumclientConnectionPolicyAllowedOperationProp, bool)`

GetAllowedOperationOk returns a tuple with the AllowedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperation

`func (o *ClientConnectionPolicyResponse) SetAllowedOperation(v []EnumclientConnectionPolicyAllowedOperationProp)`

SetAllowedOperation sets AllowedOperation field to given value.


### GetRequiredOperationRequestCriteria

`func (o *ClientConnectionPolicyResponse) GetRequiredOperationRequestCriteria() string`

GetRequiredOperationRequestCriteria returns the RequiredOperationRequestCriteria field if non-nil, zero value otherwise.

### GetRequiredOperationRequestCriteriaOk

`func (o *ClientConnectionPolicyResponse) GetRequiredOperationRequestCriteriaOk() (*string, bool)`

GetRequiredOperationRequestCriteriaOk returns a tuple with the RequiredOperationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOperationRequestCriteria

`func (o *ClientConnectionPolicyResponse) SetRequiredOperationRequestCriteria(v string)`

SetRequiredOperationRequestCriteria sets RequiredOperationRequestCriteria field to given value.

### HasRequiredOperationRequestCriteria

`func (o *ClientConnectionPolicyResponse) HasRequiredOperationRequestCriteria() bool`

HasRequiredOperationRequestCriteria returns a boolean if a field has been set.

### GetProhibitedOperationRequestCriteria

`func (o *ClientConnectionPolicyResponse) GetProhibitedOperationRequestCriteria() string`

GetProhibitedOperationRequestCriteria returns the ProhibitedOperationRequestCriteria field if non-nil, zero value otherwise.

### GetProhibitedOperationRequestCriteriaOk

`func (o *ClientConnectionPolicyResponse) GetProhibitedOperationRequestCriteriaOk() (*string, bool)`

GetProhibitedOperationRequestCriteriaOk returns a tuple with the ProhibitedOperationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProhibitedOperationRequestCriteria

`func (o *ClientConnectionPolicyResponse) SetProhibitedOperationRequestCriteria(v string)`

SetProhibitedOperationRequestCriteria sets ProhibitedOperationRequestCriteria field to given value.

### HasProhibitedOperationRequestCriteria

`func (o *ClientConnectionPolicyResponse) HasProhibitedOperationRequestCriteria() bool`

HasProhibitedOperationRequestCriteria returns a boolean if a field has been set.

### GetAllowedRequestControl

`func (o *ClientConnectionPolicyResponse) GetAllowedRequestControl() []string`

GetAllowedRequestControl returns the AllowedRequestControl field if non-nil, zero value otherwise.

### GetAllowedRequestControlOk

`func (o *ClientConnectionPolicyResponse) GetAllowedRequestControlOk() (*[]string, bool)`

GetAllowedRequestControlOk returns a tuple with the AllowedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestControl

`func (o *ClientConnectionPolicyResponse) SetAllowedRequestControl(v []string)`

SetAllowedRequestControl sets AllowedRequestControl field to given value.

### HasAllowedRequestControl

`func (o *ClientConnectionPolicyResponse) HasAllowedRequestControl() bool`

HasAllowedRequestControl returns a boolean if a field has been set.

### GetDeniedRequestControl

`func (o *ClientConnectionPolicyResponse) GetDeniedRequestControl() []string`

GetDeniedRequestControl returns the DeniedRequestControl field if non-nil, zero value otherwise.

### GetDeniedRequestControlOk

`func (o *ClientConnectionPolicyResponse) GetDeniedRequestControlOk() (*[]string, bool)`

GetDeniedRequestControlOk returns a tuple with the DeniedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedRequestControl

`func (o *ClientConnectionPolicyResponse) SetDeniedRequestControl(v []string)`

SetDeniedRequestControl sets DeniedRequestControl field to given value.

### HasDeniedRequestControl

`func (o *ClientConnectionPolicyResponse) HasDeniedRequestControl() bool`

HasDeniedRequestControl returns a boolean if a field has been set.

### GetAllowedExtendedOperation

`func (o *ClientConnectionPolicyResponse) GetAllowedExtendedOperation() []string`

GetAllowedExtendedOperation returns the AllowedExtendedOperation field if non-nil, zero value otherwise.

### GetAllowedExtendedOperationOk

`func (o *ClientConnectionPolicyResponse) GetAllowedExtendedOperationOk() (*[]string, bool)`

GetAllowedExtendedOperationOk returns a tuple with the AllowedExtendedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtendedOperation

`func (o *ClientConnectionPolicyResponse) SetAllowedExtendedOperation(v []string)`

SetAllowedExtendedOperation sets AllowedExtendedOperation field to given value.

### HasAllowedExtendedOperation

`func (o *ClientConnectionPolicyResponse) HasAllowedExtendedOperation() bool`

HasAllowedExtendedOperation returns a boolean if a field has been set.

### GetDeniedExtendedOperation

`func (o *ClientConnectionPolicyResponse) GetDeniedExtendedOperation() []string`

GetDeniedExtendedOperation returns the DeniedExtendedOperation field if non-nil, zero value otherwise.

### GetDeniedExtendedOperationOk

`func (o *ClientConnectionPolicyResponse) GetDeniedExtendedOperationOk() (*[]string, bool)`

GetDeniedExtendedOperationOk returns a tuple with the DeniedExtendedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedExtendedOperation

`func (o *ClientConnectionPolicyResponse) SetDeniedExtendedOperation(v []string)`

SetDeniedExtendedOperation sets DeniedExtendedOperation field to given value.

### HasDeniedExtendedOperation

`func (o *ClientConnectionPolicyResponse) HasDeniedExtendedOperation() bool`

HasDeniedExtendedOperation returns a boolean if a field has been set.

### GetAllowedAuthType

`func (o *ClientConnectionPolicyResponse) GetAllowedAuthType() []EnumclientConnectionPolicyAllowedAuthTypeProp`

GetAllowedAuthType returns the AllowedAuthType field if non-nil, zero value otherwise.

### GetAllowedAuthTypeOk

`func (o *ClientConnectionPolicyResponse) GetAllowedAuthTypeOk() (*[]EnumclientConnectionPolicyAllowedAuthTypeProp, bool)`

GetAllowedAuthTypeOk returns a tuple with the AllowedAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthType

`func (o *ClientConnectionPolicyResponse) SetAllowedAuthType(v []EnumclientConnectionPolicyAllowedAuthTypeProp)`

SetAllowedAuthType sets AllowedAuthType field to given value.


### GetAllowedSASLMechanism

`func (o *ClientConnectionPolicyResponse) GetAllowedSASLMechanism() []string`

GetAllowedSASLMechanism returns the AllowedSASLMechanism field if non-nil, zero value otherwise.

### GetAllowedSASLMechanismOk

`func (o *ClientConnectionPolicyResponse) GetAllowedSASLMechanismOk() (*[]string, bool)`

GetAllowedSASLMechanismOk returns a tuple with the AllowedSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSASLMechanism

`func (o *ClientConnectionPolicyResponse) SetAllowedSASLMechanism(v []string)`

SetAllowedSASLMechanism sets AllowedSASLMechanism field to given value.

### HasAllowedSASLMechanism

`func (o *ClientConnectionPolicyResponse) HasAllowedSASLMechanism() bool`

HasAllowedSASLMechanism returns a boolean if a field has been set.

### GetDeniedSASLMechanism

`func (o *ClientConnectionPolicyResponse) GetDeniedSASLMechanism() []string`

GetDeniedSASLMechanism returns the DeniedSASLMechanism field if non-nil, zero value otherwise.

### GetDeniedSASLMechanismOk

`func (o *ClientConnectionPolicyResponse) GetDeniedSASLMechanismOk() (*[]string, bool)`

GetDeniedSASLMechanismOk returns a tuple with the DeniedSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedSASLMechanism

`func (o *ClientConnectionPolicyResponse) SetDeniedSASLMechanism(v []string)`

SetDeniedSASLMechanism sets DeniedSASLMechanism field to given value.

### HasDeniedSASLMechanism

`func (o *ClientConnectionPolicyResponse) HasDeniedSASLMechanism() bool`

HasDeniedSASLMechanism returns a boolean if a field has been set.

### GetAllowedFilterType

`func (o *ClientConnectionPolicyResponse) GetAllowedFilterType() []EnumclientConnectionPolicyAllowedFilterTypeProp`

GetAllowedFilterType returns the AllowedFilterType field if non-nil, zero value otherwise.

### GetAllowedFilterTypeOk

`func (o *ClientConnectionPolicyResponse) GetAllowedFilterTypeOk() (*[]EnumclientConnectionPolicyAllowedFilterTypeProp, bool)`

GetAllowedFilterTypeOk returns a tuple with the AllowedFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFilterType

`func (o *ClientConnectionPolicyResponse) SetAllowedFilterType(v []EnumclientConnectionPolicyAllowedFilterTypeProp)`

SetAllowedFilterType sets AllowedFilterType field to given value.

### HasAllowedFilterType

`func (o *ClientConnectionPolicyResponse) HasAllowedFilterType() bool`

HasAllowedFilterType returns a boolean if a field has been set.

### GetAllowUnindexedSearches

`func (o *ClientConnectionPolicyResponse) GetAllowUnindexedSearches() bool`

GetAllowUnindexedSearches returns the AllowUnindexedSearches field if non-nil, zero value otherwise.

### GetAllowUnindexedSearchesOk

`func (o *ClientConnectionPolicyResponse) GetAllowUnindexedSearchesOk() (*bool, bool)`

GetAllowUnindexedSearchesOk returns a tuple with the AllowUnindexedSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnindexedSearches

`func (o *ClientConnectionPolicyResponse) SetAllowUnindexedSearches(v bool)`

SetAllowUnindexedSearches sets AllowUnindexedSearches field to given value.

### HasAllowUnindexedSearches

`func (o *ClientConnectionPolicyResponse) HasAllowUnindexedSearches() bool`

HasAllowUnindexedSearches returns a boolean if a field has been set.

### GetAllowUnindexedSearchesWithControl

`func (o *ClientConnectionPolicyResponse) GetAllowUnindexedSearchesWithControl() bool`

GetAllowUnindexedSearchesWithControl returns the AllowUnindexedSearchesWithControl field if non-nil, zero value otherwise.

### GetAllowUnindexedSearchesWithControlOk

`func (o *ClientConnectionPolicyResponse) GetAllowUnindexedSearchesWithControlOk() (*bool, bool)`

GetAllowUnindexedSearchesWithControlOk returns a tuple with the AllowUnindexedSearchesWithControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnindexedSearchesWithControl

`func (o *ClientConnectionPolicyResponse) SetAllowUnindexedSearchesWithControl(v bool)`

SetAllowUnindexedSearchesWithControl sets AllowUnindexedSearchesWithControl field to given value.

### HasAllowUnindexedSearchesWithControl

`func (o *ClientConnectionPolicyResponse) HasAllowUnindexedSearchesWithControl() bool`

HasAllowUnindexedSearchesWithControl returns a boolean if a field has been set.

### GetMinimumSubstringLength

`func (o *ClientConnectionPolicyResponse) GetMinimumSubstringLength() int32`

GetMinimumSubstringLength returns the MinimumSubstringLength field if non-nil, zero value otherwise.

### GetMinimumSubstringLengthOk

`func (o *ClientConnectionPolicyResponse) GetMinimumSubstringLengthOk() (*int32, bool)`

GetMinimumSubstringLengthOk returns a tuple with the MinimumSubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSubstringLength

`func (o *ClientConnectionPolicyResponse) SetMinimumSubstringLength(v int32)`

SetMinimumSubstringLength sets MinimumSubstringLength field to given value.

### HasMinimumSubstringLength

`func (o *ClientConnectionPolicyResponse) HasMinimumSubstringLength() bool`

HasMinimumSubstringLength returns a boolean if a field has been set.

### GetMaximumConcurrentConnections

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentConnections() int32`

GetMaximumConcurrentConnections returns the MaximumConcurrentConnections field if non-nil, zero value otherwise.

### GetMaximumConcurrentConnectionsOk

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentConnectionsOk() (*int32, bool)`

GetMaximumConcurrentConnectionsOk returns a tuple with the MaximumConcurrentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentConnections

`func (o *ClientConnectionPolicyResponse) SetMaximumConcurrentConnections(v int32)`

SetMaximumConcurrentConnections sets MaximumConcurrentConnections field to given value.

### HasMaximumConcurrentConnections

`func (o *ClientConnectionPolicyResponse) HasMaximumConcurrentConnections() bool`

HasMaximumConcurrentConnections returns a boolean if a field has been set.

### GetMaximumConnectionDuration

`func (o *ClientConnectionPolicyResponse) GetMaximumConnectionDuration() string`

GetMaximumConnectionDuration returns the MaximumConnectionDuration field if non-nil, zero value otherwise.

### GetMaximumConnectionDurationOk

`func (o *ClientConnectionPolicyResponse) GetMaximumConnectionDurationOk() (*string, bool)`

GetMaximumConnectionDurationOk returns a tuple with the MaximumConnectionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConnectionDuration

`func (o *ClientConnectionPolicyResponse) SetMaximumConnectionDuration(v string)`

SetMaximumConnectionDuration sets MaximumConnectionDuration field to given value.

### HasMaximumConnectionDuration

`func (o *ClientConnectionPolicyResponse) HasMaximumConnectionDuration() bool`

HasMaximumConnectionDuration returns a boolean if a field has been set.

### GetMaximumIdleConnectionDuration

`func (o *ClientConnectionPolicyResponse) GetMaximumIdleConnectionDuration() string`

GetMaximumIdleConnectionDuration returns the MaximumIdleConnectionDuration field if non-nil, zero value otherwise.

### GetMaximumIdleConnectionDurationOk

`func (o *ClientConnectionPolicyResponse) GetMaximumIdleConnectionDurationOk() (*string, bool)`

GetMaximumIdleConnectionDurationOk returns a tuple with the MaximumIdleConnectionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIdleConnectionDuration

`func (o *ClientConnectionPolicyResponse) SetMaximumIdleConnectionDuration(v string)`

SetMaximumIdleConnectionDuration sets MaximumIdleConnectionDuration field to given value.

### HasMaximumIdleConnectionDuration

`func (o *ClientConnectionPolicyResponse) HasMaximumIdleConnectionDuration() bool`

HasMaximumIdleConnectionDuration returns a boolean if a field has been set.

### GetMaximumOperationCountPerConnection

`func (o *ClientConnectionPolicyResponse) GetMaximumOperationCountPerConnection() int32`

GetMaximumOperationCountPerConnection returns the MaximumOperationCountPerConnection field if non-nil, zero value otherwise.

### GetMaximumOperationCountPerConnectionOk

`func (o *ClientConnectionPolicyResponse) GetMaximumOperationCountPerConnectionOk() (*int32, bool)`

GetMaximumOperationCountPerConnectionOk returns a tuple with the MaximumOperationCountPerConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOperationCountPerConnection

`func (o *ClientConnectionPolicyResponse) SetMaximumOperationCountPerConnection(v int32)`

SetMaximumOperationCountPerConnection sets MaximumOperationCountPerConnection field to given value.

### HasMaximumOperationCountPerConnection

`func (o *ClientConnectionPolicyResponse) HasMaximumOperationCountPerConnection() bool`

HasMaximumOperationCountPerConnection returns a boolean if a field has been set.

### GetMaximumConcurrentOperationsPerConnection

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentOperationsPerConnection() int32`

GetMaximumConcurrentOperationsPerConnection returns the MaximumConcurrentOperationsPerConnection field if non-nil, zero value otherwise.

### GetMaximumConcurrentOperationsPerConnectionOk

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentOperationsPerConnectionOk() (*int32, bool)`

GetMaximumConcurrentOperationsPerConnectionOk returns a tuple with the MaximumConcurrentOperationsPerConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentOperationsPerConnection

`func (o *ClientConnectionPolicyResponse) SetMaximumConcurrentOperationsPerConnection(v int32)`

SetMaximumConcurrentOperationsPerConnection sets MaximumConcurrentOperationsPerConnection field to given value.

### HasMaximumConcurrentOperationsPerConnection

`func (o *ClientConnectionPolicyResponse) HasMaximumConcurrentOperationsPerConnection() bool`

HasMaximumConcurrentOperationsPerConnection returns a boolean if a field has been set.

### GetMaximumConcurrentOperationWaitTimeBeforeRejecting

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentOperationWaitTimeBeforeRejecting() string`

GetMaximumConcurrentOperationWaitTimeBeforeRejecting returns the MaximumConcurrentOperationWaitTimeBeforeRejecting field if non-nil, zero value otherwise.

### GetMaximumConcurrentOperationWaitTimeBeforeRejectingOk

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentOperationWaitTimeBeforeRejectingOk() (*string, bool)`

GetMaximumConcurrentOperationWaitTimeBeforeRejectingOk returns a tuple with the MaximumConcurrentOperationWaitTimeBeforeRejecting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentOperationWaitTimeBeforeRejecting

`func (o *ClientConnectionPolicyResponse) SetMaximumConcurrentOperationWaitTimeBeforeRejecting(v string)`

SetMaximumConcurrentOperationWaitTimeBeforeRejecting sets MaximumConcurrentOperationWaitTimeBeforeRejecting field to given value.

### HasMaximumConcurrentOperationWaitTimeBeforeRejecting

`func (o *ClientConnectionPolicyResponse) HasMaximumConcurrentOperationWaitTimeBeforeRejecting() bool`

HasMaximumConcurrentOperationWaitTimeBeforeRejecting returns a boolean if a field has been set.

### GetMaximumConcurrentOperationsPerConnectionExceededBehavior

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentOperationsPerConnectionExceededBehavior() EnumclientConnectionPolicyMaximumConcurrentOperationsPerConnectionExceededBehaviorProp`

GetMaximumConcurrentOperationsPerConnectionExceededBehavior returns the MaximumConcurrentOperationsPerConnectionExceededBehavior field if non-nil, zero value otherwise.

### GetMaximumConcurrentOperationsPerConnectionExceededBehaviorOk

`func (o *ClientConnectionPolicyResponse) GetMaximumConcurrentOperationsPerConnectionExceededBehaviorOk() (*EnumclientConnectionPolicyMaximumConcurrentOperationsPerConnectionExceededBehaviorProp, bool)`

GetMaximumConcurrentOperationsPerConnectionExceededBehaviorOk returns a tuple with the MaximumConcurrentOperationsPerConnectionExceededBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentOperationsPerConnectionExceededBehavior

`func (o *ClientConnectionPolicyResponse) SetMaximumConcurrentOperationsPerConnectionExceededBehavior(v EnumclientConnectionPolicyMaximumConcurrentOperationsPerConnectionExceededBehaviorProp)`

SetMaximumConcurrentOperationsPerConnectionExceededBehavior sets MaximumConcurrentOperationsPerConnectionExceededBehavior field to given value.

### HasMaximumConcurrentOperationsPerConnectionExceededBehavior

`func (o *ClientConnectionPolicyResponse) HasMaximumConcurrentOperationsPerConnectionExceededBehavior() bool`

HasMaximumConcurrentOperationsPerConnectionExceededBehavior returns a boolean if a field has been set.

### GetMaximumConnectionOperationRate

`func (o *ClientConnectionPolicyResponse) GetMaximumConnectionOperationRate() []string`

GetMaximumConnectionOperationRate returns the MaximumConnectionOperationRate field if non-nil, zero value otherwise.

### GetMaximumConnectionOperationRateOk

`func (o *ClientConnectionPolicyResponse) GetMaximumConnectionOperationRateOk() (*[]string, bool)`

GetMaximumConnectionOperationRateOk returns a tuple with the MaximumConnectionOperationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConnectionOperationRate

`func (o *ClientConnectionPolicyResponse) SetMaximumConnectionOperationRate(v []string)`

SetMaximumConnectionOperationRate sets MaximumConnectionOperationRate field to given value.

### HasMaximumConnectionOperationRate

`func (o *ClientConnectionPolicyResponse) HasMaximumConnectionOperationRate() bool`

HasMaximumConnectionOperationRate returns a boolean if a field has been set.

### GetConnectionOperationRateExceededBehavior

`func (o *ClientConnectionPolicyResponse) GetConnectionOperationRateExceededBehavior() EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp`

GetConnectionOperationRateExceededBehavior returns the ConnectionOperationRateExceededBehavior field if non-nil, zero value otherwise.

### GetConnectionOperationRateExceededBehaviorOk

`func (o *ClientConnectionPolicyResponse) GetConnectionOperationRateExceededBehaviorOk() (*EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp, bool)`

GetConnectionOperationRateExceededBehaviorOk returns a tuple with the ConnectionOperationRateExceededBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOperationRateExceededBehavior

`func (o *ClientConnectionPolicyResponse) SetConnectionOperationRateExceededBehavior(v EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp)`

SetConnectionOperationRateExceededBehavior sets ConnectionOperationRateExceededBehavior field to given value.

### HasConnectionOperationRateExceededBehavior

`func (o *ClientConnectionPolicyResponse) HasConnectionOperationRateExceededBehavior() bool`

HasConnectionOperationRateExceededBehavior returns a boolean if a field has been set.

### GetMaximumPolicyOperationRate

`func (o *ClientConnectionPolicyResponse) GetMaximumPolicyOperationRate() []string`

GetMaximumPolicyOperationRate returns the MaximumPolicyOperationRate field if non-nil, zero value otherwise.

### GetMaximumPolicyOperationRateOk

`func (o *ClientConnectionPolicyResponse) GetMaximumPolicyOperationRateOk() (*[]string, bool)`

GetMaximumPolicyOperationRateOk returns a tuple with the MaximumPolicyOperationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPolicyOperationRate

`func (o *ClientConnectionPolicyResponse) SetMaximumPolicyOperationRate(v []string)`

SetMaximumPolicyOperationRate sets MaximumPolicyOperationRate field to given value.

### HasMaximumPolicyOperationRate

`func (o *ClientConnectionPolicyResponse) HasMaximumPolicyOperationRate() bool`

HasMaximumPolicyOperationRate returns a boolean if a field has been set.

### GetPolicyOperationRateExceededBehavior

`func (o *ClientConnectionPolicyResponse) GetPolicyOperationRateExceededBehavior() EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp`

GetPolicyOperationRateExceededBehavior returns the PolicyOperationRateExceededBehavior field if non-nil, zero value otherwise.

### GetPolicyOperationRateExceededBehaviorOk

`func (o *ClientConnectionPolicyResponse) GetPolicyOperationRateExceededBehaviorOk() (*EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp, bool)`

GetPolicyOperationRateExceededBehaviorOk returns a tuple with the PolicyOperationRateExceededBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOperationRateExceededBehavior

`func (o *ClientConnectionPolicyResponse) SetPolicyOperationRateExceededBehavior(v EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp)`

SetPolicyOperationRateExceededBehavior sets PolicyOperationRateExceededBehavior field to given value.

### HasPolicyOperationRateExceededBehavior

`func (o *ClientConnectionPolicyResponse) HasPolicyOperationRateExceededBehavior() bool`

HasPolicyOperationRateExceededBehavior returns a boolean if a field has been set.

### GetMaximumSearchSizeLimit

`func (o *ClientConnectionPolicyResponse) GetMaximumSearchSizeLimit() int32`

GetMaximumSearchSizeLimit returns the MaximumSearchSizeLimit field if non-nil, zero value otherwise.

### GetMaximumSearchSizeLimitOk

`func (o *ClientConnectionPolicyResponse) GetMaximumSearchSizeLimitOk() (*int32, bool)`

GetMaximumSearchSizeLimitOk returns a tuple with the MaximumSearchSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSearchSizeLimit

`func (o *ClientConnectionPolicyResponse) SetMaximumSearchSizeLimit(v int32)`

SetMaximumSearchSizeLimit sets MaximumSearchSizeLimit field to given value.

### HasMaximumSearchSizeLimit

`func (o *ClientConnectionPolicyResponse) HasMaximumSearchSizeLimit() bool`

HasMaximumSearchSizeLimit returns a boolean if a field has been set.

### GetMaximumSearchTimeLimit

`func (o *ClientConnectionPolicyResponse) GetMaximumSearchTimeLimit() string`

GetMaximumSearchTimeLimit returns the MaximumSearchTimeLimit field if non-nil, zero value otherwise.

### GetMaximumSearchTimeLimitOk

`func (o *ClientConnectionPolicyResponse) GetMaximumSearchTimeLimitOk() (*string, bool)`

GetMaximumSearchTimeLimitOk returns a tuple with the MaximumSearchTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSearchTimeLimit

`func (o *ClientConnectionPolicyResponse) SetMaximumSearchTimeLimit(v string)`

SetMaximumSearchTimeLimit sets MaximumSearchTimeLimit field to given value.

### HasMaximumSearchTimeLimit

`func (o *ClientConnectionPolicyResponse) HasMaximumSearchTimeLimit() bool`

HasMaximumSearchTimeLimit returns a boolean if a field has been set.

### GetMaximumSearchLookthroughLimit

`func (o *ClientConnectionPolicyResponse) GetMaximumSearchLookthroughLimit() int32`

GetMaximumSearchLookthroughLimit returns the MaximumSearchLookthroughLimit field if non-nil, zero value otherwise.

### GetMaximumSearchLookthroughLimitOk

`func (o *ClientConnectionPolicyResponse) GetMaximumSearchLookthroughLimitOk() (*int32, bool)`

GetMaximumSearchLookthroughLimitOk returns a tuple with the MaximumSearchLookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSearchLookthroughLimit

`func (o *ClientConnectionPolicyResponse) SetMaximumSearchLookthroughLimit(v int32)`

SetMaximumSearchLookthroughLimit sets MaximumSearchLookthroughLimit field to given value.

### HasMaximumSearchLookthroughLimit

`func (o *ClientConnectionPolicyResponse) HasMaximumSearchLookthroughLimit() bool`

HasMaximumSearchLookthroughLimit returns a boolean if a field has been set.

### GetMaximumLDAPJoinSizeLimit

`func (o *ClientConnectionPolicyResponse) GetMaximumLDAPJoinSizeLimit() int32`

GetMaximumLDAPJoinSizeLimit returns the MaximumLDAPJoinSizeLimit field if non-nil, zero value otherwise.

### GetMaximumLDAPJoinSizeLimitOk

`func (o *ClientConnectionPolicyResponse) GetMaximumLDAPJoinSizeLimitOk() (*int32, bool)`

GetMaximumLDAPJoinSizeLimitOk returns a tuple with the MaximumLDAPJoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLDAPJoinSizeLimit

`func (o *ClientConnectionPolicyResponse) SetMaximumLDAPJoinSizeLimit(v int32)`

SetMaximumLDAPJoinSizeLimit sets MaximumLDAPJoinSizeLimit field to given value.

### HasMaximumLDAPJoinSizeLimit

`func (o *ClientConnectionPolicyResponse) HasMaximumLDAPJoinSizeLimit() bool`

HasMaximumLDAPJoinSizeLimit returns a boolean if a field has been set.

### GetMaximumSortSizeLimitWithoutVLVIndex

`func (o *ClientConnectionPolicyResponse) GetMaximumSortSizeLimitWithoutVLVIndex() int32`

GetMaximumSortSizeLimitWithoutVLVIndex returns the MaximumSortSizeLimitWithoutVLVIndex field if non-nil, zero value otherwise.

### GetMaximumSortSizeLimitWithoutVLVIndexOk

`func (o *ClientConnectionPolicyResponse) GetMaximumSortSizeLimitWithoutVLVIndexOk() (*int32, bool)`

GetMaximumSortSizeLimitWithoutVLVIndexOk returns a tuple with the MaximumSortSizeLimitWithoutVLVIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSortSizeLimitWithoutVLVIndex

`func (o *ClientConnectionPolicyResponse) SetMaximumSortSizeLimitWithoutVLVIndex(v int32)`

SetMaximumSortSizeLimitWithoutVLVIndex sets MaximumSortSizeLimitWithoutVLVIndex field to given value.

### HasMaximumSortSizeLimitWithoutVLVIndex

`func (o *ClientConnectionPolicyResponse) HasMaximumSortSizeLimitWithoutVLVIndex() bool`

HasMaximumSortSizeLimitWithoutVLVIndex returns a boolean if a field has been set.

### GetMeta

`func (o *ClientConnectionPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClientConnectionPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClientConnectionPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClientConnectionPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ClientConnectionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ClientConnectionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ClientConnectionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ClientConnectionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


