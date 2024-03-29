/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PasswordPolicyAPIService PasswordPolicyAPI service
type PasswordPolicyAPIService service

type ApiAddPasswordPolicyRequest struct {
	ctx                      context.Context
	ApiService               *PasswordPolicyAPIService
	addPasswordPolicyRequest *AddPasswordPolicyRequest
}

// Create a new Password Policy in the config
func (r ApiAddPasswordPolicyRequest) AddPasswordPolicyRequest(addPasswordPolicyRequest AddPasswordPolicyRequest) ApiAddPasswordPolicyRequest {
	r.addPasswordPolicyRequest = &addPasswordPolicyRequest
	return r
}

func (r ApiAddPasswordPolicyRequest) Execute() (*PasswordPolicyResponse, *http.Response, error) {
	return r.ApiService.AddPasswordPolicyExecute(r)
}

/*
AddPasswordPolicy Add a new Password Policy to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPasswordPolicyRequest
*/
func (a *PasswordPolicyAPIService) AddPasswordPolicy(ctx context.Context) ApiAddPasswordPolicyRequest {
	return ApiAddPasswordPolicyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PasswordPolicyResponse
func (a *PasswordPolicyAPIService) AddPasswordPolicyExecute(r ApiAddPasswordPolicyRequest) (*PasswordPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPolicyAPIService.AddPasswordPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addPasswordPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("addPasswordPolicyRequest is required and must be specified")
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
	localVarPostBody = r.addPasswordPolicyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeletePasswordPolicyRequest struct {
	ctx                context.Context
	ApiService         *PasswordPolicyAPIService
	passwordPolicyName string
}

func (r ApiDeletePasswordPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePasswordPolicyExecute(r)
}

/*
DeletePasswordPolicy Delete a Password Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordPolicyName Name of the Password Policy
	@return ApiDeletePasswordPolicyRequest
*/
func (a *PasswordPolicyAPIService) DeletePasswordPolicy(ctx context.Context, passwordPolicyName string) ApiDeletePasswordPolicyRequest {
	return ApiDeletePasswordPolicyRequest{
		ApiService:         a,
		ctx:                ctx,
		passwordPolicyName: passwordPolicyName,
	}
}

// Execute executes the request
func (a *PasswordPolicyAPIService) DeletePasswordPolicyExecute(r ApiDeletePasswordPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPolicyAPIService.DeletePasswordPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-policies/{password-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-policy-name"+"}", url.PathEscape(parameterValueToString(r.passwordPolicyName, "passwordPolicyName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetPasswordPolicyRequest struct {
	ctx                context.Context
	ApiService         *PasswordPolicyAPIService
	passwordPolicyName string
}

func (r ApiGetPasswordPolicyRequest) Execute() (*PasswordPolicyResponse, *http.Response, error) {
	return r.ApiService.GetPasswordPolicyExecute(r)
}

/*
GetPasswordPolicy Returns a single Password Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordPolicyName Name of the Password Policy
	@return ApiGetPasswordPolicyRequest
*/
func (a *PasswordPolicyAPIService) GetPasswordPolicy(ctx context.Context, passwordPolicyName string) ApiGetPasswordPolicyRequest {
	return ApiGetPasswordPolicyRequest{
		ApiService:         a,
		ctx:                ctx,
		passwordPolicyName: passwordPolicyName,
	}
}

// Execute executes the request
//
//	@return PasswordPolicyResponse
func (a *PasswordPolicyAPIService) GetPasswordPolicyExecute(r ApiGetPasswordPolicyRequest) (*PasswordPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPolicyAPIService.GetPasswordPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-policies/{password-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-policy-name"+"}", url.PathEscape(parameterValueToString(r.passwordPolicyName, "passwordPolicyName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListPasswordPoliciesRequest struct {
	ctx        context.Context
	ApiService *PasswordPolicyAPIService
	filter     *string
}

// SCIM filter
func (r ApiListPasswordPoliciesRequest) Filter(filter string) ApiListPasswordPoliciesRequest {
	r.filter = &filter
	return r
}

func (r ApiListPasswordPoliciesRequest) Execute() (*PasswordPolicyListResponse, *http.Response, error) {
	return r.ApiService.ListPasswordPoliciesExecute(r)
}

/*
ListPasswordPolicies Returns a list of all Password Policy objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListPasswordPoliciesRequest
*/
func (a *PasswordPolicyAPIService) ListPasswordPolicies(ctx context.Context) ApiListPasswordPoliciesRequest {
	return ApiListPasswordPoliciesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PasswordPolicyListResponse
func (a *PasswordPolicyAPIService) ListPasswordPoliciesExecute(r ApiListPasswordPoliciesRequest) (*PasswordPolicyListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicyListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPolicyAPIService.ListPasswordPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdatePasswordPolicyRequest struct {
	ctx                context.Context
	ApiService         *PasswordPolicyAPIService
	passwordPolicyName string
	updateRequest      *UpdateRequest
}

// Update an existing Password Policy
func (r ApiUpdatePasswordPolicyRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdatePasswordPolicyRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdatePasswordPolicyRequest) Execute() (*PasswordPolicyResponse, *http.Response, error) {
	return r.ApiService.UpdatePasswordPolicyExecute(r)
}

/*
UpdatePasswordPolicy Update an existing Password Policy by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordPolicyName Name of the Password Policy
	@return ApiUpdatePasswordPolicyRequest
*/
func (a *PasswordPolicyAPIService) UpdatePasswordPolicy(ctx context.Context, passwordPolicyName string) ApiUpdatePasswordPolicyRequest {
	return ApiUpdatePasswordPolicyRequest{
		ApiService:         a,
		ctx:                ctx,
		passwordPolicyName: passwordPolicyName,
	}
}

// Execute executes the request
//
//	@return PasswordPolicyResponse
func (a *PasswordPolicyAPIService) UpdatePasswordPolicyExecute(r ApiUpdatePasswordPolicyRequest) (*PasswordPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPolicyAPIService.UpdatePasswordPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-policies/{password-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-policy-name"+"}", url.PathEscape(parameterValueToString(r.passwordPolicyName, "passwordPolicyName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
