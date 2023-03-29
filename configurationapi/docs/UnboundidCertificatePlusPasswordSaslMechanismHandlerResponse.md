# UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn**](EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the SASL Mechanism Handler | 
**CertificateMapper** | **string** | The certificate mapper that will be used to identify the target user based on the certificate that was presented to the server. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse

`func NewUnboundidCertificatePlusPasswordSaslMechanismHandlerResponse(schemas []EnumunboundidCertificatePlusPasswordSaslMechanismHandlerSchemaUrn, id string, certificateMapper string, enabled bool, ) *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse`

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


### GetId

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


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


### GetMeta

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidCertificatePlusPasswordSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


