# AddSyslogExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsyslogExternalServerSchemaUrn**](EnumsyslogExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The address of the syslog server. | 
**ServerPort** | Pointer to **int64** | The port on which the syslog server accepts connections. | [optional] 
**TransportMechanism** | [**EnumexternalServerTransportMechanismProp**](EnumexternalServerTransportMechanismProp.md) |  | 
**ConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. This will only be used when communicating with the syslog server over TCP (with or without TLS encryption). | [optional] 
**MaxConnectionAge** | Pointer to **string** | The maximum length of time that TCP connections should remain established. This will be ignored for UDP-based connections. A zero duration indicates that no maximum age will be imposed. | [optional] 
**TrustManagerProvider** | Pointer to **string** | A trust manager provider that will be used to determine whether to trust the certificate chain presented by the syslog server when communication is encrypted with TLS. This property will be ignored when not using TLS encryption. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**ServerName** | **string** | Name of the new External Server | 

## Methods

### NewAddSyslogExternalServerRequest

`func NewAddSyslogExternalServerRequest(schemas []EnumsyslogExternalServerSchemaUrn, serverHostName string, transportMechanism EnumexternalServerTransportMechanismProp, serverName string, ) *AddSyslogExternalServerRequest`

NewAddSyslogExternalServerRequest instantiates a new AddSyslogExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSyslogExternalServerRequestWithDefaults

`func NewAddSyslogExternalServerRequestWithDefaults() *AddSyslogExternalServerRequest`

NewAddSyslogExternalServerRequestWithDefaults instantiates a new AddSyslogExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSyslogExternalServerRequest) GetSchemas() []EnumsyslogExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSyslogExternalServerRequest) GetSchemasOk() (*[]EnumsyslogExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSyslogExternalServerRequest) SetSchemas(v []EnumsyslogExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *AddSyslogExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddSyslogExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddSyslogExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddSyslogExternalServerRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddSyslogExternalServerRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddSyslogExternalServerRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddSyslogExternalServerRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetTransportMechanism

`func (o *AddSyslogExternalServerRequest) GetTransportMechanism() EnumexternalServerTransportMechanismProp`

GetTransportMechanism returns the TransportMechanism field if non-nil, zero value otherwise.

### GetTransportMechanismOk

`func (o *AddSyslogExternalServerRequest) GetTransportMechanismOk() (*EnumexternalServerTransportMechanismProp, bool)`

GetTransportMechanismOk returns a tuple with the TransportMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMechanism

`func (o *AddSyslogExternalServerRequest) SetTransportMechanism(v EnumexternalServerTransportMechanismProp)`

SetTransportMechanism sets TransportMechanism field to given value.


### GetConnectTimeout

`func (o *AddSyslogExternalServerRequest) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *AddSyslogExternalServerRequest) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *AddSyslogExternalServerRequest) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *AddSyslogExternalServerRequest) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *AddSyslogExternalServerRequest) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *AddSyslogExternalServerRequest) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *AddSyslogExternalServerRequest) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.

### HasMaxConnectionAge

`func (o *AddSyslogExternalServerRequest) HasMaxConnectionAge() bool`

HasMaxConnectionAge returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddSyslogExternalServerRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddSyslogExternalServerRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddSyslogExternalServerRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddSyslogExternalServerRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddSyslogExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSyslogExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSyslogExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSyslogExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerName

`func (o *AddSyslogExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddSyslogExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddSyslogExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


