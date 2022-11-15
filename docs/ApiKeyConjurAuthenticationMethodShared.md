# ApiKeyConjurAuthenticationMethodShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumapiKeyConjurAuthenticationMethodSchemaUrn**](EnumapiKeyConjurAuthenticationMethodSchemaUrn.md) |  | [optional] 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | Pointer to **string** | The password for the user to authenticate. This will be used to obtain an API key for the target user. | [optional] 
**ApiKey** | Pointer to **string** | The API key for the user to authenticate. | [optional] 
**Description** | Pointer to **string** | A description for this Conjur Authentication Method | [optional] 

## Methods

### NewApiKeyConjurAuthenticationMethodShared

`func NewApiKeyConjurAuthenticationMethodShared(username string, ) *ApiKeyConjurAuthenticationMethodShared`

NewApiKeyConjurAuthenticationMethodShared instantiates a new ApiKeyConjurAuthenticationMethodShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyConjurAuthenticationMethodSharedWithDefaults

`func NewApiKeyConjurAuthenticationMethodSharedWithDefaults() *ApiKeyConjurAuthenticationMethodShared`

NewApiKeyConjurAuthenticationMethodSharedWithDefaults instantiates a new ApiKeyConjurAuthenticationMethodShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ApiKeyConjurAuthenticationMethodShared) GetSchemas() []EnumapiKeyConjurAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ApiKeyConjurAuthenticationMethodShared) GetSchemasOk() (*[]EnumapiKeyConjurAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ApiKeyConjurAuthenticationMethodShared) SetSchemas(v []EnumapiKeyConjurAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ApiKeyConjurAuthenticationMethodShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetUsername

`func (o *ApiKeyConjurAuthenticationMethodShared) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiKeyConjurAuthenticationMethodShared) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiKeyConjurAuthenticationMethodShared) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ApiKeyConjurAuthenticationMethodShared) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiKeyConjurAuthenticationMethodShared) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiKeyConjurAuthenticationMethodShared) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiKeyConjurAuthenticationMethodShared) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiKey

`func (o *ApiKeyConjurAuthenticationMethodShared) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiKeyConjurAuthenticationMethodShared) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiKeyConjurAuthenticationMethodShared) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ApiKeyConjurAuthenticationMethodShared) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyConjurAuthenticationMethodShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyConjurAuthenticationMethodShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyConjurAuthenticationMethodShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyConjurAuthenticationMethodShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


