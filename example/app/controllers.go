//************************************************************************//
// API "congo": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// inited is true if initService has been called
var inited = false

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	if inited {
		return
	}
	inited = true

	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Callback(*CallbackAuthContext) error
	Oauth(*OauthAuthContext) error
	Refresh(*RefreshAuthContext) error
	Token(*TokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCallbackAuthContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Callback(rctx)
	}
	service.Mux.Handle("GET", "/api/auth/:provider/callback", ctrl.MuxHandler("Callback", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Callback", "route", "GET /api/auth/:provider/callback")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewOauthAuthContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Oauth(rctx)
	}
	service.Mux.Handle("GET", "/api/auth/:provider", ctrl.MuxHandler("Oauth", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Oauth", "route", "GET /api/auth/:provider")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRefreshAuthContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RefreshAuthPayload)
		}
		return ctrl.Refresh(rctx)
	}
	service.Mux.Handle("POST", "/api/auth/refresh", ctrl.MuxHandler("Refresh", h, unmarshalRefreshAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Refresh", "route", "POST /api/auth/refresh")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewTokenAuthContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TokenAuthPayload)
		}
		return ctrl.Token(rctx)
	}
	service.Mux.Handle("POST", "/api/auth/token", ctrl.MuxHandler("Token", h, unmarshalTokenAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Token", "route", "POST /api/auth/token")
}

// unmarshalRefreshAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalRefreshAuthPayload(ctx context.Context, req *http.Request) error {
	var payload RefreshAuthPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalTokenAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalTokenAuthPayload(ctx context.Context, req *http.Request) error {
	var payload TokenAuthPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// ProposalController is the controller interface for the Proposal actions.
type ProposalController interface {
	goa.Muxer
	Create(*CreateProposalContext) error
	Delete(*DeleteProposalContext) error
	List(*ListProposalContext) error
	Show(*ShowProposalContext) error
	Update(*UpdateProposalContext) error
}

// MountProposalController "mounts" a Proposal resource controller on the given service.
func MountProposalController(service *goa.Service, ctrl ProposalController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateProposalContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateProposalPayload)
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/api/users/:userID/proposals", ctrl.MuxHandler("Create", h, unmarshalCreateProposalPayload))
	service.LogInfo("mount", "ctrl", "Proposal", "action", "Create", "route", "POST /api/users/:userID/proposals")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteProposalContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/api/users/:userID/proposals/:proposalID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Proposal", "action", "Delete", "route", "DELETE /api/users/:userID/proposals/:proposalID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListProposalContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/users/:userID/proposals", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Proposal", "action", "List", "route", "GET /api/users/:userID/proposals")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowProposalContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/users/:userID/proposals/:proposalID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Proposal", "action", "Show", "route", "GET /api/users/:userID/proposals/:proposalID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateProposalContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateProposalPayload)
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/api/users/:userID/proposals/:proposalID", ctrl.MuxHandler("Update", h, unmarshalUpdateProposalPayload))
	service.LogInfo("mount", "ctrl", "Proposal", "action", "Update", "route", "PATCH /api/users/:userID/proposals/:proposalID")
}

// unmarshalCreateProposalPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateProposalPayload(ctx context.Context, req *http.Request) error {
	var payload CreateProposalPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalUpdateProposalPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateProposalPayload(ctx context.Context, req *http.Request) error {
	var payload UpdateProposalPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// ReviewController is the controller interface for the Review actions.
type ReviewController interface {
	goa.Muxer
	Create(*CreateReviewContext) error
	Delete(*DeleteReviewContext) error
	List(*ListReviewContext) error
	Show(*ShowReviewContext) error
	Update(*UpdateReviewContext) error
}

// MountReviewController "mounts" a Review resource controller on the given service.
func MountReviewController(service *goa.Service, ctrl ReviewController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateReviewContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateReviewPayload)
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/api/users/:userID/proposals/:proposalID/review", ctrl.MuxHandler("Create", h, unmarshalCreateReviewPayload))
	service.LogInfo("mount", "ctrl", "Review", "action", "Create", "route", "POST /api/users/:userID/proposals/:proposalID/review")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteReviewContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Review", "action", "Delete", "route", "DELETE /api/users/:userID/proposals/:proposalID/review/:reviewID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListReviewContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/users/:userID/proposals/:proposalID/review", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Review", "action", "List", "route", "GET /api/users/:userID/proposals/:proposalID/review")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowReviewContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Review", "action", "Show", "route", "GET /api/users/:userID/proposals/:proposalID/review/:reviewID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateReviewContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateReviewPayload)
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.MuxHandler("Update", h, unmarshalUpdateReviewPayload))
	service.LogInfo("mount", "ctrl", "Review", "action", "Update", "route", "PATCH /api/users/:userID/proposals/:proposalID/review/:reviewID")
}

// unmarshalCreateReviewPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateReviewPayload(ctx context.Context, req *http.Request) error {
	var payload CreateReviewPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalUpdateReviewPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateReviewPayload(ctx context.Context, req *http.Request) error {
	var payload UpdateReviewPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// UIController is the controller interface for the UI actions.
type UIController interface {
	goa.Muxer
	Bootstrap(*BootstrapUIContext) error
}

// MountUIController "mounts" a UI resource controller on the given service.
func MountUIController(service *goa.Service, ctrl UIController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewBootstrapUIContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Bootstrap(rctx)
	}
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("Bootstrap", h, nil))
	service.LogInfo("mount", "ctrl", "UI", "action", "Bootstrap", "route", "GET /")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Show(*ShowUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateUserContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserPayload)
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/api/users", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /api/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteUserContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/api/users/:userID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /api/users/:userID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListUserContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/users", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /api/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowUserContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/users/:userID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Show", "route", "GET /api/users/:userID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateUserContext(ctx)
		if err != nil {
			return err
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserPayload)
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/api/users/:userID", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PATCH /api/users/:userID")
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, req *http.Request) error {
	var payload CreateUserPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, req *http.Request) error {
	var payload UpdateUserPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}
