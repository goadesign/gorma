package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/goadesign/gorma/example/app"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// CallbackAuthOK test setup
func CallbackAuthOK(t *testing.T, ctrl app.AuthController, provider string) {
	CallbackAuthOKCtx(t, context.Background(), ctrl, provider)
}

// CallbackAuthOKCtx test setup
func CallbackAuthOKCtx(t *testing.T, ctx context.Context, ctrl app.AuthController, provider string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/auth/%v/callback", provider), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["provider"] = []string{fmt.Sprintf("%v", provider)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, prms)
	callbackCtx, err := app.NewCallbackAuthContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Callback(callbackCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// OauthAuthOK test setup
func OauthAuthOK(t *testing.T, ctrl app.AuthController, provider string) *app.Authorize {
	return OauthAuthOKCtx(t, context.Background(), ctrl, provider)
}

// OauthAuthOKCtx test setup
func OauthAuthOKCtx(t *testing.T, ctx context.Context, ctrl app.AuthController, provider string) *app.Authorize {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/auth/%v", provider), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["provider"] = []string{fmt.Sprintf("%v", provider)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, prms)
	oauthCtx, err := app.NewOauthAuthContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Oauth(oauthCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Authorize)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Authorize", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// RefreshAuthCreated test setup
func RefreshAuthCreated(t *testing.T, ctrl app.AuthController, payload *app.RefreshAuthPayload) *app.Authorize {
	return RefreshAuthCreatedCtx(t, context.Background(), ctrl, payload)
}

// RefreshAuthCreatedCtx test setup
func RefreshAuthCreatedCtx(t *testing.T, ctx context.Context, ctrl app.AuthController, payload *app.RefreshAuthPayload) *app.Authorize {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/auth/refresh"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, prms)
	refreshCtx, err := app.NewRefreshAuthContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	refreshCtx.Payload = payload

	err = ctrl.Refresh(refreshCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Authorize)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Authorize", resp)
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	return a

}

// TokenAuthCreated test setup
func TokenAuthCreated(t *testing.T, ctrl app.AuthController, payload *app.TokenAuthPayload) *app.Authorize {
	return TokenAuthCreatedCtx(t, context.Background(), ctrl, payload)
}

// TokenAuthCreatedCtx test setup
func TokenAuthCreatedCtx(t *testing.T, ctx context.Context, ctrl app.AuthController, payload *app.TokenAuthPayload) *app.Authorize {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/auth/token"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, prms)
	tokenCtx, err := app.NewTokenAuthContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	tokenCtx.Payload = payload

	err = ctrl.Token(tokenCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Authorize)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Authorize", resp)
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	return a

}
