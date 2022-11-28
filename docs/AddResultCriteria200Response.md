# AddResultCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Result Criteria | 
**Schemas** | [**[]EnumthirdPartyResultCriteriaSchemaUrn**](EnumthirdPartyResultCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for operations included in this Simple Result Criteria. | [optional] 
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
**SearchEntryReturnedCount** | Pointer to **int32** | Specifies the target number of entries returned for use when determining whether a search operation should be included in this Simple Result Criteria. This will be ignored for all operations other than search, and it will be ignored for search operations if the \&quot;search-entry-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**SearchReferenceReturnedCriteria** | Pointer to [**EnumresultCriteriaSearchReferenceReturnedCriteriaProp**](EnumresultCriteriaSearchReferenceReturnedCriteriaProp.md) |  | [optional] 
**SearchReferenceReturnedCount** | Pointer to **int32** | Specifies the target number of references returned for use when determining whether a search operation should be included in this Simple Result Criteria. This will be ignored for all operations other than search, and it will be ignored for search operations if the \&quot;search-reference-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**SearchIndexedCriteria** | Pointer to [**EnumresultCriteriaSearchIndexedCriteriaProp**](EnumresultCriteriaSearchIndexedCriteriaProp.md) |  | [optional] 
**IncludedAuthzUserBaseDN** | Pointer to **[]string** | Specifies a base DN below which authorization user entries may exist for operations included in this Simple Result Criteria. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**ExcludedAuthzUserBaseDN** | Pointer to **[]string** | Specifies a base DN below which authorization user entries may exist for operations excluded from this Simple Result Criteria. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**AllIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users must exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must be a member of all of those groups. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**AnyIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users may exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must be a member of at least one of those groups. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**NotAllIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users should not exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must not be a member of at least one of those groups (that is, the user may be a member of zero or more of those groups, but not of all of them). The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**NoneIncludedAuthzUserGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which authorization users must not exist for operations included in this Simple Result Criteria. If any group DNs are provided, then the authorization user must not be a member any of those groups. The authorization user could be the currently authenticated user on the connection (the user that performed the Bind operation), or different if proxied authorization was used to request that the operation be performed under the authorization of another user (as is the case for operations that come through a Directory Proxy Server). This property will be ignored for operations where no authentication or authorization has been performed. | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 
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

### NewAddResultCriteria200Response

`func NewAddResultCriteria200Response(id string, schemas []EnumthirdPartyResultCriteriaSchemaUrn, localAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp, remoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp, extensionClass string, ) *AddResultCriteria200Response`

NewAddResultCriteria200Response instantiates a new AddResultCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResultCriteria200ResponseWithDefaults

`func NewAddResultCriteria200ResponseWithDefaults() *AddResultCriteria200Response`

NewAddResultCriteria200ResponseWithDefaults instantiates a new AddResultCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AddResultCriteria200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddResultCriteria200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddResultCriteria200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddResultCriteria200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddResultCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddResultCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddResultCriteria200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddResultCriteria200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AddResultCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddResultCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddResultCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddResultCriteria200Response) GetSchemas() []EnumthirdPartyResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddResultCriteria200Response) GetSchemasOk() (*[]EnumthirdPartyResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddResultCriteria200Response) SetSchemas(v []EnumthirdPartyResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddResultCriteria200Response) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddResultCriteria200Response) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddResultCriteria200Response) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddResultCriteria200Response) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetResultCodeCriteria

`func (o *AddResultCriteria200Response) GetResultCodeCriteria() EnumresultCriteriaResultCodeCriteriaProp`

GetResultCodeCriteria returns the ResultCodeCriteria field if non-nil, zero value otherwise.

### GetResultCodeCriteriaOk

`func (o *AddResultCriteria200Response) GetResultCodeCriteriaOk() (*EnumresultCriteriaResultCodeCriteriaProp, bool)`

GetResultCodeCriteriaOk returns a tuple with the ResultCodeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeCriteria

`func (o *AddResultCriteria200Response) SetResultCodeCriteria(v EnumresultCriteriaResultCodeCriteriaProp)`

SetResultCodeCriteria sets ResultCodeCriteria field to given value.

### HasResultCodeCriteria

`func (o *AddResultCriteria200Response) HasResultCodeCriteria() bool`

HasResultCodeCriteria returns a boolean if a field has been set.

### GetResultCodeValue

