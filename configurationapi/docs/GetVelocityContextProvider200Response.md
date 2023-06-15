# GetVelocityContextProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Velocity Context Provider | 
**Schemas** | [**[]EnumthirdPartyVelocityContextProviderSchemaUrn**](EnumthirdPartyVelocityContextProviderSchemaUrn.md) |  | 
**RequestTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized for each request. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**SessionTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized for each session. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**ApplicationTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized once for the life of the server. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Context Provider is enabled. If set to &#39;false&#39; this Velocity Context Provider will not contribute context content for any requests. | [optional] 
**ObjectScope** | Pointer to [**EnumvelocityContextProviderObjectScopeProp**](EnumvelocityContextProviderObjectScopeProp.md) |  | [optional] 
**IncludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will contribute content. | [optional] 
**ExcludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will not contribute content. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**HttpMethod** | Pointer to **[]string** | Specifies the set of HTTP methods handled by this Velocity Context Provider, which will perform actions necessary to fulfill the request before updating the context for the response. The values of this property are not case-sensitive. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Velocity Context Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Velocity Context Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewGetVelocityContextProvider200Response

`func NewGetVelocityContextProvider200Response(id string, schemas []EnumthirdPartyVelocityContextProviderSchemaUrn, extensionClass string, ) *GetVelocityContextProvider200Response`

NewGetVelocityContextProvider200Response instantiates a new GetVelocityContextProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVelocityContextProvider200ResponseWithDefaults

`func NewGetVelocityContextProvider200ResponseWithDefaults() *GetVelocityContextProvider200Response`

NewGetVelocityContextProvider200ResponseWithDefaults instantiates a new GetVelocityContextProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVelocityContextProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVelocityContextProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVelocityContextProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GetVelocityContextProvider200Response) GetSchemas() []EnumthirdPartyVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetVelocityContextProvider200Response) GetSchemasOk() (*[]EnumthirdPartyVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetVelocityContextProvider200Response) SetSchemas(v []EnumthirdPartyVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestTool

`func (o *GetVelocityContextProvider200Response) GetRequestTool() []string`

GetRequestTool returns the RequestTool field if non-nil, zero value otherwise.

### GetRequestToolOk

`func (o *GetVelocityContextProvider200Response) GetRequestToolOk() (*[]string, bool)`

GetRequestToolOk returns a tuple with the RequestTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTool

`func (o *GetVelocityContextProvider200Response) SetRequestTool(v []string)`

SetRequestTool sets RequestTool field to given value.

### HasRequestTool

`func (o *GetVelocityContextProvider200Response) HasRequestTool() bool`

HasRequestTool returns a boolean if a field has been set.

### GetSessionTool

`func (o *GetVelocityContextProvider200Response) GetSessionTool() []string`

GetSessionTool returns the SessionTool field if non-nil, zero value otherwise.

### GetSessionToolOk

`func (o *GetVelocityContextProvider200Response) GetSessionToolOk() (*[]string, bool)`

GetSessionToolOk returns a tuple with the SessionTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTool

`func (o *GetVelocityContextProvider200Response) SetSessionTool(v []string)`

SetSessionTool sets SessionTool field to given value.

### HasSessionTool

`func (o *GetVelocityContextProvider200Response) HasSessionTool() bool`

HasSessionTool returns a boolean if a field has been set.

### GetApplicationTool

`func (o *GetVelocityContextProvider200Response) GetApplicationTool() []string`

GetApplicationTool returns the ApplicationTool field if non-nil, zero value otherwise.

### GetApplicationToolOk

`func (o *GetVelocityContextProvider200Response) GetApplicationToolOk() (*[]string, bool)`

GetApplicationToolOk returns a tuple with the ApplicationTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTool

`func (o *GetVelocityContextProvider200Response) SetApplicationTool(v []string)`

SetApplicationTool sets ApplicationTool field to given value.

### HasApplicationTool

`func (o *GetVelocityContextProvider200Response) HasApplicationTool() bool`

HasApplicationTool returns a boolean if a field has been set.

### GetEnabled

`func (o *GetVelocityContextProvider200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetVelocityContextProvider200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetVelocityContextProvider200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetVelocityContextProvider200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *GetVelocityContextProvider200Response) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *GetVelocityContextProvider200Response) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *GetVelocityContextProvider200Response) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *GetVelocityContextProvider200Response) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *GetVelocityContextProvider200Response) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *GetVelocityContextProvider200Response) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *GetVelocityContextProvider200Response) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *GetVelocityContextProvider200Response) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *GetVelocityContextProvider200Response) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *GetVelocityContextProvider200Response) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *GetVelocityContextProvider200Response) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *GetVelocityContextProvider200Response) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetResponseHeader

`func (o *GetVelocityContextProvider200Response) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *GetVelocityContextProvider200Response) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *GetVelocityContextProvider200Response) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *GetVelocityContextProvider200Response) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *GetVelocityContextProvider200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetVelocityContextProvider200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetVelocityContextProvider200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetVelocityContextProvider200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetVelocityContextProvider200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetVelocityContextProvider200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetVelocityContextProvider200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetVelocityContextProvider200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetHttpMethod

`func (o *GetVelocityContextProvider200Response) GetHttpMethod() []string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *GetVelocityContextProvider200Response) GetHttpMethodOk() (*[]string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *GetVelocityContextProvider200Response) SetHttpMethod(v []string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *GetVelocityContextProvider200Response) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetExtensionClass

`func (o *GetVelocityContextProvider200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetVelocityContextProvider200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetVelocityContextProvider200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetVelocityContextProvider200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetVelocityContextProvider200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetVelocityContextProvider200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetVelocityContextProvider200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


