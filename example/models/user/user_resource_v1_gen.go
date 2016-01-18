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

package user

import "github.com/gopheracademy/congo/app/v1"

func UserFromV1CreatePayload(ctx *v1.CreateUserContext) (*User, error) {
	payload := ctx.Payload
	var err error
	middle, err := MarshalCreateUserPayload(payload, err)
	target, err := UnmarshalUser(middle)
	return target, err
}

func UserFromV1UpdatePayload(ctx *v1.UpdateUserContext) (*User, error) {
	payload := ctx.Payload
	var err error
	middle, err := MarshalUpdateUserPayload(payload, err)
	target, err := UnmarshalUser(middle)
	return target, err
}
