# GenericWebApplicationExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Web Application Extension | 
**Schemas** | [**[]EnumgenericWebApplicationExtensionSchemaUrn**](EnumgenericWebApplicationExtensionSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Web Application Extension | [optional] 
**BaseContextPath** | **string** | Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and at least one additional character and must represent a valid HTTP context path. | 
**WarFile** | Pointer to **string** | Specifies the path to a standard web application archive (WAR) file. | [optional] 
**DocumentRootDirectory** | Pointer to **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this Web Application Extension. The path must exist, and it must be a directory. | [optional] 
**DeploymentDescriptorFile** | Pointer to **string** | Specifies the path to the deployment descriptor file when used with document-root-directory. | [optional] 
**TemporaryDirectory** | Pointer to **string** | Specifies the path to the directory that may be used to store temporary files such as extracted WAR files and compiled JSP files. | [optional] 
**InitParameter** | Pointer to **[]string** | Specifies an initialization parameter to pass into the web application during startup. | [optional] 

## Methods

### NewGenericWebApplicationExtensionResponse

`func NewGenericWebApplicationExtensionResponse(id string, schemas []EnumgenericWebApplicationExtensionSchemaUrn, baseContextPath string, ) *GenericWebApplicationExtensionResponse`

NewGenericWebApplicationExtensionResponse instantiates a new GenericWebApplicationExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericWebApplicationExtensionResponseWithDefaults

`func NewGenericWebApplicationExtensionResponseWithDefaults() *GenericWebApplicationExtensionResponse`

NewGenericWebApplicationExtensionResponseWithDefaults instantiates a new GenericWebApplicationExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GenericWebApplicationExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenericWebApplicationExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenericWebApplicationExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenericWebApplicationExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GenericWebApplicationExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GenericWebApplicationExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GenericWebApplicationExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GenericWebApplicationExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *GenericWebApplicationExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericWebApplicationExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericWebApplicationExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GenericWebApplicationExtensionResponse) GetSchemas() []EnumgenericWebApplicationExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GenericWebApplicationExtensionResponse) GetSchemasOk() (*[]EnumgenericWebApplicationExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GenericWebApplicationExtensionResponse) SetSchemas(v []EnumgenericWebApplicationExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *GenericWebApplicationExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericWebApplicationExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericWebApplicationExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericWebApplicationExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *GenericWebApplicationExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *GenericWebApplicationExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *GenericWebApplicationExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetWarFile

`func (o *GenericWebApplicationExtensionResponse) GetWarFile() string`

GetWarFile returns the WarFile field if non-nil, zero value otherwise.

### GetWarFileOk

`func (o *GenericWebApplicationExtensionResponse) GetWarFileOk() (*string, bool)`

GetWarFileOk returns a tuple with the WarFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarFile

`func (o *GenericWebApplicationExtensionResponse) SetWarFile(v string)`

SetWarFile sets WarFile field to given value.

### HasWarFile

`func (o *GenericWebApplicationExtensionResponse) HasWarFile() bool`

HasWarFile returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *GenericWebApplicationExtensionResponse) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *GenericWebApplicationExtensionResponse) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *GenericWebApplicationExtensionResponse) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.

### HasDocumentRootDirectory

`func (o *GenericWebApplicationExtensionResponse) HasDocumentRootDirectory() bool`

HasDocumentRootDirectory returns a boolean if a field has been set.

### GetDeploymentDescriptorFile

`func (o *GenericWebApplicationExtensionResponse) GetDeploymentDescriptorFile() string`

GetDeploymentDescriptorFile returns the DeploymentDescriptorFile field if non-nil, zero value otherwise.

### GetDeploymentDescriptorFileOk

`func (o *GenericWebApplicationExtensionResponse) GetDeploymentDescriptorFileOk() (*string, bool)`

GetDeploymentDescriptorFileOk returns a tuple with the DeploymentDescriptorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentDescriptorFile

`func (o *GenericWebApplicationExtensionResponse) SetDeploymentDescriptorFile(v string)`

SetDeploymentDescriptorFile sets DeploymentDescriptorFile field to given value.

### HasDeploymentDescriptorFile

`func (o *GenericWebApplicationExtensionResponse) HasDeploymentDescriptorFile() bool`

HasDeploymentDescriptorFile returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *GenericWebApplicationExtensionResponse) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *GenericWebApplicationExtensionResponse) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *GenericWebApplicationExtensionResponse) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.

### HasTemporaryDirectory

`func (o *GenericWebApplicationExtensionResponse) HasTemporaryDirectory() bool`

HasTemporaryDirectory returns a boolean if a field has been set.

### GetInitParameter

`func (o *GenericWebApplicationExtensionResponse) GetInitParameter() []string`

GetInitParameter returns the InitParameter field if non-nil, zero value otherwise.

### GetInitParameterOk

`func (o *GenericWebApplicationExtensionResponse) GetInitParameterOk() (*[]string, bool)`

GetInitParameterOk returns a tuple with the InitParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitParameter

`func (o *GenericWebApplicationExtensionResponse) SetInitParameter(v []string)`

SetInitParameter sets InitParameter field to given value.

### HasInitParameter

`func (o *GenericWebApplicationExtensionResponse) HasInitParameter() bool`

HasInitParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


