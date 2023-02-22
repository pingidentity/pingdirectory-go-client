# AddStringArrayTokenClaimValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationName** | **string** | Name of the new Token Claim Validation | 
**Schemas** | [**[]EnumstringArrayTokenClaimValidationSchemaUrn**](EnumstringArrayTokenClaimValidationSchemaUrn.md) |  | 
**AllRequiredValue** | Pointer to **[]string** | The set of all values that the claim must have to be considered valid. | [optional] 
**AnyRequiredValue** | Pointer to **[]string** | The set of values that the claim may have to be considered valid. | [optional] 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 

## Methods

### NewAddStringArrayTokenClaimValidationRequest

`func NewAddStringArrayTokenClaimValidationRequest(validationName string, schemas []EnumstringArrayTokenClaimValidationSchemaUrn, claimName string, ) *AddStringArrayTokenClaimValidationRequest`

NewAddStringArrayTokenClaimValidationRequest instantiates a new AddStringArrayTokenClaimValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStringArrayTokenClaimValidationRequestWithDefaults

`func NewAddStringArrayTokenClaimValidationRequestWithDefaults() *AddStringArrayTokenClaimValidationRequest`

NewAddStringArrayTokenClaimValidationRequestWithDefaults instantiates a new AddStringArrayTokenClaimValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationName

`func (o *AddStringArrayTokenClaimValidationRequest) GetValidationName() string`

GetValidationName returns the ValidationName field if non-nil, zero value otherwise.

### GetValidationNameOk

`func (o *AddStringArrayTokenClaimValidationRequest) GetValidationNameOk() (*string, bool)`

GetValidationNameOk returns a tuple with the ValidationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationName

`func (o *AddStringArrayTokenClaimValidationRequest) SetValidationName(v string)`

SetValidationName sets ValidationName field to given value.


### GetSchemas

`func (o *AddStringArrayTokenClaimValidationRequest) GetSchemas() []EnumstringArrayTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddStringArrayTokenClaimValidationRequest) GetSchemasOk() (*[]EnumstringArrayTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddStringArrayTokenClaimValidationRequest) SetSchemas(v []EnumstringArrayTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllRequiredValue

`func (o *AddStringArrayTokenClaimValidationRequest) GetAllRequiredValue() []string`

GetAllRequiredValue returns the AllRequiredValue field if non-nil, zero value otherwise.

### GetAllRequiredValueOk

`func (o *AddStringArrayTokenClaimValidationRequest) GetAllRequiredValueOk() (*[]string, bool)`

GetAllRequiredValueOk returns a tuple with the AllRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredValue

`func (o *AddStringArrayTokenClaimValidationRequest) SetAllRequiredValue(v []string)`

SetAllRequiredValue sets AllRequiredValue field to given value.

### HasAllRequiredValue

`func (o *AddStringArrayTokenClaimValidationRequest) HasAllRequiredValue() bool`

HasAllRequiredValue returns a boolean if a field has been set.

### GetAnyRequiredValue

`func (o *AddStringArrayTokenClaimValidationRequest) GetAnyRequiredValue() []string`

GetAnyRequiredValue returns the AnyRequiredValue field if non-nil, zero value otherwise.

### GetAnyRequiredValueOk

`func (o *AddStringArrayTokenClaimValidationRequest) GetAnyRequiredValueOk() (*[]string, bool)`

GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredValue

`func (o *AddStringArrayTokenClaimValidationRequest) SetAnyRequiredValue(v []string)`

SetAnyRequiredValue sets AnyRequiredValue field to given value.

### HasAnyRequiredValue

`func (o *AddStringArrayTokenClaimValidationRequest) HasAnyRequiredValue() bool`

HasAnyRequiredValue returns a boolean if a field has been set.

### GetDescription

`func (o *AddStringArrayTokenClaimValidationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStringArrayTokenClaimValidationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStringArrayTokenClaimValidationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStringArrayTokenClaimValidationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *AddStringArrayTokenClaimValidationRequest) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddStringArrayTokenClaimValidationRequest) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddStringArrayTokenClaimValidationRequest) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


