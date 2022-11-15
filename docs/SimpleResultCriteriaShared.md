# SimpleResultCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsimpleResultCriteriaSchemaUrn**](EnumsimpleResultCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for operations included in this Simple Result Criteria. | [optional] 
**ResultCodeCriteria** | Pointer to [**EnumresultCriteriaResultCodeCriteriaProp**](EnumresultCriteriaResultCodeCriteriaProp.md) |  | [optional] 
**ResultCodeValue** | Pointer to [**[]EnumresultCriteriaResultCodeValueProp**](EnumresultCriteriaResultCodeValueProp.md) |  | [optional] 
**ProcessingTimeCriteria** | Pointer to [**EnumresultCriteriaProcessingTimeCriteriaProp**](EnumresultCriteriaProcessingTimeCriteriaProp.md) |  | [optional] 
**ProcessingTimeValue** | Pointer to **string** | Specifies the boundary value to use for the operation processing time when determining whether to include that operation in this Simple Result Criteria. This will be ignored if the \&quot;processing-time-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**QueueTimeCriteria** | Pointer to [**EnumresultCriteriaQueueTimeCriteriaProp**](EnumresultCriteriaQueueTimeCriteriaProp.md) |  | [optional] 
**QueueTimeValue** | Pointer to **string** | Specifies the boundary value to use for the time an operation spent on the work queue when determining whether to include that operation in this Simple Result Criteria. This will be ignored if the \&quot;queue-time-criteria\&quot; property has a value of \&quot;any\&quot;. | [optional] 
**ReferralReturned** | Pointer to [**EnumresultCriteriaReferralReturnedProp**](EnumresultCriteriaReferralReturnedProp.md) |  | [optional] 
**AllIncludedResponseControl** | Pointer to **[]string** |  | [optional] 
**AnyIncludedResponseControl** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedResponseControl** | Pointer to **[]string** |  | [optional] 
**NoneIncludedResponseControl** | Pointer to **[]string** |  | [optional] 
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
**IncludedAuthzUserBaseDN** | Pointer to **[]string** |  | [optional] 
**ExcludedAuthzUserBaseDN** | Pointer to **[]string** |  | [optional] 
**AllIncludedAuthzUserGroupDN** | Pointer to **[]string** |  | [optional] 
**AnyIncludedAuthzUserGroupDN** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedAuthzUserGroupDN** | Pointer to **[]string** |  | [optional] 
**NoneIncludedAuthzUserGroupDN** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 

## Methods

### NewSimpleResultCriteriaShared

`func NewSimpleResultCriteriaShared(schemas []EnumsimpleResultCriteriaSchemaUrn, ) *SimpleResultCriteriaShared`

NewSimpleResultCriteriaShared instantiates a new SimpleResultCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleResultCriteriaSharedWithDefaults

`func NewSimpleResultCriteriaSharedWithDefaults() *SimpleResultCriteriaShared`

NewSimpleResultCriteriaSharedWithDefaults instantiates a new SimpleResultCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SimpleResultCriteriaShared) GetSchemas() []EnumsimpleResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleResultCriteriaShared) GetSchemasOk() (*[]EnumsimpleResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleResultCriteriaShared) SetSchemas(v []EnumsimpleResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *SimpleResultCriteriaShared) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *SimpleResultCriteriaShared) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *SimpleResultCriteriaShared) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *SimpleResultCriteriaShared) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetResultCodeCriteria

`func (o *SimpleResultCriteriaShared) GetResultCodeCriteria() EnumresultCriteriaResultCodeCriteriaProp`

GetResultCodeCriteria returns the ResultCodeCriteria field if non-nil, zero value otherwise.

### GetResultCodeCriteriaOk

`func (o *SimpleResultCriteriaShared) GetResultCodeCriteriaOk() (*EnumresultCriteriaResultCodeCriteriaProp, bool)`

GetResultCodeCriteriaOk returns a tuple with the ResultCodeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeCriteria

`func (o *SimpleResultCriteriaShared) SetResultCodeCriteria(v EnumresultCriteriaResultCodeCriteriaProp)`

SetResultCodeCriteria sets ResultCodeCriteria field to given value.

