# ExternalSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumexternalSaslMechanismHandlerSchemaUrn**](EnumexternalSaslMechanismHandlerSchemaUrn.md) |  | 
**CertificateValidationPolicy** | [**EnumsaslMechanismHandlerCertificateValidationPolicyProp**](EnumsaslMechanismHandlerCertificateValidationPolicyProp.md) |  | 
**CertificateAttribute** | Pointer to **string** | Specifies the name of the attribute to hold user certificates. | [optional] 
**CertificateMapper** | **string** | Specifies the name of the certificate mapper that should be used to match client certificates to user entries. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewExternalSaslMechanismHandlerResponse

`func NewExternalSaslMechanismHandlerResponse(schemas []EnumexternalSaslMechanismHandlerSchemaUrn, certificateValidationPolicy EnumsaslMechanismHandlerCertificateValidationPolicyProp, certificateMapper string, enabled bool, ) *ExternalSaslMechanismHandlerResponse`

NewExternalSaslMechanismHandlerResponse instantiates a new ExternalSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSaslMechanismHandlerResponseWithDefaults

`func NewExternalSaslMechanismHandlerResponseWithDefaults() *ExternalSaslMechanismHandlerResponse`

NewExternalSaslMechanismHandlerResponseWithDefaults instantiates a new ExternalSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExternalSaslMechanismHandlerResponse) GetSchemas() []EnumexternalSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExternalSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumexternalSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExternalSaslMechanismHandlerResponse) SetSchemas(v []EnumexternalSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCertificateValidationPolicy

`func (o *ExternalSaslMechanismHandlerResponse) GetCertificateValidationPolicy() EnumsaslMechanismHandlerCertificateValidationPolicyProp`

GetCertificateValidationPolicy returns the CertificateValidationPolicy field if non-nil, zero value otherwise.

### GetCertificateValidationPolicyOk

`func (o *ExternalSaslMechanismHandlerResponse) GetCertificateValidationPolicyOk() (*EnumsaslMechanismHandlerCertificateValidationPolicyProp, bool)`

GetCertificateValidationPolicyOk returns a tuple with the CertificateValidationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidationPolicy

`func (o *ExternalSaslMechanismHandlerResponse) SetCertificateValidationPolicy(v EnumsaslMechanismHandlerCertificateValidationPolicyProp)`

SetCertificateValidationPolicy sets CertificateValidationPolicy field to given value.


### GetCertificateAttribute

`func (o *ExternalSaslMechanismHandlerResponse) GetCertificateAttribute() string`

GetCertificateAttribute returns the CertificateAttribute field if non-nil, zero value otherwise.

### GetCertificateAttributeOk

`func (o *ExternalSaslMechanismHandlerResponse) GetCertificateAttributeOk() (*string, bool)`

GetCertificateAttributeOk returns a tuple with the CertificateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAttribute

`func (o *ExternalSaslMechanismHandlerResponse) SetCertificateAttribute(v string)`

SetCertificateAttribute sets CertificateAttribute field to given value.

### HasCertificateAttribute

`func (o *ExternalSaslMechanismHandlerResponse) HasCertificateAttribute() bool`

HasCertificateAttribute returns a boolean if a field has been set.

### GetCertificateMapper

`func (o *ExternalSaslMechanismHandlerResponse) GetCertificateMapper() string`

GetCertificateMapper returns the CertificateMapper field if non-nil, zero value otherwise.

### GetCertificateMapperOk

`func (o *ExternalSaslMechanismHandlerResponse) GetCertificateMapperOk() (*string, bool)`

GetCertificateMapperOk returns a tuple with the CertificateMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateMapper

`func (o *ExternalSaslMechanismHandlerResponse) SetCertificateMapper(v string)`

SetCertificateMapper sets CertificateMapper field to given value.


### GetDescription

`func (o *ExternalSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ExternalSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExternalSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExternalSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


