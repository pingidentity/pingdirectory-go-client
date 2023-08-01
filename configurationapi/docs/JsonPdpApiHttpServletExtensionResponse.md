# JsonPdpApiHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjsonPdpApiHttpServletExtensionSchemaUrn**](EnumjsonPdpApiHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**SharedSecretHeaderName** | **string** | The request header used to find the shared secret header for incoming JSON PDP API HTTP Servlet Extension requests. | 
**SharedSecrets** | Pointer to **[]string** | Shared secrets between the third-party Policy Enforcement Point and the JSON PDP API HTTP Servlet Extension. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewJsonPdpApiHttpServletExtensionResponse

`func NewJsonPdpApiHttpServletExtensionResponse(schemas []EnumjsonPdpApiHttpServletExtensionSchemaUrn, id string, sharedSecretHeaderName string, ) *JsonPdpApiHttpServletExtensionResponse`

NewJsonPdpApiHttpServletExtensionResponse instantiates a new JsonPdpApiHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPdpApiHttpServletExtensionResponseWithDefaults

`func NewJsonPdpApiHttpServletExtensionResponseWithDefaults() *JsonPdpApiHttpServletExtensionResponse`

NewJsonPdpApiHttpServletExtensionResponseWithDefaults instantiates a new JsonPdpApiHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JsonPdpApiHttpServletExtensionResponse) GetSchemas() []EnumjsonPdpApiHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetSchemasOk() (*[]EnumjsonPdpApiHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonPdpApiHttpServletExtensionResponse) SetSchemas(v []EnumjsonPdpApiHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *JsonPdpApiHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonPdpApiHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSharedSecretHeaderName

`func (o *JsonPdpApiHttpServletExtensionResponse) GetSharedSecretHeaderName() string`

GetSharedSecretHeaderName returns the SharedSecretHeaderName field if non-nil, zero value otherwise.

### GetSharedSecretHeaderNameOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetSharedSecretHeaderNameOk() (*string, bool)`

GetSharedSecretHeaderNameOk returns a tuple with the SharedSecretHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretHeaderName

`func (o *JsonPdpApiHttpServletExtensionResponse) SetSharedSecretHeaderName(v string)`

SetSharedSecretHeaderName sets SharedSecretHeaderName field to given value.


### GetSharedSecrets

`func (o *JsonPdpApiHttpServletExtensionResponse) GetSharedSecrets() []string`

GetSharedSecrets returns the SharedSecrets field if non-nil, zero value otherwise.

### GetSharedSecretsOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetSharedSecretsOk() (*[]string, bool)`

GetSharedSecretsOk returns a tuple with the SharedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecrets

`func (o *JsonPdpApiHttpServletExtensionResponse) SetSharedSecrets(v []string)`

SetSharedSecrets sets SharedSecrets field to given value.

### HasSharedSecrets

`func (o *JsonPdpApiHttpServletExtensionResponse) HasSharedSecrets() bool`

HasSharedSecrets returns a boolean if a field has been set.

### GetDescription

`func (o *JsonPdpApiHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonPdpApiHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonPdpApiHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *JsonPdpApiHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *JsonPdpApiHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *JsonPdpApiHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *JsonPdpApiHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *JsonPdpApiHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *JsonPdpApiHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *JsonPdpApiHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *JsonPdpApiHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *JsonPdpApiHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *JsonPdpApiHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JsonPdpApiHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JsonPdpApiHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JsonPdpApiHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonPdpApiHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JsonPdpApiHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonPdpApiHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JsonPdpApiHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


