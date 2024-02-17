package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/tests"
)

func TestHandlerUsersCreate(t *testing.T) {
	ta := test.NewTest()
	defer ta.Close()

	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	param := parameters{
		Username: "test",
		Password: "password",
	}

	r := httptest.NewRequest(http.MethodPost, "/users/", test.StructToJSONBody(param))
	w := httptest.NewRecorder()
	ta.Handler.HandlerUsersCreate(w, r)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.Status != "201 Created" {
		t.Errorf("Expected status 201, got %s", resp.Status)
	}
}
