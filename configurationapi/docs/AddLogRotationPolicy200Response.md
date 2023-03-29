# AddLogRotationPolicy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Rotation Policy | 
**Schemas** | [**[]EnumsizeLimitLogRotationPolicySchemaUrn**](EnumsizeLimitLogRotationPolicySchemaUrn.md) |  | 
**RotationInterval** | **string** | Specifies the time interval between rotations. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**TimeOfDay** | **[]string** | Specifies the time of day at which log rotation should occur. | 
**FileSizeLimit** | **string** | Specifies the maximum size that a log file can reach before it is rotated. | 

## Methods

### NewAddLogRotationPolicy200Response

`func NewAddLogRotationPolicy200Response(id string, schemas []EnumsizeLimitLogRotationPolicySchemaUrn, rotationInterval string, timeOfDay []string, fileSizeLimit string, ) *AddLogRotationPolicy200Response`

NewAddLogRotationPolicy200Response instantiates a new AddLogRotationPolicy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogRotationPolicy200ResponseWithDefaults

`func NewAddLogRotationPolicy200ResponseWithDefaults() *AddLogRotationPolicy200Response`

NewAddLogRotationPolicy200ResponseWithDefaults instantiates a new AddLogRotationPolicy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLogRotationPolicy200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLogRotationPolicy200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLogRotationPolicy200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddLogRotationPolicy200Response) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogRotationPolicy200Response) GetSchemasOk() (*[]EnumsizeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogRotationPolicy200Response) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRotationInterval

`func (o *AddLogRotationPolicy200Response) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *AddLogRotationPolicy200Response) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *AddLogRotationPolicy200Response) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.


### GetDescription

`func (o *AddLogRotationPolicy200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogRotationPolicy200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogRotationPolicy200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogRotationPolicy200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddLogRotationPolicy200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddLogRotationPolicy200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddLogRotationPolicy200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddLogRotationPolicy200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogRotationPolicy200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddLogRotationPolicy200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogRotationPolicy200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogRotationPolicy200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *AddLogRotationPolicy200Response) GetTimeOfDay() []string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *AddLogRotationPolicy200Response) GetTimeOfDayOk() (*[]string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *AddLogRotationPolicy200Response) SetTimeOfDay(v []string)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetFileSizeLimit

`func (o *AddLogRotationPolicy200Response) GetFileSizeLimit() string`

GetFileSizeLimit returns the FileSizeLimit field if non-nil, zero value otherwise.

### GetFileSizeLimitOk

`func (o *AddLogRotationPolicy200Response) GetFileSizeLimitOk() (*string, bool)`

GetFileSizeLimitOk returns a tuple with the FileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLimit

`func (o *AddLogRotationPolicy200Response) SetFileSizeLimit(v string)`

SetFileSizeLimit sets FileSizeLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


