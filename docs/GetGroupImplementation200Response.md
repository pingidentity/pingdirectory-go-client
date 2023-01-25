# GetGroupImplementation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumdynamicGroupImplementationSchemaUrn**](EnumdynamicGroupImplementationSchemaUrn.md) |  | 
**Id** | **string** | Name of the Group Implementation | 
**Description** | Pointer to **string** | A description for this Group Implementation | [optional] 
**Enabled** | **bool** | Indicates whether the Group Implementation is enabled. | 

## Methods

### NewGetGroupImplementation200Response

`func NewGetGroupImplementation200Response(schemas []EnumdynamicGroupImplementationSchemaUrn, id string, enabled bool, ) *GetGroupImplementation200Response`

NewGetGroupImplementation200Response instantiates a new GetGroupImplementation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupImplementation200ResponseWithDefaults

`func NewGetGroupImplementation200ResponseWithDefaults() *GetGroupImplementation200Response`

NewGetGroupImplementation200ResponseWithDefaults instantiates a new GetGroupImplementation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetGroupImplementation200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGroupImplementation200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGroupImplementation200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetGroupImplementation200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetGroupImplementation200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetGroupImplementation200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetGroupImplementation200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetGroupImplementation200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GetGroupImplementation200Response) GetSchemas() []EnumdynamicGroupImplementationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetGroupImplementation200Response) GetSchemasOk() (*[]EnumdynamicGroupImplementationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetGroupImplementation200Response) SetSchemas(v []EnumdynamicGroupImplementationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetGroupImplementation200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGroupImplementation200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGroupImplementation200Response) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *GetGroupImplementation200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGroupImplementation200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGroupImplementation200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGroupImplementation200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetGroupImplementation200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetGroupImplementation200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetGroupImplementation200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


