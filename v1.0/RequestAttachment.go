// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AttachmentRequestBuilder is request builder for Attachment
type AttachmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttachmentRequest
func (b *AttachmentRequestBuilder) Request() *AttachmentRequest {
	return &AttachmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttachmentRequest is request for Attachment
type AttachmentRequest struct{ BaseRequest }

// Get performs GET request for Attachment
func (r *AttachmentRequest) Get(ctx context.Context) (resObj *Attachment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Attachment
func (r *AttachmentRequest) Update(ctx context.Context, reqObj *Attachment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Attachment
func (r *AttachmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
