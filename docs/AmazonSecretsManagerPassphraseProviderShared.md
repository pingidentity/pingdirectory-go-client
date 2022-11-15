# AmazonSecretsManagerPassphraseProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonSecretsManagerPassphraseProviderSchemaUrn**](EnumamazonSecretsManagerPassphraseProviderSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be retrieved. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Vault. | [optional] 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewAmazonSecretsManagerPassphraseProviderShared

`func NewAmazonSecretsManagerPassphraseProviderShared(schemas []EnumamazonSecretsManagerPassphraseProviderSchemaUrn, awsExternalServer string, secretID string, secretFieldName string, enabled bool, ) *AmazonSecretsManagerPassphraseProviderShared`

NewAmazonSecretsManagerPassphraseProviderShared instantiates a new AmazonSecretsManagerPassphraseProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSecretsManagerPassphraseProviderSharedWithDefaults

`func NewAmazonSecretsManagerPassphraseProviderSharedWithDefaults() *AmazonSecretsManagerPassphraseProviderShared`

NewAmazonSecretsManagerPassphraseProviderSharedWithDefaults instantiates a new AmazonSecretsManagerPassphraseProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSchemas() []EnumamazonSecretsManagerPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSchemasOk() (*[]EnumamazonSecretsManagerPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetSchemas(v []EnumamazonSecretsManagerPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetSecretID

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *AmazonSecretsManagerPassphraseProviderShared) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *AmazonSecretsManagerPassphraseProviderShared) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetMaxCacheDuration

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *AmazonSecretsManagerPassphraseProviderShared) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonSecretsManagerPassphraseProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmazonSecretsManagerPassphraseProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmazonSecretsManagerPassphraseProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


