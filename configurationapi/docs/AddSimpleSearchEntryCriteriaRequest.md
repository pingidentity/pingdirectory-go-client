# AddSimpleSearchEntryCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsimpleSearchEntryCriteriaSchemaUrn**](EnumsimpleSearchEntryCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for entries included in this Simple Search Entry Criteria. of them. | [optional] 
**AllIncludedEntryControl** | Pointer to **[]string** | Specifies the OID of a control that must be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must contain all of those controls. | [optional] 
**AnyIncludedEntryControl** | Pointer to **[]string** | Specifies the OID of a control that may be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must contain at least one of those controls. | [optional] 
**NotAllIncludedEntryControl** | Pointer to **[]string** | Specifies the OID of a control that should not be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them). | [optional] 
**NoneIncludedEntryControl** | Pointer to **[]string** | Specifies the OID of a control that must not be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must not contain any of those controls. | [optional] 
**IncludedEntryBaseDN** | Pointer to **[]string** | Specifies a base DN below which entries included in this Simple Search Entry Criteria may exist. | [optional] 
**ExcludedEntryBaseDN** | Pointer to **[]string** | Specifies a base DN below which entries included in this Simple Search Entry Criteria may not exist. | [optional] 
**AllIncludedEntryFilter** | Pointer to **[]string** | Specifies a search filter that must match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the returned entry must match all of those filters. | [optional] 
**AnyIncludedEntryFilter** | Pointer to **[]string** | Specifies a search filter that may match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the entry must match at least one of those filters. | [optional] 
**NotAllIncludedEntryFilter** | Pointer to **[]string** | Specifies a search filter that should not match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the entry must not match at least one of those filters (that is, the entry may match zero or more of those filters, but not of all of them). | [optional] 
**NoneIncludedEntryFilter** | Pointer to **[]string** | Specifies a search filter that must not match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the entry must not match any of those filters. | [optional] 
**AllIncludedEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the entry must be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must be a member of all of them. | [optional] 
**AnyIncludedEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the entry may be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must be a member of at least one of them. | [optional] 
**NotAllIncludedEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the entry should not be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must not be a member of at least one of them (that is, the entry may be a member of zero or more of the specified groups, but not of all of them). | [optional] 
**NoneIncludedEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the entry must not be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must not be a member of any of them. | [optional] 
**Description** | Pointer to **string** | A description for this Search Entry Criteria | [optional] 
**CriteriaName** | **string** | Name of the new Search Entry Criteria | 

## Methods

### NewAddSimpleSearchEntryCriteriaRequest

`func NewAddSimpleSearchEntryCriteriaRequest(schemas []EnumsimpleSearchEntryCriteriaSchemaUrn, criteriaName string, ) *AddSimpleSearchEntryCriteriaRequest`

NewAddSimpleSearchEntryCriteriaRequest instantiates a new AddSimpleSearchEntryCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSimpleSearchEntryCriteriaRequestWithDefaults

`func NewAddSimpleSearchEntryCriteriaRequestWithDefaults() *AddSimpleSearchEntryCriteriaRequest`

NewAddSimpleSearchEntryCriteriaRequestWithDefaults instantiates a new AddSimpleSearchEntryCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSimpleSearchEntryCriteriaRequest) GetSchemas() []EnumsimpleSearchEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetSchemasOk() (*[]EnumsimpleSearchEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSimpleSearchEntryCriteriaRequest) SetSchemas(v []EnumsimpleSearchEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddSimpleSearchEntryCriteriaRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddSimpleSearchEntryCriteriaRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddSimpleSearchEntryCriteriaRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryControl() []string`

GetAllIncludedEntryControl returns the AllIncludedEntryControl field if non-nil, zero value otherwise.

### GetAllIncludedEntryControlOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryControlOk() (*[]string, bool)`

GetAllIncludedEntryControlOk returns a tuple with the AllIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) SetAllIncludedEntryControl(v []string)`

SetAllIncludedEntryControl sets AllIncludedEntryControl field to given value.

### HasAllIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) HasAllIncludedEntryControl() bool`

