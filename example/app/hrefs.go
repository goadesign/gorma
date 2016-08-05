//************************************************************************//
// API "congo": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/gorma/example/design
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// ProposalHref returns the resource href.
func ProposalHref(userID, proposalID interface{}) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// ReviewHref returns the resource href.
func ReviewHref(userID, proposalID, reviewID interface{}) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	return fmt.Sprintf("/api/users/%v", userID)
}
