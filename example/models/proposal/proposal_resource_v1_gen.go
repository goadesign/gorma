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

package proposal

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/raphael/goa"
)

func ProposalFromV1CreatePayload(ctx *v1.CreateProposalContext) Proposal {
	payload := ctx.Payload
	var err error
	target, _ := MarshalCreateProposalPayload(payload, err)

	return target
}

// MarshalCreateProposalPayload validates and renders an instance of CreateProposalPayload into a interface{}
func MarshalCreateProposalPayload(source *v1.CreateProposalPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp9 := map[string]interface{}{
		"abstract":   source.Abstract,
		"created_at": source.CreatedAt,
		"deleted_at": source.DeletedAt,
		"detail":     source.Detail,
		"first_name": source.FirstName,
		"id":         source.ID,
		"m2reviews":  source.M2reviews,
		"reviews":    source.Reviews,
		"title":      source.Title,
		"updated_at": source.UpdatedAt,
		"user_id":    source.UserId,
		"withdrawn":  source.Withdrawn,
	}
	target = tmp9
	return
}

func ProposalFromV1UpdatePayload(ctx *v1.UpdateProposalContext) Proposal {
	payload := ctx.Payload
	var err error
	target, _ := MarshalUpdateProposalPayload(payload, err)

	return target
}

// MarshalUpdateProposalPayload validates and renders an instance of UpdateProposalPayload into a interface{}
func MarshalUpdateProposalPayload(source *v1.UpdateProposalPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp10 := map[string]interface{}{
		"abstract":   source.Abstract,
		"created_at": source.CreatedAt,
		"deleted_at": source.DeletedAt,
		"detail":     source.Detail,
		"first_name": source.FirstName,
		"id":         source.ID,
		"m2reviews":  source.M2reviews,
		"reviews":    source.Reviews,
		"title":      source.Title,
		"updated_at": source.UpdatedAt,
		"user_id":    source.UserId,
		"withdrawn":  source.Withdrawn,
	}
	target = tmp10
	return
}
