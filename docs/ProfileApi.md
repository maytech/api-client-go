# \ProfileApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Profile2faGenerateGet**](ProfileApi.md#Profile2faGenerateGet) | **Get** /profile/2fa/generate | Generate a new 2FA code
[**ProfileGet**](ProfileApi.md#ProfileGet) | **Get** /profile | Get profile metadata
[**ProfileInfoGet**](ProfileApi.md#ProfileInfoGet) | **Get** /profile/info | Retrieve additional profile info
[**ProfileRemoveMfaPost**](ProfileApi.md#ProfileRemoveMfaPost) | **Post** /profile/remove-mfa | Disable MFA for the logged-in user
[**ProfileSetMfaPost**](ProfileApi.md#ProfileSetMfaPost) | **Post** /profile/set-mfa | Enable MFA for the logged-in user
[**ProfileSetPasswordPost**](ProfileApi.md#ProfileSetPasswordPost) | **Post** /profile/set-password | Change profile password
[**ProfileSetPost**](ProfileApi.md#ProfileSetPost) | **Post** /profile/set | Update profile metadata


# **Profile2faGenerateGet**
> Profile2faGenerateGet(ctx, )
Generate a new 2FA code

Generate a new verification code for 2FA with QR code. 

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

# **ProfileGet**
> ProfileResp ProfileGet(ctx, )
Get profile metadata

Retrieve profile information of the current user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProfileResp**](ProfileResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfileInfoGet**
> ProfileInfoResp ProfileInfoGet(ctx, )
Retrieve additional profile info

Get additional details about the account e.g. number of used and available user licenses. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProfileInfoResp**](ProfileInfoResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfileRemoveMfaPost**
> ProfileRemoveMfaResp ProfileRemoveMfaPost(ctx, body)
Disable MFA for the logged-in user

Trun off MFA (multifactor authentication) for the user who requested 2FA deactivation while editing their profile. This operation is possible if it was not forced by the administrator. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ProfileRemoveMfaReq**](ProfileRemoveMfaReq.md)|  | 

### Return type

[**ProfileRemoveMfaResp**](ProfileRemoveMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfileSetMfaPost**
> ProfileSetMfaResp ProfileSetMfaPost(ctx, body)
Enable MFA for the logged-in user

Turn on MFA (multifactor authentication) for the user who requested 2FA activation while editing their profile. MFA adds an additional secure step on the way to log in to the account by using one more authentication method beyond the email and password. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ProfileSetMfaReq**](ProfileSetMfaReq.md)|  | 

### Return type

[**ProfileSetMfaResp**](ProfileSetMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfileSetPasswordPost**
> ProfileSetPasswordResp ProfileSetPasswordPost(ctx, body)
Change profile password

Change the account password for the logged-in user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ProfileSetPasswordReq**](ProfileSetPasswordReq.md)|  | 

### Return type

[**ProfileSetPasswordResp**](ProfileSetPasswordResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfileSetPost**
> ProfileSetResp ProfileSetPost(ctx, optional)
Update profile metadata

Edit profile information of the current user including name, email, language and message signature. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProfileSetReq**](ProfileSetReq.md)|  | 

### Return type

[**ProfileSetResp**](ProfileSetResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

