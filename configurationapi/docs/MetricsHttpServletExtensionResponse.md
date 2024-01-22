# MetricsHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnummetricsHttpServletExtensionSchemaUrn**](EnummetricsHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**RequireAPIAuthentication** | Pointer to **bool** | Require authentication when accessing the REST API. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the identity mapper that is to be used for associating user entries with basic authentication user names. | [optional] 
**OmitErrorMessageDetails** | Pointer to **bool** | Specifies that API error messages for invalid queries, unknown resources, service unavailable, and internal server errors are generic in nature. | [optional] 
**ApiAuthenticationTimeout** | Pointer to **string** | Length of time before a REST API authentication session expires. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewMetricsHttpServletExtensionResponse

`func NewMetricsHttpServletExtensionResponse(schemas []EnummetricsHttpServletExtensionSchemaUrn, id string, ) *MetricsHttpServletExtensionResponse`

NewMetricsHttpServletExtensionResponse instantiates a new MetricsHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsHttpServletExtensionResponseWithDefaults

`func NewMetricsHttpServletExtensionResponseWithDefaults() *MetricsHttpServletExtensionResponse`

NewMetricsHttpServletExtensionResponseWithDefaults instantiates a new MetricsHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *MetricsHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MetricsHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MetricsHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MetricsHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MetricsHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MetricsHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MetricsHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MetricsHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *MetricsHttpServletExtensionResponse) GetSchemas() []EnummetricsHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MetricsHttpServletExtensionResponse) GetSchemasOk() (*[]EnummetricsHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MetricsHttpServletExtensionResponse) SetSchemas(v []EnummetricsHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *MetricsHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRequireAPIAuthentication

`func (o *MetricsHttpServletExtensionResponse) GetRequireAPIAuthentication() bool`

GetRequireAPIAuthentication returns the RequireAPIAuthentication field if non-nil, zero value otherwise.

### GetRequireAPIAuthenticationOk

`func (o *MetricsHttpServletExtensionResponse) GetRequireAPIAuthenticationOk() (*bool, bool)`

GetRequireAPIAuthenticationOk returns a tuple with the RequireAPIAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAPIAuthentication

`func (o *MetricsHttpServletExtensionResponse) SetRequireAPIAuthentication(v bool)`

SetRequireAPIAuthentication sets RequireAPIAuthentication field to given value.

### HasRequireAPIAuthentication

`func (o *MetricsHttpServletExtensionResponse) HasRequireAPIAuthentication() bool`

HasRequireAPIAuthentication returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *MetricsHttpServletExtensionResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *MetricsHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *MetricsHttpServletExtensionResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *MetricsHttpServletExtensionResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetOmitErrorMessageDetails

`func (o *MetricsHttpServletExtensionResponse) GetOmitErrorMessageDetails() bool`

GetOmitErrorMessageDetails returns the OmitErrorMessageDetails field if non-nil, zero value otherwise.

### GetOmitErrorMessageDetailsOk

`func (o *MetricsHttpServletExtensionResponse) GetOmitErrorMessageDetailsOk() (*bool, bool)`

GetOmitErrorMessageDetailsOk returns a tuple with the OmitErrorMessageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitErrorMessageDetails

`func (o *MetricsHttpServletExtensionResponse) SetOmitErrorMessageDetails(v bool)`

SetOmitErrorMessageDetails sets OmitErrorMessageDetails field to given value.

### HasOmitErrorMessageDetails

`func (o *MetricsHttpServletExtensionResponse) HasOmitErrorMessageDetails() bool`

HasOmitErrorMessageDetails returns a boolean if a field has been set.

### GetApiAuthenticationTimeout

`func (o *MetricsHttpServletExtensionResponse) GetApiAuthenticationTimeout() string`

GetApiAuthenticationTimeout returns the ApiAuthenticationTimeout field if non-nil, zero value otherwise.

### GetApiAuthenticationTimeoutOk

`func (o *MetricsHttpServletExtensionResponse) GetApiAuthenticationTimeoutOk() (*string, bool)`

GetApiAuthenticationTimeoutOk returns a tuple with the ApiAuthenticationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAuthenticationTimeout

`func (o *MetricsHttpServletExtensionResponse) SetApiAuthenticationTimeout(v string)`

SetApiAuthenticationTimeout sets ApiAuthenticationTimeout field to given value.

### HasApiAuthenticationTimeout

`func (o *MetricsHttpServletExtensionResponse) HasApiAuthenticationTimeout() bool`

HasApiAuthenticationTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *MetricsHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *MetricsHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *MetricsHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *MetricsHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *MetricsHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *MetricsHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *MetricsHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *MetricsHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *MetricsHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *MetricsHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *MetricsHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *MetricsHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *MetricsHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


