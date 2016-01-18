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

import "github.com/gopheracademy/congo/app/v1"

func ReviewFromV1CreatePayload(ctx *v1.CreateReviewContext) (*Review, error) {
	payload := ctx.Payload
	var err error
	middle, err := MarshalCreateReviewPayload(payload, err)
	target, err := UnmarshalReview(middle)
	return target, err
}

func ReviewFromV1UpdatePayload(ctx *v1.UpdateReviewContext) (*Review, error) {
	payload := ctx.Payload
	var err error
	middle, err := MarshalUpdateReviewPayload(payload, err)
	target, err := UnmarshalReview(middle)
	return target, err
}
