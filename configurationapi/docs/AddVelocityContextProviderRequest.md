# AddVelocityContextProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Velocity Context Provider | 
**Schemas** | [**[]EnumthirdPartyVelocityContextProviderSchemaUrn**](EnumthirdPartyVelocityContextProviderSchemaUrn.md) |  | 
**RequestTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized for each request. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**SessionTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized for each session. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**ApplicationTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized once for the life of the server. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Context Provider is enabled. If set to &#39;false&#39; this Velocity Context Provider will not contribute context content for any requests. | [optional] 
**ObjectScope** | Pointer to [**EnumvelocityContextProviderObjectScopeProp**](EnumvelocityContextProviderObjectScopeProp.md) |  | [optional] 
**IncludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will contribute content. | [optional] 
**ExcludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will not contribute content. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Velocity Context Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Velocity Context Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**HttpMethod** | Pointer to **[]string** | Specifies the set of HTTP methods handled by this Velocity Context Provider, which will perform actions necessary to fulfill the request before updating the context for the response. The values of this property are not case-sensitive. | [optional] 

## Methods

### NewAddVelocityContextProviderRequest

`func NewAddVelocityContextProviderRequest(providerName string, schemas []EnumthirdPartyVelocityContextProviderSchemaUrn, extensionClass string, ) *AddVelocityContextProviderRequest`

NewAddVelocityContextProviderRequest instantiates a new AddVelocityContextProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVelocityContextProviderRequestWithDefaults

`func NewAddVelocityContextProviderRequestWithDefaults() *AddVelocityContextProviderRequest`

NewAddVelocityContextProviderRequestWithDefaults instantiates a new AddVelocityContextProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddVelocityContextProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddVelocityContextProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddVelocityContextProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddVelocityContextProviderRequest) GetSchemas() []EnumthirdPartyVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVelocityContextProviderRequest) GetSchemasOk() (*[]EnumthirdPartyVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVelocityContextProviderRequest) SetSchemas(v []EnumthirdPartyVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestTool

`func (o *AddVelocityContextProviderRequest) GetRequestTool() []string`

GetRequestTool returns the RequestTool field if non-nil, zero value otherwise.

### GetRequestToolOk

`func (o *AddVelocityContextProviderRequest) GetRequestToolOk() (*[]string, bool)`

GetRequestToolOk returns a tuple with the RequestTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTool

`func (o *AddVelocityContextProviderRequest) SetRequestTool(v []string)`

SetRequestTool sets RequestTool field to given value.

### HasRequestTool

`func (o *AddVelocityContextProviderRequest) HasRequestTool() bool`

HasRequestTool returns a boolean if a field has been set.

### GetSessionTool

`func (o *AddVelocityContextProviderRequest) GetSessionTool() []string`

GetSessionTool returns the SessionTool field if non-nil, zero value otherwise.

### GetSessionToolOk

`func (o *AddVelocityContextProviderRequest) GetSessionToolOk() (*[]string, bool)`

GetSessionToolOk returns a tuple with the SessionTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTool

`func (o *AddVelocityContextProviderRequest) SetSessionTool(v []string)`

SetSessionTool sets SessionTool field to given value.

### HasSessionTool

`func (o *AddVelocityContextProviderRequest) HasSessionTool() bool`

HasSessionTool returns a boolean if a field has been set.

### GetApplicationTool

`func (o *AddVelocityContextProviderRequest) GetApplicationTool() []string`

GetApplicationTool returns the ApplicationTool field if non-nil, zero value otherwise.

### GetApplicationToolOk

`func (o *AddVelocityContextProviderRequest) GetApplicationToolOk() (*[]string, bool)`

GetApplicationToolOk returns a tuple with the ApplicationTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTool

`func (o *AddVelocityContextProviderRequest) SetApplicationTool(v []string)`

SetApplicationTool sets ApplicationTool field to given value.

### HasApplicationTool

`func (o *AddVelocityContextProviderRequest) HasApplicationTool() bool`

HasApplicationTool returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVelocityContextProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVelocityContextProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVelocityContextProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddVelocityContextProviderRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *AddVelocityContextProviderRequest) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *AddVelocityContextProviderRequest) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *AddVelocityContextProviderRequest) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *AddVelocityContextProviderRequest) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *AddVelocityContextProviderRequest) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *AddVelocityContextProviderRequest) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *AddVelocityContextProviderRequest) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *AddVelocityContextProviderRequest) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *AddVelocityContextProviderRequest) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *AddVelocityContextProviderRequest) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *AddVelocityContextProviderRequest) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *AddVelocityContextProviderRequest) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddVelocityContextProviderRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddVelocityContextProviderRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddVelocityContextProviderRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddVelocityContextProviderRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddVelocityContextProviderRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddVelocityContextProviderRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddVelocityContextProviderRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddVelocityContextProviderRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddVelocityContextProviderRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddVelocityContextProviderRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddVelocityContextProviderRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetHttpMethod

`func (o *AddVelocityContextProviderRequest) GetHttpMethod() []string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *AddVelocityContextProviderRequest) GetHttpMethodOk() (*[]string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *AddVelocityContextProviderRequest) SetHttpMethod(v []string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *AddVelocityContextProviderRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


