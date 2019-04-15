# \MFAApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfileRemoveMfaPost**](MFAApi.md#ProfileRemoveMfaPost) | **Post** /profile/remove-mfa | Disable MFA for the logged-in user
[**ProfileSetMfaPost**](MFAApi.md#ProfileSetMfaPost) | **Post** /profile/set-mfa | Enable MFA for the logged-in user
[**SessionLoginPost**](MFAApi.md#SessionLoginPost) | **Post** /session/login | Log in to the account using MFA
[**UserRemoveMfaPost**](MFAApi.md#UserRemoveMfaPost) | **Post** /user/remove-mfa | Disable MFA for users
[**UserResetMfaPost**](MFAApi.md#UserResetMfaPost) | **Post** /user/reset-mfa | Update existing MFA settings for users
[**UserSetMfaPost**](MFAApi.md#UserSetMfaPost) | **Post** /user/set-mfa | Enable MFA for users


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

# **UserRemoveMfaPost**
> ProfileRemoveMfaResp UserRemoveMfaPost(ctx, body)
Disable MFA for users

Deactivate MFA for a user by specified user ID. If 2FA is forced for the account, the user won’t be able to deactivate it using this API call. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserRemoveMfaReq**](UserRemoveMfaReq.md)|  | 

### Return type

[**ProfileRemoveMfaResp**](ProfileRemoveMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserResetMfaPost**
> ProfileRemoveMfaResp UserResetMfaPost(ctx, body)
Update existing MFA settings for users

Remove already set MFA code values, the MFA will be left activated. All logged-in sessions of users will be stopped. On the next login the user should set MFA again. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserResetMfaReq**](UserResetMfaReq.md)|  | 

### Return type

[**ProfileRemoveMfaResp**](ProfileRemoveMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSetMfaPost**
> ProfileRemoveMfaResp UserSetMfaPost(ctx, body)
Enable MFA for users

Activate MFA for a user by specified user ID. If auth type is 2FA, the user will be forced to use MFA. All logged-in sessions of users will be stopped. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserSetMfaReq**](UserSetMfaReq.md)|  | 

### Return type

[**ProfileRemoveMfaResp**](ProfileRemoveMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

