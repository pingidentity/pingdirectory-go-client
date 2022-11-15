# ThirdPartyPassThroughAuthenticationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Pass Through Authentication Handler | 
**Schemas** | [**[]EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn**](EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Pass Through Authentication Handler. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Pass Through Authentication Handler | [optional] 

## Methods

### NewThirdPartyPassThroughAuthenticationHandlerResponse

`func NewThirdPartyPassThroughAuthenticationHandlerResponse(id string, schemas []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, extensionClass string, ) *ThirdPartyPassThroughAuthenticationHandlerResponse`

NewThirdPartyPassThroughAuthenticationHandlerResponse instantiates a new ThirdPartyPassThroughAuthenticationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyPassThroughAuthenticationHandlerResponseWithDefaults

`func NewThirdPartyPassThroughAuthenticationHandlerResponseWithDefaults() *ThirdPartyPassThroughAuthenticationHandlerResponse`

NewThirdPartyPassThroughAuthenticationHandlerResponseWithDefaults instantiates a new ThirdPartyPassThroughAuthenticationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetSchemas() []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetSchemasOk() (*[]EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) SetSchemas(v []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyPassThroughAuthenticationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


