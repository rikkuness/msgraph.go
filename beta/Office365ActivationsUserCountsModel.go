// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Office365ActivationsUserCounts undocumented
type Office365ActivationsUserCounts struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// ProductType undocumented
	ProductType *string `json:"productType,omitempty"`
	// Assigned undocumented
	Assigned *int `json:"assigned,omitempty"`
	// Activated undocumented
	Activated *int `json:"activated,omitempty"`
	// SharedComputerActivation undocumented
	SharedComputerActivation *int `json:"sharedComputerActivation,omitempty"`
}