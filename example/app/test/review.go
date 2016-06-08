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

// CreateReviewCreated test setup
func CreateReviewCreated(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, payload *app.CreateReviewPayload) {
	CreateReviewCreatedCtx(t, context.Background(), ctrl, userID, proposalID, payload)
}

// CreateReviewCreatedCtx test setup
func CreateReviewCreatedCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, payload *app.CreateReviewPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 201 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	createCtx, err := app.NewCreateReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createCtx.Payload = payload

	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

}

// DeleteReviewNoContent test setup
func DeleteReviewNoContent(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int) {
	DeleteReviewNoContentCtx(t, context.Background(), ctrl, userID, proposalID, reviewID)
}

// DeleteReviewNoContentCtx test setup
func DeleteReviewNoContentCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteReviewContext(goaCtx, service)
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

// DeleteReviewNotFound test setup
func DeleteReviewNotFound(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int) {
	DeleteReviewNotFoundCtx(t, context.Background(), ctrl, userID, proposalID, reviewID)
}

// DeleteReviewNotFoundCtx test setup
func DeleteReviewNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteReviewContext(goaCtx, service)
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

// ListReviewOK test setup
func ListReviewOK(t *testing.T, ctrl app.ReviewController, userID string, proposalID string) *app.ReviewCollection {
	return ListReviewOKCtx(t, context.Background(), ctrl, userID, proposalID)
}

// ListReviewOKCtx test setup
func ListReviewOKCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string) *app.ReviewCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	listCtx, err := app.NewListReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.ReviewCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.ReviewCollection", resp)
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

// ListReviewOKLink test setup
func ListReviewOKLink(t *testing.T, ctrl app.ReviewController, userID string, proposalID string) *app.ReviewLinkCollection {
	return ListReviewOKLinkCtx(t, context.Background(), ctrl, userID, proposalID)
}

// ListReviewOKLinkCtx test setup
func ListReviewOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string) *app.ReviewLinkCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	listCtx, err := app.NewListReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.ReviewLinkCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.ReviewLinkCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// ShowReviewNotFound test setup
func ShowReviewNotFound(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int) {
	ShowReviewNotFoundCtx(t, context.Background(), ctrl, userID, proposalID, reviewID)
}

// ShowReviewNotFoundCtx test setup
func ShowReviewNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	showCtx, err := app.NewShowReviewContext(goaCtx, service)
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

// ShowReviewOK test setup
func ShowReviewOK(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int) *app.Review {
	return ShowReviewOKCtx(t, context.Background(), ctrl, userID, proposalID, reviewID)
}

// ShowReviewOKCtx test setup
func ShowReviewOKCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int) *app.Review {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	showCtx, err := app.NewShowReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Review)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Review", resp)
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

// ShowReviewOKLink test setup
func ShowReviewOKLink(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int) *app.ReviewLink {
	return ShowReviewOKLinkCtx(t, context.Background(), ctrl, userID, proposalID, reviewID)
}

// ShowReviewOKLinkCtx test setup
func ShowReviewOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int) *app.ReviewLink {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	showCtx, err := app.NewShowReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.ReviewLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.ReviewLink", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// UpdateReviewNoContent test setup
func UpdateReviewNoContent(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int, payload *app.UpdateReviewPayload) {
	UpdateReviewNoContentCtx(t, context.Background(), ctrl, userID, proposalID, reviewID, payload)
}

// UpdateReviewNoContentCtx test setup
func UpdateReviewNoContentCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int, payload *app.UpdateReviewPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 204 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// UpdateReviewNotFound test setup
func UpdateReviewNotFound(t *testing.T, ctrl app.ReviewController, userID string, proposalID string, reviewID int, payload *app.UpdateReviewPayload) {
	UpdateReviewNotFoundCtx(t, context.Background(), ctrl, userID, proposalID, reviewID, payload)
}

// UpdateReviewNotFoundCtx test setup
func UpdateReviewNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ReviewController, userID string, proposalID string, reviewID int, payload *app.UpdateReviewPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["userID"] = []string{fmt.Sprintf("%v", userID)}
	prms["proposalID"] = []string{fmt.Sprintf("%v", proposalID)}
	prms["reviewID"] = []string{fmt.Sprintf("%v", reviewID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReviewTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateReviewContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
