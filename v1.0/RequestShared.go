// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SharedDriveItemRequestBuilder is request builder for SharedDriveItem
type SharedDriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharedDriveItemRequest
func (b *SharedDriveItemRequestBuilder) Request() *SharedDriveItemRequest {
	return &SharedDriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SharedDriveItemRequest is request for SharedDriveItem
type SharedDriveItemRequest struct{ BaseRequest }

// Get performs GET request for SharedDriveItem
func (r *SharedDriveItemRequest) Get(ctx context.Context) (resObj *SharedDriveItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharedDriveItem
func (r *SharedDriveItemRequest) Update(ctx context.Context, reqObj *SharedDriveItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharedDriveItem
func (r *SharedDriveItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SharedInsightRequestBuilder is request builder for SharedInsight
type SharedInsightRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharedInsightRequest
func (b *SharedInsightRequestBuilder) Request() *SharedInsightRequest {
	return &SharedInsightRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SharedInsightRequest is request for SharedInsight
type SharedInsightRequest struct{ BaseRequest }

// Get performs GET request for SharedInsight
func (r *SharedInsightRequest) Get(ctx context.Context) (resObj *SharedInsight, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharedInsight
func (r *SharedInsightRequest) Update(ctx context.Context, reqObj *SharedInsight) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharedInsight
func (r *SharedInsightRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}