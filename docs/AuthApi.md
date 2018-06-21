# \AuthApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionKeepaliveGet**](AuthApi.md#SessionKeepaliveGet) | **Get** /session/keepalive | Keepalive
[**SessionLoginGet**](AuthApi.md#SessionLoginGet) | **Get** /session/login | Login and get session ID
[**SessionLoginPost**](AuthApi.md#SessionLoginPost) | **Post** /session/login | MFA
[**SessionLogoutGet**](AuthApi.md#SessionLogoutGet) | **Get** /session/logout | Logout
[**SessionUnblockCaptchaPost**](AuthApi.md#SessionUnblockCaptchaPost) | **Post** /session/unblock-captcha | Unblock captcha


# **SessionKeepaliveGet**
> SessionKeepaliveGet(ctx, )
Keepalive

Keep alive current session 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionLoginGet**
> SessionLoginResp SessionLoginGet(ctx, )
Login and get session ID

Basic Authentication with the Authorization header The Authorization header is constructed as follows   1. The user email and password are combined with a single colon. (:)   2. The resulting string is encoded using a variant of Base64.   3. The authorization method and a space is then prepended to the encoded string, separated with a space (e.g. \"Basic \").   For example, 'test@example.com' as the user email and 'qwerty' as the password, then the field's value is the   base64-encoding of test@example.com:qwerty, or dGVzdEBleGFtcGxlLmNvbTpxd2VydHk=.   Then the Authorization header will appear as   'Authorization: Basic dGVzdEBleGFtcGxlLmNvbTpxd2VydHk=' 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SessionLoginResp**](SessionLoginResp.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionLoginPost**
> SessionLoginResp SessionLoginPost(ctx, body)
MFA

Login with MFA 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SessionLoginPostResp**](SessionLoginPostResp.md)|  | 

### Return type

[**SessionLoginResp**](SessionLoginResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionLogoutGet**
> SessionLogoutGet(ctx, )
Logout

Logout of current session 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionUnblockCaptchaPost**
> SessionUnblockCaptchaResp SessionUnblockCaptchaPost(ctx, body)
Unblock captcha

Unblock captcha 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UnblockCaptchaReq**](UnblockCaptchaReq.md)|  | 

### Return type

[**SessionUnblockCaptchaResp**](SessionUnblockCaptchaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

