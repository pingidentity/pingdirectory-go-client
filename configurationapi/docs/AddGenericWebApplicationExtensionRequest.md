# AddGenericWebApplicationExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgenericWebApplicationExtensionSchemaUrn**](EnumgenericWebApplicationExtensionSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Web Application Extension | [optional] 
**BaseContextPath** | **string** | Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and at least one additional character and must represent a valid HTTP context path. | 
**WarFile** | Pointer to **string** | Specifies the path to a standard web application archive (WAR) file. | [optional] 
**DocumentRootDirectory** | Pointer to **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this Web Application Extension. The path must exist, and it must be a directory. | [optional] 
**DeploymentDescriptorFile** | Pointer to **string** | Specifies the path to the deployment descriptor file when used with document-root-directory. | [optional] 
**TemporaryDirectory** | Pointer to **string** | Specifies the path to the directory that may be used to store temporary files such as extracted WAR files and compiled JSP files. | [optional] 
**InitParameter** | Pointer to **[]string** | Specifies an initialization parameter to pass into the web application during startup. | [optional] 
**ExtensionName** | **string** | Name of the new Web Application Extension | 

## Methods

### NewAddGenericWebApplicationExtensionRequest

`func NewAddGenericWebApplicationExtensionRequest(schemas []EnumgenericWebApplicationExtensionSchemaUrn, baseContextPath string, extensionName string, ) *AddGenericWebApplicationExtensionRequest`

NewAddGenericWebApplicationExtensionRequest instantiates a new AddGenericWebApplicationExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGenericWebApplicationExtensionRequestWithDefaults

`func NewAddGenericWebApplicationExtensionRequestWithDefaults() *AddGenericWebApplicationExtensionRequest`

NewAddGenericWebApplicationExtensionRequestWithDefaults instantiates a new AddGenericWebApplicationExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddGenericWebApplicationExtensionRequest) GetSchemas() []EnumgenericWebApplicationExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGenericWebApplicationExtensionRequest) GetSchemasOk() (*[]EnumgenericWebApplicationExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGenericWebApplicationExtensionRequest) SetSchemas(v []EnumgenericWebApplicationExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddGenericWebApplicationExtensionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGenericWebApplicationExtensionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGenericWebApplicationExtensionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGenericWebApplicationExtensionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *AddGenericWebApplicationExtensionRequest) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AddGenericWebApplicationExtensionRequest) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AddGenericWebApplicationExtensionRequest) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetWarFile

`func (o *AddGenericWebApplicationExtensionRequest) GetWarFile() string`

GetWarFile returns the WarFile field if non-nil, zero value otherwise.

### GetWarFileOk

`func (o *AddGenericWebApplicationExtensionRequest) GetWarFileOk() (*string, bool)`

GetWarFileOk returns a tuple with the WarFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarFile

`func (o *AddGenericWebApplicationExtensionRequest) SetWarFile(v string)`

SetWarFile sets WarFile field to given value.

### HasWarFile

`func (o *AddGenericWebApplicationExtensionRequest) HasWarFile() bool`

HasWarFile returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *AddGenericWebApplicationExtensionRequest) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *AddGenericWebApplicationExtensionRequest) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *AddGenericWebApplicationExtensionRequest) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.

### HasDocumentRootDirectory

`func (o *AddGenericWebApplicationExtensionRequest) HasDocumentRootDirectory() bool`

HasDocumentRootDirectory returns a boolean if a field has been set.

### GetDeploymentDescriptorFile

`func (o *AddGenericWebApplicationExtensionRequest) GetDeploymentDescriptorFile() string`

GetDeploymentDescriptorFile returns the DeploymentDescriptorFile field if non-nil, zero value otherwise.

### GetDeploymentDescriptorFileOk

`func (o *AddGenericWebApplicationExtensionRequest) GetDeploymentDescriptorFileOk() (*string, bool)`

GetDeploymentDescriptorFileOk returns a tuple with the DeploymentDescriptorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentDescriptorFile

`func (o *AddGenericWebApplicationExtensionRequest) SetDeploymentDescriptorFile(v string)`

SetDeploymentDescriptorFile sets DeploymentDescriptorFile field to given value.

### HasDeploymentDescriptorFile

`func (o *AddGenericWebApplicationExtensionRequest) HasDeploymentDescriptorFile() bool`

HasDeploymentDescriptorFile returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *AddGenericWebApplicationExtensionRequest) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *AddGenericWebApplicationExtensionRequest) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *AddGenericWebApplicationExtensionRequest) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.

### HasTemporaryDirectory

`func (o *AddGenericWebApplicationExtensionRequest) HasTemporaryDirectory() bool`

HasTemporaryDirectory returns a boolean if a field has been set.

### GetInitParameter

`func (o *AddGenericWebApplicationExtensionRequest) GetInitParameter() []string`

GetInitParameter returns the InitParameter field if non-nil, zero value otherwise.

### GetInitParameterOk

`func (o *AddGenericWebApplicationExtensionRequest) GetInitParameterOk() (*[]string, bool)`

GetInitParameterOk returns a tuple with the InitParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitParameter

`func (o *AddGenericWebApplicationExtensionRequest) SetInitParameter(v []string)`

SetInitParameter sets InitParameter field to given value.

### HasInitParameter

`func (o *AddGenericWebApplicationExtensionRequest) HasInitParameter() bool`

HasInitParameter returns a boolean if a field has been set.

### GetExtensionName

`func (o *AddGenericWebApplicationExtensionRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *AddGenericWebApplicationExtensionRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *AddGenericWebApplicationExtensionRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


