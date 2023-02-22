# AddTokenClaimValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationName** | **string** | Name of the new Token Claim Validation | 
**Schemas** | [**[]EnumstringTokenClaimValidationSchemaUrn**](EnumstringTokenClaimValidationSchemaUrn.md) |  | 
**AllRequiredValue** | Pointer to **[]string** | The set of all values that the claim must have to be considered valid. | [optional] 
**AnyRequiredValue** | **[]string** | The set of values that the claim may have to be considered valid. | 
**Description** | Pointer to **string** | A description for this Token Claim Validation | [optional] 
**ClaimName** | **string** | The name of the claim to be validated. | 
**RequiredValue** | [**EnumtokenClaimValidationRequiredValueProp**](EnumtokenClaimValidationRequiredValueProp.md) |  | 

## Methods

### NewAddTokenClaimValidationRequest

`func NewAddTokenClaimValidationRequest(validationName string, schemas []EnumstringTokenClaimValidationSchemaUrn, anyRequiredValue []string, claimName string, requiredValue EnumtokenClaimValidationRequiredValueProp, ) *AddTokenClaimValidationRequest`

NewAddTokenClaimValidationRequest instantiates a new AddTokenClaimValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTokenClaimValidationRequestWithDefaults

`func NewAddTokenClaimValidationRequestWithDefaults() *AddTokenClaimValidationRequest`

NewAddTokenClaimValidationRequestWithDefaults instantiates a new AddTokenClaimValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationName

`func (o *AddTokenClaimValidationRequest) GetValidationName() string`

GetValidationName returns the ValidationName field if non-nil, zero value otherwise.

### GetValidationNameOk

`func (o *AddTokenClaimValidationRequest) GetValidationNameOk() (*string, bool)`

GetValidationNameOk returns a tuple with the ValidationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationName

`func (o *AddTokenClaimValidationRequest) SetValidationName(v string)`

SetValidationName sets ValidationName field to given value.


### GetSchemas

`func (o *AddTokenClaimValidationRequest) GetSchemas() []EnumstringTokenClaimValidationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTokenClaimValidationRequest) GetSchemasOk() (*[]EnumstringTokenClaimValidationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTokenClaimValidationRequest) SetSchemas(v []EnumstringTokenClaimValidationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllRequiredValue

`func (o *AddTokenClaimValidationRequest) GetAllRequiredValue() []string`

GetAllRequiredValue returns the AllRequiredValue field if non-nil, zero value otherwise.

### GetAllRequiredValueOk

`func (o *AddTokenClaimValidationRequest) GetAllRequiredValueOk() (*[]string, bool)`

GetAllRequiredValueOk returns a tuple with the AllRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredValue

`func (o *AddTokenClaimValidationRequest) SetAllRequiredValue(v []string)`

SetAllRequiredValue sets AllRequiredValue field to given value.

### HasAllRequiredValue

`func (o *AddTokenClaimValidationRequest) HasAllRequiredValue() bool`

HasAllRequiredValue returns a boolean if a field has been set.

### GetAnyRequiredValue

`func (o *AddTokenClaimValidationRequest) GetAnyRequiredValue() []string`

GetAnyRequiredValue returns the AnyRequiredValue field if non-nil, zero value otherwise.

### GetAnyRequiredValueOk

`func (o *AddTokenClaimValidationRequest) GetAnyRequiredValueOk() (*[]string, bool)`

GetAnyRequiredValueOk returns a tuple with the AnyRequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredValue

`func (o *AddTokenClaimValidationRequest) SetAnyRequiredValue(v []string)`

SetAnyRequiredValue sets AnyRequiredValue field to given value.


### GetDescription

`func (o *AddTokenClaimValidationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTokenClaimValidationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTokenClaimValidationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTokenClaimValidationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClaimName

`func (o *AddTokenClaimValidationRequest) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddTokenClaimValidationRequest) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddTokenClaimValidationRequest) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetRequiredValue

`func (o *AddTokenClaimValidationRequest) GetRequiredValue() EnumtokenClaimValidationRequiredValueProp`

GetRequiredValue returns the RequiredValue field if non-nil, zero value otherwise.

### GetRequiredValueOk

`func (o *AddTokenClaimValidationRequest) GetRequiredValueOk() (*EnumtokenClaimValidationRequiredValueProp, bool)`

GetRequiredValueOk returns a tuple with the RequiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredValue

`func (o *AddTokenClaimValidationRequest) SetRequiredValue(v EnumtokenClaimValidationRequiredValueProp)`

SetRequiredValue sets RequiredValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