### HasResultCodeCriteria

`func (o *SimpleResultCriteriaShared) HasResultCodeCriteria() bool`

HasResultCodeCriteria returns a boolean if a field has been set.

### GetResultCodeValue

`func (o *SimpleResultCriteriaShared) GetResultCodeValue() []EnumresultCriteriaResultCodeValueProp`

GetResultCodeValue returns the ResultCodeValue field if non-nil, zero value otherwise.

### GetResultCodeValueOk

`func (o *SimpleResultCriteriaShared) GetResultCodeValueOk() (*[]EnumresultCriteriaResultCodeValueProp, bool)`

GetResultCodeValueOk returns a tuple with the ResultCodeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeValue

`func (o *SimpleResultCriteriaShared) SetResultCodeValue(v []EnumresultCriteriaResultCodeValueProp)`

SetResultCodeValue sets ResultCodeValue field to given value.

### HasResultCodeValue

`func (o *SimpleResultCriteriaShared) HasResultCodeValue() bool`

HasResultCodeValue returns a boolean if a field has been set.

### GetProcessingTimeCriteria

`func (o *SimpleResultCriteriaShared) GetProcessingTimeCriteria() EnumresultCriteriaProcessingTimeCriteriaProp`

GetProcessingTimeCriteria returns the ProcessingTimeCriteria field if non-nil, zero value otherwise.

### GetProcessingTimeCriteriaOk

`func (o *SimpleResultCriteriaShared) GetProcessingTimeCriteriaOk() (*EnumresultCriteriaProcessingTimeCriteriaProp, bool)`

GetProcessingTimeCriteriaOk returns a tuple with the ProcessingTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeCriteria

`func (o *SimpleResultCriteriaShared) SetProcessingTimeCriteria(v EnumresultCriteriaProcessingTimeCriteriaProp)`

SetProcessingTimeCriteria sets ProcessingTimeCriteria field to given value.

### HasProcessingTimeCriteria

`func (o *SimpleResultCriteriaShared) HasProcessingTimeCriteria() bool`

HasProcessingTimeCriteria returns a boolean if a field has been set.

### GetProcessingTimeValue

`func (o *SimpleResultCriteriaShared) GetProcessingTimeValue() string`

GetProcessingTimeValue returns the ProcessingTimeValue field if non-nil, zero value otherwise.

### GetProcessingTimeValueOk

`func (o *SimpleResultCriteriaShared) GetProcessingTimeValueOk() (*string, bool)`

GetProcessingTimeValueOk returns a tuple with the ProcessingTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeValue

`func (o *SimpleResultCriteriaShared) SetProcessingTimeValue(v string)`

SetProcessingTimeValue sets ProcessingTimeValue field to given value.

### HasProcessingTimeValue

`func (o *SimpleResultCriteriaShared) HasProcessingTimeValue() bool`

HasProcessingTimeValue returns a boolean if a field has been set.

### GetQueueTimeCriteria

`func (o *SimpleResultCriteriaShared) GetQueueTimeCriteria() EnumresultCriteriaQueueTimeCriteriaProp`

GetQueueTimeCriteria returns the QueueTimeCriteria field if non-nil, zero value otherwise.

### GetQueueTimeCriteriaOk

`func (o *SimpleResultCriteriaShared) GetQueueTimeCriteriaOk() (*EnumresultCriteriaQueueTimeCriteriaProp, bool)`

GetQueueTimeCriteriaOk returns a tuple with the QueueTimeCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeCriteria

`func (o *SimpleResultCriteriaShared) SetQueueTimeCriteria(v EnumresultCriteriaQueueTimeCriteriaProp)`

SetQueueTimeCriteria sets QueueTimeCriteria field to given value.

### HasQueueTimeCriteria

`func (o *SimpleResultCriteriaShared) HasQueueTimeCriteria() bool`

HasQueueTimeCriteria returns a boolean if a field has been set.

### GetQueueTimeValue

`func (o *SimpleResultCriteriaShared) GetQueueTimeValue() string`

GetQueueTimeValue returns the QueueTimeValue field if non-nil, zero value otherwise.

### GetQueueTimeValueOk

