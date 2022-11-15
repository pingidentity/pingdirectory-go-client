# ThirdPartyAccessTokenValidatorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyAccessTokenValidatorSchemaUrn**](EnumthirdPartyAccessTokenValidatorSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Access Token Validator. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 
**EvaluationOrderIndex** | **int32** | When multiple Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Access Token Validators defined within Directory Server but not necessarily contiguous. Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 

## Methods

### NewThirdPartyAccessTokenValidatorShared

`func NewThirdPartyAccessTokenValidatorShared(schemas []EnumthirdPartyAccessTokenValidatorSchemaUrn, extensionClass string, enabled bool, evaluationOrderIndex int32, ) *ThirdPartyAccessTokenValidatorShared`

NewThirdPartyAccessTokenValidatorShared instantiates a new ThirdPartyAccessTokenValidatorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyAccessTokenValidatorSharedWithDefaults

`func NewThirdPartyAccessTokenValidatorSharedWithDefaults() *ThirdPartyAccessTokenValidatorShared`

NewThirdPartyAccessTokenValidatorSharedWithDefaults instantiates a new ThirdPartyAccessTokenValidatorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ThirdPartyAccessTokenValidatorShared) GetSchemas() []EnumthirdPartyAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetSchemasOk() (*[]EnumthirdPartyAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyAccessTokenValidatorShared) SetSchemas(v []EnumthirdPartyAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyAccessTokenValidatorShared) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyAccessTokenValidatorShared) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyAccessTokenValidatorShared) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyAccessTokenValidatorShared) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyAccessTokenValidatorShared) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *ThirdPartyAccessTokenValidatorShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *ThirdPartyAccessTokenValidatorShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *ThirdPartyAccessTokenValidatorShared) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *ThirdPartyAccessTokenValidatorShared) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *ThirdPartyAccessTokenValidatorShared) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *ThirdPartyAccessTokenValidatorShared) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *ThirdPartyAccessTokenValidatorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyAccessTokenValidatorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyAccessTokenValidatorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyAccessTokenValidatorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyAccessTokenValidatorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEvaluationOrderIndex

`func (o *ThirdPartyAccessTokenValidatorShared) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *ThirdPartyAccessTokenValidatorShared) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *ThirdPartyAccessTokenValidatorShared) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


