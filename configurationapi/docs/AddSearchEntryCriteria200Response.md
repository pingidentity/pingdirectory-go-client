# AddSearchEntryCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Entry Criteria | 
**Schemas** | [**[]EnumthirdPartySearchEntryCriteriaSchemaUrn**](EnumthirdPartySearchEntryCriteriaSchemaUrn.md) |  | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**AllIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that must match the associated search result entry in order to match the aggregate search entry criteria. If one or more all-included search entry criteria objects are provided, then a search result entry must match all of them in order to match the aggregate search entry criteria. | [optional] 
**AnyIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that may match the associated search result entry in order to match the aggregate search entry criteria. If one or more any-included search entry criteria objects are provided, then a search result entry must match at least one of them in order to match the aggregate search entry criteria. | [optional] 
**NotAllIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that should not match the associated search result entry in order to match the aggregate search entry criteria. If one or more not-all-included search entry criteria objects are provided, then a search result entry must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate search entry criteria. | [optional] 
**NoneIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that must not match the associated search result entry in order to match the aggregate search entry criteria. If one or more none-included search entry criteria objects are provided, then a search result entry must not match any of them in order to match the aggregate search entry criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Search Entry Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Search Entry Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddSearchEntryCriteria200Response

`func NewAddSearchEntryCriteria200Response(id string, schemas []EnumthirdPartySearchEntryCriteriaSchemaUrn, extensionClass string, ) *AddSearchEntryCriteria200Response`

NewAddSearchEntryCriteria200Response instantiates a new AddSearchEntryCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSearchEntryCriteria200ResponseWithDefaults

`func NewAddSearchEntryCriteria200ResponseWithDefaults() *AddSearchEntryCriteria200Response`

NewAddSearchEntryCriteria200ResponseWithDefaults instantiates a new AddSearchEntryCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSearchEntryCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSearchEntryCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSearchEntryCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddSearchEntryCriteria200Response) GetSchemas() []EnumthirdPartySearchEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSearchEntryCriteria200Response) GetSchemasOk() (*[]EnumthirdPartySearchEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSearchEntryCriteria200Response) SetSchemas(v []EnumthirdPartySearchEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddSearchEntryCriteria200Response) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddSearchEntryCriteria200Response) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddSearchEntryCriteria200Response) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddSearchEntryCriteria200Response) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedEntryControl() []string`

GetAllIncludedEntryControl returns the AllIncludedEntryControl field if non-nil, zero value otherwise.

### GetAllIncludedEntryControlOk

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedEntryControlOk() (*[]string, bool)`

GetAllIncludedEntryControlOk returns a tuple with the AllIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) SetAllIncludedEntryControl(v []string)`

SetAllIncludedEntryControl sets AllIncludedEntryControl field to given value.

### HasAllIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) HasAllIncludedEntryControl() bool`

HasAllIncludedEntryControl returns a boolean if a field has been set.

### GetAnyIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedEntryControl() []string`

GetAnyIncludedEntryControl returns the AnyIncludedEntryControl field if non-nil, zero value otherwise.

### GetAnyIncludedEntryControlOk

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedEntryControlOk() (*[]string, bool)`

GetAnyIncludedEntryControlOk returns a tuple with the AnyIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) SetAnyIncludedEntryControl(v []string)`

SetAnyIncludedEntryControl sets AnyIncludedEntryControl field to given value.

### HasAnyIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) HasAnyIncludedEntryControl() bool`

HasAnyIncludedEntryControl returns a boolean if a field has been set.

### GetNotAllIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedEntryControl() []string`

GetNotAllIncludedEntryControl returns the NotAllIncludedEntryControl field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryControlOk

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedEntryControlOk() (*[]string, bool)`

GetNotAllIncludedEntryControlOk returns a tuple with the NotAllIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) SetNotAllIncludedEntryControl(v []string)`

SetNotAllIncludedEntryControl sets NotAllIncludedEntryControl field to given value.

### HasNotAllIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) HasNotAllIncludedEntryControl() bool`

HasNotAllIncludedEntryControl returns a boolean if a field has been set.

### GetNoneIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedEntryControl() []string`

GetNoneIncludedEntryControl returns the NoneIncludedEntryControl field if non-nil, zero value otherwise.

### GetNoneIncludedEntryControlOk

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedEntryControlOk() (*[]string, bool)`

GetNoneIncludedEntryControlOk returns a tuple with the NoneIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) SetNoneIncludedEntryControl(v []string)`

