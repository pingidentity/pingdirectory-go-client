# GatewayHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumgatewayHttpServletExtensionSchemaUrn**](EnumgatewayHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**ExcludedAPIEndpoint** | Pointer to **[]string** | Specifies any Gateway API Endpoints that will not be handled by the Gateway HTTP Servlet Extension. | [optional] 
**RequestLimit** | Pointer to **string** | The maximum number of bytes allowed per request body. | [optional] 
**ResponseLimit** | Pointer to **string** | The maximum number of bytes allowed per response body. | [optional] 
**NumForwardThreads** | Pointer to **int64** | The number of threads used to forward responses to the API backend. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 

## Methods

### NewGatewayHttpServletExtensionResponse

`func NewGatewayHttpServletExtensionResponse(schemas []EnumgatewayHttpServletExtensionSchemaUrn, id string, ) *GatewayHttpServletExtensionResponse`

NewGatewayHttpServletExtensionResponse instantiates a new GatewayHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayHttpServletExtensionResponseWithDefaults

`func NewGatewayHttpServletExtensionResponseWithDefaults() *GatewayHttpServletExtensionResponse`

NewGatewayHttpServletExtensionResponseWithDefaults instantiates a new GatewayHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GatewayHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GatewayHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GatewayHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GatewayHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GatewayHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GatewayHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GatewayHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GatewayHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GatewayHttpServletExtensionResponse) GetSchemas() []EnumgatewayHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GatewayHttpServletExtensionResponse) GetSchemasOk() (*[]EnumgatewayHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GatewayHttpServletExtensionResponse) SetSchemas(v []EnumgatewayHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GatewayHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetExcludedAPIEndpoint

`func (o *GatewayHttpServletExtensionResponse) GetExcludedAPIEndpoint() []string`

GetExcludedAPIEndpoint returns the ExcludedAPIEndpoint field if non-nil, zero value otherwise.

### GetExcludedAPIEndpointOk

`func (o *GatewayHttpServletExtensionResponse) GetExcludedAPIEndpointOk() (*[]string, bool)`

GetExcludedAPIEndpointOk returns a tuple with the ExcludedAPIEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAPIEndpoint

`func (o *GatewayHttpServletExtensionResponse) SetExcludedAPIEndpoint(v []string)`

SetExcludedAPIEndpoint sets ExcludedAPIEndpoint field to given value.

### HasExcludedAPIEndpoint

`func (o *GatewayHttpServletExtensionResponse) HasExcludedAPIEndpoint() bool`

HasExcludedAPIEndpoint returns a boolean if a field has been set.

### GetRequestLimit

`func (o *GatewayHttpServletExtensionResponse) GetRequestLimit() string`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *GatewayHttpServletExtensionResponse) GetRequestLimitOk() (*string, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *GatewayHttpServletExtensionResponse) SetRequestLimit(v string)`

SetRequestLimit sets RequestLimit field to given value.

### HasRequestLimit

`func (o *GatewayHttpServletExtensionResponse) HasRequestLimit() bool`

HasRequestLimit returns a boolean if a field has been set.

### GetResponseLimit

`func (o *GatewayHttpServletExtensionResponse) GetResponseLimit() string`

GetResponseLimit returns the ResponseLimit field if non-nil, zero value otherwise.

### GetResponseLimitOk

`func (o *GatewayHttpServletExtensionResponse) GetResponseLimitOk() (*string, bool)`

GetResponseLimitOk returns a tuple with the ResponseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLimit

`func (o *GatewayHttpServletExtensionResponse) SetResponseLimit(v string)`

SetResponseLimit sets ResponseLimit field to given value.

### HasResponseLimit

`func (o *GatewayHttpServletExtensionResponse) HasResponseLimit() bool`

HasResponseLimit returns a boolean if a field has been set.

### GetNumForwardThreads

`func (o *GatewayHttpServletExtensionResponse) GetNumForwardThreads() int64`

GetNumForwardThreads returns the NumForwardThreads field if non-nil, zero value otherwise.

### GetNumForwardThreadsOk

`func (o *GatewayHttpServletExtensionResponse) GetNumForwardThreadsOk() (*int64, bool)`

GetNumForwardThreadsOk returns a tuple with the NumForwardThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumForwardThreads

`func (o *GatewayHttpServletExtensionResponse) SetNumForwardThreads(v int64)`

SetNumForwardThreads sets NumForwardThreads field to given value.

### HasNumForwardThreads

`func (o *GatewayHttpServletExtensionResponse) HasNumForwardThreads() bool`

HasNumForwardThreads returns a boolean if a field has been set.

### GetDescription

`func (o *GatewayHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *GatewayHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *GatewayHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *GatewayHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *GatewayHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *GatewayHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *GatewayHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *GatewayHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *GatewayHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


