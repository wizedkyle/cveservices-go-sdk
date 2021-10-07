package cveservices_go_sdk

import (
	"fmt"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
UpdateUser
Updates a user record from the organization. No roles are needed to access the endpoint.
Expected Behavior:
Secretariat - Can update a user record for any Organization.
Admin User - Can only update a user record for users that belongs to the same Organization.
Regular User - Can only update its own user record.
*/
func (a *APIClient) UpdateUser(username string, localVarOptionals *types.UpdateUserOpts) (types.UpdatedUserResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue types.UpdatedUserResponse
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/org/{organization}/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", fmt.Sprintf("%v", a.Cfg.Organization), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Active.IsSet() {
		localVarQueryParams.Add("active", parameterToString(localVarOptionals.Active.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NewUsername.IsSet() {
		localVarQueryParams.Add("new_username", parameterToString(localVarOptionals.NewUsername.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgShortname.IsSet() {
		localVarQueryParams.Add("org_shortname", parameterToString(localVarOptionals.OrgShortname.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameFirst.IsSet() {
		localVarQueryParams.Add("name.first", parameterToString(localVarOptionals.NameFirst.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameLast.IsSet() {
		localVarQueryParams.Add("name.last", parameterToString(localVarOptionals.NameLast.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameMiddle.IsSet() {
		localVarQueryParams.Add("name.middle", parameterToString(localVarOptionals.NameMiddle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameSuffix.IsSet() {
		localVarQueryParams.Add("name.suffix", parameterToString(localVarOptionals.NameSuffix.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActiveRolesAdd.IsSet() {
		localVarQueryParams.Add("active_roles.add", parameterToString(localVarOptionals.ActiveRolesAdd.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActiveRolesRemove.IsSet() {
		localVarQueryParams.Add("active_roles.remove", parameterToString(localVarOptionals.ActiveRolesRemove.Value(), ""))
	}
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	} else if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v types.UpdatedUserResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		} else if localVarHttpResponse.StatusCode >= 400 {
			var v types.ErrorResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			if v.Message != "" {
				newErr.error = v.Message
			} else {
				newErr.error = v.Message
			}
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
