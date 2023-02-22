# AddLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationName** | **string** | Name of the new Location | 
**Schemas** | Pointer to [**[]EnumlocationSchemaUrn**](EnumlocationSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Location | [optional] 

## Methods

### NewAddLocationRequest

`func NewAddLocationRequest(locationName string, ) *AddLocationRequest`

NewAddLocationRequest instantiates a new AddLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLocationRequestWithDefaults

`func NewAddLocationRequestWithDefaults() *AddLocationRequest`

NewAddLocationRequestWithDefaults instantiates a new AddLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationName

`func (o *AddLocationRequest) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *AddLocationRequest) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *AddLocationRequest) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.


### GetSchemas

`func (o *AddLocationRequest) GetSchemas() []EnumlocationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLocationRequest) GetSchemasOk() (*[]EnumlocationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLocationRequest) SetSchemas(v []EnumlocationSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddLocationRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddLocationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLocationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLocationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLocationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


