/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AmazonAwsExternalServerResponse struct for AmazonAwsExternalServerResponse
type AmazonAwsExternalServerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the External Server
	Id      string                                 `json:"id"`
	Schemas []EnumamazonAwsExternalServerSchemaUrn `json:"schemas"`
	// The access key ID that will be used if authentication should use an access key. If this is provided, then an aws-secret-access-key must also be provided. If this is not provided, then no aws-secret-access-key may be configured, and the server must be running in an EC2 instance that is configured with an IAM role with permission to perform the necessary operations.
	AwsAccessKeyID *string `json:"awsAccessKeyID,omitempty"`
	// The secret access key that will be used if authentication should use an access key. If this is provided, then an aws-access-key-id must also be provided. If this is not provided, then no aws-access-key-id may be configured, and the server must be running in an EC2 instance that is configured with an IAM role with permission to perform the necessary operations.
	AwsSecretAccessKey *string `json:"awsSecretAccessKey,omitempty"`
	// The name of the AWS region containing the resources that will be accessed.
	AwsRegionName string `json:"awsRegionName"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewAmazonAwsExternalServerResponse instantiates a new AmazonAwsExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAwsExternalServerResponse(id string, schemas []EnumamazonAwsExternalServerSchemaUrn, awsRegionName string) *AmazonAwsExternalServerResponse {
	this := AmazonAwsExternalServerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AwsRegionName = awsRegionName
	return &this
}

// NewAmazonAwsExternalServerResponseWithDefaults instantiates a new AmazonAwsExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAwsExternalServerResponseWithDefaults() *AmazonAwsExternalServerResponse {
	this := AmazonAwsExternalServerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AmazonAwsExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AmazonAwsExternalServerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AmazonAwsExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AmazonAwsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AmazonAwsExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AmazonAwsExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AmazonAwsExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AmazonAwsExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AmazonAwsExternalServerResponse) GetSchemas() []EnumamazonAwsExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumamazonAwsExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetSchemasOk() ([]EnumamazonAwsExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AmazonAwsExternalServerResponse) SetSchemas(v []EnumamazonAwsExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetAwsAccessKeyID returns the AwsAccessKeyID field value if set, zero value otherwise.
func (o *AmazonAwsExternalServerResponse) GetAwsAccessKeyID() string {
	if o == nil || isNil(o.AwsAccessKeyID) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyID
}

// GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetAwsAccessKeyIDOk() (*string, bool) {
	if o == nil || isNil(o.AwsAccessKeyID) {
		return nil, false
	}
	return o.AwsAccessKeyID, true
}

// HasAwsAccessKeyID returns a boolean if a field has been set.
func (o *AmazonAwsExternalServerResponse) HasAwsAccessKeyID() bool {
	if o != nil && !isNil(o.AwsAccessKeyID) {
		return true
	}

	return false
}

// SetAwsAccessKeyID gets a reference to the given string and assigns it to the AwsAccessKeyID field.
func (o *AmazonAwsExternalServerResponse) SetAwsAccessKeyID(v string) {
	o.AwsAccessKeyID = &v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise.
func (o *AmazonAwsExternalServerResponse) GetAwsSecretAccessKey() string {
	if o == nil || isNil(o.AwsSecretAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil || isNil(o.AwsSecretAccessKey) {
		return nil, false
	}
	return o.AwsSecretAccessKey, true
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *AmazonAwsExternalServerResponse) HasAwsSecretAccessKey() bool {
	if o != nil && !isNil(o.AwsSecretAccessKey) {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given string and assigns it to the AwsSecretAccessKey field.
func (o *AmazonAwsExternalServerResponse) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = &v
}

// GetAwsRegionName returns the AwsRegionName field value
func (o *AmazonAwsExternalServerResponse) GetAwsRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsRegionName
}

// GetAwsRegionNameOk returns a tuple with the AwsRegionName field value
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetAwsRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRegionName, true
}

// SetAwsRegionName sets field value
func (o *AmazonAwsExternalServerResponse) SetAwsRegionName(v string) {
	o.AwsRegionName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmazonAwsExternalServerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAwsExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmazonAwsExternalServerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmazonAwsExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

func (o AmazonAwsExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.AwsAccessKeyID) {
		toSerialize["awsAccessKeyID"] = o.AwsAccessKeyID
	}
	if !isNil(o.AwsSecretAccessKey) {
		toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey
	}
	if true {
		toSerialize["awsRegionName"] = o.AwsRegionName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonAwsExternalServerResponse struct {
	value *AmazonAwsExternalServerResponse
	isSet bool
}

func (v NullableAmazonAwsExternalServerResponse) Get() *AmazonAwsExternalServerResponse {
	return v.value
}

func (v *NullableAmazonAwsExternalServerResponse) Set(val *AmazonAwsExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAwsExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAwsExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAwsExternalServerResponse(val *AmazonAwsExternalServerResponse) *NullableAmazonAwsExternalServerResponse {
	return &NullableAmazonAwsExternalServerResponse{value: val, isSet: true}
}

func (v NullableAmazonAwsExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAwsExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
