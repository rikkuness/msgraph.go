// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TextColumn undocumented
type TextColumn struct {
	// Object is the base model of TextColumn
	Object
	// AllowMultipleLines undocumented
	AllowMultipleLines *bool `json:"allowMultipleLines,omitempty"`
	// AppendChangesToExistingText undocumented
	AppendChangesToExistingText *bool `json:"appendChangesToExistingText,omitempty"`
	// LinesForEditing undocumented
	LinesForEditing *int `json:"linesForEditing,omitempty"`
	// MaxLength undocumented
	MaxLength *int `json:"maxLength,omitempty"`
	// TextType undocumented
	TextType *string `json:"textType,omitempty"`
}