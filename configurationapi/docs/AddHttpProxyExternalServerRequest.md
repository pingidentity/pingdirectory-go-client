# AddHttpProxyExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumhttpProxyExternalServerSchemaUrn**](EnumhttpProxyExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The host name or IP address of the HTTP Proxy External Server. | 
**ServerPort** | **int64** | The port on which the HTTP Proxy External Server is listening for connections. | 
**BasicAuthenticationUsername** | Pointer to **string** | The username to use to authenticate to the HTTP Proxy External Server. | [optional] 
**BasicAuthenticationPassphraseProvider** | Pointer to **string** | A passphrase provider that provides access to the password to use to authenticate to the HTTP Proxy External Server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewAddHttpProxyExternalServerRequest

`func NewAddHttpProxyExternalServerRequest(serverName string, schemas []EnumhttpProxyExternalServerSchemaUrn, serverHostName string, serverPort int64, ) *AddHttpProxyExternalServerRequest`

NewAddHttpProxyExternalServerRequest instantiates a new AddHttpProxyExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddHttpProxyExternalServerRequestWithDefaults

`func NewAddHttpProxyExternalServerRequestWithDefaults() *AddHttpProxyExternalServerRequest`

NewAddHttpProxyExternalServerRequestWithDefaults instantiates a new AddHttpProxyExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddHttpProxyExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddHttpProxyExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddHttpProxyExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddHttpProxyExternalServerRequest) GetSchemas() []EnumhttpProxyExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddHttpProxyExternalServerRequest) GetSchemasOk() (*[]EnumhttpProxyExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddHttpProxyExternalServerRequest) SetSchemas(v []EnumhttpProxyExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *AddHttpProxyExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddHttpProxyExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddHttpProxyExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddHttpProxyExternalServerRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddHttpProxyExternalServerRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddHttpProxyExternalServerRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.


### GetBasicAuthenticationUsername

`func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationUsername() string`

GetBasicAuthenticationUsername returns the BasicAuthenticationUsername field if non-nil, zero value otherwise.

### GetBasicAuthenticationUsernameOk

`func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationUsernameOk() (*string, bool)`

GetBasicAuthenticationUsernameOk returns a tuple with the BasicAuthenticationUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationUsername

`func (o *AddHttpProxyExternalServerRequest) SetBasicAuthenticationUsername(v string)`

SetBasicAuthenticationUsername sets BasicAuthenticationUsername field to given value.

### HasBasicAuthenticationUsername

`func (o *AddHttpProxyExternalServerRequest) HasBasicAuthenticationUsername() bool`

HasBasicAuthenticationUsername returns a boolean if a field has been set.

### GetBasicAuthenticationPassphraseProvider

`func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationPassphraseProvider() string`

GetBasicAuthenticationPassphraseProvider returns the BasicAuthenticationPassphraseProvider field if non-nil, zero value otherwise.

### GetBasicAuthenticationPassphraseProviderOk

`func (o *AddHttpProxyExternalServerRequest) GetBasicAuthenticationPassphraseProviderOk() (*string, bool)`

GetBasicAuthenticationPassphraseProviderOk returns a tuple with the BasicAuthenticationPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationPassphraseProvider

`func (o *AddHttpProxyExternalServerRequest) SetBasicAuthenticationPassphraseProvider(v string)`

SetBasicAuthenticationPassphraseProvider sets BasicAuthenticationPassphraseProvider field to given value.

### HasBasicAuthenticationPassphraseProvider

`func (o *AddHttpProxyExternalServerRequest) HasBasicAuthenticationPassphraseProvider() bool`

HasBasicAuthenticationPassphraseProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddHttpProxyExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddHttpProxyExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddHttpProxyExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddHttpProxyExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


