//************************************************************************//
// API "congo" version v1: Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "fmt"

// ProposalHref returns the resource href.
func ProposalHref(api_version, userID, proposalID interface{}) string {
	return fmt.Sprintf("/%v/api/users/%v/proposals/%v", api_version, userID, proposalID)
}

// ReviewHref returns the resource href.
func ReviewHref(api_version, userID, proposalID, reviewID interface{}) string {
	return fmt.Sprintf("/%v/api/users/%v/proposals/%v/review/%v", api_version, userID, proposalID, reviewID)
}
