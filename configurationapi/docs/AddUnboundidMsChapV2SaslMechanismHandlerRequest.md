# AddUnboundidMsChapV2SaslMechanismHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new SASL Mechanism Handler | 
**Schemas** | [**[]EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn**](EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the entry associated with the username provided in the bind request. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewAddUnboundidMsChapV2SaslMechanismHandlerRequest

`func NewAddUnboundidMsChapV2SaslMechanismHandlerRequest(handlerName string, schemas []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, ) *AddUnboundidMsChapV2SaslMechanismHandlerRequest`

NewAddUnboundidMsChapV2SaslMechanismHandlerRequest instantiates a new AddUnboundidMsChapV2SaslMechanismHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUnboundidMsChapV2SaslMechanismHandlerRequestWithDefaults

`func NewAddUnboundidMsChapV2SaslMechanismHandlerRequestWithDefaults() *AddUnboundidMsChapV2SaslMechanismHandlerRequest`

NewAddUnboundidMsChapV2SaslMechanismHandlerRequestWithDefaults instantiates a new AddUnboundidMsChapV2SaslMechanismHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetSchemas() []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetSchemasOk() (*[]EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetSchemas(v []EnumunboundidMsChapV2SaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetDescription

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUnboundidMsChapV2SaslMechanismHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


