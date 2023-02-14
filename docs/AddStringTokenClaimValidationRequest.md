# AddStringTokenClaimValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationName** | **string** | Name of the new Token Claim Validation | 
**Schemas** | [**[]EnumstringTokenClaimValidationSchemaUrn**](EnumstringTokenClaimValidationSchemaUrn.md) |  | 
**AnyRequiredValue** | **[]string** | The set of values that the claim may have to be considered valid. | 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 

## Methods

### NewAddStringTokenClaimValidationRequest

`func NewAddStringTokenClaimValidationRequest(validationName string, schemas []EnumstringTokenClaimValidationSchemaUrn, anyRequiredValue []string, claimName string, ) *AddStringTokenClaimValidationRequest`

NewAddStringTokenClaimValidationRequest instantiates a new AddStringTokenClaimValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStringTokenClaimValidationRequestWithDefaults

`func NewAddStringTokenClaimValidationRequestWithDefaults() *AddStringTokenClaimValidationRequest`

NewAddStringTokenClaimValidationRequestWithDefaults instantiates a new AddStringTokenClaimValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationName

`func (o *AddStringTokenClaimValidationRequest) GetValidationName() string`

GetValidationName returns the ValidationName field if non-nil, zero value otherwise.

### GetValidationNameOk

`func (o *AddStringTokenClaimValidationRequest) GetValidationNameOk() (*string, bool)`

GetValidationNameOk returns a tuple with the ValidationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationName

`func (o *AddStringTokenClaimValidationRequest) SetValidationName(v string)`

SetValidationName sets ValidationName field to given value.


### GetSchemas

`func (o *AddStringTokenClaimValidationRequest) GetSchemas() []EnumstringTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddStringTokenClaimValidationRequest) GetSchemasOk() (*[]EnumstringTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddStringTokenClaimValidationRequest) SetSchemas(v []EnumstringTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAnyRequiredValue

`func (o *AddStringTokenClaimValidationRequest) GetAnyRequiredValue() []string`

GetAnyRequiredValue returns the AnyRequiredValue field if non-nil, zero value otherwise.

### GetAnyRequiredValueOk

`func (o *AddStringTokenClaimValidationRequest) GetAnyRequiredValueOk() (*[]string, bool)`

GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredValue

`func (o *AddStringTokenClaimValidationRequest) SetAnyRequiredValue(v []string)`

SetAnyRequiredValue sets AnyRequiredValue field to given value.


### GetDescription

`func (o *AddStringTokenClaimValidationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStringTokenClaimValidationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStringTokenClaimValidationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStringTokenClaimValidationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *AddStringTokenClaimValidationRequest) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddStringTokenClaimValidationRequest) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddStringTokenClaimValidationRequest) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


