# BooleanTokenClaimValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Token Claim Validation | 
**Schemas** | [**[]EnumbooleanTokenClaimValidationSchemaUrn**](EnumbooleanTokenClaimValidationSchemaUrn.md) |  | 
**RequiredValue** | [**EnumtokenClaimValidationRequiredValueProp**](EnumtokenClaimValidationRequiredValueProp.md) |  | 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewBooleanTokenClaimValidationResponse

`func NewBooleanTokenClaimValidationResponse(id string, schemas []EnumbooleanTokenClaimValidationSchemaUrn, requiredValue EnumtokenClaimValidationRequiredValueProp, claimName string, ) *BooleanTokenClaimValidationResponse`

NewBooleanTokenClaimValidationResponse instantiates a new BooleanTokenClaimValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanTokenClaimValidationResponseWithDefaults

`func NewBooleanTokenClaimValidationResponseWithDefaults() *BooleanTokenClaimValidationResponse`

NewBooleanTokenClaimValidationResponseWithDefaults instantiates a new BooleanTokenClaimValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BooleanTokenClaimValidationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BooleanTokenClaimValidationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BooleanTokenClaimValidationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *BooleanTokenClaimValidationResponse) GetSchemas() []EnumbooleanTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BooleanTokenClaimValidationResponse) GetSchemasOk() (*[]EnumbooleanTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BooleanTokenClaimValidationResponse) SetSchemas(v []EnumbooleanTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequiredValue

`func (o *BooleanTokenClaimValidationResponse) GetRequiredValue() EnumtokenClaimValidationRequiredValueProp`

GetRequiredValue returns the RequiredValue field if non-nil, zero value otherwise.

### GetRequiredValueOk

`func (o *BooleanTokenClaimValidationResponse) GetRequiredValueOk() (*EnumtokenClaimValidationRequiredValueProp, bool)`

GetRequiredValueOk returns a tuple with the RequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredValue

`func (o *BooleanTokenClaimValidationResponse) SetRequiredValue(v EnumtokenClaimValidationRequiredValueProp)`

SetRequiredValue sets RequiredValue field to given value.


### GetDescription

`func (o *BooleanTokenClaimValidationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BooleanTokenClaimValidationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BooleanTokenClaimValidationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BooleanTokenClaimValidationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *BooleanTokenClaimValidationResponse) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *BooleanTokenClaimValidationResponse) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *BooleanTokenClaimValidationResponse) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetMeta

`func (o *BooleanTokenClaimValidationResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BooleanTokenClaimValidationResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BooleanTokenClaimValidationResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BooleanTokenClaimValidationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *BooleanTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *BooleanTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *BooleanTokenClaimValidationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *BooleanTokenClaimValidationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


