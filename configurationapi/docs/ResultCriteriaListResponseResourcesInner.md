# ResultCriteriaListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Result Criteria | 
**Schemas** | [**[]EnumthirdPartyResultCriteriaSchemaUrn**](EnumthirdPartyResultCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for operations included in this Successful Bind Result Criteria. | [optional] 
**ResultCodeCriteria** | Pointer to [**EnumresultCriteriaResultCodeCriteriaProp**](EnumresultCriteriaResultCodeCriteriaProp.md) |  | [optional] 
**ResultCodeValue** | Pointer to [**[]EnumresultCriteriaResultCodeValueProp**](EnumresultCriteriaResultCodeValueProp.md) |  | [optional] 
**ProcessingTimeCriteria** | Pointer to [**EnumresultCriteriaProcessingTimeCriteriaProp**](EnumresultCriteriaProcessingTimeCriteriaProp.md) |  | [optional] 
**ProcessingTimeValue** | Pointer to **string** | Specifies the boundary value to use for the operation processing time when determining whether to include that operation in this Simple Result Criteria. This will be ignored if the \&quot;processing-time-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**QueueTimeCriteria** | Pointer to [**EnumresultCriteriaQueueTimeCriteriaProp**](EnumresultCriteriaQueueTimeCriteriaProp.md) |  | [optional] 
**QueueTimeValue** | Pointer to **string** | Specifies the boundary value to use for the time an operation spent on the work queue when determining whether to include that operation in this Simple Result Criteria. This will be ignored if the \&quot;queue-time-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**ReferralReturned** | Pointer to [**EnumresultCriteriaReferralReturnedProp**](EnumresultCriteriaReferralReturnedProp.md) |  | [optional] 
**AllIncludedResponseControl** | Pointer to **[]string** | Specifies the OID of a control that must be present in the response to the client for operations included in this Simple Result Criteria. If any control OIDs are provided, then the response must contain all of those controls. | [optional] 
**AnyIncludedResponseControl** | Pointer to **[]string** | Specifies the OID of a control that may be present in the response to the client for operations included in this Simple Result Criteria. If any control OIDs are provided, then the response must contain at least one of those controls. | [optional] 
**NotAllIncludedResponseControl** | Pointer to **[]string** | Specifies the OID of a control that should not be present in the response to the client for operations included in this Simple Result Criteria. If any control OIDs are provided, then the response must not contain at least one of those controls (that is, the response may contain zero or more of those controls, but not all of them). | [optional] 
**NoneIncludedResponseControl** | Pointer to **[]string** | Specifies the OID of a control that must not be present in the response to the client for operations included in this Simple Result Criteria. If any control OIDs are provided, then the response must not contain any of those controls. | [optional] 
**UsedAlternateAuthzid** | Pointer to [**EnumresultCriteriaUsedAlternateAuthzidProp**](EnumresultCriteriaUsedAlternateAuthzidProp.md) |  | [optional] 
**UsedAnyPrivilege** | Pointer to [**EnumresultCriteriaUsedAnyPrivilegeProp**](EnumresultCriteriaUsedAnyPrivilegeProp.md) |  | [optional] 
**UsedPrivilege** | Pointer to [**[]EnumresultCriteriaUsedPrivilegeProp**](EnumresultCriteriaUsedPrivilegeProp.md) |  | [optional] 
**MissingAnyPrivilege** | Pointer to [**EnumresultCriteriaMissingAnyPrivilegeProp**](EnumresultCriteriaMissingAnyPrivilegeProp.md) |  | [optional] 
**MissingPrivilege** | Pointer to [**[]EnumresultCriteriaMissingPrivilegeProp**](EnumresultCriteriaMissingPrivilegeProp.md) |  | [optional] 
**RetiredPasswordUsedForBind** | Pointer to [**EnumresultCriteriaRetiredPasswordUsedForBindProp**](EnumresultCriteriaRetiredPasswordUsedForBindProp.md) |  | [optional] 
**SearchEntryReturnedCriteria** | Pointer to [**EnumresultCriteriaSearchEntryReturnedCriteriaProp**](EnumresultCriteriaSearchEntryReturnedCriteriaProp.md) |  | [optional] 
**SearchEntryReturnedCount** | Pointer to **int64** | Specifies the target number of entries returned for use when determining whether a search operation should be included in this Simple Result Criteria. This will be ignored for all operations other than search, and it will be ignored for search operations if the \&quot;search-entry-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**SearchReferenceReturnedCriteria** | Pointer to [**EnumresultCriteriaSearchReferenceReturnedCriteriaProp**](EnumresultCriteriaSearchReferenceReturnedCriteriaProp.md) |  | [optional] 
**SearchReferenceReturnedCount** | Pointer to **int64** | Specifies the target number of references returned for use when determining whether a search operation should be included in this Simple Result Criteria. This will be ignored for all operations other than search, and it will be ignored for search operations if the \&quot;search-reference-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**SearchIndexedCriteria** | Pointer to [**EnumresultCriteriaSearchIndexedCriteriaProp**](EnumresultCriteriaSearchIndexedCriteriaProp.md) |  | [optional] 
**IncludedAuthzUserBaseDN** | Pointer to **[]string** | Specifies a base DN below which authorization user entries may exist for operations included in this Simple Result Criteria. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**ExcludedAuthzUserBaseDN** | Pointer to **[]string** | Specifies a base DN below which authorization user entries may exist for operations excluded from this Simple Result Criteria. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**AllIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users must exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must be a member of all of those groups. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**AnyIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users may exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must be a member of at least one of those groups. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**NotAllIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users should not exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must not be a member of at least one of those groups (that is, the user may be a member of zero or more of those groups, but not of all of them). The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**NoneIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users must not exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must not be a member any of those groups. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**IncludeAnonymousBinds** | Pointer to **bool** | Indicates whether this criteria will be permitted to match bind operations that resulted in anonymous authentication. | [optional] 
**IncludedUserBaseDN** | Pointer to **[]string** | A set of base DNs for authenticated users that will be permitted to match this criteria. | [optional] 
**ExcludedUserBaseDN** | Pointer to **[]string** | A set of base DNs for authenticated users that will not be permitted to match this criteria. | [optional] 
**IncludedUserFilter** | Pointer to **[]string** | A set of filters that may be used to identify entries for authenticated users that will be permitted to match this criteria. | [optional] 
**ExcludedUserFilter** | Pointer to **[]string** | A set of filters that may be used to identify entries for authenticated users that will not be permitted to match this criteria. | [optional] 
**IncludedUserGroupDN** | Pointer to **[]string** | The DNs of the groups whose members will be permitted to match this criteria. | [optional] 
**ExcludedUserGroupDN** | Pointer to **[]string** | The DNs of the groups whose members will not be permitted to match this criteria. | [optional] 
**AllIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that must match the associated operation result in order to match the aggregate result criteria. If one or more all-included result criteria objects are provided, then an operation result must match all of them in order to match the aggregate result criteria. | [optional] 
**AnyIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that may match the associated operation result in order to match the aggregate result criteria. If one or more any-included result criteria objects are provided, then an operation result must match at least one of them in order to match the aggregate result criteria. | [optional] 
**NotAllIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that should not match the associated operation result in order to match the aggregate result criteria. If one or more not-all-included result criteria objects are provided, then an operation result must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate result criteria. | [optional] 
**NoneIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that must not match the associated operation result in order to match the aggregate result criteria. If one or more none-included result criteria objects are provided, then an operation result must not match any of them in order to match the aggregate result criteria. | [optional] 
**LocalAssuranceLevel** | [**[]EnumresultCriteriaLocalAssuranceLevelProp**](EnumresultCriteriaLocalAssuranceLevelProp.md) |  | 
**RemoteAssuranceLevel** | [**[]EnumresultCriteriaRemoteAssuranceLevelProp**](EnumresultCriteriaRemoteAssuranceLevelProp.md) |  | 
**AssuranceTimeoutCriteria** | Pointer to [**EnumresultCriteriaAssuranceTimeoutCriteriaProp**](EnumresultCriteriaAssuranceTimeoutCriteriaProp.md) |  | [optional] 
**AssuranceTimeoutValue** | Pointer to **string** | The value to use for performing matching based on the assurance timeout. This will be ignored if the assurance-timeout-criteria is \&quot;any\&quot;. | [optional] 
**ResponseDelayedByAssurance** | Pointer to [**EnumresultCriteriaResponseDelayedByAssuranceProp**](EnumresultCriteriaResponseDelayedByAssuranceProp.md) |  | [optional] 
**AssuranceBehaviorAlteredByControl** | Pointer to [**EnumresultCriteriaAssuranceBehaviorAlteredByControlProp**](EnumresultCriteriaAssuranceBehaviorAlteredByControlProp.md) |  | [optional] 
**AssuranceSatisfied** | Pointer to [**EnumresultCriteriaAssuranceSatisfiedProp**](EnumresultCriteriaAssuranceSatisfiedProp.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Result Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Result Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewResultCriteriaListResponseResourcesInner

`func NewResultCriteriaListResponseResourcesInner(id string, schemas []EnumthirdPartyResultCriteriaSchemaUrn, localAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp, remoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp, extensionClass string, ) *ResultCriteriaListResponseResourcesInner`

NewResultCriteriaListResponseResourcesInner instantiates a new ResultCriteriaListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCriteriaListResponseResourcesInnerWithDefaults

`func NewResultCriteriaListResponseResourcesInnerWithDefaults() *ResultCriteriaListResponseResourcesInner`

NewResultCriteriaListResponseResourcesInnerWithDefaults instantiates a new ResultCriteriaListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultCriteriaListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultCriteriaListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultCriteriaListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ResultCriteriaListResponseResourcesInner) GetSchemas() []EnumthirdPartyResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ResultCriteriaListResponseResourcesInner) GetSchemasOk() (*[]EnumthirdPartyResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ResultCriteriaListResponseResourcesInner) SetSchemas(v []EnumthirdPartyResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetResultCodeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetResultCodeCriteria() EnumresultCriteriaResultCodeCriteriaProp`

GetResultCodeCriteria returns the ResultCodeCriteria field if non-nil, zero value otherwise.

### GetResultCodeCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetResultCodeCriteriaOk() (*EnumresultCriteriaResultCodeCriteriaProp, bool)`

GetResultCodeCriteriaOk returns a tuple with the ResultCodeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetResultCodeCriteria(v EnumresultCriteriaResultCodeCriteriaProp)`

SetResultCodeCriteria sets ResultCodeCriteria field to given value.

### HasResultCodeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasResultCodeCriteria() bool`

HasResultCodeCriteria returns a boolean if a field has been set.

### GetResultCodeValue

`func (o *ResultCriteriaListResponseResourcesInner) GetResultCodeValue() []EnumresultCriteriaResultCodeValueProp`

GetResultCodeValue returns the ResultCodeValue field if non-nil, zero value otherwise.

### GetResultCodeValueOk

`func (o *ResultCriteriaListResponseResourcesInner) GetResultCodeValueOk() (*[]EnumresultCriteriaResultCodeValueProp, bool)`

GetResultCodeValueOk returns a tuple with the ResultCodeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeValue

`func (o *ResultCriteriaListResponseResourcesInner) SetResultCodeValue(v []EnumresultCriteriaResultCodeValueProp)`

SetResultCodeValue sets ResultCodeValue field to given value.

### HasResultCodeValue

`func (o *ResultCriteriaListResponseResourcesInner) HasResultCodeValue() bool`

HasResultCodeValue returns a boolean if a field has been set.

### GetProcessingTimeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetProcessingTimeCriteria() EnumresultCriteriaProcessingTimeCriteriaProp`

GetProcessingTimeCriteria returns the ProcessingTimeCriteria field if non-nil, zero value otherwise.

### GetProcessingTimeCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetProcessingTimeCriteriaOk() (*EnumresultCriteriaProcessingTimeCriteriaProp, bool)`

GetProcessingTimeCriteriaOk returns a tuple with the ProcessingTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetProcessingTimeCriteria(v EnumresultCriteriaProcessingTimeCriteriaProp)`

SetProcessingTimeCriteria sets ProcessingTimeCriteria field to given value.

### HasProcessingTimeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasProcessingTimeCriteria() bool`

HasProcessingTimeCriteria returns a boolean if a field has been set.

### GetProcessingTimeValue

`func (o *ResultCriteriaListResponseResourcesInner) GetProcessingTimeValue() string`

GetProcessingTimeValue returns the ProcessingTimeValue field if non-nil, zero value otherwise.

### GetProcessingTimeValueOk

`func (o *ResultCriteriaListResponseResourcesInner) GetProcessingTimeValueOk() (*string, bool)`

GetProcessingTimeValueOk returns a tuple with the ProcessingTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeValue

`func (o *ResultCriteriaListResponseResourcesInner) SetProcessingTimeValue(v string)`

SetProcessingTimeValue sets ProcessingTimeValue field to given value.

### HasProcessingTimeValue

`func (o *ResultCriteriaListResponseResourcesInner) HasProcessingTimeValue() bool`

HasProcessingTimeValue returns a boolean if a field has been set.

### GetQueueTimeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetQueueTimeCriteria() EnumresultCriteriaQueueTimeCriteriaProp`

GetQueueTimeCriteria returns the QueueTimeCriteria field if non-nil, zero value otherwise.

### GetQueueTimeCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetQueueTimeCriteriaOk() (*EnumresultCriteriaQueueTimeCriteriaProp, bool)`

GetQueueTimeCriteriaOk returns a tuple with the QueueTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetQueueTimeCriteria(v EnumresultCriteriaQueueTimeCriteriaProp)`

SetQueueTimeCriteria sets QueueTimeCriteria field to given value.

### HasQueueTimeCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasQueueTimeCriteria() bool`

HasQueueTimeCriteria returns a boolean if a field has been set.

### GetQueueTimeValue

`func (o *ResultCriteriaListResponseResourcesInner) GetQueueTimeValue() string`

GetQueueTimeValue returns the QueueTimeValue field if non-nil, zero value otherwise.

### GetQueueTimeValueOk

`func (o *ResultCriteriaListResponseResourcesInner) GetQueueTimeValueOk() (*string, bool)`

GetQueueTimeValueOk returns a tuple with the QueueTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeValue

`func (o *ResultCriteriaListResponseResourcesInner) SetQueueTimeValue(v string)`

SetQueueTimeValue sets QueueTimeValue field to given value.

### HasQueueTimeValue

`func (o *ResultCriteriaListResponseResourcesInner) HasQueueTimeValue() bool`

HasQueueTimeValue returns a boolean if a field has been set.

### GetReferralReturned

`func (o *ResultCriteriaListResponseResourcesInner) GetReferralReturned() EnumresultCriteriaReferralReturnedProp`

GetReferralReturned returns the ReferralReturned field if non-nil, zero value otherwise.

### GetReferralReturnedOk

`func (o *ResultCriteriaListResponseResourcesInner) GetReferralReturnedOk() (*EnumresultCriteriaReferralReturnedProp, bool)`

GetReferralReturnedOk returns a tuple with the ReferralReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralReturned

`func (o *ResultCriteriaListResponseResourcesInner) SetReferralReturned(v EnumresultCriteriaReferralReturnedProp)`

SetReferralReturned sets ReferralReturned field to given value.

### HasReferralReturned

`func (o *ResultCriteriaListResponseResourcesInner) HasReferralReturned() bool`

HasReferralReturned returns a boolean if a field has been set.

### GetAllIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) GetAllIncludedResponseControl() []string`

GetAllIncludedResponseControl returns the AllIncludedResponseControl field if non-nil, zero value otherwise.

### GetAllIncludedResponseControlOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAllIncludedResponseControlOk() (*[]string, bool)`

GetAllIncludedResponseControlOk returns a tuple with the AllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) SetAllIncludedResponseControl(v []string)`

SetAllIncludedResponseControl sets AllIncludedResponseControl field to given value.

### HasAllIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) HasAllIncludedResponseControl() bool`

