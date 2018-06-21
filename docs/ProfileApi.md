# \ProfileApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Profile2faGenerateGet**](ProfileApi.md#Profile2faGenerateGet) | **Get** /profile/2fa/generate | Generate new 2fa code
[**ProfileGet**](ProfileApi.md#ProfileGet) | **Get** /profile | Profile metadata
[**ProfileInfoGet**](ProfileApi.md#ProfileInfoGet) | **Get** /profile/info | Additional profile info
[**ProfileRemoveMfaPost**](ProfileApi.md#ProfileRemoveMfaPost) | **Post** /profile/remove-mfa | Remove MFA for account
[**ProfileSetMfaPost**](ProfileApi.md#ProfileSetMfaPost) | **Post** /profile/set-mfa | Set MFA enabled for account
[**ProfileSetPasswordPost**](ProfileApi.md#ProfileSetPasswordPost) | **Post** /profile/set-password | Change profile password
[**ProfileSetPost**](ProfileApi.md#ProfileSetPost) | **Post** /profile/set | Update profile metadata


# **Profile2faGenerateGet**
> Profile2faGenerateGet(ctx, )
Generate new 2fa code

Generate new 2fa code QR code 

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
Profile metadata

Get profile metadata 

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
Additional profile info

Get additional profile info 

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
Remove MFA for account

Remove MFA for account if it was not forced by admin 

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
Set MFA enabled for account

Set multi factor autorization method (MFA) enabled for account 

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

Change profile password 

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

Update profile metadata 

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

