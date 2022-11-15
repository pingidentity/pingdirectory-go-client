# GroovyScriptedHttpServletExtensionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedHttpServletExtensionSchemaUrn**](EnumgroovyScriptedHttpServletExtensionSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Servlet Extension. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewGroovyScriptedHttpServletExtensionShared

`func NewGroovyScriptedHttpServletExtensionShared(schemas []EnumgroovyScriptedHttpServletExtensionSchemaUrn, scriptClass string, ) *GroovyScriptedHttpServletExtensionShared`

NewGroovyScriptedHttpServletExtensionShared instantiates a new GroovyScriptedHttpServletExtensionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedHttpServletExtensionSharedWithDefaults

`func NewGroovyScriptedHttpServletExtensionSharedWithDefaults() *GroovyScriptedHttpServletExtensionShared`

NewGroovyScriptedHttpServletExtensionSharedWithDefaults instantiates a new GroovyScriptedHttpServletExtensionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedHttpServletExtensionShared) GetSchemas() []EnumgroovyScriptedHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetSchemasOk() (*[]EnumgroovyScriptedHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedHttpServletExtensionShared) SetSchemas(v []EnumgroovyScriptedHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedHttpServletExtensionShared) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedHttpServletExtensionShared) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedHttpServletExtensionShared) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedHttpServletExtensionShared) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedHttpServletExtensionShared) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedHttpServletExtensionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedHttpServletExtensionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedHttpServletExtensionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *GroovyScriptedHttpServletExtensionShared) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *GroovyScriptedHttpServletExtensionShared) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *GroovyScriptedHttpServletExtensionShared) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *GroovyScriptedHttpServletExtensionShared) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *GroovyScriptedHttpServletExtensionShared) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *GroovyScriptedHttpServletExtensionShared) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *GroovyScriptedHttpServletExtensionShared) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *GroovyScriptedHttpServletExtensionShared) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *GroovyScriptedHttpServletExtensionShared) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *GroovyScriptedHttpServletExtensionShared) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


