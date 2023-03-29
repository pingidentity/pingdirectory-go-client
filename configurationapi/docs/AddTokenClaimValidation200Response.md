# AddTokenClaimValidation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the ID Token Validator | 
**Schemas** | [**[]EnumstringTokenClaimValidationSchemaUrn**](EnumstringTokenClaimValidationSchemaUrn.md) |  | 
**AllRequiredValue** | Pointer to **[]string** | The set of all values that the claim must have to be considered valid. | [optional] 
**AnyRequiredValue** | **[]string** | The set of values that the claim may have to be considered valid. | 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**RequiredValue** | [**EnumtokenClaimValidationRequiredValueProp**](EnumtokenClaimValidationRequiredValueProp.md) |  | 

## Methods

### NewAddTokenClaimValidation200Response

`func NewAddTokenClaimValidation200Response(id string, schemas []EnumstringTokenClaimValidationSchemaUrn, anyRequiredValue []string, claimName string, requiredValue EnumtokenClaimValidationRequiredValueProp, ) *AddTokenClaimValidation200Response`

NewAddTokenClaimValidation200Response instantiates a new AddTokenClaimValidation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTokenClaimValidation200ResponseWithDefaults

`func NewAddTokenClaimValidation200ResponseWithDefaults() *AddTokenClaimValidation200Response`

NewAddTokenClaimValidation200ResponseWithDefaults instantiates a new AddTokenClaimValidation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTokenClaimValidation200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTokenClaimValidation200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTokenClaimValidation200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddTokenClaimValidation200Response) GetSchemas() []EnumstringTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTokenClaimValidation200Response) GetSchemasOk() (*[]EnumstringTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTokenClaimValidation200Response) SetSchemas(v []EnumstringTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllRequiredValue

`func (o *AddTokenClaimValidation200Response) GetAllRequiredValue() []string`

GetAllRequiredValue returns the AllRequiredValue field if non-nil, zero value otherwise.

### GetAllRequiredValueOk

`func (o *AddTokenClaimValidation200Response) GetAllRequiredValueOk() (*[]string, bool)`

GetAllRequiredValueOk returns a tuple with the AllRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredValue

`func (o *AddTokenClaimValidation200Response) SetAllRequiredValue(v []string)`

SetAllRequiredValue sets AllRequiredValue field to given value.

### HasAllRequiredValue

`func (o *AddTokenClaimValidation200Response) HasAllRequiredValue() bool`

HasAllRequiredValue returns a boolean if a field has been set.

### GetAnyRequiredValue

`func (o *AddTokenClaimValidation200Response) GetAnyRequiredValue() []string`

GetAnyRequiredValue returns the AnyRequiredValue field if non-nil, zero value otherwise.

### GetAnyRequiredValueOk

`func (o *AddTokenClaimValidation200Response) GetAnyRequiredValueOk() (*[]string, bool)`

GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredValue

`func (o *AddTokenClaimValidation200Response) SetAnyRequiredValue(v []string)`

SetAnyRequiredValue sets AnyRequiredValue field to given value.


### GetDescription

`func (o *AddTokenClaimValidation200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTokenClaimValidation200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTokenClaimValidation200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTokenClaimValidation200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *AddTokenClaimValidation200Response) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddTokenClaimValidation200Response) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddTokenClaimValidation200Response) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetMeta

`func (o *AddTokenClaimValidation200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddTokenClaimValidation200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddTokenClaimValidation200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddTokenClaimValidation200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddTokenClaimValidation200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddTokenClaimValidation200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddTokenClaimValidation200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddTokenClaimValidation200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetRequiredValue

`func (o *AddTokenClaimValidation200Response) GetRequiredValue() EnumtokenClaimValidationRequiredValueProp`

GetRequiredValue returns the RequiredValue field if non-nil, zero value otherwise.

### GetRequiredValueOk

`func (o *AddTokenClaimValidation200Response) GetRequiredValueOk() (*EnumtokenClaimValidationRequiredValueProp, bool)`

GetRequiredValueOk returns a tuple with the RequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredValue

`func (o *AddTokenClaimValidation200Response) SetRequiredValue(v EnumtokenClaimValidationRequiredValueProp)`

SetRequiredValue sets RequiredValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


