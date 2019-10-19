// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OutlookItemRequestBuilder is request builder for OutlookItem
type OutlookItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookItemRequest
func (b *OutlookItemRequestBuilder) Request() *OutlookItemRequest {
	return &OutlookItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookItemRequest is request for OutlookItem
type OutlookItemRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookItem
func (r *OutlookItemRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookItem, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OutlookItem
func (r *OutlookItemRequest) Get() (*OutlookItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OutlookItem
func (r *OutlookItemRequest) Update(reqObj *OutlookItem) (*OutlookItem, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OutlookItem
func (r *OutlookItemRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}