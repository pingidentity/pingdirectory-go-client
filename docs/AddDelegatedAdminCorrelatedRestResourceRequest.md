# AddDelegatedAdminCorrelatedRestResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | Name of the new Delegated Admin Correlated REST Resource | 
**Schemas** | Pointer to [**[]EnumdelegatedAdminCorrelatedRestResourceSchemaUrn**](EnumdelegatedAdminCorrelatedRestResourceSchemaUrn.md) |  | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Correlated REST Resource. | 
**CorrelatedRESTResource** | **string** | The REST Resource Type that will be linked to this REST Resource Type. | 
**PrimaryRESTResourceCorrelationAttribute** | **string** | The LDAP attribute from the parent REST Resource Type whose value will be used to match objects in the Delegated Admin Correlated REST Resource. This attribute must be writeable when use-secondary-value-for-linking is enabled. | 
**SecondaryRESTResourceCorrelationAttribute** | **string** | The LDAP attribute from the Delegated Admin Correlated REST Resource whose value will be matched with the primary-rest-resource-correlation-attribute. This attribute must be writeable when use-secondary-value-for-linking is disabled. | 
**UseSecondaryValueForLinking** | Pointer to **bool** | Indicates whether links should be created using the secondary correlation attribute value. | [optional] 

## Methods

### NewAddDelegatedAdminCorrelatedRestResourceRequest

`func NewAddDelegatedAdminCorrelatedRestResourceRequest(resourceName string, displayName string, correlatedRESTResource string, primaryRESTResourceCorrelationAttribute string, secondaryRESTResourceCorrelationAttribute string, ) *AddDelegatedAdminCorrelatedRestResourceRequest`

NewAddDelegatedAdminCorrelatedRestResourceRequest instantiates a new AddDelegatedAdminCorrelatedRestResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelegatedAdminCorrelatedRestResourceRequestWithDefaults

`func NewAddDelegatedAdminCorrelatedRestResourceRequestWithDefaults() *AddDelegatedAdminCorrelatedRestResourceRequest`

NewAddDelegatedAdminCorrelatedRestResourceRequestWithDefaults instantiates a new AddDelegatedAdminCorrelatedRestResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetSchemas

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetSchemas() []EnumdelegatedAdminCorrelatedRestResourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetSchemasOk() (*[]EnumdelegatedAdminCorrelatedRestResourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetSchemas(v []EnumdelegatedAdminCorrelatedRestResourceSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCorrelatedRESTResource

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetCorrelatedRESTResource() string`

GetCorrelatedRESTResource returns the CorrelatedRESTResource field if non-nil, zero value otherwise.

### GetCorrelatedRESTResourceOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetCorrelatedRESTResourceOk() (*string, bool)`

GetCorrelatedRESTResourceOk returns a tuple with the CorrelatedRESTResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedRESTResource

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetCorrelatedRESTResource(v string)`

SetCorrelatedRESTResource sets CorrelatedRESTResource field to given value.


### GetPrimaryRESTResourceCorrelationAttribute

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetPrimaryRESTResourceCorrelationAttribute() string`

GetPrimaryRESTResourceCorrelationAttribute returns the PrimaryRESTResourceCorrelationAttribute field if non-nil, zero value otherwise.

### GetPrimaryRESTResourceCorrelationAttributeOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetPrimaryRESTResourceCorrelationAttributeOk() (*string, bool)`

GetPrimaryRESTResourceCorrelationAttributeOk returns a tuple with the PrimaryRESTResourceCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRESTResourceCorrelationAttribute

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetPrimaryRESTResourceCorrelationAttribute(v string)`

SetPrimaryRESTResourceCorrelationAttribute sets PrimaryRESTResourceCorrelationAttribute field to given value.


### GetSecondaryRESTResourceCorrelationAttribute

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetSecondaryRESTResourceCorrelationAttribute() string`

GetSecondaryRESTResourceCorrelationAttribute returns the SecondaryRESTResourceCorrelationAttribute field if non-nil, zero value otherwise.

### GetSecondaryRESTResourceCorrelationAttributeOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetSecondaryRESTResourceCorrelationAttributeOk() (*string, bool)`

GetSecondaryRESTResourceCorrelationAttributeOk returns a tuple with the SecondaryRESTResourceCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRESTResourceCorrelationAttribute

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetSecondaryRESTResourceCorrelationAttribute(v string)`

SetSecondaryRESTResourceCorrelationAttribute sets SecondaryRESTResourceCorrelationAttribute field to given value.


### GetUseSecondaryValueForLinking

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetUseSecondaryValueForLinking() bool`

GetUseSecondaryValueForLinking returns the UseSecondaryValueForLinking field if non-nil, zero value otherwise.

### GetUseSecondaryValueForLinkingOk

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) GetUseSecondaryValueForLinkingOk() (*bool, bool)`

GetUseSecondaryValueForLinkingOk returns a tuple with the UseSecondaryValueForLinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSecondaryValueForLinking

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) SetUseSecondaryValueForLinking(v bool)`

SetUseSecondaryValueForLinking sets UseSecondaryValueForLinking field to given value.

### HasUseSecondaryValueForLinking

`func (o *AddDelegatedAdminCorrelatedRestResourceRequest) HasUseSecondaryValueForLinking() bool`

HasUseSecondaryValueForLinking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


