# StringArrayTokenClaimValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Token Claim Validation | 
**Schemas** | [**[]EnumstringArrayTokenClaimValidationSchemaUrn**](EnumstringArrayTokenClaimValidationSchemaUrn.md) |  | 
**AllRequiredValue** | Pointer to **[]string** | The set of all values that the claim must have to be considered valid. | [optional] 
**AnyRequiredValue** | Pointer to **[]string** | The set of values that the claim may have to be considered valid. | [optional] 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewStringArrayTokenClaimValidationResponse

`func NewStringArrayTokenClaimValidationResponse(id string, schemas []EnumstringArrayTokenClaimValidationSchemaUrn, claimName string, ) *StringArrayTokenClaimValidationResponse`

NewStringArrayTokenClaimValidationResponse instantiates a new StringArrayTokenClaimValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringArrayTokenClaimValidationResponseWithDefaults

`func NewStringArrayTokenClaimValidationResponseWithDefaults() *StringArrayTokenClaimValidationResponse`

NewStringArrayTokenClaimValidationResponseWithDefaults instantiates a new StringArrayTokenClaimValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StringArrayTokenClaimValidationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StringArrayTokenClaimValidationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StringArrayTokenClaimValidationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *StringArrayTokenClaimValidationResponse) GetSchemas() []EnumstringArrayTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StringArrayTokenClaimValidationResponse) GetSchemasOk() (*[]EnumstringArrayTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StringArrayTokenClaimValidationResponse) SetSchemas(v []EnumstringArrayTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllRequiredValue

`func (o *StringArrayTokenClaimValidationResponse) GetAllRequiredValue() []string`

GetAllRequiredValue returns the AllRequiredValue field if non-nil, zero value otherwise.

### GetAllRequiredValueOk

`func (o *StringArrayTokenClaimValidationResponse) GetAllRequiredValueOk() (*[]string, bool)`

GetAllRequiredValueOk returns a tuple with the AllRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredValue

`func (o *StringArrayTokenClaimValidationResponse) SetAllRequiredValue(v []string)`

SetAllRequiredValue sets AllRequiredValue field to given value.

### HasAllRequiredValue

`func (o *StringArrayTokenClaimValidationResponse) HasAllRequiredValue() bool`

HasAllRequiredValue returns a boolean if a field has been set.

### GetAnyRequiredValue

`func (o *StringArrayTokenClaimValidationResponse) GetAnyRequiredValue() []string`

GetAnyRequiredValue returns the AnyRequiredValue field if non-nil, zero value otherwise.

### GetAnyRequiredValueOk

`func (o *StringArrayTokenClaimValidationResponse) GetAnyRequiredValueOk() (*[]string, bool)`

GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredValue

`func (o *StringArrayTokenClaimValidationResponse) SetAnyRequiredValue(v []string)`

SetAnyRequiredValue sets AnyRequiredValue field to given value.

### HasAnyRequiredValue

`func (o *StringArrayTokenClaimValidationResponse) HasAnyRequiredValue() bool`

HasAnyRequiredValue returns a boolean if a field has been set.

### GetDescription

`func (o *StringArrayTokenClaimValidationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringArrayTokenClaimValidationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringArrayTokenClaimValidationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringArrayTokenClaimValidationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *StringArrayTokenClaimValidationResponse) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *StringArrayTokenClaimValidationResponse) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *StringArrayTokenClaimValidationResponse) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetMeta

`func (o *StringArrayTokenClaimValidationResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StringArrayTokenClaimValidationResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StringArrayTokenClaimValidationResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StringArrayTokenClaimValidationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *StringArrayTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *StringArrayTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *StringArrayTokenClaimValidationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *StringArrayTokenClaimValidationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


