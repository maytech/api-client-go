# \PasswordResetApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetPasswordMetadataIdGet**](PasswordResetApi.md#ResetPasswordMetadataIdGet) | **Get** /reset-password/metadata/{id} | Get password reset request metadata
[**ResetPasswordRequestPost**](PasswordResetApi.md#ResetPasswordRequestPost) | **Post** /reset-password/request | Request password reset
[**ResetPasswordResetIdPost**](PasswordResetApi.md#ResetPasswordResetIdPost) | **Post** /reset-password/reset/{id} | Reset password


# **ResetPasswordMetadataIdGet**
> ResetPasswordMetadataResp ResetPasswordMetadataIdGet(ctx, id)
Get password reset request metadata

Retrieve information about the password reset request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of the password reset request | 

### Return type

[**ResetPasswordMetadataResp**](ResetPasswordMetadataResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPasswordRequestPost**
> ResetPasswordRequestResp ResetPasswordRequestPost(ctx, body)
Request password reset

Send an email with the request to reset password (including the link) to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ResetPasswordRequestReq**](ResetPasswordRequestReq.md)|  | 

### Return type

[**ResetPasswordRequestResp**](ResetPasswordRequestResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPasswordResetIdPost**
> IdResp ResetPasswordResetIdPost(ctx, id, body)
Reset password

Change the password based on the existing password reset request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a password reset request | 
  **body** | [**ResetPasswordResetReq**](ResetPasswordResetReq.md)|  | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

