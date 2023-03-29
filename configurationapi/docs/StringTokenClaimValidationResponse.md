# StringTokenClaimValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the ID Token Validator | 
**Schemas** | [**[]EnumstringTokenClaimValidationSchemaUrn**](EnumstringTokenClaimValidationSchemaUrn.md) |  | 
**AnyRequiredValue** | **[]string** | The set of values that the claim may have to be considered valid. | 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewStringTokenClaimValidationResponse

`func NewStringTokenClaimValidationResponse(id string, schemas []EnumstringTokenClaimValidationSchemaUrn, anyRequiredValue []string, claimName string, ) *StringTokenClaimValidationResponse`

NewStringTokenClaimValidationResponse instantiates a new StringTokenClaimValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringTokenClaimValidationResponseWithDefaults

`func NewStringTokenClaimValidationResponseWithDefaults() *StringTokenClaimValidationResponse`

NewStringTokenClaimValidationResponseWithDefaults instantiates a new StringTokenClaimValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StringTokenClaimValidationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StringTokenClaimValidationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StringTokenClaimValidationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *StringTokenClaimValidationResponse) GetSchemas() []EnumstringTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StringTokenClaimValidationResponse) GetSchemasOk() (*[]EnumstringTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StringTokenClaimValidationResponse) SetSchemas(v []EnumstringTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAnyRequiredValue

`func (o *StringTokenClaimValidationResponse) GetAnyRequiredValue() []string`

GetAnyRequiredValue returns the AnyRequiredValue field if non-nil, zero value otherwise.

### GetAnyRequiredValueOk

`func (o *StringTokenClaimValidationResponse) GetAnyRequiredValueOk() (*[]string, bool)`

GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredValue

`func (o *StringTokenClaimValidationResponse) SetAnyRequiredValue(v []string)`

SetAnyRequiredValue sets AnyRequiredValue field to given value.


### GetDescription

`func (o *StringTokenClaimValidationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringTokenClaimValidationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringTokenClaimValidationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringTokenClaimValidationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *StringTokenClaimValidationResponse) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *StringTokenClaimValidationResponse) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *StringTokenClaimValidationResponse) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetMeta

`func (o *StringTokenClaimValidationResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StringTokenClaimValidationResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StringTokenClaimValidationResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StringTokenClaimValidationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *StringTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *StringTokenClaimValidationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *StringTokenClaimValidationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *StringTokenClaimValidationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


