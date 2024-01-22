# AddDnMapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumdnMapSchemaUrn**](EnumdnMapSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this DN Map | [optional] 
**FromDNPattern** | **string** | Specifies the DN pattern to match when determining whether this map applies to a specific source DN. If the provided bind DN matches this pattern, then the to-dn-pattern will be used to perform the mapping. If the provided bind DN does not match this pattern, then no mapping will be performed. | 
**ToDNPattern** | **string** | Specifies a pattern for constructing the DN value using fixed text, DN components matching wild-card values in from-dn-pattern, and attribute values from the source entry. | 
**MapName** | **string** | Name of the new DN Map | 

## Methods

### NewAddDnMapRequest

`func NewAddDnMapRequest(fromDNPattern string, toDNPattern string, mapName string, ) *AddDnMapRequest`

NewAddDnMapRequest instantiates a new AddDnMapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDnMapRequestWithDefaults

`func NewAddDnMapRequestWithDefaults() *AddDnMapRequest`

NewAddDnMapRequestWithDefaults instantiates a new AddDnMapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDnMapRequest) GetSchemas() []EnumdnMapSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDnMapRequest) GetSchemasOk() (*[]EnumdnMapSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDnMapRequest) SetSchemas(v []EnumdnMapSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddDnMapRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddDnMapRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDnMapRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDnMapRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDnMapRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromDNPattern

`func (o *AddDnMapRequest) GetFromDNPattern() string`

GetFromDNPattern returns the FromDNPattern field if non-nil, zero value otherwise.

### GetFromDNPatternOk

`func (o *AddDnMapRequest) GetFromDNPatternOk() (*string, bool)`

GetFromDNPatternOk returns a tuple with the FromDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDNPattern

`func (o *AddDnMapRequest) SetFromDNPattern(v string)`

SetFromDNPattern sets FromDNPattern field to given value.


### GetToDNPattern

`func (o *AddDnMapRequest) GetToDNPattern() string`

GetToDNPattern returns the ToDNPattern field if non-nil, zero value otherwise.

### GetToDNPatternOk

`func (o *AddDnMapRequest) GetToDNPatternOk() (*string, bool)`

GetToDNPatternOk returns a tuple with the ToDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDNPattern

`func (o *AddDnMapRequest) SetToDNPattern(v string)`

SetToDNPattern sets ToDNPattern field to given value.


### GetMapName

`func (o *AddDnMapRequest) GetMapName() string`

GetMapName returns the MapName field if non-nil, zero value otherwise.

### GetMapNameOk

`func (o *AddDnMapRequest) GetMapNameOk() (*string, bool)`

GetMapNameOk returns a tuple with the MapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapName

`func (o *AddDnMapRequest) SetMapName(v string)`

SetMapName sets MapName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


