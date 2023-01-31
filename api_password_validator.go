/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PasswordValidatorApiService PasswordValidatorApi service
type PasswordValidatorApiService service

type ApiAddPasswordValidatorRequest struct {
	ctx                         context.Context
	ApiService                  *PasswordValidatorApiService
	addPasswordValidatorRequest *AddPasswordValidatorRequest
}

// Create a new Password Validator in the config
func (r ApiAddPasswordValidatorRequest) AddPasswordValidatorRequest(addPasswordValidatorRequest AddPasswordValidatorRequest) ApiAddPasswordValidatorRequest {
	r.addPasswordValidatorRequest = &addPasswordValidatorRequest
	return r
}

func (r ApiAddPasswordValidatorRequest) Execute() (*AddPasswordValidator200Response, *http.Response, error) {
	return r.ApiService.AddPasswordValidatorExecute(r)
}

/*
AddPasswordValidator Add a new Password Validator to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPasswordValidatorRequest
*/
func (a *PasswordValidatorApiService) AddPasswordValidator(ctx context.Context) ApiAddPasswordValidatorRequest {
	return ApiAddPasswordValidatorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddPasswordValidator200Response
func (a *PasswordValidatorApiService) AddPasswordValidatorExecute(r ApiAddPasswordValidatorRequest) (*AddPasswordValidator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordValidator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordValidatorApiService.AddPasswordValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-validators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addPasswordValidatorRequest == nil {
		return localVarReturnValue, nil, reportError("addPasswordValidatorRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addPasswordValidatorRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeletePasswordValidatorRequest struct {
	ctx                   context.Context
	ApiService            *PasswordValidatorApiService
	passwordValidatorName string
}

func (r ApiDeletePasswordValidatorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePasswordValidatorExecute(r)
}

/*
DeletePasswordValidator Delete a Password Validator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordValidatorName Name of the Password Validator to be deleted
	@return ApiDeletePasswordValidatorRequest
*/
func (a *PasswordValidatorApiService) DeletePasswordValidator(ctx context.Context, passwordValidatorName string) ApiDeletePasswordValidatorRequest {
	return ApiDeletePasswordValidatorRequest{
		ApiService:            a,
		ctx:                   ctx,
		passwordValidatorName: passwordValidatorName,
	}
}

// Execute executes the request
func (a *PasswordValidatorApiService) DeletePasswordValidatorExecute(r ApiDeletePasswordValidatorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordValidatorApiService.DeletePasswordValidator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-validators/{password-validator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-validator-name"+"}", url.PathEscape(parameterToString(r.passwordValidatorName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPasswordValidatorRequest struct {
	ctx                   context.Context
	ApiService            *PasswordValidatorApiService
	passwordValidatorName string
}

func (r ApiGetPasswordValidatorRequest) Execute() (*AddPasswordValidator200Response, *http.Response, error) {
	return r.ApiService.GetPasswordValidatorExecute(r)
}

/*
GetPasswordValidator Returns a single Password Validator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordValidatorName Name of the Password Validator to be read
	@return ApiGetPasswordValidatorRequest
*/
func (a *PasswordValidatorApiService) GetPasswordValidator(ctx context.Context, passwordValidatorName string) ApiGetPasswordValidatorRequest {
	return ApiGetPasswordValidatorRequest{
		ApiService:            a,
		ctx:                   ctx,
		passwordValidatorName: passwordValidatorName,
	}
}

// Execute executes the request
//
//	@return AddPasswordValidator200Response
func (a *PasswordValidatorApiService) GetPasswordValidatorExecute(r ApiGetPasswordValidatorRequest) (*AddPasswordValidator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordValidator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordValidatorApiService.GetPasswordValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-validators/{password-validator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-validator-name"+"}", url.PathEscape(parameterToString(r.passwordValidatorName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePasswordValidatorRequest struct {
	ctx                   context.Context
	ApiService            *PasswordValidatorApiService
	passwordValidatorName string
	updateRequest         *UpdateRequest
}

// Update an existing Password Validator
func (r ApiUpdatePasswordValidatorRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdatePasswordValidatorRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdatePasswordValidatorRequest) Execute() (*AddPasswordValidator200Response, *http.Response, error) {
	return r.ApiService.UpdatePasswordValidatorExecute(r)
}

/*
UpdatePasswordValidator Update an existing Password Validator by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordValidatorName Name of the Password Validator to be updated
	@return ApiUpdatePasswordValidatorRequest
*/
func (a *PasswordValidatorApiService) UpdatePasswordValidator(ctx context.Context, passwordValidatorName string) ApiUpdatePasswordValidatorRequest {
	return ApiUpdatePasswordValidatorRequest{
		ApiService:            a,
		ctx:                   ctx,
		passwordValidatorName: passwordValidatorName,
	}
}

// Execute executes the request
//
//	@return AddPasswordValidator200Response
func (a *PasswordValidatorApiService) UpdatePasswordValidatorExecute(r ApiUpdatePasswordValidatorRequest) (*AddPasswordValidator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordValidator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordValidatorApiService.UpdatePasswordValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-validators/{password-validator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-validator-name"+"}", url.PathEscape(parameterToString(r.passwordValidatorName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
