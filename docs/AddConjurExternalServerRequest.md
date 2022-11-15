# AddConjurExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumconjurExternalServerSchemaUrn**](EnumconjurExternalServerSchemaUrn.md) |  | 
**ConjurServerBaseURI** | **[]string** |  | 
**ConjurAuthenticationMethod** | **string** | The mechanism used to authenticate to the Conjur server. | 
**ConjurAccountName** | **string** | The name of the account with which the desired secrets are associated. | 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Conjur servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The PIN needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot;, \&quot;PKCS12\&quot;, or \&quot;BCFKS\&quot;. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewAddConjurExternalServerRequest

`func NewAddConjurExternalServerRequest(serverName string, schemas []EnumconjurExternalServerSchemaUrn, conjurServerBaseURI []string, conjurAuthenticationMethod string, conjurAccountName string, ) *AddConjurExternalServerRequest`

NewAddConjurExternalServerRequest instantiates a new AddConjurExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConjurExternalServerRequestWithDefaults

`func NewAddConjurExternalServerRequestWithDefaults() *AddConjurExternalServerRequest`

NewAddConjurExternalServerRequestWithDefaults instantiates a new AddConjurExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddConjurExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddConjurExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddConjurExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddConjurExternalServerRequest) GetSchemas() []EnumconjurExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConjurExternalServerRequest) GetSchemasOk() (*[]EnumconjurExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConjurExternalServerRequest) SetSchemas(v []EnumconjurExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConjurServerBaseURI

`func (o *AddConjurExternalServerRequest) GetConjurServerBaseURI() []string`

GetConjurServerBaseURI returns the ConjurServerBaseURI field if non-nil, zero value otherwise.

### GetConjurServerBaseURIOk

`func (o *AddConjurExternalServerRequest) GetConjurServerBaseURIOk() (*[]string, bool)`

GetConjurServerBaseURIOk returns a tuple with the ConjurServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurServerBaseURI

`func (o *AddConjurExternalServerRequest) SetConjurServerBaseURI(v []string)`

SetConjurServerBaseURI sets ConjurServerBaseURI field to given value.


### GetConjurAuthenticationMethod

`func (o *AddConjurExternalServerRequest) GetConjurAuthenticationMethod() string`

GetConjurAuthenticationMethod returns the ConjurAuthenticationMethod field if non-nil, zero value otherwise.

### GetConjurAuthenticationMethodOk

`func (o *AddConjurExternalServerRequest) GetConjurAuthenticationMethodOk() (*string, bool)`

GetConjurAuthenticationMethodOk returns a tuple with the ConjurAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurAuthenticationMethod

`func (o *AddConjurExternalServerRequest) SetConjurAuthenticationMethod(v string)`

SetConjurAuthenticationMethod sets ConjurAuthenticationMethod field to given value.


### GetConjurAccountName

`func (o *AddConjurExternalServerRequest) GetConjurAccountName() string`

GetConjurAccountName returns the ConjurAccountName field if non-nil, zero value otherwise.

### GetConjurAccountNameOk

`func (o *AddConjurExternalServerRequest) GetConjurAccountNameOk() (*string, bool)`

GetConjurAccountNameOk returns a tuple with the ConjurAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurAccountName

`func (o *AddConjurExternalServerRequest) SetConjurAccountName(v string)`

SetConjurAccountName sets ConjurAccountName field to given value.


### GetTrustStoreFile

`func (o *AddConjurExternalServerRequest) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *AddConjurExternalServerRequest) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *AddConjurExternalServerRequest) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *AddConjurExternalServerRequest) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *AddConjurExternalServerRequest) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *AddConjurExternalServerRequest) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *AddConjurExternalServerRequest) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *AddConjurExternalServerRequest) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *AddConjurExternalServerRequest) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *AddConjurExternalServerRequest) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *AddConjurExternalServerRequest) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *AddConjurExternalServerRequest) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetDescription

`func (o *AddConjurExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddConjurExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddConjurExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddConjurExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


