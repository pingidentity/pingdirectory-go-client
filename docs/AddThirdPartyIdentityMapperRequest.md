# AddThirdPartyIdentityMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Identity Mapper | 
**Schemas** | [**[]EnumthirdPartyIdentityMapperSchemaUrn**](EnumthirdPartyIdentityMapperSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Identity Mapper. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewAddThirdPartyIdentityMapperRequest

`func NewAddThirdPartyIdentityMapperRequest(mapperName string, schemas []EnumthirdPartyIdentityMapperSchemaUrn, extensionClass string, enabled bool, ) *AddThirdPartyIdentityMapperRequest`

NewAddThirdPartyIdentityMapperRequest instantiates a new AddThirdPartyIdentityMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyIdentityMapperRequestWithDefaults

`func NewAddThirdPartyIdentityMapperRequestWithDefaults() *AddThirdPartyIdentityMapperRequest`

NewAddThirdPartyIdentityMapperRequestWithDefaults instantiates a new AddThirdPartyIdentityMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddThirdPartyIdentityMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddThirdPartyIdentityMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddThirdPartyIdentityMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddThirdPartyIdentityMapperRequest) GetSchemas() []EnumthirdPartyIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyIdentityMapperRequest) GetSchemasOk() (*[]EnumthirdPartyIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyIdentityMapperRequest) SetSchemas(v []EnumthirdPartyIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyIdentityMapperRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyIdentityMapperRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyIdentityMapperRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyIdentityMapperRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyIdentityMapperRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyIdentityMapperRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyIdentityMapperRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyIdentityMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyIdentityMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyIdentityMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyIdentityMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyIdentityMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyIdentityMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyIdentityMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


