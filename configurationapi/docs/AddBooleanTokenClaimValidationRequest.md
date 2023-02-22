# AddBooleanTokenClaimValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationName** | **string** | Name of the new Token Claim Validation | 
**Schemas** | [**[]EnumbooleanTokenClaimValidationSchemaUrn**](EnumbooleanTokenClaimValidationSchemaUrn.md) |  | 
**RequiredValue** | [**EnumtokenClaimValidationRequiredValueProp**](EnumtokenClaimValidationRequiredValueProp.md) |  | 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 

## Methods

### NewAddBooleanTokenClaimValidationRequest

`func NewAddBooleanTokenClaimValidationRequest(validationName string, schemas []EnumbooleanTokenClaimValidationSchemaUrn, requiredValue EnumtokenClaimValidationRequiredValueProp, claimName string, ) *AddBooleanTokenClaimValidationRequest`

NewAddBooleanTokenClaimValidationRequest instantiates a new AddBooleanTokenClaimValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBooleanTokenClaimValidationRequestWithDefaults

`func NewAddBooleanTokenClaimValidationRequestWithDefaults() *AddBooleanTokenClaimValidationRequest`

NewAddBooleanTokenClaimValidationRequestWithDefaults instantiates a new AddBooleanTokenClaimValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationName

`func (o *AddBooleanTokenClaimValidationRequest) GetValidationName() string`

GetValidationName returns the ValidationName field if non-nil, zero value otherwise.

### GetValidationNameOk

`func (o *AddBooleanTokenClaimValidationRequest) GetValidationNameOk() (*string, bool)`

GetValidationNameOk returns a tuple with the ValidationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationName

`func (o *AddBooleanTokenClaimValidationRequest) SetValidationName(v string)`

SetValidationName sets ValidationName field to given value.


### GetSchemas

`func (o *AddBooleanTokenClaimValidationRequest) GetSchemas() []EnumbooleanTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddBooleanTokenClaimValidationRequest) GetSchemasOk() (*[]EnumbooleanTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddBooleanTokenClaimValidationRequest) SetSchemas(v []EnumbooleanTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequiredValue

`func (o *AddBooleanTokenClaimValidationRequest) GetRequiredValue() EnumtokenClaimValidationRequiredValueProp`

GetRequiredValue returns the RequiredValue field if non-nil, zero value otherwise.

### GetRequiredValueOk

`func (o *AddBooleanTokenClaimValidationRequest) GetRequiredValueOk() (*EnumtokenClaimValidationRequiredValueProp, bool)`

GetRequiredValueOk returns a tuple with the RequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredValue

`func (o *AddBooleanTokenClaimValidationRequest) SetRequiredValue(v EnumtokenClaimValidationRequiredValueProp)`

SetRequiredValue sets RequiredValue field to given value.


### GetDescription

`func (o *AddBooleanTokenClaimValidationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBooleanTokenClaimValidationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBooleanTokenClaimValidationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBooleanTokenClaimValidationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *AddBooleanTokenClaimValidationRequest) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddBooleanTokenClaimValidationRequest) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddBooleanTokenClaimValidationRequest) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