`func (o *AddResultCriteria200Response) GetResultCodeValue() []EnumresultCriteriaResultCodeValueProp`

GetResultCodeValue returns the ResultCodeValue field if non-nil, zero value otherwise.

### GetResultCodeValueOk

`func (o *AddResultCriteria200Response) GetResultCodeValueOk() (*[]EnumresultCriteriaResultCodeValueProp, bool)`

GetResultCodeValueOk returns a tuple with the ResultCodeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeValue

`func (o *AddResultCriteria200Response) SetResultCodeValue(v []EnumresultCriteriaResultCodeValueProp)`

SetResultCodeValue sets ResultCodeValue field to given value.

### HasResultCodeValue

`func (o *AddResultCriteria200Response) HasResultCodeValue() bool`

HasResultCodeValue returns a boolean if a field has been set.

### GetProcessingTimeCriteria

`func (o *AddResultCriteria200Response) GetProcessingTimeCriteria() EnumresultCriteriaProcessingTimeCriteriaProp`

GetProcessingTimeCriteria returns the ProcessingTimeCriteria field if non-nil, zero value otherwise.

### GetProcessingTimeCriteriaOk

`func (o *AddResultCriteria200Response) GetProcessingTimeCriteriaOk() (*EnumresultCriteriaProcessingTimeCriteriaProp, bool)`

GetProcessingTimeCriteriaOk returns a tuple with the ProcessingTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeCriteria

`func (o *AddResultCriteria200Response) SetProcessingTimeCriteria(v EnumresultCriteriaProcessingTimeCriteriaProp)`

SetProcessingTimeCriteria sets ProcessingTimeCriteria field to given value.

### HasProcessingTimeCriteria

`func (o *AddResultCriteria200Response) HasProcessingTimeCriteria() bool`

HasProcessingTimeCriteria returns a boolean if a field has been set.

### GetProcessingTimeValue

`func (o *AddResultCriteria200Response) GetProcessingTimeValue() string`

GetProcessingTimeValue returns the ProcessingTimeValue field if non-nil, zero value otherwise.

### GetProcessingTimeValueOk

`func (o *AddResultCriteria200Response) GetProcessingTimeValueOk() (*string, bool)`

GetProcessingTimeValueOk returns a tuple with the ProcessingTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeValue

`func (o *AddResultCriteria200Response) SetProcessingTimeValue(v string)`

SetProcessingTimeValue sets ProcessingTimeValue field to given value.

### HasProcessingTimeValue

`func (o *AddResultCriteria200Response) HasProcessingTimeValue() bool`

HasProcessingTimeValue returns a boolean if a field has been set.

### GetQueueTimeCriteria

`func (o *AddResultCriteria200Response) GetQueueTimeCriteria() EnumresultCriteriaQueueTimeCriteriaProp`

GetQueueTimeCriteria returns the QueueTimeCriteria field if non-nil, zero value otherwise.

### GetQueueTimeCriteriaOk

`func (o *AddResultCriteria200Response) GetQueueTimeCriteriaOk() (*EnumresultCriteriaQueueTimeCriteriaProp, bool)`

GetQueueTimeCriteriaOk returns a tuple with the QueueTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeCriteria

`func (o *AddResultCriteria200Response) SetQueueTimeCriteria(v EnumresultCriteriaQueueTimeCriteriaProp)`

SetQueueTimeCriteria sets QueueTimeCriteria field to given value.

### HasQueueTimeCriteria

`func (o *AddResultCriteria200Response) HasQueueTimeCriteria() bool`

HasQueueTimeCriteria returns a boolean if a field has been set.

### GetQueueTimeValue

`func (o *AddResultCriteria200Response) GetQueueTimeValue() string`

GetQueueTimeValue returns the QueueTimeValue field if non-nil, zero value otherwise.

### GetQueueTimeValueOk

`func (o *AddResultCriteria200Response) GetQueueTimeValueOk() (*string, bool)`

GetQueueTimeValueOk returns a tuple with the QueueTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeValue

`func (o *AddResultCriteria200Response) SetQueueTimeValue(v string)`

SetQueueTimeValue sets QueueTimeValue field to given value.

### HasQueueTimeValue

`func (o *AddResultCriteria200Response) HasQueueTimeValue() bool`

HasQueueTimeValue returns a boolean if a field has been set.

### GetReferralReturned

`func (o *AddResultCriteria200Response) GetReferralReturned() EnumresultCriteriaReferralReturnedProp`

