//************************************************************************//
// API "congo" version v1: Resource and Payload Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package review

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/raphael/goa"
)

func ReviewFromV1CreatePayload(ctx *v1.CreateReviewContext) Review {
	payload := ctx.Payload
	var err error
	target, _ := MarshalCreateReviewPayload(payload, err)

	return target
}

// MarshalCreateReviewPayload validates and renders an instance of CreateReviewPayload into a interface{}
func MarshalCreateReviewPayload(source *v1.CreateReviewPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp11 := map[string]interface{}{
		"User":        source.User,
		"comment":     source.Comment,
		"id":          source.ID,
		"proposal_id": source.ProposalId,
		"rating":      source.Rating,
		"reviewers":   source.Reviewers,
	}
	target = tmp11
	return
}

func ReviewFromV1UpdatePayload(ctx *v1.UpdateReviewContext) Review {
	payload := ctx.Payload
	var err error
	target, _ := MarshalUpdateReviewPayload(payload, err)

	return target
}

// MarshalUpdateReviewPayload validates and renders an instance of UpdateReviewPayload into a interface{}
func MarshalUpdateReviewPayload(source *v1.UpdateReviewPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp12 := map[string]interface{}{
		"User":        source.User,
		"comment":     source.Comment,
		"id":          source.ID,
		"proposal_id": source.ProposalId,
		"rating":      source.Rating,
		"reviewers":   source.Reviewers,
	}
	target = tmp12
	return
}
