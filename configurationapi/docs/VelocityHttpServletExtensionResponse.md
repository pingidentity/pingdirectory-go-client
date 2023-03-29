# VelocityHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvelocityHttpServletExtensionSchemaUrn**](EnumvelocityHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**BaseContextPath** | **string** | The context path to use to access all template-based and static content. The value must start with a forward slash and must represent a valid HTTP context path. | 
**StaticContextPath** | Pointer to **string** | The path below the base context path by which static, non-template content such as images, CSS, and Javascript files are accessible. | [optional] 
**StaticContentDirectory** | Pointer to **string** | Specifies the base directory in which static, non-template content such as images, CSS, and Javascript files are stored on the filesystem. | [optional] 
**StaticCustomDirectory** | Pointer to **string** | Specifies the base directory in which custom static, non-template content such as images, CSS, and Javascript files are stored on the filesystem. Files in this directory will override those with the same name in the directory specified by the static-content-directory property. | [optional] 
**TemplateDirectory** | **[]string** | Specifies an ordered list of directories in which to search for the template files. | 
**ExposeRequestAttributes** | Pointer to **bool** | Specifies whether the HTTP request will be exposed to templates. | [optional] 
**ExposeSessionAttributes** | Pointer to **bool** | Specifies whether the HTTP session will be exposed to templates. | [optional] 
**ExposeServerContext** | Pointer to **bool** | Specifies whether a server context will be exposed under context key &#39;ubid_server&#39; for all template contexts. | [optional] 
**AllowContextOverride** | Pointer to **bool** | Indicates whether context providers may override existing context objects with new values. | [optional] 
**MimeTypesFile** | Pointer to **string** | Specifies the path to a file that contains MIME type mappings that will be used to determine the appropriate value to return for the Content-Type header based on the extension of the requested static content file. | [optional] 
**DefaultMIMEType** | Pointer to **string** | Specifies the default value that will be used in the response&#39;s Content-Type header that indicates the type of content to return. | [optional] 
**CharacterEncoding** | Pointer to **string** | Specifies the value that will be used for all responses&#39; Content-Type headers&#39; charset parameter that indicates the character encoding of the document. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all template page requests. | [optional] 
**StaticResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for static content requests such as images and scripts. | [optional] 
**RequireAuthentication** | Pointer to **bool** | Require authentication when accessing Velocity templates. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the identity mapper that is to be used for associating basic authentication credentials with user entries. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewVelocityHttpServletExtensionResponse

`func NewVelocityHttpServletExtensionResponse(schemas []EnumvelocityHttpServletExtensionSchemaUrn, id string, baseContextPath string, templateDirectory []string, ) *VelocityHttpServletExtensionResponse`

NewVelocityHttpServletExtensionResponse instantiates a new VelocityHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityHttpServletExtensionResponseWithDefaults

`func NewVelocityHttpServletExtensionResponseWithDefaults() *VelocityHttpServletExtensionResponse`

