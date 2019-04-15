# \LanguageApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LanguagesDefaultGet**](LanguageApi.md#LanguagesDefaultGet) | **Get** /languages/default | Get default language
[**LanguagesGet**](LanguageApi.md#LanguagesGet) | **Get** /languages | List available languages


# **LanguagesDefaultGet**
> LanguagesDefaultResp LanguagesDefaultGet(ctx, )
Get default language

Get the default language ID. English (GB) is set as the default language for Quatrix. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LanguagesDefaultResp**](LanguagesDefaultResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LanguagesGet**
> []LanguagesRespItems LanguagesGet(ctx, )
List available languages

Get the list of languages that can be set for the account. English and Chinese are currently available. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LanguagesRespItems**](LanguagesRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

