// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MediaContentRatingAustralia undocumented
type MediaContentRatingAustralia struct {
	// MovieRating Movies rating selected for Australia
	MovieRating *RatingAustraliaMoviesType `json:"movieRating,omitempty"`
	// TvRating TV rating selected for Australia
	TvRating *RatingAustraliaTelevisionType `json:"tvRating,omitempty"`
}