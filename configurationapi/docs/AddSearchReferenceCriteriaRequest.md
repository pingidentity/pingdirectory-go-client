# AddSearchReferenceCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Search Reference Criteria | 
**Schemas** | [**[]EnumthirdPartySearchReferenceCriteriaSchemaUrn**](EnumthirdPartySearchReferenceCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria. | [optional] 
**AllIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that must be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain all of those controls. | [optional] 
**AnyIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that may be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain at least one of those controls. | [optional] 
**NotAllIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that should not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them). | [optional] 
**NoneIncludedReferenceControl** | Pointer to **[]string** | Specifies the OID of a control that must not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain any of those controls. | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 
**AllIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that must match the associated search result reference in order to match the aggregate search reference criteria. If one or more all-included search reference criteria objects are provided, then a search result reference must match all of them in order to match the aggregate search reference criteria. | [optional] 
**AnyIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that may match the associated search result reference in order to match the aggregate search reference criteria. If one or more any-included search reference criteria objects are provided, then a search result reference must match at least one of them in order to match the aggregate search reference criteria. | [optional] 
**NotAllIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that should not match the associated search result reference in order to match the aggregate search reference criteria. If one or more not-all-included search reference criteria objects are provided, then a search result reference must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate search reference criteria. | [optional] 
**NoneIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that must not match the associated search result reference in order to match the aggregate search reference criteria. If one or more none-included search reference criteria objects are provided, then a search result reference must not match any of them in order to match the aggregate search reference criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Search Reference Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Search Reference Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddSearchReferenceCriteriaRequest

`func NewAddSearchReferenceCriteriaRequest(criteriaName string, schemas []EnumthirdPartySearchReferenceCriteriaSchemaUrn, extensionClass string, ) *AddSearchReferenceCriteriaRequest`

NewAddSearchReferenceCriteriaRequest instantiates a new AddSearchReferenceCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSearchReferenceCriteriaRequestWithDefaults

`func NewAddSearchReferenceCriteriaRequestWithDefaults() *AddSearchReferenceCriteriaRequest`

NewAddSearchReferenceCriteriaRequestWithDefaults instantiates a new AddSearchReferenceCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddSearchReferenceCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddSearchReferenceCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddSearchReferenceCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddSearchReferenceCriteriaRequest) GetSchemas() []EnumthirdPartySearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSearchReferenceCriteriaRequest) GetSchemasOk() (*[]EnumthirdPartySearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSearchReferenceCriteriaRequest) SetSchemas(v []EnumthirdPartySearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddSearchReferenceCriteriaRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddSearchReferenceCriteriaRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddSearchReferenceCriteriaRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddSearchReferenceCriteriaRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) GetAllIncludedReferenceControl() []string`

GetAllIncludedReferenceControl returns the AllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAllIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteriaRequest) GetAllIncludedReferenceControlOk() (*[]string, bool)`

GetAllIncludedReferenceControlOk returns a tuple with the AllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) SetAllIncludedReferenceControl(v []string)`

SetAllIncludedReferenceControl sets AllIncludedReferenceControl field to given value.

### HasAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) HasAllIncludedReferenceControl() bool`

HasAllIncludedReferenceControl returns a boolean if a field has been set.

### GetAnyIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) GetAnyIncludedReferenceControl() []string`

GetAnyIncludedReferenceControl returns the AnyIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAnyIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteriaRequest) GetAnyIncludedReferenceControlOk() (*[]string, bool)`

GetAnyIncludedReferenceControlOk returns a tuple with the AnyIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) SetAnyIncludedReferenceControl(v []string)`

SetAnyIncludedReferenceControl sets AnyIncludedReferenceControl field to given value.

### HasAnyIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) HasAnyIncludedReferenceControl() bool`

HasAnyIncludedReferenceControl returns a boolean if a field has been set.

### GetNotAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) GetNotAllIncludedReferenceControl() []string`