GetReferralReturned returns the ReferralReturned field if non-nil, zero value otherwise.

### GetReferralReturnedOk

`func (o *AddResultCriteria200Response) GetReferralReturnedOk() (*EnumresultCriteriaReferralReturnedProp, bool)`

GetReferralReturnedOk returns a tuple with the ReferralReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralReturned

`func (o *AddResultCriteria200Response) SetReferralReturned(v EnumresultCriteriaReferralReturnedProp)`

SetReferralReturned sets ReferralReturned field to given value.

### HasReferralReturned

`func (o *AddResultCriteria200Response) HasReferralReturned() bool`

HasReferralReturned returns a boolean if a field has been set.

### GetAllIncludedResponseControl

`func (o *AddResultCriteria200Response) GetAllIncludedResponseControl() []string`

GetAllIncludedResponseControl returns the AllIncludedResponseControl field if non-nil, zero value otherwise.

### GetAllIncludedResponseControlOk

`func (o *AddResultCriteria200Response) GetAllIncludedResponseControlOk() (*[]string, bool)`

GetAllIncludedResponseControlOk returns a tuple with the AllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResponseControl

`func (o *AddResultCriteria200Response) SetAllIncludedResponseControl(v []string)`

SetAllIncludedResponseControl sets AllIncludedResponseControl field to given value.

### HasAllIncludedResponseControl

`func (o *AddResultCriteria200Response) HasAllIncludedResponseControl() bool`

HasAllIncludedResponseControl returns a boolean if a field has been set.

### GetAnyIncludedResponseControl

`func (o *AddResultCriteria200Response) GetAnyIncludedResponseControl() []string`

GetAnyIncludedResponseControl returns the AnyIncludedResponseControl field if non-nil, zero value otherwise.

### GetAnyIncludedResponseControlOk

`func (o *AddResultCriteria200Response) GetAnyIncludedResponseControlOk() (*[]string, bool)`

GetAnyIncludedResponseControlOk returns a tuple with the AnyIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResponseControl

`func (o *AddResultCriteria200Response) SetAnyIncludedResponseControl(v []string)`

SetAnyIncludedResponseControl sets AnyIncludedResponseControl field to given value.

### HasAnyIncludedResponseControl

`func (o *AddResultCriteria200Response) HasAnyIncludedResponseControl() bool`

HasAnyIncludedResponseControl returns a boolean if a field has been set.

### GetNotAllIncludedResponseControl

`func (o *AddResultCriteria200Response) GetNotAllIncludedResponseControl() []string`

GetNotAllIncludedResponseControl returns the NotAllIncludedResponseControl field if non-nil, zero value otherwise.

### GetNotAllIncludedResponseControlOk

`func (o *AddResultCriteria200Response) GetNotAllIncludedResponseControlOk() (*[]string, bool)`

GetNotAllIncludedResponseControlOk returns a tuple with the NotAllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResponseControl

`func (o *AddResultCriteria200Response) SetNotAllIncludedResponseControl(v []string)`

SetNotAllIncludedResponseControl sets NotAllIncludedResponseControl field to given value.

### HasNotAllIncludedResponseControl

`func (o *AddResultCriteria200Response) HasNotAllIncludedResponseControl() bool`

HasNotAllIncludedResponseControl returns a boolean if a field has been set.

### GetNoneIncludedResponseControl

`func (o *AddResultCriteria200Response) GetNoneIncludedResponseControl() []string`

GetNoneIncludedResponseControl returns the NoneIncludedResponseControl field if non-nil, zero value otherwise.

### GetNoneIncludedResponseControlOk

`func (o *AddResultCriteria200Response) GetNoneIncludedResponseControlOk() (*[]string, bool)`

GetNoneIncludedResponseControlOk returns a tuple with the NoneIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResponseControl

`func (o *AddResultCriteria200Response) SetNoneIncludedResponseControl(v []string)`

SetNoneIncludedResponseControl sets NoneIncludedResponseControl field to given value.

### HasNoneIncludedResponseControl

`func (o *AddResultCriteria200Response) HasNoneIncludedResponseControl() bool`

HasNoneIncludedResponseControl returns a boolean if a field has been set.

### GetUsedAlternateAuthzid

`func (o *AddResultCriteria200Response) GetUsedAlternateAuthzid() EnumresultCriteriaUsedAlternateAuthzidProp`

GetUsedAlternateAuthzid returns the UsedAlternateAuthzid field if non-nil, zero value otherwise.

