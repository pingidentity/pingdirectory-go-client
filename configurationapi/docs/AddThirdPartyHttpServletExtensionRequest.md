# AddThirdPartyHttpServletExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | **string** | Name of the new HTTP Servlet Extension | 
**Schemas** | [**[]EnumthirdPartyHttpServletExtensionSchemaUrn**](EnumthirdPartyHttpServletExtensionSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party HTTP Servlet Extension. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewAddThirdPartyHttpServletExtensionRequest

`func NewAddThirdPartyHttpServletExtensionRequest(extensionName string, schemas []EnumthirdPartyHttpServletExtensionSchemaUrn, extensionClass string, ) *AddThirdPartyHttpServletExtensionRequest`

NewAddThirdPartyHttpServletExtensionRequest instantiates a new AddThirdPartyHttpServletExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyHttpServletExtensionRequestWithDefaults

`func NewAddThirdPartyHttpServletExtensionRequestWithDefaults() *AddThirdPartyHttpServletExtensionRequest`

NewAddThirdPartyHttpServletExtensionRequestWithDefaults instantiates a new AddThirdPartyHttpServletExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *AddThirdPartyHttpServletExtensionRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *AddThirdPartyHttpServletExtensionRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetSchemas

`func (o *AddThirdPartyHttpServletExtensionRequest) GetSchemas() []EnumthirdPartyHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetSchemasOk() (*[]EnumthirdPartyHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyHttpServletExtensionRequest) SetSchemas(v []EnumthirdPartyHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyHttpServletExtensionRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyHttpServletExtensionRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyHttpServletExtensionRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyHttpServletExtensionRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyHttpServletExtensionRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyHttpServletExtensionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyHttpServletExtensionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyHttpServletExtensionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AddThirdPartyHttpServletExtensionRequest) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AddThirdPartyHttpServletExtensionRequest) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AddThirdPartyHttpServletExtensionRequest) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddThirdPartyHttpServletExtensionRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddThirdPartyHttpServletExtensionRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddThirdPartyHttpServletExtensionRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddThirdPartyHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddThirdPartyHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddThirdPartyHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddThirdPartyHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


