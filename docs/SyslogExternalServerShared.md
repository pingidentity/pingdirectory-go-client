# SyslogExternalServerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsyslogExternalServerSchemaUrn**](EnumsyslogExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The address of the syslog server. | 
**ServerPort** | Pointer to **int32** | The port on which the syslog server accepts connections. | [optional] 
**TransportMechanism** | [**EnumexternalServerTransportMechanismProp**](EnumexternalServerTransportMechanismProp.md) |  | 
**ConnectTimeout** | **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. This will only be used when communicating with the syslog server over TCP (with or without TLS encryption). | 
**MaxConnectionAge** | **string** | The maximum length of time that TCP connections should remain established. This will be ignored for UDP-based connections. A zero duration indicates that no maximum age will be imposed. | 
**TrustManagerProvider** | **string** | A trust manager provider that will be used to determine whether to trust the certificate chain presented by the syslog server when communication is encrypted with TLS. This property will be ignored when not using TLS encryption. | 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewSyslogExternalServerShared

`func NewSyslogExternalServerShared(schemas []EnumsyslogExternalServerSchemaUrn, serverHostName string, transportMechanism EnumexternalServerTransportMechanismProp, connectTimeout string, maxConnectionAge string, trustManagerProvider string, ) *SyslogExternalServerShared`

NewSyslogExternalServerShared instantiates a new SyslogExternalServerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogExternalServerSharedWithDefaults

`func NewSyslogExternalServerSharedWithDefaults() *SyslogExternalServerShared`

NewSyslogExternalServerSharedWithDefaults instantiates a new SyslogExternalServerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SyslogExternalServerShared) GetSchemas() []EnumsyslogExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SyslogExternalServerShared) GetSchemasOk() (*[]EnumsyslogExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SyslogExternalServerShared) SetSchemas(v []EnumsyslogExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *SyslogExternalServerShared) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *SyslogExternalServerShared) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *SyslogExternalServerShared) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *SyslogExternalServerShared) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *SyslogExternalServerShared) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *SyslogExternalServerShared) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *SyslogExternalServerShared) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetTransportMechanism

`func (o *SyslogExternalServerShared) GetTransportMechanism() EnumexternalServerTransportMechanismProp`

GetTransportMechanism returns the TransportMechanism field if non-nil, zero value otherwise.

### GetTransportMechanismOk

`func (o *SyslogExternalServerShared) GetTransportMechanismOk() (*EnumexternalServerTransportMechanismProp, bool)`

GetTransportMechanismOk returns a tuple with the TransportMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMechanism

`func (o *SyslogExternalServerShared) SetTransportMechanism(v EnumexternalServerTransportMechanismProp)`

SetTransportMechanism sets TransportMechanism field to given value.


### GetConnectTimeout

`func (o *SyslogExternalServerShared) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *SyslogExternalServerShared) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *SyslogExternalServerShared) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetMaxConnectionAge

`func (o *SyslogExternalServerShared) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *SyslogExternalServerShared) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *SyslogExternalServerShared) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.


### GetTrustManagerProvider

`func (o *SyslogExternalServerShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *SyslogExternalServerShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *SyslogExternalServerShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.


### GetDescription

`func (o *SyslogExternalServerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyslogExternalServerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyslogExternalServerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyslogExternalServerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