HasAllIncludedEntryControl returns a boolean if a field has been set.

### GetAnyIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryControl() []string`

GetAnyIncludedEntryControl returns the AnyIncludedEntryControl field if non-nil, zero value otherwise.

### GetAnyIncludedEntryControlOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryControlOk() (*[]string, bool)`

GetAnyIncludedEntryControlOk returns a tuple with the AnyIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) SetAnyIncludedEntryControl(v []string)`

SetAnyIncludedEntryControl sets AnyIncludedEntryControl field to given value.

### HasAnyIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) HasAnyIncludedEntryControl() bool`

HasAnyIncludedEntryControl returns a boolean if a field has been set.

### GetNotAllIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryControl() []string`

GetNotAllIncludedEntryControl returns the NotAllIncludedEntryControl field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryControlOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryControlOk() (*[]string, bool)`

GetNotAllIncludedEntryControlOk returns a tuple with the NotAllIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) SetNotAllIncludedEntryControl(v []string)`

SetNotAllIncludedEntryControl sets NotAllIncludedEntryControl field to given value.

### HasNotAllIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) HasNotAllIncludedEntryControl() bool`

HasNotAllIncludedEntryControl returns a boolean if a field has been set.

### GetNoneIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryControl() []string`

GetNoneIncludedEntryControl returns the NoneIncludedEntryControl field if non-nil, zero value otherwise.

### GetNoneIncludedEntryControlOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryControlOk() (*[]string, bool)`

GetNoneIncludedEntryControlOk returns a tuple with the NoneIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) SetNoneIncludedEntryControl(v []string)`

SetNoneIncludedEntryControl sets NoneIncludedEntryControl field to given value.

### HasNoneIncludedEntryControl

`func (o *AddSimpleSearchEntryCriteriaRequest) HasNoneIncludedEntryControl() bool`

HasNoneIncludedEntryControl returns a boolean if a field has been set.

### GetIncludedEntryBaseDN

`func (o *AddSimpleSearchEntryCriteriaRequest) GetIncludedEntryBaseDN() []string`

GetIncludedEntryBaseDN returns the IncludedEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedEntryBaseDNOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetIncludedEntryBaseDNOk() (*[]string, bool)`

GetIncludedEntryBaseDNOk returns a tuple with the IncludedEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedEntryBaseDN

`func (o *AddSimpleSearchEntryCriteriaRequest) SetIncludedEntryBaseDN(v []string)`

SetIncludedEntryBaseDN sets IncludedEntryBaseDN field to given value.

### HasIncludedEntryBaseDN

`func (o *AddSimpleSearchEntryCriteriaRequest) HasIncludedEntryBaseDN() bool`

HasIncludedEntryBaseDN returns a boolean if a field has been set.

### GetExcludedEntryBaseDN

`func (o *AddSimpleSearchEntryCriteriaRequest) GetExcludedEntryBaseDN() []string`

GetExcludedEntryBaseDN returns the ExcludedEntryBaseDN field if non-nil, zero value otherwise.

### GetExcludedEntryBaseDNOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetExcludedEntryBaseDNOk() (*[]string, bool)`

GetExcludedEntryBaseDNOk returns a tuple with the ExcludedEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntryBaseDN

`func (o *AddSimpleSearchEntryCriteriaRequest) SetExcludedEntryBaseDN(v []string)`

SetExcludedEntryBaseDN sets ExcludedEntryBaseDN field to given value.

### HasExcludedEntryBaseDN

`func (o *AddSimpleSearchEntryCriteriaRequest) HasExcludedEntryBaseDN() bool`

HasExcludedEntryBaseDN returns a boolean if a field has been set.

### GetAllIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryFilter() []string`

GetAllIncludedEntryFilter returns the AllIncludedEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedEntryFilterOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryFilterOk() (*[]string, bool)`

GetAllIncludedEntryFilterOk returns a tuple with the AllIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) SetAllIncludedEntryFilter(v []string)`

