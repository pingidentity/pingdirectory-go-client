# ConjurCipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Cipher Stream Provider | 
**Schemas** | [**[]EnumconjurCipherStreamProviderSchemaUrn**](EnumconjurCipherStreamProviderSchemaUrn.md) |  | 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur server. | 
**ConjurSecretRelativePath** | **string** | The portion of the path that follows the account name in the URI needed to obtain the secret passphrase to use to generate the encryption key. Any special characters in the path must be URL-encoded. | 
**EncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this Conjur Cipher Stream Provider. | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewConjurCipherStreamProviderResponse

`func NewConjurCipherStreamProviderResponse(id string, schemas []EnumconjurCipherStreamProviderSchemaUrn, conjurExternalServer string, conjurSecretRelativePath string, encryptionMetadataFile string, enabled bool, ) *ConjurCipherStreamProviderResponse`

NewConjurCipherStreamProviderResponse instantiates a new ConjurCipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConjurCipherStreamProviderResponseWithDefaults

`func NewConjurCipherStreamProviderResponseWithDefaults() *ConjurCipherStreamProviderResponse`

NewConjurCipherStreamProviderResponseWithDefaults instantiates a new ConjurCipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ConjurCipherStreamProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConjurCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConjurCipherStreamProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConjurCipherStreamProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConjurCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConjurCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConjurCipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConjurCipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ConjurCipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConjurCipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConjurCipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ConjurCipherStreamProviderResponse) GetSchemas() []EnumconjurCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConjurCipherStreamProviderResponse) GetSchemasOk() (*[]EnumconjurCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConjurCipherStreamProviderResponse) SetSchemas(v []EnumconjurCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConjurExternalServer

`func (o *ConjurCipherStreamProviderResponse) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *ConjurCipherStreamProviderResponse) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *ConjurCipherStreamProviderResponse) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetConjurSecretRelativePath

`func (o *ConjurCipherStreamProviderResponse) GetConjurSecretRelativePath() string`

GetConjurSecretRelativePath returns the ConjurSecretRelativePath field if non-nil, zero value otherwise.

### GetConjurSecretRelativePathOk

`func (o *ConjurCipherStreamProviderResponse) GetConjurSecretRelativePathOk() (*string, bool)`

GetConjurSecretRelativePathOk returns a tuple with the ConjurSecretRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurSecretRelativePath

`func (o *ConjurCipherStreamProviderResponse) SetConjurSecretRelativePath(v string)`

SetConjurSecretRelativePath sets ConjurSecretRelativePath field to given value.


### GetEncryptionMetadataFile

`func (o *ConjurCipherStreamProviderResponse) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *ConjurCipherStreamProviderResponse) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *ConjurCipherStreamProviderResponse) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.


### GetDescription

`func (o *ConjurCipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConjurCipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConjurCipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConjurCipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ConjurCipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConjurCipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConjurCipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


