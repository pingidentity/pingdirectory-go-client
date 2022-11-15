/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AddPasswordStorageScheme200Response - struct for AddPasswordStorageScheme200Response
type AddPasswordStorageScheme200Response struct {
	Aes256PasswordStorageSchemeResponse *Aes256PasswordStorageSchemeResponse
	AmazonSecretsManagerPasswordStorageSchemeResponse *AmazonSecretsManagerPasswordStorageSchemeResponse
	Argon2PasswordStorageSchemeResponse *Argon2PasswordStorageSchemeResponse
	AzureKeyVaultPasswordStorageSchemeResponse *AzureKeyVaultPasswordStorageSchemeResponse
	BcryptPasswordStorageSchemeResponse *BcryptPasswordStorageSchemeResponse
	ConjurPasswordStorageSchemeResponse *ConjurPasswordStorageSchemeResponse
	CryptPasswordStorageSchemeResponse *CryptPasswordStorageSchemeResponse
	Pbkdf2PasswordStorageSchemeResponse *Pbkdf2PasswordStorageSchemeResponse
	ScryptPasswordStorageSchemeResponse *ScryptPasswordStorageSchemeResponse
	ThirdPartyEnhancedPasswordStorageSchemeResponse *ThirdPartyEnhancedPasswordStorageSchemeResponse
	ThirdPartyPasswordStorageSchemeResponse *ThirdPartyPasswordStorageSchemeResponse
	VaultPasswordStorageSchemeResponse *VaultPasswordStorageSchemeResponse
}

// Aes256PasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns Aes256PasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func Aes256PasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *Aes256PasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		Aes256PasswordStorageSchemeResponse: v,
	}
}

// AmazonSecretsManagerPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns AmazonSecretsManagerPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func AmazonSecretsManagerPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *AmazonSecretsManagerPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		AmazonSecretsManagerPasswordStorageSchemeResponse: v,
	}
}

// Argon2PasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns Argon2PasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func Argon2PasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *Argon2PasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		Argon2PasswordStorageSchemeResponse: v,
	}
}

// AzureKeyVaultPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns AzureKeyVaultPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func AzureKeyVaultPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *AzureKeyVaultPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		AzureKeyVaultPasswordStorageSchemeResponse: v,
	}
}

// BcryptPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns BcryptPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func BcryptPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *BcryptPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		BcryptPasswordStorageSchemeResponse: v,
	}
}

// ConjurPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns ConjurPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func ConjurPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *ConjurPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		ConjurPasswordStorageSchemeResponse: v,
	}
}

// CryptPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns CryptPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func CryptPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *CryptPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		CryptPasswordStorageSchemeResponse: v,
	}
}

// Pbkdf2PasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns Pbkdf2PasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func Pbkdf2PasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *Pbkdf2PasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		Pbkdf2PasswordStorageSchemeResponse: v,
	}
}

// ScryptPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns ScryptPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func ScryptPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *ScryptPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		ScryptPasswordStorageSchemeResponse: v,
	}
}

// ThirdPartyEnhancedPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns ThirdPartyEnhancedPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func ThirdPartyEnhancedPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *ThirdPartyEnhancedPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		ThirdPartyEnhancedPasswordStorageSchemeResponse: v,
	}
}

// ThirdPartyPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns ThirdPartyPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func ThirdPartyPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *ThirdPartyPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		ThirdPartyPasswordStorageSchemeResponse: v,
	}
}

// VaultPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response is a convenience function that returns VaultPasswordStorageSchemeResponse wrapped in AddPasswordStorageScheme200Response
func VaultPasswordStorageSchemeResponseAsAddPasswordStorageScheme200Response(v *VaultPasswordStorageSchemeResponse) AddPasswordStorageScheme200Response {
	return AddPasswordStorageScheme200Response{
		VaultPasswordStorageSchemeResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPasswordStorageScheme200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Aes256PasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.Aes256PasswordStorageSchemeResponse)
	if err == nil {
		jsonAes256PasswordStorageSchemeResponse, _ := json.Marshal(dst.Aes256PasswordStorageSchemeResponse)
		if string(jsonAes256PasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.Aes256PasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.Aes256PasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into AmazonSecretsManagerPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.AmazonSecretsManagerPasswordStorageSchemeResponse)
	if err == nil {
		jsonAmazonSecretsManagerPasswordStorageSchemeResponse, _ := json.Marshal(dst.AmazonSecretsManagerPasswordStorageSchemeResponse)
		if string(jsonAmazonSecretsManagerPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.AmazonSecretsManagerPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.AmazonSecretsManagerPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into Argon2PasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.Argon2PasswordStorageSchemeResponse)
	if err == nil {
		jsonArgon2PasswordStorageSchemeResponse, _ := json.Marshal(dst.Argon2PasswordStorageSchemeResponse)
		if string(jsonArgon2PasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.Argon2PasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.Argon2PasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into AzureKeyVaultPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.AzureKeyVaultPasswordStorageSchemeResponse)
	if err == nil {
		jsonAzureKeyVaultPasswordStorageSchemeResponse, _ := json.Marshal(dst.AzureKeyVaultPasswordStorageSchemeResponse)
		if string(jsonAzureKeyVaultPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.AzureKeyVaultPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.AzureKeyVaultPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into BcryptPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.BcryptPasswordStorageSchemeResponse)
	if err == nil {
		jsonBcryptPasswordStorageSchemeResponse, _ := json.Marshal(dst.BcryptPasswordStorageSchemeResponse)
		if string(jsonBcryptPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.BcryptPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.BcryptPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into ConjurPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.ConjurPasswordStorageSchemeResponse)
	if err == nil {
		jsonConjurPasswordStorageSchemeResponse, _ := json.Marshal(dst.ConjurPasswordStorageSchemeResponse)
		if string(jsonConjurPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.ConjurPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConjurPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into CryptPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.CryptPasswordStorageSchemeResponse)
	if err == nil {
		jsonCryptPasswordStorageSchemeResponse, _ := json.Marshal(dst.CryptPasswordStorageSchemeResponse)
		if string(jsonCryptPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.CryptPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.CryptPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into Pbkdf2PasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.Pbkdf2PasswordStorageSchemeResponse)
	if err == nil {
		jsonPbkdf2PasswordStorageSchemeResponse, _ := json.Marshal(dst.Pbkdf2PasswordStorageSchemeResponse)
		if string(jsonPbkdf2PasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.Pbkdf2PasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.Pbkdf2PasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into ScryptPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.ScryptPasswordStorageSchemeResponse)
	if err == nil {
		jsonScryptPasswordStorageSchemeResponse, _ := json.Marshal(dst.ScryptPasswordStorageSchemeResponse)
		if string(jsonScryptPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.ScryptPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ScryptPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into ThirdPartyEnhancedPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyEnhancedPasswordStorageSchemeResponse)
	if err == nil {
		jsonThirdPartyEnhancedPasswordStorageSchemeResponse, _ := json.Marshal(dst.ThirdPartyEnhancedPasswordStorageSchemeResponse)
		if string(jsonThirdPartyEnhancedPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.ThirdPartyEnhancedPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyEnhancedPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into ThirdPartyPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyPasswordStorageSchemeResponse)
	if err == nil {
		jsonThirdPartyPasswordStorageSchemeResponse, _ := json.Marshal(dst.ThirdPartyPasswordStorageSchemeResponse)
		if string(jsonThirdPartyPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.ThirdPartyPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyPasswordStorageSchemeResponse = nil
	}

	// try to unmarshal data into VaultPasswordStorageSchemeResponse
	err = newStrictDecoder(data).Decode(&dst.VaultPasswordStorageSchemeResponse)
	if err == nil {
		jsonVaultPasswordStorageSchemeResponse, _ := json.Marshal(dst.VaultPasswordStorageSchemeResponse)
		if string(jsonVaultPasswordStorageSchemeResponse) == "{}" { // empty struct
			dst.VaultPasswordStorageSchemeResponse = nil
		} else {
			match++
		}
	} else {
		dst.VaultPasswordStorageSchemeResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Aes256PasswordStorageSchemeResponse = nil
		dst.AmazonSecretsManagerPasswordStorageSchemeResponse = nil
		dst.Argon2PasswordStorageSchemeResponse = nil
		dst.AzureKeyVaultPasswordStorageSchemeResponse = nil
		dst.BcryptPasswordStorageSchemeResponse = nil
		dst.ConjurPasswordStorageSchemeResponse = nil
		dst.CryptPasswordStorageSchemeResponse = nil
		dst.Pbkdf2PasswordStorageSchemeResponse = nil
		dst.ScryptPasswordStorageSchemeResponse = nil
		dst.ThirdPartyEnhancedPasswordStorageSchemeResponse = nil
		dst.ThirdPartyPasswordStorageSchemeResponse = nil
		dst.VaultPasswordStorageSchemeResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPasswordStorageScheme200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPasswordStorageScheme200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPasswordStorageScheme200Response) MarshalJSON() ([]byte, error) {
	if src.Aes256PasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.Aes256PasswordStorageSchemeResponse)
	}

	if src.AmazonSecretsManagerPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.AmazonSecretsManagerPasswordStorageSchemeResponse)
	}

	if src.Argon2PasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.Argon2PasswordStorageSchemeResponse)
	}

	if src.AzureKeyVaultPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.AzureKeyVaultPasswordStorageSchemeResponse)
	}

	if src.BcryptPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.BcryptPasswordStorageSchemeResponse)
	}

	if src.ConjurPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.ConjurPasswordStorageSchemeResponse)
	}

	if src.CryptPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.CryptPasswordStorageSchemeResponse)
	}

	if src.Pbkdf2PasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.Pbkdf2PasswordStorageSchemeResponse)
	}

	if src.ScryptPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.ScryptPasswordStorageSchemeResponse)
	}

	if src.ThirdPartyEnhancedPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.ThirdPartyEnhancedPasswordStorageSchemeResponse)
	}

	if src.ThirdPartyPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.ThirdPartyPasswordStorageSchemeResponse)
	}

	if src.VaultPasswordStorageSchemeResponse != nil {
		return json.Marshal(&src.VaultPasswordStorageSchemeResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPasswordStorageScheme200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Aes256PasswordStorageSchemeResponse != nil {
		return obj.Aes256PasswordStorageSchemeResponse
	}

	if obj.AmazonSecretsManagerPasswordStorageSchemeResponse != nil {
		return obj.AmazonSecretsManagerPasswordStorageSchemeResponse
	}

	if obj.Argon2PasswordStorageSchemeResponse != nil {
		return obj.Argon2PasswordStorageSchemeResponse
	}

	if obj.AzureKeyVaultPasswordStorageSchemeResponse != nil {
		return obj.AzureKeyVaultPasswordStorageSchemeResponse
	}

	if obj.BcryptPasswordStorageSchemeResponse != nil {
		return obj.BcryptPasswordStorageSchemeResponse
	}

	if obj.ConjurPasswordStorageSchemeResponse != nil {
		return obj.ConjurPasswordStorageSchemeResponse
	}

	if obj.CryptPasswordStorageSchemeResponse != nil {
		return obj.CryptPasswordStorageSchemeResponse
	}

	if obj.Pbkdf2PasswordStorageSchemeResponse != nil {
		return obj.Pbkdf2PasswordStorageSchemeResponse
	}

	if obj.ScryptPasswordStorageSchemeResponse != nil {
		return obj.ScryptPasswordStorageSchemeResponse
	}

	if obj.ThirdPartyEnhancedPasswordStorageSchemeResponse != nil {
		return obj.ThirdPartyEnhancedPasswordStorageSchemeResponse
	}

	if obj.ThirdPartyPasswordStorageSchemeResponse != nil {
		return obj.ThirdPartyPasswordStorageSchemeResponse
	}

	if obj.VaultPasswordStorageSchemeResponse != nil {
		return obj.VaultPasswordStorageSchemeResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddPasswordStorageScheme200Response struct {
	value *AddPasswordStorageScheme200Response
	isSet bool
}

func (v NullableAddPasswordStorageScheme200Response) Get() *AddPasswordStorageScheme200Response {
	return v.value
}

func (v *NullableAddPasswordStorageScheme200Response) Set(val *AddPasswordStorageScheme200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPasswordStorageScheme200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPasswordStorageScheme200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPasswordStorageScheme200Response(val *AddPasswordStorageScheme200Response) *NullableAddPasswordStorageScheme200Response {
	return &NullableAddPasswordStorageScheme200Response{value: val, isSet: true}
}

func (v NullableAddPasswordStorageScheme200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPasswordStorageScheme200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


