// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder is request builder for AndroidManagedStoreAccountEnterpriseSettings
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidManagedStoreAccountEnterpriseSettingsRequest
func (b *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Request() *AndroidManagedStoreAccountEnterpriseSettingsRequest {
	return &AndroidManagedStoreAccountEnterpriseSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidManagedStoreAccountEnterpriseSettingsRequest is request for AndroidManagedStoreAccountEnterpriseSettings
type AndroidManagedStoreAccountEnterpriseSettingsRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidManagedStoreAccountEnterpriseSettings
func (r *AndroidManagedStoreAccountEnterpriseSettingsRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidManagedStoreAccountEnterpriseSettings, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidManagedStoreAccountEnterpriseSettings
func (r *AndroidManagedStoreAccountEnterpriseSettingsRequest) Get() (*AndroidManagedStoreAccountEnterpriseSettings, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidManagedStoreAccountEnterpriseSettings
func (r *AndroidManagedStoreAccountEnterpriseSettingsRequest) Update(reqObj *AndroidManagedStoreAccountEnterpriseSettings) (*AndroidManagedStoreAccountEnterpriseSettings, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidManagedStoreAccountEnterpriseSettings
func (r *AndroidManagedStoreAccountEnterpriseSettingsRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}