SetNoneIncludedEntryControl sets NoneIncludedEntryControl field to given value.

### HasNoneIncludedEntryControl

`func (o *AddSearchEntryCriteria200Response) HasNoneIncludedEntryControl() bool`

HasNoneIncludedEntryControl returns a boolean if a field has been set.

### GetIncludedEntryBaseDN

`func (o *AddSearchEntryCriteria200Response) GetIncludedEntryBaseDN() []string`

GetIncludedEntryBaseDN returns the IncludedEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedEntryBaseDNOk

`func (o *AddSearchEntryCriteria200Response) GetIncludedEntryBaseDNOk() (*[]string, bool)`

GetIncludedEntryBaseDNOk returns a tuple with the IncludedEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedEntryBaseDN

`func (o *AddSearchEntryCriteria200Response) SetIncludedEntryBaseDN(v []string)`

SetIncludedEntryBaseDN sets IncludedEntryBaseDN field to given value.

### HasIncludedEntryBaseDN

`func (o *AddSearchEntryCriteria200Response) HasIncludedEntryBaseDN() bool`

HasIncludedEntryBaseDN returns a boolean if a field has been set.

### GetExcludedEntryBaseDN

`func (o *AddSearchEntryCriteria200Response) GetExcludedEntryBaseDN() []string`

GetExcludedEntryBaseDN returns the ExcludedEntryBaseDN field if non-nil, zero value otherwise.

### GetExcludedEntryBaseDNOk

`func (o *AddSearchEntryCriteria200Response) GetExcludedEntryBaseDNOk() (*[]string, bool)`

GetExcludedEntryBaseDNOk returns a tuple with the ExcludedEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntryBaseDN

`func (o *AddSearchEntryCriteria200Response) SetExcludedEntryBaseDN(v []string)`

SetExcludedEntryBaseDN sets ExcludedEntryBaseDN field to given value.

### HasExcludedEntryBaseDN

`func (o *AddSearchEntryCriteria200Response) HasExcludedEntryBaseDN() bool`

HasExcludedEntryBaseDN returns a boolean if a field has been set.

### GetAllIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedEntryFilter() []string`

GetAllIncludedEntryFilter returns the AllIncludedEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedEntryFilterOk

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedEntryFilterOk() (*[]string, bool)`

GetAllIncludedEntryFilterOk returns a tuple with the AllIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) SetAllIncludedEntryFilter(v []string)`

SetAllIncludedEntryFilter sets AllIncludedEntryFilter field to given value.

### HasAllIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) HasAllIncludedEntryFilter() bool`

HasAllIncludedEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedEntryFilter() []string`

GetAnyIncludedEntryFilter returns the AnyIncludedEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedEntryFilterOk

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedEntryFilterOk() (*[]string, bool)`

GetAnyIncludedEntryFilterOk returns a tuple with the AnyIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) SetAnyIncludedEntryFilter(v []string)`

SetAnyIncludedEntryFilter sets AnyIncludedEntryFilter field to given value.

### HasAnyIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) HasAnyIncludedEntryFilter() bool`

HasAnyIncludedEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedEntryFilter() []string`

GetNotAllIncludedEntryFilter returns the NotAllIncludedEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryFilterOk

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedEntryFilterOk returns a tuple with the NotAllIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) SetNotAllIncludedEntryFilter(v []string)`

SetNotAllIncludedEntryFilter sets NotAllIncludedEntryFilter field to given value.

### HasNotAllIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) HasNotAllIncludedEntryFilter() bool`

HasNotAllIncludedEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedEntryFilter() []string`

GetNoneIncludedEntryFilter returns the NoneIncludedEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedEntryFilterOk

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedEntryFilterOk() (*[]string, bool)`

GetNoneIncludedEntryFilterOk returns a tuple with the NoneIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) SetNoneIncludedEntryFilter(v []string)`

SetNoneIncludedEntryFilter sets NoneIncludedEntryFilter field to given value.

### HasNoneIncludedEntryFilter

`func (o *AddSearchEntryCriteria200Response) HasNoneIncludedEntryFilter() bool`

HasNoneIncludedEntryFilter returns a boolean if a field has been set.

### GetAllIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedEntryGroupDN() []string`

GetAllIncludedEntryGroupDN returns the AllIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedEntryGroupDNOk

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedEntryGroupDNOk returns a tuple with the AllIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) SetAllIncludedEntryGroupDN(v []string)`

SetAllIncludedEntryGroupDN sets AllIncludedEntryGroupDN field to given value.

### HasAllIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) HasAllIncludedEntryGroupDN() bool`

HasAllIncludedEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedEntryGroupDN() []string`

GetAnyIncludedEntryGroupDN returns the AnyIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedEntryGroupDNOk

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedEntryGroupDNOk returns a tuple with the AnyIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) SetAnyIncludedEntryGroupDN(v []string)`

SetAnyIncludedEntryGroupDN sets AnyIncludedEntryGroupDN field to given value.

### HasAnyIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) HasAnyIncludedEntryGroupDN() bool`

HasAnyIncludedEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedEntryGroupDN() []string`

GetNotAllIncludedEntryGroupDN returns the NotAllIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryGroupDNOk

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedEntryGroupDNOk returns a tuple with the NotAllIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) SetNotAllIncludedEntryGroupDN(v []string)`

SetNotAllIncludedEntryGroupDN sets NotAllIncludedEntryGroupDN field to given value.

### HasNotAllIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) HasNotAllIncludedEntryGroupDN() bool`

HasNotAllIncludedEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedEntryGroupDN() []string`

GetNoneIncludedEntryGroupDN returns the NoneIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedEntryGroupDNOk

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedEntryGroupDNOk returns a tuple with the NoneIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) SetNoneIncludedEntryGroupDN(v []string)`

SetNoneIncludedEntryGroupDN sets NoneIncludedEntryGroupDN field to given value.

### HasNoneIncludedEntryGroupDN

`func (o *AddSearchEntryCriteria200Response) HasNoneIncludedEntryGroupDN() bool`

HasNoneIncludedEntryGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddSearchEntryCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSearchEntryCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSearchEntryCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSearchEntryCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddSearchEntryCriteria200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddSearchEntryCriteria200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddSearchEntryCriteria200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddSearchEntryCriteria200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddSearchEntryCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddSearchEntryCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddSearchEntryCriteria200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddSearchEntryCriteria200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetAllIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedSearchEntryCriteria() []string`

GetAllIncludedSearchEntryCriteria returns the AllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchEntryCriteriaOk

`func (o *AddSearchEntryCriteria200Response) GetAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchEntryCriteriaOk returns a tuple with the AllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) SetAllIncludedSearchEntryCriteria(v []string)`

SetAllIncludedSearchEntryCriteria sets AllIncludedSearchEntryCriteria field to given value.

### HasAllIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) HasAllIncludedSearchEntryCriteria() bool`

HasAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedSearchEntryCriteria() []string`

GetAnyIncludedSearchEntryCriteria returns the AnyIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchEntryCriteriaOk

`func (o *AddSearchEntryCriteria200Response) GetAnyIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchEntryCriteriaOk returns a tuple with the AnyIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) SetAnyIncludedSearchEntryCriteria(v []string)`

SetAnyIncludedSearchEntryCriteria sets AnyIncludedSearchEntryCriteria field to given value.

### HasAnyIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) HasAnyIncludedSearchEntryCriteria() bool`

HasAnyIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedSearchEntryCriteria() []string`

GetNotAllIncludedSearchEntryCriteria returns the NotAllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchEntryCriteriaOk

`func (o *AddSearchEntryCriteria200Response) GetNotAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchEntryCriteriaOk returns a tuple with the NotAllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) SetNotAllIncludedSearchEntryCriteria(v []string)`

SetNotAllIncludedSearchEntryCriteria sets NotAllIncludedSearchEntryCriteria field to given value.

### HasNotAllIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) HasNotAllIncludedSearchEntryCriteria() bool`

HasNotAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedSearchEntryCriteria() []string`

GetNoneIncludedSearchEntryCriteria returns the NoneIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchEntryCriteriaOk

`func (o *AddSearchEntryCriteria200Response) GetNoneIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchEntryCriteriaOk returns a tuple with the NoneIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) SetNoneIncludedSearchEntryCriteria(v []string)`

SetNoneIncludedSearchEntryCriteria sets NoneIncludedSearchEntryCriteria field to given value.

### HasNoneIncludedSearchEntryCriteria

`func (o *AddSearchEntryCriteria200Response) HasNoneIncludedSearchEntryCriteria() bool`

HasNoneIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddSearchEntryCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddSearchEntryCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddSearchEntryCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddSearchEntryCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddSearchEntryCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddSearchEntryCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddSearchEntryCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