### GetUsedAlternateAuthzidOk

`func (o *AddResultCriteria200Response) GetUsedAlternateAuthzidOk() (*EnumresultCriteriaUsedAlternateAuthzidProp, bool)`

GetUsedAlternateAuthzidOk returns a tuple with the UsedAlternateAuthzid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAlternateAuthzid

`func (o *AddResultCriteria200Response) SetUsedAlternateAuthzid(v EnumresultCriteriaUsedAlternateAuthzidProp)`

SetUsedAlternateAuthzid sets UsedAlternateAuthzid field to given value.

### HasUsedAlternateAuthzid

`func (o *AddResultCriteria200Response) HasUsedAlternateAuthzid() bool`

HasUsedAlternateAuthzid returns a boolean if a field has been set.

### GetUsedAnyPrivilege

`func (o *AddResultCriteria200Response) GetUsedAnyPrivilege() EnumresultCriteriaUsedAnyPrivilegeProp`

GetUsedAnyPrivilege returns the UsedAnyPrivilege field if non-nil, zero value otherwise.

### GetUsedAnyPrivilegeOk

`func (o *AddResultCriteria200Response) GetUsedAnyPrivilegeOk() (*EnumresultCriteriaUsedAnyPrivilegeProp, bool)`

GetUsedAnyPrivilegeOk returns a tuple with the UsedAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAnyPrivilege

`func (o *AddResultCriteria200Response) SetUsedAnyPrivilege(v EnumresultCriteriaUsedAnyPrivilegeProp)`

SetUsedAnyPrivilege sets UsedAnyPrivilege field to given value.

### HasUsedAnyPrivilege

`func (o *AddResultCriteria200Response) HasUsedAnyPrivilege() bool`

HasUsedAnyPrivilege returns a boolean if a field has been set.

### GetUsedPrivilege

`func (o *AddResultCriteria200Response) GetUsedPrivilege() []EnumresultCriteriaUsedPrivilegeProp`

GetUsedPrivilege returns the UsedPrivilege field if non-nil, zero value otherwise.

### GetUsedPrivilegeOk

`func (o *AddResultCriteria200Response) GetUsedPrivilegeOk() (*[]EnumresultCriteriaUsedPrivilegeProp, bool)`

GetUsedPrivilegeOk returns a tuple with the UsedPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPrivilege

`func (o *AddResultCriteria200Response) SetUsedPrivilege(v []EnumresultCriteriaUsedPrivilegeProp)`

SetUsedPrivilege sets UsedPrivilege field to given value.

### HasUsedPrivilege

`func (o *AddResultCriteria200Response) HasUsedPrivilege() bool`

HasUsedPrivilege returns a boolean if a field has been set.

### GetMissingAnyPrivilege

`func (o *AddResultCriteria200Response) GetMissingAnyPrivilege() EnumresultCriteriaMissingAnyPrivilegeProp`

GetMissingAnyPrivilege returns the MissingAnyPrivilege field if non-nil, zero value otherwise.

### GetMissingAnyPrivilegeOk

`func (o *AddResultCriteria200Response) GetMissingAnyPrivilegeOk() (*EnumresultCriteriaMissingAnyPrivilegeProp, bool)`

GetMissingAnyPrivilegeOk returns a tuple with the MissingAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAnyPrivilege

`func (o *AddResultCriteria200Response) SetMissingAnyPrivilege(v EnumresultCriteriaMissingAnyPrivilegeProp)`

SetMissingAnyPrivilege sets MissingAnyPrivilege field to given value.

### HasMissingAnyPrivilege

`func (o *AddResultCriteria200Response) HasMissingAnyPrivilege() bool`

HasMissingAnyPrivilege returns a boolean if a field has been set.

### GetMissingPrivilege

`func (o *AddResultCriteria200Response) GetMissingPrivilege() []EnumresultCriteriaMissingPrivilegeProp`

GetMissingPrivilege returns the MissingPrivilege field if non-nil, zero value otherwise.

### GetMissingPrivilegeOk

`func (o *AddResultCriteria200Response) GetMissingPrivilegeOk() (*[]EnumresultCriteriaMissingPrivilegeProp, bool)`

GetMissingPrivilegeOk returns a tuple with the MissingPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPrivilege

`func (o *AddResultCriteria200Response) SetMissingPrivilege(v []EnumresultCriteriaMissingPrivilegeProp)`

SetMissingPrivilege sets MissingPrivilege field to given value.

