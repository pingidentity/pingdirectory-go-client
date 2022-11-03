# AddLocation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the new location | 
**Description** | Pointer to **string** | Description of the new location | [optional] 

## Methods

### NewAddLocation200Response

`func NewAddLocation200Response(id string, ) *AddLocation200Response`

NewAddLocation200Response instantiates a new AddLocation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLocation200ResponseWithDefaults

`func NewAddLocation200ResponseWithDefaults() *AddLocation200Response`

NewAddLocation200ResponseWithDefaults instantiates a new AddLocation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLocation200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLocation200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLocation200Response) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *AddLocation200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLocation200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLocation200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLocation200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


