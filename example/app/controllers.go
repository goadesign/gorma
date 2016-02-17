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

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Controller
	Callback(*CallbackAuthContext) error
	Oauth(*OauthAuthContext) error
	Refresh(*RefreshAuthContext) error
	Token(*TokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service goa.Service, ctrl AuthController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCallbackAuthContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Callback(rctx)
	}
	mux.Handle("GET", "/api/auth/:provider/callback", ctrl.HandleFunc("Callback", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Auth"}, goa.KV{"action", "Callback"}, goa.KV{"route", "GET /api/auth/:provider/callback"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewOauthAuthContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Oauth(rctx)
	}
	mux.Handle("GET", "/api/auth/:provider", ctrl.HandleFunc("Oauth", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Auth"}, goa.KV{"action", "Oauth"}, goa.KV{"route", "GET /api/auth/:provider"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRefreshAuthContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RefreshAuthPayload)
		}
		return ctrl.Refresh(rctx)
	}
	mux.Handle("POST", "/api/auth/refresh", ctrl.HandleFunc("Refresh", h, unmarshalRefreshAuthPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Auth"}, goa.KV{"action", "Refresh"}, goa.KV{"route", "POST /api/auth/refresh"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewTokenAuthContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TokenAuthPayload)
		}
		return ctrl.Token(rctx)
	}
	mux.Handle("POST", "/api/auth/token", ctrl.HandleFunc("Token", h, unmarshalTokenAuthPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Auth"}, goa.KV{"action", "Token"}, goa.KV{"route", "POST /api/auth/token"})
}

// unmarshalRefreshAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalRefreshAuthPayload(ctx context.Context, req *http.Request) error {
	payload := &RefreshAuthPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}

// unmarshalTokenAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalTokenAuthPayload(ctx context.Context, req *http.Request) error {
	payload := &TokenAuthPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}

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
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateProposalContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateProposalPayload)
		}
		return ctrl.Create(rctx)
	}
	mux.Handle("POST", "/api/users/:userID/proposals", ctrl.HandleFunc("Create", h, unmarshalCreateProposalPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Proposal"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /api/users/:userID/proposals"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteProposalContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(rctx)
	}
	mux.Handle("DELETE", "/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Delete", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Proposal"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /api/users/:userID/proposals/:proposalID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListProposalContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(rctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals", ctrl.HandleFunc("List", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Proposal"}, goa.KV{"action", "List"}, goa.KV{"route", "GET /api/users/:userID/proposals"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowProposalContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Proposal"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /api/users/:userID/proposals/:proposalID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateProposalContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateProposalPayload)
		}
		return ctrl.Update(rctx)
	}
	mux.Handle("PATCH", "/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Update", h, unmarshalUpdateProposalPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Proposal"}, goa.KV{"action", "Update"}, goa.KV{"route", "PATCH /api/users/:userID/proposals/:proposalID"})
}

// unmarshalCreateProposalPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateProposalPayload(ctx context.Context, req *http.Request) error {
	payload := &CreateProposalPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}

// unmarshalUpdateProposalPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateProposalPayload(ctx context.Context, req *http.Request) error {
	payload := &UpdateProposalPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
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
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateReviewContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateReviewPayload)
		}
		return ctrl.Create(rctx)
	}
	mux.Handle("POST", "/api/users/:userID/proposals/:proposalID/review", ctrl.HandleFunc("Create", h, unmarshalCreateReviewPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Review"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /api/users/:userID/proposals/:proposalID/review"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteReviewContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(rctx)
	}
	mux.Handle("DELETE", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Delete", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Review"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /api/users/:userID/proposals/:proposalID/review/:reviewID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListReviewContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(rctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals/:proposalID/review", ctrl.HandleFunc("List", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Review"}, goa.KV{"action", "List"}, goa.KV{"route", "GET /api/users/:userID/proposals/:proposalID/review"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowReviewContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Review"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /api/users/:userID/proposals/:proposalID/review/:reviewID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateReviewContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateReviewPayload)
		}
		return ctrl.Update(rctx)
	}
	mux.Handle("PATCH", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Update", h, unmarshalUpdateReviewPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Review"}, goa.KV{"action", "Update"}, goa.KV{"route", "PATCH /api/users/:userID/proposals/:proposalID/review/:reviewID"})
}

// unmarshalCreateReviewPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateReviewPayload(ctx context.Context, req *http.Request) error {
	payload := &CreateReviewPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}

// unmarshalUpdateReviewPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateReviewPayload(ctx context.Context, req *http.Request) error {
	payload := &UpdateReviewPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}

// UiController is the controller interface for the Ui actions.
type UiController interface {
	goa.Controller
	Bootstrap(*BootstrapUiContext) error
}

// MountUiController "mounts" a Ui resource controller on the given service.
func MountUiController(service goa.Service, ctrl UiController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewBootstrapUiContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Bootstrap(rctx)
	}
	mux.Handle("GET", "/", ctrl.HandleFunc("Bootstrap", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Ui"}, goa.KV{"action", "Bootstrap"}, goa.KV{"route", "GET /"})
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Controller
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Show(*ShowUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service goa.Service, ctrl UserController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCreateUserContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserPayload)
		}
		return ctrl.Create(rctx)
	}
	mux.Handle("POST", "/api/users", ctrl.HandleFunc("Create", h, unmarshalCreateUserPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "User"}, goa.KV{"action", "Create"}, goa.KV{"route", "POST /api/users"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewDeleteUserContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(rctx)
	}
	mux.Handle("DELETE", "/api/users/:userID", ctrl.HandleFunc("Delete", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "User"}, goa.KV{"action", "Delete"}, goa.KV{"route", "DELETE /api/users/:userID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewListUserContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(rctx)
	}
	mux.Handle("GET", "/api/users", ctrl.HandleFunc("List", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "User"}, goa.KV{"action", "List"}, goa.KV{"route", "GET /api/users"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowUserContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	mux.Handle("GET", "/api/users/:userID", ctrl.HandleFunc("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "User"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /api/users/:userID"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewUpdateUserContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserPayload)
		}
		return ctrl.Update(rctx)
	}
	mux.Handle("PATCH", "/api/users/:userID", ctrl.HandleFunc("Update", h, unmarshalUpdateUserPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "User"}, goa.KV{"action", "Update"}, goa.KV{"route", "PATCH /api/users/:userID"})
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, req *http.Request) error {
	payload := &CreateUserPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, req *http.Request) error {
	payload := &UpdateUserPayload{}
	if err := goa.RequestService(ctx).DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = payload
	return nil
}