### HasMissingPrivilege

`func (o *AddResultCriteria200Response) HasMissingPrivilege() bool`

HasMissingPrivilege returns a boolean if a field has been set.

### GetRetiredPasswordUsedForBind

`func (o *AddResultCriteria200Response) GetRetiredPasswordUsedForBind() EnumresultCriteriaRetiredPasswordUsedForBindProp`

GetRetiredPasswordUsedForBind returns the RetiredPasswordUsedForBind field if non-nil, zero value otherwise.

### GetRetiredPasswordUsedForBindOk

`func (o *AddResultCriteria200Response) GetRetiredPasswordUsedForBindOk() (*EnumresultCriteriaRetiredPasswordUsedForBindProp, bool)`

GetRetiredPasswordUsedForBindOk returns a tuple with the RetiredPasswordUsedForBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiredPasswordUsedForBind

`func (o *AddResultCriteria200Response) SetRetiredPasswordUsedForBind(v EnumresultCriteriaRetiredPasswordUsedForBindProp)`

SetRetiredPasswordUsedForBind sets RetiredPasswordUsedForBind field to given value.

### HasRetiredPasswordUsedForBind

`func (o *AddResultCriteria200Response) HasRetiredPasswordUsedForBind() bool`

HasRetiredPasswordUsedForBind returns a boolean if a field has been set.

### GetSearchEntryReturnedCriteria

`func (o *AddResultCriteria200Response) GetSearchEntryReturnedCriteria() EnumresultCriteriaSearchEntryReturnedCriteriaProp`

GetSearchEntryReturnedCriteria returns the SearchEntryReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCriteriaOk

`func (o *AddResultCriteria200Response) GetSearchEntryReturnedCriteriaOk() (*EnumresultCriteriaSearchEntryReturnedCriteriaProp, bool)`

GetSearchEntryReturnedCriteriaOk returns a tuple with the SearchEntryReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCriteria

`func (o *AddResultCriteria200Response) SetSearchEntryReturnedCriteria(v EnumresultCriteriaSearchEntryReturnedCriteriaProp)`

SetSearchEntryReturnedCriteria sets SearchEntryReturnedCriteria field to given value.

### HasSearchEntryReturnedCriteria

`func (o *AddResultCriteria200Response) HasSearchEntryReturnedCriteria() bool`

HasSearchEntryReturnedCriteria returns a boolean if a field has been set.

### GetSearchEntryReturnedCount

`func (o *AddResultCriteria200Response) GetSearchEntryReturnedCount() int32`

GetSearchEntryReturnedCount returns the SearchEntryReturnedCount field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCountOk

`func (o *AddResultCriteria200Response) GetSearchEntryReturnedCountOk() (*int32, bool)`

GetSearchEntryReturnedCountOk returns a tuple with the SearchEntryReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCount

`func (o *AddResultCriteria200Response) SetSearchEntryReturnedCount(v int32)`

SetSearchEntryReturnedCount sets SearchEntryReturnedCount field to given value.

### HasSearchEntryReturnedCount

`func (o *AddResultCriteria200Response) HasSearchEntryReturnedCount() bool`

HasSearchEntryReturnedCount returns a boolean if a field has been set.

### GetSearchReferenceReturnedCriteria

`func (o *AddResultCriteria200Response) GetSearchReferenceReturnedCriteria() EnumresultCriteriaSearchReferenceReturnedCriteriaProp`

GetSearchReferenceReturnedCriteria returns the SearchReferenceReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCriteriaOk

`func (o *AddResultCriteria200Response) GetSearchReferenceReturnedCriteriaOk() (*EnumresultCriteriaSearchReferenceReturnedCriteriaProp, bool)`

GetSearchReferenceReturnedCriteriaOk returns a tuple with the SearchReferenceReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCriteria

`func (o *AddResultCriteria200Response) SetSearchReferenceReturnedCriteria(v EnumresultCriteriaSearchReferenceReturnedCriteriaProp)`

SetSearchReferenceReturnedCriteria sets SearchReferenceReturnedCriteria field to given value.

### HasSearchReferenceReturnedCriteria

`func (o *AddResultCriteria200Response) HasSearchReferenceReturnedCriteria() bool`

HasSearchReferenceReturnedCriteria returns a boolean if a field has been set.

### GetSearchReferenceReturnedCount

`func (o *AddResultCriteria200Response) GetSearchReferenceReturnedCount() int32`

