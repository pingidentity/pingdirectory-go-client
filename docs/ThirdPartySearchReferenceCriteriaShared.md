# ThirdPartySearchReferenceCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartySearchReferenceCriteriaSchemaUrn**](EnumthirdPartySearchReferenceCriteriaSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Search Reference Criteria. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 

## Methods

### NewThirdPartySearchReferenceCriteriaShared

`func NewThirdPartySearchReferenceCriteriaShared(schemas []EnumthirdPartySearchReferenceCriteriaSchemaUrn, extensionClass string, ) *ThirdPartySearchReferenceCriteriaShared`

NewThirdPartySearchReferenceCriteriaShared instantiates a new ThirdPartySearchReferenceCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartySearchReferenceCriteriaSharedWithDefaults

`func NewThirdPartySearchReferenceCriteriaSharedWithDefaults() *ThirdPartySearchReferenceCriteriaShared`

NewThirdPartySearchReferenceCriteriaSharedWithDefaults instantiates a new ThirdPartySearchReferenceCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartySearchReferenceCriteriaShared) GetSchemas() []EnumthirdPartySearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartySearchReferenceCriteriaShared) GetSchemasOk() (*[]EnumthirdPartySearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartySearchReferenceCriteriaShared) SetSchemas(v []EnumthirdPartySearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartySearchReferenceCriteriaShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartySearchReferenceCriteriaShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartySearchReferenceCriteriaShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartySearchReferenceCriteriaShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartySearchReferenceCriteriaShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartySearchReferenceCriteriaShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartySearchReferenceCriteriaShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartySearchReferenceCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartySearchReferenceCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartySearchReferenceCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartySearchReferenceCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