HasAllIncludedResponseControl returns a boolean if a field has been set.

### GetAnyIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) GetAnyIncludedResponseControl() []string`

GetAnyIncludedResponseControl returns the AnyIncludedResponseControl field if non-nil, zero value otherwise.

### GetAnyIncludedResponseControlOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAnyIncludedResponseControlOk() (*[]string, bool)`

GetAnyIncludedResponseControlOk returns a tuple with the AnyIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) SetAnyIncludedResponseControl(v []string)`

SetAnyIncludedResponseControl sets AnyIncludedResponseControl field to given value.

### HasAnyIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) HasAnyIncludedResponseControl() bool`

HasAnyIncludedResponseControl returns a boolean if a field has been set.

### GetNotAllIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) GetNotAllIncludedResponseControl() []string`

GetNotAllIncludedResponseControl returns the NotAllIncludedResponseControl field if non-nil, zero value otherwise.

### GetNotAllIncludedResponseControlOk

`func (o *ResultCriteriaListResponseResourcesInner) GetNotAllIncludedResponseControlOk() (*[]string, bool)`

GetNotAllIncludedResponseControlOk returns a tuple with the NotAllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) SetNotAllIncludedResponseControl(v []string)`

SetNotAllIncludedResponseControl sets NotAllIncludedResponseControl field to given value.

### HasNotAllIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) HasNotAllIncludedResponseControl() bool`

