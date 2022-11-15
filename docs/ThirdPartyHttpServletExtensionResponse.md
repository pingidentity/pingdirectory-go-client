# ThirdPartyHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Schemas** | [**[]EnumthirdPartyHttpServletExtensionSchemaUrn**](EnumthirdPartyHttpServletExtensionSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party HTTP Servlet Extension. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewThirdPartyHttpServletExtensionResponse

`func NewThirdPartyHttpServletExtensionResponse(id string, schemas []EnumthirdPartyHttpServletExtensionSchemaUrn, extensionClass string, ) *ThirdPartyHttpServletExtensionResponse`

NewThirdPartyHttpServletExtensionResponse instantiates a new ThirdPartyHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyHttpServletExtensionResponseWithDefaults

`func NewThirdPartyHttpServletExtensionResponseWithDefaults() *ThirdPartyHttpServletExtensionResponse`

NewThirdPartyHttpServletExtensionResponseWithDefaults instantiates a new ThirdPartyHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartyHttpServletExtensionResponse) GetSchemas() []EnumthirdPartyHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetSchemasOk() (*[]EnumthirdPartyHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyHttpServletExtensionResponse) SetSchemas(v []EnumthirdPartyHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyHttpServletExtensionResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyHttpServletExtensionResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyHttpServletExtensionResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyHttpServletExtensionResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyHttpServletExtensionResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *ThirdPartyHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *ThirdPartyHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *ThirdPartyHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *ThirdPartyHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *ThirdPartyHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *ThirdPartyHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *ThirdPartyHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *ThirdPartyHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *ThirdPartyHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *ThirdPartyHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


