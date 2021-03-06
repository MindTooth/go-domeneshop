/*
 * Domeneshop API Documentation
 *
 * # Overview Domeneshop offers a simple, REST-based API, which currently supports the following  features: - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including: - Web hosting administration - Email address and email user/account administration  # Authentication The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website: <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package domeneshop

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// DdnsApiService DdnsApi service
type DdnsApiService service

// DyndnsUpdateGetOpts Optional parameters for the method 'DyndnsUpdateGet'
type DyndnsUpdateGetOpts struct {
    Myip optional.String
}

/*
DyndnsUpdateGet Update
Update DNS using the \&quot;IP update protocol\&quot;.  A DNS record for the given hostname will be created if it does not exist, or updated if it does. The record type (&#x60;A&#x60; or &#x60;AAAA&#x60; will automatically be detected).  If the DDNS implementation does not allow you to specify authentication, it can usually be specified inline in the URL:    &#x60;&#x60;&#x60;   https://{token}:{secret}@api.domeneshop.no/v0/dyndns/update?hostname&#x3D;example.com&amp;myip&#x3D;127.0.0.1   &#x60;&#x60;&#x60; 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param hostname The fully qualified domain (FQDN) to be updated, without trailing dot.
 * @param optional nil or *DyndnsUpdateGetOpts - Optional Parameters:
 * @param "Myip" (optional.String) -  The new IPv4 or IPv6 address to set. If not provided, the IP of the client making the API request will be used.
*/
func (a *DdnsApiService) DyndnsUpdateGet(ctx _context.Context, hostname string, localVarOptionals *DyndnsUpdateGetOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dyndns/update"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("hostname", parameterToString(hostname, ""))
	if localVarOptionals != nil && localVarOptionals.Myip.IsSet() {
		localVarQueryParams.Add("myip", parameterToString(localVarOptionals.Myip.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
