# AmazonSecretsManagerPassphraseProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumamazonSecretsManagerPassphraseProviderSchemaUrn**](EnumamazonSecretsManagerPassphraseProviderSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be retrieved. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Vault. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewAmazonSecretsManagerPassphraseProviderResponse

`func NewAmazonSecretsManagerPassphraseProviderResponse(id string, schemas []EnumamazonSecretsManagerPassphraseProviderSchemaUrn, awsExternalServer string, secretID string, secretFieldName string, enabled bool, ) *AmazonSecretsManagerPassphraseProviderResponse`

NewAmazonSecretsManagerPassphraseProviderResponse instantiates a new AmazonSecretsManagerPassphraseProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSecretsManagerPassphraseProviderResponseWithDefaults

`func NewAmazonSecretsManagerPassphraseProviderResponseWithDefaults() *AmazonSecretsManagerPassphraseProviderResponse`

NewAmazonSecretsManagerPassphraseProviderResponseWithDefaults instantiates a new AmazonSecretsManagerPassphraseProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSchemas() []EnumamazonSecretsManagerPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSchemasOk() (*[]EnumamazonSecretsManagerPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetSchemas(v []EnumamazonSecretsManagerPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetSecretID

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *AmazonSecretsManagerPassphraseProviderResponse) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *AmazonSecretsManagerPassphraseProviderResponse) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetMaxCacheDuration

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *AmazonSecretsManagerPassphraseProviderResponse) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonSecretsManagerPassphraseProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AmazonSecretsManagerPassphraseProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AmazonSecretsManagerPassphraseProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AmazonSecretsManagerPassphraseProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


