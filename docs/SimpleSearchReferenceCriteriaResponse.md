# SimpleSearchReferenceCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Search Reference Criteria | 
**Schemas** | [**[]EnumsimpleSearchReferenceCriteriaSchemaUrn**](EnumsimpleSearchReferenceCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria. | [optional] 
**AllIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that must be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain all of those controls. | [optional] 
**AnyIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that may be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain at least one of those controls. | [optional] 
**NotAllIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that should not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them). | [optional] 
**NoneIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that must not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain any of those controls. | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 

## Methods

### NewSimpleSearchReferenceCriteriaResponse

`func NewSimpleSearchReferenceCriteriaResponse(id string, schemas []EnumsimpleSearchReferenceCriteriaSchemaUrn, ) *SimpleSearchReferenceCriteriaResponse`

NewSimpleSearchReferenceCriteriaResponse instantiates a new SimpleSearchReferenceCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSearchReferenceCriteriaResponseWithDefaults

`func NewSimpleSearchReferenceCriteriaResponseWithDefaults() *SimpleSearchReferenceCriteriaResponse`

NewSimpleSearchReferenceCriteriaResponseWithDefaults instantiates a new SimpleSearchReferenceCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SimpleSearchReferenceCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SimpleSearchReferenceCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SimpleSearchReferenceCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleSearchReferenceCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SimpleSearchReferenceCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleSearchReferenceCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleSearchReferenceCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *SimpleSearchReferenceCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleSearchReferenceCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SimpleSearchReferenceCriteriaResponse) GetSchemas() []EnumsimpleSearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetSchemasOk() (*[]EnumsimpleSearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleSearchReferenceCriteriaResponse) SetSchemas(v []EnumsimpleSearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *SimpleSearchReferenceCriteriaResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *SimpleSearchReferenceCriteriaResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *SimpleSearchReferenceCriteriaResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) GetAllIncludedReferenceControl() []string`

GetAllIncludedReferenceControl returns the AllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAllIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetAllIncludedReferenceControlOk() (*[]string, bool)`

GetAllIncludedReferenceControlOk returns a tuple with the AllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) SetAllIncludedReferenceControl(v []string)`

SetAllIncludedReferenceControl sets AllIncludedReferenceControl field to given value.

### HasAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) HasAllIncludedReferenceControl() bool`

HasAllIncludedReferenceControl returns a boolean if a field has been set.

### GetAnyIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) GetAnyIncludedReferenceControl() []string`

GetAnyIncludedReferenceControl returns the AnyIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAnyIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetAnyIncludedReferenceControlOk() (*[]string, bool)`

GetAnyIncludedReferenceControlOk returns a tuple with the AnyIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) SetAnyIncludedReferenceControl(v []string)`

SetAnyIncludedReferenceControl sets AnyIncludedReferenceControl field to given value.

### HasAnyIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) HasAnyIncludedReferenceControl() bool`

HasAnyIncludedReferenceControl returns a boolean if a field has been set.

### GetNotAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) GetNotAllIncludedReferenceControl() []string`

GetNotAllIncludedReferenceControl returns the NotAllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNotAllIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetNotAllIncludedReferenceControlOk() (*[]string, bool)`

GetNotAllIncludedReferenceControlOk returns a tuple with the NotAllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) SetNotAllIncludedReferenceControl(v []string)`

SetNotAllIncludedReferenceControl sets NotAllIncludedReferenceControl field to given value.

### HasNotAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) HasNotAllIncludedReferenceControl() bool`

HasNotAllIncludedReferenceControl returns a boolean if a field has been set.

### GetNoneIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) GetNoneIncludedReferenceControl() []string`

GetNoneIncludedReferenceControl returns the NoneIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNoneIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetNoneIncludedReferenceControlOk() (*[]string, bool)`

GetNoneIncludedReferenceControlOk returns a tuple with the NoneIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) SetNoneIncludedReferenceControl(v []string)`

SetNoneIncludedReferenceControl sets NoneIncludedReferenceControl field to given value.

### HasNoneIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaResponse) HasNoneIncludedReferenceControl() bool`

HasNoneIncludedReferenceControl returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleSearchReferenceCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleSearchReferenceCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleSearchReferenceCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleSearchReferenceCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


