# MockAccessTokenValidatorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummockAccessTokenValidatorSchemaUrn**](EnummockAccessTokenValidatorSchemaUrn.md) |  | 
**ClientIDClaimName** | Pointer to **string** | The name of the token claim that contains the OAuth2 client ID. | [optional] 
**ScopeClaimName** | Pointer to **string** | The name of the token claim that contains the scopes granted by the token. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple Mock Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Mock Access Token Validators defined within Directory Server but not necessarily contiguous. Mock Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 

## Methods

### NewMockAccessTokenValidatorShared

`func NewMockAccessTokenValidatorShared(schemas []EnummockAccessTokenValidatorSchemaUrn, evaluationOrderIndex int32, enabled bool, ) *MockAccessTokenValidatorShared`

NewMockAccessTokenValidatorShared instantiates a new MockAccessTokenValidatorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMockAccessTokenValidatorSharedWithDefaults

`func NewMockAccessTokenValidatorSharedWithDefaults() *MockAccessTokenValidatorShared`

NewMockAccessTokenValidatorSharedWithDefaults instantiates a new MockAccessTokenValidatorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MockAccessTokenValidatorShared) GetSchemas() []EnummockAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MockAccessTokenValidatorShared) GetSchemasOk() (*[]EnummockAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MockAccessTokenValidatorShared) SetSchemas(v []EnummockAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetClientIDClaimName

`func (o *MockAccessTokenValidatorShared) GetClientIDClaimName() string`

GetClientIDClaimName returns the ClientIDClaimName field if non-nil, zero value otherwise.

### GetClientIDClaimNameOk

`func (o *MockAccessTokenValidatorShared) GetClientIDClaimNameOk() (*string, bool)`

GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIDClaimName

`func (o *MockAccessTokenValidatorShared) SetClientIDClaimName(v string)`

SetClientIDClaimName sets ClientIDClaimName field to given value.

### HasClientIDClaimName

`func (o *MockAccessTokenValidatorShared) HasClientIDClaimName() bool`

HasClientIDClaimName returns a boolean if a field has been set.

### GetScopeClaimName

`func (o *MockAccessTokenValidatorShared) GetScopeClaimName() string`

GetScopeClaimName returns the ScopeClaimName field if non-nil, zero value otherwise.

### GetScopeClaimNameOk

`func (o *MockAccessTokenValidatorShared) GetScopeClaimNameOk() (*string, bool)`

GetScopeClaimNameOk returns a tuple with the ScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClaimName

`func (o *MockAccessTokenValidatorShared) SetScopeClaimName(v string)`

SetScopeClaimName sets ScopeClaimName field to given value.

### HasScopeClaimName

`func (o *MockAccessTokenValidatorShared) HasScopeClaimName() bool`

HasScopeClaimName returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *MockAccessTokenValidatorShared) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *MockAccessTokenValidatorShared) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *MockAccessTokenValidatorShared) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetIdentityMapper

`func (o *MockAccessTokenValidatorShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *MockAccessTokenValidatorShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *MockAccessTokenValidatorShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *MockAccessTokenValidatorShared) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *MockAccessTokenValidatorShared) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *MockAccessTokenValidatorShared) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *MockAccessTokenValidatorShared) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *MockAccessTokenValidatorShared) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *MockAccessTokenValidatorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MockAccessTokenValidatorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MockAccessTokenValidatorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MockAccessTokenValidatorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MockAccessTokenValidatorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MockAccessTokenValidatorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MockAccessTokenValidatorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