GetNotAllIncludedReferenceControl returns the NotAllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNotAllIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteriaRequest) GetNotAllIncludedReferenceControlOk() (*[]string, bool)`

GetNotAllIncludedReferenceControlOk returns a tuple with the NotAllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) SetNotAllIncludedReferenceControl(v []string)`

SetNotAllIncludedReferenceControl sets NotAllIncludedReferenceControl field to given value.

### HasNotAllIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) HasNotAllIncludedReferenceControl() bool`

HasNotAllIncludedReferenceControl returns a boolean if a field has been set.

### GetNoneIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) GetNoneIncludedReferenceControl() []string`

GetNoneIncludedReferenceControl returns the NoneIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNoneIncludedReferenceControlOk

`func (o *AddSearchReferenceCriteriaRequest) GetNoneIncludedReferenceControlOk() (*[]string, bool)`

GetNoneIncludedReferenceControlOk returns a tuple with the NoneIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) SetNoneIncludedReferenceControl(v []string)`

SetNoneIncludedReferenceControl sets NoneIncludedReferenceControl field to given value.

### HasNoneIncludedReferenceControl

`func (o *AddSearchReferenceCriteriaRequest) HasNoneIncludedReferenceControl() bool`

HasNoneIncludedReferenceControl returns a boolean if a field has been set.

### GetDescription

`func (o *AddSearchReferenceCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSearchReferenceCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSearchReferenceCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSearchReferenceCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) GetAllIncludedSearchReferenceCriteria() []string`

GetAllIncludedSearchReferenceCriteria returns the AllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteriaRequest) GetAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchReferenceCriteriaOk returns a tuple with the AllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) SetAllIncludedSearchReferenceCriteria(v []string)`

SetAllIncludedSearchReferenceCriteria sets AllIncludedSearchReferenceCriteria field to given value.

### HasAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) HasAllIncludedSearchReferenceCriteria() bool`

HasAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) GetAnyIncludedSearchReferenceCriteria() []string`

GetAnyIncludedSearchReferenceCriteria returns the AnyIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteriaRequest) GetAnyIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchReferenceCriteriaOk returns a tuple with the AnyIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) SetAnyIncludedSearchReferenceCriteria(v []string)`

SetAnyIncludedSearchReferenceCriteria sets AnyIncludedSearchReferenceCriteria field to given value.

### HasAnyIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) HasAnyIncludedSearchReferenceCriteria() bool`

HasAnyIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) GetNotAllIncludedSearchReferenceCriteria() []string`

GetNotAllIncludedSearchReferenceCriteria returns the NotAllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteriaRequest) GetNotAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchReferenceCriteriaOk returns a tuple with the NotAllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) SetNotAllIncludedSearchReferenceCriteria(v []string)`

SetNotAllIncludedSearchReferenceCriteria sets NotAllIncludedSearchReferenceCriteria field to given value.

### HasNotAllIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) HasNotAllIncludedSearchReferenceCriteria() bool`

HasNotAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) GetNoneIncludedSearchReferenceCriteria() []string`

GetNoneIncludedSearchReferenceCriteria returns the NoneIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchReferenceCriteriaOk

`func (o *AddSearchReferenceCriteriaRequest) GetNoneIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchReferenceCriteriaOk returns a tuple with the NoneIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) SetNoneIncludedSearchReferenceCriteria(v []string)`

SetNoneIncludedSearchReferenceCriteria sets NoneIncludedSearchReferenceCriteria field to given value.

### HasNoneIncludedSearchReferenceCriteria

`func (o *AddSearchReferenceCriteriaRequest) HasNoneIncludedSearchReferenceCriteria() bool`

HasNoneIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddSearchReferenceCriteriaRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddSearchReferenceCriteriaRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddSearchReferenceCriteriaRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddSearchReferenceCriteriaRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddSearchReferenceCriteriaRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddSearchReferenceCriteriaRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddSearchReferenceCriteriaRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


