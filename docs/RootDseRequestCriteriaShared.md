# RootDseRequestCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumrootDseRequestCriteriaSchemaUrn**](EnumrootDseRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaOperationTypeProp**](EnumrequestCriteriaOperationTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 

## Methods

### NewRootDseRequestCriteriaShared

`func NewRootDseRequestCriteriaShared(schemas []EnumrootDseRequestCriteriaSchemaUrn, ) *RootDseRequestCriteriaShared`

NewRootDseRequestCriteriaShared instantiates a new RootDseRequestCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootDseRequestCriteriaSharedWithDefaults

`func NewRootDseRequestCriteriaSharedWithDefaults() *RootDseRequestCriteriaShared`

NewRootDseRequestCriteriaSharedWithDefaults instantiates a new RootDseRequestCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *RootDseRequestCriteriaShared) GetSchemas() []EnumrootDseRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RootDseRequestCriteriaShared) GetSchemasOk() (*[]EnumrootDseRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RootDseRequestCriteriaShared) SetSchemas(v []EnumrootDseRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *RootDseRequestCriteriaShared) GetOperationType() []EnumrequestCriteriaOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RootDseRequestCriteriaShared) GetOperationTypeOk() (*[]EnumrequestCriteriaOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RootDseRequestCriteriaShared) SetOperationType(v []EnumrequestCriteriaOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *RootDseRequestCriteriaShared) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetDescription

`func (o *RootDseRequestCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RootDseRequestCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RootDseRequestCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RootDseRequestCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


