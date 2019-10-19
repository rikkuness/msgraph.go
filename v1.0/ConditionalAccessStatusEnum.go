// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConditionalAccessStatus undocumented
type ConditionalAccessStatus int

const (
	// ConditionalAccessStatusVSuccess undocumented
	ConditionalAccessStatusVSuccess ConditionalAccessStatus = 0
	// ConditionalAccessStatusVFailure undocumented
	ConditionalAccessStatusVFailure ConditionalAccessStatus = 1
	// ConditionalAccessStatusVNotApplied undocumented
	ConditionalAccessStatusVNotApplied ConditionalAccessStatus = 2
	// ConditionalAccessStatusVUnknownFutureValue undocumented
	ConditionalAccessStatusVUnknownFutureValue ConditionalAccessStatus = 3
)

// ConditionalAccessStatusPSuccess returns a pointer to ConditionalAccessStatusVSuccess
func ConditionalAccessStatusPSuccess() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVSuccess
	return &v
}

// ConditionalAccessStatusPFailure returns a pointer to ConditionalAccessStatusVFailure
func ConditionalAccessStatusPFailure() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVFailure
	return &v
}

// ConditionalAccessStatusPNotApplied returns a pointer to ConditionalAccessStatusVNotApplied
func ConditionalAccessStatusPNotApplied() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVNotApplied
	return &v
}

// ConditionalAccessStatusPUnknownFutureValue returns a pointer to ConditionalAccessStatusVUnknownFutureValue
func ConditionalAccessStatusPUnknownFutureValue() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVUnknownFutureValue
	return &v
}