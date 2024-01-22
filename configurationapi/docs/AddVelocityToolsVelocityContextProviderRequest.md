# AddVelocityToolsVelocityContextProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvelocityToolsVelocityContextProviderSchemaUrn**](EnumvelocityToolsVelocityContextProviderSchemaUrn.md) |  | 
**RequestTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized for each request. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**SessionTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized for each session. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**ApplicationTool** | Pointer to **[]string** | The fully-qualified name of a Velocity Tool class that will be initialized once for the life of the server. May optionally include a path to a properties file used to configure this tool separated from the class name by a semi-colon (;). The path may absolute or relative to the server root. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Context Provider is enabled. If set to &#39;false&#39; this Velocity Context Provider will not contribute context content for any requests. | [optional] 
**ObjectScope** | Pointer to [**EnumvelocityContextProviderObjectScopeProp**](EnumvelocityContextProviderObjectScopeProp.md) |  | [optional] 
**IncludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will contribute content. | [optional] 
**ExcludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will not contribute content. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content. | [optional] 
**ProviderName** | **string** | Name of the new Velocity Context Provider | 

## Methods

### NewAddVelocityToolsVelocityContextProviderRequest

`func NewAddVelocityToolsVelocityContextProviderRequest(schemas []EnumvelocityToolsVelocityContextProviderSchemaUrn, providerName string, ) *AddVelocityToolsVelocityContextProviderRequest`

NewAddVelocityToolsVelocityContextProviderRequest instantiates a new AddVelocityToolsVelocityContextProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVelocityToolsVelocityContextProviderRequestWithDefaults

`func NewAddVelocityToolsVelocityContextProviderRequestWithDefaults() *AddVelocityToolsVelocityContextProviderRequest`

NewAddVelocityToolsVelocityContextProviderRequestWithDefaults instantiates a new AddVelocityToolsVelocityContextProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetSchemas() []EnumvelocityToolsVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetSchemasOk() (*[]EnumvelocityToolsVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetSchemas(v []EnumvelocityToolsVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetRequestTool() []string`

GetRequestTool returns the RequestTool field if non-nil, zero value otherwise.

### GetRequestToolOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetRequestToolOk() (*[]string, bool)`

GetRequestToolOk returns a tuple with the RequestTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetRequestTool(v []string)`

SetRequestTool sets RequestTool field to given value.

### HasRequestTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasRequestTool() bool`

HasRequestTool returns a boolean if a field has been set.

### GetSessionTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetSessionTool() []string`

GetSessionTool returns the SessionTool field if non-nil, zero value otherwise.

### GetSessionToolOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetSessionToolOk() (*[]string, bool)`

GetSessionToolOk returns a tuple with the SessionTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetSessionTool(v []string)`

SetSessionTool sets SessionTool field to given value.

### HasSessionTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasSessionTool() bool`

HasSessionTool returns a boolean if a field has been set.

### GetApplicationTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetApplicationTool() []string`

GetApplicationTool returns the ApplicationTool field if non-nil, zero value otherwise.

### GetApplicationToolOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetApplicationToolOk() (*[]string, bool)`

GetApplicationToolOk returns a tuple with the ApplicationTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetApplicationTool(v []string)`

SetApplicationTool sets ApplicationTool field to given value.

### HasApplicationTool

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasApplicationTool() bool`

HasApplicationTool returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddVelocityToolsVelocityContextProviderRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetProviderName

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddVelocityToolsVelocityContextProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddVelocityToolsVelocityContextProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


