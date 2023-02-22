# DnMapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the DN Map | 
**Schemas** | Pointer to [**[]EnumdnMapSchemaUrn**](EnumdnMapSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this DN Map | [optional] 
**FromDNPattern** | **string** | Specifies the DN pattern to match when determining whether this map applies to a specific source DN. If the provided bind DN matches this pattern, then the to-dn-pattern will be used to perform the mapping. If the provided bind DN does not match this pattern, then no mapping will be performed. | 
**ToDNPattern** | **string** | Specifies a pattern for constructing the DN value using fixed text, DN components matching wild-card values in from-dn-pattern, and attribute values from the source entry. | 

## Methods

### NewDnMapResponse

`func NewDnMapResponse(id string, fromDNPattern string, toDNPattern string, ) *DnMapResponse`

NewDnMapResponse instantiates a new DnMapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnMapResponseWithDefaults

`func NewDnMapResponseWithDefaults() *DnMapResponse`

NewDnMapResponseWithDefaults instantiates a new DnMapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DnMapResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DnMapResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DnMapResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DnMapResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DnMapResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DnMapResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DnMapResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DnMapResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DnMapResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnMapResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnMapResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DnMapResponse) GetSchemas() []EnumdnMapSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DnMapResponse) GetSchemasOk() (*[]EnumdnMapSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DnMapResponse) SetSchemas(v []EnumdnMapSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DnMapResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DnMapResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DnMapResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DnMapResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DnMapResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromDNPattern

`func (o *DnMapResponse) GetFromDNPattern() string`

GetFromDNPattern returns the FromDNPattern field if non-nil, zero value otherwise.

### GetFromDNPatternOk

`func (o *DnMapResponse) GetFromDNPatternOk() (*string, bool)`

GetFromDNPatternOk returns a tuple with the FromDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDNPattern

`func (o *DnMapResponse) SetFromDNPattern(v string)`

SetFromDNPattern sets FromDNPattern field to given value.


### GetToDNPattern

`func (o *DnMapResponse) GetToDNPattern() string`

GetToDNPattern returns the ToDNPattern field if non-nil, zero value otherwise.

### GetToDNPatternOk

`func (o *DnMapResponse) GetToDNPatternOk() (*string, bool)`

GetToDNPatternOk returns a tuple with the ToDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDNPattern

`func (o *DnMapResponse) SetToDNPattern(v string)`

SetToDNPattern sets ToDNPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


