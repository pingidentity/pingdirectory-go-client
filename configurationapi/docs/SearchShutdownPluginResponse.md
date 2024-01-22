# SearchShutdownPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsearchShutdownPluginSchemaUrn**](EnumsearchShutdownPluginSchemaUrn.md) |  | 
**BaseDN** | Pointer to **string** | The base DN to use for the search. | [optional] 
**Scope** | [**EnumpluginScopeProp**](EnumpluginScopeProp.md) |  | 
**Filter** | **string** | The filter to use for the search. | 
**IncludeAttribute** | Pointer to **[]string** | The name of an attribute that should be included in the results. This may include any token which is allowed as a requested attribute in search requests, including the name of an attribute, an asterisk (to indicate all user attributes), a plus sign (to indicate all operational attributes), an object class name preceded with an at symbol (to indicate all attributes associated with that object class), an attribute name preceded by a caret (to indicate that attribute should be excluded), or an object class name preceded by a caret and an at symbol (to indicate that all attributes associated with that object class should be excluded). | [optional] 
**OutputFile** | **string** | The path of an LDIF file that should be created with the results of the search. | 
**PreviousFileExtension** | Pointer to **string** | An extension that should be appended to the name of an existing output file rather than deleting it. If a file already exists with the full previous file name, then it will be deleted before the current file is renamed to become the previous file. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewSearchShutdownPluginResponse

`func NewSearchShutdownPluginResponse(schemas []EnumsearchShutdownPluginSchemaUrn, scope EnumpluginScopeProp, filter string, outputFile string, enabled bool, id string, ) *SearchShutdownPluginResponse`

NewSearchShutdownPluginResponse instantiates a new SearchShutdownPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchShutdownPluginResponseWithDefaults

`func NewSearchShutdownPluginResponseWithDefaults() *SearchShutdownPluginResponse`

NewSearchShutdownPluginResponseWithDefaults instantiates a new SearchShutdownPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SearchShutdownPluginResponse) GetSchemas() []EnumsearchShutdownPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SearchShutdownPluginResponse) GetSchemasOk() (*[]EnumsearchShutdownPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SearchShutdownPluginResponse) SetSchemas(v []EnumsearchShutdownPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseDN

`func (o *SearchShutdownPluginResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *SearchShutdownPluginResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *SearchShutdownPluginResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *SearchShutdownPluginResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetScope

`func (o *SearchShutdownPluginResponse) GetScope() EnumpluginScopeProp`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SearchShutdownPluginResponse) GetScopeOk() (*EnumpluginScopeProp, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SearchShutdownPluginResponse) SetScope(v EnumpluginScopeProp)`

SetScope sets Scope field to given value.


### GetFilter

`func (o *SearchShutdownPluginResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchShutdownPluginResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchShutdownPluginResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetIncludeAttribute

`func (o *SearchShutdownPluginResponse) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *SearchShutdownPluginResponse) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *SearchShutdownPluginResponse) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *SearchShutdownPluginResponse) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetOutputFile

`func (o *SearchShutdownPluginResponse) GetOutputFile() string`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *SearchShutdownPluginResponse) GetOutputFileOk() (*string, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *SearchShutdownPluginResponse) SetOutputFile(v string)`

SetOutputFile sets OutputFile field to given value.


### GetPreviousFileExtension

`func (o *SearchShutdownPluginResponse) GetPreviousFileExtension() string`

GetPreviousFileExtension returns the PreviousFileExtension field if non-nil, zero value otherwise.

### GetPreviousFileExtensionOk

`func (o *SearchShutdownPluginResponse) GetPreviousFileExtensionOk() (*string, bool)`

GetPreviousFileExtensionOk returns a tuple with the PreviousFileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFileExtension

`func (o *SearchShutdownPluginResponse) SetPreviousFileExtension(v string)`

SetPreviousFileExtension sets PreviousFileExtension field to given value.

### HasPreviousFileExtension

`func (o *SearchShutdownPluginResponse) HasPreviousFileExtension() bool`

HasPreviousFileExtension returns a boolean if a field has been set.

### GetDescription

`func (o *SearchShutdownPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchShutdownPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchShutdownPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchShutdownPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SearchShutdownPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SearchShutdownPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SearchShutdownPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *SearchShutdownPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SearchShutdownPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SearchShutdownPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SearchShutdownPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SearchShutdownPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SearchShutdownPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SearchShutdownPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SearchShutdownPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *SearchShutdownPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchShutdownPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchShutdownPluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


