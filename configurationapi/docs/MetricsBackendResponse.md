# MetricsBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummetricsBackendSchemaUrn**](EnummetricsBackendSchemaUrn.md) |  | 
**Id** | **string** | Name of the Backend | 
**BackendID** | **string** | Specifies a name to identify the associated backend. | 
**StorageDir** | **string** | Specifies the path to the directory that will be used to store queued samples. | 
**MetricsDir** | **string** | Specifies the path to the directory that contains metric definitions. | 
**SampleFlushInterval** | Pointer to **string** | Period when samples are flushed to disk. | [optional] 
**RetentionPolicy** | **[]string** | The retention policy to use for the Metrics Backend . | 
**Description** | Pointer to **string** | A description for this Backend | [optional] 
**Enabled** | **bool** | Indicates whether the backend is enabled in the server. | 
**WritabilityMode** | [**EnumbackendWritabilityModeProp**](EnumbackendWritabilityModeProp.md) |  | 
**ReturnUnavailableWhenDisabled** | Pointer to **bool** | Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled. | [optional] 
**NotificationManager** | Pointer to **string** | Specifies a notification manager for changes resulting from operations processed through this Backend | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewMetricsBackendResponse

`func NewMetricsBackendResponse(schemas []EnummetricsBackendSchemaUrn, id string, backendID string, storageDir string, metricsDir string, retentionPolicy []string, enabled bool, writabilityMode EnumbackendWritabilityModeProp, ) *MetricsBackendResponse`

NewMetricsBackendResponse instantiates a new MetricsBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsBackendResponseWithDefaults

`func NewMetricsBackendResponseWithDefaults() *MetricsBackendResponse`

NewMetricsBackendResponseWithDefaults instantiates a new MetricsBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MetricsBackendResponse) GetSchemas() []EnummetricsBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MetricsBackendResponse) GetSchemasOk() (*[]EnummetricsBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MetricsBackendResponse) SetSchemas(v []EnummetricsBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *MetricsBackendResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsBackendResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsBackendResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBackendID

`func (o *MetricsBackendResponse) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *MetricsBackendResponse) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *MetricsBackendResponse) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetStorageDir

`func (o *MetricsBackendResponse) GetStorageDir() string`

GetStorageDir returns the StorageDir field if non-nil, zero value otherwise.

### GetStorageDirOk

`func (o *MetricsBackendResponse) GetStorageDirOk() (*string, bool)`

GetStorageDirOk returns a tuple with the StorageDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDir

`func (o *MetricsBackendResponse) SetStorageDir(v string)`

SetStorageDir sets StorageDir field to given value.


### GetMetricsDir

`func (o *MetricsBackendResponse) GetMetricsDir() string`

GetMetricsDir returns the MetricsDir field if non-nil, zero value otherwise.

### GetMetricsDirOk

`func (o *MetricsBackendResponse) GetMetricsDirOk() (*string, bool)`

GetMetricsDirOk returns a tuple with the MetricsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsDir

`func (o *MetricsBackendResponse) SetMetricsDir(v string)`

SetMetricsDir sets MetricsDir field to given value.


### GetSampleFlushInterval

`func (o *MetricsBackendResponse) GetSampleFlushInterval() string`

GetSampleFlushInterval returns the SampleFlushInterval field if non-nil, zero value otherwise.

### GetSampleFlushIntervalOk

`func (o *MetricsBackendResponse) GetSampleFlushIntervalOk() (*string, bool)`

GetSampleFlushIntervalOk returns a tuple with the SampleFlushInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFlushInterval

`func (o *MetricsBackendResponse) SetSampleFlushInterval(v string)`

SetSampleFlushInterval sets SampleFlushInterval field to given value.

### HasSampleFlushInterval

`func (o *MetricsBackendResponse) HasSampleFlushInterval() bool`

HasSampleFlushInterval returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *MetricsBackendResponse) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *MetricsBackendResponse) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *MetricsBackendResponse) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetDescription

`func (o *MetricsBackendResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsBackendResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsBackendResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsBackendResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricsBackendResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsBackendResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsBackendResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWritabilityMode

`func (o *MetricsBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp`

GetWritabilityMode returns the WritabilityMode field if non-nil, zero value otherwise.

### GetWritabilityModeOk

`func (o *MetricsBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool)`

GetWritabilityModeOk returns a tuple with the WritabilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritabilityMode

`func (o *MetricsBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp)`

SetWritabilityMode sets WritabilityMode field to given value.


### GetReturnUnavailableWhenDisabled

`func (o *MetricsBackendResponse) GetReturnUnavailableWhenDisabled() bool`

GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field if non-nil, zero value otherwise.

### GetReturnUnavailableWhenDisabledOk

`func (o *MetricsBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool)`

GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUnavailableWhenDisabled

`func (o *MetricsBackendResponse) SetReturnUnavailableWhenDisabled(v bool)`

SetReturnUnavailableWhenDisabled sets ReturnUnavailableWhenDisabled field to given value.

### HasReturnUnavailableWhenDisabled

`func (o *MetricsBackendResponse) HasReturnUnavailableWhenDisabled() bool`

HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.

### GetNotificationManager

`func (o *MetricsBackendResponse) GetNotificationManager() string`

GetNotificationManager returns the NotificationManager field if non-nil, zero value otherwise.

### GetNotificationManagerOk

`func (o *MetricsBackendResponse) GetNotificationManagerOk() (*string, bool)`

GetNotificationManagerOk returns a tuple with the NotificationManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationManager

`func (o *MetricsBackendResponse) SetNotificationManager(v string)`

SetNotificationManager sets NotificationManager field to given value.

### HasNotificationManager

`func (o *MetricsBackendResponse) HasNotificationManager() bool`

HasNotificationManager returns a boolean if a field has been set.

### GetMeta

`func (o *MetricsBackendResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MetricsBackendResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MetricsBackendResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MetricsBackendResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MetricsBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MetricsBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MetricsBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MetricsBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


