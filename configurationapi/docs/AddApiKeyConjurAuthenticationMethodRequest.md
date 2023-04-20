# AddApiKeyConjurAuthenticationMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodName** | **string** | Name of the new Conjur Authentication Method | 
**Schemas** | [**[]EnumapiKeyConjurAuthenticationMethodSchemaUrn**](EnumapiKeyConjurAuthenticationMethodSchemaUrn.md) |  | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | Pointer to **string** | The password for the user to authenticate. This will be used to obtain an API key for the target user. | [optional] 
**ApiKey** | Pointer to **string** | The API key for the user to authenticate. | [optional] 
**Description** | Pointer to **string** | A description for this Conjur Authentication Method | [optional] 

## Methods

### NewAddApiKeyConjurAuthenticationMethodRequest

`func NewAddApiKeyConjurAuthenticationMethodRequest(methodName string, schemas []EnumapiKeyConjurAuthenticationMethodSchemaUrn, username string, ) *AddApiKeyConjurAuthenticationMethodRequest`

NewAddApiKeyConjurAuthenticationMethodRequest instantiates a new AddApiKeyConjurAuthenticationMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddApiKeyConjurAuthenticationMethodRequestWithDefaults

`func NewAddApiKeyConjurAuthenticationMethodRequestWithDefaults() *AddApiKeyConjurAuthenticationMethodRequest`

NewAddApiKeyConjurAuthenticationMethodRequestWithDefaults instantiates a new AddApiKeyConjurAuthenticationMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodName

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *AddApiKeyConjurAuthenticationMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetSchemas

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetSchemas() []EnumapiKeyConjurAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetSchemasOk() (*[]EnumapiKeyConjurAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddApiKeyConjurAuthenticationMethodRequest) SetSchemas(v []EnumapiKeyConjurAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetUsername

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddApiKeyConjurAuthenticationMethodRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddApiKeyConjurAuthenticationMethodRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddApiKeyConjurAuthenticationMethodRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiKey

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AddApiKeyConjurAuthenticationMethodRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AddApiKeyConjurAuthenticationMethodRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDescription

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddApiKeyConjurAuthenticationMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddApiKeyConjurAuthenticationMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddApiKeyConjurAuthenticationMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


