# ThirdPartySearchReferenceCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Reference Criteria | 
**Schemas** | [**[]EnumthirdPartySearchReferenceCriteriaSchemaUrn**](EnumthirdPartySearchReferenceCriteriaSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Search Reference Criteria. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 

## Methods

### NewThirdPartySearchReferenceCriteriaResponse

`func NewThirdPartySearchReferenceCriteriaResponse(id string, schemas []EnumthirdPartySearchReferenceCriteriaSchemaUrn, extensionClass string, ) *ThirdPartySearchReferenceCriteriaResponse`

NewThirdPartySearchReferenceCriteriaResponse instantiates a new ThirdPartySearchReferenceCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartySearchReferenceCriteriaResponseWithDefaults

`func NewThirdPartySearchReferenceCriteriaResponseWithDefaults() *ThirdPartySearchReferenceCriteriaResponse`

NewThirdPartySearchReferenceCriteriaResponseWithDefaults instantiates a new ThirdPartySearchReferenceCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartySearchReferenceCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetSchemas() []EnumthirdPartySearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetSchemasOk() (*[]EnumthirdPartySearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartySearchReferenceCriteriaResponse) SetSchemas(v []EnumthirdPartySearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartySearchReferenceCriteriaResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartySearchReferenceCriteriaResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartySearchReferenceCriteriaResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartySearchReferenceCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartySearchReferenceCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartySearchReferenceCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


