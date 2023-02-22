# AddQuickstartHttpServletExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | **string** | Name of the new HTTP Servlet Extension | 
**Schemas** | [**[]EnumquickstartHttpServletExtensionSchemaUrn**](EnumquickstartHttpServletExtensionSchemaUrn.md) |  | 
**Server** | Pointer to **string** | Specifies the PingFederate server to be configured. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewAddQuickstartHttpServletExtensionRequest

`func NewAddQuickstartHttpServletExtensionRequest(extensionName string, schemas []EnumquickstartHttpServletExtensionSchemaUrn, ) *AddQuickstartHttpServletExtensionRequest`

NewAddQuickstartHttpServletExtensionRequest instantiates a new AddQuickstartHttpServletExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddQuickstartHttpServletExtensionRequestWithDefaults

`func NewAddQuickstartHttpServletExtensionRequestWithDefaults() *AddQuickstartHttpServletExtensionRequest`

NewAddQuickstartHttpServletExtensionRequestWithDefaults instantiates a new AddQuickstartHttpServletExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *AddQuickstartHttpServletExtensionRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *AddQuickstartHttpServletExtensionRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetSchemas

`func (o *AddQuickstartHttpServletExtensionRequest) GetSchemas() []EnumquickstartHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetSchemasOk() (*[]EnumquickstartHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddQuickstartHttpServletExtensionRequest) SetSchemas(v []EnumquickstartHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *AddQuickstartHttpServletExtensionRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddQuickstartHttpServletExtensionRequest) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *AddQuickstartHttpServletExtensionRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetDescription

`func (o *AddQuickstartHttpServletExtensionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddQuickstartHttpServletExtensionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddQuickstartHttpServletExtensionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AddQuickstartHttpServletExtensionRequest) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AddQuickstartHttpServletExtensionRequest) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AddQuickstartHttpServletExtensionRequest) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddQuickstartHttpServletExtensionRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddQuickstartHttpServletExtensionRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddQuickstartHttpServletExtensionRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddQuickstartHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddQuickstartHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddQuickstartHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddQuickstartHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


