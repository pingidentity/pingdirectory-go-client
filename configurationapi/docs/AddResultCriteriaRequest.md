# AddResultCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Result Criteria | 
**Schemas** | [**[]EnumthirdPartyResultCriteriaSchemaUrn**](EnumthirdPartyResultCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for operations included in this Simple Result Criteria. | [optional] 
**IncludeAnonymousBinds** | Pointer to **bool** | Indicates whether this criteria will be permitted to match bind operations that resulted in anonymous authentication. | [optional] 
**IncludedUserBaseDN** | Pointer to **[]string** | A set of base DNs for authenticated users that will be permitted to match this criteria. | [optional] 
**ExcludedUserBaseDN** | Pointer to **[]string** | A set of base DNs for authenticated users that will not be permitted to match this criteria. | [optional] 
**IncludedUserFilter** | Pointer to **[]string** | A set of filters that may be used to identify entries for authenticated users that will be permitted to match this criteria. | [optional] 
**ExcludedUserFilter** | Pointer to **[]string** | A set of filters that may be used to identify entries for authenticated users that will not be permitted to match this criteria. | [optional] 
**IncludedUserGroupDN** | Pointer to **[]string** | The DNs of the groups whose members will be permitted to match this criteria. | [optional] 
**ExcludedUserGroupDN** | Pointer to **[]string** | The DNs of the groups whose members will not be permitted to match this criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 
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
**AllIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that must match the associated operation result in order to match the aggregate result criteria. If one or more all-included result criteria objects are provided, then an operation result must match all of them in order to match the aggregate result criteria. | [optional] 
**AnyIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that may match the associated operation result in order to match the aggregate result criteria. If one or more any-included result criteria objects are provided, then an operation result must match at least one of them in order to match the aggregate result criteria. | [optional] 
**NotAllIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that should not match the associated operation result in order to match the aggregate result criteria. If one or more not-all-included result criteria objects are provided, then an operation result must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate result criteria. | [optional] 
**NoneIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that must not match the associated operation result in order to match the aggregate result criteria. If one or more none-included result criteria objects are provided, then an operation result must not match any of them in order to match the aggregate result criteria. | [optional] 
**LocalAssuranceLevel** | Pointer to [**[]EnumresultCriteriaLocalAssuranceLevelProp**](EnumresultCriteriaLocalAssuranceLevelProp.md) |  | [optional] 
**RemoteAssuranceLevel** | Pointer to [**[]EnumresultCriteriaRemoteAssuranceLevelProp**](EnumresultCriteriaRemoteAssuranceLevelProp.md) |  | [optional] 
**AssuranceTimeoutCriteria** | Pointer to [**EnumresultCriteriaAssuranceTimeoutCriteriaProp**](EnumresultCriteriaAssuranceTimeoutCriteriaProp.md) |  | [optional] 
**AssuranceTimeoutValue** | Pointer to **string** | The value to use for performing matching based on the assurance timeout. This will be ignored if the assurance-timeout-criteria is \&quot;any\&quot;. | [optional] 
**ResponseDelayedByAssurance** | Pointer to [**EnumresultCriteriaResponseDelayedByAssuranceProp**](EnumresultCriteriaResponseDelayedByAssuranceProp.md) |  | [optional] 
**AssuranceBehaviorAlteredByControl** | Pointer to [**EnumresultCriteriaAssuranceBehaviorAlteredByControlProp**](EnumresultCriteriaAssuranceBehaviorAlteredByControlProp.md) |  | [optional] 
**AssuranceSatisfied** | Pointer to [**EnumresultCriteriaAssuranceSatisfiedProp**](EnumresultCriteriaAssuranceSatisfiedProp.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Result Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Result Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddResultCriteriaRequest

`func NewAddResultCriteriaRequest(criteriaName string, schemas []EnumthirdPartyResultCriteriaSchemaUrn, extensionClass string, ) *AddResultCriteriaRequest`

NewAddResultCriteriaRequest instantiates a new AddResultCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResultCriteriaRequestWithDefaults

`func NewAddResultCriteriaRequestWithDefaults() *AddResultCriteriaRequest`

NewAddResultCriteriaRequestWithDefaults instantiates a new AddResultCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddResultCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddResultCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddResultCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddResultCriteriaRequest) GetSchemas() []EnumthirdPartyResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddResultCriteriaRequest) GetSchemasOk() (*[]EnumthirdPartyResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddResultCriteriaRequest) SetSchemas(v []EnumthirdPartyResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddResultCriteriaRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddResultCriteriaRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddResultCriteriaRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddResultCriteriaRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetIncludeAnonymousBinds

`func (o *AddResultCriteriaRequest) GetIncludeAnonymousBinds() bool`

GetIncludeAnonymousBinds returns the IncludeAnonymousBinds field if non-nil, zero value otherwise.

### GetIncludeAnonymousBindsOk

`func (o *AddResultCriteriaRequest) GetIncludeAnonymousBindsOk() (*bool, bool)`

GetIncludeAnonymousBindsOk returns a tuple with the IncludeAnonymousBinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnonymousBinds

`func (o *AddResultCriteriaRequest) SetIncludeAnonymousBinds(v bool)`

SetIncludeAnonymousBinds sets IncludeAnonymousBinds field to given value.

### HasIncludeAnonymousBinds

`func (o *AddResultCriteriaRequest) HasIncludeAnonymousBinds() bool`

HasIncludeAnonymousBinds returns a boolean if a field has been set.

### GetIncludedUserBaseDN

`func (o *AddResultCriteriaRequest) GetIncludedUserBaseDN() []string`

GetIncludedUserBaseDN returns the IncludedUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedUserBaseDNOk

`func (o *AddResultCriteriaRequest) GetIncludedUserBaseDNOk() (*[]string, bool)`

GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserBaseDN

`func (o *AddResultCriteriaRequest) SetIncludedUserBaseDN(v []string)`

SetIncludedUserBaseDN sets IncludedUserBaseDN field to given value.

### HasIncludedUserBaseDN

`func (o *AddResultCriteriaRequest) HasIncludedUserBaseDN() bool`

HasIncludedUserBaseDN returns a boolean if a field has been set.

### GetExcludedUserBaseDN

`func (o *AddResultCriteriaRequest) GetExcludedUserBaseDN() []string`

GetExcludedUserBaseDN returns the ExcludedUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedUserBaseDNOk

`func (o *AddResultCriteriaRequest) GetExcludedUserBaseDNOk() (*[]string, bool)`

GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserBaseDN

`func (o *AddResultCriteriaRequest) SetExcludedUserBaseDN(v []string)`

SetExcludedUserBaseDN sets ExcludedUserBaseDN field to given value.

### HasExcludedUserBaseDN

`func (o *AddResultCriteriaRequest) HasExcludedUserBaseDN() bool`

HasExcludedUserBaseDN returns a boolean if a field has been set.

### GetIncludedUserFilter

`func (o *AddResultCriteriaRequest) GetIncludedUserFilter() []string`

GetIncludedUserFilter returns the IncludedUserFilter field if non-nil, zero value otherwise.

### GetIncludedUserFilterOk

`func (o *AddResultCriteriaRequest) GetIncludedUserFilterOk() (*[]string, bool)`

GetIncludedUserFilterOk returns a tuple with the IncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilter

`func (o *AddResultCriteriaRequest) SetIncludedUserFilter(v []string)`

SetIncludedUserFilter sets IncludedUserFilter field to given value.

### HasIncludedUserFilter

`func (o *AddResultCriteriaRequest) HasIncludedUserFilter() bool`

HasIncludedUserFilter returns a boolean if a field has been set.

### GetExcludedUserFilter

`func (o *AddResultCriteriaRequest) GetExcludedUserFilter() []string`

GetExcludedUserFilter returns the ExcludedUserFilter field if non-nil, zero value otherwise.

### GetExcludedUserFilterOk

`func (o *AddResultCriteriaRequest) GetExcludedUserFilterOk() (*[]string, bool)`

GetExcludedUserFilterOk returns a tuple with the ExcludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilter

`func (o *AddResultCriteriaRequest) SetExcludedUserFilter(v []string)`

SetExcludedUserFilter sets ExcludedUserFilter field to given value.

### HasExcludedUserFilter

`func (o *AddResultCriteriaRequest) HasExcludedUserFilter() bool`

HasExcludedUserFilter returns a boolean if a field has been set.

### GetIncludedUserGroupDN

`func (o *AddResultCriteriaRequest) GetIncludedUserGroupDN() []string`

GetIncludedUserGroupDN returns the IncludedUserGroupDN field if non-nil, zero value otherwise.

### GetIncludedUserGroupDNOk

`func (o *AddResultCriteriaRequest) GetIncludedUserGroupDNOk() (*[]string, bool)`

GetIncludedUserGroupDNOk returns a tuple with the IncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserGroupDN

`func (o *AddResultCriteriaRequest) SetIncludedUserGroupDN(v []string)`

SetIncludedUserGroupDN sets IncludedUserGroupDN field to given value.

### HasIncludedUserGroupDN

`func (o *AddResultCriteriaRequest) HasIncludedUserGroupDN() bool`

HasIncludedUserGroupDN returns a boolean if a field has been set.

### GetExcludedUserGroupDN

`func (o *AddResultCriteriaRequest) GetExcludedUserGroupDN() []string`

GetExcludedUserGroupDN returns the ExcludedUserGroupDN field if non-nil, zero value otherwise.

### GetExcludedUserGroupDNOk

`func (o *AddResultCriteriaRequest) GetExcludedUserGroupDNOk() (*[]string, bool)`

GetExcludedUserGroupDNOk returns a tuple with the ExcludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserGroupDN

`func (o *AddResultCriteriaRequest) SetExcludedUserGroupDN(v []string)`

SetExcludedUserGroupDN sets ExcludedUserGroupDN field to given value.

### HasExcludedUserGroupDN

`func (o *AddResultCriteriaRequest) HasExcludedUserGroupDN() bool`

HasExcludedUserGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddResultCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddResultCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddResultCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddResultCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResultCodeCriteria

`func (o *AddResultCriteriaRequest) GetResultCodeCriteria() EnumresultCriteriaResultCodeCriteriaProp`

GetResultCodeCriteria returns the ResultCodeCriteria field if non-nil, zero value otherwise.

### GetResultCodeCriteriaOk

`func (o *AddResultCriteriaRequest) GetResultCodeCriteriaOk() (*EnumresultCriteriaResultCodeCriteriaProp, bool)`

GetResultCodeCriteriaOk returns a tuple with the ResultCodeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeCriteria

`func (o *AddResultCriteriaRequest) SetResultCodeCriteria(v EnumresultCriteriaResultCodeCriteriaProp)`

SetResultCodeCriteria sets ResultCodeCriteria field to given value.

### HasResultCodeCriteria

`func (o *AddResultCriteriaRequest) HasResultCodeCriteria() bool`

HasResultCodeCriteria returns a boolean if a field has been set.

### GetResultCodeValue

`func (o *AddResultCriteriaRequest) GetResultCodeValue() []EnumresultCriteriaResultCodeValueProp`

GetResultCodeValue returns the ResultCodeValue field if non-nil, zero value otherwise.

### GetResultCodeValueOk

`func (o *AddResultCriteriaRequest) GetResultCodeValueOk() (*[]EnumresultCriteriaResultCodeValueProp, bool)`

GetResultCodeValueOk returns a tuple with the ResultCodeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeValue

`func (o *AddResultCriteriaRequest) SetResultCodeValue(v []EnumresultCriteriaResultCodeValueProp)`

SetResultCodeValue sets ResultCodeValue field to given value.

### HasResultCodeValue

`func (o *AddResultCriteriaRequest) HasResultCodeValue() bool`

HasResultCodeValue returns a boolean if a field has been set.

### GetProcessingTimeCriteria

`func (o *AddResultCriteriaRequest) GetProcessingTimeCriteria() EnumresultCriteriaProcessingTimeCriteriaProp`

GetProcessingTimeCriteria returns the ProcessingTimeCriteria field if non-nil, zero value otherwise.

### GetProcessingTimeCriteriaOk

`func (o *AddResultCriteriaRequest) GetProcessingTimeCriteriaOk() (*EnumresultCriteriaProcessingTimeCriteriaProp, bool)`

GetProcessingTimeCriteriaOk returns a tuple with the ProcessingTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeCriteria

`func (o *AddResultCriteriaRequest) SetProcessingTimeCriteria(v EnumresultCriteriaProcessingTimeCriteriaProp)`

SetProcessingTimeCriteria sets ProcessingTimeCriteria field to given value.

### HasProcessingTimeCriteria

`func (o *AddResultCriteriaRequest) HasProcessingTimeCriteria() bool`

HasProcessingTimeCriteria returns a boolean if a field has been set.

### GetProcessingTimeValue

`func (o *AddResultCriteriaRequest) GetProcessingTimeValue() string`

GetProcessingTimeValue returns the ProcessingTimeValue field if non-nil, zero value otherwise.

### GetProcessingTimeValueOk

`func (o *AddResultCriteriaRequest) GetProcessingTimeValueOk() (*string, bool)`

GetProcessingTimeValueOk returns a tuple with the ProcessingTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeValue

`func (o *AddResultCriteriaRequest) SetProcessingTimeValue(v string)`

SetProcessingTimeValue sets ProcessingTimeValue field to given value.

### HasProcessingTimeValue

`func (o *AddResultCriteriaRequest) HasProcessingTimeValue() bool`

HasProcessingTimeValue returns a boolean if a field has been set.

### GetQueueTimeCriteria

`func (o *AddResultCriteriaRequest) GetQueueTimeCriteria() EnumresultCriteriaQueueTimeCriteriaProp`

GetQueueTimeCriteria returns the QueueTimeCriteria field if non-nil, zero value otherwise.

### GetQueueTimeCriteriaOk

`func (o *AddResultCriteriaRequest) GetQueueTimeCriteriaOk() (*EnumresultCriteriaQueueTimeCriteriaProp, bool)`

GetQueueTimeCriteriaOk returns a tuple with the QueueTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeCriteria

`func (o *AddResultCriteriaRequest) SetQueueTimeCriteria(v EnumresultCriteriaQueueTimeCriteriaProp)`

SetQueueTimeCriteria sets QueueTimeCriteria field to given value.

### HasQueueTimeCriteria

`func (o *AddResultCriteriaRequest) HasQueueTimeCriteria() bool`

HasQueueTimeCriteria returns a boolean if a field has been set.

### GetQueueTimeValue

`func (o *AddResultCriteriaRequest) GetQueueTimeValue() string`

GetQueueTimeValue returns the QueueTimeValue field if non-nil, zero value otherwise.

### GetQueueTimeValueOk

`func (o *AddResultCriteriaRequest) GetQueueTimeValueOk() (*string, bool)`

GetQueueTimeValueOk returns a tuple with the QueueTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeValue

`func (o *AddResultCriteriaRequest) SetQueueTimeValue(v string)`

SetQueueTimeValue sets QueueTimeValue field to given value.

### HasQueueTimeValue

`func (o *AddResultCriteriaRequest) HasQueueTimeValue() bool`

HasQueueTimeValue returns a boolean if a field has been set.

### GetReferralReturned

`func (o *AddResultCriteriaRequest) GetReferralReturned() EnumresultCriteriaReferralReturnedProp`

GetReferralReturned returns the ReferralReturned field if non-nil, zero value otherwise.

### GetReferralReturnedOk

`func (o *AddResultCriteriaRequest) GetReferralReturnedOk() (*EnumresultCriteriaReferralReturnedProp, bool)`

GetReferralReturnedOk returns a tuple with the ReferralReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralReturned

`func (o *AddResultCriteriaRequest) SetReferralReturned(v EnumresultCriteriaReferralReturnedProp)`

SetReferralReturned sets ReferralReturned field to given value.

### HasReferralReturned

`func (o *AddResultCriteriaRequest) HasReferralReturned() bool`

HasReferralReturned returns a boolean if a field has been set.

### GetAllIncludedResponseControl

`func (o *AddResultCriteriaRequest) GetAllIncludedResponseControl() []string`

GetAllIncludedResponseControl returns the AllIncludedResponseControl field if non-nil, zero value otherwise.

### GetAllIncludedResponseControlOk

`func (o *AddResultCriteriaRequest) GetAllIncludedResponseControlOk() (*[]string, bool)`

GetAllIncludedResponseControlOk returns a tuple with the AllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResponseControl

`func (o *AddResultCriteriaRequest) SetAllIncludedResponseControl(v []string)`

SetAllIncludedResponseControl sets AllIncludedResponseControl field to given value.

### HasAllIncludedResponseControl

`func (o *AddResultCriteriaRequest) HasAllIncludedResponseControl() bool`

HasAllIncludedResponseControl returns a boolean if a field has been set.

### GetAnyIncludedResponseControl

`func (o *AddResultCriteriaRequest) GetAnyIncludedResponseControl() []string`

GetAnyIncludedResponseControl returns the AnyIncludedResponseControl field if non-nil, zero value otherwise.

### GetAnyIncludedResponseControlOk

`func (o *AddResultCriteriaRequest) GetAnyIncludedResponseControlOk() (*[]string, bool)`

GetAnyIncludedResponseControlOk returns a tuple with the AnyIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResponseControl

`func (o *AddResultCriteriaRequest) SetAnyIncludedResponseControl(v []string)`

SetAnyIncludedResponseControl sets AnyIncludedResponseControl field to given value.

### HasAnyIncludedResponseControl

`func (o *AddResultCriteriaRequest) HasAnyIncludedResponseControl() bool`

HasAnyIncludedResponseControl returns a boolean if a field has been set.

### GetNotAllIncludedResponseControl

`func (o *AddResultCriteriaRequest) GetNotAllIncludedResponseControl() []string`

GetNotAllIncludedResponseControl returns the NotAllIncludedResponseControl field if non-nil, zero value otherwise.

### GetNotAllIncludedResponseControlOk

`func (o *AddResultCriteriaRequest) GetNotAllIncludedResponseControlOk() (*[]string, bool)`

GetNotAllIncludedResponseControlOk returns a tuple with the NotAllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResponseControl

`func (o *AddResultCriteriaRequest) SetNotAllIncludedResponseControl(v []string)`

SetNotAllIncludedResponseControl sets NotAllIncludedResponseControl field to given value.

### HasNotAllIncludedResponseControl

`func (o *AddResultCriteriaRequest) HasNotAllIncludedResponseControl() bool`

HasNotAllIncludedResponseControl returns a boolean if a field has been set.

### GetNoneIncludedResponseControl

`func (o *AddResultCriteriaRequest) GetNoneIncludedResponseControl() []string`

GetNoneIncludedResponseControl returns the NoneIncludedResponseControl field if non-nil, zero value otherwise.

### GetNoneIncludedResponseControlOk

`func (o *AddResultCriteriaRequest) GetNoneIncludedResponseControlOk() (*[]string, bool)`

GetNoneIncludedResponseControlOk returns a tuple with the NoneIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResponseControl

`func (o *AddResultCriteriaRequest) SetNoneIncludedResponseControl(v []string)`

SetNoneIncludedResponseControl sets NoneIncludedResponseControl field to given value.

### HasNoneIncludedResponseControl

`func (o *AddResultCriteriaRequest) HasNoneIncludedResponseControl() bool`

HasNoneIncludedResponseControl returns a boolean if a field has been set.

### GetUsedAlternateAuthzid

`func (o *AddResultCriteriaRequest) GetUsedAlternateAuthzid() EnumresultCriteriaUsedAlternateAuthzidProp`

GetUsedAlternateAuthzid returns the UsedAlternateAuthzid field if non-nil, zero value otherwise.

### GetUsedAlternateAuthzidOk

`func (o *AddResultCriteriaRequest) GetUsedAlternateAuthzidOk() (*EnumresultCriteriaUsedAlternateAuthzidProp, bool)`

GetUsedAlternateAuthzidOk returns a tuple with the UsedAlternateAuthzid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAlternateAuthzid

`func (o *AddResultCriteriaRequest) SetUsedAlternateAuthzid(v EnumresultCriteriaUsedAlternateAuthzidProp)`

SetUsedAlternateAuthzid sets UsedAlternateAuthzid field to given value.

### HasUsedAlternateAuthzid

`func (o *AddResultCriteriaRequest) HasUsedAlternateAuthzid() bool`

HasUsedAlternateAuthzid returns a boolean if a field has been set.

### GetUsedAnyPrivilege

`func (o *AddResultCriteriaRequest) GetUsedAnyPrivilege() EnumresultCriteriaUsedAnyPrivilegeProp`

GetUsedAnyPrivilege returns the UsedAnyPrivilege field if non-nil, zero value otherwise.

### GetUsedAnyPrivilegeOk

`func (o *AddResultCriteriaRequest) GetUsedAnyPrivilegeOk() (*EnumresultCriteriaUsedAnyPrivilegeProp, bool)`

GetUsedAnyPrivilegeOk returns a tuple with the UsedAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAnyPrivilege

`func (o *AddResultCriteriaRequest) SetUsedAnyPrivilege(v EnumresultCriteriaUsedAnyPrivilegeProp)`

SetUsedAnyPrivilege sets UsedAnyPrivilege field to given value.

### HasUsedAnyPrivilege

`func (o *AddResultCriteriaRequest) HasUsedAnyPrivilege() bool`

HasUsedAnyPrivilege returns a boolean if a field has been set.

### GetUsedPrivilege

`func (o *AddResultCriteriaRequest) GetUsedPrivilege() []EnumresultCriteriaUsedPrivilegeProp`

GetUsedPrivilege returns the UsedPrivilege field if non-nil, zero value otherwise.

### GetUsedPrivilegeOk

`func (o *AddResultCriteriaRequest) GetUsedPrivilegeOk() (*[]EnumresultCriteriaUsedPrivilegeProp, bool)`

GetUsedPrivilegeOk returns a tuple with the UsedPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPrivilege

`func (o *AddResultCriteriaRequest) SetUsedPrivilege(v []EnumresultCriteriaUsedPrivilegeProp)`

SetUsedPrivilege sets UsedPrivilege field to given value.

### HasUsedPrivilege

`func (o *AddResultCriteriaRequest) HasUsedPrivilege() bool`

HasUsedPrivilege returns a boolean if a field has been set.

### GetMissingAnyPrivilege

`func (o *AddResultCriteriaRequest) GetMissingAnyPrivilege() EnumresultCriteriaMissingAnyPrivilegeProp`

GetMissingAnyPrivilege returns the MissingAnyPrivilege field if non-nil, zero value otherwise.

### GetMissingAnyPrivilegeOk

`func (o *AddResultCriteriaRequest) GetMissingAnyPrivilegeOk() (*EnumresultCriteriaMissingAnyPrivilegeProp, bool)`

GetMissingAnyPrivilegeOk returns a tuple with the MissingAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAnyPrivilege

`func (o *AddResultCriteriaRequest) SetMissingAnyPrivilege(v EnumresultCriteriaMissingAnyPrivilegeProp)`

SetMissingAnyPrivilege sets MissingAnyPrivilege field to given value.

### HasMissingAnyPrivilege

`func (o *AddResultCriteriaRequest) HasMissingAnyPrivilege() bool`

HasMissingAnyPrivilege returns a boolean if a field has been set.

### GetMissingPrivilege

`func (o *AddResultCriteriaRequest) GetMissingPrivilege() []EnumresultCriteriaMissingPrivilegeProp`

GetMissingPrivilege returns the MissingPrivilege field if non-nil, zero value otherwise.

### GetMissingPrivilegeOk

`func (o *AddResultCriteriaRequest) GetMissingPrivilegeOk() (*[]EnumresultCriteriaMissingPrivilegeProp, bool)`

GetMissingPrivilegeOk returns a tuple with the MissingPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPrivilege

`func (o *AddResultCriteriaRequest) SetMissingPrivilege(v []EnumresultCriteriaMissingPrivilegeProp)`

SetMissingPrivilege sets MissingPrivilege field to given value.

### HasMissingPrivilege

`func (o *AddResultCriteriaRequest) HasMissingPrivilege() bool`

HasMissingPrivilege returns a boolean if a field has been set.

### GetRetiredPasswordUsedForBind

`func (o *AddResultCriteriaRequest) GetRetiredPasswordUsedForBind() EnumresultCriteriaRetiredPasswordUsedForBindProp`

GetRetiredPasswordUsedForBind returns the RetiredPasswordUsedForBind field if non-nil, zero value otherwise.

### GetRetiredPasswordUsedForBindOk

`func (o *AddResultCriteriaRequest) GetRetiredPasswordUsedForBindOk() (*EnumresultCriteriaRetiredPasswordUsedForBindProp, bool)`

GetRetiredPasswordUsedForBindOk returns a tuple with the RetiredPasswordUsedForBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiredPasswordUsedForBind

`func (o *AddResultCriteriaRequest) SetRetiredPasswordUsedForBind(v EnumresultCriteriaRetiredPasswordUsedForBindProp)`

SetRetiredPasswordUsedForBind sets RetiredPasswordUsedForBind field to given value.

### HasRetiredPasswordUsedForBind

`func (o *AddResultCriteriaRequest) HasRetiredPasswordUsedForBind() bool`

HasRetiredPasswordUsedForBind returns a boolean if a field has been set.

### GetSearchEntryReturnedCriteria

`func (o *AddResultCriteriaRequest) GetSearchEntryReturnedCriteria() EnumresultCriteriaSearchEntryReturnedCriteriaProp`

GetSearchEntryReturnedCriteria returns the SearchEntryReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCriteriaOk

`func (o *AddResultCriteriaRequest) GetSearchEntryReturnedCriteriaOk() (*EnumresultCriteriaSearchEntryReturnedCriteriaProp, bool)`

GetSearchEntryReturnedCriteriaOk returns a tuple with the SearchEntryReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCriteria

`func (o *AddResultCriteriaRequest) SetSearchEntryReturnedCriteria(v EnumresultCriteriaSearchEntryReturnedCriteriaProp)`

SetSearchEntryReturnedCriteria sets SearchEntryReturnedCriteria field to given value.

### HasSearchEntryReturnedCriteria

`func (o *AddResultCriteriaRequest) HasSearchEntryReturnedCriteria() bool`

HasSearchEntryReturnedCriteria returns a boolean if a field has been set.

### GetSearchEntryReturnedCount

`func (o *AddResultCriteriaRequest) GetSearchEntryReturnedCount() int64`

GetSearchEntryReturnedCount returns the SearchEntryReturnedCount field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCountOk

`func (o *AddResultCriteriaRequest) GetSearchEntryReturnedCountOk() (*int64, bool)`

GetSearchEntryReturnedCountOk returns a tuple with the SearchEntryReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCount

`func (o *AddResultCriteriaRequest) SetSearchEntryReturnedCount(v int64)`

SetSearchEntryReturnedCount sets SearchEntryReturnedCount field to given value.

### HasSearchEntryReturnedCount

`func (o *AddResultCriteriaRequest) HasSearchEntryReturnedCount() bool`

HasSearchEntryReturnedCount returns a boolean if a field has been set.

### GetSearchReferenceReturnedCriteria

`func (o *AddResultCriteriaRequest) GetSearchReferenceReturnedCriteria() EnumresultCriteriaSearchReferenceReturnedCriteriaProp`

GetSearchReferenceReturnedCriteria returns the SearchReferenceReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCriteriaOk

`func (o *AddResultCriteriaRequest) GetSearchReferenceReturnedCriteriaOk() (*EnumresultCriteriaSearchReferenceReturnedCriteriaProp, bool)`

GetSearchReferenceReturnedCriteriaOk returns a tuple with the SearchReferenceReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCriteria

`func (o *AddResultCriteriaRequest) SetSearchReferenceReturnedCriteria(v EnumresultCriteriaSearchReferenceReturnedCriteriaProp)`

SetSearchReferenceReturnedCriteria sets SearchReferenceReturnedCriteria field to given value.

### HasSearchReferenceReturnedCriteria

`func (o *AddResultCriteriaRequest) HasSearchReferenceReturnedCriteria() bool`

HasSearchReferenceReturnedCriteria returns a boolean if a field has been set.

### GetSearchReferenceReturnedCount

`func (o *AddResultCriteriaRequest) GetSearchReferenceReturnedCount() int64`

GetSearchReferenceReturnedCount returns the SearchReferenceReturnedCount field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCountOk

`func (o *AddResultCriteriaRequest) GetSearchReferenceReturnedCountOk() (*int64, bool)`

GetSearchReferenceReturnedCountOk returns a tuple with the SearchReferenceReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCount

`func (o *AddResultCriteriaRequest) SetSearchReferenceReturnedCount(v int64)`

SetSearchReferenceReturnedCount sets SearchReferenceReturnedCount field to given value.

### HasSearchReferenceReturnedCount

`func (o *AddResultCriteriaRequest) HasSearchReferenceReturnedCount() bool`

HasSearchReferenceReturnedCount returns a boolean if a field has been set.

### GetSearchIndexedCriteria

`func (o *AddResultCriteriaRequest) GetSearchIndexedCriteria() EnumresultCriteriaSearchIndexedCriteriaProp`

GetSearchIndexedCriteria returns the SearchIndexedCriteria field if non-nil, zero value otherwise.

### GetSearchIndexedCriteriaOk

`func (o *AddResultCriteriaRequest) GetSearchIndexedCriteriaOk() (*EnumresultCriteriaSearchIndexedCriteriaProp, bool)`

GetSearchIndexedCriteriaOk returns a tuple with the SearchIndexedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexedCriteria

`func (o *AddResultCriteriaRequest) SetSearchIndexedCriteria(v EnumresultCriteriaSearchIndexedCriteriaProp)`

SetSearchIndexedCriteria sets SearchIndexedCriteria field to given value.

### HasSearchIndexedCriteria

`func (o *AddResultCriteriaRequest) HasSearchIndexedCriteria() bool`

HasSearchIndexedCriteria returns a boolean if a field has been set.

### GetIncludedAuthzUserBaseDN

`func (o *AddResultCriteriaRequest) GetIncludedAuthzUserBaseDN() []string`

GetIncludedAuthzUserBaseDN returns the IncludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedAuthzUserBaseDNOk

`func (o *AddResultCriteriaRequest) GetIncludedAuthzUserBaseDNOk() (*[]string, bool)`

GetIncludedAuthzUserBaseDNOk returns a tuple with the IncludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAuthzUserBaseDN

`func (o *AddResultCriteriaRequest) SetIncludedAuthzUserBaseDN(v []string)`

SetIncludedAuthzUserBaseDN sets IncludedAuthzUserBaseDN field to given value.

### HasIncludedAuthzUserBaseDN

`func (o *AddResultCriteriaRequest) HasIncludedAuthzUserBaseDN() bool`

HasIncludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetExcludedAuthzUserBaseDN

`func (o *AddResultCriteriaRequest) GetExcludedAuthzUserBaseDN() []string`

GetExcludedAuthzUserBaseDN returns the ExcludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedAuthzUserBaseDNOk

`func (o *AddResultCriteriaRequest) GetExcludedAuthzUserBaseDNOk() (*[]string, bool)`

GetExcludedAuthzUserBaseDNOk returns a tuple with the ExcludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAuthzUserBaseDN

`func (o *AddResultCriteriaRequest) SetExcludedAuthzUserBaseDN(v []string)`

SetExcludedAuthzUserBaseDN sets ExcludedAuthzUserBaseDN field to given value.

### HasExcludedAuthzUserBaseDN

`func (o *AddResultCriteriaRequest) HasExcludedAuthzUserBaseDN() bool`

HasExcludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) GetAllIncludedAuthzUserGroupDN() []string`

GetAllIncludedAuthzUserGroupDN returns the AllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteriaRequest) GetAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAllIncludedAuthzUserGroupDNOk returns a tuple with the AllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) SetAllIncludedAuthzUserGroupDN(v []string)`

