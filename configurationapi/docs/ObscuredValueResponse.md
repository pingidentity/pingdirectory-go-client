# ObscuredValueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Obscured Value | 
**Schemas** | Pointer to [**[]EnumobscuredValueSchemaUrn**](EnumobscuredValueSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Obscured Value | [optional] 
**ObscuredValue** | **string** | The value to be stored in an obscured form. | 

## Methods

### NewObscuredValueResponse

`func NewObscuredValueResponse(id string, obscuredValue string, ) *ObscuredValueResponse`

NewObscuredValueResponse instantiates a new ObscuredValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObscuredValueResponseWithDefaults

`func NewObscuredValueResponseWithDefaults() *ObscuredValueResponse`

NewObscuredValueResponseWithDefaults instantiates a new ObscuredValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ObscuredValueResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ObscuredValueResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ObscuredValueResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ObscuredValueResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ObscuredValueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ObscuredValueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ObscuredValueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ObscuredValueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ObscuredValueResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObscuredValueResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObscuredValueResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ObscuredValueResponse) GetSchemas() []EnumobscuredValueSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ObscuredValueResponse) GetSchemasOk() (*[]EnumobscuredValueSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ObscuredValueResponse) SetSchemas(v []EnumobscuredValueSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ObscuredValueResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ObscuredValueResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObscuredValueResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObscuredValueResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObscuredValueResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObscuredValue

`func (o *ObscuredValueResponse) GetObscuredValue() string`

GetObscuredValue returns the ObscuredValue field if non-nil, zero value otherwise.

### GetObscuredValueOk

`func (o *ObscuredValueResponse) GetObscuredValueOk() (*string, bool)`

GetObscuredValueOk returns a tuple with the ObscuredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredValue

`func (o *ObscuredValueResponse) SetObscuredValue(v string)`

SetObscuredValue sets ObscuredValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