NewVelocityHttpServletExtensionResponseWithDefaults instantiates a new VelocityHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VelocityHttpServletExtensionResponse) GetSchemas() []EnumvelocityHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VelocityHttpServletExtensionResponse) GetSchemasOk() (*[]EnumvelocityHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VelocityHttpServletExtensionResponse) SetSchemas(v []EnumvelocityHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *VelocityHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VelocityHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VelocityHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBaseContextPath

`func (o *VelocityHttpServletExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *VelocityHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *VelocityHttpServletExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetStaticContextPath

`func (o *VelocityHttpServletExtensionResponse) GetStaticContextPath() string`

GetStaticContextPath returns the StaticContextPath field if non-nil, zero value otherwise.

### GetStaticContextPathOk

`func (o *VelocityHttpServletExtensionResponse) GetStaticContextPathOk() (*string, bool)`

GetStaticContextPathOk returns a tuple with the StaticContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContextPath

`func (o *VelocityHttpServletExtensionResponse) SetStaticContextPath(v string)`

SetStaticContextPath sets StaticContextPath field to given value.

### HasStaticContextPath

`func (o *VelocityHttpServletExtensionResponse) HasStaticContextPath() bool`

HasStaticContextPath returns a boolean if a field has been set.

### GetStaticContentDirectory

`func (o *VelocityHttpServletExtensionResponse) GetStaticContentDirectory() string`

GetStaticContentDirectory returns the StaticContentDirectory field if non-nil, zero value otherwise.

### GetStaticContentDirectoryOk

`func (o *VelocityHttpServletExtensionResponse) GetStaticContentDirectoryOk() (*string, bool)`

GetStaticContentDirectoryOk returns a tuple with the StaticContentDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContentDirectory

`func (o *VelocityHttpServletExtensionResponse) SetStaticContentDirectory(v string)`

SetStaticContentDirectory sets StaticContentDirectory field to given value.

### HasStaticContentDirectory

`func (o *VelocityHttpServletExtensionResponse) HasStaticContentDirectory() bool`

HasStaticContentDirectory returns a boolean if a field has been set.

### GetStaticCustomDirectory

`func (o *VelocityHttpServletExtensionResponse) GetStaticCustomDirectory() string`

GetStaticCustomDirectory returns the StaticCustomDirectory field if non-nil, zero value otherwise.

### GetStaticCustomDirectoryOk

`func (o *VelocityHttpServletExtensionResponse) GetStaticCustomDirectoryOk() (*string, bool)`

GetStaticCustomDirectoryOk returns a tuple with the StaticCustomDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticCustomDirectory

`func (o *VelocityHttpServletExtensionResponse) SetStaticCustomDirectory(v string)`

SetStaticCustomDirectory sets StaticCustomDirectory field to given value.

### HasStaticCustomDirectory

`func (o *VelocityHttpServletExtensionResponse) HasStaticCustomDirectory() bool`

HasStaticCustomDirectory returns a boolean if a field has been set.

### GetTemplateDirectory

`func (o *VelocityHttpServletExtensionResponse) GetTemplateDirectory() []string`

GetTemplateDirectory returns the TemplateDirectory field if non-nil, zero value otherwise.

### GetTemplateDirectoryOk

`func (o *VelocityHttpServletExtensionResponse) GetTemplateDirectoryOk() (*[]string, bool)`

GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDirectory

`func (o *VelocityHttpServletExtensionResponse) SetTemplateDirectory(v []string)`

SetTemplateDirectory sets TemplateDirectory field to given value.


### GetExposeRequestAttributes

`func (o *VelocityHttpServletExtensionResponse) GetExposeRequestAttributes() bool`

GetExposeRequestAttributes returns the ExposeRequestAttributes field if non-nil, zero value otherwise.

### GetExposeRequestAttributesOk

`func (o *VelocityHttpServletExtensionResponse) GetExposeRequestAttributesOk() (*bool, bool)`

GetExposeRequestAttributesOk returns a tuple with the ExposeRequestAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeRequestAttributes

`func (o *VelocityHttpServletExtensionResponse) SetExposeRequestAttributes(v bool)`

SetExposeRequestAttributes sets ExposeRequestAttributes field to given value.

### HasExposeRequestAttributes

`func (o *VelocityHttpServletExtensionResponse) HasExposeRequestAttributes() bool`

HasExposeRequestAttributes returns a boolean if a field has been set.

### GetExposeSessionAttributes

`func (o *VelocityHttpServletExtensionResponse) GetExposeSessionAttributes() bool`

GetExposeSessionAttributes returns the ExposeSessionAttributes field if non-nil, zero value otherwise.

### GetExposeSessionAttributesOk

`func (o *VelocityHttpServletExtensionResponse) GetExposeSessionAttributesOk() (*bool, bool)`

GetExposeSessionAttributesOk returns a tuple with the ExposeSessionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeSessionAttributes

`func (o *VelocityHttpServletExtensionResponse) SetExposeSessionAttributes(v bool)`

SetExposeSessionAttributes sets ExposeSessionAttributes field to given value.

### HasExposeSessionAttributes

`func (o *VelocityHttpServletExtensionResponse) HasExposeSessionAttributes() bool`

HasExposeSessionAttributes returns a boolean if a field has been set.

### GetExposeServerContext

`func (o *VelocityHttpServletExtensionResponse) GetExposeServerContext() bool`

GetExposeServerContext returns the ExposeServerContext field if non-nil, zero value otherwise.

### GetExposeServerContextOk

`func (o *VelocityHttpServletExtensionResponse) GetExposeServerContextOk() (*bool, bool)`

GetExposeServerContextOk returns a tuple with the ExposeServerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeServerContext

`func (o *VelocityHttpServletExtensionResponse) SetExposeServerContext(v bool)`

SetExposeServerContext sets ExposeServerContext field to given value.

### HasExposeServerContext

`func (o *VelocityHttpServletExtensionResponse) HasExposeServerContext() bool`

HasExposeServerContext returns a boolean if a field has been set.

### GetAllowContextOverride

`func (o *VelocityHttpServletExtensionResponse) GetAllowContextOverride() bool`

GetAllowContextOverride returns the AllowContextOverride field if non-nil, zero value otherwise.

### GetAllowContextOverrideOk

`func (o *VelocityHttpServletExtensionResponse) GetAllowContextOverrideOk() (*bool, bool)`

GetAllowContextOverrideOk returns a tuple with the AllowContextOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowContextOverride

`func (o *VelocityHttpServletExtensionResponse) SetAllowContextOverride(v bool)`

SetAllowContextOverride sets AllowContextOverride field to given value.

### HasAllowContextOverride

`func (o *VelocityHttpServletExtensionResponse) HasAllowContextOverride() bool`

HasAllowContextOverride returns a boolean if a field has been set.

### GetMimeTypesFile

`func (o *VelocityHttpServletExtensionResponse) GetMimeTypesFile() string`

GetMimeTypesFile returns the MimeTypesFile field if non-nil, zero value otherwise.

### GetMimeTypesFileOk

`func (o *VelocityHttpServletExtensionResponse) GetMimeTypesFileOk() (*string, bool)`

GetMimeTypesFileOk returns a tuple with the MimeTypesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypesFile

`func (o *VelocityHttpServletExtensionResponse) SetMimeTypesFile(v string)`

SetMimeTypesFile sets MimeTypesFile field to given value.

### HasMimeTypesFile

`func (o *VelocityHttpServletExtensionResponse) HasMimeTypesFile() bool`

HasMimeTypesFile returns a boolean if a field has been set.

### GetDefaultMIMEType

`func (o *VelocityHttpServletExtensionResponse) GetDefaultMIMEType() string`

GetDefaultMIMEType returns the DefaultMIMEType field if non-nil, zero value otherwise.

### GetDefaultMIMETypeOk

`func (o *VelocityHttpServletExtensionResponse) GetDefaultMIMETypeOk() (*string, bool)`

GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMIMEType

`func (o *VelocityHttpServletExtensionResponse) SetDefaultMIMEType(v string)`

SetDefaultMIMEType sets DefaultMIMEType field to given value.

### HasDefaultMIMEType

`func (o *VelocityHttpServletExtensionResponse) HasDefaultMIMEType() bool`

HasDefaultMIMEType returns a boolean if a field has been set.

### GetCharacterEncoding

`func (o *VelocityHttpServletExtensionResponse) GetCharacterEncoding() string`

GetCharacterEncoding returns the CharacterEncoding field if non-nil, zero value otherwise.

### GetCharacterEncodingOk

`func (o *VelocityHttpServletExtensionResponse) GetCharacterEncodingOk() (*string, bool)`

GetCharacterEncodingOk returns a tuple with the CharacterEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterEncoding

`func (o *VelocityHttpServletExtensionResponse) SetCharacterEncoding(v string)`

SetCharacterEncoding sets CharacterEncoding field to given value.

### HasCharacterEncoding

`func (o *VelocityHttpServletExtensionResponse) HasCharacterEncoding() bool`

HasCharacterEncoding returns a boolean if a field has been set.

### GetResponseHeader

`func (o *VelocityHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *VelocityHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *VelocityHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *VelocityHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetStaticResponseHeader

`func (o *VelocityHttpServletExtensionResponse) GetStaticResponseHeader() []string`

GetStaticResponseHeader returns the StaticResponseHeader field if non-nil, zero value otherwise.

### GetStaticResponseHeaderOk

`func (o *VelocityHttpServletExtensionResponse) GetStaticResponseHeaderOk() (*[]string, bool)`

GetStaticResponseHeaderOk returns a tuple with the StaticResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticResponseHeader

`func (o *VelocityHttpServletExtensionResponse) SetStaticResponseHeader(v []string)`

SetStaticResponseHeader sets StaticResponseHeader field to given value.

### HasStaticResponseHeader

`func (o *VelocityHttpServletExtensionResponse) HasStaticResponseHeader() bool`

HasStaticResponseHeader returns a boolean if a field has been set.

### GetRequireAuthentication

`func (o *VelocityHttpServletExtensionResponse) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *VelocityHttpServletExtensionResponse) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *VelocityHttpServletExtensionResponse) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *VelocityHttpServletExtensionResponse) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *VelocityHttpServletExtensionResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *VelocityHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *VelocityHttpServletExtensionResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *VelocityHttpServletExtensionResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetDescription

`func (o *VelocityHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VelocityHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VelocityHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VelocityHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *VelocityHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *VelocityHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *VelocityHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *VelocityHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *VelocityHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *VelocityHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *VelocityHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *VelocityHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *VelocityHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VelocityHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VelocityHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VelocityHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *VelocityHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