HasNotAllIncludedResponseControl returns a boolean if a field has been set.

### GetNoneIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) GetNoneIncludedResponseControl() []string`

GetNoneIncludedResponseControl returns the NoneIncludedResponseControl field if non-nil, zero value otherwise.

### GetNoneIncludedResponseControlOk

`func (o *ResultCriteriaListResponseResourcesInner) GetNoneIncludedResponseControlOk() (*[]string, bool)`

GetNoneIncludedResponseControlOk returns a tuple with the NoneIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) SetNoneIncludedResponseControl(v []string)`

SetNoneIncludedResponseControl sets NoneIncludedResponseControl field to given value.

### HasNoneIncludedResponseControl

`func (o *ResultCriteriaListResponseResourcesInner) HasNoneIncludedResponseControl() bool`

HasNoneIncludedResponseControl returns a boolean if a field has been set.

### GetUsedAlternateAuthzid

`func (o *ResultCriteriaListResponseResourcesInner) GetUsedAlternateAuthzid() EnumresultCriteriaUsedAlternateAuthzidProp`

GetUsedAlternateAuthzid returns the UsedAlternateAuthzid field if non-nil, zero value otherwise.

### GetUsedAlternateAuthzidOk

`func (o *ResultCriteriaListResponseResourcesInner) GetUsedAlternateAuthzidOk() (*EnumresultCriteriaUsedAlternateAuthzidProp, bool)`

GetUsedAlternateAuthzidOk returns a tuple with the UsedAlternateAuthzid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAlternateAuthzid

`func (o *ResultCriteriaListResponseResourcesInner) SetUsedAlternateAuthzid(v EnumresultCriteriaUsedAlternateAuthzidProp)`

SetUsedAlternateAuthzid sets UsedAlternateAuthzid field to given value.

### HasUsedAlternateAuthzid

`func (o *ResultCriteriaListResponseResourcesInner) HasUsedAlternateAuthzid() bool`

HasUsedAlternateAuthzid returns a boolean if a field has been set.

### GetUsedAnyPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) GetUsedAnyPrivilege() EnumresultCriteriaUsedAnyPrivilegeProp`

