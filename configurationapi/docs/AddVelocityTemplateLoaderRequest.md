# AddVelocityTemplateLoaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoaderName** | **string** | Name of the new Velocity Template Loader | 
**Schemas** | Pointer to [**[]EnumvelocityTemplateLoaderSchemaUrn**](EnumvelocityTemplateLoaderSchemaUrn.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Template Loader is enabled. | [optional] 
**EvaluationOrderIndex** | **int32** | This property determines the evaluation order for determining the correct Velocity Template Loader to load a template for generating content for a particular request. | 
**MimeTypeMatcher** | **string** | Specifies a media type for matching Accept request-header values. | 
**MimeType** | Pointer to **string** | Specifies a the value that will be used in the response&#39;s Content-Type header that indicates the type of content to return. | [optional] 
**TemplateSuffix** | Pointer to **string** | Specifies the suffix to append to the requested resource name when searching for the template file with which to form a response. | [optional] 
**TemplateDirectory** | Pointer to **string** | Specifies the directory in which to search for the template files. | [optional] 

## Methods

### NewAddVelocityTemplateLoaderRequest

`func NewAddVelocityTemplateLoaderRequest(loaderName string, evaluationOrderIndex int32, mimeTypeMatcher string, ) *AddVelocityTemplateLoaderRequest`

NewAddVelocityTemplateLoaderRequest instantiates a new AddVelocityTemplateLoaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVelocityTemplateLoaderRequestWithDefaults

`func NewAddVelocityTemplateLoaderRequestWithDefaults() *AddVelocityTemplateLoaderRequest`

NewAddVelocityTemplateLoaderRequestWithDefaults instantiates a new AddVelocityTemplateLoaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoaderName

`func (o *AddVelocityTemplateLoaderRequest) GetLoaderName() string`

GetLoaderName returns the LoaderName field if non-nil, zero value otherwise.

### GetLoaderNameOk

`func (o *AddVelocityTemplateLoaderRequest) GetLoaderNameOk() (*string, bool)`

GetLoaderNameOk returns a tuple with the LoaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoaderName

`func (o *AddVelocityTemplateLoaderRequest) SetLoaderName(v string)`

SetLoaderName sets LoaderName field to given value.


### GetSchemas

`func (o *AddVelocityTemplateLoaderRequest) GetSchemas() []EnumvelocityTemplateLoaderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVelocityTemplateLoaderRequest) GetSchemasOk() (*[]EnumvelocityTemplateLoaderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVelocityTemplateLoaderRequest) SetSchemas(v []EnumvelocityTemplateLoaderSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddVelocityTemplateLoaderRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVelocityTemplateLoaderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVelocityTemplateLoaderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVelocityTemplateLoaderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddVelocityTemplateLoaderRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *AddVelocityTemplateLoaderRequest) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *AddVelocityTemplateLoaderRequest) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *AddVelocityTemplateLoaderRequest) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetMimeTypeMatcher

`func (o *AddVelocityTemplateLoaderRequest) GetMimeTypeMatcher() string`

GetMimeTypeMatcher returns the MimeTypeMatcher field if non-nil, zero value otherwise.

### GetMimeTypeMatcherOk

`func (o *AddVelocityTemplateLoaderRequest) GetMimeTypeMatcherOk() (*string, bool)`

GetMimeTypeMatcherOk returns a tuple with the MimeTypeMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypeMatcher

`func (o *AddVelocityTemplateLoaderRequest) SetMimeTypeMatcher(v string)`

SetMimeTypeMatcher sets MimeTypeMatcher field to given value.


### GetMimeType

`func (o *AddVelocityTemplateLoaderRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AddVelocityTemplateLoaderRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AddVelocityTemplateLoaderRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AddVelocityTemplateLoaderRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetTemplateSuffix

`func (o *AddVelocityTemplateLoaderRequest) GetTemplateSuffix() string`

GetTemplateSuffix returns the TemplateSuffix field if non-nil, zero value otherwise.

### GetTemplateSuffixOk

`func (o *AddVelocityTemplateLoaderRequest) GetTemplateSuffixOk() (*string, bool)`

GetTemplateSuffixOk returns a tuple with the TemplateSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSuffix

`func (o *AddVelocityTemplateLoaderRequest) SetTemplateSuffix(v string)`

SetTemplateSuffix sets TemplateSuffix field to given value.

### HasTemplateSuffix

`func (o *AddVelocityTemplateLoaderRequest) HasTemplateSuffix() bool`

HasTemplateSuffix returns a boolean if a field has been set.

### GetTemplateDirectory

`func (o *AddVelocityTemplateLoaderRequest) GetTemplateDirectory() string`

GetTemplateDirectory returns the TemplateDirectory field if non-nil, zero value otherwise.

### GetTemplateDirectoryOk

`func (o *AddVelocityTemplateLoaderRequest) GetTemplateDirectoryOk() (*string, bool)`

GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDirectory

`func (o *AddVelocityTemplateLoaderRequest) SetTemplateDirectory(v string)`

SetTemplateDirectory sets TemplateDirectory field to given value.

### HasTemplateDirectory

`func (o *AddVelocityTemplateLoaderRequest) HasTemplateDirectory() bool`

HasTemplateDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


