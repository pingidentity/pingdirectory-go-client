# AddThirdPartyPassThroughAuthenticationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn**](EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Pass Through Authentication Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Pass Through Authentication Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Pass Through Authentication Handler | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**HandlerName** | **string** | Name of the new Pass Through Authentication Handler | 

## Methods

### NewAddThirdPartyPassThroughAuthenticationHandlerRequest

`func NewAddThirdPartyPassThroughAuthenticationHandlerRequest(schemas []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, extensionClass string, handlerName string, ) *AddThirdPartyPassThroughAuthenticationHandlerRequest`

NewAddThirdPartyPassThroughAuthenticationHandlerRequest instantiates a new AddThirdPartyPassThroughAuthenticationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyPassThroughAuthenticationHandlerRequestWithDefaults

`func NewAddThirdPartyPassThroughAuthenticationHandlerRequestWithDefaults() *AddThirdPartyPassThroughAuthenticationHandlerRequest`

NewAddThirdPartyPassThroughAuthenticationHandlerRequestWithDefaults instantiates a new AddThirdPartyPassThroughAuthenticationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetSchemas() []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetSchemasOk() (*[]EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetSchemas(v []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetHandlerName

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddThirdPartyPassThroughAuthenticationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


