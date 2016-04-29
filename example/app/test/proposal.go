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
	"testing"
)

// CreateProposalCreated test setup
func CreateProposalCreated(t *testing.T, ctrl app.ProposalController, userID string, payload *app.CreateProposalPayload) {
	CreateProposalCreatedCtx(t, context.Background(), ctrl, userID, payload)
}

// CreateProposalCreatedCtx test setup
func CreateProposalCreatedCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, payload *app.CreateProposalPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/users/%v/proposals", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	createCtx, err := app.NewCreateProposalContext(goaCtx, service)
	createCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

}

// UpdateProposalNoContent test setup
func UpdateProposalNoContent(t *testing.T, ctrl app.ProposalController, userID string, proposalID int, payload *app.UpdateProposalPayload) {
	UpdateProposalNoContentCtx(t, context.Background(), ctrl, userID, proposalID, payload)
}

// UpdateProposalNoContentCtx test setup
func UpdateProposalNoContentCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int, payload *app.UpdateProposalPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateProposalContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// UpdateProposalNotFound test setup
func UpdateProposalNotFound(t *testing.T, ctrl app.ProposalController, userID string, proposalID int, payload *app.UpdateProposalPayload) {
	UpdateProposalNotFoundCtx(t, context.Background(), ctrl, userID, proposalID, payload)
}

// UpdateProposalNotFoundCtx test setup
func UpdateProposalNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int, payload *app.UpdateProposalPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateProposalContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// DeleteProposalNoContent test setup
func DeleteProposalNoContent(t *testing.T, ctrl app.ProposalController, userID string, proposalID int) {
	DeleteProposalNoContentCtx(t, context.Background(), ctrl, userID, proposalID)
}

// DeleteProposalNoContentCtx test setup
func DeleteProposalNoContentCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteProposalContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// DeleteProposalNotFound test setup
func DeleteProposalNotFound(t *testing.T, ctrl app.ProposalController, userID string, proposalID int) {
	DeleteProposalNotFoundCtx(t, context.Background(), ctrl, userID, proposalID)
}

// DeleteProposalNotFoundCtx test setup
func DeleteProposalNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteProposalContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ListProposalOK test setup
func ListProposalOK(t *testing.T, ctrl app.ProposalController, userID string) *app.ProposalCollection {
	return ListProposalOKCtx(t, context.Background(), ctrl, userID)
}

// ListProposalOKCtx test setup
func ListProposalOKCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string) *app.ProposalCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	listCtx, err := app.NewListProposalContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.ProposalCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.ProposalCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowProposalOK test setup
func ShowProposalOK(t *testing.T, ctrl app.ProposalController, userID string, proposalID int) *app.Proposal {
	return ShowProposalOKCtx(t, context.Background(), ctrl, userID, proposalID)
}

// ShowProposalOKCtx test setup
func ShowProposalOKCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int) *app.Proposal {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	showCtx, err := app.NewShowProposalContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Proposal)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Proposal", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowProposalOKLink test setup
func ShowProposalOKLink(t *testing.T, ctrl app.ProposalController, userID string, proposalID int) *app.ProposalLink {
	return ShowProposalOKLinkCtx(t, context.Background(), ctrl, userID, proposalID)
}

// ShowProposalOKLinkCtx test setup
func ShowProposalOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int) *app.ProposalLink {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	showCtx, err := app.NewShowProposalContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.ProposalLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.ProposalLink", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowProposalNotFound test setup
func ShowProposalNotFound(t *testing.T, ctrl app.ProposalController, userID string, proposalID int) {
	ShowProposalNotFoundCtx(t, context.Background(), ctrl, userID, proposalID)
}

// ShowProposalNotFoundCtx test setup
func ShowProposalNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ProposalController, userID string, proposalID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ProposalTest"), rw, req, nil)
	showCtx, err := app.NewShowProposalContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
