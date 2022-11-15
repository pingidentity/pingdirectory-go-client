# RootDseRequestCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Request Criteria | 
**Schemas** | [**[]EnumrootDseRequestCriteriaSchemaUrn**](EnumrootDseRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaOperationTypeProp**](EnumrequestCriteriaOperationTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 

## Methods

### NewRootDseRequestCriteriaResponse

`func NewRootDseRequestCriteriaResponse(id string, schemas []EnumrootDseRequestCriteriaSchemaUrn, ) *RootDseRequestCriteriaResponse`

NewRootDseRequestCriteriaResponse instantiates a new RootDseRequestCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootDseRequestCriteriaResponseWithDefaults

`func NewRootDseRequestCriteriaResponseWithDefaults() *RootDseRequestCriteriaResponse`

NewRootDseRequestCriteriaResponseWithDefaults instantiates a new RootDseRequestCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RootDseRequestCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RootDseRequestCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RootDseRequestCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *RootDseRequestCriteriaResponse) GetSchemas() []EnumrootDseRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RootDseRequestCriteriaResponse) GetSchemasOk() (*[]EnumrootDseRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RootDseRequestCriteriaResponse) SetSchemas(v []EnumrootDseRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *RootDseRequestCriteriaResponse) GetOperationType() []EnumrequestCriteriaOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RootDseRequestCriteriaResponse) GetOperationTypeOk() (*[]EnumrequestCriteriaOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RootDseRequestCriteriaResponse) SetOperationType(v []EnumrequestCriteriaOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *RootDseRequestCriteriaResponse) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetDescription

`func (o *RootDseRequestCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RootDseRequestCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RootDseRequestCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RootDseRequestCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


