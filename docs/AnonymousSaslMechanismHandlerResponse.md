# AnonymousSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumanonymousSaslMechanismHandlerSchemaUrn**](EnumanonymousSaslMechanismHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewAnonymousSaslMechanismHandlerResponse

`func NewAnonymousSaslMechanismHandlerResponse(schemas []EnumanonymousSaslMechanismHandlerSchemaUrn, enabled bool, ) *AnonymousSaslMechanismHandlerResponse`

NewAnonymousSaslMechanismHandlerResponse instantiates a new AnonymousSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymousSaslMechanismHandlerResponseWithDefaults

`func NewAnonymousSaslMechanismHandlerResponseWithDefaults() *AnonymousSaslMechanismHandlerResponse`

NewAnonymousSaslMechanismHandlerResponseWithDefaults instantiates a new AnonymousSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AnonymousSaslMechanismHandlerResponse) GetSchemas() []EnumanonymousSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AnonymousSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumanonymousSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AnonymousSaslMechanismHandlerResponse) SetSchemas(v []EnumanonymousSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AnonymousSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnonymousSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnonymousSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnonymousSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AnonymousSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AnonymousSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AnonymousSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


