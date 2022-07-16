package tests

import (
	"net/http"
	"testing"
)

func TestGetNameWithNameMustReturnOk(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/name/Dara", nil)

	want := http.StatusOK
	got := executeRequest(req).Code

	if got != want {
		t.Errorf("TestGetNameWithNameMustReturnOk got: %v, want: %v", got, want)
	}
}

func TestGetNameWithNameMustReturnMessage(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/name/Dara", nil)

	want := "Hello, Dara!"
	got := executeRequest(req).Body.String()

	if got != want {
		t.Errorf("TestGetNameWithNameMustReturnMessage got: %v, want: %v", got, want)
	}
}

func TestBadGetMustReturnStatus500(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/bad", nil)

	want := http.StatusInternalServerError
	got := executeRequest(req).Code

	if got != want {
		t.Errorf("TestBadGetMustReturnStatus500 got: %v, want: %v", got, want)
	}
}