GetSearchReferenceReturnedCount returns the SearchReferenceReturnedCount field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCountOk

`func (o *AddResultCriteria200Response) GetSearchReferenceReturnedCountOk() (*int32, bool)`

GetSearchReferenceReturnedCountOk returns a tuple with the SearchReferenceReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCount

`func (o *AddResultCriteria200Response) SetSearchReferenceReturnedCount(v int32)`

SetSearchReferenceReturnedCount sets SearchReferenceReturnedCount field to given value.

### HasSearchReferenceReturnedCount

`func (o *AddResultCriteria200Response) HasSearchReferenceReturnedCount() bool`

HasSearchReferenceReturnedCount returns a boolean if a field has been set.

### GetSearchIndexedCriteria

`func (o *AddResultCriteria200Response) GetSearchIndexedCriteria() EnumresultCriteriaSearchIndexedCriteriaProp`

GetSearchIndexedCriteria returns the SearchIndexedCriteria field if non-nil, zero value otherwise.

### GetSearchIndexedCriteriaOk

`func (o *AddResultCriteria200Response) GetSearchIndexedCriteriaOk() (*EnumresultCriteriaSearchIndexedCriteriaProp, bool)`

GetSearchIndexedCriteriaOk returns a tuple with the SearchIndexedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexedCriteria

`func (o *AddResultCriteria200Response) SetSearchIndexedCriteria(v EnumresultCriteriaSearchIndexedCriteriaProp)`

SetSearchIndexedCriteria sets SearchIndexedCriteria field to given value.

### HasSearchIndexedCriteria

`func (o *AddResultCriteria200Response) HasSearchIndexedCriteria() bool`

HasSearchIndexedCriteria returns a boolean if a field has been set.

### GetIncludedAuthzUserBaseDN

`func (o *AddResultCriteria200Response) GetIncludedAuthzUserBaseDN() []string`

GetIncludedAuthzUserBaseDN returns the IncludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedAuthzUserBaseDNOk

`func (o *AddResultCriteria200Response) GetIncludedAuthzUserBaseDNOk() (*[]string, bool)`

GetIncludedAuthzUserBaseDNOk returns a tuple with the IncludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAuthzUserBaseDN

`func (o *AddResultCriteria200Response) SetIncludedAuthzUserBaseDN(v []string)`

SetIncludedAuthzUserBaseDN sets IncludedAuthzUserBaseDN field to given value.

### HasIncludedAuthzUserBaseDN

`func (o *AddResultCriteria200Response) HasIncludedAuthzUserBaseDN() bool`

HasIncludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetExcludedAuthzUserBaseDN

`func (o *AddResultCriteria200Response) GetExcludedAuthzUserBaseDN() []string`

GetExcludedAuthzUserBaseDN returns the ExcludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedAuthzUserBaseDNOk

`func (o *AddResultCriteria200Response) GetExcludedAuthzUserBaseDNOk() (*[]string, bool)`

GetExcludedAuthzUserBaseDNOk returns a tuple with the ExcludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAuthzUserBaseDN

`func (o *AddResultCriteria200Response) SetExcludedAuthzUserBaseDN(v []string)`

SetExcludedAuthzUserBaseDN sets ExcludedAuthzUserBaseDN field to given value.

### HasExcludedAuthzUserBaseDN

`func (o *AddResultCriteria200Response) HasExcludedAuthzUserBaseDN() bool`

HasExcludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) GetAllIncludedAuthzUserGroupDN() []string`

GetAllIncludedAuthzUserGroupDN returns the AllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteria200Response) GetAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAllIncludedAuthzUserGroupDNOk returns a tuple with the AllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) SetAllIncludedAuthzUserGroupDN(v []string)`

SetAllIncludedAuthzUserGroupDN sets AllIncludedAuthzUserGroupDN field to given value.

### HasAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) HasAllIncludedAuthzUserGroupDN() bool`

HasAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) GetAnyIncludedAuthzUserGroupDN() []string`

GetAnyIncludedAuthzUserGroupDN returns the AnyIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteria200Response) GetAnyIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedAuthzUserGroupDNOk returns a tuple with the AnyIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) SetAnyIncludedAuthzUserGroupDN(v []string)`

SetAnyIncludedAuthzUserGroupDN sets AnyIncludedAuthzUserGroupDN field to given value.

### HasAnyIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) HasAnyIncludedAuthzUserGroupDN() bool`

HasAnyIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) GetNotAllIncludedAuthzUserGroupDN() []string`

GetNotAllIncludedAuthzUserGroupDN returns the NotAllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteria200Response) GetNotAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedAuthzUserGroupDNOk returns a tuple with the NotAllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) SetNotAllIncludedAuthzUserGroupDN(v []string)`

SetNotAllIncludedAuthzUserGroupDN sets NotAllIncludedAuthzUserGroupDN field to given value.

### HasNotAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) HasNotAllIncludedAuthzUserGroupDN() bool`

HasNotAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) GetNoneIncludedAuthzUserGroupDN() []string`

GetNoneIncludedAuthzUserGroupDN returns the NoneIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteria200Response) GetNoneIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedAuthzUserGroupDNOk returns a tuple with the NoneIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) SetNoneIncludedAuthzUserGroupDN(v []string)`

SetNoneIncludedAuthzUserGroupDN sets NoneIncludedAuthzUserGroupDN field to given value.

### HasNoneIncludedAuthzUserGroupDN

`func (o *AddResultCriteria200Response) HasNoneIncludedAuthzUserGroupDN() bool`

HasNoneIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddResultCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddResultCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddResultCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddResultCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllIncludedResultCriteria

`func (o *AddResultCriteria200Response) GetAllIncludedResultCriteria() []string`

GetAllIncludedResultCriteria returns the AllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAllIncludedResultCriteriaOk

`func (o *AddResultCriteria200Response) GetAllIncludedResultCriteriaOk() (*[]string, bool)`

GetAllIncludedResultCriteriaOk returns a tuple with the AllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResultCriteria

`func (o *AddResultCriteria200Response) SetAllIncludedResultCriteria(v []string)`

SetAllIncludedResultCriteria sets AllIncludedResultCriteria field to given value.

### HasAllIncludedResultCriteria

`func (o *AddResultCriteria200Response) HasAllIncludedResultCriteria() bool`

HasAllIncludedResultCriteria returns a boolean if a field has been set.

### GetAnyIncludedResultCriteria

`func (o *AddResultCriteria200Response) GetAnyIncludedResultCriteria() []string`

GetAnyIncludedResultCriteria returns the AnyIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedResultCriteriaOk

`func (o *AddResultCriteria200Response) GetAnyIncludedResultCriteriaOk() (*[]string, bool)`

GetAnyIncludedResultCriteriaOk returns a tuple with the AnyIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResultCriteria

`func (o *AddResultCriteria200Response) SetAnyIncludedResultCriteria(v []string)`

SetAnyIncludedResultCriteria sets AnyIncludedResultCriteria field to given value.

### HasAnyIncludedResultCriteria

`func (o *AddResultCriteria200Response) HasAnyIncludedResultCriteria() bool`

HasAnyIncludedResultCriteria returns a boolean if a field has been set.

### GetNotAllIncludedResultCriteria

`func (o *AddResultCriteria200Response) GetNotAllIncludedResultCriteria() []string`

GetNotAllIncludedResultCriteria returns the NotAllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedResultCriteriaOk

`func (o *AddResultCriteria200Response) GetNotAllIncludedResultCriteriaOk() (*[]string, bool)`

GetNotAllIncludedResultCriteriaOk returns a tuple with the NotAllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResultCriteria

`func (o *AddResultCriteria200Response) SetNotAllIncludedResultCriteria(v []string)`

SetNotAllIncludedResultCriteria sets NotAllIncludedResultCriteria field to given value.

### HasNotAllIncludedResultCriteria

`func (o *AddResultCriteria200Response) HasNotAllIncludedResultCriteria() bool`

HasNotAllIncludedResultCriteria returns a boolean if a field has been set.

### GetNoneIncludedResultCriteria

`func (o *AddResultCriteria200Response) GetNoneIncludedResultCriteria() []string`

GetNoneIncludedResultCriteria returns the NoneIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedResultCriteriaOk

`func (o *AddResultCriteria200Response) GetNoneIncludedResultCriteriaOk() (*[]string, bool)`

GetNoneIncludedResultCriteriaOk returns a tuple with the NoneIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResultCriteria

`func (o *AddResultCriteria200Response) SetNoneIncludedResultCriteria(v []string)`

SetNoneIncludedResultCriteria sets NoneIncludedResultCriteria field to given value.

### HasNoneIncludedResultCriteria

`func (o *AddResultCriteria200Response) HasNoneIncludedResultCriteria() bool`

HasNoneIncludedResultCriteria returns a boolean if a field has been set.

