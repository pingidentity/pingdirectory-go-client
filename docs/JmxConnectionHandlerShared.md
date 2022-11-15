# JmxConnectionHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjmxConnectionHandlerSchemaUrn**](EnumjmxConnectionHandlerSchemaUrn.md) |  | 
**ListenPort** | **int32** | Specifies the port number on which the JMX Connection Handler will listen for connections from clients. | 
**UseSSL** | Pointer to **bool** | Indicates whether the JMX Connection Handler should use SSL. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the JMX Connection Handler should use when performing SSL communication. | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies the name of the key manager that should be used with this JMX Connection Handler . | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** |  | [optional] 
**DeniedClient** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJmxConnectionHandlerShared

`func NewJmxConnectionHandlerShared(schemas []EnumjmxConnectionHandlerSchemaUrn, listenPort int32, enabled bool, ) *JmxConnectionHandlerShared`

NewJmxConnectionHandlerShared instantiates a new JmxConnectionHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJmxConnectionHandlerSharedWithDefaults

`func NewJmxConnectionHandlerSharedWithDefaults() *JmxConnectionHandlerShared`

NewJmxConnectionHandlerSharedWithDefaults instantiates a new JmxConnectionHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JmxConnectionHandlerShared) GetSchemas() []EnumjmxConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JmxConnectionHandlerShared) GetSchemasOk() (*[]EnumjmxConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JmxConnectionHandlerShared) SetSchemas(v []EnumjmxConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenPort

`func (o *JmxConnectionHandlerShared) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *JmxConnectionHandlerShared) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *JmxConnectionHandlerShared) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *JmxConnectionHandlerShared) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *JmxConnectionHandlerShared) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *JmxConnectionHandlerShared) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *JmxConnectionHandlerShared) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *JmxConnectionHandlerShared) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *JmxConnectionHandlerShared) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *JmxConnectionHandlerShared) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *JmxConnectionHandlerShared) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *JmxConnectionHandlerShared) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *JmxConnectionHandlerShared) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *JmxConnectionHandlerShared) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *JmxConnectionHandlerShared) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetDescription

`func (o *JmxConnectionHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JmxConnectionHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JmxConnectionHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JmxConnectionHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JmxConnectionHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JmxConnectionHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JmxConnectionHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *JmxConnectionHandlerShared) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *JmxConnectionHandlerShared) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *JmxConnectionHandlerShared) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *JmxConnectionHandlerShared) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *JmxConnectionHandlerShared) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *JmxConnectionHandlerShared) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *JmxConnectionHandlerShared) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *JmxConnectionHandlerShared) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


