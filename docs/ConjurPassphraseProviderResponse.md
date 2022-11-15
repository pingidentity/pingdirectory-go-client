# ConjurPassphraseProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumconjurPassphraseProviderSchemaUrn**](EnumconjurPassphraseProviderSchemaUrn.md) |  | 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur instance containing the passphrase. | 
**ConjurSecretRelativePath** | **string** | The portion of the path that follows the account name in the URI needed to obtain the desired secret. Any special characters in the path must be URL-encoded. | 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Conjur. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Conjur. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewConjurPassphraseProviderResponse

`func NewConjurPassphraseProviderResponse(id string, schemas []EnumconjurPassphraseProviderSchemaUrn, conjurExternalServer string, conjurSecretRelativePath string, enabled bool, ) *ConjurPassphraseProviderResponse`

NewConjurPassphraseProviderResponse instantiates a new ConjurPassphraseProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConjurPassphraseProviderResponseWithDefaults

`func NewConjurPassphraseProviderResponseWithDefaults() *ConjurPassphraseProviderResponse`

NewConjurPassphraseProviderResponseWithDefaults instantiates a new ConjurPassphraseProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConjurPassphraseProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConjurPassphraseProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConjurPassphraseProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ConjurPassphraseProviderResponse) GetSchemas() []EnumconjurPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConjurPassphraseProviderResponse) GetSchemasOk() (*[]EnumconjurPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConjurPassphraseProviderResponse) SetSchemas(v []EnumconjurPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConjurExternalServer

`func (o *ConjurPassphraseProviderResponse) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *ConjurPassphraseProviderResponse) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *ConjurPassphraseProviderResponse) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetConjurSecretRelativePath

`func (o *ConjurPassphraseProviderResponse) GetConjurSecretRelativePath() string`

GetConjurSecretRelativePath returns the ConjurSecretRelativePath field if non-nil, zero value otherwise.

### GetConjurSecretRelativePathOk

`func (o *ConjurPassphraseProviderResponse) GetConjurSecretRelativePathOk() (*string, bool)`

GetConjurSecretRelativePathOk returns a tuple with the ConjurSecretRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurSecretRelativePath

`func (o *ConjurPassphraseProviderResponse) SetConjurSecretRelativePath(v string)`

SetConjurSecretRelativePath sets ConjurSecretRelativePath field to given value.


### GetMaxCacheDuration

`func (o *ConjurPassphraseProviderResponse) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *ConjurPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *ConjurPassphraseProviderResponse) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *ConjurPassphraseProviderResponse) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *ConjurPassphraseProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConjurPassphraseProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConjurPassphraseProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConjurPassphraseProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ConjurPassphraseProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConjurPassphraseProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConjurPassphraseProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


