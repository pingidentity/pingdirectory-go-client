# SidebandApiHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumsidebandApiHttpServletExtensionSchemaUrn**](EnumsidebandApiHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**RequestLimit** | Pointer to **string** | The maximum number of bytes allowed per request body. | [optional] 
**RequestContextMethod** | Pointer to [**EnumhttpServletExtensionRequestContextMethodProp**](EnumhttpServletExtensionRequestContextMethodProp.md) |  | [optional] 
**SharedSecretHeaderName** | **string** | The request header used to find the shared secret header for incoming Sideband API HTTP Servlet Extension requests. | 
**SharedSecrets** | Pointer to **[]string** | Shared secrets between the third-party API Gateway and Sideband API HTTP Servlet Extension. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewSidebandApiHttpServletExtensionResponse

`func NewSidebandApiHttpServletExtensionResponse(schemas []EnumsidebandApiHttpServletExtensionSchemaUrn, id string, sharedSecretHeaderName string, ) *SidebandApiHttpServletExtensionResponse`

NewSidebandApiHttpServletExtensionResponse instantiates a new SidebandApiHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidebandApiHttpServletExtensionResponseWithDefaults

`func NewSidebandApiHttpServletExtensionResponseWithDefaults() *SidebandApiHttpServletExtensionResponse`

NewSidebandApiHttpServletExtensionResponseWithDefaults instantiates a new SidebandApiHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SidebandApiHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SidebandApiHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SidebandApiHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SidebandApiHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SidebandApiHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SidebandApiHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SidebandApiHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SidebandApiHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *SidebandApiHttpServletExtensionResponse) GetSchemas() []EnumsidebandApiHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SidebandApiHttpServletExtensionResponse) GetSchemasOk() (*[]EnumsidebandApiHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SidebandApiHttpServletExtensionResponse) SetSchemas(v []EnumsidebandApiHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SidebandApiHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SidebandApiHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SidebandApiHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRequestLimit

`func (o *SidebandApiHttpServletExtensionResponse) GetRequestLimit() string`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *SidebandApiHttpServletExtensionResponse) GetRequestLimitOk() (*string, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *SidebandApiHttpServletExtensionResponse) SetRequestLimit(v string)`

SetRequestLimit sets RequestLimit field to given value.

### HasRequestLimit

`func (o *SidebandApiHttpServletExtensionResponse) HasRequestLimit() bool`

HasRequestLimit returns a boolean if a field has been set.

### GetRequestContextMethod

`func (o *SidebandApiHttpServletExtensionResponse) GetRequestContextMethod() EnumhttpServletExtensionRequestContextMethodProp`

GetRequestContextMethod returns the RequestContextMethod field if non-nil, zero value otherwise.

### GetRequestContextMethodOk

`func (o *SidebandApiHttpServletExtensionResponse) GetRequestContextMethodOk() (*EnumhttpServletExtensionRequestContextMethodProp, bool)`

GetRequestContextMethodOk returns a tuple with the RequestContextMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContextMethod

`func (o *SidebandApiHttpServletExtensionResponse) SetRequestContextMethod(v EnumhttpServletExtensionRequestContextMethodProp)`

SetRequestContextMethod sets RequestContextMethod field to given value.

### HasRequestContextMethod

`func (o *SidebandApiHttpServletExtensionResponse) HasRequestContextMethod() bool`

HasRequestContextMethod returns a boolean if a field has been set.

### GetSharedSecretHeaderName

`func (o *SidebandApiHttpServletExtensionResponse) GetSharedSecretHeaderName() string`

GetSharedSecretHeaderName returns the SharedSecretHeaderName field if non-nil, zero value otherwise.

### GetSharedSecretHeaderNameOk

`func (o *SidebandApiHttpServletExtensionResponse) GetSharedSecretHeaderNameOk() (*string, bool)`

GetSharedSecretHeaderNameOk returns a tuple with the SharedSecretHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretHeaderName

`func (o *SidebandApiHttpServletExtensionResponse) SetSharedSecretHeaderName(v string)`

SetSharedSecretHeaderName sets SharedSecretHeaderName field to given value.


### GetSharedSecrets

`func (o *SidebandApiHttpServletExtensionResponse) GetSharedSecrets() []string`

GetSharedSecrets returns the SharedSecrets field if non-nil, zero value otherwise.

### GetSharedSecretsOk

`func (o *SidebandApiHttpServletExtensionResponse) GetSharedSecretsOk() (*[]string, bool)`

GetSharedSecretsOk returns a tuple with the SharedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecrets

`func (o *SidebandApiHttpServletExtensionResponse) SetSharedSecrets(v []string)`

SetSharedSecrets sets SharedSecrets field to given value.

### HasSharedSecrets

`func (o *SidebandApiHttpServletExtensionResponse) HasSharedSecrets() bool`

HasSharedSecrets returns a boolean if a field has been set.

### GetDescription

`func (o *SidebandApiHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SidebandApiHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SidebandApiHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SidebandApiHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *SidebandApiHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *SidebandApiHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *SidebandApiHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *SidebandApiHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *SidebandApiHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *SidebandApiHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *SidebandApiHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *SidebandApiHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *SidebandApiHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *SidebandApiHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *SidebandApiHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *SidebandApiHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


