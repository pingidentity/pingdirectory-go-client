# QuickstartHttpServletExtensionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumquickstartHttpServletExtensionSchemaUrn**](EnumquickstartHttpServletExtensionSchemaUrn.md) |  | 
**Server** | Pointer to **string** | Specifies the PingFederate server to be configured. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewQuickstartHttpServletExtensionShared

`func NewQuickstartHttpServletExtensionShared(schemas []EnumquickstartHttpServletExtensionSchemaUrn, ) *QuickstartHttpServletExtensionShared`

NewQuickstartHttpServletExtensionShared instantiates a new QuickstartHttpServletExtensionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickstartHttpServletExtensionSharedWithDefaults

`func NewQuickstartHttpServletExtensionSharedWithDefaults() *QuickstartHttpServletExtensionShared`

NewQuickstartHttpServletExtensionSharedWithDefaults instantiates a new QuickstartHttpServletExtensionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *QuickstartHttpServletExtensionShared) GetSchemas() []EnumquickstartHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *QuickstartHttpServletExtensionShared) GetSchemasOk() (*[]EnumquickstartHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *QuickstartHttpServletExtensionShared) SetSchemas(v []EnumquickstartHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *QuickstartHttpServletExtensionShared) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *QuickstartHttpServletExtensionShared) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *QuickstartHttpServletExtensionShared) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *QuickstartHttpServletExtensionShared) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetDescription

`func (o *QuickstartHttpServletExtensionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickstartHttpServletExtensionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickstartHttpServletExtensionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickstartHttpServletExtensionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *QuickstartHttpServletExtensionShared) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *QuickstartHttpServletExtensionShared) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *QuickstartHttpServletExtensionShared) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *QuickstartHttpServletExtensionShared) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *QuickstartHttpServletExtensionShared) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *QuickstartHttpServletExtensionShared) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *QuickstartHttpServletExtensionShared) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *QuickstartHttpServletExtensionShared) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *QuickstartHttpServletExtensionShared) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *QuickstartHttpServletExtensionShared) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *QuickstartHttpServletExtensionShared) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *QuickstartHttpServletExtensionShared) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