GetUsedAnyPrivilege returns the UsedAnyPrivilege field if non-nil, zero value otherwise.

### GetUsedAnyPrivilegeOk

`func (o *ResultCriteriaListResponseResourcesInner) GetUsedAnyPrivilegeOk() (*EnumresultCriteriaUsedAnyPrivilegeProp, bool)`

GetUsedAnyPrivilegeOk returns a tuple with the UsedAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAnyPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) SetUsedAnyPrivilege(v EnumresultCriteriaUsedAnyPrivilegeProp)`

SetUsedAnyPrivilege sets UsedAnyPrivilege field to given value.

### HasUsedAnyPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) HasUsedAnyPrivilege() bool`

HasUsedAnyPrivilege returns a boolean if a field has been set.

### GetUsedPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) GetUsedPrivilege() []EnumresultCriteriaUsedPrivilegeProp`

GetUsedPrivilege returns the UsedPrivilege field if non-nil, zero value otherwise.

### GetUsedPrivilegeOk

`func (o *ResultCriteriaListResponseResourcesInner) GetUsedPrivilegeOk() (*[]EnumresultCriteriaUsedPrivilegeProp, bool)`

GetUsedPrivilegeOk returns a tuple with the UsedPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) SetUsedPrivilege(v []EnumresultCriteriaUsedPrivilegeProp)`

SetUsedPrivilege sets UsedPrivilege field to given value.

### HasUsedPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) HasUsedPrivilege() bool`

HasUsedPrivilege returns a boolean if a field has been set.

### GetMissingAnyPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) GetMissingAnyPrivilege() EnumresultCriteriaMissingAnyPrivilegeProp`

GetMissingAnyPrivilege returns the MissingAnyPrivilege field if non-nil, zero value otherwise.

### GetMissingAnyPrivilegeOk

