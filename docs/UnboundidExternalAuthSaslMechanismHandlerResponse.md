# UnboundidExternalAuthSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumunboundidExternalAuthSaslMechanismHandlerSchemaUrn**](EnumunboundidExternalAuthSaslMechanismHandlerSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the SASL Mechanism Handler | [optional] 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user targeted by the authentication ID contained in the bind request. This will only be used for \&quot;u:\&quot;-style authentication ID values. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewUnboundidExternalAuthSaslMechanismHandlerResponse

`func NewUnboundidExternalAuthSaslMechanismHandlerResponse(schemas []EnumunboundidExternalAuthSaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, ) *UnboundidExternalAuthSaslMechanismHandlerResponse`

NewUnboundidExternalAuthSaslMechanismHandlerResponse instantiates a new UnboundidExternalAuthSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidExternalAuthSaslMechanismHandlerResponseWithDefaults

`func NewUnboundidExternalAuthSaslMechanismHandlerResponseWithDefaults() *UnboundidExternalAuthSaslMechanismHandlerResponse`

NewUnboundidExternalAuthSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidExternalAuthSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidExternalAuthSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumunboundidExternalAuthSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidExternalAuthSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetDescription

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidExternalAuthSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


