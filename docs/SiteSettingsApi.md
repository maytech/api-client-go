# \SiteSettingsApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsAuthMethodsGet**](SiteSettingsApi.md#SettingsAuthMethodsGet) | **Get** /settings/auth-methods | Get available authentication methods
[**SettingsGet**](SiteSettingsApi.md#SettingsGet) | **Get** /settings | Get site settings
[**SettingsSetPost**](SiteSettingsApi.md#SettingsSetPost) | **Post** /settings/set | Set site settings
[**SettingsUploadLogoLinkGet**](SiteSettingsApi.md#SettingsUploadLogoLinkGet) | **Get** /settings/upload-logo-link | Get a new logo upload link


# **SettingsAuthMethodsGet**
> []SettingsAuthMethodsRespItems SettingsAuthMethodsGet(ctx, )
Get available authentication methods

Get available authentication methods that can be set for the account. 

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

Retrieve information about the service settings adjusted for the account e.g. language, banner text, email footer, billing emails, PGP, permitted share types. 

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

Update service settings for the account e.g. enable or disable PGP, force 2FA for all users of the account. 

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
Get a new logo upload link

Get a unique key for uploading a new logo 

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