`func (o *ResultCriteriaListResponseResourcesInner) GetMissingAnyPrivilegeOk() (*EnumresultCriteriaMissingAnyPrivilegeProp, bool)`

GetMissingAnyPrivilegeOk returns a tuple with the MissingAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAnyPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) SetMissingAnyPrivilege(v EnumresultCriteriaMissingAnyPrivilegeProp)`

SetMissingAnyPrivilege sets MissingAnyPrivilege field to given value.

### HasMissingAnyPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) HasMissingAnyPrivilege() bool`

HasMissingAnyPrivilege returns a boolean if a field has been set.

### GetMissingPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) GetMissingPrivilege() []EnumresultCriteriaMissingPrivilegeProp`

GetMissingPrivilege returns the MissingPrivilege field if non-nil, zero value otherwise.

### GetMissingPrivilegeOk

`func (o *ResultCriteriaListResponseResourcesInner) GetMissingPrivilegeOk() (*[]EnumresultCriteriaMissingPrivilegeProp, bool)`

GetMissingPrivilegeOk returns a tuple with the MissingPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) SetMissingPrivilege(v []EnumresultCriteriaMissingPrivilegeProp)`

SetMissingPrivilege sets MissingPrivilege field to given value.

### HasMissingPrivilege

`func (o *ResultCriteriaListResponseResourcesInner) HasMissingPrivilege() bool`

HasMissingPrivilege returns a boolean if a field has been set.

### GetRetiredPasswordUsedForBind

`func (o *ResultCriteriaListResponseResourcesInner) GetRetiredPasswordUsedForBind() EnumresultCriteriaRetiredPasswordUsedForBindProp`

GetRetiredPasswordUsedForBind returns the RetiredPasswordUsedForBind field if non-nil, zero value otherwise.

### GetRetiredPasswordUsedForBindOk

`func (o *ResultCriteriaListResponseResourcesInner) GetRetiredPasswordUsedForBindOk() (*EnumresultCriteriaRetiredPasswordUsedForBindProp, bool)`

GetRetiredPasswordUsedForBindOk returns a tuple with the RetiredPasswordUsedForBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiredPasswordUsedForBind

`func (o *ResultCriteriaListResponseResourcesInner) SetRetiredPasswordUsedForBind(v EnumresultCriteriaRetiredPasswordUsedForBindProp)`

SetRetiredPasswordUsedForBind sets RetiredPasswordUsedForBind field to given value.

### HasRetiredPasswordUsedForBind

`func (o *ResultCriteriaListResponseResourcesInner) HasRetiredPasswordUsedForBind() bool`

HasRetiredPasswordUsedForBind returns a boolean if a field has been set.

### GetSearchEntryReturnedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchEntryReturnedCriteria() EnumresultCriteriaSearchEntryReturnedCriteriaProp`

GetSearchEntryReturnedCriteria returns the SearchEntryReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchEntryReturnedCriteriaOk() (*EnumresultCriteriaSearchEntryReturnedCriteriaProp, bool)`

GetSearchEntryReturnedCriteriaOk returns a tuple with the SearchEntryReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetSearchEntryReturnedCriteria(v EnumresultCriteriaSearchEntryReturnedCriteriaProp)`

SetSearchEntryReturnedCriteria sets SearchEntryReturnedCriteria field to given value.

### HasSearchEntryReturnedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasSearchEntryReturnedCriteria() bool`

HasSearchEntryReturnedCriteria returns a boolean if a field has been set.

### GetSearchEntryReturnedCount

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchEntryReturnedCount() int64`

GetSearchEntryReturnedCount returns the SearchEntryReturnedCount field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCountOk

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchEntryReturnedCountOk() (*int64, bool)`

GetSearchEntryReturnedCountOk returns a tuple with the SearchEntryReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCount

`func (o *ResultCriteriaListResponseResourcesInner) SetSearchEntryReturnedCount(v int64)`

SetSearchEntryReturnedCount sets SearchEntryReturnedCount field to given value.

### HasSearchEntryReturnedCount

`func (o *ResultCriteriaListResponseResourcesInner) HasSearchEntryReturnedCount() bool`

HasSearchEntryReturnedCount returns a boolean if a field has been set.

### GetSearchReferenceReturnedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchReferenceReturnedCriteria() EnumresultCriteriaSearchReferenceReturnedCriteriaProp`

GetSearchReferenceReturnedCriteria returns the SearchReferenceReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchReferenceReturnedCriteriaOk() (*EnumresultCriteriaSearchReferenceReturnedCriteriaProp, bool)`

GetSearchReferenceReturnedCriteriaOk returns a tuple with the SearchReferenceReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetSearchReferenceReturnedCriteria(v EnumresultCriteriaSearchReferenceReturnedCriteriaProp)`

SetSearchReferenceReturnedCriteria sets SearchReferenceReturnedCriteria field to given value.

### HasSearchReferenceReturnedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasSearchReferenceReturnedCriteria() bool`

HasSearchReferenceReturnedCriteria returns a boolean if a field has been set.

### GetSearchReferenceReturnedCount

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchReferenceReturnedCount() int64`

GetSearchReferenceReturnedCount returns the SearchReferenceReturnedCount field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCountOk

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchReferenceReturnedCountOk() (*int64, bool)`

GetSearchReferenceReturnedCountOk returns a tuple with the SearchReferenceReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCount

`func (o *ResultCriteriaListResponseResourcesInner) SetSearchReferenceReturnedCount(v int64)`

SetSearchReferenceReturnedCount sets SearchReferenceReturnedCount field to given value.

### HasSearchReferenceReturnedCount

`func (o *ResultCriteriaListResponseResourcesInner) HasSearchReferenceReturnedCount() bool`

HasSearchReferenceReturnedCount returns a boolean if a field has been set.

### GetSearchIndexedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchIndexedCriteria() EnumresultCriteriaSearchIndexedCriteriaProp`

GetSearchIndexedCriteria returns the SearchIndexedCriteria field if non-nil, zero value otherwise.

### GetSearchIndexedCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetSearchIndexedCriteriaOk() (*EnumresultCriteriaSearchIndexedCriteriaProp, bool)`

GetSearchIndexedCriteriaOk returns a tuple with the SearchIndexedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetSearchIndexedCriteria(v EnumresultCriteriaSearchIndexedCriteriaProp)`

SetSearchIndexedCriteria sets SearchIndexedCriteria field to given value.

### HasSearchIndexedCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasSearchIndexedCriteria() bool`

HasSearchIndexedCriteria returns a boolean if a field has been set.

### GetIncludedAuthzUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedAuthzUserBaseDN() []string`

GetIncludedAuthzUserBaseDN returns the IncludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedAuthzUserBaseDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedAuthzUserBaseDNOk() (*[]string, bool)`

GetIncludedAuthzUserBaseDNOk returns a tuple with the IncludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAuthzUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) SetIncludedAuthzUserBaseDN(v []string)`

SetIncludedAuthzUserBaseDN sets IncludedAuthzUserBaseDN field to given value.

### HasIncludedAuthzUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) HasIncludedAuthzUserBaseDN() bool`

HasIncludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetExcludedAuthzUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedAuthzUserBaseDN() []string`

GetExcludedAuthzUserBaseDN returns the ExcludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedAuthzUserBaseDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedAuthzUserBaseDNOk() (*[]string, bool)`

GetExcludedAuthzUserBaseDNOk returns a tuple with the ExcludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAuthzUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) SetExcludedAuthzUserBaseDN(v []string)`

SetExcludedAuthzUserBaseDN sets ExcludedAuthzUserBaseDN field to given value.

### HasExcludedAuthzUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) HasExcludedAuthzUserBaseDN() bool`

HasExcludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) GetAllIncludedAuthzUserGroupDN() []string`

GetAllIncludedAuthzUserGroupDN returns the AllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedAuthzUserGroupDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAllIncludedAuthzUserGroupDNOk returns a tuple with the AllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) SetAllIncludedAuthzUserGroupDN(v []string)`

SetAllIncludedAuthzUserGroupDN sets AllIncludedAuthzUserGroupDN field to given value.

### HasAllIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) HasAllIncludedAuthzUserGroupDN() bool`

HasAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) GetAnyIncludedAuthzUserGroupDN() []string`

GetAnyIncludedAuthzUserGroupDN returns the AnyIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedAuthzUserGroupDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAnyIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedAuthzUserGroupDNOk returns a tuple with the AnyIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) SetAnyIncludedAuthzUserGroupDN(v []string)`

SetAnyIncludedAuthzUserGroupDN sets AnyIncludedAuthzUserGroupDN field to given value.

### HasAnyIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) HasAnyIncludedAuthzUserGroupDN() bool`

HasAnyIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) GetNotAllIncludedAuthzUserGroupDN() []string`

GetNotAllIncludedAuthzUserGroupDN returns the NotAllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedAuthzUserGroupDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetNotAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedAuthzUserGroupDNOk returns a tuple with the NotAllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) SetNotAllIncludedAuthzUserGroupDN(v []string)`

SetNotAllIncludedAuthzUserGroupDN sets NotAllIncludedAuthzUserGroupDN field to given value.

### HasNotAllIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) HasNotAllIncludedAuthzUserGroupDN() bool`

HasNotAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) GetNoneIncludedAuthzUserGroupDN() []string`

GetNoneIncludedAuthzUserGroupDN returns the NoneIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedAuthzUserGroupDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetNoneIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedAuthzUserGroupDNOk returns a tuple with the NoneIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) SetNoneIncludedAuthzUserGroupDN(v []string)`

SetNoneIncludedAuthzUserGroupDN sets NoneIncludedAuthzUserGroupDN field to given value.

### HasNoneIncludedAuthzUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) HasNoneIncludedAuthzUserGroupDN() bool`

HasNoneIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *ResultCriteriaListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultCriteriaListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultCriteriaListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResultCriteriaListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *ResultCriteriaListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ResultCriteriaListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ResultCriteriaListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ResultCriteriaListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ResultCriteriaListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ResultCriteriaListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ResultCriteriaListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetIncludeAnonymousBinds

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludeAnonymousBinds() bool`

GetIncludeAnonymousBinds returns the IncludeAnonymousBinds field if non-nil, zero value otherwise.

### GetIncludeAnonymousBindsOk

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludeAnonymousBindsOk() (*bool, bool)`

GetIncludeAnonymousBindsOk returns a tuple with the IncludeAnonymousBinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnonymousBinds

`func (o *ResultCriteriaListResponseResourcesInner) SetIncludeAnonymousBinds(v bool)`

SetIncludeAnonymousBinds sets IncludeAnonymousBinds field to given value.

