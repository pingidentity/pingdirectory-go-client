# AddSearchReferenceCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Reference Criteria | 
**Schemas** | [**[]EnumthirdPartySearchReferenceCriteriaSchemaUrn**](EnumthirdPartySearchReferenceCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria. | [optional] 
**AllIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that must be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain all of those controls. | [optional] 
**AnyIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that may be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain at least one of those controls. | [optional] 
**NotAllIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that should not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them). | [optional] 
**NoneIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that must not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain any of those controls. | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**AllIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that must match the associated search result reference in order to match the aggregate search reference criteria. If one or more all-included search reference criteria objects are provided, then a search result reference must match all of them in order to match the aggregate search reference criteria. | [optional] 
**AnyIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that may match the associated search result reference in order to match the aggregate search reference criteria. If one or more any-included search reference criteria objects are provided, then a search result reference must match at least one of them in order to match the aggregate search reference criteria. | [optional] 
**NotAllIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that should not match the associated search result reference in order to match the aggregate search reference criteria. If one or more not-all-included search reference criteria objects are provided, then a search result reference must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate search reference criteria. | [optional] 
**NoneIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that must not match the associated search result reference in order to match the aggregate search reference criteria. If one or more none-included search reference criteria objects are provided, then a search result reference must not match any of them in order to match the aggregate search reference criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Search Reference Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Search Reference Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddSearchReferenceCriteria200Response

`func NewAddSearchReferenceCriteria200Response(id string, schemas []EnumthirdPartySearchReferenceCriteriaSchemaUrn, extensionClass string, ) *AddSearchReferenceCriteria200Response`

NewAddSearchReferenceCriteria200Response instantiates a new AddSearchReferenceCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSearchReferenceCriteria200ResponseWithDefaults

`func NewAddSearchReferenceCriteria200ResponseWithDefaults() *AddSearchReferenceCriteria200Response`

NewAddSearchReferenceCriteria200ResponseWithDefaults instantiates a new AddSearchReferenceCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSearchReferenceCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSearchReferenceCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSearchReferenceCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddSearchReferenceCriteria200Response) GetSchemas() []EnumthirdPartySearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSearchReferenceCriteria200Response) GetSchemasOk() (*[]EnumthirdPartySearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSearchReferenceCriteria200Response) SetSchemas(v []EnumthirdPartySearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddSearchReferenceCriteria200Response) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddSearchReferenceCriteria200Response) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddSearchReferenceCriteria200Response) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddSearchReferenceCriteria200Response) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) GetAllIncludedReferenceControl() []string`

GetAllIncludedReferenceControl returns the AllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAllIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteria200Response) GetAllIncludedReferenceControlOk() (*[]string, bool)`

GetAllIncludedReferenceControlOk returns a tuple with the AllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) SetAllIncludedReferenceControl(v []string)`

SetAllIncludedReferenceControl sets AllIncludedReferenceControl field to given value.

### HasAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) HasAllIncludedReferenceControl() bool`

HasAllIncludedReferenceControl returns a boolean if a field has been set.

### GetAnyIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) GetAnyIncludedReferenceControl() []string`

GetAnyIncludedReferenceControl returns the AnyIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAnyIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteria200Response) GetAnyIncludedReferenceControlOk() (*[]string, bool)`

GetAnyIncludedReferenceControlOk returns a tuple with the AnyIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) SetAnyIncludedReferenceControl(v []string)`

SetAnyIncludedReferenceControl sets AnyIncludedReferenceControl field to given value.

### HasAnyIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) HasAnyIncludedReferenceControl() bool`

HasAnyIncludedReferenceControl returns a boolean if a field has been set.

### GetNotAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) GetNotAllIncludedReferenceControl() []string`

GetNotAllIncludedReferenceControl returns the NotAllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNotAllIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteria200Response) GetNotAllIncludedReferenceControlOk() (*[]string, bool)`

GetNotAllIncludedReferenceControlOk returns a tuple with the NotAllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) SetNotAllIncludedReferenceControl(v []string)`

SetNotAllIncludedReferenceControl sets NotAllIncludedReferenceControl field to given value.

### HasNotAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) HasNotAllIncludedReferenceControl() bool`

HasNotAllIncludedReferenceControl returns a boolean if a field has been set.

### GetNoneIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) GetNoneIncludedReferenceControl() []string`

GetNoneIncludedReferenceControl returns the NoneIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNoneIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteria200Response) GetNoneIncludedReferenceControlOk() (*[]string, bool)`

GetNoneIncludedReferenceControlOk returns a tuple with the NoneIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) SetNoneIncludedReferenceControl(v []string)`

SetNoneIncludedReferenceControl sets NoneIncludedReferenceControl field to given value.

### HasNoneIncludedReferenceControl

`func (o *AddSearchReferenceCriteria200Response) HasNoneIncludedReferenceControl() bool`

HasNoneIncludedReferenceControl returns a boolean if a field has been set.

### GetDescription

`func (o *AddSearchReferenceCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSearchReferenceCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSearchReferenceCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSearchReferenceCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddSearchReferenceCriteria200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddSearchReferenceCriteria200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddSearchReferenceCriteria200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddSearchReferenceCriteria200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) GetAllIncludedSearchReferenceCriteria() []string`

GetAllIncludedSearchReferenceCriteria returns the AllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteria200Response) GetAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchReferenceCriteriaOk returns a tuple with the AllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) SetAllIncludedSearchReferenceCriteria(v []string)`

SetAllIncludedSearchReferenceCriteria sets AllIncludedSearchReferenceCriteria field to given value.

### HasAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) HasAllIncludedSearchReferenceCriteria() bool`

HasAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) GetAnyIncludedSearchReferenceCriteria() []string`

GetAnyIncludedSearchReferenceCriteria returns the AnyIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteria200Response) GetAnyIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchReferenceCriteriaOk returns a tuple with the AnyIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) SetAnyIncludedSearchReferenceCriteria(v []string)`

SetAnyIncludedSearchReferenceCriteria sets AnyIncludedSearchReferenceCriteria field to given value.

### HasAnyIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) HasAnyIncludedSearchReferenceCriteria() bool`

HasAnyIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) GetNotAllIncludedSearchReferenceCriteria() []string`

GetNotAllIncludedSearchReferenceCriteria returns the NotAllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteria200Response) GetNotAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchReferenceCriteriaOk returns a tuple with the NotAllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) SetNotAllIncludedSearchReferenceCriteria(v []string)`

SetNotAllIncludedSearchReferenceCriteria sets NotAllIncludedSearchReferenceCriteria field to given value.

### HasNotAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) HasNotAllIncludedSearchReferenceCriteria() bool`

HasNotAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) GetNoneIncludedSearchReferenceCriteria() []string`

GetNoneIncludedSearchReferenceCriteria returns the NoneIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteria200Response) GetNoneIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchReferenceCriteriaOk returns a tuple with the NoneIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) SetNoneIncludedSearchReferenceCriteria(v []string)`

SetNoneIncludedSearchReferenceCriteria sets NoneIncludedSearchReferenceCriteria field to given value.

### HasNoneIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteria200Response) HasNoneIncludedSearchReferenceCriteria() bool`

HasNoneIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddSearchReferenceCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddSearchReferenceCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddSearchReferenceCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddSearchReferenceCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddSearchReferenceCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddSearchReferenceCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddSearchReferenceCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


