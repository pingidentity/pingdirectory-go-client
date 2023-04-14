# AddJmxConnectionHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Connection Handler | 
**Schemas** | [**[]EnumjmxConnectionHandlerSchemaUrn**](EnumjmxConnectionHandlerSchemaUrn.md) |  | 
**ListenPort** | **int64** | Specifies the port number on which the JMX Connection Handler will listen for connections from clients. | 
**UseSSL** | Pointer to **bool** | Indicates whether the JMX Connection Handler should use SSL. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the JMX Connection Handler should use when performing SSL communication. | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies the name of the key manager that should be used with this JMX Connection Handler . | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are allowed to establish connections to this connection handler. | [optional] 
**DeniedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are not allowed to establish connections to this connection handler. | [optional] 

## Methods

### NewAddJmxConnectionHandlerRequest

`func NewAddJmxConnectionHandlerRequest(handlerName string, schemas []EnumjmxConnectionHandlerSchemaUrn, listenPort int64, enabled bool, ) *AddJmxConnectionHandlerRequest`

NewAddJmxConnectionHandlerRequest instantiates a new AddJmxConnectionHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJmxConnectionHandlerRequestWithDefaults

`func NewAddJmxConnectionHandlerRequestWithDefaults() *AddJmxConnectionHandlerRequest`

NewAddJmxConnectionHandlerRequestWithDefaults instantiates a new AddJmxConnectionHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddJmxConnectionHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddJmxConnectionHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddJmxConnectionHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddJmxConnectionHandlerRequest) GetSchemas() []EnumjmxConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJmxConnectionHandlerRequest) GetSchemasOk() (*[]EnumjmxConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJmxConnectionHandlerRequest) SetSchemas(v []EnumjmxConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenPort

`func (o *AddJmxConnectionHandlerRequest) GetListenPort() int64`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *AddJmxConnectionHandlerRequest) GetListenPortOk() (*int64, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *AddJmxConnectionHandlerRequest) SetListenPort(v int64)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *AddJmxConnectionHandlerRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *AddJmxConnectionHandlerRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *AddJmxConnectionHandlerRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *AddJmxConnectionHandlerRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *AddJmxConnectionHandlerRequest) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddJmxConnectionHandlerRequest) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddJmxConnectionHandlerRequest) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *AddJmxConnectionHandlerRequest) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *AddJmxConnectionHandlerRequest) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddJmxConnectionHandlerRequest) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddJmxConnectionHandlerRequest) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddJmxConnectionHandlerRequest) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddJmxConnectionHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJmxConnectionHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJmxConnectionHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJmxConnectionHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddJmxConnectionHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddJmxConnectionHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddJmxConnectionHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *AddJmxConnectionHandlerRequest) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *AddJmxConnectionHandlerRequest) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *AddJmxConnectionHandlerRequest) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *AddJmxConnectionHandlerRequest) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *AddJmxConnectionHandlerRequest) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *AddJmxConnectionHandlerRequest) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *AddJmxConnectionHandlerRequest) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *AddJmxConnectionHandlerRequest) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