### HasIncludeAnonymousBinds

`func (o *ResultCriteriaListResponseResourcesInner) HasIncludeAnonymousBinds() bool`

HasIncludeAnonymousBinds returns a boolean if a field has been set.

### GetIncludedUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedUserBaseDN() []string`

GetIncludedUserBaseDN returns the IncludedUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedUserBaseDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedUserBaseDNOk() (*[]string, bool)`

GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) SetIncludedUserBaseDN(v []string)`

SetIncludedUserBaseDN sets IncludedUserBaseDN field to given value.

### HasIncludedUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) HasIncludedUserBaseDN() bool`

HasIncludedUserBaseDN returns a boolean if a field has been set.

### GetExcludedUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedUserBaseDN() []string`

GetExcludedUserBaseDN returns the ExcludedUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedUserBaseDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedUserBaseDNOk() (*[]string, bool)`

GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) SetExcludedUserBaseDN(v []string)`

SetExcludedUserBaseDN sets ExcludedUserBaseDN field to given value.

### HasExcludedUserBaseDN

`func (o *ResultCriteriaListResponseResourcesInner) HasExcludedUserBaseDN() bool`

HasExcludedUserBaseDN returns a boolean if a field has been set.

### GetIncludedUserFilter

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedUserFilter() []string`

GetIncludedUserFilter returns the IncludedUserFilter field if non-nil, zero value otherwise.

### GetIncludedUserFilterOk

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedUserFilterOk() (*[]string, bool)`

GetIncludedUserFilterOk returns a tuple with the IncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilter

`func (o *ResultCriteriaListResponseResourcesInner) SetIncludedUserFilter(v []string)`

SetIncludedUserFilter sets IncludedUserFilter field to given value.

### HasIncludedUserFilter

`func (o *ResultCriteriaListResponseResourcesInner) HasIncludedUserFilter() bool`

HasIncludedUserFilter returns a boolean if a field has been set.

### GetExcludedUserFilter

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedUserFilter() []string`

GetExcludedUserFilter returns the ExcludedUserFilter field if non-nil, zero value otherwise.

### GetExcludedUserFilterOk

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedUserFilterOk() (*[]string, bool)`

GetExcludedUserFilterOk returns a tuple with the ExcludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilter

`func (o *ResultCriteriaListResponseResourcesInner) SetExcludedUserFilter(v []string)`

SetExcludedUserFilter sets ExcludedUserFilter field to given value.

### HasExcludedUserFilter

`func (o *ResultCriteriaListResponseResourcesInner) HasExcludedUserFilter() bool`

HasExcludedUserFilter returns a boolean if a field has been set.

### GetIncludedUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedUserGroupDN() []string`

GetIncludedUserGroupDN returns the IncludedUserGroupDN field if non-nil, zero value otherwise.

### GetIncludedUserGroupDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetIncludedUserGroupDNOk() (*[]string, bool)`

GetIncludedUserGroupDNOk returns a tuple with the IncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) SetIncludedUserGroupDN(v []string)`

SetIncludedUserGroupDN sets IncludedUserGroupDN field to given value.

### HasIncludedUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) HasIncludedUserGroupDN() bool`

HasIncludedUserGroupDN returns a boolean if a field has been set.

### GetExcludedUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedUserGroupDN() []string`

GetExcludedUserGroupDN returns the ExcludedUserGroupDN field if non-nil, zero value otherwise.

### GetExcludedUserGroupDNOk

`func (o *ResultCriteriaListResponseResourcesInner) GetExcludedUserGroupDNOk() (*[]string, bool)`

GetExcludedUserGroupDNOk returns a tuple with the ExcludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) SetExcludedUserGroupDN(v []string)`

SetExcludedUserGroupDN sets ExcludedUserGroupDN field to given value.

### HasExcludedUserGroupDN

`func (o *ResultCriteriaListResponseResourcesInner) HasExcludedUserGroupDN() bool`

HasExcludedUserGroupDN returns a boolean if a field has been set.

### GetAllIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetAllIncludedResultCriteria() []string`

GetAllIncludedResultCriteria returns the AllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAllIncludedResultCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAllIncludedResultCriteriaOk() (*[]string, bool)`

GetAllIncludedResultCriteriaOk returns a tuple with the AllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetAllIncludedResultCriteria(v []string)`

SetAllIncludedResultCriteria sets AllIncludedResultCriteria field to given value.

### HasAllIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasAllIncludedResultCriteria() bool`

HasAllIncludedResultCriteria returns a boolean if a field has been set.

### GetAnyIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetAnyIncludedResultCriteria() []string`

GetAnyIncludedResultCriteria returns the AnyIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedResultCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAnyIncludedResultCriteriaOk() (*[]string, bool)`

GetAnyIncludedResultCriteriaOk returns a tuple with the AnyIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetAnyIncludedResultCriteria(v []string)`

SetAnyIncludedResultCriteria sets AnyIncludedResultCriteria field to given value.

### HasAnyIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasAnyIncludedResultCriteria() bool`

HasAnyIncludedResultCriteria returns a boolean if a field has been set.

### GetNotAllIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetNotAllIncludedResultCriteria() []string`

GetNotAllIncludedResultCriteria returns the NotAllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedResultCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetNotAllIncludedResultCriteriaOk() (*[]string, bool)`

GetNotAllIncludedResultCriteriaOk returns a tuple with the NotAllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetNotAllIncludedResultCriteria(v []string)`

