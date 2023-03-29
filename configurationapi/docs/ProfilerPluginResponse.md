# ProfilerPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumprofilerPluginSchemaUrn**](EnumprofilerPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin Root | 
**ProfileSampleInterval** | **string** | Specifies the sample interval in milliseconds to be used when capturing profiling information in the server. | 
**ProfileDirectory** | **string** | Specifies the path to the directory where profile information is to be written. This path may be either an absolute path or a path that is relative to the root of the Directory Server instance. | 
**EnableProfilingOnStartup** | **bool** | Indicates whether the profiler plug-in is to start collecting data automatically when the Directory Server is started. | 
**ProfileAction** | Pointer to [**EnumpluginProfileActionProp**](EnumpluginProfileActionProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewProfilerPluginResponse

`func NewProfilerPluginResponse(schemas []EnumprofilerPluginSchemaUrn, id string, profileSampleInterval string, profileDirectory string, enableProfilingOnStartup bool, enabled bool, ) *ProfilerPluginResponse`

NewProfilerPluginResponse instantiates a new ProfilerPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilerPluginResponseWithDefaults

`func NewProfilerPluginResponseWithDefaults() *ProfilerPluginResponse`

NewProfilerPluginResponseWithDefaults instantiates a new ProfilerPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ProfilerPluginResponse) GetSchemas() []EnumprofilerPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ProfilerPluginResponse) GetSchemasOk() (*[]EnumprofilerPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ProfilerPluginResponse) SetSchemas(v []EnumprofilerPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ProfilerPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfilerPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfilerPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetProfileSampleInterval

`func (o *ProfilerPluginResponse) GetProfileSampleInterval() string`

GetProfileSampleInterval returns the ProfileSampleInterval field if non-nil, zero value otherwise.

### GetProfileSampleIntervalOk

`func (o *ProfilerPluginResponse) GetProfileSampleIntervalOk() (*string, bool)`

GetProfileSampleIntervalOk returns a tuple with the ProfileSampleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSampleInterval

`func (o *ProfilerPluginResponse) SetProfileSampleInterval(v string)`

SetProfileSampleInterval sets ProfileSampleInterval field to given value.


### GetProfileDirectory

`func (o *ProfilerPluginResponse) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *ProfilerPluginResponse) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *ProfilerPluginResponse) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetEnableProfilingOnStartup

`func (o *ProfilerPluginResponse) GetEnableProfilingOnStartup() bool`

GetEnableProfilingOnStartup returns the EnableProfilingOnStartup field if non-nil, zero value otherwise.

### GetEnableProfilingOnStartupOk

`func (o *ProfilerPluginResponse) GetEnableProfilingOnStartupOk() (*bool, bool)`

GetEnableProfilingOnStartupOk returns a tuple with the EnableProfilingOnStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfilingOnStartup

`func (o *ProfilerPluginResponse) SetEnableProfilingOnStartup(v bool)`

SetEnableProfilingOnStartup sets EnableProfilingOnStartup field to given value.


### GetProfileAction

`func (o *ProfilerPluginResponse) GetProfileAction() EnumpluginProfileActionProp`

GetProfileAction returns the ProfileAction field if non-nil, zero value otherwise.

### GetProfileActionOk

`func (o *ProfilerPluginResponse) GetProfileActionOk() (*EnumpluginProfileActionProp, bool)`

GetProfileActionOk returns a tuple with the ProfileAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAction

`func (o *ProfilerPluginResponse) SetProfileAction(v EnumpluginProfileActionProp)`

SetProfileAction sets ProfileAction field to given value.

### HasProfileAction

`func (o *ProfilerPluginResponse) HasProfileAction() bool`

HasProfileAction returns a boolean if a field has been set.

### GetDescription

`func (o *ProfilerPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfilerPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfilerPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProfilerPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ProfilerPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProfilerPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProfilerPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ProfilerPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProfilerPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProfilerPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProfilerPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ProfilerPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ProfilerPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ProfilerPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ProfilerPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


