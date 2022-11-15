# SimpleSearchReferenceCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsimpleSearchReferenceCriteriaSchemaUrn**](EnumsimpleSearchReferenceCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria. | [optional] 
**AllIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**AnyIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**NoneIncludedReferenceControl** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 

## Methods

### NewSimpleSearchReferenceCriteriaShared

`func NewSimpleSearchReferenceCriteriaShared(schemas []EnumsimpleSearchReferenceCriteriaSchemaUrn, ) *SimpleSearchReferenceCriteriaShared`

NewSimpleSearchReferenceCriteriaShared instantiates a new SimpleSearchReferenceCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSearchReferenceCriteriaSharedWithDefaults

`func NewSimpleSearchReferenceCriteriaSharedWithDefaults() *SimpleSearchReferenceCriteriaShared`

NewSimpleSearchReferenceCriteriaSharedWithDefaults instantiates a new SimpleSearchReferenceCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SimpleSearchReferenceCriteriaShared) GetSchemas() []EnumsimpleSearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleSearchReferenceCriteriaShared) GetSchemasOk() (*[]EnumsimpleSearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleSearchReferenceCriteriaShared) SetSchemas(v []EnumsimpleSearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *SimpleSearchReferenceCriteriaShared) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *SimpleSearchReferenceCriteriaShared) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *SimpleSearchReferenceCriteriaShared) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *SimpleSearchReferenceCriteriaShared) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) GetAllIncludedReferenceControl() []string`

GetAllIncludedReferenceControl returns the AllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAllIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaShared) GetAllIncludedReferenceControlOk() (*[]string, bool)`

GetAllIncludedReferenceControlOk returns a tuple with the AllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) SetAllIncludedReferenceControl(v []string)`

SetAllIncludedReferenceControl sets AllIncludedReferenceControl field to given value.

### HasAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) HasAllIncludedReferenceControl() bool`

HasAllIncludedReferenceControl returns a boolean if a field has been set.

### GetAnyIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) GetAnyIncludedReferenceControl() []string`

GetAnyIncludedReferenceControl returns the AnyIncludedReferenceControl field if non-nil, zero value otherwise.

### GetAnyIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaShared) GetAnyIncludedReferenceControlOk() (*[]string, bool)`

GetAnyIncludedReferenceControlOk returns a tuple with the AnyIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) SetAnyIncludedReferenceControl(v []string)`

SetAnyIncludedReferenceControl sets AnyIncludedReferenceControl field to given value.

### HasAnyIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) HasAnyIncludedReferenceControl() bool`

HasAnyIncludedReferenceControl returns a boolean if a field has been set.

### GetNotAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) GetNotAllIncludedReferenceControl() []string`

GetNotAllIncludedReferenceControl returns the NotAllIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNotAllIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaShared) GetNotAllIncludedReferenceControlOk() (*[]string, bool)`

GetNotAllIncludedReferenceControlOk returns a tuple with the NotAllIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) SetNotAllIncludedReferenceControl(v []string)`

SetNotAllIncludedReferenceControl sets NotAllIncludedReferenceControl field to given value.

### HasNotAllIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) HasNotAllIncludedReferenceControl() bool`

HasNotAllIncludedReferenceControl returns a boolean if a field has been set.

### GetNoneIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) GetNoneIncludedReferenceControl() []string`

GetNoneIncludedReferenceControl returns the NoneIncludedReferenceControl field if non-nil, zero value otherwise.

### GetNoneIncludedReferenceControlOk

`func (o *SimpleSearchReferenceCriteriaShared) GetNoneIncludedReferenceControlOk() (*[]string, bool)`

GetNoneIncludedReferenceControlOk returns a tuple with the NoneIncludedReferenceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) SetNoneIncludedReferenceControl(v []string)`

SetNoneIncludedReferenceControl sets NoneIncludedReferenceControl field to given value.

### HasNoneIncludedReferenceControl

`func (o *SimpleSearchReferenceCriteriaShared) HasNoneIncludedReferenceControl() bool`

HasNoneIncludedReferenceControl returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleSearchReferenceCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleSearchReferenceCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleSearchReferenceCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleSearchReferenceCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


