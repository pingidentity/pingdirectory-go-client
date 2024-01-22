# ApiKeyConjurAuthenticationMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumapiKeyConjurAuthenticationMethodSchemaUrn**](EnumapiKeyConjurAuthenticationMethodSchemaUrn.md) |  | 
**Username** | **string** | The username for the user to authenticate. | 
**Password** | Pointer to **string** | The password for the user to authenticate. This will be used to obtain an API key for the target user. | [optional] 
**ApiKey** | Pointer to **string** | The API key for the user to authenticate. | [optional] 
**Description** | Pointer to **string** | A description for this Conjur Authentication Method | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Conjur Authentication Method | 

## Methods

### NewApiKeyConjurAuthenticationMethodResponse

`func NewApiKeyConjurAuthenticationMethodResponse(schemas []EnumapiKeyConjurAuthenticationMethodSchemaUrn, username string, id string, ) *ApiKeyConjurAuthenticationMethodResponse`

NewApiKeyConjurAuthenticationMethodResponse instantiates a new ApiKeyConjurAuthenticationMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyConjurAuthenticationMethodResponseWithDefaults

`func NewApiKeyConjurAuthenticationMethodResponseWithDefaults() *ApiKeyConjurAuthenticationMethodResponse`

NewApiKeyConjurAuthenticationMethodResponseWithDefaults instantiates a new ApiKeyConjurAuthenticationMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetSchemas() []EnumapiKeyConjurAuthenticationMethodSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetSchemasOk() (*[]EnumapiKeyConjurAuthenticationMethodSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetSchemas(v []EnumapiKeyConjurAuthenticationMethodSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetUsername

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiKeyConjurAuthenticationMethodResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiKey

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ApiKeyConjurAuthenticationMethodResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyConjurAuthenticationMethodResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApiKeyConjurAuthenticationMethodResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ApiKeyConjurAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyConjurAuthenticationMethodResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyConjurAuthenticationMethodResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


