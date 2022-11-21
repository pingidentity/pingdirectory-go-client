# AddThirdPartyPasswordGeneratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratorName** | **string** | Name of the new Password Generator | 
**Schemas** | [**[]EnumthirdPartyPasswordGeneratorSchemaUrn**](EnumthirdPartyPasswordGeneratorSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Password Generator. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Password Generator. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewAddThirdPartyPasswordGeneratorRequest

`func NewAddThirdPartyPasswordGeneratorRequest(generatorName string, schemas []EnumthirdPartyPasswordGeneratorSchemaUrn, extensionClass string, enabled bool, ) *AddThirdPartyPasswordGeneratorRequest`

NewAddThirdPartyPasswordGeneratorRequest instantiates a new AddThirdPartyPasswordGeneratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyPasswordGeneratorRequestWithDefaults

`func NewAddThirdPartyPasswordGeneratorRequestWithDefaults() *AddThirdPartyPasswordGeneratorRequest`

NewAddThirdPartyPasswordGeneratorRequestWithDefaults instantiates a new AddThirdPartyPasswordGeneratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorName

`func (o *AddThirdPartyPasswordGeneratorRequest) GetGeneratorName() string`

GetGeneratorName returns the GeneratorName field if non-nil, zero value otherwise.

### GetGeneratorNameOk

`func (o *AddThirdPartyPasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool)`

GetGeneratorNameOk returns a tuple with the GeneratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorName

`func (o *AddThirdPartyPasswordGeneratorRequest) SetGeneratorName(v string)`

SetGeneratorName sets GeneratorName field to given value.


### GetSchemas

`func (o *AddThirdPartyPasswordGeneratorRequest) GetSchemas() []EnumthirdPartyPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyPasswordGeneratorRequest) GetSchemasOk() (*[]EnumthirdPartyPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyPasswordGeneratorRequest) SetSchemas(v []EnumthirdPartyPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyPasswordGeneratorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyPasswordGeneratorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyPasswordGeneratorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyPasswordGeneratorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyPasswordGeneratorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyPasswordGeneratorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyPasswordGeneratorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyPasswordGeneratorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyPasswordGeneratorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyPasswordGeneratorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