SetAllIncludedEntryFilter sets AllIncludedEntryFilter field to given value.

### HasAllIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) HasAllIncludedEntryFilter() bool`

HasAllIncludedEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryFilter() []string`

GetAnyIncludedEntryFilter returns the AnyIncludedEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedEntryFilterOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryFilterOk() (*[]string, bool)`

GetAnyIncludedEntryFilterOk returns a tuple with the AnyIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) SetAnyIncludedEntryFilter(v []string)`

SetAnyIncludedEntryFilter sets AnyIncludedEntryFilter field to given value.

### HasAnyIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) HasAnyIncludedEntryFilter() bool`

HasAnyIncludedEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryFilter() []string`

GetNotAllIncludedEntryFilter returns the NotAllIncludedEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryFilterOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedEntryFilterOk returns a tuple with the NotAllIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) SetNotAllIncludedEntryFilter(v []string)`

SetNotAllIncludedEntryFilter sets NotAllIncludedEntryFilter field to given value.

### HasNotAllIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) HasNotAllIncludedEntryFilter() bool`

HasNotAllIncludedEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryFilter() []string`

GetNoneIncludedEntryFilter returns the NoneIncludedEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedEntryFilterOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryFilterOk() (*[]string, bool)`

GetNoneIncludedEntryFilterOk returns a tuple with the NoneIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) SetNoneIncludedEntryFilter(v []string)`

SetNoneIncludedEntryFilter sets NoneIncludedEntryFilter field to given value.

### HasNoneIncludedEntryFilter

`func (o *AddSimpleSearchEntryCriteriaRequest) HasNoneIncludedEntryFilter() bool`

HasNoneIncludedEntryFilter returns a boolean if a field has been set.

### GetAllIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryGroupDN() []string`

GetAllIncludedEntryGroupDN returns the AllIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedEntryGroupDNOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedEntryGroupDNOk returns a tuple with the AllIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) SetAllIncludedEntryGroupDN(v []string)`

SetAllIncludedEntryGroupDN sets AllIncludedEntryGroupDN field to given value.

### HasAllIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) HasAllIncludedEntryGroupDN() bool`

HasAllIncludedEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryGroupDN() []string`

GetAnyIncludedEntryGroupDN returns the AnyIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedEntryGroupDNOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedEntryGroupDNOk returns a tuple with the AnyIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) SetAnyIncludedEntryGroupDN(v []string)`

SetAnyIncludedEntryGroupDN sets AnyIncludedEntryGroupDN field to given value.

### HasAnyIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) HasAnyIncludedEntryGroupDN() bool`

HasAnyIncludedEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryGroupDN() []string`

GetNotAllIncludedEntryGroupDN returns the NotAllIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryGroupDNOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedEntryGroupDNOk returns a tuple with the NotAllIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) SetNotAllIncludedEntryGroupDN(v []string)`

SetNotAllIncludedEntryGroupDN sets NotAllIncludedEntryGroupDN field to given value.

### HasNotAllIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) HasNotAllIncludedEntryGroupDN() bool`

HasNotAllIncludedEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryGroupDN() []string`

GetNoneIncludedEntryGroupDN returns the NoneIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedEntryGroupDNOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedEntryGroupDNOk returns a tuple with the NoneIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) SetNoneIncludedEntryGroupDN(v []string)`

SetNoneIncludedEntryGroupDN sets NoneIncludedEntryGroupDN field to given value.

### HasNoneIncludedEntryGroupDN

`func (o *AddSimpleSearchEntryCriteriaRequest) HasNoneIncludedEntryGroupDN() bool`

HasNoneIncludedEntryGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddSimpleSearchEntryCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSimpleSearchEntryCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSimpleSearchEntryCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteriaName

`func (o *AddSimpleSearchEntryCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddSimpleSearchEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddSimpleSearchEntryCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