`func (o *SimpleResultCriteriaShared) GetQueueTimeValueOk() (*string, bool)`

GetQueueTimeValueOk returns a tuple with the QueueTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTimeValue

`func (o *SimpleResultCriteriaShared) SetQueueTimeValue(v string)`

SetQueueTimeValue sets QueueTimeValue field to given value.

### HasQueueTimeValue

`func (o *SimpleResultCriteriaShared) HasQueueTimeValue() bool`

HasQueueTimeValue returns a boolean if a field has been set.

### GetReferralReturned

`func (o *SimpleResultCriteriaShared) GetReferralReturned() EnumresultCriteriaReferralReturnedProp`

GetReferralReturned returns the ReferralReturned field if non-nil, zero value otherwise.

### GetReferralReturnedOk

`func (o *SimpleResultCriteriaShared) GetReferralReturnedOk() (*EnumresultCriteriaReferralReturnedProp, bool)`

GetReferralReturnedOk returns a tuple with the ReferralReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralReturned

`func (o *SimpleResultCriteriaShared) SetReferralReturned(v EnumresultCriteriaReferralReturnedProp)`

SetReferralReturned sets ReferralReturned field to given value.

### HasReferralReturned

`func (o *SimpleResultCriteriaShared) HasReferralReturned() bool`

HasReferralReturned returns a boolean if a field has been set.

### GetAllIncludedResponseControl

`func (o *SimpleResultCriteriaShared) GetAllIncludedResponseControl() []string`

GetAllIncludedResponseControl returns the AllIncludedResponseControl field if non-nil, zero value otherwise.

### GetAllIncludedResponseControlOk

`func (o *SimpleResultCriteriaShared) GetAllIncludedResponseControlOk() (*[]string, bool)`

GetAllIncludedResponseControlOk returns a tuple with the AllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResponseControl

`func (o *SimpleResultCriteriaShared) SetAllIncludedResponseControl(v []string)`

SetAllIncludedResponseControl sets AllIncludedResponseControl field to given value.

### HasAllIncludedResponseControl

`func (o *SimpleResultCriteriaShared) HasAllIncludedResponseControl() bool`

HasAllIncludedResponseControl returns a boolean if a field has been set.

### GetAnyIncludedResponseControl

`func (o *SimpleResultCriteriaShared) GetAnyIncludedResponseControl() []string`

GetAnyIncludedResponseControl returns the AnyIncludedResponseControl field if non-nil, zero value otherwise.

### GetAnyIncludedResponseControlOk

`func (o *SimpleResultCriteriaShared) GetAnyIncludedResponseControlOk() (*[]string, bool)`

GetAnyIncludedResponseControlOk returns a tuple with the AnyIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResponseControl

`func (o *SimpleResultCriteriaShared) SetAnyIncludedResponseControl(v []string)`

SetAnyIncludedResponseControl sets AnyIncludedResponseControl field to given value.

### HasAnyIncludedResponseControl

`func (o *SimpleResultCriteriaShared) HasAnyIncludedResponseControl() bool`

HasAnyIncludedResponseControl returns a boolean if a field has been set.

### GetNotAllIncludedResponseControl

`func (o *SimpleResultCriteriaShared) GetNotAllIncludedResponseControl() []string`

GetNotAllIncludedResponseControl returns the NotAllIncludedResponseControl field if non-nil, zero value otherwise.

### GetNotAllIncludedResponseControlOk

`func (o *SimpleResultCriteriaShared) GetNotAllIncludedResponseControlOk() (*[]string, bool)`

GetNotAllIncludedResponseControlOk returns a tuple with the NotAllIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResponseControl

`func (o *SimpleResultCriteriaShared) SetNotAllIncludedResponseControl(v []string)`

SetNotAllIncludedResponseControl sets NotAllIncludedResponseControl field to given value.

### HasNotAllIncludedResponseControl

`func (o *SimpleResultCriteriaShared) HasNotAllIncludedResponseControl() bool`

HasNotAllIncludedResponseControl returns a boolean if a field has been set.

### GetNoneIncludedResponseControl

`func (o *SimpleResultCriteriaShared) GetNoneIncludedResponseControl() []string`

