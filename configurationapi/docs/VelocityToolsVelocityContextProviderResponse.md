# VelocityToolsVelocityContextProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Velocity Context Provider | 
**Schemas** | [**[]EnumvelocityToolsVelocityContextProviderSchemaUrn**](EnumvelocityToolsVelocityContextProviderSchemaUrn.md) |  | 
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

## Methods

### NewVelocityToolsVelocityContextProviderResponse

`func NewVelocityToolsVelocityContextProviderResponse(id string, schemas []EnumvelocityToolsVelocityContextProviderSchemaUrn, ) *VelocityToolsVelocityContextProviderResponse`

NewVelocityToolsVelocityContextProviderResponse instantiates a new VelocityToolsVelocityContextProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityToolsVelocityContextProviderResponseWithDefaults

`func NewVelocityToolsVelocityContextProviderResponseWithDefaults() *VelocityToolsVelocityContextProviderResponse`

NewVelocityToolsVelocityContextProviderResponseWithDefaults instantiates a new VelocityToolsVelocityContextProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VelocityToolsVelocityContextProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VelocityToolsVelocityContextProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *VelocityToolsVelocityContextProviderResponse) GetSchemas() []EnumvelocityToolsVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetSchemasOk() (*[]EnumvelocityToolsVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VelocityToolsVelocityContextProviderResponse) SetSchemas(v []EnumvelocityToolsVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestTool

`func (o *VelocityToolsVelocityContextProviderResponse) GetRequestTool() []string`

GetRequestTool returns the RequestTool field if non-nil, zero value otherwise.

### GetRequestToolOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetRequestToolOk() (*[]string, bool)`

GetRequestToolOk returns a tuple with the RequestTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTool

`func (o *VelocityToolsVelocityContextProviderResponse) SetRequestTool(v []string)`

SetRequestTool sets RequestTool field to given value.

### HasRequestTool

`func (o *VelocityToolsVelocityContextProviderResponse) HasRequestTool() bool`

HasRequestTool returns a boolean if a field has been set.

### GetSessionTool

`func (o *VelocityToolsVelocityContextProviderResponse) GetSessionTool() []string`

GetSessionTool returns the SessionTool field if non-nil, zero value otherwise.

### GetSessionToolOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetSessionToolOk() (*[]string, bool)`

GetSessionToolOk returns a tuple with the SessionTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTool

`func (o *VelocityToolsVelocityContextProviderResponse) SetSessionTool(v []string)`

SetSessionTool sets SessionTool field to given value.

### HasSessionTool

`func (o *VelocityToolsVelocityContextProviderResponse) HasSessionTool() bool`

HasSessionTool returns a boolean if a field has been set.

### GetApplicationTool

`func (o *VelocityToolsVelocityContextProviderResponse) GetApplicationTool() []string`

GetApplicationTool returns the ApplicationTool field if non-nil, zero value otherwise.

### GetApplicationToolOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetApplicationToolOk() (*[]string, bool)`

GetApplicationToolOk returns a tuple with the ApplicationTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTool

`func (o *VelocityToolsVelocityContextProviderResponse) SetApplicationTool(v []string)`

SetApplicationTool sets ApplicationTool field to given value.

### HasApplicationTool

`func (o *VelocityToolsVelocityContextProviderResponse) HasApplicationTool() bool`

HasApplicationTool returns a boolean if a field has been set.

### GetEnabled

`func (o *VelocityToolsVelocityContextProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VelocityToolsVelocityContextProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VelocityToolsVelocityContextProviderResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *VelocityToolsVelocityContextProviderResponse) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *VelocityToolsVelocityContextProviderResponse) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *VelocityToolsVelocityContextProviderResponse) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *VelocityToolsVelocityContextProviderResponse) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *VelocityToolsVelocityContextProviderResponse) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *VelocityToolsVelocityContextProviderResponse) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *VelocityToolsVelocityContextProviderResponse) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *VelocityToolsVelocityContextProviderResponse) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *VelocityToolsVelocityContextProviderResponse) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetResponseHeader

`func (o *VelocityToolsVelocityContextProviderResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *VelocityToolsVelocityContextProviderResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *VelocityToolsVelocityContextProviderResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *VelocityToolsVelocityContextProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VelocityToolsVelocityContextProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VelocityToolsVelocityContextProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VelocityToolsVelocityContextProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityToolsVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *VelocityToolsVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityToolsVelocityContextProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityToolsVelocityContextProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


