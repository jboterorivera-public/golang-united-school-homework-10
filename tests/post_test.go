package tests

import (
	"net/http"
	"strings"
	"testing"
)

func TestPostDataWithoutBodyMustReturn400(t *testing.T) {
	body := ""
	req, _ := http.NewRequest(http.MethodPost, "/data", strings.NewReader(body))

	want := http.StatusBadRequest
	got := executeRequest(req).Code

	if got != want {
		t.Errorf("TestBadGetMustReturnStatus500 got: %v, want: %v", got, want)
	}
}

func TestPostDataWithoutBodyMustReturnAlertMessage(t *testing.T) {
	body := ""
	req, _ := http.NewRequest(http.MethodPost, "/data", strings.NewReader(body))

	want := "Empty body"
	got := executeRequest(req).Body.String()

	if got != want {
		t.Errorf("TestPostDataWithoutBodyMustReturnAlertMessage got: %v, want: %v", got, want)
	}
}

func TestPostDataWithBodyMustReturnMessage(t *testing.T) {
	body := "Dara"
	req, _ := http.NewRequest(http.MethodPost, "/data", strings.NewReader(body))

	want := "I got message:\nDara"
	got := executeRequest(req).Body.String()

	if got != want {
		t.Errorf("TestPostDataWithBodyMustReturnMessage got: %v, want: %v", got, want)
	}
}

func TestPostDataWithBodyMustReturnOk(t *testing.T) {
	body := "Dara"
	req, _ := http.NewRequest(http.MethodPost, "/data", strings.NewReader(body))

	want := http.StatusOK
	got := executeRequest(req).Code

	if got != want {
		t.Errorf("TestPostDataWithBodyMustReturnOk got: %v, want: %v", got, want)
	}
}

func TestPostHeaderWithoutHeader(t *testing.T) {
	body := ""
	req, _ := http.NewRequest(http.MethodPost, "/headers", strings.NewReader(body))

	want := http.StatusBadRequest
	got := executeRequest(req).Code

	if got != want {
		t.Errorf("TestPostHeaderWithoutHeader got: %v, want: %v", got, want)
	}
}

func TestPostHeaderWithBadHeaderA(t *testing.T) {
	body := ""
	req, _ := http.NewRequest(http.MethodPost, "/headers", strings.NewReader(body))
	req.Header.Add("a", "z")
	req.Header.Add("b", "1")

	want := "Invalid val for header a"
	got := executeRequest(req).Body.String()

	if got != want {
		t.Errorf("TestPostHeaderWithBadHeaderA got: %v, want: %v", got, want)
	}
}

func TestPostHeaderWithBadHeaderB(t *testing.T) {
	body := ""
	req, _ := http.NewRequest(http.MethodPost, "/headers", strings.NewReader(body))
	req.Header.Add("a", "1")
	req.Header.Add("b", "z")

	want := "Invalid val for header b"
	got := executeRequest(req).Body.String()

	if got != want {
		t.Errorf("TestPostHeaderWithBadHeaderB got: %v, want: %v", got, want)
	}
}

func TestPostHeaderSumOk(t *testing.T) {
	body := ""
	req, _ := http.NewRequest(http.MethodPost, "/headers", strings.NewReader(body))
	req.Header.Add("a", "10")
	req.Header.Add("b", "2")

	want := "12"
	got := executeRequest(req).Header().Get("a+b")

	if got != want {
		t.Errorf("TestPostHeaderSumOk got: %v, want: %v", got, want)
	}
}
