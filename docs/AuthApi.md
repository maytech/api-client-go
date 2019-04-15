# \AuthApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionKeepaliveGet**](AuthApi.md#SessionKeepaliveGet) | **Get** /session/keepalive | Refresh session expiration time
[**SessionLoginGet**](AuthApi.md#SessionLoginGet) | **Get** /session/login | Log in and get session ID details
[**SessionLoginPost**](AuthApi.md#SessionLoginPost) | **Post** /session/login | Log in to the account using MFA
[**SessionLogoutGet**](AuthApi.md#SessionLogoutGet) | **Get** /session/logout | Close the user&#39;s session
[**SessionUnblockCaptchaPost**](AuthApi.md#SessionUnblockCaptchaPost) | **Post** /session/unblock-captcha | Unblock the session using CAPTCHA


# **SessionKeepaliveGet**
> SessionKeepaliveGet(ctx, )
Refresh session expiration time

Refresh the existing session using the ID for the session. The session will expire if there were no API actions for 15 minutes. 

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
Log in and get session ID details

Get session information for the given session ID and log in to the account.  The easiest way to authenticate is using [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).  To proceed with the authorization, you need to have Quatrix account. If you don’t have the one, you can set up a free trial account [here](https://www.maytech.net/freetrial.html#trial=qtrx).  As the basic authentication requires the authentication of the user with the user ID and password, the Authorization header should be constructed as follows:    1. The user email and password are combined with a single colon. (:)    2. The resulting string is encoded using a variant of Base64.    3. The authorization method and a space is then prepended to the encoded string, separated with a space (e.g. \"Basic \").    For example, 'test@example.com' as the user email and 'qwerty' as the password, then the field's value is the following:    base64-encoding of test@example.com:qwerty, or dGVzdEBleGFtcGxlLmNvbTpxd2VydHk=.    The authorization header will appear as    'Authorization: Basic dGVzdEBleGFtcGxlLmNvbTpxd2VydHk=' 

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
Log in to the account using MFA

Use to generate a session login token in scenarios in which 2FA or PIN are required. 

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
Close the user's session

Close the session for the currently logged in user. 

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
Unblock the session using CAPTCHA

Enter CAPTCHA to log in the user to the current session. 

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

