// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CommsNotification undocumented
type CommsNotification struct {
	// Object is the base model of CommsNotification
	Object
	// ChangeType undocumented
	ChangeType *ChangeType `json:"changeType,omitempty"`
	// ResourceURL undocumented
	ResourceURL *string `json:"resourceUrl,omitempty"`
}

// CommsNotifications undocumented
type CommsNotifications struct {
	// Object is the base model of CommsNotifications
	Object
	// Value undocumented
	Value []CommsNotification `json:"value,omitempty"`
}

// CommsOperation undocumented
type CommsOperation struct {
	// Entity is the base model of CommsOperation
	Entity
	// Status undocumented
	Status *OperationStatus `json:"status,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
	// ResultInfo undocumented
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`
}