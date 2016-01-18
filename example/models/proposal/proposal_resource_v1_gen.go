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

import "github.com/gopheracademy/congo/app/v1"

func ProposalFromV1CreatePayload(ctx *v1.CreateProposalContext) (*Proposal, error) {
	payload := ctx.Payload
	var err error
	middle, err := MarshalCreateProposalPayload(payload, err)
	target, err := UnmarshalProposal(middle)
	return target, err
}

func ProposalFromV1UpdatePayload(ctx *v1.UpdateProposalContext) (*Proposal, error) {
	payload := ctx.Payload
	var err error
	middle, err := MarshalUpdateProposalPayload(payload, err)
	target, err := UnmarshalProposal(middle)
	return target, err
}
