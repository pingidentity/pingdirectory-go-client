# DnMapShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumdnMapSchemaUrn**](EnumdnMapSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this DN Map | [optional] 
**FromDNPattern** | **string** | Specifies the DN pattern to match when determining whether this map applies to a specific source DN. If the provided bind DN matches this pattern, then the to-dn-pattern will be used to perform the mapping. If the provided bind DN does not match this pattern, then no mapping will be performed. | 
**ToDNPattern** | **string** | Specifies a pattern for constructing the DN value using fixed text, DN components matching wild-card values in from-dn-pattern, and attribute values from the source entry. | 

## Methods

### NewDnMapShared

`func NewDnMapShared(fromDNPattern string, toDNPattern string, ) *DnMapShared`

NewDnMapShared instantiates a new DnMapShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnMapSharedWithDefaults

`func NewDnMapSharedWithDefaults() *DnMapShared`

NewDnMapSharedWithDefaults instantiates a new DnMapShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DnMapShared) GetSchemas() []EnumdnMapSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DnMapShared) GetSchemasOk() (*[]EnumdnMapSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DnMapShared) SetSchemas(v []EnumdnMapSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DnMapShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DnMapShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DnMapShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DnMapShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DnMapShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromDNPattern

`func (o *DnMapShared) GetFromDNPattern() string`

GetFromDNPattern returns the FromDNPattern field if non-nil, zero value otherwise.

### GetFromDNPatternOk

`func (o *DnMapShared) GetFromDNPatternOk() (*string, bool)`

GetFromDNPatternOk returns a tuple with the FromDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDNPattern

`func (o *DnMapShared) SetFromDNPattern(v string)`

SetFromDNPattern sets FromDNPattern field to given value.


### GetToDNPattern

`func (o *DnMapShared) GetToDNPattern() string`

GetToDNPattern returns the ToDNPattern field if non-nil, zero value otherwise.

### GetToDNPatternOk

`func (o *DnMapShared) GetToDNPatternOk() (*string, bool)`

GetToDNPatternOk returns a tuple with the ToDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDNPattern

`func (o *DnMapShared) SetToDNPattern(v string)`

SetToDNPattern sets ToDNPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


