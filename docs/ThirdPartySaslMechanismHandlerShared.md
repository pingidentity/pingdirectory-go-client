# ThirdPartySaslMechanismHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartySaslMechanismHandlerSchemaUrn**](EnumthirdPartySaslMechanismHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party SASL Mechanism Handler. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**IdentityMapper** | Pointer to **string** | The identity mapper that may be used to map usernames to user entries. If the custom SASL mechanism involves a username or some other form of authentication and/or authorization identity, then this may be used to map that ID to an entry for that user. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewThirdPartySaslMechanismHandlerShared

`func NewThirdPartySaslMechanismHandlerShared(schemas []EnumthirdPartySaslMechanismHandlerSchemaUrn, extensionClass string, enabled bool, ) *ThirdPartySaslMechanismHandlerShared`

NewThirdPartySaslMechanismHandlerShared instantiates a new ThirdPartySaslMechanismHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartySaslMechanismHandlerSharedWithDefaults

`func NewThirdPartySaslMechanismHandlerSharedWithDefaults() *ThirdPartySaslMechanismHandlerShared`

NewThirdPartySaslMechanismHandlerSharedWithDefaults instantiates a new ThirdPartySaslMechanismHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartySaslMechanismHandlerShared) GetSchemas() []EnumthirdPartySaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartySaslMechanismHandlerShared) GetSchemasOk() (*[]EnumthirdPartySaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartySaslMechanismHandlerShared) SetSchemas(v []EnumthirdPartySaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartySaslMechanismHandlerShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartySaslMechanismHandlerShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartySaslMechanismHandlerShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartySaslMechanismHandlerShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartySaslMechanismHandlerShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartySaslMechanismHandlerShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartySaslMechanismHandlerShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *ThirdPartySaslMechanismHandlerShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *ThirdPartySaslMechanismHandlerShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *ThirdPartySaslMechanismHandlerShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *ThirdPartySaslMechanismHandlerShared) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartySaslMechanismHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartySaslMechanismHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartySaslMechanismHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartySaslMechanismHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartySaslMechanismHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartySaslMechanismHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartySaslMechanismHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


