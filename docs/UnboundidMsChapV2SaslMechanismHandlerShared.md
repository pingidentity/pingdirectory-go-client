# UnboundidMsChapV2SaslMechanismHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn**](EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the entry associated with the username provided in the bind request. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewUnboundidMsChapV2SaslMechanismHandlerShared

`func NewUnboundidMsChapV2SaslMechanismHandlerShared(schemas []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, ) *UnboundidMsChapV2SaslMechanismHandlerShared`

NewUnboundidMsChapV2SaslMechanismHandlerShared instantiates a new UnboundidMsChapV2SaslMechanismHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidMsChapV2SaslMechanismHandlerSharedWithDefaults

`func NewUnboundidMsChapV2SaslMechanismHandlerSharedWithDefaults() *UnboundidMsChapV2SaslMechanismHandlerShared`

NewUnboundidMsChapV2SaslMechanismHandlerSharedWithDefaults instantiates a new UnboundidMsChapV2SaslMechanismHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetSchemas() []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetSchemasOk() (*[]EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) SetSchemas(v []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetDescription

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidMsChapV2SaslMechanismHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


