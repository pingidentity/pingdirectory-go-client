# AddRootDseRequestCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Request Criteria | 
**Schemas** | [**[]EnumrootDseRequestCriteriaSchemaUrn**](EnumrootDseRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaRootDseOperationTypeProp**](EnumrequestCriteriaRootDseOperationTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 

## Methods

### NewAddRootDseRequestCriteriaRequest

`func NewAddRootDseRequestCriteriaRequest(criteriaName string, schemas []EnumrootDseRequestCriteriaSchemaUrn, ) *AddRootDseRequestCriteriaRequest`

NewAddRootDseRequestCriteriaRequest instantiates a new AddRootDseRequestCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRootDseRequestCriteriaRequestWithDefaults

`func NewAddRootDseRequestCriteriaRequestWithDefaults() *AddRootDseRequestCriteriaRequest`

NewAddRootDseRequestCriteriaRequestWithDefaults instantiates a new AddRootDseRequestCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddRootDseRequestCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddRootDseRequestCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddRootDseRequestCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddRootDseRequestCriteriaRequest) GetSchemas() []EnumrootDseRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRootDseRequestCriteriaRequest) GetSchemasOk() (*[]EnumrootDseRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRootDseRequestCriteriaRequest) SetSchemas(v []EnumrootDseRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *AddRootDseRequestCriteriaRequest) GetOperationType() []EnumrequestCriteriaRootDseOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AddRootDseRequestCriteriaRequest) GetOperationTypeOk() (*[]EnumrequestCriteriaRootDseOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AddRootDseRequestCriteriaRequest) SetOperationType(v []EnumrequestCriteriaRootDseOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *AddRootDseRequestCriteriaRequest) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetDescription

`func (o *AddRootDseRequestCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRootDseRequestCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRootDseRequestCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRootDseRequestCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