GetNoneIncludedResponseControl returns the NoneIncludedResponseControl field if non-nil, zero value otherwise.

### GetNoneIncludedResponseControlOk

`func (o *SimpleResultCriteriaShared) GetNoneIncludedResponseControlOk() (*[]string, bool)`

GetNoneIncludedResponseControlOk returns a tuple with the NoneIncludedResponseControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResponseControl

`func (o *SimpleResultCriteriaShared) SetNoneIncludedResponseControl(v []string)`

SetNoneIncludedResponseControl sets NoneIncludedResponseControl field to given value.

### HasNoneIncludedResponseControl

`func (o *SimpleResultCriteriaShared) HasNoneIncludedResponseControl() bool`

HasNoneIncludedResponseControl returns a boolean if a field has been set.

### GetUsedAlternateAuthzid

`func (o *SimpleResultCriteriaShared) GetUsedAlternateAuthzid() EnumresultCriteriaUsedAlternateAuthzidProp`

GetUsedAlternateAuthzid returns the UsedAlternateAuthzid field if non-nil, zero value otherwise.

### GetUsedAlternateAuthzidOk

`func (o *SimpleResultCriteriaShared) GetUsedAlternateAuthzidOk() (*EnumresultCriteriaUsedAlternateAuthzidProp, bool)`

GetUsedAlternateAuthzidOk returns a tuple with the UsedAlternateAuthzid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAlternateAuthzid

`func (o *SimpleResultCriteriaShared) SetUsedAlternateAuthzid(v EnumresultCriteriaUsedAlternateAuthzidProp)`

SetUsedAlternateAuthzid sets UsedAlternateAuthzid field to given value.

### HasUsedAlternateAuthzid

`func (o *SimpleResultCriteriaShared) HasUsedAlternateAuthzid() bool`

HasUsedAlternateAuthzid returns a boolean if a field has been set.

### GetUsedAnyPrivilege

`func (o *SimpleResultCriteriaShared) GetUsedAnyPrivilege() EnumresultCriteriaUsedAnyPrivilegeProp`

GetUsedAnyPrivilege returns the UsedAnyPrivilege field if non-nil, zero value otherwise.

### GetUsedAnyPrivilegeOk

`func (o *SimpleResultCriteriaShared) GetUsedAnyPrivilegeOk() (*EnumresultCriteriaUsedAnyPrivilegeProp, bool)`

GetUsedAnyPrivilegeOk returns a tuple with the UsedAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAnyPrivilege

`func (o *SimpleResultCriteriaShared) SetUsedAnyPrivilege(v EnumresultCriteriaUsedAnyPrivilegeProp)`

SetUsedAnyPrivilege sets UsedAnyPrivilege field to given value.

### HasUsedAnyPrivilege

`func (o *SimpleResultCriteriaShared) HasUsedAnyPrivilege() bool`

HasUsedAnyPrivilege returns a boolean if a field has been set.

### GetUsedPrivilege

`func (o *SimpleResultCriteriaShared) GetUsedPrivilege() []EnumresultCriteriaUsedPrivilegeProp`

GetUsedPrivilege returns the UsedPrivilege field if non-nil, zero value otherwise.

### GetUsedPrivilegeOk

`func (o *SimpleResultCriteriaShared) GetUsedPrivilegeOk() (*[]EnumresultCriteriaUsedPrivilegeProp, bool)`

GetUsedPrivilegeOk returns a tuple with the UsedPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPrivilege

`func (o *SimpleResultCriteriaShared) SetUsedPrivilege(v []EnumresultCriteriaUsedPrivilegeProp)`

SetUsedPrivilege sets UsedPrivilege field to given value.

### HasUsedPrivilege

`func (o *SimpleResultCriteriaShared) HasUsedPrivilege() bool`

HasUsedPrivilege returns a boolean if a field has been set.

### GetMissingAnyPrivilege

`func (o *SimpleResultCriteriaShared) GetMissingAnyPrivilege() EnumresultCriteriaMissingAnyPrivilegeProp`

GetMissingAnyPrivilege returns the MissingAnyPrivilege field if non-nil, zero value otherwise.

### GetMissingAnyPrivilegeOk

