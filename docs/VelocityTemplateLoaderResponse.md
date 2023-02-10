# VelocityTemplateLoaderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Velocity Template Loader | 
**Schemas** | Pointer to [**[]EnumvelocityTemplateLoaderSchemaUrn**](EnumvelocityTemplateLoaderSchemaUrn.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Template Loader is enabled. | [optional] 
**EvaluationOrderIndex** | **int32** | This property determines the evaluation order for determining the correct Velocity Template Loader to load a template for generating content for a particular request. | 
**MimeTypeMatcher** | **string** | Specifies a media type for matching Accept request-header values. | 
**MimeType** | Pointer to **string** | Specifies a the value that will be used in the response&#39;s Content-Type header that indicates the type of content to return. | [optional] 
**TemplateSuffix** | Pointer to **string** | Specifies the suffix to append to the requested resource name when searching for the template file with which to form a response. | [optional] 
**TemplateDirectory** | Pointer to **string** | Specifies the directory in which to search for the template files. | [optional] 

## Methods

### NewVelocityTemplateLoaderResponse

`func NewVelocityTemplateLoaderResponse(id string, evaluationOrderIndex int32, mimeTypeMatcher string, ) *VelocityTemplateLoaderResponse`

NewVelocityTemplateLoaderResponse instantiates a new VelocityTemplateLoaderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityTemplateLoaderResponseWithDefaults

`func NewVelocityTemplateLoaderResponseWithDefaults() *VelocityTemplateLoaderResponse`

NewVelocityTemplateLoaderResponseWithDefaults instantiates a new VelocityTemplateLoaderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *VelocityTemplateLoaderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VelocityTemplateLoaderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VelocityTemplateLoaderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VelocityTemplateLoaderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityTemplateLoaderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *VelocityTemplateLoaderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityTemplateLoaderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *VelocityTemplateLoaderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *VelocityTemplateLoaderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VelocityTemplateLoaderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VelocityTemplateLoaderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *VelocityTemplateLoaderResponse) GetSchemas() []EnumvelocityTemplateLoaderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VelocityTemplateLoaderResponse) GetSchemasOk() (*[]EnumvelocityTemplateLoaderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VelocityTemplateLoaderResponse) SetSchemas(v []EnumvelocityTemplateLoaderSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *VelocityTemplateLoaderResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetEnabled

`func (o *VelocityTemplateLoaderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VelocityTemplateLoaderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VelocityTemplateLoaderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VelocityTemplateLoaderResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *VelocityTemplateLoaderResponse) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *VelocityTemplateLoaderResponse) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *VelocityTemplateLoaderResponse) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetMimeTypeMatcher

`func (o *VelocityTemplateLoaderResponse) GetMimeTypeMatcher() string`

GetMimeTypeMatcher returns the MimeTypeMatcher field if non-nil, zero value otherwise.

### GetMimeTypeMatcherOk

`func (o *VelocityTemplateLoaderResponse) GetMimeTypeMatcherOk() (*string, bool)`

GetMimeTypeMatcherOk returns a tuple with the MimeTypeMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypeMatcher

`func (o *VelocityTemplateLoaderResponse) SetMimeTypeMatcher(v string)`

SetMimeTypeMatcher sets MimeTypeMatcher field to given value.


### GetMimeType

`func (o *VelocityTemplateLoaderResponse) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *VelocityTemplateLoaderResponse) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *VelocityTemplateLoaderResponse) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *VelocityTemplateLoaderResponse) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetTemplateSuffix

`func (o *VelocityTemplateLoaderResponse) GetTemplateSuffix() string`

GetTemplateSuffix returns the TemplateSuffix field if non-nil, zero value otherwise.

### GetTemplateSuffixOk

`func (o *VelocityTemplateLoaderResponse) GetTemplateSuffixOk() (*string, bool)`

GetTemplateSuffixOk returns a tuple with the TemplateSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSuffix

`func (o *VelocityTemplateLoaderResponse) SetTemplateSuffix(v string)`

SetTemplateSuffix sets TemplateSuffix field to given value.

### HasTemplateSuffix

`func (o *VelocityTemplateLoaderResponse) HasTemplateSuffix() bool`

HasTemplateSuffix returns a boolean if a field has been set.

### GetTemplateDirectory

`func (o *VelocityTemplateLoaderResponse) GetTemplateDirectory() string`

GetTemplateDirectory returns the TemplateDirectory field if non-nil, zero value otherwise.

### GetTemplateDirectoryOk

`func (o *VelocityTemplateLoaderResponse) GetTemplateDirectoryOk() (*string, bool)`

GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDirectory

`func (o *VelocityTemplateLoaderResponse) SetTemplateDirectory(v string)`

SetTemplateDirectory sets TemplateDirectory field to given value.

### HasTemplateDirectory

`func (o *VelocityTemplateLoaderResponse) HasTemplateDirectory() bool`

HasTemplateDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


