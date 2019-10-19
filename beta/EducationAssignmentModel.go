// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationAssignment undocumented
type EducationAssignment struct {
	Entity
	// ClassID undocumented
	ClassID *string `json:"classId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Instructions undocumented
	Instructions *EducationItemBody `json:"instructions,omitempty"`
	// CloseDateTime undocumented
	CloseDateTime *time.Time `json:"closeDateTime,omitempty"`
	// DueDateTime undocumented
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// AssignDateTime undocumented
	AssignDateTime *time.Time `json:"assignDateTime,omitempty"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
	// AssignTo undocumented
	AssignTo *EducationAssignmentRecipient `json:"assignTo,omitempty"`
	// AllowLateSubmissions undocumented
	AllowLateSubmissions *bool `json:"allowLateSubmissions,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// AllowStudentsToAddResourcesToSubmission undocumented
	AllowStudentsToAddResourcesToSubmission *bool `json:"allowStudentsToAddResourcesToSubmission,omitempty"`
	// Status undocumented
	Status *EducationAssignmentStatus `json:"status,omitempty"`
	// Resources undocumented
	Resources []EducationAssignmentResource `json:"resources,omitempty"`
	// Submissions undocumented
	Submissions []EducationSubmission `json:"submissions,omitempty"`
	// Categories undocumented
	Categories []EducationCategory `json:"categories,omitempty"`
	// Rubric undocumented
	Rubric *EducationRubric `json:"rubric,omitempty"`
}