### GetLocalAssuranceLevel

`func (o *AddResultCriteria200Response) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp`

GetLocalAssuranceLevel returns the LocalAssuranceLevel field if non-nil, zero value otherwise.

### GetLocalAssuranceLevelOk

`func (o *AddResultCriteria200Response) GetLocalAssuranceLevelOk() (*[]EnumresultCriteriaLocalAssuranceLevelProp, bool)`

GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAssuranceLevel

`func (o *AddResultCriteria200Response) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp)`

SetLocalAssuranceLevel sets LocalAssuranceLevel field to given value.


### GetRemoteAssuranceLevel

`func (o *AddResultCriteria200Response) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp`

GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field if non-nil, zero value otherwise.

### GetRemoteAssuranceLevelOk

`func (o *AddResultCriteria200Response) GetRemoteAssuranceLevelOk() (*[]EnumresultCriteriaRemoteAssuranceLevelProp, bool)`

GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssuranceLevel

`func (o *AddResultCriteria200Response) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp)`

SetRemoteAssuranceLevel sets RemoteAssuranceLevel field to given value.


### GetAssuranceTimeoutCriteria

`func (o *AddResultCriteria200Response) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp`

GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field if non-nil, zero value otherwise.

### GetAssuranceTimeoutCriteriaOk

`func (o *AddResultCriteria200Response) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool)`

GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutCriteria

`func (o *AddResultCriteria200Response) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp)`

SetAssuranceTimeoutCriteria sets AssuranceTimeoutCriteria field to given value.

### HasAssuranceTimeoutCriteria

`func (o *AddResultCriteria200Response) HasAssuranceTimeoutCriteria() bool`

HasAssuranceTimeoutCriteria returns a boolean if a field has been set.

### GetAssuranceTimeoutValue

`func (o *AddResultCriteria200Response) GetAssuranceTimeoutValue() string`

GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field if non-nil, zero value otherwise.

### GetAssuranceTimeoutValueOk

`func (o *AddResultCriteria200Response) GetAssuranceTimeoutValueOk() (*string, bool)`

GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutValue

`func (o *AddResultCriteria200Response) SetAssuranceTimeoutValue(v string)`

SetAssuranceTimeoutValue sets AssuranceTimeoutValue field to given value.

### HasAssuranceTimeoutValue

`func (o *AddResultCriteria200Response) HasAssuranceTimeoutValue() bool`

HasAssuranceTimeoutValue returns a boolean if a field has been set.

### GetResponseDelayedByAssurance

`func (o *AddResultCriteria200Response) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp`

GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field if non-nil, zero value otherwise.

### GetResponseDelayedByAssuranceOk

`func (o *AddResultCriteria200Response) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool)`

GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelayedByAssurance

`func (o *AddResultCriteria200Response) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp)`

SetResponseDelayedByAssurance sets ResponseDelayedByAssurance field to given value.

### HasResponseDelayedByAssurance

`func (o *AddResultCriteria200Response) HasResponseDelayedByAssurance() bool`

HasResponseDelayedByAssurance returns a boolean if a field has been set.

### GetAssuranceBehaviorAlteredByControl

`func (o *AddResultCriteria200Response) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp`

GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field if non-nil, zero value otherwise.

### GetAssuranceBehaviorAlteredByControlOk

`func (o *AddResultCriteria200Response) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool)`

GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceBehaviorAlteredByControl

`func (o *AddResultCriteria200Response) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp)`

SetAssuranceBehaviorAlteredByControl sets AssuranceBehaviorAlteredByControl field to given value.

### HasAssuranceBehaviorAlteredByControl

`func (o *AddResultCriteria200Response) HasAssuranceBehaviorAlteredByControl() bool`

HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.

### GetAssuranceSatisfied

`func (o *AddResultCriteria200Response) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp`

GetAssuranceSatisfied returns the AssuranceSatisfied field if non-nil, zero value otherwise.

### GetAssuranceSatisfiedOk

`func (o *AddResultCriteria200Response) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool)`

GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceSatisfied

`func (o *AddResultCriteria200Response) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp)`

SetAssuranceSatisfied sets AssuranceSatisfied field to given value.

### HasAssuranceSatisfied

`func (o *AddResultCriteria200Response) HasAssuranceSatisfied() bool`

HasAssuranceSatisfied returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddResultCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddResultCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddResultCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddResultCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddResultCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddResultCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddResultCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