`func (o *SimpleResultCriteriaShared) GetMissingAnyPrivilegeOk() (*EnumresultCriteriaMissingAnyPrivilegeProp, bool)`

GetMissingAnyPrivilegeOk returns a tuple with the MissingAnyPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAnyPrivilege

`func (o *SimpleResultCriteriaShared) SetMissingAnyPrivilege(v EnumresultCriteriaMissingAnyPrivilegeProp)`

SetMissingAnyPrivilege sets MissingAnyPrivilege field to given value.

### HasMissingAnyPrivilege

`func (o *SimpleResultCriteriaShared) HasMissingAnyPrivilege() bool`

HasMissingAnyPrivilege returns a boolean if a field has been set.

### GetMissingPrivilege

`func (o *SimpleResultCriteriaShared) GetMissingPrivilege() []EnumresultCriteriaMissingPrivilegeProp`

GetMissingPrivilege returns the MissingPrivilege field if non-nil, zero value otherwise.

### GetMissingPrivilegeOk

`func (o *SimpleResultCriteriaShared) GetMissingPrivilegeOk() (*[]EnumresultCriteriaMissingPrivilegeProp, bool)`

GetMissingPrivilegeOk returns a tuple with the MissingPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPrivilege

`func (o *SimpleResultCriteriaShared) SetMissingPrivilege(v []EnumresultCriteriaMissingPrivilegeProp)`

SetMissingPrivilege sets MissingPrivilege field to given value.

### HasMissingPrivilege

`func (o *SimpleResultCriteriaShared) HasMissingPrivilege() bool`

HasMissingPrivilege returns a boolean if a field has been set.

### GetRetiredPasswordUsedForBind

`func (o *SimpleResultCriteriaShared) GetRetiredPasswordUsedForBind() EnumresultCriteriaRetiredPasswordUsedForBindProp`

GetRetiredPasswordUsedForBind returns the RetiredPasswordUsedForBind field if non-nil, zero value otherwise.

### GetRetiredPasswordUsedForBindOk

`func (o *SimpleResultCriteriaShared) GetRetiredPasswordUsedForBindOk() (*EnumresultCriteriaRetiredPasswordUsedForBindProp, bool)`

GetRetiredPasswordUsedForBindOk returns a tuple with the RetiredPasswordUsedForBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiredPasswordUsedForBind

`func (o *SimpleResultCriteriaShared) SetRetiredPasswordUsedForBind(v EnumresultCriteriaRetiredPasswordUsedForBindProp)`

SetRetiredPasswordUsedForBind sets RetiredPasswordUsedForBind field to given value.

### HasRetiredPasswordUsedForBind

`func (o *SimpleResultCriteriaShared) HasRetiredPasswordUsedForBind() bool`

HasRetiredPasswordUsedForBind returns a boolean if a field has been set.

### GetSearchEntryReturnedCriteria

`func (o *SimpleResultCriteriaShared) GetSearchEntryReturnedCriteria() EnumresultCriteriaSearchEntryReturnedCriteriaProp`

GetSearchEntryReturnedCriteria returns the SearchEntryReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCriteriaOk

`func (o *SimpleResultCriteriaShared) GetSearchEntryReturnedCriteriaOk() (*EnumresultCriteriaSearchEntryReturnedCriteriaProp, bool)`

GetSearchEntryReturnedCriteriaOk returns a tuple with the SearchEntryReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCriteria

`func (o *SimpleResultCriteriaShared) SetSearchEntryReturnedCriteria(v EnumresultCriteriaSearchEntryReturnedCriteriaProp)`

SetSearchEntryReturnedCriteria sets SearchEntryReturnedCriteria field to given value.

### HasSearchEntryReturnedCriteria

`func (o *SimpleResultCriteriaShared) HasSearchEntryReturnedCriteria() bool`

HasSearchEntryReturnedCriteria returns a boolean if a field has been set.

### GetSearchEntryReturnedCount

`func (o *SimpleResultCriteriaShared) GetSearchEntryReturnedCount() int32`

GetSearchEntryReturnedCount returns the SearchEntryReturnedCount field if non-nil, zero value otherwise.

### GetSearchEntryReturnedCountOk

