package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	routerutils "github.com/GolangUnited/helloweb/router-utils"
	"github.com/gorilla/mux"
)

var router *mux.Router

func TestMain(m *testing.M) {
	setup()
	code := m.Run()

	//teardown()

	os.Exit(code)
}

func setup() {
	router = routerutils.Create()
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr
}
