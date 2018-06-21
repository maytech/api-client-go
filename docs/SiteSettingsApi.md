# \SiteSettingsApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsAuthMethodsGet**](SiteSettingsApi.md#SettingsAuthMethodsGet) | **Get** /settings/auth-methods | Get account auth-methods settings
[**SettingsGet**](SiteSettingsApi.md#SettingsGet) | **Get** /settings | Get site settings
[**SettingsSetPost**](SiteSettingsApi.md#SettingsSetPost) | **Post** /settings/set | Set site settings
[**SettingsUploadLogoLinkGet**](SiteSettingsApi.md#SettingsUploadLogoLinkGet) | **Get** /settings/upload-logo-link | Get logo upload link


# **SettingsAuthMethodsGet**
> []SettingsAuthMethodsRespItems SettingsAuthMethodsGet(ctx, )
Get account auth-methods settings

Get account auth-methods settings 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SettingsAuthMethodsRespItems**](SettingsAuthMethodsRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsGet**
> SettingsResp SettingsGet(ctx, )
Get site settings

Get site settings 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsResp**](SettingsResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsSetPost**
> SettingsResp SettingsSetPost(ctx, body)
Set site settings

Set site settings 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SettingsSetReq**](SettingsSetReq.md)|  | 

### Return type

[**SettingsResp**](SettingsResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsUploadLogoLinkGet**
> SettingsUploadLogoLinkResp SettingsUploadLogoLinkGet(ctx, )
Get logo upload link

Get account logo upload link 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsUploadLogoLinkResp**](SettingsUploadLogoLinkResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

