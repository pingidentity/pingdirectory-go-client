# UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn**](EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn.md) |  | 
**CertificateMapper** | **string** | The certificate mapper that will be used to identify the target user based on the certificate that was presented to the server. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse

`func NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse(schemas []EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn, certificateMapper string, enabled bool, ) *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse`

NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse instantiates a new UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponseWithDefaults

`func NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponseWithDefaults() *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse`

NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCertificateMapper

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetCertificateMapper() string`

GetCertificateMapper returns the CertificateMapper field if non-nil, zero value otherwise.

### GetCertificateMapperOk

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetCertificateMapperOk() (*string, bool)`

GetCertificateMapperOk returns a tuple with the CertificateMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateMapper

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetCertificateMapper(v string)`

SetCertificateMapper sets CertificateMapper field to given value.


### GetDescription

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


