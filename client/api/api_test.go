package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var testFileName = "../../ports_test.json"
var testFileName2Records = "../../ports_test_2_records.json"
var e = GetEnvTest()
var rt = ServerWithMiddlewares()
var assertCorrectMessage = func(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %s[%T] want %s[%T]", got, got, want, want)
	}
}
var assertNil = func(t *testing.T, got interface{}) {
	t.Helper()
	if got != nil {
		t.Errorf("got %q want nil", got)
	}
}
var assertTrue = func(t *testing.T, got interface{}) {
	t.Helper()
	if got != true {
		t.Errorf("got %q want true", got)
	}
}

type ReqTest struct {
	Verb string
	URL  string
	Body string
}

func GetEnvTest() *Environment {
	testing.Init()
	env := &Environment{}
	return env
}

func ServerWithMiddlewares() http.Handler {
	router := mux.NewRouter()
	e.GetRoutes(router)
	e.GetMiddlewares(router)

	return router
}

func MakeRequest(r ReqTest) (res *httptest.ResponseRecorder) {
	b := strings.NewReader(r.Body)
	req, _ := http.NewRequest(r.Verb, r.URL, b)
	if r.Verb == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	res = httptest.NewRecorder()
	rt.ServeHTTP(res, req)

	return res
}