SetNotAllIncludedResultCriteria sets NotAllIncludedResultCriteria field to given value.

### HasNotAllIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasNotAllIncludedResultCriteria() bool`

HasNotAllIncludedResultCriteria returns a boolean if a field has been set.

### GetNoneIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetNoneIncludedResultCriteria() []string`

GetNoneIncludedResultCriteria returns the NoneIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedResultCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetNoneIncludedResultCriteriaOk() (*[]string, bool)`

GetNoneIncludedResultCriteriaOk returns a tuple with the NoneIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetNoneIncludedResultCriteria(v []string)`

SetNoneIncludedResultCriteria sets NoneIncludedResultCriteria field to given value.

### HasNoneIncludedResultCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasNoneIncludedResultCriteria() bool`

HasNoneIncludedResultCriteria returns a boolean if a field has been set.

### GetLocalAssuranceLevel

`func (o *ResultCriteriaListResponseResourcesInner) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp`

GetLocalAssuranceLevel returns the LocalAssuranceLevel field if non-nil, zero value otherwise.

### GetLocalAssuranceLevelOk

`func (o *ResultCriteriaListResponseResourcesInner) GetLocalAssuranceLevelOk() (*[]EnumresultCriteriaLocalAssuranceLevelProp, bool)`

GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAssuranceLevel

`func (o *ResultCriteriaListResponseResourcesInner) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp)`

SetLocalAssuranceLevel sets LocalAssuranceLevel field to given value.


### GetRemoteAssuranceLevel

`func (o *ResultCriteriaListResponseResourcesInner) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp`

GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field if non-nil, zero value otherwise.

### GetRemoteAssuranceLevelOk

`func (o *ResultCriteriaListResponseResourcesInner) GetRemoteAssuranceLevelOk() (*[]EnumresultCriteriaRemoteAssuranceLevelProp, bool)`

GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssuranceLevel

`func (o *ResultCriteriaListResponseResourcesInner) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp)`

SetRemoteAssuranceLevel sets RemoteAssuranceLevel field to given value.


### GetAssuranceTimeoutCriteria

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp`

GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field if non-nil, zero value otherwise.

### GetAssuranceTimeoutCriteriaOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool)`

GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutCriteria

`func (o *ResultCriteriaListResponseResourcesInner) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp)`

SetAssuranceTimeoutCriteria sets AssuranceTimeoutCriteria field to given value.

### HasAssuranceTimeoutCriteria

`func (o *ResultCriteriaListResponseResourcesInner) HasAssuranceTimeoutCriteria() bool`

HasAssuranceTimeoutCriteria returns a boolean if a field has been set.

### GetAssuranceTimeoutValue

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceTimeoutValue() string`

GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field if non-nil, zero value otherwise.

### GetAssuranceTimeoutValueOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceTimeoutValueOk() (*string, bool)`

GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutValue

`func (o *ResultCriteriaListResponseResourcesInner) SetAssuranceTimeoutValue(v string)`

SetAssuranceTimeoutValue sets AssuranceTimeoutValue field to given value.

### HasAssuranceTimeoutValue

`func (o *ResultCriteriaListResponseResourcesInner) HasAssuranceTimeoutValue() bool`

HasAssuranceTimeoutValue returns a boolean if a field has been set.

### GetResponseDelayedByAssurance

`func (o *ResultCriteriaListResponseResourcesInner) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp`

GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field if non-nil, zero value otherwise.

### GetResponseDelayedByAssuranceOk

`func (o *ResultCriteriaListResponseResourcesInner) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool)`

GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelayedByAssurance

`func (o *ResultCriteriaListResponseResourcesInner) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp)`

SetResponseDelayedByAssurance sets ResponseDelayedByAssurance field to given value.

### HasResponseDelayedByAssurance

`func (o *ResultCriteriaListResponseResourcesInner) HasResponseDelayedByAssurance() bool`

HasResponseDelayedByAssurance returns a boolean if a field has been set.

### GetAssuranceBehaviorAlteredByControl

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp`

GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field if non-nil, zero value otherwise.

### GetAssuranceBehaviorAlteredByControlOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool)`

GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceBehaviorAlteredByControl

`func (o *ResultCriteriaListResponseResourcesInner) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp)`

SetAssuranceBehaviorAlteredByControl sets AssuranceBehaviorAlteredByControl field to given value.

### HasAssuranceBehaviorAlteredByControl

`func (o *ResultCriteriaListResponseResourcesInner) HasAssuranceBehaviorAlteredByControl() bool`

HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.

### GetAssuranceSatisfied

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp`

GetAssuranceSatisfied returns the AssuranceSatisfied field if non-nil, zero value otherwise.

### GetAssuranceSatisfiedOk

`func (o *ResultCriteriaListResponseResourcesInner) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool)`

GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceSatisfied

`func (o *ResultCriteriaListResponseResourcesInner) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp)`

SetAssuranceSatisfied sets AssuranceSatisfied field to given value.

### HasAssuranceSatisfied

`func (o *ResultCriteriaListResponseResourcesInner) HasAssuranceSatisfied() bool`

HasAssuranceSatisfied returns a boolean if a field has been set.

### GetExtensionClass

`func (o *ResultCriteriaListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ResultCriteriaListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ResultCriteriaListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ResultCriteriaListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ResultCriteriaListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ResultCriteriaListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ResultCriteriaListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