SetAllIncludedAuthzUserGroupDN sets AllIncludedAuthzUserGroupDN field to given value.

### HasAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) HasAllIncludedAuthzUserGroupDN() bool`

HasAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) GetAnyIncludedAuthzUserGroupDN() []string`

GetAnyIncludedAuthzUserGroupDN returns the AnyIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteriaRequest) GetAnyIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedAuthzUserGroupDNOk returns a tuple with the AnyIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) SetAnyIncludedAuthzUserGroupDN(v []string)`

SetAnyIncludedAuthzUserGroupDN sets AnyIncludedAuthzUserGroupDN field to given value.

### HasAnyIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) HasAnyIncludedAuthzUserGroupDN() bool`

HasAnyIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) GetNotAllIncludedAuthzUserGroupDN() []string`

GetNotAllIncludedAuthzUserGroupDN returns the NotAllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteriaRequest) GetNotAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedAuthzUserGroupDNOk returns a tuple with the NotAllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) SetNotAllIncludedAuthzUserGroupDN(v []string)`

SetNotAllIncludedAuthzUserGroupDN sets NotAllIncludedAuthzUserGroupDN field to given value.

### HasNotAllIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) HasNotAllIncludedAuthzUserGroupDN() bool`

HasNotAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) GetNoneIncludedAuthzUserGroupDN() []string`

GetNoneIncludedAuthzUserGroupDN returns the NoneIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedAuthzUserGroupDNOk

`func (o *AddResultCriteriaRequest) GetNoneIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedAuthzUserGroupDNOk returns a tuple with the NoneIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) SetNoneIncludedAuthzUserGroupDN(v []string)`

SetNoneIncludedAuthzUserGroupDN sets NoneIncludedAuthzUserGroupDN field to given value.

### HasNoneIncludedAuthzUserGroupDN

`func (o *AddResultCriteriaRequest) HasNoneIncludedAuthzUserGroupDN() bool`

HasNoneIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetAllIncludedResultCriteria

`func (o *AddResultCriteriaRequest) GetAllIncludedResultCriteria() []string`

GetAllIncludedResultCriteria returns the AllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAllIncludedResultCriteriaOk

`func (o *AddResultCriteriaRequest) GetAllIncludedResultCriteriaOk() (*[]string, bool)`

GetAllIncludedResultCriteriaOk returns a tuple with the AllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResultCriteria

`func (o *AddResultCriteriaRequest) SetAllIncludedResultCriteria(v []string)`

SetAllIncludedResultCriteria sets AllIncludedResultCriteria field to given value.

### HasAllIncludedResultCriteria

`func (o *AddResultCriteriaRequest) HasAllIncludedResultCriteria() bool`

HasAllIncludedResultCriteria returns a boolean if a field has been set.

### GetAnyIncludedResultCriteria

`func (o *AddResultCriteriaRequest) GetAnyIncludedResultCriteria() []string`

GetAnyIncludedResultCriteria returns the AnyIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedResultCriteriaOk

`func (o *AddResultCriteriaRequest) GetAnyIncludedResultCriteriaOk() (*[]string, bool)`

GetAnyIncludedResultCriteriaOk returns a tuple with the AnyIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResultCriteria

`func (o *AddResultCriteriaRequest) SetAnyIncludedResultCriteria(v []string)`

SetAnyIncludedResultCriteria sets AnyIncludedResultCriteria field to given value.

### HasAnyIncludedResultCriteria

`func (o *AddResultCriteriaRequest) HasAnyIncludedResultCriteria() bool`

HasAnyIncludedResultCriteria returns a boolean if a field has been set.

### GetNotAllIncludedResultCriteria

`func (o *AddResultCriteriaRequest) GetNotAllIncludedResultCriteria() []string`

GetNotAllIncludedResultCriteria returns the NotAllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedResultCriteriaOk

`func (o *AddResultCriteriaRequest) GetNotAllIncludedResultCriteriaOk() (*[]string, bool)`

GetNotAllIncludedResultCriteriaOk returns a tuple with the NotAllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResultCriteria

`func (o *AddResultCriteriaRequest) SetNotAllIncludedResultCriteria(v []string)`

SetNotAllIncludedResultCriteria sets NotAllIncludedResultCriteria field to given value.

### HasNotAllIncludedResultCriteria

`func (o *AddResultCriteriaRequest) HasNotAllIncludedResultCriteria() bool`

HasNotAllIncludedResultCriteria returns a boolean if a field has been set.

### GetNoneIncludedResultCriteria

`func (o *AddResultCriteriaRequest) GetNoneIncludedResultCriteria() []string`

GetNoneIncludedResultCriteria returns the NoneIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedResultCriteriaOk

`func (o *AddResultCriteriaRequest) GetNoneIncludedResultCriteriaOk() (*[]string, bool)`

GetNoneIncludedResultCriteriaOk returns a tuple with the NoneIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResultCriteria

`func (o *AddResultCriteriaRequest) SetNoneIncludedResultCriteria(v []string)`

SetNoneIncludedResultCriteria sets NoneIncludedResultCriteria field to given value.

### HasNoneIncludedResultCriteria

`func (o *AddResultCriteriaRequest) HasNoneIncludedResultCriteria() bool`

HasNoneIncludedResultCriteria returns a boolean if a field has been set.

### GetLocalAssuranceLevel

`func (o *AddResultCriteriaRequest) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp`

GetLocalAssuranceLevel returns the LocalAssuranceLevel field if non-nil, zero value otherwise.

### GetLocalAssuranceLevelOk

`func (o *AddResultCriteriaRequest) GetLocalAssuranceLevelOk() (*[]EnumresultCriteriaLocalAssuranceLevelProp, bool)`

GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAssuranceLevel

`func (o *AddResultCriteriaRequest) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp)`

SetLocalAssuranceLevel sets LocalAssuranceLevel field to given value.

### HasLocalAssuranceLevel

`func (o *AddResultCriteriaRequest) HasLocalAssuranceLevel() bool`

HasLocalAssuranceLevel returns a boolean if a field has been set.

### GetRemoteAssuranceLevel

`func (o *AddResultCriteriaRequest) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp`

GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field if non-nil, zero value otherwise.

### GetRemoteAssuranceLevelOk

`func (o *AddResultCriteriaRequest) GetRemoteAssuranceLevelOk() (*[]EnumresultCriteriaRemoteAssuranceLevelProp, bool)`

GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssuranceLevel

`func (o *AddResultCriteriaRequest) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp)`

SetRemoteAssuranceLevel sets RemoteAssuranceLevel field to given value.

### HasRemoteAssuranceLevel

`func (o *AddResultCriteriaRequest) HasRemoteAssuranceLevel() bool`

HasRemoteAssuranceLevel returns a boolean if a field has been set.

### GetAssuranceTimeoutCriteria

`func (o *AddResultCriteriaRequest) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp`

GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field if non-nil, zero value otherwise.

### GetAssuranceTimeoutCriteriaOk

`func (o *AddResultCriteriaRequest) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool)`

GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutCriteria

`func (o *AddResultCriteriaRequest) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp)`

SetAssuranceTimeoutCriteria sets AssuranceTimeoutCriteria field to given value.

### HasAssuranceTimeoutCriteria

`func (o *AddResultCriteriaRequest) HasAssuranceTimeoutCriteria() bool`

HasAssuranceTimeoutCriteria returns a boolean if a field has been set.

### GetAssuranceTimeoutValue

`func (o *AddResultCriteriaRequest) GetAssuranceTimeoutValue() string`

GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field if non-nil, zero value otherwise.

### GetAssuranceTimeoutValueOk

`func (o *AddResultCriteriaRequest) GetAssuranceTimeoutValueOk() (*string, bool)`

GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutValue

`func (o *AddResultCriteriaRequest) SetAssuranceTimeoutValue(v string)`

SetAssuranceTimeoutValue sets AssuranceTimeoutValue field to given value.

### HasAssuranceTimeoutValue

`func (o *AddResultCriteriaRequest) HasAssuranceTimeoutValue() bool`

HasAssuranceTimeoutValue returns a boolean if a field has been set.

### GetResponseDelayedByAssurance

`func (o *AddResultCriteriaRequest) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp`

GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field if non-nil, zero value otherwise.

### GetResponseDelayedByAssuranceOk

`func (o *AddResultCriteriaRequest) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool)`

GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelayedByAssurance

`func (o *AddResultCriteriaRequest) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp)`

SetResponseDelayedByAssurance sets ResponseDelayedByAssurance field to given value.

### HasResponseDelayedByAssurance

`func (o *AddResultCriteriaRequest) HasResponseDelayedByAssurance() bool`

HasResponseDelayedByAssurance returns a boolean if a field has been set.

### GetAssuranceBehaviorAlteredByControl

`func (o *AddResultCriteriaRequest) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp`

GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field if non-nil, zero value otherwise.

### GetAssuranceBehaviorAlteredByControlOk

`func (o *AddResultCriteriaRequest) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool)`

GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceBehaviorAlteredByControl

`func (o *AddResultCriteriaRequest) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp)`

SetAssuranceBehaviorAlteredByControl sets AssuranceBehaviorAlteredByControl field to given value.

### HasAssuranceBehaviorAlteredByControl

`func (o *AddResultCriteriaRequest) HasAssuranceBehaviorAlteredByControl() bool`

HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.

### GetAssuranceSatisfied

`func (o *AddResultCriteriaRequest) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp`

GetAssuranceSatisfied returns the AssuranceSatisfied field if non-nil, zero value otherwise.

### GetAssuranceSatisfiedOk

`func (o *AddResultCriteriaRequest) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool)`

GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceSatisfied

`func (o *AddResultCriteriaRequest) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp)`

SetAssuranceSatisfied sets AssuranceSatisfied field to given value.

### HasAssuranceSatisfied

`func (o *AddResultCriteriaRequest) HasAssuranceSatisfied() bool`

HasAssuranceSatisfied returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddResultCriteriaRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddResultCriteriaRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddResultCriteriaRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddResultCriteriaRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddResultCriteriaRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddResultCriteriaRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddResultCriteriaRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