`func (o *SimpleResultCriteriaShared) GetSearchEntryReturnedCountOk() (*int32, bool)`

GetSearchEntryReturnedCountOk returns a tuple with the SearchEntryReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEntryReturnedCount

`func (o *SimpleResultCriteriaShared) SetSearchEntryReturnedCount(v int32)`

SetSearchEntryReturnedCount sets SearchEntryReturnedCount field to given value.

### HasSearchEntryReturnedCount

`func (o *SimpleResultCriteriaShared) HasSearchEntryReturnedCount() bool`

HasSearchEntryReturnedCount returns a boolean if a field has been set.

### GetSearchReferenceReturnedCriteria

`func (o *SimpleResultCriteriaShared) GetSearchReferenceReturnedCriteria() EnumresultCriteriaSearchReferenceReturnedCriteriaProp`

GetSearchReferenceReturnedCriteria returns the SearchReferenceReturnedCriteria field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCriteriaOk

`func (o *SimpleResultCriteriaShared) GetSearchReferenceReturnedCriteriaOk() (*EnumresultCriteriaSearchReferenceReturnedCriteriaProp, bool)`

GetSearchReferenceReturnedCriteriaOk returns a tuple with the SearchReferenceReturnedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCriteria

`func (o *SimpleResultCriteriaShared) SetSearchReferenceReturnedCriteria(v EnumresultCriteriaSearchReferenceReturnedCriteriaProp)`

SetSearchReferenceReturnedCriteria sets SearchReferenceReturnedCriteria field to given value.

### HasSearchReferenceReturnedCriteria

`func (o *SimpleResultCriteriaShared) HasSearchReferenceReturnedCriteria() bool`

HasSearchReferenceReturnedCriteria returns a boolean if a field has been set.

### GetSearchReferenceReturnedCount

`func (o *SimpleResultCriteriaShared) GetSearchReferenceReturnedCount() int32`

GetSearchReferenceReturnedCount returns the SearchReferenceReturnedCount field if non-nil, zero value otherwise.

### GetSearchReferenceReturnedCountOk

`func (o *SimpleResultCriteriaShared) GetSearchReferenceReturnedCountOk() (*int32, bool)`

GetSearchReferenceReturnedCountOk returns a tuple with the SearchReferenceReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchReferenceReturnedCount

`func (o *SimpleResultCriteriaShared) SetSearchReferenceReturnedCount(v int32)`

SetSearchReferenceReturnedCount sets SearchReferenceReturnedCount field to given value.

### HasSearchReferenceReturnedCount

`func (o *SimpleResultCriteriaShared) HasSearchReferenceReturnedCount() bool`

HasSearchReferenceReturnedCount returns a boolean if a field has been set.

### GetSearchIndexedCriteria

`func (o *SimpleResultCriteriaShared) GetSearchIndexedCriteria() EnumresultCriteriaSearchIndexedCriteriaProp`

GetSearchIndexedCriteria returns the SearchIndexedCriteria field if non-nil, zero value otherwise.

### GetSearchIndexedCriteriaOk

`func (o *SimpleResultCriteriaShared) GetSearchIndexedCriteriaOk() (*EnumresultCriteriaSearchIndexedCriteriaProp, bool)`

GetSearchIndexedCriteriaOk returns a tuple with the SearchIndexedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexedCriteria

`func (o *SimpleResultCriteriaShared) SetSearchIndexedCriteria(v EnumresultCriteriaSearchIndexedCriteriaProp)`

SetSearchIndexedCriteria sets SearchIndexedCriteria field to given value.

### HasSearchIndexedCriteria

`func (o *SimpleResultCriteriaShared) HasSearchIndexedCriteria() bool`

HasSearchIndexedCriteria returns a boolean if a field has been set.

### GetIncludedAuthzUserBaseDN

`func (o *SimpleResultCriteriaShared) GetIncludedAuthzUserBaseDN() []string`

GetIncludedAuthzUserBaseDN returns the IncludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedAuthzUserBaseDNOk

`func (o *SimpleResultCriteriaShared) GetIncludedAuthzUserBaseDNOk() (*[]string, bool)`

