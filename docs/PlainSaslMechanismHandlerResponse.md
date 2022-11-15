# PlainSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumplainSaslMechanismHandlerSchemaUrn**](EnumplainSaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to match the authentication or authorization ID included in the SASL bind request to the corresponding user in the directory. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewPlainSaslMechanismHandlerResponse

`func NewPlainSaslMechanismHandlerResponse(schemas []EnumplainSaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, ) *PlainSaslMechanismHandlerResponse`

NewPlainSaslMechanismHandlerResponse instantiates a new PlainSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlainSaslMechanismHandlerResponseWithDefaults

`func NewPlainSaslMechanismHandlerResponseWithDefaults() *PlainSaslMechanismHandlerResponse`

NewPlainSaslMechanismHandlerResponseWithDefaults instantiates a new PlainSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PlainSaslMechanismHandlerResponse) GetSchemas() []EnumplainSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PlainSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumplainSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PlainSaslMechanismHandlerResponse) SetSchemas(v []EnumplainSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *PlainSaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *PlainSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *PlainSaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetDescription

`func (o *PlainSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlainSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlainSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlainSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PlainSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlainSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlainSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


