# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Vv210SandboxDataImport**](DefaultApi.md#Vv210SandboxDataImport) | **Post** /obp/v2.1.0/sandbox/data-import | Create sandbox

# **Vv210SandboxDataImport**
> Object Vv210SandboxDataImport(ctx, )
Create sandbox

<p>Import bulk data into the sandbox (Authenticated access).</p><p>This call can be used to create banks, users, accounts and transactions which are stored in the local RDBMS.</p><p>The user needs to have CanCreateSandbox entitlement.</p><p>Note: This is a monolithic call. You could also use a combination of endpoints including create bank, create user, create account and create transaction request to create similar data.</p><p>An example of an import set of data (json) can be found <a href=\"htt

### Required Parameters
This endpoint does not need any parameter.

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

