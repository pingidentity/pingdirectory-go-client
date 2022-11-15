# SimpleSearchReferenceCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Reference Criteria | 
**Schemas** | [**[]EnumsimpleSearchReferenceCriteriaSchemaUrn**](EnumsimpleSearchReferenceCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria. | [optional] 
**AllIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**AnyIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**NoneIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
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