GetIncludedAuthzUserBaseDNOk returns a tuple with the IncludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAuthzUserBaseDN

`func (o *SimpleResultCriteriaShared) SetIncludedAuthzUserBaseDN(v []string)`

SetIncludedAuthzUserBaseDN sets IncludedAuthzUserBaseDN field to given value.

### HasIncludedAuthzUserBaseDN

`func (o *SimpleResultCriteriaShared) HasIncludedAuthzUserBaseDN() bool`

HasIncludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetExcludedAuthzUserBaseDN

`func (o *SimpleResultCriteriaShared) GetExcludedAuthzUserBaseDN() []string`

GetExcludedAuthzUserBaseDN returns the ExcludedAuthzUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedAuthzUserBaseDNOk

`func (o *SimpleResultCriteriaShared) GetExcludedAuthzUserBaseDNOk() (*[]string, bool)`

GetExcludedAuthzUserBaseDNOk returns a tuple with the ExcludedAuthzUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAuthzUserBaseDN

`func (o *SimpleResultCriteriaShared) SetExcludedAuthzUserBaseDN(v []string)`

SetExcludedAuthzUserBaseDN sets ExcludedAuthzUserBaseDN field to given value.

### HasExcludedAuthzUserBaseDN

`func (o *SimpleResultCriteriaShared) HasExcludedAuthzUserBaseDN() bool`

HasExcludedAuthzUserBaseDN returns a boolean if a field has been set.

### GetAllIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) GetAllIncludedAuthzUserGroupDN() []string`

GetAllIncludedAuthzUserGroupDN returns the AllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedAuthzUserGroupDNOk

`func (o *SimpleResultCriteriaShared) GetAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAllIncludedAuthzUserGroupDNOk returns a tuple with the AllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) SetAllIncludedAuthzUserGroupDN(v []string)`

SetAllIncludedAuthzUserGroupDN sets AllIncludedAuthzUserGroupDN field to given value.

### HasAllIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) HasAllIncludedAuthzUserGroupDN() bool`

HasAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetAnyIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) GetAnyIncludedAuthzUserGroupDN() []string`

GetAnyIncludedAuthzUserGroupDN returns the AnyIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedAuthzUserGroupDNOk

`func (o *SimpleResultCriteriaShared) GetAnyIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetAnyIncludedAuthzUserGroupDNOk returns a tuple with the AnyIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) SetAnyIncludedAuthzUserGroupDN(v []string)`

SetAnyIncludedAuthzUserGroupDN sets AnyIncludedAuthzUserGroupDN field to given value.

### HasAnyIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) HasAnyIncludedAuthzUserGroupDN() bool`

HasAnyIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) GetNotAllIncludedAuthzUserGroupDN() []string`

GetNotAllIncludedAuthzUserGroupDN returns the NotAllIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedAuthzUserGroupDNOk

`func (o *SimpleResultCriteriaShared) GetNotAllIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNotAllIncludedAuthzUserGroupDNOk returns a tuple with the NotAllIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) SetNotAllIncludedAuthzUserGroupDN(v []string)`

SetNotAllIncludedAuthzUserGroupDN sets NotAllIncludedAuthzUserGroupDN field to given value.

### HasNotAllIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) HasNotAllIncludedAuthzUserGroupDN() bool`

HasNotAllIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetNoneIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) GetNoneIncludedAuthzUserGroupDN() []string`

GetNoneIncludedAuthzUserGroupDN returns the NoneIncludedAuthzUserGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedAuthzUserGroupDNOk

`func (o *SimpleResultCriteriaShared) GetNoneIncludedAuthzUserGroupDNOk() (*[]string, bool)`

GetNoneIncludedAuthzUserGroupDNOk returns a tuple with the NoneIncludedAuthzUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) SetNoneIncludedAuthzUserGroupDN(v []string)`

SetNoneIncludedAuthzUserGroupDN sets NoneIncludedAuthzUserGroupDN field to given value.

### HasNoneIncludedAuthzUserGroupDN

`func (o *SimpleResultCriteriaShared) HasNoneIncludedAuthzUserGroupDN() bool`

HasNoneIncludedAuthzUserGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleResultCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleResultCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleResultCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleResultCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


