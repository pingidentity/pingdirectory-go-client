# ReplaceCertificateExtendedOperationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumreplaceCertificateExtendedOperationHandlerSchemaUrn**](EnumreplaceCertificateExtendedOperationHandlerSchemaUrn.md) |  | 
**AllowRemotelyProvidedCertificates** | Pointer to **bool** | Indicates whether clients should be allowed to directly provide a new listener or inter-server certificate chain in the extended request. | [optional] 
**AllowedOperation** | Pointer to [**[]EnumextendedOperationHandlerAllowedOperationProp**](EnumextendedOperationHandlerAllowedOperationProp.md) |  | [optional] 
**ConnectionCriteria** | Pointer to **string** | A set of criteria that client connections must satisfy before they will be allowed to request the associated extended operations. | [optional] 
**RequestCriteria** | Pointer to **string** | A set of criteria that the extended requests must satisfy before they will be processed by the server. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewReplaceCertificateExtendedOperationHandlerShared

`func NewReplaceCertificateExtendedOperationHandlerShared(schemas []EnumreplaceCertificateExtendedOperationHandlerSchemaUrn, enabled bool, ) *ReplaceCertificateExtendedOperationHandlerShared`

NewReplaceCertificateExtendedOperationHandlerShared instantiates a new ReplaceCertificateExtendedOperationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceCertificateExtendedOperationHandlerSharedWithDefaults

`func NewReplaceCertificateExtendedOperationHandlerSharedWithDefaults() *ReplaceCertificateExtendedOperationHandlerShared`

NewReplaceCertificateExtendedOperationHandlerSharedWithDefaults instantiates a new ReplaceCertificateExtendedOperationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetSchemas() []EnumreplaceCertificateExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetSchemasOk() (*[]EnumreplaceCertificateExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetSchemas(v []EnumreplaceCertificateExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowRemotelyProvidedCertificates

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetAllowRemotelyProvidedCertificates() bool`

GetAllowRemotelyProvidedCertificates returns the AllowRemotelyProvidedCertificates field if non-nil, zero value otherwise.

### GetAllowRemotelyProvidedCertificatesOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetAllowRemotelyProvidedCertificatesOk() (*bool, bool)`

GetAllowRemotelyProvidedCertificatesOk returns a tuple with the AllowRemotelyProvidedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemotelyProvidedCertificates

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetAllowRemotelyProvidedCertificates(v bool)`

SetAllowRemotelyProvidedCertificates sets AllowRemotelyProvidedCertificates field to given value.

### HasAllowRemotelyProvidedCertificates

`func (o *ReplaceCertificateExtendedOperationHandlerShared) HasAllowRemotelyProvidedCertificates() bool`

HasAllowRemotelyProvidedCertificates returns a boolean if a field has been set.

### GetAllowedOperation

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetAllowedOperation() []EnumextendedOperationHandlerAllowedOperationProp`

GetAllowedOperation returns the AllowedOperation field if non-nil, zero value otherwise.

### GetAllowedOperationOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetAllowedOperationOk() (*[]EnumextendedOperationHandlerAllowedOperationProp, bool)`

GetAllowedOperationOk returns a tuple with the AllowedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperation

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetAllowedOperation(v []EnumextendedOperationHandlerAllowedOperationProp)`

SetAllowedOperation sets AllowedOperation field to given value.

### HasAllowedOperation

`func (o *ReplaceCertificateExtendedOperationHandlerShared) HasAllowedOperation() bool`

HasAllowedOperation returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ReplaceCertificateExtendedOperationHandlerShared) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ReplaceCertificateExtendedOperationHandlerShared) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplaceCertificateExtendedOperationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplaceCertificateExtendedOperationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplaceCertificateExtendedOperationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


