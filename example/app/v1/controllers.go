//************************************************************************//
// API "congo" version v1: Application Controllers
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

import "github.com/goadesign/goa"

// ProposalController is the controller interface for the Proposal actions.
type ProposalController interface {
	goa.Controller
	Create(*CreateProposalContext) error
	Delete(*DeleteProposalContext) error
	List(*ListProposalContext) error
	Show(*ShowProposalContext) error
	Update(*UpdateProposalContext) error
}

// MountProposalController "mounts" a Proposal resource controller on the given service.
func MountProposalController(service goa.Service, ctrl ProposalController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.Version("v1").SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.Version("v1").SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.Version("v1").SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.Version("v1").SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.Version("v1").SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.Version("v1").SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.Version("v1").ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCreateProposalContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		ctx.Payload = ctx.RawPayload().(*CreateProposalPayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/:api_version/api/users/:userID/proposals", ctrl.HandleFunc("Create", h, unmarshalCreateProposalPayload))
	service.Info("mount", "ctrl", "Proposal", "version", "v1", "action", "Create", "route", "POST /:api_version/api/users/:userID/proposals")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteProposalContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/:api_version/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Delete", h, nil))
	service.Info("mount", "ctrl", "Proposal", "version", "v1", "action", "Delete", "route", "DELETE /:api_version/api/users/:userID/proposals/:proposalID")
	h = func(c *goa.Context) error {
		ctx, err := NewListProposalContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/:api_version/api/users/:userID/proposals", ctrl.HandleFunc("List", h, nil))
	service.Info("mount", "ctrl", "Proposal", "version", "v1", "action", "List", "route", "GET /:api_version/api/users/:userID/proposals")
	h = func(c *goa.Context) error {
		ctx, err := NewShowProposalContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/:api_version/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Show", h, nil))
	service.Info("mount", "ctrl", "Proposal", "version", "v1", "action", "Show", "route", "GET /:api_version/api/users/:userID/proposals/:proposalID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateProposalContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		ctx.Payload = ctx.RawPayload().(*UpdateProposalPayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/:api_version/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Update", h, unmarshalUpdateProposalPayload))
	service.Info("mount", "ctrl", "Proposal", "version", "v1", "action", "Update", "route", "PATCH /:api_version/api/users/:userID/proposals/:proposalID")
}

// unmarshalCreateProposalPayload unmarshals the request body.
func unmarshalCreateProposalPayload(ctx *goa.Context) error {
	payload := &CreateProposalPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// unmarshalUpdateProposalPayload unmarshals the request body.
func unmarshalUpdateProposalPayload(ctx *goa.Context) error {
	payload := &UpdateProposalPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// ReviewController is the controller interface for the Review actions.
type ReviewController interface {
	goa.Controller
	Create(*CreateReviewContext) error
	Delete(*DeleteReviewContext) error
	List(*ListReviewContext) error
	Show(*ShowReviewContext) error
	Update(*UpdateReviewContext) error
}

// MountReviewController "mounts" a Review resource controller on the given service.
func MountReviewController(service goa.Service, ctrl ReviewController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.Version("v1").SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.Version("v1").SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.Version("v1").SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.Version("v1").SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.Version("v1").SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.Version("v1").SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.Version("v1").ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCreateReviewContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		ctx.Payload = ctx.RawPayload().(*CreateReviewPayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/:api_version/api/users/:userID/proposals/:proposalID/review", ctrl.HandleFunc("Create", h, unmarshalCreateReviewPayload))
	service.Info("mount", "ctrl", "Review", "version", "v1", "action", "Create", "route", "POST /:api_version/api/users/:userID/proposals/:proposalID/review")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteReviewContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/:api_version/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Delete", h, nil))
	service.Info("mount", "ctrl", "Review", "version", "v1", "action", "Delete", "route", "DELETE /:api_version/api/users/:userID/proposals/:proposalID/review/:reviewID")
	h = func(c *goa.Context) error {
		ctx, err := NewListReviewContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/:api_version/api/users/:userID/proposals/:proposalID/review", ctrl.HandleFunc("List", h, nil))
	service.Info("mount", "ctrl", "Review", "version", "v1", "action", "List", "route", "GET /:api_version/api/users/:userID/proposals/:proposalID/review")
	h = func(c *goa.Context) error {
		ctx, err := NewShowReviewContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/:api_version/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Show", h, nil))
	service.Info("mount", "ctrl", "Review", "version", "v1", "action", "Show", "route", "GET /:api_version/api/users/:userID/proposals/:proposalID/review/:reviewID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateReviewContext(c)
		ctx.APIVersion = service.Version("v1").VersionName()
		ctx.Payload = ctx.RawPayload().(*UpdateReviewPayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/:api_version/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Update", h, unmarshalUpdateReviewPayload))
	service.Info("mount", "ctrl", "Review", "version", "v1", "action", "Update", "route", "PATCH /:api_version/api/users/:userID/proposals/:proposalID/review/:reviewID")
}

// unmarshalCreateReviewPayload unmarshals the request body.
func unmarshalCreateReviewPayload(ctx *goa.Context) error {
	payload := &CreateReviewPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// unmarshalUpdateReviewPayload unmarshals the request body.
func unmarshalUpdateReviewPayload(ctx *goa.Context) error {
	payload := &UpdateReviewPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}
