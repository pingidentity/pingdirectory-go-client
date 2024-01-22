# KafkaClusterExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumkafkaClusterExternalServerSchemaUrn**](EnumkafkaClusterExternalServerSchemaUrn.md) |  | 
**Id** | **string** | Name of the External Server | 
**BootstrapServer** | **[]string** | List of Kafka brokers to use for this Kafka Cluster External Server, following the host:port format. | 
**ProducerProperty** | Pointer to **[]string** | Specifies extra properties to use when constructing the KafkaProducer for sending messages. | [optional] 
**UseSSL** | Pointer to **bool** | If enabled, the Kafka Cluster External Server will use SSL to encrypt communication with the Kafka brokers. | [optional] 
**TrustManagerProvider** | Pointer to **string** | Specifies the file-based trust manager that should be used with the Kafka Cluster External Server for connecting to the Kafka cluster over SSL. | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies the file-based key manager that should be used with the Kafka Cluster External Server for connecting to the Kafka cluster over SSL. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewKafkaClusterExternalServerResponse

`func NewKafkaClusterExternalServerResponse(schemas []EnumkafkaClusterExternalServerSchemaUrn, id string, bootstrapServer []string, ) *KafkaClusterExternalServerResponse`

NewKafkaClusterExternalServerResponse instantiates a new KafkaClusterExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaClusterExternalServerResponseWithDefaults

`func NewKafkaClusterExternalServerResponseWithDefaults() *KafkaClusterExternalServerResponse`

NewKafkaClusterExternalServerResponseWithDefaults instantiates a new KafkaClusterExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *KafkaClusterExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *KafkaClusterExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *KafkaClusterExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *KafkaClusterExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *KafkaClusterExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *KafkaClusterExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *KafkaClusterExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *KafkaClusterExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *KafkaClusterExternalServerResponse) GetSchemas() []EnumkafkaClusterExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *KafkaClusterExternalServerResponse) GetSchemasOk() (*[]EnumkafkaClusterExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *KafkaClusterExternalServerResponse) SetSchemas(v []EnumkafkaClusterExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *KafkaClusterExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaClusterExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaClusterExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBootstrapServer

`func (o *KafkaClusterExternalServerResponse) GetBootstrapServer() []string`

GetBootstrapServer returns the BootstrapServer field if non-nil, zero value otherwise.

### GetBootstrapServerOk

`func (o *KafkaClusterExternalServerResponse) GetBootstrapServerOk() (*[]string, bool)`

GetBootstrapServerOk returns a tuple with the BootstrapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServer

`func (o *KafkaClusterExternalServerResponse) SetBootstrapServer(v []string)`

SetBootstrapServer sets BootstrapServer field to given value.


### GetProducerProperty

`func (o *KafkaClusterExternalServerResponse) GetProducerProperty() []string`

GetProducerProperty returns the ProducerProperty field if non-nil, zero value otherwise.

### GetProducerPropertyOk

`func (o *KafkaClusterExternalServerResponse) GetProducerPropertyOk() (*[]string, bool)`

GetProducerPropertyOk returns a tuple with the ProducerProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerProperty

`func (o *KafkaClusterExternalServerResponse) SetProducerProperty(v []string)`

SetProducerProperty sets ProducerProperty field to given value.

### HasProducerProperty

`func (o *KafkaClusterExternalServerResponse) HasProducerProperty() bool`

HasProducerProperty returns a boolean if a field has been set.

### GetUseSSL

`func (o *KafkaClusterExternalServerResponse) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *KafkaClusterExternalServerResponse) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *KafkaClusterExternalServerResponse) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *KafkaClusterExternalServerResponse) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *KafkaClusterExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *KafkaClusterExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *KafkaClusterExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *KafkaClusterExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *KafkaClusterExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *KafkaClusterExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *KafkaClusterExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *KafkaClusterExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetDescription

`func (o *KafkaClusterExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KafkaClusterExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KafkaClusterExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KafkaClusterExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


