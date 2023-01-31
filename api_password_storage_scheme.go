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

// PasswordStorageSchemeApiService PasswordStorageSchemeApi service
type PasswordStorageSchemeApiService service

type ApiAddPasswordStorageSchemeRequest struct {
	ctx                             context.Context
	ApiService                      *PasswordStorageSchemeApiService
	addPasswordStorageSchemeRequest *AddPasswordStorageSchemeRequest
}

// Create a new Password Storage Scheme in the config
func (r ApiAddPasswordStorageSchemeRequest) AddPasswordStorageSchemeRequest(addPasswordStorageSchemeRequest AddPasswordStorageSchemeRequest) ApiAddPasswordStorageSchemeRequest {
	r.addPasswordStorageSchemeRequest = &addPasswordStorageSchemeRequest
	return r
}

func (r ApiAddPasswordStorageSchemeRequest) Execute() (*AddPasswordStorageScheme200Response, *http.Response, error) {
	return r.ApiService.AddPasswordStorageSchemeExecute(r)
}

/*
AddPasswordStorageScheme Add a new Password Storage Scheme to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPasswordStorageSchemeRequest
*/
func (a *PasswordStorageSchemeApiService) AddPasswordStorageScheme(ctx context.Context) ApiAddPasswordStorageSchemeRequest {
	return ApiAddPasswordStorageSchemeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddPasswordStorageScheme200Response
func (a *PasswordStorageSchemeApiService) AddPasswordStorageSchemeExecute(r ApiAddPasswordStorageSchemeRequest) (*AddPasswordStorageScheme200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordStorageScheme200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordStorageSchemeApiService.AddPasswordStorageScheme")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-storage-schemes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addPasswordStorageSchemeRequest == nil {
		return localVarReturnValue, nil, reportError("addPasswordStorageSchemeRequest is required and must be specified")
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
	localVarPostBody = r.addPasswordStorageSchemeRequest
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

type ApiDeletePasswordStorageSchemeRequest struct {
	ctx                       context.Context
	ApiService                *PasswordStorageSchemeApiService
	passwordStorageSchemeName string
}

func (r ApiDeletePasswordStorageSchemeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePasswordStorageSchemeExecute(r)
}

/*
DeletePasswordStorageScheme Delete a Password Storage Scheme

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordStorageSchemeName Name of the Password Storage Scheme to be deleted
	@return ApiDeletePasswordStorageSchemeRequest
*/
func (a *PasswordStorageSchemeApiService) DeletePasswordStorageScheme(ctx context.Context, passwordStorageSchemeName string) ApiDeletePasswordStorageSchemeRequest {
	return ApiDeletePasswordStorageSchemeRequest{
		ApiService:                a,
		ctx:                       ctx,
		passwordStorageSchemeName: passwordStorageSchemeName,
	}
}

// Execute executes the request
func (a *PasswordStorageSchemeApiService) DeletePasswordStorageSchemeExecute(r ApiDeletePasswordStorageSchemeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordStorageSchemeApiService.DeletePasswordStorageScheme")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-storage-schemes/{password-storage-scheme-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-storage-scheme-name"+"}", url.PathEscape(parameterToString(r.passwordStorageSchemeName, "")), -1)

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

type ApiGetPasswordStorageSchemeRequest struct {
	ctx                       context.Context
	ApiService                *PasswordStorageSchemeApiService
	passwordStorageSchemeName string
}

func (r ApiGetPasswordStorageSchemeRequest) Execute() (*GetPasswordStorageScheme200Response, *http.Response, error) {
	return r.ApiService.GetPasswordStorageSchemeExecute(r)
}

/*
GetPasswordStorageScheme Returns a single Password Storage Scheme

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordStorageSchemeName Name of the Password Storage Scheme to be read
	@return ApiGetPasswordStorageSchemeRequest
*/
func (a *PasswordStorageSchemeApiService) GetPasswordStorageScheme(ctx context.Context, passwordStorageSchemeName string) ApiGetPasswordStorageSchemeRequest {
	return ApiGetPasswordStorageSchemeRequest{
		ApiService:                a,
		ctx:                       ctx,
		passwordStorageSchemeName: passwordStorageSchemeName,
	}
}

// Execute executes the request
//
//	@return GetPasswordStorageScheme200Response
func (a *PasswordStorageSchemeApiService) GetPasswordStorageSchemeExecute(r ApiGetPasswordStorageSchemeRequest) (*GetPasswordStorageScheme200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPasswordStorageScheme200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordStorageSchemeApiService.GetPasswordStorageScheme")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-storage-schemes/{password-storage-scheme-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-storage-scheme-name"+"}", url.PathEscape(parameterToString(r.passwordStorageSchemeName, "")), -1)

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

type ApiUpdatePasswordStorageSchemeRequest struct {
	ctx                       context.Context
	ApiService                *PasswordStorageSchemeApiService
	passwordStorageSchemeName string
	updateRequest             *UpdateRequest
}

// Update an existing Password Storage Scheme
func (r ApiUpdatePasswordStorageSchemeRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdatePasswordStorageSchemeRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdatePasswordStorageSchemeRequest) Execute() (*GetPasswordStorageScheme200Response, *http.Response, error) {
	return r.ApiService.UpdatePasswordStorageSchemeExecute(r)
}

/*
UpdatePasswordStorageScheme Update an existing Password Storage Scheme by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordStorageSchemeName Name of the Password Storage Scheme to be updated
	@return ApiUpdatePasswordStorageSchemeRequest
*/
func (a *PasswordStorageSchemeApiService) UpdatePasswordStorageScheme(ctx context.Context, passwordStorageSchemeName string) ApiUpdatePasswordStorageSchemeRequest {
	return ApiUpdatePasswordStorageSchemeRequest{
		ApiService:                a,
		ctx:                       ctx,
		passwordStorageSchemeName: passwordStorageSchemeName,
	}
}

// Execute executes the request
//
//	@return GetPasswordStorageScheme200Response
func (a *PasswordStorageSchemeApiService) UpdatePasswordStorageSchemeExecute(r ApiUpdatePasswordStorageSchemeRequest) (*GetPasswordStorageScheme200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPasswordStorageScheme200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordStorageSchemeApiService.UpdatePasswordStorageScheme")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-storage-schemes/{password-storage-scheme-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-storage-scheme-name"+"}", url.PathEscape(parameterToString(r.passwordStorageSchemeName, "")), -1)

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
